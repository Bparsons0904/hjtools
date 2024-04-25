package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type MainMenu struct {
	cursor int
	menu   []string
}

func NewMainMenu() tea.Model {
	return &MainMenu{
		menu: []string{"Eligibility File Tool", "Quit"},
	}
}

func (m MainMenu) Init() tea.Cmd {
	return nil
}

func (m MainMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}
		switch msg.String() {
		case "enter":
			switch m.cursor {
			case 0:
				return &EligibilityFileTool{}, nil
			case 1:
				fmt.Println("Exiting program.")
				return m, tea.Quit
			}
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
			return m, nil
		case "down", "j":
			if m.cursor < len(m.menu)-1 {
				m.cursor++
			}
			return m, nil
		}
	}
	return m, nil
}

func (m MainMenu) View() string {
	return "Main Menu:\n" + renderMenu(m.menu, m.cursor)
}
