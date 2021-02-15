package main

import (
	"fmt"

	"github.com/mebyus/ffd/cli"
	"github.com/mebyus/ffd/resource"
	"github.com/mebyus/ffd/resource/fiction"
)

func download(command *cli.Command) (err error) {
	if command.Target == "" {
		return fmt.Errorf("\"download\" command: target is not specified")
	}
	_, save := command.Flags["s"]
	format := fiction.RenderFormat(command.Flags["format"])
	if format == "" {
		format = fiction.TXT
	}
	err = resource.Download(command.Target, save, format)
	if err != nil {
		return
	}
	return
}
