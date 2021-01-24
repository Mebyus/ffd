package resource

import (
	"net/url"
)

// Download fetches a fic from a given target.
// An appropriate target is fic page URL.
// SaveSource flag indicates whether responses will be saved before parsing
func Download(target string, saveSource bool) error {
	u, err := url.Parse(target)
	if err != nil {
		return err
	}
	hostname := u.Hostname()
	t, err := Choose(hostname)
	if err != nil {
		return err
	}
	t.Download(target, saveSource)
	return nil
}
