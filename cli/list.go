package cli

import (
	"github.com/mebyus/ffd/track"
)

func list(c *Command) (err error) {
	trackpath := c.Flags["track"]
	err = track.List(trackpath)
	if err != nil {
		return err
	}
	return
}