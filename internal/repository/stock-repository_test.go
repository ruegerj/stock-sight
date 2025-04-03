package repository_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSqlcStockRepository_AddTrackedStock(t *testing.T) {
	ticker := "AAPL"
	date := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)

	stock, err := stockRepo.AddTrackedStock(t.Context(), ticker, date)

	assert.NoError(t, err)
	assert.Equal(t, ticker, stock.Ticker)
	assert.NotZero(t, stock.DateAdded)
}

func TestSqlcStockRepository_AddTrackedStock_FailForEmptyTicker(t *testing.T) {
	ticker := ""
	date := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)

	_, err := stockRepo.AddTrackedStock(t.Context(), ticker, date)

	assert.ErrorContains(t, err, "supply a valid stock ticker")
}

func TestSqlcStockRepository_GetTrackedStocks(t *testing.T) {
	_, err := stockRepo.AddTrackedStock(t.Context(), "AAPL", time.Now())
	assert.NoError(t, err)
	_, err = stockRepo.AddTrackedStock(t.Context(), "GOOGL", time.Now())
	assert.NoError(t, err)

	stocks, err := stockRepo.GetTrackedStocks(t.Context())

	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(stocks), 2)
	// check for descending order
	assert.Equal(t, "GOOGL", stocks[0].Ticker)
	assert.Equal(t, "AAPL", stocks[1].Ticker)
}
