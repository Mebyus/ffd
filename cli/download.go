package cli

import (
	"fmt"
	"strings"

	"github.com/mebyus/ffd/cli/command"
	"github.com/mebyus/ffd/resource"
	"github.com/mebyus/ffd/resource/fiction"
)

type downloadExecutor struct{}

func NewDownloadExecutor() *downloadExecutor {
	return &downloadExecutor{}
}

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

func (e *downloadExecutor) Execute(cmd *command.Command) (err error) {
	if len(cmd.Targets) == 0 {
		return fmt.Errorf("target is not specified")
	}
	saveSource := cmd.BoolFlags["save-source"]
	format := fiction.RenderFormat(strings.ToUpper(cmd.ValueFlags["format"]))
	err = resource.Download(cmd.Targets[0], saveSource, format)
	return
}
