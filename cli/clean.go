package cli

import (
	"github.com/mebyus/ffd/cli/command"
	"github.com/mebyus/ffd/resource"
)

func NewCleanTemplate() (template *command.Template) {
	template = &command.Template{
		Name: "clean",
		BoolFlags: []command.BoolFlag{
			{
				Flag: command.Flag{
					Aliases: map[string]command.AliasType{
						"s":      command.SingleChar,
						"source": command.MultipleChars,
					},
				},
				Default: false,
			},
			{
				Flag: command.Flag{
					Aliases: map[string]command.AliasType{
						"h":       command.SingleChar,
						"history": command.MultipleChars,
					},
				},
				Default: false,
			},
		},
	}
	return
}

func clean(c *Command) (err error) {
	_, cleanHistory := c.Flags["h"]
	_, cleanSource := c.Flags["s"]
	err = resource.Clean(cleanHistory, cleanSource)
	if err != nil {
		return
	}
	return
}
