package oracle

import (
	"context"
	"errors"
	"math/big"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	// Import other necessary packages
	"github.com/Taraxa-project/taraxa-go-client/taraxa_client/dpos_contract_client/dpos_interface"
	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/contracts"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage/pebble"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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
	contract         *bind.BoundContract
	validatorsMutex  sync.Mutex
	latestValidators []YieldedValidator
}

type YieldedValidator struct {
	Account           string
	Yield             string
	Commisson         *uint64
	Rank              uint64
	RegistrationBlock uint64
	PbftCount         uint64
	Rating            uint64
}

type RawValidator struct {
	Address string
	Yield   string
}

func MakeOracle(blockchain_ws, signing_key, oracle_address string, chainId int, storage pebble.Storage) *Oracle {
	o := new(Oracle)
	o.storage = storage
	o.Eth = connect(blockchain_ws)
	o.signer = makeSigner(signing_key, chainId)
	o.oracleAddress = oracle_address
	o.chainId = chainId
	o.contract = o.makeContract()
	return o
}

func makeSigner(signingKey string, chainId int) *bind.TransactOpts {
	// Load your private key (securely)
	privateKey, err := crypto.HexToECDSA(signingKey)
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	// Create an auth object to use for the transaction
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(chainId)))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	return auth
}

func (o *Oracle) makeContract() *bind.BoundContract {
	// Define the contract address
	contractAddress := common.HexToAddress(o.oracleAddress)

	// Create an instance of your contract
	oracleAbi, err := abi.JSON(strings.NewReader(contracts.ApyOracle)) // Assuming contracts.ApyOracle is your ABI
	if err != nil {
		log.Fatalf("Failed to read ABI: %v", err)
	}
	contractInstance := bind.NewBoundContract(contractAddress, oracleAbi, o.Eth, o.Eth, o.Eth)
	return contractInstance
}

func (o *Oracle) UpdateValidators(validators []YieldedValidator) {
	log.Infof("Updating validators: %d", len(validators))
	log.Infof("Validators before: %v", len(o.latestValidators))
	o.latestValidators = validators

}

func (o *Oracle) PushValidators(validators []RawValidator) {
	o.validatorsMutex.Lock()
	defer o.validatorsMutex.Unlock()
	yieldedValidators := make([]YieldedValidator, 0)
	for _, validator := range validators {
		validatorData, err := FetchValidatorInfo(o.Eth, validator.Address)
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
		yieldedValidators[i].Rank = uint64(i + 1)
		yield, err := strconv.ParseFloat(yieldedValidators[i].Yield, 64)
		if err != nil {
			log.Fatalf("Failed to parse yield: %v", err)
		}
		yieldInt := uint64(yield * 1000)
		yieldedValidators[i].Rating = yieldedValidators[i].Rank * yieldInt
		yieldedValidators[i].Rating = yieldedValidators[i].Rank * yieldInt
	}
	o.UpdateValidators(yieldedValidators)
	log.Infof("Loading validators into oracle instance: %d", len(yieldedValidators))
}

func (o *Oracle) pushDataToContract() {
	if len(o.latestValidators) == 0 {
		log.Warn("No validator data to push")
		return
	}
	// Obtain the nonce for the account
	nonce, err := o.Eth.PendingNonceAt(context.Background(), o.signer.From)
	if err != nil {
		log.Fatalf("Failed to get nonce: %v", err)
	}
	o.signer.Nonce = big.NewInt(int64(nonce))

	// Suggest a gas price
	gasPrice, err := o.Eth.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to suggest gas price: %v", err)
	}
	o.signer.GasPrice = gasPrice

	validatorAddresses := make([]common.Address, 0)

	for _, validator := range o.latestValidators {
		validatorAddresses = append(validatorAddresses, common.HexToAddress(validator.Account))
	}

	log.Infof("Pushing validator data to contract: %v", validatorAddresses)

	_, err = o.contract.Transact(o.signer, "batchUpdateNodeData", validatorAddresses, o.latestValidators)
	if err != nil {
		log.Fatalf("Failed to transact with contract: %v", err)
	}
}

func connect(url string) *ethclient.Client {
	var err error
	var client *ethclient.Client
	for {
		client, err = ethclient.Dial(url)
		if err != nil {
			log.WithError(err).Error("Failed to connect to eth client")
			time.Sleep(5 * time.Second)
			continue
		}
		_, err := client.BlockNumber(context.Background())
		if err != nil {
			log.WithError(err).Error("Failed to get current block")
			break
		} else {
			break
		}
	}
	return client
}

func RegisterCron(o *Oracle, yield_saving_interval int) {
	s := gocron.NewScheduler(time.UTC)
	var err error
	_, err = s.Every(10).Seconds().Do(func() {
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

// func getValidatorDatas(storage storage.Storage, client *ethclient.Client) []YieldedValidator {
// 	tm, _ := goment.New()
// 	year := int32(tm.ISOWeekYear())
// 	week := int32(tm.ISOWeek())

// 	// print the year & week
// 	log.Infof("Current year: %d", year)
// 	log.Infof("Current week: %d", week)

// 	stats := storage.GetWeekStats(year, week)
// 	stats.Sort()
// 	validators := make([]YieldedValidator, 0)

// 	for k, v := range stats.Validators {
// 		yieldedValidator := YieldedValidator{
// 			Address:           v.Address,
// 			Rank:              uint64(k + 1),
// 			Yield:             v.Yield,
// 			RegistrationBlock: *v.RegistrationBlock,
// 			PbftCount:         v.PbftCount,
// 		}

// 		statsResponse := storage.GetAddressStats(v.Address).StatsResponse
// 		yieldedValidator.Commisson = statsResponse.Commission
// 		yieldedValidator.Rating = yieldedValidator.Rank * taracommon.ParseUInt(yieldedValidator.Yield)
// 		validators = append(validators, yieldedValidator)
// 	}
// 	return validators
// }

// will not be used in the first primitive version
func calculateRating(validator YieldedValidator, commission *uint64, client *ethclient.Client) float64 {
	if commission == nil {
		return 0
	}

	currentBlock, err := client.BlockByNumber(context.Background(), nil)

	if err != nil {
		log.Fatalf("Failed to get current block: %v", err)
	}

	blocksSinceRegistration := currentBlock.NumberU64() - validator.RegistrationBlock
	commission_float := float64(*validator.Commisson)
	yield_float, err := strconv.ParseFloat(validator.Yield, 64)
	if err != nil {
		log.Fatalf("Failed to parse yield: %v", err)
	}
	commission_percentage := commission_float / float64(100000)
	adjusted_apy := (1 - commission_percentage) * yield_float * 100
	continuity := float64(blocksSinceRegistration) / float64(currentBlock.NumberU64()-validator.RegistrationBlock)

	//w1 * (APY) - (Commission * w2) + w3 * Continuity + w4 * stake
	score := float64(0.4)*adjusted_apy - float64(0.1)*commission_float + float64(0.5)*continuity
	return score
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
