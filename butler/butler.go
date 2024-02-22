package butler

import (
	"fmt"
	"os"
	"ddevbutler/wordpress"
)

func InitializeProjectFolder(folder string) {
	err := os.Mkdir(folder, 0755)
	if err != nil {
		fmt.Println("Error creating the project directory")
	}
}

func PrintSupportedProjectTypes() {
	fmt.Println("I currently support the following project types");
	fmt.Println("1. Wordpress")
}

func InitializeProject(project_name string, project_type string) {
	switch project_type {
	case "wordpress":
		InitializeProjectFolder(project_name)
		os.Chdir(project_name)
		wordpress.InitializeProject(project_name)
	default:
		fmt.Println("Sorry, I don't know how to initialize a project of type: " + project_type);
		PrintSupportedProjectTypes()
	}
}