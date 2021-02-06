package main

import (
	"github.com/mebyus/ffd/cli"
	"github.com/mebyus/ffd/track"
)

func list(command *cli.Command) (err error) {
	trackpath := command.Flags["track"]
	err = track.List(trackpath)
	if err != nil {
		return err
	}
	return
}
