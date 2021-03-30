package cli

import (
	"fmt"

	"github.com/mebyus/ffd/cli/command"
	"github.com/mebyus/ffd/track"
)

func NewAddTemplate() (template *command.Template) {
	template = &command.Template{
		Name: "add",
		ValueFlags: []command.ValueFlag{
			{
				Flag: command.Flag{
					Aliases: map[string]command.AliasType{
						"t":     command.SingleChar,
						"track": command.MultipleChars,
					},
				},
				Default: "",
			},
		},
	}
	return
}

func add(c *Command) (err error) {
	if c.Target == "" {
		return fmt.Errorf("\"add\" command: target is not specified")
	}
	trackpath := c.Flags["track"]
	err = track.Add(c.Target, trackpath)
	if err != nil {
		return
	}
	return
}
