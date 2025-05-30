package api

import (
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/ethereum/go-ethereum/rlp"
	log "github.com/sirupsen/logrus"
)

func wasAccountActive(s storage.Storage, address models.AddressParam, from_date, to_date uint64) (found bool) {
	trx := models.Transaction{}
	// check for transaction from address in the interval. start from most recent
	s.ForEachBackwards(&trx, "", nil, func(key []byte, res []byte) (stop bool) {
		err := rlp.DecodeBytes(res, &trx)
		if err != nil {
			log.WithError(err).Fatal("Error decoding data from db")
			return false
		}

		// we should only count transactions from the account
		if trx.From != address {
			return false
		}

		if trx.Timestamp > to_date {
			return false
		}

		if trx.Timestamp < from_date {
			return true
		}

		found = true
		return true
	})

	return
}

func receivedTransactionsCount(s storage.Storage, address models.AddressParam, from_date, to_date uint64) (count models.Counter) {
	trx := models.Transaction{}

	// check for transaction from address in the interval. start from most recent
	s.ForEach(&trx, address, nil, func(key []byte, res []byte) (stop bool) {
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
