package migration

import (
	"github.com/Taraxa-project/taraxa-indexer/internal/storage/pebble"
	log "github.com/sirupsen/logrus"
)

const migration_prefix = "mm"

type Migration interface {
	GetId() string
	Apply(s *pebble.Storage) error
}

type Manager struct {
	storage    *pebble.Storage
	migrations []Migration
}

func NewManager(s *pebble.Storage) *Manager {
	m := Manager{
		storage: s,
	}
	m.RegisterMigration(&RemoveSenderMigration{id: "0_dag_removeSender"})
	m.RegisterMigration(&RemoveNonceTxIndexAddFeeMigration{id: "1_tx_remove_nonce_txIndex_gas"})
	m.RegisterMigration(&RemovePbftHashMigration{id: "2_pbft_removePbftHash"})
	return &m
}

func (m *Manager) RegisterMigration(migration Migration) {
	m.migrations = append(m.migrations, migration)
}

func (m *Manager) IsApplied(migration Migration) bool {
	migrationId := ""
	err := m.storage.GetFromDB(&migrationId, []byte(migration_prefix+migration.GetId()))
	return err == nil
}

func (m *Manager) ApplyAll() (err error) {
	log.Info("Migration Manager: Running migrations")
	for _, migration := range m.migrations {
		isApplied := m.IsApplied(migration)
		if !isApplied {
			log.WithFields(log.Fields{"migration": migration.GetId()}).Info("Running migration")
			err = migration.Apply(m.storage)
			if err != nil {
				return
			}
			b := m.storage.NewBatch()
			err = b.AddToBatchFullKey(migration.GetId(), []byte(migration_prefix+migration.GetId()))
			if err != nil {
				return
			}
			b.CommitBatch()
			log.WithFields(log.Fields{"migration": migration.GetId()}).Info("Applied migration")
		} else {
			log.WithFields(log.Fields{"migration": migration.GetId()}).Info("Skipping migration")
		}
	}
	log.Info("Migration Manager: Applied all migrations")
	return nil
}
