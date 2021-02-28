package cli

import (
	"fmt"

)

func help(c *Command) (err error) {
	fmt.Println(`FanFiction Dissector
	Available commands:

	help
	add
	check
	download
	list
	parse
	suppress
	tidy`)
	return
}
