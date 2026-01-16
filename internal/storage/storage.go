package storage

import (
	"fmt"

	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/ethereum/go-ethereum/rlp"
	log "github.com/sirupsen/logrus"
)

type Direction int

const (
	Forward Direction = iota
	Backward
)

type Storage interface {
	Clean() error
	Close() error
	ForEach(o any, key_prefix string, start *uint64, direction Direction, fn func(key, res []byte) (stop bool))
	ForEachFromKey(prefix, start_key []byte, direction Direction, fn func(key, res []byte) (stop bool))
	NewBatch() Batch
	GetTotalSupply() *TotalSupply
	GetAccounts() Accounts
	GetWeekStats(year, week int32) WeekStats
	GetFinalizationData() *common.FinalizationData
	GetDayStats(timestamp uint64) DayStatsWithTimestamp
	GetAddressStats(addr string) *AddressStats
	GenesisHashExist() bool
	GetGenesisHash() GenesisHash
	GetTransactionByHash(hash string) models.Transaction
	GetInternalTransactions(hash string) models.InternalTransactionsResponse
	GetTransactionLogs(hash string) models.TransactionLogsResponse
	GetValidatorYield(validator string, block uint64) (res Yield)
	GetTotalYield(block uint64) (res Yield)
	GetMonthlyActiveAddresses(to_date uint64) *uint64
	GetDailyContractUsers(address string, timestamp uint64) DailyContractUsersList
	GetYieldIntervals(from_block, to_block uint64) []uint64
	GetYieldInterval(block *uint64) (uint64, uint64)
	GetLatestYieldSaving() *YieldSaving
	GetLambda() *uint64
}

func GetTotal[T Paginated](s Storage, address string) (r uint64) {
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

func GetObjectsPage[T Paginated](s Storage, address string, from, count uint64) (ret []T, pagination *models.PaginatedResponse) {
	var o T

	pagination = new(models.PaginatedResponse)
	pagination.Start = from
	pagination.Total = GetTotal[T](s, address)
	end := from + count
	pagination.HasNext = (end < pagination.Total)
	if end > pagination.Total {
		end = pagination.Total
	}
	pagination.End = end

	ret = make([]T, 0, count)
	start := pagination.Total - from
	s.ForEach(&o, address, &start, Backward, func(_, res []byte) (stop bool) {
		err := rlp.DecodeBytes(res, &o)
		if err != nil {
			log.WithFields(log.Fields{"type": GetTypeName[T](), "error": err}).Fatal("Error decoding data from db")
		}
		ret = append(ret, o)
		if uint64(len(ret)) == count {
			return true
		}
		return
	})
	return
}

func GetHoldersPage(s Storage, from, count uint64) (ret []models.Account, pagination *models.PaginatedResponse) {
	holders := s.GetAccounts()
	pagination = new(models.PaginatedResponse)
	pagination.Start = from
	pagination.Total = uint64(holders.Total)
	end := from + count
	pagination.HasNext = (end < uint64(len(holders.Accounts)))
	if end > pagination.Total {
		end = pagination.Total
	}
	pagination.End = end

	ret = make([]models.Account, 0, count)
	for i := from; i < end; i++ {
		ret = append(ret, holders.Accounts[i].ToModel())
	}
	return
}

func ProcessIntervalData[T Yields](s Storage, start uint64, fn func([]byte, T) (stop bool)) {
	var o T
	s.ForEach(&o, "", &start, Forward, func(key, res []byte) bool {
		err := rlp.DecodeBytes(res, &o)
		if err != nil {
			log.WithFields(log.Fields{"type": GetTypeName[T](), "error": err}).Fatal("Error decoding data from db")
		}
		return fn(key, o)
	})
}

func GetUIntKey(key uint64) string {
	return fmt.Sprintf("%020d", key)
}
