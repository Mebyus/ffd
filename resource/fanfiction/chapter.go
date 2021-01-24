package fanfiction

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/net/html"
)

const timeout = 15.0 * time.Second
const outdir = "out"

func (t *ffTools) Download(target string, saveSource bool) {
	client := &http.Client{
		Timeout: timeout,
	}
	urls, err := geturls(target, client)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = os.MkdirAll(outdir, 0766)
	if err != nil {
		fmt.Println(err)
		return
	}
	outpath := filepath.Join(outdir, getFicName(target)+".txt")
	outfile, err := os.Create(outpath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer closeFile(outfile)
	for _, url := range urls {
		err = getchapter(url, client, outfile)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func getFicName(target string) string {
	split := strings.Split(target, "/")
	name := split[len(split)-1]
	return strings.ToLower(name)
}

func geturls(tocurl string, client *http.Client) (urls []string, err error) {
	// request, err := http.NewRequest("GET", tocurl, bytes.NewReader([]byte{}))
	// if err != nil {
	// 	err = fmt.Errorf("Composing table of contents request: %v", err)
	// 	return
	// }
	// response, err := client.Do(request)
	// if err != nil {
	// 	err = fmt.Errorf("Doing table of contents request: %v", err)
	// 	return
	// }
	// if response.StatusCode != http.StatusOK {
	// 	err = fmt.Errorf("Table of contents response: %s", response.Status)
	// 	return
	// }
	// urls = parseTableOfContents(response.Body)
	urls = []string{"https://www.fanfiction.net/s/12735662/1/Pinky-and-the-Brain"}
	// closeErr := response.Body.Close()
	// if closeErr != nil {
	// 	fmt.Printf("Closing table of contents response body: %v\n", closeErr)
	// }
	return
}

func getchapter(url string, client *http.Client, destination io.Writer) (err error) {
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
	// contentStr := parseChapter(response.Body)
	// _, err = io.Copy(destination, strings.NewReader(contentStr))
	_, err = io.Copy(destination, response.Body)
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
	tokenizer := html.NewTokenizer(page)
	depth := 0
	insideChapter := false
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
			if !insideChapter && token.Data == "div" {
				insideChapter = true
				chapterDepth = depth
			}
			depth++
		case html.EndTagToken:
			depth--
			token := tokenizer.Token()
			if insideChapter && token.Data == "div" && chapterDepth == depth {
				return
			}
		case html.TextToken:
			if insideChapter {
				token := tokenizer.Token()
				content += token.Data + "\n\n"
			}
		}
	}
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		fmt.Println(err)
	}
	return
}
