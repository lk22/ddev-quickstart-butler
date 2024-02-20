package main

import "fmt"
import "os"
import "os/exec"

func renderWelcomeMessage() {
	fmt.Println("Hello")
	fmt.Println("|----------------------------------|")
	fmt.Println("| Welcome to ddev cli butler       |")
	fmt.Println("| Version: 0.0.1		   |")
	fmt.Println("|----------------------------------|")
	fmt.Println("")
	fmt.Println("i am your quickstart helper for ddev, I will help you to initialize your next project")
	fmt.Println("I currently support the following project types")
	fmt.Println("1. WordPress")
	fmt.Println("Tell me what project you want to start and I will help you to initialize it")
	fmt.Println("")
}

func initializeProjectFolder(folder string) {
	err := os.Mkdir(folder, 0755)
	if err != nil {
		fmt.Println("Error creating the project directory")
	}
}

func runDdevConfigCommand() {
	cmd := exec.Command("ddev", "config", "--project-type=wordpress")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func runDdevStartCommand() {
	cmd := exec.Command("ddev", "start")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func runDdevWpCoreDownloadCommand() {
	cmd := exec.Command("ddev", "wp", "core", "download")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func runDdevLaunchCommand() {
	cmd := exec.Command("ddev", "launch")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func initializeWordPressProject(project_name string) {
	// create the project directory
	initializeProjectFolder(project_name)
	os.Chdir(project_name)
	runDdevConfigCommand()
	runDdevStartCommand()
	runDdevWpCoreDownloadCommand()
	runDdevLaunchCommand()
}

/**
* main function for ddev cli butler
*/
func main() {
	renderWelcomeMessage()
	// read the input from the user
	var project_name string
	var project_type string
	fmt.Println("Please enter the project name")
	fmt.Scanln(&project_name)
	fmt.Println("we have received the project name to initialize: ", project_name)
	fmt.Scanln(&project_type)
	fmt.Println("You have selected the project type: " + project_type)
	fmt.Println("Please wait while I initialise the project for you")
	// create the project directory

	switch project_type {
	case "wordpress":
		initializeWordPressProject(project_name)
	default:
		fmt.Println("Project type not supported")
	}
}