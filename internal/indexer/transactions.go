package indexer

import (
	"math/big"
	"time"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/ethereum/go-ethereum/rlp"
	log "github.com/sirupsen/logrus"
)

func (bc *blockContext) processTransactions() (err error) {
	if len(bc.Block.Pbft.Transactions) == 0 {
		return
	}

	start := time.Now()
	if len(bc.Block.Pbft.Transactions) != len(bc.Block.Transactions) {
		log.WithFields(log.Fields{"in_block": len(bc.Block.Pbft.Transactions), "transactions": len(bc.Block.Transactions), "traces": len(bc.Block.Traces)}).Error("Transactions count mismatch")
	}
	feeReward := big.NewInt(0)

	start_tp := time.Now()
	for t_idx := 0; t_idx < len(bc.Block.Transactions); t_idx++ {
		trx_fee := bc.Block.Transactions[t_idx].GetFee()
		feeReward.Add(feeReward, trx_fee)
		// Remove fee from sender balance
		bc.accounts.AddToBalance(bc.Block.Transactions[t_idx].From, big.NewInt(0).Neg(trx_fee))
		err = bc.processTransaction(t_idx)
		if err != nil {
			return
		}
	}
	elapsed_tp := time.Since(start_tp)
	log.WithFields(log.Fields{"func": "scheduleTransactions", "period": bc.Block.Pbft.Number, "elapsed": elapsed_tp}).Debug("Schedule transactions time")
	// add total fee to the block producer balance before the magnolia hardfork
	if bc.Config.Chain != nil && (bc.Block.Pbft.Number < bc.Config.Chain.Hardforks.MagnoliaHf.BlockNum) {
		bc.accounts.AddToBalance(bc.Block.Pbft.Author, feeReward)
	}
	elapsed := time.Since(start)
	log.WithFields(log.Fields{"func": "processTransactions", "elapsed": elapsed}).Debug("Process transactions time")
	return
}

func (bc *blockContext) processTransaction(t_idx int) (err error) {
	start_transaction := time.Now()
	bc.Block.Transactions[t_idx].SetTimestamp(bc.Block.Pbft.Timestamp)

	bc.SaveTransaction(*bc.Block.Transactions[t_idx].GetModel(), false)
	elapsed_transaction := time.Since(start_transaction)
	log.WithFields(log.Fields{"func": "SaveTransaction", "elapsed": elapsed_transaction}).Debug("Save transaction time")

	start_transaction = time.Now()

	if !bc.Block.Transactions[t_idx].Status {
		return
	}
	// remove value from sender and add to receiver
	receiver := bc.Block.Transactions[t_idx].To
	// handle contract creation
	if receiver == "" {
		receiver = bc.Block.Transactions[t_idx].ContractAddress
	}
	bc.accounts.UpdateBalances(bc.Block.Transactions[t_idx].From, receiver, bc.Block.Transactions[t_idx].Value)
	elapsed_update_balances := time.Since(start_transaction)
	log.WithFields(log.Fields{"func": "UpdateBalances", "elapsed": elapsed_update_balances}).Debug("Update balances time")

	start_transaction = time.Now()
	// process logs
	err = bc.processTransactionLogs(bc.Block.Transactions[t_idx])
	if err != nil {
		return
	}
	elapsed_process_logs := time.Since(start_transaction)
	log.WithFields(log.Fields{"func": "processTransactionLogs", "elapsed": elapsed_process_logs}).Debug("Process logs time")

	start_transaction = time.Now()
	if len(bc.Block.Traces) > 0 {
		if internal_transactions := bc.processInternalTransactions(t_idx); internal_transactions != nil {
			bc.Batch.AddSingleKey(internal_transactions, bc.Block.Transactions[t_idx].Hash)
		}
	}
	elapsed_process_internal_transactions := time.Since(start_transaction)
	log.WithFields(log.Fields{"func": "processInternalTransactions", "elapsed": elapsed_process_internal_transactions}).Debug("Process internal transactions time")

	return
}

func (bc *blockContext) processInternalTransactions(t_idx int) (internal_transactions *models.InternalTransactionsResponse) {
	trace := &bc.Block.Traces[t_idx]
	trx := &bc.Block.Transactions[t_idx]
	if len(trace.Trace) <= 1 {
		return
	}
	internal_transactions = new(models.InternalTransactionsResponse)
	internal_transactions.Data = make([]models.Transaction, 0, len(trace.Trace)-1)

	for e_idx, entry := range trace.Trace {
		if e_idx == 0 {
			continue
		}
		internal := makeInternal(*trx.GetModel(), entry, trx.GasPrice)
		internal_transactions.Data = append(internal_transactions.Data, internal)

		bc.SaveTransaction(internal, true)
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

func (bc *blockContext) SaveTransaction(trx models.Transaction, internal bool) {
	log.WithFields(log.Fields{"from": trx.From, "to": trx.To, "hash": trx.Hash}).Trace("Saving transaction")

	// As the same data is saved with a different keys, it is better to serialize it only once
	trx_bytes, err := rlp.EncodeToBytes(trx)
	if err != nil {
		log.WithFields(log.Fields{"from": trx.From, "to": trx.To, "hash": trx.Hash}).Error("Failed to encode transaction")
	}
	from_index := bc.addressStats.GetAddress(bc.Storage, trx.From).AddTransaction(trx.Timestamp)
	bc.Batch.AddSerialized(trx, trx_bytes, trx.From, from_index)
	if trx.To != "" {
		to_index := bc.addressStats.GetAddress(bc.Storage, trx.To).AddTransaction(trx.Timestamp)
		bc.Batch.AddSerialized(trx, trx_bytes, trx.To, to_index)
	}

	if !internal {
		bc.Batch.AddSerializedSingleKey(trx, trx_bytes, trx.Hash)
	}
}
