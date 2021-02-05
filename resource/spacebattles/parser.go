package spacebattles

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

type sbParser struct {
}

func (p *sbParser) parsePiece(source io.Reader) (result io.Reader, err error) {
	tokenizer := html.NewTokenizer(source)
	text := ""
	depth := 0
	insideChapter := false
	insideArticle := false
	insideThreadmark := false
	chapterDepth := 0
	articleDepth := 0
	insertNewline := false
	insertSpace := false

	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			err = tokenizer.Err()
			if err == io.EOF {
				err = nil
				result = strings.NewReader(text)
				return
			}
			err = fmt.Errorf("Page tokenization: %v", err)
			return
		case html.StartTagToken:
			token := tokenizer.Token()
			if token.Data == "article" {
				insideArticle = true
				articleDepth = depth
			} else if insideArticle && !insideChapter && token.Data == "span" && isThreadmark(token.Attr) {
				insideThreadmark = true
			} else if insideArticle && !insideChapter && token.Data == "div" && isPostWrapper(token.Attr) {
				insideChapter = true
				chapterDepth = depth
			} else if insideChapter && (token.Data == "i" || token.Data == "b") {
				text += " "
				insertSpace = true
			}
			depth++
		case html.SelfClosingTagToken:
			token := tokenizer.Token()
			if insideChapter && token.Data == "br" {
				text += "\n"
				insertNewline = true
			}
		case html.EndTagToken:
			depth--
			token := tokenizer.Token()
			if token.Data == "article" && articleDepth == depth {
				insideArticle = false
			} else if insideThreadmark && token.Data == "span" {
				insideThreadmark = false
			} else if insideChapter && token.Data == "div" && chapterDepth == depth {
				insideChapter = false
				text += "\n\n"
			}
		case html.TextToken:
			token := tokenizer.Token()
			if insideThreadmark {
				text += strings.TrimSpace(token.Data) + "\n\n"
			}
			if insideChapter {
				if insertNewline {
					text += strings.TrimSpace(token.Data)
					insertNewline = false
					insertSpace = false
				} else if insertSpace {
					text += strings.TrimSpace(token.Data) + " "
					insertSpace = false
				} else {
					text += strings.TrimSpace(token.Data)
				}
			}
		}
	}
}
