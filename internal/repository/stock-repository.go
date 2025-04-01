package repository

import (
	"context"
	"errors"
	"time"

	"github.com/ruegerj/stock-sight/internal/db"
	"github.com/ruegerj/stock-sight/internal/queries"
)

var emptyStock = queries.TrackedStock{}

type StockRepository interface {
	GetTrackedStocks(ctx context.Context) ([]queries.TrackedStock, error)
	AddTrackedStock(ctx context.Context, ticker string, date time.Time) (queries.TrackedStock, error)
}

func NewSqlcStockRepository(connection db.DbConnection) StockRepository {
	return &SqlcStockRepository{
		queries: queries.New(connection.Database()),
	}
}

type SqlcStockRepository struct {
	queries *queries.Queries
}

func (ssr *SqlcStockRepository) GetTrackedStocks(ctx context.Context) ([]queries.TrackedStock, error) {
	return ssr.queries.ListTrackedStocks(ctx)
}

func (ssr *SqlcStockRepository) AddTrackedStock(ctx context.Context, ticker string, date time.Time) (queries.TrackedStock, error) {
	if ticker == "" {
		return emptyStock, errors.New("supply a valid stock ticker")
	}

	createParams := queries.AddTrackedStockParams{
		Ticker:    ticker,
		DateAdded: date,
	}

	return ssr.queries.AddTrackedStock(ctx, createParams)
}
