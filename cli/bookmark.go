package cli

import (
	"fmt"

	"github.com/mebyus/ffd/cli/command"
	"github.com/mebyus/ffd/track"
)

type bookmarkExecutor struct{}

func NewBookmarkExecutor() *bookmarkExecutor {
	return &bookmarkExecutor{}
}

func NewBookmarkTemplate() (template *command.Template) {
	template = &command.Template{
		Name: "bookmark",
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
			{
				Flag: command.Flag{
					Aliases: map[string]command.AliasType{
						"c":       command.SingleChar,
						"chapter": command.MultipleChars,
					},
				},
				Default: "latest",
			},
		},
	}
	return
}

func (e *bookmarkExecutor) Execute(cmd *command.Command) (err error) {
	if len(cmd.Targets) == 0 {
		return fmt.Errorf("target is not specified")
	}
	trackpath := cmd.ValueFlags["track"]
	chapter := cmd.ValueFlags["chapter"]
	err = track.Bookmark(trackpath, cmd.Targets[0], chapter)
	return
}
