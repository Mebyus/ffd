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

func (d *Dispatcher) displayHelp(args []string) {

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
