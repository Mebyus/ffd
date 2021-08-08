package command

import "fmt"

func isHelpCommand(args []string) (isHelp bool, helpArgs []string) {
	if len(args) == 0 {
		isHelp = true
		return
	}
	if args[0] == "help" || args[0] == "-h" || args[0] == "--help" {
		isHelp = true
		if len(args) > 1 {
			helpArgs = append(helpArgs, args[1])
		}
	}
	return
}

func (d *Dispatcher) displayGeneralHelp() {
	fmt.Printf("%s\n\n", d.description)
	fmt.Printf("List of recognized commands:\n\n")
	for _, pair := range d.pairs {
		fmt.Printf("    %-16s %s\n", pair.template.Name, pair.template.Description.Short)
	}
	fmt.Printf("    %-16s %s\n", "version", "print version and exit")
	fmt.Printf("    %-16s %s\n\n", "help", "print this message and exit")
	fmt.Printf("Use \"%s help <command>\" for more information about a command.\n", d.name)
}

func (d *Dispatcher) displayCommandHelp(name string) {
	pair, ok := d.pairs[name]
	if !ok {
		d.displayGeneralHelp()
		return
	}
	pair.template.displayHelp(d.name)
}

func (d *Dispatcher) displayHelp(args []string) {
	if len(args) == 0 {
		d.displayGeneralHelp()
		return
	}
	d.displayCommandHelp(args[0])
}

func (d *Dispatcher) displayCommandNotFound(name string) {
	fmt.Printf("command [ %s ] not found\n", name)
}

func (d *Dispatcher) displayTemplateParseError(name string, err error) {
	fmt.Printf("error parsing arguments for command [ %s ]: %v\n", name, err)
}

func (d *Dispatcher) displayCommandExecutionError(name string, err error) {
	fmt.Printf("error executing command [ %s ]: %v\n", name, err)
}
