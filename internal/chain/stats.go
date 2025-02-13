package chain

import (
	"github.com/Taraxa-project/taraxa-indexer/models"
)

type Stats struct {
	blocks        []models.Pbft
	interval      int
	totalTrxCount uint64
	blockTimeDiff uint64
	stats         *models.ChainStats
}

func MakeStats(interval int) *Stats {
	return &Stats{
		blocks:   make([]models.Pbft, 0, interval+1),
		interval: interval,
	}
}

func (s *Stats) AddPbft(lastPbft *models.Pbft) {
	s.blocks = append(s.blocks, *lastPbft)
	block_len := len(s.blocks)
	if block_len != s.interval+1 {
		return
	}
	if s.stats == nil {
		s.stats = new(models.ChainStats)
		// lowest key in map is the start block
		s.stats.StartBlock = s.blocks[0].Number
	}

	s.totalTrxCount -= s.blocks[0].TransactionCount
	s.blocks = s.blocks[1:]

	firstPbft := s.blocks[0]
	s.stats.StartBlock = firstPbft.Number

	s.totalTrxCount += lastPbft.TransactionCount

	s.blockTimeDiff = lastPbft.Timestamp - firstPbft.Timestamp
	s.stats.EndBlock = lastPbft.Number

	s.stats.BlockInterval = float32(s.blockTimeDiff) / float32(s.interval)
	s.stats.Tps = float32(s.totalTrxCount) / float32(s.blockTimeDiff)
}

func (s *Stats) GetStats() *models.ChainStats {
	return s.stats
}
