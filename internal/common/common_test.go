package common

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestIntervalEnd(t *testing.T) {
	interval := uint64(100)

	for _, block_num := range []uint64{100, 300, 1500} {
		end := GetYieldIntervalEnd(10000, &block_num, interval)
		assert.Equal(t, block_num, end)
	}

	for _, block_num := range []uint64{101, 150, 199} {
		end := GetYieldIntervalEnd(10000, &block_num, interval)
		assert.Equal(t, uint64(200), end)
	}

	for _, block_num := range []uint64{1001, 1050, 1099} {
		end := GetYieldIntervalEnd(10000, &block_num, interval)
		assert.Equal(t, uint64(1100), end)
	}
	{
		end := GetYieldIntervalEnd(150, nil, interval)
		assert.Equal(t, uint64(100), end)
	}
	{
		end := GetYieldIntervalEnd(50, nil, interval)
		assert.Equal(t, uint64(100), end)
	}
}

func TestDistributionFrequency(t *testing.T) {
	config := HardforksConfig{RewardsDistributionFrequency: map[uint64]uint32{100: 10, 1000: 20, 2000: 30}}
	for period, frequency := range map[uint64]uint32{50: 1, 99: 1, 100: 10, 101: 10, 1000: 20, 2000: 30, 10001: 30} {
		df := config.GetDistributionFrequency(period)
		if frequency != df {
			t.Errorf("Expected frequency %d for period %d, but got %d", frequency, period, df)
		}
	}

}
