package main

import (
	"fmt"
	"os"

	"github.com/mebyus/ffd/cli"
	"github.com/mebyus/ffd/resource"
	"github.com/mebyus/ffd/resource/fiction"
)

func parse(command *cli.Command) (err error) {
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
	if command.Target == "" {
		err = resource.ParseReader(os.Stdin, resourceID, format)
	} else {
		err = resource.Parse(command.Target, resourceID, separate, format)
	}
	return
}
