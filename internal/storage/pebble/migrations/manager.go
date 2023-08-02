package migration

import (
	"github.com/Taraxa-project/taraxa-indexer/internal/storage/pebble"
	"github.com/ethereum/go-ethereum/common"
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

func NewManager(s *pebble.Storage, nodeAddr common.Address) *Manager {
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

func (m *Manager) ApplyAll() error {
	for _, migration := range m.migrations {
		if !m.IsApplied(migration) {
			err := migration.Apply(m.s)
			if err != nil {
				return err
			}
			b := m.s.NewBatch()
			b.AddToBatchFullKey(migration.GetId(), []byte(migration_prefix+migration.GetId()))
			b.CommitBatch()
		}
	}
	return nil
}
