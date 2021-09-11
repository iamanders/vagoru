package terminal

import (
	"os"
	"strings"
	"testing"
)

func TestArgv(t *testing.T) {
	// Preserve old args
	originalArgs := os.Args
	defer func() {
		os.Args = originalArgs
	}()

	// No args
	os.Args = []string{"vagoru"}
	out := ParseArgv(false)
	if !strings.Contains(out, "No input") {
		t.Error("No input args does not work as it should")
	}

	// Help
	os.Args = []string{"vagoru", "help"}
	out = ParseArgv(false)
	if !strings.Contains(out, "Usage: vagoru") {
		t.Error("Help command does not return help")
	}

	// Version
	os.Args = []string{"vagoru", "version"}
	out = ParseArgv(false)
	if !strings.Contains(out, "Version") {
		t.Error("Version command does not return version")
	}

}
