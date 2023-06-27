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
	Account   common.Address
	Validator common.Address
	Amount    *big.Int
}

type CommissionRewardsClaimedEvent struct {
	Account   common.Address
	Validator common.Address
	Amount    *big.Int
}

const dposABI = `[{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"account","type":"address"},{"indexed":true,"internalType":"address","name":"validator","type":"address"},{"indexed":false,"internalType":"uint256","name":"amount","type":"uint256"}],"name":"RewardsClaimed","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"account","type":"address"},{"indexed":true,"internalType":"address","name":"validator","type":"address"},{"indexed":false,"internalType":"uint256","name":"amount","type":"uint256"}],"name":"CommissionRewardsClaimed","type":"event"}]`

func decodeEvent(log models.EventLog) (interface{}, error) {
	// Convert the hex-encoded data to bytes
	data, err := hex.DecodeString(log.Data)
	if err != nil {
		return nil, err
	}

	contractABI, error := abi.JSON(strings.NewReader(dposABI))
	if error != nil {
		return nil, error
	}
	// Decode the event based on its topic
	switch log.Topics[0] {
	case crypto.Keccak256Hash([]byte("RewardsClaimed(address,address,uint256)")).Hex():
		var event RewardsClaimedEvent
		err := contractABI.UnpackIntoInterface(&event, "RewardsClaimed", data)
		if err != nil {
			return nil, err
		}
		return &event, nil

	case crypto.Keccak256Hash([]byte("CommissionRewardsClaimed(address,address,uint256)")).Hex():
		var event CommissionRewardsClaimedEvent
		err := contractABI.UnpackIntoInterface(&event, "CommissionRewardsClaimed", data)
		if err != nil {
			return nil, err
		}
		return &event, nil

	default:
		return nil, fmt.Errorf("unknown event topic")
	}
}

func decodeRewardsTopics(logs []models.EventLog) (decodedEvents []LogReward, err error) {
	for _, log := range logs {
		decoded, err := decodeEvent(log)
		if err != nil {
			return nil, err
		}

		switch event := decoded.(type) {
		case *RewardsClaimedEvent:
			decodedEvents = append(decodedEvents, LogReward{
				EventName: "RewardsClaimed",
				Account:   event.Account.Hex(),
				Validator: event.Validator.Hex(),
				Value:     event.Amount,
			})

		case *CommissionRewardsClaimedEvent:
			decodedEvents = append(decodedEvents, LogReward{
				EventName: "CommissionRewardsClaimed",
				Account:   event.Account.Hex(),
				Validator: event.Validator.Hex(),
				Value:     event.Amount,
			})
		}
	}

	return decodedEvents, err
}
