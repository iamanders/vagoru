package command

import (
	"strings"
	"testing"
)

func TestStatusCommandNoProject(t *testing.T) {
	output := Status()
	testString := "No project started"
	if !strings.Contains(output, testString) {
		t.Errorf("\"%s\" should contain \"%s\"", output, testString)
	}
}

func TestStatusCommandProject(t *testing.T) {
	StartTracking("Project 1")
	output := Status()
	testString := "Project 1"
	if !strings.Contains(output, testString) {
		t.Errorf("\"%s\" should contain \"%s\"", output, testString)
	}
}
