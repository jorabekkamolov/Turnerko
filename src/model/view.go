package model

import (
	"fmt"
)

func (m model) View() string {
	var s string
	var s1 string
	if m.choices[m.cursor] == "Topic" {
		for _, choice := range m.choices {
			s += fmt.Sprintf("%s\n", choice)
		}
	} else if m.choices[m.cursor] == "Tasks" {
		for _, choice := range m.menuTasksModel.choices {
			s += fmt.Sprintf("%s\n", choice)
		}
		s += string(m.textEditorModel.content)
	} else if m.choices[m.cursor] == "Editor" {
		if m.textEditorModel.cursor < len(m.textEditorModel.content) &&
			m.textEditorModel.content[m.textEditorModel.cursor] == '\t' {
			s1 = string(m.textEditorModel.content)
		} else {
			s1 = string(m.textEditorModel.content[:m.textEditorModel.cursor]) +
				"|" +
				string(m.textEditorModel.content[m.textEditorModel.cursor:])
		}
	}
	return s1
}
