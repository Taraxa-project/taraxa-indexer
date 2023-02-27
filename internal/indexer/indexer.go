package indexer

import (
	"fmt"
	"log"
	"time"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
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
		fmt.Println("Checking genesis local", local_hash, "remote", remote_hash)
		if local_hash != remote_hash {
			fmt.Println("Genesis hash mismatch. Cleaning DB and restart syncing")
			if err := i.storage.Clean(); err != nil {
				log.Fatal("init storage.Clean() ", err)
			}
			db_clean = true
		}
	} else {
		db_clean = true
	}
	if !db_clean {
		return
	}

	genesis := MakeGenesis(i.storage, i.client, string(remote_hash))
	// Genesis hash and finalized period(0) is set inside
	genesis.process()
}

func (i *Indexer) sync() {
	// start processing blocks from the next one
	start := i.storage.GetFinalizedPeriod() + 1
	end := i.client.GetLatestPeriod()
	fmt.Println("Starting sync from", start, "to", end)
	prev := time.Now()
	for p := uint64(start); p <= end; p++ {
		blk := i.client.GetBlockByNumber(p)
		if p%100 == 0 {
			fmt.Println(p, "elapsed", time.Since(prev).Milliseconds(), "ms")
			prev = time.Now()
		}
		err := MakeBlockContext(i.storage, i.client).process(blk)
		if err != nil {
			log.Fatal("processBlock", err)
		}
	}
}

func (i *Indexer) Start() {
	i.init()
	i.sync()
	ch, sub, err := i.client.SubscribeNewHeads()
	if err != nil {
		log.Fatal("Subscription failed")
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case blk := <-ch:
			blk_num := chain.ParseInt(blk.Number)
			fmt.Println("Processing event block", blk_num)
			if blk_num != uint64(i.storage.GetFinalizedPeriod())+1 {
				i.sync()
				continue
			}
			err = MakeBlockContext(i.storage, i.client).process(blk)
			if err != nil {
				log.Fatal("processBlock", err)
			}
		}
	}
}
