package track

import (
	"fmt"
	"strconv"

	"github.com/mebyus/ffd/track/fic"
)

func Remove(target, trackpath string) (err error) {
	fics, originpath, err := fic.Load(trackpath)
	if err != nil {
		return
	}
	ficNumber, err := strconv.ParseInt(target, 10, 64)
	if err != nil {
		err = fmt.Errorf("incorrect fic number: %v", err)
		return
	}
	err = fic.Remove(&fics, int(ficNumber))
	if err != nil {
		return
	}
	err = fic.Save(originpath, fics)
	if err != nil {
		return
	}
	return
}
