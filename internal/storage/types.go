package storage

import (
	"log"
	"sync"
)

// AddressStats defines the model for an address aggregate.
type AddressStats struct {
	Address   string       `json:"address"`
	TxTotal   uint64       `json:"txTotal"`
	DagTotal  uint64       `json:"dagTotal"`
	PbftTotal uint64       `json:"pbftTotal"`
	mutex     sync.RWMutex `rlp:"-"`
}

func (a *AddressStats) AddTx() uint64 {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.TxTotal++
	return a.TxTotal
}

func (a *AddressStats) AddPbft() uint64 {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.PbftTotal++
	return a.PbftTotal
}

func (a *AddressStats) AddDag() uint64 {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.DagTotal++
	return a.DagTotal
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
	if a.Address == b.Address && a.TxTotal == b.TxTotal && a.DagTotal == b.DagTotal && a.PbftTotal == b.PbftTotal {
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

type GenesisHash string
