package indexer

import (
	"math/big"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/models"
	log "github.com/sirupsen/logrus"
)

func (bc *blockContext) processTransactions() (err error) {
	if len(bc.Block.Pbft.Transactions) == 0 {
		return
	}

	if err != nil || len(bc.Block.Pbft.Transactions) != len(bc.Block.Transactions) || len(bc.Block.Traces) != len(bc.Block.Transactions) {
		return
	}

	for t_idx := 0; t_idx < len(bc.Block.Transactions); t_idx++ {
		bc.Block.Transactions[t_idx].SetTimestamp(bc.Block.Pbft.Timestamp)

		bc.SaveTransaction(*bc.Block.Transactions[t_idx].GetModel())

		trx_fee := bc.Block.Transactions[t_idx].GetFee()
		// Remove fee from sender balance
		bc.accounts.AddToBalance(bc.Block.Transactions[t_idx].From, big.NewInt(0).Neg(trx_fee))
		if !bc.Block.Transactions[t_idx].Status {
			continue
		}
		// remove value from sender and add to receiver
		receiver := bc.Block.Transactions[t_idx].To
		if receiver == "" {
			receiver = bc.Block.Transactions[t_idx].ContractAddress
		}
		bc.accounts.UpdateBalances(bc.Block.Transactions[t_idx].From, receiver, bc.Block.Transactions[t_idx].Value)

		// process logs
		err = bc.processTransactionLogs(bc.Block.Transactions[t_idx])
		if err != nil {
			return
		}

		if internal_transactions := bc.processInternalTransactions(bc.Block.Traces[t_idx], t_idx, bc.Block.Transactions[t_idx].GasPrice); internal_transactions != nil {
			bc.Batch.AddToBatchSingleKey(internal_transactions, bc.Block.Transactions[t_idx].Hash)
		}
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
		internal := makeInternal(*bc.Block.Transactions[t_idx].GetModel(), entry, gasPrice)
		internal_transactions.Data = append(internal_transactions.Data, internal)

		bc.SaveTransaction(internal)
		// TODO: hotfix, remove after fix in taraxa-node
		if entry.Action.CallType != "delegatecall" {
			bc.accounts.UpdateBalances(internal.From, internal.To, internal.Value)
		}
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
