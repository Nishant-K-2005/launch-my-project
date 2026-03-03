package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

var MyConfig Config

func main() {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		fmt.Println("\n\n Stopping all Services....")
		os.Exit(0)
	}()

	if len(os.Args) < 2 {
		fmt.Println("Error: Please provide a project name.")
		fmt.Println("Usage: launch <project_name>")
		os.Exit(1)
	}
	projectName := os.Args[1]

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Failed to find home directory", err)
		os.Exit(1)
	}
	configPath := filepath.Join(homeDir, ".config", "launcher.json")
	fileBytes, err := os.ReadFile(configPath)

	if err != nil {
		fmt.Print("Failed to read Json file", err)
		os.Exit(1)
	}

	err = json.Unmarshal(fileBytes, &MyConfig)
	if err != nil {
		fmt.Println("Failed to unmarshal json data:", err)
		os.Exit(1)
	}
	_, ok := MyConfig[projectName]
	if !ok {
		fmt.Println("Project not found! Please check your project name for typo")
		os.Exit(1)
	}
	LaunchProject(projectName)
}
