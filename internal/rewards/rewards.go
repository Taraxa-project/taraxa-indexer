package rewards

import (
	"math/big"
	"strings"

	"github.com/Taraxa-project/taraxa-go-client/taraxa_client/dpos_contract_client/dpos_interface"
	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage/pebble"
	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/ethereum/go-ethereum/rlp"
	log "github.com/sirupsen/logrus"
)

var multiplier = big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)
var percentage_multiplier = big.NewInt(10000)

type Validators struct {
	config     *common.Config
	validators map[string]dpos_interface.DposInterfaceValidatorBasicInfo
}

func MakeValidators(config *common.Config, validators []dpos_interface.DposInterfaceValidatorData) *Validators {
	v := Validators{config, make(map[string]dpos_interface.DposInterfaceValidatorBasicInfo)}
	for _, val := range validators {
		v.validators[strings.ToLower(val.Account.Hex())] = val.Info
	}
	return &v
}

func (v *Validators) IsEligible(address string) bool {
	validator, ok := v.validators[strings.ToLower(address)]
	if ok {
		return v.config.IsEligible(validator.TotalStake)
	}
	return false
}

func (v *Validators) Exists(address string) bool {
	_, ok := v.validators[strings.ToLower(address)]
	return ok
}

type PeriodRewards struct {
	ValidatorRewards map[string]*big.Int
	TotalReward      *big.Int
	BlockFee         *big.Int
}

func (p *PeriodRewards) ToStorage() storage.PeriodRewards {
	sr := storage.PeriodRewards{TotalReward: p.TotalReward, BlockFee: p.BlockFee}
	sr.ValidatorRewards = make([]storage.ValidatorReward, 0, len(p.ValidatorRewards))
	for addr, reward := range p.ValidatorRewards {
		sr.ValidatorRewards = append(sr.ValidatorRewards, storage.ValidatorReward{Validator: addr, Reward: reward})
	}
	return sr
}

func MakePeriodRewards() (r PeriodRewards) {
	r.ValidatorRewards = make(map[string]*big.Int)
	r.TotalReward = big.NewInt(0)
	return
}

type Rewards struct {
	storage    storage.Storage
	batch      storage.Batch
	config     *common.Config
	validators *Validators

	blockNum    uint64
	blockAuthor string
	blockFee    *big.Int
}

func MakeRewards(storage storage.Storage, batch storage.Batch, config *common.Config, block *models.Pbft, blockFee *big.Int, validators []dpos_interface.DposInterfaceValidatorData) *Rewards {
	r := Rewards{storage, batch, config, MakeValidators(config, validators), block.Number, strings.ToLower(block.Author), blockFee}
	return &r
}

func (r *Rewards) Process(total_minted *big.Int, dags []chain.DagBlock, trxs []models.Transaction, votes chain.VotesResponse) (blockFee *big.Int) {
	totalStake := CalculateTotalStake(r.validators)
	if r.blockNum%r.config.TotalYieldSavingInterval == 0 {
		log.WithFields(log.Fields{"total_stake": totalStake}).Info("totalStake")
	}

	rewards := r.calculateValidatorsRewards(dags, votes, trxs, totalStake)
	totalReward, totalBlockFee := r.ProcessRewards(rewards, total_minted, totalStake)

	if totalReward.Cmp(totalReward) != 0 {
		log.WithFields(log.Fields{"period": r.blockNum, "total_reward_check": rewards.TotalReward, "total_minted": total_minted}).Fatal("Total reward check failed")
	}
	r.addTotalMinted(totalReward)
	return totalBlockFee
}

