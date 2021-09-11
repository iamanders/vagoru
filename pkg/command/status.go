package command

import "github.com/iamanders/vagoru/pkg/database"

func Status() string {
	project, err := database.CurrentTracking()
	if err != nil {
		return "No project started"
	} else {
		return "Current " + project.Title
	}
}
