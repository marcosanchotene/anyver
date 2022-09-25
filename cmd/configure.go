/*
Copyright © 2022 Marco Sanchotene <marco.sanchotene@outlook.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure the name and directory of current tool",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tools := getTools()
		currentTool := viper.GetString("current-tool")
		delete(tools, currentTool)
		tools[name] = directory
		if name != "" {
			viper.Set("current-tool", name)
			fmt.Printf("Tool name has been changed to %s.\n", name)
		}
		viper.Set("current-directory", directory)
		fmt.Printf("Tool directory has been changed to %s.\n", name)
		viper.Set("tools", tools)
		viper.WriteConfig()
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)

	configureCmd.Flags().StringVarP(&directory, "directory", "d", "", "set the directory")
	configureCmd.Flags().StringVarP(&name, "name", "n", "", "set the name")
	configureCmd.MarkFlagRequired("directory")
}
