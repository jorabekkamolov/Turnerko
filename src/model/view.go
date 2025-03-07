package model

import (
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
	}

	return s
}
