package indexer

import (
	"testing"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage/pebble"
	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/stretchr/testify/assert"
)

func prepareBlockContext(timestamp uint64, contractAddresses []string) *blockContext {
	st := pebble.NewStorage("")
	dayStats := storage.MakeDayStatsWithTimestamp(common.DayStart(timestamp))
	config := &common.Config{}
	bc := MakeBlockContext(st, nil, config, dayStats)

	// Create mock block data with the specified timestamp
	bd := &chain.BlockData{
		Pbft: &common.Block{
			Pbft: models.Pbft{
				Timestamp: timestamp,
				Number:    1,
			},
		},
		Transactions: []common.Transaction{},
	}
	bc.SetBlockData(bd)

	for _, contractAddr := range contractAddresses {
		bc.addressStats.GetAddress(st, contractAddr).ContractRegisteredTimestamp = &timestamp
	}

	return bc
}

func TestTrackContractUser_BasicFunctionality(t *testing.T) {
	contractAddr := "0x1234567890abcdef1234567890abcdef12345678"

	bc := prepareBlockContext(1640995200, []string{contractAddr}) // 2022-01-01 00:00:00

	userAddr := "0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"

	// Track a user interacting with a contract
	bc.addContractUser(userAddr, contractAddr)

	// Verify the user was added
	users := bc.getDailyContractUsers(contractAddr)
	assert.True(t, users.Users.Contains(userAddr))
	assert.Equal(t, 1, users.Users.Cardinality())
}

func TestTrackContractUser_MultipleUsers(t *testing.T) {
	contractAddr := "0x1234567890abcdef1234567890abcdef12345678"

	bc := prepareBlockContext(1640995200, []string{contractAddr}) // 2022-01-01 00:00:00

	user1 := "0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"
	user2 := "0x1111111111111111111111111111111111111111"
	user3 := "0x2222222222222222222222222222222222222222"

	// Track multiple users
	bc.addContractUser(user1, contractAddr)
	bc.addContractUser(user2, contractAddr)
	bc.addContractUser(user3, contractAddr)

	// Verify all users are tracked
	users := bc.getDailyContractUsers(contractAddr)
	assert.True(t, users.Users.Contains(user1))
	assert.True(t, users.Users.Contains(user2))
	assert.True(t, users.Users.Contains(user3))
	assert.Equal(t, 3, users.Users.Cardinality())
}

func TestTrackContractUser_DuplicateUser(t *testing.T) {
	contractAddr := "0x1234567890abcdef1234567890abcdef12345678"

	bc := prepareBlockContext(1640995200, []string{contractAddr}) // 2022-01-01 00:00:00

	userAddr := "0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"

	// Track the same user multiple times
	bc.addContractUser(userAddr, contractAddr)
	bc.addContractUser(userAddr, contractAddr)
	bc.addContractUser(userAddr, contractAddr)

	// Verify only one entry (sets handle duplicates)
	users := bc.getDailyContractUsers(contractAddr)
	assert.True(t, users.Users.Contains(userAddr))
	assert.Equal(t, 1, users.Users.Cardinality())
}

func TestTrackContractUser_MultipleContracts(t *testing.T) {
	contract1 := "0x1234567890abcdef1234567890abcdef12345678"
	contract2 := "0x8765432187654321876543218765432187654321"

	bc := prepareBlockContext(1640995200, []string{contract1, contract2}) // 2022-01-01 00:00:00

	user1 := "0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"
	user2 := "0x1111111111111111111111111111111111111111"

	// Track users for different contracts
	bc.addContractUser(user1, contract1)
	bc.addContractUser(user2, contract2)
	bc.addContractUser(user2, contract1) // user2 uses both contracts

	// Verify contract1 tracking
	users1 := bc.getDailyContractUsers(contract1)
	assert.True(t, users1.Users.Contains(user1))
	assert.True(t, users1.Users.Contains(user2))
	assert.Equal(t, 2, users1.Users.Cardinality())

	// Verify contract2 tracking
	users2 := bc.getDailyContractUsers(contract2)
	assert.True(t, users2.Users.Contains(user2))
	assert.False(t, users2.Users.Contains(user1))
	assert.Equal(t, 1, users2.Users.Cardinality())
}

func TestTrackContractUser_EmptyContractAddress(t *testing.T) {
	bc := prepareBlockContext(1640995200, []string{}) // 2022-01-01 00:00:00

	userAddr := "0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"

	// Track with empty contract address (should be ignored)
	bc.addContractUser("", userAddr)

	// Verify no data was created
	assert.Equal(t, 0, len(bc.dailyContractUsers))
}

func TestSaveDailyContractUsers_Persistence(t *testing.T) {
	contractAddr := "0x1234567890abcdef1234567890abcdef12345678"

	bc := prepareBlockContext(1640995200, []string{contractAddr}) // 2022-01-01 00:00:00

	user1 := "0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"
	user2 := "0x1111111111111111111111111111111111111111"

	// Track users
	bc.addContractUser(user1, contractAddr)
	bc.addContractUser(user2, contractAddr)

	// Save to storage
	bc.saveDailyContractUsers()
	bc.Batch.CommitBatch()

	// Verify data was saved by reading from storage
	dayStart := common.DayStart(bc.Block.Pbft.Timestamp)
	savedUsers := bc.Storage.GetDailyContractUsers(contractAddr, dayStart)

	assert.Equal(t, 2, len(savedUsers.Users), "Should have saved two users")
	assert.Contains(t, savedUsers.Users, user1, "User1 should be in saved data")
	assert.Contains(t, savedUsers.Users, user2, "User2 should be in saved data")
}

