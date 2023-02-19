package tx

import (
	"encoding/json"
	"fmt"

	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	m "github.com/Taraxa-project/taraxa-indexer/models"
)

var prefix = "tx"

func getKey(address string, age int64) string {
	return fmt.Sprintf("%s:%s:%d", prefix, address, age)
}

// Tx defines the model for a Transaction.
type Tx struct {
	m.Transaction
}

func (t *Tx) AddToDB(s *storage.Storage) error {
	b, err := json.Marshal(t)

	if err != nil {
		return err
	}

	err = s.Add([]byte(getKey(t.From, t.Age)), b)
	if err != nil {
		return err
	}

	err = s.Add([]byte(getKey(t.To, t.Age)), b)
	if err != nil {
		return err
	}

	return nil
}
