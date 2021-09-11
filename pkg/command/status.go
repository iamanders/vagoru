package command

func StatusCommand() string {
	if a, b := activeTime(); a {
		return "Current " + b
	} else {
		return "No project started"
	}
}

func activeTime() (bool, string) {
	return false, ""
	// return true, "Vagoru"
}
