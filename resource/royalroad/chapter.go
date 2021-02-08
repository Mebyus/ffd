package royalroad

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

type result struct {
	index   int
	url     string
	err     error
	content io.Reader
}

func parseChapter(page io.Reader) (content string) {
	tokenizer := html.NewTokenizer(page)
	depth := 0
	insideChapter := false
	insideTitle := false
	chapterDepth := 0
	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			err := tokenizer.Err()
			if err == io.EOF {
				return
			}
			fmt.Printf("Page tokenization: %v\n", err)
		case html.StartTagToken:
			token := tokenizer.Token()
			if !insideChapter && token.Data == "div" && isChapterContent(token.Attr) {
				insideChapter = true
				chapterDepth = depth
			} else if !insideChapter && token.Data == "h1" {
				insideTitle = true
			}
			depth++
		case html.EndTagToken:
			depth--
			token := tokenizer.Token()
			if insideChapter && token.Data == "div" && chapterDepth == depth {
				return
			} else if insideChapter && token.Data == "p" {
				content += "\n\n"
			} else if insideTitle && token.Data == "h1" {
				insideTitle = false
			}
		case html.TextToken:
			token := tokenizer.Token()
			if insideChapter {
				content += strings.TrimSpace(token.Data)
			} else if insideTitle {
				content += strings.TrimSpace(token.Data) + "\n\n"
			}
		}
	}
}

func isChapterContent(attrs []html.Attribute) bool {
	for _, attr := range attrs {
		if attr.Key == "class" && attr.Val == "chapter-inner chapter-content" {
			return true
		}
	}
	return false
}
