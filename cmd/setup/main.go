package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Starting setup for Go project...")

	// 1. ตรวจสอบ Go ติดตั้งหรือยัง
	checkGoInstalled()

	// 2. ติดตั้ง nodemon
	installNodemon()

	// 3. ดึง dependencies ด้วย go mod tidy
	runCommand("go", "mod", "tidy")

	// 4. สร้าง nodemon.json
	createNodemonConfig()

	// 5. สร้าง Binary (optional)
	fmt.Println("Building the project...")
	runCommand("go", "build", "-o", "app", "./cmd/app")

	fmt.Println("Setup completed! Use the following commands:")
	fmt.Println("- Run with nodemon: go run ./cmd/dev")
	fmt.Println("- Run binary: ./app")
}

func checkGoInstalled() {
	_, err := exec.LookPath("go")
	if err != nil {
		fmt.Println("Error: Go is not installed. Please install Go first.")
		os.Exit(1)
	}
	fmt.Println("Go is installed.")
}

func installNodemon() {
	_, err := exec.LookPath("nodemon")
	if err != nil {
		fmt.Println("Installing nodemon...")
		runCommand("npm", "install", "-g", "nodemon")
	} else {
		fmt.Println("nodemon is already installed.")
	}
}

func createNodemonConfig() {
	config := `{
  "watch": ["**/*.go"],
  "ext": "go",
  "exec": "go run ./cmd/app/main.go"
}`
	if _, err := os.Stat("nodemon.json"); os.IsNotExist(err) {
		fmt.Println("Creating nodemon.json...")
		if err := os.WriteFile("nodemon.json", []byte(config), 0644); err != nil {
			fmt.Println("Error creating nodemon.json:", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("nodemon.json already exists.")
	}
}

func runCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running command %s: %v\n", name, err)
		os.Exit(1)
	}
}
