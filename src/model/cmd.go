package model

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func Run() {
	p := tea.NewProgram(initModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}
}
