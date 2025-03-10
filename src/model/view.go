package model

import (
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/common-nighthawk/go-figure"
)

func (m model) View() string {
	var output string

	switch m.choices[m.cursor] {
	case "Topic":
		selectedStyle, defaultStyle := m.styleTopic()
		titleStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#28DF99"))
		subtitleStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#9400D3"))
		title := titleStyle.Render(figure.NewFigure("School 21", "slant", true).String())
		subtitle := subtitleStyle.Render(figure.NewFigure("Turnerko", "slant", true).String())

		styledTopics := styleChoiceList(m.menuTopicModel.choices, m.menuTopicModel.cursor, selectedStyle, defaultStyle)
		asciiBlock := lipgloss.JoinVertical(lipgloss.Center, title, subtitle)
		output = lipgloss.JoinVertical(lipgloss.Center, lipgloss.JoinHorizontal(lipgloss.Center, styledTopics...), asciiBlock)

	case "Tasks":
		selectedStyle, defaultStyle := m.styleTask()
		styledTasks := styleChoiceList(m.menuTasksModel.choices, m.menuTasksModel.cursor, selectedStyle, defaultStyle)
		taskContent, _ := glamour.Render(string(m.textEditorModel.content), "dark")

		taskList := m.styleTaskBorder().Render(lipgloss.JoinVertical(lipgloss.Center, styledTasks...))
		contentBox := m.styleTaskBorder().Render(m.styleText().Render(lipgloss.JoinVertical(lipgloss.Center, taskContent)))

		output = lipgloss.JoinHorizontal(lipgloss.Top, taskList, contentBox)

	case "Editor":
		selectedStyle, defaultStyle := m.styleTask()
		styledTasks := styleChoiceList(m.menuTasksModel.choices, m.menuTasksModel.cursor, selectedStyle, defaultStyle)
		editorContent := addCursorMarker(m.textEditorModel.content, m.textEditorModel.cursor)
		formattedContent, _ := glamour.Render(editorContent, "dark")

		taskList := m.styleTaskBorder().Render(lipgloss.JoinVertical(lipgloss.Center, styledTasks...))
		editorBox := m.styleTaskBorder().Render(lipgloss.JoinVertical(lipgloss.Center, formattedContent))
		output = lipgloss.JoinHorizontal(lipgloss.Top, taskList, editorBox)
	}

	return output
}

func styleChoiceList(choices []string, cursorIndex int, selectedStyle lipgloss.Style, defaultStyle lipgloss.Style) []string {
	styledList := make([]string, len(choices))

	for i, choice := range choices {
		if i == cursorIndex {
			styledList[i] = selectedStyle.Render(choice)
		} else {
			styledList[i] = defaultStyle.Render(choice)
		}
	}
	return styledList
}

func addCursorMarker(content []rune, cursorIndex int) string {
	return string(content[:cursorIndex]) + "|" + string(content[cursorIndex:])
}
