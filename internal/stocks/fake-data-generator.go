package stocks

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type FakeDataGenerator struct {
	InitialPrice float64
	Volatility   float64
	rand         *rand.Rand
}

func NewFakeDataGenerator(initialPrice, volatility float64, seed int64) *FakeDataGenerator {
	return &FakeDataGenerator{
		InitialPrice: initialPrice,
		Volatility:   volatility,
		rand:         rand.New(rand.NewSource(seed)),
	}
}

func (g *FakeDataGenerator) GenerateForTimeSpan(span string) []StockDataPoint {
	now := time.Now()

	switch span {
	case "d":
		return g.generateData(now.Add(-24*time.Hour), now, 5*time.Minute)
	case "w":
		return g.generateData(now.Add(-7*24*time.Hour), now, 1*time.Hour)
	case "m":
		return g.generateData(now.Add(-30*24*time.Hour), now, 4*time.Hour)
	case "y":
		return g.generateData(now.Add(-365*24*time.Hour), now, 12*time.Hour)
	case "y2d":
		startOfYear := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
		return g.generateData(startOfYear, now, 12*time.Hour)
	default:
		panic(fmt.Errorf("unknown timespan option: %s", span))
	}
}

func (g *FakeDataGenerator) generateData(start, end time.Time, interval time.Duration) []StockDataPoint {
	result := make([]StockDataPoint, 0)
	price := g.InitialPrice

	for t := start; t.Before(end); t = t.Add(interval) {
		if t.Weekday() == time.Saturday || t.Weekday() == time.Sunday {
			continue
		}

		hour, min := t.Hour(), t.Minute()
		isMarketOpen := hour > 9 || (hour == 9 && min >= 30)
		notMarketClosed := hour < 16

		if isMarketOpen && notMarketClosed {
			randomFactor := g.rand.NormFloat64() * g.Volatility

			price *= (1 + randomFactor)

			if price < 1.0 {
				price = 1.0 + g.rand.Float64()
			}

			result = append(result, StockDataPoint{
				Timestamp: t,
				Price:     roundToTwoDecimals(price),
				Currency:  "USD",
			})
		}
	}

	return result
}

func roundToTwoDecimals(val float64) float64 {
	return math.Round(val*100) / 100
}
