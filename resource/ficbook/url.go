package ficbook

import (
	"fmt"
	"strings"
)

func analyze(url string) (base, id string, err error) {
	split := strings.Split(url, "/")
	if len(split) < 2 {
		err = fmt.Errorf("incorrect url")
		return
	}
	parts := 3
	if split[0] == "https:" || split[0] == "http:" {
		parts = 5
	}
	if len(split) < parts || split[parts-1] == "" {
		err = fmt.Errorf("incorrect url")
		return
	}
	id = split[parts-1]
	base = strings.Join(split[:parts], "/")
	return
}

func indexPageURL(baseURL string) string {
	return baseURL
}
