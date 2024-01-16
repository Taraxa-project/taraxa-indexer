package rewards

import (
	"math/big"
	"strings"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/oracle"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	ethcommon "github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
)

var multiplier = big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)
var percentage_multiplier = big.NewInt(10000)
var YieldFractionDecimalPrecision = big.NewInt(1e+6)

type Rewards struct {
	oracle      *oracle.Oracle
	storage     storage.Storage
	batch       storage.Batch
	config      *common.Config
	validators  *Validators
	blockNum    uint64
	blockAuthor string
}

func MakeRewards(oracle *oracle.Oracle, storage storage.Storage, batch storage.Batch, config *common.Config, block *chain.Block, validators []chain.Validator) *Rewards {
	r := Rewards{oracle, storage, batch, config, MakeValidators(config, validators), block.Number, strings.ToLower(block.Author)}
	return &r
}

func (r *Rewards) Process(total_minted *big.Int, dags []chain.DagBlock, trxs []chain.Transaction, votes chain.VotesResponse) (currentBlockFee *big.Int) {
	totalStake := CalculateTotalStake(r.validators)
	if r.blockNum%r.config.TotalYieldSavingInterval == 0 {
		log.WithFields(log.Fields{"total_stake": totalStake}).Info("totalStake")
	}

	rewards := r.calculateValidatorsRewards(dags, votes, trxs, totalStake)
	totalReward, currentBlockFee := r.ProcessRewards(rewards, total_minted, totalStake)

	if totalReward.Cmp(totalReward) != 0 {
		log.WithFields(log.Fields{"period": r.blockNum, "total_reward_check": totalReward, "total_minted": total_minted}).Fatal("Total reward check failed")
	}
	r.addTotalMinted(totalReward)

	return
}

func (r *Rewards) ProcessRewards(periodRewards PeriodRewards, total_minted *big.Int, totalStake *big.Int) (*big.Int, *big.Int) {
	distributionFrequency := r.config.Chain.Hardforks.GetDistributionFrequency(r.blockNum)

	if r.blockNum%uint64(distributionFrequency) == 0 {
		// distribute rewards for whole interval
		periodRewards = r.GetIntervalRewards(periodRewards, distributionFrequency)
	} else if distributionFrequency > 1 {
		// Save blockFee to db to process it later and return it from this method to avoid yield double counting
		toStore := periodRewards.ToStorage()
		r.batch.AddToBatchSingleKey(toStore, storage.FormatIntToKey(r.blockNum))
		return big.NewInt(0), big.NewInt(0)
	}

	validators_yield := GetValidatorsYield(periodRewards.ValidatorRewards, r.validators)
	r.batch.AddToBatchSingleKey(storage.ValidatorsYield{Yields: validators_yield}, storage.FormatIntToKey(r.blockNum))
	r.batch.AddToBatchSingleKey(storage.MultipliedYield{Yield: GetMultipliedYield(total_minted, totalStake)}, storage.FormatIntToKey(r.blockNum))

	return periodRewards.TotalReward, periodRewards.BlockFee
}

func (r *Rewards) addTotalMinted(amount *big.Int) {
	current := r.storage.GetTotalSupply()
	current.Add(current, amount)

	r.batch.SetTotalSupply(current)
}

func (r *Rewards) calculateValidatorsRewards(
	dags []chain.DagBlock, votes chain.VotesResponse,
	trxs []chain.Transaction, totalStake *big.Int) PeriodRewards {
	stats := makeStats(dags, votes, trxs, r.config.Chain.CommitteeSize.Int64())
	return r.rewardsFromStats(totalStake, stats)
}

func (r *Rewards) calculateCurrentYield(current_total_tara_supply *big.Int) *big.Int {
	// Current yield = (max supply - current total supply) / current total supply
	current_yield := big.NewInt(0).Sub(r.config.Chain.Hardforks.AspenHf.MaxSupply, current_total_tara_supply)
	current_yield.Mul(current_yield, YieldFractionDecimalPrecision)
	current_yield.Div(current_yield, current_total_tara_supply)

	return current_yield
}

