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
	s          *pebble.Storage
	migrations []Migration
}

func NewManager(s *pebble.Storage) *Manager {
	m := Manager{
		s: s,
	}
	m.RegisterMigration(&RemoveSenderMigration{id: "0_dag_removeSender"})
	return &m
}

func (m *Manager) RegisterMigration(migration Migration) {
	m.migrations = append(m.migrations, migration)
}

func (m *Manager) IsApplied(migration Migration) bool {
	migrationId := ""
	err := m.s.GetFromDB(&migrationId, []byte(migration_prefix+migration.GetId()))
	return err == nil
}

func (m *Manager) ApplyAll() (err error) {
	log.Info("Migration Manager: Running migrations")
	for _, migration := range m.migrations {
		isApplied := m.IsApplied(migration)
		if !isApplied {
			log.WithFields(log.Fields{"migration": migration.GetId()}).Info("Running migration")
			err = migration.Apply(m.s)
			if err != nil {
				return
			}
			b := m.s.NewBatch()
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
