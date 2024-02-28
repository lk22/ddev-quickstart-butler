package butler

import (
	"ddevbutler/wordpress"
	"fmt"
	"os"
	"os/exec"

	"github.com/nexidian/gocliselect"
)

type Butler interface {
	Trigger(command string, args string)
	checkError(err error)
}

type ButlerCommand struct {
	Command  string
	Arg      string
	Multiple []string
}

var GlobalButlerCommand ButlerCommand

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

func ProcessStrings(strings []string) {
	for _, str := range strings {
		fmt.Println(str)
	}
}

func UtilStartDdevProject() {
	GlobalButlerCommand.Trigger("start", "")
}

func UtilStopDdevProject() {
	GlobalButlerCommand.Trigger("stop", "")
}

func UtilImportDatabase(db_file string) {
	GlobalButlerCommand.Trigger("import-db", db_file)
}

func (butlerCommand *ButlerCommand) Trigger(command string, arg ...string) {
	cmd := exec.Command("ddev", command)
	if len(arg) > 0 {
		cmd = exec.Command("ddev", command, arg[0])
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func (bc *ButlerCommand) checkError(err error) {
	if err != nil {
		fmt.Println("Error running the command")
	}
}
