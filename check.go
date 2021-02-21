package main

import (
	"github.com/mebyus/ffd/cli"
	"github.com/mebyus/ffd/track"
)

func check(command *cli.Command) error {
	trackpath := command.Flags["track"]
	err := track.Check(trackpath, command.Target)
	if err != nil {
		return err
	}
	return nil
}
