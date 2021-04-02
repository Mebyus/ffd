package cli

import (
	"fmt"

	"github.com/mebyus/ffd/cli/command"
	"github.com/mebyus/ffd/track"
)

type suppressExecutor struct{}

func NewSuppressExecutor() *suppressExecutor {
	return &suppressExecutor{}
}

func NewSuppressTemplate() (template *command.Template) {
	template = &command.Template{
		Name: "suppress",
		BoolFlags: []command.BoolFlag{
			{
				Flag: command.Flag{
					Aliases: map[string]command.AliasType{
						"r":      command.SingleChar,
						"resume": command.MultipleChars,
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

func (e *suppressExecutor) Execute(cmd *command.Command) (err error) {
	if len(cmd.Targets) == 0 {
		return fmt.Errorf("target is not specified")
	}
	resume := cmd.BoolFlags["resume"]
	trackpath := cmd.ValueFlags["track"]
	err = track.Suppress(cmd.Targets[0], trackpath, resume)
	return
}
