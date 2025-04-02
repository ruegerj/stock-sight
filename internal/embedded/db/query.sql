-- name: GetTransaction :one
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
    1;

-- name: ListTransactions :many
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
    date desc;

-- name: CreateTransaction :one
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
    ) returning *;

-- name: UpdateTransaction :exec
update transactions
set
    ticker = ?,
    price_per_unit = ?,
    currency = ?,
    amount = ?,
    date = ?,
    is_buy = ?
where
    id = ?;

-- name: DeleteTransaction :exec
delete from transactions
where
    id = ?;

-- name: AddTrackedStock :one
insert into
    tracked_stocks (ticker, date_added)
values
    (?, ?) returning *;

-- name: ListTrackedStocks :many
select
    *
from
    tracked_stocks
order by
    date_added desc;

-- name: GetTrackedStockByTicker :one
select
    id,
    ticker,
    date_added
from
    tracked_stocks
where
    ticker = ?
limit
    1;
