package main

import (
	"fmt"
	"os"

	"github.com/mebyus/ffd/cli"
	"github.com/mebyus/ffd/planner"
)

func unknown(command *cli.Command) (err error) {
	return fmt.Errorf("unknown command")
}

func main() {
	command := cli.Parse(os.Args[1:])

	var executor func(command *cli.Command) error
	switch command.Name {
	case "download":
		executor = download
	case "parse":
		executor = parse
	case "help":
		executor = help
	case "add":
		executor = add
	case "check":
		executor = check
	case "suppress":
		executor = suppress
	case "list":
		executor = list
	case "tidy":
		executor = tidy
	default:
		executor = unknown
	}
	go planner.Planner()
	err := executor(command)
	if err != nil {
		fmt.Printf("Command execution: %v\n", err)
		return
	}

	return
}
