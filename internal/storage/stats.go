package storage

import (
	"math/big"
	"strings"

	"github.com/Taraxa-project/taraxa-indexer/internal/common"
)

type TrxGasStats struct {
	TrxCount uint64   `json:"trxCount"`
	GasUsed  *big.Int `json:"gasUsed"`
}

func EmptyTrxGasStats() TrxGasStats {
	return TrxGasStats{
		TrxCount: 0,
		GasUsed:  big.NewInt(0),
	}
}

func (d *TrxGasStats) Add(other TrxGasStats) {
	d.TrxCount += other.TrxCount
	d.GasUsed.Add(d.GasUsed, other.GasUsed)
}

type DayStatsWithTimestamp struct {
	TrxGasStats
	Timestamp uint64 `json:"timestamp"`
}

func (d *DayStatsWithTimestamp) AddBlock(blk *common.Block) {
	day_start := common.DayStart(blk.Timestamp)
	if day_start > d.Timestamp {
		*d = *MakeDayStatsWithTimestamp(day_start)
	}
	d.TrxCount += blk.TransactionCount
	d.GasUsed.Add(d.GasUsed, blk.GasUsed)
}
func MakeDayStatsWithTimestamp(ts uint64) *DayStatsWithTimestamp {
	return &DayStatsWithTimestamp{
		TrxGasStats: EmptyTrxGasStats(),
		Timestamp:   ts,
	}
}

func GetTimestampFromKey(key []byte) uint64 {
	ts := strings.Split(string(key), "|")
	return common.ParseUInt(strings.TrimLeft(ts[1], "0"))
}
