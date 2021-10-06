/*
Copyright © 2021 blacktop

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/apex/log"
	clihander "github.com/apex/log/handlers/cli"
	"github.com/blacktop/go-macho"
	"github.com/blacktop/go-macho/types"

	"github.com/blacktop/arm64-cgo/disassemble"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

type opName uint32

const (
	AMXLDX opName = iota
	AMXLDY
	AMXSTX
	AMXSTY
	AMXLDZ
	AMXSTZ
	AMXLDZI
	AMXSTZI
	AMXEXTRX // amxextrx?
	AMXEXTRY // amxextry?
	AMXFMA64
	AMXFMS64
	AMXFMA32
	AMXFMS32
	AMXMAC16
	AMXFMA16
	AMXFMS16
	AMX17 // amxset / amxclr
	AMXVECINT
	AMXVECFP
	AMXMATINT
	AMXMATFP
	AMXGENLUT
)

func (o opName) String() string {
	switch o {
	case AMXLDX:
		return "amx_ldx"
	case AMXLDY:
		return "amx_ldy"
	case AMXSTX:
		return "amx_stx"
	case AMXSTY:
		return "amx_sty"
	case AMXLDZ:
		return "amx_ldz"
	case AMXSTZ:
		return "amx_stz"
	case AMXLDZI:
		return "amx_ldzi"
	case AMXSTZI:
		return "amx_stzi"
	case AMXEXTRX:
		return "amx_extrx"
	case AMXEXTRY:
		return "amx_extry"
	case AMXFMA64:
		return "amx_fma64"
	case AMXFMS64:
		return "amx_fms64"
	case AMXFMA32:
		return "amx_fma32"
	case AMXFMS32:
		return "amx_fms32"
	case AMXMAC16:
		return "amx_mac16"
	case AMXFMA16:
		return "amx_fma16"
	case AMXFMS16:
		return "amx_fms16"
	case AMX17:
		return "amx_op17"
	case AMXVECINT:
		return "amx_vecint"
	case AMXVECFP:
		return "amx_vecfp"
	case AMXMATINT:
		return "amx_matint"
	case AMXMATFP:
		return "amx_matfp"
	case AMXGENLUT:
		return "amx_genlut"
	default:
		return "unk"
	}
}

func init() {
	log.SetHandler(clihander.Default)
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.disass.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("json", "j", false, "Output as JSON")

	rootCmd.Flags().BoolP("all", "", false, "Disassemble all functions")
	rootCmd.Flags().StringP("symbol", "s", "", "Function to disassemble")
	rootCmd.Flags().Uint64P("vaddr", "a", 0, "Virtual address to disassemble")
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "disass",
	Short: "MachO AARCH64 disassembler",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		startVMAddr, _ := cmd.Flags().GetUint64("vaddr")
		symbolName, _ := cmd.Flags().GetString("symbol")
		asJSON, _ := cmd.Flags().GetBool("json")
		allFuncs, _ := cmd.Flags().GetBool("all")

		var isMiddle bool
		var symAddr uint64
		var data []byte
		var funcs []types.Function

		addr2SymMap := make(map[uint64]string)

		var m *macho.File
		var instructions []disassemble.Instruction

		machoPath := filepath.Clean(args[0])

		fileInfo, err := os.Lstat(machoPath)
		if err != nil {
			log.Fatal(fmt.Sprintf("file %s does not exist: %v", machoPath, err))
		}

		// Check if file is a symlink
		if fileInfo.Mode()&os.ModeSymlink != 0 {
			symlinkPath, err := os.Readlink(machoPath)
			if err != nil {
				log.Fatal(fmt.Sprintf("failed to read symlink %s: %v", machoPath, err))
			}
			// TODO: this seems like it would break
			linkParent := filepath.Dir(machoPath)
			linkRoot := filepath.Dir(linkParent)

			machoPath = filepath.Join(linkRoot, symlinkPath)
		}

		fat, err := macho.OpenFat(machoPath)
		if err != nil && err != macho.ErrNotFat {
			log.Fatal(err.Error())
		}
		if err == macho.ErrNotFat {
			m, err = macho.Open(machoPath)
			if err != nil {
				log.Fatal(err.Error())
			}
		} else {
			for _, arch := range fat.Arches {
				if strings.Contains(strings.ToLower(arch.SubCPU.String(arch.CPU)), "arm64") {
					m = arch.File
					break
				}
			}
		}

		if !strings.Contains(strings.ToLower(m.FileHeader.SubCPU.String(m.CPU)), "arm64") {
			log.Fatal("[ERROR] can only disassemble arm64 binaries")
		}

		for _, sym := range m.Symtab.Syms {
			addr2SymMap[sym.Value] = sym.Name
		}

		if !allFuncs {
			if len(symbolName) > 0 && startVMAddr != 0 {
				log.Fatal("[ERROR] you can only use --all OR --symbol OR --vaddr (not a combination)")
			} else if len(symbolName) == 0 && startVMAddr == 0 {
				if m.Symtab != nil && len(m.Symtab.Syms) > 0 {
					var syms []string
					for _, sym := range m.Symtab.Syms {
						if _, err := m.GetFunctionForVMAddr(sym.Value); err == nil {
							syms = append(syms, sym.Name)
						}
					}
					promptVer := &survey.Select{
						Message:  "Choose a symbol to disassemble:",
						Options:  syms,
						PageSize: 20,
					}
					if err := survey.AskOne(promptVer, &symbolName); err != nil {
						if err == terminal.InterruptErr {
							fmt.Println("Exiting...")
							os.Exit(0)
						}
						log.Fatal(err.Error())
					}
				} else {
					log.Fatal("[ERROR] you must supply a --symbol OR --vaddr to disassemble")
				}
			}

			if len(symbolName) > 0 {
				symAddr, err = m.FindSymbolAddress(symbolName)
				if err != nil {
					log.Fatal(err.Error())
				}

				fn, err := m.GetFunctionForVMAddr(symAddr)
				if err != nil {
					log.Fatal(err.Error())
				}

				funcs = append(funcs, fn)

			} else if startVMAddr > 0 {
				fn, err := m.GetFunctionForVMAddr(startVMAddr)
				if err != nil {
					log.Fatal(err.Error())
				}

				funcs = append(funcs, fn)

				if startVMAddr != fn.StartAddr {
					isMiddle = true
				}
			}
		} else {
			funcs = m.GetFunctions()
		}

		var instrStr string
		var instrValue uint32
		var resutls [1024]byte
		var prevInstr *disassemble.Instruction

		for idx, fn := range funcs {

			data, err = m.GetFunctionData(fn)
			if err != nil {
				log.Fatalf("failed to get data for function %v: %v", fn, err)
			}

			r := bytes.NewReader(data)

			if name, ok := addr2SymMap[fn.StartAddr]; ok && !asJSON {
				if idx == 0 {
					fmt.Printf("%s:\n", name)
				} else {
					fmt.Printf("\n%s:\n", name)
				}
			} else {
				if idx == 0 {
					fmt.Printf("sub_%x:\n", fn.StartAddr)
				} else {
					fmt.Printf("\nsub_%x:\n", fn.StartAddr)
				}
			}

			symAddr = fn.StartAddr

			for {
				err = binary.Read(r, binary.LittleEndian, &instrValue)

				if err == io.EOF {
					break
				}

				if asJSON {
					instruction, err := disassemble.Decompose(symAddr, instrValue, &resutls)
					if err != nil {
						log.Fatal(err.Error())
					}

					instructions = append(instructions, *instruction)
				} else {
					// instruction, err := disassemble.Disassemble(symAddr, instrValue, &resutls)
					instruction, err := disassemble.Decompose(symAddr, instrValue, &resutls)
					if err != nil {
						if instrValue == 0xfeedfacf {
							fmt.Printf("%#08x:  %s\t.long\t%#x ; (possible embedded MachO)\n", uint64(symAddr), disassemble.GetOpCodeByteString(instrValue), instrValue)
							break
						} else if instrValue == 0x201420 {
							fmt.Printf("%#08x:  %s\tgenter\n", uint64(symAddr), disassemble.GetOpCodeByteString(instrValue))
							continue
						} else if instrValue == 0x00201400 {
							fmt.Printf("%#08x:  %s\tgexit\n", uint64(symAddr), disassemble.GetOpCodeByteString(instrValue))
							continue
						} else if instrValue == 0xe7ffdefe || instrValue == 0xe7ffdeff {
							fmt.Printf("%#08x:  %s\ttrap\n", uint64(symAddr), disassemble.GetOpCodeByteString(instrValue))
							continue
						} else if instrValue > 0xffff0000 {
							fmt.Printf("%#08x:  %s\t.long\t%#x ; (probably a jump-table)\n", uint64(symAddr), disassemble.GetOpCodeByteString(instrValue), instrValue)
							break
						} else if prevInstr != nil && strings.Contains(prevInstr.Operation.String(), "braa") {
							break
						} else if (instrValue & 0xfffffC00) == 0x00201000 {
							Xr := disassemble.Register((instrValue & 0x1F) + 34)
							m := (instrValue >> 5) & 0x1F
							if m == 17 {
								if instrValue&0x1F == 0 {
									fmt.Printf("%#08x:  %s\tamxset\n", uint64(symAddr), disassemble.GetOpCodeByteString(instrValue))
								} else {
									fmt.Printf("%#08x:  %s\tamxclr\n", uint64(symAddr), disassemble.GetOpCodeByteString(instrValue))
								}
							} else {
								fmt.Printf("%#08x:  %s\t%s\t%s\n", uint64(symAddr), disassemble.GetOpCodeByteString(instrValue), opName(m), Xr.String())
							}
							continue
						} else if instrValue>>21 == 1 {
							fmt.Printf("%#08x:  %s\t.long\t%#x ; (possible unknown Apple instruction)\n", uint64(symAddr), disassemble.GetOpCodeByteString(instrValue), instrValue)
							continue
						}
						fmt.Printf("%#08x:  %s\t.long\t%#x ; (%s)\n", uint64(symAddr), disassemble.GetOpCodeByteString(instrValue), instrValue, err.Error())
						break
					}

					instrStr = instruction.String()

					if instruction.Operation == disassemble.ARM64_MRS || instruction.Operation == disassemble.ARM64_MSR {
						var ops []string
						replaced := false
						for _, op := range instruction.Operands {
							if op.Class == disassemble.REG {
								ops = append(ops, op.Registers[0].String())
							} else if op.Class == disassemble.IMPLEMENTATION_SPECIFIC {
								sysRegFix := op.ImplSpec.GetSysReg().String()
								if len(sysRegFix) > 0 {
									ops = append(ops, sysRegFix)
									replaced = true
								}
							}
							if replaced {
								instrStr = fmt.Sprintf("%s\t%s", instruction.Operation, strings.Join(ops, ", "))
							}
						}
					}

					if instruction.Encoding == disassemble.ENC_BL_ONLY_BRANCH_IMM || instruction.Encoding == disassemble.ENC_B_ONLY_BRANCH_IMM {
						if name, ok := addr2SymMap[uint64(instruction.Operands[0].Immediate)]; ok {
							instrStr = fmt.Sprintf("%s\t%s", instruction.Operation, name)
						}
					}

					if instruction.Encoding == disassemble.ENC_CBZ_64_COMPBRANCH {
						if name, ok := addr2SymMap[uint64(instruction.Operands[1].Immediate)]; ok {
							instrStr += fmt.Sprintf(" ; %s", name)
						}
					}

					if (prevInstr != nil && prevInstr.Operation == disassemble.ARM64_ADRP) && (instruction.Operation == disassemble.ARM64_ADD || instruction.Operation == disassemble.ARM64_LDR) {
						adrpRegister := prevInstr.Operands[0].Registers[0]
						adrpImm := prevInstr.Operands[1].Immediate
						if instruction.Operation == disassemble.ARM64_LDR && adrpRegister == instruction.Operands[1].Registers[0] {
							adrpImm += instruction.Operands[1].Immediate
						} else if instruction.Operation == disassemble.ARM64_ADD && adrpRegister == instruction.Operands[1].Registers[0] {
							adrpImm += instruction.Operands[2].Immediate
						}
						if name, ok := addr2SymMap[uint64(adrpImm)]; ok {
							instrStr += fmt.Sprintf(" ; %s", name)
						} else if cstr, err := m.GetCString(adrpImm); err == nil {
							if len(cstr) > 200 {
								instrStr += fmt.Sprintf(" ; %#v...", cstr[:200])
							} else if len(cstr) > 1 {
								instrStr += fmt.Sprintf(" ; %#v", cstr)
							}
						}
					}

					if isMiddle && startVMAddr == symAddr {
						fmt.Printf("👉%08x:  %s\t%s\n", uint64(symAddr), disassemble.GetOpCodeByteString(instrValue), instrStr)
					} else {
						fmt.Printf("%#08x:  %s\t%s\n", uint64(symAddr), disassemble.GetOpCodeByteString(instrValue), instrStr)
					}

					prevInstr = instruction
				}

				symAddr += uint64(binary.Size(uint32(0)))
			}

			if asJSON {
				if dat, err := json.MarshalIndent(instructions, "", "   "); err == nil {
					fmt.Println(string(dat))
				}
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".arm64-cgo" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".disass")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
