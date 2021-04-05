package spacebattles

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/mebyus/ffd/cmn"
	"github.com/mebyus/ffd/document"
	"github.com/mebyus/ffd/planner"
	"github.com/mebyus/ffd/track/fic"
	"golang.org/x/net/html"
)

func (t *sbTools) Check(target string) (info *fic.Info, err error) {
	baseURL, _, id, err := analyze(target)
	if err != nil {
		return
	}

	fmt.Printf("Downloading index page...")
	start := time.Now()
	indexPage, err := cmn.GetBody(indexPageURL(baseURL), planner.Client)
	if err != nil {
		fmt.Printf("\n")
		return
	}
	fmt.Printf(" [ OK ] %v\n", time.Since(start))
	defer cmn.SmartClose(indexPage)

	info, err = parseThreadmarksPage(indexPage)
	if err != nil {
		fmt.Println(err)
		return
	}
	info.ID = id
	info.BaseURL = baseURL
	info.Location = fic.SpaceBattles
	return
}

func parseThreadmarksPage(source io.Reader) (f *fic.Info, err error) {
	n, err := html.Parse(source)
	if err != nil {
		return
	}
	d := document.FromNode(n)
	name := extractFicName(d)
	author := extractAuthor(d)
	annotation := extractAnnotation(d)
	created := extractCreatedTime(d)
	chapters := extractThreadmarkList(d)
	f = &fic.Info{
		Name:       name,
		Created:    created,
		Annotation: annotation,
		Author:     author,
		Chapters:   chapters,
	}
	return
}

func extractThreadmarkList(d *document.Document) (list []fic.Chapter) {
	containers := d.GetNodesByClass("structItemContainer")
	if len(containers) == 0 {
		fmt.Println("unable to locate list container")
		return
	} else if len(containers) > 1 {
		fmt.Println("located several potential list containers")
	}
	datetimes := document.FindAttributeValues(containers[0], "datetime")
	listItems := d.GetNodesByClass("structItem--threadmark")
	j := 0
	for _, node := range listItems {
		if j == len(datetimes) {
			fmt.Printf("array of dates (%d) have been exhausted\n", len(datetimes))
			break
		}
		row := document.FindNonSpaceTexts(node)
		if len(row) != 4 && len(row) != 5 {
			// Skip row with dots sign (...)
			continue
		} else if len(row) == 5 {
			row = row[1:]
		}
		link := document.FindFirstByTag(node, "a")
		created, err := time.Parse("2006-01-02T15:04:05-0700", datetimes[j])
		if err != nil {
			fmt.Printf("unable to parse chapter creation time [ %s ]: %v\n", datetimes[j], err)
		}
		name := row[0]
		words := convertWordCount(row[2])
		list = append(list, fic.Chapter{
			ID:      exctractChapterID(link),
			Name:    name,
			Words:   words,
			Created: created,
		})
		j++
	}
	return
}

func exctractChapterID(n *html.Node) (id string) {
	val := document.GetAttributeValue(n, "href")
	if val == "" {
		return
	}
	split := strings.Split(val, "#")
	if len(split) != 2 {
		return
	}
	id = strings.TrimPrefix(split[1], "post-")
	return
}

func extractAnnotation(d *document.Document) (annotation string) {
	nodes := d.GetNodesByClass("bbWrapper")
	if len(nodes) == 0 {
		fmt.Println("unable to locate annotation container node")
		return
	}
	annotation = strings.Join(document.FindNonSpaceTexts(nodes[0]), " ")
	return
}

func extractAuthor(d *document.Document) (author string) {
	nodes := d.GetNodesByClass("username")
	if len(nodes) == 0 {
		fmt.Println("unable to locate author username node")
		return
	}
	author = document.FindFirstNonSpaceText(nodes[0])
	return
}

func extractFicName(d *document.Document) (name string) {
	nodes := d.GetNodesByClass("threadmarkListingHeader-name")
	if len(nodes) == 0 {
		fmt.Println("unable to locate name node")
		return
	} else if len(nodes) > 1 {
		fmt.Println("located several potential name nodes")
	}
	name = document.FindFirstNonSpaceText(nodes[0])
	return
}

func extractCreatedTime(d *document.Document) (t time.Time) {
	nodes := d.GetNodesByClass("threadmarkListingHeader-stats")
	if len(nodes) == 0 {
		fmt.Println("unable to locate stats node")
		return
	} else if len(nodes) > 1 {
		fmt.Println("located several potential stats nodes")
	}
	datetimes := document.FindAttributeValues(nodes[0], "datetime")
	if len(datetimes) == 0 {
		fmt.Println("unable to locate creation time node")
		return
	} else if len(datetimes) > 1 {
		fmt.Println("located several potential creation time nodes")
	}
	t, err := time.Parse("2006-01-02T15:04:05-0700", datetimes[0])
	if err != nil {
		fmt.Printf("unable to parse fic creation time [ %s ]: %v\n", datetimes[0], err)
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
