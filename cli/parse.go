package cli

import (
	"fmt"
	"os"

	"github.com/mebyus/ffd/cli/command"
	"github.com/mebyus/ffd/resource"
	"github.com/mebyus/ffd/resource/fiction"
)

func NewParseTemplate() (template *command.Template) {
	template = &command.Template{
		Name: "parse",
		BoolFlags: []command.BoolFlag{
			{
				Flag: command.Flag{
					Aliases: map[string]command.AliasType{
						"s":        command.SingleChar,
						"separate": command.MultipleChars,
					},
				},
				Default: false,
			},
		},
		ValueFlags: []command.ValueFlag{
			{
				Flag: command.Flag{
					Aliases: map[string]command.AliasType{
						"r":        command.SingleChar,
						"res":      command.MultipleChars,
						"resource": command.MultipleChars,
					},
				},
				Default: "",
			},
			{
				Flag: command.Flag{
					Aliases: map[string]command.AliasType{
						"f":      command.SingleChar,
						"format": command.MultipleChars,
					},
				},
				Default: "txt",
			},
		},
	}
	return
}

func parse(c *Command) (err error) {
	resourceID := c.Flags["resource"]
	if resourceID == "" {
		resourceID = c.Flags["res"]
	}
	if resourceID == "" {
		return fmt.Errorf("\"parse\" command: resource is not specified")
	}
	format := fiction.RenderFormat(c.Flags["format"])
	if format == "" {
		format = fiction.TXT
	}
	_, separate := c.Flags["s"]
	if c.Target == "" {
		err = resource.ParseReader(os.Stdin, resourceID, format)
	} else {
		err = resource.Parse(c.Target, resourceID, separate, format)
	}
	return
}
