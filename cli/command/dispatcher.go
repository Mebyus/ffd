package command

import (
	"fmt"
	"os"
)

type Executor interface {
	Execute(cmd *Command) (err error)
}

type pair struct {
	template *Template
	executor Executor
}

type Dispatcher struct {
	version string
	pairs   map[string]*pair
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		pairs: make(map[string]*pair),
	}
}

func (d *Dispatcher) SetVersion(version fmt.Stringer) {
	d.version = version.String()
}

func (d *Dispatcher) Register(template *Template, executor Executor) {
	err := template.prepare()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_, ok := d.pairs[template.Name]
	if ok {
		fmt.Printf("template [ %s ] already registered\n", template.Name)
		os.Exit(1)
	}
	d.pairs[template.Name] = &pair{
		template: template,
		executor: executor,
	}
}

func (d *Dispatcher) Dispatch(args []string) (err error) {
	isHelp, helpArgs := isHelpCommand(args)
	if isHelp {
		d.displayHelp(helpArgs)
		return
	}
	if isVersionCommand(args) {
		d.displayVersion()
		return
	}
	pair, ok := d.pairs[args[0]]
	if !ok {
		d.displayCommandNotFound(args[0])
		return
	}
	command, err := pair.template.Parse(args[1:])
	if err != nil {
		d.displayTemplateParseError(pair.template.Name, err)
		return
	}
	err = pair.executor.Execute(command)
	if err != nil {
		d.displayCommandExecutionError(command.Name, err)
	}
	return
}

func (d *Dispatcher) SetStringVersion(version string) {
	d.version = version
}
