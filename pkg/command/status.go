package command

func Status() string {
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
