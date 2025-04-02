package repository

import (
	"context"
	"os"
	"slices"
	"testing"
	"time"

	"github.com/ruegerj/stock-sight/internal/db"
	"github.com/ruegerj/stock-sight/internal/queries"
	"github.com/stretchr/testify/assert"
)

var dbConn db.DbConnection
var repo TransactionRepository
var buyFixture queries.Transaction
var sellFixture queries.Transaction

func TestMain(m *testing.M) {
	ctx := context.Background()
	dbConn = db.NewInMemorySQLite(ctx)
	repo = NewSqlcTransactionRepository(dbConn)

	err := createTransactionFixtures(ctx)
	if err != nil {
		panic(err)
	}

	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestGetById(t *testing.T) {
	trx, err := repo.GetById(t.Context(), sellFixture.ID)

	assert.NoError(t, err)
	assertEqualTransaction(t, sellFixture, trx, true)
}

func TestGetById_FailForNegativeId(t *testing.T) {
	var id int64 = -5

	_, err := repo.GetById(t.Context(), id)

	assert.Error(t, err)
	assert.ErrorContains(t, err, "supply a valid transaction id")
}

func TestGetAll(t *testing.T) {
	trxs, err := repo.GetAll(t.Context())

	assert.NoError(t, err)
	// in order to prevent any flackyness -> only check for guaranteed records (pre-defined fixtures)
	assert.GreaterOrEqual(t, 2, len(trxs))
	assertContainsTransaction(t, trxs, buyFixture)
	assertContainsTransaction(t, trxs, sellFixture)
}

func TestCreate(t *testing.T) {
	expected := queries.Transaction{
		Ticker:       "NVDA",
		Amount:       10,
		Currency:     "USD",
		PricePerUnit: 110.42,
		Date:         time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC),
		IsBuy:        true,
	}

	actual, err := repo.Create(t.Context(),
		expected.Ticker,
		expected.Amount,
		expected.Currency,
		expected.PricePerUnit,
		expected.Date,
		expected.IsBuy)

	assert.NoError(t, err)
	assertEqualTransaction(t, expected, actual, false)
}

func TestUpdate(t *testing.T) {
	trxToUpdate, err := repo.Create(t.Context(), "PLTR", 1, "USD", 87.45, time.Now(), true)
	assert.NoError(t, err)
	expected := trxToUpdate
	expected.Amount = 5

	err = repo.Update(t.Context(), expected)

	assert.NoError(t, err)
	actual, _ := repo.GetById(t.Context(), trxToUpdate.ID)
	assertEqualTransaction(t, expected, actual, true)
}

func TestDelete(t *testing.T) {
	trxToDelete, err := repo.Create(t.Context(), "PLTR", 1, "USD", 87.45, time.Now(), true)
	assert.NoError(t, err)

	err = repo.Delete(t.Context(), trxToDelete.ID)

	assert.NoError(t, err)
	actual, _ := repo.GetById(t.Context(), trxToDelete.ID)
	assert.Empty(t, actual)
}

func TestDelete_FailForNegativeId(t *testing.T) {
	var id int64 = -5

	err := repo.Delete(t.Context(), id)

	assert.Error(t, err)
	assert.ErrorContains(t, err, "supply a valid transaction id")
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
	buyDate := time.Date(2020, 1, 1, 16, 20, 0, 0, time.UTC)
	aaplBuy, err := repo.Create(ctx, "AAPL", 10, "USD", 150.5, buyDate, true)
	if err != nil {
		return err
	}
	buyFixture = aaplBuy

	sellDate := time.Date(2024, 11, 5, 0, 0, 0, 0, time.UTC)
	teslaSell, err := repo.Create(ctx, "TSLA", 1000, "USD", 0.5, sellDate, false)
	if err != nil {
		return err
	}
	sellFixture = teslaSell

	return nil
}
