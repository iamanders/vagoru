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
		out = command.NoArgs()
	} else {
		// Match commands
		switch strings.ToLower(args[0]) {
		case "help":
			out = command.Help()
		case "version":
			out = command.Version()
		case "status":
			out = command.Status()
		case "start":
			out = command.StartTracking()
		default:
			out = command.NotFound()
		}
	}

	// Return or print
	if output {
		fmt.Println(out)
		return ""
	}
	return out
}
