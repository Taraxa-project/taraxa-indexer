package indexer

import (
	"math/big"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/internal/utils"
	"github.com/Taraxa-project/taraxa-indexer/models"
)

func UpdateBalancesInternal(storage storage.Storage, trx models.Transaction) (err error) {
	accounts := storage.GetAccounts()

	parsedValue, ok := new(big.Int).SetString(trx.Value, 10)

	if ok && parsedValue.Cmp(big.NewInt(0)) == 1 {
		fromBalance, _ := utils.FindBalance(accounts, trx.From)
		toBalance, _ := utils.FindBalance(accounts, trx.To)
		utils.SubstractFromBalance(fromBalance, *parsedValue)
		utils.AddToBalance(toBalance, *parsedValue)
	}
	return
}

func UpdateBalances(storage storage.Storage, trx *chain.Transaction) error {
	accounts := storage.GetAccounts()

	logs := trx.ExtractLogs()

	if len(logs) > 0 {
		events, err := utils.DecodeRewardsTopics(logs)
		if err != nil {
			return err
		}
		for _, event := range events {
			account, _ := utils.FindBalance(accounts, event.Account)
			utils.AddToBalance(account, *event.Value)
			if utils.IsZero(*account) {
				utils.RemoveBalance(&accounts, event.Account)
			}
		}
	}
	parsedValue, ok := new(big.Int).SetString(trx.Value, 10)
	if ok && parsedValue.Cmp(big.NewInt(0)) == 1 {
		fromBalance, _ := utils.FindBalance(accounts, trx.From)
		toBalance, _ := utils.FindBalance(accounts, trx.To)
		utils.SubstractFromBalance(fromBalance, *parsedValue)
		utils.AddToBalance(toBalance, *parsedValue)
		if utils.IsZero(*fromBalance) {
			utils.RemoveBalance(&accounts, trx.From)
		}
		if utils.IsZero(*toBalance) {
			utils.RemoveBalance(&accounts, trx.To)
		}
	}
	return nil
}
