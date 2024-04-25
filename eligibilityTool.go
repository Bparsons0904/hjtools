package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/filepicker"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	MenuStage = iota
	FilePickerStage
)

const (
	ConvertToCSVOption = iota
	BackOption
	QuitOption
)

type EligibilityFileTool struct {
	picker       filepicker.Model
	menuCursor   int
	currentStage int
	selectedFile string
	err          error
}

func NewEligibilityFileTool() *EligibilityFileTool {
	picker := filepicker.New()
	picker.AllowedTypes = []string{"psv"}
	picker.CurrentDirectory, _ = os.UserHomeDir()

	return &EligibilityFileTool{
		picker:       picker,
		currentStage: MenuStage,
	}
}

type clearErrorMsg struct{}

func clearErrorAfter(t time.Duration) tea.Cmd {
	return tea.Tick(t, func(_ time.Time) tea.Msg {
		return clearErrorMsg{}
	})
}

func (m *EligibilityFileTool) Init() tea.Cmd {
	return m.picker.Init()
}

func (m *EligibilityFileTool) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC || msg.String() == "q" {
			return m, tea.Quit
		}
		if m.currentStage == MenuStage {
			return m.updateMenu(msg)
		} else if m.currentStage == FilePickerStage {
			var cmd tea.Cmd
			m.updateFilePicker(msg)
			return m, cmd
		}
	}
	return m, nil
}

func (m *EligibilityFileTool) updateMenu(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.Type {
	case tea.KeyEnter:
		switch m.menuCursor {
		case ConvertToCSVOption:
			m.currentStage = FilePickerStage
			// Attempt to reinitialize the file picker to ensure it is in a ready state.
			return m, nil
		case BackOption:
			// Handle the back option if necessary
		case QuitOption:
			return m, tea.Quit
		}
	case tea.KeyUp:
		if m.menuCursor > 0 {
			m.menuCursor--
		}
	case tea.KeyDown:
		if m.menuCursor < 2 {
			m.menuCursor++
		}
	}
	return m, nil
}

func (m *EligibilityFileTool) updateFilePicker(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.picker, cmd = m.picker.Update(msg)
	if didSelect, path := m.picker.DidSelectFile(msg); didSelect {
		m.selectedFile = path
	}

	if didSelect, path := m.picker.DidSelectDisabledFile(msg); didSelect {
		m.err = errors.New(path + " is not valid.")
		m.selectedFile = ""
		return m, tea.Batch(cmd, clearErrorAfter(2*time.Second))
	}
	return m, cmd
}

func (m *EligibilityFileTool) View() string {
	if m.currentStage == MenuStage {
		return fmt.Sprintf(
			"\nChoose an option:\n%s Convert to CSV\n%s Back\n%s Quit\n",
			ifCursor(m.menuCursor == ConvertToCSVOption), ifCursor(m.menuCursor == BackOption), ifCursor(m.menuCursor == QuitOption),
		)
	} else if m.currentStage == FilePickerStage {
		log.Println("FilePickerStage")
		return m.picker.View() // Make sure this correctly shows the file picker
	}
	return "Unexpected stage"
}

func ifCursor(cond bool) string {
	if cond {
		return ">"
	}
	return " "
}
