package indexer

import (
	"strings"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/models"
)

func (bc *blockContext) processTransactionLogs(tx chain.Transaction) (err error) {
	logs := tx.ExtractLogs()
	if len(logs) == 0 {
		return
	}
	logsResponse := models.TransactionLogsResponse{
		Data: logs,
	}
	bc.Batch.AddToBatchSingleKey(logsResponse, tx.Hash)
	err = bc.balances.UpdateEvents(logs)
	if err != nil {
		return err
	}
	err = bc.handleValidatorRegistrations(logs)
	if err != nil {
		return err
	}
	return
}

func (bc *blockContext) handleValidatorRegistrations(logs []models.EventLog) (err error) {
	const registerValidatorTopic = "0xd09501348473474a20c772c79c653e1fd7e8b437e418fe235d277d2c88853251"
	for _, log := range logs {
		if strings.Compare(log.Topics[0], registerValidatorTopic) != 0 {
			continue
		}
		addressStats := bc.addressStats[log.Address]
		if addressStats == nil {
			addressStats = &storage.AddressStats{
				Address: log.Address,
				StatsResponse: models.StatsResponse{
					ValidatorRegisteredBlock: &bc.block.Number,
				},
			}
		} else {
			addressStats.ValidatorRegisteredBlock = &bc.block.Number
		}
		bc.addressStats[log.Address] = addressStats
	}
	bc.addAddressStatsToBatch()
	return nil
}
