package wq

import (
	"html"

	"github.com/manvalls/wit"
)

// Node represents one or more HTML nodes
type Node struct {
	Send func(wit.Delta) // Send will be called as a result of this node's methods
}

// Selector encapsulates a CSS selector. Must be initialised by Node.S()
type Selector struct {
	selector wit.Selector
	parent   Node
	Node
}

// First returns a Node that applies deltas to the first matching element
func (s Selector) First() Node {
	return Node{func(delta wit.Delta) {
		s.parent.Send(wit.First{
			Selector: s.selector,
			Delta:    delta,
		})
	}}
}

// S returns a Selector. Accepts either a string or a wit.Selector as parameter.
func (n Node) S(selector interface{}) Selector {
	var witSelector wit.Selector

	switch s := selector.(type) {
	case string:
		witSelector = wit.S(s)
	case wit.Selector:
		witSelector = s
	default:
		witSelector = wit.S("*")
	}

	return Selector{
		selector: witSelector,
		parent:   n,
		Node: Node{
			Send: func(delta wit.Delta) {
				n.Send(wit.First{
					Selector: witSelector,
					Delta:    delta,
				})
			},
		},
	}
}

// Root matches the root node
func (n Node) Root() Node {
	return Node{func(delta wit.Delta) {
		n.Send(wit.Root{
			Delta: delta,
		})
	}}
}

// Parent matches the parents of matched elements
func (n Node) Parent() Node {
	return Node{func(delta wit.Delta) {
		n.Send(wit.Root{
			Delta: delta,
		})
	}}
}

// FirstChild matches the first child of matched elements
func (n Node) FirstChild() Node {
	return Node{func(delta wit.Delta) {
		n.Send(wit.FirstChild{
			Delta: delta,
		})
	}}
}

// LastChild matches the last child of matched elements
func (n Node) LastChild() Node {
	return Node{func(delta wit.Delta) {
		n.Send(wit.LastChild{
			Delta: delta,
		})
	}}
}

// PrevSibling matches the previous sibling of matched elements
func (n Node) PrevSibling() Node {
	return Node{func(delta wit.Delta) {
		n.Send(wit.PrevSibling{
			Delta: delta,
		})
	}}
}

// NextSibling matches the previous sibling of matched elements
func (n Node) NextSibling() Node {
	return Node{func(delta wit.Delta) {
		n.Send(wit.NextSibling{
			Delta: delta,
		})
	}}
}

// Remove removes matched elements
func (n Node) Remove() Node {
	n.Send(wit.Remove{})
	return n
}

// Clear empties matched elements
func (n Node) Clear() Node {
	n.Send(wit.Clear{})
	return n
}

// Text sets the text content of matched elements
func (n Node) Text(text string) Node {
	n.Send(wit.HTML{
		HTMLSource: wit.HTMLFromString(html.EscapeString(text)),
	})

	return n
}

func getHTMLSource(html interface{}) wit.HTMLSource {
	switch h := html.(type) {
	case string:
		return wit.HTMLFromString(h)
	case wit.HTMLSource:
		return h
	default:
		return wit.HTMLFromString("")
	}
}

// HTML sets the inner HTML of matched elements. Accepts either a string or an HTMLSource as argument.
func (n Node) HTML(html interface{}) Node {
	n.Send(wit.HTML{
		HTMLSource: getHTMLSource(html),
	})

	return n
}

// Replace replaces matching elements with the provided HTML
func (n Node) Replace(html interface{}) Node {
	n.Send(wit.Replace{
		HTMLSource: getHTMLSource(html),
	})

	return n
}

// Append appends content to matched elements
func (n Node) Append(html interface{}) Node {
	n.Send(wit.Append{
		HTMLSource: getHTMLSource(html),
	})

	return n
}

// Prepend prepends content to matched elements
func (n Node) Prepend(html interface{}) Node {
	n.Send(wit.Prepend{
		HTMLSource: getHTMLSource(html),
	})

	return n
}

// InsertAfter inserts content after matched elements
func (n Node) InsertAfter(html interface{}) Node {
	n.Send(wit.InsertAfter{
		HTMLSource: getHTMLSource(html),
	})

	return n
}

// InsertBefore inserts content before matched elements
func (n Node) InsertBefore(html interface{}) Node {
	n.Send(wit.InsertBefore{
		HTMLSource: getHTMLSource(html),
	})

	return n
}

// SetAttr sets the attributes of matching elements
func (n Node) SetAttr(attr map[string]string) Node {
	n.Send(wit.SetAttr{
		Attributes: attr,
	})

	return n
}

// ReplaceAttr replaces the attributes of matching elements
func (n Node) ReplaceAttr(attr map[string]string) Node {
	n.Send(wit.ReplaceAttr{
		Attributes: attr,
	})

	return n
}

// RmAttr removes the provided attributes from matching elements
func (n Node) RmAttr(attrs ...string) Node {
	n.Send(wit.RmAttr{
		Attributes: attrs,
	})

	return n
}

// SetStyles sets the styles of matching elements
func (n Node) SetStyles(styles map[string]string) Node {
	n.Send(wit.SetStyles{
		Styles: styles,
	})

	return n
}

// RmStyles removes the provided styles from matching elements
func (n Node) RmStyles(styles ...string) Node {
	n.Send(wit.RmStyles{
		Styles: styles,
	})

	return n
}

// AddClass adds the provided class to matching elements
func (n Node) AddClass(class string) Node {
	n.Send(wit.AddClasses{
		Classes: class,
	})

	return n
}

// RmClass removes the provided class from matching elements
func (n Node) RmClass(class string) Node {
	n.Send(wit.RmClasses{
		Classes: class,
	})

	return n
}
