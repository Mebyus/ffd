package spacebattles

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/mebyus/ffd/document"
	"github.com/mebyus/ffd/resource/fiction"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type sbParser struct {
}

func parsePiece(source io.Reader) (chapters []fiction.Chapter, pages int64, err error) {
	n, err := html.Parse(source)
	if err != nil {
		return
	}
	d := document.FromNode(n)

	pageNavs := d.GetNodesByClass("pageNav-main")
	if len(pageNavs) == 0 {
		pages = 1
	} else {
		pages = extractNumberOfPages(pageNavs[0])
		if pages == 0 {
			fmt.Printf("unable to determine number of pages\n")
			pages = 1
		}
	}
	threadmarks := extractThreadmarks(d.GetNodesByClass("threadmarkLabel"))
	chapters = extractChapters(d.GetNodesByClass("bbWrapper"))

	if len(threadmarks) != len(chapters) {
		fmt.Printf("number of threadmarks (%d) and chapters (%d) does not match\n", len(threadmarks), len(chapters))
	} else {
		for i := range chapters {
			chapters[i].Title = threadmarks[i]
		}
	}
	return
}

func extractNumberOfPages(pageNav *html.Node) (pages int64) {
	numberText := document.FindLastNonSpaceText(pageNav)
	pages, err := strconv.ParseInt(numberText, 10, 64)
	if err != nil {
		pages = 0
	}
	return
}

func extractChapters(posts []*html.Node) (chapters []fiction.Chapter) {
	for _, post := range posts {
		if post.Parent != nil && !document.HasClass(post.Parent, "threadmarkListingHeader-extraInfoChild") {
			chapters = append(chapters, *extractChapter(post))
		}
	}
	return
}

func extractChapter(post *html.Node) (chapter *fiction.Chapter) {
	document.Detach(post)

	root := &html.Node{
		Data:     "section",
		DataAtom: atom.Section,
		Type:     html.ElementNode,
	}

	child := post.FirstChild
	paragraph := document.NewParagraph()
	root.AppendChild(paragraph)
	for child != nil {
		next := child.NextSibling

		child.Parent = nil
		child.PrevSibling = nil
		child.NextSibling = nil

		switch child.Type {
		case html.TextNode:
			if strings.TrimSpace(child.Data) != "" {
				paragraph.AppendChild(child)
			}
		case html.ElementNode:
			switch child.Data {
			case "br":
				if paragraph.FirstChild != nil {
					paragraph = document.NewParagraph()
					root.AppendChild(paragraph)
				}
			case "i", "b", "em", "strong", "a":
				paragraph.AppendChild(child)
			case "div":
				paragraph.AppendChild(child)
				paragraph = document.NewParagraph()
				root.AppendChild(paragraph)
			case "table":
				root.AppendChild(child)
				paragraph = document.NewParagraph()
				root.AppendChild(paragraph)
			}
		}

		child = next
	}

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

func extractThreadmarks(labels []*html.Node) (threadmarks []string) {
	for _, label := range labels {
		threadmarks = append(threadmarks, extractThreadmark(label))
	}
	return
}

func extractThreadmark(label *html.Node) (threadmark string) {
	if label.FirstChild != nil && label.FirstChild.Type == html.TextNode {
		threadmark = strings.TrimSpace(label.FirstChild.Data)
	}
	return
}
