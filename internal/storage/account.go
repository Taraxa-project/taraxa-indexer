package storage

import (
	"math/big"
	"sort"
	"strings"

	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/events"
	"github.com/Taraxa-project/taraxa-indexer/models"
)

type Account struct {
	Address string   `json:"address"`
	Balance *big.Int `json:"balance"`
}

type AccountReturn struct {
	Address string `json:"address"`
	Balance string `json:"balance"`
}

type Balances struct {
	Accounts []Account `json:"accounts"`
}

func (a *Balances) SortByBalanceDescending() {
	sort.Slice(a.Accounts, func(i, j int) bool {
		return a.Accounts[i].Balance.Cmp(a.Accounts[j].Balance) == 1
	})
}

func (a *Balances) findIndex(address string) int {
	for i := 0; i < len(a.Accounts); i++ {
		if strings.EqualFold(a.Accounts[i].Address, address) {
			return i
		}
	}
	return -1
}

func (a *Balances) FindBalance(address string) *Account {
	i := a.findIndex(address)
	if i == -1 {
		return nil
	}
	return &a.Accounts[i]
}

func (a *Balances) RegisterBalance(address string) *Account {
	// Append the new account to the array
	a.Accounts = append(a.Accounts, Account{
		Address: address,
		Balance: big.NewInt(0),
	})

	return &a.Accounts[len(a.Accounts)-1]
}

func (a *Balances) RemoveBalance(address string) {
	i := a.findIndex(address)
	if i != -1 {
		a.Accounts = append(a.Accounts[:i], a.Accounts[i+1:]...)
	}
}

func (a *Balances) AddToBalance(address string, value *big.Int) {
	address = strings.ToLower(address)
	account := a.FindBalance(address)
	if account == nil {
		account = a.RegisterBalance(address)
	}
	account.Balance.Add(account.Balance, value)
	if account.Balance.Cmp(big.NewInt(0)) == 0 {
		a.RemoveBalance(address)
	}
}

func (a *Balances) UpdateBalances(from, to, value_str string) {
	from = strings.ToLower(from)
	to = strings.ToLower(to)
	value, ok := big.NewInt(0).SetString(value_str, 0)

	if ok && value.Cmp(big.NewInt(0)) == 1 {
		a.AddToBalance(from, big.NewInt(0).Neg(value))
		a.AddToBalance(to, value)
	}
}

func (a *Balances) UpdateEvents(logs []models.EventLog) error {
	if len(logs) > 0 {
		rewards_events, err := events.DecodeRewardsTopics(logs)
		if err != nil {
			return err
		}
		for _, event := range rewards_events {
			to_account := a.FindBalance(event.Account)
			if to_account == nil {
				to_account = a.RegisterBalance(event.Account)
			}
			to_account.Balance.Add(to_account.Balance, event.Value)

			from_account := a.FindBalance(common.DposContractAddress)
			if from_account == nil {
				from_account = a.RegisterBalance(event.Account)
			}
			from_account.Balance.Sub(from_account.Balance, event.Value)
		}
	}
	return nil
}
