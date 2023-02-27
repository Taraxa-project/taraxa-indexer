package storage

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

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

type Paginated interface {
	models.Transaction | models.Dag | models.Pbft
}

func ParseKeyIndex(key, prefix string) uint64 {
	index_str := key[len(prefix):]

	index, err := strconv.ParseUint(index_str, 10, 64)
	if err != nil {
		log.Fatal("ParseKeyIndex ", key, prefix)
	}
	return index
}

func GetObjectsPage[T Paginated](s *Storage, hash string, from uint64, count int) (ret []T, pagination *models.PaginatedResponse, err error) {
	var o T
	ret = make([]T, 0, count)
	pagination = new(models.PaginatedResponse)
	prefix := getPrefixKey(getPrefix(&o), hash)
	start := getKey(getPrefix(&o), hash, from)

	iter := s.find(prefix)
	defer iter.Close()

	if from != 0 {
		iter.SeekGE(start)
	} else {
		iter.Last()
	}
	defer func() {
		pagination.HasNext = (pagination.End != 1)
	}()
	for ; iter.Valid(); iter.Prev() {
		err = rlp.DecodeBytes(iter.Value(), &o)
		if len(ret) == 0 {
			pagination.Start = ParseKeyIndex(string(iter.Key()), string(prefix))
		}
		if err != nil {
			return
		}
		ret = append(ret, o)
		pagination.End = ParseKeyIndex(string(iter.Key()), string(prefix))
		if len(ret) == count {
			return
		}
	}
	return
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
	case *GenesisHash:
		return "g"
	case *WeekStats:
		return "w"
	default:
		err := fmt.Errorf("getPrefix: Unexpected type %T", tt)
		panic(err)
	}
}

func getKey(prefix, author string, number uint64) []byte {
	author = strings.ToLower(author)
	return []byte(fmt.Sprintf("%s%s%020d", prefix, author, number))
}

func getPrefixKey(prefix, author string) []byte {
	author = strings.ToLower(author)
	return []byte(fmt.Sprintf("%s%s", prefix, author))
}

func getWeekKey(prefix string, year, week int) []byte {
	return []byte(fmt.Sprintf("%s%d%02d", prefix, year, week))
}

func (s *Storage) GetWeekStats(year, week int) WeekStats {
	ptr := MakeEmptyWeekStats()
	ptr.key = []byte(getWeekKey(getPrefix(ptr), year, week))
	err := s.getFromDB(ptr, ptr.key)
	if err != nil && err != pebble.ErrNotFound {
		log.Fatal("GetFinalizedPeriod ", err)
	}
	return *ptr
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

func (s *Storage) GetFinalizedPeriod() *FinalizationData {
	ptr := new(FinalizationData)
	err := s.getFromDB(ptr, []byte(getPrefix(ptr)))
	if err != nil && err != pebble.ErrNotFound {
		log.Fatal("GetFinalizedPeriod ", err)
	}
	return ptr
}

func (s *Storage) GetAddressStats(hash string) (ret *AddressStats, err error) {
	ret = new(AddressStats)
	err = s.getFromDB(ret, getKey(getPrefix(ret), hash, 0))
	return
}

func (s *Storage) GenesisHashExist() bool {
	ptr := new(GenesisHash)
	err := s.getFromDB(ptr, []byte(getPrefix(ptr)))
	return err == nil
}

func (s *Storage) GetGenesisHash() GenesisHash {
	ptr := new(GenesisHash)
	err := s.getFromDB(ptr, []byte(getPrefix(ptr)))
	if err != nil {
		log.Fatal("GetGenesisHash ", err)
	}
	return *ptr
}

func (s *Storage) getFromDB(o interface{}, key []byte) error {
	switch tt := o.(type) {
	case *AddressStats, *FinalizationData, *GenesisHash, *WeekStats:
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
