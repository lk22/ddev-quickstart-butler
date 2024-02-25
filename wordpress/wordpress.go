package wordpress

import (
	"os"
	"os/exec"
	"time"

	"github.com/fatih/color"
)

/*
* Run the ddev config command to fast track the project setup
 */
func RunDdevConfigCommand(project_name string) {
	color.Green("Configuring ", project_name)
	time.Sleep(2 * time.Second)
	cmd := exec.Command("ddev", "config", "--project-type=wordpress", "--php-version=8.1")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

/*
* Run the ddev start command to start the project
 */
func RunDdevStartCommand(project_name string) {
	color.Green("Starting", project_name, "..")
	time.Sleep(2 * time.Second)
	cmd := exec.Command("ddev", "start")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

/*
* Run the ddev wp core download command to download the wordpress core
 */
func RunDdevWpCoreDownloadCommand() {
	color.Green("Downloading WordPress core hold on!...")
	time.Sleep(2 * time.Second)
	cmd := exec.Command("ddev", "wp", "core", "download")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

/*
* Run the ddev launch command to launch the project
 */
func RunDdevLaunchCommand(project_name string) {
	color.Green("Launching", project_name, "..")
	time.Sleep(2 * time.Second)
	cmd := exec.Command("ddev", "launch")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

/*
* Run the ddev wp core install command to install the wordpress core in the project
 */
func RunDdevWpCoreInstallCommand(project_name string) {
	cmd := exec.Command("ddev", "wp", "core", "install", "--url="+project_name, "--title="+project_name, "--admin_user=admin", "--admin_password=password", "--admin_email=test@test.dk --prompt=admin_password")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

/*
* Initialize the WordPress project
 */
func InitializeProject(project_name string) {
	RunDdevConfigCommand(project_name)
	RunDdevStartCommand(project_name)
	RunDdevWpCoreDownloadCommand()
	RunDdevLaunchCommand(project_name)
}
