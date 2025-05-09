// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package queries

import (
	"time"
)

type TrackedStock struct {
	ID        int64
	Ticker    string
	DateAdded time.Time
}

type Transaction struct {
	ID           int64
	Ticker       string
	PricePerUnit float64
	Currency     string
	Amount       float64
	Date         time.Time
	IsBuy        bool
}
