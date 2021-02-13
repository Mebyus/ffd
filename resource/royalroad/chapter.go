package royalroad

import (
	"io"
	"strings"

	"github.com/mebyus/ffd/document"
	"github.com/mebyus/ffd/logs"
	"golang.org/x/net/html"
)

func parseChapter(source io.Reader) (result io.Reader, err error) {
	n, err := html.Parse(source)
	if err != nil {
		return
	}
	d := document.FromNode(n)

	title := extractTitle(d)
	text := extractChapterText(d)
	result = strings.NewReader("\n\n" + title + text)
	return
}

func extractTitle(d *document.Document) (title string) {
	nodes := d.GetNodesByTag("h1")
	if len(nodes) == 0 {
		logs.Warn.Println("unable to locate chapter title node")
		return
	}
	title = document.FindFirstNonSpaceText(nodes[0])
	return
}

func extractChapterText(d *document.Document) (text string) {
	nodes := d.GetNodesByClass("chapter-inner")
	if len(nodes) == 0 {
		logs.Warn.Println("unable to locate chapter text container")
		return
	} else if len(nodes) > 1 {
		logs.Warn.Println("located several potential text containers")
	}

	action := func(n *html.Node) {
		switch n.Type {
		case html.TextNode:
			text += strings.TrimSpace(n.Data)
		case html.ElementNode:
			if n.Data == "p" {
				text += "\n\n"
			}
		}
	}
	document.Walk(nodes[0], action)
	return
}
