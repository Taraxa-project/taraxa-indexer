package migration

import (
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage/pebble"
	log "github.com/sirupsen/logrus"
)

const migration_prefix = "mm"

type Migration interface {
	GetId() string
	Init(ws common.Client)
	Apply(s *pebble.Storage) error
}

type Manager struct {
	storage    *pebble.Storage
	migrations []Migration
	client     common.Client
}

func NewManager(s *pebble.Storage, client common.Client) *Manager {
	m := Manager{
		storage:    s,
		migrations: make([]Migration, 0),
		client:     client,
	}
	m.RegisterMigration(&CheckContracts{})
	m.RegisterMigration(&LastTrxTimestamp{})
	return &m
}

func (m *Manager) RegisterMigration(migration Migration) {
	m.migrations = append(m.migrations, migration)
	migration.Init(m.client)
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
			err = b.AddWithKey(migration.GetId(), []byte(migration_prefix+migration.GetId()))
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
