package rewards

import (
	"strings"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
)

type Validators struct {
	config     *common.Config
	validators map[string]chain.Validator
}

func MakeValidators(config *common.Config, validators []chain.Validator) *Validators {
	v := Validators{config, make(map[string]chain.Validator)}
	for _, val := range validators {
		v.validators[strings.ToLower(val.Address)] = val
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
