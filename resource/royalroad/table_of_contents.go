package royalroad

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/html"
)

const baseurl = "https://www.royalroad.com"

func geturls(tocurl string, client *http.Client) (urls []string, err error) {
	request, err := http.NewRequest("GET", tocurl, bytes.NewReader([]byte{}))
	if err != nil {
		err = fmt.Errorf("Composing table of contents request: %v", err)
		return
	}
	response, err := client.Do(request)
	if err != nil {
		err = fmt.Errorf("Doing table of contents request: %v", err)
		return
	}
	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("Table of contents response: %s", response.Status)
		return
	}
	urls = parseTableOfContents(response.Body)
	closeErr := response.Body.Close()
	if closeErr != nil {
		fmt.Printf("Closing table of contents response body: %v\n", closeErr)
	}
	return
}

func parseTableOfContents(page io.Reader) []string {
	tokenizer := html.NewTokenizer(page)
	depth := 0
	tableFound := false
	tableDepth := 0
	hrefs := []string{}

	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			err := tokenizer.Err()
			if err == io.EOF {
				return hrefs
			}
			fmt.Printf("Table of contents tokenization: %v\n", err)
		case html.StartTagToken:
			token := tokenizer.Token()
			if !tableFound && token.Data == "table" && isChaptersTable(token.Attr) {
				tableFound = true
				tableDepth = depth
			} else if tableFound && token.Data == "a" {
				href := extracthref(token.Attr)
				if href != "" {
					hrefs = append(hrefs, baseurl+href)
				}
			}
			depth++
		case html.EndTagToken:
			depth--
			token := tokenizer.Token()
			if tableFound && token.Data == "table" && tableDepth == depth {
				return hrefs
			}
		}
	}
}

func isChaptersTable(attrs []html.Attribute) bool {
	for _, attr := range attrs {
		if attr.Key == "id" && attr.Val == "chapters" {
			return true
		}
	}
	return false
}

func extracthref(attrs []html.Attribute) string {
	for _, attr := range attrs {
		if attr.Key == "href" {
			return attr.Val
		}
	}
	return ""
}
