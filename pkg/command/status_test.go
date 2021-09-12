package command

import (
	"strings"
	"testing"

	"github.com/iamanders/vagoru/pkg/database"
)

func TestStatusCommandNoProject(t *testing.T) {
	database.ResetDb()
	output := Status()
	testString := "No project started"
	if !strings.Contains(output, testString) {
		t.Errorf("\"%s\" should contain \"%s\"", output, testString)
	}
}

func TestStatusCommandProject(t *testing.T) {
	database.ResetDb()
	database.StartTracking("Project 1")
	output := Status()
	testString := "Project 1"
	if !strings.Contains(output, testString) {
		t.Errorf("\"%s\" should contain \"%s\"", output, testString)
	}
}
