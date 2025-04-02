package repository_test

import (
	"testing"
	"time"

	"github.com/ruegerj/stock-sight/internal/queries"
	"github.com/stretchr/testify/assert"
)

func TestSqlcTransactionRepository_GetById(t *testing.T) {
	trx, err := trxRepo.GetById(t.Context(), sellFixture.ID)

	assert.NoError(t, err)
	assertEqualTransaction(t, sellFixture, trx, true)
}

func TestSqlcTransactionRepository_GetById_FailForNegativeId(t *testing.T) {
	var id int64 = -5

	_, err := trxRepo.GetById(t.Context(), id)

	assert.Error(t, err)
	assert.ErrorContains(t, err, "supply a valid transaction id")
}

func TestSqlcTransactionRepository_GetAll(t *testing.T) {
	trxs, err := trxRepo.GetAll(t.Context())

	assert.NoError(t, err)
	// in order to prevent any flackyness -> only check for guaranteed records (pre-defined fixtures)
	assert.GreaterOrEqual(t, 2, len(trxs))
	assertContainsTransaction(t, trxs, buyFixture)
	assertContainsTransaction(t, trxs, sellFixture)
}

func TestSqlcTransactionRepository_Create(t *testing.T) {
	expected := queries.Transaction{
		Ticker:       "NVDA",
		Amount:       10,
		Currency:     "USD",
		PricePerUnit: 110.42,
		Date:         time.Date(2025, time.January, 1, 12, 0, 0, 0, time.UTC),
		IsBuy:        true,
	}

	actual, err := trxRepo.Create(t.Context(),
		expected.Ticker,
		expected.Amount,
		expected.Currency,
		expected.PricePerUnit,
		expected.Date,
		expected.IsBuy)

	assert.NoError(t, err)
	assertEqualTransaction(t, expected, actual, false)
}

func TestSqlcTransactionRepository_Update(t *testing.T) {
	trxToUpdate, err := trxRepo.Create(t.Context(), "PLTR", 1, "USD", 87.45, time.Now(), true)
	assert.NoError(t, err)
	expected := trxToUpdate
	expected.Amount = 5

	err = trxRepo.Update(t.Context(), expected)

	assert.NoError(t, err)
	actual, _ := trxRepo.GetById(t.Context(), trxToUpdate.ID)
	assertEqualTransaction(t, expected, actual, true)
}

func TestSqlcTransactionRepository_Delete(t *testing.T) {
	trxToDelete, err := trxRepo.Create(t.Context(), "PLTR", 1, "USD", 87.45, time.Now(), true)
	assert.NoError(t, err)

	err = trxRepo.Delete(t.Context(), trxToDelete.ID)

	assert.NoError(t, err)
	actual, _ := trxRepo.GetById(t.Context(), trxToDelete.ID)
	assert.Empty(t, actual)
}

func TestSqlcTransactionRepository_Delete_FailForNegativeId(t *testing.T) {
	var id int64 = -5

	err := trxRepo.Delete(t.Context(), id)

	assert.Error(t, err)
	assert.ErrorContains(t, err, "supply a valid transaction id")
}
