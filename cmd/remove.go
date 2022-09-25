/*
Copyright © 2022 Marco Sanchotene <marco.sanchotene@outlook.com>
*/
package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a command line interface (CLI) tool from configuration",
	Run: func(cmd *cobra.Command, args []string) {
		exitIfNoToolIsConfigured()

		tools := getTools()
		names := getToolsNames()
		currentTool := viper.GetString("current-tool")
		prompt := promptui.Select{
			Label: "Select tool",
			Items: names,
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		delete(tools, result)
		viper.Set("tools", tools)
		viper.WriteConfig()
		fmt.Printf("Tool %s has been removed.\n", result)

		if result == currentTool {
			fmt.Println("Select another tool for current configuration.")
			changeTool()
		}

		displayCurrentConfiguration()
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
