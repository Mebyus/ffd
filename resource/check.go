package resource

import (
	"fmt"

	"github.com/mebyus/ffd/track/fic"
)

func Check(target string) (f *fic.Info, err error) {
	t, err := ChooseByTarget(target)
	if err != nil {
		err = fmt.Errorf("choosing tool for %s: %v", target, err)
		return
	}
	f, err = t.Check(target)
	return
}
