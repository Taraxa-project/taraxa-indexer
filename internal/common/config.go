package common

import (
	"math/big"
)

type ChainConfig struct {
	CommitteeSize        *big.Int
	BlocksPerYear        *big.Int
	YieldPercentage      *big.Int
	DagProposersReward   *big.Int
	MaxBlockAuthorReward *big.Int
}

func DefaultChainConfig() *ChainConfig {
	return &ChainConfig{
		CommitteeSize:        big.NewInt(1000),
		BlocksPerYear:        big.NewInt(365 * 24 * 60 * 15),
		YieldPercentage:      big.NewInt(20),
		DagProposersReward:   big.NewInt(50),
		MaxBlockAuthorReward: big.NewInt(10),
	}
}

type Config struct {
	Chain                         *ChainConfig
	TotalYieldSavingInterval      uint64
	ValidatorsYieldSavingInterval uint64
}

func DefaultConfig() *Config {
	return &Config{
		Chain:                         DefaultChainConfig(),
		TotalYieldSavingInterval:      1000,
		ValidatorsYieldSavingInterval: 1000,
	}
}
