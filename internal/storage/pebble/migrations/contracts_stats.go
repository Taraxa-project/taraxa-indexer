package migration

import (
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage/pebble"
	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/ethereum/go-ethereum/rlp"
	log "github.com/sirupsen/logrus"
)

// StatsResponse defines model for StatsResponse.
type OldStatsResponse struct {
	DagsCount                models.Counter         `json:"dagsCount"`
	LastDagTimestamp         *models.NillableUint64 `json:"lastDagTimestamp" rlp:"nil"`
	LastPbftTimestamp        *models.NillableUint64 `json:"lastPbftTimestamp" rlp:"nil"`
	LastTransactionTimestamp *models.NillableUint64 `json:"lastTransactionTimestamp" rlp:"nil"`
	PbftCount                models.Counter         `json:"pbftCount"`
	TransactionsCount        models.Counter         `json:"transactionsCount"`
	ValidatorRegisteredBlock *models.NillableUint64 `json:"validatorRegisteredBlock" rlp:"nil"`
}
type OldAddressStats struct {
	OldStatsResponse
	Address string `json:"address"`
}

func (o *OldAddressStats) toStatsResponse(timestamp models.Timestamp) storage.AddressStats {
	return storage.AddressStats{StatsResponse: models.StatsResponse{
		DagsCount:                   o.DagsCount,
		LastDagTimestamp:            o.LastDagTimestamp,
		LastPbftTimestamp:           o.LastPbftTimestamp,
		LastTransactionTimestamp:    o.LastTransactionTimestamp,
		PbftCount:                   o.PbftCount,
		TransactionsCount:           o.TransactionsCount,
		ValidatorRegisteredBlock:    o.ValidatorRegisteredBlock,
		ContractRegisteredTimestamp: &timestamp,
	}}
}

type ContractStats struct {
}

func (m *ContractStats) GetId() string {
	return "contract_stats"
}

func (m *ContractStats) Init(common.Client) {
}

func (m *ContractStats) contractCreationTimestamp(s *pebble.Storage, address models.Address) (timestamp models.Timestamp) {
	trx := models.Transaction{}
	s.ForEach(&trx, address, nil, func(_, res []byte) (stop bool) {
		err := rlp.DecodeBytes(res, &trx)
		if err != nil {
			log.WithFields(log.Fields{"type": storage.GetTypeName[models.Transaction](), "error": err}).Fatal("Error decoding data from db")
		}

		if trx.Type != models.ContractCreation && trx.Type != models.InternalContractCreation {
			return false
		}
		if trx.To != address {
			return false
		}

		log.WithFields(log.Fields{
			"trx":     trx,
			"address": address,
		}).Info("Found creation transaction")

		timestamp = trx.Timestamp
		return true
	})

	return
}

func (m *ContractStats) Apply(s *pebble.Storage) error {
	old_stats := OldAddressStats{}
	stats := storage.AddressStats{}
	batch := s.NewBatch()
	s.ForEach(&stats, "", nil, func(key []byte, res []byte) (stop bool) {
		err := rlp.DecodeBytes(res, &old_stats)
		if err != nil {
			err_new := rlp.DecodeBytes(res, &stats)
			if err_new != nil {
				log.WithField("error", err).Fatal("Error parsing address stats")
			}
			return true
		}
		stats = old_stats.toStatsResponse(m.contractCreationTimestamp(s, old_stats.Address))
		err = batch.AddWithKey(&stats, key)
		if err != nil {
			log.WithField("error", err).Fatal("Error adding stats to batch")
		}
		return false
	})

	batch.CommitBatch()
	return nil
}
