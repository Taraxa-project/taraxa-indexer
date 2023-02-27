package chain

import (
	"log"
	"strconv"

	"github.com/Taraxa-project/taraxa-indexer/models"
)

func ParseInt(s string) uint64 {
	v, err := strconv.ParseUint(s, 0, 64)
	if err != nil {
		log.Fatal(s, "ParseInt ", err)
	}
	return v
}

func parseBool(v string) bool {
	i, err := strconv.ParseUint(v, 0, 64)
	if err != nil {
		log.Fatal("parseBool ", v)
	}
	return i > 0
}

type Block struct {
	models.Pbft
	Number       string   `json:"number"`
	Timestamp    string   `json:"timestamp"`
	Transactions []string `json:"transactions"`
}

func (b *Block) ToModel() (pbft *models.Pbft) {
	pbft = &b.Pbft
	pbft.Timestamp = ParseInt(b.Timestamp)
	pbft.Number = ParseInt(b.Number)
	pbft.TransactionCount = uint64(len(b.Transactions))

	return
}

type dagBlock struct {
	models.Dag
	Level        string   `json:"level"`
	Timestamp    string   `json:"timestamp"`
	Transactions []string `json:"transactions"`
}

func (b *dagBlock) ToModel() (dag *models.Dag) {
	dag = &b.Dag
	dag.Timestamp = ParseInt(b.Timestamp)
	dag.Level = ParseInt(b.Level)
	dag.TransactionCount = uint64(len(b.Transactions))

	return
}

type transaction struct {
	models.Transaction
	BlockNumber      string `json:"blockNumber"`
	Nonce            string `json:"nonce"`
	GasPrice         string `json:"gasPrice"`
	GasUsed          string `json:"gasUsed"`
	Status           string `json:"status"`
	TransactionIndex string `json:"transactionIndex"`
	Input            string `json:"input"`
}

const emptyInput = "0x"
const emptyReceiver = ""

func (t *transaction) GetType() models.TransactionType {
	if t.To == emptyReceiver {
		return models.ContractCreation
	}
	if t.Input != emptyInput {
		return models.ContractCall
	}

	return models.Transfer
}

func (t *transaction) ToModelWithTimestamp(timestamp uint64) (trx *models.Transaction) {
	trx = &t.Transaction
	trx.BlockNumber = ParseInt(t.BlockNumber)
	trx.Nonce = ParseInt(t.Nonce)
	trx.GasPrice = ParseInt(t.GasPrice)
	trx.GasUsed = ParseInt(t.GasUsed)
	trx.TransactionIndex = ParseInt(t.TransactionIndex)
	trx.Status = parseBool(t.Status)
	trx.Type = t.GetType()
	trx.Timestamp = timestamp

	return
}

type pbftBlockWithDags struct {
	BlockHash string `json:"block_hash"`
	Period    uint64 `json:"period"`
	Schedule  struct {
		DagBlocksOrder []string `json:"dag_blocks_order"`
	} `json:"schedule"`
}

type GenesisObject struct {
	DagGenesisBlock dagBlock          `json:"dag_genesis_block"`
	InitialBalances map[string]string `json:"initial_balances"`
}
