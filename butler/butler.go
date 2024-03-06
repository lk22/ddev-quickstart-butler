package butler

import (
	"os"
	"os/exec"
	"time"

	"github.com/fatih/color"
)

/********* Butler interface *********/
type Butler interface {
	Trigger(command string, args string)
	checkError(err error)
	InitializeScaffoldingProjectTypesProcedure()
}

/********* Project interface *********/
type Project interface {
	InitializeProject(project_name string, project_type string)
	RunDdevConfigCommand(project_name string)
	RunDdevStartCommand(project_name string)
	RunDdevLaunchCommand(project_name string)
}

/********* ButlerCommand struct *********/
type ButlerCommand struct {
	Command  string
	Arg      string
	Multiple []string
}

var Command ButlerCommand

/*
* DDEV CLI Butler Utility functions
 */
func UtilStartDdevProject() {
	Command.Trigger("start", "")
}

/********* Utility function to stop project *********/
func UtilStopDdevProject() {
	Command.Trigger("stop", "")
}

/********* Utility function to remove project *********/
func UtilRemoveProject(project_name string) {
	Command.Arg = project_name
	Command.Trigger("rm", Command.Arg)
}

/********* Utility function to import database *********/
func UtilImportDatabase(db_file string) {
	Command.Arg = db_file
	Command.Trigger("import-db", Command.Arg)
}

/********* Utility function to export database *********/
func UtilExportDatabase() {
	Command.Trigger("export-db", "")
}

/********* Utility function to list all projects *********/
func UtilListProjects() {
	Command.Trigger("list", "")
}

/********* Utility function to trigger help command *********/
func Help() {
	Command.Trigger("help", "")
}

/*
* Main function to trigger commands from the butler
 */
func (butlerCommand *ButlerCommand) Trigger(command string, arg string) {
	cmd := exec.Command("ddev", command)
	if arg != "" {
		cmd = exec.Command("ddev", command, arg)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

/********* Checking for error *********/
func (bc *ButlerCommand) checkError(err error) {
	if err != nil {
		println("Error running the command")
	}
}

/********* Initialize project folder *********/
func InitializeProjectFolder(folder string) {
	err := os.Mkdir(folder, 0755)
	if err != nil {
		println("Error creating the project directory")
	}
}

/*
* Perfomming scaffold procedure for the WordPress Quickstart
* https://ddev.readthedocs.io/en/latest/users/quickstart/#wordpress
 */
func ScaffoldWordpressProject() {
	color.Green("Scaffolding WordPress project")
	time.Sleep(2 * time.Second)
	Command.Trigger("config", "--project-type=wordpress --php-version=8.1")

	color.Green("Starting WordPRess project")
	time.Sleep(2 * time.Second)
	Command.Trigger("start", "")

	color.Green("Downloading WordPress core hold on!...")
	time.Sleep(2 * time.Second)
	Command.Trigger("wp", "core download")

	color.Green("Launching WordPress project")
	time.Sleep(2 * time.Second)
	Command.Trigger("launch", "")
}
