package cli

import (
	"github.com/mebyus/ffd/cli/command"
	"github.com/mebyus/ffd/track"
)

func NewListTemplate() (template *command.Template) {
	template = &command.Template{
		Name: "list",
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

func list(c *Command) (err error) {
	trackpath := c.Flags["track"]
	err = track.List(trackpath, c.Target)
	if err != nil {
		return err
	}
	return
}
