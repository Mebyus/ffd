package command

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

}

func (d *Dispatcher) displayTemplateParseError(name string, err error) {

}

func (d *Dispatcher) displayCommandExecutionError(name string, err error) {
	
}