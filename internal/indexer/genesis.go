package indexer

import (
	"math/big"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/models"
	log "github.com/sirupsen/logrus"
	"github.com/spiretechnology/go-pool"
)

type Genesis struct {
	storage *storage.Storage
	genesis *chain.GenesisObject
	bc      *blockContext
	hash    string
}

func MakeGenesis(s *storage.Storage, c *chain.WsClient, tp pool.Pool, genesisHash storage.GenesisHash) (*Genesis, error) {
	var genesis Genesis
	var err error
	genesis.storage = s
	genesis.genesis, err = c.GetGenesis()
	genesis.hash = string(genesisHash)
	genesis.bc = MakeBlockContext(s, c, tp)

	return &genesis, err
}

func (g *Genesis) makeInitBalanceTrx(addr, value string) *models.Transaction {
	var trx models.Transaction
	trx.Hash = "GENESIS_" + addr
	trx.From = "GENESIS"
	trx.To = addr
	trx.Value = value
	trx.BlockNumber = 0
	trx.Timestamp = chain.ParseInt(g.genesis.DagGenesisBlock.Timestamp)
	trx.Status = true
	return &trx
}

func (g *Genesis) process() {
	genesisSupply := big.NewInt(0)
	for addr, value := range g.genesis.InitialBalances {
		trx := g.makeInitBalanceTrx(addr, value)
		g.bc.SaveTransaction(trx)
		genesisSupply.Add(genesisSupply, parseStringToBigInt(trx.Value))
	}

	// TODO[45]: Old genesis structure. Remove
	for addr, value := range g.genesis.FinalChain.State.GenesisBalances {
		trx := g.makeInitBalanceTrx(addr, value)
		g.bc.SaveTransaction(trx)
		genesisSupply.Add(genesisSupply, parseStringToBigInt(trx.Value))
	}
	log.WithField("count", len(g.genesis.InitialBalances)).Info("Genesis: Init balance transactions parsed")

	// Genesis transactions isn't real transactions, so don't count it here
	g.bc.finalized.TrxCount = 0
	g.bc.batch.SetGenesisHash(storage.GenesisHash(g.hash))
	g.bc.batch.SetTotalSupply(genesisSupply)
	g.bc.commit(0)
}
