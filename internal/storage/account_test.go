package storage

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortByBalanceDescending(t *testing.T) {
	// Create test data
	accounts := MakeAccountsMap()
	accounts.AddToBalance("0x1111111111111111111111111111111111111111", big.NewInt(100))
	accounts.AddToBalance("0x2222222222222222222222222222222222222222", big.NewInt(50))
	accounts.AddToBalance("0x3333333333333333333333333333333333333333", big.NewInt(200))

	// Expected result after sorting
	expected := Accounts([]Account{
		{Address: "0x3333333333333333333333333333333333333333", Balance: big.NewInt(200)},
		{Address: "0x1111111111111111111111111111111111111111", Balance: big.NewInt(100)},
		{Address: "0x2222222222222222222222222222222222222222", Balance: big.NewInt(50)},
	})

	// Sort the accounts
	sorted := accounts.SortedSlice()

	// Compare the sorted accounts with the expected result
	assert.Equal(t, len(expected), len(sorted), "SortByBalanceDescending failed to sort the accounts correctly")

	for i, acc := range sorted {
		if acc.Address != expected[i].Address || acc.Balance.Cmp(expected[i].Balance) != 0 {
			t.Errorf("Mismatch in sorted account at index %d. Got %v, expected %v", i, acc, expected[i])
		}
	}
}

func TestModifyBalance(t *testing.T) {
	// Create test data
	accounts := MakeAccountsMap()
	accounts.AddToBalance("0x1111111111111111111111111111111111111111", big.NewInt(100))
	accounts.AddToBalance("0x2222222222222222222222222222222222222222", big.NewInt(50))
	accounts.AddToBalance("0x3333333333333333333333333333333333333333", big.NewInt(200))

	checkBalance := func(address string, expectedBalance *big.Int) {
		balance := accounts.GetBalance(address)
		assert.Equal(t, expectedBalance, balance, "ModifyBalance failed to update the balance correctly")
	}
	checkBalance("0x1111111111111111111111111111111111111111", big.NewInt(100))
	accounts.AddToBalance("0x1111111111111111111111111111111111111111", big.NewInt(100))
	checkBalance("0x1111111111111111111111111111111111111111", big.NewInt(200))
}

func TestUpdateBalances(t *testing.T) {
	// Create test data
	accounts := MakeAccountsMap()
	accounts.AddToBalance("0x1111111111111111111111111111111111111111", big.NewInt(100))
	accounts.AddToBalance("0x2222222222222222222222222222222222222222", big.NewInt(50))
	accounts.AddToBalance("0x3333333333333333333333333333333333333333", big.NewInt(200))

	checkBalance := func(address string, expectedBalance *big.Int) {
		balance := accounts.GetBalance(address)
		assert.Equal(t, expectedBalance, balance, "ModifyBalance failed to update the balance correctly")
	}
	checkBalance("0x1111111111111111111111111111111111111111", big.NewInt(100))
	checkBalance("0x2222222222222222222222222222222222222222", big.NewInt(50))
	accounts.UpdateBalances("0x1111111111111111111111111111111111111111", "0x2222222222222222222222222222222222222222", big.NewInt(50))
	checkBalance("0x1111111111111111111111111111111111111111", big.NewInt(50))
	checkBalance("0x2222222222222222222222222222222222222222", big.NewInt(100))
}

func TestUpdateBalancesInternal(t *testing.T) {
	// Prepare test data
	accounts := MakeAccountsMap()
	accounts.AddToBalance("0x1111111111111111111111111111111111111111", big.NewInt(100))
	accounts.AddToBalance("0x0DC0d841F962759DA25547c686fa440cF6C28C61", big.NewInt(50))
	trx := Transaction{
		From:    "0x1111111111111111111111111111111111111111",
		To:      "0x0DC0d841F962759DA25547c686fa440cF6C28C61",
		GasCost: big.NewInt(1),
		Value:   big.NewInt(20),
	}

	accounts.UpdateBalances(trx.From, trx.To, trx.Value)

	// Validate the updated balances
	{
		balance := accounts.GetBalance("0x1111111111111111111111111111111111111111")
		assert.Equal(t, big.NewInt(100-20), balance, "UpdateBalancesInternal failed to update 'from' balance correctly")
	}

	{
		balance := accounts.GetBalance("0x0DC0d841F962759DA25547c686fa440cF6C28C61")
		assert.Equal(t, big.NewInt(50+20), balance, "UpdateBalancesInternal failed to update 'to' balance correctly")
	}
}
