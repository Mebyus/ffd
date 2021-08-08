package cli

import (
	"github.com/mebyus/ffd/cli/command"
	"github.com/mebyus/ffd/track"
)

type listExecutor struct{}

func NewListExecutor() *listExecutor {
	return &listExecutor{}
}

func NewListTemplate() (template *command.Template) {
	template = &command.Template{
		Name: "list",
		Description: command.Description{
			Short: "print list of fics in your library",
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

func (e *listExecutor) Execute(cmd *command.Command) (err error) {
	target := ""
	if len(cmd.Targets) > 0 {
		target = cmd.Targets[0]
	}
	trackpath := cmd.ValueFlags["track"]
	err = track.List(trackpath, target)
	return
}
