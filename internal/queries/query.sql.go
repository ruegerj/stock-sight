// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: query.sql

package queries

import (
	"context"
	"time"
)

const addTrackedStock = `-- name: AddTrackedStock :one
insert into
    tracked_stocks (ticker, date_added)
values
    (?, ?) returning id, ticker, date_added
`

type AddTrackedStockParams struct {
	Ticker    string
	DateAdded time.Time
}

func (q *Queries) AddTrackedStock(ctx context.Context, arg AddTrackedStockParams) (TrackedStock, error) {
	row := q.db.QueryRowContext(ctx, addTrackedStock, arg.Ticker, arg.DateAdded)
	var i TrackedStock
	err := row.Scan(&i.ID, &i.Ticker, &i.DateAdded)
	return i, err
}

const createTransaction = `-- name: CreateTransaction :one
insert into
    transactions (
        ticker,
        price_per_unit,
        currency,
        amount,
        date,
        is_buy
    )
values
    (
        ?, -- ticker
        ?, -- price_per_unit
        ?, -- currency
        ?, -- amount
        ?, -- date
        ? -- is_buy
    ) returning id, ticker, price_per_unit, currency, amount, date, is_buy
`

type CreateTransactionParams struct {
	Ticker       string
	PricePerUnit float64
	Currency     string
	Amount       float64
	Date         time.Time
	IsBuy        bool
}

func (q *Queries) CreateTransaction(ctx context.Context, arg CreateTransactionParams) (Transaction, error) {
	row := q.db.QueryRowContext(ctx, createTransaction,
		arg.Ticker,
		arg.PricePerUnit,
		arg.Currency,
		arg.Amount,
		arg.Date,
		arg.IsBuy,
	)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.Ticker,
		&i.PricePerUnit,
		&i.Currency,
		&i.Amount,
		&i.Date,
		&i.IsBuy,
	)
	return i, err
}

const deleteTransaction = `-- name: DeleteTransaction :exec
delete from transactions
where
    id = ?
`

func (q *Queries) DeleteTransaction(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTransaction, id)
	return err
}

const getTrackedStockByTicker = `-- name: GetTrackedStockByTicker :one
select
    id,
    ticker,
    date_added
from
    tracked_stocks
where
    ticker = ?
limit
    1
`

func (q *Queries) GetTrackedStockByTicker(ctx context.Context, ticker string) (TrackedStock, error) {
	row := q.db.QueryRowContext(ctx, getTrackedStockByTicker, ticker)
	var i TrackedStock
	err := row.Scan(&i.ID, &i.Ticker, &i.DateAdded)
	return i, err
}

const getTransaction = `-- name: GetTransaction :one
select
    id,
    ticker,
    price_per_unit,
    currency,
    amount,
    date,
    is_buy
from
    transactions
where
    id = ?
limit
    1
`

func (q *Queries) GetTransaction(ctx context.Context, id int64) (Transaction, error) {
	row := q.db.QueryRowContext(ctx, getTransaction, id)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.Ticker,
		&i.PricePerUnit,
		&i.Currency,
		&i.Amount,
		&i.Date,
		&i.IsBuy,
	)
	return i, err
}

const listTrackedStocks = `-- name: ListTrackedStocks :many
select
    id, ticker, date_added
from
    tracked_stocks
order by
    date_added desc
`

func (q *Queries) ListTrackedStocks(ctx context.Context) ([]TrackedStock, error) {
	rows, err := q.db.QueryContext(ctx, listTrackedStocks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TrackedStock
	for rows.Next() {
		var i TrackedStock
		if err := rows.Scan(&i.ID, &i.Ticker, &i.DateAdded); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTransactions = `-- name: ListTransactions :many
select
    id,
    ticker,
    price_per_unit,
    currency,
    amount,
    date,
    is_buy
from
    transactions
order by
    date desc
`

func (q *Queries) ListTransactions(ctx context.Context) ([]Transaction, error) {
	rows, err := q.db.QueryContext(ctx, listTransactions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transaction
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.ID,
			&i.Ticker,
			&i.PricePerUnit,
			&i.Currency,
			&i.Amount,
			&i.Date,
			&i.IsBuy,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTransaction = `-- name: UpdateTransaction :exec
update transactions
set
    ticker = ?,
    price_per_unit = ?,
    currency = ?,
    amount = ?,
    date = ?,
    is_buy = ?
where
    id = ?
`

type UpdateTransactionParams struct {
	Ticker       string
	PricePerUnit float64
	Currency     string
	Amount       float64
	Date         time.Time
	IsBuy        bool
	ID           int64
}

func (q *Queries) UpdateTransaction(ctx context.Context, arg UpdateTransactionParams) error {
	_, err := q.db.ExecContext(ctx, updateTransaction,
		arg.Ticker,
		arg.PricePerUnit,
		arg.Currency,
		arg.Amount,
		arg.Date,
		arg.IsBuy,
		arg.ID,
	)
	return err
}
