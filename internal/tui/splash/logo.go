package splash


func(m Model) logoView() string {
	return m.Theme.TextAccent().Bold(true).
		Render("yasaworks") + m.Cursor.View()
}

