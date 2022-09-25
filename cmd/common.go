/*
Copyright © 2022 Marco Sanchotene <marco.sanchotene@outlook.com>
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/viper"
)

var (
	name      string
	directory string
)

func displayCurrentConfiguration() {
	fmt.Println("")
	fmt.Println("Current configuration")
	displayCurrentTool()
	displayCurrentDirectory()
	displayCurrentBinary()
}

func displayCurrentTool() {
	if viper.GetString("current-tool") != "" {
		fmt.Printf("Tool: %s\n", viper.GetString("current-tool"))
	}
}

func displayCurrentDirectory() {
	if viper.GetString("current-directory") != "" {
		fmt.Printf("Directory: %s\n", viper.GetString("current-directory"))
	}
}

func displayCurrentBinary() {
	if viper.GetString("current-binary") == "" {
		fmt.Println("Current binary: None. Please set it with the 'set' command.")
	} else {
		fmt.Printf("Binary: %s\n", viper.GetString("current-binary"))
	}
}

func exitIfNoToolIsConfigured() {
	tools := viper.GetStringMapString("tools")

	if len(tools) == 0 {
		fmt.Println("You have not added any tool. Please add with the 'add' command before continuing.")
		os.Exit(1)
	}
}

func getBinaries() (string, []string) {
	currentTool := viper.GetString("current-tool")
	tools := viper.GetStringMapString("tools")
	directory := tools[currentTool]
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

func getTools() map[string]string {
	tools := viper.GetStringMapString("tools")
	return tools
}

func getToolsNames() []string {
	tools := getTools()
	names := make([]string, 0, len(tools))
	for key := range tools {
		names = append(names, key)
	}
	return names
}
