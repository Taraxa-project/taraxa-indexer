package indexer

import (
	"fmt"
	"strconv"
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
	err = bc.handleValidatorCommissionChange(logs)
	if err != nil {
		return err
	}
	return
}

func (bc *blockContext) handleValidatorRegistrations(logs []models.EventLog) (err error) {
	const registerValidatorTopic = "0xd09501348473474a20c772c79c653e1fd7e8b437e418fe235d277d2c88853251"
	for _, log := range logs {
		if strings.Compare(log.Topics[0], registerValidatorTopic) != 0 && strings.Compare(log.Address, "0x00000000000000000000000000000000000000fe") != 0 {
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
	return nil
}

func (bc *blockContext) handleValidatorCommissionChange(logs []models.EventLog) (err error) {
	const commissionSetTopic = "0xc909daf778d180f43dac53b55d0de934d2f1e0b70412ca274982e4e6e894eb1a"
	for _, log := range logs {
		if strings.Compare(log.Topics[0], commissionSetTopic) != 0 && strings.Compare(log.Address, "0x00000000000000000000000000000000000000fe") != 0 {
			continue
		}
		commissionHex := log.Data
		hexString := commissionHex[2:]
		value, err := strconv.ParseUint(hexString, 16, 64)
		if err != nil {
			value = 0
			fmt.Println("Error parsing hexadecimal string:", err)
		}
		addressStats := bc.addressStats[log.Address]
		if addressStats == nil {
			addressStats = &storage.AddressStats{
				Address: log.Address,
				StatsResponse: models.StatsResponse{
					Commission: &value,
				},
			}
		} else {
			addressStats.Commission = &value
		}
		bc.addressStats[log.Address] = addressStats
	}
	return nil
}
