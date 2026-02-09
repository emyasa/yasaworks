// Package entries handles each blog entry's registration
package entries

import (
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/glamour"
	"github.com/emyasa/yasaworks/internal/tui/blog"
)

func init() {
	in := `
# Hello World
This is a simple example of Markdown rendering with Glamour!

Scroll down to see more.

- Item 1
- Item 2
- Item 3
- Item 4
- Item 5
- Item 6
- Item 7

Bye!
`

	r, _ := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(50))

	content, _ := r.Render(in)
	vp := viewport.New(50, 10)
	vp.SetContent(content)

	var entry = blog.BlogEntry{
		Name: "First Entry",
		Content: content,
		Viewport: &vp,
	}

	blog.Register(entry)
}

