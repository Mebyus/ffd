package track

import (
	"fmt"

	"github.com/mebyus/ffd/track/fic"
)

func List() (err error) {
	fics, err := fic.Load()
	if err != nil {
		return
	}
	for i := range fics {
		chapters := len(fics[i].Chapters)
		if chapters == 0 {
			fmt.Printf("%s [ unknown ]\n", fics[i].Name)
		} else {
			lastUpdate := fics[i].Chapters[chapters-1].Date
			fmt.Printf("%s [ %s ]\n", fics[i].Name, lastUpdate.Format("02.01.2006"))
		}
	}
	return
}
