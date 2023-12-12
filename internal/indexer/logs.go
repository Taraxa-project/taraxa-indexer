package indexer

import (
	"strings"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/ethereum/go-ethereum/common"
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
		if strings.Compare(log.Topics[0], registerValidatorTopic) != 0 && strings.Compare(log.Address, "0x00000000000000000000000000000000000000fe") != 0 {
			continue
		}
		address := common.HexToAddress(log.Topics[1])
		bc.addressStats.GetAddress(bc.Storage, address.Hex()).RegisterValidatorBlock(bc.block.Number)
	}
	return nil
}
