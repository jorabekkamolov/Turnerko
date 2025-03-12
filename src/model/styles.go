package model

import "github.com/charmbracelet/lipgloss"

func (m *model) styleTopic() (lipgloss.Style, lipgloss.Style) {

	return lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#28DF99")).
			Background(lipgloss.Color("#28DF99")).
			Foreground(lipgloss.Color("#062743")).
			Bold(true).
			Padding(2, 6).
			Italic(true).
			Margin(1, 2).
			Align(lipgloss.Center),

		lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#9400D3")).
			Background(lipgloss.Color("#9400D3")).
			Foreground(lipgloss.Color("#062743")).
			Bold(true).
			Faint(true).
			Italic(true).
			Padding(1, 5).
			Margin(1, 2).
			Align(lipgloss.Center)
}

func (m *model) styleTask() (lipgloss.Style, lipgloss.Style) {

	return lipgloss.NewStyle().
			Foreground(lipgloss.Color("#28DF99")).
			Bold(true).
			Italic(true),

		lipgloss.NewStyle().
			Foreground(lipgloss.Color("#9400D3")).
			Bold(true).
			Faint(true).
			Italic(true)
}

func (m *model) styleTaskBorder() lipgloss.Style {
	return lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#28DF99")).
		Padding(1, 3).
		Margin(1, 1)
}

func (m *model) styleText() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FCFFC9"))
}
