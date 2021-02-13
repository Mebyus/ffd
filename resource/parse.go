package resource

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mebyus/ffd/cmn"
	"github.com/mebyus/ffd/logs"
	"github.com/mebyus/ffd/setting"
)

func Parse(dirpath, resourceID string, separate bool) (err error) {
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
		err = parseSeparate(dirstat.Name(), dirpath, tool, dirnames)
	} else {
		err = parseTogether(dirstat.Name(), dirpath, tool, dirnames)
	}
	return
}

func parseTogether(name, dirpath string, tool tools, dirnames []string) (err error) {
	err = os.MkdirAll(setting.OutDir, 0766)
	if err != nil {
		return
	}
	outpath := filepath.Join(setting.OutDir, name+".txt")
	outfile, err := os.Create(outpath)
	if err != nil {
		return err
	}
	defer cmn.SmartClose(outfile)
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
			err := tool.Parse(file, outfile)
			if err != nil {
				return err
			}
		}
	}
	return
}

func parseSeparate(name, dirpath string, tool tools, dirnames []string) (err error) {
	outdirpath := filepath.Join(setting.OutDir, name)
	err = os.MkdirAll(outdirpath, 0766)
	if err != nil {
		return
	}
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
			outname := strings.TrimSuffix(name, "html") + "txt"
			outpath := filepath.Join(outdirpath, outname)
			outfile, err := os.Create(outpath)
			if err != nil {
				return err
			}
			defer cmn.SmartClose(outfile)
			parseErr := tool.Parse(file, outfile)
			if parseErr != nil {
				logs.Error.Println(parseErr)
			}
		}
	}
	return
}
