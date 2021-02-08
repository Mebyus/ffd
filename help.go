package main

import (
	"fmt"

	"github.com/mebyus/ffd/cli"
)

func help(command *cli.Command) (err error) {
	fmt.Println(`FanFiction Dissector
	Available commands:

	help
	add
	check
	download
	list
	parse
	suppress`)
	return
}
