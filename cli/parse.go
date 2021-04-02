package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/mebyus/ffd/cli/command"
	"github.com/mebyus/ffd/resource"
	"github.com/mebyus/ffd/resource/fiction"
)

type parseExecutor struct{}

func NewParseExecutor() *parseExecutor {
	return &parseExecutor{}
}

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

func (e *parseExecutor) Execute(cmd *command.Command) (err error) {
	resourceID := cmd.ValueFlags["resource"]
	if resourceID == "" {
		return fmt.Errorf("resource is not specified")
	}
	format := fiction.RenderFormat(strings.ToUpper(cmd.ValueFlags["format"]))
	separate := cmd.BoolFlags["separate"]
	if len(cmd.Targets) == 0 {
		err = resource.ParseReader(os.Stdin, resourceID, format)
	} else {
		err = resource.Parse(cmd.Targets[0], resourceID, separate, format)
	}
	return
}
