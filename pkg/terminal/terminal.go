package terminal

import (
	"fmt"
	"os"
	"strings"

	"github.com/iamanders/vagoru/pkg/command"
)

func ParseArgv(output bool) string {
	args := os.Args[1:]
	var out string

	if len(args) < 1 {
		// No args
		out = noArgsOutput()
	} else {
		// Match commands
		switch strings.ToLower(args[0]) {
		case "help":
			out = helpOutput()
		case "version":
			out = versionOutput()
		case "status":
			out = command.StatusCommand()
		default:
			out = commandNotFoundOutput()
		}
	}

	// Return or print
	if output {
		fmt.Println(out)
		return ""
	}
	return out
}
