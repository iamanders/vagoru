package command

import (
	"testing"

	"github.com/iamanders/vagoru/pkg/database"
)

func TestStartTracking(t *testing.T) {
	database.ResetDb()
	projectTitle := "Test project 1"
	StartTracking(projectTitle)
	tracking, _ := database.CurrentTracking()
	if tracking.Title != projectTitle {
		t.Errorf("\"%s\" should be equal to \"%s\"", tracking.Title, projectTitle)
	}
}
