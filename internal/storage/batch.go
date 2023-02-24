package storage

import (
	"github.com/cockroachdb/pebble"
	"github.com/ethereum/go-ethereum/rlp"
)

type Batch struct {
	*pebble.Batch
}

func (s *Storage) NewBatch() *Batch {
	return &Batch{s.db.NewBatch()}
}

func (b *Batch) CommitBatch() {
	b.Commit(pebble.NoSync)
}

func (s *Batch) RecordFinalizedPeriod(f FinalizationData) error {
	return s.addToBatch(&f, []byte(getPrefix(&f)))
}

func (b *Batch) AddToBatch(o interface{}, key1 string, key2 uint64) error {
	return b.addToBatch(o, getKey(getPrefix(o), key1, key2))
}

func (b *Batch) addToBatch(o interface{}, key []byte) error {
	data, err := rlp.EncodeToBytes(o)
	if err != nil {
		return err
	}
	return b.Set(key, data, nil)
}
