package chain

import (
	"log"
	"runtime/debug"
	"strconv"

	"github.com/Taraxa-project/taraxa-indexer/models"
)

func ParseInt(s string) (v uint64) {
	if len(s) == 0 {
		return
	}
	v, err := strconv.ParseUint(s, 0, 64)
	if err != nil {
		debug.PrintStack()
		log.Fatal(s, "ParseInt ", err)
	}
	return v
}

func parseBool(s string) (v bool) {
	if len(s) == 0 {
		return
	}
	i, err := strconv.ParseUint(s, 0, 64)
	if err != nil {
		log.Fatal("parseBool ", v)
	}
	return i > 0
}

type Block struct {
	models.Pbft
	Number       string    `json:"number"`
	Timestamp    string    `json:"timestamp"`
	Transactions *[]string `json:"transactions"`
	TotalReward  string    `json:"totalReward"`
}

func (b *Block) ToModel() (pbft *models.Pbft) {
	pbft = &b.Pbft
	pbft.Timestamp = ParseInt(b.Timestamp)
	pbft.Number = ParseInt(b.Number)
	pbft.TransactionCount = uint64(len(*b.Transactions))

	return
}

type DagBlock struct {
	models.Dag
	Level        string   `json:"level"`
	Timestamp    string   `json:"timestamp"`
	Transactions []string `json:"transactions"`
}

func (b *DagBlock) ToModel() (dag *models.Dag) {
	dag = &b.Dag
	dag.Timestamp = ParseInt(b.Timestamp)
	dag.Level = ParseInt(b.Level)
	dag.TransactionCount = uint64(len(b.Transactions))

	return
}

type Transaction struct {
	models.Transaction
	BlockNumber      string `json:"blockNumber"`
	Nonce            string `json:"nonce"`
	GasPrice         string `json:"gasPrice"`
	GasUsed          string `json:"gasUsed"`
	Status           string `json:"status"`
	TransactionIndex string `json:"transactionIndex"`
	Input            string `json:"input"`
	ContractAddress  string `json:"contractAddress"`
}

const emptyInput = "0x"
const emptyReceiver = ""

func GetTransactionType(to, input string, internal bool) models.TransactionType {
	trx_type := 0
	// add offset if transaction is internal
	if internal {
		trx_type = 3
	}

	if to == emptyReceiver {
		trx_type += int(models.ContractCreation)
	} else if input != emptyInput {
		trx_type += int(models.ContractCall)
	}
	return models.TransactionType(trx_type)
}

func (t *Transaction) ToModelWithTimestamp(timestamp uint64) (trx *models.Transaction) {
	trx = &t.Transaction
	trx.BlockNumber = ParseInt(t.BlockNumber)
	trx.Nonce = ParseInt(t.Nonce)
	trx.GasPrice = ParseInt(t.GasPrice)
	trx.GasUsed = ParseInt(t.GasUsed)
	trx.TransactionIndex = ParseInt(t.TransactionIndex)
	trx.Status = parseBool(t.Status)
	trx.Type = GetTransactionType(trx.To, t.Input, false)
	if trx.Type == models.ContractCreation {
		trx.To = t.ContractAddress
	}
	trx.Timestamp = timestamp

	return
}

type PbftBlockWithDags struct {
	BlockHash string `json:"block_hash"`
	Period    uint64 `json:"period"`
	Schedule  struct {
		DagBlocksOrder []string `json:"dag_blocks_order"`
	} `json:"schedule"`
}

type GenesisObject struct {
	DagGenesisBlock DagBlock          `json:"dag_genesis_block"`
	InitialBalances map[string]string `json:"initial_balances"`
}

type TransactionTrace struct {
	Trace []TraceEntry `json:"trace"`
}

type TraceEntryResult struct {
	Output  string `json:"output"`
	GasUsed string `json:"gasUsed"`
}

type TraceEntry struct {
	Action       Action           `json:"action"`
	Subtraces    uint16           `json:"subtraces"`
	TraceAddress []uint16         `json:"traceAddress"`
	Type         string           `json:"type"`
	Result       TraceEntryResult `json:"result"`
}

type Action struct {
	CallType string `json:"callType"`
	From     string `json:"from"`
	Gas      string `json:"gas"`
	Input    string `json:"input"`
	To       string `json:"to"`
	Value    string `json:"value"`
}
