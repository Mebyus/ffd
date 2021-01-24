package royalroad

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/mebyus/ffd/cmn"
)

const timeout = 15.0 * time.Second
const Hostname = "www.royalroad.com"
const sourcedir = "source"

func (t *rrTools) Download(target string, saveSource bool) {
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
	for i, url := range urls {
		fmt.Println(url)
		err = t.getchapter(url, i+1, client, outfile)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("page %3d [OK] ( %s )\n", i+1, t.ficName)
	}
	// resOut := make(chan result, len(urls))
	// for index, url := range urls {
	// 	go getchapter(url, index, resOut, client)
	// }

	// inRes := make(chan result, len(urls))
	// successRes := make(chan error, 1)
	// go gatherall(len(urls), 1, inRes, successRes)

	// for range urls {
	// 	newResult := <-resOut
	// 	if newResult.err != nil {
	// 		fmt.Printf("%d [%s]: %v", newResult.index, newResult.url, newResult.err)
	// 		return
	// 	}
	// 	inRes <- newResult
	// }
	// err = <-successRes
	// if err != nil {
	// 	fmt.Println(err)
	// }

	return
}

func getFicName(target string) string {
	split := strings.Split(strings.Trim(target, "/"), "/")
	if len(split) == 0 {
		return ""
	}
	name := split[len(split)-1]
	return strings.ToLower(name)
}

func (t *rrTools) getchapter(url string, index int, client *http.Client, destination io.Writer) (err error) {
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
