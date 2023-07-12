package storage

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortByBalanceDescending(t *testing.T) {
	// Create test data
	accounts := &Balances{Accounts: []Account{
		{Address: "0x1111111111111111111111111111111111111111", Balance: big.NewInt(100)},
		{Address: "0x2222222222222222222222222222222222222222", Balance: big.NewInt(50)},
		{Address: "0x3333333333333333333333333333333333333333", Balance: big.NewInt(200)},
	}}

	// Expected result after sorting
	expected := &Balances{Accounts: []Account{
		{Address: "0x3333333333333333333333333333333333333333", Balance: big.NewInt(200)},
		{Address: "0x1111111111111111111111111111111111111111", Balance: big.NewInt(100)},
		{Address: "0x2222222222222222222222222222222222222222", Balance: big.NewInt(50)},
	}}

	// Sort the accounts
	accounts.SortByBalanceDescending()

	// Compare the sorted accounts with the expected result
	assert.Equal(t, len(expected.Accounts), len(accounts.Accounts), "SortByBalanceDescending failed to sort the accounts correctly")

	for i, acc := range accounts.Accounts {
		if acc.Address != expected.Accounts[i].Address || acc.Balance.Cmp(expected.Accounts[i].Balance) != 0 {
			t.Errorf("Mismatch in sorted account at index %d. Got %v, expected %v", i, acc, expected.Accounts[i])
		}
	}
}

func TestFindBalance(t *testing.T) {
	// Create test data
	accounts := &Balances{Accounts: []Account{
		{Address: "0x1111111111111111111111111111111111111111", Balance: big.NewInt(100)},
		{Address: "0x2222222222222222222222222222222222222222", Balance: big.NewInt(50)},
		{Address: "0x3333333333333333333333333333333333333333", Balance: big.NewInt(200)},
	}}

	// Test cases
	tests := []struct {
		address     string
		expectedIdx int
	}{
		{"0x1111111111111111111111111111111111111111", 0},
		{"0x2222222222222222222222222222222222222222", 1},
		{"0x3333333333333333333333333333333333333333", 2},
		{"0x4444444444444444444444444444444444444444", -1},
	}

	// Run the tests
	for _, test := range tests {
		idx := accounts.findIndex(test.address)
		if idx != test.expectedIdx {
			t.Errorf("Mismatch in FindBalance result for address %s. Got %d, expected %d", test.address, idx, test.expectedIdx)
		}
	}
}

func TestRegisterBalance(t *testing.T) {
	// Create test data
	accounts := &Balances{Accounts: []Account{
		{Address: "0x1111111111111111111111111111111111111111", Balance: big.NewInt(100)},
		{Address: "0x2222222222222222222222222222222222222222", Balance: big.NewInt(50)},
	}}

	// Test cases
	tests := []struct {
		address     string
		expectedIdx int
	}{
		{"0x3333333333333333333333333333333333333333", 2},
		{"0x4444444444444444444444444444444444444444", 3},
	}

	// Run the tests
	for _, test := range tests {
		bal := accounts.RegisterBalance(test.address)
		idx := accounts.findIndex(bal.Address)
		if idx != test.expectedIdx {
			t.Errorf("Mismatch in RegisterBalance result for address %s. Got %d, expected %d", test.address, idx, test.expectedIdx)
		}
	}

	// Verify that the accounts array has been modified
	expectedAccounts := &Balances{Accounts: []Account{
		{Address: "0x1111111111111111111111111111111111111111", Balance: big.NewInt(100)},
		{Address: "0x2222222222222222222222222222222222222222", Balance: big.NewInt(50)},
		{Address: "0x3333333333333333333333333333333333333333", Balance: big.NewInt(0)},
		{Address: "0x4444444444444444444444444444444444444444", Balance: big.NewInt(0)},
	}}

	if len(accounts.Accounts) != len(expectedAccounts.Accounts) {
		t.Fatalf("Unexpected length of accounts array. Got %d, expected %d", len(accounts.Accounts), len(expectedAccounts.Accounts))
	}

	for i, acc := range accounts.Accounts {
		if acc.Address != expectedAccounts.Accounts[i].Address || acc.Balance.Cmp(expectedAccounts.Accounts[i].Balance) != 0 {
			t.Errorf("Mismatch in account at index %d. Got %v, expected %v", i, acc, expectedAccounts.Accounts[i])
		}
	}
}

func TestRemoveBalance(t *testing.T) {
	// Create test data
	accounts := &Balances{Accounts: []Account{
		{Address: "0x1111111111111111111111111111111111111111", Balance: big.NewInt(100)},
		{Address: "0x2222222222222222222222222222222222222222", Balance: big.NewInt(50)},
		{Address: "0x3333333333333333333333333333333333333333", Balance: big.NewInt(200)},
	}}

	// Test cases
	tests := []struct {
		address       string
		expectedArray []Account
	}{
		{"0x2222222222222222222222222222222222222222", []Account{
			{Address: "0x1111111111111111111111111111111111111111", Balance: big.NewInt(100)},
			{Address: "0x3333333333333333333333333333333333333333", Balance: big.NewInt(200)},
		}},
		{"0x4444444444444444444444444444444444444444", []Account{
			{Address: "0x1111111111111111111111111111111111111111", Balance: big.NewInt(100)},
			{Address: "0x3333333333333333333333333333333333333333", Balance: big.NewInt(200)},
		}},
	}

	// Run the tests
	for _, test := range tests {
		accounts.RemoveBalance(test.address)

		// Verify that the accounts array has been modified
		if len(accounts.Accounts) != len(test.expectedArray) {
			t.Fatalf("Unexpected length of accounts array. Got %d, expected %d", len(accounts.Accounts), len(test.expectedArray))
		}

		for i, acc := range accounts.Accounts {
			if acc.Address != test.expectedArray[i].Address || acc.Balance.Cmp(test.expectedArray[i].Balance) != 0 {
				t.Errorf("Mismatch in account at index %d. Got %v, expected %v", i, acc, test.expectedArray[i])
			}
		}
	}
}
