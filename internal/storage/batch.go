package storage

type Batch interface {
	CommitBatch()
	SetTotalSupply(s *TotalSupply)
	SetFinalizationData(f *FinalizationData)
	SetGenesisHash(h GenesisHash)
	UpdateWeekStats(w WeekStats)
	SaveAccounts(a Accounts)
	Add(o interface{}, key1 string, key2 uint64)
	AddSerialized(o interface{}, data []byte, key1 string, key2 uint64)
	AddSingleKey(o interface{}, key string)
	AddSerializedSingleKey(o interface{}, data []byte, key string)
	AddWithKey(o interface{}, key []byte) error
	AddSerializedWithKey(o interface{}, data, key []byte) error
	Remove(key []byte)
}
