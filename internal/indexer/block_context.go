package indexer

import (
	"strings"
	"sync"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/nleeper/goment"
	log "github.com/sirupsen/logrus"
	"github.com/spiretechnology/go-pool"
)

type blockContext struct {
	storage        *storage.Storage
	batch          *storage.Batch
	client         *chain.WsClient
	tp             pool.Pool
	blockTimestamp uint64
	addressStats   map[string]*storage.AddressStats
	finalized      *storage.FinalizationData
	statsMutex     sync.RWMutex
}

func MakeBlockContext(s *storage.Storage, client *chain.WsClient, tp pool.Pool) *blockContext {
	var bc blockContext
	bc.storage = s
	bc.batch = s.NewBatch()
	bc.client = client
	// pool limit is limit of concurrent ws requests to the node
	bc.tp = tp
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

	bc.tp.Go(func() { bc.updateValidatorStats(block) })

	for _, trx_hash := range *transactions {
		bc.tp.Go(MakeTask(bc.processTransaction, trx_hash, &err).Run)
	}

	block_with_dags, pbft_err := bc.client.GetPbftBlockWithDagBlocks(block.Number)
	block.PbftHash = block_with_dags.BlockHash
	if pbft_err != nil {
		err = pbft_err
		return
	}
	dags_count = uint64(len(block_with_dags.Schedule.DagBlocksOrder))
	bc.finalized.DagCount += dags_count

	for _, dag_hash := range block_with_dags.Schedule.DagBlocksOrder {
		bc.tp.Go(MakeTask(bc.processDag, dag_hash, &err).Run)
	}

	bc.tp.Wait()
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
	tn, _ := goment.Unix(int64(block.Timestamp))
	weekStats := bc.storage.GetWeekStats(int32(tn.ISOWeekYear()), int32(tn.ISOWeek()))
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
