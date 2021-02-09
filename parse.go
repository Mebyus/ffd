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
	resourceID := command.Flags["resource"]
	if resourceID == "" {
		resourceID = command.Flags["res"]
	}
	if resourceID == "" {
		return fmt.Errorf("\"parse\" command: resource is not specified")
	}
	_, separate := command.Flags["s"]
	err = resource.Parse(command.Target, resourceID, separate)
	if err != nil {
		return
	}
	return nil
}
