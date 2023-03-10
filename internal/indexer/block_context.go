package indexer

import (
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/models"
	log "github.com/sirupsen/logrus"
	"github.com/spiretechnology/go-pool"
)

type blockContext struct {
	storage        *storage.Storage
	batch          *storage.Batch
	client         *chain.WsClient
	pool           pool.Pool
	blockTimestamp uint64
	addressStats   map[string]*storage.AddressStats
	finalized      *storage.FinalizationData
	statsMutex     sync.RWMutex
}

func MakeBlockContext(s *storage.Storage, client *chain.WsClient) *blockContext {
	var bc blockContext
	bc.storage = s
	bc.batch = s.NewBatch()
	bc.client = client
	// pool limit is limit of concurrent ws requests to the node
	bc.pool = pool.New(uint(10 * runtime.NumCPU()))
	bc.addressStats = make(map[string]*storage.AddressStats)
	bc.finalized = s.GetFinalizationData()

	return &bc
}

func (bc *blockContext) commit(period uint64) {
	bc.batch.SaveFinalizedPeriod(bc.finalized)
	bc.addAddressStatsToBatch()
	bc.batch.CommitBatch()
}

func (bc *blockContext) process(raw *chain.Block) (dags_count, trx_count uint64, err error) {
	block := raw.ToModel()
	transactions := raw.Transactions

	trx_count = block.TransactionCount
	bc.finalized.TrxCount += trx_count
	bc.blockTimestamp = block.Timestamp

	bc.pool.Go(func() { bc.updateValidatorStats(block) })

	for _, trx_hash := range *transactions {
		bc.pool.Go(MakeTask(bc.processTransaction, trx_hash, &err).Run)
	}

	block_with_dags, err := bc.client.GetPbftBlockWithDagBlocks(block.Number)
	if err != nil {
		return
	}
	dags_count = uint64(len(block_with_dags.Schedule.DagBlocksOrder))
	bc.finalized.DagCount += dags_count

	for _, dag_hash := range block_with_dags.Schedule.DagBlocksOrder {
		bc.pool.Go(MakeTask(bc.processDag, dag_hash, &err).Run)
	}

	bc.pool.Wait()
	if err != nil {
		return
	}

	bc.finalized.PbftCount++
	author_pbft_index := bc.getAddress(bc.storage, block.Author).AddPbft(block.Timestamp)
	log.WithFields(log.Fields{"author": block.Author, "hash": block.Hash}).Debug("Saving PBFT block")
	bc.batch.AddToBatch(block, block.Author, author_pbft_index)

	// If stats is available check for consistency
	remote_stats, stats_err := bc.client.GetChainStats()
	if stats_err == nil {
		bc.finalized.Check(remote_stats)
	}
	bc.commit(block.Number)
	return
}

func (bc *blockContext) processTransaction(hash string) error {
	trx, err := bc.client.GetTransactionByHash(hash)
	if err != nil {
		return err
	}

	bc.SaveTransaction(trx.ToModelWithTimestamp(bc.blockTimestamp))
	return nil
}

func (bc *blockContext) updateValidatorStats(block *models.Pbft) {
	tn := time.Unix(int64(block.Timestamp), 0)
	year, week := tn.ISOWeek()
	weekStats := bc.storage.GetWeekStats(year, week)
	weekStats.AddPbftBlock(block)
	bc.batch.UpdateWeekStats(weekStats)
}

func (bc *blockContext) processDag(hash string) error {
	raw_dag, err := bc.client.GetDagBlockByHash(hash)
	if err != nil {
		return err
	}
	dag := raw_dag.ToModel()
	log.WithFields(log.Fields{"sender": dag.Sender, "hash": dag.Hash}).Trace("Saving DAG block")
	dag_index := bc.getAddress(bc.storage, dag.Sender).AddDag(dag.Timestamp)
	bc.batch.AddToBatch(dag, dag.Sender, dag_index)
	return nil
}

func (bc *blockContext) addAddressStatsToBatch() {
	for _, stats := range bc.addressStats {
		bc.batch.AddToBatch(stats, stats.Address, 0)
	}
}

func (bc *blockContext) getAddress(s *storage.Storage, addr string) *storage.AddressStats {
	addr = strings.ToLower(addr)
	bc.statsMutex.Lock()
	defer bc.statsMutex.Unlock()
	stats := bc.addressStats[addr]
	if stats != nil {
		return stats
	}

	bc.addressStats[addr] = s.GetAddressStats(addr)

	return bc.addressStats[addr]
}

func (bc *blockContext) SaveTransaction(trx *models.Transaction) {
	log.WithFields(log.Fields{"from": trx.From, "to": trx.To, "hash": trx.Hash}).Trace("Saving transaction")
	from_index := bc.getAddress(bc.storage, trx.From).AddTransaction(trx.Timestamp)
	to_index := bc.getAddress(bc.storage, trx.To).AddTransaction(trx.Timestamp)

	bc.batch.AddToBatch(trx, trx.From, from_index)
	bc.batch.AddToBatch(trx, trx.To, to_index)
}
