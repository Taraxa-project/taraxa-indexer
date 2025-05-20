package storage

import "github.com/Taraxa-project/taraxa-indexer/internal/common"

type Batch interface {
	CommitBatch()
	SetTotalSupply(s *TotalSupply)
	SetFinalizationData(f *common.FinalizationData)
	SetGenesisHash(h GenesisHash)
	UpdateWeekStats(w WeekStats)
	SaveAccounts(a *AccountsMap)
	Add(o any, key1 string, key2 uint64)
	AddSerialized(o any, data []byte, key1 string, key2 uint64)
	AddSingleKey(o any, key string)
	AddSerializedSingleKey(o any, data []byte, key string)
	AddWithKey(o any, key []byte) error
	AddSerializedWithKey(o any, data, key []byte) error
	Remove(key []byte)
}
