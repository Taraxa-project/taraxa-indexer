package migration

import (
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
)

// Define the contract ABI for the Solidity smart contract
var contractABI = `[{"anonymous":false,"inputs":[{"indexed":true,"name":"validator","type":"address"}],"name":"ValidatorRegistered","type":"event"}]`

type ValidatorRegistration struct {
	Validator   ethcommon.Address
	BlockHeight uint64
}

func GetValidatorsRegisteredInBlock(client *chain.WsClient, from, to uint64) ([]ValidatorRegistration, error) {
	if from > to {
		return nil, fmt.Errorf("from block %d is greater than to block %d", from, to)
	}
	// Ethereum contract address
	contractAddress := ethcommon.HexToAddress("0x00000000000000000000000000000000000000fe")

	// Create an Ethereum client
	// Parse the contract ABI
	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		log.Fatal(err)
	}

	var query ethereum.FilterQuery
	if from == 0 && to == 0 {
		query = ethereum.FilterQuery{
			Addresses: []ethcommon.Address{contractAddress},
			Topics: [][]ethcommon.Hash{
				{
					parsedABI.Events["ValidatorRegistered"].ID,
				},
			},
		}
	} else {
		// Create a filter query
		query = ethereum.FilterQuery{
			Addresses: []ethcommon.Address{contractAddress},
			Topics: [][]ethcommon.Hash{
				{
					parsedABI.Events["ValidatorRegistered"].ID,
				},
			},
			FromBlock: big.NewInt(int64(from)),
			ToBlock:   big.NewInt(int64(to)),
		}
	}

	topicStrings := [][]string{}
	for i, topic := range query.Topics {
		for j, hash := range topic {
			topicStrings[i][j] = hash.Hex()
		}
	}

	addressStrings := []string{}
	for _, address := range query.Addresses {
		addressStrings = append(addressStrings, address.Hex())
	}

	// Fetch events
	logs, err := client.GetLogs(from, to, addressStrings, topicStrings)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found logs: %d\n", len(logs))
	var validators []ValidatorRegistration
	// Process and print the events
	for _, eLog := range logs {
		event := struct {
			Validator ethcommon.Address `json:"validator"`
		}{}

		event.Validator = ethcommon.HexToAddress(eLog.Topics[1])

		fmt.Printf("Validator Registered: %s\n", event.Validator.Hex())
		validators = append(validators, ValidatorRegistration{Validator: event.Validator, BlockHeight: common.ParseUInt(eLog.BlockNumber)})
	}
	return validators, nil
}
