package royalroad

import (
	"fmt"
	"io"

	"github.com/mebyus/ffd/resource/internal"
)

func (t *rrTools) Parse(src io.Reader) (book *internal.Book, err error) {
	_, err = parseChapter(src)
	if err != nil {
		err = fmt.Errorf("Parsing piece: %v", err)
		return
	}
	return
}
