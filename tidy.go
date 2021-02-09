package main

import (
	"github.com/mebyus/ffd/cli"
	"github.com/mebyus/ffd/track"
)

func tidy(command *cli.Command) (err error) {
	trackpath := command.Flags["track"]
	err = track.Tidy(trackpath)
	if err != nil {
		return
	}
	return
}
