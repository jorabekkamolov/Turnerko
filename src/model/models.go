package models

type MenuTopicModel struct {
	choices  []string
	cursor   int
	selected map[int]string
}

type MenuTasksModel struct {
	choices  []string
	cursor   int
	topic    string
	selected map[int]string
}

type TextEditorModel struct {
	filepath string
	content  []byte
	cursor   int
}
