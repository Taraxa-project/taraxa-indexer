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
}
