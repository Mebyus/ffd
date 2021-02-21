package logs

import (
	"io/ioutil"
	"log"
	"os"
)

// Level represents logging level
type Level uint

const (
	// I - Info level
	I Level = iota + 3
	// W - Warning level
	W
	// E - Error level
	E
)

var (
	level Level = W

	Info  *log.Logger = log.New(ioutil.Discard, "", 0)
	Warn  *log.Logger = log.New(os.Stdout, "Warning: ", 0)
	Error *log.Logger = log.New(os.Stderr, "Error: ", 0)
)

func CurrentLevel() Level {
	return level
}

func (l Level) String() string {
	switch l {
	case I:
		return "Info"
	case W:
		return "Warn"
	case E:
		return "Error"
	default:
		return "Unknown"
	}
}

// Init sets logging level for loggers.
// Default level is Warning.
// To clarify: loggers are already initialized to Warning level,
// you only need to use Init to set other levels of logging
func Init(l Level) {
	level = l
	switch l {
	case I:
		Info = log.New(os.Stdout, "", 0)
	case E:
		Warn = log.New(ioutil.Discard, "Warning: ", 0)
	}
	return
}
