package main

import (
	"fmt"

	"github.com/mebyus/ffd/cli"
	"github.com/mebyus/ffd/track"
)

func remove(command *cli.Command) (err error) {
	if command.Target == "" {
		return fmt.Errorf("\"remove\" command: target is not specified")
	}
	trackpath := command.Flags["track"]
	err = track.Remove(command.Target, trackpath)
	if err != nil {
		return
	}
	return
}
