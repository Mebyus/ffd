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

func (c *Command) setBoolDefault(flag BoolFlag) {
	if flag.Default {
		for k := range flag.Aliases {
			c.BoolFlags[k] = flag.Default
		}
	}
}

func (c *Command) setValue(flag ValueFlag, value string) {
	for k := range flag.Aliases {
		c.ValueFlags[k] = value
	}
}

func (c *Command) setValueDefault(flag ValueFlag) {
	if flag.Default != "" {
		for k := range flag.Aliases {
			c.ValueFlags[k] = flag.Default
		}
	}
}
