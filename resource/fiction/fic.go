package fiction

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/mebyus/ffd/document"
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

func (b *Book) Save(dirpath string, format RenderFormat) (err error) {
	file, err := os.Create(filepath.Join(dirpath, b.Filename(format)))
	if err != nil {
		return
	}
	err = b.Format(file, format)
	return
}

func (b *Book) Filename(format RenderFormat) (filename string) {
	ext := ""
	switch format {
	case TXT:
		ext = ".txt"
	case FB2:
		ext = ".fb2"
	}
	filename = b.Title + ext
	return
}

func (b *Book) Format(dst io.Writer, format RenderFormat) (err error) {
	switch format {
	case TXT:
		err = b.FormatTXT(dst)
	case FB2:
		err = b.FormatFB2(dst)
	default:
		err = fmt.Errorf("unknown format [ %s ]", format)
	}
	return
}

func (b *Book) FormatFB2(dst io.Writer) (err error) {
	return
}

func (b *Book) FormatTXT(dst io.Writer) (err error) {
	for _, c := range b.Chapters {
		err = c.FormatTXT(dst)
		if err != nil {
			return
		}
	}
	return
}

func (c *Chapter) FormatTXT(dst io.Writer) (err error) {
	text := "\n\n" + c.Title
	first := true
	space := false
	action := func(n *html.Node) {
		switch n.Type {
		case html.TextNode:
			p := strings.TrimSpace(n.Data)
			if first && p != "" {
				first = false
				if p == c.Title {
					p = ""
				}
			}
			text += p
			if space {
				space = false
				text += " "
			}
		case html.ElementNode:
			if n.Data == "p" {
				text += "\n\n"
			} else if n.Data == "i" || n.Data == "b" {
				space = true
				text += " "
			}
		}
	}
	document.Walk(c.Body, action)
	_, err = io.Copy(dst, strings.NewReader(text))
	return
}
