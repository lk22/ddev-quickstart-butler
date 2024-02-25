package main

import (
	"ddevbutler/butler"
	"fmt"

	"github.com/fatih/color"
	"github.com/nexidian/gocliselect"
)

func renderWelcomeMessage() {
	fmt.Println("")
	color.Blue("|----------------------------------|")
	color.Blue("| Welcome to ddev cli butler       |")
	color.Blue("| Version: 0.0.1		   |")
	color.Blue("|----------------------------------|")
	fmt.Println("")
	fmt.Println("i am your quickstart helper for ddev, I will help you to initialize your next project")
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
	menu := gocliselect.NewMenu("Pick your project type")
	menu.AddItem("Wordpress", "wordpress")
	menu.AddItem("Backdrop", "backdrop")
	menu.AddItem("Craft CMS", "craftcms")
	menu.AddItem("Drupal", "drupal")

	project_type := menu.Display()
	var project_name string

	fmt.Println("You have selected the project type: " + project_type)
	fmt.Println("Please enter the project name")
	fmt.Scanln(&project_name)
	fmt.Println("Initializing project:", project_name)
	butler.InitializeProject(project_name, project_type)

}
