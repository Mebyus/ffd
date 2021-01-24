package royalroad

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type result struct {
	index   int
	url     string
	err     error
	content io.Reader
}

func getchapter(url string, index int, out chan<- result, client *http.Client) {
	res := result{
		index: index,
		url:   url,
	}
	request, err := http.NewRequest("GET", url, bytes.NewReader([]byte{}))
	if err != nil {
		err = fmt.Errorf("Composing chapter request: %v", err)
		res.err = err
		out <- res
		return
	}
	response, err := client.Do(request)
	if err != nil {
		err = fmt.Errorf("Doing chapter request: %v", err)
		res.err = err
		out <- res
		return
	}
	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("Chapter response: %s", response.Status)
		res.err = err
		out <- res
		return
	}
	contentStr := parseChapter(response.Body)
	res.content = strings.NewReader(contentStr)
	closeErr := response.Body.Close()
	if closeErr != nil {
		fmt.Printf("Closing chapter response body: %v\n", closeErr)
	}
	out <- res
	return
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
