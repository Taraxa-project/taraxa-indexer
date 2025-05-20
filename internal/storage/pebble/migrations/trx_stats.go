package migration

import (
	"strconv"
	"sync"
	"time"

	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage/pebble"
	"github.com/ethereum/go-ethereum/rlp"
	log "github.com/sirupsen/logrus"
	"github.com/spiretechnology/go-pool"
)

type BlockQueue struct {
	blocks   map[uint64]*common.Block
	client   common.Client
	finished bool
	tp       pool.Pool
	mu       sync.Mutex
}

func MakeBlockQueue(client common.Client) *BlockQueue {
	return &BlockQueue{
		blocks: make(map[uint64]*common.Block),
		client: client,
	}
}

func (q *BlockQueue) Add(block *common.Block) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.blocks[block.Number] = block
}

func (q *BlockQueue) GetAndRemove(number uint64) *common.Block {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.blocks[number] == nil {
		return nil
	}
	block := *q.blocks[number]
	delete(q.blocks, number)
	return &block
}

func (q *BlockQueue) Start(block_num, earlisest_block_ts uint64) {
	batch_size := uint64(100)
	q.tp = common.MakeThreadPool()
	q.tp.Go(func() {
		block_to_get := block_num
		for {
			current_block, err := q.client.GetBlocks(block_to_get-batch_size, block_to_get)
			if err != nil {
				log.WithField("error", err).Fatal("Error getting block")
			}
			if current_block[0].Timestamp < earlisest_block_ts {
				q.finished = true
				break
			}
			for _, block := range current_block {
				q.Add(block)
			}
			block_to_get -= batch_size
		}
	})
}

type TrxStats struct {
	client common.Client
}

func (m *TrxStats) GetId() string {
	return "trx_stats"
}

func (m *TrxStats) Init(client common.Client) {
	m.client = client
}

func commitDayStats(s *pebble.Storage, day_stats *storage.DayStatsWithTimestamp) {
	log.WithField("stats", day_stats).Info("Committing day stats")
	b := s.NewBatch()
	b.AddSingleKey(day_stats.TrxGasStats, strconv.FormatUint(day_stats.Timestamp, 10))
	b.CommitBatch()
}

func estimateBlockNumber(last_block *common.Block, day_stats *storage.DayStatsWithTimestamp) uint64 {
	block_interval := 3.5
	return uint64(int64(last_block.Number) - int64(float64(int64(last_block.Timestamp)-int64(day_stats.Timestamp))/block_interval))
}

func findStart(last_block *common.Block, c common.Client, s *pebble.Storage) (day_stats *storage.DayStatsWithTimestamp, block_num uint64) {
	s.ForEach(storage.TrxGasStats{}, "", nil, func(key []byte, res []byte) (stop bool) {
		db_stats := storage.TrxGasStats{}
		err := rlp.DecodeBytes(res, &db_stats)
		if err != nil {
			log.WithField("error", err).Fatal("Error unmarshalling day stats")
		}

		log.WithFields(log.Fields{
			"key":      string(key),
			"db_stats": db_stats,
		}).Info("DB stats")
		ts := storage.GetTimestampFromKey(key)
		if day_stats == nil || ts < day_stats.Timestamp {
			day_stats = storage.MakeDayStatsWithTimestamp(ts)
		}
		return false
	})

	if day_stats == nil {
		day_stats = storage.MakeDayStatsWithTimestamp(common.DayStart(last_block.Timestamp))
		block_num = last_block.Number
		return
	}

	block_num = estimateBlockNumber(last_block, day_stats)
	for {
		block, err := c.GetBlockByNumber(block_num)
		if err != nil {
			log.WithField("error", err).Fatal("Error getting block")
		}
		if block.Timestamp <= day_stats.Timestamp {
			next_block, err := c.GetBlockByNumber(block_num + 1)
			if err != nil {
				log.WithField("error", err).Fatal("Error getting block")
			}
			if next_block.Timestamp >= day_stats.Timestamp {
				break
			}
		}
		block_num = estimateBlockNumber(block, day_stats)
	}
	day_stats.Timestamp -= common.Day
	return day_stats, block_num
}

func (m *TrxStats) Apply(s *pebble.Storage) error {
	finalized := s.GetFinalizationData()
	if finalized == nil {
		return nil
	}

	current_block_num := finalized.PbftCount
	current_block, err := m.client.GetBlockByNumber(current_block_num)
	if err != nil {
		return err
	}
	first_day := common.DayStart(current_block.Timestamp) - common.Days30
	day_stats, current_block_num := findStart(current_block, m.client, s)
	if day_stats.Timestamp <= first_day {
		return nil
	}

	block_queue := MakeBlockQueue(m.client)
	block_queue.Start(current_block_num, first_day)

	for {
		block := block_queue.GetAndRemove(current_block_num)
		if block == nil {
			time.Sleep(10 * time.Millisecond)
			continue
		}
		current_day_start := common.DayStart(block.Timestamp)

		if current_day_start < first_day {
			break
		}
		if current_day_start < day_stats.Timestamp {
			commitDayStats(s, day_stats)
			day_stats = storage.MakeDayStatsWithTimestamp(current_day_start)
		}

		day_stats.AddBlock(block)
		current_block_num--
	}
	return nil
}
