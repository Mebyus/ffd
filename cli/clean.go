package cli

import (
	"github.com/mebyus/ffd/cli/command"
	"github.com/mebyus/ffd/resource"
)

type cleanExecutor struct{}

func NewCleanExecutor() *cleanExecutor {
	return &cleanExecutor{}
}

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

func (e *cleanExecutor) Execute(cmd *command.Command) (err error) {
	cleanHistory := cmd.BoolFlags["history"]
	cleanSource := cmd.BoolFlags["source"]
	err = resource.Clean(cleanHistory, cleanSource)
	return
}
