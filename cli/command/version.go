package command

import "fmt"

func isVersionCommand(args []string) (isVersion bool) {
	isVersion = args[0] == "version" || args[0] == "-v" || args[0] == "--version"
	return
}

func (d *Dispatcher) displayVersion() {
	if d.version == "" {
		fmt.Println("version information is not available")
		return
	}
	fmt.Println(d.version)
}
