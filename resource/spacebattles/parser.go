package spacebattles

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/mebyus/ffd/document"
	"golang.org/x/net/html"
)

type sbParser struct {
}

func parsePiece(source io.Reader) (result io.Reader, pages int64, err error) {
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
	chapters := extractChapters(d.GetNodesByClass("bbWrapper"))

	readers := make([]io.Reader, len(chapters))
	if len(threadmarks) != len(chapters) {
		fmt.Printf("number of threadmarks (%d) and chapters (%d) does not match\n", len(threadmarks), len(chapters))
		for i := range chapters {
			readers[i] = strings.NewReader(chapters[i])
		}
	} else {
		for i := range chapters {
			readers[i] = strings.NewReader(threadmarks[i] + "\n\n" + chapters[i])
		}
	}
	result = io.MultiReader(readers...)
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

func extractChapters(posts []*html.Node) (chapters []string) {
	for _, post := range posts {
		chapters = append(chapters, extractChapter(post))
	}
	return
}

func extractChapter(post *html.Node) (chapter string) {
	action := func(n *html.Node) {
		switch n.Type {
		case html.TextNode:
			chapter += strings.TrimSpace(n.Data)
		case html.ElementNode:
			if n.Data == "br" {
				chapter += "\n"
			}
		}
	}
	document.Walk(post, action)
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
