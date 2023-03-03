package chain

import (
	"context"
	"fmt"

	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/ethereum/go-ethereum/rpc"

	log "github.com/sirupsen/logrus"
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

// Call calls an RPC method on the chain
func (client *WsClient) Call(method string, args ...interface{}) (res map[string]interface{}) {
	err := client.rpc.Call(&res, method, args...)
	if err != nil {
		log.WithField("error", err).Fatal("Call failed")
	}
	return
}

func (client *WsClient) GetBlockByNumber(number uint64) (blk *Block) {
	blk = new(Block)
	err := client.rpc.Call(blk, "eth_getBlockByNumber", fmt.Sprintf("0x%x", number), false)
	if err != nil {
		log.WithFields(log.Fields{"number": number, "error": err}).Fatal("GetBlockByNumber failed")
	}
	return
}

func (client *WsClient) GetLatestPeriod() uint64 {
	blk := new(Block)
	err := client.rpc.Call(blk, "eth_getBlockByNumber", "latest", false)
	if err != nil {
		log.WithField("error", err).Fatal("GetLatestPeriod failed")
	}
	return ParseInt(blk.Number)
}

func (client *WsClient) GetTransactionByHash(hash string) (trx *transaction) {
	trx = new(transaction)
	err := client.rpc.Call(trx, "eth_getTransactionByHash", hash)
	if err != nil {
		log.WithField("error", err).Fatal("GetTransactionByHash failed")
	}
	client.AddTransactionReceiptData(trx)
	return
}

func (client *WsClient) AddTransactionReceiptData(trx *transaction) {
	err := client.rpc.Call(&trx, "eth_getTransactionReceipt", trx.Hash)
	if err != nil {
		log.WithField("error", err).Fatal("AddTransactionReceiptData failed")
	}
}

func (client *WsClient) GetPbftBlockWithDagBlocks(period uint64) (pbftWithDags *pbftBlockWithDags) {
	pbftWithDags = new(pbftBlockWithDags)
	err := client.rpc.Call(&pbftWithDags, "taraxa_getScheduleBlockByPeriod", fmt.Sprintf("0x%x", period))
	if err != nil {
		log.WithFields(log.Fields{"number": period, "error": err}).Fatal("GetPbftBlockWithDagBlocks failed")
	}
	return
}

func (client *WsClient) GetDagBlockByHash(hash string) (dag *dagBlock) {
	dag = new(dagBlock)
	err := client.rpc.Call(&dag, "taraxa_getDagBlockByHash", hash, false)
	if err != nil {
		log.WithField("error", err).Fatal("GetDagBlockByHash failed")
	}
	return
}

func (client *WsClient) GetGenesis() (genesis *GenesisObject) {
	genesis = new(GenesisObject)
	err := client.rpc.Call(&genesis, "taraxa_getConfig")
	if err != nil {
		log.WithField("error", err).Fatal("GetGenesis failed")
	}
	return
}

func (client *WsClient) GetNodeStats() (ns *storage.FinalizationData, err error) {
	ns = new(storage.FinalizationData)
	err = client.rpc.Call(&ns, "get_node_status")
	return
}

func (client *WsClient) SubscribeNewHeads() (chan *Block, *rpc.ClientSubscription) {
	ch := make(chan *Block)
	sub, err := client.rpc.Subscribe(client.ctx, "eth", ch, "newHeads")
	if err != nil {
		log.WithField("error", err).Fatal("SubscribeNewHeads failed")
	}
	return ch, sub
}

// Close disconnects from the node
func (client *WsClient) Close() {
	client.rpc.Close()
}
