package main

import (
	"ddevbutler/butler"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/nexidian/gocliselect"
)

/********* Outputting welcome message *********/
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

/********* Outputting project type menu *********/
func renderProjectTypeMenu() {
	println("Please choose a project type")

	numberOftypesChannel := make(chan int)

	menu := gocliselect.NewMenu("Choose a project type")
	menu.AddItem("WordPress", "wordpress")
	menu.AddItem("Drupal", "drupal")
	menu.AddItem("Laravel", "laravel")
	menu.AddItem("Symfony", "symfony")
	menu.AddItem("Go back", "back")
	action := menu.Display()

	if action == "back" {
		renderActionsMenu()
	} else {
		var project_name string
		fmt.Scanln(&project_name)

		// increment the number of times the user has entered a project name if the project name is empty
		// TODO: make sure the user has entered a project and validate if
		// TODO: its empty and make the user have 3 chances to enter a project name
		go func() {

			// does the channel exist?
			if numberOftypesChannel == nil {
				numberOftypesChannel = make(chan int)
			}

			select {
			case number := <-numberOftypesChannel:
				fmt.Println("Number of times the user has entered a project name: ", number)

				// is project name empty?
				if project_name == "" {
					number++
					numberOftypesChannel <- number

					if number == 1 {
						fmt.Println("Please provide me a project name")
						fmt.Scanln(&project_name)
					} else if number == 2 {
						fmt.Println("Please provide me a project name")
						fmt.Scanln(&project_name)
					} else if number == 3 || number > 3 {
						fmt.Println("You have entered a project name more than 3 times")
					}
				}
			}
		}()

		// close the channel
		defer close(numberOftypesChannel) // using defer for cleanup of the channel

		// go back to the main menu
		// closing the channel
		close(numberOftypesChannel)
		// and re render the main menu
		renderActionsMenu()
		// i want to count the number of times the user has entered a project name
		// for each time the enters a project name, i want to check if the project name is empty
	}
}

func reRenderButlerMenu() {
	fmt.Println("Command executed successfully")
	fmt.Println("What can i help you with next?")
	renderActionsMenu()
}

/*
* Render actions menu
* if action is chosen render project type menu
* else if action is start, perfom start project command from the directory directly
* else if action is stop, perform stop project command from the directory directly
* else if action is dbimport, ask for the database file and perform import database command from the directory directly
* else if action is quit, exit the program
 */
func renderActionsMenu() {
	menu := gocliselect.NewMenu("Choose an action i should perform")
	menu.AddItem("Scaffold new project", "scaffold")
	menu.AddItem("Import database", "dbimport")
	menu.AddItem("Export database", "dbexport")
	menu.AddItem("Starting project", "start")
	menu.AddItem("Stop project", "stop")
	menu.AddItem("remove project", "remove")
	menu.AddItem("List all projects", "list")
	menu.AddItem("Help", "help")
	menu.AddItem("Quit", "quit")
	action := menu.Display()

	switch action {
	case "scaffold":
		renderProjectTypeMenu()
	case "start":
		fmt.Println("Starting project")
		butler.UtilStartDdevProject()
		reRenderButlerMenu()
	case "stop":
		fmt.Println("Stopping project")
		butler.UtilStopDdevProject()
		reRenderButlerMenu()
	case "remove":
		var project_name string
		fmt.Println("Give me the name of the project you want to remove")
		fmt.Scanln(&project_name)
		butler.UtilRemoveProject(project_name)
		reRenderButlerMenu()
	case "dbimport":
		fmt.Println("Give me the path to the database file")
		var db_file string
		fmt.Scanln(&db_file)
		butler.UtilImportDatabase(db_file)
		reRenderButlerMenu()
	case "dbexport":
		fmt.Println("Exporting database")
		butler.UtilExportDatabase()
		reRenderButlerMenu()
	case "list":
		fmt.Println("Listing all projects")
		butler.UtilListProjects()
		reRenderButlerMenu()
	case "help":
		fmt.Println("Help")
		butler.Help()
		reRenderButlerMenu()
	case "quit":
		fmt.Println("Goodbye!")
		os.Exit(0)
		return
	}
}

/********* Main function *********/
func main() {
	renderWelcomeMessage()
	renderActionsMenu()
}
