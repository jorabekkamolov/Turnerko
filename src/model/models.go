package model

// import "time"

type menuTopicModel struct {
	choices []string
	cursor  int
}

type menuTasksModel struct {
	choices []string
	cursor  int
	topic   string
}

type textEditorModel struct {
	filepath string
	content  []rune
	cursor   int
	// lastTime  time.Time
	// saveStack []int
}

type checkResult struct {
	filepath string
	answer   int
	success  []string
	warning  []string
	errors   []string
}

type model struct {
	choices         []string
	cursor          int
	width           int
	height          int
	menuTopicModel  menuTopicModel
	menuTasksModel  menuTasksModel
	textEditorModel textEditorModel
	checkResult     checkResult
}
