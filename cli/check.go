package cli

import (
	"github.com/mebyus/ffd/track"
)

func check(c *Command) error {
	trackpath := c.Flags["track"]
	err := track.Check(trackpath, c.Target)
	if err != nil {
		return err
	}
	return nil
}
