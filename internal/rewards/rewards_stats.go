package rewards

import (
	"math/big"
	"strings"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
)

type RewardsStats struct {
	storage.RewardsStats
	ValidatorsStats map[string]storage.ValidatorStats
}

func (s *RewardsStats) ToStorage() *storage.RewardsStats {
	rs := storage.RewardsStats{TotalRewardsStats: s.TotalRewardsStats}
	rs.ValidatorsStats = make([]storage.ValidatorStatsWithAddress, 0, len(s.ValidatorsStats))
	for addr, stats := range s.ValidatorsStats {
		rs.ValidatorsStats = append(rs.ValidatorsStats, storage.ValidatorStatsWithAddress{ValidatorStats: stats, Address: addr})
	}
	return &rs
}

func (s *RewardsStats) processDags(dags []chain.DagBlock, trxs []chain.Transaction) {
	transaction_fees := getPeriodTransactionsFees(trxs)
	total_dag_count := int64(0)
	for _, d := range dags {
		total_dag_count += 1
		feeReward := dagFeeReward(transaction_fees, d)
		has_unique_trx := feeReward.Cmp(big.NewInt(0)) > 0
		if has_unique_trx {
			sender := strings.ToLower(d.Sender)
			entry := s.ValidatorsStats[sender]
			entry.DagBlocksCount += 1
			if entry.FeeReward == nil {
				entry.FeeReward = big.NewInt(0)
			}
			entry.FeeReward.Add(entry.FeeReward, feeReward)
			s.ValidatorsStats[sender] = entry
			s.TotalDagCount += 1
		}
	}
}

func (s *RewardsStats) processDagsAspen(dags []chain.DagBlock, trxs []chain.Transaction) {
	transaction_fees := getPeriodTransactionsFees(trxs)
	min_difficulty := ^uint16(0)
	for _, d := range dags {
		if d.Vdf.Difficulty < min_difficulty {
			min_difficulty = d.Vdf.Difficulty
		}
	}

	for _, d := range dags {
		author := d.Sender
		entry := s.ValidatorsStats[author]
		if d.Vdf.Difficulty == min_difficulty {
			entry.DagBlocksCount += 1
			s.TotalDagCount += 1
		}
		if entry.FeeReward == nil {
			entry.FeeReward = big.NewInt(0)
		}
		entry.FeeReward.Add(entry.FeeReward, dagFeeReward(transaction_fees, d))
		s.ValidatorsStats[author] = entry
	}
}

func dagFeeReward(fees map[string]*big.Int, d chain.DagBlock) *big.Int {
	feeReward := big.NewInt(0)
	for _, th := range d.Transactions {
		// if we don't have fee for this transaction, it means that it was processed before
		if fees[th] != nil {
			feeReward.Add(feeReward, fees[th])
			delete(fees, th)
		}
	}

	return feeReward
}

func getPeriodTransactionsFees(trxs []chain.Transaction) map[string]*big.Int {
	period_transactions := make(map[string]*big.Int, 0)
	for _, t := range trxs {
		period_transactions[t.Hash] = t.GetFee()
	}

	return period_transactions
}

func makeRewardsStats(is_aspen_dag_rewards bool, dags []chain.DagBlock, votes chain.VotesResponse, trxs []chain.Transaction, committee_size uint64, blockAuthor string) (s *RewardsStats) {
	s = new(RewardsStats)
	s.ValidatorsStats = make(map[string]storage.ValidatorStats)
	s.MaxVotesWeight = common.Min(votes.PeriodTotalVotesCount, committee_size)
	s.BlockAuthor = blockAuthor

	for _, v := range votes.Votes {
		voter := strings.ToLower(v.Voter)
		entry := s.ValidatorsStats[voter]
		entry.VoteWeight = common.ParseUInt(v.Weight)
		s.TotalVotesWeight += entry.VoteWeight

		s.ValidatorsStats[voter] = entry
	}

	if is_aspen_dag_rewards {
		s.processDagsAspen(dags, trxs)
	} else {
		s.processDags(dags, trxs)
	}
	return
}

type totalPeriodRewards struct {
	dags  *big.Int
	votes *big.Int
	bonus *big.Int
}

func (tr totalPeriodRewards) MarshalJSON() ([]byte, error) {
	return []byte(`{"dags":` + tr.dags.String() + `,"votes":` + tr.votes.String() + `,"bonus":` + tr.bonus.String() + `}`), nil
}

func ZeroTotalRewards() (tr totalPeriodRewards) {
	tr.dags = big.NewInt(0)
	tr.votes = big.NewInt(0)
	tr.bonus = big.NewInt(0)
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
