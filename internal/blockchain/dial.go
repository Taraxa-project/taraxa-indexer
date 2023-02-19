package blockchain

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/rpc"
)

type block struct {
	Author       string   `json:"author"`
	Hash         string   `json:"hash"`
	Number       string   `json:"number"`
	Timestamp    string   `json:"timestamp"`
	Transactions []string `json:"transactions"`
}

type tx struct {
	Hash             string `json:"hash"`
	BlockNumber      string `json:"blockNumber"`
	TransactionIndex string `json:"transactionIndex"`
	From             string `json:"from"`
	To               string `json:"to"`
	Value            string `json:"value"`
	GasPrice         string `json:"gasPrice"`
}

type txReceipt struct {
	GasUsed string `json:"gasUsed"`
	Status  string `json:"status"`
}

// Blockchain is a struct that connects to a Taraxa node.
type Blockchain struct {
	client *rpc.Client
	ctx    context.Context
}

// NewBlockchain creates a new instance of the Blockchain struct.
func NewBlockchain(url string) (*Blockchain, error) {
	ctx := context.Background()
	client, err := rpc.DialWebsocket(ctx, url, "")
	if err != nil {
		return nil, err
	}
	return &Blockchain{client: client, ctx: ctx}, nil
}

// Call calls an RPC method on the chain
func (l *Blockchain) Call(method string, args ...interface{}) map[string]interface{} {
	var result map[string]interface{}
	err := l.client.Call(&result, method, args...)
	if err != nil {
		log.Fatal(err.Error())
	}
	return result
}

func (l *Blockchain) GetBlockByNumber(number uint64) block {
	var result block
	err := l.client.Call(&result, "eth_getBlockByNumber", fmt.Sprintf("0x%x", number), false)
	if err != nil {
		log.Fatal(err.Error())
	}
	return result
}

func (l *Blockchain) GetTransactionByHash(hash string) tx {
	var result tx
	err := l.client.Call(&result, "eth_getTransactionByHash", hash)
	if err != nil {
		log.Fatal(err.Error())
	}
	return result
}

func (l *Blockchain) GetTransactionReceipt(hash string) txReceipt {
	var result txReceipt
	err := l.client.Call(&result, "eth_getTransactionReceipt", hash)
	if err != nil {
		log.Fatal(err.Error())
	}
	return result
}

// Close disconnects from the blockchain node
func (l *Blockchain) Close() {
	l.client.Close()
}
