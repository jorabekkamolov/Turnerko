package models

import (
	"fmt"
	"os"
)

func initMenuTopicModel() MenuTopicModel {
	return MenuTopicModel{
		choices: []string{
			"IF", "LOOP", "NESTED LOOP",
			"ARRAY", "STRING",
		},
		selected: make(map[int]string),
	}
}

func initMenuTasksModel(topic string) MenuTasksModel {
	tasks := getTopicTasks(topic)
	return MenuTasksModel{
		topic:    topic,
		choices:  tasks,
		selected: make(map[int]string),
	}
}

func inittextEditorModel(filepath string) TextEditorModel {
	data, err := os.ReadFile(filepath)
	if err != nil {
		data = []byte("Yangi fayl. Matnni o'zgartiring...\n")
	}
	return TextEditorModel{
		filepath: filepath,
		content:  []byte(string(data)),
	}
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
