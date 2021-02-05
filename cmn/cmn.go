package cmn

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func SmartClose(c io.Closer) {
	err := c.Close()
	if err != nil {
		fmt.Printf("closing: %v\n", err)
	}
	return
}

func GetBody(url string, client *http.Client) (body io.ReadCloser, err error) {
	request, err := http.NewRequest("GET", url, bytes.NewReader([]byte{}))
	if err != nil {
		err = fmt.Errorf("Preparing request: %v", err)
		return
	}
	response, err := client.Do(request)
	if err != nil {
		err = fmt.Errorf("Making request: %v", err)
		return
	}
	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("Response status: %s [ %s ]", response.Status, url)
		return
	}
	body = response.Body
	return
}
