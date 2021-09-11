package terminal

import (
	"os"
	"strings"

	"github.com/iamanders/vagoru/pkg/command"
)

func ParseArgv() {
	args := os.Args[1:]

	// No args
	if len(args) < 1 {
		noArgsOutput()
		return
	}

	// Match commands
	switch strings.ToLower(args[0]) {
	case "help":
		helpOutput()
	case "version":
		versionOutput()
	case "status":
		command.StatusCommand()
	default:
		commandNotFoundOutput()
	}
}
