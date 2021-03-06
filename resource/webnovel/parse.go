package webnovel

import (
	"fmt"
	"io"

	"github.com/mebyus/ffd/resource/fiction"
)

func (t *wnTools) Parse(src io.Reader) (book *fiction.Book, err error) {
	chapter, err := parseChapter(src)
	if err != nil {
		err = fmt.Errorf("Parsing piece: %v", err)
		return
	}
	book = &fiction.Book{
		Chapters: []fiction.Chapter{*chapter},
	}
	return
}