func (r *Rewards) calculateBlockReward(totalStake, current_total_tara_supply *big.Int) (block_reward *big.Int, yield *big.Int) {
	yield = r.calculateCurrentYield(current_total_tara_supply)
	block_reward = big.NewInt(0).Mul(totalStake, yield)
	block_reward.Div(block_reward, big.NewInt(0).Mul(YieldFractionDecimalPrecision, r.config.Chain.BlocksPerYear))
	return
}

func (r *Rewards) calculateTotalPeriodReward(totalStake *big.Int) *big.Int {
	if r.config.Chain.Hardforks.AspenHf.BlockNumPartTwo >= r.blockNum {
		current_supply := r.storage.GetTotalSupply()

		blockReward, _ := r.calculateBlockReward(totalStake, current_supply)

		return blockReward
	} else {
		// Original fixed yield curve
		blockReward := big.NewInt(0).Mul(totalStake, r.config.Chain.YieldPercentage)
		blockReward.Div(blockReward, big.NewInt(0).Mul(big.NewInt(100), r.config.Chain.BlocksPerYear))
		return blockReward
	}
}

func (r *Rewards) rewardsFromStats(totalStake *big.Int, stats *stats) (rewards PeriodRewards) {
	totalRewards := r.calculateTotalPeriodReward(totalStake)
	totalRewardsParts := calculatePeriodRewardsParts(r.config.Chain, totalRewards, stats.TotalVotesWeight == 0)
	rewards = MakePeriodRewards()
	for addr, s := range stats.ValidatorStats {
		if !r.validators.Exists(addr) {
			continue
		}
		if rewards.ValidatorRewards[addr] == nil {
			rewards.ValidatorRewards[addr] = big.NewInt(0)
		}

		// Add dags reward
		if s.DagBlocksCount > 0 {
			dag_reward := big.NewInt(0)
			dag_reward.Mul(big.NewInt(s.DagBlocksCount), totalRewardsParts.dags)
			dag_reward.Div(dag_reward, big.NewInt(stats.TotalDagCount))
			rewards.TotalReward.Add(rewards.TotalReward, dag_reward)
			rewards.ValidatorRewards[addr].Add(rewards.ValidatorRewards[addr], dag_reward)
		}

		// Add voting reward
		if s.VoteWeight > 0 {
			// total_votes_reward * validator_vote_weight / total_votes_weight
			vote_reward := big.NewInt(0).Mul(totalRewardsParts.votes, big.NewInt(s.VoteWeight))
			vote_reward.Div(vote_reward, big.NewInt(stats.TotalVotesWeight))
			rewards.TotalReward.Add(rewards.TotalReward, vote_reward)
			rewards.ValidatorRewards[addr].Add(rewards.ValidatorRewards[addr], vote_reward)
		}

		if s.FeeReward != nil && s.FeeReward.Cmp(big.NewInt(0)) > 0 {
			rewards.BlockFee.Add(rewards.BlockFee, s.FeeReward)
		}
	}
	blockAuthorReward := big.NewInt(0)
	{
		maxVotesWeight := Max(stats.MaxVotesWeight, stats.TotalVotesWeight)
		// In case all reward votes are included we will just pass block author whole reward, this should improve rounding issues
		if maxVotesWeight == stats.TotalVotesWeight {
			blockAuthorReward = totalRewardsParts.bonus
		} else {
			twoTPlusOne := maxVotesWeight*2/3 + 1
			bonusVotesWeight := int64(0)
			bonusVotesWeight = stats.TotalVotesWeight - twoTPlusOne
			// should be zero if rewardsStats.TotalVotesWeight == twoTPlusOne
			blockAuthorReward.Div(big.NewInt(0).Mul(totalRewardsParts.bonus, big.NewInt(bonusVotesWeight)), big.NewInt(maxVotesWeight-twoTPlusOne))
		}
	}
	if blockAuthorReward.Cmp(big.NewInt(0)) > 0 {
		if rewards.ValidatorRewards[r.blockAuthor] == nil {
			rewards.ValidatorRewards[r.blockAuthor] = big.NewInt(0)
		}
		rewards.TotalReward.Add(rewards.TotalReward, blockAuthorReward)
		rewards.ValidatorRewards[r.blockAuthor].Add(rewards.ValidatorRewards[r.blockAuthor], blockAuthorReward)
	}
	return
}

