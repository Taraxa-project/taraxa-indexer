package storage

import (
	"testing"
	"time"

	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/stretchr/testify/assert"
)

func TestDailyContractUsers_Add(t *testing.T) {
	dcu := MakeDailyContractUsersFromList(DailyContractUsersList{Users: []string{}})

	// Test adding first user
	dcu.Add("0x1111111111111111111111111111111111111111")
	users := dcu.GetList().Users
	assert.Len(t, users, 1)
	assert.Contains(t, users, "0x1111111111111111111111111111111111111111")

	// Test adding duplicate user (should not increase count)
	dcu.Add("0x1111111111111111111111111111111111111111")
	users = dcu.GetList().Users
	assert.Len(t, users, 1)

	// Test adding second user
	dcu.Add("0x2222222222222222222222222222222222222222")
	users = dcu.GetList().Users
	assert.Len(t, users, 2)
	assert.Contains(t, users, "0x1111111111111111111111111111111111111111")
	assert.Contains(t, users, "0x2222222222222222222222222222222222222222")
}

func TestDailyContractUsers_MakeFromList(t *testing.T) {
	usersList := DailyContractUsersList{
		Users: []string{
			"0x1111111111111111111111111111111111111111",
			"0x2222222222222222222222222222222222222222",
			"0x1111111111111111111111111111111111111111", // duplicate
		},
	}

	dcu := MakeDailyContractUsersFromList(usersList)
	users := dcu.GetList().Users

	// Should only have 2 unique users
	assert.Len(t, users, 2)
	assert.Contains(t, users, "0x1111111111111111111111111111111111111111")
	assert.Contains(t, users, "0x2222222222222222222222222222222222222222")
}

func TestDayStart_Utility(t *testing.T) {
	// Test day start calculation (should align with our time calculations)
	now := uint64(time.Now().Unix())
	dayStart := common.DayStart(now)

	// Day start should be less than or equal to now
	assert.LessOrEqual(t, dayStart, now)

	// Should be aligned to day boundary (divisible by 24*60*60)
	assert.Equal(t, uint64(0), dayStart%common.Day)

	// Test specific case
	// January 1, 2024 15:30:45 UTC
	specificTime := uint64(1704120645)     // 2024-01-01T15:30:45Z
	expectedDayStart := uint64(1704067200) // 2024-01-01T00:00:00Z

	actualDayStart := common.DayStart(specificTime)
	assert.Equal(t, expectedDayStart, actualDayStart)
}

func TestDays30_Constant(t *testing.T) {
	// Verify that Days30 constant is correct
	assert.Equal(t, 30*24*60*60, common.Days30)
	assert.Equal(t, 2592000, common.Days30) // 30 days in seconds
}
