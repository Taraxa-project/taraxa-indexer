package oracle

import (
	"context"
	"errors"
	"math/big"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	// Import other necessary packages
	"github.com/Taraxa-project/taraxa-go-client/taraxa_client/dpos_contract_client/dpos_interface"
	apy_oracle "github.com/Taraxa-project/taraxa-indexer/abi/oracle"
	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/contracts"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage/pebble"
	"github.com/Taraxa-project/taraxa-indexer/internal/transact"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-co-op/gocron"
	log "github.com/sirupsen/logrus"
)

type Oracle struct {
	storage          pebble.Storage
	Eth              *ethclient.Client
	signer           *bind.TransactOpts
	oracleAddress    string
	chainId          int
	contract         *apy_oracle.ApyOracle
	validatorsMutex  sync.Mutex
	LatestValidators []YieldedValidator
}

func MakeOracle(rpc *ethclient.Client, signing_key, oracle_address string, chainId int, storage pebble.Storage) *Oracle {
	o := new(Oracle)
	o.storage = storage
	o.Eth = rpc
	o.signer = transact.MakeSigner(signing_key, chainId)
	o.oracleAddress = oracle_address
	o.chainId = chainId
	contract, err := apy_oracle.NewApyOracle(common.HexToAddress(o.oracleAddress), o.Eth)
	if err != nil {
		log.Fatalf("Failed to create contract: %v", err)
	}
	o.contract = contract
	log.Info("Oracle initialized")
	return o
}

func (o *Oracle) UpdateValidators(validators []YieldedValidator) {
	log.Infof("Updating validators: %d", len(validators))
	log.Infof("Validators before: %v", len(o.LatestValidators))
	o.LatestValidators = validators
	o.pushDataToContract()
}

func (o *Oracle) PushValidators(validators []RawValidator) {
	o.validatorsMutex.Lock()
	defer o.validatorsMutex.Unlock()
	yieldedValidators := make([]YieldedValidator, 0)
	for _, validator := range validators {
		validatorData, err := FetchValidatorInfo(o.Eth, validator.Address.Hex())
		if err != nil {
			if err.Error() == "Validator does not exist" {
				continue
			}
			log.Fatalf("Failed to fetch validator info: %v", err)
		}
		commission := uint64(validatorData.Commission)
		yieldedValidator := YieldedValidator{
			Account:           validator.Address,
			Yield:             validator.Yield,
			Commisson:         &commission,
			Rank:              0,
			RegistrationBlock: 0,
			PbftCount:         0,
			Rating:            0,
		}
		yieldedValidators = append(yieldedValidators, yieldedValidator)
	}
	// sort by yield and add positions as rank and rating
	for i := range yieldedValidators {
		yieldedValidators[i].Rank = uint16(i + 1)
		yield, err := strconv.ParseFloat(yieldedValidators[i].Yield, 64)
		if err != nil {
			log.Fatalf("Failed to parse yield: %v", err)
		}
		yieldInt := uint64(yield * 1000)
		yieldedValidators[i].Rating = uint64(yieldedValidators[i].Rank) * yieldInt
	}
	o.UpdateValidators(yieldedValidators)
	log.Infof("Loading validators into oracle instance: %d", len(yieldedValidators))
}

func (o *Oracle) pushDataToContract() {
	if len(o.LatestValidators) == 0 {
		log.Warn("No validator data to push")
		return
	}

	validatorDatas := make([]NodeData, 0)

	for _, validator := range o.LatestValidators {
		data := validator.ToNodeData(o.Eth)
		validatorDatas = append(validatorDatas, data)
	}

	// sort data by rating in descending order
	sort.Slice(validatorDatas, func(i, j int) bool {
		return validatorDatas[i].Rating.Cmp(validatorDatas[j].Rating) == 1
	})

	for {
		_, err := o.contract.BatchUpdateNodeData(o.signer, validatorDatas)
		if err != nil {
			if strings.Contains(err.Error(), "Transaction already in transactions pool") || strings.Contains(err.Error(), "nonce too low") || strings.Contains(err.Error(), "out of gas") {
				time.Sleep(1 * time.Second)
				continue
			}
			log.Fatalf("Failed to batch update node data: %v", err)
		}

		log.Infof("Pushed %d validators to contract", len(validatorDatas))
		// wait 1 second
		time.Sleep(1 * time.Second)
		nodeCount, err := o.contract.NodeCount(nil)
		if err != nil {
			log.Errorf("Failed to get node count: %v", err)
		}
		if nodeCount.Cmp(big.NewInt(int64(len(validatorDatas)))) != 0 {
			log.Errorf("Node count mismatch: %d != %d", nodeCount, len(validatorDatas))
		} else {
			break
		}
	}
}

