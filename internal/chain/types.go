package chain

import (
	"log"
	"strconv"

	"github.com/Taraxa-project/taraxa-indexer/models"
)

func parseBool(v string) bool {
	i, err := strconv.ParseUint(v, 0, 64)
	if err != nil {
		log.Fatal("parseBool ", v)
	}
	return i > 0
}

type block struct {
	models.Pbft
	Number       string   `json:"number"`
	Timestamp    string   `json:"timestamp"`
	Transactions []string `json:"transactions"`
}

func (b *block) ToModel() (pbft *models.Pbft) {
	pbft = &b.Pbft
	pbft.Age, _ = strconv.ParseUint(b.Timestamp, 0, 64)
	pbft.Number, _ = strconv.ParseUint(b.Number, 0, 64)
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
	dag.Age, _ = strconv.ParseUint(b.Timestamp, 0, 64)
	dag.Level, _ = strconv.ParseUint(b.Level, 0, 64)
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

func (t *transaction) ToModelWithAge(age uint64) (trx *models.Transaction) {
	trx = &t.Transaction
	trx.BlockNumber, _ = strconv.ParseUint(t.BlockNumber, 0, 64)
	trx.Nonce, _ = strconv.ParseUint(t.Nonce, 0, 64)
	trx.GasPrice, _ = strconv.ParseUint(t.GasPrice, 0, 64)
	trx.GasUsed, _ = strconv.ParseUint(t.GasUsed, 0, 64)
	trx.TransactionIndex, _ = strconv.ParseUint(t.TransactionIndex, 0, 64)
	trx.Status = parseBool(t.Status)
	trx.Type = t.GetType()
	trx.Age = age

	return
}

type pbftBlockWithDags struct {
	BlockHash string `json:"block_hash"`
	Period    uint64 `json:"period"`
	Schedule  struct {
		DagBlocksOrder []string `json:"dag_blocks_order"`
	} `json:"schedule"`
}
