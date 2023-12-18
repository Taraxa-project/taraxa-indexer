package oracle

import (
	"context"
	"math/big"
	"strconv"

	apy_oracle "github.com/Taraxa-project/taraxa-indexer/abi/oracle"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

type NodeData = apy_oracle.IApyOracleNodeData

type YieldedValidator struct {
	Account           common.Address
	Rank              uint16
	Rating            uint64
	Yield             string
	Commisson         *uint64
	RegistrationBlock uint64
	PbftCount         uint64
}

type RawValidator struct {
	Address common.Address
	Yield   string
}

func (r *RawValidator) ToYieldedValidator() YieldedValidator {
	return YieldedValidator{
		Account: r.Address,
		Yield:   r.Yield,
	}
}

func (y *YieldedValidator) ToRawValidator() RawValidator {
	return RawValidator{
		Address: y.Account,
		Yield:   y.Yield,
	}
}

func (v *YieldedValidator) ToNodeData(cli *ethclient.Client) NodeData {
	rating64, from, to := v.calculateRating(cli)
	yield, err := strconv.ParseFloat(v.Yield, 64)
	if err != nil {
		log.Fatalf("Failed to parse yield: %v", err)
	}

	return NodeData{
		Rating:    big.NewInt(rating64),
		Account:   v.Account,
		Rank:      v.Rank,
		Apy:       uint16(yield * 1000),
		FromBlock: from,
		ToBlock:   to,
	}
}

// will not be used in the first primitive version
func (validator *YieldedValidator) calculateRating(client *ethclient.Client) (int64, uint64, uint64) {
	currentBlock, err := client.BlockByNumber(context.Background(), nil)

	if err != nil {
		log.Fatalf("Failed to get current block: %v", err)
	}

	blocksSinceRegistration := currentBlock.NumberU64() - validator.RegistrationBlock
	commission_float := float64(*validator.Commisson)
	yield_float, err := strconv.ParseFloat(validator.Yield, 64)
	if err != nil {
		log.Fatalf("Failed to parse yield: %v", err)
	}
	commission_percentage := commission_float / float64(100000)
	adjusted_apy := (1 - commission_percentage) * yield_float * 100
	continuity := float64(blocksSinceRegistration) / float64(currentBlock.NumberU64()-validator.RegistrationBlock)

	//w1 * (APY) - (Commission * w2) + w3 * Continuity + w4 * stake
	score := float64(0.4)*adjusted_apy - float64(0.1)*commission_float + float64(0.5)*continuity
	return int64(score * 1000), validator.RegistrationBlock, currentBlock.NumberU64()
}
