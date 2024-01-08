package rewards

import (
	"math/big"
	"strings"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
)

type validatorStats struct {
	// count of rewardable(with 1 or more unique transactions) DAG blocks produced by this validator
	DagBlocksCount int64

	// Validator cert voted block weight
	VoteWeight int64
}

type totalPeriodRewards struct {
	dags  *big.Int
	votes *big.Int
	bonus *big.Int
}

func ZeroTotalRewards() (tr totalPeriodRewards) {
	tr.dags = big.NewInt(0)
	tr.votes = big.NewInt(0)
	tr.bonus = big.NewInt(0)
	return
}

type stats struct {
	TotalVotesWeight int64
	MaxVotesWeight   int64
	TotalDagCount    int64
	ValidatorStats   map[string]validatorStats
}

func getPeriodTransactionsMap(trxs []chain.Transaction) map[string]bool {
	period_transactions := make(map[string]bool, 0)
	for _, t := range trxs {
		period_transactions[t.Hash] = true
	}

	return period_transactions
}

func makeStats(dags []chain.DagBlock, votes chain.VotesResponse, trxs []chain.Transaction, committee_size int64) (s *stats) {
	s = new(stats)
	s.ValidatorStats = make(map[string]validatorStats)
	s.MaxVotesWeight = Min(votes.PeriodTotalVotesCount, committee_size)

	for _, v := range votes.Votes {
		voter := strings.ToLower(v.Voter)
		entry := s.ValidatorStats[voter]
		entry.VoteWeight = int64(common.ParseInt(v.Weight))
		s.TotalVotesWeight += entry.VoteWeight

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
			entry.DagBlocksCount += 1
			s.ValidatorStats[sender] = entry
			s.TotalDagCount += 1
		}
	}
	return
}

func calculatePeriodRewardsParts(config *common.ChainConfig, totalRewards *big.Int, noVotes bool) (tr totalPeriodRewards) {
	tr = ZeroTotalRewards()

	// Should only happen for block 1, so we are distributing all rewards to dag blocks producers
	if noVotes {
		tr.dags = totalRewards
		return
	}

	// calculate dags rewards
	tr.dags.Mul(totalRewards, config.DagProposersReward).Div(tr.dags, big.NewInt(100))

	// calculate bonus reward
	tr.bonus.Div(big.NewInt(0).Mul(totalRewards, config.MaxBlockAuthorReward), big.NewInt(100))

	// calculate votes rewards
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
