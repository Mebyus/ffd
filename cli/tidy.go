package cli

import (
	"github.com/mebyus/ffd/cli/command"
	"github.com/mebyus/ffd/track"
)

func NewTidyTemplate() (template *command.Template) {
	template = &command.Template{
		Name: "tidy",
		BoolFlags: []command.BoolFlag{
			{
				Flag: command.Flag{
					Aliases: map[string]command.AliasType{
						"c":        command.SingleChar,
						"chapters": command.MultipleChars,
					},
				},
				Default: false,
			},
			{
				Flag: command.Flag{
					Aliases: map[string]command.AliasType{
						"u":       command.SingleChar,
						"updates": command.MultipleChars,
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

func tidy(c *Command) (err error) {
	trackpath := c.Flags["track"]
	_, cleanChapters := c.Flags["chapters"]
	_, cleanUpdates := c.Flags["updates"]
	err = track.Tidy(trackpath, cleanChapters, cleanUpdates)
	if err != nil {
		return
	}
	return
}
