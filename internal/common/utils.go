package common

import (
	"math/big"
	"runtime"

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