func (r *Rewards) ProcessRewards(periodRewards PeriodRewards, total_minted *big.Int, totalStake *big.Int) (*big.Int, *big.Int) {
	if r.blockNum > r.config.Chain.Hardforks.MagnoliaHf.BlockNum {
		distributionFrequency := r.config.Chain.Hardforks.GetDistributionFrequency(r.blockNum)
		// distribute it right away
		log.WithFields(log.Fields{"period": r.blockNum, "distribution_frequency": distributionFrequency}).Info("Distributing rewards")
		if distributionFrequency == 1 {
		} else if r.blockNum%uint64(distributionFrequency) == 0 {
			// distribute whole interval rewards
			periodRewards.BlockFee = r.blockFee
			fromKey := storage.FormatIntToKey(r.blockNum - uint64(distributionFrequency))
			r.storage.ForEachFromKey([]byte(pebble.GetPrefix(storage.PeriodRewards{})), []byte(fromKey), func(key, res []byte) (stop bool) {
				var pr storage.PeriodRewards
				err := rlp.DecodeBytes(res, &pr)
				if err != nil {
					log.WithError(err).Fatal("Error decoding data from db")
				}
				log.WithFields(log.Fields{"key": string(key), "blockFee": pr.BlockFee.String()}).Info("Reward from db")
				for _, reward := range pr.ValidatorRewards {
					if periodRewards.ValidatorRewards[reward.Validator] == nil {
						periodRewards.ValidatorRewards[reward.Validator] = big.NewInt(0)
					}
					periodRewards.ValidatorRewards[reward.Validator].Add(periodRewards.ValidatorRewards[reward.Validator], reward.Reward)
				}
				periodRewards.TotalReward.Add(periodRewards.TotalReward, pr.TotalReward)
				periodRewards.BlockFee.Add(periodRewards.BlockFee, pr.BlockFee)
				r.batch.Remove(key)
				return false
			})
		} else {
			// save rewards for this period
			periodRewards.BlockFee = r.blockFee
			r.batch.AddToBatchSingleKey(periodRewards.ToStorage(), storage.FormatIntToKey(r.blockNum))
			periodRewards.BlockFee = nil
			log.WithFields(log.Fields{"period": r.blockNum, "distribution_frequency": distributionFrequency, "key": storage.FormatIntToKey(r.blockNum)}).Info("Saving rewards")
		}
	}

	if periodRewards.TotalReward.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0), nil
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
	trxs []models.Transaction, totalStake *big.Int) PeriodRewards {
	stats := makeStats(dags, votes, trxs, r.config.Chain.CommitteeSize.Int64())
	return r.rewardsFromStats(totalStake, stats)
}

func (r *Rewards) rewardsFromStats(totalStake *big.Int, stats *stats) (periodRewards PeriodRewards) {
	totalPeriodRewards := calculateTotalPeriodRewards(r.config.Chain, totalStake, stats.TotalVotesWeight == 0)
	periodRewards = MakePeriodRewards()
	for addr, s := range stats.ValidatorStats {
		if !r.validators.Exists(addr) {
			continue
		}
		if periodRewards.ValidatorRewards[addr] == nil {
			periodRewards.ValidatorRewards[addr] = big.NewInt(0)
		}

		// Add dags reward
		if s.DagBlocksCount > 0 {
			dag_reward := big.NewInt(0)
			dag_reward.Mul(big.NewInt(s.DagBlocksCount), totalPeriodRewards.dags)
			dag_reward.Div(dag_reward, big.NewInt(stats.TotalDagCount))
			periodRewards.TotalReward.Add(periodRewards.TotalReward, dag_reward)
			periodRewards.ValidatorRewards[addr].Add(periodRewards.ValidatorRewards[addr], dag_reward)
		}

		// Add voting reward
		if s.VoteWeight > 0 {
			// total_votes_reward * validator_vote_weight / total_votes_weight
			vote_reward := big.NewInt(0).Mul(totalPeriodRewards.votes, big.NewInt(s.VoteWeight))
			vote_reward.Div(vote_reward, big.NewInt(stats.TotalVotesWeight))
			periodRewards.TotalReward.Add(periodRewards.TotalReward, vote_reward)
			periodRewards.ValidatorRewards[addr].Add(periodRewards.ValidatorRewards[addr], vote_reward)
		}
	}
	blockAuthorReward := big.NewInt(0)
	{
		maxVotesWeight := Max(stats.MaxVotesWeight, stats.TotalVotesWeight)
		// In case all reward votes are included we will just pass block author whole reward, this should improve rounding issues
		if maxVotesWeight == stats.TotalVotesWeight {
			blockAuthorReward = totalPeriodRewards.bonus
		} else {
			twoTPlusOne := maxVotesWeight*2/3 + 1
			bonusVotesWeight := int64(0)
			bonusVotesWeight = stats.TotalVotesWeight - twoTPlusOne
			// should be zero if rewardsStats.TotalVotesWeight == twoTPlusOne
			blockAuthorReward.Div(big.NewInt(0).Mul(totalPeriodRewards.bonus, big.NewInt(bonusVotesWeight)), big.NewInt(maxVotesWeight-twoTPlusOne))
		}
	}
	if blockAuthorReward.Cmp(big.NewInt(0)) > 0 {
		if periodRewards.ValidatorRewards[r.blockAuthor] == nil {
			periodRewards.ValidatorRewards[r.blockAuthor] = big.NewInt(0)
		}
		periodRewards.TotalReward.Add(periodRewards.TotalReward, blockAuthorReward)
		periodRewards.ValidatorRewards[r.blockAuthor].Add(periodRewards.ValidatorRewards[r.blockAuthor], blockAuthorReward)
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

	for val, sum := range sum_by_validator {
		yield := GetYieldForInterval(sum, r.config.Chain.BlocksPerYear, int64(r.config.ValidatorsYieldSavingInterval))
		log.WithFields(log.Fields{"validator": val, "yield": yield}).Info("processValidatorsIntervalYield")
		batch.AddToBatch(&storage.Yield{Yield: common.FormatFloat(yield)}, val, r.blockNum)
	}
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
