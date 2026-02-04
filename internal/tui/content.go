package tui

func (m model) getContent() string {
	page := "unknown"
	switch m.page {
	case blogPage:
		page = m.blogView()
	}

	return page
}
