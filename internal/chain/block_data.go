package chain

import (
	"math/big"
)

type BlockData struct {
	Block        *Block
	Dags         []DagBlock
	Transactions []Transaction
	Traces       []TransactionTrace
	Votes        VotesResponse
	Validators   []Validator
	BlockFee     *big.Int
}

func MakeBlockData() *BlockData {
	return &BlockData{BlockFee: big.NewInt(0)}
}
