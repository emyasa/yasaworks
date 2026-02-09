// Package entries handles each blog entry's registration
package entries

import (
	"github.com/charmbracelet/glamour"
	"github.com/emyasa/yasaworks/internal/tui/blog"
)

func init() {
	in := 
`# Hello World
This is a simple example of Markdown rendering with Glamour!

Check out the other examples too.

Bye!`

	r, _ := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(50))

	content, _ := r.Render(in)
	var entry = blog.BlogEntry{
		Name: "First Entry",
		Content: content,
	}

	blog.Register(entry)
}

