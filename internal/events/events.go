package events

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/contracts"
	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
)

var rewardTopics = map[string]bool{
	"CommissionRewardsClaimed(address,address,uint256)":     true,
	"RewardsClaimed(address,address,uint256)":               true,
	"UndelegateConfirmed(address,address,uint256)":          true,
	"UndelegateConfirmedV2(address,address,uint64,uint256)": true,
}

func DecodeEventDynamic(log models.EventLog) (string, any, error) {
	relevantAbi := contracts.ContractABIs[strings.ToLower(log.Address)]
	if relevantAbi == "" {
		return "", nil, nil
	}
	// Convert the hex-encoded data to bytes
	trimmed := strings.TrimPrefix(log.Data, "0x")
	data, err := hex.DecodeString(trimmed)

	if err != nil {
		return "", nil, err
	}

	contractABI, error := abi.JSON(strings.NewReader(relevantAbi))
	if error != nil {
		return "", nil, error
	}

	// Get the event for the topic

	event, err := contractABI.EventByID(ethcommon.HexToHash(log.Topics[0]))

	if err != nil {
		return "", nil, err
	}

	unpacked, err := contractABI.Unpack(event.Name, data)

	if err != nil {
		return "", nil, err
	}

	params, err := common.ParseToString(unpacked)

	if err != nil {
		return "", nil, err
	}

	return event.Sig, params, nil
}

func DecodeRewardsTopics(logs []models.EventLog) (decodedEvents []LogReward, err error) {
	for _, log := range logs {
		if !strings.EqualFold(log.Address, common.DposContractAddress) {
			continue
		}
		name, d, err := DecodeEventDynamic(log)
		if err != nil {
			return nil, err
		}
		decoded := d.([]any)
		if rewardTopics[name] {
			account := ethcommon.HexToAddress(log.Topics[1])
			validator := ethcommon.HexToAddress(log.Topics[2])
			value, _ := big.NewInt(0).SetString(fmt.Sprintf("%v", decoded[0]), 10)

			decodedEvents = append(decodedEvents, LogReward{
				EventName: name,
				Account:   account.Hex(),
				Validator: validator.Hex(),
				Value:     value,
			})
		}
	}
	return decodedEvents, err
}
