package storage

import (
	"fmt"

	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/ethereum/go-ethereum/rlp"
	log "github.com/sirupsen/logrus"
)

type Storage interface {
	Clean() error
	Close() error
	ForEach(o interface{}, key_prefix string, start uint64, fn func(key, res []byte) (stop bool))
	ForEachBackwards(o interface{}, key_prefix string, start uint64, fn func(key, res []byte) (stop bool))
	NewBatch() Batch
	GetTotalSupply() *TotalSupply
	GetAccounts() []Account
	GetWeekStats(year, week int32) WeekStats
	GetFinalizationData() *FinalizationData
	GetAddressStats(addr string) *AddressStats
	GenesisHashExist() bool
	GetGenesisHash() GenesisHash
	GetTransactionByHash(hash string) models.Transaction
	GetInternalTransactions(hash string) models.InternalTransactionsResponse
	GetTransactionLogs(hash string) models.TransactionLogsResponse
	GetValidatorYield(validator string, block uint64) (res Yield)
	GetTotalYield(block uint64) (res Yield)
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
	s.ForEachBackwards(&o, address, pagination.Total-from, func(_, res []byte) (stop bool) {
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
	pagination.Total = uint64(len(holders))
	end := from + count
	pagination.HasNext = (end < pagination.Total)
	if end > pagination.Total {
		end = pagination.Total
	}
	pagination.End = end

	ret = make([]models.Account, 0, count)
	for i := from; i < end; i++ {
		ret = append(ret, holders[i].ToModel())
	}
	return
}

func ProcessIntervalData[T Yields](s Storage, start uint64, fn func(string, T) (stop bool)) {
	var o T
	s.ForEach(&o, "", start, func(key, res []byte) (stop bool) {
		err := rlp.DecodeBytes(res, &o)
		if err != nil {
			log.WithFields(log.Fields{"type": GetTypeName[T](), "error": err}).Fatal("Error decoding data from db")
		}
		return fn(string(key), o)
	})
}

func GetUIntKey(key uint64) string {
	return fmt.Sprintf("%020d", key)
}
