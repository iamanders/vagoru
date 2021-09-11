package terminal

// No argument output
func noArgsOutput() string {
	return "No input, aborting.."
}

// Help output
func helpOutput() string {
	var out = ""
	out += "Usage: vagoru <command>\n\n"
	out += "Commands:\n"
	out += "help\t\tShow help\n"
	out += "version\t\tShow version\n"
	return out
}

// Version output
func versionOutput() string {
	return "Version 1.0"
}

// Command not found output
func commandNotFoundOutput() string {
	return "Command not found"
}
