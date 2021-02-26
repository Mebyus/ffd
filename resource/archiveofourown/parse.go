package archiveofourown

import (
	"fmt"
	"io"

	"github.com/mebyus/ffd/document"
	"github.com/mebyus/ffd/resource/fiction"
	"golang.org/x/net/html"
)

func (t *ao3Tools) Parse(src io.Reader) (book *fiction.Book, err error) {
	chapters, err := parseWhole(src)
	if err != nil {
		err = fmt.Errorf("Parsing piece: %v", err)
		return
	}
	book = &fiction.Book{
		Chapters: chapters,
	}
	return
}

func parseWhole(src io.Reader) (chapters []fiction.Chapter, err error) {
	n, err := html.Parse(src)
	if err != nil {
		return
	}
	d := document.FromNode(n)

	node := d.GetNodeById("chapters")
	if node == nil {
		err = fmt.Errorf("unable to find chapters container")
		return
	}

	document.Detach(node)

	allowed := map[string]bool{
		"p":      true,
		"i":      true,
		"b":      true,
		"table":  true,
		"tr":     true,
		"td":     true,
		"em":     true,
		"strong": true,
	}
	document.Flatten(node, allowed)
	chapter := &fiction.Chapter{
		Body: node,
	}
	chapters = []fiction.Chapter{*chapter}
	return
}
