package cmn

import (
	"fmt"
	"io"
)

func SmartClose(c io.Closer) {
	err := c.Close()
	if err != nil {
		fmt.Println(err)
	}
	return
}
