package api

import (
	"fmt"
	"testing"
	"time"

	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage/pebble"
	"github.com/stretchr/testify/assert"
)

func TestApiHandler_calculateContract30DayAverage(t *testing.T) {
	// Create test storage
	st := pebble.NewStorage("")
	defer st.Close()

	config := &common.Config{}
	handler := NewApiHandler(st, config, nil)

	contractAddress := "0x1111111111111111111111111111111111111111"
	timestamp := uint64(time.Now().Unix())

	// Test with no data - should return 0 count
	average := handler.calculateContract30DayAverage(contractAddress, timestamp)
	assert.InDelta(t, 0.0, float64(average.Count), 1e-5)
}

func TestApiHandler_calculateContract30DayAverage_WithData(t *testing.T) {
	// Create test storage
	st := pebble.NewStorage("")
	defer st.Close()

	config := &common.Config{}
	handler := NewApiHandler(st, config, nil)

	contractAddress := "0x1111111111111111111111111111111111111111"
	timestamp := uint64(time.Now().Unix())
	endDay := common.DayStart(timestamp)

	// Create test data for multiple days
	batch := st.NewBatch()

	// Add more days with users to get a meaningful average
	totalUsers := 0
	for i := 1; i <= 30; i++ {
		day := endDay - (uint64(i) * common.Day)
		users := storage.MakeDailyContractUsers()

		// Add 10 users per day for a clear average of 10
		for j := 1; j <= 10; j++ {
			users.Add(makeTestAddress(i*100 + j))
		}
		batch.AddDailyContractUsers(contractAddress, day, users)
		totalUsers += 10
	}

	batch.CommitBatch()

	// Calculate average
	average := handler.calculateContract30DayAverage(contractAddress, endDay)

	// Should be (10 users/day * 30 days) / 30 = 10
	assert.InDelta(t, 10.0, float64(average.Count), 1e-5)
}

func TestApiHandler_calculateContract30DayAverage_NoDataDays(t *testing.T) {
	// Create test storage
	st := pebble.NewStorage("")
	defer st.Close()

	config := &common.Config{}
	handler := NewApiHandler(st, config, nil)

	contractAddress := "0x1111111111111111111111111111111111111111"
	timestamp := uint64(time.Now().Unix())
	endDay := common.DayStart(timestamp)

	// Create test data with gaps (days with no users)
	batch := st.NewBatch()

	// Only add data for 2 days out of 30
	day1 := endDay - (10 * common.Day)
	users1 := storage.MakeDailyContractUsers()
	users1.Add("0x1111111111111111111111111111111111111111")
	users1.Add("0x2222222222222222222222222222222222222222")
	batch.AddDailyContractUsers(contractAddress, day1, users1)

	day2 := endDay - (5 * common.Day)
	users2 := storage.MakeDailyContractUsers()
	users2.Add("0x3333333333333333333333333333333333333333")
	users2.Add("0x4444444444444444444444444444444444444444")
	users2.Add("0x5555555555555555555555555555555555555555")
	users2.Add("0x6666666666666666666666666666666666666666")
	batch.AddDailyContractUsers(contractAddress, day2, users2)

	batch.CommitBatch()

	// Calculate average
	average := handler.calculateContract30DayAverage(contractAddress, endDay)

	// Should be (2 + 4) / 30 = 0.2
	assert.InDelta(t, 0.2, float64(average.Count), 1e-5)
}

func TestApiHandler_calculateContract30DayAverage_EmptyDays(t *testing.T) {
	// Create test storage
	st := pebble.NewStorage("")
	defer st.Close()

	config := &common.Config{}
	handler := NewApiHandler(st, config, nil)

	contractAddress := "0x1111111111111111111111111111111111111111"
	timestamp := uint64(time.Now().Unix())
	endDay := common.DayStart(timestamp)

	// Create test data with some days having empty user lists
	batch := st.NewBatch()

	// Day with users
	day1 := endDay - (5 * common.Day)
	users1 := storage.MakeDailyContractUsers()
	users1.Add("0x1111111111111111111111111111111111111111")
	users1.Add("0x2222222222222222222222222222222222222222")
	batch.AddDailyContractUsers(contractAddress, day1, users1)

	// Day with empty users list (should be ignored in calculation)
	day2 := endDay - (3 * common.Day)
	batch.AddDailyContractUsers(contractAddress, day2, storage.MakeDailyContractUsers())

	batch.CommitBatch()

	// Calculate average
	average := handler.calculateContract30DayAverage(contractAddress, endDay)

	// Should be 2 / 30 = 0.0666667
	assert.InDelta(t, 2.0/30.0, float64(average.Count), 1e-5)
}

