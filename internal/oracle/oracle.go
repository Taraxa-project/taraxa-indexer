package oracle

import (
	"context"
	"math/big"
	"strconv"
	"strings"
	"time"

	// Import other necessary packages
	taracommon "github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/contracts"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-co-op/gocron"
	"github.com/nleeper/goment"
	log "github.com/sirupsen/logrus"
)

type YieldedValidator struct {
	Address           string
	Yield             string
	Commisson         *uint64
	Rank              uint64
	RegistrationBlock uint64
	PbftCount         uint64
	Rating            uint64
}

func pushDataToContract(ws, signingKey, oracleAddress string, chainID int64, client *ethclient.Client, storage storage.Storage) {
	// Load your private key (securely)
	privateKey, err := crypto.HexToECDSA(signingKey)
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	// Create an auth object to use for the transaction
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	// Obtain the nonce for the account
	nonce, err := client.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		log.Fatalf("Failed to get nonce: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))

	// Suggest a gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to suggest gas price: %v", err)
	}
	auth.GasPrice = gasPrice

	// Define the contract address
	contractAddress := common.HexToAddress(oracleAddress)

	// Create an instance of your contract
	oracleAbi, err := abi.JSON(strings.NewReader(contracts.ApyOracle)) // Assuming contracts.ApyOracle is your ABI
	if err != nil {
		log.Fatalf("Failed to read ABI: %v", err)
	}
	contractInstance := bind.NewBoundContract(contractAddress, oracleAbi, client, client, client)

	validators := getValidatorDatas(storage, client)

	log.Infof("Pushing  amount of validator data to contract: %d", len(validators))

	validatorAddresses := make([]common.Address, 0)

	for _, validator := range validators {
		validatorAddresses = append(validatorAddresses, common.HexToAddress(validator.Address))
	}

	_, err = contractInstance.Transact(auth, "batchUpdateNodeData", validatorAddresses, validators)
	if err != nil {
		log.Fatalf("Failed to transact with contract: %v", err)
	}
}

func RegisterOracleCron(blockchain_ws, signing_key, oracle_address string, yield_saving_interval int, chainId int, storage storage.Storage) {
	s := gocron.NewScheduler(time.UTC)

	client, err := ethclient.Dial(blockchain_ws)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	_, err = s.Every(yield_saving_interval * 4).Seconds().Do(func() {
		log.Info("Oracle cron started")
		pushDataToContract(blockchain_ws, signing_key, oracle_address, int64(chainId), client, storage)
	})
	if err != nil {
		log.Fatalf("Failed to schedule cron: %v", err)
	}
	// check if job was successfully ran
	_, t := s.NextRun()
	log.Info("Oracle cron scheduled at ", t)
	s.StartAsync()
}

func getValidatorDatas(storage storage.Storage, client *ethclient.Client) []YieldedValidator {
	tm, _ := goment.New()
	year := int32(tm.ISOWeekYear())
	week := int32(tm.ISOWeek())

	stats := storage.GetWeekStats(year, week)
	stats.Sort()
	validators := make([]YieldedValidator, 0)

	for k, v := range stats.Validators {
		yieldedValidator := YieldedValidator{
			Address:           v.Address,
			Rank:              uint64(k + 1),
			Yield:             v.Yield,
			RegistrationBlock: *v.RegistrationBlock,
			PbftCount:         v.PbftCount,
		}

		statsResponse := storage.GetAddressStats(v.Address).StatsResponse
		yieldedValidator.Commisson = statsResponse.Commission
		yieldedValidator.Rating = yieldedValidator.Rank * taracommon.ParseUInt(yieldedValidator.Yield)
		validators = append(validators, yieldedValidator)
	}
	return validators
}

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
