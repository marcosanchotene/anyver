/*
Copyright © 2022 Marco Sanchotene <marco.sanchotene@outlook.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a command line interface (CLI) tool for usage",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		tools := getTools()
		if _, ok := tools[args[0]]; ok {
			fmt.Println("Tool is already configured.")
		} else {
			if len(tools) == 0 {
				viper.Set("current-tool", args[0])
				viper.Set("current-directory", directory)
			}
			tools[args[0]] = directory
			viper.Set("tools", tools)
			viper.WriteConfig()
			fmt.Printf("Tool %s has been added.\n", args[0])
		}
		displayCurrentConfiguration()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&directory, "directory", "d", "", "set the directory where your binaries are")
	addCmd.MarkFlagRequired("directory")
}
