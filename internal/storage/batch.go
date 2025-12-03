package storage

import (
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
)

type Batch interface {
	CommitBatch()
	SetTotalSupply(s *TotalSupply)
	SetFinalizationData(f *common.FinalizationData)
	SetGenesisHash(h GenesisHash)
	UpdateWeekStats(w WeekStats)
	SaveHoldersLeaderboard(a Accounts)
	AddDailyContractUsers(address string, timestamp uint64, users *DailyContractUsers)
	AddDayStats(d *DayStatsWithTimestamp)
	AddYieldSaving(period, timestamp uint64)
	AddLambda(lambdaMs uint64)

	Add(o any, key1 string, key2 uint64)
	AddSerialized(o any, data []byte, key1 string, key2 uint64)
	AddSingleKey(o any, key string)
	AddSerializedSingleKey(o any, data []byte, key string)
	AddWithKey(o any, key []byte) error
	AddSerializedWithKey(o any, data, key []byte) error
	Remove(key []byte)
}
