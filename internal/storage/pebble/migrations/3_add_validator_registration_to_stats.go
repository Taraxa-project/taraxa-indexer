package migration

import (
	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage/pebble"
	"github.com/Taraxa-project/taraxa-indexer/models"
	log "github.com/sirupsen/logrus"
)

type AddValidatorRegistrationBlock struct {
	id            string
	blockchain_ws string
}

func (m *AddValidatorRegistrationBlock) GetId() string {
	return m.id
}

func getCurrentBlockNumber(client *chain.WsClient) uint64 {
	// Ethereum node URL
	blockNumber, err := client.GetLatestPeriod()
	if err != nil {
		log.Fatal(err)
	}
	return blockNumber
}

func (m *AddValidatorRegistrationBlock) Apply(s *pebble.Storage) error {
	// Ethereum node URL
	// Create an Ethereum client
	client, err := chain.NewWsClient(m.blockchain_ws)

	if err != nil {
		log.Fatal(err)
	}

	startBlock := uint64(0)
	endBlock := uint64(999)
	current := getCurrentBlockNumber(client)

	for {
		validators, err := GetValidatorsRegisteredInBlock(client, startBlock, endBlock)
		if err != nil {
			log.Fatal(err)
		}

		for _, validator := range validators {
			addressStats := s.GetAddressStats(validator.Validator.Hex())
			if addressStats == nil {
				addressStats = &storage.AddressStats{
					Address: validator.Validator.Hex(),
					StatsResponse: models.StatsResponse{
						ValidatorRegisteredBlock: &validator.BlockHeight,
					},
				}
			} else {
				addressStats.ValidatorRegisteredBlock = &validator.BlockHeight
			}
		}

		// Process the validators registered in this batch

		// Update the start and end block numbers for the next iteration
		startBlock = endBlock + 1
		endBlock += 1000

		// You can add a condition to break the loop when you reach a specific block number or any other criteria.
		// For example, to stop after reaching block 5000, you can add:
		if startBlock > current {
			break
		}
	}
	return nil
}
