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

type AddValidatorRegistrationBlock struct {
	id            string
	blockchain_ws string
}

type ValidatorRegistration struct {
	Validator   string
	BlockHeight uint64
}

type OldStatsResponse struct {
	DagsCount                models.Counter        `json:"dagsCount"`
	LastDagTimestamp         *models.NilableUint64 `json:"lastDagTimestamp" rlp:"nil"`
	LastPbftTimestamp        *models.NilableUint64 `json:"lastPbftTimestamp" rlp:"nil"`
	LastTransactionTimestamp *models.NilableUint64 `json:"lastTransactionTimestamp" rlp:"nil"`
	PbftCount                models.Counter        `json:"pbftCount"`
	TransactionsCount        models.Counter        `json:"transactionsCount"`
}

type OldValidator struct {
	Address   string `json:"address"`
	PbftCount uint64 `json:"pbftCount"`
	Rank      uint64 `json:"rank" rlp:"-"`
	Yield     string `json:"yield,omitempty" rlp:"-"`
}

type OldAddressStats struct {
	OldStatsResponse
	Address string `json:"address"`
}

type OldWeekStats struct {
	Validators []OldValidator
	Total      uint32
	Key        []byte `rlp:"-"`
}

type ValidatorRegistrationEvent struct {
	Validator   string `json:"validator"`
	BlockHeight uint64 `json:"blockHeight"`
}

func (m *AddValidatorRegistrationBlock) GetId() string {
	return m.id
}

func (m *AddValidatorRegistrationBlock) Apply(s *pebble.Storage) error {
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
	const MAX_BATCH_THRESHOLD = 1000

	registeredEvent, err := GetValidatorRegistrationLogs(client, 0, currentHead)

	if err != nil {
		log.Fatal(err)
	}

	batch := s.NewBatch()
	var last_key []byte
	var o OldAddressStats
	for {
		count := 0
		s.ForEachFromKey([]byte(pebble.GetPrefix(storage.AddressStats{})), last_key, func(key, res []byte) (stop bool) {
			err := rlp.DecodeBytes(res, &o)
			// fmt.Println(o)
			if err != nil {
				if err.Error() == "rlp: input list has too many elements for migration.OldStatsResponse, decoding into (migration.OldAddressStats).OldStatsResponse" {
					return false
				} else {
					log.WithFields(log.Fields{"migration": m.id, "error": err}).Fatal("Error decoding OldStatsResponse")
				}
			}

			validatorRegistrationData := registeredEvent[o.Address]
			fmt.Printf("Address: %s, RegisteredBlock: %d\n", o.Address, validatorRegistrationData.BlockHeight)

			newAddressStats := storage.AddressStats{
				Address: o.Address,
				StatsResponse: models.StatsResponse{
					DagsCount:                o.DagsCount,
					LastDagTimestamp:         o.LastDagTimestamp,
					LastPbftTimestamp:        o.LastPbftTimestamp,
					LastTransactionTimestamp: o.LastTransactionTimestamp,
					PbftCount:                o.PbftCount,
					ValidatorRegisteredBlock: &validatorRegistrationData.BlockHeight,
				},
			}
			if validatorRegistrationData.BlockHeight != 0 {
				fmt.Printf("Address updated : %s, RegisteredBlock: %d\n", o.Address, *newAddressStats.ValidatorRegisteredBlock)
			}
			err = batch.AddToBatchFullKey(&newAddressStats, key)

			if err != nil {
				log.WithFields(log.Fields{"migration": m.id, "error": err}).Fatal("Error adding AddressStats to batch")
			}
			fmt.Printf("Address updated : %s, RegisteredBlock: %d\n", o.Address, *newAddressStats.ValidatorRegisteredBlock)

			last_key = key
			count++
			return count == MAX_BATCH_THRESHOLD
		})
		batch.CommitBatch()
		batch = s.NewBatch()
		if count < MAX_BATCH_THRESHOLD {
			break
		}
	}

	batch = s.NewBatch()
	var last_key_week []byte
	var o_week OldWeekStats

	for {
		count := 0
		s.ForEachFromKey([]byte(pebble.GetPrefix(storage.WeekStats{})), last_key_week, func(key, res []byte) (stop bool) {
			err := rlp.DecodeBytes(res, &o_week)
			if err != nil {
				if err.Error() == "rlp: input list has too many elements for migration.OldValidator, decoding into (migration.OldWeekStats).OldValidator" {
					return false
				} else {
					log.WithFields(log.Fields{"migration": m.id, "error": err}).Fatal("Error decoding OldValidator")
				}
			}

			var weekValidators []models.Validator

			for _, validator := range o_week.Validators {
				registrationBlock := registeredEvent[validator.Address].BlockHeight
				weekValidators = append(weekValidators, models.Validator{
					Address:           validator.Address,
					PbftCount:         validator.PbftCount,
					Rank:              validator.Rank,
					Yield:             validator.Yield,
					RegistrationBlock: &registrationBlock,
				})
				fmt.Printf("Week stats updated : %s, RegisteredBlock: %d\n", o.Address, weekValidators[len(weekValidators)-1].RegistrationBlock)
			}

			newWeekStats := storage.WeekStats{
				Total:      o_week.Total,
				Validators: weekValidators,
				Key:        o_week.Key,
			}

			err = batch.AddToBatchFullKey(&newWeekStats, key)

			if err != nil {
				log.WithFields(log.Fields{"migration": m.id, "error": err}).Fatal("Error adding AddressStats to batch")
			}
			last_key = key
			count++
			return count == MAX_BATCH_THRESHOLD
		})
		batch.CommitBatch()
		batch = s.NewBatch()
		if count < MAX_BATCH_THRESHOLD {
			break
		}
	}
	batch.CommitBatch()
	return nil
}

func GetValidatorRegistrationLogs(client *chain.WsClient, from, to uint64) (map[string]ValidatorRegistration, error) {

	logs, err := client.GetLogs(from, to, []string{"0x00000000000000000000000000000000000000fe"}, [][]string{{"0xd09501348473474a20c772c79c653e1fd7e8b437e418fe235d277d2c88853251"}})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found %d logs\n", len(logs))

	validators := make(map[string]ValidatorRegistration)
	if len(logs) == 0 {
		return validators, nil
	}
	for _, eLog := range logs {
		validatorAddress := ethcommon.HexToAddress(eLog.Topics[1])

		validators[strings.ToLower(validatorAddress.Hex())] = ValidatorRegistration{Validator: strings.ToLower(validatorAddress.Hex()), BlockHeight: common.ParseUInt(eLog.BlockNumber)}
	}
	return validators, nil
}
