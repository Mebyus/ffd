package cli

import (
	"github.com/mebyus/ffd/cli/command"
	"github.com/mebyus/ffd/track"
)

func NewCheckTemplate() (template *command.Template) {
	template = &command.Template{
		Name: "check",
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

func check(c *Command) error {
	trackpath := c.Flags["track"]
	err := track.Check(trackpath, c.Target)
	if err != nil {
		return err
	}
	return nil
}
