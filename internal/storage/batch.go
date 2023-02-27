package storage

import (
	"log"
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

func (b *Batch) CommitBatch() {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	err := b.Commit(pebble.NoSync)
	if err != nil {
		log.Fatal("CommitBatch ", err)
	}
}

func (b *Batch) SaveFinalizedPeriod(f FinalizationData) {
	err := b.addToBatch(&f, []byte(getPrefix(&f)))
	if err != nil {
		log.Fatal("SaveFinalizedPeriod ", err)
	}
}

func (b *Batch) SaveGenesisHash(h GenesisHash) {
	err := b.addToBatch(&h, []byte(getPrefix(&h)))
	if err != nil {
		log.Fatal("SaveGenesisHash ", err)
	}
}

func (b *Batch) AddToBatch(o interface{}, key1 string, key2 uint64) {
	err := b.addToBatch(o, getKey(getPrefix(o), key1, key2))
	if err != nil {
		log.Fatal("AddToBatch ", err)
	}
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
