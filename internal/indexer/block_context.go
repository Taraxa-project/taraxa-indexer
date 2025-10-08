package indexer

import (
	"fmt"
	"math/big"
	"time"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/metrics"
	"github.com/Taraxa-project/taraxa-indexer/internal/rewards"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/nleeper/goment"
	log "github.com/sirupsen/logrus"
)

type blockContext struct {
	Storage            storage.Storage
	Batch              storage.Batch
	Config             *common.Config
	Client             common.Client
	Block              *chain.BlockData
	addressStats       *storage.AddressStatsMap
	finalized          *common.FinalizationData
	dayStats           *storage.DayStatsWithTimestamp
	dailyContractUsers map[string]*storage.DailyContractUsers // key: contract_address
}

func MakeBlockContext(s storage.Storage, client common.Client, config *common.Config, dayStats *storage.DayStatsWithTimestamp) *blockContext {
	var bc blockContext
	bc.Storage = s
	bc.Batch = s.NewBatch()
	bc.Config = config
	bc.addressStats = storage.MakeAddressStatsMap()
	bc.finalized = s.GetFinalizationData()
	bc.Client = client
	bc.dayStats = dayStats
	bc.dailyContractUsers = make(map[string]*storage.DailyContractUsers)

	return &bc
}

func (bc *blockContext) SetBlockData(bd *chain.BlockData) {
	bc.Block = bd
}

func (bc *blockContext) commit() {
	bc.Batch.SetFinalizationData(bc.finalized)
	bc.addressStats.AddToBatch(bc.Batch)
	// Save daily contract users before committing
	bc.saveDailyContractUsers()
	bc.Batch.CommitBatch()

	metrics.StorageCommitCounter.Inc()
}

// getDailyContractUsers gets or creates DailyContractUsers for a contract address
func (bc *blockContext) getDailyContractUsers(contractAddress string) *storage.DailyContractUsers {
	if _, exists := bc.dailyContractUsers[contractAddress]; !exists {
		// Load existing data from storage or create new
		dayStart := common.DayStart(bc.Block.Pbft.Timestamp)
		existingUsers := bc.Storage.GetDailyContractUsers(contractAddress, dayStart)

		bc.dailyContractUsers[contractAddress] = storage.MakeDailyContractUsers()

		if len(existingUsers.Users) > 0 {
			bc.dailyContractUsers[contractAddress] = storage.MakeDailyContractUsersFromList(existingUsers)
		}
	}
	return bc.dailyContractUsers[contractAddress]
}

// addContractUser adds a user to the daily contract users tracking
func (bc *blockContext) addContractUser(sender, receiver string) {
	// check if receiver is a contract address
	if !bc.addressStats.GetAddress(bc.Storage, receiver).IsContract() {
		return
	}

	users := bc.getDailyContractUsers(receiver)
	users.Add(sender)
}

// saveDailyContractUsers saves all tracked daily contract users to storage
func (bc *blockContext) saveDailyContractUsers() {
	dayStart := common.DayStart(bc.Block.Pbft.Timestamp)
	for contractAddress, users := range bc.dailyContractUsers {
		bc.Batch.AddDailyContractUsers(contractAddress, dayStart, users)
	}
}

