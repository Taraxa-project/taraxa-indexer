package chain

import (
	"math/big"

	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/spiretechnology/go-pool"
)

type BlockData struct {
	Pbft                 *common.Block
	Dags                 []common.DagBlock
	Transactions         []common.Transaction
	Traces               []common.TransactionTrace
	Votes                common.VotesResponse
	Validators           []common.Validator
	TotalAmountDelegated *big.Int
	TotalSupply          *big.Int
}

func MakeEmptyBlockData() *BlockData {
	bd := new(BlockData)
	bd.Pbft = new(common.Block)
	bd.Dags = make([]common.DagBlock, 0)
	bd.Transactions = make([]common.Transaction, 0)
	bd.Traces = make([]common.TransactionTrace, 0)
	bd.Votes = common.VotesResponse{Votes: make([]common.Vote, 0)}
	bd.Validators = make([]common.Validator, 0)
	return bd
}

// Move common parts to the function, so we won't need change this it in two places
func scheduleBlockDataTasks(tp pool.Pool, c common.Client, period uint64, bd *BlockData, err *error) {
	tp.Go(common.MakeTaskWithResult(c.GetPeriodDagBlocks, period, &bd.Dags, err).Run)
	tp.Go(common.MakeTaskWithResult(c.GetPeriodTransactions, period, &bd.Transactions, err).Run)
	tp.Go(common.MakeTaskWithResult(c.TraceBlockTransactions, period, &bd.Traces, err).Run)
	tp.Go(common.MakeTaskWithResult(c.GetPreviousBlockCertVotes, period, &bd.Votes, err).Run)
	tp.Go(common.MakeTaskWithResult(c.GetValidatorsAtBlock, period, &bd.Validators, err).Run)
	tp.Go(common.MakeTaskWithResult(c.GetTotalAmountDelegated, period, &bd.TotalAmountDelegated, err).Run)
	supplyPeriod := period
	if period >= 100 {
		supplyPeriod = period - 100
	}
	tp.Go(common.MakeTaskWithResult(c.GetTotalSupply, supplyPeriod, &bd.TotalSupply, err).Run)
}

func GetBlockData(c common.Client, period uint64) (bd *BlockData, err error) {
	bd = MakeEmptyBlockData()
	bd.Pbft.Number = period
	tp := common.MakeThreadPool()
	tp.Go(common.MakeTaskWithResult(c.GetBlockByNumber, period, &bd.Pbft, &err).Run)
	scheduleBlockDataTasks(tp, c, period, bd, &err)

	tp.Wait()

	if err != nil {
		return nil, err
	}
	return
}

func GetBlockDataFromPbft(c common.Client, pbft *common.Block) (bd *BlockData, err error) {
	bd = MakeEmptyBlockData()
	bd.Pbft = pbft

	tp := common.MakeThreadPool()
	scheduleBlockDataTasks(tp, c, pbft.Number, bd, &err)

	tp.Wait()
	if err != nil {
		return nil, err
	}
	return
}
