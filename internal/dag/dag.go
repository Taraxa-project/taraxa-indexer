package dag

import (
	"encoding/json"
	"fmt"

	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	m "github.com/Taraxa-project/taraxa-indexer/models"
)

var prefix = "dag"

func getKey(sender string, age uint64) string {
	return fmt.Sprintf("%s:%s:%d", prefix, sender, age)
}

// Dag defines the model for a DAG Block.
type Dag struct {
	m.Dag
}

func (d *Dag) AddToDB(s *storage.Storage) error {
	b, err := json.Marshal(d)

	if err != nil {
		return err
	}

	err = s.Add([]byte(getKey(d.Sender, d.Age)), b)
	if err != nil {
		return err
	}

	return nil
}
