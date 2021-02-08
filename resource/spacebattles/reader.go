package spacebattles

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/mebyus/ffd/cmn"
	"github.com/mebyus/ffd/planner"
)

const Hostname = "forums.spacebattles.com"

const timeout = 15 * time.Second
const outdir = "out"
const sourcedir = "source"

func (t *sbTools) Download(target string, saveSource bool) {
	fmt.Printf("Analyzing URL\n")
	baseURL, name, err := analyze(target)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("URL is correct. Base part: [ %s ]\n", baseURL)
	fmt.Printf("Fic designation set to [ %s ]\n", name)

	fmt.Printf("Started downloading [ %s ]\n", name)
	err = downloadSync(baseURL, name, saveSource, planner.Client)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Finished downloading [ %s ]\n", name)
	return
	// } else {
	// 	results := make(chan *planner.Result, 10)
	// 	for _, url := range urls {
	// 		task := gettask(url, results)
	// 		planner.Tasks <- task
	// 	}
	// 	for _, url := range urls {
	// 		result := <-results
	// 		if result.Err != nil {
	// 			fmt.Printf("Obtaining result from %s: %v\n", url, result.Err)
	// 			return
	// 		}
	// 		contentStr := parseChapter(result.Content)
	// 		_, err = io.Copy(outfile, strings.NewReader(contentStr))
	// 		if err != nil {
	// 			err = fmt.Errorf("Saving chapter to destination: %v", err)
	// 			return
	// 		}
	// 		closeErr := result.Content.Close()
	// 		if closeErr != nil {
	// 			fmt.Printf("Closing chapter response body: %v\n", closeErr)
	// 		}
	// 	}
	// }
}

func downloadAsync(target string, saveSource bool) (err error) {
	return
}

func downloadSync(baseURL, name string, saveSource bool, client *http.Client) (err error) {
	fmt.Printf("Downloading first page...")
	start := time.Now()
	firstPage, err := cmn.GetBody(readerPageURL(baseURL, 1), client)
	if err != nil {
		fmt.Println()
		return
	}
	fmt.Printf(" [ OK ] %v\n", time.Since(start))
	defer cmn.SmartClose(firstPage)

	err = os.MkdirAll(outdir, 0774)
	if err != nil {
		return
	}
	outpath := filepath.Join(outdir, name+".txt")
	outfile, err := os.Create(outpath)
	if err != nil {
		return
	}
	defer cmn.SmartClose(outfile)
	fmt.Printf("Output file: %s\n", outpath)

	var teeFirstPage io.Reader
	savedir := filepath.Join(sourcedir, name)
	if saveSource {
		err = os.MkdirAll(savedir, 0774)
		if err != nil {
			return
		}
		fmt.Printf("Source files will be saved to: %s\n", savedir)
		fp := filepath.Join(savedir, "1.html")
		sourcefile, err := os.Create(fp)
		if err != nil {
			return err
		}
		defer cmn.SmartClose(sourcefile)
		teeFirstPage = io.TeeReader(firstPage, sourcefile)
	} else {
		teeFirstPage = firstPage
	}

	fmt.Printf("Parsing first page...\n")
	start = time.Now()
	parsedFirstPage, pages, err := parsePiece(teeFirstPage)
	if err != nil {
		return
	}
	parsingDuration := time.Since(start)
	_, err = io.Copy(outfile, parsedFirstPage)
	if err != nil {
		return
	}
	fmt.Printf("First page parsed. Fic contains %d pages total\n", pages)

	filenames := sourceFilenames(pages)
	for i := int64(2); i <= pages; i++ {
		fmt.Printf("Downloading page %3d / %d", i, pages)
		start = time.Now()
		page, err := cmn.GetBody(readerPageURL(baseURL, i), client)
		if err != nil {
			fmt.Println()
			return err
		}
		fmt.Printf("  [ OK ] %v\n", time.Since(start))
		defer cmn.SmartClose(page)

		var teePage io.Reader
		if saveSource {
			fp := filepath.Join(savedir, filenames[i-1])
			sourcefile, err := os.Create(fp)
			if err != nil {
				return err
			}
			defer cmn.SmartClose(sourcefile)
			teePage = io.TeeReader(page, sourcefile)
		} else {
			teePage = page
		}

		start = time.Now()
		parsedPage, _, err := parsePiece(teePage)
		if err != nil {
			return err
		}
		parsingDuration += time.Since(start)
		_, err = io.Copy(outfile, parsedPage)
		if err != nil {
			return err
		}
	}
	fmt.Printf("Parsing %d pages took: %v (%v per page)\n", pages, parsingDuration,
		parsingDuration/time.Duration(pages))
	return
}

func sourceFilenames(maxIndex int64) (filenames []string) {
	maxIndexStr := fmt.Sprintf("%d", maxIndex)
	for i := int64(1); i <= maxIndex; i++ {
		indexStr := fmt.Sprintf("%d", i)
		formatStr := ""
		if len(maxIndexStr)-len(indexStr) > 0 {
			formatStr = strings.Repeat("0", len(maxIndexStr)-len(indexStr)) + "%d.html"
		} else {
			formatStr = "%d.html"
		}
		filenames = append(filenames, fmt.Sprintf(formatStr, i))
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
