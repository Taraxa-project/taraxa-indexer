package indexer

import (
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	log "github.com/sirupsen/logrus"
	"github.com/spiretechnology/go-pool"
)

type SyncQueue struct {
	client     common.Client
	blocks     sync.Map
	current    atomic.Uint64
	latest     atomic.Uint64
	end        uint64
	queueLimit uint64
	tp         pool.Pool
}

func MakeSyncQueue(start, end, queueLimit uint64, client common.Client) *SyncQueue {
	sq := new(SyncQueue)
	sq.client = client
	sq.blocks = sync.Map{}
	sq.current.Store(start)
	sq.latest.Store(start)
	sq.end = end
	sq.queueLimit = queueLimit
	sq.tp = common.MakeThreadPool()
	return sq
}

func (sq *SyncQueue) PopNext() *chain.BlockData {
	if (sq.latest.Load() - sq.current.Load()) == 0 {
		return nil
	}
	ret, ok := sq.blocks.LoadAndDelete(sq.current.Load())
	if !ok {
		return nil
	}
	sq.current.Add(1)
	if ret != nil {
		return ret.(*chain.BlockData)
	} else {
		return nil
	}
}

func (sq *SyncQueue) Push(bd *chain.BlockData) {
	log.WithField("number", bd.Pbft.Number).Debug("Pushing block")
	sq.blocks.Store(bd.Pbft.Number, bd)
}

func (sq *SyncQueue) GetCurrent() uint64 {
	return sq.current.Load()
}

func (sq *SyncQueue) Start() {
	for {
		latest := sq.latest.Load()
		current := sq.current.Load()
		if latest-current >= sq.queueLimit {
			time.Sleep(1 * time.Millisecond)
			continue
		}
		if latest > sq.end {
			break
		}
		sq.latest.Add(1)
		go sq.RequestBlock(latest)
	}
}

func (sq *SyncQueue) RequestBlock(number uint64) {
	bd, err := chain.GetBlockData(sq.client, number)
	// check for Requested blk num substring

	if err != nil {
		if strings.Contains(err.Error(), "Requested blk num") {
			log.WithField("number", number).Info("Stop syncing")
			return
		}
		log.WithError(err).Fatal("Failed to get block data")
	}
	sq.Push(bd)
}
