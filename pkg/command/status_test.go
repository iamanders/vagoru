package command

import (
	"strings"
	"testing"
)

func TestStatusCommandNoProject(t *testing.T) {
	output := Status()
	testString := "No project started"
	if !strings.Contains(output, testString) {
		t.Errorf("\"%s\" should be \"%s\"", output, testString)
	}
}

func TestStatusCommandProject(t *testing.T) {
	// output := Status()
	// testString := "No project started"
	// if !strings.Contains(output, testString) {
	// 	t.Errorf("\"%s\" should be \"%s\"", output, testString)
	// }
}
