package storage

import "sync"

// AddressStats defines the model for an address aggregate.
type AddressStats struct {
	Address   string `json:"address"`
	TxTotal   uint64 `json:"txTotal"`
	DagTotal  uint64 `json:"dagTotal"`
	PbftTotal uint64 `json:"pbftTotal"`
	mutex     sync.RWMutex
}

func (a *AddressStats) AddTx() uint64 {
	a.mutex.Lock()
	a.TxTotal++
	a.mutex.Unlock()
	return a.TxTotal
}

func (a *AddressStats) AddPbft() {
	a.mutex.Lock()
	a.PbftTotal++
	a.mutex.Unlock()
}

func (a *AddressStats) AddDag() uint64 {
	a.mutex.Lock()
	a.DagTotal++
	a.mutex.Unlock()
	return a.DagTotal
}

func MakeEmptyAddressStats(addr string) *AddressStats {
	data := new(AddressStats)
	data.Address = addr
	return data
}

type FinalizationData uint64
