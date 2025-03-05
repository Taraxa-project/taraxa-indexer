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

func (a Accounts) ToMap() AccountsMap {
	am := make(AccountsMap)
	for _, account := range a {
		am[account.Address] = account.Balance
	}
	return am
}

type AccountsMap map[string]*big.Int

func (am AccountsMap) toSlice() Accounts {
	slice := make(Accounts, 0, len(am))
	for address, balance := range am {
		slice = append(slice, Account{Address: address, Balance: balance})
	}
	return slice
}

func (am AccountsMap) SortedSlice() Accounts {
	sl := am.toSlice()
	sort.Slice(sl, func(i, j int) bool {
		return sl[i].Balance.Cmp(sl[j].Balance) == 1
	})
	return sl
}

func (am AccountsMap) GetBalance(address string) *big.Int {
	address = strings.ToLower(address)
	return am[address]
}

func (am AccountsMap) AddToBalance(address string, value *big.Int) {
	address = strings.ToLower(address)
	if _, ok := am[address]; !ok {
		am[address] = big.NewInt(0)
	}
	am[address].Add(am[address], value)
}

func (am AccountsMap) UpdateBalances(from, to, valueStr string) {
	value, ok := big.NewInt(0).SetString(valueStr, 0)

	if ok && value.Cmp(big.NewInt(0)) > 0 {
		am.AddToBalance(from, big.NewInt(0).Neg(value))
		am.AddToBalance(to, value)
	}
}

func (am AccountsMap) UpdateEvents(logs []models.EventLog) error {
	if len(logs) > 0 {
		rewardsEvents, err := events.DecodeRewardsTopics(logs)
		if err != nil {
			return err
		}
		for _, event := range rewardsEvents {
			am.AddToBalance(common.DposContractAddress, big.NewInt(0).Neg(event.Value))
			am.AddToBalance(event.Account, event.Value)
		}
	}
	return nil
}
