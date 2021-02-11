package main

import (
	"github.com/mebyus/ffd/cli"
	"github.com/mebyus/ffd/track"
)

func tidy(command *cli.Command) (err error) {
	trackpath := command.Flags["track"]
	_, cleanChapters := command.Flags["chapters"]
	_, cleanUpdates := command.Flags["updates"]
	err = track.Tidy(trackpath, cleanChapters, cleanUpdates)
	if err != nil {
		return
	}
	return
}
