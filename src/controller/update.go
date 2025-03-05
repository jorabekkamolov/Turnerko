package controllers

import (
	tea "github.com/charmbracelet/bubbletea"
	model "github.com/jorabekkamolov/Turnerko/src/model"
)

type Updatable interface {
	Update(msg tea.Msg) (tea.Model, tea.Cmd)
}

func (m model.MenuTopicModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m model.MenuTasksModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m model.TextEditorModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}
