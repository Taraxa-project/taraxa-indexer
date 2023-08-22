package rewards

import (
	"math/big"
	"strings"

	"github.com/Taraxa-project/taraxa-go-client/taraxa_client/dpos_contract_client/dpos_interface"
	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/models"
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

type Rewards struct {
	storage    storage.Storage
	batch      storage.Batch
	config     *common.Config
	validators *Validators

	blockNum    uint64
	blockAuthor string
}

func MakeRewards(storage storage.Storage, batch storage.Batch, config *common.Config, block *models.Pbft, validators []dpos_interface.DposInterfaceValidatorData) *Rewards {
	r := Rewards{storage, batch, config, MakeValidators(config, validators), block.Number, strings.ToLower(block.Author)}
	return &r
}

func (r *Rewards) Process(total_minted *big.Int, dags []chain.DagBlock, trxs []models.Transaction, votes chain.VotesResponse) {
	r.addTotalMinted(total_minted)

	totalStake := CalculateTotalStake(r.validators)
	if r.blockNum%r.config.TotalYieldSavingInterval == 0 {
		log.WithFields(log.Fields{"total_stake": totalStake}).Info("totalStake")
	}
	rewards, total_reward := r.calculateValidatorsRewards(dags, votes, trxs, totalStake)
	validators_yield := GetValidatorsYield(rewards, r.validators)
	if total_reward.Cmp(total_minted) != 0 {
		log.WithFields(log.Fields{"total_reward_check": total_reward, "total_minted": total_minted}).Fatal("Total reward check failed")
	}
	r.batch.AddToBatchSingleKey(storage.ValidatorsYield{Yields: validators_yield}, storage.FormatIntToKey(r.blockNum))
	r.batch.AddToBatchSingleKey(storage.MultipliedYield{Yield: GetMultipliedYield(total_minted, totalStake)}, storage.FormatIntToKey(r.blockNum))
}

func (r *Rewards) addTotalMinted(amount *big.Int) {
	current := r.storage.GetTotalSupply()
	current.Add(current, amount)

	r.batch.SetTotalSupply(current)
}

func (r *Rewards) calculateValidatorsRewards(dags []chain.DagBlock, votes chain.VotesResponse, trxs []models.Transaction, totalStake *big.Int) (map[string]*big.Int, *big.Int) {
	stats := makeStats(dags, votes, trxs, r.config.Chain.CommitteeSize.Int64())
	return r.rewardsFromStats(totalStake, stats)
}

func (r *Rewards) rewardsFromStats(totalStake *big.Int, stats *stats) (rewards map[string]*big.Int, total_reward *big.Int) {
	rewards = make(map[string]*big.Int)
	total_reward = big.NewInt(0)

	totalRewards := calculateTotalRewards(r.config.Chain, totalStake, stats.TotalVotesWeight == 0)
	for addr, s := range stats.ValidatorStats {
		if !r.validators.Exists(addr) {
			continue
		}
		if rewards[addr] == nil {
			rewards[addr] = big.NewInt(0)
		}

		// Add dags reward
		if s.dagBlocksCount > 0 {
			dag_reward := big.NewInt(0)
			dag_reward.Mul(big.NewInt(s.dagBlocksCount), totalRewards.dags)
			dag_reward.Div(dag_reward, big.NewInt(stats.TotalDagCount))
			total_reward.Add(total_reward, dag_reward)
			rewards[addr].Add(rewards[addr], dag_reward)
		}

		// Add voting reward
		if s.voteWeight > 0 {
			// total_votes_reward * validator_vote_weight / total_votes_weight
			vote_reward := big.NewInt(0).Mul(totalRewards.votes, big.NewInt(s.voteWeight))
			vote_reward.Div(vote_reward, big.NewInt(stats.TotalVotesWeight))
			total_reward.Add(total_reward, vote_reward)
			rewards[addr].Add(rewards[addr], vote_reward)
		}
	}
	blockAuthorReward := big.NewInt(0)
	{
		maxVotesWeight := Max(stats.MaxVotesWeight, stats.TotalVotesWeight)
		// In case all reward votes are included we will just pass block author whole reward, this should improve rounding issues
		if maxVotesWeight == stats.TotalVotesWeight {
			blockAuthorReward = totalRewards.bonus
		} else {
			twoTPlusOne := maxVotesWeight*2/3 + 1
			bonusVotesWeight := int64(0)
			bonusVotesWeight = stats.TotalVotesWeight - twoTPlusOne
			// should be zero if rewardsStats.TotalVotesWeight == twoTPlusOne
			blockAuthorReward.Div(big.NewInt(0).Mul(totalRewards.bonus, big.NewInt(bonusVotesWeight)), big.NewInt(maxVotesWeight-twoTPlusOne))
		}
	}
	if blockAuthorReward.Cmp(big.NewInt(0)) > 0 {
		if rewards[r.blockAuthor] == nil {
			rewards[r.blockAuthor] = big.NewInt(0)
		}
		total_reward.Add(total_reward, blockAuthorReward)
		rewards[r.blockAuthor].Add(rewards[r.blockAuthor], blockAuthorReward)
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
		} else if validators.IsEligible(v_addr) {
			ret = append(ret, storage.ValidatorYield{Validator: v_addr, Yield: big.NewInt(0)})
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
	storage.ProcessIntervalData(r.storage, r.blockNum-r.config.TotalYieldSavingInterval, func(key string, o storage.MultipliedYield) (stop bool) {
		sum.Add(sum, o.Yield)
		batch.Remove(key)
		return false
	})

	yield := GetYieldForInterval(sum, r.config.Chain.BlocksPerYear, int64(r.config.TotalYieldSavingInterval))
	log.WithFields(log.Fields{"total_yield": yield}).Info("processIntervalYield")
	batch.AddToBatchSingleKey(&storage.Yield{Yield: common.FormatFloat(yield)}, storage.FormatIntToKey(r.blockNum))
}

func (r *Rewards) processValidatorsIntervalYield(batch storage.Batch) {
	start := uint64(0)
	if r.blockNum > r.config.TotalYieldSavingInterval {
		start = r.blockNum - r.config.TotalYieldSavingInterval
	}

	sum_by_validator := make(map[string]*big.Int)
	count_by_validator := make(map[string]int64)

	storage.ProcessIntervalData(r.storage, start, func(key string, o storage.ValidatorsYield) (stop bool) {
		for _, y := range o.Yields {
			if sum_by_validator[y.Validator] == nil {
				sum_by_validator[y.Validator] = big.NewInt(0)
			}
			sum_by_validator[y.Validator].Add(sum_by_validator[y.Validator], y.Yield)
			count_by_validator[y.Validator]++
		}
		batch.Remove(string(key))
		return false
	})

	for val, sum := range sum_by_validator {
		yield := GetYieldForInterval(sum, r.config.Chain.BlocksPerYear, count_by_validator[val])
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
