package royalroad

import (
	"fmt"
	"strings"
)

type rrTools struct {
	ficName    string
	chapters   int
	saveSource bool
}

func NewTools() *rrTools {
	return &rrTools{}
}

func analyze(url string) (base, name string, err error) {
	split := strings.Split(url, "/")
	if len(split) < 2 {
		err = fmt.Errorf("incorrect url")
		return
	}
	parts := 4
	if split[0] == "https:" || split[0] == "http" {
		parts = 6
	}
	if len(split) < parts || split[parts-1] == "" {
		err = fmt.Errorf("incorrect url")
		return
	}
	name = split[parts-1]
	base = strings.Join(split[:parts], "/")
	return
}
