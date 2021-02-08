package spacebattles

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/mebyus/ffd/cmn"
	"github.com/mebyus/ffd/planner"
	"golang.org/x/net/html"
)

const Hostname = "forums.spacebattles.com"

const timeout = 15 * time.Second
const outdir = "out"
const sourcedir = "source"

func (t *sbTools) Download(target string, saveSource bool) {
	client := &http.Client{
		Timeout: timeout,
	}
	t.saveSource = saveSource
	t.ficName = getFicName(target)
	urls, err := geturls(target, client)
	if err != nil {
		fmt.Println(err)
		return
	}
	t.chapters = len(urls)
	fmt.Printf("%d pages ( %s ) from %s\n", len(urls), t.ficName, Hostname)
	err = os.MkdirAll(outdir, 0766)
	if err != nil {
		fmt.Println(err)
		return
	}
	outpath := filepath.Join(outdir, t.ficName+".txt")
	outfile, err := os.Create(outpath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer cmn.SmartClose(outfile)
	if saveSource {
		for i, url := range urls {
			fmt.Println(url)
			err = t.getchapter(url, i+1, client, outfile)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("page %3d [OK] ( %s )\n", i+1, t.ficName)
		}
	} else {
		results := make(chan *planner.Result, 10)
		for _, url := range urls {
			task := gettask(url, results)
			planner.Tasks <- task
		}
		for _, url := range urls {
			result := <-results
			if result.Err != nil {
				fmt.Printf("Obtaining result from %s: %v\n", url, result.Err)
				return
			}
			contentStr := parseChapter(result.Content)
			_, err = io.Copy(outfile, strings.NewReader(contentStr))
			if err != nil {
				err = fmt.Errorf("Saving chapter to destination: %v", err)
				return
			}
			closeErr := result.Content.Close()
			if closeErr != nil {
				fmt.Printf("Closing chapter response body: %v\n", closeErr)
			}
		}
	}
}

func geturls(target string, client *http.Client) (urls []string, err error) {
	readerTarget := getReaderURL(target)
	request, err := http.NewRequest("GET", readerTarget, bytes.NewReader([]byte{}))
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
	pageNumbers := parseTableOfContents(response.Body)

	// determine max page number
	var maxNumber int64 = 1
	for _, pageNumber := range pageNumbers {
		number, parseErr := strconv.ParseInt(pageNumber, 10, 64)
		if parseErr == nil && number > maxNumber {
			maxNumber = number
		}
	}

	for i := int64(1); i <= maxNumber; i++ {
		urls = append(urls, readerTarget+convert(i))
	}

	closeErr := response.Body.Close()
	if closeErr != nil {
		fmt.Printf("Closing table of contents response body: %v\n", closeErr)
	}
	return
}

func gettask(url string, ch chan<- *planner.Result) *planner.Task {
	return &planner.Task{
		Label:       "SB",
		URL:         url,
		Destination: ch,
	}
}

func (t *sbTools) getchapter(url string, index int, client *http.Client, destination io.Writer) (err error) {
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
	var contentStr string
	if t.saveSource {
		savedir := filepath.Join(sourcedir, t.ficName)
		err = os.MkdirAll(savedir, 0766)
		if err != nil {
			fmt.Println(err)
			return
		}
		indexStr := fmt.Sprintf("%d", index)
		maxIndexStr := fmt.Sprintf("%d", t.chapters)
		formatStr := ""
		if len(maxIndexStr)-len(indexStr) > 0 {
			formatStr = strings.Repeat("0", len(maxIndexStr)-len(indexStr)) + "%d.html"
		} else {
			formatStr = "%d.html"
		}
		filename := fmt.Sprintf(formatStr, index)
		fp := filepath.Join(savedir, filename)
		var file *os.File
		file, err = os.Create(fp)
		if err != nil {
			fmt.Println(err)
			return
		}
		teeBody := io.TeeReader(response.Body, file)
		contentStr = parseChapter(teeBody)

	} else {
		contentStr = parseChapter(response.Body)
	}
	_, err = io.Copy(destination, strings.NewReader(contentStr))
	if err != nil {
		err = fmt.Errorf("Saving chapter to destination: %v", err)
		return
	}
	closeErr := response.Body.Close()
	if closeErr != nil {
		fmt.Printf("Closing chapter response body: %v\n", closeErr)
	}
	return
}

func parseChapter(page io.Reader) (content string) {
	return
	// tokenizer := html.NewTokenizer(page)
	// depth := 0
	// insideChapter := false
	// insideArticle := false
	// insideThreadmark := false
	// chapterDepth := 0
	// articleDepth := 0
	// insertNewline := false
	// insertSpace := false
	// for {
	// 	tokenType := tokenizer.Next()
	// 	switch tokenType {
	// 	case html.ErrorToken:
	// 		err := tokenizer.Err()
	// 		if err == io.EOF {
	// 			return
	// 		}
	// 		fmt.Printf("Page tokenization: %v\n", err)
	// 		return
	// 	case html.StartTagToken:
	// 		token := tokenizer.Token()
	// 		if token.Data == "article" {
	// 			insideArticle = true
	// 			articleDepth = depth
	// 		} else if insideArticle && !insideChapter && token.Data == "span" && isThreadmark(token.Attr) {
	// 			insideThreadmark = true
	// 		} else if insideArticle && !insideChapter && token.Data == "div" && isPostWrapper(token.Attr) {
	// 			insideChapter = true
	// 			chapterDepth = depth
	// 		} else if insideChapter && (token.Data == "i" || token.Data == "b") {
	// 			content += " "
	// 			insertSpace = true
	// 		}
	// 		depth++
	// 	case html.SelfClosingTagToken:
	// 		token := tokenizer.Token()
	// 		if insideChapter && token.Data == "br" {
	// 			content += "\n"
	// 			insertNewline = true
	// 		}
	// 	case html.EndTagToken:
	// 		depth--
	// 		token := tokenizer.Token()
	// 		if token.Data == "article" && articleDepth == depth {
	// 			insideArticle = false
	// 		} else if insideThreadmark && token.Data == "span" {
	// 			insideThreadmark = false
	// 		} else if insideChapter && token.Data == "div" && chapterDepth == depth {
	// 			insideChapter = false
	// 			content += "\n\n"
	// 		}
	// 	case html.TextToken:
	// 		token := tokenizer.Token()
	// 		if insideThreadmark {
	// 			content += strings.TrimSpace(token.Data) + "\n\n"
	// 		}
	// 		if insideChapter {
	// 			if insertNewline {
	// 				content += strings.TrimSpace(token.Data)
	// 				insertNewline = false
	// 				insertSpace = false
	// 			} else if insertSpace {
	// 				content += strings.TrimSpace(token.Data) + " "
	// 				insertSpace = false
	// 			} else {
	// 				content += strings.TrimSpace(token.Data)
	// 			}
	// 		}
	// 	}
	// }
}

func getFicName(target string) string {
	split := strings.Split(strings.Trim(target, "/"), "/")
	if len(split) == 0 {
		return ""
	}
	name := split[len(split)-1]
	if name == "reader" && len(split) > 1 {
		name = split[len(split)-2]
	}
	split = strings.Split(name, ".")
	if len(split) == 0 {
		return ""
	} else if len(split) == 1 {
		name = split[len(split)-1]
	} else {
		name = split[len(split)-2]
	}
	return strings.ToLower(name)
}

func getReaderURL(target string) string {
	trim := strings.Trim(target, "/")
	if strings.HasSuffix(trim, "/reader") {
		return trim + "/"
	}
	return trim + "/reader/"
}

func convert(pageNumber int64) string {
	if pageNumber == 1 {
		return ""
	}
	return fmt.Sprintf("page-%d", pageNumber)
}

func parseTableOfContents(page io.Reader) []string {
	tokenizer := html.NewTokenizer(page)
	insidePageNav := false
	insideLink := false
	pageNumbers := []string{}
	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			err := tokenizer.Err()
			if err == io.EOF {
				return pageNumbers
			}
			fmt.Printf("Table of contents tokenization: %v\n", err)
			return []string{}
		case html.StartTagToken:
			token := tokenizer.Token()
			if token.Data == "ul" && isPageNav(token.Attr) {
				insidePageNav = true
			} else if insidePageNav && token.Data == "a" {
				insideLink = true
			}
		case html.EndTagToken:
			token := tokenizer.Token()
			if insideLink && token.Data == "a" {
				insideLink = false
			}
			if insidePageNav && token.Data == "ul" {
				return pageNumbers
			}
		case html.TextToken:
			token := tokenizer.Token()
			if insideLink {
				pageNumbers = append(pageNumbers, token.Data)
			}
		}
	}
}

func isPageNav(attrs []html.Attribute) bool {
	for _, attr := range attrs {
		if attr.Key == "class" && attr.Val == "pageNav-main" {
			return true
		}
	}
	return false
}

func isThreadmark(attrs []html.Attribute) bool {
	for _, attr := range attrs {
		if attr.Key == "class" && strings.Contains(attr.Val, "threadmarkLabel") {
			return true
		}
	}
	return false
}

func isPostWrapper(attrs []html.Attribute) bool {
	for _, attr := range attrs {
		if attr.Key == "class" && attr.Val == "bbWrapper" {
			return true
		}
	}
	return false
}
