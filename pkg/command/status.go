package command

import "fmt"

func StatusCommand() {
	if a, b := activeTime(); a {
		fmt.Println("Current " + b)
	} else {
		fmt.Println("No project started")
	}
}

func activeTime() (bool, string) {
	return false, ""
	// return true, "Vagoru"
}
