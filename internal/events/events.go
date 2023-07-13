package events

import (
	"encoding/hex"
	"math/big"
	"strings"

	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/contracts"
	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
)

const commissionRewardsClaimedName = "CommissionRewardsClaimed(address,address,uint256)"
const rewardsClaimedName = "RewardsClaimed(address,address,uint256)"

func DecodeEventDynamic(log models.EventLog) (string, []string, error) {

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

	params, err := common.ParseToStringSlice(unpacked)

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

		if name == rewardsClaimedName || name == commissionRewardsClaimedName {
			account := ethcommon.HexToAddress(log.Topics[1])
			validator := ethcommon.HexToAddress(log.Topics[2])
			value, _ := big.NewInt(0).SetString(decoded[0], 10)

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
