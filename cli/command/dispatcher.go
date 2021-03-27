package command

type Executor interface {
	Execute(cmd *Command) (err error)
}

type pair struct {
	template *Template
	executor Executor
}

type Dispatcher struct {
	pairs map[string]*pair
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		pairs: make(map[string]*pair),
	}
}

func (d *Dispatcher) Register(template *Template, executor Executor) {
	d.pairs[template.Name] = &pair{
		template: template,
		executor: executor,
	}
}

func (d *Dispatcher) Dispatch(args []string) (err error) {
	pair := d.pairs[args[0]]
	command := pair.template.Parse(args)
	err = pair.executor.Execute(command)
	return
}

func (d *Dispatcher) Describe(name string) (description string) {
	return
}
