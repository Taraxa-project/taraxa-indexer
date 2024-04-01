package rewards

import (
	"math/big"

	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage/pebble"
	"github.com/ethereum/go-ethereum/rlp"
	log "github.com/sirupsen/logrus"
)

type IntervalRewards struct {
	ValidatorRewards map[string]*big.Int
	TotalReward      *big.Int
	BlockFee         *big.Int
}

func makeIntervalRewards() (r IntervalRewards) {
	r.ValidatorRewards = make(map[string]*big.Int)
	r.TotalReward = big.NewInt(0)
	r.BlockFee = big.NewInt(0)
	return
}

func (r *Rewards) GetIntervalRewards(fullBlockReward *big.Int, periodStats *storage.RewardsStats, distributionFrequency uint32) IntervalRewards {
	// Calculate rewards for the last interval

	intervalRewards := r.rewardsFromStats(fullBlockReward, periodStats)
	// Get stats for the previous intervals and accumulate rewards
	fromKey := storage.FormatIntToKey(r.blockNum - uint64(distributionFrequency))
	r.storage.ForEachFromKey([]byte(pebble.GetPrefix(storage.RewardsStats{})), []byte(fromKey), func(key, res []byte) (stop bool) {
		rs := new(storage.RewardsStats)
		err := rlp.DecodeBytes(res, rs)
		if err != nil {
			log.WithError(err).Fatal("Error decoding data from db")
		}
		pr := r.rewardsFromStats(fullBlockReward, rs)
		for validator, reward := range pr.ValidatorRewards {
			if intervalRewards.ValidatorRewards[validator] == nil {
				intervalRewards.ValidatorRewards[validator] = big.NewInt(0)
			}
			intervalRewards.ValidatorRewards[validator].Add(intervalRewards.ValidatorRewards[validator], reward)
		}
		intervalRewards.TotalReward.Add(intervalRewards.TotalReward, pr.TotalReward)
		intervalRewards.BlockFee.Add(intervalRewards.BlockFee, pr.BlockFee)
		r.batch.Remove(key)
		return false
	})

	return intervalRewards
}
