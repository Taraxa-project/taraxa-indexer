package storage

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/cockroachdb/pebble"
	"github.com/cockroachdb/pebble/vfs"
	"github.com/ethereum/go-ethereum/rlp"
)

type Storage struct {
	db   *pebble.DB
	path string
}

func NewStorage(file string) *Storage {
	db, err := open(file)
	if err != nil {
		log.Fatal(err)
	}

	return &Storage{
		db:   db,
		path: file,
	}
}

func open(file string) (*pebble.DB, error) {
	var ops pebble.Options
	if file == "" {
		ops.FS = vfs.NewMem()
	}
	return pebble.Open(file, &ops)
}

func (s *Storage) Clean() error {
	if err := s.db.Close(); err != nil {
		return err
	}
	if err := os.RemoveAll(s.path); err != nil {
		return err
	}
	db, err := open(s.path)
	if err != nil {
		return err
	}
	s.db = db
	return nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}

func (s *Storage) add(key, value []byte) error {
	err := s.db.Set(key, value, pebble.NoSync)
	return err
}

func (s *Storage) get(key []byte) ([]byte, io.Closer, error) {
	value, closer, err := s.db.Get(key)
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

	iter := s.db.NewIter(prefixIterOptions(prefix))
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
	case *FinalizationData:
		return "f"
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
	return s.addToDB(getKey(getPrefix(o), key1, key2), o)
}

func (s *Storage) addToDB(key []byte, o interface{}) error {
	b, err := rlp.EncodeToBytes(o)
	if err != nil {
		return err
	}

	err = s.add(key, b)
	return err
}

func (s *Storage) RecordFinalizedPeriod(f FinalizationData) error {
	return s.addToDB([]byte(getPrefix(&f)), &f)
}

func (s *Storage) FinalizedPeriodExists() bool {
	ptr := new(FinalizationData)
	err := s.getFromDB(ptr, []byte(getPrefix(ptr)))
	fmt.Println("FinalizedPeriodExists", err)
	return err == nil
}

func (s *Storage) GetFinalizedPeriod() FinalizationData {
	ptr := new(FinalizationData)
	err := s.getFromDB(ptr, []byte(getPrefix(ptr)))
	if err != nil {
		log.Fatal("GetFinalizedPeriod ", err)
	}
	return *ptr
}

func (s *Storage) GetFromDB(o interface{}, hash string) error {
	return s.getFromDB(o, getKey(getPrefix(o), hash, 0))
}

func (s *Storage) getFromDB(o interface{}, key []byte) error {
	switch tt := o.(type) {
	case *AddressStats, *FinalizationData:
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

type Batch struct {
	*pebble.Batch
}

func (s *Storage) NewBatch() *Batch {
	return &Batch{s.db.NewBatch()}
}

func (b *Batch) AddToBatch(o interface{}, key1 string, key2 uint64) error {
	data, err := rlp.EncodeToBytes(o)
	if err != nil {
		return err
	}
	return b.Set(getKey(getPrefix(o), key1, key2), data, nil)
}