func CalculateTotalStake(validators *Validators) *big.Int {
	totalStake := big.NewInt(0)
	for _, v := range validators.validators {
		totalStake.Add(totalStake, v.TotalStake)
	}
	return totalStake
}

func GetValidatorsYield(rewards map[string]*big.Int, validators *Validators) []storage.ValidatorYield {
	ret := make([]storage.ValidatorYield, 0, len(validators.validators))
	for v_addr, v := range validators.validators {
		if v.TotalStake.Cmp(big.NewInt(0)) == 0 {
			continue
		}
		if rewards[v_addr] != nil {
			ret = append(ret, storage.ValidatorYield{Validator: v_addr, Yield: GetMultipliedYield(rewards[v_addr], v.TotalStake)})
		}
	}

	return ret
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

func (r *Rewards) processIntervalYield(batch storage.Batch) {
	sum := big.NewInt(0)
	storage.ProcessIntervalData(r.storage, r.blockNum-r.config.TotalYieldSavingInterval, func(key []byte, o storage.MultipliedYield) (stop bool) {
		sum.Add(sum, o.Yield)
		batch.Remove([]byte(key))
		return false
	})

	yield := GetYieldForInterval(sum, r.config.Chain.BlocksPerYear, int64(r.config.TotalYieldSavingInterval))
	log.WithFields(log.Fields{"total_yield": yield}).Info("processIntervalYield")
	batch.AddToBatchSingleKey(&storage.Yield{Yield: common.FormatFloat(yield)}, storage.FormatIntToKey(r.blockNum))
}

func (r *Rewards) processValidatorsIntervalYield(batch storage.Batch) {
	start := uint64(0)
	if r.blockNum > r.config.ValidatorsYieldSavingInterval {
		start = r.blockNum - r.config.ValidatorsYieldSavingInterval
	}

	sum_by_validator := make(map[string]*big.Int)

	storage.ProcessIntervalData(r.storage, start, func(key []byte, o storage.ValidatorsYield) (stop bool) {
		for _, y := range o.Yields {
			if sum_by_validator[y.Validator] == nil {
				sum_by_validator[y.Validator] = big.NewInt(0)
			}
			sum_by_validator[y.Validator].Add(sum_by_validator[y.Validator], y.Yield)
		}
		batch.Remove(key)
		return false
	})

	log.WithFields(log.Fields{"validators": len(sum_by_validator)}).Info("processValidatorsIntervalYield")
	yields := make([]oracle.RawValidator, 0, len(sum_by_validator))
	for val, sum := range sum_by_validator {
		yield := GetYieldForInterval(sum, r.config.Chain.BlocksPerYear, int64(r.config.ValidatorsYieldSavingInterval))
		log.WithFields(log.Fields{"validator": val, "yield": yield}).Info("processValidatorsIntervalYield")
		batch.AddToBatch(&storage.Yield{Yield: common.FormatFloat(yield)}, val, r.blockNum)
		yields = append(yields, oracle.RawValidator{Yield: common.FormatFloat(yield), Address: ethcommon.HexToAddress(val)})
	}

	r.oracle.PushValidators(yields)
}

func GetYieldForInterval(yields_sum, blocks_per_year *big.Int, elem_count int64) float64 {
	res := big.NewInt(0)
	res.Mul(yields_sum, blocks_per_year)
	res.Mul(res, percentage_multiplier)
	res.Div(res, big.NewInt(int64(elem_count)))
	res.Div(res, multiplier)

	ret := float64(res.Uint64())
	ret /= float64(percentage_multiplier.Uint64())
	return ret
}

func GetMultipliedYield(reward, stake *big.Int) *big.Int {
	r := big.NewInt(0)
	r.Mul(reward, multiplier)
	r.Div(r, stake)

	return r
}
