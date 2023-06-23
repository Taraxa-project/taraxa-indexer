package indexer

import (
	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/models"
	log "github.com/sirupsen/logrus"
)

func (bc *blockContext) processTransactions(trxHashes *[]string) (err error) {
	var traces []chain.TransactionTrace
	var transactions []chain.Transaction

	tp := makeThreadPool()
	tp.Go(MakeTaskWithResult(bc.client.TraceBlockTransactions, bc.block.Number, &traces, &err).Run)
	tp.Go(MakeTaskWithResult(bc.getTransactions, trxHashes, &transactions, &err).Run)
	tp.Wait()

	internal_transactions := new(models.InternalTransactionsResponse)
	logs := new(models.TransactionLogsResponse)
	for i, trx := range transactions {
		trx_model := trx.ToModelWithTimestamp(bc.block.Timestamp)
		bc.SaveTransaction(trx_model)
		trace := traces[i]
		if len(trace.Trace) > 1 {
			for i, entry := range trace.Trace {
				if i == 0 {
					continue
				}
				internal := makeInternal(trx_model, entry)
				internal_transactions.Data = append(internal_transactions.Data, internal)
				eventLog := makeEventLog(bc, trx, entry)
				logs.Data = eventLog
				bc.SaveTransaction(&internal)
			}
			bc.batch.AddToBatchSingleKey(internal_transactions, trx_model.Hash)
			bc.batch.AddToBatchSingleKey(logs, trx_model.Hash)
		}
	}
	return
}

func (bc *blockContext) getTransactions(trxHashes *[]string) (trxs []chain.Transaction, err error) {
	trxs, err = bc.client.GetPeriodTransactions(bc.block.Number)
	if err != nil {
		log.WithError(err).Debug("GetPeriodTransactions error")
		return bc.getTransactionsOld(trxHashes)
	}

	return
}

func (bc *blockContext) getTransactionsOld(trxHashes *[]string) (trxs []chain.Transaction, err error) {
	trxs = make([]chain.Transaction, len(*trxHashes))

	tp := makeThreadPool()
	for i, trx_hash := range *trxHashes {
		tp.Go(MakeTaskWithResult(bc.client.GetTransactionByHash, trx_hash, &trxs[i], &err).Run)
	}
	tp.Wait()
	return
}

func makeInternal(trx *models.Transaction, entry chain.TraceEntry) (internal models.Transaction) {
	internal = *trx
	internal.From = entry.Action.From
	internal.To = entry.Action.To
	internal.Value = entry.Action.Value
	internal.GasUsed = chain.ParseInt(entry.Result.GasUsed)
	internal.Type = chain.GetTransactionType(trx.To, entry.Action.Input, true)
	internal.BlockNumber = 0

	return internal
}

func (bc *blockContext) SaveTransaction(trx *models.Transaction) {
	log.WithFields(log.Fields{"from": trx.From, "to": trx.To, "hash": trx.Hash}).Trace("Saving transaction")

	from_index := bc.getAddress(bc.storage, trx.From).AddTransaction(trx.Timestamp)
	to_index := bc.getAddress(bc.storage, trx.To).AddTransaction(trx.Timestamp)

	bc.batch.AddToBatch(trx, trx.From, from_index)
	bc.batch.AddToBatch(trx, trx.To, to_index)
}

func (bc *blockContext) SaveEventLog(eventLog *models.EventLog) {
	log.WithFields(log.Fields{"address": eventLog.Address, "trnxHash": eventLog.TransactionHash}).Trace("Saving Event Log")

	bc.batch.AddToBatch(eventLog, eventLog.TransactionHash, eventLog.LogIndex)
}

func (bc *blockContext) addAddressStatsToBatch() {
	for _, stats := range bc.addressStats {
		bc.batch.AddToBatch(stats, stats.Address, 0)
	}
}

func makeEventLog(bc *blockContext, trx chain.Transaction, entry chain.TraceEntry) (logs []models.EventLog) {
	for _, log := range trx.Logs {
		log := models.EventLog{
			LogIndex: log.LogIndex,
			Address: log.Address,
			Data: log.Data,
			Removed: log.Removed,
			Topics: log.Topics,
			TransactionHash: log.TransactionHash,
			TransactionIndex: log.TransactionIndex,
		}
		bc.SaveEventLog(&log)
		logs = append(logs, log)
	}
	return logs
}
