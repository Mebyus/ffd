package resource

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/mebyus/ffd/cmn"
	"github.com/mebyus/ffd/resource/fiction"
	"github.com/mebyus/ffd/setting"
)

func ParseReader(src io.Reader, resourceID string, format fiction.RenderFormat) (err error) {
	tool, err := ChooseByID(resourceID)
	if err != nil {
		err = fmt.Errorf("choosing tool for %s: %v", resourceID, err)
		return
	}
	book, err := tool.Parse(src)
	if err != nil {
		return err
	}
	err = book.SaveAs(setting.OutDir, "stdin", format)
	if err != nil {
		return err
	}
	return
}

func Parse(dirpath, resourceID string, separate bool, format fiction.RenderFormat) (err error) {
	dir, err := os.Open(dirpath)
	if err != nil {
		return
	}
	defer cmn.SmartClose(dir)
	dirstat, err := dir.Stat()
	if err != nil {
		return
	}
	if !dirstat.IsDir() {
		err = fmt.Errorf("\"%s\" is not a directory", dirpath)
		return
	}
	dirnames, err := dir.Readdirnames(0)
	if err != nil {
		return
	}
	tool, err := ChooseByID(resourceID)
	if err != nil {
		err = fmt.Errorf("choosing tool for %s: %v", resourceID, err)
		return
	}
	if separate {
		err = parseSeparate(dirstat.Name(), dirpath, tool, dirnames, format)
	} else {
		err = parseTogether(dirstat.Name(), dirpath, tool, dirnames, format)
	}
	return
}

func parseTogether(name, dirpath string, tool tools, dirnames []string, format fiction.RenderFormat) (err error) {
	chapters := []fiction.Chapter{}
	for _, partname := range dirnames {
		file, err := os.Open(filepath.Join(dirpath, partname))
		if err != nil {
			return err
		}
		defer cmn.SmartClose(file)
		fstat, err := file.Stat()
		if err != nil {
			return err
		}
		if !fstat.IsDir() {
			book, err := tool.Parse(file)
			if err != nil {
				return err
			}
			chapters = append(chapters, book.Chapters...)
		}
	}
	combinedBook := &fiction.Book{
		Chapters: chapters,
	}
	err = combinedBook.SaveAs(setting.OutDir, name, format)
	if err != nil {
		return err
	}
	return
}

func parseSeparate(name, dirpath string, tool tools, dirnames []string, format fiction.RenderFormat) (err error) {
	outdirpath := filepath.Join(setting.OutDir, name)
	for _, partname := range dirnames {
		file, err := os.Open(filepath.Join(dirpath, partname))
		if err != nil {
			return err
		}
		defer cmn.SmartClose(file)
		fstat, err := file.Stat()
		if err != nil {
			return err
		}
		if !fstat.IsDir() {
			book, parseErr := tool.Parse(file)
			if parseErr != nil {
				fmt.Println(parseErr)
				continue
			}
			saveErr := book.SaveAs(outdirpath, partname, format)
			if saveErr != nil {
				fmt.Println(saveErr)
			}
		}
	}
	return
}
