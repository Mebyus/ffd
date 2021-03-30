package cli

import (
	"fmt"
	"strings"

	"github.com/mebyus/ffd/cli/command"
	"github.com/mebyus/ffd/resource"
	"github.com/mebyus/ffd/resource/fiction"
)

func NewDownloadTemplate() (template *command.Template) {
	template = &command.Template{
		Name: "download",
		ValueFlags: []command.ValueFlag{
			{
				Flag: command.Flag{
					Aliases: map[string]command.AliasType{
						"o":   command.SingleChar,
						"out": command.MultipleChars,
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
		BoolFlags: []command.BoolFlag{
			{
				Flag: command.Flag{
					Aliases: map[string]command.AliasType{
						"s":           command.SingleChar,
						"save-source": command.MultipleChars,
					},
				},
				Default: false,
			},
		},
	}
	return
}

func download(c *Command) (err error) {
	if c.Target == "" {
		return fmt.Errorf("\"download\" command: target is not specified")
	}
	_, save := c.Flags["s"]
	format := fiction.RenderFormat(strings.ToUpper(c.Flags["format"]))
	if format == "" {
		format = fiction.TXT
	}
	err = resource.Download(c.Target, save, format)
	if err != nil {
		return
	}
	return
}
