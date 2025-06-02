package migration

import (
	"sync"

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

func (o *OldAddressStats) toStatsResponse(timestamp *models.NillableUint64) storage.AddressStats {
	return storage.AddressStats{StatsResponse: models.StatsResponse{
		DagsCount:                   o.DagsCount,
		LastDagTimestamp:            o.LastDagTimestamp,
		LastPbftTimestamp:           o.LastPbftTimestamp,
		LastTransactionTimestamp:    o.LastTransactionTimestamp,
		PbftCount:                   o.PbftCount,
		TransactionsCount:           o.TransactionsCount,
		ValidatorRegisteredBlock:    o.ValidatorRegisteredBlock,
		ContractRegisteredTimestamp: timestamp,
	}}
}

type Contracts struct {
}

type ContractStats struct {
	client           common.Client
	contracts        []models.Address
	genesisTimestamp models.Timestamp
	mutex            sync.Mutex
}

func (m *ContractStats) GetId() string {
	return "contract_stats"
}

func (m *ContractStats) Init(client common.Client) {
	block, err := client.GetBlockByNumber(0)
	if err != nil {
		log.WithField("error", err).Fatal("Error getting genesis block")
	}
	m.genesisTimestamp = block.Timestamp
	log.WithField("genesis_timestamp", m.genesisTimestamp).Info("Genesis timestamp")
	m.client = client
	m.contracts = make([]models.Address, 0)
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

		timestamp = trx.Timestamp
		return true
	})

	return
}

func (m *ContractStats) saveContracts(addresses []models.Address) {
	res, err := m.client.FilterContracts(addresses)
	if err != nil {
		log.WithField("error", err).Fatal("Error filtering contracts")
	}

	m.mutex.Lock()
	m.contracts = append(m.contracts, res...)
	m.mutex.Unlock()
}

func (m *ContractStats) Apply(s *pebble.Storage) error {
	old_stats := OldAddressStats{}
	stats := storage.AddressStats{}
	total_addresses := 0

	tp := common.MakeThreadPool()
	batch := s.NewBatch()
	addresses := make([]models.Address, 0, 100)
	s.ForEach(&stats, "", nil, func(key []byte, res []byte) (stop bool) {
		err := rlp.DecodeBytes(res, &old_stats)
		if err != nil {
			err_new := rlp.DecodeBytes(res, &stats)
			if err_new != nil {
				log.WithField("error", err).Fatal("Error parsing address stats")
			}
			return true
		}
		addresses = append(addresses, old_stats.Address)
		if len(addresses) == 100 {
			// make addresses copy and save only contract addresses
			addressesCopy := make([]models.Address, len(addresses))
			copy(addressesCopy, addresses)
			tp.Go(func() {
				m.saveContracts(addressesCopy)
			})
			addresses = make([]models.Address, 0, 100)
		}
		stats = old_stats.toStatsResponse(nil)
		err = batch.AddWithKey(&stats, key)
		if err != nil {
			log.WithField("error", err).Fatal("Error adding stats to batch")
		}
		total_addresses++
		return false
	})
	tp.Wait()

	batch.CommitBatch()
	batch = s.NewBatch()

	log.WithFields(log.Fields{
		"contracts_found": len(m.contracts),
		"total_addresses": total_addresses,
	}).Info("Found contracts")

	processed := 0
	for _, contract := range m.contracts {
		timestamp := m.contractCreationTimestamp(s, contract)
		if timestamp == 0 {
			log.WithField("contract", contract).Info("No creation timestamp found, using genesis timestamp")
			timestamp = m.genesisTimestamp
			continue
		}
		stats := s.GetAddressStats(contract)
		stats.ContractRegisteredTimestamp = &timestamp
		batch.AddSingleKey(stats, contract)
		processed++
		if processed%100 == 0 {
			log.WithFields(log.Fields{
				"processed": processed,
				"contracts": len(m.contracts),
			}).Info("Processed contracts")
		}
	}

	batch.CommitBatch()
	return nil
}
