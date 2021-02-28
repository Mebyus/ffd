package main

import (
	"fmt"
	"os"

	"github.com/mebyus/ffd/cli"
	"github.com/mebyus/ffd/planner"
	"github.com/mebyus/ffd/setting"
)

func main() {
	if len(os.Args) == 0 {
		fmt.Println("no args to process")
		return
	}
	command := cli.Parse(os.Args[1:])

	executor, err := cli.CreateExecutor(command)
	if err != nil {
		fmt.Println(err)
		return
	}

	setting.Load()
	go planner.Planner()
	err = executor.Execute()
	if err != nil {
		fmt.Printf("Command execution: %v\n", err)
		return
	}

	return
}
