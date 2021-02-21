package main

import (
	"fmt"

	"github.com/mebyus/ffd/cli"
	"github.com/mebyus/ffd/track"
)

func suppress(command *cli.Command) (err error) {
	if command.Target == "" {
		return fmt.Errorf("\"suppress\" command: target is not specified")
	}
	_, resume := command.Flags["r"]
	trackpath := command.Flags["track"]
	err = track.Suppress(command.Target, trackpath, resume)
	if err != nil {
		return
	}
	return
}
