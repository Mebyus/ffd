package royalroad

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/mebyus/ffd/cmn"
	"github.com/mebyus/ffd/planner"
	"github.com/mebyus/ffd/resource/internal"
	"github.com/mebyus/ffd/setting"
)

const Hostname = "www.royalroad.com"
const pause = time.Second

func (t *rrTools) Download(target string, saveSource bool) (book *internal.Book, err error) {
	fmt.Printf("Analyzing URL\n")
	baseURL, name, err := analyze(target)
	if err != nil {
		return
	}
	fmt.Printf("URL is correct. Base part: [ %s ]\n", baseURL)
	fmt.Printf("Fic designation set to [ %s ]\n", name)

	fmt.Printf("Started downloading [ %s ]\n", name)
	err = downloadSync(baseURL, name, saveSource, planner.Client)
	if err != nil {
		return
	}
	fmt.Printf("Finished downloading [ %s ]\n", name)
	return
}

func downloadSync(baseURL, name string, saveSource bool, client *http.Client) (err error) {
	urls, err := getChapterURLs(baseURL, client)
	if err != nil {
		return
	}

	err = os.MkdirAll(setting.OutDir, 0774)
	if err != nil {
		return
	}
	outpath := filepath.Join(setting.OutDir, name+".txt")
	outfile, err := os.Create(outpath)
	if err != nil {
		return
	}
	defer cmn.SmartClose(outfile)
	fmt.Printf("Output file: %s\n", outpath)

	savedir := filepath.Join(setting.SourceSaveDir, name)
	if saveSource {
		err = os.MkdirAll(savedir, 0774)
		if err != nil {
			return
		}
		fmt.Printf("Source files will be saved to: %s\n", savedir)
	}

	parsingDuration := time.Duration(0)
	pages := int64(len(urls))
	filenames := cmn.GenerateFilenames(pages, "html")
	for i, url := range urls {
		fmt.Printf("Downloading chapter %3d / %d", i+1, pages)
		start := time.Now()
		page, err := cmn.GetBody(url, client)
		if err != nil {
			fmt.Println()
			return err
		}
		fmt.Printf("  [ OK ] %v\n", time.Since(start))
		defer cmn.SmartClose(page)

		var teePage io.Reader
		if saveSource {
			fp := filepath.Join(savedir, filenames[i])
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
		parsedPage, err := parseChapter(teePage)
		if err != nil {
			return err
		}
		parsingDuration += time.Since(start)
		_, err = io.Copy(outfile, parsedPage)
		if err != nil {
			return err
		}

		// wait to not spook server DOS (or whatever) protection
		time.Sleep(pause)
	}
	fmt.Printf("Parsing %d pages took: %v (%v per page)\n", pages, parsingDuration,
		parsingDuration/time.Duration(pages))
	return
}
