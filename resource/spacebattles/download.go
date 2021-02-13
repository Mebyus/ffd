package spacebattles

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/mebyus/ffd/cmn"
	"github.com/mebyus/ffd/logs"
	"github.com/mebyus/ffd/planner"
	"github.com/mebyus/ffd/setting"
)

const Hostname = "forums.spacebattles.com"

func (t *sbTools) Download(target string, saveSource bool) {
	fmt.Printf("Analyzing URL\n")
	baseURL, name, err := analyze(target)
	if err != nil {
		logs.Error.Println(err)
		return
	}
	fmt.Printf("URL is correct. Base part: [ %s ]\n", baseURL)
	fmt.Printf("Fic designation set to [ %s ]\n", name)

	fmt.Printf("Started downloading [ %s ]\n", name)
	err = downloadSync(baseURL, name, saveSource, planner.Client)
	if err != nil {
		logs.Error.Println(err)
		return
	}
	fmt.Printf("Finished downloading [ %s ]\n", name)
	return
}

func downloadAsync(target string, saveSource bool) (err error) {
	return
}

func downloadSync(baseURL, name string, saveSource bool, client *http.Client) (err error) {
	logs.Info.Printf("Downloading first page...")
	start := time.Now()
	firstPage, err := cmn.GetBody(readerPageURL(baseURL, 1), client)
	if err != nil {
		return
	}
	logs.Info.Printf(" [ OK ] %v\n", time.Since(start))
	defer cmn.SmartClose(firstPage)

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
	logs.Info.Printf("Output file: %s\n", outpath)

	var teeFirstPage io.Reader
	savedir := filepath.Join(setting.SourceSaveDir, name)
	if saveSource {
		err = os.MkdirAll(savedir, 0774)
		if err != nil {
			return
		}
		logs.Info.Printf("Source files will be saved to: %s\n", savedir)
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

	logs.Info.Printf("Parsing first page...\n")
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
	logs.Info.Printf("First page parsed. Fic contains %d pages total\n", pages)

	filenames := cmn.GenerateFilenames(pages, "html")
	for i := int64(2); i <= pages; i++ {
		fmt.Printf("Downloading page %3d / %d", i, pages)
		start = time.Now()
		page, err := cmn.GetBody(readerPageURL(baseURL, i), client)
		if err != nil {
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
	logs.Info.Printf("Parsing %d pages took: %v (%v per page)\n", pages, parsingDuration,
		parsingDuration/time.Duration(pages))
	return
}

func gettask(url string, ch chan<- *planner.Result) *planner.Task {
	return &planner.Task{
		Label:       "SB",
		URL:         url,
		Destination: ch,
	}
}
