package cli

import (
	"fmt"

	"github.com/mebyus/ffd/cli/command"
	"github.com/mebyus/ffd/track"
)

func NewSuppressTemplate() (template *command.Template) {
	template = &command.Template{
		Name: "suppress",
		BoolFlags: []command.BoolFlag{
			{
				Flag: command.Flag{
					Aliases: map[string]command.AliasType{
						"r":      command.SingleChar,
						"resume": command.MultipleChars,
					},
				},
				Default: false,
			},
		},
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

func suppress(c *Command) (err error) {
	if c.Target == "" {
		return fmt.Errorf("\"suppress\" command: target is not specified")
	}
	_, resume := c.Flags["r"]
	trackpath := c.Flags["track"]
	err = track.Suppress(c.Target, trackpath, resume)
	if err != nil {
		return
	}
	return
}
