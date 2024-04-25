package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	currentModule tea.Model
}

func main() {
	initialModel := model{currentModule: NewMainMenu()}
	p := tea.NewProgram(initialModel)
	if _, err := p.Run(); err != nil {

		fmt.Println("Error starting program:", err)
		os.Exit(1)
	}
}

func (m model) Init() tea.Cmd {
	return m.currentModule.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.currentModule, cmd = m.currentModule.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.currentModule.View()
}

func renderMenu(options []string, cursor int) string {
	var sb strings.Builder
	for i, option := range options {
		if i == cursor {
			sb.WriteString(fmt.Sprintf("> %s\n", option))
		} else {
			sb.WriteString(fmt.Sprintf("  %s\n", option))
		}
	}
	return sb.String()
}
