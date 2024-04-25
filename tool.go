package main

import tea "github.com/charmbracelet/bubbletea"

type ToolModule interface {
	Init() tea.Cmd
	Update(msg tea.Msg) (ToolModule, tea.Cmd)
	View() string
}
