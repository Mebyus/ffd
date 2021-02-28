package cli

import (
	"github.com/mebyus/ffd/track"
)

func tidy(c *Command) (err error) {
	trackpath := c.Flags["track"]
	_, cleanChapters := c.Flags["chapters"]
	_, cleanUpdates := c.Flags["updates"]
	err = track.Tidy(trackpath, cleanChapters, cleanUpdates)
	if err != nil {
		return
	}
	return
}
