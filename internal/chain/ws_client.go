package chain

import (
	"context"
	"fmt"
	"math/big"

	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/metrics"
	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/gorilla/websocket"

	"github.com/ethereum/go-ethereum/rpc"
	log "github.com/sirupsen/logrus"
)

// WsClient is a struct that connects to a Taraxa node.
type WsClient struct {
	rpc     *rpc.Client
	ctx     context.Context
	ChainId *big.Int
}

// NewWsClient creates a new instance of the WsClient struct.
func NewWsClient(url string) (*WsClient, error) {
	ctx := context.Background()
	options := []rpc.ClientOption{
		rpc.WithWebsocketMessageSizeLimit(0),
		rpc.WithWebsocketDialer(websocket.Dialer{
			EnableCompression: true,
		}),
	}
	ws, err := rpc.DialOptions(ctx, url, options...)

	if err != nil {
		return nil, err
	}
	client := &WsClient{rpc: ws, ctx: ctx}
	client.GetChainId()

	return client, nil
}

func (client *WsClient) GetChainId() *big.Int {
	if client.ChainId != nil {
		return client.ChainId
	}

	var str string
	err := client.rpc.Call(&str, "eth_chainId")
	if err != nil {
		log.WithError(err).Panic("GetChainId error")
	}
	metrics.RpcCallsCounter.Inc()
	client.ChainId = big.NewInt(0)
	client.ChainId.SetString(str, 0)
	return client.ChainId
}

func (client *WsClient) GetBalanceAtBlock(address string, blockNumber uint64) (balance string, err error) {
	blkNumberHex := fmt.Sprintf("0x%x", blockNumber)
	err = client.rpc.Call(&balance, "eth_getBalance", address, blkNumberHex)
	metrics.RpcCallsCounter.Inc()
	return
}

func (client *WsClient) GetBlocks(start, end uint64) (blocks []*common.Block, err error) {
	blocks = make([]*common.Block, end-start+1)
	batch := make([]rpc.BatchElem, end-start+1)
	for i := end; i >= start; i-- {
		batch[i-start] = rpc.BatchElem{
			Method: "eth_getBlockByNumber",
			Args:   []any{fmt.Sprintf("0x%x", i), false},
			Result: &blocks[end-i],
		}
	}
	err = client.rpc.BatchCall(batch)
	metrics.RpcCallsCounter.Inc()
	return
}

func (client *WsClient) GetBlockByNumber(number uint64) (blk *common.Block, err error) {
	err = client.rpc.Call(&blk, "eth_getBlockByNumber", fmt.Sprintf("0x%x", number), false)
	metrics.RpcCallsCounter.Inc()
	return
}

func (client *WsClient) GetLatestPeriod() (uint64, error) {
	blk := new(common.Block)
	err := client.rpc.Call(blk, "eth_getBlockByNumber", "latest", false)
	metrics.RpcCallsCounter.Inc()
	if err != nil {
		return 0, err
	}
	return blk.Number, err
}

func (client *WsClient) TraceBlockTransactions(number uint64) (traces []common.TransactionTrace, err error) {
	err = client.rpc.Call(&traces, "trace_replayBlockTransactions", fmt.Sprintf("0x%x", number), []string{"trace"})
	defer metrics.RpcCallsCounter.Inc()
	return
}

// TODO: Optimize this. We are making two requests here, so its pretty slow
func (client *WsClient) GetTransactionByHash(hash string) (trx common.Transaction, err error) {
	err = client.rpc.Call(&trx, "eth_getTransactionByHash", hash)
	metrics.RpcCallsCounter.Inc()
	if err != nil {
		return
	}
	err = client.addTransactionReceiptData(&trx)

	return
}

func (client *WsClient) addTransactionReceiptData(trx *common.Transaction) (err error) {
	err = client.rpc.Call(&trx, "eth_getTransactionReceipt", trx.Hash)
	metrics.RpcCallsCounter.Inc()
	return
}

func (client *WsClient) GetPeriodTransactions(number uint64) (trxs []common.Transaction, err error) {
	err = client.rpc.Call(&trxs, "debug_getPeriodTransactionsWithReceipts", fmt.Sprintf("0x%x", number))
	metrics.RpcCallsCounter.Inc()
	return
}

func (client *WsClient) GetPbftBlockWithDagBlocks(period uint64) (pbftWithDags common.PbftBlockWithDags, err error) {
	err = client.rpc.Call(&pbftWithDags, "taraxa_getScheduleBlockByPeriod", fmt.Sprintf("0x%x", period))
	metrics.RpcCallsCounter.Inc()
	return
}

