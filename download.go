package main

import (
	"fmt"

	"github.com/mebyus/ffd/cli"
	"github.com/mebyus/ffd/resource"
)

func download(command *cli.Command) (err error) {
	if command.Target == "" {
		return fmt.Errorf("\"download\" command: target is not specified")
	}
	_, save := command.Flags["s"]
	err = resource.Download(command.Target, save)
	if err != nil {
		return
	}
	return
}
