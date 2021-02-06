package track

import (
	"fmt"
	"net/url"

	"github.com/mebyus/ffd/track/fic"
)

func Add(target, trackpath string) (err error) {
	_, err = url.Parse(target)
	if err != nil {
		return
	}
	fics, originpath, err := fic.Load(trackpath)
	if err != nil {
		return
	}
	i := fic.Find(fics, target)
	if i != -1 {
		err = fmt.Errorf("already in the list")
		return
	}
	fics = append(fics, fic.Info{
		BaseURL: target,
	})
	err = fic.Save(originpath, fics)
	if err != nil {
		return
	}
	return
}
