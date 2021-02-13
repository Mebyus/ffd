package main

import (
	"fmt"
	"os"

	"github.com/mebyus/ffd/cli"
	"github.com/mebyus/ffd/logs"
	"github.com/mebyus/ffd/planner"
	"github.com/mebyus/ffd/setting"
)

func unknown(command *cli.Command) (err error) {
	return fmt.Errorf("unknown command")
}

func main() {
	command := cli.Parse(os.Args[1:])

	_, verbose := command.Flags["v"]
	if verbose {
		logs.Init(logs.I)
	}

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
	case "remove":
		executor = remove
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

	setting.Load()
	go planner.Planner()
	err := executor(command)
	if err != nil {
		logs.Error.Printf("command execution: %v\n", err)
		return
	}

	return
}
