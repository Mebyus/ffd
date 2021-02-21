package archiveofourown

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
	parts := 3
	if split[0] == "https:" || split[0] == "http" {
		parts = 5
	}
	if len(split) < parts || split[parts-1] == "" {
		err = fmt.Errorf("incorrect url")
		return
	}
	name = split[parts-1]
	base = strings.Join(split[:parts], "/")
	return
}

func readerPageURL(baseURL string, chapterID string) string {
	return baseURL + "/chapters/" + chapterID
}

func indexPageURL(baseURL string) string {
	return baseURL + "/navigate"
}
