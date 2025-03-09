package model

import (
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/common-nighthawk/go-figure"
)

func (m model) View() string {
	var s string

	if m.choices[m.cursor] == "Topic" {
		var styledChoices []string
		greenStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#28DF99"))
		purpleStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#9400D3"))
		fig1 := greenStyle.Render(figure.NewFigure("School 21", "slant", true).String())
		fig2 := purpleStyle.Render(figure.NewFigure("Turnerko", "slant", true).String())

		style1, style2 := m.styleTopic()
		styledChoices = make([]string, len(m.menuTopicModel.choices))
		for i, choice := range m.menuTopicModel.choices {
			if i == m.menuTopicModel.cursor {
				styledChoices[i] = style1.Render(choice)
			} else {
				styledChoices[i] = style2.Render(choice)
			}
		}
		asciiBlock := lipgloss.JoinVertical(lipgloss.Center, fig1, fig2)
		s = lipgloss.JoinVertical(lipgloss.Center, lipgloss.JoinHorizontal(lipgloss.Center, styledChoices...), asciiBlock)

	} else if m.choices[m.cursor] == "Tasks" {
		styledChoices := make([]string, len(m.menuTasksModel.choices))
		style1, style2 := m.styleTask()
		style3 := m.styleTaskBorder()
		style4 := m.styleText()

		for i, choice := range m.menuTasksModel.choices {
			if i == m.menuTasksModel.cursor {
				styledChoices[i] = style1.Render("> " + choice)
			} else {
				styledChoices[i] = style2.Render("  " + choice)
			}
		}
		list1 := style3.Render(lipgloss.JoinVertical(lipgloss.Center, styledChoices...))
		str, _ := glamour.Render(string(m.textEditorModel.content), "dark")
		list2 := style3.Render(style4.Render(lipgloss.JoinVertical(lipgloss.Center, str)))
		s = lipgloss.JoinHorizontal(lipgloss.Top, list1, list2)

	} else if m.choices[m.cursor] == "Editor" {
		styledChoices := make([]string, len(m.menuTasksModel.choices))
		style1, style2 := m.styleTask()
		style3 := m.styleTaskBorder()

		for i, choice := range m.menuTasksModel.choices {
			if i == m.menuTasksModel.cursor {
				styledChoices[i] = style1.Render("> " + choice)
			} else {
				styledChoices[i] = style2.Render("  " + choice)
			}
		}

		var s1 string
		if m.textEditorModel.cursor < len(m.textEditorModel.content) &&
			m.textEditorModel.content[m.textEditorModel.cursor] == '\t' {
			s1 = string(m.textEditorModel.content)
		} else {
			s1 = string(m.textEditorModel.content[:m.textEditorModel.cursor]) +
				"|" +
				string(m.textEditorModel.content[m.textEditorModel.cursor:])
		}
		str, _ := glamour.Render(s1, "dark")
		list1 := style3.Render(lipgloss.JoinVertical(lipgloss.Center, styledChoices...))
		list2 := style3.Render(lipgloss.JoinVertical(lipgloss.Center, str))
		s = lipgloss.JoinHorizontal(lipgloss.Top, list1, list2)
	}

	return s
}
