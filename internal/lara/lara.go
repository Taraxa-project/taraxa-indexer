package lara

import (
	"math/big"
	"strings"

	dpos_contract "github.com/Taraxa-project/taraxa-indexer/abi/dpos"
	lara_contract "github.com/Taraxa-project/taraxa-indexer/abi/lara"
	apy_oracle "github.com/Taraxa-project/taraxa-indexer/abi/oracle"
	"github.com/Taraxa-project/taraxa-indexer/internal/contracts"
	"github.com/Taraxa-project/taraxa-indexer/internal/oracle"
	"github.com/Taraxa-project/taraxa-indexer/internal/transact"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
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
	dpos              *dpos_contract.DposContract
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
	l.oracle, err = apy_oracle.NewApyOracle(common.HexToAddress(l.oracleAddress), l.Eth)
	if err != nil {
		log.Fatalf("Failed to create oracle: %v", err)
	}
	l.dpos, err = dpos_contract.NewDposContract(common.HexToAddress("0x00000000000000000000000000000000000000fe"), l.Eth)
	if err != nil {
		log.Fatalf("Failed to create dpos: %v", err)
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

func (l *Lara) Evaluate(newValidators []oracle.NodeData) {
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

	log.Infof("Evaluating new validators: %d", len(newValidators))

	// we need to go through lara's validators that have stake
	for address, stake := range l.state.validatorStakes {
		log.Infof("Evaluating validator: %s", address.Hex())
		node, err := l.oracle.Nodes(nil, address)
		if err != nil {
			log.Fatalf("Failed to get node: %v", err)
		}

		// compare on-chain data with new score in state
		newValidatorData := findNode(newValidators, node)
		if newValidatorData.Account.Hex() != "" {
			// if score is smaller with at least 10% compared to on-chain score
			if newValidatorData.Rating.Cmp(node.Rating) == -1 && newValidatorData.Rating.Cmp(node.Rating.Div(node.Rating, big.NewInt(10))) == -1 {
				// we need to redelegate to the higest score new node, which is the first one in the list
				for _, validator := range newValidators {
					// check if the node has enough stake room
					info, err := l.dpos.GetValidator(callOpts, validator.Account)
					if err != nil {
						log.Fatalf("Failed to get validator info: %v", err)
					}
					// if there's a node that can fit & is still 10% better than the current node
					if info.TotalStake.Cmp(stake) >= 0 && validator.Rating.Cmp(node.Rating.Div(node.Rating, big.NewInt(10))) == 1 {
						// redelegate
						log.Infof("Redelegating %s from %s to %s", stake, address.Hex(), validator.Account.Hex())
						_, err := l.contract.ReDelegate(opts, address, validator.Account, stake)
						if err != nil {
							log.Fatalf("Failed to delegate: %v", err)
						}
					} else {
						// we do this until we find a node with enough stake room
						continue
					}
				}
			}
		}
	}
}

func findNode(nodes []oracle.NodeData, node oracle.NodeData) oracle.NodeData {
	for _, n := range nodes {
		if n.Account == node.Account {
			return n
		}
	}
	return oracle.NodeData{}
}
