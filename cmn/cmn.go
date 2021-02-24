package cmn

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func SmartClose(c io.Closer) {
	err := c.Close()
	if err != nil {
		fmt.Printf("closing: %v\n", err)
	}
	return
}

// Plural returns "s" when there is more than one number n is given
// and returns empty string otherwise.
// It is meant for use with format string like fmt.Printf("%d apple%s", n, Plural(n)).
// Another alternative is just using string concatenation "apple" + Plural(n)
func Plural(n int) string {
	if n == 1 {
		return ""
	}
	return "s"
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

func GetBytes(url string, client *http.Client) (b []byte, err error) {
	body, err := GetBody(url, client)
	if err != nil {
		return
	}
	defer SmartClose(body)
	b, err = ioutil.ReadAll(body)
	return
}

func GenerateFilenames(maxIndex int, ext string) (filenames []string) {
	maxIndexStr := fmt.Sprintf("%d", maxIndex)
	for i := 1; i <= maxIndex; i++ {
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
