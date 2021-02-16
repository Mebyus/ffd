package royalroad

import (
	"fmt"
	"io"

	"github.com/mebyus/ffd/document"
	"github.com/mebyus/ffd/resource/fiction"
	"golang.org/x/net/html"
)

func parseChapter(source io.Reader) (chapter *fiction.Chapter, err error) {
	n, err := html.Parse(source)
	if err != nil {
		return
	}
	d := document.FromNode(n)

	title := extractTitle(d)
	chapterBody := extractChapterBody(d)
	chapter = &fiction.Chapter{
		Title: title,
		Body:  chapterBody,
	}
	return
}

func extractTitle(d *document.Document) (title string) {
	nodes := d.GetNodesByTag("h1")
	if len(nodes) == 0 {
		fmt.Println("unable to locate chapter title node")
		return
	}
	title = document.FindFirstNonSpaceText(nodes[0])
	return
}

func extractChapterBody(d *document.Document) (root *html.Node) {
	nodes := d.GetNodesByClass("chapter-inner")
	if len(nodes) == 0 {
		fmt.Println("unable to locate chapter text container")
		return
	} else if len(nodes) > 1 {
		fmt.Println("located several potential text containers")
	}
	root = nodes[0]
	document.Detach(root)
	allowed := map[string]bool{
		"p":     true,
		"i":     true,
		"b":     true,
		"table": true,
		"tr":    true,
		"td":    true,
	}
	document.Flatten(root, allowed)
	return
}
