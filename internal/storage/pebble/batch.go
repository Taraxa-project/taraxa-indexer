package pebble

import (
	"encoding/json"
	"sync"

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
	err := b.AddToBatchFullKey(s, []byte(GetPrefix((*storage.TotalSupply)(s))))
	if err != nil {
		log.WithError(err).Fatal("SetTotalSupply failed")
	}
}

func (b *Batch) SetFinalizationData(f *storage.FinalizationData) {
	err := b.AddToBatchFullKey(f, []byte(GetPrefix(f)))
	if err != nil {
		log.WithError(err).Fatal("SetFinalizationData failed")
	}
}

func (b *Batch) SetGenesisHash(h storage.GenesisHash) {
	err := b.AddToBatchFullKey(&h, []byte(GetPrefix(&h)))
	if err != nil {
		log.WithError(err).Fatal("SetGenesisHash failed")
	}
}

func (b *Batch) UpdateWeekStats(w storage.WeekStats) {
	err := b.AddToBatchFullKey(&w, w.Key)
	if err != nil {
		log.WithError(err).Fatal("UpdateWeekStats failed")
	}
}

func (b *Batch) SaveAccounts(a storage.Accounts) {
	a.SortByBalanceDescending()
	b.AddToBatchSingleKey(a, "")
}

func (b *Batch) AddToBatch(o interface{}, key1 string, key2 uint64) {
	err := b.AddToBatchFullKey(o, getKey(GetPrefix(o), key1, key2))
	if err != nil {
		log.WithError(err).Fatal("AddToBatch failed")
	}
}

func (b *Batch) AddToBatchSingleKey(o interface{}, key string) {
	err := b.AddToBatchFullKey(o, GetPrefixKey(GetPrefix(o), key))
	if err != nil {
		obj_json, _ := json.Marshal(o)
		log.WithError(err).WithField("object", string(obj_json)).WithField("key", string(GetPrefixKey(GetPrefix(o), key))).Fatal("AddToBatchSingleKey failed")
	}
}

func (b *Batch) AddToBatchFullKey(o interface{}, key []byte) error {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()

	data, err := rlp.EncodeToBytes(o)
	if err != nil {
		return err
	}
	return b.Set(key, data, nil)
}

func (b *Batch) Remove(key []byte) {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()

	err := b.Delete(key, nil)
	if err != nil {
		log.WithError(err).Fatal("Remove failed")
	}
}
