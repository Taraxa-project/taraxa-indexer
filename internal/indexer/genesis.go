package indexer

import (
	"fmt"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/models"
)

type Genesis struct {
	storage *storage.Storage
	genesis *chain.GenesisObject
	bc      *blockContext
	hash    string
}

func MakeGenesis(s *storage.Storage, c *chain.WsClient, genesisHash string) *Genesis {
	var genesis Genesis
	genesis.storage = s
	genesis.genesis = c.GetGenesis()
	genesis.hash = genesisHash
	genesis.bc = MakeBlockContext(s, c)

	return &genesis
}

func (g *Genesis) makeInitBalanceTrx(addr, value string) *models.Transaction {
	var trx models.Transaction
	trx.Hash = "GENESIS_" + addr
	trx.From = "GENESIS"
	trx.To = addr
	trx.Value = value
	trx.BlockNumber = 0
	trx.Timestamp = chain.ParseInt(g.genesis.DagGenesisBlock.Timestamp)
	return &trx
}

func (g *Genesis) process() {
	for addr, value := range g.genesis.InitialBalances {
		trx := g.makeInitBalanceTrx(addr, value)
		g.bc.SaveTransaction(trx)
	}
	fmt.Println("GENESIS:", len(g.genesis.InitialBalances), "init balance transactions parsed")

	// Genesis transactions isn't real transactions, so don't count it here
	g.bc.finalized.TrxCount = 0
	g.bc.batch.SaveGenesisHash(storage.GenesisHash(g.hash))
	g.bc.commit(0)
}
