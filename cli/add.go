package cli

import (
	"fmt"

	"github.com/mebyus/ffd/track"
)

func add(c *Command) (err error) {
	if c.Target == "" {
		return fmt.Errorf("\"add\" command: target is not specified")
	}
	trackpath := c.Flags["track"]
	err = track.Add(c.Target, trackpath)
	if err != nil {
		return
	}
	return
}
