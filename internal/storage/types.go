package storage

import (
	"reflect"
	"sort"
	"strings"
	"sync"

	"github.com/Taraxa-project/taraxa-indexer/models"
	log "github.com/sirupsen/logrus"
)

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

func (a *AddressStats) isEqual(b *AddressStats) bool {
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

type WeekStats struct {
	Validators []models.Validator
	Total      uint32
	key        []byte `rlp:"-"`
}

func MakeEmptyWeekStats() *WeekStats {
	data := new(WeekStats)
	return data
}

func (w *WeekStats) Sort() {
	sort.Slice(w.Validators, func(i, j int) bool {
		return w.Validators[i].PbftCount > w.Validators[j].PbftCount
	})
}

func (w *WeekStats) AddPbftBlock(block *models.Pbft) {
	w.Total++
	for k, v := range w.Validators {
		if v.Address == block.Author {
			w.Validators[k].PbftCount++
			return
		}
	}
	w.Validators = append(w.Validators, models.Validator{Address: block.Author, PbftCount: 1})
}

func (w *WeekStats) GetPaginated(from, count uint64) ([]models.Validator, *models.PaginatedResponse) {
	pagination := new(models.PaginatedResponse)
	pagination.Total = uint64(len(w.Validators))
	pagination.Start = from
	end := from + count
	pagination.HasNext = (end < pagination.Total)
	if end > pagination.Total {
		end = pagination.Total
	}
	pagination.End = end

	w.Sort()
	var validators []models.Validator

	for k, v := range w.Validators[from:end] {
		v.Rank = uint64(k + 1)
		validators = append(validators, v)
	}

	return validators, pagination
}

type GenesisHash string
