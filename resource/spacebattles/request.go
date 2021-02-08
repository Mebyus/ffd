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
	target, _, err := analyze(rawtarget)
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

func analyze(url string) (base, name string, err error) {
	split := strings.Split(url, "/")
	if len(split) < 2 {
		err = fmt.Errorf("incorrect url")
		return
	}
	parts := 3
	if split[0] == "https:" || split[0] == "http" {
		parts = 5
	}
	if len(split) < parts || split[parts-1] == "" {
		err = fmt.Errorf("incorrect url")
		return
	}
	name = strings.Split(split[parts-1], ".")[0]
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

func (a *HTTPRequestAccess) GetPieceContainer(pieceNumber int64) (container io.ReadCloser, err error) {
	container, err = cmn.GetBody(readerPageURL(a.target, pieceNumber), a.client)
	return
}

func readerPageURL(baseURL string, pageNumber int64) string {
	url := baseURL + "/reader"
	if pageNumber == 1 {
		return url
	}
	return url + fmt.Sprintf("/page-%d", pageNumber)
}
