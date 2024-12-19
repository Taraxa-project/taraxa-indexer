package chain

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/ethereum/go-ethereum/rpc"
)

type ClientMock struct {
	Blocks            map[uint64]*Block
	Traces            map[string][]TransactionTrace
	Transactions      map[string]Transaction
	BlockTransactions map[uint64][]string
	EventLogs         map[string][]EventLog
}

var ErrNotImplemented = fmt.Errorf("Not implemented")

func MakeMockClient() *ClientMock {
	m := new(ClientMock)
	m.Traces = make(map[string][]TransactionTrace)
	m.Transactions = make(map[string]Transaction)
	m.BlockTransactions = make(map[uint64][]string)
	m.EventLogs = make(map[string][]EventLog)
	m.Blocks = make(map[uint64]*Block)
	return m
}

func (c *ClientMock) GetBalanceAtBlock(address string, blockNumber uint64) (balance string, err error) {
	return "", ErrNotImplemented
}

func (c *ClientMock) GetBlockByNumber(number uint64) (blk *Block, err error) {
	return c.Blocks[number], nil
}

func (c *ClientMock) GetLatestPeriod() (p uint64, e error) {
	return 0, ErrNotImplemented
}

func (c *ClientMock) TraceBlockTransactions(num uint64) (traces []TransactionTrace, err error) {
	hashes := c.BlockTransactions[num]
	for _, h := range hashes {
		traces = append(traces, c.Traces[h]...)
	}
	return
}

func (c *ClientMock) GetTransactionByHash(hash string) (trx Transaction, err error) {
	return c.Transactions[hash], nil
}

func (c *ClientMock) GetPeriodTransactions(num uint64) (trxs []Transaction, err error) {
	hashes := c.BlockTransactions[num]
	for _, h := range hashes {
		trxs = append(trxs, c.Transactions[h])
	}
	return trxs, nil
}

func (c *ClientMock) GetPbftBlockWithDagBlocks(period uint64) (pbftWithDags PbftBlockWithDags, err error) {
	return PbftBlockWithDags{}, ErrNotImplemented
}

func (c *ClientMock) GetDagBlockByHash(hash string) (dag DagBlock, err error) {
	return DagBlock{}, ErrNotImplemented
}

func (c *ClientMock) GetPeriodDagBlocks(period uint64) (dags []DagBlock, err error) {
	return []DagBlock{}, nil
}
func (c *ClientMock) GetVersion() (version string, err error) {
	return "", nil
}

func (c *ClientMock) GetGenesis() (genesis GenesisObject, err error) {
	return GenesisObject{}, ErrNotImplemented
}

func (c *ClientMock) GetLogs(fromBlock, toBlock uint64, addresses []string, topics [][]string) (logs []EventLog, err error) {
	return nil, ErrNotImplemented
}

func (c *ClientMock) GetChainStats() (ns storage.FinalizationData, err error) {
	return storage.FinalizationData{}, ErrNotImplemented
}

func (c *ClientMock) GetPreviousBlockCertVotes(period uint64) (vr VotesResponse, err error) {
	return VotesResponse{}, nil
}

func (c *ClientMock) GetValidatorsAtBlock(uint64) (validators []Validator, err error) {
	return []Validator{}, nil
}

func (c *ClientMock) GetTotalAmountDelegated(uint64) (totalAmountDelegated *big.Int, err error) {
	return big.NewInt(0), nil
}

func (c *ClientMock) GetTotalSupply(uint64) (totalSupply *big.Int, err error) {
	return big.NewInt(0), nil
}

func (c *ClientMock) SubscribeNewHeads() (chan Block, *rpc.ClientSubscription, error) {
	return nil, nil, nil
}

func (c *ClientMock) Close() {
}

func (c *ClientMock) AddTransactionFromJson(trx_json string) {
	var trx Transaction
	err := json.Unmarshal([]byte(trx_json), &trx)
	if err != nil {
		fmt.Println("ClientMock.AddTransactionFromJson", err)
	}

	trx.SetTimestamp(1)
	bn := trx.GetStorage().BlockNumber
	c.BlockTransactions[bn] = append(c.BlockTransactions[bn], trx.Hash)
	c.Transactions[trx.Hash] = trx
}

func (c *ClientMock) AddLogsFromJson(trx_json string) {
	var trx Transaction
	err := json.Unmarshal([]byte(trx_json), &trx)
	if err != nil {
		fmt.Println(err)
	}

	c.EventLogs[trx.Hash] = trx.Logs
	c.Transactions[trx.Hash] = trx
}

func (c *ClientMock) AddTracesFromJson(hash, traces_json string) {
	var traces []TransactionTrace
	err := json.Unmarshal([]byte(traces_json), &traces)
	if err != nil {
		fmt.Println("ClientMock.AddTracesFromJson", err)
	}

	c.Traces[hash] = traces
}

func (c *ClientMock) AddPbftBlock(period uint64, block *Block) {
	c.Blocks[period] = block
}