func (client *WsClient) GetDagBlockByHash(hash string) (dag common.DagBlock, err error) {
	err = client.rpc.Call(&dag, "taraxa_getDagBlockByHash", hash, false)
	metrics.RpcCallsCounter.Inc()
	return
}

func (client *WsClient) GetPeriodDagBlocks(period uint64) (dags []common.DagBlock, err error) {
	err = client.rpc.Call(&dags, "debug_getPeriodDagBlocks", fmt.Sprintf("0x%x", period))
	metrics.RpcCallsCounter.Inc()
	return
}

func (client *WsClient) GetGenesis() (genesis common.GenesisObject, err error) {
	err = client.rpc.Call(&genesis, "taraxa_getConfig")
	metrics.RpcCallsCounter.Inc()
	return
}

func (client *WsClient) GetVersion() (version string, err error) {
	versionResponse := make(map[string]string)
	err = client.rpc.Call(&versionResponse, "taraxa_getVersion")
	version = versionResponse["version"]
	metrics.RpcCallsCounter.Inc()
	return
}

func (client *WsClient) GetChainStats() (fd common.FinalizationData, err error) {
	err = client.rpc.Call(&fd, "taraxa_getChainStats")
	metrics.RpcCallsCounter.Inc()
	return
}

func (client *WsClient) GetPreviousBlockCertVotes(period uint64) (vr common.VotesResponse, err error) {
	err = client.rpc.Call(&vr, "debug_getPreviousBlockCertVotes", fmt.Sprintf("0x%x", period))
	metrics.RpcCallsCounter.Inc()
	return
}

func (client *WsClient) GetLogs(fromBlock, toBlock uint64, addresses []string, topics [][]string) (logs []common.EventLog, err error) {
	err = client.rpc.Call(&logs, "eth_getLogs", map[string]any{
		"fromBlock": fmt.Sprintf("0x%x", fromBlock),
		"toBlock":   fmt.Sprintf("0x%x", toBlock),
		"address":   addresses,
		"topics":    topics,
	})
	metrics.RpcCallsCounter.Inc()
	return
}

func (client *WsClient) GetValidatorsAtBlock(period uint64) (validators []common.Validator, err error) {
	err = client.rpc.Call(&validators, "debug_dposValidatorTotalStakes", fmt.Sprintf("0x%x", period))
	return
}

func (client *WsClient) GetTotalAmountDelegated(block_num uint64) (totalAmountDelegated *big.Int, err error) {
	delegatedStr := ""
	err = client.rpc.Call(&delegatedStr, "debug_dposTotalAmountDelegated", fmt.Sprintf("0x%x", block_num))
	totalAmountDelegated = common.ParseStringToBigInt(delegatedStr)
	metrics.RpcCallsCounter.Inc()
	return
}

func (client *WsClient) GetTotalSupply(block_num uint64) (totalSupply *big.Int, err error) {
	supplyStr := ""
	err = client.rpc.Call(&supplyStr, "taraxa_totalSupply", fmt.Sprintf("0x%x", block_num))
	totalSupply = common.ParseStringToBigInt(supplyStr)
	metrics.RpcCallsCounter.Inc()
	return
}

func (client *WsClient) GetPeriodLambda(period uint64) (lambdaMs *uint64, err error) {
	lambdaStr := ""
	err = client.rpc.Call(&lambdaStr, "taraxa_getPeriodLambda", fmt.Sprintf("0x%x", period))
	if err != nil {
		return
	}
	if lambdaStr == "" {
		return
	}
	lambdaMs = new(uint64)
	*lambdaMs = common.ParseUInt(lambdaStr)
	metrics.RpcCallsCounter.Inc()
	return
}

func (client *WsClient) FilterContracts(addresses []models.Address) (contracts []models.Address, err error) {
	contracts = make([]models.Address, 0, len(addresses))
	batch := make([]rpc.BatchElem, len(addresses))
	results := make([]string, len(addresses))
	for i, address := range addresses {
		batch[i] = rpc.BatchElem{
			Method: "eth_getCode",
			Args:   []any{address, "latest"},
			Result: &results[i],
		}
	}
	err = client.rpc.BatchCall(batch)
	if err != nil {
		return
	}
	for i, result := range results {
		if result != "0x" {
			contracts = append(contracts, addresses[i])
		}
	}
	metrics.RpcCallsCounter.Inc()
	return
}

func (client *WsClient) SubscribeNewHeads() (chan common.Block, *rpc.ClientSubscription, error) {
	ch := make(chan common.Block)
	sub, err := client.rpc.Subscribe(client.ctx, "eth", ch, "newHeads")
	metrics.RpcCallsCounter.Inc()
	return ch, sub, err
}

// Close disconnects from the node
func (client *WsClient) Close() {
	client.rpc.Close()
}
