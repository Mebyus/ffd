package track

import (
	"fmt"

	"github.com/mebyus/ffd/track/fic"
)

func List(trackpath string) (err error) {
	fics, _, err := fic.Load(trackpath)
	if err != nil {
		return
	}
	for i := range fics {
		shortName := fics[i].Name
		if len(shortName) > 20 {
			shortName = shortName[:20]
		}
		fmt.Printf("%3d) %3s -- %30s -- %4dk [ %s ]\n", i+1, fics[i].Location, shortName, fics[i].Words/1000,
			fics[i].Updated.Format("02.01.2006"))
	}
	return
}
