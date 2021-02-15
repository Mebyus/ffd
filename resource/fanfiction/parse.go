package fanfiction

import (
	"fmt"
	"io"
	"strings"

	"github.com/mebyus/ffd/document"
	"github.com/mebyus/ffd/resource/fiction"
	"golang.org/x/net/html"
)

func (t *ffTools) Parse(src io.Reader) (book *fiction.Book, err error) {
	_, _, err = parsePage(src)
	if err != nil {
		err = fmt.Errorf("Parsing piece: %v", err)
		return
	}
	return
}

func parsePage(source io.Reader) (result io.Reader, pages int64, err error) {
	n, err := html.Parse(source)
	if err != nil {
		return
	}
	d := document.FromNode(n)
	storytextDiv := d.GetNodeById("storytext")
	text := extractChapter(storytextDiv)
	result = strings.NewReader(text)
	chapterSelector := d.GetNodeById("chap_select")
	options := document.FindByTag(chapterSelector, "option")
	pages = int64(len(options))
	return
}

func extractChapter(chapterContainer *html.Node) (chapter string) {
	action := func(n *html.Node) {
		switch n.Type {
		case html.TextNode:
			chapter += strings.TrimSpace(n.Data)
		case html.ElementNode:
			if n.Data == "p" {
				chapter += "\n\n"
			}
		}
	}
	document.Walk(chapterContainer, action)
	return
}
