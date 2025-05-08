package stocks

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

var _ StockDataProvider = (*FakeDataGenerator)(nil)

type FakeDataGenerator struct {
	rand *rand.Rand
}

func NewFakeDataGenerator() StockDataProvider {
	return &FakeDataGenerator{
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (g *FakeDataGenerator) ProvideFor(ticker string, timespan Timespan, currency string) ([]StockDataPoint, error) {
	now := time.Now()

	switch timespan {
	case LAST_DAY:
		return g.generateData(ticker, now.Add(-24*time.Hour), now, 5*time.Minute, currency), nil
	case LAST_WEEK:
		return g.generateData(ticker, now.Add(-7*24*time.Hour), now, 1*time.Hour, currency), nil
	case LAST_MONTH:
		return g.generateData(ticker, now.Add(-30*24*time.Hour), now, 4*time.Hour, currency), nil
	case LAST_YEAR:
		return g.generateData(ticker, now.Add(-365*24*time.Hour), now, 12*time.Hour, currency), nil
	case YEAR_TO_DAY:
		startOfYear := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
		return g.generateData(ticker, startOfYear, now, 12*time.Hour, currency), nil
	default:
		return nil, fmt.Errorf("unknown timespan option: %d", timespan)
	}
}

func (g *FakeDataGenerator) generateData(ticker string, start, end time.Time, interval time.Duration, currency string) []StockDataPoint {
	result := make([]StockDataPoint, 0)
	initialPrice, volatitlity := determineStockDataFor(ticker)
	price := initialPrice

	for t := start; t.Before(end); t = t.Add(interval) {
		if t.Weekday() == time.Saturday || t.Weekday() == time.Sunday {
			continue
		}

		hour, min := t.Hour(), t.Minute()
		isMarketOpen := hour > 9 || (hour == 9 && min >= 30)
		notMarketClosed := hour < 16

		if isMarketOpen && notMarketClosed {
			randomFactor := g.rand.NormFloat64() * volatitlity

			price *= (1 + randomFactor)

			if price < 1.0 {
				price = 1.0 + g.rand.Float64()
			}

			result = append(result, StockDataPoint{
				Timestamp: t,
				Price:     roundToTwoDecimals(price),
				Currency:  currency,
			})
		}
	}

	return result
}

func determineStockDataFor(ticker string) (initialPrice, volatility float64) {
	ticker = strings.ToUpper(strings.TrimSpace(ticker))
	tickerSeed := int64(0)
	for _, r := range ticker {
		tickerSeed += int64(r)
	}
	tickerRand := rand.New(rand.NewSource(tickerSeed))

	// 1. determine price range based on ticker characteristics
	// - longer tickers tend to be newer/smaller companies -> lower prices
	// - tickers with more numbers tend to be newer -> lower prices
	// - special characters often indicate foreign or special stocks -> more variability

	// price between: 15-300
	basePrice := 15.0 + tickerRand.Float64()*285.0

	// length factor - longer names tend to be newer/smaller companies
	lengthFactor := 1.0
	if len(ticker) > 4 {
		lengthFactor = 0.7
	} else if len(ticker) < 3 {
		lengthFactor = 1.3
	}

	numberCount := 0
	specialCharCount := 0
	for _, r := range ticker {
		if unicode.IsDigit(r) {
			numberCount++
		} else if !unicode.IsLetter(r) {
			specialCharCount++
		}
	}

	// numbers in ticker often indicate newer companies (lower price)
	numberFactor := 1.0
	if numberCount > 0 {
		numberFactor = 0.6
	}

	// sSpecial characters often indicate ADRs or special stock classes
	specialFactor := 1.0
	if specialCharCount > 0 {
		specialFactor = 0.8 + tickerRand.Float64()*0.4 // More variability for these
	}

	adjustedPrice := basePrice * lengthFactor * numberFactor * specialFactor

	// 2. volatility based on ticker
	// - longer tickers - newer companies, more volatile
	// - tech-like prefixes - more volatile
	// - single letter tickers - usually established companies, less volatile

	baseVolatility := 0.01 + tickerRand.Float64()*0.02

	// length volatility adjustment
	if len(ticker) > 4 {
		baseVolatility *= 1.3
	} else if len(ticker) < 3 {
		baseVolatility *= 0.8
	}

	// tech company heuristic - common prefixes suggest tech/newer companies
	techPrefixes := []string{"A", "E", "I", "Q", "X", "Z", "NET", "BIT", "DATA", "TECH", "SOFT"}
	for _, prefix := range techPrefixes {
		if strings.HasPrefix(ticker, prefix) {
			baseVolatility *= 1.25
			break
		}
	}

	// industry heuristic based on common patterns
	financeSuffixes := []string{"AC", "AM", "AL", "BK", "FG", "FC", "INS"}
	for _, suffix := range financeSuffixes {
		if strings.HasSuffix(ticker, suffix) {
			baseVolatility *= 0.85 // financial companies often less volatile
			adjustedPrice *= 1.2   // financial companies often higher priced
			break
		}
	}

	// Round the price reasonably based on magnitude
	if adjustedPrice > 1000 {
		adjustedPrice = math.Round(adjustedPrice/10) * 10
	} else if adjustedPrice > 100 {
		adjustedPrice = math.Round(adjustedPrice)
	} else if adjustedPrice > 10 {
		adjustedPrice = math.Round(adjustedPrice*10) / 10
	} else {
		adjustedPrice = math.Round(adjustedPrice*100) / 100
	}

	return adjustedPrice, baseVolatility
}

func roundToTwoDecimals(val float64) float64 {
	return math.Round(val*100) / 100
}
