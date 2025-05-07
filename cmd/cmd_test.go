package cmd_test

import (
	"context"
	"testing"
	"time"

	"github.com/ruegerj/stock-sight/internal/queries"
	"github.com/ruegerj/stock-sight/internal/stocks"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/mock"
)

type MockStockRepository struct {
	mock.Mock
}

func (m *MockStockRepository) AddTrackedStock(ctx context.Context, ticker string, date time.Time) (queries.TrackedStock, error) {
	args := m.Called(ctx, ticker, date)
	return args.Get(0).(queries.TrackedStock), args.Error(1)
}

func (m *MockStockRepository) GetTrackedStocks(ctx context.Context) ([]queries.TrackedStock, error) {
	args := m.Called(ctx)
	return args.Get(0).([]queries.TrackedStock), args.Error(1)
}

type MockTerminalAccessor struct {
	mock.Mock
}

func (mta *MockTerminalAccessor) ResolveDimensions() (width int, height int) {
	args := mta.Called()
	return args.Get(0).(int), args.Get(1).(int)
}

type MockStockDataProvider struct {
	mock.Mock
}

func (msdp *MockStockDataProvider) ProvideFor(ticker string, timespan stocks.Timespan, currency string) ([]stocks.StockDataPoint, error) {
	args := msdp.Called(ticker, timespan, currency)
	return args.Get(0).([]stocks.StockDataPoint), args.Error(1)
}

func applyFlagsTo(t *testing.T, cmd *cobra.Command, flags map[string]string) {
	t.Helper()

	for key, value := range flags {
		flag := cmd.Flag(key)
		if flag == nil {
			t.Fatalf("failed to set flag %q", key)
		}

		flag.Value.Set(value)
	}
}

func runCobraCmd(t *testing.T, cmd *cobra.Command, args []string, flags map[string]string) error {
	t.Helper()
	applyFlagsTo(t, cmd, flags)

	if cmd.PreRunE != nil {
		err := cmd.PreRunE(cmd, args)
		if err != nil {
			return err
		}

	}

	if cmd.RunE != nil {
		err := cmd.RunE(cmd, args)
		if err != nil {
			return err
		}
	}

	if cmd.PostRunE != nil {
		err := cmd.PostRunE(cmd, args)
		if err != nil {
			return err
		}
	}

	return nil
}
