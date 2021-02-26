package resource

import (
	"os"

	"github.com/mebyus/ffd/setting"
)

func Clean(cleanHistory, cleanSource bool) (err error) {
	if cleanHistory {
		err = os.Remove(setting.HistoryPath)
		if err != nil {
			return
		}
	}
	if cleanSource {
		err = os.RemoveAll(setting.SourceSaveDir)
		if err != nil {
			return
		}
	}
	err = os.RemoveAll(setting.OutDir)
	return
}
