package storage

import (
	"sync"

	"github.com/cockroachdb/pebble"
	"github.com/ethereum/go-ethereum/rlp"
	log "github.com/sirupsen/logrus"
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
		log.WithError(err).Fatal("CommitBatch failed")
	}
}

func (b *Batch) SaveFinalizedPeriod(f *FinalizationData) {
	err := b.addToBatch(f, []byte(getPrefix(f)))
	if err != nil {
		log.WithError(err).Fatal("SaveFinalizedPeriod failed")
	}
}

func (b *Batch) SaveGenesisHash(h GenesisHash) {
	err := b.addToBatch(&h, []byte(getPrefix(&h)))
	if err != nil {
		log.WithError(err).Fatal("SaveGenesisHashSaveGenesisHash failed")
	}
}

func (b *Batch) UpdateWeekStats(w WeekStats) {
	err := b.addToBatch(&w, w.key)
	if err != nil {
		log.WithError(err).Fatal("UpdateWeekStats failed")
	}
}

func (b *Batch) AddToBatch(o interface{}, key1 string, key2 uint64) {
	err := b.addToBatch(o, getKey(getPrefix(o), key1, key2))
	if err != nil {
		log.WithError(err).Fatal("AddToBatch failed")
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
