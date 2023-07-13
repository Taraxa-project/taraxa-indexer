package common

import (
	"fmt"
	"math/big"
	"runtime"
	"strconv"

	"github.com/spiretechnology/go-pool"
)

const DposContractAddress = "0x00000000000000000000000000000000000000fe"

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

func GetYieldIntervalEnd(pbft_count uint64, block_num *uint64, interval uint64) uint64 {
	block := uint64(0)
	if block_num == nil {
		block = pbft_count
	} else {
		block = *block_num
	}

	if block%interval == 0 {
		return block
	}
	block = block - block%interval + interval
	return block
}

func ParseToStringSlice(data []interface{}) ([]string, error) {
	result := make([]string, len(data))

	for i, item := range data {
		switch val := item.(type) {
		case string:
			result[i] = val
		case int:
			result[i] = strconv.Itoa(val)
		case int64:
			result[i] = strconv.FormatInt(val, 10)
		case float64:
			result[i] = strconv.FormatFloat(val, 'f', -1, 64)
		case *big.Int:
			result[i] = val.String()
		default:
			return nil, fmt.Errorf("failed to convert element at index %d to string", i)
		}
	}

	return result, nil
}