func TestApiHandler_calculateContract30DayAverage_MixedData(t *testing.T) {
	// Create test storage
	st := pebble.NewStorage("")
	defer st.Close()

	config := &common.Config{}
	handler := NewApiHandler(st, config, nil)

	contractAddress := "0x1111111111111111111111111111111111111111"
	timestamp := uint64(time.Now().Unix())
	endDay := common.DayStart(timestamp)

	// Create test data with mixed user counts - some days high, some low, some zero
	batch := st.NewBatch()

	totalUsers := 0
	// Week 1: High activity (days 1-7)
	for i := 1; i <= 7; i++ {
		day := endDay - (uint64(i) * common.Day)
		users := storage.MakeDailyContractUsers()

		// Add 20 users per day
		for j := 1; j <= 20; j++ {
			users.Add(makeTestAddress(i*1000 + j))
		}
		batch.AddDailyContractUsers(contractAddress, day, users)
		totalUsers += 20
	}

	// Week 2: Medium activity (days 8-14)
	for i := 8; i <= 14; i++ {
		day := endDay - (uint64(i) * common.Day)
		users := storage.MakeDailyContractUsers()

		// Add 5 users per day
		for j := 1; j <= 5; j++ {
			users.Add(makeTestAddress(i*1000 + j))
		}
		batch.AddDailyContractUsers(contractAddress, day, users)
		totalUsers += 5
	}

	// Week 3: Low activity (days 15-21)
	for i := 15; i <= 21; i++ {
		day := endDay - (uint64(i) * common.Day)
		users := storage.MakeDailyContractUsers()

		// Add 1 user per day
		users.Add(makeTestAddress(i*1000 + 1))
		batch.AddDailyContractUsers(contractAddress, day, users)
		totalUsers += 1
	}

	// Week 4: No activity (days 22-28) - don't add any data
	// Days 29-30: Some activity
	for i := 29; i <= 30; i++ {
		day := endDay - (uint64(i) * common.Day)
		users := storage.MakeDailyContractUsers()

		// Add 3 users per day
		for j := 1; j <= 3; j++ {
			users.Add(makeTestAddress(i*1000 + j))
		}
		batch.AddDailyContractUsers(contractAddress, day, users)
		totalUsers += 3
	}

	batch.CommitBatch()

	// Calculate average
	average := handler.calculateContract30DayAverage(contractAddress, endDay)

	// Should be totalUsers / 30 days
	// totalUsers = (7*20) + (7*5) + (7*1) + (0*7) + (2*3) = 140 + 35 + 7 + 0 + 6 = 188
	expectedAverage := float64(188) / 30.0 // 188/30 = 6.2666667
	assert.InDelta(t, expectedAverage, float64(average.Count), 1e-5)
}

func TestApiHandler_calculateContract30DayAverage_RecentActivity(t *testing.T) {
	// Test scenario where most activity is recent
	st := pebble.NewStorage("")
	defer st.Close()

	config := &common.Config{}
	handler := NewApiHandler(st, config, nil)

	contractAddress := "0x2222222222222222222222222222222222222222"
	timestamp := uint64(time.Now().Unix())
	endDay := common.DayStart(timestamp)

	batch := st.NewBatch()

	totalUsers := 0
	// Only last 5 days have activity
	for i := 1; i <= 5; i++ {
		day := endDay - (uint64(i) * common.Day)
		users := storage.MakeDailyContractUsers()

		// Add 30 users per day for recent days
		for j := 1; j <= 30; j++ {
			users.Add(makeTestAddress(i*2000 + j))
		}
		batch.AddDailyContractUsers(contractAddress, day, users)
		totalUsers += 30
	}

	batch.CommitBatch()

	// Calculate average
	average := handler.calculateContract30DayAverage(contractAddress, endDay)

	// Should be (5*30) / 30 = 150/30 = 5
	assert.InDelta(t, 5.0, float64(average.Count), 1e-5)
}

func TestApiHandler_calculateContract30DayAverage_OldActivity(t *testing.T) {
	// Test scenario where most activity is old
	st := pebble.NewStorage("")
	defer st.Close()

	config := &common.Config{}
	handler := NewApiHandler(st, config, nil)

	contractAddress := "0x3333333333333333333333333333333333333333"
	timestamp := uint64(time.Now().Unix())
	endDay := common.DayStart(timestamp)

	batch := st.NewBatch()

	totalUsers := 0
	// Only days 25-30 (oldest in our 30-day window) have activity
	for i := 25; i <= 30; i++ {
		day := endDay - (uint64(i) * common.Day)
		users := storage.MakeDailyContractUsers()

		// Add 15 users per day for old days
		for j := 1; j <= 15; j++ {
			users.Add(makeTestAddress(i*3000 + j))
		}
		batch.AddDailyContractUsers(contractAddress, day, users)
		totalUsers += 15
	}

	batch.CommitBatch()

	// Calculate average
	average := handler.calculateContract30DayAverage(contractAddress, endDay)

	// Should be (6*15) / 30 = 90/30 = 3
	assert.InDelta(t, 3.0, float64(average.Count), 1e-5)
}

