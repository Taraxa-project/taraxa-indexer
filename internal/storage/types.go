package storage

import (
	"fmt"
	"math/big"
	"reflect"
	"strings"

	"github.com/Taraxa-project/taraxa-indexer/models"
)

type TotalSupply = big.Int

type Paginated interface {
	models.Transaction | models.Dag | models.Pbft
}

type Yields interface {
	ValidatorsYield | MultipliedYield
}

func GetTypeName[T any]() string {
	var t T
	tt := reflect.TypeOf(t)
	// Don't include package name in this returned value
	return strings.Split(tt.String(), ".")[1]
}

type GenesisHash string

type ValidatorYield struct {
	Validator string   `json:"validator"`
	Yield     *big.Int `json:"yield"`
}

type ValidatorsYield struct {
	Yields []ValidatorYield `json:"yields"`
}

type MultipliedYield struct {
	Yield *big.Int `json:"yield"`
}

type Yield struct {
	Yield string `json:"yield"`
}

type ValidatorStats struct {
	// count of rewardable(with 1 or more unique transactions) DAG blocks produced by this validator
	DagBlocksCount uint64
	// Validator cert voted block weight
	VoteWeight uint64
	// Validator fee reward amount
	FeeReward *big.Int
}

type ValidatorStatsWithAddress struct {
	ValidatorStats
	Address string
}

type TotalRewardsStats struct {
	BlockAuthor      string
	TotalVotesWeight uint64
	MaxVotesWeight   uint64
	TotalDagCount    uint64
}

// map can't be serialized to rlp, so we need to use slice of structs
type RewardsStats struct {
	TotalRewardsStats
	ValidatorsStats []ValidatorStatsWithAddress
}

func FormatIntToKey(i uint64) string {
	return fmt.Sprintf("%020d", i)
}

type YieldSaving struct {
	Time   uint64
	Period uint64
}