func RegisterCron(o *Oracle, yield_saving_interval int) {
	s := gocron.NewScheduler(time.UTC)
	var err error
	_, err = s.Every(yield_saving_interval).Seconds().Do(func() {
		log.Info("Oracle cron started")
		o.pushDataToContract()
	})
	if err != nil {
		log.Fatalf("Failed to schedule cron: %v", err)
	}
	// check if job was successfully ran
	_, t := s.NextRun()
	log.Info("Oracle cron scheduled at ", t)
	s.StartAsync()
}

// FetchValidatorInfo fetches the ValidatorBasicInfo for a given validator address.
func FetchValidatorInfo(client chain.EthereumClient, validatorAddress string) (*dpos_interface.DposInterfaceValidatorBasicInfo, error) {
	if client == nil {
		return nil, errors.New("Ethereum client is not available")
	}

	// Define the contract address
	contractAddress := common.HexToAddress("0x00000000000000000000000000000000000000fe")

	// Parse the ABI
	parsedABI, err := abi.JSON(strings.NewReader(contracts.ContractABIs["0x00000000000000000000000000000000000000fe"]))
	if err != nil {
		return nil, err
	}

	// Prepare the call input data
	data, err := parsedABI.Pack("getValidator", common.HexToAddress(validatorAddress))
	if err != nil {
		return nil, err
	}

	// Call the contract
	msg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: data,
	}
	result, err := client.CallContract(context.Background(), msg, nil)
	// we need to treat the case where Error fetching validator info: Validator does not exist
	if err != nil {
		if err.Error() == "Error fetching validator info: Validator does not exist" {
			return nil, nil
		}
		return nil, err
	}

	// Unpack the result
	unpacked, err := parsedABI.Unpack("getValidator", result)
	if err != nil {
		return nil, err
	}
	unpackedValue := reflect.ValueOf(unpacked[0])
	if unpackedValue.Kind() != reflect.Struct {
		return nil, errors.New("unpacked result is not a struct")
	}
	var validatorInfo dpos_interface.DposInterfaceValidatorBasicInfo
	for i := 0; i < unpackedValue.NumField(); i++ {
		field := unpackedValue.Field(i)

		switch i {
		case 0:
			validatorInfo.TotalStake = field.Interface().(*big.Int)
		case 1:
			validatorInfo.CommissionReward = field.Interface().(*big.Int)
		case 2:
			validatorInfo.Commission = field.Interface().(uint16)
		case 3:
			validatorInfo.LastCommissionChange = field.Interface().(uint64)
		case 4:
			validatorInfo.UndelegationsCount = field.Interface().(uint16)
		case 5:
			validatorInfo.Owner = field.Interface().(common.Address)
		case 6:
			validatorInfo.Description = field.Interface().(string)
		case 7:
			validatorInfo.Endpoint = field.Interface().(string)
		}
	}
	return &validatorInfo, nil
}

// go run main.go --blockchain_ws=ws://localhost:8777 --log_level=debug --chain_id=842 --signing_key=472a3f59fe3d81cda76dbb2a64825e46c4b067ae559cd4dfc784869da80bd05e --oracle_address=0x4076f9669fd33e55545823c4cB9f1abA7cfa480B

func MakeMockOracle(eth *ethclient.Client) *Oracle {
	return &Oracle{
		Eth: eth,
	}
}
