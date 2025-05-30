package api

import (
	"testing"

	"github.com/Taraxa-project/taraxa-indexer/internal/storage/pebble"
	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/go-playground/assert/v2"
)

func TestWasAccountActive(t *testing.T) {
	address := "0x123"
	db := pebble.NewStorage(t.TempDir())
	defer db.Close()

	batch := db.NewBatch()
	batch.Add(&models.Transaction{
		From:      address,
		Timestamp: 200,
	}, address, 1)
	batch.CommitBatch()

	assert.Equal(t, wasAccountActive(db, address, 100, 200), true)
	assert.Equal(t, wasAccountActive(db, address, 100, 1000), true)
	assert.Equal(t, wasAccountActive(db, address, 1000, 2000), false)
	assert.Equal(t, wasAccountActive(db, address, 1000, 10000), false)
}

func TestReceivedTransactionsCount(t *testing.T) {
	address := "0x123"
	db := pebble.NewStorage(t.TempDir())
	defer db.Close()

	batch := db.NewBatch()
	i := uint64(1)
	batch.Add(models.Transaction{
		To:        address,
		Timestamp: 200,
	}, address, i)
	i++
	batch.Add(models.Transaction{
		To:        address,
		Timestamp: 100,
	}, address, i)
	i++

	// shouldn't be counted
	batch.Add(models.Transaction{
		From:      address,
		Timestamp: 200,
	}, address, i)
	i++
	batch.Add(models.Transaction{
		To:        address,
		Timestamp: 201,
	}, address, i)
	i++

	batch.Add(models.Transaction{
		To:        address,
		Timestamp: 99,
	}, address, i)

	batch.CommitBatch()

	assert.Equal(t, receivedTransactionsCount(db, address, 100, 200), uint64(2))
	assert.Equal(t, receivedTransactionsCount(db, address, 90, 100), uint64(2))
	assert.Equal(t, receivedTransactionsCount(db, address, 200, 202), uint64(1))
}
