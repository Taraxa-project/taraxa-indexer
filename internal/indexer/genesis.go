package indexer

import (
	"math/big"

	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/models"
	log "github.com/sirupsen/logrus"
)

type Genesis struct {
	storage storage.Storage
	genesis common.GenesisObject
	bc      blockContext
	hash    string
}

func MakeGenesis(s storage.Storage, c common.Client, gen_obj common.GenesisObject, genesisHash storage.GenesisHash, dayStats *storage.DayStatsWithTimestamp) *Genesis {
	var genesis Genesis
	genesis.storage = s
	genesis.genesis = gen_obj
	genesis.hash = string(genesisHash)
	genesis.bc = *MakeBlockContext(s, c, &common.Config{Chain: gen_obj.ToChainConfig()}, dayStats)

	return &genesis
}

func (g *Genesis) makeInitBalanceTrx(addr, value string) (trx models.Transaction) {
	trx.Hash = "GENESIS_" + addr
	trx.From = "GENESIS"
	trx.To = addr
	trx.Value = value
	trx.BlockNumber = 0
	trx.Timestamp = g.genesis.DagGenesisBlock.Timestamp
	trx.Status = true
	return
}

func (g *Genesis) process() {
	genesisSupply := big.NewInt(0)
	for addr, value := range g.genesis.InitialBalances {
		trx := g.makeInitBalanceTrx(addr, value)
		g.bc.SaveTransaction(trx, false)
		value := common.ParseStringToBigInt(trx.Value)
		genesisSupply.Add(genesisSupply, value)
		g.bc.addressStats.AddToBalance(g.bc.Storage, trx.To, value)
		// g.accounts.AddToBalance(trx.To, value)
	}
	for _, validator := range g.genesis.Dpos.InitialValidators {
		for addr, value := range validator.Delegations {
			delegation := common.ParseStringToBigInt(value)
			g.bc.addressStats.AddToBalance(g.bc.Storage, addr, big.NewInt(0).Neg(delegation))
			g.bc.addressStats.AddToBalance(g.bc.Storage, common.DposContractAddress, delegation)
		}
	}
	log.WithField("count", len(g.genesis.InitialBalances)).Info("Genesis: Init balance transactions parsed")

	// Genesis transactions isn't real transactions, so don't count it here
	g.bc.finalized.TrxCount = 0
	g.bc.Batch.SetGenesisHash(storage.GenesisHash(g.hash))
	g.bc.Batch.SetTotalSupply(genesisSupply)
	g.bc.commit()
}
