package resource

import (
	"fmt"
	"io"

	"github.com/mebyus/ffd/resource/fanfiction"
	"github.com/mebyus/ffd/resource/royalroad"
	"github.com/mebyus/ffd/resource/spacebattles"
	"github.com/mebyus/ffd/track/fic"
)

type tools interface {
	Download(target string, saveSource bool)
	Check(target string) []fic.Chapter
	Parse(src io.Reader, dst io.Writer) error
}

func Choose(hostname string) (t tools, err error) {
	switch hostname {
	case spacebattles.Hostname:
		t = spacebattles.NewTools()
	case "forums.sufficientvelocity.com":
		t = spacebattles.NewTools()
	case "forum.questionablequesting.com":
		err = notImplemented(hostname)
	case "www.fanfiction.net":
		t = fanfiction.NewTools()
	case "archiveofourown.org":
		err = notImplemented(hostname)
	case royalroad.Hostname:
		t = royalroad.NewTools()
	default:
		err = notImplemented(hostname)
	}
	return
}

func notImplemented(hostname string) error {
	return fmt.Errorf("resource ( %s ) not implemented", hostname)
}
