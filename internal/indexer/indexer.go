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
	if !i.storage.FinalizedPeriodExists() {
		i.storage.RecordFinalizedPeriod(1)
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
			fmt.Println(p, "time:", time.Now().Unix(), "diff", time.Now().Sub(prev).Milliseconds(), "ms")
			prev = time.Now()
		}
		i.processBlock(blk)
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
		case block := <-ch:
			fmt.Println("Processing event block", block.ToModel().Number)
			if chain.ParseHexInt(block.Number) != uint64(i.storage.GetFinalizedPeriod())+1 {
				i.sync()
				continue
			}
			i.processBlock(block)
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

func (i *Indexer) ProcessTransaction(bc *blockContext, hash string) (err error) {
	trx := i.Client.GetTransactionByHash(hash)
	bc.SaveTransaction(trx.ToModelWithAge(bc.age))
	bc.wg.Done()
	return
}

func (i *Indexer) ProcessDag(bc *blockContext, hash string) (err error) {
	dag := i.Client.GetDagBlockByHash(hash)

	dag_index := bc.GetAddress(i.storage, dag.Sender).AddDag()
	err = bc.batch.AddToBatch(dag.ToModel(), dag.Sender, dag_index)

	bc.wg.Done()
	return
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

	s.GetFromDB(bc.addressStats[addr], addr)

	bc.statsMutex.Unlock()
	return bc.addressStats[addr]
}

func (bc *blockContext) SaveTransaction(trx *models.Transaction) {

	from_index := bc.GetAddress(bc.storage, trx.From).AddTx()
	to_index := bc.GetAddress(bc.storage, trx.To).AddTx()

	err1 := bc.batch.AddToBatch(trx, trx.From, from_index)
	err2 := bc.batch.AddToBatch(trx, trx.To, to_index)
	if err1 != nil || err2 != nil {
		log.Fatal("Something wrong saving transaction to DB")
	}
}

func MakeBlockContext(s *storage.Storage, age uint64) *blockContext {
	var bc blockContext
	bc.age = age
	bc.addressStats = make(map[string]*storage.AddressStats)
	bc.batch = s.NewBatch()
	bc.storage = s
	return &bc
}
