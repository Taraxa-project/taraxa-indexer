package pbft

import (
	"encoding/json"
	"fmt"

	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	m "github.com/Taraxa-project/taraxa-indexer/models"
)

var prefix = "pbft"

func getKey(author string, number uint64) string {
	return fmt.Sprintf("%s:%s:%d", prefix, author, number)
}

// Pbft defines the model for a PBFT Block.
type Pbft struct {
	m.Pbft
}

func (p *Pbft) AddToDB(s *storage.Storage) error {
	b, err := json.Marshal(p)
	if err != nil {
		return err
	}

	err = s.Add([]byte(getKey(p.Author, p.Number)), b)
	if err != nil {
		return err
	}

	return nil
}