func TestGetDailyContractUsers_LoadExistingData(t *testing.T) {
	contractAddr := "0x1234567890abcdef1234567890abcdef12345678"

	bc := prepareBlockContext(1640995200, []string{contractAddr}) // 2022-01-01 00:00:00

	existingUser := "0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"
	newUser := "0x1111111111111111111111111111111111111111"

	// Pre-populate storage with existing data
	dayStart := common.DayStart(bc.Block.Pbft.Timestamp)
	existingUsers := storage.MakeDailyContractUsers()
	existingUsers.Add(existingUser)
	bc.Batch.AddDailyContractUsers(contractAddr, dayStart, existingUsers)
	bc.Batch.CommitBatch()

	// Create new block context (simulates new block processing)
	bc2 := prepareBlockContext(1640995200, []string{contractAddr}) // Same day
	bc2.Storage = bc.Storage                                       // Use same storage

	// Track a new user (should merge with existing)
	bc2.addContractUser(newUser, contractAddr)

	// Verify both users are present
	users := bc2.getDailyContractUsers(contractAddr)
	assert.True(t, users.Users.Contains(existingUser), "Existing user should be loaded")
	assert.True(t, users.Users.Contains(newUser), "New user should be added")
	assert.Equal(t, 2, users.Users.Cardinality(), "Should have both users")
}

func TestProcessTransactionContractCall_Integration(t *testing.T) {
	contractAddr := "0x1234567890abcdef1234567890abcdef12345678"

	bc := prepareBlockContext(1640995200, []string{contractAddr}) // 2022-01-01 00:00:00

	userAddr := "0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"

	// Create a contract call transaction
	trx := common.Transaction{
		Transaction: models.Transaction{
			From:   userAddr,
			To:     contractAddr,
			Type:   models.ContractCall,
			Status: true,
		},
	}

	// Add transaction to block
	bc.Block.Transactions = []common.Transaction{trx}

	// Process the transaction
	err := bc.processTransaction(0)
	assert.NoError(t, err)

	// Verify user was tracked
	users := bc.getDailyContractUsers(contractAddr)
	assert.True(t, users.Users.Contains(userAddr), "User should be tracked for contract call")
}

func TestProcessTransactionContractCreation_Integration(t *testing.T) {
	contractAddr := "0x1234567890abcdef1234567890abcdef12345678"
	bc := prepareBlockContext(1640995200, []string{contractAddr}) // 2022-01-01 00:00:00

	userAddr := "0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"

	// Create a contract creation transaction
	trx := common.Transaction{
		Transaction: models.Transaction{
			From:   userAddr,
			To:     "", // Empty for contract creation
			Type:   models.ContractCreation,
			Status: true,
		},
		ContractAddress: contractAddr,
	}

	// Add transaction to block
	bc.Block.Transactions = []common.Transaction{trx}

	// Process the transaction
	err := bc.processTransaction(0)
	assert.NoError(t, err)

	// Verify user was tracked for contract creation
	users := bc.getDailyContractUsers(contractAddr)
	assert.True(t, users.Users.Contains(userAddr), "User should be tracked for contract creation")
}

func TestProcessTransactionNonContractCall_NotTracked(t *testing.T) {
	userAddr := "0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"
	receiverAddr := "0x1234567890abcdef1234567890abcdef12345678"
	bc := prepareBlockContext(1640995200, []string{}) // 2022-01-01 00:00:00

	// Create a regular transfer transaction
	trx := common.Transaction{
		Transaction: models.Transaction{
			From:   userAddr,
			To:     receiverAddr,
			Type:   models.Transfer, // Not a contract interaction
			Status: true,
		},
	}

	// Add transaction to block
	bc.Block.Transactions = []common.Transaction{trx}

	// Process the transaction
	err := bc.processTransaction(0)
	assert.NoError(t, err)

	// Verify no contract users were tracked
	assert.Equal(t, 0, len(bc.dailyContractUsers), "No contract users should be tracked for regular transfer")
}

func TestProcessTransactionFailedTransaction_NotTracked(t *testing.T) {
	contractAddr := "0x1234567890abcdef1234567890abcdef12345678"
	bc := prepareBlockContext(1640995200, []string{contractAddr}) // 2022-01-01 00:00:00

	userAddr := "0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"

	// Create a failed contract call transaction
	trx := common.Transaction{
		Transaction: models.Transaction{
			From:   userAddr,
			To:     contractAddr,
			Type:   models.ContractCall,
			Status: false, // Failed transaction
		},
	}

	// Add transaction to block
	bc.Block.Transactions = []common.Transaction{trx}

	// Process the transaction
	err := bc.processTransaction(0)
	assert.NoError(t, err)

	// Verify no contract users were tracked (failed transactions are skipped)
	assert.Equal(t, 0, len(bc.dailyContractUsers), "No contract users should be tracked for failed transaction")
}
