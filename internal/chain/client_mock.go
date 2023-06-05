package chain

import (
	"encoding/json"
	"fmt"

	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/ethereum/go-ethereum/rpc"
)

type ClientMock struct {
	Traces            map[string][]TransactionTrace
	Transactions      map[string]Transaction
	BlockTransactions map[uint64][]string
}

var ErrNotImplemented = fmt.Errorf("Not implemented")

func MakeMockClient() *ClientMock {
	m := new(ClientMock)
	m.Traces = make(map[string][]TransactionTrace)
	m.Transactions = make(map[string]Transaction)
	m.BlockTransactions = make(map[uint64][]string)
	return m
}

func (c *ClientMock) GetBlockByNumber(number uint64) (blk *Block, err error) {
	return nil, ErrNotImplemented
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

func (c *ClientMock) GetPeriodTransactions(p uint64) (trx []Transaction, err error) {
	return nil, ErrNotImplemented
}

func (c *ClientMock) GetPbftBlockWithDagBlocks(period uint64) (pbftWithDags *PbftBlockWithDags, err error) {
	return nil, ErrNotImplemented
}

func (c *ClientMock) GetDagBlockByHash(hash string) (dag *DagBlock, err error) {
	return nil, ErrNotImplemented
}

func (c *ClientMock) GetPeriodDagBlocks(period uint64) (dags []DagBlock, err error) {
	return nil, ErrNotImplemented
}

func (c *ClientMock) GetGenesis() (genesis *GenesisObject, err error) {
	return nil, ErrNotImplemented
}

func (c *ClientMock) GetChainStats() (ns *storage.FinalizationData, err error) {
	return nil, ErrNotImplemented
}

func (c *ClientMock) SubscribeNewHeads() (chan *Block, *rpc.ClientSubscription, error) {
	return nil, nil, nil
}

func (c *ClientMock) Close() {
}

func (c *ClientMock) AddTransactionFromJson(trx_json string) {
	var trx Transaction
	err := json.Unmarshal([]byte(trx_json), &trx)
	if err != nil {
		fmt.Println(err)
	}

	tm := trx.ToModelWithTimestamp(1)
	c.BlockTransactions[tm.BlockNumber] = append(c.BlockTransactions[tm.BlockNumber], trx.Hash)
	c.Transactions[trx.Hash] = trx
}

func (c *ClientMock) AddTracesFromJson(hash, traces_json string) {
	var traces []TransactionTrace
	err := json.Unmarshal([]byte(traces_json), &traces)
	if err != nil {
		fmt.Println(err)
	}

	c.Traces[hash] = traces
}
