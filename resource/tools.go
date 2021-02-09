package resource

import (
	"fmt"
	"io"
	"net/url"
	"strings"

	"github.com/mebyus/ffd/resource/fanfiction"
	"github.com/mebyus/ffd/resource/royalroad"
	"github.com/mebyus/ffd/resource/spacebattles"
	"github.com/mebyus/ffd/track/fic"
)

type tools interface {
	Download(target string, saveSource bool)
	Check(target string) *fic.Info
	Parse(src io.Reader, dst io.Writer) error
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
	case spacebattles.Hostname:
		location = fic.SpaceBattles
	case "forums.sufficientvelocity.com":
		location = fic.SufficientVelocity
	case "forum.questionablequesting.com":
		location = fic.QuestionableQuesting
	case "www.fanfiction.net":
		location = fic.FanFiction
	case "archiveofourown.org":
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
		err = unknown(hostname)
	}
	return
}

func ChooseByLocation(location string) (t tools, err error) {
	loc := strings.ToUpper(location)
	switch fic.Location(loc) {
	case fic.SpaceBattles:
		t = spacebattles.NewTools()
	case fic.SufficientVelocity:
		t = spacebattles.NewTools()
	case fic.QuestionableQuesting:
		err = notImplemented(loc)
	case fic.FanFiction:
		t = fanfiction.NewTools()
	case fic.ArchiveOfOurOwn:
		err = notImplemented(loc)
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
