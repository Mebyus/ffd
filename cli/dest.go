package cli

import (
	"github.com/mebyus/ffd/cli/command"
	"github.com/mebyus/ffd/open"
	"github.com/mebyus/ffd/setting"
)

func NewDestTemplate() (template *command.Template) {
	template = &command.Template{
		Name: "dest",
	}
	return
}

func dest(c *Command) (err error) {
	downloadDir := setting.OutDir
	err = open.Start(downloadDir)
	if err != nil {
		return
	}
	return
}
