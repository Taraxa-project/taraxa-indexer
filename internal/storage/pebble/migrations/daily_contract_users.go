package migration

import (
	"time"

	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage/pebble"
	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/ethereum/go-ethereum/rlp"
	log "github.com/sirupsen/logrus"
)

type DailyContractUsers struct {
	client common.Client
}

func (m *DailyContractUsers) GetId() string {
	return "daily_contract_users"
}

func (m *DailyContractUsers) Init(client common.Client) {
	m.client = client
}

func (m *DailyContractUsers) Apply(s *pebble.Storage) error {
	log.Info("DailyContractUsers: Starting migration to backfill contract users data for the last month")

	// Calculate the date range for the last month
	now := time.Now()
	startTimestamp := common.DayStart(uint64(now.AddDate(0, -1, 0).Unix())) // 1 month ago

	log.WithFields(log.Fields{
		"start_timestamp": startTimestamp,
		"start_date":      time.Unix(int64(startTimestamp), 0).Format("2006-01-02"),
	}).Info("DailyContractUsers: Processing date range")

	// Step 1: Collect all contract addresses
	log.Info("DailyContractUsers: Collecting contract addresses")
	contractAddresses := make([]string, 0)
	stats := storage.AddressStats{}

	s.ForEach(&stats, "", nil, storage.Forward, func(key []byte, res []byte) (stop bool) {
		err := rlp.DecodeBytes(res, &stats)
		if err != nil {
			log.WithFields(log.Fields{"error": err}).Error("DailyContractUsers: Error decoding address stats")
			return false
		}

		// Only include addresses that are contracts
		if stats.IsContract() {
			contractAddresses = append(contractAddresses, stats.Address)
		}

		return false
	})

	log.WithFields(log.Fields{
		"count": len(contractAddresses),
	}).Info("DailyContractUsers: Found contracts")

	// Step 2: Process transactions for each contract
	dailyContractUsers := make(map[uint64]map[string]*storage.DailyContractUsers)
	totalProcessed := 0
	contractInteractions := 0

	for contractIndex, contractAddress := range contractAddresses {
		if contractIndex%100 == 0 && contractIndex > 0 {
			log.WithFields(log.Fields{
				"processed_contracts":   contractIndex,
				"total_contracts":       len(contractAddresses),
				"transactions":          totalProcessed,
				"contract_interactions": contractInteractions,
			}).Info("DailyContractUsers: Progress")
		}

		// Process all transactions for this contract starting from newest
		trx := models.Transaction{}
		s.ForEach(&trx, contractAddress, nil, storage.Backward, func(_, res []byte) (stop bool) {
			err := rlp.DecodeBytes(res, &trx)
			if err != nil {
				log.WithFields(log.Fields{
					"error":    err,
					"contract": contractAddress,
				}).Error("DailyContractUsers: Error decoding transaction")
				return false
			}

			totalProcessed++

			// If too old, stop processing this contract
			if trx.Timestamp < startTimestamp {
				return true
			}

			// Only track contract calls to the contract
			if trx.Type == models.ContractCall && trx.To == contractAddress {
				contractInteractions++

				// Get the day start timestamp for this transaction
				dayStart := common.DayStart(trx.Timestamp)

				// Initialize the nested map if needed
				if dailyContractUsers[dayStart] == nil {
					dailyContractUsers[dayStart] = make(map[string]*storage.DailyContractUsers)
				}

				// Initialize the contract users set if needed
				if dailyContractUsers[dayStart][contractAddress] == nil {
					dailyContractUsers[dayStart][contractAddress] = storage.MakeDailyContractUsers()
				}

				// Add the user to the set
				dailyContractUsers[dayStart][contractAddress].Add(trx.From)
			}

			return false
		})
	}

	log.WithFields(log.Fields{
		"processed_contracts":    len(contractAddresses),
		"processed_transactions": totalProcessed,
		"contract_interactions":  contractInteractions,
		"days_with_activity":     len(dailyContractUsers),
	}).Info("DailyContractUsers: Finished processing transactions")

	// Step 3: Save the aggregated data to storage
	batch := s.NewBatch()
	savedRecords := 0

	for dayStart, contractsMap := range dailyContractUsers {
		for contractAddress, users := range contractsMap {
			batch.AddDailyContractUsers(contractAddress, dayStart, users)
			savedRecords++
		}
	}

	batch.CommitBatch()

	log.WithFields(log.Fields{
		"saved_records": savedRecords,
	}).Info("DailyContractUsers: Migration completed successfully")

	return nil
}
