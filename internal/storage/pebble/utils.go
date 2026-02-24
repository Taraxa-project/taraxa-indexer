package pebble

import (
	"bytes"

	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/ethereum/go-ethereum/rlp"
	log "github.com/sirupsen/logrus"
)

type prefixStat struct {
	Name  string
	Count uint64
	Bytes uint64
}

func PrintDbStats(s *Storage) {
	// Build mapping from on-disk prefix (e.g. "t|") to stats bucket.
	prefixBuckets := map[string]*prefixStat{
		GetPrefix(storage.Accounts{}):               {Name: "accounts"},
		GetPrefix(models.TransactionLogsResponse{}): {Name: "transaction_logs"},
		GetPrefix(models.Transaction{}):             {Name: "transactions"},
		GetPrefix(models.Pbft{}):                    {Name: "pbft"},
		GetPrefix(models.Dag{}):                     {Name: "dags"},
		GetPrefix(storage.AddressStats{}):           {Name: "address_stats"},
		GetPrefix(common.FinalizationData{}):        {Name: "finalization_data"},
		GetPrefix(new(storage.GenesisHash)):         {Name: "genesis_hash"},
		GetPrefix(storage.WeekStats{}):              {Name: "week_stats"},
		GetPrefix(new(storage.TotalSupply)):         {Name: "total_supply"},
		GetPrefix(models.InternalTransactionsResponse{}): {
			Name: "internal_transactions",
		},
		GetPrefix(storage.Yield{}):           {Name: "yield"},
		GetPrefix(storage.ValidatorsYield{}): {Name: "validators_yield"},
		GetPrefix(storage.MultipliedYield{}): {Name: "multiplied_yield"},
		GetPrefix(storage.RewardsStats{}):    {Name: "rewards_stats"},
		GetPrefix(storage.TrxGasStats{}):     {Name: "day_stats"},
		GetPrefix(storage.MonthlyActiveAddresses{}): {
			Name: "monthly_active_addresses",
		},
		GetPrefix(storage.DailyContractUsersList{}): {
			Name: "daily_contract_users",
		},
		GetPrefix(storage.YieldSaving{}): {Name: "yield_saving"},
		GetPrefix(storage.Lambda{}):      {Name: "lambda"},
	}

	var totalCount uint64
	var totalBytes uint64

	s.ForEachKeyAll(func(key, value []byte) (stop bool) {
		sepIdx := bytes.IndexByte(key, '|')
		if sepIdx == -1 {
			return false
		}

		if bucket, ok := prefixBuckets[string(key[:sepIdx+1])]; ok {
			bucket.Count++
			bucket.Bytes += uint64(len(value))
		}

		totalCount++
		totalBytes += uint64(len(value))

		return false
	})

	for prefixKey, bucket := range prefixBuckets {
		log.WithFields(log.Fields{
			"db_prefix": prefixKey,
			"name":      bucket.Name,
			"count":     bucket.Count,
			"bytes":     bucket.Bytes,
		}).Info("Prefix statistics")
	}

	log.WithFields(log.Fields{
		"total_count": totalCount,
		"total_bytes": totalBytes,
	}).Info("Overall prefix statistics")

}

func PrintAddressStats(s *Storage) {
	s.ForEach(storage.AddressStats{}, "", nil, storage.Forward, func(key, res []byte) (stop bool) {
		var stat storage.AddressStats
		err := rlp.DecodeBytes(res, &stat)
		if err != nil {
			return true
		}
		log.WithFields(log.Fields{"address": stat.Address, "balance": stat.Balance.String(), "transactions": stat.TransactionsCount, "pbfts": stat.PbftCount, "dags": stat.DagsCount}).Info("Address stats")
		return false
	})
}
