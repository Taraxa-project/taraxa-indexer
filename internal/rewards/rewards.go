package rewards

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/models"
	log "github.com/sirupsen/logrus"
)

var multiplier = big.NewInt(10 ^ 12)

type Rewards struct {
	Storage storage.Storage
	Batch   storage.Batch
	Config  *common.Config
	Client  chain.Client

	blockNum    uint64
	blockAuthor string
}

func MakeRewards(storage storage.Storage, batch storage.Batch, config *common.Config, client chain.Client, block_num uint64, block_author string) *Rewards {
	r := Rewards{storage, batch, config, client, block_num, block_author}
	return &r
}

func (r *Rewards) addTotalMinted(amount *big.Int) {
	current := r.Storage.GetTotalSupply()
	current.Add(current, amount)

	r.Batch.SetTotalSupply(current)
}

func (r *Rewards) Process(total_minted *big.Int, dags []*chain.DagBlock, trxs []*models.Transaction, votes *chain.VotesResponse) {
	r.addTotalMinted(total_minted)
	fmt.Println("Process rewards at", r.blockNum)
	validators, err := r.Client.GetValidatorsAtBlock(big.NewInt(0).SetUint64(r.blockNum))
	if err != nil {
		log.WithError(err).Fatal("Error GetValidators")
	}

	totalStake := big.NewInt(0)
	for _, v := range validators {
		totalStake.Add(totalStake, v.Info.TotalStake)
	}
	rewards := r.calculateValidatorsRewards(dags, votes, trxs, totalStake)
	for k, v := range rewards {
		fmt.Println("rewards", k, v)
	}
	total_reward_check := big.NewInt(0)
	for _, v := range validators {
		validator := strings.ToLower(v.Account.Hex())
		validator_reward := rewards[validator]
		if validator_reward == nil {
			continue
		}
		total_reward_check.Add(total_reward_check, validator_reward)
		r.saveValidatorYield(validator, getMultipliedYield(validator_reward, v.Info.TotalStake))

		totalStake.Add(totalStake, v.Info.TotalStake)
	}
	if total_reward_check.Cmp(total_minted) != 0 {
		log.WithFields(log.Fields{"total_reward_check": total_reward_check, "total_minted": total_minted}).Fatal("Total reward check failed")
	}
	r.saveTotalYield(getMultipliedYield(total_minted, totalStake))
}

func (r *Rewards) saveTotalYield(yield *storage.Yield) {
	// r.Batch.AddToBatchSingleKey(yield, fmt.Sprint(r.blockNum))
}

func (r *Rewards) saveValidatorYield(validator string, yield *storage.Yield) {
	// r.Batch.AddToBatch(yield, validator, r.blockNum)
}

type ValidatorsRewards map[string]*big.Int

func (r *Rewards) calculateValidatorsRewards(dags []*chain.DagBlock, votes *chain.VotesResponse, trxs []*models.Transaction, totalStake *big.Int) ValidatorsRewards {
	stats := makeStats(dags, votes, trxs, r.Config.Chain.CommitteeSize.Int64())
	fmt.Println("stats", stats.MaxVotesWeight, stats.TotalVotesWeight, stats.TotalDagCount)
	for addr, s := range stats.ValidatorStats {
		fmt.Println("stats", addr, s)
	}
	return r.rewardsFromStats(totalStake, stats)
}

func (r *Rewards) rewardsFromStats(totalStake *big.Int, stats *stats) (rewards ValidatorsRewards) {
	rewards = make(ValidatorsRewards)

	totalRewards := calculateTotalRewards(r.Config.Chain, totalStake)
	for addr, s := range stats.ValidatorStats {
		if rewards[addr] == nil {
			rewards[addr] = big.NewInt(0)
		}

		// Add dags reward
		if s.dagBlocksCount > 0 {
			dag_reward := big.NewInt(0)
			dag_reward.Mul(big.NewInt(s.dagBlocksCount), totalRewards.dags)
			dag_reward.Div(dag_reward, big.NewInt(stats.TotalDagCount))
			fmt.Println("dag_reward", dag_reward, s.dagBlocksCount, totalRewards.dags, stats.TotalDagCount)
			rewards[addr].Add(rewards[addr], dag_reward)
		}

		// Add voting reward
		if s.voteWeight > 0 {
			// total_votes_reward * validator_vote_weight / total_votes_weight
			vote_reward := big.NewInt(0).Mul(totalRewards.votes, big.NewInt(s.voteWeight))
			vote_reward.Div(vote_reward, big.NewInt(stats.TotalVotesWeight))
			fmt.Println("vote_reward", vote_reward, totalRewards.votes, s.voteWeight, stats.TotalVotesWeight)
			rewards[addr].Add(rewards[addr], vote_reward)
		}
	}
	blockAuthorReward := big.NewInt(0)
	{
		maxVotesWeight := Max(stats.MaxVotesWeight, stats.TotalVotesWeight)
		fmt.Println("maxVotesWeight", maxVotesWeight, stats.MaxVotesWeight, stats.TotalVotesWeight)
		// In case all reward votes are included we will just pass block author whole reward, this should improve rounding issues
		if maxVotesWeight == stats.TotalVotesWeight {
			blockAuthorReward = totalRewards.bonus
		} else {
			twoTPlusOne := maxVotesWeight*2/3 + 1
			fmt.Println("twoTPlusOne", twoTPlusOne)
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
		rewards[r.blockAuthor].Add(rewards[r.blockAuthor], blockAuthorReward)
	}
	return
}

func getMultipliedYield(reward, total *big.Int) *storage.Yield {
	r := big.NewInt(0)
	r.Mul(reward, multiplier)
	r.Div(r, total)

	return &storage.Yield{Yield: r}
}
