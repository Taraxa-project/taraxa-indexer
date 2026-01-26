package rewards

import (
	"math/big"
	"testing"
	"time"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage/pebble"
	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/stretchr/testify/assert"
)

func TestYieldSavingStruct(t *testing.T) {
	ys := storage.YieldSaving{
		Time:   1640995200, // 2022-01-01 00:00:00 UTC
		Period: 1000,
	}

	assert.Equal(t, uint64(1640995200), ys.Time)
	assert.Equal(t, uint64(1000), ys.Period)
}

func TestRewardsWithYieldSaving(t *testing.T) {
	st := pebble.NewStorage("")
	defer func() { _ = st.Close() }()

	config := makeTestConfig()
	config.Chain.BlocksPerYear = big.NewInt(100)

	// Create a block with timestamp
	blockTime := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	block := common.Block{
		Pbft: models.Pbft{
			Number:    1000,
			Author:    "0x1",
			Timestamp: uint64(blockTime.Unix()),
		},
	}

	bd := &chain.BlockData{
		Pbft:                 &block,
		TotalAmountDelegated: big.NewInt(1000000),
		TotalSupply:          big.NewInt(1000000),
		Validators:           []common.Validator{},
	}

	prevYieldsSaving := storage.YieldSaving{
		Time:   uint64(blockTime.Add(-7 * 24 * time.Hour).Unix()), // 1 week ago
		Period: 900,
	}
	pys := prevYieldsSaving

	r := MakeRewards(st, st.NewBatch(), config.Chain, bd, &pys)

	assert.Equal(t, uint64(1000), r.blockNum)
	assert.Equal(t, prevYieldsSaving, pys)
}

func TestAfterCommit_NoRewardsDistribution(t *testing.T) {
	st := pebble.NewStorage("")
	defer func() { _ = st.Close() }()

	config := makeTestConfig()
	// Set distribution frequency to 100, so block 50 won't trigger distribution
	config.Chain.Hardforks.RewardsDistributionFrequency = map[uint64]uint32{0: 100}
	config.Chain.Hardforks.AspenHf.BlockNumPartOne = 0

	blockTime := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	block := common.Block{
		Pbft: models.Pbft{
			Number:    50, // Not divisible by distribution frequency
			Author:    "0x1",
			Timestamp: uint64(blockTime.Unix()),
		},
	}

	bd := &chain.BlockData{
		Pbft:                 &block,
		TotalAmountDelegated: big.NewInt(1000000),
		TotalSupply:          big.NewInt(1000000),
		Validators:           []common.Validator{},
	}

	prevYieldsSaving := storage.YieldSaving{
		Time:   uint64(blockTime.Add(-7 * 24 * time.Hour).Unix()),
		Period: 0,
	}

	pys := &prevYieldsSaving
	r := MakeRewards(st, st.NewBatch(), config.Chain, bd, pys)

	r.AfterCommit()
	assert.Equal(t, *pys, *r.prevYieldsSaving) // No yield saving should occur
}

func TestAfterCommit_WithRewardsDistributionAndWeekChange(t *testing.T) {
	st := pebble.NewStorage("")
	defer func() { _ = st.Close() }()

	config := makeTestConfig()
	config.Chain.Hardforks.AspenHf.BlockNumPartOne = 0

	// Current time - week 2 of 2022
	currentTime := time.Date(2022, 1, 10, 0, 0, 0, 0, time.UTC) // Monday of week 2
	// Previous time - week 1 of 2022
	prevTime := time.Date(2022, 1, 2, 0, 0, 0, 0, time.UTC) // Monday of week 1

	block := common.Block{
		Pbft: models.Pbft{
			Number:    100, // Divisible by distribution frequency (100)
			Author:    "0x1",
			Timestamp: uint64(currentTime.Unix()),
		},
	}

	bd := &chain.BlockData{
		Pbft:                 &block,
		TotalAmountDelegated: big.NewInt(1000000),
		TotalSupply:          big.NewInt(1000000),
		Validators:           []common.Validator{},
	}

	prevYieldsSaving := storage.YieldSaving{
		Time:   uint64(prevTime.Unix()),
		Period: 50,
	}

	// Add some multiplied yield data for the interval
	batch := st.NewBatch()
	for i := uint64(51); i <= 100; i++ {
		multipliedYield := storage.MultipliedYield{Yield: big.NewInt(1000)}
		batch.AddSingleKey(&multipliedYield, storage.FormatIntToKey(i))
	}
	batch.CommitBatch()

	pys := prevYieldsSaving
	r := MakeRewards(st, st.NewBatch(), config.Chain, bd, &pys)

	r.AfterCommit()
	assert.NotEqual(t, *r.prevYieldsSaving, prevYieldsSaving)
}

