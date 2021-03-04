package resource

import (
	"fmt"
	"io"
	"net/url"
	"strings"

	"github.com/mebyus/ffd/resource/archiveofourown"
	"github.com/mebyus/ffd/resource/fanfiction"
	"github.com/mebyus/ffd/resource/ficbook"
	"github.com/mebyus/ffd/resource/fiction"
	"github.com/mebyus/ffd/resource/royalroad"
	"github.com/mebyus/ffd/resource/samlib"
	"github.com/mebyus/ffd/resource/spacebattles"
	"github.com/mebyus/ffd/resource/webnovel"
	"github.com/mebyus/ffd/track/fic"
)

type tools interface {
	Download(target string, saveSource bool) (book *fiction.Book, err error)
	Check(target string) (info *fic.Info, err error)
	Parse(src io.Reader) (book *fiction.Book, err error)
}

func GetLocationForTarget(target string) (location fic.Location, err error) {
	u, err := url.Parse(target)
	if err != nil {
		return
	}
	hostname := u.Hostname()
	location, err = GetLocationForHostname(hostname)
	return
}

func GetLocationForHostname(hostname string) (location fic.Location, err error) {
	switch hostname {
	case ficbook.Hostname:
		location = fic.FicBook
	case webnovel.Hostname:
		location = fic.WebNovel
	case spacebattles.Hostname:
		location = fic.SpaceBattles
	case "forums.sufficientvelocity.com":
		location = fic.SufficientVelocity
	case "forum.questionablequesting.com":
		location = fic.QuestionableQuesting
	case fanfiction.Hostname:
		location = fic.FanFiction
	case archiveofourown.Hostname:
		location = fic.ArchiveOfOurOwn
	case royalroad.Hostname:
		location = fic.RoyalRoad
	default:
		err = unknown(hostname)
	}
	return
}

func ChooseByTarget(target string) (t tools, err error) {
	u, err := url.Parse(target)
	if err != nil {
		return
	}
	hostname := u.Hostname()
	t, err = ChooseByHostname(hostname)
	return
}

func ChooseByID(resourceID string) (t tools, err error) {
	if len(resourceID) > 3 {
		// resource id is a hostname
		t, err = ChooseByHostname(resourceID)
	} else {
		// resource id is a location
		t, err = ChooseByLocation(resourceID)
	}
	return
}

func ChooseByHostname(hostname string) (t tools, err error) {
	switch hostname {
	case ficbook.Hostname:
		t = ficbook.NewTools()
	case webnovel.Hostname:
		t = webnovel.NewTools()
	case samlib.Hostname:
		t = samlib.NewTools()
	case spacebattles.Hostname:
		t = spacebattles.NewTools()
	case "forums.sufficientvelocity.com":
		t = spacebattles.NewTools()
	case "forum.questionablequesting.com":
		err = notImplemented(hostname)
	case fanfiction.Hostname:
		t = fanfiction.NewTools()
	case archiveofourown.Hostname:
		t = archiveofourown.NewTools()
	case royalroad.Hostname:
		t = royalroad.NewTools()
	default:
		err = unknown(hostname)
	}
	return
}

func ChooseByLocation(location string) (t tools, err error) {
	loc := strings.ToUpper(location)
	switch fic.Location(loc) {
	case fic.FicBook:
		t = ficbook.NewTools()
	case fic.WebNovel:
		t = webnovel.NewTools()
	case fic.SpaceBattles:
		t = spacebattles.NewTools()
	case fic.SufficientVelocity:
		t = spacebattles.NewTools()
	case fic.QuestionableQuesting:
		err = notImplemented(loc)
	case fic.FanFiction:
		t = fanfiction.NewTools()
	case fic.ArchiveOfOurOwn:
		t = archiveofourown.NewTools()
	case fic.RoyalRoad:
		t = royalroad.NewTools()
	default:
		err = unknown(loc)
	}
	return
}

func notImplemented(resourceID string) error {
	return fmt.Errorf("resource [ %s ] not implemented", resourceID)
}

func unknown(resourceID string) error {
	return fmt.Errorf("unknown resource [ %s ]", resourceID)
}
