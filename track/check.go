package track

import (
	"fmt"
	"strconv"
	"time"

	"github.com/mebyus/ffd/resource"
	"github.com/mebyus/ffd/track/fic"
)

func Check(trackpath string, target string) (err error) {
	ficNumber, err := strconv.ParseInt(target, 10, 64)
	if err != nil {
		err = checkAll(trackpath)
	} else {
		err = checkByNumber(trackpath, int(ficNumber))
	}
	return
}

func updateFic(f, u *fic.Info) (updated bool) {
	newChapters := fic.Compare(f.Chapters, u.Chapters)
	f.Check.NewChapters = newChapters
	f.Check.Words = fic.CountWords(newChapters)
	f.Chapters = append(f.Chapters, newChapters...)
	f.Words = fic.CountWords(f.Chapters)
	f.Updated = fic.UpdatedTime(f.Chapters)
	f.Name = u.Name
	f.Author = u.Author
	f.BaseURL = u.BaseURL
	f.Annotation = u.Annotation
	f.Created = u.Created
	f.Finished = u.Finished
	f.Location = u.Location
	f.Check.Time = time.Now()
	if len(newChapters) > 0 {
		updated = true
	}
	return
}

func updatedMsg(f *fic.Info) string {
	if len(f.Check.NewChapters) == 0 {
		return ""
	}
	return fmt.Sprintf("%d new chapters (%dk words) in %s\n", len(f.Check.NewChapters),
		f.Check.Words/1000, f.BaseURL)
}

func checkByNumber(trackpath string, n int) (err error) {
	fics, originpath, err := fic.Load(trackpath)
	if err != nil {
		return
	}
	if n < 1 || n > len(fics) {
		err = fmt.Errorf("fic number = %d exceeds boundaries [%d, %d]", n, 1, len(fics))
		return
	}
	f := &fics[n-1]
	fmt.Printf("Checking [ %s ]\n", f.BaseURL)
	updatedFic, err := resource.Check(f.BaseURL)
	if err != nil {
		return err
	}
	updated := updateFic(f, updatedFic)
	fmt.Println()
	if updated {
		fmt.Println(updatedMsg(f))
	}
	err = fic.Save(originpath, fics)
	if err != nil {
		return
	}
	return
}

func checkAll(trackpath string) (err error) {
	fics, originpath, err := fic.Load(trackpath)
	if err != nil {
		return
	}
	updatedMessages := make([]string, 0)
	for i := range fics {
		if fics[i].Check.Suppressed || fics[i].Finished {
			continue
		}
		fmt.Printf("Checking [ %s ]\n", fics[i].BaseURL)
		updatedFic, err := resource.Check(fics[i].BaseURL)
		if err != nil {
			return err
		}
		updated := updateFic(&fics[i], updatedFic)
		if updated {
			updatedMessages = append(updatedMessages, updatedMsg(&fics[i]))
		}
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
