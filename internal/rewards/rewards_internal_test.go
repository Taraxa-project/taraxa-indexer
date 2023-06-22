package rewards

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage/pebble"
	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/stretchr/testify/assert"
)

type AddressCount map[string]int

func makeTransactions(count int) (trxs []*models.Transaction) {
	for i := 0; i < count; i++ {
		trxs = append(trxs, &models.Transaction{Hash: fmt.Sprintf("0x%x", i)})
	}
	return
}

func makeDags(ac AddressCount) (dags []*chain.DagBlock) {
	total_count := 0
	for addr, c := range ac {
		for i := 0; i < c; i++ {
			dags = append(dags, &chain.DagBlock{Dag: models.Dag{Sender: addr, Hash: fmt.Sprintf("0x%x", total_count)}, Transactions: []string{fmt.Sprintf("0x%x", total_count)}})
			total_count++
		}
	}
	return
}

func makeVotes(ac AddressCount) (votes *chain.VotesResponse) {
	votes = new(chain.VotesResponse)
	votes.Votes = make([]chain.Vote, 0)
	total_weight := int64(0)
	for addr, weight := range ac {
		votes.Votes = append(votes.Votes, chain.Vote{Voter: addr, Weight: fmt.Sprintf("0x%x", weight)})
		total_weight += int64(weight)
	}
	votes.PeriodTotalVotesCount = total_weight
	return
}

func makeTestConfig() (config *common.Config) {
	config = common.DefaultConfig()
	config.Chain.BlocksPerYear = big.NewInt(1)
	config.Chain.YieldPercentage = big.NewInt(100)

	return
}

func TestMakeStats(t *testing.T) {
	trxs := makeTransactions(6)
	dags := makeDags(AddressCount{"0x1": 1, "0x2": 2, "0x3": 3})
	votes := makeVotes(AddressCount{"0x1": 1, "0x2": 2, "0x3": 3})
	assert.Equal(t, 6, len(trxs))
	assert.Equal(t, 6, len(dags))
	assert.Equal(t, 3, len(votes.Votes))
	assert.Equal(t, int64(6), votes.PeriodTotalVotesCount)

	s := makeStats(dags, votes, trxs, 100)
	assert.Equal(t, 3, len(s.ValidatorStats))
	assert.Equal(t, 6, int(s.TotalDagCount))
	assert.Equal(t, 6, int(s.TotalVotesWeight))
}

func TestCalculateTotalRewards(t *testing.T) {
	config := makeTestConfig().Chain

	totalStake := big.NewInt(1000000000)
	blockReward := big.NewInt(1000000000)
	blockReward.Mul(blockReward, config.YieldPercentage)
	blockReward.Div(blockReward, big.NewInt(100))
	blockReward.Div(blockReward, config.BlocksPerYear)

	dags_reward := big.NewInt(0)
	dags_reward.Mul(blockReward, config.DagProposersReward).Div(dags_reward, big.NewInt(100))

	votes_rewards_part := big.NewInt(100)
	votes_rewards_part.Sub(votes_rewards_part, config.DagProposersReward)
	votes_rewards_part.Sub(votes_rewards_part, config.MaxBlockAuthorReward)

	votes_reward := big.NewInt(0)
	votes_reward.Mul(blockReward, votes_rewards_part).Div(votes_reward, big.NewInt(100))

	totalRewards := calculateTotalRewards(config, totalStake)
	assert.Equal(t, dags_reward, totalRewards.dags)
	assert.Equal(t, votes_reward, totalRewards.votes)
}

