package migration

import (
	"testing"
	"time"

	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage/pebble"
	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/assert"
)

func TestDailyContractUsers_Apply(t *testing.T) {
	// Create in-memory storage
	s := pebble.NewStorage("")
	batch := s.NewBatch()

	// Create test transactions from the last month
	now := time.Now()
	oneDayAgo := uint64(now.AddDate(0, 0, -1).Unix())
	twoDaysAgo := uint64(now.AddDate(0, 0, -2).Unix())

	contractAddr := "0x1234567890abcdef1234567890abcdef12345678"
	user1 := "0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"
	user2 := "0x1111111111111111111111111111111111111111"

	// Step 1: Create AddressStats for the contract so it can be found
	contractStats := storage.MakeEmptyAddressStats(contractAddr)
	contractStats.ContractRegisteredTimestamp = &oneDayAgo // Mark as contract
	batch.Add(contractStats, contractAddr, 0)

	// Test transactions - only ContractCall transactions are tracked now
	testTransactions := []models.Transaction{
		{
			From:      user1,
			To:        contractAddr,
			Type:      models.ContractCall, // Only contract calls are tracked
			Timestamp: oneDayAgo,
			Hash:      "0xhash1",
		},
		{
			From:      user2,
			To:        contractAddr,
			Type:      models.ContractCall, // Only contract calls are tracked
			Timestamp: oneDayAgo,           // Same day as first transaction
			Hash:      "0xhash2",
		},
		{
			From:      user1,
			To:        contractAddr,
			Type:      models.ContractCall, // Only contract calls are tracked
			Timestamp: twoDaysAgo,          // Different day
			Hash:      "0xhash3",
		},
		{
			From:      user1,
			To:        "0x9999999999999999999999999999999999999999", // Different contract
			Type:      models.ContractCall,
			Timestamp: oneDayAgo,
			Hash:      "0xhash4",
		},
	}

	// Store test transactions - use the contract address as key to make them findable by contract
	contractTrxIndex := uint64(1) // Start from 1 for contract transactions
	for _, trx := range testTransactions {
		trxBytes, err := rlp.EncodeToBytes(trx)
		assert.NoError(t, err)

		// Store transactions for the contract address with incremental index
		if trx.To == contractAddr {
			batch.AddSerialized(trx, trxBytes, trx.To, contractTrxIndex) // Use contract address as key with index
			contractTrxIndex++                                           // Increment index for next transaction
		}
		batch.AddSerializedSingleKey(trx, trxBytes, trx.Hash) // Also store by hash for completeness
	}
	batch.CommitBatch()

	// Create and run the migration
	migration := &DailyContractUsers{}
	migration.Init(nil) // No client needed for this test

	err := migration.Apply(s)
	assert.NoError(t, err, "Migration should complete without error")

	// Verify the data was saved correctly
	day1Start := common.DayStart(oneDayAgo)
	day2Start := common.DayStart(twoDaysAgo)

	// Check day 1 data (should have both users)
	day1Users := s.GetDailyContractUsers(contractAddr, day1Start)
	assert.Equal(t, 2, len(day1Users.Users), "Day 1 should have 2 users")
	assert.Contains(t, day1Users.Users, user1, "Day 1 should contain user1")
	assert.Contains(t, day1Users.Users, user2, "Day 1 should contain user2")

	// Check day 2 data (should have only user1)
	day2Users := s.GetDailyContractUsers(contractAddr, day2Start)
	assert.Equal(t, 1, len(day2Users.Users), "Day 2 should have 1 user")
	assert.Contains(t, day2Users.Users, user1, "Day 2 should contain user1")
	assert.NotContains(t, day2Users.Users, user2, "Day 2 should not contain user2")
}

func TestDailyContractUsers_GetId(t *testing.T) {
	migration := &DailyContractUsers{}
	assert.Equal(t, "daily_contract_users", migration.GetId())
}

func TestDailyContractUsers_OnlyContractCalls(t *testing.T) {
	// Create in-memory storage
	s := pebble.NewStorage("")
	batch := s.NewBatch()

	now := time.Now()
	yesterday := uint64(now.AddDate(0, 0, -1).Unix())

	contractAddr := "0x1234567890abcdef1234567890abcdef12345678"
	user := "0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"

	// Step 1: Create AddressStats for the contract so it can be found
	contractStats := storage.MakeEmptyAddressStats(contractAddr)
	contractStats.ContractRegisteredTimestamp = &yesterday // Mark as contract
	batch.Add(contractStats, contractAddr, 0)

	// Test different transaction types - only ContractCall should be tracked
	testTransactions := []models.Transaction{
		{
			From:      user,
			To:        contractAddr,
			Type:      models.ContractCall, // This should be tracked
			Timestamp: yesterday,
			Hash:      "0xcontractcall",
		},
		{
			From:      user,
			To:        contractAddr,
			Type:      models.ContractCreation, // This should NOT be tracked
			Timestamp: yesterday,
			Hash:      "0xcontractcreation",
		},
		{
			From:      user,
			To:        contractAddr,
			Type:      models.Transfer, // This should NOT be tracked
			Timestamp: yesterday,
			Hash:      "0xtransfer",
		},
	}

	// Store transactions for the contract address
	for i, trx := range testTransactions {
		trxBytes, err := rlp.EncodeToBytes(trx)
		assert.NoError(t, err)

		if trx.To == contractAddr {
			batch.AddSerialized(trx, trxBytes, trx.To, uint64(i+1))
		}
		batch.AddSerializedSingleKey(trx, trxBytes, trx.Hash)
	}
	batch.CommitBatch()

	// Run migration
	migration := &DailyContractUsers{}
	migration.Init(nil)

	err := migration.Apply(s)
	assert.NoError(t, err)

	// Verify only the ContractCall was tracked
	dayStart := common.DayStart(yesterday)
	users := s.GetDailyContractUsers(contractAddr, dayStart)
	assert.Equal(t, 1, len(users.Users), "Should track only ContractCall transactions")
	assert.Contains(t, users.Users, user, "Should contain the user from ContractCall")
}

func TestDailyContractUsers_NoContracts(t *testing.T) {
	// Create in-memory storage with no contracts
	s := pebble.NewStorage("")

	// Create and run the migration
	migration := &DailyContractUsers{}
	migration.Init(nil) // No client needed for this test

	err := migration.Apply(s)
	assert.NoError(t, err, "Migration should complete without error even with no contracts")

	// Verify no data was saved (since there were no contracts)
	now := time.Now()
	yesterday := uint64(now.AddDate(0, 0, -1).Unix())
	dayStart := common.DayStart(yesterday)

	// Try to get data for a non-existent contract
	users := s.GetDailyContractUsers("0x1234567890abcdef1234567890abcdef12345678", dayStart)
	assert.Equal(t, 0, len(users.Users), "Should have no users for non-existent contract")
}
