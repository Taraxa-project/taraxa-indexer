package indexer

import (
	"math/big"
	"sort"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/internal/utils"
	"github.com/Taraxa-project/taraxa-indexer/models"
	log "github.com/sirupsen/logrus"
)

func (bc *blockContext) processTransactions(trxHashes *[]string) (err error) {
	var traces []chain.TransactionTrace
	var transactions []chain.Transaction
	var balances map[string]*storage.Account

	tp := utils.MakeThreadPool()
	tp.Go(utils.MakeTaskWithResult(bc.client.TraceBlockTransactions, bc.block.Number, &traces, &err).Run)
	tp.Go(utils.MakeTaskWithResult(bc.getTransactions, trxHashes, &transactions, &err).Run)
	tp.Go(utils.MakeTaskWithResult(bc.readBalanceMap, balances, &balances, &err).Run)
	tp.Wait()

	if err != nil {
		return
	}

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
				tp.Go(utils.MakeTask(bc.updateInternalHolderBalances, internal, &err).Run)
				internal_transactions.Data = append(internal_transactions.Data, internal)
				bc.SaveTransaction(&internal)
			}
			bc.batch.AddToBatchSingleKey(internal_transactions, trx_model.Hash)
		}
		logs := models.TransactionLogsResponse{
			Data: trx.ExtractLogs(),
		}
		bc.batch.AddToBatchSingleKey(logs, trx_model.Hash)
		tp.Go(utils.MakeTask(bc.updateHolderBalances, chain.AccountParams{Transaction: &trx, BalanceMap: &balances}, &err).Run)
		tp.Wait()
	}
	sortBalances(&balances)
	bc.batch.AddToBatchSingleKey(balances, "0x0")
	return
}

func (bc *blockContext) readBalanceMap(balance map[string]*storage.Account) (newBalances map[string]*storage.Account, err error) {
	newBalances = *bc.storage.GetAccounts()
	if newBalances == nil {
		panic("cannot read balances from storage")
	}
	return newBalances, nil
}

func sortBalances(balances *map[string]*storage.Account) {
	var accountSlice []*storage.Account
	for _, account := range *balances {
		account.Mutex.RLock()
		accountSlice = append(accountSlice, account)
		account.Mutex.RUnlock()
	}

	// Sort the slice based on balances
	sort.Slice(accountSlice, func(i, j int) bool {
		return accountSlice[i].Balance.Cmp(accountSlice[j].Balance) < 0
	})

	// Rebuild the sorted accounts map if necessary
	unwrappedBalances := *balances
	for _, account := range accountSlice {
		unwrappedBalances[account.Address] = account
	}
}

func (bc *blockContext) updateHolderBalances(params chain.AccountParams) (err error) {
	balances, trx := *params.BalanceMap, *params.Transaction
	logs := trx.ExtractLogs()

	if len(logs) > 0 {
		events, err := utils.DecodeRewardsTopics(logs)
		if err != nil {
			return err
		}
		for _, event := range events {
			balances[event.Account].AddToBalance(*event.Value)
			if balances[event.Account].Balance.Cmp(big.NewInt(0)) != 1 {
				delete(balances, event.Account)
			}
		}
	}
	parsedValue, ok := new(big.Int).SetString(trx.Value, 10)
	if ok && parsedValue.Cmp(big.NewInt(0)) == 1 {
		fromBalance := balances[trx.From]
		toBalance := balances[trx.To]
		fromBalance.SubstractFromBalance(*parsedValue)
		toBalance.AddToBalance(*parsedValue)
		if fromBalance.Balance.Cmp(big.NewInt(0)) != 1 {
			delete(balances, trx.From)
		}
		if toBalance.Balance.Cmp(big.NewInt(0)) != 1 {
			delete(balances, trx.To)
		}
	}
	return
}

func (bc *blockContext) updateInternalHolderBalances(trx models.Transaction) (err error) {
	parsedValue, ok := new(big.Int).SetString(trx.Value, 10)

	if ok && parsedValue.Cmp(big.NewInt(0)) == 1 {
		fromBalance := bc.storage.GetAccount(trx.From)
		toBalance := bc.storage.GetAccount(trx.To)
		fromBalance.SubstractFromBalance(*parsedValue)
		toBalance.AddToBalance(*parsedValue)
		bc.batch.AddToBatchSingleKey(fromBalance.ToModel(), trx.From)
		bc.batch.AddToBatchSingleKey(toBalance.ToModel(), trx.To)
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

	tp := utils.MakeThreadPool()
	for i, trx_hash := range *trxHashes {
		tp.Go(utils.MakeTaskWithResult(bc.client.GetTransactionByHash, trx_hash, &trxs[i], &err).Run)
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

func (bc *blockContext) addAddressStatsToBatch() {
	for _, stats := range bc.addressStats {
		bc.batch.AddToBatch(stats, stats.Address, 0)
	}
}
