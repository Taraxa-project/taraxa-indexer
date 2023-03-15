package storage

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/cockroachdb/pebble"
	"github.com/cockroachdb/pebble/vfs"
	"github.com/ethereum/go-ethereum/rlp"
	log "github.com/sirupsen/logrus"
)

type Storage struct {
	db   *pebble.DB
	path string
}

func NewStorage(file string) *Storage {
	db, err := open(file)
	if err != nil {
		log.WithError(err).Fatal("Can't create storage")
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

func GetTotal[T Paginated](s *Storage, address string) (r uint64) {
	stats := s.GetAddressStats(address)

	var o T
	switch t := any(o).(type) {
	case models.Dag:
		r = stats.DagsCount
	case models.Pbft:
		r = stats.PbftCount
	case models.Transaction:
		r = stats.TransactionsCount
	default:
		log.WithField("type", t).Fatal("GetCount incorrect type passed")
	}
	return
}

func GetObjectsPage[T Paginated](s *Storage, address string, from, count uint64) (ret []T, pagination *models.PaginatedResponse) {
	var o T
	ret = make([]T, 0, count)

	pagination = new(models.PaginatedResponse)
	pagination.Start = from
	pagination.Total = GetTotal[T](s, address)
	end := from + count
	pagination.HasNext = (end < pagination.Total)
	if end > pagination.Total {
		end = pagination.Total
	}
	pagination.End = end

	prefix := getPrefixKey(getPrefix(&o), address)
	start := getKey(getPrefix(&o), address, pagination.Total-from)

	iter := s.find(prefix)
	defer iter.Close()
	iter.SeekGE(start)

	for ; iter.Valid(); iter.Prev() {
		err := rlp.DecodeBytes(iter.Value(), &o)
		if err != nil {
			log.WithFields(log.Fields{"type": GetTypeName[T](), "error": err}).Fatal("Error decoding data from db")
		}
		ret = append(ret, o)
		if uint64(len(ret)) == count {
			return
		}
	}
	return
}

func getPrefix(o interface{}) (ret string) {
	switch tt := o.(type) {
	case *models.Transaction:
		ret = "t"
	case *models.Pbft:
		ret = "p"
	case *models.Dag:
		ret = "d"
	case *AddressStats:
		ret = "s"
	case *FinalizationData:
		ret = "f"
	case *GenesisHash:
		ret = "g"
	case *WeekStats:
		ret = "w"
	default:
		log.WithField("type", tt).Fatal("getPrefix: Unexpected type")
	}
	return
}

func getKey(prefix, author string, number uint64) []byte {
	author = strings.ToLower(author)
	return []byte(fmt.Sprintf("%s%s%020d", prefix, author, number))
}

func getPrefixKey(prefix, author string) []byte {
	author = strings.ToLower(author)
	return []byte(fmt.Sprintf("%s%s", prefix, author))
}

func getWeekKey(prefix string, year, week int32) []byte {
	return []byte(fmt.Sprintf("%s%d%02d", prefix, year, week))
}

func (s *Storage) GetWeekStats(year, week int32) WeekStats {
	ptr := MakeEmptyWeekStats()
	ptr.key = []byte(getWeekKey(getPrefix(ptr), year, week))
	err := s.getFromDB(ptr, ptr.key)
	if err != nil && err != pebble.ErrNotFound {
		log.WithError(err).Fatal("GetWeekStats failed")
	}
	return *ptr
}

func (s *Storage) addToDBTest(o interface{}, key1 string, key2 uint64) error {
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

func (s *Storage) GetFinalizationData() *FinalizationData {
	ptr := new(FinalizationData)
	err := s.getFromDB(ptr, []byte(getPrefix(ptr)))
	if err != nil && err != pebble.ErrNotFound {
		log.WithError(err).Fatal("GetFinalizationData failed")
	}
	return ptr
}

func (s *Storage) GetAddressStats(addr string) *AddressStats {
	ptr := MakeEmptyAddressStats(addr)
	err := s.getFromDB(ptr, getKey(getPrefix(ptr), addr, 0))
	if err != nil && err != pebble.ErrNotFound {
		log.Fatal("GetAddressStats ", err)
	}
	return ptr
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
		log.WithError(err).Fatal("GetGenesisHash failed")
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
		log.WithField("type", tt).Fatal("getFromDB: Unexpected type")
	}
	return nil
}
