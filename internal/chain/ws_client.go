package chain

import (
	"context"
	"fmt"

	"github.com/Taraxa-project/taraxa-indexer/internal/metrics"

	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/ethereum/go-ethereum/rpc"
)

// WsClient is a struct that connects to a Taraxa node.
type WsClient struct {
	rpc *rpc.Client
	ctx context.Context
}

// NewWsClient creates a new instance of the WsClient struct.
func NewWsClient(url string) (*WsClient, error) {
	ctx := context.Background()
	client, err := rpc.DialWebsocket(ctx, url, "")
	if err != nil {
		return nil, err
	}
	return &WsClient{rpc: client, ctx: ctx}, nil
}

func (client *WsClient) GetBlockByNumber(number uint64) (blk *Block, err error) {
	blk = new(Block)
	err = client.rpc.Call(blk, "eth_getBlockByNumber", fmt.Sprintf("0x%x", number), false)
	metrics.RpcCallsCounter.Inc()
	return
}

func (client *WsClient) GetLatestPeriod() (uint64, error) {
	blk := new(Block)
	err := client.rpc.Call(blk, "eth_getBlockByNumber", "latest", false)
	metrics.RpcCallsCounter.Inc()
	if err != nil {
		return 0, err
	}
	return ParseInt(blk.Number), err
}

func (client *WsClient) TraceBlockTransactions(number uint64) (traces []TransactionTrace, err error) {
	err = client.rpc.Call(&traces, "trace_replayBlockTransactions", fmt.Sprintf("0x%x", number), []string{"trace"})
	defer metrics.RpcCallsCounter.Inc()
	return
}

// TODO: Optimize this. We are making two requests here, so its pretty slow
func (client *WsClient) GetTransactionByHash(hash string) (trx Transaction, err error) {
	err = client.rpc.Call(&trx, "eth_getTransactionByHash", hash)
	metrics.RpcCallsCounter.Inc()
	if err != nil {
		return
	}
	err = client.addTransactionReceiptData(&trx)

	return
}

func (client *WsClient) addTransactionReceiptData(trx *Transaction) (err error) {
	err = client.rpc.Call(&trx, "eth_getTransactionReceipt", trx.Hash)
	metrics.RpcCallsCounter.Inc()
	return
}

func (client *WsClient) GetPeriodTransactions(number uint64) (trxs []Transaction, err error) {
	err = client.rpc.Call(&trxs, "taraxa_getPeriodTransactionsWithReceipts", fmt.Sprintf("0x%x", number))
	metrics.RpcCallsCounter.Inc()
	return
}

func (client *WsClient) GetPbftBlockWithDagBlocks(period uint64) (pbftWithDags *PbftBlockWithDags, err error) {
	pbftWithDags = new(PbftBlockWithDags)
	err = client.rpc.Call(&pbftWithDags, "taraxa_getScheduleBlockByPeriod", fmt.Sprintf("0x%x", period))
	metrics.RpcCallsCounter.Inc()
	return
}

func (client *WsClient) GetDagBlockByHash(hash string) (dag *DagBlock, err error) {
	dag = new(DagBlock)
	err = client.rpc.Call(&dag, "taraxa_getDagBlockByHash", hash, false)
	metrics.RpcCallsCounter.Inc()
	return
}

func (client *WsClient) GetPeriodDagBlocks(period uint64) (dags []DagBlock, err error) {
	err = client.rpc.Call(&dags, "taraxa_getPeriodDagBlocks", fmt.Sprintf("0x%x", period))
	metrics.RpcCallsCounter.Inc()
	return
}

func (client *WsClient) GetGenesis() (genesis *GenesisObject, err error) {
	genesis = new(GenesisObject)
	err = client.rpc.Call(&genesis, "taraxa_getConfig")
	metrics.RpcCallsCounter.Inc()
	return
}

func (client *WsClient) GetChainStats() (ns *storage.FinalizationData, err error) {
	ns = new(storage.FinalizationData)
	err = client.rpc.Call(&ns, "taraxa_getChainStats")
	metrics.RpcCallsCounter.Inc()
	return
}

func (client *WsClient) SubscribeNewHeads() (chan *Block, *rpc.ClientSubscription, error) {
	ch := make(chan *Block)
	sub, err := client.rpc.Subscribe(client.ctx, "eth", ch, "newHeads")
	metrics.RpcCallsCounter.Inc()
	return ch, sub, err
}

// Close disconnects from the node
func (client *WsClient) Close() {
	client.rpc.Close()
}
