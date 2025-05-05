package stocks

import "time"

type StockDataPoint struct {
	Timestamp time.Time
	Price     float64
	Currency  string
}
