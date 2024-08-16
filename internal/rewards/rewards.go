package rewards

import (
	"math/big"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/oracle"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	log "github.com/sirupsen/logrus"
)

type Rewards struct {
	oracle      *oracle.Oracle
	storage     storage.Storage
	batch       storage.Batch
	config      *common.Config
	validators  *Validators
	totalStake  *big.Int
	totalSupply *big.Int

	blockNum uint64
}

func MakeRewards(oracle *oracle.Oracle, storage storage.Storage, batch storage.Batch, config *common.Config, block *chain.BlockData) *Rewards {
	r := Rewards{oracle, storage, batch, config, MakeValidators(config, block.Validators), block.TotalAmountDelegated, block.TotalSupply, block.Pbft.Number}
	return &r
}

func (r *Rewards) addTotalMinted(amount *big.Int) {
	current := r.storage.GetTotalSupply()
	current.Add(current, amount)

	r.batch.SetTotalSupply(current)
}

func (r *Rewards) Process(total_minted *big.Int, dags []chain.DagBlock, trxs []chain.Transaction, votes chain.VotesResponse, block_author string) (currentBlockFee *big.Int) {
	if r.blockNum%r.config.TotalYieldSavingInterval == 0 {
		log.WithFields(log.Fields{"total_stake": r.totalStake}).Info("totalStake")
	}
	rewardsStats := r.makeRewardsStats(dags, votes, trxs, block_author)
	totalReward, currentBlockFee := r.ProcessStats(rewardsStats, total_minted, r.totalStake)

	// if totalReward.Cmp(total_minted) != 0 {
	// 	log.WithFields(log.Fields{"period": r.blockNum, "total_reward_check": totalReward, "total_minted": total_minted}).Fatal("Total reward check failed")
	// }
	r.addTotalMinted(totalReward)

	return
}

func (r *Rewards) ProcessStats(periodStats *storage.RewardsStats, total_minted *big.Int, totalStake *big.Int) (*big.Int, *big.Int) {
	distributionFrequency := r.config.Chain.Hardforks.GetDistributionFrequency(r.blockNum)

	if r.blockNum%uint64(distributionFrequency) != 0 {
		// Save blockFee to db to process it later and return it from this method to avoid yield double counting
		toStore := periodStats
		r.batch.AddToBatchSingleKey(toStore, storage.FormatIntToKey(r.blockNum))
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
	r.batch.AddToBatchSingleKey(storage.ValidatorsYield{Yields: validators_yield}, storage.FormatIntToKey(r.blockNum))
	r.batch.AddToBatchSingleKey(storage.MultipliedYield{Yield: GetMultipliedYield(total_minted, totalStake)}, storage.FormatIntToKey(r.blockNum))
	return periodRewards.TotalReward, periodRewards.BlockFee
}

func (r *Rewards) makeRewardsStats(
	dags []chain.DagBlock, votes chain.VotesResponse,
	trxs []chain.Transaction, block_author string) *storage.RewardsStats {
	return makeRewardsStats(r.config.Chain.Hardforks.IsAspenHfOne(r.blockNum), dags, votes, trxs, r.config.Chain.CommitteeSize.Uint64(), block_author).ToStorage()
}

func (r *Rewards) calculateBlockReward(total_stake, current_total_tara_supply *big.Int) (block_reward *big.Int, yield *big.Int) {
	yield = r.calculateCurrentYield(current_total_tara_supply)
	block_reward = big.NewInt(0).Mul(total_stake, yield)
	block_reward.Div(block_reward, big.NewInt(0).Mul(YieldFractionDecimalPrecision, r.config.Chain.BlocksPerYear))
	return
}

func (r *Rewards) calculateFullBlockReward() *big.Int {
	if r.config.Chain.Hardforks.IsAspenHfTwo(r.blockNum) {
		fullReward, _ := r.calculateBlockReward(r.totalStake, r.totalSupply)
		return fullReward
	} else {
		fullReward := big.NewInt(0).Mul(r.totalStake, r.config.Chain.YieldPercentage)
		fullReward.Div(fullReward, big.NewInt(0).Mul(big.NewInt(100), r.config.Chain.BlocksPerYear))

		return fullReward
	}
}

func (r *Rewards) rewardsFromStats(stats *storage.RewardsStats) (rewards IntervalRewards) {
	full_block_reward := r.calculateFullBlockReward()
	totalRewardsParts := calculatePeriodRewardsParts(r.config.Chain, full_block_reward, stats.TotalVotesWeight == 0)

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
	if r.blockNum%r.config.TotalYieldSavingInterval == 0 {
		r.processIntervalYield(b)
	}
	if r.blockNum%r.config.ValidatorsYieldSavingInterval == 0 {
		r.processValidatorsIntervalYield(b)
	}
	b.CommitBatch()
}
