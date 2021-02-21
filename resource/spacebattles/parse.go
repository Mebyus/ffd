package spacebattles

import (
	"fmt"
	"io"

	"github.com/mebyus/ffd/resource/fiction"
)

func (t *sbTools) Parse(src io.Reader) (book *fiction.Book, err error) {
	chapters, _, err := parsePiece(src)
	if err != nil {
		err = fmt.Errorf("Parsing piece: %v", err)
		return
	}
	book = &fiction.Book{
		Chapters: chapters,
	}
	return
}
