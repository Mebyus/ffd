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
		if len(shortName) > 30 {
			shortName = shortName[:30]
		}
		if fics[i].Bookmark.Chapter > 0 {
			if fics[i].Bookmark.Chapter == len(fics[i].Chapters)+1 {
				fmt.Printf("%3d) %3s -- %30s -- %4dk [ %s ] latest\n", i+1, fics[i].Location, shortName, fics[i].Words/1000,
					fics[i].Updated.Format("02.01.2006"))
			} else {
				fmt.Printf("%3d) %3s -- %30s -- %4dk [ %s ] %d new\n", i+1, fics[i].Location, shortName, fics[i].Words/1000,
					fics[i].Updated.Format("02.01.2006"), 1+len(fics[i].Chapters)-fics[i].Bookmark.Chapter)
			}
		} else {
			fmt.Printf("%3d) %3s -- %30s -- %4dk [ %s ]\n", i+1, fics[i].Location, shortName, fics[i].Words/1000,
				fics[i].Updated.Format("02.01.2006"))
		}
	}
	return
}
