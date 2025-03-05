package controllers

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jorabekkamolov/Turnerko/model"
)

type Updatable interface {
	Update(msg tea.Msg) (tea.Model, tea.Cmd)
}

func (m model.MenuTopicModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Bu yerda modelni yangilash kodini yozing
	return m, nil
}

func (m model.MenuTasksModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Bu yerda modelni yangilash kodini yozing
	return m, nil
}

func (m model.TextEditorModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Bu yerda modelni yangilash kodini yozing
	return m, nil
}
