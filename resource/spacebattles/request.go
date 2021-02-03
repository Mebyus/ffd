package spacebattles

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/mebyus/ffd/cmn"
)

type HTTPRequestAccess struct {
	client *http.Client
	target string
}

func NewHTTPRequestAccess(rawtarget string) (a *HTTPRequestAccess, err error) {
	target, err := baseURL(rawtarget)
	if err != nil {
		return
	}
	a = &HTTPRequestAccess{
		client: &http.Client{
			Timeout: timeout,
		},
		target: target,
	}
	return
}

func baseURL(url string) (base string, err error) {
	split := strings.Split(url, "/")
	if len(split) == 0 {
		err = fmt.Errorf("incorrect url")
		return
	}
	parts := 3
	if split[0] == "https:" || split[0] == "http" {
		parts = 5
	}
	if len(split) < parts {
		err = fmt.Errorf("incorrect url")
		return
	}
	base = strings.Join(split[:parts], "/")
	return
}

func (a *HTTPRequestAccess) GetPiecesListContainer() (container io.ReadCloser, err error) {
	container, err = cmn.GetBody(a.target+"/reader", a.client)
	return
}

func (a *HTTPRequestAccess) GetChaptersListContainer() (container io.ReadCloser, err error) {
	container, err = cmn.GetBody(a.target+"/threadmarks", a.client)
	return
}

func (a *HTTPRequestAccess) GetPieceContainer(pieceNumber int) (container io.ReadCloser, err error) {
	container, err = cmn.GetBody(a.readerPageURL(pieceNumber), a.client)
	return
}

func (a *HTTPRequestAccess) readerPageURL(pageNumber int) string {
	url := a.target + "/reader"
	if pageNumber == 1 {
		return url
	}
	return url + fmt.Sprintf("/page-%d", pageNumber)
}
