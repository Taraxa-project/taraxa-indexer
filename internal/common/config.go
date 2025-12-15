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

type AspenHfConfig struct {
	BlockNumPartOne  uint64   `json:"block_num_part_one"`
	BlockNumPartTwo  uint64   `json:"block_num_part_two"`
	MaxSupply        *big.Int `json:"max_supply"`
	GeneratedRewards *big.Int `json:"generated_rewards"`
}

func (hf *AspenHfConfig) UnmarshalJSON(data []byte) error {
	var res map[string]string

	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	hf.BlockNumPartOne = ParseUInt(res["block_num_part_one"])
	hf.BlockNumPartTwo = ParseUInt(res["block_num_part_two"])
	hf.MaxSupply = ParseStringToBigInt(res["max_supply"])
	hf.GeneratedRewards = ParseStringToBigInt(res["generated_rewards"])

	return nil
}

type CactiHfConfig struct {
	BlockNum       uint64 `json:"block_num"`
	DefaultLambda  uint64 `json:"lambda_default"`
	ConsensusDelay uint64 `json:"consensus_delay"`
}

func (hf *CactiHfConfig) UnmarshalJSON(data []byte) error {
	var res map[string]string

	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	hf.BlockNum = ParseUInt(res["block_num"])
	hf.DefaultLambda = ParseUInt(res["lambda_default"])
	hf.ConsensusDelay = ParseUInt(res["consensus_delay"])

	return nil
}

type HardforksConfig struct {
	RewardsDistributionFrequency map[uint64]uint32 `json:"rewards_distribution_frequency"`
	MagnoliaHf                   MagnoliaHfConfig  `json:"magnolia_hf"`
	AspenHf                      AspenHfConfig     `json:"aspen_hf"`
	CactiHf                      CactiHfConfig     `json:"cacti_hf"`
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

func (c *HardforksConfig) IsAspenHfOne(blockNum uint64) bool {
	return blockNum >= c.AspenHf.BlockNumPartOne
}

func (c *HardforksConfig) IsAspenHfTwo(blockNum uint64) bool {
	return blockNum >= c.AspenHf.BlockNumPartTwo
}

func (c *HardforksConfig) IsCactiHf(blockNum uint64) bool {
	return blockNum >= c.CactiHf.BlockNum
}

type ChainConfig struct {
	CommitteeSize               *big.Int
	BlocksPerYear               *big.Int
	YieldPercentage             *big.Int
	DagProposersReward          *big.Int
	MaxBlockAuthorReward        *big.Int
	EligibilityBalanceThreshold *big.Int
	Hardforks                   HardforksConfig
	DagGenesisBlock             DagBlock
	LambdaMs                    uint64
}

func (cc *ChainConfig) IsEligible(stake *big.Int) bool {
	if cc.EligibilityBalanceThreshold != nil && stake.Cmp(cc.EligibilityBalanceThreshold) >= 0 {
		return true
	}
	return false
}

func (cc *ChainConfig) InitLambda(period uint64, dbLambdaMs *uint64) {
	if cc.Hardforks.CactiHf.BlockNum <= period {
		if dbLambdaMs != nil {
			cc.LambdaMs = *dbLambdaMs
		} else {
			cc.LambdaMs = cc.Hardforks.CactiHf.DefaultLambda
		}
	}
}

func (cc *ChainConfig) AdjustLambda(period uint64, lambda *uint64) {
	if cc.Hardforks.CactiHf.BlockNum <= period {
		if lambda != nil {
			cc.LambdaMs = *lambda
		}
	}
}

func (cc *ChainConfig) GetLambda(round uint64) uint64 {
	if round == 1 {
		return cc.LambdaMs
	}
	return cc.Hardforks.CactiHf.DefaultLambda
}

func DefaultChainConfig() *ChainConfig {
	return &ChainConfig{
		CommitteeSize:               big.NewInt(1000),
		BlocksPerYear:               big.NewInt(365 * 24 * 60 * 15),
		YieldPercentage:             big.NewInt(20),
		DagProposersReward:          big.NewInt(50),
		MaxBlockAuthorReward:        big.NewInt(10),
		EligibilityBalanceThreshold: ParseStringToBigInt("0x69E10DE76676D0800000"),
		LambdaMs:                    1500,
	}
}

type Config struct {
	Chain              *ChainConfig
	SyncQueueLimit     uint64
	ChainStatsInterval int
	// Auth credentials for protected endpoints
	AuthUsername string
	AuthPassword string
}

func DefaultConfig() *Config {
	return &Config{
		Chain:          DefaultChainConfig(),
		SyncQueueLimit: 10,
	}
}
