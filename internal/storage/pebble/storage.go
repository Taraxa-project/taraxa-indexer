package pebble

import (
	"fmt"
	"io"
	"math/big"
	"os"
	"strings"
	"sync"

	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
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

func (s *Storage) NewBatch() storage.Batch {
	return &Batch{Batch: s.db.NewBatch(), Mutex: new(sync.RWMutex)}
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

func getPrefix(o interface{}) (ret string) {

	switch tt := o.(type) {
	case *[]models.Account, []models.Account:
		ret = "a"
	case *models.TransactionLogsResponse, models.TransactionLogsResponse:
		ret = "e"
	case *models.Transaction:
		ret = "t"
	case *models.Pbft:
		ret = "p"
	case *models.Dag:
		ret = "d"
	case *storage.AddressStats:
		ret = "s"
	case *storage.FinalizationData:
		ret = "f"
	case *storage.GenesisHash:
		ret = "g"
	case *storage.WeekStats:
		ret = "w"
	case *storage.TotalSupply:
		ret = "ts"
	case *models.InternalTransactionsResponse:
		ret = "i"
	// hack if we aren't passing original type directly to this method, but passing interface from other one
	case *interface{}:
		ret = getPrefix(*o.(*interface{}))
	default:
		log.WithFields(log.Fields{"type": tt, "value": o}).Fatalf("getPrefix: Unexpected type %T", tt)
	}
	return
}

func getKey(prefix, key1 string, key2 uint64) []byte {
	key1 = strings.ToLower(key1)
	return []byte(fmt.Sprintf("%s%s%020d", prefix, key1, key2))
}

func getPrefixKey(prefix, author string) []byte {
	author = strings.ToLower(author)
	return []byte(fmt.Sprintf("%s%s", prefix, author))
}

func getWeekKey(prefix string, year, week int32) []byte {
	return []byte(fmt.Sprintf("%s%d%02d", prefix, year, week))
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

	prefixIterOptions := pebble.IterOptions{
		LowerBound: prefix,
		UpperBound: keyUpperBound(prefix),
	}

	iter := s.db.NewIter(&prefixIterOptions)
	return iter
}

func (s *Storage) ForEach(o interface{}, key_prefix string, start uint64, fn func([]byte) (stop bool)) {
	prefix := getPrefixKey(getPrefix(&o), key_prefix)
	start_key := getKey(getPrefix(&o), key_prefix, start)

	iter := s.find(prefix)
	defer iter.Close()
	iter.SeekGE(start_key)

	for ; iter.Valid(); iter.Prev() {
		if fn(iter.Value()) {
			break
		}
	}
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

func (s *Storage) GetTotalSupply() *storage.TotalSupply {
	ptr := big.NewInt(0)
	err := s.getFromDB(ptr, []byte(getPrefix((*storage.TotalSupply)(ptr))))
	if err != nil {
		log.Fatal("GetTotalSupply ", err)
	}
	return ptr
}

func (s *Storage) GetAccounts() []models.Account {
	ptr := new([]models.Account)
	err := s.getFromDB(ptr, getKey(getPrefix(ptr), "0x0", 0))
	if err != nil && err != pebble.ErrNotFound {
		log.Fatal("GetAccounts failed: ", err)
	}
	return *ptr
}

func (s *Storage) GetWeekStats(year, week int32) storage.WeekStats {
	ptr := storage.MakeEmptyWeekStats()
	ptr.Key = []byte(getWeekKey(getPrefix(ptr), year, week))
	err := s.getFromDB(ptr, ptr.Key)
	if err != nil && err != pebble.ErrNotFound {
		log.WithError(err).Fatal("GetWeekStats failed")
	}
	return *ptr
}

func (s *Storage) GetFinalizationData() *storage.FinalizationData {
	ptr := new(storage.FinalizationData)
	err := s.getFromDB(ptr, []byte(getPrefix(ptr)))
	if err != nil && err != pebble.ErrNotFound {
		log.WithError(err).Fatal("GetFinalizationData failed")
	}
	return ptr
}

func (s *Storage) GetAddressStats(addr string) *storage.AddressStats {
	ptr := storage.MakeEmptyAddressStats(addr)
	err := s.getFromDB(ptr, getKey(getPrefix(ptr), addr, 0))
	if err != nil && err != pebble.ErrNotFound {
		log.Fatal("GetAddressStats ", err)
	}
	return ptr
}

func (s *Storage) GenesisHashExist() bool {
	ptr := new(storage.GenesisHash)
	err := s.getFromDB(ptr, []byte(getPrefix(ptr)))
	return err == nil
}

func (s *Storage) GetGenesisHash() storage.GenesisHash {
	ptr := new(storage.GenesisHash)
	err := s.getFromDB(ptr, []byte(getPrefix(ptr)))
	if err != nil {
		log.WithError(err).Fatal("GetGenesisHash failed")
	}
	return *ptr
}

func (s *Storage) GetInternalTransactions(hash string) models.InternalTransactionsResponse {
	ptr := new(models.InternalTransactionsResponse)
	err := s.getFromDB(ptr, getPrefixKey(getPrefix(ptr), hash))
	if err != nil && err != pebble.ErrNotFound {
		log.WithError(err).Fatal("GetInternalTransactions failed")
	}
	return *ptr
}

func (s *Storage) GetTransactionLogs(hash string) models.TransactionLogsResponse {
	ptr := new(models.TransactionLogsResponse)
	err := s.getFromDB(ptr, getPrefixKey(getPrefix(ptr), hash))
	if err != nil && err != pebble.ErrNotFound {
		log.WithError(err).Fatal("GetTransactionLogs failed")
	}
	return *ptr
}

func (s *Storage) getFromDB(o interface{}, key []byte) error {
	switch tt := o.(type) {
	case *[]models.Account, *storage.AddressStats, *storage.FinalizationData, *storage.GenesisHash, *storage.WeekStats, *storage.TotalSupply, *models.InternalTransactionsResponse, *models.TransactionLogsResponse:
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
