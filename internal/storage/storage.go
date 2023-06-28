package storage

import (
	"reflect"

	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/ethereum/go-ethereum/rlp"
	log "github.com/sirupsen/logrus"
)

type Storage interface {
	Clean() error
	Close() error
	ForEach(o interface{}, key_prefix string, start uint64, fn func([]byte) (stop bool))
	NewBatch() Batch
	GetTotalSupply() *TotalSupply
	GetWeekStats(year, week int32) WeekStats
	GetFinalizationData() *FinalizationData
	GetAddressStats(addr string) *AddressStats
	GetAccount(addr string) *Account
	GetAccounts() *map[string]*Account
	GenesisHashExist() bool
	GetGenesisHash() GenesisHash
	GetInternalTransactions(hash string) models.InternalTransactionsResponse
	GetTransactionLogs(hash string) models.TransactionLogsResponse
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
	s.ForEach(&o, address, 0, func(res []byte) (stop bool) {
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

func GetAccountsPage[T Paginated](s Storage, from, count uint64) (ret []T, pagination *models.PaginatedResponse) {

	pagination = new(models.PaginatedResponse)
	pagination.Start = from
	accountMap := s.GetAccounts()
	pagination.Total = uint64(len(*accountMap))
	end := from + count
	pagination.HasNext = (end < pagination.Total)
	if end > pagination.Total {
		end = pagination.Total
	}
	pagination.End = end

	ret = make([]T, 0, count)
	for _, account := range *accountMap {
		if count < from {
			count++
			continue
		}

		if count >= end {
			break
		}

		account.Mutex.RLock()
		acc := account.ToModel()
		t := reflect.TypeOf(acc)
		value := reflect.New(t).Elem()

		// Copy values from acc to the new instance
		value.Set(reflect.ValueOf(acc))

		// Append the new instance to the ret slice
		ret = append(ret, value.Interface().(T))
		account.Mutex.RUnlock()

		count++
	}
	return
}
