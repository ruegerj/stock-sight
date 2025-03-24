-- name: GetTransaction :one
select *
from transactions
where id = ? limit 1;

-- name: ListTransactions :many
select *
from transactions
order by date desc;


-- name: CreateTransaction :one
insert into transactions (
    ticker,
    price_per_unit,
    currency,
    amount,
    date,
    is_buy
) values (
    ?,  -- ticker
    ?,  -- price_per_unit
    ?,  -- currency
    ?,  -- amount
    ?,  -- date
    ?   -- is_buy
)
returning *;

-- name: UpdateTransaction :exec
update transactions
set ticker = ?,
    price_per_unit = ?,
    currency = ?,
    amount = ?,
    date = ?,
    is_buy = ?
where id = ?;

-- name: DeleteTransaction :exec
delete from transactions
where id = ?;
