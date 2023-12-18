package pebble

import (
	"fmt"
	"io"
	"math/big"
	"os"
	"runtime/debug"
	"strings"
	"sync"

	"github.com/Taraxa-project/taraxa-indexer/internal/events"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/cockroachdb/pebble"
	"github.com/cockroachdb/pebble/vfs"
	"github.com/ethereum/go-ethereum/rlp"
	log "github.com/sirupsen/logrus"
)

const prefixSeparator = "|"
const accountPrefix = "b"
const logsPrefix = "e"
const transactionPrefix = "t"
const pbftPrefix = "p"
const dagsPrefix = "d"
const statsPrefix = "s"
const finalizationDataPrefix = "f"
const genesisHashPrefix = "g"
const weekStatsPrefix = "w"
const totalSupplyPrefix = "ts"
const internalTransactionsPrefix = "i"
const yieldPrefix = "y"
const validatorsYieldPrefix = "vy"
const multipliedYieldPrefix = "my"
const PeriodRewardsPrefix = "pr"

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

func GetPrefix(o interface{}) (ret string) {
	switch tt := o.(type) {
	case *storage.Accounts, storage.Accounts:
		ret = accountPrefix
	case *models.TransactionLogsResponse, models.TransactionLogsResponse:
		ret = logsPrefix
	case *models.Transaction, models.Transaction:
		ret = transactionPrefix
	case *models.Pbft, models.Pbft:
		ret = pbftPrefix
	case *models.Dag, models.Dag:
		ret = dagsPrefix
	case *storage.AddressStats, storage.AddressStats:
		ret = statsPrefix
	case *storage.FinalizationData, storage.FinalizationData:
		ret = finalizationDataPrefix
	case *storage.GenesisHash, storage.GenesisHash:
		ret = genesisHashPrefix
	case *storage.WeekStats, storage.WeekStats:
		ret = weekStatsPrefix
	case *storage.TotalSupply, storage.TotalSupply:
		ret = totalSupplyPrefix
	case *models.InternalTransactionsResponse, models.InternalTransactionsResponse:
		ret = internalTransactionsPrefix
	case *storage.Yield, storage.Yield:
		ret = yieldPrefix
	case *storage.ValidatorsYield, storage.ValidatorsYield:
		ret = validatorsYieldPrefix
	case *storage.MultipliedYield, storage.MultipliedYield:
		ret = multipliedYieldPrefix
	case *storage.PeriodRewards, storage.PeriodRewards:
		ret = PeriodRewardsPrefix
	// hack if we aren't passing original type directly to this function, but passing interface{} from other function
	case *interface{}:
		ret = GetPrefix(*o.(*interface{}))
		// We don't need to add separator in this case, so return from here
		return
	default:
		debug.PrintStack()
		log.WithFields(log.Fields{"type": tt, "value": o}).Fatalf("getPrefix: Unexpected type %T", tt)
	}
	ret += prefixSeparator
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

func (s *Storage) forEach(prefix, start_key []byte, fn func(key, res []byte) (stop bool), navigate func(iter *pebble.Iterator)) {
	iter := s.find(prefix)
	defer iter.Close()
	if len(start_key) == 0 {
		start_key = prefix
	}
	iter.SeekGE(start_key)

	for ; iter.Valid(); navigate(iter) {
		if fn(iter.Key(), iter.Value()) {
			break
		}
	}
}

func (s *Storage) ForEachFromKey(prefix, start_key []byte, fn func(key, res []byte) (stop bool)) {
	s.forEach(prefix, start_key, fn, func(iter *pebble.Iterator) { iter.Next() })
}

func (s *Storage) forEachPrefix(o interface{}, key_prefix string, start *uint64, fn func(key, res []byte) (stop bool), navigate func(iter *pebble.Iterator)) {
	prefix := getPrefixKey(GetPrefix(&o), key_prefix)
	start_key := prefix
	if start != nil {
		start_key = getKey(GetPrefix(&o), key_prefix, *start)
	}
	s.forEach(prefix, start_key, fn, navigate)
}

func (s *Storage) ForEach(o interface{}, key_prefix string, start *uint64, fn func(key, res []byte) (stop bool)) {
	s.forEachPrefix(o, key_prefix, start, fn, func(iter *pebble.Iterator) { iter.Next() })
}

func (s *Storage) ForEachBackwards(o interface{}, key_prefix string, start *uint64, fn func(key, res []byte) (stop bool)) {
	s.forEachPrefix(o, key_prefix, start, fn, func(iter *pebble.Iterator) { iter.Prev() })
}

func (s *Storage) addToDBTest(o interface{}, key1 string, key2 uint64) error {
	return s.addToDB(getKey(GetPrefix(o), key1, key2), o)
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
	err := s.GetFromDB(ptr, []byte(GetPrefix((*storage.TotalSupply)(ptr))))
	if err != nil {
		log.Fatal("GetTotalSupply ", err)
	}
	return ptr
}

func (s *Storage) GetAccounts() storage.Accounts {
	ptr := new(storage.Accounts)
	err := s.GetFromDB(ptr, getPrefixKey(GetPrefix(ptr), ""))
	if err != nil && err != pebble.ErrNotFound {
		log.Fatal("GetAccounts failed: ", err)
	}
	return *ptr
}

func (s *Storage) GetWeekStats(year, week int32) storage.WeekStats {
	ptr := storage.MakeEmptyWeekStats()
	ptr.Key = []byte(getWeekKey(GetPrefix(ptr), year, week))
	err := s.GetFromDB(ptr, ptr.Key)
	if err != nil && err != pebble.ErrNotFound {
		log.WithError(err).Fatal("GetWeekStats failed")
	}
	return *ptr
}

func (s *Storage) GetFinalizationData() *storage.FinalizationData {
	ptr := new(storage.FinalizationData)
	err := s.GetFromDB(ptr, []byte(GetPrefix(ptr)))
	if err != nil && err != pebble.ErrNotFound {
		log.WithError(err).Fatal("GetFinalizationData failed")
	}
	return ptr
}

func (s *Storage) GetAddressStats(addr string) *storage.AddressStats {
	ptr := storage.MakeEmptyAddressStats(addr)
	err := s.GetFromDB(ptr, getKey(GetPrefix(ptr), addr, 0))
	if err != nil && err != pebble.ErrNotFound {
		log.Fatal("GetAddressStats ", err)
	}
	return ptr
}

func (s *Storage) GenesisHashExist() bool {
	ptr := new(storage.GenesisHash)
	err := s.GetFromDB(ptr, []byte(GetPrefix(ptr)))
	return err == nil
}

func (s *Storage) GetGenesisHash() storage.GenesisHash {
	ptr := new(storage.GenesisHash)
	err := s.GetFromDB(ptr, []byte(GetPrefix(ptr)))
	if err != nil {
		log.WithError(err).Fatal("GetGenesisHash failed")
	}
	return *ptr
}

func (s *Storage) GetInternalTransactions(hash string) models.InternalTransactionsResponse {
	ptr := new(models.InternalTransactionsResponse)
	err := s.GetFromDB(ptr, getPrefixKey(GetPrefix(ptr), hash))
	if err != nil && err != pebble.ErrNotFound {
		log.WithError(err).Fatal("GetInternalTransactions failed")
	}
	return *ptr
}

func (s *Storage) GetTransactionLogs(hash string) models.TransactionLogsResponse {
	ptr := new(models.TransactionLogsResponse)
	err := s.GetFromDB(ptr, getPrefixKey(GetPrefix(ptr), hash))
	for i, eventLog := range ptr.Data {
		name, params, err := events.DecodeEventDynamic(eventLog)
		if err != nil {
			log.WithError(err).WithField("name", name).WithField("params", params).Error(err)
		}
		eventLog.Name = name
		eventLog.Params = params
		ptr.Data[i] = eventLog
	}
	if err != nil && err != pebble.ErrNotFound {
		log.WithError(err).Fatal("GetTransactionLogs failed")
	}
	return *ptr
}

func (s *Storage) GetValidatorYield(validator string, block uint64) (res storage.Yield) {
	err := s.GetFromDB(&res, getKey(GetPrefix(&res), validator, block))
	if err != nil && err != pebble.ErrNotFound {
		log.WithError(err).Fatal("GetValidatorYield failed")
	}
	return
}

func (s *Storage) GetTotalYield(block uint64) (res storage.Yield) {
	// total yield is stored under empty address
	return s.GetValidatorYield("", block)
}

func (s *Storage) GetTransactionByHash(hash string) (res models.Transaction) {
	err := s.GetFromDB(&res, getPrefixKey(GetPrefix(&res), strings.ToLower(hash)))
	if err != nil && err != pebble.ErrNotFound {
		log.WithError(err).Fatal("GetTransactionByHash failed")
	}
	return
}

func (s *Storage) GetFromDB(o interface{}, key []byte) error {
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
	return nil
}
