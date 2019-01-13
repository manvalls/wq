package wq

import (
	"html"
	"sync"

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
	mux     sync.Mutex
	delta   *wit.Delta
	parent  *Action
	action  wit.Action
	wrapper func(wit.Action) wit.Action
}

// Delta returns the resolved delta
func (action *Action) Delta() wit.Delta {
	action.mux.Lock()
	defer action.mux.Unlock()

	if action.delta != nil {
		return *action.delta
	}

	root := action
	actions := []wit.Action{}

	for action != nil {
		if action.wrapper != nil {
			actions = []wit.Action{action.wrapper(wit.List(actions...))}
		} else {
			actions = append(actions, action.action)
		}

		action = action.parent
	}

	delta := wit.List(actions...).Delta()
	root.delta = &delta
	return delta
}

// Procedure returns a procedure wrapping the resolved delta
func (action *Action) Procedure() wok.Procedure {
	return wok.Action(action).Procedure()
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
func (s Selector) First() *Action {
	return &Action{
		wrapper: func(action wit.Action) wit.Action {
			return s.selector.One(action)
		},
	}
}

// Apply applies the given actions
func (action *Action) Apply(actions ...wit.Action) *Action {
	return &Action{
		parent: action,
		action: wit.List(actions...),
	}
}

// Root matches the root element
func (action *Action) Root() *Action {
	return &Action{
		parent: action,
		wrapper: func(action wit.Action) wit.Action {
			return wit.Root(action)
		},
	}
}

// Parent matches the parents of matched elements
func (action *Action) Parent() *Action {
	return &Action{
		parent: action,
		wrapper: func(action wit.Action) wit.Action {
			return wit.Parent(action)
		},
	}
}

// FirstChild matches the first child of matched elements
func (action *Action) FirstChild() *Action {
	return &Action{
		parent: action,
		wrapper: func(action wit.Action) wit.Action {
			return wit.FirstChild(action)
		},
	}
}

// LastChild matches the last child of matched elements
func (action *Action) LastChild() *Action {
	return &Action{
		parent: action,
		wrapper: func(action wit.Action) wit.Action {
			return wit.LastChild(action)
		},
	}
}

// PrevSibling matches the previous sibling of matched elements
func (action *Action) PrevSibling() *Action {
	return &Action{
		parent: action,
		wrapper: func(action wit.Action) wit.Action {
			return wit.PrevSibling(action)
		},
	}
}

// NextSibling matches the next sibling of matched elements
func (action *Action) NextSibling() *Action {
	return &Action{
		parent: action,
		wrapper: func(action wit.Action) wit.Action {
			return wit.NextSibling(action)
		},
	}
}

// Remove removes matched elements
func (action *Action) Remove() *Action {
	return &Action{
		parent: action,
		action: wit.Remove,
	}
}

// Clear empties matched elements
func (action *Action) Clear() *Action {
	return &Action{
		parent: action,
		action: wit.Clear,
	}
}

// Set sets the contents of matched elements
func (action *Action) Set(factory wit.Factory) *Action {
	return &Action{
		parent: action,
		action: wit.HTML(factory),
	}
}

// SetText sets the text content of matched elements
func (action *Action) SetText(text string) *Action {
	return action.Set(wit.FromString(html.EscapeString(text)))
}

// SetHTML sets the HTML content of matched elements
func (action *Action) SetHTML(text string) *Action {
	return action.Set(wit.FromString(text))
}

// Replace replaces matching elements
func (action *Action) Replace(factory wit.Factory) *Action {
	return &Action{
		parent: action,
		action: wit.Replace(factory),
	}
}

// ReplaceText replaces matching elements with the provided text
func (action *Action) ReplaceText(text string) *Action {
	return action.Replace(wit.FromString(html.EscapeString(text)))
}

// ReplaceHTML replaces matching elements with the provided HTML
func (action *Action) ReplaceHTML(text string) *Action {
	return action.Replace(wit.FromString(text))
}

// Append appends content to matched elements
func (action *Action) Append(factory wit.Factory) *Action {
	return &Action{
		parent: action,
		action: wit.Append(factory),
	}
}

// AppendText appends text to elements
func (action *Action) AppendText(text string) *Action {
	return action.Append(wit.FromString(html.EscapeString(text)))
}

// AppendHTML appends HTML to matched elements
func (action *Action) AppendHTML(text string) *Action {
	return action.Append(wit.FromString(text))
}

// Prepend prepends content to matched elements
func (action *Action) Prepend(factory wit.Factory) *Action {
	return &Action{
		parent: action,
		action: wit.Prepend(factory),
	}
}

// PrependText prepends text to elements
func (action *Action) PrependText(text string) *Action {
	return action.Prepend(wit.FromString(html.EscapeString(text)))
}

// PrependHTML prepends HTML to matched elements
func (action *Action) PrependHTML(text string) *Action {
	return action.Prepend(wit.FromString(text))
}

// InsertAfter inserts content after matched elements
func (action *Action) InsertAfter(factory wit.Factory) *Action {
	return &Action{
		parent: action,
		action: wit.InsertAfter(factory),
	}
}

// InsertTextAfter inserts text after elements
func (action *Action) InsertTextAfter(text string) *Action {
	return action.InsertAfter(wit.FromString(html.EscapeString(text)))
}

// InsertHTMLAfter inserts HTML after matched elements
func (action *Action) InsertHTMLAfter(text string) *Action {
	return action.InsertAfter(wit.FromString(text))
}

// InsertBefore inserts content before matched elements
func (action *Action) InsertBefore(factory wit.Factory) *Action {
	return &Action{
		parent: action,
		action: wit.InsertBefore(factory),
	}
}

// InsertTextBefore inserts text before elements
func (action *Action) InsertTextBefore(text string) *Action {
	return action.InsertBefore(wit.FromString(html.EscapeString(text)))
}

// InsertHTMLBefore inserts HTML before matched elements
func (action *Action) InsertHTMLBefore(text string) *Action {
	return action.InsertBefore(wit.FromString(text))
}

// AddAttr adds the provided attributes to the matching elements
func (action *Action) AddAttr(attr map[string]string) *Action {
	return &Action{
		parent: action,
		action: wit.AddAttr(attr),
	}
}

// SetAttr sets the attributes of the matching elements
func (action *Action) SetAttr(attr map[string]string) *Action {
	return &Action{
		parent: action,
		action: wit.SetAttr(attr),
	}
}

// RmAttr removes the provided attributes from the matching elements
func (action *Action) RmAttr(attrs ...string) *Action {
	return &Action{
		parent: action,
		action: wit.RmAttr(attrs...),
	}
}

// AddStyles adds the provided styles to the matching elements
func (action *Action) AddStyles(styles map[string]string) *Action {
	return &Action{
		parent: action,
		action: wit.AddStyles(styles),
	}
}

// RmStyles removes the provided styles from the matching elements
func (action *Action) RmStyles(styles ...string) *Action {
	return &Action{
		parent: action,
		action: wit.RmStyles(styles...),
	}
}

// AddClass adds the provided class to the matching elements
func (action *Action) AddClass(class string) *Action {
	return &Action{
		parent: action,
		action: wit.AddClass(class),
	}
}

// RmClass removes the provided class from the matching elements
func (action *Action) RmClass(class string) *Action {
	return &Action{
		parent: action,
		action: wit.RmClass(class),
	}
}
