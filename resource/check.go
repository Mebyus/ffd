package resource

import (
	"fmt"

	"github.com/mebyus/ffd/track/fic"
)

func Check(target string) (c []fic.Chapter, err error) {
	t, err := ChooseByTarget(target)
	if err != nil {
		err = fmt.Errorf("choosing tool for %s: %v", target, err)
		return
	}
	c = t.Check(target)
	return
}
