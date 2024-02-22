package main

import (
	"fmt"
	"ddevbutler/butler"
)

func renderWelcomeMessage() {
	fmt.Println("")
	fmt.Println("|----------------------------------|")
	fmt.Println("| Welcome to ddev cli butler       |")
	fmt.Println("| Version: 0.0.1		   |")
	fmt.Println("|----------------------------------|")
	fmt.Println("")
	fmt.Println("i am your quickstart helper for ddev, I will help you to initialize your next project")
	fmt.Println("I currently support the following project types")
	fmt.Println("1. WordPress")
	fmt.Println("")
	fmt.Println("Want to read more about the ddev quickstarts visit: https://ddev.readthedocs.io/en/stable/users/quickstart/")
	fmt.Println("Tell me what project you want to start and I will help you to initialize it")
	fmt.Println("")
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

	butler.InitializeProject(project_name, project_type)
}