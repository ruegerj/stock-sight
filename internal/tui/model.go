package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ruegerj/stock-sight/internal/watchlist"
)

type Model struct {
	watchlist watchlist.Watchlist
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.watchlist.SetHeight(msg.Height)
		m.watchlist.SetWidth(msg.Width)
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() string {
	return m.watchlist.View()
}

func Start() error {
	entry := watchlist.WatchlistEntry{
		Ticker:   "APPL",
		Name:     "Apple Inc.",
		Ppu:      32.0,
		Currency: "USD",
		Amount:   12,
	}

	m := Model{
		watchlist: watchlist.New(),
	}

	m.watchlist.AddWatchlistEntry(entry)
	p := tea.NewProgram(m, tea.WithAltScreen())
	_, err := p.Run()
	return err
}
