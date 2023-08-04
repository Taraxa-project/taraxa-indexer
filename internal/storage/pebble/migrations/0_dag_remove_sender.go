package migration

import (
	"github.com/Taraxa-project/taraxa-indexer/internal/storage/pebble"
	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/ethereum/go-ethereum/rlp"
	log "github.com/sirupsen/logrus"
)

type OldDag struct {
	Hash             models.Hash      `json:"hash"`
	Level            models.Counter   `json:"level"`
	Sender           models.Address   `json:"sender"`
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
	const DAG_BATCH_THRESHOLD = 1000
	batch := s.NewBatch()
	var last_key []byte

	for {
		var o OldDag
		count := 0
		s.ForEachFromKey([]byte(pebble.GetPrefix(models.Dag{})), last_key, func(key, res []byte) (stop bool) {
			err := rlp.DecodeBytes(res, &o)
			if err != nil {
				if err.Error() == "rlp: too few elements for migration.OldDag" {
					return false
				}
				log.WithFields(log.Fields{"migration": m.id, "error": err}).Fatal("Error decoding Dag")
			}
			dag := models.Dag{
				Hash:             o.Hash,
				Level:            o.Level,
				Timestamp:        o.Timestamp,
				TransactionCount: o.TransactionCount,
			}
			err = batch.AddToBatchFullKey(&dag, key)

			if err != nil {
				log.WithFields(log.Fields{"migration": m.id, "error": err}).Fatal("Error adding Dag to batch")
			}

			last_key = key
			count++
			return count == DAG_BATCH_THRESHOLD
		})
		batch.CommitBatch()
		batch = s.NewBatch()
		if count < DAG_BATCH_THRESHOLD {
			break
		}
	}

	return nil
}
