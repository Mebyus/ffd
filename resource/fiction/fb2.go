package fiction

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func fb2base(chapters []Chapter) (root *html.Node) {
	root = &html.Node{
		Type: html.ElementNode,
		Data: "FictionBook",
	}
	description := &html.Node{
		Type:   html.ElementNode,
		Data:   "description",
		Parent: root,
	}
	root.FirstChild = description
	body := &html.Node{
		Type:     html.ElementNode,
		Data:     "body",
		DataAtom: atom.Body,
		Parent:   root,
	}
	body.PrevSibling = description
	description.NextSibling = body
	root.LastChild = body

	if len(chapters) == 0 {
		return
	}
	section := chapters[0].Body
	body.FirstChild = section
	section.Parent = body
	i := 0
	for i < len(chapters) {
		prev := section
		section = chapters[i].Body
		prev.NextSibling = section
		section.Parent = body
		section.PrevSibling = prev
		i++
	}
	body.LastChild = section
	return
}
