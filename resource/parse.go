package resource

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mebyus/ffd/cmn"
)

func Parse(dirpath, hostname string, separate bool) (err error) {
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
	if separate {
		err = parseSeparate(dirstat.Name(), dirpath, hostname, dirnames)
	} else {
		err = parseTogether(dirstat.Name(), dirpath, hostname, dirnames)
	}
	return
}

func parseTogether(name, dirpath, hostname string, dirnames []string) (err error) {
	t, err := Choose(hostname)
	if err != nil {
		return
	}
	outpath := filepath.Join("out", name+".txt")
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
			err := t.Parse(file, outfile)
			if err != nil {
				return err
			}
		}
	}
	return
}

func parseSeparate(name, dirpath, hostname string, dirnames []string) (err error) {
	t, err := Choose(hostname)
	if err != nil {
		return
	}
	outdirpath := filepath.Join("out", name)
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
			parseErr := t.Parse(file, outfile)
			if parseErr != nil {
				fmt.Println(parseErr)
			}
		}
	}
	return
}
