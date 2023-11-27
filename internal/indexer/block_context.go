package indexer

import (
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/metrics"
	"github.com/Taraxa-project/taraxa-indexer/internal/oracle"
	"github.com/Taraxa-project/taraxa-indexer/internal/rewards"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/nleeper/goment"
	log "github.com/sirupsen/logrus"
)

type blockContext struct {
	Storage      storage.Storage
	Batch        storage.Batch
	Config       *common.Config
	Client       chain.Client
	Oracle       *oracle.Oracle
	block        *models.Pbft
	dags         []chain.DagBlock
	transactions []models.Transaction
	finalized    *storage.FinalizationData
	statsMutex   sync.RWMutex
	addressStats map[string]*storage.AddressStats
	balances     *storage.Balances
	blockFee     *big.Int
}

func MakeBlockContext(s storage.Storage, client chain.Client, oracle *oracle.Oracle, config *common.Config) *blockContext {
	var bc blockContext
	bc.Storage = s
	bc.Batch = s.NewBatch()
	bc.Client = client
	bc.Oracle = oracle
	bc.addressStats = make(map[string]*storage.AddressStats)
	bc.finalized = s.GetFinalizationData()
	bc.Config = config
	bc.balances = &storage.Balances{Accounts: bc.Storage.GetAccounts()}
	bc.blockFee = big.NewInt(0)

	return &bc
}

func (bc *blockContext) commit() {
	bc.Batch.SetFinalizationData(bc.finalized)
	bc.addAddressStatsToBatch()
	bc.Batch.CommitBatch()

	metrics.StorageCommitCounter.Inc()
}

func (bc *blockContext) process(raw chain.Block) (dags_count, trx_count uint64, err error) {
	start_processing := time.Now()
	bc.block = raw.ToModel()

	tp := common.MakeThreadPool()
	tp.Go(func() { bc.updateValidatorStats(bc.block) })
	tp.Go(common.MakeTaskWithoutParams(bc.processDags, &err).Run)
	tp.Go(common.MakeTask(bc.processTransactions, raw.Transactions, &err).Run)
	votes := new(chain.VotesResponse)
	tp.Go(common.MakeTaskWithResult(bc.Client.GetPreviousBlockCertVotes, bc.block.Number, votes, &err).Run)

	validators := make([]chain.Validator, 0)
	tp.Go(common.MakeTaskWithResult(bc.Client.GetValidatorsAtBlock, bc.block.Number, &validators, &err).Run)

	tp.Wait()
	if err != nil {
		return
	}

	totalReward := common.ParseStringToBigInt(raw.TotalReward)

	r := rewards.MakeRewards(bc.Oracle, bc.Storage, bc.Batch, bc.Config, bc.block, bc.blockFee, validators)
	blockFee := r.Process(totalReward, bc.dags, bc.transactions, *votes)
	if blockFee != nil {
		bc.balances.AddToBalance(common.DposContractAddress, blockFee)
	}

	bc.balances.AddToBalance(common.DposContractAddress, totalReward)

	if bc.block.Number%1000 == 0 {
		err = bc.checkIndexedBalances()
		if err != nil {
			return
		}
	}
	bc.Batch.SaveAccounts(bc.balances)

	dags_count = uint64(len(bc.dags))
	trx_count = bc.block.TransactionCount
	bc.finalized.TrxCount += trx_count
	bc.finalized.DagCount += dags_count
	bc.finalized.PbftCount++

	pbft_author_index := bc.getAddress(bc.Storage, bc.block.Author).AddPbft(bc.block.Timestamp)
	log.WithFields(log.Fields{"author": bc.block.Author, "hash": bc.block.Hash}).Debug("Saving PBFT block")
	bc.Batch.AddToBatch(bc.block, bc.block.Author, pbft_author_index)

	// If stats is available check for consistency
	remote_stats, stats_err := bc.Client.GetChainStats()
	if stats_err == nil {
		bc.finalized.Check(remote_stats)
	}
	bc.commit()
	r.AfterCommit()

	metrics.Save(start_processing, dags_count, trx_count, bc.finalized)

	return
}

func (bc *blockContext) checkIndexedBalances() (err error) {
	tp := common.MakeThreadPool()
	for _, balance := range bc.balances.Accounts {
		address := balance.Address
		balance := balance.Balance
		tp.Go(func() {
			b, get_err := bc.Client.GetBalanceAtBlock(address, bc.block.Number)
			if get_err != nil {
				err = get_err
				return
			}
			chain_balance := common.ParseStringToBigInt(b)
			if balance.Cmp(chain_balance) != 0 {
				err = fmt.Errorf("balance of %s: calc(%s) != chain(%s)", address, balance, chain_balance)
			}
		})
	}
	tp.Wait()

	return
}

func (bc *blockContext) updateValidatorStats(block *models.Pbft) {
	tn, _ := goment.Unix(int64(block.Timestamp))
	weekStats := bc.Storage.GetWeekStats(int32(tn.ISOWeekYear()), int32(tn.ISOWeek()))
	weekStats.AddPbftBlock(block)
	bc.Batch.UpdateWeekStats(weekStats)
}

func (bc *blockContext) processDags() (err error) {
	dag_blocks, err := bc.Client.GetPeriodDagBlocks(bc.block.Number)
	if err != nil {
		log.WithError(err).Debug("GetPeriodDagBlocks error")
		return bc.processDagsOld()
	}
	bc.dags = make([]chain.DagBlock, len(dag_blocks))
	for i, dag := range dag_blocks {
		bc.dags[i] = dag
		bc.saveDag(&dag)
	}
	return
}

func (bc *blockContext) processDagsOld() (err error) {
	block_with_dags, err := bc.Client.GetPbftBlockWithDagBlocks(bc.block.Number)
	if err != nil {
		return
	}
	tp := common.MakeThreadPool()
	for i, dag_hash := range block_with_dags.Schedule.DagBlocksOrder {
		tp.Go(common.MakeTaskWithResult(bc.processDag, dag_hash, &bc.dags[i], &err).Run)
	}
	tp.Wait()
	return
}

func (bc *blockContext) processDag(hash string) (dag chain.DagBlock, err error) {
	dag, err = bc.Client.GetDagBlockByHash(hash)
	if err != nil {
		return chain.DagBlock{}, err
	}
	bc.saveDag(&dag)
	return
}

func (bc *blockContext) saveDag(dag *chain.DagBlock) {
	log.WithFields(log.Fields{"sender": dag.Sender, "hash": dag.Hash}).Trace("Saving DAG block")
	dagModel := dag.ToModel()
	dag_index := bc.getAddress(bc.Storage, dag.Sender).AddDag(dagModel.Timestamp)
	bc.Batch.AddToBatch(dagModel, dag.Sender, dag_index)
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
