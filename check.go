package main

import (
	"github.com/mebyus/ffd/cli"
	"github.com/mebyus/ffd/track"
)

func check(command *cli.Command) error {
	err := track.Check()
	if err != nil {
		return err
	}
	return nil
}
