package migration

import (
	"github.com/Taraxa-project/taraxa-indexer/internal/storage/pebble"
	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/ethereum/go-ethereum/rlp"
	log "github.com/sirupsen/logrus"
)

type OldDag struct {
	Hash             models.Hash      `json:"hash"`
	Sender           models.Address   `json:"sender"`
	Level            models.Counter   `json:"level"`
	Timestamp        models.Timestamp `json:"timestamp"`
	TransactionCount models.Counter   `json:"transactionCount"`
}

// RemoveSenderMigration is a migration that removes the Sender attribute from the Dag struct.
type RemoveSenderMigration struct {
	id string
}

func (m *RemoveSenderMigration) GetId() string {
	return m.id
}

// Apply is the implementation of the Migration interface for the RemoveSenderMigration.
func (m *RemoveSenderMigration) Apply(s *pebble.Storage) error {
	// Retrieve all Dags from the database
	const DAG_BATCH_THRESHOLD = 1000
	batch := s.NewBatch()
	var last_key []byte

	var done = false

	for !done {
		var o OldDag
		count := 0
		s.ForEachFromKey(&o, last_key, func(key, res []byte) bool {
			err := rlp.DecodeBytes(res, &o)
			if err != nil {
				log.WithFields(log.Fields{"migration": m.id, "error": err}).Fatal("Error decoding Dag")
			}
			dag := models.Dag{
				Hash:             o.Hash,
				Level:            o.Level,
				Timestamp:        o.Timestamp,
				TransactionCount: o.TransactionCount,
			}
			batch.AddToBatchFullKey(&dag, key)

			last_key = key
			count++
			if count == DAG_BATCH_THRESHOLD {
				return true
			}
			return false
		})

		if count < DAG_BATCH_THRESHOLD {
			batch.CommitBatch()
			done = true
			break
		}
	}

	return nil
}
