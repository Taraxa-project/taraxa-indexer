package storage

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortByBalanceDescending(t *testing.T) {
	// Create test data
	accounts_map := MakeAccountBalancesMap()
	accounts := accounts_map.GetAccounts()
	accounts["0x1111111111111111111111111111111111111111"] = big.NewInt(100)
	accounts["0x2222222222222222222222222222222222222222"] = big.NewInt(50)
	accounts["0x3333333333333333333333333333333333333333"] = big.NewInt(200)

	// Expected result after sorting
	expected := []Account{
		{Address: "0x3333333333333333333333333333333333333333", Balance: big.NewInt(200)},
		{Address: "0x1111111111111111111111111111111111111111", Balance: big.NewInt(100)},
		{Address: "0x2222222222222222222222222222222222222222", Balance: big.NewInt(50)},
	}

	// Sort the accounts
	sorted := accounts_map.Sorted()

	// Compare the sorted accounts with the expected result
	assert.Equal(t, len(expected), len(sorted.Accounts), "SortByBalanceDescending failed to sort the accounts correctly")

	for i, acc := range sorted.Accounts {
		if acc.Address != expected[i].Address || acc.Balance.Cmp(expected[i].Balance) != 0 {
			t.Errorf("Mismatch in sorted account at index %d. Got %v, expected %v", i, acc, expected[i])
		}
	}
}
