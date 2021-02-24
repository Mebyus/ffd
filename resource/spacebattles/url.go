package spacebattles

import (
	"fmt"
	"strings"
)

func analyze(url string) (base, name, id string, err error) {
	split := strings.Split(url, "/")
	if len(split) < 2 {
		err = fmt.Errorf("incorrect url")
		return
	}
	parts := 3
	if split[0] == "https:" || split[0] == "http" {
		parts = 5
	}
	if len(split) < parts || split[parts-1] == "" {
		err = fmt.Errorf("incorrect url")
		return
	}
	threadNameSplit := strings.Split(split[parts-1], ".")
	name = threadNameSplit[0]
	if len(threadNameSplit) == 2 {
		id = threadNameSplit[1]
	}
	base = strings.Join(split[:parts], "/")
	return
}

func readerPageURL(baseURL string, pageNumber int64) string {
	url := baseURL + "/reader"
	if pageNumber == 1 {
		return url
	}
	return url + fmt.Sprintf("/page-%d", pageNumber)
}

func indexPageURL(baseURL string) string {
	return baseURL + "/threadmarks"
}
