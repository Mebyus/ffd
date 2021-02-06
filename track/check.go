package track

import (
	"fmt"
	"time"

	"github.com/mebyus/ffd/resource"
	"github.com/mebyus/ffd/track/fic"
)

func Check(trackpath string) (err error) {
	fics, originpath, err := fic.Load(trackpath)
	if err != nil {
		return
	}
	for i := range fics {
		if fics[i].Check.Suppressed {
			continue
		}
		chapters, err := resource.Check(fics[i].BaseURL)
		if err != nil {
			return err
		}
		newChapters := fic.Compare(fics[i].Chapters, chapters)
		fics[i].Check.NewChapters = newChapters
		fics[i].Chapters = append(chapters, newChapters...)
		fics[i].Words = fic.CountWords(fics[i].Chapters)
		fics[i].Check.Time = time.Now()
		if len(newChapters) != 0 {
			fmt.Printf("%d new chapters in %s\n", len(newChapters), fics[i].BaseURL)
		}
	}
	err = fic.Save(originpath, fics)
	if err != nil {
		return
	}
	return
}
