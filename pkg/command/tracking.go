package command

import "github.com/iamanders/vagoru/pkg/database"

func StartTracking() string {
	// TODO: Args

	// TODO: Already tracking?
	currentTracking, _ := database.CurrentTracking()
	if currentTracking.Title != "" {
		return "Already tracking!!1111111§§§"
	}

	// TODO: Insert
	database.StartTracking("New project")
	return "-"
}
