package migration

import (
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage/pebble"
	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/ethereum/go-ethereum/rlp"
	log "github.com/sirupsen/logrus"
)

type LastTrxTimestamp struct {
}

func (m *LastTrxTimestamp) GetId() string {
	return "last_trx_timestamp"
}

func (m *LastTrxTimestamp) Init(client common.Client) {
}

func (m *LastTrxTimestamp) getLastTrxTimestamp(s *pebble.Storage, address models.Address) (timestamp *models.Timestamp) {
	trx := new(models.Transaction)
	s.ForEach(trx, address, nil, storage.Backward, func(_, res []byte) (stop bool) {
		err := rlp.DecodeBytes(res, &trx)
		if err != nil {
			log.WithFields(log.Fields{"type": storage.GetTypeName[models.Transaction](), "error": err}).Fatal("Error decoding data from db")
		}
		if trx.From != address {
			return false
		}
		timestamp = &trx.Timestamp
		return true
	})

	return
}

func (m *LastTrxTimestamp) Apply(s *pebble.Storage) error {

	batch := s.NewBatch()
	tp := common.MakeThreadPool()
	s.ForEach(storage.AddressStats{}, "", nil, storage.Forward, func(key []byte, res []byte) (stop bool) {
		tp.Go(func() {
			stats := storage.AddressStats{}
			err := rlp.DecodeBytes(res, &stats)
			if err != nil {
				log.WithField("error", err).Fatal("Error parsing address stats")
			}
			timestamp := m.getLastTrxTimestamp(s, stats.Address)
			stats.LastTransactionTimestamp = timestamp
			batch.Add(&stats, stats.Address, 0)

			if stats.LastTransactionTimestamp == nil || *stats.LastTransactionTimestamp != *timestamp {
				ts := uint64(0)
				if timestamp != nil {
					ts = *timestamp
				}
				log.WithFields(log.Fields{"address": stats.Address, "timestamp": ts}).Info("Updated last transaction timestamp")
			}
		})
		return false
	})
	tp.Wait()

	batch.CommitBatch()
	return nil
}
