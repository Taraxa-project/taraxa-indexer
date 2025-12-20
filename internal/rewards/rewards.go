package rewards

import (
	"math/big"
	"time"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	log "github.com/sirupsen/logrus"
)

type Rewards struct {
	storage          storage.Storage
	batch            storage.Batch
	config           *common.ChainConfig
	validators       *Validators
	totalStake       *big.Int
	totalSupply      *big.Int
	blockNum         uint64
	blockTime        uint64
	round            uint64
	prevYieldsSaving *storage.YieldSaving
}

func MakeRewards(storage storage.Storage, batch storage.Batch, config *common.ChainConfig, block *chain.BlockData, prevYieldsSaving *storage.YieldSaving) *Rewards {
	r := Rewards{storage, batch, config, MakeValidators(config, block.Validators), block.TotalAmountDelegated, block.TotalSupply, block.Pbft.Number, block.Pbft.Timestamp, block.Round, prevYieldsSaving}
	// special case for  the networks without aspen hf part1 (incorrect initialization of the supply without aspen hf part1)
	if r.totalSupply.Sign() == 0 {
		r.totalSupply = r.storage.GetTotalSupply()
	}

	return &r
}

func (r *Rewards) addTotalMinted(amount *big.Int) {
	current := r.storage.GetTotalSupply()
	current.Add(current, amount)

	r.batch.SetTotalSupply(current)
}

func (r *Rewards) Process(totalMinted *big.Int, block *chain.BlockData) (currentBlockFee *big.Int) {
	rewardsStats := r.makeRewardsStats(block.Dags, block.Votes, block.Transactions, block.Pbft.Author)

	totalReward, currentBlockFee := r.ProcessStats(rewardsStats, totalMinted)

	if totalReward.Cmp(totalMinted) != 0 {
		log.WithFields(log.Fields{"period": r.blockNum, "total_reward_check": totalReward, "total_minted": totalMinted}).Fatal("Total reward check failed")
	}
	r.addTotalMinted(totalReward)

	return
}

func (r *Rewards) ProcessStats(periodStats *storage.RewardsStats, totalMinted *big.Int) (*big.Int, *big.Int) {
	distributionFrequency := r.config.Hardforks.GetDistributionFrequency(r.blockNum)

	if r.blockNum%uint64(distributionFrequency) != 0 {
		// Save blockFee to db to process it later and return it from this method to avoid yield double counting
		toStore := periodStats
		r.batch.AddSingleKey(toStore, storage.FormatIntToKey(r.blockNum))
		return big.NewInt(0), big.NewInt(0)
	}

	periodRewards := makeIntervalRewards()
	if distributionFrequency > 1 {
		// distribute rewards for whole interval
		periodRewards = r.GetIntervalRewards(periodStats, distributionFrequency)
	} else {
		periodRewards = r.rewardsFromStats(periodStats)
	}

	validators_yield := GetValidatorsYield(periodRewards.ValidatorRewards, r.validators)
	r.batch.AddSingleKey(storage.ValidatorsYield{Yields: validators_yield}, storage.FormatIntToKey(r.blockNum))
	r.batch.AddSingleKey(storage.MultipliedYield{Yield: GetMultipliedYield(totalMinted, r.totalStake)}, storage.FormatIntToKey(r.blockNum))
	return periodRewards.TotalReward, periodRewards.BlockFee
}

func (r *Rewards) makeRewardsStats(
	dags []common.DagBlock, votes common.VotesResponse,
	trxs []common.Transaction, block_author string) *storage.RewardsStats {
	return makeRewardsStats(r.config.Hardforks.IsAspenHfOne(r.blockNum), dags, votes, trxs, r.config.CommitteeSize.Uint64(), block_author, r.config.GetLambda(r.round)).ToStorage()
}

func (r *Rewards) calculateBlockReward(total_stake, current_total_tara_supply, blocks_per_year *big.Int) (block_reward *big.Int, yield *big.Int) {
	yield = r.calculateCurrentYield(current_total_tara_supply)
	block_reward = big.NewInt(0).Mul(total_stake, yield)
	block_reward.Div(block_reward, big.NewInt(0).Mul(YieldFractionDecimalPrecision, blocks_per_year))
	return
}

func blocksPerYearFromLambdaMs(lambda_ms, delay_ms uint64) *big.Int {
	var year_ms uint64 = 365 * 24 * 60 * 60
	year_ms *= 1000
	// we have fixed 2*lambda time for proposing step and adding approx delay it takes to receive 2t+1 soft and cert votes
	expected_block_time := 2*lambda_ms + delay_ms
	return big.NewInt(0).SetUint64(year_ms / expected_block_time)
}

func (r *Rewards) calculateFullBlockReward(stats *storage.RewardsStats) *big.Int {
	if r.config.Hardforks.IsCactiHf(r.blockNum) {
		blocks_per_year := blocksPerYearFromLambdaMs(stats.LambdaMs, r.config.Hardforks.CactiHf.ConsensusDelay)
		fullReward, _ := r.calculateBlockReward(r.totalStake, r.totalSupply, blocks_per_year)
		return fullReward
	} else if r.config.Hardforks.IsAspenHfTwo(r.blockNum) {
		fullReward, _ := r.calculateBlockReward(r.totalStake, r.totalSupply, r.config.BlocksPerYear)
		return fullReward
	} else {
		fullReward := big.NewInt(0).Mul(r.totalStake, r.config.YieldPercentage)
		fullReward.Div(fullReward, big.NewInt(0).Mul(big.NewInt(100), r.config.BlocksPerYear))

		return fullReward
	}
}

