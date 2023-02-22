package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/cockroachdb/pebble"
	"github.com/cockroachdb/pebble/vfs"
)

type Storage struct {
	DB *pebble.DB
}

func NewStorage(file string) *Storage {
	db, err := pebble.Open(file, &pebble.Options{FS: vfs.NewMem()})
	if err != nil {
		log.Fatal(err)
	}

	return &Storage{
		DB: db,
	}
}

func (s *Storage) add(key, value []byte) error {
	err := s.DB.Set(key, value, pebble.Sync)
	return err
}

func (s *Storage) get(key []byte) ([]byte, io.Closer, error) {
	value, closer, err := s.DB.Get(key)
	if err != nil {
		return nil, nil, err
	}

	return value, closer, nil
}

func (s *Storage) find(prefix []byte) *pebble.Iterator {
	keyUpperBound := func(b []byte) []byte {
		end := make([]byte, len(b))
		copy(end, b)
		for i := len(end) - 1; i >= 0; i-- {
			end[i] = end[i] + 1
			if end[i] != 0 {
				return end[:i+1]
			}
		}
		return nil // no upper-bound
	}

	prefixIterOptions := func(prefix []byte) *pebble.IterOptions {
		return &pebble.IterOptions{
			LowerBound: prefix,
			UpperBound: keyUpperBound(prefix),
		}
	}

	iter := s.DB.NewIter(prefixIterOptions(prefix))
	return iter
}

func getPrefix(o interface{}) string {
	switch tt := o.(type) {
	case *models.Transaction:
		return "tx"
	case *models.Pbft:
		return "pbft"
	case *models.Dag:
		return "dag"
	case *Address:
		return "address"
	default:
		err := fmt.Errorf("getPrefix: Unexpected type %T", tt)
		panic(err)
	}
}

func getKey(prefix, author string, number uint64) string {
	return fmt.Sprintf("%s:%s:%d", prefix, author, number)
}

func (s *Storage) AddToDB(o interface{}) error {
	b, err := json.Marshal(o)

	if err != nil {
		return err
	}

	switch tt := o.(type) {
	case *models.Transaction:
		err = s.add([]byte(getKey(getPrefix(o), tt.From, tt.Age)), b)
		if err != nil {
			return err
		}
		err = s.add([]byte(getKey(getPrefix(o), tt.To, tt.Age)), b)
		if err != nil {
			return err
		}
	case *models.Pbft:
		err = s.add([]byte(getKey(getPrefix(o), tt.Author, tt.Number)), b)
		if err != nil {
			return err
		}
	case *models.Dag:
		err = s.add([]byte(getKey(getPrefix(o), tt.Sender, tt.Age)), b)
		if err != nil {
			return err
		}
	case *Address:
		err = s.add([]byte(getKey(getPrefix(o), tt.Address, 0)), b)
		if err != nil {
			return err
		}
	default:
		err := fmt.Errorf("AddToDB: Unexpected type %T", tt)
		panic(err)
	}
	return nil
}

func (s *Storage) GetFromDB(o interface{}, hash string) error {
	switch tt := o.(type) {
	case *Address:
		key := []byte(getKey(getPrefix(o), hash, 0))
		value, closer, err := s.get(key)
		if err != nil {
			return err
		}
		err = json.Unmarshal(value, &o)
		if err != nil {
			return err
		}
		if err := closer.Close(); err != nil {
			return err
		}
	default:
		err := fmt.Errorf("AddToDB: Unexpected type %T", tt)
		panic(err)
	}
	return nil
}
