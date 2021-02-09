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
	updatedMessages := make([]string, 0)
	for i := range fics {
		if fics[i].Check.Suppressed {
			continue
		}
		fmt.Printf("Checking [ %s ]\n", fics[i].BaseURL)
		updatedFic, err := resource.Check(fics[i].BaseURL)
		if err != nil {
			return err
		}
		newChapters := fic.Compare(fics[i].Chapters, updatedFic.Chapters)
		updatedFic.Check.NewChapters = newChapters
		updatedFic.Chapters = append(fics[i].Chapters, newChapters...)
		updatedFic.Words = fic.CountWords(fics[i].Chapters)
		updatedFic.Updated = fic.UpdatedTime(updatedFic.Chapters)
		updatedFic.Check.Time = time.Now()
		updatedFic.Check.Words = fic.CountWords(newChapters)
		if len(newChapters) != 0 {
			updatedMessages = append(updatedMessages,
				fmt.Sprintf("%d new chapters (%dk words) in %s\n", len(newChapters),
					updatedFic.Check.Words/1000, fics[i].BaseURL),
			)
		}
		fics[i] = *updatedFic
	}
	fmt.Println()
	for _, msg := range updatedMessages {
		fmt.Printf(msg)
	}
	err = fic.Save(originpath, fics)
	if err != nil {
		return
	}
	return
}
