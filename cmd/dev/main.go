package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("nodemon", "--exec", "go", "run", "./cmd/server", "--signal", "SIGTERM")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running command %s: %v\n", "nodemon", err)
		os.Exit(1)
	}
}
