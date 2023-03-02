package storage

import (
	"log"
	"sort"
	"sync"

	"github.com/Taraxa-project/taraxa-indexer/models"
)

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
	DagCount  uint64 `json:"blk_executed"`
	TrxCount  uint64 `json:"trx_executed"`
	PbftCount uint64 `json:"pbft_size"`
}

func (f1 *FinalizationData) Check(f2 *FinalizationData) {
	// Perform this check only if we are getting data for the same block from node
	if f1.PbftCount != f2.PbftCount {
		return
	}
	if f1.DagCount != f2.DagCount {
		log.Fatal("Dag consistency check failed", f1.DagCount, "!=", f2.DagCount)
	}

	if f1.TrxCount != f2.TrxCount {
		log.Fatal("Transactions consistency check failed ", f1.TrxCount, "!=", f2.TrxCount)
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

	// Sort
	sort.Slice(w.Validators, func(i, j int) bool {
		return w.Validators[i].PbftCount > w.Validators[j].PbftCount
	})
	return w.Validators[from:end], pagination
}

type GenesisHash string
