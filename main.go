package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/filepicker"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	currentModule tea.Model
	filepicker    filepicker.Model
}

func main() {
	fp := filepicker.New()
	fp.CurrentDirectory, _ = os.UserHomeDir()

	// Create the initial model, setting the filepicker before passing the model to NewMainMenu
	initialModel := &model{ // Make sure initialModel is a pointer to model
		filepicker: fp,
	}

	// Now pass initialModel to NewMainMenu, which needs a reference to model
	initialModel.currentModule = NewMainMenu(initialModel)

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
