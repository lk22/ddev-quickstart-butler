package wordpress

import (
	"os"
	"os/exec"
)

/*
* Run the ddev config command to fast track the project setup
*/
func RunDdevConfigCommand() {
	cmd := exec.Command("ddev", "config", "--project-type=wordpress")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

/*
* Run the ddev start command to start the project
*/
func RunDdevStartCommand() {
	cmd := exec.Command("ddev", "start")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

/*
* Run the ddev wp core download command to download the wordpress core
*/
func RunDdevWpCoreDownloadCommand() {
	cmd := exec.Command("ddev", "wp", "core", "download")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

/*
* Run the ddev launch command to launch the project
*/
func RunDdevLaunchCommand() {
	cmd := exec.Command("ddev", "launch")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

/*
* Initialize the WordPress project
*/
func InitializeProject(project_name string) {
	RunDdevConfigCommand()
	RunDdevStartCommand()
	RunDdevWpCoreDownloadCommand()
	RunDdevLaunchCommand()
}