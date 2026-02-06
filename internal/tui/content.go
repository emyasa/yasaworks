package tui

import "github.com/emyasa/yasaworks/internal/tui/blog"

func (m model) getContent() string {
	page := "unknown"
	switch m.page {
	case blogPage:
		page = blog.BlogView()
	}

	return page
}
