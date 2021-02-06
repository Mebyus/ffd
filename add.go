package main

import (
	"fmt"

	"github.com/mebyus/ffd/cli"
	"github.com/mebyus/ffd/track"
)

func add(command *cli.Command) (err error) {
	if command.Target == "" {
		return fmt.Errorf("\"add\" command: target is not specified")
	}
	trackpath := command.Flags["track"]
	err = track.Add(command.Target, trackpath)
	if err != nil {
		return
	}
	return
}
