package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() string {
	return "Hello, from the TUI!"
}

func Start() error {
	p := tea.NewProgram(Model{}, tea.WithAltScreen())
	_, err := p.Run()
	return err
}
