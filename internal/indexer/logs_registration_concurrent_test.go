package indexer

import (
	"sync"
	"testing"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/stretchr/testify/assert"
)

func TestHandleValidatorRegistrations(t *testing.T) {
	var wg sync.WaitGroup
	concurrency := 5
	wg.Add(concurrency)
	// Create a mock blockContext
	mc := chain.MakeMockClient()
	bc := MakeTestBlockContext(mc, 1)

	// Create a slice of EventLog for testing
	logs := []models.EventLog{
		{
			Address:  "0x00000000000000000000000000000000000000fe",
			Data:     "0x000000000000000000000000000000000000000000000000a7a44a964be1f30a",
			LogIndex: 1,
			Removed:  false,
			Topics: []string{
				"0xd09501348473474a20c772c79c653e1fd7e8b437e418fe235d277d2c88853251",
				"0x000000000000000000000000fc43217e71ec0a1cc480f3d210cd07cbde7374ec",
				"0x000000000000000000000000e50b5452b2e8435404dbe06e6a05410c47b7583d",
			},
			TransactionHash:  "0x689811a0705b89add2cd02d8a713bbd43c31c5afc123aeaca264494b375d6968",
			TransactionIndex: 105,
		},
		{
			Address:  "0x00000000000000000000000000000000000000fe",
			Data:     "0x000000000000000000000000000000000000000000000000a7a44a964be1f30a",
			LogIndex: 2,
			Removed:  false,
			Topics: []string{
				"0xd09501348473474a20c772c79c653e1fd7e8b437e418fe235d277d2c88853251",
				"0x000000000000000000000000e50b5452b2e8435404dbe06e6a05410c47b7583d",
				"0x000000000000000000000000fc43217e71ec0a1cc480f3d210cd07cbde7374ec",
			},
			TransactionHash:  "0x689811a0705b89add2cd02d8a713bbd43c31c5afc123aeaca264494b375d6968",
			TransactionIndex: 105,
		},
	}

	// Call the function in a goroutine to simulate concurrency
	writeStats := func() {
		defer wg.Done()
		err := bc.handleValidatorRegistrations(logs)
		assert.Nil(t, err)
	}

	for concurrency > 0 {
		go writeStats()
		concurrency--
	}

	wg.Wait()

	bc.commit()
}
