package rewards

import (
	"math/big"
	"strings"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/models"
)

type validatorStats struct {
	// count of rewardable(with 1 or more unique transactions) DAG blocks produced by this validator
	dagBlocksCount int64

	// Validator cert voted block weight
	voteWeight int64
}

type totalRewards struct {
	dags  *big.Int
	votes *big.Int
	bonus *big.Int
}

type stats struct {
	TotalVotesWeight int64
	MaxVotesWeight   int64
	TotalDagCount    int64
	ValidatorStats   map[string]validatorStats
}

func getPeriodTransactionsMap(trxs []models.Transaction) map[string]bool {
	period_transactions := make(map[string]bool, 0)
	for _, t := range trxs {
		period_transactions[t.Hash] = true
	}

	return period_transactions
}

func makeStats(dags []chain.DagBlock, votes chain.VotesResponse, trxs []models.Transaction, committee_size int64) (s *stats) {
	s = new(stats)
	s.ValidatorStats = make(map[string]validatorStats)
	s.MaxVotesWeight = Min(votes.PeriodTotalVotesCount, committee_size)

	for _, v := range votes.Votes {
		voter := strings.ToLower(v.Voter)
		entry := s.ValidatorStats[voter]
		entry.voteWeight = int64(chain.ParseInt(v.Weight))
		s.TotalVotesWeight += entry.voteWeight

		s.ValidatorStats[voter] = entry
	}

	period_transactions := getPeriodTransactionsMap(trxs)
	is_tx_seen := make(map[string]bool, 0)
	total_dag_count := int64(0)
	for _, d := range dags {
		total_dag_count += 1
		has_unique_transactions := false
		for _, th := range d.Transactions {
			if is_tx_seen[th] {
				continue
			}
			if period_transactions[th] {
				has_unique_transactions = true
				period_transactions[th] = false
			}
			is_tx_seen[th] = true
		}
		if has_unique_transactions {
			sender := strings.ToLower(d.Sender)
			entry := s.ValidatorStats[sender]
			entry.dagBlocksCount += 1
			s.ValidatorStats[sender] = entry
			s.TotalDagCount += 1
		}
	}
	return
}

func calculateTotalRewards(config *common.ChainConfig, totalStake *big.Int) (tr totalRewards) {
	// calculate total rewards
	totalRewards := big.NewInt(0).Mul(totalStake, config.YieldPercentage)
	totalRewards.Div(totalRewards, big.NewInt(0).Mul(big.NewInt(100), config.BlocksPerYear))
	// calculate dags rewards
	tr.dags = big.NewInt(0)
	tr.dags.Mul(totalRewards, config.DagProposersReward).Div(tr.dags, big.NewInt(100))

	// calculate bonus reward
	tr.bonus = big.NewInt(0)
	tr.bonus.Div(big.NewInt(0).Mul(totalRewards, config.MaxBlockAuthorReward), big.NewInt(100))

	// calculate votes rewards
	tr.votes = big.NewInt(0)
	tr.votes.Sub(totalRewards, tr.dags)
	tr.votes.Sub(tr.votes, tr.bonus)

	return
}

func Max(x, y int64) int64 {
	if x < y {
		return y
	}
	return x
}

func Min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}
