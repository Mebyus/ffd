package spacebattles

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/mebyus/ffd/track/fic"
	"golang.org/x/net/html"
)

func (t *sbTools) Check(target string) (cs []fic.Chapter) {
	client := &http.Client{
		Timeout: timeout,
	}
	cs, err := getThreadmarks(getThreadmarksURL(target), client)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func getThreadmarksURL(target string) string {
	return target + "threadmarks/"
}

func getThreadmarks(url string, client *http.Client) (cs []fic.Chapter, err error) {
	request, err := http.NewRequest("GET", url, bytes.NewReader([]byte{}))
	if err != nil {
		err = fmt.Errorf("Composing chapter request: %v", err)
		return
	}
	response, err := client.Do(request)
	if err != nil {
		err = fmt.Errorf("Doing chapter request: %v", err)
		return
	}
	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("Chapter response: %s [ %s ]", response.Status, url)
		return
	}
	cs = parseThreadmarksPage(response.Body)
	closeErr := response.Body.Close()
	if closeErr != nil {
		fmt.Printf("Closing chapter response body: %v\n", closeErr)
	}
	return
}

func parseThreadmarksPage(page io.Reader) (cs []fic.Chapter) {
	tokenizer := html.NewTokenizer(page)
	depth := 0
	insideBody := false
	insideLink := false
	insideTitle := false
	insideWordCount := false
	bodyDepth := 0
	titleDepth := 0
	itemDepth := 0
	threadmark := fic.Chapter{}
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
			if !insideBody && token.Data == "div" && isBody(token.Attr) {
				insideBody = true
				bodyDepth = depth
			} else if insideBody && token.Data == "div" && isThreadmarkItem(token.Attr) {
				threadmark = fic.Chapter{}
				itemDepth = depth
			} else if insideBody && token.Data == "div" && isThreadmarkTitle(token.Attr) {
				insideTitle = true
				titleDepth = depth
			} else if insideBody && token.Data == "a" {
				insideLink = true
			} else if insideBody && token.Data == "dd" {
				insideWordCount = true
			} else if insideBody && token.Data == "time" {
				threadmark.Date = extractDatetime(token.Attr)
			}
			depth++
		case html.EndTagToken:
			depth--
			token := tokenizer.Token()
			if insideBody && bodyDepth == depth && token.Data == "div" {
				insideBody = false
			} else if insideBody && itemDepth == depth && token.Data == "div" {
				if threadmark.Name != "" {
					cs = append(cs, threadmark)
				}
			} else if insideBody && titleDepth == depth && token.Data == "div" {
				insideTitle = false
			} else if insideBody && token.Data == "a" {
				insideLink = false
			} else if insideBody && token.Data == "dd" {
				insideWordCount = false
			}
		case html.TextToken:
			token := tokenizer.Token()
			if insideTitle && insideLink {
				threadmark.Name = token.Data
			}
			if insideWordCount {
				threadmark.Words = convertWordCount(token.Data)
			}
		}
	}
}

func isBody(attrs []html.Attribute) bool {
	for _, attr := range attrs {
		if attr.Key == "class" && strings.Contains(attr.Val, "block-body--threadmarkBody") {
			return true
		}
	}
	return false
}

func isThreadmarkItem(attrs []html.Attribute) bool {
	for _, attr := range attrs {
		if attr.Key == "class" && strings.Contains(attr.Val, "structItem--threadmark") {
			return true
		}
	}
	return false
}

func isThreadmarkTitle(attrs []html.Attribute) bool {
	for _, attr := range attrs {
		if attr.Key == "class" && strings.Contains(attr.Val, "structItem-title") {
			return true
		}
	}
	return false
}

func extractDatetime(attrs []html.Attribute) (t time.Time) {
	for _, attr := range attrs {
		if attr.Key == "datetime" {
			t, _ = time.Parse("2006-01-02T15:04:05-0700", attr.Val)
			return
		}
	}
	return
}

func convertWordCount(str string) int64 {
	if strings.HasSuffix(str, "k") {
		count, err := strconv.ParseFloat(strings.TrimRight(str, "k"), 64)
		if err != nil {
			return 0
		}
		return int64(1000 * count)
	}
	count, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return count
}
