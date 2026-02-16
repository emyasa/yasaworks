package tui


func (m model) getContent() string {
	page := "unknown"
	switch m.page {
	case blogPage:
		page = m.blog.View()
	case termsPage:
		page = m.terms.View()
	}

	return page
}
