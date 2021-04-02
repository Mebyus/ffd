package cli

import (
	"github.com/mebyus/ffd/cli/command"
	"github.com/mebyus/ffd/track"
)

type tidyExecutor struct{}

func NewTidyExecutor() *tidyExecutor {
	return &tidyExecutor{}
}

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

func (e *tidyExecutor) Execute(cmd *command.Command) (err error) {
	trackpath := cmd.ValueFlags["track"]
	cleanChapters := cmd.BoolFlags["chapters"]
	cleanUpdates := cmd.BoolFlags["updates"]
	err = track.Tidy(trackpath, cleanChapters, cleanUpdates)
	return
}
