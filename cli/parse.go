package cli

import (
	"fmt"
	"os"

	"github.com/mebyus/ffd/resource"
	"github.com/mebyus/ffd/resource/fiction"
)

func parse(c *Command) (err error) {
	resourceID := c.Flags["resource"]
	if resourceID == "" {
		resourceID = c.Flags["res"]
	}
	if resourceID == "" {
		return fmt.Errorf("\"parse\" command: resource is not specified")
	}
	format := fiction.RenderFormat(c.Flags["format"])
	if format == "" {
		format = fiction.TXT
	}
	_, separate := c.Flags["s"]
	if c.Target == "" {
		err = resource.ParseReader(os.Stdin, resourceID, format)
	} else {
		err = resource.Parse(c.Target, resourceID, separate, format)
	}
	return
}
