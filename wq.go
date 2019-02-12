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

// Command wraps a set of commands and wrappers
type Command struct {
	mux     sync.Mutex
	delta   *wit.Delta
	parent  *Command
	command wit.Command
	wrapper func(wit.Command) wit.Command
}

// Delta returns the resolved delta
func (command *Command) Delta() wit.Delta {
	command.mux.Lock()
	defer command.mux.Unlock()

	if command.delta != nil {
		return *command.delta
	}

	root := command
	commands := []wit.Command{}

	for command != nil {
		if command.wrapper != nil {
			commands = []wit.Command{command.wrapper(wit.List(commands...))}
		} else {
			commands = append(commands, command.command)
		}

		command = command.parent
	}

	delta := wit.List(commands...).Delta()
	root.delta = &delta
	return delta
}

// Procedure returns a procedure wrapping the resolved delta
func (command *Command) Procedure() wok.Procedure {
	return wok.Command(command).Procedure()
}

// Selector wraps a CSS selector
type Selector struct {
	selector wit.Selector
	*Command
}

// S builds a wrapper based on the provided selector
func S(selector string) Selector {
	s := wit.S(selector)
	return Selector{
		selector: s,
		Command: &Command{
			wrapper: func(command wit.Command) wit.Command {
				return s.All(command)
			},
		},
	}
}

// First matches only the first element
func (s Selector) First() *Command {
	return &Command{
		wrapper: func(command wit.Command) wit.Command {
			return s.selector.One(command)
		},
	}
}

// Apply applies the given commands
func (command *Command) Apply(commands ...wit.Command) *Command {
	return &Command{
		parent:  command,
		command: wit.List(commands...),
	}
}

// Root matches the root element
func (command *Command) Root() *Command {
	return &Command{
		parent: command,
		wrapper: func(command wit.Command) wit.Command {
			return wit.Root(command)
		},
	}
}

// Parent matches the parents of matched elements
func (command *Command) Parent() *Command {
	return &Command{
		parent: command,
		wrapper: func(command wit.Command) wit.Command {
			return wit.Parent(command)
		},
	}
}

// FirstChild matches the first child of matched elements
func (command *Command) FirstChild() *Command {
	return &Command{
		parent: command,
		wrapper: func(command wit.Command) wit.Command {
			return wit.FirstChild(command)
		},
	}
}

// LastChild matches the last child of matched elements
func (command *Command) LastChild() *Command {
	return &Command{
		parent: command,
		wrapper: func(command wit.Command) wit.Command {
			return wit.LastChild(command)
		},
	}
}

// PrevSibling matches the previous sibling of matched elements
func (command *Command) PrevSibling() *Command {
	return &Command{
		parent: command,
		wrapper: func(command wit.Command) wit.Command {
			return wit.PrevSibling(command)
		},
	}
}

// NextSibling matches the next sibling of matched elements
func (command *Command) NextSibling() *Command {
	return &Command{
		parent: command,
		wrapper: func(command wit.Command) wit.Command {
			return wit.NextSibling(command)
		},
	}
}

// Remove removes matched elements
func (command *Command) Remove() *Command {
	return &Command{
		parent:  command,
		command: wit.Remove,
	}
}

// Clear empties matched elements
func (command *Command) Clear() *Command {
	return &Command{
		parent:  command,
		command: wit.Clear,
	}
}

// Set sets the contents of matched elements
func (command *Command) Set(factory wit.Factory) *Command {
	return &Command{
		parent:  command,
		command: wit.HTML(factory),
	}
}

// SetText sets the text content of matched elements
func (command *Command) SetText(text string) *Command {
	return command.Set(wit.FromString(html.EscapeString(text)))
}

// SetHTML sets the HTML content of matched elements
func (command *Command) SetHTML(text string) *Command {
	return command.Set(wit.FromString(text))
}

// Replace replaces matching elements
func (command *Command) Replace(factory wit.Factory) *Command {
	return &Command{
		parent:  command,
		command: wit.Replace(factory),
	}
}

