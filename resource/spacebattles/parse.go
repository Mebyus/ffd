package spacebattles

import (
	"fmt"
	"io"
	"strings"
)

func (t *sbTools) Parse(src io.Reader, dst io.Writer) (err error) {
	text := parseChapter(src)
	_, err = io.Copy(dst, strings.NewReader(text))
	if err != nil {
		err = fmt.Errorf("Saving chapter to destination: %v", err)
		return
	}
	return
}
