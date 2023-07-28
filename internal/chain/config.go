package chain

import (
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
)

type initialValidator struct {
	Address     string            `json:"address"`
	Delegations map[string]string `json:"delegations"`
}

type DposConfig struct {
	BlocksPerYear               string             `json:"blocks_per_year"`
	DagProposersReward          string             `json:"dag_proposers_reward"`
	MaxBlockAuthorReward        string             `json:"max_block_author_reward"`
	EligibilityBalanceThreshold string             `json:"eligibility_balance_threshold"`
	YieldPercentage             string             `json:"yield_percentage"`
	InitialValidators           []initialValidator `json:"initial_validators"`
}

type PbftConfig struct {
	CommitteeSize string `json:"committee_size"`
	LambdaMs      string `json:"lambda_ms"`
}

type GenesisObject struct {
	DagGenesisBlock DagBlock          `json:"dag_genesis_block"`
	InitialBalances map[string]string `json:"initial_balances"`
	Pbft            PbftConfig        `json:"pbft"`
	Dpos            DposConfig        `json:"dpos"`
}

func (g *GenesisObject) ToChainConfig() (c *common.ChainConfig) {
	c = new(common.ChainConfig)
	c.CommitteeSize = common.ParseStringToBigInt(g.Pbft.CommitteeSize)
	c.BlocksPerYear = common.ParseStringToBigInt(g.Dpos.BlocksPerYear)
	c.YieldPercentage = common.ParseStringToBigInt(g.Dpos.YieldPercentage)
	c.DagProposersReward = common.ParseStringToBigInt(g.Dpos.DagProposersReward)
	c.MaxBlockAuthorReward = common.ParseStringToBigInt(g.Dpos.MaxBlockAuthorReward)
	c.EligibilityBalanceThreshold = common.ParseStringToBigInt(g.Dpos.EligibilityBalanceThreshold)
	return
}
