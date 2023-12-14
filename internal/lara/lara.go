package lara

import (
	"context"
	"math/big"
	"strings"
	"time"

	dpos_contract "github.com/Taraxa-project/taraxa-indexer/abi/dpos"
	lara_contract "github.com/Taraxa-project/taraxa-indexer/abi/lara"
	apy_oracle "github.com/Taraxa-project/taraxa-indexer/abi/oracle"
	"github.com/Taraxa-project/taraxa-indexer/internal/oracle"
	"github.com/Taraxa-project/taraxa-indexer/internal/transact"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

type State struct {
	epochDuration                 *big.Int
	lastEpochStartBlock           *big.Int
	isEpochRunning                bool
	lastEpochTotalDelegatedAmount *big.Int
	validatorStakes               map[common.Address]*big.Int
	validators                    []oracle.NodeData
	canRebalance                  bool
	isStartingEpoch               bool
	isEndingEpoch                 bool
	isRebalancing                 bool
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
	l.SyncState()
	return l
}

func (l *Lara) Run() {
	if l.Eth == nil {
		log.Fatalf("Eth client is nil")
	}
	ticker := time.NewTicker(3 * time.Second)
	for range ticker.C {
		ctx := context.Background()
		currentBlock, err := l.Eth.BlockNumber(ctx)
		if err != nil {
			log.Fatalf("Failed to get current block: %v", err)
		}
		if err != nil {
			log.Fatalf("Failed to get block by number: %v", err)
		}
		// if we pass the time to end epoch
		expectedEpochEnd := l.state.lastEpochStartBlock.Int64() + l.state.epochDuration.Int64()
		l.SyncState()
		if int64(currentBlock) > expectedEpochEnd {
			// if the epoch is running
			if l.state.isEpochRunning {
				// end the epoch
				l.EndEpoch()
				l.state.canRebalance = true
				l.Rebalance()
				// wait 3 sec
				time.Sleep(3 * time.Second)
			} else {
				// start the epoch
				l.StartEpoch()
				l.state.canRebalance = false
				// wait 3 sec
				time.Sleep(3 * time.Second)
			}
		}
	}

}

func (l *Lara) SyncState() {
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

	epochStartBlock, err := l.contract.LastEpochStartBlock(opts)
	if err != nil {
		log.Fatalf("Failed to get epoch start block: %v", err)
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
			if strings.Contains(err.Error(), "reverted") {
				break
			} else {
				log.Fatalf("Failed to get validator: %v", err)
				break
			}
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
		pos.Add(pos, big.NewInt(1))
	}

	//set commission to 10%
	commission, err := l.contract.Commission(opts)
	if err != nil {
		log.Fatalf("Failed to get commission: %v", err)
	}
	if commission.Cmp(big.NewInt(2)) != 0 {
		_, err = l.contract.SetCommission(l.signer, big.NewInt(2))

		if err != nil && !strings.Contains(err.Error(), "Transaction already in transactions pool") {
			log.Fatalf("Failed to set commission: %v", err)
		}
		// wait 1 sec
		time.Sleep(1 * time.Second)
	}
	l.state = State{
		epochDuration:                 epochDuration,
		lastEpochStartBlock:           epochStartBlock,
		isEpochRunning:                isEpochRunning,
		lastEpochTotalDelegatedAmount: lastEpochTotalDelegatedAmount,
		validatorStakes:               validatorStakes,
		validators:                    validators,
	}
	epochEndBlock := big.NewInt(0).Add(epochStartBlock, epochDuration)
	currentBlock, err := l.Eth.BlockNumber(context.Background())
	if err != nil {
		log.Fatalf("Failed to get current block: %v", err)
	}
	log.WithFields(log.Fields{"isRunning": l.state.isEpochRunning, "currentBlock": currentBlock, "epochStartBlock": l.state.lastEpochStartBlock, "epochEndBlock": epochEndBlock, "nodesDelegatedTo": len(l.state.validators), "totalDelegated": l.state.lastEpochTotalDelegatedAmount}).Info("LARA STATE: ")
}

func (l *Lara) StartEpoch() {
	if l.state.isStartingEpoch {
		log.Warn("WARN: PENDING START EPOCH")
		return
	}
	opts := &bind.TransactOpts{
		From:     l.signer.From,
		Signer:   l.signer.Signer,
		GasLimit: 0,
		Context:  nil,
	}
	oracleNodeCount, err := l.oracle.NodeCount(nil)
	if err != nil {
		log.Fatalf("Failed to get oracle node count: %v", err)
	}
	if oracleNodeCount.Cmp(big.NewInt(0)) == 0 {
		log.Warn("LARA == No oracle nodes")
		return
	}
	l.state.isStartingEpoch = true
	tx, err := l.contract.StartEpoch(opts)
	if err != nil {
		if strings.Contains(err.Error(), "Transaction already in transactions pool") {
			log.Warn("Start epoch tx already in pool")
		} else {
			log.Fatalf("Failed to start epoch: %v", err)
		}
	}
	if tx != nil {
		log.Warnf("Started epoch at timestamp: %d", tx.Time().Unix())
		l.state.isStartingEpoch = false
	}
	// wait 3 sec
	time.Sleep(3 * time.Second)
	l.SyncState()
	log.Warnf("Started epoch: %s", l.state.lastEpochStartBlock)
}

func (l *Lara) EndEpoch() {
	if l.state.isEndingEpoch {
		log.Warn("WARN: PENDING END EPOCH")
		return
	}
	opts := &bind.TransactOpts{
		From:     l.signer.From,
		Signer:   l.signer.Signer,
		GasLimit: 0,
		Context:  nil,
	}
	l.state.isEndingEpoch = true
	tx, err := l.contract.EndEpoch(opts)
	if err != nil {
		if strings.Contains(err.Error(), "Transaction already in transactions pool") {
			log.Warn("End epoch tx already in pool")
		} else {
			log.Fatalf("Failed to end epoch: %v", err)
		}
	}
	if tx != nil {
		l.state.isEndingEpoch = false
		log.Warnf("Ended epoch at timestamp: %d", tx.Time().Unix())
	}
	// wait one block
	time.Sleep(3 * time.Second)
	l.SyncState()
}

func (l *Lara) GetState() State {
	return l.state
}

func (l *Lara) Rebalance() {
	if l.state.canRebalance || l.state.isRebalancing {
		log.Warn("WARN: PENDING REBALANCE")
		return
	}
	opts := &bind.TransactOpts{
		From:     l.signer.From,
		Signer:   l.signer.Signer,
		GasLimit: 0,
		Context:  nil,
	}
	l.state.isRebalancing = true
	tx, err := l.contract.Rebalance(opts)
	if err != nil {
		if strings.Contains(err.Error(), "Transaction already in transactions pool") {
			log.Warn("Rebalance tx already in pool")
		} else {
			log.Fatalf("Failed to rebalance: %v", err)
		}
	}
	if tx != nil {
		l.state.isRebalancing = false
		log.Warnf("Rebalanced at timestamp: %d", tx.Time().Unix())
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
