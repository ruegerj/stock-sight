package repository_test

import (
	"context"
	"os"
	"slices"
	"testing"
	"time"

	"github.com/ruegerj/stock-sight/internal/db"
	"github.com/ruegerj/stock-sight/internal/queries"
	"github.com/ruegerj/stock-sight/internal/repository"
	"github.com/stretchr/testify/assert"
)

var stockRepo repository.StockRepository
var trxRepo repository.TransactionRepository

var buyFixture queries.Transaction
var sellFixture queries.Transaction

// Initializes the test environment for the repository package.
func TestMain(m *testing.M) {
	ctx := context.Background()
	dbConn := db.NewInMemorySQLite(ctx)
	stockRepo = repository.NewSqlcStockRepository(dbConn)
	trxRepo = repository.NewSqlcTransactionRepository(dbConn)

	err := createTransactionFixtures(ctx)
	if err != nil {
		panic(err)
	}

	exitCode := m.Run()
	os.Exit(exitCode)
}

func assertContainsTransaction(t *testing.T, trxs []queries.Transaction, expected queries.Transaction) {
	idx := slices.IndexFunc(trxs, func(trx queries.Transaction) bool {
		return trx.ID == expected.ID
	})

	assert.GreaterOrEqual(t, idx, 0)
	actual := trxs[idx]
	assertEqualTransaction(t, expected, actual, true)
}

func assertEqualTransaction(t *testing.T, expected, actual queries.Transaction, checkId bool) {
	if checkId {
		assert.Equal(t, expected.ID, actual.ID)
	}
	assert.Equal(t, expected.Ticker, actual.Ticker)
	assert.Equal(t, expected.Amount, actual.Amount)
	assert.Equal(t, expected.Currency, actual.Currency)
	assert.Equal(t, expected.PricePerUnit, actual.PricePerUnit)
	assert.Equal(t, expected.Date, actual.Date)
	assert.Equal(t, expected.IsBuy, actual.IsBuy)
}

func createTransactionFixtures(ctx context.Context) error {
	buyDate := time.Date(2020, time.January, 1, 16, 20, 0, 0, time.UTC)
	aaplBuy, err := trxRepo.Create(ctx, "AAPL", 10, "USD", 150.5, buyDate, true)
	if err != nil {
		return err
	}
	buyFixture = aaplBuy

	sellDate := time.Date(2024, time.November, 5, 0, 0, 0, 0, time.UTC)
	teslaSell, err := trxRepo.Create(ctx, "TSLA", 1000, "USD", 0.5, sellDate, false)
	if err != nil {
		return err
	}
	sellFixture = teslaSell

	return nil
}
