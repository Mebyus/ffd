package archiveofourown

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/mebyus/ffd/cmn"
	"github.com/mebyus/ffd/document"
	"github.com/mebyus/ffd/planner"
	"github.com/mebyus/ffd/track/fic"
	"golang.org/x/net/html"
)

func (t *ao3Tools) Check(target string) (info *fic.Info, err error) {
	baseURL, ficID, err := analyze(target)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Downloading index page...")
	start := time.Now()
	indexPage, err := cmn.GetBody(indexPageURL(baseURL), planner.Client)
	if err != nil {
		fmt.Printf("\n%v\n", err)
		return
	}
	fmt.Printf(" [ OK ] %v\n", time.Since(start))
	defer cmn.SmartClose(indexPage)

	_, info, err = parseIndex(indexPage)
	if err != nil {
		fmt.Println(err)
		return
	}
	info.ID = ficID
	info.BaseURL = baseURL
	info.Location = fic.ArchiveOfOurOwn
	return
}

func parseIndex(source io.Reader) (hrefs []string, f *fic.Info, err error) {
	n, err := html.Parse(source)
	if err != nil {
		return
	}
	d := document.FromNode(n)
	hrefs, chapters := extractChaptersInfo(d)
	name, author := extractNameAndAuthor(d)
	var created time.Time
	if len(chapters) != 0 {
		created = chapters[0].Created
	}
	f = &fic.Info{
		Name:     name,
		Author:   author,
		Created:  created,
		Chapters: chapters,
	}
	return
}

func extractNameAndAuthor(d *document.Document) (name, author string) {
	nodes := d.GetNodesByTagClass("h2", "heading")
	if len(nodes) == 0 {
		fmt.Println("unable to locate index heading node")
		return
	}
	heading := nodes[0]
	links := document.FindByTag(heading, "a")
	if len(links) > 0 {
		name = document.FindFirstNonSpaceText(links[0])
	} else {
		fmt.Println("unable to locate fic name node")
	}
	if len(links) > 1 {
		author = document.FindFirstNonSpaceText(links[1])
	} else {
		fmt.Println("unable to locate author node")
	}
	return
}

func extractChaptersInfo(d *document.Document) (hrefs []string, chapters []fic.Chapter) {
	nodes := d.GetNodesByTag("ol")
	if len(nodes) == 0 {
		fmt.Println("unable to locate index container node")
		return
	}
	ol := nodes[0]
	hrefs = document.FindAttributeValues(ol, "href")
	namesAndDates := document.FindNonSpaceTexts(ol)
	if len(namesAndDates)%2 != 0 {
		fmt.Printf("odd number (%d) of texts in names and dates slice\n", len(namesAndDates))
	}
	if 2*len(hrefs) != len(namesAndDates) {
		fmt.Printf("number of hrefs (%d) and chapters (%d) does not match\n", len(hrefs), len(namesAndDates)/2)
	}
	for i, href := range hrefs {
		if 2*i+2 >= len(namesAndDates) {
			break
		}
		t, err := time.Parse("(2006-01-02)", namesAndDates[2*i+1])
		if err != nil {
			fmt.Printf("unable to parse chapter creation time [ %s ]: %v\n", namesAndDates[2*i+1], err)
		}
		chapters = append(chapters, fic.Chapter{
			ID:      extractChapterID(href),
			Name:    namesAndDates[2*i],
			Created: t,
		})
	}
	return
}

func extractChapterID(href string) (id string) {
	split := strings.Split(href, "/")
	if len(split) == 0 {
		return
	}
	id = split[len(split)-1]
	return
}
