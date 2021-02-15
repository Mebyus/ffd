package spacebattles

import (
	"fmt"
	"io"

	"github.com/mebyus/ffd/resource/internal"
)

func (t *sbTools) Parse(src io.Reader) (book *internal.Book, err error) {
	_, _, err = parsePiece(src)
	if err != nil {
		err = fmt.Errorf("Parsing piece: %v", err)
		return
	}
	return
}
