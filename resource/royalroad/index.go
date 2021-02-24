package royalroad

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/mebyus/ffd/cmn"
	"github.com/mebyus/ffd/document"
	"github.com/mebyus/ffd/track/fic"
	"golang.org/x/net/html"
)

const rootURL = "https://" + Hostname

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
	hrefs, _, err := parseIndex(indexPage)
	if err != nil {
		return
	}
	urls = make([]string, len(hrefs))
	for i := range hrefs {
		urls[i] = rootURL + hrefs[i]
	}
	fmt.Printf("Index page parsed. Fic contains %d chapter%s total\n", len(urls), cmn.Plural(len(urls)))
	return
}

func parseIndex(source io.Reader) (hrefs []string, f *fic.Info, err error) {
	n, err := html.Parse(source)
	if err != nil {
		return
	}
	d := document.FromNode(n)
	table := d.GetNodeById("chapters")
	hrefs = document.FindAttributeValues(table, "href")
	chapters := exctactChaptersInfo(table)
	var created time.Time
	if len(chapters) != 0 {
		created = chapters[0].Created
	}
	f = &fic.Info{
		Name:       extractFicName(d),
		Annotation: extractAnnotation(d),
		Author:     extractAuthor(d),
		Finished:   extractFinished(d),
		Created:    created,
		Chapters:   chapters,
	}
	return
}

func extractFicName(d *document.Document) (name string) {
	nodes := d.GetNodesByTag("h1")
	if len(nodes) == 0 {
		fmt.Println("unable to locate fic title node")
		return
	}
	name = document.FindFirstNonSpaceText(nodes[0])
	return
}

func extractAnnotation(d *document.Document) (annotation string) {
	nodes := d.GetNodesByClass("description")
	if len(nodes) == 0 {
		fmt.Println("unable to locate annotation container node")
		return
	}
	annotation = document.FindFirstNonSpaceText(nodes[0])
	return
}

func extractAuthor(d *document.Document) (author string) {
	nodes := d.GetNodesByTag("h4")
	if len(nodes) == 0 {
		fmt.Println("unable to locate author username node")
		return
	}
	author = document.FindLastNonSpaceText(nodes[0])
	return
}

func extractFinished(d *document.Document) (finished bool) {
	nodes := d.GetNodesByTagClass("span", "label-default")
	for _, n := range nodes {
		if document.FindFirstNonSpaceText(n) == "COMPLETED" {
			return true
		}
	}
	return
}

func exctactChaptersInfo(table *html.Node) (chapters []fic.Chapter) {
	tbody := document.FindFirstByTag(table, "tbody")
	texts := document.FindNonSpaceTexts(tbody)
	for i := 0; i < len(texts); i += 3 {
		chapters = append(chapters, fic.Chapter{
			Name: texts[i],
		})
	}
	datetimes := document.FindTagAttributeValues(tbody, "time", "title")
	if len(datetimes) != len(chapters) {
		fmt.Printf("number of timestamps (%d) and chapters (%d) does not match\n", len(datetimes), len(chapters))
	} else {
		for i := range chapters {
			t, err := time.Parse("Monday, January 2, 2006 3:04 PM", datetimes[i])
			if err != nil {
				fmt.Printf("unable to parse chapter creation time [ %s ]: %v\n", datetimes[i], err)
			} else {
				chapters[i].Created = t
			}
		}
	}
	return
}
