package migration

import (
	"github.com/Taraxa-project/taraxa-indexer/internal/storage/pebble"
	"github.com/ethereum/go-ethereum/common"
)

type Migration interface {
	GetId() string
	Apply(s *pebble.Storage) error
}

type Manager struct {
	s          *pebble.Storage
	migrations []Migration
}

func NewManager(s *pebble.Storage, nodeAddr common.Address) *Manager {
	return &Manager{
		s: s,
	}
}

func (m *Manager) RegisterMigration(migration Migration) {
	m.migrations = append(m.migrations, migration)
}

func (m *Manager) IsApplied(migration Migration) bool {
	migrationId := m.s.GetMigration(migration.GetId())
	return migrationId != ""
}

func (m *Manager) ApplyAll() error {
	for _, migration := range m.migrations {
		err := migration.Apply(m.s)
		if err != nil {
			return err
		}
	}
	return nil
}
