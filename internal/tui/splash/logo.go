package splash


func(m Model) logoView() string {
	return m.theme.TextAccent().Bold(true).
		Render("yasaworks") + m.cursor.View()
}

