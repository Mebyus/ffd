package cli

import "strings"

type Command struct {
	Name   string
	Target string
	Flags  map[string]string
}

func Parse(args []string) (c *Command) {
	c = &Command{
		Name:  "help", // default command
		Flags: make(map[string]string),
	}
	if len(args) == 0 {
		return
	}
	c.Name = args[0]
	for _, arg := range args[1:] {
		c.parseArg(arg)
	}
	return
}

func (c *Command) parseArg(arg string) {
	if strings.HasPrefix(arg, "--") {
		split := strings.SplitN(strings.TrimPrefix(arg, "--"), "=", 2)
		if len(split) == 0 {
			return
		} else if len(split) == 1 {
			c.Flags[split[0]] = ""
		} else {
			c.Flags[split[0]] = split[1]
		}
	} else if strings.HasPrefix(arg, "-") {
		flags := strings.TrimPrefix(arg, "-")
		if flags == "" {
			return
		}
		split := strings.Split(flags, "")
		for _, flag := range split {
			c.Flags[flag] = ""
		}
	} else {
		c.Target = arg
	}
	return
}
