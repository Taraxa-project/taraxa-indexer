package indexer

import (
	"fmt"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/models"
	log "github.com/sirupsen/logrus"
)

func (bc *blockContext) processTransactions(trxHashes *[]string) (err error) {
	var traces []chain.TransactionTrace
	var transactions []chain.Transaction

	tp := common.MakeThreadPool()
	tp.Go(common.MakeTaskWithResult(bc.Client.TraceBlockTransactions, bc.block.Number, &traces, &err).Run)
	tp.Go(common.MakeTaskWithResult(bc.getTransactions, trxHashes, &transactions, &err).Run)
	tp.Wait()

	if err != nil {
		return
	}

	internal_transactions := new(models.InternalTransactionsResponse)
	bc.transactions = make([]*models.Transaction, len(transactions))

	for t_idx, trx := range transactions {
		fmt.Println("processTransactions", t_idx, trx.Hash)
		bc.transactions[t_idx] = trx.ToModelWithTimestamp(bc.block.Timestamp)
		bc.SaveTransaction(bc.transactions[t_idx])
		trace := traces[t_idx]
		if len(trace.Trace) <= 1 {
			continue
		}
		for e_idx, entry := range trace.Trace {
			if e_idx == 0 {
				continue
			}
			internal := makeInternal(bc.transactions[t_idx], entry)
			internal_transactions.Data = append(internal_transactions.Data, *internal)
		}
		logs := models.TransactionLogsResponse{
			Data: trx.ExtractLogs(),
		}
		bc.batch.AddToBatchSingleKey(logs, trx_model.Hash)
	}
	bc.Batch.AddToBatchSingleKey(internal_transactions, "internal_transactions")
	return
}

func (bc *blockContext) getTransactions(trxHashes *[]string) (trxs []chain.Transaction, err error) {
	trxs, err = bc.Client.GetPeriodTransactions(bc.block.Number)
	// fmt.Println("[")
	// for _, trx := range trxs {
	// 	fmt.Println("\"", trx.Hash, "\",")
	// }
	// fmt.Println("]")
	if err != nil {
		log.WithError(err).Debug("GetPeriodTransactions error")
		return bc.getTransactionsOld(trxHashes)
	}
	fmt.Println("getTransactions")

	return
}

func (bc *blockContext) getTransactionsOld(trxHashes *[]string) (trxs []chain.Transaction, err error) {
	fmt.Println("getTransactionsOld")
	trxs = make([]chain.Transaction, len(*trxHashes))

	tp := common.MakeThreadPool()
	for i, trx_hash := range *trxHashes {
		tp.Go(common.MakeTaskWithResult(bc.Client.GetTransactionByHash, trx_hash, &trxs[i], &err).Run)
	}
	tp.Wait()
	return
}

func makeInternal(trx *models.Transaction, entry chain.TraceEntry) (internal *models.Transaction) {
	internal = new(models.Transaction)
	*internal = *trx
	internal.From = entry.Action.From
	internal.To = entry.Action.To
	internal.Value = entry.Action.Value
	internal.GasUsed = chain.ParseUInt(entry.Result.GasUsed)
	internal.Type = chain.GetTransactionType(trx.To, entry.Action.Input, true)
	internal.BlockNumber = 0

	return internal
}

func (bc *blockContext) SaveTransaction(trx *models.Transaction) {
	log.WithFields(log.Fields{"from": trx.From, "to": trx.To, "hash": trx.Hash}).Trace("Saving transaction")

	from_index := bc.getAddress(bc.Storage, trx.From).AddTransaction(trx.Timestamp)
	to_index := bc.getAddress(bc.Storage, trx.To).AddTransaction(trx.Timestamp)

	bc.Batch.AddToBatch(trx, trx.From, from_index)
	bc.Batch.AddToBatch(trx, trx.To, to_index)
}

// func (bc *blockContext) SaveEventLog(eventLog *models.EventLog) {
// 	log.WithFields(log.Fields{"address": eventLog.Address, "trnxHash": eventLog.TransactionHash}).Trace("Saving Event Log")

// 	bc.batch.AddToBatch(eventLog, eventLog.TransactionHash)
// }

func (bc *blockContext) addAddressStatsToBatch() {
	for _, stats := range bc.addressStats {
		bc.Batch.AddToBatch(stats, stats.Address, 0)
	}
}
