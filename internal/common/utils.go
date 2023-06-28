package common

import (
	"math/big"
	"runtime"
	"strconv"

	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/spiretechnology/go-pool"
)

// isn't creating threads, but limiting goroutines count. Mostly used for RPC and db related tasks
func MakeThreadPool() pool.Pool {
	return pool.New(uint(runtime.NumCPU()))
}

func ParseStringToBigInt(v string) *big.Int {
	a := big.NewInt(0)
	a.SetString(v, 0)
	return a
}

func FormatFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', 4, 64)
}

func GetYieldIntervalEnd(storage storage.Storage, block_num *uint64, interval uint64) uint64 {
	block := uint64(0)
	if block_num == nil {
		block = storage.GetFinalizationData().PbftCount
	} else {
		block = *block_num
	}

	if block%interval == 0 {
		return block
	}
	block = block - block%interval + interval
	return block
}
