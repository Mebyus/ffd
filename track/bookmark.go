package track

import (
	"fmt"
	"strconv"

	"github.com/mebyus/ffd/track/fic"
)

func Bookmark(trackpath, target, chapter string) (err error) {
	ficNumber, err := strconv.ParseInt(target, 10, 64)
	if err != nil {
		err = fmt.Errorf("only numbers on the list are allowed: %v", err)
	} else {
		err = bookmarkByNumber(trackpath, int(ficNumber), chapter)
	}
	return
}

func bookmarkByNumber(trackpath string, n int, chapter string) (err error) {
	fics, originpath, err := fic.Load(trackpath)
	if err != nil {
		return
	}
	if n < 1 || n > len(fics) {
		err = fmt.Errorf("fic number = %d exceeds boundaries [%d, %d]", n, 1, len(fics))
		return
	}
	f := &fics[n-1]
	if chapter == "latest" {
		f.Bookmark.Chapter = len(f.Chapters) + 1
	} else {
		chapterNumber, err := strconv.ParseInt(chapter, 10, 64)
		if err != nil {
			return fmt.Errorf("%s is not a chapter number: %v", chapter, err)
		}
		if chapterNumber < 0 || int(chapterNumber) > len(f.Chapters) {
			return fmt.Errorf("number = %d is out of fic chapter bounds [%d, %d]",
				chapterNumber, 1, len(f.Chapters))
		}
		f.Bookmark.Chapter = int(chapterNumber)
	}
	err = fic.Save(originpath, fics)
	if err != nil {
		return
	}
	return
}
