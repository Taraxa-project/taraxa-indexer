package utils

import (
	"math/big"
	"testing"

	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
)

func TestSortByBalanceDescending(t *testing.T) {
	// Create test data
	accounts := []storage.Account{
		{Address: "0x1111111111111111111111111111111111111111", Balance: big.NewInt(100)},
		{Address: "0x2222222222222222222222222222222222222222", Balance: big.NewInt(50)},
		{Address: "0x3333333333333333333333333333333333333333", Balance: big.NewInt(200)},
	}

	// Expected result after sorting
	expected := []storage.Account{
		{Address: "0x3333333333333333333333333333333333333333", Balance: big.NewInt(200)},
		{Address: "0x1111111111111111111111111111111111111111", Balance: big.NewInt(100)},
		{Address: "0x2222222222222222222222222222222222222222", Balance: big.NewInt(50)},
	}

	// Sort the accounts
	SortByBalanceDescending(accounts)

	// Compare the sorted accounts with the expected result
	if len(accounts) != len(expected) {
		t.Fatalf("Unexpected length of sorted accounts. Got %d, expected %d", len(accounts), len(expected))
	}

	for i, acc := range accounts {
		if acc.Address != expected[i].Address || acc.Balance.Cmp(expected[i].Balance) != 0 {
			t.Errorf("Mismatch in sorted account at index %d. Got %v, expected %v", i, acc, expected[i])
		}
	}
}

func TestFindBalance(t *testing.T) {
	// Create test data
	accounts := []storage.Account{
		{Address: "0x1111111111111111111111111111111111111111", Balance: big.NewInt(100)},
		{Address: "0x2222222222222222222222222222222222222222", Balance: big.NewInt(50)},
		{Address: "0x3333333333333333333333333333333333333333", Balance: big.NewInt(200)},
	}

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
		idx := FindBalance(accounts, test.address)
		if idx != test.expectedIdx {
			t.Errorf("Mismatch in FindBalance result for address %s. Got %d, expected %d", test.address, idx, test.expectedIdx)
		}
	}
}

func TestRegisterBalance(t *testing.T) {
	// Create test data
	accounts := []storage.Account{
		{Address: "0x1111111111111111111111111111111111111111", Balance: big.NewInt(100)},
		{Address: "0x2222222222222222222222222222222222222222", Balance: big.NewInt(50)},
	}

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
		idx := RegisterBalance(&accounts, test.address)
		t.Log(accounts)
		if idx != test.expectedIdx {
			t.Errorf("Mismatch in RegisterBalance result for address %s. Got %d, expected %d", test.address, idx, test.expectedIdx)
		}
	}

	// Verify that the accounts array has been modified
	expectedAccounts := []storage.Account{
		{Address: "0x1111111111111111111111111111111111111111", Balance: big.NewInt(100)},
		{Address: "0x2222222222222222222222222222222222222222", Balance: big.NewInt(50)},
		{Address: "0x3333333333333333333333333333333333333333", Balance: big.NewInt(0)},
		{Address: "0x4444444444444444444444444444444444444444", Balance: big.NewInt(0)},
	}

	if len(accounts) != len(expectedAccounts) {
		t.Fatalf("Unexpected length of accounts array. Got %d, expected %d", len(accounts), len(expectedAccounts))
	}

	for i, acc := range accounts {
		if acc.Address != expectedAccounts[i].Address || acc.Balance.Cmp(expectedAccounts[i].Balance) != 0 {
			t.Errorf("Mismatch in account at index %d. Got %v, expected %v", i, acc, expectedAccounts[i])
		}
	}
}

func TestRemoveBalance(t *testing.T) {
	// Create test data
	accounts := []storage.Account{
		{Address: "0x1111111111111111111111111111111111111111", Balance: big.NewInt(100)},
		{Address: "0x2222222222222222222222222222222222222222", Balance: big.NewInt(50)},
		{Address: "0x3333333333333333333333333333333333333333", Balance: big.NewInt(200)},
	}

	// Test cases
	tests := []struct {
		address       string
		expectedArray []storage.Account
	}{
		{"0x2222222222222222222222222222222222222222", []storage.Account{
			{Address: "0x1111111111111111111111111111111111111111", Balance: big.NewInt(100)},
			{Address: "0x3333333333333333333333333333333333333333", Balance: big.NewInt(200)},
		}},
		{"0x4444444444444444444444444444444444444444", []storage.Account{
			{Address: "0x1111111111111111111111111111111111111111", Balance: big.NewInt(100)},
			{Address: "0x3333333333333333333333333333333333333333", Balance: big.NewInt(200)},
		}},
	}

	// Run the tests
	for _, test := range tests {
		RemoveBalance(&accounts, test.address)

		// Verify that the accounts array has been modified
		t.Log(accounts)
		if len(accounts) != len(test.expectedArray) {
			t.Fatalf("Unexpected length of accounts array. Got %d, expected %d", len(accounts), len(test.expectedArray))
		}

		for i, acc := range accounts {
			if acc.Address != test.expectedArray[i].Address || acc.Balance.Cmp(test.expectedArray[i].Balance) != 0 {
				t.Errorf("Mismatch in account at index %d. Got %v, expected %v", i, acc, test.expectedArray[i])
			}
		}
	}
}

func TestAddToBalance(t *testing.T) {
	// Create test data
	account := &storage.Account{Address: "0x1111111111111111111111111111111111111111", Balance: big.NewInt(100)}
	value := big.NewInt(50)

	// Add to balance
	ModifyBalance(account, *value)

	// Verify the updated balance
	expectedBalance := big.NewInt(150)
	if account.Balance.Cmp(expectedBalance) != 0 {
		t.Errorf("Mismatch in account balance. Got %s, expected %s", account.Balance.String(), expectedBalance.String())
	}
}

func TestSubstractFromBalance(t *testing.T) {
	// Create test data
	account := &storage.Account{Address: "0x1111111111111111111111111111111111111111", Balance: big.NewInt(100)}
	value := big.NewInt(-50)

	// Subtract from balance
	ModifyBalance(account, *value)

	// Verify the updated balance
	expectedBalance := big.NewInt(50)
	if account.Balance.Cmp(expectedBalance) != 0 {
		t.Errorf("Mismatch in account balance. Got %s, expected %s", account.Balance.String(), expectedBalance.String())
	}
}

func TestIsZero(t *testing.T) {
	// Create test data
	nonZeroAccount := storage.Account{Address: "0x1111111111111111111111111111111111111111", Balance: big.NewInt(100)}
	zeroAccount := storage.Account{Address: "0x1111111111111111111111111111111111111111", Balance: big.NewInt(0)}

	// Test non-zero account
	result := IsZero(nonZeroAccount)
	if result {
		t.Errorf("Unexpected result for non-zero account. Got %t, expected false", result)
	}

	// Test zero account
	result = IsZero(zeroAccount)
	if !result {
		t.Errorf("Unexpected result for zero account. Got %t, expected true", result)
	}
}
