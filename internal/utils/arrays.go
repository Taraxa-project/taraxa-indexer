package utils

import (
	"sort"
	"strings"

	"math/big"

	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
)

func SortByBalanceDescending(array []storage.Account) {
	sort.Slice(array, func(i, j int) bool {
		return array[i].Balance.Cmp(array[j].Balance) == 1
	})
}

func FindBalance(array []storage.Account, address string) int {
	for i, account := range array {
		if strings.EqualFold(account.Address, address) {
			return i
		}
	}

	return -1
}

func RegisterBalance(ptr *[]storage.Account, address string) int {
	array := *ptr
	newAccount := &storage.Account{
		Address: address,
		Balance: big.NewInt(0),
	}

	// Append the new account to the array
	*ptr = append(array, *newAccount)

	// Get the index of the newly added account
	index := len(*ptr) - 1

	return index
}

func RemoveBalance(array *[]storage.Account, address string) {
	unwrapped := *array
	i := FindBalance(unwrapped, address)
	if i != -1 {
		*array = append(unwrapped[:i], unwrapped[i+1:]...)
	}
}

func ModifyBalance(acc *storage.Account, value big.Int) {
	acc.Balance = acc.Balance.Add(acc.Balance, &value)
}

func IsZero(account storage.Account) bool {
	return account.Balance.Cmp(big.NewInt(0)) == 0
}
