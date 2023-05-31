package indexer

import (
	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/models"
	log "github.com/sirupsen/logrus"
)

func (bc *blockContext) processTransactions(trxHashes *[]string) (err error) {
	var traces []chain.TransactionTrace
	bc.tp.Go(MakeTaskWithResult(bc.client.TraceBlockTransactions, bc.block.Number, &traces, &err).Run)
	err, transactions := bc.getTransactions(trxHashes)

	bc.tp.Wait()

	internal_transactions := new(models.InternalTransactionsResponse)
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
				bc.SaveTransaction(&internal)
			}
			bc.batch.AddToBatchSingleKey(internal_transactions, trx_model.Hash)
		}
	}

	return
}

func (bc *blockContext) getTransactions(trxHashes *[]string) (err error, trxs []chain.Transaction) {
	trxs = make([]chain.Transaction, len(*trxHashes))

	for i, trx_hash := range *trxHashes {
		bc.tp.Go(MakeTaskWithResult(bc.client.GetTransactionByHash, trx_hash, &trxs[i], &err).Run)
	}

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

func (bc *blockContext) addAddressStatsToBatch() {
	for _, stats := range bc.addressStats {
		bc.batch.AddToBatch(stats, stats.Address, 0)
	}
}
