package chain

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/ethereum/go-ethereum/rpc"
)

type ClientMock struct {
	Blocks            map[uint64]*common.Block
	Traces            map[string][]common.TransactionTrace
	Transactions      map[string]common.Transaction
	BlockTransactions map[uint64][]string
	EventLogs         map[string][]common.EventLog
}

var ErrNotImplemented = fmt.Errorf("not implemented")

func MakeMockClient() *ClientMock {
	m := new(ClientMock)
	m.Traces = make(map[string][]common.TransactionTrace)
	m.Transactions = make(map[string]common.Transaction)
	m.BlockTransactions = make(map[uint64][]string)
	m.EventLogs = make(map[string][]common.EventLog)
	m.Blocks = make(map[uint64]*common.Block)
	return m
}

func (c *ClientMock) GetBalanceAtBlock(address string, blockNumber uint64) (balance string, err error) {
	return "", ErrNotImplemented
}

func (c *ClientMock) GetBlocks(start, end uint64) (blks []*common.Block, err error) {
	for i := start; i <= end; i++ {
		blks = append(blks, c.Blocks[i])
	}
	return blks, nil
}

func (c *ClientMock) GetBlockByNumber(number uint64) (blk *common.Block, err error) {
	return c.Blocks[number], nil
}

func (c *ClientMock) GetLatestPeriod() (p uint64, e error) {
	return 0, ErrNotImplemented
}

func (c *ClientMock) TraceBlockTransactions(num uint64) (traces []common.TransactionTrace, err error) {
	hashes := c.BlockTransactions[num]
	for _, h := range hashes {
		traces = append(traces, c.Traces[h]...)
	}
	return
}

func (c *ClientMock) GetTransactionByHash(hash string) (trx common.Transaction, err error) {
	return c.Transactions[hash], nil
}

func (c *ClientMock) GetPeriodTransactions(num uint64) (trxs []common.Transaction, err error) {
	hashes := c.BlockTransactions[num]
	for _, h := range hashes {
		trxs = append(trxs, c.Transactions[h])
	}
	return trxs, nil
}

func (c *ClientMock) GetPbftBlockWithDagBlocks(period uint64) (pbftWithDags common.PbftBlockWithDags, err error) {
	return common.PbftBlockWithDags{}, ErrNotImplemented
}

func (c *ClientMock) GetDagBlockByHash(hash string) (dag common.DagBlock, err error) {
	return common.DagBlock{}, ErrNotImplemented
}

func (c *ClientMock) GetPeriodDagBlocks(period uint64) (dags []common.DagBlock, err error) {
	return []common.DagBlock{}, nil
}
func (c *ClientMock) GetVersion() (version string, err error) {
	return "", nil
}

func (c *ClientMock) GetGenesis() (genesis common.GenesisObject, err error) {
	return common.GenesisObject{}, ErrNotImplemented
}

func (c *ClientMock) GetLogs(fromBlock, toBlock uint64, addresses []string, topics [][]string) (logs []common.EventLog, err error) {
	return nil, ErrNotImplemented
}

func (c *ClientMock) GetChainStats() (ns common.FinalizationData, err error) {
	return common.FinalizationData{}, ErrNotImplemented
}

func (c *ClientMock) GetPreviousBlockCertVotes(period uint64) (vr common.VotesResponse, err error) {
	return common.VotesResponse{}, nil
}

func (c *ClientMock) GetValidatorsAtBlock(uint64) (validators []common.Validator, err error) {
	return []common.Validator{}, nil
}

func (c *ClientMock) GetTotalAmountDelegated(uint64) (totalAmountDelegated *big.Int, err error) {
	return big.NewInt(0), nil
}

func (c *ClientMock) GetTotalSupply(uint64) (totalSupply *big.Int, err error) {
	return big.NewInt(0), nil
}

func (c *ClientMock) SubscribeNewHeads() (chan common.Block, *rpc.ClientSubscription, error) {
	return nil, nil, nil
}

func (c *ClientMock) Close() {
}

func (c *ClientMock) AddTransactionFromJson(trx_json string) {
	var trx common.Transaction
	err := json.Unmarshal([]byte(trx_json), &trx)
	if err != nil {
		fmt.Println("ClientMock.AddTransactionFromJson", err)
	}

	trx.SetTimestamp(1)
	c.BlockTransactions[trx.GetModel().BlockNumber] = append(c.BlockTransactions[trx.GetModel().BlockNumber], trx.Hash)
	c.Transactions[trx.Hash] = trx
}

func (c *ClientMock) AddLogsFromJson(trx_json string) {
	var trx common.Transaction
	err := json.Unmarshal([]byte(trx_json), &trx)
	if err != nil {
		fmt.Println(err)
	}

	c.EventLogs[trx.Hash] = trx.Logs
	c.Transactions[trx.Hash] = trx
}

func (c *ClientMock) AddTracesFromJson(hash, traces_json string) {
	var traces []common.TransactionTrace
	err := json.Unmarshal([]byte(traces_json), &traces)
	if err != nil {
		fmt.Println("ClientMock.AddTracesFromJson", err)
	}

	c.Traces[hash] = traces
}

func (c *ClientMock) AddPbftBlock(period uint64, block *common.Block) {
	c.Blocks[period] = block
}
