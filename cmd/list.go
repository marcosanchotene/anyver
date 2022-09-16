/*
Copyright © 2022 Marco Sanchotene <marco.sanchotene@outlook.com>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List binaries available on directory set",
	Long:  `Display a list with the files on the directory set with the 'directory' command.`,
	Run: func(cmd *cobra.Command, args []string) {
		checkTools()
		displayCurrentConfiguration()

		directory, files := getBinaries()
		fmt.Printf("\nBinaries on %s:\n\n", directory)

		for _, file := range files {
			fmt.Println(file)
		}

		fmt.Println("")
		displayCurrentConfiguration()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