func (bc *blockContext) process(bd *chain.BlockData, stats *chain.Stats) (dags_count, trx_count uint64, err error) {
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

	r := rewards.MakeRewards(bc.Storage, bc.Batch, bc.Config, bc.Block)
	blockFee := r.Process(totalReward, bc.Block.Dags, bc.Block.Transactions, bc.Block.Votes, bc.Block.Pbft.Author)

	// add total fee to the dpos contract balance after the magnolia hardfork(it is added to block producers commission pools)
	if bc.Config.Chain != nil && (bc.Block.Pbft.Number >= bc.Config.Chain.Hardforks.MagnoliaHf.BlockNum) {
		if blockFee != nil && blockFee.Cmp(big.NewInt(0)) > 0 {
			bc.addressStats.AddToBalance(bc.Storage, common.DposContractAddress, blockFee)
		}
	}

	bc.addressStats.AddToBalance(bc.Storage, common.DposContractAddress, totalReward)

	// disable balance check for now as it's taking too long
	// if bc.Block.Pbft.Number%1000 == 0 {
	// 	bc.checkIndexedBalances()
	// }
	if bc.Block.Pbft.Number%5000 == 0 {
		bc.SaveHoldersLeaderboard()
	}

	bc.dayStats.AddBlock(bc.Block.Pbft)
	bc.Batch.AddDayStats(bc.dayStats)

	dags_count = uint64(len(bc.Block.Dags))
	trx_count = uint64(len(bc.Block.Transactions))
	bc.finalized.TrxCount += trx_count
	bc.finalized.DagCount += dags_count
	bc.finalized.PbftCount++

	pbft_author_index := bc.addressStats.GetAddress(bc.Storage, bc.Block.Pbft.Author).AddPbft(bc.Block.Pbft.Timestamp)
	log.WithFields(log.Fields{"author": bc.Block.Pbft.Author, "hash": bc.Block.Pbft.Hash}).Debug("Saving PBFT block")
	pbft_model := bc.Block.Pbft.GetModel()
	bc.Batch.Add(pbft_model, bc.Block.Pbft.Author, pbft_author_index)
	stats.AddPbft(pbft_model)

	bc.commit()
	r.AfterCommit()
	metrics.Save(start_processing, dags_count, trx_count, bc.finalized)
	return
}

func (bc *blockContext) SaveHoldersLeaderboard() {
	accounts := storage.MakeAccountBalancesMap()
	start := time.Now()
	stats := storage.AddressStats{}
	bc.Storage.ForEach(&stats, "", nil, storage.Forward, func(key, res []byte) (stop bool) {
		err := rlp.DecodeBytes(res, &stats)
		if err != nil {
			log.WithError(err).Fatal("storage.ForEach failed")
		}
		// skip zero balances
		if stats.Balance.Cmp(big.NewInt(0)) == 0 {
			return false
		}
		accounts.Set(stats.Address, stats.Balance)
		return false
	})
	log.WithFields(log.Fields{"time": time.Since(start)}).Info("Iterate over address stats for leaderboard")
	bc.Batch.SaveHoldersLeaderboard(accounts.Sorted())
	log.WithFields(log.Fields{"time": time.Since(start)}).Info("Saved holders leaderboard")
}

// func (bc *blockContext) checkIndexedBalances() {
// 	if bc.accounts.GetLength() == 0 {
// 		log.Fatal("checkIndexedBalances: No balances in the storage, something is wrong")
// 	}
// 	tp := common.MakeThreadPool()
// 	for a, b := range bc.accounts.GetAccounts() {
// 		address := a
// 		balance := b
// 		tp.Go(func() {
// 			b, get_err := bc.Client.GetBalanceAtBlock(address, bc.Block.Pbft.Number)
// 			if get_err != nil {
// 				log.WithError(get_err).WithField("address", address).Warn("GetBalanceAtBlock error for address")
// 				return
// 			}
// 			chain_balance := common.ParseStringToBigInt(b)
// 			if balance.Cmp(chain_balance) != 0 {
// 				log.WithFields(log.Fields{"address": address, "balance": balance, "chain_balance": chain_balance}).Error("Balance check failed")
// 			}
// 		})
// 	}
// 	tp.Wait()
// }

func (bc *blockContext) updateValidatorStats(block *common.Block) {
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

func (bc *blockContext) saveDag(dag *common.DagBlock) {
	log.WithFields(log.Fields{"sender": dag.Sender, "hash": dag.Hash}).Trace("Saving DAG block")
	dag_index := bc.addressStats.GetAddress(bc.Storage, dag.Sender).AddDag(dag.GetModel().Timestamp)
	bc.Batch.Add(dag.GetModel(), dag.Sender, dag_index)
}
