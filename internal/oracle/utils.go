package oracle

import (
	"encoding/csv"
	"math/big"
	"os"
	"strconv"

	apy_oracle "github.com/Taraxa-project/taraxa-indexer/abi/oracle"
	"github.com/ethereum/go-ethereum/common"
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

func (v *YieldedValidator) ToNodeData(currentBlock uint64) NodeData {
	rating64, from, to := v.calculateRating(currentBlock)
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

func (validator *YieldedValidator) calculateRating(currentBlock uint64) (int64, uint64, uint64) {
	commission_float := float64(*validator.Commisson)
	yield_float, err := strconv.ParseFloat(validator.Yield, 64)
	if err != nil {
		log.Fatalf("Failed to parse yield: %v", err)
	}
	commission_percentage := commission_float / float64(10000)
	adjusted_apy := (1 - commission_percentage) * yield_float * 100
	continuity := float64(validator.PbftCount) / float64(currentBlock-validator.RegistrationBlock)

	// Check for 100% commission and zero yield
	score := float64(0)
	if commission_float == 10000 || yield_float == 0 {
		score = 0
	} else {
		//w1 * (APY) - (Commission * w2) + w3 * Continuity + w4 * stake
		score = float64(0.4)*adjusted_apy - float64(0.1)*commission_float + float64(0.5)*continuity
	}
	log.WithFields(log.Fields{"validator": validator.Account.String(), "currentBlock": currentBlock, "score": score, "continuity": continuity, "apy": adjusted_apy, "commission": commission_float}).Info("Validator score")
	file, err := os.OpenFile("validator_scores.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open csv file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	record := []string{
		validator.Account.String(),
		strconv.FormatFloat(score, 'f', 6, 64),
		strconv.FormatFloat(adjusted_apy, 'f', 6, 64),
		strconv.FormatFloat(commission_float, 'f', 6, 64),
		strconv.FormatFloat(continuity, 'f', 6, 64),
		strconv.FormatUint(currentBlock, 10),
	}
	err = writer.Write(record)
	if err != nil {
		log.Fatalf("Failed to write to csv file: %v", err)
	}

	return int64(score * 1000), validator.RegistrationBlock, currentBlock
}
