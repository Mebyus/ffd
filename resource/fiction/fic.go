package fiction

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/mebyus/ffd/cmn"
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

func ext(format RenderFormat) string {
	switch format {
	case TXT:
		return ".txt"
	case FB2:
		return ".fb2"
	default:
		return ""
	}
}

func filename(name string, format RenderFormat) string {
	return name + ext(format)
}

func (b *Book) save(path string, format RenderFormat) (err error) {
	err = os.MkdirAll(filepath.Dir(path), 0774)
	if err != nil {
		return
	}
	file, err := os.Create(path)
	if err != nil {
		return
	}
	defer cmn.SmartClose(file)
	err = b.Format(file, format)
	return
}

func (b *Book) SaveAs(dirpath, name string, format RenderFormat) (err error) {
	err = b.save(filepath.Join(dirpath, filename(name, format)), format)
	return
}

func (b *Book) Save(dirpath string, format RenderFormat) (err error) {
	err = b.save(filepath.Join(dirpath, b.Filename(format)), format)
	return
}

func (b *Book) Filename(format RenderFormat) string {
	return filename(b.Title, format)
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
	// write something better to ensure the whole string will be written
	_, err = io.WriteString(dst, "<?xml version=\"1.0\" encoding=\"utf-8\"?>\n")
	if err != nil {
		return
	}
	root := fb2base(b.Chapters)
	err = html.Render(dst, root)
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
			if space && p != "" {
				space = false
				text += " "
			}
		case html.ElementNode:
			switch n.Data {
			case "p":
				text += "\n\n"
			case "i", "b", "em", "strong":
				if !space {
					space = true
					text += " "
				}
			}
		}
	}
	document.Walk(c.Body, action)
	_, err = io.Copy(dst, strings.NewReader(text))
	return
}
