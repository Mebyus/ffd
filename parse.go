package main

import (
	"fmt"

	"github.com/mebyus/ffd/cli"
	"github.com/mebyus/ffd/resource"
)

func parse(command *cli.Command) (err error) {
	if command.Target == "" {
		return fmt.Errorf("\"parse\" command: target is not specified")
	}
	hostname := command.Flags["hostname"]
	if hostname == "" {
		return fmt.Errorf("\"parse\" command: hostname is not specified")
	}
	_, separate := command.Flags["s"]
	err = resource.Parse(command.Target, hostname, separate)
	if err != nil {
		return
	}
	return nil
}
