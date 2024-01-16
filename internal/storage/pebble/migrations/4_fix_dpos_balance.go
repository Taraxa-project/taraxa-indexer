package migration

import (
	"math/big"

	log "github.com/sirupsen/logrus"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage/pebble"
)

// FixDposBalance is a migration that removes the Sender attribute from the Dag struct.
type FixDposBalance struct {
	id            string
	blockchain_ws string
}

func (m *FixDposBalance) GetId() string {
	return m.id
}

// Apply is the implementation of the Migration interface for the FixDposBalance.
func (m *FixDposBalance) Apply(s *pebble.Storage) error {
	currentPeriod := s.GetFinalizationData().PbftCount
	diffAtBlock := uint64(6284000)
	balanceDiff := big.NewInt(378000)
	if currentPeriod < diffAtBlock {
		log.Info("FixDposBalance: Skipping migration as problematic block wasn't indexed yet")
		return nil
	}
	client, err := chain.NewWsClient(m.blockchain_ws)
	if err != nil {
		log.Fatal(err)
	}
	strBalance, _ := client.GetBalanceAtBlock(common.DposContractAddress, currentPeriod)
	chainBalance := common.ParseStringToBigInt(strBalance)

	accounts := s.GetAccounts()
	for i := 0; i < len(accounts); i++ {
		if accounts[i].Address != common.DposContractAddress {
			continue
		}
		// if chainBalance of dpos contract is smaller by balanceDiff, then sub it from the balance
		chainBalance.Add(chainBalance, balanceDiff)
		if accounts[i].Balance.Cmp(chainBalance) == 0 {
			before := big.NewInt(0).Set(accounts[i].Balance)
			accounts[i].Balance.Sub(accounts[i].Balance, balanceDiff)
			log.WithFields(log.Fields{"address": accounts[i].Address, "before": before, "after": accounts[i].Balance}).Info("FixDposBalance migration: Fixed balance")
		} else {
			log.Info("FixDposBalance: Skipping because balance is different from expected")
		}
	}
	b := s.NewBatch()
	b.SaveAccounts(accounts)
	b.CommitBatch()

	return nil
}
