package royalroad

import (
	"fmt"
	"time"

	"github.com/mebyus/ffd/cmn"
	"github.com/mebyus/ffd/planner"
	"github.com/mebyus/ffd/track/fic"
)

func (t *rrTools) Check(target string) (f *fic.Info) {
	baseURL, _, err := analyze(target)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Downloading index page...")
	start := time.Now()
	indexPage, err := cmn.GetBody(baseURL, planner.Client)
	if err != nil {
		fmt.Printf("\n%v\n", err)
		return
	}
	fmt.Printf(" [ OK ] %v\n", time.Since(start))
	defer cmn.SmartClose(indexPage)

	_, f, err = parseIndex(indexPage)
	if err != nil {
		fmt.Println(err)
		return
	}
	f.BaseURL = baseURL
	f.Location = fic.RoyalRoad
	return
}
