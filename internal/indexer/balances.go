package indexer

import (
	"fmt"
	"math/big"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/utils"
	"github.com/Taraxa-project/taraxa-indexer/models"
)

func UpdateBalancesInternal(ptr *[]models.Account, trx models.Transaction) (err error) {

	accounts := *ptr
	parsedValue, ok := new(big.Int).SetString(trx.Value, 10)

	if ok && parsedValue.Cmp(big.NewInt(0)) == 1 {
		i := utils.FindBalance(ptr, trx.From)
		fmt.Println(accounts[i])
		if i == -1 {
			utils.RegisterBalance(ptr, trx.From)
		}
		j := utils.FindBalance(ptr, trx.To)
		if j == -1 {
			utils.RegisterBalance(ptr, trx.To)
		}
		utils.SubstractFromBalance(&accounts[i], *parsedValue)
		if utils.IsZero(accounts[i]) {
			utils.RemoveBalance(ptr, trx.From)
		}
		utils.AddToBalance(&accounts[j], *parsedValue)
	}
	return
}

func UpdateBalances(ptr *[]models.Account, trx *chain.Transaction) error {
	logs := trx.ExtractLogs()

	accounts := *ptr

	if len(logs) > 0 {
		events, err := utils.DecodeRewardsTopics(logs)
		if err != nil {
			return err
		}
		for _, event := range events {
			i := utils.FindBalance(ptr, event.Account)
			utils.AddToBalance(&accounts[i], *event.Value)
		}
	}
	parsedValue, ok := new(big.Int).SetString(trx.Value, 10)
	if ok && parsedValue.Cmp(big.NewInt(0)) == 1 {
		j := utils.FindBalance(ptr, trx.From)
		z := utils.FindBalance(ptr, trx.To)
		utils.SubstractFromBalance(&accounts[j], *parsedValue)
		utils.AddToBalance(&accounts[z], *parsedValue)
		if utils.IsZero(accounts[j]) {
			utils.RemoveBalance(ptr, trx.From)
		}
	}
	return nil
}
