package track

import (
	"fmt"
	"time"

	"github.com/mebyus/ffd/resource"
	"github.com/mebyus/ffd/track/fic"
)

func Add(target, trackpath string) (err error) {
	location, err := resource.GetLocationForTarget(target)
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
	newfic, err := resource.Check(target)
	if err != nil {
		return
	}
	newfic.Location = location
	newfic.Words = fic.CountWords(newfic.Chapters)
	newfic.Updated = fic.UpdatedTime(newfic.Chapters)
	newfic.Check = fic.Check{
		Time: time.Now(),
	}
	fics = append(fics, *newfic)
	err = fic.Save(originpath, fics)
	if err != nil {
		return
	}
	fmt.Printf("Fic added under index %d\n", len(fics))
	return
}
