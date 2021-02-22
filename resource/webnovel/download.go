package webnovel

import (
	"fmt"
	"io"

	"github.com/mebyus/ffd/cmn"
	"github.com/mebyus/ffd/document"
	"github.com/mebyus/ffd/planner"
	"github.com/mebyus/ffd/resource/fiction"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func (t *wnTools) Download(target string, saveSource bool) (book *fiction.Book, err error) {
	page, err := cmn.GetBody(target, planner.Client)
	if err != nil {
		return
	}
	defer cmn.SmartClose(page)
	// book, err = parseChapter(charmap.Windows1251.NewDecoder().Reader(page))
	return
}

func parseChapter(source io.Reader) (chapter *fiction.Chapter, err error) {
	n, err := html.Parse(source)
	if err != nil {
		return
	}
	d := document.FromNode(n)
	chapter, err = extractChapter(d)
	if err != nil {
		return
	}
	return
}

func extractChapter(d *document.Document) (chapter *fiction.Chapter, err error) {
	nodes := d.GetNodesByClass("cha-words")
	if len(nodes) == 0 {
		err = fmt.Errorf("unable to locate chapter text container")
		return
	}
	root := nodes[0]
	root.Data = "section"
	root.DataAtom = atom.Section
	document.Detach(root)
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
	document.Flatten(root, allowed)
	chapter = &fiction.Chapter{
		Body: root,
	}
	return
}
