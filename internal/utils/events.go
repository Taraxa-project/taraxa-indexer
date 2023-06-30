package utils

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// LogReward ..
type LogReward struct {
	Account   string
	Validator string
	Value     *big.Int
	EventName string
}

type RewardsClaimedEvent struct {
	Account   string
	Validator string
	Amount    *big.Int
}

type CommissionRewardsClaimedEvent struct {
	Account   string
	Validator string
	Amount    *big.Int
}

const dposABI = `[{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"account","type":"address"},{"indexed":true,"internalType":"address","name":"validator","type":"address"},{"indexed":false,"internalType":"uint256","name":"amount","type":"uint256"}],"name":"RewardsClaimed","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"account","type":"address"},{"indexed":true,"internalType":"address","name":"validator","type":"address"},{"indexed":false,"internalType":"uint256","name":"amount","type":"uint256"}],"name":"CommissionRewardsClaimed","type":"event"},{"inputs":[{"internalType":"address","name":"validator","type":"address"}],"name":"claimCommissionRewards","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"validator","type":"address"}],"name":"claimRewards","outputs":[],"stateMutability":"nonpayable","type":"function"}]`

func DecodeEvent(log models.EventLog) (interface{}, error) {
	// Convert the hex-encoded data to bytes
	trimmed := strings.TrimPrefix(log.Data, "0x")
	data, err := hex.DecodeString(trimmed)

	if err != nil {
		return nil, err
	}

	contractABI, error := abi.JSON(strings.NewReader(dposABI))
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
		account := common.HexToAddress(log.Topics[1])
		validator := common.HexToAddress(log.Topics[2])

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
		account := common.HexToAddress(log.Topics[1])
		validator := common.HexToAddress(log.Topics[2])

		// Set the addresses in the event struct
		event.Account = account.Hex()
		event.Validator = validator.Hex()
		return &event, nil
	}
	return nil, fmt.Errorf("no matching event topic found")
}

func DecodeRewardsTopics(logs []models.EventLog) (decodedEvents []LogReward, err error) {
	for _, log := range logs {
		if strings.ToLower(log.Address) != strings.ToLower("0x00000000000000000000000000000000000000fe") {
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
