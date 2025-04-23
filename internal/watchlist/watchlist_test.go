package watchlist_test

import (
	"testing"

	"github.com/ruegerj/stock-sight/internal/watchlist"
)

func TestDimensions(t *testing.T) {
	w := watchlist.New()

	tests := []struct {
		name      string
		width     int
		height    int
		minWidth  int
		minHeight int
		wantErr   bool
	}{
		{
			name:      "valid dimensions",
			width:     100,
			height:    50,
			minWidth:  30,
			minHeight: 10,
			wantErr:   false,
		},
		{
			name:      "width below minimum",
			width:     20,
			height:    50,
			minWidth:  30,
			minHeight: 10,
			wantErr:   true,
		},
		{
			name:      "height below minimum",
			width:     100,
			height:    5,
			minWidth:  30,
			minHeight: 10,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w.SetMinWidth(tt.minWidth)
			w.SetMinHeight(tt.minHeight)

			errWidth := w.SetWidth(tt.width)
			errHeight := w.SetHeight(tt.height)

			if (errWidth != nil || errHeight != nil) != tt.wantErr {
				t.Errorf("SetWidth/Height() error = %v/%v, wantErr %v",
					errWidth, errHeight, tt.wantErr)
			}

			if w.MinWidth() != tt.minWidth {
				t.Errorf("MinWidth() = %v, want %v", w.MinWidth(), tt.minWidth)
			}

			if w.MinHeight() != tt.minHeight {
				t.Errorf("MinHeight() = %v, want %v", w.MinHeight(), tt.minHeight)
			}
		})
	}
}

func TestTableSelectedEntry(t *testing.T) {
	w := watchlist.New()
	testEntries := []watchlist.WatchlistEntry{
		{
			Ticker:   "TEST",
			Name:     "Test Stock",
			Ppu:      42.42,
			Currency: "USD",
			Amount:   100,
		},
	}

	w.SetWatchlistEntries(testEntries)

	// Test selected entry matches input
	selected := w.SelectedEntry()
	if selected.Ticker != testEntries[0].Ticker {
		t.Errorf("SelectedEntry() Ticker = %v, want %v",
			selected.Ticker, testEntries[0].Ticker)
	}
}
