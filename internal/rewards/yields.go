package rewards

import (
	"math/big"

	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	log "github.com/sirupsen/logrus"
)

var multiplier = big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)
var percentage_multiplier = big.NewInt(10000)
var YieldFractionDecimalPrecision = big.NewInt(1e+6)

func (r *Rewards) calculateCurrentYield(current_total_tara_supply *big.Int) *big.Int {
	// Current yield = (max supply - current total supply) / current total supply
	current_yield := big.NewInt(0).Sub(r.config.Hardforks.AspenHf.MaxSupply, current_total_tara_supply)
	current_yield.Mul(current_yield, YieldFractionDecimalPrecision)
	current_yield.Div(current_yield, current_total_tara_supply)

	return current_yield
}

func GetMultipliedYield(reward, stake *big.Int) *big.Int {
	r := big.NewInt(0)
	r.Mul(reward, multiplier)
	r.Div(r, stake)

	return r
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

func (r *Rewards) processIntervalYield(intervalStart uint64, batch storage.Batch) {
	sum := big.NewInt(0)
	storage.ProcessIntervalData(r.storage, intervalStart, func(key []byte, o storage.MultipliedYield) (stop bool) {
		sum.Add(sum, o.Yield)
		batch.Remove([]byte(key))
		return false
	})

	yield := GetYieldForInterval(sum, r.config.BlocksPerYear, int64(r.blockNum-intervalStart))
	log.WithFields(log.Fields{"total_yield": yield}).Info("processIntervalYield")
	batch.AddSingleKey(&storage.Yield{Yield: common.FormatFloat(yield)}, storage.FormatIntToKey(r.blockNum))
}

func (r *Rewards) processValidatorsIntervalYield(intervalStart uint64, batch storage.Batch) {

	sum_by_validator := make(map[string]*big.Int)

	storage.ProcessIntervalData(r.storage, intervalStart, func(key []byte, o storage.ValidatorsYield) (stop bool) {
		for _, y := range o.Yields {
			if sum_by_validator[y.Validator] == nil {
				sum_by_validator[y.Validator] = big.NewInt(0)
			}
			sum_by_validator[y.Validator].Add(sum_by_validator[y.Validator], y.Yield)
		}
		batch.Remove(key)
		return false
	})

	log.WithFields(log.Fields{"validators": len(sum_by_validator)}).Info("processValidatorsIntervalYield")

	for val, sum := range sum_by_validator {
		yield := GetYieldForInterval(sum, r.config.BlocksPerYear, int64(r.blockNum-intervalStart))
		log.WithFields(log.Fields{"validator": val, "yield": yield}).Info("processValidatorsIntervalYield")
		batch.Add(&storage.Yield{Yield: common.FormatFloat(yield)}, val, r.blockNum)
	}
}
