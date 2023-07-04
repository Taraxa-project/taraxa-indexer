package indexer

import (
	"strings"
	"sync"
	"time"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/metrics"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/internal/utils"
	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/nleeper/goment"
	log "github.com/sirupsen/logrus"
)

type blockContext struct {
	storage      storage.Storage
	batch        storage.Batch
	client       chain.Client
	block        *models.Pbft
	finalized    *storage.FinalizationData
	statsMutex   sync.RWMutex
	addressStats map[string]*storage.AddressStats
}

func MakeBlockContext(s storage.Storage, client chain.Client) *blockContext {
	var bc blockContext
	bc.storage = s
	bc.batch = s.NewBatch()
	bc.client = client
	bc.addressStats = make(map[string]*storage.AddressStats)
	bc.finalized = s.GetFinalizationData()

	return &bc
}

func (bc *blockContext) commit() {
	bc.batch.SetFinalizationData(bc.finalized)
	bc.addAddressStatsToBatch()
	bc.batch.CommitBatch()

	metrics.StorageCommitCounter.Inc()
}

func (bc *blockContext) process(raw *chain.Block) (dags_count, trx_count uint64, err error) {
	// Add reward minted in this block to TotalSupply
	bc.addToTotalSupply(raw.TotalReward)

	start_processing := time.Now()
	bc.block = raw.ToModel()

	tp := utils.MakeThreadPool()
	tp.Go(func() { bc.updateValidatorStats(bc.block) })
	tp.Go(func() { dags_count, err = bc.processDags() })
	tp.Go(func() { err = bc.processTransactions(raw.Transactions) })
	tp.Wait()
	if bc.block.Number%1000 == 0 {
		tp.Go(func() { err = bc.CheckIndexedBalances() })
	}
	if err != nil {
		return
	}

	trx_count = bc.block.TransactionCount
	bc.finalized.TrxCount += trx_count
	bc.finalized.DagCount += dags_count
	bc.finalized.PbftCount++

	pbft_author_index := bc.getAddress(bc.storage, bc.block.Author).AddPbft(bc.block.Timestamp)
	log.WithFields(log.Fields{"author": bc.block.Author, "hash": bc.block.Hash}).Debug("Saving PBFT block")
	bc.batch.AddToBatch(bc.block, bc.block.Author, pbft_author_index)

	// If stats is available check for consistency
	remote_stats, stats_err := bc.client.GetChainStats()
	if stats_err == nil {
		bc.finalized.Check(remote_stats)
	}
	bc.commit()

	metrics.Save(start_processing, dags_count, trx_count, bc.finalized)

	return
}

func (bc *blockContext) addToTotalSupply(amount string) {
	a := parseStringToBigInt(amount)

	current := bc.storage.GetTotalSupply()
	current.Add(current, a)

	bc.batch.SetTotalSupply(current)
}

func (bc *blockContext) updateValidatorStats(block *models.Pbft) {
	tn, _ := goment.Unix(int64(block.Timestamp))
	weekStats := bc.storage.GetWeekStats(int32(tn.ISOWeekYear()), int32(tn.ISOWeek()))
	weekStats.AddPbftBlock(block)
	bc.batch.UpdateWeekStats(weekStats)
}

func (bc *blockContext) processDags() (dags_count uint64, err error) {
	dag_blocks, err := bc.client.GetPeriodDagBlocks(bc.block.Number)
	if err != nil {
		log.WithError(err).Debug("GetPeriodDagBlocks error")
		return bc.processDagsOld()
	}
	dags_count = uint64(len(dag_blocks))
	for _, dag := range dag_blocks {
		bc.saveDag(dag.ToModel())
	}
	return
}

func (bc *blockContext) processDagsOld() (dags_count uint64, err error) {
	block_with_dags, err := bc.client.GetPbftBlockWithDagBlocks(bc.block.Number)
	if err != nil {
		return
	}
	bc.block.PbftHash = block_with_dags.BlockHash

	dags_count = uint64(len(block_with_dags.Schedule.DagBlocksOrder))
	tp := utils.MakeThreadPool()
	for _, dag_hash := range block_with_dags.Schedule.DagBlocksOrder {
		tp.Go(utils.MakeTask(bc.processDag, dag_hash, &err).Run)
	}
	tp.Wait()
	return
}

func (bc *blockContext) processDag(hash string) error {
	raw_dag, err := bc.client.GetDagBlockByHash(hash)
	if err != nil {
		return err
	}
	bc.saveDag(raw_dag.ToModel())
	return nil
}

func (bc *blockContext) saveDag(dag *models.Dag) {
	log.WithFields(log.Fields{"sender": dag.Sender, "hash": dag.Hash}).Trace("Saving DAG block")
	dag_index := bc.getAddress(bc.storage, dag.Sender).AddDag(dag.Timestamp)
	bc.batch.AddToBatch(dag, dag.Sender, dag_index)
}

func (bc *blockContext) getAddress(s storage.Storage, addr string) *storage.AddressStats {
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
