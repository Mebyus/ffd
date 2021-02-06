package fic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type Chapter struct {
	Name    string
	Created time.Time
	Words   int64
}

type Check struct {
	Suppressed  bool
	Time        time.Time
	Words       int64
	NewChapters []Chapter
}

type Info struct {
	BaseURL  string
	Location Location
	Name     string
	Author   string
	Words    int64
	Finished bool
	Created  time.Time
	Updated  time.Time
	Chapters []Chapter
	Check    Check
}

func CountWords(cs []Chapter) (count int64) {
	for i := range cs {
		count += cs[i].Words
	}
	return
}

func Save(path string, fics []Info) (err error) {
	b, err := json.MarshalIndent(fics, "", "    ")
	if err != nil {
		err = fmt.Errorf("encoding fics list: %v", err)
		return
	}
	err = ioutil.WriteFile(path, b, 0664)
	if err != nil {
		err = fmt.Errorf("saving fics list: %v", err)
		return
	}
	return
}

func Load(path string) (fics []Info, originpath string, err error) {
	var b []byte
	if path != "" {
		b, err = ioutil.ReadFile(path)
		if err != nil {
			err = fmt.Errorf("reading fics list: %v", err)
			return
		}
		fics = make([]Info, 0)
		err = json.Unmarshal(b, &fics)
		if err != nil {
			err = fmt.Errorf("unmarshalling fics list: %v", err)
			fics = nil
			return
		}
		originpath = path
		return
	}
	execpath, err := os.Executable()
	if err != nil {
		err = fmt.Errorf("cannot locate executable: %v", err)
		return
	}
	execdir := filepath.Dir(execpath)
	wdpath, err := os.Getwd()
	if err != nil {
		fmt.Printf("cannot locate working directory: %v\n", err)
	}
	if err != nil || execdir == wdpath {
		err = nil
		originpath = filepath.Join(execdir, "track.json")
		b, err = ioutil.ReadFile(originpath)
		if err != nil {
			if !os.IsNotExist(err) {
				err = fmt.Errorf("reading fics list: %v", err)
				originpath = ""
				return
			}
			fics = make([]Info, 0)
			fmt.Printf("track file doesn't exist, fic list is treated as empty\n")
			return
		}
		fics = make([]Info, 0)
		err = json.Unmarshal(b, &fics)
		if err != nil {
			err = fmt.Errorf("unmarshalling fics list: %v", err)
			originpath = ""
			fics = nil
			return
		}
		return
	}
	originpath = filepath.Join(execdir, "track.json")
	b, err = ioutil.ReadFile(originpath)
	if err == nil {
		fics = make([]Info, 0)
		err = json.Unmarshal(b, &fics)
		if err != nil {
			err = fmt.Errorf("unmarshalling fics list: %v", err)
			originpath = ""
			fics = nil
			return
		}
		return

	}
	if !os.IsNotExist(err) {
		err = fmt.Errorf("reading fics list: %v", err)
		originpath = ""
		return
	}
	originpath = filepath.Join(wdpath, "track.json")
	b, err = ioutil.ReadFile(originpath)
	if err == nil {
		fics = make([]Info, 0)
		err = json.Unmarshal(b, &fics)
		if err != nil {
			err = fmt.Errorf("unmarshalling fics list: %v", err)
			originpath = ""
			fics = nil
			return
		}
		return

	}
	if !os.IsNotExist(err) {
		err = fmt.Errorf("reading fics list: %v", err)
		originpath = ""
		return
	}
	err = nil
	originpath = filepath.Join(execdir, "track.json")
	fics = make([]Info, 0)
	fmt.Printf("track file doesn't exist, fic list is treated as empty\n")
	return
}

func Compare(o, n []Chapter) (diff []Chapter) {
	m := make(map[string]Chapter)
	for _, oc := range o {
		m[oc.Name] = oc
	}
	for _, nc := range n {
		ec, exists := m[nc.Name]
		if !exists {
			diff = append(diff, nc)
		} else if ec.Words != nc.Words {
			diff = append(diff, nc)
		}
	}
	return
}

func Find(fics []Info, target string) (index int) {
	index = -1
	for i := range fics {
		if fics[i].BaseURL == target {
			index = i
			return
		}
	}
	return
}
