package terminal

import "fmt"

// No argument output
func noArgsOutput() {
	fmt.Println("No input, aborting..")
}

// Help output
func helpOutput() {
	fmt.Println("Help!")
}

// Version output
func versionOutput() {
	fmt.Println("Version 1.0")
}
