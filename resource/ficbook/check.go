package ficbook

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

func (t *fbTools) Check(target string) (info *fic.Info, err error) {
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
	info.Location = fic.FicBook
	return
}

func parseIndex(source io.Reader) (hrefs []string, f *fic.Info, err error) {
	n, err := html.Parse(source)
	if err != nil {
		return
	}
	d := document.FromNode(n)
	hrefs, chapters := extractChaptersInfo(d)
	name := extractFicName(d)
	author := extractAuthor(d)
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

func extractAuthor(d *document.Document) (author string) {
	aNodes := d.GetNodesByTagClass("a", "creator-nickname")
	if len(aNodes) == 0 {
		fmt.Println("unable to locate author container node")
		return
	}
	author = document.GetFirstChildText(aNodes[0])
	return
}

func extractFicName(d *document.Document) (name string) {
	h1Nodes := d.GetNodesByTag("h1")
	if len(h1Nodes) == 0 {
		fmt.Println("unable to locate fic name container node")
		return
	}
	name = document.GetFirstChildText(h1Nodes[0])
	return
}

func extractChaptersInfo(d *document.Document) (hrefs []string, chapters []fic.Chapter) {
	nodes := d.GetNodesByClass("list-of-fanfic-parts")
	if len(nodes) == 0 {
		fmt.Println("unable to locate index container node")
		return
	}
	ul := nodes[0]
	hrefs = document.FindTagAttributeValues(ul, "a", "href")
	names := extractChapterNames(ul)
	times := extractChapterTimes(ul)
	if len(hrefs) != 2*len(names) || len(hrefs) != 2*len(times) {
		fmt.Printf("Unmatched number of key values in fic index page: %d, %d, %d\n",
			len(hrefs), len(names), len(times))
		return
	}
	for i, name := range names {
		href := hrefs[2*i]
		chapters = append(chapters, fic.Chapter{
			ID:      extractChapterID(href),
			Name:    name,
			Created: times[i],
		})
	}
	return
}

func extractChapterNames(root *html.Node) (names []string) {
	nodes := document.FindByTag(root, "h3")
	for _, node := range nodes {
		name := document.GetFirstChildText(node)
		if name != "" {
			names = append(names, name)
		} else {
			fmt.Printf("strange chapter name container encountered\n")
		}
	}
	return
}

func extractChapterTimes(root *html.Node) (times []time.Time) {
	timesInStrings := document.FindAttributeValues(root, "title")
	for _, str := range timesInStrings {
		t, err := convertToTime(str)
		if err != nil {
			fmt.Printf("Unable to parse string [ %s ] as time: %v\n", str, err)
		}
		times = append(times, t)
	}
	return
}

var ruToEnMonth = map[string]string{
	"января":   "Jan",
	"февраля":  "Feb",
	"марта":    "Mar",
	"апреля":   "Apr",
	"мая":      "May",
	"июня":     "Jun",
	"июля":     "Jul",
	"августа":  "Aug",
	"сентября": "Sep",
	"октября":  "Oct",
	"ноября":   "Nov",
	"декабря":  "Dec",
}

func convertToTime(str string) (t time.Time, err error) {
	split := strings.Split(str, " ")
	if len(split) != 4 {
		err = fmt.Errorf("unknown format")
		return
	}
	split[1] = ruToEnMonth[split[1]]
	t, err = time.Parse("02 Jan 2006, 15:04", strings.Join(split, " "))
	return
}

func extractChapterID(href string) (id string) {
	split := strings.Split(href, "/")
	if len(split) != 4 {
		return
	}
	hashSplit := strings.Split(split[3], "#")
	if len(hashSplit) != 2 {
		return
	}
	id = hashSplit[0]
	return
}
