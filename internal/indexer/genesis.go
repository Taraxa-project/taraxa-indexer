package indexer

import (
	"math/big"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/models"
	log "github.com/sirupsen/logrus"
)

type Genesis struct {
	storage storage.Storage
	genesis *chain.GenesisObject
	bc      *blockContext
	hash    string
}

func MakeGenesis(s storage.Storage, c *chain.WsClient, gen_obj *chain.GenesisObject, genesisHash storage.GenesisHash) (*Genesis, error) {
	var genesis Genesis
	var err error
	genesis.storage = s
	genesis.genesis = gen_obj
	genesis.hash = string(genesisHash)
	genesis.bc = MakeBlockContext(s, c, &common.Config{Chain: gen_obj.ToChainConfig()})

	return &genesis, err
}

func (g *Genesis) makeInitBalanceTrx(addr, value string) *models.Transaction {
	var trx models.Transaction
	trx.Hash = "GENESIS_" + addr
	trx.From = "GENESIS"
	trx.To = addr
	trx.Value = value
	trx.BlockNumber = 0
	trx.Timestamp = chain.ParseUInt(g.genesis.DagGenesisBlock.Timestamp)
	trx.Status = true
	return &trx
}

func (g *Genesis) process() {
	genesisSupply := big.NewInt(0)
	for addr, value := range g.genesis.InitialBalances {
		trx := g.makeInitBalanceTrx(addr, value)
		g.bc.SaveTransaction(trx)
		genesisSupply.Add(genesisSupply, common.ParseStringToBigInt(trx.Value))
	}
	log.WithField("count", len(g.genesis.InitialBalances)).Info("Genesis: Init balance transactions parsed")

	// Genesis transactions isn't real transactions, so don't count it here
	g.bc.finalized.TrxCount = 0
	g.bc.Batch.SetGenesisHash(storage.GenesisHash(g.hash))
	g.bc.Batch.SetTotalSupply(genesisSupply)
	g.bc.commit()
}
