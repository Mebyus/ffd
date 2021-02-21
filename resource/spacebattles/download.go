package spacebattles

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/mebyus/ffd/cmn"
	"github.com/mebyus/ffd/planner"
	"github.com/mebyus/ffd/resource/fiction"
	"github.com/mebyus/ffd/setting"
)

const Hostname = "forums.spacebattles.com"

func (t *sbTools) Download(target string, saveSource bool) (book *fiction.Book, err error) {
	fmt.Printf("Analyzing URL\n")
	baseURL, name, err := analyze(target)
	if err != nil {
		return
	}
	fmt.Printf("URL is correct. Base part: [ %s ]\n", baseURL)
	fmt.Printf("Fic designation set to [ %s ]\n", name)

	fmt.Printf("Started downloading [ %s ]\n", name)
	book, err = downloadSync(baseURL, name, saveSource, planner.Client)
	if err != nil {
		return
	}
	fmt.Printf("Finished downloading [ %s ]\n", name)
	return
}

func downloadAsync(target string, saveSource bool) (err error) {
	return
}

func downloadSync(baseURL, name string, saveSource bool, client *http.Client) (book *fiction.Book, err error) {
	fmt.Printf("Downloading first page...")
	start := time.Now()
	firstPage, err := cmn.GetBody(readerPageURL(baseURL, 1), client)
	if err != nil {
		fmt.Println()
		return
	}
	fmt.Printf(" [ OK ] %v\n", time.Since(start))
	defer cmn.SmartClose(firstPage)

	var teeFirstPage io.Reader
	saveSourceDir := filepath.Join(setting.SourceSaveDir, name)
	if saveSource {
		err = os.MkdirAll(saveSourceDir, 0774)
		if err != nil {
			return
		}
		fmt.Printf("Source files will be saved to: %s\n", saveSourceDir)
		fp := filepath.Join(saveSourceDir, "1.html")
		sourcefile, err := os.Create(fp)
		if err != nil {
			return nil, err
		}
		defer cmn.SmartClose(sourcefile)
		teeFirstPage = io.TeeReader(firstPage, sourcefile)
	} else {
		teeFirstPage = firstPage
	}

	fmt.Printf("Parsing first page...\n")
	start = time.Now()
	firstChapters, pages, err := parsePiece(teeFirstPage)
	if err != nil {
		return
	}
	parsingDuration := time.Since(start)
	fmt.Printf("First page parsed. Fic contains %d pages total\n", pages)

	filenames := cmn.GenerateFilenames(pages, "html")
	chapters := []fiction.Chapter{}
	chapters = append(chapters, firstChapters...)
	for i := int64(2); i <= pages; i++ {
		fmt.Printf("Downloading page %3d / %d", i, pages)
		start = time.Now()
		page, err := cmn.GetBody(readerPageURL(baseURL, i), client)
		if err != nil {
			fmt.Println()
			return nil, err
		}
		fmt.Printf("  [ OK ] %v\n", time.Since(start))
		defer cmn.SmartClose(page)

		var teePage io.Reader
		if saveSource {
			fp := filepath.Join(saveSourceDir, filenames[i-1])
			sourcefile, err := os.Create(fp)
			if err != nil {
				return nil, err
			}
			defer cmn.SmartClose(sourcefile)
			teePage = io.TeeReader(page, sourcefile)
		} else {
			teePage = page
		}

		start = time.Now()
		pieceChapters, _, err := parsePiece(teePage)
		if err != nil {
			return nil, err
		}
		parsingDuration += time.Since(start)
		chapters = append(chapters, pieceChapters...)
	}
	fmt.Printf("Parsing %d pages took: %v (%v per page)\n", pages, parsingDuration,
		parsingDuration/time.Duration(pages))
	book = &fiction.Book{
		Title:    name,
		Chapters: chapters,
	}
	return
}
