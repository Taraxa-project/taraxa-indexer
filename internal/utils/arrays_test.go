package utils

import (
	"math/big"
	"testing"

	"github.com/Taraxa-project/taraxa-indexer/models"
)

func TestSortByBalanceDescending(t *testing.T) {
	// Create test data
	accounts := []models.Account{
		{Address: "0x1111111111111111111111111111111111111111", Balance: "100"},
		{Address: "0x2222222222222222222222222222222222222222", Balance: "50"},
		{Address: "0x3333333333333333333333333333333333333333", Balance: "200"},
	}

	// Expected result after sorting
	expected := []models.Account{
		{Address: "0x3333333333333333333333333333333333333333", Balance: "200"},
		{Address: "0x1111111111111111111111111111111111111111", Balance: "100"},
		{Address: "0x2222222222222222222222222222222222222222", Balance: "50"},
	}

	// Sort the accounts
	SortByBalanceDescending(&accounts)

	// Compare the sorted accounts with the expected result
	if len(accounts) != len(expected) {
		t.Fatalf("Unexpected length of sorted accounts. Got %d, expected %d", len(accounts), len(expected))
	}

	for i, acc := range accounts {
		if acc.Address != expected[i].Address || acc.Balance != expected[i].Balance {
			t.Errorf("Mismatch in sorted account at index %d. Got %v, expected %v", i, acc, expected[i])
		}
	}
}

func TestFindBalance(t *testing.T) {
	// Create test data
	accounts := []models.Account{
		{Address: "0x1111111111111111111111111111111111111111", Balance: "100"},
		{Address: "0x2222222222222222222222222222222222222222", Balance: "50"},
		{Address: "0x3333333333333333333333333333333333333333", Balance: "200"},
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
		idx := FindBalance(&accounts, test.address)
		if idx != test.expectedIdx {
			t.Errorf("Mismatch in FindBalance result for address %s. Got %d, expected %d", test.address, idx, test.expectedIdx)
		}
	}
}

func TestRegisterBalance(t *testing.T) {
	// Create test data
	accounts := []models.Account{
		{Address: "0x1111111111111111111111111111111111111111", Balance: "100"},
		{Address: "0x2222222222222222222222222222222222222222", Balance: "50"},
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
	expectedAccounts := []models.Account{
		{Address: "0x1111111111111111111111111111111111111111", Balance: "100"},
		{Address: "0x2222222222222222222222222222222222222222", Balance: "50"},
		{Address: "0x3333333333333333333333333333333333333333", Balance: "0"},
		{Address: "0x4444444444444444444444444444444444444444", Balance: "0"},
	}

	if len(accounts) != len(expectedAccounts) {
		t.Fatalf("Unexpected length of accounts array. Got %d, expected %d", len(accounts), len(expectedAccounts))
	}

	for i, acc := range accounts {
		if acc.Address != expectedAccounts[i].Address || acc.Balance != expectedAccounts[i].Balance {
			t.Errorf("Mismatch in account at index %d. Got %v, expected %v", i, acc, expectedAccounts[i])
		}
	}
}

func TestRemoveBalance(t *testing.T) {
	// Create test data
	accounts := []models.Account{
		{Address: "0x1111111111111111111111111111111111111111", Balance: "100"},
		{Address: "0x2222222222222222222222222222222222222222", Balance: "50"},
		{Address: "0x3333333333333333333333333333333333333333", Balance: "200"},
	}

	// Test cases
	tests := []struct {
		address       string
		expectedArray []models.Account
	}{
		{"0x2222222222222222222222222222222222222222", []models.Account{
			{Address: "0x1111111111111111111111111111111111111111", Balance: "100"},
			{Address: "0x3333333333333333333333333333333333333333", Balance: "200"},
		}},
		{"0x4444444444444444444444444444444444444444", []models.Account{
			{Address: "0x1111111111111111111111111111111111111111", Balance: "100"},
			{Address: "0x3333333333333333333333333333333333333333", Balance: "200"},
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
			if acc.Address != test.expectedArray[i].Address || acc.Balance != test.expectedArray[i].Balance {
				t.Errorf("Mismatch in account at index %d. Got %v, expected %v", i, acc, test.expectedArray[i])
			}
		}
	}
}

func TestAddToBalance(t *testing.T) {
	// Create test data
	account := &models.Account{Address: "0x1111111111111111111111111111111111111111", Balance: "100"}
	value := big.NewInt(50)

	// Add to balance
	AddToBalance(account, *value)

	// Verify the updated balance
	expectedBalance := big.NewInt(150)
	actualBalance, _ := new(big.Int).SetString(account.Balance, 10)
	if actualBalance.Cmp(expectedBalance) != 0 {
		t.Errorf("Mismatch in account balance. Got %s, expected %s", account.Balance, expectedBalance.String())
	}
}

func TestSubstractFromBalance(t *testing.T) {
	// Create test data
	account := &models.Account{Address: "0x1111111111111111111111111111111111111111", Balance: "100"}
	value := big.NewInt(50)

	// Subtract from balance
	SubstractFromBalance(account, *value)

	// Verify the updated balance
	expectedBalance := big.NewInt(50)
	actualBalance, _ := new(big.Int).SetString(account.Balance, 10)
	if actualBalance.Cmp(expectedBalance) != 0 {
		t.Errorf("Mismatch in account balance. Got %s, expected %s", account.Balance, expectedBalance.String())
	}
}

func TestCompareAccounts(t *testing.T) {
	// Create test data
	account1 := models.Account{Address: "0x1111111111111111111111111111111111111111", Balance: "100"}
	account2 := models.Account{Address: "0x2222222222222222222222222222222222222222", Balance: "200"}

	// Compare the accounts
	result := CompareAccounts(account1, account2)

	// Verify the result
	expectedResult := -1 // account1 < account2
	if result != expectedResult {
		t.Errorf("Mismatch in CompareAccounts result. Got %d, expected %d", result, expectedResult)
	}
}

func TestIsZero(t *testing.T) {
	// Create test data
	nonZeroAccount := models.Account{Address: "0x1111111111111111111111111111111111111111", Balance: "100"}
	zeroAccount := models.Account{Address: "0x2222222222222222222222222222222222222222", Balance: "0"}

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
