package model

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd = nil
	switch m.choices[m.cursor] {
	case "Topic":
		m.funcUpdateTopic(msg)
	case "Tasks":
		m.funcUpdateTasks(msg)
	case "Editor":
		cmd = m.funcUpdateEditor(msg)
	}
	return m, cmd
}

func (m *model) funcUpdateTopic(msg tea.Msg) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case ".", "right":
			if m.menuTopicModel.cursor < len(m.menuTopicModel.choices)-1 {
				m.menuTopicModel.cursor++
			}
		case ",", "left":
			if m.menuTopicModel.cursor > 0 {
				m.menuTopicModel.cursor--
			}
		case "enter", " ":
			m.updateTaskModel()
			m.updateEditorModel()
			m.cursor++
		}
	}
}

func (m *model) funcUpdateTasks(msg tea.Msg) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "down", "S", "s":
			if m.menuTasksModel.cursor < len(m.menuTasksModel.choices)-1 {
				m.menuTasksModel.cursor++
				m.updateEditorModel()
			}
		case "up", "W", "w":
			if m.menuTasksModel.cursor > 0 {
				m.menuTasksModel.cursor--
				m.updateEditorModel()
			}
		case "Q", "q":
			m.cursor--
		case "enter", " ":
			m.cursor++
		}
	}
}

func (m *model) funcUpdateEditor(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd = nil
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "right":
			if m.textEditorModel.cursor < len(m.textEditorModel.content) {
				m.textEditorModel.cursor++
			}
		case "left":
			if m.textEditorModel.cursor > 0 {
				m.textEditorModel.cursor--
			}
		case "backspace":
			if m.textEditorModel.cursor > 0 {
				if m.textEditorModel.content[m.textEditorModel.cursor-1] == '\t' ||
					m.textEditorModel.content[m.textEditorModel.cursor-1] == '\n' {
					cmd = tea.ClearScreen
				}
				m.textEditorModel.content = append(
					m.textEditorModel.content[:m.textEditorModel.cursor-1],
					m.textEditorModel.content[m.textEditorModel.cursor:]...,
				)
				m.textEditorModel.cursor--

			}
		case "enter":
			m.textEditorModel.content = append(
				m.textEditorModel.content[:m.textEditorModel.cursor],
				append([]rune{'\n'}, m.textEditorModel.content[m.textEditorModel.cursor:]...)...,
			)
			m.textEditorModel.cursor++
		case "down":
			var back int = 0
			var temp int = m.textEditorModel.cursor

			for temp > 0 && m.textEditorModel.content[temp-1] != '\n' {
				temp--
				back++
			}
			if temp+back >= len(m.textEditorModel.content) {
				break
			}

			for m.textEditorModel.cursor < len(m.textEditorModel.content) &&
				m.textEditorModel.content[m.textEditorModel.cursor] != '\n' {
				m.textEditorModel.cursor++
			}

			if m.textEditorModel.cursor < len(m.textEditorModel.content) {
				m.textEditorModel.cursor++
			}
			for back > 0 && m.textEditorModel.cursor < len(m.textEditorModel.content) &&
				m.textEditorModel.content[m.textEditorModel.cursor] != '\n' {
				m.textEditorModel.cursor++
				back--
			}

		case "up":
			var back int = 0
			var temp int = m.textEditorModel.cursor
			for temp > 0 && m.textEditorModel.content[temp-1] != '\n' {
				temp--
				back++
			}
			if temp == 0 {
				break
			}
			temp--
			for temp > 0 && m.textEditorModel.content[temp-1] != '\n' {
				temp--
			}
			var newCursor = temp
			for back > 0 && newCursor < len(m.textEditorModel.content) &&
				m.textEditorModel.content[newCursor] != '\n' {
				newCursor++
				back--
			}

			m.textEditorModel.cursor = newCursor
		case "tab":
			m.textEditorModel.content = append(
				m.textEditorModel.content[:m.textEditorModel.cursor],
				append([]rune{'\t'}, m.textEditorModel.content[m.textEditorModel.cursor:]...)...,
			)
			m.textEditorModel.cursor++
			cmd = tea.ClearScreen
		case "insert", "delete":
		default:
			m.textEditorModel.content = append(
				m.textEditorModel.content[:m.textEditorModel.cursor],
				append([]rune(msg.String()), m.textEditorModel.content[m.textEditorModel.cursor:]...)...,
			)
			m.textEditorModel.cursor++
		}
	}
	return cmd
}

func (m *model) updateTaskModel() {
	m.menuTasksModel.choices = getTopicTasks(m.menuTopicModel.choices[m.menuTopicModel.cursor])
	m.menuTasksModel.topic = m.menuTopicModel.choices[m.menuTopicModel.cursor]
}

func (m *model) updateEditorModel() {
	filepath := fmt.Sprintf("C/%s/%s.txt",
		m.menuTopicModel.choices[m.menuTopicModel.cursor],
		m.menuTasksModel.choices[m.menuTasksModel.cursor])
	m.textEditorModel.filepath = filepath
	data, err := os.ReadFile(filepath)
	if err != nil {
		data = []byte("Yangi fayl. Matnni o'zgartiring...\n")
	}
	// data = append(data, ' '\)
	m.textEditorModel.content = []rune(string(data))
}

func getTopicTasks(topic string) []string {
	var tasks []string
	var taskCount int

	switch topic {
	case "IF":
		taskCount = 20
	case "LOOP", "NESTED LOOP":
		taskCount = 25
	case "ARRAY":
		taskCount = 20
	case "STRING":
		taskCount = 15
	}

	for i := 1; i <= taskCount; i++ {
		tasks = append(tasks, fmt.Sprintf("%d-Task", i))
	}
	return tasks
}
