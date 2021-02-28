package cli

import (
	"fmt"

	"github.com/mebyus/ffd/track"
)

func suppress(c *Command) (err error) {
	if c.Target == "" {
		return fmt.Errorf("\"suppress\" command: target is not specified")
	}
	_, resume := c.Flags["r"]
	trackpath := c.Flags["track"]
	err = track.Suppress(c.Target, trackpath, resume)
	if err != nil {
		return
	}
	return
}
