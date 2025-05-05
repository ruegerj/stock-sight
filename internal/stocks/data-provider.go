package stocks

type Timespan int

const (
	LAST_DAY Timespan = iota
	LAST_WEEK
	LAST_MONTH
	LAST_YEAR
	YEAR_TO_DAY
)

type StockDataProvider interface {
	ProvideFor(ticker string, timespan Timespan, currency string) ([]StockDataPoint, error)
}
