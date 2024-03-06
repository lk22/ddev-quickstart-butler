package wordpress

import (
	Butler "ddevbutler/butler"
	"time"

	"github.com/fatih/color"
)

type WordPress interface {
	RunDdevConfigCommand(project_name string)
	RunDdevStartCommand(project_name string)
	RunDdevWpCoreDownloadCommand()
	RunDdevLaunchCommand(project_name string)
}

func RunDdevConfigCommand(project_name string) *WordPress {
	color.Green("Configuring ", project_name)
	time.Sleep(2 * time.Second)
	Butler.GlobalButlerCommand.Trigger("config", "--project-type=wordpress", "--php-version=8.1")
	return nil
}

func RunDdevStartCommand(project_name string) {
	color.Green("Starting", project_name, "..")
	time.Sleep(2 * time.Second)
	Butler.GlobalButlerCommand.Trigger("start", "")
}

func RunDdevWpCoreDownloadCommand() {
	color.Green("Downloading WordPress core hold on!...")
	time.Sleep(2 * time.Second)
	Butler.GlobalButlerCommand.Trigger("wp", "core", "download")
}

func RunDdevLaunchCommand(project_name string) {
	color.Green("Launching", project_name, "..")
	time.Sleep(2 * time.Second)
	Butler.GlobalButlerCommand.Trigger("launch", "")
}

func InittializeWordpressProject(project_name string) {
	RunDdevConfigCommand(project_name)
	RunDdevStartCommand(project_name)
	RunDdevWpCoreDownloadCommand()
	RunDdevLaunchCommand(project_name)
}
