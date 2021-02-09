package document

import (
	"strings"

	"golang.org/x/net/html"
)

func FindFirstByTag(root *html.Node, tag string) (node *html.Node) {
	check := func(n *html.Node) bool {
		return n.Type == html.ElementNode && n.Data == tag
	}
	return FindFirst(root, check)
}

func FindByTag(root *html.Node, tag string) (nodes []*html.Node) {
	check := func(n *html.Node) bool {
		return n.Type == html.ElementNode && n.Data == tag
	}
	return Find(root, check)
}

func FindFirstNonSpaceText(root *html.Node) (text string) {
	check := func(n *html.Node) bool {
		return n.Type == html.TextNode && strings.TrimSpace(n.Data) != ""
	}
	n := FindFirst(root, check)
	if n != nil {
		text = strings.TrimSpace(n.Data)
	}
	return
}

func FindLastNonSpaceText(root *html.Node) (text string) {
	check := func(n *html.Node) bool {
		return n.Type == html.TextNode && strings.TrimSpace(n.Data) != ""
	}
	n := FindLast(root, check)
	if n != nil {
		text = strings.TrimSpace(n.Data)
	}
	return
}

func FindNonSpaceTexts(root *html.Node) (texts []string) {
	check := func(n *html.Node) bool {
		return n.Type == html.TextNode && strings.TrimSpace(n.Data) != ""
	}
	nodes := Find(root, check)
	for _, n := range nodes {
		texts = append(texts, strings.TrimSpace(n.Data))
	}
	return
}

func FindAttributeValues(root *html.Node, key string) (values []string) {
	if root == nil {
		return
	}
	tip := root.FirstChild

	asc := false
	for tip != root && tip != nil {
		if !asc {
			for _, attr := range tip.Attr {
				if attr.Key == key {
					values = append(values, attr.Val)
				}
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

func Find(root *html.Node, check func(n *html.Node) bool) (nodes []*html.Node) {
	if root == nil {
		return
	}
	tip := root.FirstChild

	asc := false
	for tip != root && tip != nil {
		if !asc && check(tip) {
			nodes = append(nodes, tip)
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

func FindFirst(root *html.Node, check func(n *html.Node) bool) (node *html.Node) {
	if root == nil {
		return
	}
	tip := root.FirstChild

	asc := false
	for tip != root && tip != nil {
		if !asc && check(tip) {
			return tip
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

func FindLast(root *html.Node, check func(n *html.Node) bool) (node *html.Node) {
	if root == nil {
		return
	}
	tip := root.LastChild

	asc := false
	for tip != root && tip != nil {
		if !asc && check(tip) {
			return tip
		}

		if !asc && tip.LastChild != nil {
			tip = tip.LastChild
			asc = false
		} else if tip.PrevSibling != nil {
			tip = tip.PrevSibling
			asc = false
		} else {
			tip = tip.Parent
			asc = true
		}
	}
	return
}

func Walk(root *html.Node, action func(n *html.Node)) {
	if root == nil {
		return
	}
	tip := root.FirstChild

	asc := false
	for tip != root && tip != nil {
		if !asc {
			action(tip)
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
}
