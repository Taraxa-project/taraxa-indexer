package pebble

import (
	"sync"

	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/cockroachdb/pebble"
	"github.com/ethereum/go-ethereum/rlp"
	log "github.com/sirupsen/logrus"
)

type Batch struct {
	*pebble.Batch
	Mutex *sync.RWMutex
}

func (b *Batch) CommitBatch() {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()

	err := b.Commit(pebble.NoSync)
	if err != nil {
		log.WithError(err).Fatal("CommitBatch failed")
	}
}

func (b *Batch) SetTotalSupply(s *storage.TotalSupply) {
	err := b.AddWithKey(s, []byte(GetPrefix((*storage.TotalSupply)(s))))
	if err != nil {
		log.WithError(err).Fatal("SetTotalSupply failed")
	}
}

func (b *Batch) SetFinalizationData(f *common.FinalizationData) {
	err := b.AddWithKey(f, []byte(GetPrefix(f)))
	if err != nil {
		log.WithError(err).Fatal("SetFinalizationData failed")
	}
}

func (b *Batch) SetGenesisHash(h storage.GenesisHash) {
	err := b.AddWithKey(&h, []byte(GetPrefix(&h)))
	if err != nil {
		log.WithError(err).Fatal("SetGenesisHash failed")
	}
}

func (b *Batch) UpdateWeekStats(w storage.WeekStats) {
	err := b.AddWithKey(&w, w.Key)
	if err != nil {
		log.WithError(err).Fatal("UpdateWeekStats failed")
	}
}

func (b *Batch) SaveAccounts(a *storage.AccountsMap) {
	b.AddSingleKey(a.SortedSlice(), "")
}

func (b *Batch) AddDayStats(d *storage.DayStatsWithTimestamp) {
	err := b.AddWithKey(d.TrxGasStats, getKey(GetPrefix(d.TrxGasStats), "", d.Timestamp))
	if err != nil {
		log.WithError(err).Fatal("AddDayStats failed")
	}
}

func (b *Batch) Add(o any, key1 string, key2 uint64) {
	err := b.AddWithKey(o, getKey(GetPrefix(o), key1, key2))
	if err != nil {
		log.WithError(err).WithFields(log.Fields{"prefix": GetPrefix(o), "key1": key1, "key2": key2}).Fatal("Batch.Add failed")
	}
}

func (b *Batch) AddSerialized(o any, data []byte, key1 string, key2 uint64) {
	err := b.AddSerializedWithKey(o, data, getKey(GetPrefix(o), key1, key2))
	if err != nil {
		log.WithError(err).WithFields(log.Fields{"object": o, "prefix": GetPrefix(o), "key1": key1, "key2": key2}).Fatal("Batch.AddSerialized failed")
	}
}

func (b *Batch) AddSerializedSingleKey(o any, data []byte, key string) {
	err := b.AddSerializedWithKey(o, data, GetPrefixKey(GetPrefix(o), key))
	if err != nil {
		log.WithError(err).WithFields(log.Fields{"object": o, "prefix": GetPrefix(o), "key": GetPrefixKey(GetPrefix(o), key)}).Fatal("Batch.AddSerializedSingleKey failed")
	}
}

func (b *Batch) AddSingleKey(o any, key string) {
	err := b.AddWithKey(o, GetPrefixKey(GetPrefix(o), key))
	if err != nil {
		log.WithError(err).WithFields(log.Fields{"object": o, "prefix": GetPrefix(o), "key": GetPrefixKey(GetPrefix(o), key)}).Fatal("Batch.AddSingleKey failed")
	}
}

func (b *Batch) AddSerializedWithKey(o any, data, key []byte) error {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()

	return b.Set(key, data, nil)
}

func (b *Batch) AddWithKey(o any, key []byte) error {
	data, err := rlp.EncodeToBytes(o)
	if err != nil {
		return err
	}
	return b.AddSerializedWithKey(o, data, key)
}

func (b *Batch) Remove(key []byte) {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()

	err := b.Delete(key, nil)
	if err != nil {
		log.WithError(err).Fatal("Remove failed")
	}
}
