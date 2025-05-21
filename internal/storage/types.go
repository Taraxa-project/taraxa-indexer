package storage

import (
	"fmt"
	"math/big"
	"reflect"
	"strings"
	"sync"

	"github.com/Taraxa-project/taraxa-indexer/internal/common"
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

// AddressStats defines the model for an address aggregate.
type AddressStats struct {
	models.StatsResponse
	Address string       `json:"address"`
	mutex   sync.RWMutex `rlp:"-"`
}

func (a *AddressStats) RegisterValidatorBlock(blockHeight uint64) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.ValidatorRegisteredBlock = &blockHeight
}

func (a *AddressStats) AddTransaction(timestamp models.Timestamp) uint64 {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.TransactionsCount++
	a.LastTransactionTimestamp = &timestamp
	return a.TransactionsCount
}

func (a *AddressStats) AddPbft(timestamp models.Timestamp) uint64 {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.PbftCount++
	a.LastPbftTimestamp = &timestamp
	return a.PbftCount
}

func (a *AddressStats) AddDag(timestamp models.Timestamp) uint64 {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.DagsCount++
	a.LastDagTimestamp = &timestamp
	return a.DagsCount
}

func MakeEmptyAddressStats(addr string) *AddressStats {
	data := new(AddressStats)
	data.Address = addr
	return data
}

func (a *AddressStats) IsEqual(b *AddressStats) bool {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	b.mutex.Lock()
	defer b.mutex.Unlock()
	if a.Address == b.Address && a.TransactionsCount == b.TransactionsCount && a.DagsCount == b.DagsCount && a.PbftCount == b.PbftCount {
		return true
	}
	return false
}

type AddressStatsMap struct {
	m            sync.RWMutex
	addressStats map[string]*AddressStats
}

func (a *AddressStatsMap) AddToBatch(b Batch) {
	for _, stats := range a.addressStats {
		b.Add(stats, stats.Address, 0)
	}
}

func (a *AddressStatsMap) GetAddress(s Storage, addr string) *AddressStats {
	addr = strings.ToLower(addr)
	a.m.Lock()
	defer a.m.Unlock()
	stats := a.addressStats[addr]
	if stats != nil {
		return stats
	}

	a.addressStats[addr] = s.GetAddressStats(addr)

	return a.addressStats[addr]
}

func MakeAddressStatsMap() *AddressStatsMap {
	return &AddressStatsMap{
		addressStats: make(map[string]*AddressStats),
	}
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

type TrxGasStats struct {
	TrxCount uint64   `json:"trxCount"`
	GasUsed  *big.Int `json:"gasUsed"`
}

func EmptyTrxGasStats() TrxGasStats {
	return TrxGasStats{
		TrxCount: 0,
		GasUsed:  big.NewInt(0),
	}
}

type DayStatsWithTimestamp struct {
	TrxGasStats
	Timestamp uint64 `json:"timestamp"`
}

func (d *DayStatsWithTimestamp) AddBlock(blk *common.Block) {
	day_start := common.DayStart(blk.Timestamp)
	if day_start > d.Timestamp {
		*d = *MakeDayStatsWithTimestamp(day_start)
	}
	d.TrxCount += blk.TransactionCount
	d.GasUsed.Add(d.GasUsed, blk.GasUsed)
}
func MakeDayStatsWithTimestamp(ts uint64) *DayStatsWithTimestamp {
	return &DayStatsWithTimestamp{
		TrxGasStats: EmptyTrxGasStats(),
		Timestamp:   ts,
	}
}

func GetTimestampFromKey(key []byte) uint64 {
	ts := strings.Split(string(key), "|")
	return common.ParseUInt(strings.TrimLeft(ts[1], "0"))
}

func (d *TrxGasStats) Add(other TrxGasStats) {
	d.TrxCount += other.TrxCount
	d.GasUsed.Add(d.GasUsed, other.GasUsed)
}
