package cli

import (
	"github.com/mebyus/ffd/open"
	"github.com/mebyus/ffd/setting"
)

func dest(c *Command) (err error) {
	downloadDir := setting.OutDir
	err = open.Start(downloadDir)
	if err != nil {
		return
	}
	return
}
