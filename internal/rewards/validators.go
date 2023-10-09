package rewards

import (
	"strings"

	"github.com/Taraxa-project/taraxa-go-client/taraxa_client/dpos_contract_client/dpos_interface"
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
)

type Validators struct {
	config     *common.Config
	validators map[string]dpos_interface.DposInterfaceValidatorBasicInfo
}

func MakeValidators(config *common.Config, validators []dpos_interface.DposInterfaceValidatorData) *Validators {
	v := Validators{config, make(map[string]dpos_interface.DposInterfaceValidatorBasicInfo)}
	for _, val := range validators {
		v.validators[strings.ToLower(val.Account.Hex())] = val.Info
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
