package butler

import (
	"ddevbutler/wordpress"
	"fmt"
	"os"
	"os/exec"

	"github.com/nexidian/gocliselect"
)

func InitializeProjectFolder(folder string) {
	err := os.Mkdir(folder, 0755)
	if err != nil {
		fmt.Println("Error creating the project directory")
	}
}

func InitializeScaffoldingProjectTypesProcedure() {
	menu := gocliselect.NewMenu("Select the project type you want to initialize")
	menu.AddItem("Wordpress", "wordpress")
	menu.AddItem("Drupal", "drupal")
	menu.AddItem("Laravel", "laravel")
	menu.AddItem("Go back", "back")

	project_type := menu.Display()

	if project_type == "back" {
		return
	}

	var project_name string
	fmt.Println("You have selected the project type: " + project_type)
	fmt.Println("Please enter the project name")
	fmt.Scanln(&project_name)

	if project_name == "" {
		fmt.Println("Please enter a valid project name")
		return
	}

	InitializeProject(project_name, project_type)
}

func InitializeProject(project_name string, project_type string) {
	switch project_type {
	case "wordpress":
		InitializeProjectFolder(project_name)
		os.Chdir(project_name)
		wordpress.InitializeProject(project_name)
	default:
		fmt.Println("Sorry, I don't know how to initialize a project of type: " + project_type)
		InitializeScaffoldingProjectTypesProcedure()
	}
}

func UtilStartDdevProject() {
	cmd := exec.Command("ddev", "start")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func UtilStopDdevProject() {
	cmd := exec.Command("ddev", "stop")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
