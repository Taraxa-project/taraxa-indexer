package indexer

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	log "github.com/sirupsen/logrus"
	"github.com/spiretechnology/go-pool"
)

type SyncQueue struct {
	client     chain.Client
	blocks     sync.Map
	current    atomic.Uint64
	latest     atomic.Uint64
	queueLimit uint64
	tp         pool.Pool
}

func MakeSyncQueue(start, queueLimit uint64, client chain.Client) *SyncQueue {
	sq := new(SyncQueue)
	sq.current.Store(start)
	sq.latest.Store(start)
	sq.queueLimit = queueLimit
	sq.client = client
	sq.blocks = sync.Map{}
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
			log.WithFields(log.Fields{"latest": latest, "current": current}).Debug("Syncing: queue limit reached")
			time.Sleep(1 * time.Millisecond)
			continue
		}
		sq.latest.Add(1)
		go func(toRequest uint64) {
			bd, err := chain.GetBlockData(sq.client, toRequest)
			if err == chain.ErrFutureBlock {
				log.WithField("number", toRequest).Debug("Stop syncing")
				return
			} else if err != nil {
				log.WithError(err).Fatal("Failed to get block data")
			}
			sq.Push(bd)
		}(latest)
	}
}
