package resource

import (
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
		return
	}
	c = t.Check(target)
	return
}
