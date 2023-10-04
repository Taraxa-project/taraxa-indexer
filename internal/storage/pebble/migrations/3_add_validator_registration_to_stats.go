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

	step := uint64(100000)

	batch := s.NewBatch()

	for startBlock := uint64(0); startBlock < currentHead; startBlock += step {
		endBlock := startBlock + step
		validators, err := GetValidatorsRegisteredInBlock(client, startBlock, endBlock)
		if err != nil {
			log.Fatal(err)
		}

		for _, validator := range validators {
			addressStats := s.GetAddressStats(strings.ToLower(validator.Validator))
			if addressStats == nil {
				addressStats = &storage.AddressStats{
					Address: strings.ToLower(validator.Validator),
					StatsResponse: models.StatsResponse{
						ValidatorRegisteredBlock: &validator.BlockHeight,
					},
				}
			} else {
				addressStats.ValidatorRegisteredBlock = &validator.BlockHeight
			}
			batch.AddToBatch(addressStats, addressStats.Address, 0)
		}
	}
	batch.CommitBatch()
	return nil
}

func GetValidatorsRegisteredInBlock(client *chain.WsClient, from, to uint64) ([]ValidatorRegistration, error) {
	if from > to {
		return nil, fmt.Errorf("from block %d is greater than to block %d", from, to)
	}

	logs, err := client.GetLogs(from, to, []string{"0x00000000000000000000000000000000000000fe"}, [][]string{{"0xd09501348473474a20c772c79c653e1fd7e8b437e418fe235d277d2c88853251"}})
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("Found %d validator registrations in blocks %d-%d", len(logs), from, to)

	var validators []ValidatorRegistration
	for _, eLog := range logs {
		event := struct {
			Validator string `json:"validator"`
		}{}

		event.Validator = strings.ToLower(ethcommon.HexToAddress(eLog.Topics[1]).Hex())
		validators = append(validators, ValidatorRegistration{Validator: event.Validator, BlockHeight: common.ParseUInt(eLog.BlockNumber)})
		log.Infof("Found validator %s registered in block %d", event.Validator, common.ParseUInt(eLog.BlockNumber))
	}
	return validators, nil
}
