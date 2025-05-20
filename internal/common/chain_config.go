package common

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
	Hardforks       HardforksConfig   `json:"hardforks"`
}

func (g *GenesisObject) ToChainConfig() (c *ChainConfig) {
	c = new(ChainConfig)
	c.CommitteeSize = ParseStringToBigInt(g.Pbft.CommitteeSize)
	c.BlocksPerYear = ParseStringToBigInt(g.Dpos.BlocksPerYear)
	c.YieldPercentage = ParseStringToBigInt(g.Dpos.YieldPercentage)
	c.DagProposersReward = ParseStringToBigInt(g.Dpos.DagProposersReward)
	c.MaxBlockAuthorReward = ParseStringToBigInt(g.Dpos.MaxBlockAuthorReward)
	c.EligibilityBalanceThreshold = ParseStringToBigInt(g.Dpos.EligibilityBalanceThreshold)
	c.Hardforks = g.Hardforks
	return
}
