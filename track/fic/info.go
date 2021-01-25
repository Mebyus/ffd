package fic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type Chapter struct {
	Name  string
	Date  time.Time
	Words int64
}

type Check struct {
	Date        time.Time
	NewChapters []Chapter
}

type Info struct {
	URL        string
	Name       string
	Words      int64
	Suppressed bool
	Created    time.Time
	Chapters   []Chapter
	Check      Check
}

func CountWords(cs []Chapter) (count int64) {
	for i := range cs {
		count += cs[i].Words
	}
	return
}

func Save(fics []Info) (err error) {
	b, err := json.MarshalIndent(fics, "", "    ")
	if err != nil {
		err = fmt.Errorf("encoding fics list: %v", err)
		return
	}
	err = ioutil.WriteFile("track.json", b, 0333)
	if err != nil {
		err = fmt.Errorf("saving fics list: %v", err)
		return
	}
	return
}

func Load() (fics []Info, err error) {
	fics = make([]Info, 0)
	b, err := ioutil.ReadFile("track.json")
	if err != nil {
		fmt.Printf("reading fics list: %v\nlist is treated as empty\n", err)
		err = nil
		return
	}
	err = json.Unmarshal(b, &fics)
	if err != nil {
		err = fmt.Errorf("unmarshalling fics list: %v", err)
		return
	}
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
		if fics[i].URL == target {
			index = i
			return
		}
	}
	return
}