func TestRewards(t *testing.T) {
	config := makeTestConfig()

	mc := chain.MakeMockClient()
	st := pebble.NewStorage("")

	r := MakeRewards(st, st.NewBatch(), config, mc, 1, "0x4")

	trxs := makeTransactions(5)
	dags := makeDags(AddressCount{"0x1": 1, "0x2": 2, "0x3": 2})
	votes := makeVotes(AddressCount{"0x1": 1, "0x2": 2, "0x3": 2})
	assert.Equal(t, 5, len(dags))
	assert.Equal(t, 3, len(votes.Votes))
	assert.Equal(t, 5, len(trxs))
	totalStake := big.NewInt(1000000000000)
	rewards := r.calculateValidatorsRewards(dags, votes, trxs, totalStake)
	assert.Equal(t, 4, len(rewards))
	// Calculate total reward for the block
	total_reward := totalStake.Int64() * config.Chain.YieldPercentage.Int64() / 100 / config.Chain.BlocksPerYear.Int64()
	// Calculate reward for DAG proposer
	reward1_dag_part := (total_reward * config.Chain.DagProposersReward.Int64() / 100) / int64(len(dags))
	// Calculate reward for voter
	reward1_vote_part := (total_reward * (100 - config.Chain.DagProposersReward.Int64() - config.Chain.MaxBlockAuthorReward.Int64()) / 100) / votes.PeriodTotalVotesCount
	// Calculate total reward for validator
	reward1 := big.NewInt(reward1_dag_part + reward1_vote_part)
	assert.Equal(t, reward1, rewards["0x1"])
	// validator 2 and 3 should have the same reward that is two times bigger than reward1, because they have two times more dags and votes
	assert.Equal(t, big.NewInt(0).Mul(reward1, big.NewInt(2)), rewards["0x2"])
	assert.Equal(t, big.NewInt(0).Mul(reward1, big.NewInt(2)), rewards["0x3"])
}

