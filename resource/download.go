package resource

import (
	"fmt"
	"strconv"

	"github.com/mebyus/ffd/track/fic"
)

// Download fetches a fic from a given target.
// An appropriate target is fic page URL.
// SaveSource flag indicates whether responses will be saved before parsing
func Download(target string, saveSource bool) (err error) {
	ficNumber, err := strconv.ParseInt(target, 10, 64)
	if err != nil {
		// target is not a number, thus treat it as URL
		err = downloadFromURL(target, saveSource)
	} else {
		// target is the number of a fic in the list
		err = downloadFromList(int(ficNumber), saveSource)
	}
	return
}

func downloadFromURL(target string, saveSource bool) (err error) {
	t, err := ChooseByTarget(target)
	if err != nil {
		err = fmt.Errorf("choosing tool for %s: %v", target, err)
		return
	}
	t.Download(target, saveSource)
	return
}

func downloadFromList(ficNumber int, saveSource bool) (err error) {
	f, err := fic.Get(ficNumber)
	if err != nil {
		return
	}
	err = downloadFromURL(f.BaseURL, saveSource)
	if err != nil {
		return
	}
	return
}
