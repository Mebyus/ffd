package resource

import "fmt"

// Download fetches a fic from a given target.
// An appropriate target is fic page URL.
// SaveSource flag indicates whether responses will be saved before parsing
func Download(target string, saveSource bool) (err error) {
	t, err := ChooseByTarget(target)
	if err != nil {
		err = fmt.Errorf("choosing tool for %s: %v", target, err)
		return
	}
	t.Download(target, saveSource)
	return nil
}
