package storage

import (
	"math/big"
	"sort"

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

type Accounts struct {
	Accounts []Account `json:"accounts"`
	Total    uint64    `json:"total"`
}

func (a Accounts) ToMap() *AccountBalancesMap {
	am := &AccountBalancesMap{
		accounts: make(map[string]*big.Int),
	}
	for _, account := range a.Accounts {
		am.accounts[account.Address] = account.Balance
	}
	return am
}

type AccountBalancesMap struct {
	accounts map[string]*big.Int
}

func MakeAccountBalancesMap() *AccountBalancesMap {
	return &AccountBalancesMap{
		accounts: make(map[string]*big.Int),
	}
}

func (am *AccountBalancesMap) GetAccounts() map[string]*big.Int {
	return am.accounts
}

func (am *AccountBalancesMap) GetLength() int {
	return len(am.accounts)
}

func (am *AccountBalancesMap) ToSlice() Accounts {
	slice := make([]Account, 0, len(am.accounts))
	for address, balance := range am.accounts {
		slice = append(slice, Account{Address: address, Balance: balance})
	}
	return Accounts{Accounts: slice, Total: uint64(len(am.accounts))}
}

func (am *AccountBalancesMap) Sorted() Accounts {
	sl := am.ToSlice()
	sort.Slice(sl.Accounts, func(i, j int) bool {
		return sl.Accounts[i].Balance.Cmp(sl.Accounts[j].Balance) == 1
	})
	return sl
}

func (am *AccountBalancesMap) GetBalance(address string) *big.Int {
	return am.accounts[address]
}

func (am *AccountBalancesMap) Set(address string, value *big.Int) {
	// make a copy!
	am.accounts[address] = big.NewInt(0).Set(value)
}
