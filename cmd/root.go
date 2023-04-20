/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

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
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

type GrepResult struct {
	LineNumber int
	Line       string
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gogrep",
	Short: "This tool is simple grep command.",
	Long: `This tool searches for lines that contain the specified keyword .
	How to use this tool is below.
	gogrep <search keyword> <path>
	<path> is optional. if <path> is not specified, current directory is used.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("error! This tool usage is below. please try again.")
			fmt.Println("gogrep <search keyword> <path>")
			fmt.Println("<path> is optional. if <path> is not specified, current directory is used.")
			return
		}
		keyword := args[0]
		var path string
		if len(args) == 1 {
			path = "."
		} else {
			path = args[1]
		}

		if _, err := os.Stat(path); os.IsNotExist(err) {
			fmt.Printf("gogrep: no matches found: %s", path)
			return
		}

		err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {

			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			file, err := os.Open(path)

			if err != nil {
				return nil
			}

			defer file.Close()

			scanner := bufio.NewScanner(file)

			showedFileName := false

			i := 1
			for scanner.Scan() {
				line := scanner.Text()
				if strings.Contains(line, keyword) {
					if !showedFileName {
						fmt.Println(path)
						showedFileName = true
					}
					fmt.Println(i, line)
				}
				i += 1
			}

			return nil
		})

		if err != nil {
			fmt.Println(err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gogrep.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
