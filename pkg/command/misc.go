package command

// Version output
func Version() string {
	return "Version 1.0"
}

// No argument output
func NoArgs() string {
	return "No input, aborting.."
}

// Command not found output
func NotFound() string {
	return "Command not found"
}

// Help output
func Help() string {
	var out = ""
	out += "Usage: vagoru <command>\n\n"
	out += "Commands:\n"
	out += "help\t\tShow help\n"
	out += "version\t\tShow version\n"
	return out
}
