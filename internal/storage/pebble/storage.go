package pebble

import (
	"bytes"
	"fmt"
	"io"
	"math/big"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
	"sync"

	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/events"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/cockroachdb/pebble"
	"github.com/cockroachdb/pebble/vfs"
	"github.com/ethereum/go-ethereum/rlp"
	log "github.com/sirupsen/logrus"
)

const PrefixSeparator = "|"
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
const RewardsStatsPrefix = "rs"
const dayStatsPrefix = "ds"
const monthlyActiveAddressesPrefix = "ma"
const dailyContractUsersPrefix = "cu"
const yieldSavingPrefix = "ys"
const lambdaPrefix = "l"

var ErrNotFound = pebble.ErrNotFound

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

func GetPrefix(o any) (ret string) {
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
	case *common.FinalizationData, common.FinalizationData:
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
	case *storage.RewardsStats, storage.RewardsStats:
		ret = RewardsStatsPrefix
	case *storage.TrxGasStats, storage.TrxGasStats:
		ret = dayStatsPrefix
	case *storage.MonthlyActiveAddresses, storage.MonthlyActiveAddresses:
		ret = monthlyActiveAddressesPrefix
	case *storage.DailyContractUsersList, storage.DailyContractUsersList:
		ret = dailyContractUsersPrefix
	case *storage.YieldSaving, storage.YieldSaving:
		ret = yieldSavingPrefix
	case *storage.Lambda, storage.Lambda:
		ret = lambdaPrefix
	// hack if we aren't passing original type directly to this function, but passing any from other function
	case *any:
		ret = GetPrefix(*o.(*any))
		// We don't need to add separator in this case, so return from here
		return
	default:
		debug.PrintStack()
		log.WithFields(log.Fields{"type": tt, "value": o}).Fatalf("getPrefix: Unexpected type %T", tt)
	}
	ret += PrefixSeparator
	return
}

func getKey(prefix, key1 string, key2 uint64) []byte {
	key1 = strings.ToLower(key1)
	return fmt.Appendf(nil, "%s%s%020d", prefix, key1, key2)
}

func GetPrefixKey(prefix, author string) []byte {
	author = strings.ToLower(author)
	return fmt.Appendf(nil, "%s%s", prefix, author)
}

func getWeekKey(prefix string, year, week int32) []byte {
	return fmt.Appendf(nil, "%s%d%02d", prefix, year, week)
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

	iter, err := s.db.NewIter(&prefixIterOptions)
	if err != nil {
		log.WithError(err).Fatal("NewIter failed")
	}
	return iter
}

func navigate(direction storage.Direction) func(iter *pebble.Iterator) {
	if direction == storage.Backward {
		return func(iter *pebble.Iterator) {
			iter.Prev()
		}
	}

	return func(iter *pebble.Iterator) {
		iter.Next()
	}
}

func (s *Storage) forEachKey(prefix, start_key []byte, fn func(key, res []byte) (stop bool), navigate func(iter *pebble.Iterator)) {
	iter := s.find(prefix)
	defer func() { _ = iter.Close() }()
	if len(start_key) == 0 {
		start_key = prefix
	}
	iter.SeekGE(start_key)

	s.forEach(iter, fn, navigate)
}

func (s *Storage) forEach(iter *pebble.Iterator, fn func(key, res []byte) (stop bool), navigate func(iter *pebble.Iterator)) {
	for ; iter.Valid(); navigate(iter) {
		if fn(iter.Key(), iter.Value()) {
			break
		}
	}
}

func (s *Storage) ForEachFromKey(prefix, start_key []byte, direction storage.Direction, fn func(key, res []byte) (stop bool)) {
	start_key = bytes.Join([][]byte{prefix, start_key}, []byte(""))
	s.forEachKey(prefix, start_key, fn, navigate(direction))
}

func (s *Storage) ForEach(o any, address string, start *uint64, direction storage.Direction, fn func(key, res []byte) (stop bool)) {
	prefix := GetPrefixKey(GetPrefix(&o), address)

	iter := s.find(prefix)
	defer func() { _ = iter.Close() }()

	if start == nil {
		if direction == storage.Backward {
			iter.Last()
		} else {
			iter.First()
		}
	} else {
		start_key := getKey(GetPrefix(&o), address, *start)
		iter.SeekGE(start_key)
	}

	s.forEach(iter, fn, navigate(direction))
}

func (s *Storage) addToDBTest(o any, key1 string, key2 uint64) error {
	return s.addToDB(getKey(GetPrefix(o), key1, key2), o)
}

