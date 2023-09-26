package indexer

import (
	"github.com/Taraxa-project/taraxa-indexer/internal/events"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/models"
)

func (bc *blockContext) processValidatorRegistrations(block *models.Pbft) (err error) {

	validators, err := events.GetValidatorsRegisteredInBlock("https://rpc.mainnet.taraxa.io", block.Number, block.Number)

	if err != nil {
		return err
	}
	for _, validator := range validators {
		addressStats := bc.addressStats[validator.Validator.Hex()]
		if addressStats == nil {
			addressStats = &storage.AddressStats{
				Address: validator.Validator.Hex(),
				StatsResponse: models.StatsResponse{
					ValidatorRegisteredBlock: &block.Number,
				},
			}
		} else {
			addressStats.ValidatorRegisteredBlock = &block.Number
		}
		bc.addressStats[validator.Validator.Hex()] = addressStats
	}
	bc.addAddressStatsToBatch()
	return
}