// ReplaceText replaces matching elements with the provided text
func (command *Command) ReplaceText(text string) *Command {
	return command.Replace(wit.FromString(html.EscapeString(text)))
}

// ReplaceHTML replaces matching elements with the provided HTML
func (command *Command) ReplaceHTML(text string) *Command {
	return command.Replace(wit.FromString(text))
}

// Append appends content to matched elements
func (command *Command) Append(factory wit.Factory) *Command {
	return &Command{
		parent:  command,
		command: wit.Append(factory),
	}
}

// AppendText appends text to elements
func (command *Command) AppendText(text string) *Command {
	return command.Append(wit.FromString(html.EscapeString(text)))
}

// AppendHTML appends HTML to matched elements
func (command *Command) AppendHTML(text string) *Command {
	return command.Append(wit.FromString(text))
}

// Prepend prepends content to matched elements
func (command *Command) Prepend(factory wit.Factory) *Command {
	return &Command{
		parent:  command,
		command: wit.Prepend(factory),
	}
}

// PrependText prepends text to elements
func (command *Command) PrependText(text string) *Command {
	return command.Prepend(wit.FromString(html.EscapeString(text)))
}

// PrependHTML prepends HTML to matched elements
func (command *Command) PrependHTML(text string) *Command {
	return command.Prepend(wit.FromString(text))
}

// InsertAfter inserts content after matched elements
func (command *Command) InsertAfter(factory wit.Factory) *Command {
	return &Command{
		parent:  command,
		command: wit.InsertAfter(factory),
	}
}

// InsertTextAfter inserts text after elements
func (command *Command) InsertTextAfter(text string) *Command {
	return command.InsertAfter(wit.FromString(html.EscapeString(text)))
}

// InsertHTMLAfter inserts HTML after matched elements
func (command *Command) InsertHTMLAfter(text string) *Command {
	return command.InsertAfter(wit.FromString(text))
}

// InsertBefore inserts content before matched elements
func (command *Command) InsertBefore(factory wit.Factory) *Command {
	return &Command{
		parent:  command,
		command: wit.InsertBefore(factory),
	}
}

// InsertTextBefore inserts text before elements
func (command *Command) InsertTextBefore(text string) *Command {
	return command.InsertBefore(wit.FromString(html.EscapeString(text)))
}

// InsertHTMLBefore inserts HTML before matched elements
func (command *Command) InsertHTMLBefore(text string) *Command {
	return command.InsertBefore(wit.FromString(text))
}

// AddAttr adds the provided attributes to the matching elements
func (command *Command) AddAttr(attr map[string]string) *Command {
	return &Command{
		parent:  command,
		command: wit.AddAttr(attr),
	}
}

// SetAttr sets the attributes of the matching elements
func (command *Command) SetAttr(attr map[string]string) *Command {
	return &Command{
		parent:  command,
		command: wit.SetAttr(attr),
	}
}

// RmAttr removes the provided attributes from the matching elements
func (command *Command) RmAttr(attrs ...string) *Command {
	return &Command{
		parent:  command,
		command: wit.RmAttr(attrs...),
	}
}

// AddStyles adds the provided styles to the matching elements
func (command *Command) AddStyles(styles map[string]string) *Command {
	return &Command{
		parent:  command,
		command: wit.AddStyles(styles),
	}
}

// RmStyles removes the provided styles from the matching elements
func (command *Command) RmStyles(styles ...string) *Command {
	return &Command{
		parent:  command,
		command: wit.RmStyles(styles...),
	}
}

// AddClass adds the provided class to the matching elements
func (command *Command) AddClass(class string) *Command {
	return &Command{
		parent:  command,
		command: wit.AddClass(class),
	}
}

// RmClass removes the provided class from the matching elements
func (command *Command) RmClass(class string) *Command {
	return &Command{
		parent:  command,
		command: wit.RmClass(class),
	}
}
