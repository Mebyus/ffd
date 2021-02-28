package cli

import (
	"fmt"
	"strings"

	"github.com/mebyus/ffd/resource"
	"github.com/mebyus/ffd/resource/fiction"
)

func download(c *Command) (err error) {
	if c.Target == "" {
		return fmt.Errorf("\"download\" command: target is not specified")
	}
	_, save := c.Flags["s"]
	format := fiction.RenderFormat(strings.ToUpper(c.Flags["format"]))
	if format == "" {
		format = fiction.TXT
	}
	err = resource.Download(c.Target, save, format)
	if err != nil {
		return
	}
	return
}
