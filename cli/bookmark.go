package cli

import (
	"github.com/mebyus/ffd/cli/command"
	"github.com/mebyus/ffd/track"
)

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

func bookmark(c *Command) error {
	trackpath := c.Flags["track"]
	chapter, ok := c.Flags["chapter"]
	if !ok {
		chapter = "latest"
	}
	err := track.Bookmark(trackpath, c.Target, chapter)
	if err != nil {
		return err
	}
	return nil
}
