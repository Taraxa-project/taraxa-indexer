package chain

import (
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
)

type BlockData struct {
	Pbft         *Block
	Dags         []DagBlock
	Transactions []Transaction
	Traces       []TransactionTrace
	Votes        VotesResponse
	Validators   []Validator
}

func MakeEmptyBlockData() *BlockData {
	bd := new(BlockData)
	bd.Pbft = new(Block)
	bd.Dags = make([]DagBlock, 0)
	bd.Transactions = make([]Transaction, 0)
	bd.Traces = make([]TransactionTrace, 0)
	bd.Votes = VotesResponse{Votes: make([]Vote, 0)}
	bd.Validators = make([]Validator, 0)
	return bd
}

func GetBlockData(c Client, period uint64) (bd *BlockData, err error) {
	bd = MakeEmptyBlockData()
	tp := common.MakeThreadPool()
	tp.Go(common.MakeTaskWithResult(c.GetBlockByNumber, period, bd.Pbft, &err).Run)
	tp.Go(common.MakeTaskWithResult(c.GetPeriodDagBlocks, period, &bd.Dags, &err).Run)
	tp.Go(common.MakeTaskWithResult(c.GetPeriodTransactions, period, &bd.Transactions, &err).Run)
	tp.Go(common.MakeTaskWithResult(c.TraceBlockTransactions, period, &bd.Traces, &err).Run)
	tp.Go(common.MakeTaskWithResult(c.GetPreviousBlockCertVotes, period, &bd.Votes, &err).Run)
	tp.Go(common.MakeTaskWithResult(c.GetValidatorsAtBlock, period, &bd.Validators, &err).Run)
	tp.Wait()
	if err != nil {
		return nil, err
	}
	return
}

func GetBlockDataFromPbft(c Client, pbft *Block) (bd *BlockData, err error) {
	bd = MakeEmptyBlockData()
	bd.Pbft = pbft

	tp := common.MakeThreadPool()
	tp.Go(common.MakeTaskWithResult(c.GetPeriodDagBlocks, pbft.Number, &bd.Dags, &err).Run)
	tp.Go(common.MakeTaskWithResult(c.GetPeriodTransactions, pbft.Number, &bd.Transactions, &err).Run)
	tp.Go(common.MakeTaskWithResult(c.TraceBlockTransactions, pbft.Number, &bd.Traces, &err).Run)
	tp.Go(common.MakeTaskWithResult(c.GetPreviousBlockCertVotes, pbft.Number, &bd.Votes, &err).Run)
	tp.Go(common.MakeTaskWithResult(c.GetValidatorsAtBlock, pbft.Number, &bd.Validators, &err).Run)
	tp.Wait()
	if err != nil {
		return nil, err
	}
	return
}
