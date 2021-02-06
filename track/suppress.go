package track

import (
	"fmt"
	"net/url"

	"github.com/mebyus/ffd/track/fic"
)

func Suppress(target, trackpath string, resume bool) (err error) {
	_, err = url.Parse(target)
	if err != nil {
		return
	}
	fics, originpath, err := fic.Load(trackpath)
	if err != nil {
		return
	}
	i := fic.Find(fics, target)
	if i == -1 {
		err = fmt.Errorf("not on the list")
		return
	}
	fics[i].Check.Suppressed = !resume
	err = fic.Save(originpath, fics)
	if err != nil {
		return
	}
	return
}
