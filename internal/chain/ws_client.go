package chain

import (
	"context"
	"fmt"

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
	return
}

func (client *WsClient) GetLatestPeriod() (uint64, error) {
	blk := new(Block)
	err := client.rpc.Call(blk, "eth_getBlockByNumber", "latest", false)
	if err != nil {
		return 0, err
	}
	return ParseInt(blk.Number), err
}

func (client *WsClient) GetTransactionByHash(hash string) (trx *transaction, err error) {
	trx = new(transaction)
	err = client.rpc.Call(trx, "eth_getTransactionByHash", hash)
	if err != nil {
		return
	}
	err = client.AddTransactionReceiptData(trx)
	return
}

func (client *WsClient) AddTransactionReceiptData(trx *transaction) (err error) {
	err = client.rpc.Call(&trx, "eth_getTransactionReceipt", trx.Hash)
	return
}

func (client *WsClient) GetPbftBlockWithDagBlocks(period uint64) (pbftWithDags *pbftBlockWithDags, err error) {
	pbftWithDags = new(pbftBlockWithDags)
	err = client.rpc.Call(&pbftWithDags, "taraxa_getScheduleBlockByPeriod", fmt.Sprintf("0x%x", period))
	return
}

func (client *WsClient) GetDagBlockByHash(hash string) (dag *dagBlock, err error) {
	dag = new(dagBlock)
	err = client.rpc.Call(&dag, "taraxa_getDagBlockByHash", hash, false)
	return
}

func (client *WsClient) GetGenesis() (genesis *GenesisObject, err error) {
	genesis = new(GenesisObject)
	err = client.rpc.Call(&genesis, "taraxa_getConfig")
	return
}

func (client *WsClient) GetChainStats() (ns *storage.FinalizationData, err error) {
	ns = new(storage.FinalizationData)
	err = client.rpc.Call(&ns, "taraxa_getChainStats")
	return
}

func (client *WsClient) SubscribeNewHeads() (chan *Block, *rpc.ClientSubscription, error) {
	ch := make(chan *Block)
	sub, err := client.rpc.Subscribe(client.ctx, "eth", ch, "newHeads")
	return ch, sub, err
}

// Close disconnects from the node
func (client *WsClient) Close() {
	client.rpc.Close()
}
