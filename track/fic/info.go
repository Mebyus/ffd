package fic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/mebyus/ffd/logs"
	"github.com/mebyus/ffd/setting"
)

type Chapter struct {
	ID      string
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
	ID         string
	BaseURL    string
	Location   Location
	Name       string
	Author     string
	Annotation string
	Words      int64
	Finished   bool
	Created    time.Time
	Updated    time.Time
	Chapters   []Chapter
	Check      Check
}

func CountWords(chapters []Chapter) (count int64) {
	for i := range chapters {
		count += chapters[i].Words
	}
	return
}

func UpdatedTime(chapters []Chapter) (t time.Time) {
	if len(chapters) == 0 {
		return
	}
	t = chapters[len(chapters)-1].Created
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

func Get(n int) (f *Info, err error) {
	fics, _, err := Load("")
	if err != nil {
		return
	}
	if n < 1 || n > len(fics) {
		err = fmt.Errorf("fic number = %d exceeds boundaries [%d, %d]", n, 1, len(fics))
		return
	}
	f = &fics[n-1]
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
		logs.Error.Printf("cannot locate working directory: %v\n", err)
	}
	if err != nil || execdir == wdpath {
		err = nil
		originpath = filepath.Join(execdir, setting.TrackPath)
		b, err = ioutil.ReadFile(originpath)
		if err != nil {
			if !os.IsNotExist(err) {
				err = fmt.Errorf("reading fics list: %v", err)
				originpath = ""
				return
			}
			fics = make([]Info, 0)
			logs.Warn.Printf("track file doesn't exist, fic list is treated as empty\n")
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
	originpath = filepath.Join(execdir, setting.TrackPath)
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
	originpath = filepath.Join(wdpath, setting.TrackPath)
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
	originpath = filepath.Join(execdir, setting.TrackPath)
	fics = make([]Info, 0)
	logs.Warn.Printf("track file doesn't exist, fic list is treated as empty\n")
	return
}

func Sort(fics []Info) {
	sort.Slice(fics, func(i, j int) bool {
		return fics[i].BaseURL < fics[j].BaseURL
	})
}

func Remove(fics *[]Info, n int) (err error) {
	if n < 1 || n > len(*fics) {
		err = fmt.Errorf("fic number = %d exceeds boundaries [%d, %d]", n, 1, len(*fics))
		return
	}
	if len(*fics) == 1 {
		fics = &[]Info{}
	} else {
		(*fics)[n-1] = (*fics)[len(*fics)-1]
		*fics = (*fics)[:len(*fics)-1]
		Sort(*fics)
	}
	return
}

func Compare(o, n []Chapter) (diff []Chapter) {
	m := make(map[string]Chapter)
	for _, oc := range o {
		if oc.ID != "" {
			m[oc.ID] = oc
		}
	}
	for _, nc := range n {
		if nc.ID != "" {
			_, exists := m[nc.ID]
			if !exists {
				diff = append(diff, nc)
			}
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
