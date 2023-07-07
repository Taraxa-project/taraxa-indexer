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
	"github.com/ethereum/go-ethereum/crypto"
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

	return event.Name, params, nil
}

func DecodeEvent(log models.EventLog) (interface{}, error) {
	// Convert the hex-encoded data to bytes
	trimmed := strings.TrimPrefix(log.Data, "0x")
	data, err := hex.DecodeString(trimmed)

	if err != nil {
		return nil, err
	}

	contractABI, error := abi.JSON(strings.NewReader(contracts.ContractABIs[log.Address]))
	if error != nil {
		return nil, error
	}

	rewardsClaimedTopic := crypto.Keccak256Hash([]byte("RewardsClaimed(address,address,uint256)"))
	commissionRewardsClaimedTopic := crypto.Keccak256Hash([]byte("CommissionRewardsClaimed(address,address,uint256)"))
	// Decode the event based on its topic
	switch log.Topics[0] {
	case rewardsClaimedTopic.Hex():
		var event RewardsClaimedEvent
		err := contractABI.UnpackIntoInterface(&event, "RewardsClaimed", data)

		if err != nil {
			return nil, err
		}
		account := ethcommon.HexToAddress(log.Topics[1])
		validator := ethcommon.HexToAddress(log.Topics[2])

		// Set the addresses in the event struct
		event.Account = account.Hex()
		event.Validator = validator.Hex()
		return &event, nil

	case commissionRewardsClaimedTopic.Hex():
		var event CommissionRewardsClaimedEvent
		err := contractABI.UnpackIntoInterface(&event, "CommissionRewardsClaimed", data)
		if err != nil {
			return nil, err
		}
		account := ethcommon.HexToAddress(log.Topics[1])
		validator := ethcommon.HexToAddress(log.Topics[2])

		// Set the addresses in the event struct
		event.Account = account.Hex()
		event.Validator = validator.Hex()
		return &event, nil
	}
	return nil, nil
}

func DecodeRewardsTopics(logs []models.EventLog) (decodedEvents []LogReward, err error) {
	for _, log := range logs {
		if !strings.EqualFold(log.Address, common.DposContractAddress) {
			continue
		}
		decoded, err := DecodeEvent(log)
		if err != nil {
			return nil, err
		}

		switch event := decoded.(type) {
		case *RewardsClaimedEvent:
			decodedEvents = append(decodedEvents, LogReward{
				EventName: "RewardsClaimed",
				Account:   event.Account,
				Validator: event.Validator,
				Value:     event.Amount,
			})

		case *CommissionRewardsClaimedEvent:
			decodedEvents = append(decodedEvents, LogReward{
				EventName: "CommissionRewardsClaimed",
				Account:   event.Account,
				Validator: event.Validator,
				Value:     event.Amount,
			})
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
