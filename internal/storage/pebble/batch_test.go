package pebble

import (
	"testing"
	"time"

	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestBatch_AddDailyContractUsers(t *testing.T) {
	// Create test storage
	st := NewStorage("")
	defer func() { _ = st.Close() }()

	contractAddress := "0x1111111111111111111111111111111111111111"
	timestamp := uint64(time.Now().Unix())

	// Create DailyContractUsers with test data
	users := storage.MakeDailyContractUsers()
	users.Add("0x1111111111111111111111111111111111111111")
	users.Add("0x2222222222222222222222222222222222222222")
	users.Add("0x3333333333333333333333333333333333333333")

	// Store using batch
	batch := st.NewBatch()
	batch.AddDailyContractUsers(contractAddress, timestamp, users)
	batch.CommitBatch()

	// Retrieve and verify
	retrieved := st.GetDailyContractUsers(contractAddress, timestamp)

	assert.Len(t, retrieved.Users, 3)
	assert.Contains(t, retrieved.Users, "0x1111111111111111111111111111111111111111")
	assert.Contains(t, retrieved.Users, "0x2222222222222222222222222222222222222222")
	assert.Contains(t, retrieved.Users, "0x3333333333333333333333333333333333333333")
}

func TestBatch_AddDailyContractUsers_EmptyUsers(t *testing.T) {
	// Create test storage
	st := NewStorage("")
	defer func() { _ = st.Close() }()

	contractAddress := "0x1111111111111111111111111111111111111111"
	timestamp := common.DayStart(uint64(time.Now().Unix()))

	// Create empty DailyContractUsers
	users := storage.MakeDailyContractUsers()

	// Store using batch
	batch := st.NewBatch()
	batch.AddDailyContractUsers(contractAddress, timestamp, users)
	batch.CommitBatch()

	// Retrieve and verify
	retrieved := st.GetDailyContractUsers(contractAddress, timestamp)

	assert.Len(t, retrieved.Users, 0)
}

func TestBatch_AddDailyContractUsers_DuplicateUsers(t *testing.T) {
	// Create test storage
	st := NewStorage("")
	defer func() { _ = st.Close() }()

	contractAddress := "0x1111111111111111111111111111111111111111"
	timestamp := common.DayStart(uint64(time.Now().Unix()))

	// Create DailyContractUsers with duplicate users
	users := storage.MakeDailyContractUsers()
	users.Add("0x1111111111111111111111111111111111111111")
	users.Add("0x2222222222222222222222222222222222222222")
	users.Add("0x1111111111111111111111111111111111111111") // duplicate
	users.Add("0x2222222222222222222222222222222222222222") // duplicate
	users.Add("0x3333333333333333333333333333333333333333")

	// Store using batch
	batch := st.NewBatch()
	batch.AddDailyContractUsers(contractAddress, timestamp, users)
	batch.CommitBatch()

	// Retrieve and verify - should only have unique users
	retrieved := st.GetDailyContractUsers(contractAddress, timestamp)

	assert.Len(t, retrieved.Users, 3)
	assert.Contains(t, retrieved.Users, "0x1111111111111111111111111111111111111111")
	assert.Contains(t, retrieved.Users, "0x2222222222222222222222222222222222222222")
	assert.Contains(t, retrieved.Users, "0x3333333333333333333333333333333333333333")
}

func TestBatch_AddDailyContractUsers_MultipleContracts(t *testing.T) {
	// Create test storage
	st := NewStorage("")
	defer func() { _ = st.Close() }()

	contract1 := "0x1111111111111111111111111111111111111111"
	contract2 := "0x2222222222222222222222222222222222222222"
	timestamp := common.DayStart(uint64(time.Now().Unix()))

	// Create DailyContractUsers for contract 1
	users1 := storage.MakeDailyContractUsers()
	users1.Add("0xaaaa000000000000000000000000000000000000")
	users1.Add("0xbbbb000000000000000000000000000000000000")

	// Create DailyContractUsers for contract 2
	users2 := storage.MakeDailyContractUsers()
	users2.Add("0xcccc000000000000000000000000000000000000")
	users2.Add("0xdddd000000000000000000000000000000000000")
	users2.Add("0xeeee000000000000000000000000000000000000")

	// Store using batch
	batch := st.NewBatch()
	batch.AddDailyContractUsers(contract1, timestamp, users1)
	batch.AddDailyContractUsers(contract2, timestamp, users2)
	batch.CommitBatch()

	// Retrieve and verify contract 1
	retrieved1 := st.GetDailyContractUsers(contract1, timestamp)
	assert.Len(t, retrieved1.Users, 2)
	assert.Contains(t, retrieved1.Users, "0xaaaa000000000000000000000000000000000000")
	assert.Contains(t, retrieved1.Users, "0xbbbb000000000000000000000000000000000000")

	// Retrieve and verify contract 2
	retrieved2 := st.GetDailyContractUsers(contract2, timestamp)
	assert.Len(t, retrieved2.Users, 3)
	assert.Contains(t, retrieved2.Users, "0xcccc000000000000000000000000000000000000")
	assert.Contains(t, retrieved2.Users, "0xdddd000000000000000000000000000000000000")
	assert.Contains(t, retrieved2.Users, "0xeeee000000000000000000000000000000000000")
}

func TestBatch_AddDailyContractUsers_MultipleDays(t *testing.T) {
	// Create test storage
	st := NewStorage("")
	defer func() { _ = st.Close() }()

	contractAddress := "0x1111111111111111111111111111111111111111"
	baseTimestamp := common.DayStart(uint64(time.Now().Unix()))

	// Create DailyContractUsers for day 1
	day1 := baseTimestamp - (1 * common.Day)
	users1 := storage.MakeDailyContractUsers()
	users1.Add("0x1111111111111111111111111111111111111111")
	users1.Add("0x2222222222222222222222222222222222222222")

	// Create DailyContractUsers for day 2
	day2 := baseTimestamp - (2 * common.Day)
	users2 := storage.MakeDailyContractUsers()
	users2.Add("0x3333333333333333333333333333333333333333")
	users2.Add("0x4444444444444444444444444444444444444444")
	users2.Add("0x5555555555555555555555555555555555555555")

	// Store using batch
	batch := st.NewBatch()
	batch.AddDailyContractUsers(contractAddress, day1, users1)
	batch.AddDailyContractUsers(contractAddress, day2, users2)
	batch.CommitBatch()

	// Retrieve and verify day 1
	retrieved1 := st.GetDailyContractUsers(contractAddress, day1)
	assert.Len(t, retrieved1.Users, 2)
	assert.Contains(t, retrieved1.Users, "0x1111111111111111111111111111111111111111")
	assert.Contains(t, retrieved1.Users, "0x2222222222222222222222222222222222222222")

	// Retrieve and verify day 2
	retrieved2 := st.GetDailyContractUsers(contractAddress, day2)
	assert.Len(t, retrieved2.Users, 3)
	assert.Contains(t, retrieved2.Users, "0x3333333333333333333333333333333333333333")
	assert.Contains(t, retrieved2.Users, "0x4444444444444444444444444444444444444444")
	assert.Contains(t, retrieved2.Users, "0x5555555555555555555555555555555555555555")
}

func TestBatch_AddDailyContractUsers_Overwrite(t *testing.T) {
	// Create test storage
	st := NewStorage("")
	defer func() { _ = st.Close() }()

	contractAddress := "0x1111111111111111111111111111111111111111"
	timestamp := common.DayStart(uint64(time.Now().Unix()))

	// Create first set of users
	users1 := storage.MakeDailyContractUsers()
	users1.Add("0x1111111111111111111111111111111111111111")
	users1.Add("0x2222222222222222222222222222222222222222")

	// Store first set
	batch1 := st.NewBatch()
	batch1.AddDailyContractUsers(contractAddress, timestamp, users1)
	batch1.CommitBatch()

	// Create second set of users (should overwrite)
	users2 := storage.MakeDailyContractUsers()
	users2.Add("0x3333333333333333333333333333333333333333")
	users2.Add("0x4444444444444444444444444444444444444444")
	users2.Add("0x5555555555555555555555555555555555555555")

	// Store second set
	batch2 := st.NewBatch()
	batch2.AddDailyContractUsers(contractAddress, timestamp, users2)
	batch2.CommitBatch()

	// Retrieve and verify - should only have the second set
	retrieved := st.GetDailyContractUsers(contractAddress, timestamp)
	assert.Len(t, retrieved.Users, 3)
	assert.Contains(t, retrieved.Users, "0x3333333333333333333333333333333333333333")
	assert.Contains(t, retrieved.Users, "0x4444444444444444444444444444444444444444")
	assert.Contains(t, retrieved.Users, "0x5555555555555555555555555555555555555555")

	// Should not contain first set
	assert.NotContains(t, retrieved.Users, "0x1111111111111111111111111111111111111111")
	assert.NotContains(t, retrieved.Users, "0x2222222222222222222222222222222222222222")
}

func TestStorage_GetDailyContractUsers_NotFound(t *testing.T) {
	// Create test storage
	st := NewStorage("")
	defer func() { _ = st.Close() }()

	contractAddress := "0x1111111111111111111111111111111111111111"
	timestamp := common.DayStart(uint64(time.Now().Unix()))

	// Try to retrieve data that doesn't exist
	retrieved := st.GetDailyContractUsers(contractAddress, timestamp)

	// Should return empty list, not error
	assert.Len(t, retrieved.Users, 0)
}
