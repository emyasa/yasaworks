package entries

import "github.com/emyasa/yasaworks/internal/tui/blog"

func init() {
	var entry = blog.BlogEntry{
		Name: "Second Entry",
		Content: "No content. 2",
	}

	blog.Register(entry)
}

