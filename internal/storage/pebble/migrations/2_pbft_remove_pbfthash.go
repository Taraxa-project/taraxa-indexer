package migration

import (
	"github.com/Taraxa-project/taraxa-indexer/internal/storage/pebble"
	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/ethereum/go-ethereum/rlp"
	log "github.com/sirupsen/logrus"
)

type OldPbft struct {
	Author           models.Address   `json:"author"`
	Hash             models.Hash      `json:"hash"`
	Number           models.Counter   `json:"number"`
	PbftHash         models.Hash      `json:"pbftHash"`
	Timestamp        models.Timestamp `json:"timestamp"`
	TransactionCount models.Counter   `json:"transactionCount"`
}

// RemoveSenderMigration is a migration that removes the Sender attribute from the Dag struct.
type RemovePbftHashMigration struct {
	id string
}

func (m *RemovePbftHashMigration) GetId() string {
	return m.id
}

// Apply is the implementation of the Migration interface for the RemoveSenderMigration.
func (m *RemovePbftHashMigration) Apply(s *pebble.Storage) error {
	const PBFT_BATCH_THRESHOLD = 1000
	batch := s.NewBatch()
	var last_key []byte

	for {
		var o OldPbft
		count := 0
		s.ForEachFromKey([]byte(pebble.PbftPrefix), last_key, func(key, res []byte) (stop bool) {
			err := rlp.DecodeBytes(res, &o)
			if err != nil {
				if err.Error() == "rlp: too few elements for migration.OldPbft" {
					return false
				}
				log.WithFields(log.Fields{"migration": m.id, "error": err}).Fatal("Error decoding Pbft")
			}
			pbft := models.Pbft{
				Author:           o.Author,
				Hash:             o.Hash,
				Number:           o.Number,
				Timestamp:        o.Timestamp,
				TransactionCount: o.TransactionCount,
			}
			err = batch.AddToBatchFullKey(&pbft, key)

			if err != nil {
				log.WithFields(log.Fields{"migration": m.id, "error": err}).Fatal("Error adding Pbft to batch")
			}

			last_key = key
			count++
			return count == PBFT_BATCH_THRESHOLD
		})
		batch.CommitBatch()
		batch = s.NewBatch()
		if count < PBFT_BATCH_THRESHOLD {
			break
		}
	}

	return nil
}
