package database

import "testing"

func TestStartTracking(t *testing.T) {
	ResetDb()
	projectTitle := "Test project 1"
	StartTracking(projectTitle)
	tracking, _ := CurrentTracking()
	if tracking.Title != projectTitle {
		t.Errorf("\"%s\" should be equal to \"%s\"", tracking.Title, projectTitle)
	}
}
