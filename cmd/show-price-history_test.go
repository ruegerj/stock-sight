package cmd_test

import (
	"testing"
	"time"

	"github.com/ruegerj/stock-sight/cmd"
	"github.com/ruegerj/stock-sight/internal/queries"
	"github.com/ruegerj/stock-sight/internal/stocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestShowPriceHistoryCmd_NoTimespanProvided_YieldError(t *testing.T) {
	// GIVEN
	noTimespanError := "Must supply value for timespan flag"
	stockDataProvider := new(MockStockDataProvider)
	stockRepository := new(MockStockRepository)
	terminalAccessor := new(MockTerminalAccessor)

	// WHEN
	showPriceHistCmd := cmd.ShowPriceHistoryCmd(t.Context(), stockRepository, stockDataProvider, terminalAccessor)
	applyFlagsTo(t, showPriceHistCmd.Command(), map[string]string{"timespan": ""})
	err := showPriceHistCmd.Command().PreRunE(showPriceHistCmd.Command(), []string{})

	// THEN
	assert.EqualError(t, err, noTimespanError)
}

func TestShowPriceHistoryCmd_YieldErrorForInvalidTimespanFlag(t *testing.T) {
	// GIVEN
	invalidTimespanError := "Must supply a valid value for the timespan flag"
	tests := []struct {
		timespan         string
		expectedErrorMsg *string
	}{
		{"d", nil},
		{"w", nil},
		{"m", nil},
		{"y", nil},
		{"y2d", nil},
		{"day", &invalidTimespanError},
		{"x", &invalidTimespanError},
		{"y2m", &invalidTimespanError},
	}

	stockDataProvider := new(MockStockDataProvider)
	stockRepository := new(MockStockRepository)
	terminalAccessor := new(MockTerminalAccessor)

	for _, tt := range tests {
		t.Run(tt.timespan, func(t *testing.T) {
			// GIVEN
			stock := "AAPL"
			flags := map[string]string{"timespan": tt.timespan}
			cmd := cmd.ShowPriceHistoryCmd(t.Context(), stockRepository, stockDataProvider, terminalAccessor)
			applyFlagsTo(t, cmd.Command(), flags)

			// WHEN
			err := cmd.Command().PreRunE(cmd.Command(), []string{stock})

			// THEN
			if tt.expectedErrorMsg != nil {
				assert.EqualError(t, err, *tt.expectedErrorMsg)
				return
			}

			assert.NoError(t, err)
		})
	}
}

func TestShowPriceHistoryCmd_RunForUntrackedStock_YieldError(t *testing.T) {
	// GIVEN
	calledForUntrackedStockError := "Price history can only be shown for tracked stocks"
	stockDataProvider := new(MockStockDataProvider)
	stockRepository := new(MockStockRepository)
	stockRepository.On("GetTrackedStocks", mock.Anything).Return([]queries.TrackedStock{
		{Ticker: "AAPL", DateAdded: time.Now()},
	}, nil).Once()
	terminalAccessor := new(MockTerminalAccessor)

	// WHEN
	untrackedTicker := "TSLA"
	showPriceHistCmd := cmd.ShowPriceHistoryCmd(t.Context(), stockRepository, stockDataProvider, terminalAccessor)
	err := runCobraCmd(t, showPriceHistCmd.Command(), []string{untrackedTicker}, map[string]string{})

	// THEN
	stockRepository.AssertCalled(t, "GetTrackedStocks", mock.Anything)
	assert.EqualError(t, err, calledForUntrackedStockError)
}

func TestShowPriceHistoryCmd_ResolveDataForCorrectTimespan(t *testing.T) {
	// GIVEN
	trackedTicker := "AAPL"
	tests := []struct {
		givenTimespan    string
		expectedTimespan stocks.Timespan
	}{
		{"d", stocks.LAST_DAY},
		{"w", stocks.LAST_WEEK},
		{"m", stocks.LAST_MONTH},
		{"y", stocks.LAST_YEAR},
		{"y2d", stocks.YEAR_TO_DAY},
	}
	stockRepository := new(MockStockRepository)
	stockRepository.On("GetTrackedStocks", mock.Anything).Return([]queries.TrackedStock{
		{Ticker: trackedTicker, DateAdded: time.Now()},
	}, nil).Times(len(tests))
	terminalAccessor := new(MockTerminalAccessor)
	terminalAccessor.On("ResolveDimensions", mock.Anything).Return(100, 50, nil).Times(len(tests))
	stockDataProvider := new(MockStockDataProvider)
	stockDataProvider.On("ProvideFor", mock.Anything, mock.Anything, mock.Anything).
		Return([]stocks.StockDataPoint{}, nil).
		Times(len(tests))

	for _, tt := range tests {
		t.Run(tt.givenTimespan, func(t *testing.T) {
			// GIVEN
			flags := map[string]string{"timespan": tt.givenTimespan}
			showPriceHistCmd := cmd.ShowPriceHistoryCmd(t.Context(), stockRepository, stockDataProvider, terminalAccessor)

			// WHEN
			err := runCobraCmd(t, showPriceHistCmd.Command(), []string{trackedTicker}, flags)

			// THEN
			assert.NoError(t, err)
			stockDataProvider.AssertCalled(t, "ProvideFor", trackedTicker, tt.expectedTimespan, mock.Anything)
		})
	}
}
