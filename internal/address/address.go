package address

import (
	"encoding/json"
	"fmt"

	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
)

var prefix = "address"

func getKey(addr string) string {
	return fmt.Sprintf("%s:%s", prefix, addr)
}

// Address defines the model for an address aggregate.
type Address struct {
	Address   string `json:"address"`
	TxTotal   uint64 `json:"txTotal"`
	DagTotal  int64  `json:"dagTotal"`
	PbftTotal int64  `json:"pbftTotal"`
}

func (a *Address) AddToDB(s *storage.Storage) error {
	b, err := json.Marshal(a)

	if err != nil {
		return err
	}

	err = s.Add([]byte(getKey(a.Address)), b)
	if err != nil {
		return err
	}

	return nil
}

func GetFromDB(s *storage.Storage, addr string) (Address, error) {
	key := []byte(getKey(addr))

	value, closer, err := s.Get(key)
	if err != nil {
		return Address{}, err
	}

	var a Address
	err = json.Unmarshal(value, &a)
	if err != nil {
		return Address{}, err
	}

	if err := closer.Close(); err != nil {
		return Address{}, err
	}

	return a, nil
}
