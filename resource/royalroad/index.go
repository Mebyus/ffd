package royalroad

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/mebyus/ffd/cmn"
	"github.com/mebyus/ffd/document"
	"golang.org/x/net/html"
)

const rootURL = "https://www.royalroad.com"

func getChapterURLs(indexPageURL string, client *http.Client) (urls []string, err error) {
	fmt.Printf("Downloading index page...")
	start := time.Now()
	indexPage, err := cmn.GetBody(indexPageURL, client)
	if err != nil {
		fmt.Println()
		return
	}
	fmt.Printf(" [ OK ] %v\n", time.Since(start))
	defer cmn.SmartClose(indexPage)

	fmt.Printf("Parsing index page...\n")
	urls, err = parseIndex(indexPage)
	if err != nil {
		return
	}
	for i := range urls {
		urls[i] = rootURL + urls[i]
	}
	fmt.Printf("Index page parsed. Fic contains %d chapters total\n", len(urls))
	return
}

func parseIndex(source io.Reader) (hrefs []string, err error) {
	n, err := html.Parse(source)
	if err != nil {
		return
	}
	d := document.FromNode(n)
	hrefs = document.FindAttributeValues(d.GetNodeById("chapters"), "href")
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
					hrefs = append(hrefs, rootURL+href)
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
