package cli

import "strings"

type Command struct {
	Name   string
	Target string
	Flags  map[string]string
}

func Parse(args []string) (c *Command) {
	c = &Command{
		Name:  "help",
		Flags: make(map[string]string),
	}
	if len(args) == 0 {
		return
	}
	c.Name = args[0]
	for _, arg := range args[1:] {
		if strings.HasPrefix(arg, "--") {
			split := strings.SplitN(strings.TrimPrefix(arg, "--"), "=", 2)
			if len(split) == 0 {
				continue
			} else if len(split) == 1 {
				c.Flags[split[0]] = ""
			} else {
				c.Flags[split[0]] = split[1]
			}
		} else if strings.HasPrefix(arg, "-") {
			flag := strings.TrimPrefix(arg, "-")
			if flag == "" {
				continue
			} else {
				c.Flags[flag] = ""
			}
		} else {
			c.Target = arg
		}
	}
	return
}
