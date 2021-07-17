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
	"os"

	"github.com/blacktop/go-macho"

	"github.com/blacktop/arm64-cgo/disassemble"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "disass",
	Short: "A brief description of your application",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		m, err := macho.Open(args[0])
		if err != nil {
			panic(err)
		}

		symAddr, err := m.FindSymbolAddress(args[1])
		if err != nil {
			panic(err)
		}

		fn, err := m.GetFunctionForVMAddr(symAddr)
		if err != nil {
			panic(err)
		}

		data, err := m.GetFunctionData(fn)
		if err != nil {
			panic(err)
		}

		var instrValue uint32
		r := bytes.NewReader(data)

		fmt.Println(args[1] + ":")

		for {
			addr, _ := r.Seek(0, io.SeekCurrent)

			err = binary.Read(r, binary.LittleEndian, &instrValue)

			if err == io.EOF {
				break
			}

			symAddr += uint64(addr)

			instruction, err := disassemble.Disassemble(symAddr, instrValue, true)
			if err != nil {
				panic(err)
			}

			fmt.Printf("%#08x:  %s\t%s\n", uint64(symAddr), disassemble.GetOpCodeByteString(instrValue), instruction)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.disass.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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
