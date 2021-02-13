package cmn

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/mebyus/ffd/logs"
)

func SmartClose(c io.Closer) {
	err := c.Close()
	if err != nil {
		logs.Warn.Printf("closing: %v\n", err)
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

func GenerateFilenames(maxIndex int64, ext string) (filenames []string) {
	maxIndexStr := fmt.Sprintf("%d", maxIndex)
	for i := int64(1); i <= maxIndex; i++ {
		indexStr := fmt.Sprintf("%d", i)
		formatStr := ""
		if len(maxIndexStr)-len(indexStr) > 0 {
			formatStr = strings.Repeat("0", len(maxIndexStr)-len(indexStr)) + "%d.%s"
		} else {
			formatStr = "%d.%s"
		}
		filenames = append(filenames, fmt.Sprintf(formatStr, i, ext))
	}
	return
}
