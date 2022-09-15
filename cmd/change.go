/*
Copyright © 2022 Marco Sanchotene <marco.sanchotene@outlook.com>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var changeCmd = &cobra.Command{
	Use:   "change",
	Short: "Change the command line interface (CLI) tool for usage",
	Run: func(cmd *cobra.Command, args []string) {
		checkTools()
		exitIfOnlyOneTool()
		displayCurrentConfiguration()

		tools := viper.GetStringMapString("tools")
		names := make([]string, 0, len(tools))
		for key := range tools {
			names = append(names, key)
		}

		prompt := promptui.Select{
			Label: "Select tool",
			Items: names,
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		viper.Set("current-tool", result)
		viper.Set("current-directory", tools[result])
		viper.WriteConfig()
		displayCurrentConfiguration()
	},
}

func init() {
	rootCmd.AddCommand(changeCmd)
}

func exitIfOnlyOneTool() {
	tools := viper.GetStringMapString("tools")
	if len(tools) == 1 {
		fmt.Println("Cannot change, as there is only one tool configured.")
		os.Exit(1)
	}
}
