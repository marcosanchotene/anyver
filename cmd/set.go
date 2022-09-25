/*
Copyright © 2022 Marco Sanchotene <marco.sanchotene@outlook.com>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Select the binary you want from a list",
	Long: `Display a list with the available binaries and let
	you select which one should be configured for global usage.`,
	Run: func(cmd *cobra.Command, args []string) {
		exitIfNoToolIsConfigured()
		displayCurrentConfiguration()
		fmt.Println("")
		setBinary()
		displayCurrentConfiguration()
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}

func setBinary() {
	directory, files := getBinaries()
	prompt := promptui.Select{
		Label: "Select binary",
		Items: files,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	user, err := user.Current()

	if err != nil {
		log.Fatalf(err.Error())
	}

	homeDirectory := user.HomeDir

	linkPath := homeDirectory + "/.local/bin/" + viper.GetString("current-tool")

	if _, err := os.Lstat(linkPath); err == nil {
		if err := os.Remove(linkPath); err != nil {
			fmt.Printf("failed to unlink: %v\n", err)
		}
	}

	binaryPath := directory + "/" + result

	fmt.Printf("Setting link for %s at %s\n", binaryPath, linkPath)

	err = os.Symlink(binaryPath, linkPath)

	if err != nil {
		fmt.Printf("Error %v\n", err)
		return
	}

	currentTool := viper.GetString("current-tool")

	fmt.Printf("Version %q configured for %s\n", result, currentTool)
	viper.Set("current-binary", result)
	viper.WriteConfig()
}
