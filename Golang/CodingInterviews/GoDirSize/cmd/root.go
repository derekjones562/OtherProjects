/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "CodingInterviews",
	Short: "Command line utility for getting the size of directories",
	Long: `Command line utility in Go which takes as arguments a list of directories.
The program should output the sizes of each of the individual directories 
passed as well as a cumulative total.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: GoDirSize,
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.CodingInterviews.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("recursive", "r", false, "If a --recursive flag is provided, output the sizes of each of the individual directories passed and sub-directories recursively as well as a cumulative total")
	rootCmd.Flags().BoolP("human", "", false, "If a \"--human\" flag is passed, format the sizes to be human friendly by outputting the size in the most appropriate unit of bytes. For example, 304K for 304,000 bytes and 300M for 300000000 bytes.")
}

func GoDirSize(cmd *cobra.Command, args []string) {
	recursive, _ := cmd.Flags().GetBool("recursive")
	human, _ := cmd.Flags().GetBool("human")
	totalBytes := int64(0)
	for _, arg := range args {
		if string(arg[len(arg)-1]) == "/" {
			arg = arg[:(len(arg) - 1)] // remove trailing "/"
		}
		size, err := getDirSize(arg, recursive, human)
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			continue
		}
		totalBytes += *size
	}
	fmt.Printf("Total Bytes: %s\n", humanFlag(human, totalBytes))

}

func getDirSize(dirPath string, recursive, human bool) (*int64, error) {
	// dir, err := os.Open(dirPath)
	// if os.IsNotExist(err) {
	// 	return nil, fmt.Errorf("the directory named %s does not exist", dirPath)
	// }
	// defer dir.Close()

	files, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	var size int64
	var dirSizeTotal int64
	for _, file := range files {
		if !file.IsDir() {
			info, err := file.Info()
			if err != nil {
				return nil, err
			}
			size += info.Size()
		} else if recursive {
			dirDize, err := getDirSize(fmt.Sprintf("%s/%s", dirPath, file.Name()), recursive, human)
			if err != nil {
				return nil, err
			}
			dirSizeTotal += *dirDize
		}
	}

	fmt.Printf("%s: %s\n", dirPath, humanFlag(human, size))
	size += dirSizeTotal
	return &size, nil

}

type SizePostfix int

const (
	Byte SizePostfix = iota + 1 // EnumIndex = 1
	Kilo                        // EnumIndex = 2
	Mega                        // EnumIndex = 3
	Giga                        // EnumIndex = 4
	Tera                        // EnumIndex = 5
	Peta                        // EnumIndex = 6
	Exa                         // EnumIndex = 7
)

func (s SizePostfix) String() string {
	return [...]string{"", "K", "M", "G", "T", "P", "E"}[s-1]
}

func humanFlag(flag bool, size int64) string {
	postfix := SizePostfix(1)
	if flag {
		for size > 1000 {
			size = size / 1000
			postfix++
		}
		return fmt.Sprintf("%d %s", size, postfix.String())
	}
	return fmt.Sprintf("%d", size)
}
