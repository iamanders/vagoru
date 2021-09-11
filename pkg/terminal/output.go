package terminal

import "fmt"

// No argument output
func noArgsOutput() {
	fmt.Println("No input, aborting..")
}

// Help output
func helpOutput() {
	fmt.Println("Usage: vagoru <command>")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("help		Show help")
	fmt.Println("version		Show version")
}

// Version output
func versionOutput() {
	fmt.Println("Version 1.0")
}

// Command not found output
func commandNotFoundOutput() {
	fmt.Println("Command not found")
}
