package storage

type Batch interface {
	CommitBatch()
	SetTotalSupply(s *TotalSupply)
	SetFinalizationData(f *FinalizationData)
	SetGenesisHash(h GenesisHash)
	UpdateWeekStats(w WeekStats)
	AddToBatch(o interface{}, key1 string, key2 uint64)
	AddToBatchSingleKey(o interface{}, key string)
	Remove(key string)
}
