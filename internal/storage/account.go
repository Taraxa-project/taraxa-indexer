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

func (a *Account) ToModel() models.Account {
	return models.Account{
		Address: a.Address,
		Balance: a.Balance.String(),
	}
}

type Accounts []Account

func (a Accounts) SortByBalanceDescending() {
	sort.Slice(a, func(i, j int) bool {
		return a[i].Balance.Cmp(a[j].Balance) == 1
	})
}

func (a Accounts) findIndex(address string) int {
	for i := 0; i < len(a); i++ {
		if strings.EqualFold(a[i].Address, address) {
			return i
		}
	}
	return -1
}

func (a Accounts) FindBalance(address string) *Account {
	i := a.findIndex(address)
	if i == -1 {
		return nil
	}
	return &a[i]
}

func (a *Accounts) RegisterBalance(address string) *Account {
	// Append the new account to the array
	*a = append(*a, Account{
		Address: address,
		Balance: big.NewInt(0),
	})

	return &(*a)[len(*a)-1]
}

func (a *Accounts) RemoveBalance(address string) {
	i := a.findIndex(address)
	if i != -1 {
		*a = append((*a)[:i], (*a)[i+1:]...)
	}
}

func (a Accounts) AddToBalance(address string, value *big.Int) {
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

func (a Accounts) UpdateBalances(from, to, value_str string) {
	from = strings.ToLower(from)
	to = strings.ToLower(to)
	value, ok := big.NewInt(0).SetString(value_str, 0)

	if ok && value.Cmp(big.NewInt(0)) == 1 {
		a.AddToBalance(from, big.NewInt(0).Neg(value))
		a.AddToBalance(to, value)
	}
}

func (a Accounts) UpdateEvents(logs []models.EventLog) error {
	if len(logs) > 0 {
		rewards_events, err := events.DecodeRewardsTopics(logs)
		if err != nil {
			return err
		}
		for _, event := range rewards_events {
			a.AddToBalance(common.DposContractAddress, big.NewInt(0).Neg(event.Value))
			a.AddToBalance(event.Account, event.Value)
		}
	}
	return nil
}
