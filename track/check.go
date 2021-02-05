package track

import (
	"fmt"
	"time"

	"github.com/mebyus/ffd/resource"
	"github.com/mebyus/ffd/track/fic"
)

func Check() (err error) {
	fics, originpath, err := fic.Load("")
	if err != nil {
		return
	}
	for i := range fics {
		if fics[i].Suppressed {
			continue
		}
		chapters, err := resource.Check(fics[i].URL)
		if err != nil {
			return err
		}
		newChapters := fic.Compare(fics[i].Chapters, chapters)
		fics[i].Check.NewChapters = newChapters
		fics[i].Chapters = append(chapters, newChapters...)
		fics[i].Words = fic.CountWords(fics[i].Chapters)
		fics[i].Check.Date = time.Now()
		if len(newChapters) != 0 {
			fmt.Printf("%d new chapters in %s\n", len(newChapters), fics[i].URL)
		}
	}
	err = fic.Save(originpath, fics)
	if err != nil {
		return
	}
	return
}
