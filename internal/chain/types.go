package chain

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/models"
)

type Block struct {
	models.Pbft
	Transactions []string `json:"transactions"`
	TotalReward  string   `json:"totalReward"`
}

func (b *Block) UnmarshalJSON(data []byte) error {
	var strMap map[string]interface{}
	if err := json.Unmarshal(data, &strMap); err != nil {
		panic(err)
	}
	fmt.Println(strMap)

	b.Transactions = strMap["transactions"].([]string)
	b.TotalReward = strMap["totalReward"].(string)

	b.Pbft.Author = strMap["author"].(string)
	b.Pbft.Hash = strMap["hash"].(string)
	b.Pbft.Number = common.ParseUInt(strMap["timestamp"].(string))
	b.Pbft.Timestamp = common.ParseUInt(strMap["timestamp"].(string))

	b.Pbft.TransactionCount = uint64(len(b.Transactions))
	return nil
}

func (b *Block) GetModel() (pbft *models.Pbft) {
	return &b.Pbft
}

type DagBlock struct {
	models.Dag
	Sender       string   `json:"sender"`
	Transactions []string `json:"transactions"`
}

func (b *DagBlock) UnmarshalJSON(data []byte) error {
	var strMap map[string]interface{}
	if err := json.Unmarshal(data, &strMap); err != nil {
		panic(err)
	}
	fmt.Println(strMap)
	b.Sender = strMap["sender"].(string)
	b.Transactions = strMap["transactions"].([]string)

	b.Dag.Hash = strMap["hash"].(string)
	b.Dag.Level = common.ParseUInt(strMap["level"].(string))
	b.Dag.Timestamp = common.ParseUInt(strMap["timestamp"].(string))
	b.Dag.TransactionCount = uint64(len(b.Transactions))

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
