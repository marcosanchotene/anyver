/*
Copyright © 2022 Marco Sanchotene <marco.sanchotene@outlook.com>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a command line interface (CLI) tool from configuration",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		tools := viper.GetStringMapString("tools")
		tool := args[0]
		if _, ok := tools[args[0]]; ok {
			if viper.Get("current-tool") == tool {
				fmt.Println("Cannot remove currently configured tool.")
				fmt.Println("Please change it first with the 'change' command.")
				os.Exit(1)
			}
			delete(tools, tool)
			viper.Set("tools", tools)
			viper.WriteConfig()
			fmt.Printf("Tool %s has been removed.\n", tool)
		} else {
			fmt.Printf("There is no tool configured with name %s.\n", tool)
			os.Exit(1)
		}
		displayCurrentConfiguration()
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