func TestAfterCommit_SameWeek(t *testing.T) {
	st := pebble.NewStorage("")
	defer func() { _ = st.Close() }()

	config := makeTestConfig()
	config.Chain.Hardforks.AspenHf.BlockNumPartOne = 0

	// Both times in the same week
	currentTime := time.Date(2022, 1, 10, 0, 0, 0, 0, time.UTC) // Monday of week 2
	prevTime := time.Date(2022, 1, 11, 0, 0, 0, 0, time.UTC)    // Tuesday of week 2

	block := common.Block{
		Pbft: models.Pbft{
			Number:    100, // Divisible by distribution frequency
			Author:    "0x1",
			Timestamp: uint64(currentTime.Unix()),
		},
	}

	bd := &chain.BlockData{
		Pbft:                 &block,
		TotalAmountDelegated: big.NewInt(1000000),
		TotalSupply:          big.NewInt(1000000),
		Validators:           []common.Validator{},
	}

	prevYieldsSaving := storage.YieldSaving{
		Time:   uint64(prevTime.Unix()),
		Period: 50,
	}

	pys := &prevYieldsSaving
	r := MakeRewards(st, st.NewBatch(), config.Chain, bd, pys)

	r.AfterCommit()
	assert.Equal(t, *pys, *r.prevYieldsSaving) // No yield saving should occur (same week)
}

func TestProcessIntervalYield_WithCustomInterval(t *testing.T) {
	st := pebble.NewStorage("")
	defer func() { _ = st.Close() }()

	config := makeTestConfig()
	config.Chain.BlocksPerYear = big.NewInt(100)

	// Add multiplied yield data
	batch := st.NewBatch()
	intervalStart := uint64(50)
	currentBlock := uint64(100)

	expectedSum := big.NewInt(0)
	for i := intervalStart; i < currentBlock; i++ {
		multipliedYield := storage.MultipliedYield{Yield: big.NewInt(int64(i * 10))}
		batch.AddSingleKey(&multipliedYield, storage.FormatIntToKey(i))
		expectedSum.Add(expectedSum, big.NewInt(int64(i*10)))
	}
	batch.CommitBatch()

	block := common.Block{
		Pbft: models.Pbft{
			Number:    currentBlock,
			Author:    "0x1",
			Timestamp: uint64(time.Now().Unix()),
		},
	}

	bd := &chain.BlockData{
		Pbft:                 &block,
		TotalAmountDelegated: big.NewInt(1000000),
		TotalSupply:          big.NewInt(1000000),
		Validators:           []common.Validator{},
	}

	prevYieldsSaving := storage.YieldSaving{
		Time:   uint64(time.Now().Add(-time.Hour).Unix()),
		Period: intervalStart,
	}

	pys := &prevYieldsSaving
	r := MakeRewards(st, st.NewBatch(), config.Chain, bd, pys)

	// Test the interval yield processing
	testBatch := st.NewBatch()
	r.processIntervalYield(intervalStart, testBatch)
	testBatch.CommitBatch()

	// Verify that total yield was saved
	totalYield := st.GetTotalYield(currentBlock)
	assert.NotEmpty(t, totalYield.Yield)

	// Verify that multiplied yield data was removed
	count := 0
	storage.ProcessIntervalData(st, intervalStart, func(key []byte, o storage.MultipliedYield) (stop bool) {
		count++
		return false
	})
	assert.Equal(t, 0, count)
}

func TestProcessValidatorsIntervalYield_WithCustomInterval(t *testing.T) {
	st := pebble.NewStorage("")
	defer func() { _ = st.Close() }()

	config := makeTestConfig()
	config.Chain.BlocksPerYear = big.NewInt(100)

	// Add validators yield data
	batch := st.NewBatch()
	intervalStart := uint64(50)
	currentBlock := uint64(100)

	validators := []string{"validator1", "validator2", "validator3"}

	for i := intervalStart; i < currentBlock; i++ {
		var yields []storage.ValidatorYield
		for _, validator := range validators {
			yields = append(yields, storage.ValidatorYield{
				Validator: validator,
				Yield:     big.NewInt(int64(i * 5)),
			})
		}
		validatorsYield := storage.ValidatorsYield{Yields: yields}
		batch.AddSingleKey(&validatorsYield, storage.FormatIntToKey(i))
	}
	batch.CommitBatch()

	block := common.Block{
		Pbft: models.Pbft{
			Number:    currentBlock,
			Author:    "0x1",
			Timestamp: uint64(time.Now().Unix()),
		},
	}

	bd := &chain.BlockData{
		Pbft:                 &block,
		TotalAmountDelegated: big.NewInt(1000000),
		TotalSupply:          big.NewInt(1000000),
		Validators:           []common.Validator{},
	}

	prevYieldsSaving := storage.YieldSaving{
		Time:   uint64(time.Now().Add(-time.Hour).Unix()),
		Period: intervalStart,
	}

	pys := &prevYieldsSaving
	r := MakeRewards(st, st.NewBatch(), config.Chain, bd, pys)

	// Test the validators interval yield processing
	testBatch := st.NewBatch()
	r.processValidatorsIntervalYield(intervalStart, testBatch)
	testBatch.CommitBatch()

	// Verify that validator yields were saved
	for _, validator := range validators {
		validatorYield := st.GetValidatorYield(validator, currentBlock)
		assert.NotEmpty(t, validatorYield.Yield)
	}

	// Verify that validators yield data was removed
	count := 0
	storage.ProcessIntervalData(st, intervalStart, func(key []byte, o storage.ValidatorsYield) (stop bool) {
		count++
		return false
	})
	assert.Equal(t, 0, count)
}
