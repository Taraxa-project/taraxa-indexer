package storage

import "sync"

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

type FinalizationData uint64

type GenesisHash string
