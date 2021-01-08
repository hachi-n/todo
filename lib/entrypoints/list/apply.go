package list

import (
	"fmt"
	"github.com/charmbracelet/bubbletea"
	"github.com/hachi-n/todo/lib/models"
	"os"
)

func Apply() error {
	currentTodoList := models.LoadTodoList()
	// currentTodoList defined?
	if currentTodoList == nil {
		return fmt.Errorf("No TodoList.")
	}

	p := tea.NewProgram(currentTodoList.ViewModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	return nil
}

