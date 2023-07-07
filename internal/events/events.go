package events

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/contracts"
	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
)

func DecodeEventDynamic(log models.EventLog) (string, []string, error) {

	relevantAbi := contracts.ContractABIs[log.Address]
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

	params, err := parseToStringSlice(unpacked)

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
		name, decoded, err := DecodeEventDynamic(log)
		if err != nil {
			return nil, err
		}

		if name == "RewardsClaimed" || name == "CommissionRewardsClaimed" {
			account := ethcommon.HexToAddress(log.Topics[1])
			validator := ethcommon.HexToAddress(log.Topics[2])
			value, _ := big.NewInt(0).SetString(decoded[0], 10)

			if name == "RewardsClaimed" && decoded[0] != "" {
				decodedEvents = append(decodedEvents, LogReward{
					EventName: "RewardsClaimed",
					Account:   account.Hex(),
					Validator: validator.Hex(),
					Value:     value,
				})
			}

			if name == "CommissionRewardsClaimed" && decoded[0] != "" {
				decodedEvents = append(decodedEvents, LogReward{
					EventName: "CommissionRewardsClaimed",
					Account:   account.Hex(),
					Validator: validator.Hex(),
					Value:     value,
				})
			}
		}
	}
	return decodedEvents, err
}

func parseToStringSlice(data []interface{}) ([]string, error) {
	result := make([]string, len(data))

	for i, item := range data {
		switch val := item.(type) {
		case string:
			result[i] = val
		case int:
			result[i] = strconv.Itoa(val)
		case int64:
			result[i] = strconv.FormatInt(val, 10)
		case float64:
			result[i] = strconv.FormatFloat(val, 'f', -1, 64)
		case *big.Int:
			result[i] = val.String()
		default:
			return nil, fmt.Errorf("failed to convert element at index %d to string", i)
		}
	}

	return result, nil
}
