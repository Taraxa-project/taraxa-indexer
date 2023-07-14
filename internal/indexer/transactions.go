package indexer

import (
	"math/big"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/models"
	log "github.com/sirupsen/logrus"
)

func (bc *blockContext) processTransactions(trxHashes []string) (err error) {
	if len(trxHashes) == 0 {
		return
	}
	var traces []chain.TransactionTrace
	var transactions []chain.Transaction

	tp := common.MakeThreadPool()
	tp.Go(common.MakeTaskWithResult(bc.Client.TraceBlockTransactions, bc.block.Number, &traces, &err).Run)
	tp.Go(common.MakeTaskWithResult(bc.getTransactions, trxHashes, &transactions, &err).Run)
	tp.Wait()

	if err != nil || len(traces) != len(transactions) || len(trxHashes) != len(transactions) {
		return
	}

	block_fee := big.NewInt(0)

	bc.transactions = make([]models.Transaction, len(transactions))
	for t_idx := 0; t_idx < len(transactions); t_idx++ {
		bc.transactions[t_idx] = transactions[t_idx].ToModelWithTimestamp(bc.block.Timestamp)
		bc.SaveTransaction(bc.transactions[t_idx])

		trx_fee := transactions[t_idx].GetFee()
		block_fee.Add(block_fee, trx_fee)
		// Remove fee from sender balance
		bc.balances.AddToBalance(transactions[t_idx].From, big.NewInt(0).Neg(trx_fee))
		if transactions[t_idx].Status == "0x0" {
			continue
		}
		// remove value from sender and add to receiver
		receiver := transactions[t_idx].To
		if receiver == "" {
			receiver = transactions[t_idx].ContractAddress
		}
		bc.balances.UpdateBalances(transactions[t_idx].From, receiver, transactions[t_idx].Value)

		// process logs
		logs := models.TransactionLogsResponse{
			Data: transactions[t_idx].ExtractLogs(),
		}
		bc.Batch.AddToBatchSingleKey(logs, bc.transactions[t_idx].Hash)
		err := bc.balances.UpdateEvents(logs.Data)
		if err != nil {
			return err
		}

		if internal_transactions := bc.processInternalTransactions(traces[t_idx], t_idx); internal_transactions != nil {
			bc.Batch.AddToBatchSingleKey(internal_transactions, bc.transactions[t_idx].Hash)
		}
	}
	// add total fee from the block to block producer balance
	bc.balances.AddToBalance(bc.block.Author, block_fee)
	return
}

func (bc *blockContext) processInternalTransactions(trace chain.TransactionTrace, t_idx int) (internal_transactions *models.InternalTransactionsResponse) {
	if len(trace.Trace) <= 1 {
		return
	}
	internal_transactions = new(models.InternalTransactionsResponse)
	internal_transactions.Data = make([]models.Transaction, 0, len(trace.Trace)-1)

	for e_idx, entry := range trace.Trace {
		if e_idx == 0 {
			continue
		}
		internal := makeInternal(bc.transactions[t_idx], entry)
		internal_transactions.Data = append(internal_transactions.Data, internal)
		bc.SaveTransaction(internal)
		bc.balances.UpdateBalances(internal.From, internal.To, internal.Value)
	}
	return
}

func (bc *blockContext) getTransactions(trxHashes []string) (trxs []chain.Transaction, err error) {
	trxs, err = bc.Client.GetPeriodTransactions(bc.block.Number)
	if err != nil {
		log.WithError(err).Debug("GetPeriodTransactions error")
		return bc.getTransactionsOld(trxHashes)
	}

	return
}

func (bc *blockContext) getTransactionsOld(trxHashes []string) (trxs []chain.Transaction, err error) {
	trxs = make([]chain.Transaction, len(trxHashes))

	tp := common.MakeThreadPool()
	for i, trx_hash := range trxHashes {
		tp.Go(common.MakeTaskWithResult(bc.Client.GetTransactionByHash, trx_hash, &trxs[i], &err).Run)
	}
	tp.Wait()
	return
}

func makeInternal(trx models.Transaction, entry chain.TraceEntry) (internal models.Transaction) {
	internal = trx
	internal.From = entry.Action.From
	internal.To = entry.Action.To
	internal.Value = entry.Action.Value
	internal.GasUsed = chain.ParseUInt(entry.Result.GasUsed)
	internal.Type = chain.GetTransactionType(trx.To, entry.Action.Input, true)
	internal.BlockNumber = trx.BlockNumber
	return
}

func (bc *blockContext) SaveTransaction(trx models.Transaction) {
	log.WithFields(log.Fields{"from": trx.From, "to": trx.To, "hash": trx.Hash}).Trace("Saving transaction")

	from_index := bc.getAddress(bc.Storage, trx.From).AddTransaction(trx.Timestamp)
	to_index := bc.getAddress(bc.Storage, trx.To).AddTransaction(trx.Timestamp)

	bc.Batch.AddToBatch(trx, trx.From, from_index)
	bc.Batch.AddToBatch(trx, trx.To, to_index)
}

func (bc *blockContext) addAddressStatsToBatch() {
	for _, stats := range bc.addressStats {
		bc.Batch.AddToBatch(stats, stats.Address, 0)
	}
}
