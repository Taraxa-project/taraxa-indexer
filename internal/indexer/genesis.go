package indexer

import (
	"math/big"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	log "github.com/sirupsen/logrus"
)

type Genesis struct {
	storage  storage.Storage
	genesis  chain.GenesisObject
	bc       blockContext
	hash     string
	accounts *storage.AccountsMap
}

func MakeGenesis(s storage.Storage, c chain.Client, gen_obj chain.GenesisObject, genesisHash storage.GenesisHash, accounts *storage.AccountsMap) *Genesis {
	var genesis Genesis
	genesis.storage = s
	genesis.genesis = gen_obj
	genesis.hash = string(genesisHash)
	genesis.accounts = accounts
	genesis.bc = *MakeBlockContext(s, c, &common.Config{Chain: gen_obj.ToChainConfig()}, accounts)

	return &genesis
}

func (g *Genesis) makeInitBalanceTrx(addr string, value *big.Int) (trx *storage.Transaction) {
	trx = &storage.Transaction{}
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
		value := common.ParseStringToBigInt(value)
		trx := g.makeInitBalanceTrx(addr, value)
		g.bc.SaveTransaction(trx, false)
		genesisSupply.Add(genesisSupply, value)
		g.accounts.AddToBalance(trx.To, value)
	}
	for _, validator := range g.genesis.Dpos.InitialValidators {
		for addr, value := range validator.Delegations {
			delegation := common.ParseStringToBigInt(value)
			g.accounts.AddToBalance(addr, big.NewInt(0).Neg(delegation))
			g.accounts.AddToBalance(common.DposContractAddress, delegation)
		}
	}
	log.WithField("count", len(g.genesis.InitialBalances)).Info("Genesis: Init balance transactions parsed")

	// Genesis transactions isn't real transactions, so don't count it here
	g.bc.Batch.SaveAccounts(g.accounts)
	g.bc.finalized.TrxCount = 0
	g.bc.Batch.SetGenesisHash(storage.GenesisHash(g.hash))
	g.bc.Batch.SetTotalSupply(genesisSupply)
	g.bc.commit()
}
