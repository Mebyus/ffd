package resource

import (
	"fmt"
	"net/url"

	"github.com/mebyus/ffd/track/fic"
)

func Check(target string) (c []fic.Chapter, err error) {
	u, err := url.Parse(target)
	if err != nil {
		return
	}
	hostname := u.Hostname()
	t, err := Choose(hostname)
	if err != nil {
		err = fmt.Errorf("choosing tool for %s: %v", target, err)
		return
	}
	c = t.Check(target)
	return
}
