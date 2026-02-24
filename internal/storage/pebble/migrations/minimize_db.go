package migration

import (
	"runtime"

	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage/pebble"
	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/ethereum/go-ethereum/rlp"
	log "github.com/sirupsen/logrus"
	"github.com/spiretechnology/go-pool"
)

type MinimizeDb struct {
	new_db           *pebble.Storage
	commitBatchSize  uint64
	maxTrxPerAccount uint64
	pool             pool.Pool
}

func (m *MinimizeDb) GetId() string {
	return "minimize_db"
}

func (m *MinimizeDb) Init() {
	m.new_db = pebble.NewStorage("./data/min_db")
	m.commitBatchSize = 10000
	m.maxTrxPerAccount = 100
	m.pool = pool.New(uint(5 * runtime.NumCPU()))
}

func (m *MinimizeDb) Apply(s *pebble.Storage) error {
	m.pool.Go(func() { m.MigrateSignificantData(s) })

	m.MigrateTransactions(s)

	m.pool.Wait()

	log.Info("Migrated, preparing toprint stats")
	pebble.PrintDbStats(m.new_db)
	return nil
}

func (m *MinimizeDb) MigrateObject(obj any, s *pebble.Storage) error {
	batch := m.new_db.NewBatch()
	s.ForEach(obj, "", nil, storage.Forward, func(key, res []byte) (stop bool) {
		err := batch.AddSerializedWithKey(obj, res, key)
		if err != nil {
			log.WithError(err).WithField("object_type", storage.GetObjectTypeName(obj)).Error("Failed to add serialized object to batch")
			return true
		}
		return false
	})
	batch.CommitBatch()
	return nil
}

func (m *MinimizeDb) MigrateSignificantData(s *pebble.Storage) {
	toMigrate := []any{
		&storage.Accounts{},
		&storage.TotalSupply{},
		&storage.YieldSaving{},
		storage.GenesisHash(""),
		&storage.Lambda{},
		&common.FinalizationData{},
		&storage.WeekStats{},
		&storage.RewardsStats{},
		&models.Pbft{},
		&models.Dag{}, //10
		&storage.AddressStats{},
		&storage.Yield{},
		&storage.ValidatorsYield{},
		&storage.MultipliedYield{},
		&storage.MonthlyActiveAddresses{},
	}
	for _, obj := range toMigrate {
		m.pool.Go(func() {
			err := m.MigrateObject(obj, s)
			if err != nil {
				log.WithError(err).WithField("object_type", storage.GetObjectTypeName(obj)).Error("Failed to migrate object type")
			} else {
				log.WithField("object_type", storage.GetObjectTypeName(obj)).Info("Successfully migrated")
			}
		})
	}
}

func (m *MinimizeDb) MigrateTransactions(s *pebble.Storage) {
	s.ForEach(storage.AddressStats{}, "", nil, storage.Forward, func(key, res []byte) (stop bool) {
		localStats := storage.AddressStats{}
		err := rlp.DecodeBytes(res, &localStats)
		if err != nil {
			log.WithError(err).WithField("address", localStats.Address).Error("Failed to decode address stats")
			return true
		}

		balance := localStats.Balance
		// do not copy transactions for accounts with zero balance, as they are not significant and only take space in the db
		if balance == nil || balance.Cmp(common.TARA) < 0 {
			return false
		}
		// log.WithField("acc", localStats.Address).Info("Migrating transactions for account with non-zero balance")
		m.pool.Go(func() { m.migrateAccountTransactions(&localStats, s) })
		return false
	})
}

func (m *MinimizeDb) migrateAccountTransactions(stats *storage.AddressStats, s *pebble.Storage) {
	batch := m.new_db.NewBatch()
	trx := models.Transaction{}
	count := uint64(0)
	s.ForEach(trx, stats.Address, nil, storage.Backward, func(key, res []byte) (stop bool) {
		err := batch.AddSerializedWithKey(&trx, res, key)
		if err != nil {
			log.WithError(err).WithField("address", stats.Address).Error("Failed to add transaction to batch")
			return true
		}
		count++
		return count >= m.maxTrxPerAccount
	})

	batch.CommitBatch()

	if count == m.maxTrxPerAccount {
		log.WithFields(log.Fields{
			"address":               stats.Address,
			"migrated_transactions": count,
			"total":                 stats.TransactionsCount,
		}).Info("CleanupInternalTrx migrated")
	}

}
