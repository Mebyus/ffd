package main

import (
	"fmt"
	"os"

	"github.com/mebyus/ffd/cli"
	"github.com/mebyus/ffd/cli/command"
	"github.com/mebyus/ffd/setting"
)

func main() {
	if len(os.Args) == 0 {
		fmt.Println("no args to process")
		return
	}
	setting.Load()
	command.Register(cli.NewAddTemplate(), cli.NewAddExecutor())
	command.Register(cli.NewBookmarkTemplate(), cli.NewBookmarkExecutor())
	command.Register(cli.NewCheckTemplate(), cli.NewCheckExecutor())
	command.Register(cli.NewCleanTemplate(), cli.NewCleanExecutor())
	command.Register(cli.NewDestTemplate(), cli.NewDestExecutor())
	command.Register(cli.NewDownloadTemplate(), cli.NewDownloadExecutor())
	command.Register(cli.NewListTemplate(), cli.NewListExecutor())
	command.Register(cli.NewParseTemplate(), cli.NewParseExecutor())
	command.Register(cli.NewRemoveTemplate(), cli.NewRemoveExecutor())
	command.Register(cli.NewSuppressTemplate(), cli.NewSuppressExecutor())
	command.Register(cli.NewTidyTemplate(), cli.NewTidyExecutor())
	err := command.Dispatch(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		return
	}
}
