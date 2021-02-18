package main

import (
	"fmt"

	"github.com/mebyus/ffd/cli"
	"github.com/mebyus/ffd/resource"
	"github.com/mebyus/ffd/resource/fiction"
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
	format := fiction.RenderFormat(command.Flags["format"])
	if format == "" {
		format = fiction.TXT
	}
	_, separate := command.Flags["s"]
	err = resource.Parse(command.Target, resourceID, separate, format)
	if err != nil {
		return
	}
	return nil
}
