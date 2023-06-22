package chain

import (
	"math/big"

	"github.com/Taraxa-project/taraxa-indexer/internal/common"
)

type DposConfig struct {
	BlocksPerYear        string `json:"blocks_per_year"`
	DagProposersReward   string `json:"dag_proposers_reward"`
	MaxBlockAuthorReward string `json:"max_block_author_reward"`
	YieldPercentage      string `json:"yield_percentage"`
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
	c.CommitteeSize = big.NewInt(ParseInt(g.Pbft.CommitteeSize))
	c.BlocksPerYear = big.NewInt(ParseInt(g.Dpos.BlocksPerYear))
	c.YieldPercentage = big.NewInt(ParseInt(g.Dpos.YieldPercentage))
	c.DagProposersReward = big.NewInt(ParseInt(g.Dpos.DagProposersReward))
	c.MaxBlockAuthorReward = big.NewInt(ParseInt(g.Dpos.MaxBlockAuthorReward))
	return
}
