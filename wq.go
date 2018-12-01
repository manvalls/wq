package wq

import (
	"html"

	"github.com/manvalls/wok"

	"github.com/manvalls/wit"
)

// Head matches the document head
var Head = S("head").First()

// Body matches the document body
var Body = S("body").First()

// Title matches the document title
var Title = S("head > title").First()

// Action wraps a set of actions and wrappers
type Action struct {
	parent  *Action
	action  wit.Action
	wrapper func(wit.Action) wit.Action
}

// Delta returns the resolved delta
func (a Action) Delta() wit.Delta {
	actions := []wit.Action{}
	action := &a

	for action != nil {
		if action.wrapper != nil {
			actions = []wit.Action{action.wrapper(wit.List(actions...))}
		} else {
			actions = append(actions, action.action)
		}

		action = action.parent
	}

	return wit.List(actions...).Delta()
}

// Procedure returns a procedure wrapping the resolved delta
func (a Action) Procedure() wok.Procedure {
	return wok.Action(a).Procedure()
}

// Selector wraps a CSS selector
type Selector struct {
	selector wit.Selector
	*Action
}

// S builds a wrapper based on the provided selector
func S(selector string) Selector {
	s := wit.S(selector)
	return Selector{
		selector: s,
		Action: &Action{
			wrapper: func(action wit.Action) wit.Action {
				return s.All(action)
			},
		},
	}
}

// First matches only the first element
func (s Selector) First() Action {
	return Action{
		wrapper: func(action wit.Action) wit.Action {
			return s.selector.One(action)
		},
	}
}

// Apply applies the given actions
func (a Action) Apply(actions ...wit.Action) Action {
	return Action{
		parent: &a,
		action: wit.List(actions...),
	}
}

// Root matches the root element
func (a Action) Root() Action {
	return Action{
		parent: &a,
		wrapper: func(action wit.Action) wit.Action {
			return wit.Root(action)
		},
	}
}

// Parent matches the parents of matched elements
func (a Action) Parent() Action {
	return Action{
		parent: &a,
		wrapper: func(action wit.Action) wit.Action {
			return wit.Parent(action)
		},
	}
}

// FirstChild matches the first child of matched elements
func (a Action) FirstChild() Action {
	return Action{
		parent: &a,
		wrapper: func(action wit.Action) wit.Action {
			return wit.FirstChild(action)
		},
	}
}

// LastChild matches the last child of matched elements
func (a Action) LastChild() Action {
	return Action{
		parent: &a,
		wrapper: func(action wit.Action) wit.Action {
			return wit.LastChild(action)
		},
	}
}

// PrevSibling matches the previous sibling of matched elements
func (a Action) PrevSibling() Action {
	return Action{
		parent: &a,
		wrapper: func(action wit.Action) wit.Action {
			return wit.PrevSibling(action)
		},
	}
}

// NextSibling matches the next sibling of matched elements
func (a Action) NextSibling() Action {
	return Action{
		parent: &a,
		wrapper: func(action wit.Action) wit.Action {
			return wit.NextSibling(action)
		},
	}
}

// Remove removes matched elements
func (a Action) Remove() Action {
	return Action{
		parent: &a,
		action: wit.Remove,
	}
}

// Clear empties matched elements
func (a Action) Clear() Action {
	return Action{
		parent: &a,
		action: wit.Clear,
	}
}

// Set sets the contents of matched elements
func (a Action) Set(factory wit.Factory) Action {
	return Action{
		parent: &a,
		action: wit.HTML(factory),
	}
}

// SetText sets the text content of matched elements
func (a Action) SetText(text string) Action {
	return a.Set(wit.FromString(html.EscapeString(text)))
}

// SetHTML sets the HTML content of matched elements
func (a Action) SetHTML(text string) Action {
	return a.Set(wit.FromString(text))
}

// Replace replaces matching elements
func (a Action) Replace(factory wit.Factory) Action {
	return Action{
		parent: &a,
		action: wit.Replace(factory),
	}
}

// ReplaceText replaces matching elements with the provided text
func (a Action) ReplaceText(text string) Action {
	return a.Replace(wit.FromString(html.EscapeString(text)))
}

// ReplaceHTML replaces matching elements with the provided HTML
func (a Action) ReplaceHTML(text string) Action {
	return a.Replace(wit.FromString(text))
}

// Append appends content to matched elements
func (a Action) Append(factory wit.Factory) Action {
	return Action{
		parent: &a,
		action: wit.Append(factory),
	}
}

// AppendText appends text to elements
func (a Action) AppendText(text string) Action {
	return a.Append(wit.FromString(html.EscapeString(text)))
}

// AppendHTML appends HTML to matched elements
func (a Action) AppendHTML(text string) Action {
	return a.Append(wit.FromString(text))
}

// Prepend prepends content to matched elements
func (a Action) Prepend(factory wit.Factory) Action {
	return Action{
		parent: &a,
		action: wit.Prepend(factory),
	}
}

// PrependText prepends text to elements
func (a Action) PrependText(text string) Action {
	return a.Prepend(wit.FromString(html.EscapeString(text)))
}

// PrependHTML prepends HTML to matched elements
func (a Action) PrependHTML(text string) Action {
	return a.Prepend(wit.FromString(text))
}

// InsertAfter inserts content after matched elements
func (a Action) InsertAfter(factory wit.Factory) Action {
	return Action{
		parent: &a,
		action: wit.InsertAfter(factory),
	}
}

// InsertTextAfter inserts text after elements
func (a Action) InsertTextAfter(text string) Action {
	return a.InsertAfter(wit.FromString(html.EscapeString(text)))
}

// InsertHTMLAfter inserts HTML after matched elements
func (a Action) InsertHTMLAfter(text string) Action {
	return a.InsertAfter(wit.FromString(text))
}

// InsertBefore inserts content before matched elements
func (a Action) InsertBefore(factory wit.Factory) Action {
	return Action{
		parent: &a,
		action: wit.InsertBefore(factory),
	}
}

// InsertTextBefore inserts text before elements
func (a Action) InsertTextBefore(text string) Action {
	return a.InsertBefore(wit.FromString(html.EscapeString(text)))
}

// InsertHTMLBefore inserts HTML before matched elements
func (a Action) InsertHTMLBefore(text string) Action {
	return a.InsertBefore(wit.FromString(text))
}

// AddAttr adds the provided attributes to the matching elements
func (a Action) AddAttr(attr map[string]string) Action {
	return Action{
		parent: &a,
		action: wit.AddAttr(attr),
	}
}

// SetAttr sets the attributes of the matching elements
func (a Action) SetAttr(attr map[string]string) Action {
	return Action{
		parent: &a,
		action: wit.SetAttr(attr),
	}
}

// RmAttr removes the provided attributes from the matching elements
func (a Action) RmAttr(attrs ...string) Action {
	return Action{
		parent: &a,
		action: wit.RmAttr(attrs...),
	}
}

// AddStyles adds the provided styles to the matching elements
func (a Action) AddStyles(styles map[string]string) Action {
	return Action{
		parent: &a,
		action: wit.AddStyles(styles),
	}
}

// RmStyles removes the provided styles from the matching elements
func (a Action) RmStyles(styles ...string) Action {
	return Action{
		parent: &a,
		action: wit.RmStyles(styles...),
	}
}

// AddClass adds the provided class to the matching elements
func (a Action) AddClass(class string) Action {
	return Action{
		parent: &a,
		action: wit.AddClass(class),
	}
}

// RmClass removes the provided class from the matching elements
func (a Action) RmClass(class string) Action {
	return Action{
		parent: &a,
		action: wit.RmClass(class),
	}
}
