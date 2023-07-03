package indexer

import (
	"fmt"
	"math/big"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/utils"
	"github.com/Taraxa-project/taraxa-indexer/models"
	log "github.com/sirupsen/logrus"
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

	if len(logs) > 0 {
		events, err := utils.DecodeRewardsTopics(logs)
		if err != nil {
			return err
		}
		for _, event := range events {
			i := utils.FindBalance(ptr, event.Account)
			if i == -1 {
				i = utils.RegisterBalance(ptr, event.Account)
			}
			newArr := *ptr
			utils.AddToBalance(&newArr[i], *event.Value)
		}
	}
	parsedValue, ok := new(big.Int).SetString(trx.Value, 10)
	if ok && parsedValue.Cmp(big.NewInt(0)) == 1 {
		j := utils.FindBalance(ptr, trx.From)
		if j == -1 {
			log.Debug("Could not find balance for subtracting", trx.From)
		}
		z := utils.FindBalance(ptr, trx.To)
		if z == -1 {
			z = utils.RegisterBalance(ptr, trx.From)
		}
		newArr := *ptr
		utils.SubstractFromBalance(&newArr[j], *parsedValue)
		utils.AddToBalance(&newArr[z], *parsedValue)
		if utils.IsZero(newArr[j]) {
			utils.RemoveBalance(ptr, trx.From)
		}
	}
	return nil
}
