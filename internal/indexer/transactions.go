package indexer

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
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

	accounts := &storage.Balances{Accounts: bc.Storage.GetAccounts()}
	block_fee := big.NewInt(0)

	internal_transactions := new(models.InternalTransactionsResponse)
	bc.transactions = make([]models.Transaction, len(transactions))
	for t_idx := 0; t_idx < len(transactions); t_idx++ {
		bc.transactions[t_idx] = transactions[t_idx].ToModelWithTimestamp(bc.block.Timestamp)
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
			internal_transactions.Data = append(internal_transactions.Data, internal)
			bc.SaveTransaction(internal)
			accounts.UpdateBalances(internal.From, internal.To, internal.Value)
		}
		logs := models.TransactionLogsResponse{
			Data: transactions[t_idx].ExtractLogs(),
		}
		bc.Batch.AddToBatchSingleKey(logs, bc.transactions[t_idx].Hash)
		bc.Batch.AddToBatchSingleKey(internal_transactions, bc.transactions[t_idx].Hash)

		trx_fee := transactions[t_idx].GetFee()
		block_fee.Add(block_fee, trx_fee)
		accounts.AddToBalance(transactions[t_idx].From, big.NewInt(0).Neg(trx_fee))
		if transactions[t_idx].Status == "0x0" {
			continue
		}
		accounts.UpdateBalances(transactions[t_idx].From, transactions[t_idx].To, transactions[t_idx].Value)
		err := accounts.UpdateEvents(logs.Data)
		if err != nil {
			return err
		}
	}
	accounts.AddToBalance(bc.block.Author, block_fee)
	// if bc.block.Number%1000 == 0 {
	err = bc.checkIndexedBalances(accounts)
	// }
	bc.Batch.SaveAccounts(accounts)
	return
}

func (bc *blockContext) checkIndexedBalances(accounts *storage.Balances) (err error) {
	chainBalances := make(map[string]*big.Int)
	var mutex sync.RWMutex
	tp := common.MakeThreadPool()
	for _, balance := range accounts.Accounts {
		if balance.IsGenesis || balance.Address == common.DposContractAddress {
			continue
		}
		address := balance.Address
		tp.Go(func() {
			mutex.Lock()
			defer mutex.Unlock()
			b, get_err := bc.Client.GetBalanceAtBlock(address, bc.block.Number)
			if get_err != nil {
				err = get_err
				return
			}
			chainBalances[address], _ = big.NewInt(0).SetString(b, 0)
		})
	}
	tp.Wait()

	for _, balance := range accounts.Accounts {
		if balance.IsGenesis || balance.Address == common.DposContractAddress {
			continue
		}
		if balance.Balance.Cmp(chainBalances[balance.Address]) != 0 {
			return fmt.Errorf("balance of %s: %s != %s", balance.Address, balance.Balance, chainBalances[balance.Address])
		}
	}
	return nil
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
	internal.BlockNumber = 0
	return
}

func (bc *blockContext) SaveTransaction(trx models.Transaction) {
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
