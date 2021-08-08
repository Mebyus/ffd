package cli

import (
	"github.com/mebyus/ffd/cli/command"
	"github.com/mebyus/ffd/track"
)

type checkExecutor struct{}

func NewCheckExecutor() *checkExecutor {
	return &checkExecutor{}
}

func NewCheckTemplate() (template *command.Template) {
	template = &command.Template{
		Name: "check",
		Description: command.Description{
			Short: "check for fic updates in your library",
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

func (e *checkExecutor) Execute(cmd *command.Command) (err error) {
	target := ""
	if len(cmd.Targets) > 0 {
		target = cmd.Targets[0]
	}
	trackpath := cmd.ValueFlags["track"]
	err = track.Check(trackpath, target)
	return
}
