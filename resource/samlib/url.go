package samlib

import (
	"fmt"
	"strings"
)

func analyze(url string) (base, name string, err error) {
	split := strings.Split(url, "/")
	if len(split) < 2 {
		err = fmt.Errorf("incorrect url")
		return
	}
	parts := 4
	if split[0] == "https:" || split[0] == "http:" {
		parts = 6
	}
	if len(split) < parts || split[parts-1] == "" {
		err = fmt.Errorf("incorrect url")
		return
	}
	shtmlNameSplit := strings.Split(split[parts-1], ".")
	name = strings.Replace(shtmlNameSplit[0], "_", "-", -1) + "_" + strings.Replace(split[parts-2], "_", "-", -1)
	base = strings.Join(split[:parts], "/")
	return
}
