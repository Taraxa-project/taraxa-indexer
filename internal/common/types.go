package common

import (
	"encoding/json"
	"math/big"

	"github.com/Taraxa-project/taraxa-indexer/models"
	log "github.com/sirupsen/logrus"
)

type Block struct {
	models.Pbft
	Transactions []string `json:"transactions"`
	TotalReward  string   `json:"totalReward"`
	GasUsed      *big.Int `json:"gasUsed"`
}

func (b *Block) UnmarshalJSON(data []byte) error {
	var rawStruct struct {
		Author    string `json:"author"`
		Hash      string `json:"hash"`
		Number    string `json:"number"`
		Timestamp string `json:"timestamp"`

		Transactions []string `json:"transactions"`
		TotalReward  string   `json:"totalReward"`
		GasUsed      string   `json:"gasUsed"`
	}
	if err := json.Unmarshal(data, &rawStruct); err != nil {
		panic(err)
	}
	b.Transactions = rawStruct.Transactions
	b.TotalReward = rawStruct.TotalReward
	b.GasUsed = ParseStringToBigInt(rawStruct.GasUsed)

	b.Author = rawStruct.Author
	b.Hash = rawStruct.Hash
	b.Number = ParseUInt(rawStruct.Number)
	b.Timestamp = ParseUInt(rawStruct.Timestamp)

	b.TransactionCount = uint64(len(b.Transactions))
	return nil
}

func (b *Block) GetModel() (pbft *models.Pbft) {
	return &b.Pbft
}

type DagBlock struct {
	models.Dag
	Sender       string   `json:"sender"`
	Transactions []string `json:"transactions"`
	Vdf          struct {
		Difficulty uint16 `json:"difficulty"`
	} `json:"vdf"`
}

func (b *DagBlock) UnmarshalJSON(data []byte) error {
	var rawStruct struct {
		Hash      string `json:"hash"`
		Level     string `json:"level"`
		Timestamp string `json:"timestamp"`

		Sender       string   `json:"sender"`
		Transactions []string `json:"transactions"`

		Vdf struct {
			Difficulty string `json:"difficulty"`
		} `json:"vdf"`
	}
	if err := json.Unmarshal(data, &rawStruct); err != nil {
		panic(err)
	}
	b.Sender = rawStruct.Sender
	b.Transactions = rawStruct.Transactions

	b.Hash = rawStruct.Hash
	b.Level = ParseUInt(rawStruct.Level)
	b.Timestamp = ParseUInt(rawStruct.Timestamp)
	b.TransactionCount = uint64(len(b.Transactions))
	b.Vdf.Difficulty = uint16(ParseUInt(rawStruct.Vdf.Difficulty))

	return nil
}

func (b *DagBlock) GetModel() (pbft *models.Dag) {
	return &b.Dag
}

type PbftBlockWithDags struct {
	BlockHash string `json:"block_hash"`
	Period    uint64 `json:"period"`
	Schedule  struct {
		DagBlocksOrder []string `json:"dag_blocks_order"`
	} `json:"schedule"`
}

type Action struct {
	CallType string `json:"callType"`
	From     string `json:"from"`
	Gas      string `json:"gas"`
	Input    string `json:"input"`
	To       string `json:"to"`
	Value    string `json:"value"`
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

type TransactionTrace struct {
	Trace []TraceEntry `json:"trace"`
}

type VotesResponse struct {
	PeriodTotalVotesCount uint64 `json:"total_votes_count,omitempty"`
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

type FinalizationData struct {
	DagCount  uint64 `json:"dag_blocks_executed"`
	TrxCount  uint64 `json:"transactions_executed"`
	PbftCount uint64 `json:"pbft_period"`
}

func (local *FinalizationData) Check(remote FinalizationData) {
	// Perform this check only if we are getting data for the same block from node
	if local.PbftCount != remote.PbftCount {
		return
	}
	if local.DagCount != remote.DagCount {
		log.WithFields(log.Fields{"local": local, "remote": remote}).Fatal("Dag consistency check failed")
	}

	if local.TrxCount != remote.TrxCount {
		log.WithFields(log.Fields{"local": local, "remote": remote}).Fatal("Transactions consistency check failed ")
	}
}
