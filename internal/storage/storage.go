package storage

import (
	"fmt"
	"io"
	"log"

	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/cockroachdb/pebble"
	"github.com/cockroachdb/pebble/vfs"
	"github.com/ethereum/go-ethereum/rlp"
)

type Storage struct {
	DB *pebble.DB
}

func NewStorage(file string) *Storage {
	var ops pebble.Options
	if file == "" {
		ops.FS = vfs.NewMem()
	}
	db, err := pebble.Open(file, &ops)
	if err != nil {
		log.Fatal(err)
	}

	return &Storage{
		DB: db,
	}
}

func (s *Storage) add(key, value []byte) error {
	err := s.DB.Set(key, value, pebble.NoSync)
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

// in case from = 0 we return the upper bound
func (s *Storage) GetObjects(o interface{}, hash string, from uint64, count int) ([]interface{}, error) {
	ret := make([]interface{}, 0, count)

	upper := getPrefixKey(getPrefix(o), hash)
	start := getKey(getPrefix(o), hash, from)

	iter := s.find(upper)
	defer iter.Close()

	if from != 0 {
		iter.SeekGE(start)
	} else {
		iter.Last()
	}

	for ; iter.Valid(); iter.Prev() {
		err := rlp.DecodeBytes(iter.Value(), o)
		if err != nil {
			return nil, err
		}
		ret = append(ret, o)
		if len(ret) == count {
			return ret, nil
		}
	}
	return ret, nil
}

func getPrefix(o interface{}) string {
	switch tt := o.(type) {
	case *models.Transaction:
		return "t"
	case *models.Pbft:
		return "p"
	case *models.Dag:
		return "d"
	case *AddressStats:
		return "s"
	default:
		err := fmt.Errorf("getPrefix: Unexpected type %T", tt)
		panic(err)
	}
}

func getKey(prefix, author string, number uint64) []byte {
	return []byte(fmt.Sprintf("%s%s%020d", prefix, author, number))
}

func getPrefixKey(prefix, author string) []byte {
	return []byte(fmt.Sprintf("%s%s", prefix, author))
}

func (s *Storage) AddToDB(o interface{}, key1 string, key2 uint64) error {
	b, err := rlp.EncodeToBytes(o)
	if err != nil {
		return err
	}

	err = s.add(getKey(getPrefix(o), key1, key2), b)
	return err
}

func (s *Storage) GetFromDB(o interface{}, hash string) error {
	switch tt := o.(type) {
	case *AddressStats:
		key := getKey(getPrefix(o), hash, 0)
		value, closer, err := s.get(key)
		if err != nil {
			return err
		}
		err = rlp.DecodeBytes(value, o)
		if err != nil {
			return err
		}
		if err := closer.Close(); err != nil {
			return err
		}
	default:
		err := fmt.Errorf("GetFromDB: Unexpected type %T", tt)
		panic(err)
	}
	return nil
}
