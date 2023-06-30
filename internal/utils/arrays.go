package utils

import (
	"sort"
	"strings"

	"math/big"

	"github.com/Taraxa-project/taraxa-indexer/models"
)

func SortByBalanceDescending(ptr *[]models.Account) {
	array := *ptr
	sort.Slice(array, func(i, j int) bool {
		firstBalance, _ := new(big.Int).SetString(array[i].Balance, 10)
		secondBalance, _ := new(big.Int).SetString(array[j].Balance, 10)
		return firstBalance.Cmp(secondBalance) == 1
	})
	*ptr = array
}

func FindBalance(ptr *[]models.Account, address string) int {
	array := *ptr
	for i, account := range array {
		if strings.ToLower(account.Address) == strings.ToLower(address) {
			return i
		}
	}

	return -1
}

func RegisterBalance(ptr *[]models.Account, address string) int {
	array := *ptr
	newAccount := &models.Account{
		Address: address,
		Balance: "0",
	}

	// Append the new account to the array
	*ptr = append(array, *newAccount)

	// Get the index of the newly added account
	index := len(*ptr) - 1

	return index
}

func RemoveBalance(array *[]models.Account, address string) {
	i := FindBalance(array, address)
	if i != -1 {
		unwrapped := *array
		*array = append(unwrapped[:i], unwrapped[i+1:]...)
	}
}

func AddToBalance(acc *models.Account, value big.Int) {
	balance, _ := new(big.Int).SetString(acc.Balance, 10)
	balance.Add(balance, &value)
	acc.Balance = balance.String()
}

func SubstractFromBalance(acc *models.Account, value big.Int) {
	balance, _ := new(big.Int).SetString(acc.Balance, 10)
	balance.Sub(balance, &value)
	acc.Balance = balance.String()
}

func CompareAccounts(first, second models.Account) int {
	firstBalance, _ := new(big.Int).SetString(first.Balance, 10)
	secondBalance, _ := new(big.Int).SetString(second.Balance, 10)
	return firstBalance.Cmp(secondBalance)
}

func IsZero(account models.Account) bool {
	firstBalance, _ := new(big.Int).SetString(account.Balance, 10)
	return firstBalance.Cmp(big.NewInt(0)) == 0
}
