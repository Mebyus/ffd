package document

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func NewParagraph() *html.Node {
	return &html.Node{
		Data:     "p",
		DataAtom: atom.P,
		Type:     html.ElementNode,
	}
}
