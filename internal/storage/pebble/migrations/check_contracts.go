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

type CheckContracts struct {
	client           common.Client
	contracts        []models.Address
	genesisTimestamp models.Timestamp
	mutex            sync.Mutex
}

func (m *CheckContracts) GetId() string {
	return "check_contracts"
}

func (m *CheckContracts) Init(client common.Client) {
	block, err := client.GetBlockByNumber(0)
	if err != nil {
		log.WithField("error", err).Fatal("Error getting genesis block")
	}
	m.genesisTimestamp = block.Timestamp
	log.WithField("genesis_timestamp", m.genesisTimestamp).Info("Genesis timestamp")
	m.client = client
	m.contracts = make([]models.Address, 0)
}

func (m *CheckContracts) contractCreationTimestamp(s *pebble.Storage, address models.Address) (timestamp models.Timestamp) {
	trx := models.Transaction{}
	first_assigned := false
	s.ForEach(&trx, address, nil, storage.Forward, func(_, res []byte) (stop bool) {
		err := rlp.DecodeBytes(res, &trx)
		if err != nil {
			log.WithFields(log.Fields{"type": storage.GetTypeName[models.Transaction](), "error": err}).Fatal("Error decoding data from db")
		}
		// set first trx timestamp as contract creation timestamp if creation trx wasn't found
		if !first_assigned {
			first_assigned = true
			timestamp = trx.Timestamp
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

func (m *CheckContracts) saveContracts(addresses []models.Address) {
	res, err := m.client.FilterContracts(addresses)
	if err != nil {
		log.WithField("error", err).Fatal("Error filtering contracts")
	}

	m.mutex.Lock()
	m.contracts = append(m.contracts, res...)
	m.mutex.Unlock()
}

func (m *CheckContracts) Apply(s *pebble.Storage) error {
	stats := storage.AddressStats{}
	total_addresses := 0

	tp := common.MakeThreadPool()
	addresses := make([]models.Address, 0, 100)
	s.ForEach(&stats, "", nil, storage.Forward, func(key []byte, res []byte) (stop bool) {
		err := rlp.DecodeBytes(res, &stats)
		if err != nil {
			err_new := rlp.DecodeBytes(res, &stats)
			if err_new != nil {
				log.WithField("error", err).Fatal("Error parsing address stats")
			}
			return true
		}
		if stats.ContractRegisteredTimestamp != nil {
			return false
		}
		addresses = append(addresses, stats.Address)
		if len(addresses) == 100 {
			// make addresses copy and save only contract addresses
			addressesCopy := make([]models.Address, len(addresses))
			copy(addressesCopy, addresses)
			tp.Go(func() {
				m.saveContracts(addressesCopy)
			})
			addresses = make([]models.Address, 0, 100)
		}
		total_addresses++
		return false
	})
	tp.Wait()

	log.WithFields(log.Fields{
		"count":     len(m.contracts),
		"contracts": m.contracts,
	}).Info("Found contracts without registered timestamp")

	batch := s.NewBatch()
	processed := 0
	for _, contract := range m.contracts {
		timestamp := m.contractCreationTimestamp(s, contract)
		if timestamp == 0 {
			log.WithField("contract", contract).Info("No creation timestamp found, using genesis timestamp")
			timestamp = m.genesisTimestamp
		}
		stats := s.GetAddressStats(contract)
		stats.ContractRegisteredTimestamp = &timestamp
		batch.Add(stats, contract, 0)
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
