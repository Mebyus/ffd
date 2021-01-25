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
	err = track.Add(command.Target)
	if err != nil {
		return
	}
	return
}
