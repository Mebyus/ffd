package document

import (
	"strings"

	"golang.org/x/net/html"
)

type Document struct {
	Root *html.Node

	// im maps node ids to nodes
	im map[string]*html.Node

	// cm maps node classes to a slice of nodes which possess specified class
	cm map[string][]*html.Node

	// tm maps node tags to a slice of nodes with the specified tag
	tm map[string][]*html.Node
}

func FromNode(root *html.Node) (d *Document) {
	d = &Document{
		Root: root,
		im:   make(map[string]*html.Node),
		cm:   make(map[string][]*html.Node),
		tm:   make(map[string][]*html.Node),
	}
	d.traverse()
	return
}

func (d *Document) GetNodeById(id string) *html.Node {
	return d.im[id]
}

func (d *Document) GetNodesByClass(class string) []*html.Node {
	return d.cm[class]
}

func (d *Document) GetNodesByTag(tag string) []*html.Node {
	return d.tm[tag]
}

func (d *Document) GetNodesByTagClass(tag string, class string) (nodes []*html.Node) {
	tagNodes := d.tm[tag]
	for _, n := range tagNodes {
		if HasClass(n, class) {
			nodes = append(nodes, n)
		}
	}
	return
}

func (d *Document) traverse() {
	asc := false
	tip := d.Root
	for tip != nil {
		if !asc {
			switch tip.Type {
			case html.ElementNode:
				id, classes := extract(tip.Attr)
				if id != "" {
					d.im[id] = tip
				}
				if len(classes) != 0 {
					for _, class := range classes {
						d.cm[class] = append(d.cm[class], tip)
					}
				}
				d.tm[tip.Data] = append(d.tm[tip.Data], tip)
			case html.TextNode:
			case html.DoctypeNode:
			case html.DocumentNode:
			case html.CommentNode:
			}
		}

		if !asc && tip.FirstChild != nil {
			tip = tip.FirstChild
			asc = false
		} else if tip.NextSibling != nil {
			tip = tip.NextSibling
			asc = false
		} else {
			tip = tip.Parent
			asc = true
		}
	}
	return
}

func HasClass(node *html.Node, class string) bool {
	for _, attr := range node.Attr {
		if attr.Key == "class" && strings.Contains(attr.Val, class) {
			return true
		}
	}
	return false
}

func GetAttributeValue(node *html.Node, key string) (value string) {
	if node == nil {
		return
	}
	for _, attr := range node.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return
}

func GetChildren(node *html.Node) (nodes []*html.Node) {
	if node == nil {
		return
	}
	tip := node.FirstChild
	for tip != nil {
		nodes = append(nodes, tip)
		tip = tip.NextSibling
	}
	return
}

func extract(attrs []html.Attribute) (id string, classes []string) {
	for _, attr := range attrs {
		switch attr.Key {
		case "id":
			id = attr.Val
		case "class":
			classes = strings.Fields(attr.Val)
		}
	}
	return
}
