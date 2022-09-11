/*
Copyright © 2022 Marco Sanchotene <marco.sanchotene@outlook.com>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var directory string
var name string

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure different options for this tool",
	Long: `Configure the directory and the name for your global binary.
	
	Directory: This is the absolute path pointing to your app binaries.
	Name: This is the name your app will have globally.`,
	Run: func(cmd *cobra.Command, args []string) {

		checkDirectory()
		checkName()
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)

	configureCmd.Flags().StringVarP(&directory, "directory", "d", "", "Set the directory where your binaries are.")
	configureCmd.Flags().StringVarP(&name, "name", "n", "", "Set the global name for your binary.")
}

func checkDirectory() {
	currentDirectory := viper.GetString("currentDirectory")

	if directory == "" {
		if currentDirectory == "" {
			fmt.Println("You have not set your binaries directory. Please set it before continuing.")
			fmt.Scanln(&directory)
		}
	}

	if directory != "" {
		viper.Set("currentDirectory", directory)
		viper.WriteConfig()
		fmt.Printf("Directory set at %s\n", directory)
	} else {
		fmt.Println("Current directory for binaries:", currentDirectory)
	}
}

func checkName() {
	currentName := viper.GetString("currentName")

	if name == "" {
		if currentName == "" {
			fmt.Println("You have not set your binary name. Please set it before continuing.")
			fmt.Scanln(&name)
		}
	}

	if name != "" {
		viper.Set("currentName", name)
		viper.WriteConfig()
		fmt.Printf("Name set as %s\n", name)
	} else {
		fmt.Println("Current name for binaries:", currentName)
	}
}
