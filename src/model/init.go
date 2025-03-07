package model

import (
	tea "github.com/charmbracelet/bubbletea"
)

func initMenuTopicModel() menuTopicModel {
	return menuTopicModel{
		choices: []string{"IF", "LOOP", "NESTED LOOP", "ARRAY", "STRING"},
	}
}

func initModel() model {
	return model{
		choices:         []string{"Topic", "Tasks", "Editor", "Result"},
		menuTopicModel:  initMenuTopicModel(),
		menuTasksModel:  menuTasksModel{},
		textEditorModel: textEditorModel{},
		checkResult:     checkResult{},
	}
}

func (m model) Init() tea.Cmd {
	// Ekranni tozalash va boshlang‘ich interfeysni qaytarish
	return tea.ClearScreen
}
