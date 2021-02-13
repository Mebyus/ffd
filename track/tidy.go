package track

import (
	"sort"

	"github.com/mebyus/ffd/logs"
	"github.com/mebyus/ffd/resource"
	"github.com/mebyus/ffd/track/fic"
)

func Tidy(trackpath string, cleanChapters, cleanUpdates bool) (err error) {
	oldfics, originpath, err := fic.Load(trackpath)
	if err != nil {
		return
	}
	newfics := make([]fic.Info, 0)
	for _, f := range oldfics {
		if cleanChapters {
			f.Chapters = nil
		}
		if cleanUpdates {
			f.Check = fic.Check{}
		}
		location, err := resource.GetLocationForTarget(f.BaseURL)
		if err != nil {
			logs.Warn.Printf("removed [ %s ]: %v\n", f.BaseURL, err)
		} else {
			f.Location = location
			newfics = append(newfics, f)
		}
	}
	sort.Slice(newfics, func(i, j int) bool {
		return newfics[i].BaseURL < newfics[j].BaseURL
	})
	err = fic.Save(originpath, newfics)
	if err != nil {
		return
	}
	return
}
