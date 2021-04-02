package cli

import (
	"github.com/mebyus/ffd/cli/command"
	"github.com/mebyus/ffd/open"
	"github.com/mebyus/ffd/setting"
)

type destExecutor struct{}

func NewDestExecutor() *destExecutor {
	return &destExecutor{}
}

func NewDestTemplate() (template *command.Template) {
	template = &command.Template{
		Name: "dest",
	}
	return
}

func (e *destExecutor) Execute(cmd *command.Command) (err error) {
	downloadDir := setting.OutDir
	err = open.Start(downloadDir)
	return
}
