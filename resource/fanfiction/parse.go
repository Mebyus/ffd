package fanfiction

import (
	"fmt"
	"io"

	"github.com/mebyus/ffd/document"
	"github.com/mebyus/ffd/resource/fiction"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func (t *ffTools) Parse(src io.Reader) (book *fiction.Book, err error) {
	chapter, _, err := parsePage(src)
	if err != nil {
		err = fmt.Errorf("Parsing piece: %v", err)
		return
	}
	book = &fiction.Book{
		Chapters: []fiction.Chapter{*chapter},
	}
	return
}

func parsePage(source io.Reader) (chapter *fiction.Chapter, pages int64, err error) {
	n, err := html.Parse(source)
	if err != nil {
		return
	}
	d := document.FromNode(n)

	chapterRoot, err := extractChapter(d)
	if err != nil {
		return
	}
	chapter = &fiction.Chapter{
		Body: chapterRoot,
	}
	chapterSelector := d.GetNodeById("chap_select")
	options := document.FindByTag(chapterSelector, "option")
	pages = int64(len(options))
	return
}

func extractChapter(d *document.Document) (root *html.Node, err error) {
	storytextDiv := d.GetNodeById("storytext")
	if storytextDiv == nil {
		err = fmt.Errorf("unable to locate chapter text container")
		return
	}
	document.Detach(storytextDiv)
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
	root = storytextDiv
	document.Flatten(root, allowed)
	root.Data = "section"
	root.DataAtom = atom.Section
	return
}
