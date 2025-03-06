package model

import (
	"fmt"
)

func (m model) View() string {
	var s string
	if m.choices[m.cursor] == "Topic" {
		for _, choice := range m.choices {
			s += fmt.Sprintf("%s\n", choice)
		}
	} else if m.choices[m.cursor] == "Tasks" {
		for _, choice := range m.menuTasksModel.choices {
			s += fmt.Sprintf("%s\n", choice)
		}
		content := string(m.textEditorModel.content[:m.textEditorModel.cursor]) +
			"|" +
			string(m.textEditorModel.content[m.textEditorModel.cursor:])
		s += content
	} else if m.choices[m.cursor] == "Editor" {
		content := string(m.textEditorModel.content[:m.textEditorModel.cursor]) +
			"|" +
			string(m.textEditorModel.content[m.textEditorModel.cursor:])
		s += content

	}
	return s
}
