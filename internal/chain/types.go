package chain

import (
	"context"
	"encoding/json"
	"math/big"

	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/ethereum/go-ethereum"
)

type Block struct {
	models.Pbft
	Number       string   `json:"number"`
	Timestamp    string   `json:"timestamp"`
	Transactions []string `json:"transactions"`
	TotalReward  string   `json:"totalReward"`
}

func (b *Block) ToModel() (pbft *models.Pbft) {
	pbft = &b.Pbft
	pbft.Timestamp = common.ParseUInt(b.Timestamp)
	pbft.Number = common.ParseUInt(b.Number)
	pbft.TransactionCount = uint64(len(b.Transactions))

	return
}

type DagBlock struct {
	models.Dag
	Sender       string   `json:"sender"`
	Level        string   `json:"level"`
	Timestamp    string   `json:"timestamp"`
	Transactions []string `json:"transactions"`
}

func (b *DagBlock) ToModel() (dag *models.Dag) {
	dag = &b.Dag
	dag.Timestamp = common.ParseUInt(b.Timestamp)
	dag.Level = common.ParseUInt(b.Level)
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
	BlockNumber      string   `json:"blockNumber"`
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
	ContractAddress  string     `json:"contractAddress"`
}

const emptyInput = "0x"
const emptyReceiver = ""

func GetInternalTransactionTarget(trace TraceEntry) string {
	if trace.Action.To != "" {
		return trace.Action.To
	}
	return trace.Result.Address
}

func GetTransactionType(to, input, txType string, internal bool) models.TransactionType {
	if internal {
		if txType == "create" {
			return models.InternalContractCreation
		} else if txType == "call" && input != emptyInput {
			return models.InternalContractCall
		}
		return models.InternalTransfer
	} else {
		if to == emptyReceiver {
			return models.ContractCreation
		} else if input != emptyInput {
			return models.ContractCall
		}
		return models.Transfer
	}
}

func (t *Transaction) ToModelWithTimestamp(timestamp uint64) (trx models.Transaction) {
	trx = t.Transaction
	trx.BlockNumber = common.ParseUInt(t.BlockNumber)
	trx.GasCost = common.ParseUInt(t.GasPrice) * common.ParseUInt(t.GasUsed)
	trx.Status = common.ParseBool(t.Status)
	trx.Type = GetTransactionType(trx.To, t.Input, "", false)
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
			LogIndex:         common.ParseUInt(log.LogIndex),
			Name:             "",
			Params:           []string{},
			Removed:          log.Removed,
			Topics:           log.Topics,
			TransactionHash:  log.TransactionHash,
			TransactionIndex: common.ParseUInt(log.TransactionIndex),
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
	Address string `json:"address,omitempty"`
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

type Validator struct {
	Address    string   `json:"address"`
	TotalStake *big.Int `json:"stake"`
}

func (v *Validator) UnmarshalJSON(data []byte) error {
	var res map[string]string

	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	v.Address = res["address"]
	v.TotalStake = big.NewInt(0)
	v.TotalStake.SetString(res["total_stake"], 10)

	return nil
}

type EthereumClient interface {
	CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error)
	// ... other methods as needed
}
