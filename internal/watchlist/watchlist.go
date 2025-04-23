package watchlist

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

// TODO: prob refactor this into seperate file & make the entry constructable
type WatchlistEntry struct {
	Ticker string
	Name   string
	Ppu    float64
	//TODO: fix this into a enum
	Currency string
	Amount   int
}

// TODO: fix hardcoded sizing
// table columns
var columns = []table.Column{
	{Title: "Ticker", Width: 20},
	{Title: "Name", Width: 20},
	{Title: "Price Per Share", Width: 20},
	{Title: "Currency", Width: 20},
	{Title: "Amount", Width: 20},
}

type Watchlist struct {
	watchlist table.Model
	minWidth  int
	minHeight int
}

// create a new watchlist
func New() Watchlist {
	w := Watchlist{
		watchlist: table.New(table.WithColumns(columns)),
	}

	//default to default table styles if no styles added
	s := table.DefaultStyles()
	w.watchlist.SetStyles(s)

	return w
}

// str representation of the watchlist
func (w *Watchlist) View() string {
	return w.watchlist.View()
}

// focus the watchlist
func (w *Watchlist) Focus() {
	w.watchlist.Focus()
}

// register input updates
func (w *Watchlist) Update(msg tea.Msg) {
	t, _ := w.watchlist.Update(msg)
	w.watchlist = t
}

// set the style of the watchlist
func (w *Watchlist) SetStyles(styles table.Styles) {
	w.watchlist.SetStyles(styles)
}

// add a entry to the watchlist
func (w *Watchlist) AddWatchlistEntry(entry WatchlistEntry) {
	r := w.watchlist.Rows()

	nr := []string{
		entry.Ticker,
		entry.Name,
		fmt.Sprintf("%.2f", entry.Ppu),
		entry.Currency,
		strconv.Itoa(entry.Amount),
	}

	r = append(r, nr)

	w.watchlist.SetRows(r)
}

// useful for loadup
func (w *Watchlist) SetWatchlistEntries(entries []WatchlistEntry) {
	w.watchlist.SetRows([]table.Row{})
	for _, e := range entries {
		w.AddWatchlistEntry(e)
	}
}

// retrive selected entry as WatchListEntry
func (w *Watchlist) SelectedEntry() WatchlistEntry {
	selected := w.watchlist.SelectedRow()

	//world class error handling
	amount, _ := strconv.Atoi(selected[4])
	ppu, _ := strconv.ParseFloat(selected[2], 64)

	entry := WatchlistEntry{
		Ticker:   selected[0],
		Name:     selected[1],
		Ppu:      ppu,
		Currency: selected[3],
		Amount:   amount,
	}

	return entry
}

// width
func (w *Watchlist) SetWidth(width int) error {
	if width < w.minWidth {
		return fmt.Errorf("width can't be less than min width")
	}

	//TODO: refactor later
	columnWidth := width / len(columns)
	columns := []table.Column{
		{Title: "Ticker", Width: columnWidth},
		{Title: "Name", Width: columnWidth},
		{Title: "Price Per Share", Width: columnWidth},
		{Title: "Currency", Width: columnWidth},
		{Title: "Amount", Width: columnWidth},
	}

	w.watchlist.SetColumns(columns)
	w.watchlist.SetWidth(width)

	return nil
}

func (w *Watchlist) SetMinWidth(min int) {
	w.minWidth = min
}

func (w *Watchlist) MinWidth() int {
	return w.minWidth
}

// height
func (w *Watchlist) SetHeight(height int) error {
	if height < w.minHeight {
		return fmt.Errorf("height can't be less than min height")
	}
	w.watchlist.SetHeight(height)
	return nil
}

func (w *Watchlist) SetMinHeight(min int) {
	w.minHeight = min
}

func (w *Watchlist) MinHeight() int {
	return w.minHeight
}
