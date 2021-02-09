package track

import (
	"fmt"
	"sort"

	"github.com/mebyus/ffd/resource"
	"github.com/mebyus/ffd/track/fic"
)

func Tidy(trackpath string) (err error) {
	oldfics, originpath, err := fic.Load(trackpath)
	if err != nil {
		return
	}
	newfics := make([]fic.Info, 0)
	for _, fic := range oldfics {
		location, err := resource.GetLocationForTarget(fic.BaseURL)
		if err != nil {
			fmt.Printf("removed [ %s ]: %v\n", fic.BaseURL, err)
		} else {
			fic.Location = location
			newfics = append(newfics, fic)
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
