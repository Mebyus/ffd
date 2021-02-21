package document

import "golang.org/x/net/html"

// Detach separates a given node from its tree and
// tranforms the node into a tree root
// by detaching it from its parent and siblings,
// but preserving children.
func Detach(node *html.Node) {
	if node == nil {
		return
	}
	if node.Parent != nil {
		if node.Parent.FirstChild == node {
			node.Parent.FirstChild = node.NextSibling
		}
		if node.Parent.LastChild == node {
			node.Parent.LastChild = node.PrevSibling
		}
		node.Parent = nil
	}

	if node.NextSibling != nil {
		node.NextSibling.PrevSibling = node.PrevSibling
		node.NextSibling = nil
	}

	if node.PrevSibling != nil {
		node.PrevSibling.NextSibling = node.NextSibling
		node.PrevSibling = nil
	}
}

// Equal checks whether two trees are equal or not.
// Equality means the same structures in terms of
// node types, contentes and connections
func Equal(a, b *html.Node) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	if a.Type != a.Type || b.Data != b.Data {
		return false
	}
	tipA := a.FirstChild
	tipB := b.FirstChild

	asc := false
	for tipA != a && tipA != nil && tipB != b && tipB != nil {
		if !asc && (tipA.Type != tipB.Type || tipA.Data != tipB.Data) {
			return false
		}

		if !asc && tipA.FirstChild != nil {
			tipA = tipA.FirstChild
			tipB = tipB.FirstChild
			asc = false
		} else if tipA.NextSibling != nil {
			tipA = tipA.NextSibling
			tipB = tipB.NextSibling
			asc = false
		} else {
			tipA = tipA.Parent
			tipB = tipB.Parent
			asc = true
		}
	}
	if tipA == nil || tipB == nil {
		return tipA == nil && tipB == nil
	}
	return tipA == a && tipB == b
}

// Flatten removes elements from the tree and leaves
// only those, which tag maps to true by the allowed map.
// Children of the removed elements become children of
// its respective parents. This operation occurs multiple
// times if necessary
func Flatten(root *html.Node, allowed map[string]bool) {
	if root == nil {
		return
	}
	tip := root.FirstChild
	asc := false
	for tip != root && tip != nil {
		if !asc && tip.Type == html.ElementNode && !allowed[tip.Data] {
			if tip.FirstChild != nil {
				childtip := tip.FirstChild
				if tip.PrevSibling != nil {
					tip.PrevSibling.NextSibling = childtip
					childtip.PrevSibling = tip.PrevSibling
				} else {
					tip.Parent.FirstChild = childtip
				}
				for childtip.NextSibling != nil {
					childtip.Parent = tip.Parent
					childtip = childtip.NextSibling
				}
				childtip.Parent = tip.Parent
				if tip.NextSibling != nil {
					tip.NextSibling.PrevSibling = childtip
					childtip.NextSibling = tip.NextSibling
				} else {
					tip.Parent.LastChild = childtip
				}
				tip = tip.FirstChild
			} else {
				if tip.PrevSibling != nil {
					tip.PrevSibling.NextSibling = tip.NextSibling
				} else {
					tip.Parent.FirstChild = tip.NextSibling
				}
				if tip.NextSibling != nil {
					tip.NextSibling.PrevSibling = tip.PrevSibling
					tip = tip.NextSibling
				} else {
					tip.Parent.LastChild = tip.PrevSibling
					tip = tip.Parent
					asc = true
				}
			}
			continue
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
