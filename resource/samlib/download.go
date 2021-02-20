package samlib

import (
	"fmt"
	"io"
	"strings"

	"github.com/mebyus/ffd/cmn"
	"github.com/mebyus/ffd/document"
	"github.com/mebyus/ffd/planner"
	"github.com/mebyus/ffd/resource/fiction"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"golang.org/x/text/encoding/charmap"
)

func (t *slTools) Download(target string, saveSource bool) (book *fiction.Book, err error) {
	page, err := cmn.GetBody(target, planner.Client)
	if err != nil {
		return
	}
	defer cmn.SmartClose(page)
	book, err = parsePage(charmap.Windows1251.NewDecoder().Reader(page))
	return
}

func parsePage(source io.Reader) (book *fiction.Book, err error) {
	n, err := html.Parse(source)
	if err != nil {
		return
	}
	d := document.FromNode(n)
	chapters, err := extractChapters(d)
	if err != nil {
		return
	}
	book = &fiction.Book{
		Title:    "1111",
		Chapters: chapters,
	}
	return
}

func extractChapters(d *document.Document) (chapters []fiction.Chapter, err error) {
	nodes := d.GetNodesByTag("xxx7")
	if len(nodes) == 0 {
		err = fmt.Errorf("unable to locate chapter text container")
		return
	}
	root := nodes[0]
	root.Data = "section"
	root.DataAtom = atom.Section
	document.Detach(root)
	child := root.FirstChild
	for child != nil {
		if child.Type == html.ElementNode && child.Data == "dd" {
			child.Data = "p"
			child.DataAtom = atom.P
			textNode := child.FirstChild
			if textNode != nil && textNode.Type == html.TextNode {
				textNode.Data = strings.ReplaceAll(textNode.Data, "\n", "")
			}
		}
		child = child.NextSibling
	}
	chapters = append(chapters, fiction.Chapter{
		Body: root,
	})
	return
}
