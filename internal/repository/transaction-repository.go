package repository

import (
	"context"
	"errors"
	"github.com/ruegerj/stock-sight/internal/db"
	"github.com/ruegerj/stock-sight/internal/queries"
	"time"
)

var empty = queries.Transaction{}

type TransactionRepository interface {
	GetById(ctx context.Context, id int64) (queries.Transaction, error)
	GetAll(ctx context.Context) ([]queries.Transaction, error)
	Create(ctx context.Context, ticker string, amount float64, currency string, ppu float64, date time.Time, isBuy bool) (queries.Transaction, error)
	Update(ctx context.Context, transaction queries.Transaction) error
	Delete(ctx context.Context, id int64) error
}

func NewSqlcTransactionRepository(connection db.DbConnection) TransactionRepository {
	return &SqlcTransactionRepository{
		queries: queries.New(connection.Database()),
	}
}

type SqlcTransactionRepository struct {
	queries *queries.Queries
}

func (sar *SqlcTransactionRepository) GetAll(ctx context.Context) ([]queries.Transaction, error) {
	return sar.queries.ListTransactions(ctx)
}

func (sar *SqlcTransactionRepository) GetById(ctx context.Context, id int64) (queries.Transaction, error) {
	if id < 0 {
		return empty, errors.New("supply a valid transaction id")
	}

	return sar.queries.GetTransaction(ctx, id)
}

func (sar *SqlcTransactionRepository) Create(ctx context.Context, ticker string, amount float64, currency string, ppu float64, date time.Time, isBuy bool) (queries.Transaction, error) {

	createParams := queries.CreateTransactionParams{
		Ticker:       ticker,
		Amount:       amount,
		Currency:     currency,
		PricePerUnit: ppu,
		Date:         date,
		IsBuy:        isBuy,
	}

	return sar.queries.CreateTransaction(ctx, createParams)
}

func (sar *SqlcTransactionRepository) Update(ctx context.Context, transaction queries.Transaction) error {

	updateParams := queries.UpdateTransactionParams{
		Ticker:       transaction.Ticker,
		Amount:       transaction.Amount,
		Currency:     transaction.Currency,
		PricePerUnit: transaction.PricePerUnit,
		Date:         transaction.Date,
		IsBuy:        transaction.IsBuy,
	}

	return sar.queries.UpdateTransaction(ctx, updateParams)
}

func (sar *SqlcTransactionRepository) Delete(ctx context.Context, id int64) error {
	if id < 0 {
		return errors.New("supply a valid transaction id")
	}

	return sar.queries.DeleteTransaction(ctx, id)
}
