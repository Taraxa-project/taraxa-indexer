package events

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Define the contract ABI for the Solidity smart contract
var contractABI = `[{"anonymous":false,"inputs":[{"indexed":true,"name":"validator","type":"address"}],"name":"ValidatorRegistered","type":"event"}]`

type ValidatorRegistration struct {
	Validator   common.Address
	BlockHeight uint64
}

func GetValidatorsRegisteredInBlock(url string, from, to uint64) ([]ValidatorRegistration, error) {
	if from > to {
		return nil, fmt.Errorf("from block %d is greater than to block %d", from, to)
	}
	// Ethereum contract address
	contractAddress := common.HexToAddress("0x00000000000000000000000000000000000000fe")

	// Create an Ethereum client
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	// Parse the contract ABI
	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		log.Fatal(err)
	}

	var query ethereum.FilterQuery
	if from == 0 && to == 0 {
		query = ethereum.FilterQuery{
			Addresses: []common.Address{contractAddress},
			Topics: [][]common.Hash{
				{
					parsedABI.Events["ValidatorRegistered"].ID,
				},
			},
		}
	} else {
		// Create a filter query
		query = ethereum.FilterQuery{
			Addresses: []common.Address{contractAddress},
			Topics: [][]common.Hash{
				{
					parsedABI.Events["ValidatorRegistered"].ID,
				},
			},
			FromBlock: big.NewInt(int64(from)),
			ToBlock:   big.NewInt(int64(to)),
		}
	}

	// Fetch events
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found logs: %d\n", len(logs))
	var validators []ValidatorRegistration
	// Process and print the events
	for _, eLog := range logs {
		event := struct {
			Validator common.Address `json:"validator"`
		}{}

		event.Validator = common.BytesToAddress(eLog.Topics[1].Bytes())

		fmt.Printf("Validator Registered: %s\n", event.Validator.Hex())
		validators = append(validators, ValidatorRegistration{Validator: event.Validator, BlockHeight: eLog.BlockNumber})
	}
	return validators, nil
}
