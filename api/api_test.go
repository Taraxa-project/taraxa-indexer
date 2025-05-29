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

	api := NewApiHandler(db, nil, nil)

	assert.Equal(t, api.wasAccountActive(address, 100, 200), true)
	assert.Equal(t, api.wasAccountActive(address, 100, 1000), true)
	assert.Equal(t, api.wasAccountActive(address, 1000, 2000), false)
	assert.Equal(t, api.wasAccountActive(address, 1000, 10000), false)
}
