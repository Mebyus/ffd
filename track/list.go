package track

import (
	"fmt"
	"strconv"

	"github.com/mebyus/ffd/track/fic"
)

func List(trackpath, target string) (err error) {
	if target == "" {
		err = listAll(trackpath)
	} else {
		err = listByNumber(trackpath, target)
	}
	return
}

func listAll(trackpath string) (err error) {
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

func listByNumber(trackpath, target string) (err error) {
	ficNumber, err := strconv.ParseInt(target, 10, 64)
	if err != nil {
		err = fmt.Errorf("only numbers on the list are allowed: %v", err)
		return
	}
	fics, _, err := fic.Load(trackpath)
	if err != nil {
		return
	}
	if ficNumber < 1 || int(ficNumber) > len(fics) {
		err = fmt.Errorf("fic number = %d exceeds boundaries [%d, %d]", ficNumber, 1, len(fics))
		return
	}
	f := &fics[int(ficNumber)-1]
	for i := range f.Chapters {
		shortName := f.Chapters[i].Name
		if len(shortName) > 30 {
			shortName = shortName[:30]
		}
		if f.Bookmark.Chapter > 0 && i+1 >= f.Bookmark.Chapter {
			fmt.Printf("%3d. -- %30s -- [ %s ] new\n", i+1, shortName,
				f.Chapters[i].Created.Format("02.01.2006"))
		} else {
			fmt.Printf("%3d. -- %30s -- [ %s ]\n", i+1, shortName,
				f.Chapters[i].Created.Format("02.01.2006"))
		}
	}
	return
}
