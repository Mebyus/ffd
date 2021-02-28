package cli

import (
	"github.com/mebyus/ffd/resource"
)

func clean(c *Command) (err error) {
	_, cleanHistory := c.Flags["h"]
	_, cleanSource := c.Flags["s"]
	err = resource.Clean(cleanHistory, cleanSource)
	if err != nil {
		return
	}
	return
}
