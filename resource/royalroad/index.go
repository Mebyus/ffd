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
