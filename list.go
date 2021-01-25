package main

import (
	"github.com/mebyus/ffd/cli"
	"github.com/mebyus/ffd/track"
)

func list(command *cli.Command) (err error) {
	err = track.List()
	if err != nil {
		return err
	}
	return
}