func (s *Storage) addToDB(key []byte, o any) error {
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
	err := s.GetFromDB(ptr, GetPrefixKey(GetPrefix(ptr), ""))
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

func (s *Storage) GetFinalizationData() *common.FinalizationData {
	ptr := new(common.FinalizationData)
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
	err := s.GetFromDB(ptr, GetPrefixKey(GetPrefix(ptr), hash))
	if err != nil && err != pebble.ErrNotFound {
		log.WithError(err).Fatal("GetInternalTransactions failed")
	}
	return *ptr
}

func (s *Storage) GetTransactionLogs(hash string) models.TransactionLogsResponse {
	ptr := new(models.TransactionLogsResponse)
	err := s.GetFromDB(ptr, GetPrefixKey(GetPrefix(ptr), hash))
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
	err := s.GetFromDB(&res, GetPrefixKey(GetPrefix(&res), strings.ToLower(hash)))
	if err != nil && err != pebble.ErrNotFound {
		log.WithError(err).Fatal("GetTransactionByHash failed")
	}
	return
}

func (s *Storage) GetDayStats(timestamp uint64) storage.DayStatsWithTimestamp {
	ret := storage.MakeDayStatsWithTimestamp(timestamp)
	err := s.GetFromDB(&ret.TrxGasStats, getKey(GetPrefix(&ret.TrxGasStats), "", timestamp))
	if err != nil && err != pebble.ErrNotFound {
		log.WithError(err).Fatal("GetDayStats failed")
	}
	return *ret
}

func (s *Storage) GetMonthlyActiveAddresses(to_date uint64) *uint64 {
	res := storage.MonthlyActiveAddresses{}
	err := s.GetFromDB(&res, getKey(GetPrefix(&res), "", to_date))
	if err == pebble.ErrNotFound {
		return nil
	}
	if err != nil {
		log.WithError(err).Fatal("GetMonthlyActiveAddresses failed")
	}
	return &res.Count
}

func (s *Storage) GetDailyContractUsers(address string, timestamp uint64) storage.DailyContractUsersList {
	ret := storage.MakeDailyContractUsersList()
	dayStart := common.DayStart(timestamp)
	err := s.GetFromDB(&ret, getKey(GetPrefix(&ret), address, dayStart))
	if err != nil && err != pebble.ErrNotFound {
		log.WithError(err).Fatal("GetDailyContractUsers failed")
	}
	return ret
}

func (s *Storage) GetYieldInterval(block uint64) (uint64, uint64) {
	blocks := s.GetYieldIntervals(block, block)
	if len(blocks) == 0 {
		return 0, 0
	}
	return blocks[0], blocks[len(blocks)-1]
}

func (s *Storage) GetYieldIntervals(from_block, to_block uint64) []uint64 {
	intervals := make([]uint64, 0)
	s.ForEachFromKey([]byte(GetPrefix(storage.Yield{})), []byte(storage.FormatIntToKey(from_block)), storage.Forward, func(key, res []byte) (stop bool) {
		keyParts := strings.Split(string(key), "|")
		if len(keyParts) < 2 {
			return false
		}
		// Parse as base 10 to handle leading zeros correctly
		curr_block, err := strconv.ParseUint(keyParts[1], 10, 64)
		if err != nil {
			return false
		}
		// Only include blocks within the range [from_block, to_block]
		if curr_block >= from_block && curr_block <= to_block {
			intervals = append(intervals, curr_block)
		}
		// Stop if we've gone past the to_block
		return curr_block > to_block
	})
	return intervals
}

func (s *Storage) GetLatestYieldSaving() (res *storage.YieldSaving) {
	itr := s.find([]byte(GetPrefix(res)))
	itr.Last()
	if !itr.Valid() {
		return
	}
	key := itr.Key()
	parts := strings.Split(string(key), "|")
	if len(parts) < 2 {
		return
	}
	_, err := strconv.ParseUint(parts[1], 10, 64)
	if err != nil {
		return
	}
	return res
}

func (s *Storage) GetLambda() *uint64 {
	res := storage.Lambda{}
	err := s.GetFromDB(&res, []byte(GetPrefix(res)))
	if err != nil && err != pebble.ErrNotFound {
		log.WithError(err).Fatal("GetLambda failed")
	}
	if err == pebble.ErrNotFound {
		return nil
	}
	return &res.LambdaMs
}

func (s *Storage) GetFromDB(o any, key []byte) error {
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
