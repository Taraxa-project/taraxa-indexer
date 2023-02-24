package storage

import (
	"sync"

	"github.com/cockroachdb/pebble"
	"github.com/ethereum/go-ethereum/rlp"
)

type Batch struct {
	*pebble.Batch
	mutex *sync.RWMutex
}

func (s *Storage) NewBatch() *Batch {
	return &Batch{s.db.NewBatch(), new(sync.RWMutex)}
}

func (b *Batch) CommitBatch() error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	return b.Commit(pebble.NoSync)
}

func (b *Batch) RecordFinalizedPeriod(f FinalizationData) error {
	return b.addToBatch(&f, []byte(getPrefix(&f)))
}

func (b *Batch) AddToBatch(o interface{}, key1 string, key2 uint64) error {
	return b.addToBatch(o, getKey(getPrefix(o), key1, key2))
}

func (b *Batch) addToBatch(o interface{}, key []byte) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	data, err := rlp.EncodeToBytes(o)
	if err != nil {
		return err
	}
	return b.Set(key, data, nil)
}
