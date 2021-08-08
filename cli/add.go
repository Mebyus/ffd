package cli

import (
	"fmt"

	"github.com/mebyus/ffd/cli/command"
	"github.com/mebyus/ffd/track"
)

type addExecutor struct{}

func NewAddExecutor() *addExecutor {
	return &addExecutor{}
}

func NewAddTemplate() (template *command.Template) {
	template = &command.Template{
		Name:  "add",
		Usage: "[options] <url>",
		Description: command.Description{
			Short: "add a fic to your library",
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

func (e *addExecutor) Execute(cmd *command.Command) (err error) {
	if len(cmd.Targets) == 0 {
		return fmt.Errorf("target is not specified")
	}
	trackpath := cmd.ValueFlags["track"]
	err = track.Add(cmd.Targets[0], trackpath)
	return
}
