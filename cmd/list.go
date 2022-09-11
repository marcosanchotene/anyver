/*
Copyright © 2022 Marco Sanchotene <marco.sanchotene@outlook.com>

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List binaries available on directory set",
	Long:  `Display a list with the files on the directory set with the 'directory' command.`,
	Run: func(cmd *cobra.Command, args []string) {
		checkDirectory()
		checkName()

		directory, files := getBinaries()
		fmt.Printf("\nBinaries on %s:\n\n", directory)

		for _, file := range files {
			fmt.Println(file)
		}

		fmt.Println("")
		getCurrentBinary()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func getBinaries() (string, []string) {
	directory := viper.GetString("directory")
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}
	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}
	return directory, fileNames
}
