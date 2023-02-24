package indexer

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/models"
)

type Indexer struct {
	Client  *chain.WsClient
	storage *storage.Storage
}

func NewIndexer(url string, storage *storage.Storage) (i *Indexer, err error) {
	i = new(Indexer)
	i.storage = storage
	i.Client, err = chain.NewWsClient(url)

	return
}

func (i *Indexer) init() {
	remote_hash := storage.GenesisHash(i.Client.GetBlockByNumber(0).Hash)
	if i.storage.GenesisHashExist() {
		local_hash := i.storage.GetGenesisHash()
		fmt.Println("Checking genesis local", local_hash, "remote", remote_hash)
		if local_hash != remote_hash {
			fmt.Println("Genesis hash mismatch. Cleaning DB and restart syncing")
			if err := i.storage.Clean(); err != nil {
				log.Fatal("init storage.Clean() ", err)
			}
			if err := i.storage.SaveGenesisHash(remote_hash); err != nil {
				log.Fatal("init storage.SaveGenesisHash ", err)
			}
		}
	} else {
		if err := i.storage.SaveGenesisHash(remote_hash); err != nil {
			log.Fatal("init storage.SaveGenesisHash ", err)
		}
	}

	if !i.storage.FinalizedPeriodExists() {
		err := i.storage.RecordFinalizedPeriod(1)
		if err != nil {
			log.Fatal("init RecordFinalizedPeriod ", err)
		}
	}
}

func (i *Indexer) sync() {
	start := i.storage.GetFinalizedPeriod()
	end := i.Client.GetLatestPeriod()
	fmt.Println("Starting sync from", start, "to", end)
	prev := time.Now()
	for p := uint64(start); p <= end; p++ {
		blk := i.Client.GetBlockByNumber(p)
		if p%100 == 0 {
			fmt.Println(p, "time:", time.Now().Unix(), "diff", time.Since(prev).Milliseconds(), "ms")
			prev = time.Now()
		}
		err := i.processBlock(blk)
		if err != nil {
			log.Fatal("processBlock", err)
		}
	}
}

func (i *Indexer) Start() {
	i.init()
	i.sync()
	ch, sub, err := i.Client.SubscribeNewHeads()
	if err != nil {
		log.Fatal("Subscription failed")
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case blk := <-ch:
			blk_num := chain.ParseHexInt(blk.Number)
			fmt.Println("Processing event block", blk_num)
			if blk_num != uint64(i.storage.GetFinalizedPeriod())+1 {
				i.sync()
				continue
			}
			err := i.processBlock(blk)
			if err != nil {
				log.Fatal("processBlock", err)
			}
		}
	}
}

func (i *Indexer) processBlock(raw *chain.Block) (err error) {
	block := raw.ToModel()
	transactions := raw.Transactions
	bc := MakeBlockContext(i.storage, block.Age)

	for _, trx_hash := range transactions {
		bc.wg.Add(1)
		go i.ProcessTransaction(bc, trx_hash)
	}

	block_with_dags := i.Client.GetPbftBlockWithDagBlocks(block.Number)
	for _, dag_hash := range block_with_dags.Schedule.DagBlocksOrder {
		bc.wg.Add(1)
		go i.ProcessDag(bc, dag_hash)
	}

	bc.wg.Wait()

	bc.GetAddress(i.storage, block.Author).AddPbft()
	i.SaveAddressStats(bc)
	bc.batch.AddToBatch(block, block.Hash, block.Number)
	bc.batch.RecordFinalizedPeriod(storage.FinalizationData(block.Number))
	bc.batch.CommitBatch()

	return
}

func (i *Indexer) ProcessTransaction(bc *blockContext, hash string) {
	trx := i.Client.GetTransactionByHash(hash)
	bc.SaveTransaction(trx.ToModelWithAge(bc.age))
	bc.wg.Done()
}

func (i *Indexer) ProcessDag(bc *blockContext, hash string) {
	dag := i.Client.GetDagBlockByHash(hash)

	dag_index := bc.GetAddress(i.storage, dag.Sender).AddDag()
	bc.batch.AddToBatch(dag.ToModel(), dag.Sender, dag_index)
	bc.wg.Done()
}

func (i *Indexer) SaveAddressStats(bc *blockContext) {
	for _, stats := range bc.addressStats {
		// json, _ := json.Marshal(stats)
		// fmt.Println("SaveAddressStats", string(json))
		bc.batch.AddToBatch(stats, stats.Address, 0)
	}
}

type blockContext struct {
	wg           sync.WaitGroup
	age          uint64
	addressStats map[string]*storage.AddressStats
	storage      *storage.Storage
	batch        *storage.Batch
	statsMutex   sync.RWMutex
}

func (bc *blockContext) GetAddress(s *storage.Storage, addr string) *storage.AddressStats {
	bc.statsMutex.Lock()
	stats := bc.addressStats[addr]
	if stats != nil {
		bc.statsMutex.Unlock()
		return stats
	}
	bc.addressStats[addr] = storage.MakeEmptyAddressStats(addr)

	bc.addressStats[addr], _ = s.GetAddressStats(addr)
	bc.statsMutex.Unlock()
	return bc.addressStats[addr]
}

func (bc *blockContext) SaveTransaction(trx *models.Transaction) {

	from_index := bc.GetAddress(bc.storage, trx.From).AddTx()
	to_index := bc.GetAddress(bc.storage, trx.To).AddTx()

	bc.batch.AddToBatch(trx, trx.From, from_index)
	bc.batch.AddToBatch(trx, trx.To, to_index)
}

func MakeBlockContext(s *storage.Storage, age uint64) *blockContext {
	var bc blockContext
	bc.age = age
	bc.addressStats = make(map[string]*storage.AddressStats)
	bc.batch = s.NewBatch()
	bc.storage = s
	return &bc
}
