package command

type Command struct {
	Name       string
	Targets    []string
	BoolFlags  map[string]bool
	ValueFlags map[string]string
}

func NewCommand(name string) *Command {
	return &Command{
		Name:       name,
		Targets:    []string{},
		BoolFlags:  map[string]bool{},
		ValueFlags: map[string]string{},
	}
}

func (c *Command) setBool(flag BoolFlag) {
	for k := range flag.Aliases {
		c.BoolFlags[k] = !flag.Default
	}
}
