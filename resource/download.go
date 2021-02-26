package resource

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/mebyus/ffd/cmn"
	"github.com/mebyus/ffd/resource/fiction"
	"github.com/mebyus/ffd/setting"
	"github.com/mebyus/ffd/track/fic"
)

// Download fetches a fic from a given target.
// An appropriate target is fic page URL.
// SaveSource flag indicates whether responses will be saved before parsing.
// Format argument determines format of the resulting output file
func Download(target string, saveSource bool, format fiction.RenderFormat) (err error) {
	ficNumber, err := strconv.ParseInt(target, 10, 64)
	if err != nil {
		// target is not a number, thus treat it as URL
		err = downloadFromURL(target, saveSource, format)
	} else {
		// target is the number of a fic in the list
		err = downloadFromList(int(ficNumber), saveSource, format)
	}
	return
}

func saveHistory(target string) (err error) {
	file, err := os.OpenFile(setting.HistoryPath, os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		return
	}
	defer cmn.SmartClose(file)
	_, err = io.WriteString(file, target+"\n")
	return
}

func downloadFromURL(target string, saveSource bool, format fiction.RenderFormat) (err error) {
	t, err := ChooseByTarget(target)
	if err != nil {
		err = fmt.Errorf("choosing tool for %s: %v", target, err)
		return
	}
	err = saveHistory(target)
	if err != nil {
		fmt.Printf("Saving history: %v\n", err)
		err = nil
	}
	book, err := t.Download(target, saveSource)
	if err != nil {
		return
	}
	err = book.Save(setting.OutDir, format)
	return
}

func downloadFromList(ficNumber int, saveSource bool, format fiction.RenderFormat) (err error) {
	f, err := fic.Get(ficNumber)
	if err != nil {
		return
	}
	err = downloadFromURL(f.BaseURL, saveSource, format)
	if err != nil {
		return
	}
	return
}
