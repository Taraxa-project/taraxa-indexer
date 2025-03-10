package storage

import (
	"fmt"
	"math/big"
	"reflect"
	"strings"
	"sync"

	"github.com/Taraxa-project/taraxa-indexer/models"
	log "github.com/sirupsen/logrus"
)

type TotalSupply = big.Int

type Paginated interface {
	Transaction | models.Dag | models.Pbft
}

type Yields interface {
	ValidatorsYield | MultipliedYield
}

func GetTypeName[T any]() string {
	var o T
	tt := reflect.TypeOf(o)
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

func (a *AddressStats) AddTransaction(timestamp models.Uint64) uint64 {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.TransactionsCount++
	a.LastTransactionTimestamp = &timestamp
	return a.TransactionsCount
}

func (a *AddressStats) AddPbft(timestamp models.Uint64) uint64 {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.PbftCount++
	a.LastPbftTimestamp = &timestamp
	return a.PbftCount
}

func (a *AddressStats) AddDag(timestamp models.Uint64) uint64 {
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

type Transaction struct {
	BlockNumber models.Uint64          `json:"blockNumber"`
	Calldata    *models.CallData       `json:"calldata,omitempty" rlp:"nil"`
	From        models.Address         `json:"from"`
	GasCost     *big.Int               `json:"gas_cost"`
	Hash        models.Hash            `json:"hash"`
	Input       string                 `json:"input"`
	Status      bool                   `json:"status"`
	Timestamp   models.Uint64          `json:"timestamp"`
	To          models.Address         `json:"to"`
	Type        models.TransactionType `json:"type"`
	Value       *big.Int               `json:"value"`
}

type InternalTransactionsResponse struct {
	Data []Transaction `json:"data"`
}
