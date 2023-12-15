package common

import (
	"encoding/json"
	"math/big"
	"sort"
)

type MagnoliaHfConfig struct {
	BlockNum uint64 `json:"block_num"`
}

func (hf *MagnoliaHfConfig) UnmarshalJSON(data []byte) error {
	var res map[string]string

	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	hf.BlockNum = ParseUInt(res["block_num"])

	return nil
}

type HardforksConfig struct {
	RewardsDistributionFrequency map[uint64]uint32 `json:"rewards_distribution_frequency"`
	MagnoliaHf                   MagnoliaHfConfig  `json:"magnolia_hf"`
}

func (c *HardforksConfig) GetDistributionFrequency(period uint64) uint32 {
	keys := make([]uint64, 0)
	for k := range c.RewardsDistributionFrequency {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	lastBigger := uint64(0)
	found := false
	for _, ki := range keys {
		k := uint64(ki)
		if period >= k {
			found = true
			lastBigger = k
		} else {
			break
		}
	}
	if !found {
		return 1
	}
	return c.RewardsDistributionFrequency[lastBigger]
}

type ChainConfig struct {
	CommitteeSize               *big.Int
	BlocksPerYear               *big.Int
	YieldPercentage             *big.Int
	DagProposersReward          *big.Int
	MaxBlockAuthorReward        *big.Int
	EligibilityBalanceThreshold *big.Int
	Hardforks                   HardforksConfig
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
