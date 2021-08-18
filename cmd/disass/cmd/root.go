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
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/blacktop/go-macho"

	"github.com/blacktop/arm64-cgo/disassemble"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.disass.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("json", "j", false, "Output as JSON")

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

		var isMiddle bool
		var symAddr uint64
		var data []byte

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
			log.Fatal(err)
		}
		if err == macho.ErrNotFat {
			m, err = macho.Open(machoPath)
			if err != nil {
				log.Fatal(err)
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

		if len(symbolName) > 0 && startVMAddr != 0 {
			log.Fatal("[ERROR] you can only use --symbol OR --vaddr (not both)")
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
					log.Fatal(err)
				}
			} else {
				log.Fatal("[ERROR] you must supply a --symbol OR --vaddr to disassemble")
			}
		}

		if len(symbolName) > 0 {
			symAddr, err = m.FindSymbolAddress(symbolName)
			if err != nil {
				log.Fatal(err)
			}

			fn, err := m.GetFunctionForVMAddr(symAddr)
			if err != nil {
				log.Fatal(err)
			}

			data, err = m.GetFunctionData(fn)
			if err != nil {
				log.Fatal(err)
			}

		} else if startVMAddr > 0 {
			fn, err := m.GetFunctionForVMAddr(startVMAddr)
			if err != nil {
				log.Fatal(err)
			}

			data, err = m.GetFunctionData(fn)
			if err != nil {
				log.Fatal(err)
			}

			symAddr = fn.StartAddr

			if startVMAddr != fn.StartAddr {
				isMiddle = true
			}

			for _, sym := range m.Symtab.Syms {
				if sym.Value == fn.StartAddr {
					symbolName = sym.Name
				}
			}
		}

		var instrValue uint32
		r := bytes.NewReader(data)

		if len(symbolName) > 0 && !asJSON {
			fmt.Println(symbolName + ":")
		}

		var resutls [1024]byte

		for {
			err = binary.Read(r, binary.LittleEndian, &instrValue)

			if err == io.EOF {
				break
			}

			if asJSON {
				instruction, err := disassemble.Decompose(symAddr, instrValue, &resutls)
				if err != nil {
					log.Fatal(err)
				}

				instructions = append(instructions, *instruction)
			} else {
				instruction, err := disassemble.Disassemble(symAddr, instrValue, &resutls)
				if err != nil {
					log.Fatal(err)
				}
				if isMiddle && startVMAddr == symAddr {
					fmt.Printf("👉%08x:  %s\t%s\n", uint64(symAddr), disassemble.GetOpCodeByteString(instrValue), instruction)
				} else {
					fmt.Printf("%#08x:  %s\t%s\n", uint64(symAddr), disassemble.GetOpCodeByteString(instrValue), instruction)
				}
			}

			symAddr += uint64(binary.Size(uint32(0)))
		}

		if asJSON {
			if dat, err := json.MarshalIndent(instructions, "", "   "); err == nil {
				fmt.Println(string(dat))
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
