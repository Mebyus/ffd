package track

import (
	"fmt"

	"github.com/mebyus/ffd/track/fic"
)

func List() (err error) {
	fics, _, err := fic.Load("")
	if err != nil {
		return
	}
	for i := range fics {
		chapters := len(fics[i].Chapters)
		if chapters == 0 {
			fmt.Printf("%s %dk [ unknown ]\n", fics[i].Name, fics[i].Words/1000)
		} else {
			lastUpdate := fics[i].Chapters[chapters-1].Date
			fmt.Printf("%s %dk [ %s ]\n", fics[i].Name, fics[i].Words/1000, lastUpdate.Format("02.01.2006"))
		}
	}
	return
}
