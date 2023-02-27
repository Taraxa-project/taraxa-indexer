package indexer

import (
	"strings"
	"sync"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/models"
)

type blockContext struct {
	storage      *storage.Storage
	batch        *storage.Batch
	client       *chain.WsClient
	wg           sync.WaitGroup
	age          uint64
	addressStats map[string]*storage.AddressStats
	finalized    *storage.FinalizationData
	statsMutex   sync.RWMutex
}

func MakeBlockContext(s *storage.Storage, client *chain.WsClient) *blockContext {
	var bc blockContext
	bc.storage = s
	bc.batch = s.NewBatch()
	bc.client = client
	bc.addressStats = make(map[string]*storage.AddressStats)
	if s.FinalizedPeriodExists() {
		bc.finalized = s.GetFinalizedPeriod()
	} else {
		bc.finalized = new(storage.FinalizationData)
	}
	return &bc
}

func (bc *blockContext) commit(period uint64) {
	bc.batch.SaveFinalizedPeriod(bc.finalized)
	bc.addAddressStatsToBatch()
	bc.batch.CommitBatch()
}

func (bc *blockContext) process(raw *chain.Block) (err error) {
	block := raw.ToModel()
	transactions := &raw.Transactions

	bc.finalized.TrxCount += block.TransactionCount
	for _, trx_hash := range *transactions {
		bc.wg.Add(1)
		go bc.processTransaction(trx_hash)
	}

	block_with_dags := bc.client.GetPbftBlockWithDagBlocks(block.Number)
	bc.finalized.DagCount += uint64(len(block_with_dags.Schedule.DagBlocksOrder))
	for _, dag_hash := range block_with_dags.Schedule.DagBlocksOrder {
		bc.wg.Add(1)
		go bc.processDag(dag_hash)
	}

	bc.wg.Wait()

	bc.finalized.PbftCount++
	author_pbft_index := bc.getAddress(bc.storage, block.Author).AddPbft()
	bc.batch.AddToBatch(block, block.Author, author_pbft_index)

	// If stats is available check for consistency
	stats, stats_err := bc.client.GetNodeStats()
	if stats_err == nil {
		stats.Check(bc.finalized)
	}
	bc.commit(block.Number)

	return
}

func (bc *blockContext) processTransaction(hash string) {
	trx := bc.client.GetTransactionByHash(hash)
	bc.SaveTransaction(trx.ToModelWithTimestamp(bc.age))
	bc.wg.Done()
}

func (bc *blockContext) processDag(hash string) {
	dag := bc.client.GetDagBlockByHash(hash)

	dag_index := bc.getAddress(bc.storage, dag.Sender).AddDag()
	bc.batch.AddToBatch(dag.ToModel(), dag.Sender, dag_index)
	bc.wg.Done()
}

func (bc *blockContext) addAddressStatsToBatch() {
	for _, stats := range bc.addressStats {
		bc.batch.AddToBatch(stats, stats.Address, 0)
	}
}

func (bc *blockContext) getAddress(s *storage.Storage, addr string) *storage.AddressStats {
	addr = strings.ToLower(addr)
	bc.statsMutex.Lock()
	stats := bc.addressStats[addr]
	if stats != nil {
		bc.statsMutex.Unlock()
		return stats
	}
	bc.addressStats[addr] = storage.MakeEmptyAddressStats(addr)

	v, err := s.GetAddressStats(addr)
	if err == nil {
		bc.addressStats[addr] = v
	}
	bc.statsMutex.Unlock()
	return bc.addressStats[addr]
}

func (bc *blockContext) SaveTransaction(trx *models.Transaction) {
	from_index := bc.getAddress(bc.storage, trx.From).AddTx()
	to_index := bc.getAddress(bc.storage, trx.To).AddTx()

	bc.batch.AddToBatch(trx, trx.From, from_index)
	bc.batch.AddToBatch(trx, trx.To, to_index)
}
