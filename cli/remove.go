package cli

import (
	"fmt"

	"github.com/mebyus/ffd/cli/command"
	"github.com/mebyus/ffd/track"
)

type removeExecutor struct{}

func NewRemoveExecutor() *removeExecutor {
	return &removeExecutor{}
}

func NewRemoveTemplate() (template *command.Template) {
	template = &command.Template{
		Name: "remove",
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

func (e *removeExecutor) Execute(cmd *command.Command) (err error) {
	if len(cmd.Targets) == 0 {
		return fmt.Errorf("target is not specified")
	}
	trackpath := cmd.ValueFlags["track"]
	err = track.Remove(cmd.Targets[0], trackpath)
	return
}
