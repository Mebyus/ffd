package setting

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

const (
	defOutDir        = "out"
	defSourceSaveDir = "source"
	defTrackPath     = "track.json"
	defConfigPath    = "config.json"
	defClientTimeout = 15 * time.Second
)

var (
	OutDir        = defOutDir
	SourceSaveDir = defSourceSaveDir
	ClientTimeout = defClientTimeout
	TrackPath     = defTrackPath
)

type settings struct {
	OutDir        string
	SourceSaveDir string
	ClientTimeout time.Duration
}

func load() (s settings, useDefaults bool, err error) {
	b, err := ioutil.ReadFile(defConfigPath)
	if err != nil {
		fmt.Printf("couldn't read config file: %v\n", err)
		err = saveDefault()
		useDefaults = true
		return
	}
	err = json.Unmarshal(b, &s)
	if err != nil {
		return
	}
	return
}

func Load() {
	s, useDefaults, err := load()
	if err != nil {
		fmt.Println(err)
	}
	if useDefaults {
		fmt.Println("All settings have been set to default")
		fmt.Println()
		return
	}

	insertNewline := false
	if s.OutDir != "" {
		OutDir = s.OutDir
	} else {
		insertNewline = true
		fmt.Printf("Output directory set to default: %s\n", defOutDir)
	}
	if s.SourceSaveDir != "" {
		SourceSaveDir = s.SourceSaveDir
	} else {
		insertNewline = true
		fmt.Printf("Source saving directory set to default: %s\n", defSourceSaveDir)
	}
	if s.ClientTimeout != 0 {
		ClientTimeout = s.ClientTimeout
	} else {
		insertNewline = true
		fmt.Printf("Client timeout set to default: %v\n", defClientTimeout)
	}
	if insertNewline {
		fmt.Println()
	}
}

func saveDefault() (err error) {
	fmt.Printf("creating new config file [ %s ]\n", defConfigPath)
	s := settings{
		OutDir:        defOutDir,
		SourceSaveDir: defSourceSaveDir,
		ClientTimeout: defClientTimeout,
	}
	b, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		return
	}
	err = ioutil.WriteFile(defConfigPath, b, 0664)
	if err != nil {
		return
	}
	return
}
