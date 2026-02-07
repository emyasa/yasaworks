// Package entries handles each blog entry's registration
package entries

import "github.com/emyasa/yasaworks/internal/tui/blog"

func init() {
	var entry = blog.BlogEntry{
		Name: "First Entry",
		Content: "No content.",
	}

	blog.Register(entry)
}

