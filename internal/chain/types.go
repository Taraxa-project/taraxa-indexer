package chain

import (
	"log"
	"math/big"
	"runtime/debug"
	"strconv"

	"github.com/Taraxa-project/taraxa-indexer/models"
)

func ParseUInt(s string) (v uint64) {
	if len(s) == 0 {
		return
	}
	v, err := strconv.ParseUint(s, 0, 64)
	if err != nil {
		debug.PrintStack()
		log.Fatal(s, "ParseUInt ", err)
	}
	return v
}

func ParseInt(s string) (v int64) {
	if len(s) == 0 {
		return
	}
	v, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		debug.PrintStack()
		log.Fatal(s, "ParseUInt ", err)
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
	Number       string   `json:"number"`
	Timestamp    string   `json:"timestamp"`
	Transactions []string `json:"transactions"`
	TotalReward  string   `json:"totalReward"`
}

func (b *Block) ToModel() (pbft *models.Pbft) {
	pbft = &b.Pbft
	pbft.Timestamp = ParseUInt(b.Timestamp)
	pbft.Number = ParseUInt(b.Number)
	pbft.TransactionCount = uint64(len(b.Transactions))

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
	dag.Timestamp = ParseUInt(b.Timestamp)
	dag.Level = ParseUInt(b.Level)
	dag.TransactionCount = uint64(len(b.Transactions))

	return
}

type EventLog struct {
	Address          string   `json:"address"`
	Data             string   `json:"data"`
	LogIndex         string   `json:"logIndex"`
	Removed          bool     `json:"removed"`
	Topics           []string `json:"topics"`
	TransactionHash  string   `json:"transactionHash"`
	TransactionIndex string   `json:"transactionIndex"`
}

type Transaction struct {
	models.Transaction
	Logs             []EventLog `json:"logs"`
	BlockNumber      string     `json:"blockNumber"`
	Nonce            string     `json:"nonce"`
	GasPrice         string     `json:"gasPrice"`
	GasUsed          string     `json:"gasUsed"`
	Status           string     `json:"status"`
	TransactionIndex string     `json:"transactionIndex"`
	Input            string     `json:"input"`
	ContractAddress  string     `json:"contractAddress"`
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

func (t *Transaction) ToModelWithTimestamp(timestamp uint64) (trx models.Transaction) {
	trx = t.Transaction
	trx.BlockNumber = ParseUInt(t.BlockNumber)
	trx.Nonce = ParseUInt(t.Nonce)
	trx.GasPrice = ParseUInt(t.GasPrice)
	trx.GasUsed = ParseUInt(t.GasUsed)
	trx.TransactionIndex = ParseUInt(t.TransactionIndex)
	trx.Status = parseBool(t.Status)
	trx.Type = GetTransactionType(trx.To, t.Input, false)
	if trx.Type == models.ContractCreation {
		trx.To = t.ContractAddress
	}
	trx.Timestamp = timestamp

	return
}

func (t *Transaction) GetFee() *big.Int {
	gasUsed, _ := big.NewInt(0).SetString(t.GasUsed, 0)
	gasPrice, _ := big.NewInt(0).SetString(t.GasPrice, 0)

	return big.NewInt(0).Mul(gasUsed, gasPrice)
}

func (t *Transaction) ExtractLogs() (logs []models.EventLog) {
	for _, log := range t.Logs {
		eLog := models.EventLog{
			Address:          log.Address,
			Data:             log.Data,
			LogIndex:         ParseUInt(log.LogIndex),
			Removed:          log.Removed,
			Topics:           log.Topics,
			TransactionHash:  log.TransactionHash,
			TransactionIndex: ParseUInt(log.TransactionIndex),
		}
		logs = append(logs, eLog)
	}
	return logs
}

type PbftBlockWithDags struct {
	BlockHash string `json:"block_hash"`
	Period    uint64 `json:"period"`
	Schedule  struct {
		DagBlocksOrder []string `json:"dag_blocks_order"`
	} `json:"schedule"`
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

type VotesResponse struct {
	PeriodTotalVotesCount int64  `json:"total_votes_count,omitempty"`
	Votes                 []Vote `json:"votes,omitempty"`
}

type Vote struct {
	Voter  string `json:"voter"`
	Weight string `json:"weight"`
}
