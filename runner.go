package main

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
)

var wg sync.WaitGroup

func LaunchProject(project string) {
	for _, cmnd := range MyConfig[project].Commands {
		wg.Add(1)
		go func(command string) {
			defer wg.Done()
			cmd := exec.Command("bash", "-c", command)
			cmd.Dir = MyConfig[project].Path
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println("Error while running", cmnd, err)
				os.Exit(1)
			}
		}(cmnd)
	}
	wg.Wait()
}
