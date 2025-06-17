package storage

import (
	"fmt"
	"sync/atomic"

	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/ethereum/go-ethereum/rlp"
	log "github.com/sirupsen/logrus"
)

type MonthlyActiveAddresses struct {
	Count uint64
}

var (
	queried *uint64
)

func GetMonthlyActiveAddresses(s Storage, from_date, to_date uint64) (c uint64, err error) {
	count := s.GetMonthlyActiveAddresses(to_date)
	if count == nil {
		if queried == nil {
			queried = &to_date
			go func() {
				count := countMonthlyActiveAddresses(s, from_date, to_date)
				batch := s.NewBatch()
				batch.Add(&MonthlyActiveAddresses{Count: count}, "", to_date)
				batch.CommitBatch()
				log.WithFields(log.Fields{
					"count":     count,
					"from_date": from_date,
					"to_date":   to_date,
				}).Info("MonthlyActiveAddresses finished")
				queried = nil
			}()
		} else {
			if *queried != to_date {
				return 0, fmt.Errorf("stats are being calculated for another history period, please try again later")
			}
		}
		return 0, fmt.Errorf("stats are being calculated, please try again later")
	}
	return *count, nil
}

func countMonthlyActiveAddresses(s Storage, from_date, to_date uint64) uint64 {
	count := atomic.Uint64{}

	s.ForEach(AddressStats{}, "", nil, Forward, func(key []byte, res []byte) (stop bool) {
		stats := AddressStats{}

		err := rlp.DecodeBytes(res, &stats)
		if err != nil {
			log.WithError(err).Fatal("Error decoding data from db")
			return
		}

		// skip accounts with last transaction timestamp before from_date
		if stats.LastTransactionTimestamp == nil || *stats.LastTransactionTimestamp < from_date {
			return
		}
		// skip contracts
		if stats.ContractRegisteredTimestamp != nil {
			return
		}

		if WasAccountActive(s, stats.Address, from_date, to_date) {
			count.Add(1)
		}
		return false
	})

	return count.Load()
}

func WasAccountActive(s Storage, address models.AddressParam, from_date, to_date uint64) (found bool) {
	trx := models.Transaction{}
	// check for transaction from address in the interval. start from most recent
	s.ForEach(&trx, address, nil, Backward, func(key []byte, res []byte) (stop bool) {
		err := rlp.DecodeBytes(res, &trx)
		if err != nil {
			log.WithError(err).Fatal("Error decoding data from db")
			return false
		}

		if trx.Timestamp < from_date {
			return true
		}

		// we should only count transactions from the account
		if trx.From != address {
			return false
		}

		if trx.Timestamp > to_date {
			return false
		}

		found = true
		return true
	})

	return
}

func ReceivedTransactionsCount(s Storage, address models.AddressParam, from_date, to_date uint64) (count models.Counter) {
	trx := models.Transaction{}

	// check for transaction from address in the interval. start from most recent
	s.ForEach(&trx, address, nil, Forward, func(key []byte, res []byte) (stop bool) {
		err := rlp.DecodeBytes(res, &trx)
		if err != nil {
			log.WithError(err).Fatal("Error decoding data from db")
			return false
		}

		if trx.To != address {
			return false
		}

		if trx.Timestamp > to_date {
			return false
		}

		if trx.Timestamp < from_date {
			return true
		}

		count++
		return false
	})
	return count
}
