package storage

import (
	"math/big"
	"reflect"
	"strings"
	"sync"

	"github.com/Taraxa-project/taraxa-indexer/models"
	log "github.com/sirupsen/logrus"
)

type TotalSupply = big.Int

type Paginated interface {
	models.Transaction | models.Dag | models.Pbft
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

type Account struct {
	Address string       `json:"address"`
	Balance *big.Int     `json:"balance"`
	mutex   sync.RWMutex `rlp:"-"`
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

func (a *Account) AddToBalance(value big.Int) big.Int {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	newBalance := a.Balance.Add(a.Balance, &value)
	return *newBalance
}

func (a *Account) ToModel() *models.Account {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	balanceStr := a.Balance.String()
	return &models.Account{
		Address: &a.Address,
		Balance: &balanceStr,
	}
}

func (a *Account) SubtractFromBalance(value big.Int) big.Int {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	newBalance := a.Balance.Sub(a.Balance, &value)
	return *newBalance
}

func MakeEmptyAddressStats(addr string) *AddressStats {
	data := new(AddressStats)
	data.Address = addr
	return data
}

func MakeEmptyAccount(addr string) *Account {
	data := new(Account)
	data.Address = addr
	data.Balance = big.NewInt(0)
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

type FinalizationData struct {
	DagCount  uint64 `json:"dag_blocks_executed"`
	TrxCount  uint64 `json:"transactions_executed"`
	PbftCount uint64 `json:"pbft_period"`
}

func (local *FinalizationData) Check(remote *FinalizationData) {
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
