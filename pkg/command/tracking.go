package command

import (
	"os"

	"github.com/iamanders/vagoru/pkg/database"
)

func StartTracking() string {
	// Args
	if len(os.Args) < 3 {
		return "Missing arg 3 - project title, aborting.."
	}
	projectTitle := os.Args[2]

	// Already tracking?
	currentTracking, _ := database.CurrentTracking()
	if currentTracking.Title != "" {
		return "Already tracking, aborting.."
	}

	// Insert
	database.StartTracking(projectTitle)
	return "Started tracking " + projectTitle
}