func (r *Rewards) rewardsFromStats(stats *storage.RewardsStats) (rewards IntervalRewards) {
	full_block_reward := r.calculateFullBlockReward(stats)
	totalRewardsParts := calculatePeriodRewardsParts(r.config, full_block_reward, stats.TotalVotesWeight == 0)

	rewards = makeIntervalRewards()
	for _, s := range stats.ValidatorsStats {
		addr := s.Address
		if !r.validators.Exists(addr) {
			continue
		}
		if rewards.ValidatorRewards[addr] == nil {
			rewards.ValidatorRewards[addr] = big.NewInt(0)
		}

		// Add dags reward
		if s.DagBlocksCount > 0 {
			dag_reward := big.NewInt(0)
			dag_reward.Mul(big.NewInt(0).SetUint64(s.DagBlocksCount), totalRewardsParts.dags)
			dag_reward.Div(dag_reward, big.NewInt(0).SetUint64(stats.TotalDagCount))
			rewards.TotalReward.Add(rewards.TotalReward, dag_reward)
			rewards.ValidatorRewards[addr].Add(rewards.ValidatorRewards[addr], dag_reward)
		}

		// Add voting reward
		if s.VoteWeight > 0 {
			// total_votes_reward * validator_vote_weight / total_votes_weight
			vote_reward := big.NewInt(0).Mul(totalRewardsParts.votes, big.NewInt(0).SetUint64(s.VoteWeight))
			vote_reward.Div(vote_reward, big.NewInt(0).SetUint64(stats.TotalVotesWeight))
			rewards.TotalReward.Add(rewards.TotalReward, vote_reward)
			rewards.ValidatorRewards[addr].Add(rewards.ValidatorRewards[addr], vote_reward)
		}

		if s.FeeReward != nil && s.FeeReward.Cmp(big.NewInt(0)) > 0 {
			rewards.BlockFee.Add(rewards.BlockFee, s.FeeReward)
		}
	}
	blockAuthorReward := big.NewInt(0)
	{
		maxVotesWeight := common.Max(stats.MaxVotesWeight, stats.TotalVotesWeight)
		// fmt.Println("maxVotesWeight", maxVotesWeight, stats.MaxVotesWeight, stats.TotalVotesWeight)
		// In case all reward votes are included we will just pass block author whole reward, this should improve rounding issues
		if maxVotesWeight == stats.TotalVotesWeight {
			blockAuthorReward = totalRewardsParts.bonus
		} else {
			twoTPlusOne := maxVotesWeight*2/3 + 1
			// fmt.Println("twoTPlusOne", twoTPlusOne)
			bonusVotesWeight := uint64(0)
			bonusVotesWeight = stats.TotalVotesWeight - twoTPlusOne
			// should be zero if rewardsStats.TotalVotesWeight == twoTPlusOne
			blockAuthorReward.Div(big.NewInt(0).Mul(totalRewardsParts.bonus, big.NewInt(0).SetUint64(bonusVotesWeight)), big.NewInt(0).SetUint64(maxVotesWeight-twoTPlusOne))
			// fmt.Println("blockAuthorReward", blockAuthorReward)
		}
	}
	if blockAuthorReward.Cmp(big.NewInt(0)) > 0 {
		if rewards.ValidatorRewards[stats.BlockAuthor] == nil {
			rewards.ValidatorRewards[stats.BlockAuthor] = big.NewInt(0)
		}
		rewards.TotalReward.Add(rewards.TotalReward, blockAuthorReward)
		rewards.ValidatorRewards[stats.BlockAuthor].Add(rewards.ValidatorRewards[stats.BlockAuthor], blockAuthorReward)
	}
	return
}

func (r *Rewards) AfterCommit() {
	b := r.storage.NewBatch()
	rewardsDistributed := r.blockNum%uint64(r.config.Hardforks.GetDistributionFrequency(r.blockNum)) == 0
	_, pWeek := time.Unix(int64(r.prevYieldsSaving.Time), 0).ISOWeek()
	_, cWeek := time.Unix(int64(r.blockTime), 0).ISOWeek()
	processYields := pWeek != cWeek

	if rewardsDistributed && processYields {
		r.processIntervalYield(r.prevYieldsSaving.Period, b)
		r.processValidatorsIntervalYield(r.prevYieldsSaving.Period, b)
		// get a start of an hour from the current time
		log.WithFields(log.Fields{"period": r.blockNum, "time": r.blockTime}).Info("Adding yield saving")
		b.AddYieldSaving(r.blockNum, r.blockTime)
		*r.prevYieldsSaving = storage.YieldSaving{Time: r.blockTime, Period: r.blockNum}
	}
	b.CommitBatch()
}
