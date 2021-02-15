package royalroad

import (
	"fmt"
	"io"

	"github.com/mebyus/ffd/resource/fiction"
)

func (t *rrTools) Parse(src io.Reader) (book *fiction.Book, err error) {
	_, err = parseChapter(src)
	if err != nil {
		err = fmt.Errorf("Parsing piece: %v", err)
		return
	}
	return
}
