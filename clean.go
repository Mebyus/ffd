package main

import (
	"github.com/mebyus/ffd/cli"
	"github.com/mebyus/ffd/resource"
)

func clean(command *cli.Command) (err error) {
	_, cleanHistory := command.Flags["h"]
	_, cleanSource := command.Flags["s"]
	err = resource.Clean(cleanHistory, cleanSource)
	if err != nil {
		return
	}
	return
}
