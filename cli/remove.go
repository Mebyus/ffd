package cli

import (
	"fmt"

	"github.com/mebyus/ffd/track"
)

func remove(c *Command) (err error) {
	if c.Target == "" {
		return fmt.Errorf("\"remove\" command: target is not specified")
	}
	trackpath := c.Flags["track"]
	err = track.Remove(c.Target, trackpath)
	if err != nil {
		return
	}
	return
}
