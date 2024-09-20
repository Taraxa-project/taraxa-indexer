package indexer

import (
	"fmt"
	"math/big"
	"time"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/metrics"
	"github.com/Taraxa-project/taraxa-indexer/internal/oracle"
	"github.com/Taraxa-project/taraxa-indexer/internal/rewards"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/nleeper/goment"
	log "github.com/sirupsen/logrus"
)

type blockContext struct {
	Storage      storage.Storage
	Batch        storage.Batch
	Config       *common.Config
	Client       chain.Client
	Oracle       *oracle.Oracle
	Block        *chain.BlockData
	accounts     storage.Accounts
	addressStats *storage.AddressStatsMap
	finalized    *storage.FinalizationData
}

func MakeBlockContext(s storage.Storage, client chain.Client, oracle *oracle.Oracle, config *common.Config) *blockContext {
	var bc blockContext
	bc.Storage = s
	bc.Batch = s.NewBatch()
	bc.Config = config
	bc.accounts = bc.Storage.GetAccounts()
	bc.addressStats = storage.MakeAddressStatsMap()
	bc.finalized = s.GetFinalizationData()
	bc.Client = client
	bc.Oracle = oracle

	return &bc
}

func (bc *blockContext) SetBlockData(bd *chain.BlockData) {
	bc.Block = bd
}

func (bc *blockContext) commit() {
	bc.Batch.SetFinalizationData(bc.finalized)
	bc.addressStats.AddToBatch(bc.Batch)
	bc.Batch.CommitBatch()

	metrics.StorageCommitCounter.Inc()
}

func (bc *blockContext) process(bd *chain.BlockData) (dags_count, trx_count uint64, err error) {
	if (bc.finalized.PbftCount + 1) != bd.Pbft.Number {
		err = fmt.Errorf("block number mismatch: %d != %d", bc.finalized.PbftCount+1, bd.Pbft.Number)
		return
	}
	start_processing := time.Now()
	bc.Block = bd

	tp := common.MakeThreadPool()
	tp.Go(func() { bc.updateValidatorStats(bc.Block.Pbft) })
	tp.Go(common.MakeTaskWithoutParams(bc.processDags, &err).Run)
	tp.Go(common.MakeTaskWithoutParams(bc.processTransactions, &err).Run)

	tp.Wait()
	if err != nil {
		return
	}

	totalReward := common.ParseStringToBigInt(bd.Pbft.TotalReward)

	r := rewards.MakeRewards(bc.Oracle, bc.Storage, bc.Batch, bc.Config, bc.Block)
	blockFee := r.Process(totalReward, bc.Block.Dags, bc.Block.Transactions, bc.Block.Votes, bc.Block.Pbft.Author)

	// add total fee to the dpos contract balance after the magnolia hardfork(it is added to block producers commission pools)
	if bc.Config.Chain != nil && (bc.Block.Pbft.Number >= bc.Config.Chain.Hardforks.MagnoliaHf.BlockNum) {
		if blockFee != nil && blockFee.Cmp(big.NewInt(0)) > 0 {
			bc.accounts.AddToBalance(common.DposContractAddress, blockFee)
		}
	}

	bc.accounts.AddToBalance(common.DposContractAddress, totalReward)

	if bc.Block.Pbft.Number%1000 == 0 {
		bc.checkIndexedBalances()
	}

	bc.Batch.SaveAccounts(bc.accounts)

	dags_count = uint64(len(bc.Block.Dags))
	trx_count = uint64(len(bc.Block.Transactions))
	bc.finalized.TrxCount += trx_count
	bc.finalized.DagCount += dags_count
	bc.finalized.PbftCount++

	pbft_author_index := bc.addressStats.GetAddress(bc.Storage, bc.Block.Pbft.Author).AddPbft(bc.Block.Pbft.Timestamp)
	log.WithFields(log.Fields{"author": bc.Block.Pbft.Author, "hash": bc.Block.Pbft.Hash}).Debug("Saving PBFT block")
	bc.Batch.Add(bc.Block.Pbft.GetModel(), bc.Block.Pbft.Author, pbft_author_index)

	bc.commit()
	r.AfterCommit()
	metrics.Save(start_processing, dags_count, trx_count, bc.finalized)
	return
}

func (bc *blockContext) checkIndexedBalances() {
	if len(bc.accounts) == 0 {
		log.Fatal("checkIndexedBalances: No balances in the storage, something is wrong")
	}
	tp := common.MakeThreadPool()
	for _, account := range bc.accounts {
		address := account.Address
		balance := account.Balance
		tp.Go(func() {
			b, get_err := bc.Client.GetBalanceAtBlock(address, bc.Block.Pbft.Number)
			if get_err != nil {
				log.WithError(get_err).WithField("address", address).Warn("GetBalanceAtBlock error for address")
				return
			}
			chain_balance := common.ParseStringToBigInt(b)
			if balance.Cmp(chain_balance) != 0 {
				log.WithFields(log.Fields{"address": address, "balance": balance, "chain_balance": chain_balance}).Error("Balance check failed")
			}
		})
	}
	tp.Wait()
}

func (bc *blockContext) updateValidatorStats(block *chain.Block) {
	tn, _ := goment.Unix(int64(block.Timestamp))
	weekStats := bc.Storage.GetWeekStats(int32(tn.ISOWeekYear()), int32(tn.ISOWeek()))
	weekStats.AddPbftBlock(block.GetModel())
	bc.Batch.UpdateWeekStats(weekStats)
}

func (bc *blockContext) processDags() (err error) {
	for _, dag := range bc.Block.Dags {
		bc.saveDag(&dag)
	}
	return
}

func (bc *blockContext) saveDag(dag *chain.DagBlock) {
	log.WithFields(log.Fields{"sender": dag.Sender, "hash": dag.Hash}).Trace("Saving DAG block")
	dag_index := bc.addressStats.GetAddress(bc.Storage, dag.Sender).AddDag(dag.GetModel().Timestamp)
	bc.Batch.Add(dag.GetModel(), dag.Sender, dag_index)
}
