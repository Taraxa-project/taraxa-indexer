package indexer

import (
	"fmt"
	"math/big"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/internal/utils"
	"github.com/Taraxa-project/taraxa-indexer/models"
	log "github.com/sirupsen/logrus"
)

func UpdateBalancesInternal(ptr *[]storage.Account, trx models.Transaction) (err error) {

	accounts := *ptr
	parsedValue, ok := new(big.Int).SetString(trx.Value, 10)

	if ok && parsedValue.Cmp(big.NewInt(0)) == 1 {
		i := utils.FindBalance(accounts, trx.From)
		fmt.Println(accounts[i])
		if i == -1 {
			utils.RegisterBalance(ptr, trx.From)
		}
		j := utils.FindBalance(accounts, trx.To)
		if j == -1 {
			utils.RegisterBalance(ptr, trx.To)
		}

		utils.ModifyBalance(&accounts[i], *big.NewInt(0).Neg(parsedValue))
		if utils.IsZero(accounts[i]) {
			utils.RemoveBalance(ptr, trx.From)
		}
		utils.ModifyBalance(&accounts[j], *parsedValue)
	}
	return
}

func UpdateBalances(ptr *[]storage.Account, trx *chain.Transaction) error {
	logs := trx.ExtractLogs()
	accounts := *ptr

	if len(logs) > 0 {
		events, err := utils.DecodeRewardsTopics(logs)
		if err != nil {
			return err
		}
		for _, event := range events {
			i := utils.FindBalance(accounts, event.Account)
			if i == -1 {
				i = utils.RegisterBalance(ptr, event.Account)
			}
			newArr := *ptr
			utils.ModifyBalance(&newArr[i], *event.Value)
		}
	}
	parsedValue, ok := new(big.Int).SetString(trx.Value, 10)
	if ok && parsedValue.Cmp(big.NewInt(0)) == 1 {
		j := utils.FindBalance(accounts, trx.From)
		if j == -1 {
			log.Debug("Could not find balance for subtracting", trx.From)
		}
		z := utils.FindBalance(accounts, trx.To)
		if z == -1 {
			z = utils.RegisterBalance(ptr, trx.From)
		}
		newArr := *ptr
		utils.ModifyBalance(&newArr[j], *big.NewInt(0).Neg(parsedValue))
		utils.ModifyBalance(&newArr[z], *parsedValue)
		if utils.IsZero(newArr[j]) {
			utils.RemoveBalance(ptr, trx.From)
		}
	}
	return nil
}

func (bc *blockContext) CheckIndexedBalances() error {
	balances := bc.storage.GetAccounts()

	for _, balance := range balances {
		if balance.Balance.Cmp(big.NewInt(0)) == -1 {
			return fmt.Errorf("balance of %s is negative", balance.Address)
		}
		balanceFromChain, err := bc.client.GetBalanceFromBlock(balance.Address, bc.block.Number)
		if err != nil {
			return err
		}
		balanceInt, ok := new(big.Int).SetString(balanceFromChain, 16)
		if !ok {
			return fmt.Errorf("could not parse balance from chain %s", balanceFromChain)
		}
		if balance.Balance.Cmp(balanceInt) != 0 {
			return fmt.Errorf("balance of %s is not equal to balance from chain", balance.Address)
		}
	}
	return nil
}
