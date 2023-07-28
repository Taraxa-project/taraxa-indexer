package common

import (
	"math/big"
)

type ChainConfig struct {
	CommitteeSize               *big.Int
	BlocksPerYear               *big.Int
	YieldPercentage             *big.Int
	DagProposersReward          *big.Int
	MaxBlockAuthorReward        *big.Int
	EligibilityBalanceThreshold *big.Int
}

func DefaultChainConfig() *ChainConfig {
	return &ChainConfig{
		CommitteeSize:               big.NewInt(1000),
		BlocksPerYear:               big.NewInt(365 * 24 * 60 * 15),
		YieldPercentage:             big.NewInt(20),
		DagProposersReward:          big.NewInt(50),
		MaxBlockAuthorReward:        big.NewInt(10),
		EligibilityBalanceThreshold: ParseStringToBigInt("0x69E10DE76676D0800000"),
	}
}

type Config struct {
	Chain                         *ChainConfig
	TotalYieldSavingInterval      uint64
	ValidatorsYieldSavingInterval uint64
}

func (c *Config) IsEligible(stake *big.Int) bool {
	if c.Chain != nil && c.Chain.EligibilityBalanceThreshold != nil && stake.Cmp(c.Chain.EligibilityBalanceThreshold) >= 0 {
		return true
	}
	return false
}

func DefaultConfig() *Config {
	return &Config{
		Chain:                         DefaultChainConfig(),
		TotalYieldSavingInterval:      1000,
		ValidatorsYieldSavingInterval: 1000,
	}
}
