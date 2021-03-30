package cli

import (
	"fmt"

	"github.com/mebyus/ffd/cli/command"
	"github.com/mebyus/ffd/track"
)

func NewRemoveTemplate() (template *command.Template) {
	template = &command.Template{
		Name: "remove",
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

func remove(c *Command) (err error) {
	if c.Target == "" {
		return fmt.Errorf("\"remove\" command: target is not specified")
	}
	trackpath := c.Flags["track"]
	err = track.Remove(c.Target, trackpath)
	if err != nil {
		return
	}
	return
}
