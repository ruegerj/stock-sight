package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/ruegerj/stock-sight/internal/db"
	"github.com/ruegerj/stock-sight/internal/repository"
	"github.com/stretchr/testify/assert"
)

var dbConn db.DbConnection
var repo repository.StockRepository

func TestMain(m *testing.M) {
	ctx := context.Background()
	dbConn = db.NewInMemorySQLite(ctx)
	repo = repository.NewSqlcStockRepository(dbConn)
}

func TestSqlcStockRepository_AddTrackedStock(t *testing.T) {
	ticker := "AAPL"
	date := time.Now()

	stock, err := repo.AddTrackedStock(context.Background(), ticker, date)

	assert.NoError(t, err)
	assert.Equal(t, ticker, stock.Ticker)
	assert.NotZero(t, stock.DateAdded)
}

func TestSqlcStockRepository_GetTrackedStocks(t *testing.T) {
	_, err := repo.AddTrackedStock(context.Background(), "AAPL", time.Now())
	assert.NoError(t, err)
	_, err = repo.AddTrackedStock(context.Background(), "GOOGL", time.Now())
	assert.NoError(t, err)

	stocks, err := repo.GetTrackedStocks(context.Background())

	assert.NoError(t, err)
	assert.Len(t, stocks, 2)
	assert.Equal(t, "AAPL", stocks[0].Ticker)
	assert.Equal(t, "GOOGL", stocks[1].Ticker)
}
