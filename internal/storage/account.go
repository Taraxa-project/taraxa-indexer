package storage

import (
	"math/big"
	"sort"
	"strings"

	"github.com/Taraxa-project/taraxa-indexer/internal/events"
	"github.com/Taraxa-project/taraxa-indexer/models"
)

type Account struct {
	Address   string   `json:"address"`
	Balance   *big.Int `json:"balance"`
	IsGenesis bool     `json:"is_genesis"`
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

func (a *Balances) SetGenesis(address string) {
	account := a.FindBalance(address)
	if account != nil {
		account.IsGenesis = true
	}
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
		Address:   address,
		Balance:   big.NewInt(0),
		IsGenesis: false,
	})

	return &a.Accounts[len(a.Accounts)-1]
}

func (a *Balances) RemoveBalance(address string) {
	i := a.findIndex(address)
	if i != -1 {
		a.Accounts = append(a.Accounts[:i], a.Accounts[i+1:]...)
	}
}

func (a *Balances) AddToBalance(to string, value *big.Int) {
	to_account := a.FindBalance(to)
	if to_account == nil {
		to_account = a.RegisterBalance(to)
	}
	to_account.Balance.Add(to_account.Balance, value)
}

func (a *Balances) UpdateBalances(from, to, value_str, gas_used, gas_price string) {
	value, ok := big.NewInt(0).SetString(value_str, 0)

	if ok && value.Cmp(big.NewInt(0)) == 1 {
		from_account := a.FindBalance(from)
		if from_account == nil {
			from_account = a.RegisterBalance(from)
		}
		from_account.Balance.Sub(from_account.Balance, value)
		if gas_used != "" && gas_price != "" {
			gasUsed, _ := big.NewInt(0).SetString(gas_used, 0)
			gasPrice, _ := big.NewInt(0).SetString(gas_price, 0)
			negGas := big.NewInt(0).Neg(big.NewInt(0).Mul(gasUsed, gasPrice))
			a.AddToBalance(from, negGas)
		}
		if from_account.Balance.Cmp(big.NewInt(0)) == 0 {
			a.RemoveBalance(from)
		}
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

			from_account := a.FindBalance(events.DposContractAddress)
			if from_account == nil {
				from_account = a.RegisterBalance(event.Account)
			}
			from_account.Balance.Sub(from_account.Balance, event.Value)
		}
	}
	return nil
}
