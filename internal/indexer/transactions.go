package indexer

import (
	"math/big"
	"runtime/debug"

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

	tp := common.MakeThreadPool()
	tp.Go(common.MakeTaskWithResult(bc.Client.TraceBlockTransactions, bc.block.Number, &traces, &err).Run)
	tp.Go(common.MakeTaskWithResult(bc.getTransactions, trxHashes, &bc.transactions, &err).Run)
	tp.Wait()

	if err != nil || len(traces) != len(bc.transactions) || len(trxHashes) != len(bc.transactions) {
		return
	}

	for t_idx := 0; t_idx < len(bc.transactions); t_idx++ {
		bc.transactions[t_idx].SetTimestamp(bc.block.Timestamp)

		bc.SaveTransaction(*bc.transactions[t_idx].GetModel())

		trx_fee := bc.transactions[t_idx].GetFee()
		bc.blockFee.Add(bc.blockFee, trx_fee)
		// Remove fee from sender balance
		bc.accounts.AddToBalance(bc.transactions[t_idx].From, big.NewInt(0).Neg(trx_fee))
		if !bc.transactions[t_idx].Status {
			continue
		}
		// remove value from sender and add to receiver
		receiver := bc.transactions[t_idx].To
		if receiver == "" {
			receiver = bc.transactions[t_idx].ContractAddress
		}
		bc.accounts.UpdateBalances(bc.transactions[t_idx].From, receiver, bc.transactions[t_idx].Value)

		// process logs
		err = bc.processTransactionLogs(bc.transactions[t_idx])
		if err != nil {
			return
		}

		if internal_transactions := bc.processInternalTransactions(traces[t_idx], t_idx, bc.transactions[t_idx].GasPrice); internal_transactions != nil {
			bc.Batch.AddToBatchSingleKey(internal_transactions, bc.transactions[t_idx].Hash)
		}
	}
	// add total fee from the block to block producer balance
	if bc.Config.Chain != nil && (bc.block.Number < bc.Config.Chain.Hardforks.MagnoliaHf.BlockNum) {
		bc.accounts.AddToBalance(bc.block.Author, bc.blockFee)
	}
	return
}

func (bc *blockContext) processInternalTransactions(trace chain.TransactionTrace, t_idx int, gasPrice uint64) (internal_transactions *models.InternalTransactionsResponse) {
	if len(trace.Trace) <= 1 {
		return
	}
	internal_transactions = new(models.InternalTransactionsResponse)
	internal_transactions.Data = make([]models.Transaction, 0, len(trace.Trace)-1)

	for e_idx, entry := range trace.Trace {
		if e_idx == 0 {
			continue
		}
		internal := makeInternal(*bc.transactions[t_idx].GetModel(), entry, gasPrice)
		internal_transactions.Data = append(internal_transactions.Data, internal)

		bc.SaveTransaction(internal)
		// TODO: hotfix, remove after fix in taraxa-node
		if entry.Action.CallType != "delegatecall" {
			bc.accounts.UpdateBalances(internal.From, internal.To, internal.Value)
		}
	}
	return
}

func (bc *blockContext) getTransactions(trxHashes []string) (trxs []chain.Transaction, err error) {
	trxs, err = bc.Client.GetPeriodTransactions(bc.block.Number)
	if err != nil {
		debug.PrintStack()
		log.WithError(err).Fatal("GetPeriodTransactions error")
	}

	return
}

func makeInternal(trx models.Transaction, entry chain.TraceEntry, gasCost uint64) (internal models.Transaction) {
	internal = trx
	internal.From = entry.Action.From
	internal.To = chain.GetInternalTransactionTarget(entry)
	internal.Value = entry.Action.Value
	internal.GasCost = common.ParseUInt(entry.Result.GasUsed) * gasCost
	internal.Type = chain.GetTransactionType(trx.To, entry.Action.Input, entry.Type, true)
	internal.BlockNumber = trx.BlockNumber
	return
}

func (bc *blockContext) SaveTransaction(trx models.Transaction) {
	log.WithFields(log.Fields{"from": trx.From, "to": trx.To, "hash": trx.Hash}).Trace("Saving transaction")

	from_index := bc.addressStats.GetAddress(bc.Storage, trx.From).AddTransaction(trx.Timestamp)
	bc.Batch.AddToBatch(trx, trx.From, from_index)
	if trx.To != "" {
		to_index := bc.addressStats.GetAddress(bc.Storage, trx.To).AddTransaction(trx.Timestamp)
		bc.Batch.AddToBatch(trx, trx.To, to_index)
	}

	if (trx.Input != "0x") && (trx.Input != "") {
		bc.Batch.AddToBatchSingleKey(trx, trx.Hash)
	}
}
