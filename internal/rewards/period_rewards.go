package rewards

import (
	"math/big"

	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage/pebble"
	"github.com/ethereum/go-ethereum/rlp"
	log "github.com/sirupsen/logrus"
)

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
	r.BlockFee = big.NewInt(0)
	return
}

func (r *Rewards) GetIntervalRewards(periodRewards PeriodRewards, distributionFrequency uint32) PeriodRewards {
	// we have all needed data for 1 block interval
	if distributionFrequency == 1 {
		return periodRewards
	}

	fromKey := storage.FormatIntToKey(r.blockNum - uint64(distributionFrequency))
	r.storage.ForEachFromKey([]byte(pebble.GetPrefix(storage.PeriodRewards{})), []byte(fromKey), func(key, res []byte) (stop bool) {
		var pr storage.PeriodRewards
		err := rlp.DecodeBytes(res, &pr)
		if err != nil {
			log.WithError(err).Fatal("Error decoding data from db")
		}
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

	return periodRewards
}
