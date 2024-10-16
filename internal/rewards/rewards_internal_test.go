package rewards

import (
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage/pebble"
	"github.com/Taraxa-project/taraxa-indexer/models"
	ce "github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

type AddressCount map[string]int

func makeTransactions(count int) (trxs []chain.Transaction) {
	for i := 0; i < count; i++ {
		trxs = append(trxs, chain.Transaction{Transaction: models.Transaction{Hash: fmt.Sprintf("0x%x", i)}, GasPrice: 1, GasUsed: 21000})
	}
	return
}

func makeDags(ac AddressCount) (dags []chain.DagBlock) {
	total_count := 0
	for addr, c := range ac {
		for i := 0; i < c; i++ {
			dags = append(dags, chain.DagBlock{Dag: models.Dag{Hash: fmt.Sprintf("0x%x", total_count)}, Sender: addr, Transactions: []string{fmt.Sprintf("0x%x", total_count)}})
			total_count++
		}
	}
	return
}

func makeVotes(ac AddressCount) (votes chain.VotesResponse) {
	votes.Votes = make([]chain.Vote, 0)
	total_weight := uint64(0)
	for addr, weight := range ac {
		votes.Votes = append(votes.Votes, chain.Vote{Voter: addr, Weight: fmt.Sprintf("0x%x", weight)})
		total_weight += uint64(weight)
	}
	votes.PeriodTotalVotesCount = total_weight
	return
}

func makeTestConfig() (config *common.Config) {
	config = common.DefaultConfig()
	config.Chain.BlocksPerYear = big.NewInt(1)
	config.Chain.YieldPercentage = big.NewInt(100)
	config.Chain.EligibilityBalanceThreshold = big.NewInt(1)
	config.Chain.Hardforks.AspenHf.BlockNumPartTwo = 100
	config.Chain.Hardforks.MagnoliaHf.BlockNum = 100

	return
}

func rewardFromStake(config *common.ChainConfig, totalStake *big.Int) *big.Int {
	blockReward := big.NewInt(0).Mul(totalStake, config.YieldPercentage)
	blockReward.Div(blockReward, big.NewInt(0).Mul(big.NewInt(100), config.BlocksPerYear))
	return blockReward
}

func TestMakeStats(t *testing.T) {
	trxs := makeTransactions(6)
	dags := makeDags(AddressCount{"0x1": 1, "0x2": 2, "0x3": 3})
	votes := makeVotes(AddressCount{"0x1": 1, "0x2": 2, "0x3": 3})
	block_author := "0x4"
	assert.Equal(t, 6, len(trxs))
	assert.Equal(t, 6, len(dags))
	assert.Equal(t, 3, len(votes.Votes))
	assert.Equal(t, uint64(6), votes.PeriodTotalVotesCount)

	is_aspen_dag_rewards := false
	s := makeRewardsStats(is_aspen_dag_rewards, dags, votes, trxs, 100, block_author)
	assert.Equal(t, 3, len(s.ValidatorsStats))
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

	totalReward := rewardFromStake(config, totalStake)
	rewardsParts := calculatePeriodRewardsParts(config, totalReward, false)
	assert.Equal(t, dags_reward, rewardsParts.dags)
	assert.Equal(t, votes_reward, rewardsParts.votes)

	totalReward = rewardFromStake(config, totalStake)
	rewardsParts = calculatePeriodRewardsParts(config, totalReward, true)

	assert.Equal(t, blockReward, rewardsParts.dags)
	assert.Equal(t, big.NewInt(0), rewardsParts.votes)
}

func TestRewards(t *testing.T) {
	config := makeTestConfig()
	validator1_addr := strings.ToLower(ce.HexToAddress("0x1").Hex())
	validator2_addr := strings.ToLower(ce.HexToAddress("0x2").Hex())
	validator3_addr := strings.ToLower(ce.HexToAddress("0x3").Hex())
	validator4_addr := strings.ToLower(ce.HexToAddress("0x4").Hex())

	validators_list := []chain.Validator{
		{Address: validator1_addr, TotalStake: big.NewInt(5000000)},
		{Address: validator2_addr, TotalStake: big.NewInt(5000000)},
		{Address: validator3_addr, TotalStake: big.NewInt(5000000)},
		{Address: validator4_addr, TotalStake: big.NewInt(5000000)},
	}

	st := pebble.NewStorage("")
	block := chain.Block{Pbft: models.Pbft{Number: 1, Author: validator4_addr}}
	bd := &chain.BlockData{Pbft: &block, TotalAmountDelegated: big.NewInt(5000000 * 4), TotalSupply: big.NewInt(1), Validators: validators_list}
	r := MakeRewards(st, st.NewBatch(), config, bd)

	trxs := makeTransactions(5)
	dags := makeDags(AddressCount{validator1_addr: 1, validator2_addr: 2, validator3_addr: 2})
	votes := makeVotes(AddressCount{validator1_addr: 1, validator2_addr: 2, validator3_addr: 2})
	assert.Equal(t, 5, len(dags))
	assert.Equal(t, 3, len(votes.Votes))
	assert.Equal(t, 5, len(trxs))
	r.totalStake = big.NewInt(1000000000000)
	stats := r.makeRewardsStats(dags, votes, trxs, block.Author)
	rewards := r.rewardsFromStats(stats)
	assert.Equal(t, 4, len(rewards.ValidatorRewards))
	// Calculate total reward for the block
	total_reward := r.totalStake.Uint64() * config.Chain.YieldPercentage.Uint64() / 100 / config.Chain.BlocksPerYear.Uint64()
	// Calculate reward for DAG proposer
	reward1_dag_part := (total_reward * config.Chain.DagProposersReward.Uint64() / 100) / uint64(len(dags))
	// Calculate reward for voter
	reward1_vote_part := (total_reward * (100 - config.Chain.DagProposersReward.Uint64() - config.Chain.MaxBlockAuthorReward.Uint64()) / 100) / votes.PeriodTotalVotesCount
	// Calculate total reward for validator
	reward1 := big.NewInt(0).SetUint64(reward1_dag_part + reward1_vote_part)
	assert.Equal(t, reward1, rewards.ValidatorRewards[validator1_addr])
	// validator 2 and 3 should have the same reward that is two times bigger than reward1, because they have two times more dags and votes
	assert.Equal(t, big.NewInt(0).Mul(reward1, big.NewInt(2)), rewards.ValidatorRewards[validator2_addr])
	assert.Equal(t, big.NewInt(0).Mul(reward1, big.NewInt(2)), rewards.ValidatorRewards[validator3_addr])
}

func TestRewardsWithNodeData(t *testing.T) {
	config := common.DefaultConfig()
	config.Chain.EligibilityBalanceThreshold = big.NewInt(5000000)
	config.Chain.Hardforks.AspenHf.BlockNumPartTwo = 100
	config.Chain.Hardforks.MagnoliaHf.BlockNum = 100

	st := pebble.NewStorage("")

	TaraPrecision := big.NewInt(1e+18)
	DefaultMinimumDeposit := big.NewInt(0).Mul(big.NewInt(1000), TaraPrecision)

	validator1_addr := strings.ToLower(ce.HexToAddress("0x1").Hex())
	validator2_addr := strings.ToLower(ce.HexToAddress("0x2").Hex())
	validator3_addr := strings.ToLower(ce.HexToAddress("0x3").Hex())
	validator4_addr := strings.ToLower(ce.HexToAddress("0x4").Hex())
	validator5_addr := strings.ToLower(ce.HexToAddress("0x5").Hex())
	validators_list := []chain.Validator{
		{Address: validator1_addr, TotalStake: config.Chain.EligibilityBalanceThreshold},
		{Address: validator2_addr, TotalStake: config.Chain.EligibilityBalanceThreshold},
		{Address: validator3_addr, TotalStake: config.Chain.EligibilityBalanceThreshold},
		{Address: validator4_addr, TotalStake: config.Chain.EligibilityBalanceThreshold},
		{Address: validator5_addr, TotalStake: config.Chain.EligibilityBalanceThreshold},
	}

	// Simulated rewards statistics
	block := chain.Block{Pbft: models.Pbft{Number: 1, Author: validator3_addr}}
	bd := &chain.BlockData{Pbft: &block, TotalAmountDelegated: big.NewInt(0).Mul(DefaultMinimumDeposit, big.NewInt(8)), TotalSupply: big.NewInt(1), Validators: validators_list}
	r := MakeRewards(st, st.NewBatch(), config, bd)
	{
		rewardsStats := storage.RewardsStats{}
		rewardsStats.ValidatorsStats = []storage.ValidatorStatsWithAddress{
			{Address: validator1_addr, ValidatorStats: storage.ValidatorStats{DagBlocksCount: 8, VoteWeight: 1}},
			{Address: validator2_addr, ValidatorStats: storage.ValidatorStats{DagBlocksCount: 32, VoteWeight: 5}},
			{Address: validator4_addr, ValidatorStats: storage.ValidatorStats{VoteWeight: 1}},
		}
		rewardsStats.TotalDagCount = 40
		rewardsStats.TotalVotesWeight = 7
		rewardsStats.MaxVotesWeight = 8
		rewardsStats.BlockAuthor = block.Author

		// Expected block reward
		totalReward := rewardFromStake(config.Chain, r.totalStake)
		assert.Equal(t, totalReward, big.NewInt(202942668696093))
		rewardsParts := calculatePeriodRewardsParts(r.config.Chain, totalReward, false)
		rewards := r.rewardsFromStats(&rewardsStats)
		// We have 1 out of 2 bonus votes, so block author should get half of the bonus reward
		assert.Equal(t, big.NewInt(0).Div(rewardsParts.bonus, big.NewInt(2)), rewards.ValidatorRewards[block.Author])

		// data from node test
		expected_validator1_commission_reward := int64(31890990795100)
		expected_validator2_commission_reward := int64(139160687105891)
		expected_validator4_commission_reward := int64(11596723925491)
		assert.Equal(t, expected_validator1_commission_reward, rewards.ValidatorRewards[validator1_addr].Int64())
		assert.Equal(t, expected_validator2_commission_reward, rewards.ValidatorRewards[validator2_addr].Int64())
		assert.Equal(t, expected_validator4_commission_reward, rewards.ValidatorRewards[validator4_addr].Int64())
	}

	{
		rewardsStats := storage.RewardsStats{}
		rewardsStats.ValidatorsStats = []storage.ValidatorStatsWithAddress{
			{Address: validator1_addr, ValidatorStats: storage.ValidatorStats{DagBlocksCount: 15, VoteWeight: 3}},
			{Address: validator2_addr, ValidatorStats: storage.ValidatorStats{DagBlocksCount: 35, VoteWeight: 5}},
			{Address: validator4_addr, ValidatorStats: storage.ValidatorStats{VoteWeight: 2}},
		}
		rewardsStats.TotalDagCount = 50
		rewardsStats.TotalVotesWeight = 10
		rewardsStats.MaxVotesWeight = 13
		rewardsStats.BlockAuthor = block.Author

		// Expected block reward
		totalReward := rewardFromStake(config.Chain, r.totalStake)
		rewardsParts := calculatePeriodRewardsParts(r.config.Chain, totalReward, false)
		rewards := r.rewardsFromStats(&rewardsStats)
		// We have 1 out of 4 bonus votes, so block author should get 1/4 of the bonus reward
		assert.Equal(t, big.NewInt(0).Div(rewardsParts.bonus, big.NewInt(4)), rewards.ValidatorRewards[block.Author])
		assert.Equal(t, big.NewInt(5073566717402), rewards.ValidatorRewards[block.Author])
		// data from node test
		expected_validator1_commission_reward := int64(54794520547944)
		expected_validator2_commission_reward := int64(111618467782851)
		expected_validator4_commission_reward := int64(16235413495687)
		assert.Equal(t, expected_validator1_commission_reward, rewards.ValidatorRewards[validator1_addr].Int64())
		assert.Equal(t, expected_validator2_commission_reward, rewards.ValidatorRewards[validator2_addr].Int64())
		assert.Equal(t, expected_validator4_commission_reward, rewards.ValidatorRewards[validator4_addr].Int64())
	}

	{
		rewardsStats := storage.RewardsStats{}
		rewardsStats.ValidatorsStats = []storage.ValidatorStatsWithAddress{
			{Address: validator1_addr, ValidatorStats: storage.ValidatorStats{DagBlocksCount: 10, VoteWeight: 5}},
			{Address: validator2_addr, ValidatorStats: storage.ValidatorStats{DagBlocksCount: 10, VoteWeight: 5}},
			{Address: validator4_addr, ValidatorStats: storage.ValidatorStats{DagBlocksCount: 10, VoteWeight: 5}},
			{Address: validator5_addr, ValidatorStats: storage.ValidatorStats{DagBlocksCount: 10, VoteWeight: 5}},
		}
		rewardsStats.TotalDagCount = 40
		rewardsStats.TotalVotesWeight = 20
		rewardsStats.MaxVotesWeight = 24
		rewardsStats.BlockAuthor = block.Author

		// Expected block reward
		rewards := r.rewardsFromStats(&rewardsStats)
		// We have 1 out of 4 bonus votes, so block author should get 1/4 of the bonus reward
		// data from node test
		expected_block_author_reward := int64(8697542944118)
		expected_validator_reward := int64(45662100456620)
		assert.Equal(t, expected_block_author_reward, rewards.ValidatorRewards[block.Author].Int64())
		assert.Equal(t, expected_validator_reward, rewards.ValidatorRewards[validator1_addr].Int64())
		assert.Equal(t, expected_validator_reward, rewards.ValidatorRewards[validator2_addr].Int64())
		assert.Equal(t, expected_validator_reward, rewards.ValidatorRewards[validator4_addr].Int64())
		assert.Equal(t, expected_validator_reward, rewards.ValidatorRewards[validator4_addr].Int64())
	}

	// Block author is validator 1
	{
		block.Author = validator1_addr
		rewardsStats := storage.RewardsStats{}
		rewardsStats.ValidatorsStats = []storage.ValidatorStatsWithAddress{
			{Address: validator1_addr, ValidatorStats: storage.ValidatorStats{DagBlocksCount: 10, VoteWeight: 5}},
			{Address: validator2_addr, ValidatorStats: storage.ValidatorStats{DagBlocksCount: 10, VoteWeight: 5}},
			{Address: validator4_addr, ValidatorStats: storage.ValidatorStats{DagBlocksCount: 10, VoteWeight: 5}},
			{Address: validator5_addr, ValidatorStats: storage.ValidatorStats{DagBlocksCount: 10, VoteWeight: 5}},
		}
		rewardsStats.TotalDagCount = 40
		rewardsStats.TotalVotesWeight = 20
		rewardsStats.MaxVotesWeight = 24
		rewardsStats.BlockAuthor = block.Author

		// Expected block reward
		rewards := r.rewardsFromStats(&rewardsStats)
		// We have 1 out of 4 bonus votes, so block author should get 1/4 of the bonus reward
		// data from node test
		expected_block_author_reward := int64(8697542944118)
		expected_validator_reward := int64(45662100456620)
		assert.Equal(t, expected_validator_reward+expected_block_author_reward, rewards.ValidatorRewards[validator1_addr].Int64())
		assert.Equal(t, expected_validator_reward, rewards.ValidatorRewards[validator2_addr].Int64())
		assert.Equal(t, expected_validator_reward, rewards.ValidatorRewards[validator4_addr].Int64())
		assert.Equal(t, expected_validator_reward, rewards.ValidatorRewards[validator4_addr].Int64())
	}
}

func CalculateTotalStake(validators *Validators) *big.Int {
	totalStake := big.NewInt(0)
	for _, v := range validators.validators {
		totalStake.Add(totalStake, v.TotalStake)
	}
	return totalStake
}

func TestYieldsCalculation(t *testing.T) {
	config := makeTestConfig()
	config.Chain.BlocksPerYear = big.NewInt(10)

	total_minted := int64(15000000)
	validators_list := []chain.Validator{
		{Address: "0x1", TotalStake: big.NewInt(5000000)},
		{Address: "0x2", TotalStake: big.NewInt(10000000)},
		{Address: "0x3", TotalStake: big.NewInt(15000000)},
		{Address: "0x4", TotalStake: big.NewInt(20000000)},
		{Address: "0x5", TotalStake: big.NewInt(25000000)},
	}
	validators := MakeValidators(config, validators_list)
	totalStake := CalculateTotalStake(validators)
	rewards := make(map[string]*big.Int)

	total_check := uint64(0)
	for a, v := range validators.validators {
		rewards[a] = big.NewInt(0).Mul(big.NewInt(int64(total_minted)), v.TotalStake)
		rewards[a].Div(rewards[a], totalStake)
		total_check += rewards[a].Uint64()
	}
	assert.Equal(t, uint64(total_minted), total_check)

	validators_yield := GetValidatorsYield(rewards, validators)

	perc := float64(0)
	for _, y := range validators_yield {
		perc = GetYieldForInterval(y.Yield, big.NewInt(1), 1)
	}
	assert.Equal(t, float64(total_minted)/float64(totalStake.Int64()), perc)
}

func TestTotalYieldSaving(t *testing.T) {
	st := pebble.NewStorage("")
	config := makeTestConfig()
	config.TotalYieldSavingInterval = 10
	config.Chain.BlocksPerYear = big.NewInt(100)

	batch := st.NewBatch()
	// 10% yield per block
	multiplied_yield := GetMultipliedYield(big.NewInt(10), big.NewInt(1000))
	for i := 1; i <= 10; i++ {
		batch.AddSingleKey(storage.MultipliedYield{Yield: multiplied_yield}, storage.FormatIntToKey(uint64(i)))
	}
	batch.CommitBatch()

	totalStake := big.NewInt(0)

	block := chain.Block{Pbft: models.Pbft{Number: 10, Author: "0x4"}}
	bd := &chain.BlockData{Pbft: &block, TotalAmountDelegated: totalStake, TotalSupply: big.NewInt(1)}
	r := MakeRewards(st, st.NewBatch(), config, bd)
	b := st.NewBatch()
	assert.Equal(t, st.GetTotalYield(10), storage.Yield{})
	{
		count := 0
		storage.ProcessIntervalData[storage.MultipliedYield](r.storage, 1, func(key []byte, o storage.MultipliedYield) (stop bool) {
			count++
			return false
		})
		assert.Equal(t, 10, count)
	}
	r.processIntervalYield(b)
	b.CommitBatch()
	// check that this data was removed
	{
		count := 0
		storage.ProcessIntervalData[storage.MultipliedYield](r.storage, 1, func(key []byte, o storage.MultipliedYield) (stop bool) {
			count++
			return false
		})
		assert.Equal(t, 0, count)
	}

	// 10% yield per block will be equal to 100% for 10 blocks
	yield := st.GetTotalYield(10)
	assert.Equal(t, common.FormatFloat(10*GetYieldForInterval(multiplied_yield, config.Chain.BlocksPerYear, int64(config.TotalYieldSavingInterval))), yield.Yield)
}

func TestValidatorsYieldSaving(t *testing.T) {
	st := pebble.NewStorage("")
	config := makeTestConfig()
	config.TotalYieldSavingInterval = 10
	config.Chain.BlocksPerYear = big.NewInt(100)

	batch := st.NewBatch()
	// 10% yield per block
	multiplied_yield := GetMultipliedYield(big.NewInt(10), big.NewInt(1000))
	for i := 1; i <= 10; i++ {
		batch.AddSingleKey(storage.MultipliedYield{Yield: multiplied_yield}, storage.FormatIntToKey(uint64(i)))
	}
	batch.CommitBatch()
	totalStake := big.NewInt(0)

	block := chain.Block{Pbft: models.Pbft{Number: 10, Author: "0x4"}}
	bd := &chain.BlockData{Pbft: &block, TotalAmountDelegated: totalStake, TotalSupply: big.NewInt(1)}
	r := MakeRewards(st, st.NewBatch(), config, bd)
	b := st.NewBatch()
	assert.Equal(t, st.GetTotalYield(10), storage.Yield{})
	{
		count := 0
		storage.ProcessIntervalData[storage.MultipliedYield](r.storage, 1, func(key []byte, o storage.MultipliedYield) (stop bool) {
			count++
			return false
		})
		assert.Equal(t, 10, count)
	}
	r.processIntervalYield(b)
	b.CommitBatch()
	// check that this data was removed
	{
		count := 0
		storage.ProcessIntervalData[storage.MultipliedYield](r.storage, 1, func(key []byte, o storage.MultipliedYield) (stop bool) {
			count++
			return false
		})
		assert.Equal(t, 0, count)
	}

	// 10% yield per block will be equal to 100% for 10 blocks
	yield := st.GetTotalYield(10)
	assert.Equal(t, common.FormatFloat(10*GetYieldForInterval(multiplied_yield, config.Chain.BlocksPerYear, int64(config.TotalYieldSavingInterval))), yield.Yield)
}
