package cli

import (
	"github.com/mebyus/ffd/track"
)

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
