package royalroad

import (
	"fmt"
	"io"
)

func (t *rrTools) Parse(src io.Reader, dst io.Writer) (err error) {
	r, err := parseChapter(src)
	if err != nil {
		err = fmt.Errorf("Parsing piece: %v", err)
		return
	}
	_, err = io.Copy(dst, r)
	if err != nil {
		err = fmt.Errorf("Saving chapter to destination: %v", err)
		return
	}
	return
}
