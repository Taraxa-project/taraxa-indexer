package indexer

import (
	"time"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	log "github.com/sirupsen/logrus"
)

type Indexer struct {
	client  *chain.WsClient
	storage *storage.Storage
}

func NewIndexer(url string, storage *storage.Storage) (i *Indexer, err error) {
	i = new(Indexer)
	i.storage = storage
	i.client, err = chain.NewWsClient(url)
	return
}

func (i *Indexer) init() {
	remote_hash := storage.GenesisHash(i.client.GetBlockByNumber(0).Hash)
	db_clean := false
	if i.storage.GenesisHashExist() {
		local_hash := i.storage.GetGenesisHash()
		if local_hash != remote_hash {
			log.WithFields(log.Fields{"local_hash": local_hash, "remote_hash": remote_hash}).Warn("Genesis changed, reseting")
			if err := i.storage.Clean(); err != nil {
				log.WithField("error", err).Warn("Error during storage cleaning")
			}
			db_clean = true
		}
	} else {
		db_clean = true
	}
	if !db_clean {
		return
	}

	genesis := MakeGenesis(i.storage, i.client, remote_hash)
	// Genesis hash and finalized period(0) is set inside
	genesis.process()
}

func (i *Indexer) sync() {
	// start processing blocks from the next one
	start := i.storage.GetFinalizationData().PbftCount + 1
	end := i.client.GetLatestPeriod()
	log.WithFields(log.Fields{"start": start, "end": end}).Info("Start syncing")
	prev := time.Now()
	for p := uint64(start); p <= end; p++ {
		blk := i.client.GetBlockByNumber(p)
		if p%100 == 0 {
			log.WithFields(log.Fields{"period": p, "elapsed_ms": time.Since(prev).Milliseconds()}).Info("Syncing: block applied")
			prev = time.Now()
		}

		MakeBlockContext(i.storage, i.client).process(blk)
	}
}

func (i *Indexer) Start() {
	i.init()
	i.sync()
	ch, sub := i.client.SubscribeNewHeads()
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case sub_blk := <-ch:
			p := chain.ParseInt(sub_blk.Number)
			log.WithFields(log.Fields{"period": p}).Info("Applying block from subscription channel")
			if p != i.storage.GetFinalizationData().PbftCount+1 {
				i.sync()
				continue
			}
			// We need to get block from API one more time because chain isn't returning transactions in this subscription object
			blk := i.client.GetBlockByNumber(p)
			MakeBlockContext(i.storage, i.client).process(blk)
		}
	}
}
