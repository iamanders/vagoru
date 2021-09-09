package terminal

import (
	"os"
	"strings"
)

func ParseArgv() {
	args := os.Args[1:]

	// No args
	if len(args) < 1 {
		noArgsOutput()
		return
	}

	// Help
	if strings.EqualFold("help", args[0]) {
		helpOutput()
		return
	}

	// Version
	if strings.EqualFold("version", args[0]) {
		versionOutput()
		return
	}

	// Command not found
	// TODO
}
