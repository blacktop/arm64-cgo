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
	"fmt"
	"io"
	"log"
	"os"

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
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

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

		if len(symbolName) > 0 && startVMAddr != 0 {
			log.Fatal("[ERROR] you can only use --symbol OR --vaddr (not both)")
		} else if len(symbolName) == 0 && startVMAddr == 0 {
			log.Fatal("[ERROR] you must supply a --symbol OR --vaddr to disassemble")
		}

		var isMiddle bool
		var symAddr uint64
		var data []byte

		m, err := macho.Open(args[0])
		if err != nil {
			log.Fatal(err)
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

		if len(symbolName) > 0 {
			fmt.Println(symbolName + ":")
		}

		for {
			err = binary.Read(r, binary.LittleEndian, &instrValue)

			if err == io.EOF {
				break
			}

			instruction, err := disassemble.Disassemble(symAddr, instrValue, true)
			if err != nil {
				log.Fatal(err)
			}

			if isMiddle && startVMAddr == symAddr {
				fmt.Printf("👉%08x:  %s\t%s\n", uint64(symAddr), disassemble.GetOpCodeByteString(instrValue), instruction)
			} else {
				fmt.Printf("%#08x:  %s\t%s\n", uint64(symAddr), disassemble.GetOpCodeByteString(instrValue), instruction)
			}

			symAddr += uint64(binary.Size(uint32(0)))
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
