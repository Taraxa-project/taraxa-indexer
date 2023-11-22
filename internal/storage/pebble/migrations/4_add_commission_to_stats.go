package migration

import (
	"fmt"
	"strings"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage/pebble"
	"github.com/Taraxa-project/taraxa-indexer/models"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	log "github.com/sirupsen/logrus"
)

type AddCommission struct {
	id            string
	blockchain_ws string
}

type ValidatorCommission struct {
	Validator  string
	Commission uint64
}

type OldStatsResponseCom struct {
	DagsCount                models.Counter        `json:"dagsCount"`
	LastDagTimestamp         *models.NilableUint64 `json:"lastDagTimestamp" rlp:"nil"`
	LastPbftTimestamp        *models.NilableUint64 `json:"lastPbftTimestamp" rlp:"nil"`
	LastTransactionTimestamp *models.NilableUint64 `json:"lastTransactionTimestamp" rlp:"nil"`
	PbftCount                models.Counter        `json:"pbftCount"`
	TransactionsCount        models.Counter        `json:"transactionsCount"`
	ValidatorRegisteredBlock *models.NilableUint64 `json:"validatorRegisteredBlock" rlp:"nil"`
}

type OldAddressStatsCom struct {
	OldStatsResponseCom
	Address string `json:"address"`
}

func (m *AddCommission) GetId() string {
	return m.id
}

func (m *AddCommission) migrateStats(s *pebble.Storage) {
	const STATS_BATCH_THRESHOLD = 1000
	batch := s.NewBatch()
	var last_key []byte

	for {
		var o OldAddressStatsCom
		count := 0
		s.ForEachFromKey([]byte(pebble.GetPrefix(storage.AddressStats{})), last_key, func(key []byte, res []byte) bool {
			err := rlp.DecodeBytes(res, &o)
			if err != nil {
				if err.Error() == "rlp: input list has too many elements for migration.OldStatsResponseCom, decoding into (migration.OldAddressStatsCom).OldStatsResponseCom" {
					// Check if it's really already migrated
					var o storage.AddressStats
					err := rlp.DecodeBytes(res, &o)
					if err != nil {
						log.WithFields(log.Fields{"migration": m.id, "error": err}).Fatal("Error decoding AddressStats")
					}
					return false
				}
				log.WithFields(log.Fields{"migration": m.id, "error": err}).Fatal("Error decoding OldAddressStatsCom")
			}

			sr := models.StatsResponse{
				DagsCount:                o.DagsCount,
				LastDagTimestamp:         o.LastDagTimestamp,
				LastPbftTimestamp:        o.LastPbftTimestamp,
				LastTransactionTimestamp: o.LastTransactionTimestamp,
				PbftCount:                o.PbftCount,
				TransactionsCount:        o.TransactionsCount,
				ValidatorRegisteredBlock: o.ValidatorRegisteredBlock,
			}
			err = batch.AddToBatchFullKey(&storage.AddressStats{Address: o.Address, StatsResponse: sr}, key)

			if err != nil {
				log.WithFields(log.Fields{"migration": m.id, "error": err}).Fatal("Error adding AddressStats to batch")
			}

			last_key = key
			count++
			return count == STATS_BATCH_THRESHOLD
		})
		batch.CommitBatch()
		batch = s.NewBatch()
		if count < STATS_BATCH_THRESHOLD {
			break
		}
	}

}

func (m *AddCommission) Apply(s *pebble.Storage) error {
	m.migrateStats(s)
	client, err := chain.NewWsClient(m.blockchain_ws)
	if err != nil {
		log.Fatal(err)
	}

	currentHead, err := client.GetLatestPeriod()
	if err != nil {
		log.Fatal(err)
	}

	if currentHead == 0 {
		return nil
	}

	step := uint64(100000)

	batch := s.NewBatch()

	for startBlock := uint64(0); startBlock < currentHead; startBlock += step {
		endBlock := startBlock + step
		validators, err := GetCommissionChangesInBlock(client, startBlock, endBlock)
		if err != nil {
			log.Fatal(err)
		}

		for _, validator := range validators {
			addressStats := s.GetAddressStats(strings.ToLower(validator.Validator))
			if addressStats == nil {
				addressStats = &storage.AddressStats{
					Address: strings.ToLower(validator.Validator),
					StatsResponse: models.StatsResponse{
						Commission: &validator.Commission,
					},
				}
			} else {
				addressStats.Commission = &validator.Commission
			}
			batch.AddToBatch(addressStats, addressStats.Address, 0)
		}
	}
	batch.CommitBatch()
	return nil
}

func GetCommissionChangesInBlock(client *chain.WsClient, from, to uint64) ([]ValidatorCommission, error) {
	if from > to {
		return nil, fmt.Errorf("from block %d is greater than to block %d", from, to)
	}

	logs, err := client.GetLogs(from, to, []string{"0x00000000000000000000000000000000000000fe"}, [][]string{{"0xc909daf778d180f43dac53b55d0de934d2f1e0b70412ca274982e4e6e894eb1a"}})
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("Found %d validator registrations in blocks %d-%d", len(logs), from, to)

	var validators []ValidatorCommission
	for _, eLog := range logs {
		event := struct {
			Validator string `json:"validator"`
		}{}

		event.Validator = strings.ToLower(ethcommon.HexToAddress(eLog.Topics[1]).Hex())
		commissionHex := eLog.Data
		hexString := commissionHex[2:]
		validators = append(validators, ValidatorCommission{Validator: event.Validator, Commission: common.ParseUInt(hexString)})
		log.Infof("Found validator %s changed commission in block %d", event.Validator, common.ParseUInt(eLog.BlockNumber))
	}
	return validators, nil
}
