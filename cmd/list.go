/*
Copyright © 2022 Marco Sanchotene <marco.sanchotene@outlook.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	listToolsFlag bool
	listCmd       = &cobra.Command{
		Use:   "list",
		Short: "List tools or binaries available",
		Long: `Display a list of files on directory currently configure or 
		a list of tools configured (when the --tools flag is used).`,
		Run: func(cmd *cobra.Command, args []string) {
			exitIfNoToolIsConfigured()

			if listToolsFlag {
				listTools()
			} else {
				listBinaries()
				fmt.Print("\nTo list tools, use the --tools or -t flag.")
			}

			fmt.Println("")

			displayCurrentConfiguration()
		},
	}
)

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&listToolsFlag, "tools", "t", false, "list tools")
}

func listBinaries() {
	directory, files := getBinaries()
	fmt.Printf("\nBinaries on %s:\n\n", directory)

	for _, file := range files {
		fmt.Println(file)
	}
}

func listTools() {
	names := getToolsNames()
	fmt.Printf("\nTools available:\n\n")

	for _, name := range names {
		fmt.Println(name)
	}
}