func TestRewardsWithNodeData(t *testing.T) {
	config := common.DefaultConfig()

	TaraPrecision := big.NewInt(1e+18)
	DefaultMinimumDeposit := big.NewInt(0).Mul(big.NewInt(1000), TaraPrecision)

	// Simulated rewards statistics
	validator1_addr := "0x1"
	validator2_addr := "0x2"
	validator4_addr := "0x4"
	validator5_addr := "0x5"
	mc := chain.MakeMockClient()
	st := pebble.NewStorage("")
	r := MakeRewards(st, st.NewBatch(), config, mc, 1, "0x3")
	total_stake := big.NewInt(0).Mul(DefaultMinimumDeposit, big.NewInt(8))
	{
		rewardsStats := stats{}
		rewardsStats.ValidatorStats = map[string]validatorStats{
			validator1_addr: {dagBlocksCount: 8, voteWeight: 1},
			validator2_addr: {dagBlocksCount: 32, voteWeight: 5},
			validator4_addr: {voteWeight: 1},
		}
		rewardsStats.TotalDagCount = 40
		rewardsStats.TotalVotesWeight = 7
		rewardsStats.MaxVotesWeight = 8

		// Expected block reward
		totalRewards := calculateTotalRewards(r.Config.Chain, total_stake)
		fmt.Println("votes", totalRewards.votes, "dags", totalRewards.dags, "bonus", totalRewards.bonus)
		rewards := r.rewardsFromStats(total_stake, &rewardsStats)
		// We have 1 out of 2 bonus votes, so block author should get half of the bonus reward
		assert.Equal(t, big.NewInt(0).Div(totalRewards.bonus, big.NewInt(2)), rewards[r.blockAuthor])

		// data from node test
		expected_validator1_commission_reward := int64(31890990795100)
		expected_validator2_commission_reward := int64(139160687105891)
		expected_validator4_commission_reward := int64(11596723925491)
		assert.Equal(t, expected_validator1_commission_reward, rewards[validator1_addr].Int64())
		assert.Equal(t, expected_validator2_commission_reward, rewards[validator2_addr].Int64())
		assert.Equal(t, expected_validator4_commission_reward, rewards[validator4_addr].Int64())
	}

	{
		rewardsStats := stats{}
		rewardsStats.ValidatorStats = map[string]validatorStats{
			validator1_addr: {dagBlocksCount: 15, voteWeight: 3},
			validator2_addr: {dagBlocksCount: 35, voteWeight: 5},
			validator4_addr: {voteWeight: 2},
		}
		rewardsStats.TotalDagCount = 50
		rewardsStats.TotalVotesWeight = 10
		rewardsStats.MaxVotesWeight = 13

		// Expected block reward
		totalRewards := calculateTotalRewards(r.Config.Chain, total_stake)
		rewards := r.rewardsFromStats(total_stake, &rewardsStats)
		// We have 1 out of 4 bonus votes, so block author should get 1/4 of the bonus reward
		assert.Equal(t, big.NewInt(0).Div(totalRewards.bonus, big.NewInt(4)), rewards[r.blockAuthor])
		assert.Equal(t, big.NewInt(5073566717402), rewards[r.blockAuthor])
		// data from node test
		expected_validator1_commission_reward := int64(54794520547944)
		expected_validator2_commission_reward := int64(111618467782851)
		expected_validator4_commission_reward := int64(16235413495687)
		assert.Equal(t, expected_validator1_commission_reward, rewards[validator1_addr].Int64())
		assert.Equal(t, expected_validator2_commission_reward, rewards[validator2_addr].Int64())
		assert.Equal(t, expected_validator4_commission_reward, rewards[validator4_addr].Int64())
	}

	{
		rewardsStats := stats{}
		rewardsStats.ValidatorStats = map[string]validatorStats{
			validator1_addr: {dagBlocksCount: 10, voteWeight: 5},
			validator2_addr: {dagBlocksCount: 10, voteWeight: 5},
			validator4_addr: {dagBlocksCount: 10, voteWeight: 5},
			validator5_addr: {dagBlocksCount: 10, voteWeight: 5},
		}
		rewardsStats.TotalDagCount = 40
		rewardsStats.TotalVotesWeight = 20
		rewardsStats.MaxVotesWeight = 24

		// Expected block reward
		rewards := r.rewardsFromStats(total_stake, &rewardsStats)
		// We have 1 out of 4 bonus votes, so block author should get 1/4 of the bonus reward
		// data from node test
		expected_block_author_reward := int64(8697542944118)
		expected_validator_reward := int64(45662100456620)
		assert.Equal(t, expected_block_author_reward, rewards[r.blockAuthor].Int64())
		assert.Equal(t, expected_validator_reward, rewards[validator1_addr].Int64())
		assert.Equal(t, expected_validator_reward, rewards[validator2_addr].Int64())
		assert.Equal(t, expected_validator_reward, rewards[validator4_addr].Int64())
		assert.Equal(t, expected_validator_reward, rewards[validator4_addr].Int64())
	}

	// Block author is validator 1
	{
		r.blockAuthor = "0x1"
		rewardsStats := stats{}
		rewardsStats.ValidatorStats = map[string]validatorStats{
			validator1_addr: {dagBlocksCount: 10, voteWeight: 5},
			validator2_addr: {dagBlocksCount: 10, voteWeight: 5},
			validator4_addr: {dagBlocksCount: 10, voteWeight: 5},
			validator5_addr: {dagBlocksCount: 10, voteWeight: 5},
		}
		rewardsStats.TotalDagCount = 40
		rewardsStats.TotalVotesWeight = 20
		rewardsStats.MaxVotesWeight = 24

		// Expected block reward
		r.blockAuthor = validator1_addr
		rewards := r.rewardsFromStats(total_stake, &rewardsStats)
		// We have 1 out of 4 bonus votes, so block author should get 1/4 of the bonus reward
		// data from node test
		expected_block_author_reward := int64(8697542944118)
		expected_validator_reward := int64(45662100456620)
		assert.Equal(t, expected_validator_reward+expected_block_author_reward, rewards[validator1_addr].Int64())
		assert.Equal(t, expected_validator_reward, rewards[validator2_addr].Int64())
		assert.Equal(t, expected_validator_reward, rewards[validator4_addr].Int64())
		assert.Equal(t, expected_validator_reward, rewards[validator4_addr].Int64())
	}
}