func TestApiHandler_calculateContract30DayAverage_SingleDayHigh(t *testing.T) {
	// Test scenario with one day of very high activity
	st := pebble.NewStorage("")
	defer st.Close()

	config := &common.Config{}
	handler := NewApiHandler(st, config, nil)

	contractAddress := "0x4444444444444444444444444444444444444444"
	timestamp := uint64(time.Now().Unix())
	endDay := common.DayStart(timestamp)

	batch := st.NewBatch()

	// Only one day has activity, but very high
	day := endDay - (15 * common.Day) // Middle of the 30-day period
	users := storage.MakeDailyContractUsers()

	// Add 300 users in one day
	for j := 1; j <= 300; j++ {
		users.Add(makeTestAddress(4000 + j))
	}
	batch.AddDailyContractUsers(contractAddress, day, users)

	batch.CommitBatch()

	// Calculate average
	average := handler.calculateContract30DayAverage(contractAddress, endDay)

	// Should be 300 / 30 = 10
	assert.InDelta(t, 10.0, float64(average.Count), 1e-5)
}

func TestApiHandler_calculateContract30DayAverage_GradualIncrease(t *testing.T) {
	// Test scenario with gradually increasing activity
	st := pebble.NewStorage("")
	defer st.Close()

	config := &common.Config{}
	handler := NewApiHandler(st, config, nil)

	contractAddress := "0x5555555555555555555555555555555555555555"
	timestamp := uint64(time.Now().Unix())
	endDay := common.DayStart(timestamp)

	batch := st.NewBatch()

	totalUsers := 0
	// Activity increases each day
	for i := 1; i <= 30; i++ {
		day := endDay - (uint64(i) * common.Day)
		users := storage.MakeDailyContractUsers()

		// Add 'i' users on day 'i' (day 1 = 1 user, day 30 = 30 users)
		userCount := 31 - i // So most recent day has most users
		for j := 1; j <= userCount; j++ {
			users.Add(makeTestAddress(i*5000 + j))
		}
		batch.AddDailyContractUsers(contractAddress, day, users)
		totalUsers += userCount
	}

	batch.CommitBatch()

	// Calculate average
	average := handler.calculateContract30DayAverage(contractAddress, endDay)

	// totalUsers = 1+2+3+...+30 = 30*31/2 = 465
	// Average = 465/30 = 15.5
	assert.InDelta(t, 15.5, float64(average.Count), 1e-5)
}

func TestApiHandler_calculateContract30DayAverage_MultipleContracts(t *testing.T) {
	// Test with multiple contracts to ensure isolation
	st := pebble.NewStorage("")
	defer st.Close()

	config := &common.Config{}
	handler := NewApiHandler(st, config, nil)

	contract1 := "0x6666666666666666666666666666666666666666"
	contract2 := "0x7777777777777777777777777777777777777777"
	timestamp := uint64(time.Now().Unix())
	endDay := common.DayStart(timestamp)

	batch := st.NewBatch()

	// Contract 1: 5 users every day
	for i := 1; i <= 30; i++ {
		day := endDay - (uint64(i) * common.Day)
		users := storage.MakeDailyContractUsers()

		for j := 1; j <= 5; j++ {
			users.Add(makeTestAddress(i*6000 + j))
		}
		batch.AddDailyContractUsers(contract1, day, users)
	}

	// Contract 2: 10 users every day
	for i := 1; i <= 30; i++ {
		day := endDay - (uint64(i) * common.Day)
		users := storage.MakeDailyContractUsers()

		for j := 1; j <= 10; j++ {
			users.Add(makeTestAddress(i*7000 + j))
		}
		batch.AddDailyContractUsers(contract2, day, users)
	}

	batch.CommitBatch()

	// Calculate averages for both contracts
	avg1 := handler.calculateContract30DayAverage(contract1, endDay)
	avg2 := handler.calculateContract30DayAverage(contract2, endDay)

	// Contract 1: (5*30)/30 = 5
	assert.InDelta(t, 5.0, float64(avg1.Count), 1e-5)

	// Contract 2: (10*30)/30 = 10
	assert.InDelta(t, 10.0, float64(avg2.Count), 1e-5)
}

// Helper function to create test addresses
func makeTestAddress(i int) string {
	return "0x" + fmt.Sprintf("%040d", i)
}

func TestDailyContractUsers_Integration(t *testing.T) {
	// Test the integration between DailyContractUsers and storage
	st := pebble.NewStorage("")
	defer st.Close()

	contractAddress := "0x1111111111111111111111111111111111111111"
	timestamp := common.DayStart(uint64(time.Now().Unix()))

	// Create and populate DailyContractUsers
	users := storage.MakeDailyContractUsers()
	users.Add("0x1111111111111111111111111111111111111111")
	users.Add("0x2222222222222222222222222222222222222222")
	users.Add("0x1111111111111111111111111111111111111111") // duplicate

	// Store in database
	batch := st.NewBatch()
	batch.AddDailyContractUsers(contractAddress, timestamp, users)
	batch.CommitBatch()

	// Retrieve from database
	retrieved := st.GetDailyContractUsers(contractAddress, timestamp)

	// Should have 2 unique users
	assert.Len(t, retrieved.Users, 2)
	assert.Contains(t, retrieved.Users, "0x1111111111111111111111111111111111111111")
	assert.Contains(t, retrieved.Users, "0x2222222222222222222222222222222222222222")
}
