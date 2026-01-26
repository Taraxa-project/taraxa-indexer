package migration

import (
	"math/big"
	"time"

	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage/pebble"
	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/ethereum/go-ethereum/rlp"
	log "github.com/sirupsen/logrus"
)

type BalanceToAddressStats struct {
}

type oldAddressStats struct {
	models.StatsResponse
	Address string `json:"address"`
}

func (m *BalanceToAddressStats) GetId() string {
	return "balance_to_address_stats"
}

func (m *BalanceToAddressStats) Init(client common.Client) {
}

func (m *BalanceToAddressStats) MigrateAccounts(s *pebble.Storage) error {
	batch := s.NewBatch()
	accounts := make([]storage.Account, 0)
	err := s.GetFromDB(&accounts, pebble.GetPrefixKey(pebble.GetPrefix(&storage.Accounts{}), ""))
	if err != nil {
		return err
	}
	new_accounts := storage.Accounts{Accounts: accounts, Total: uint64(len(accounts))}
	batch.AddSingleKey(new_accounts, "")
	log.WithFields(log.Fields{"total": len(accounts)}).Info("Migrated accounts")
	batch.CommitBatch()
	return nil
}

func (m *BalanceToAddressStats) Apply(s *pebble.Storage) error {
	err := m.MigrateAccounts(s)
	// if the accounts are not found, we don't need to migrate the address stats, it means that indexer was not running yet
	if err == pebble.ErrNotFound {
		return nil
	}
	if err != nil {
		return err
	}

	batch := s.NewBatch()
	balances := s.GetAccounts().ToMap()
	count := 0
	s.ForEach(storage.AddressStats{}, "", nil, storage.Forward, func(key, res []byte) (stop bool) {

		var oldAddressStats oldAddressStats
		err := rlp.DecodeBytes(res, &oldAddressStats)
		if err != nil {

			return true
		}
		balance := balances.GetBalance(oldAddressStats.Address)
		if balance == nil {
			balance = big.NewInt(0)
		}
		addressStats := storage.AddressStats{
			StatsResponse: oldAddressStats.StatsResponse,
			Address:       oldAddressStats.Address,
			Balance:       balance,
		}
		err = batch.AddWithKey(&addressStats, key)
		count++
		if err != nil {
			log.WithError(err).WithFields(log.Fields{"address": addressStats.Address}).Error("AddWithKey failed")
			return true
		}
		return false
	})
	log.WithFields(log.Fields{"count": count}).Info("BalanceToAddressStats migrated")
	batch.CommitBatch()
	return nil
}

func (m *BalanceToAddressStats) TestIterTime(s *pebble.Storage) {
	start := time.Now()
	count := 0
	s.ForEach(storage.AddressStats{}, "", nil, storage.Forward, func(key, res []byte) (stop bool) {
		count++
		return false
	})
	log.WithFields(log.Fields{"time": time.Since(start), "count": count}).Info("TestIterTime")
}
