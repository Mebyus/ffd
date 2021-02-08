package spacebattles

import (
	"fmt"
	"io"
)

func (t *sbTools) Parse(src io.Reader, dst io.Writer) (err error) {
	r, _, err := parsePiece(src)
	_, err = io.Copy(dst, r)
	if err != nil {
		err = fmt.Errorf("Saving chapter to destination: %v", err)
		return
	}
	return
}
