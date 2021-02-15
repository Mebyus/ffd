package internal

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

type RenderFormat string

const (
	TXT RenderFormat = "TXT"
	FB2 RenderFormat = "FB2"
)

type Chapter struct {
	Title string
	Body  *html.Node
}

type Book struct {
	Title    string
	Author   string
	Chapters []Chapter
}

func (b *Book) Format(dst io.Writer, f RenderFormat) (err error) {
	switch f {
	case TXT:
		err = b.FormatTXT(dst)
	case FB2:
		err = b.FormatFB2(dst)
	default:
		err = fmt.Errorf("unknown format [ %s ]", f)
	}
	return
}

func (b *Book) FormatFB2(dst io.Writer) (err error) {
	return
}

func (b *Book) FormatTXT(dst io.Writer) (err error) {
	return
}
