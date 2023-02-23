package indexer

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

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
	for p := uint64(start); p <= end; p++ {
		blk := i.Client.GetBlockByNumber(p)
		fmt.Println("Processing", p, blk.Number)
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
			i.processBlock(block)
		}
	}
}

func (i *Indexer) processBlock(raw *chain.Block) (err error) {
	block := raw.ToModel()
	transactions := raw.Transactions
	bc := MakeBlockContext(block.Age)

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

	i.SaveAddressStats(bc)
	bc.GetAddress(i.storage, block.Author).AddPbft()
	i.storage.AddToDB(block, block.Hash, block.Number)
	i.storage.RecordFinalizedPeriod(storage.FinalizationData(block.Number))
	return
}

func (i *Indexer) ProcessTransaction(bc *blockContext, hash string) (err error) {
	trx := i.Client.GetTransactionByHash(hash)
	bc.SaveTransaction(i.storage, trx.ToModelWithAge(bc.age))
	bc.wg.Done()
	return
}

func (i *Indexer) ProcessDag(bc *blockContext, hash string) (err error) {
	dag := i.Client.GetDagBlockByHash(hash)

	dag_index := bc.GetAddress(i.storage, dag.Sender).AddDag()
	err = i.storage.AddToDB(dag.ToModel(), dag.Sender, dag_index)

	bc.wg.Done()
	return
}

func (i *Indexer) SaveAddressStats(bc *blockContext) {
	for _, stats := range bc.addressStats {
		json, _ := json.Marshal(stats)
		fmt.Println("SaveAddressStats", string(json))
		i.storage.AddToDB(stats, stats.Address, 0)
	}
}

type blockContext struct {
	wg           sync.WaitGroup
	age          uint64
	addressStats map[string]*storage.AddressStats
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

func (bc *blockContext) SaveTransaction(s *storage.Storage, trx *models.Transaction) {

	from_index := bc.GetAddress(s, trx.From).AddTx()
	to_index := bc.GetAddress(s, trx.To).AddTx()

	err1 := s.AddToDB(trx, trx.From, from_index)
	err2 := s.AddToDB(trx, trx.To, to_index)
	if err1 != nil || err2 != nil {
		log.Fatal("Something wrong saving transaction to DB")
	}
}

func MakeBlockContext(age uint64) *blockContext {
	var bc blockContext
	bc.age = age
	bc.addressStats = make(map[string]*storage.AddressStats)
	return &bc
}
