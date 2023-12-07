package chain

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
)

type EthClient interface {
	CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error)
	// Add other methods here as needed
}

type MockEthClient struct {
	CallContractFunc func(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error)
}
