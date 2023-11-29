package lara

import (
	"log"
	"math/big"
	"strings"

	lara_contract "github.com/Taraxa-project/taraxa-indexer/abi/lara"
	apy_oracle "github.com/Taraxa-project/taraxa-indexer/abi/oracle"
	"github.com/Taraxa-project/taraxa-indexer/internal/contracts"
	"github.com/Taraxa-project/taraxa-indexer/internal/oracle"
	"github.com/Taraxa-project/taraxa-indexer/internal/transact"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type State struct {
	epochDuration                 *big.Int
	isEpochRunning                bool
	lastEpochTotalDelegatedAmount *big.Int
	validatorStakes               map[common.Address]*big.Int
	validators                    []oracle.NodeData
}
type Lara struct {
	deploymentAddress string
	oracleAddress     string
	Eth               *ethclient.Client
	signer            *bind.TransactOpts
	chainID           *int
	contract          *lara_contract.LaraContract
	oracle            *apy_oracle.ApyOracle
	state             State
}

func MakeLara(rpc *ethclient.Client, signing_key, deployment_address, oracle_address string, chainID int) *Lara {
	l := new(Lara)
	l.Eth = rpc
	l.signer = transact.MakeSigner(signing_key, chainID)
	l.deploymentAddress = deployment_address
	l.oracleAddress = oracle_address
	l.chainID = &chainID
	contract, err := lara_contract.NewLaraContract(common.HexToAddress(l.deploymentAddress), l.Eth)
	if err != nil {
		log.Fatalf("Failed to create contract: %v", err)
	}

	l.contract = contract
	l.state = l.SyncState()
	return l
}

func (l *Lara) makeContract() *bind.BoundContract {
	// Define the contract address
	contractAddress := common.HexToAddress(l.deploymentAddress)

	// Create an instance of your contract
	oracleAbi, err := abi.JSON(strings.NewReader(contracts.Lara))
	if err != nil {
		log.Fatalf("Failed to read ABI: %v", err)
	}
	contractInstance := bind.NewBoundContract(contractAddress, oracleAbi, l.Eth, l.Eth, l.Eth)
	return contractInstance
}

func (l *Lara) SyncState() State {
	opts := &bind.CallOpts{
		Pending:     false,
		From:        l.signer.From,
		BlockNumber: nil,
		Context:     nil,
	}
	epochDuration, err := l.contract.EpochDuration(opts)
	if err != nil {
		log.Fatalf("Failed to get epoch duration: %v", err)
	}

	isEpochRunning, err := l.contract.IsEpochRunning(opts)
	if err != nil {
		log.Fatalf("Failed to get epoch running: %v", err)
	}

	lastEpochTotalDelegatedAmount, err := l.contract.LastEpochTotalDelegatedAmount(opts)
	if err != nil {
		log.Fatalf("Failed to get last epoch total delegated amount: %v", err)
	}

	// fetch validators from 0 until reverted and put in map
	pos := big.NewInt(0)
	validatorStakes := make(map[common.Address]*big.Int)
	validators := make([]oracle.NodeData, 0)
	for {
		validator, err := l.contract.Validators(opts, pos)
		if err != nil {
			log.Fatalf("Failed to get validator: %v", err)
			break
		}
		totalStakeAtValidator, err := l.contract.ProtocolTotalStakeAtValidator(opts, validator)
		if err != nil {
			log.Fatalf("Failed to get total stake at validator: %v", err)
		}
		validatorStakes[validator] = totalStakeAtValidator

		// fetch node data from oracle
		nodeData, err := l.oracle.Nodes(opts, validator)
		if err != nil {
			log.Fatalf("Failed to get node data: %v", err)
		}
		validators = append(validators, nodeData)
	}

	l.state = State{
		epochDuration:                 epochDuration,
		isEpochRunning:                isEpochRunning,
		lastEpochTotalDelegatedAmount: lastEpochTotalDelegatedAmount,
		validatorStakes:               validatorStakes,
		validators:                    validators,
	}

	return l.state
}

func (l *Lara) StartEpoch() {
	opts := &bind.TransactOpts{
		From:     l.signer.From,
		Signer:   l.signer.Signer,
		GasLimit: 0,
		Context:  nil,
	}
	_, err := l.contract.StartEpoch(opts)
	if err != nil {
		log.Fatalf("Failed to start epoch: %v", err)
	}

	l.SyncState()
}

func (l *Lara) EndEpoch() {
	opts := &bind.TransactOpts{
		From:     l.signer.From,
		Signer:   l.signer.Signer,
		GasLimit: 0,
		Context:  nil,
	}
	_, err := l.contract.EndEpoch(opts)
	if err != nil {
		log.Fatalf("Failed to end epoch: %v", err)
	}

	l.SyncState()
}

func (l *Lara) GetState() State {
	return l.state
}

func (l *Lara) Evaluate() {
	callOpts := &bind.CallOpts{
		Pending:     false,
		From:        l.signer.From,
		BlockNumber: nil,
		Context:     nil,
	}
	opts := &bind.TransactOpts{
		From:     l.signer.From,
		Signer:   l.signer.Signer,
		GasLimit: 0,
		Context:  nil,
	}

	// get the current validators from oracle
	nodeCount, err := l.oracle.NodeCount(nil)
	if err != nil {
		log.Fatalf("Failed to get node count: %v", err)
	}
	for i := uint64(0); i < nodeCount.Uint64(); i++ {
		nodeAddress, err := l.oracle.NodesList(callOpts, big.NewInt(int64(i)))
		if err != nil {
			log.Fatalf("Failed to get node address: %v", err)
		}
		node, err := l.oracle.Nodes(nil, nodeAddress)
		if err != nil {
			log.Fatalf("Failed to get node: %v", err)
		}
		// check if node is already in validators
		found := false
		for _, validator := range l.state.validators {
			if validator.Account == node.Account {
				found = true
				break
			}
		}

		// if found
	}
}
