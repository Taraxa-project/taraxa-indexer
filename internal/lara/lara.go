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
	protocolStartTimestamp        *big.Int
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
	l.SyncState()
	return l
}

func (l *Lara) Run() {
	if l.Eth == nil {
		log.Fatalf("Eth client is nil")
	}
	ticker := time.NewTicker(1 * time.Second)
	lastEpochStartTimestamp := int64(0)
	for range ticker.C {
		log.Infof("Last epoch start timestamp: %d", lastEpochStartTimestamp)
		ctx := context.Background()
		currentBlock, err := l.Eth.BlockNumber(ctx)
		if err != nil {
			log.Fatalf("Failed to get current block: %v", err)
		}
		blockByNumber, err := l.Eth.BlockByNumber(ctx, new(big.Int).SetUint64(currentBlock))
		if err != nil {
			log.Fatalf("Failed to get block by number: %v", err)
		}
		// if we pass the time to end epoch
		log.Infof("Current block time: %d", blockByNumber.Time())
		log.Warnf("Ending epoch at %d", lastEpochStartTimestamp+(int64(4)*l.state.epochDuration.Int64()))
		if lastEpochStartTimestamp == 0 || int64(blockByNumber.Time()) > lastEpochStartTimestamp {
			l.SyncState()
			// if the epoch is running
			if l.state.isEpochRunning {
				// end the epoch
				l.EndEpoch()
				// wait 3 sec
				time.Sleep(3 * time.Second)
			} else {
				// start the epoch
				l.StartEpoch()
				// wait 3 sec
				time.Sleep(3 * time.Second)
			}
			if lastEpochStartTimestamp == 0 {
				lastEpochStartTimestamp = l.state.protocolStartTimestamp.Int64()
			}
			timeToEndEpoch := lastEpochStartTimestamp + (int64(4) * l.state.epochDuration.Int64())
			lastEpochStartTimestamp = timeToEndEpoch
			log.Infof("Calculated new Time to end epoch: %d", timeToEndEpoch)
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

	epochStartTimestamp, err := l.contract.ProtocolStartTimestamp(opts)
	if err != nil {
		log.Fatalf("Failed to get epoch start timestamp: %v", err)
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
		log.Infof("Fetching validator at pos: %d", pos)
		validator, err := l.contract.Validators(opts, pos)
		if err != nil {
			if strings.Contains(err.Error(), "reverted") {
				log.Infof("Reached end of validators")
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
	if commission.Cmp(big.NewInt(10)) != 0 {
		_, err = l.contract.SetCommission(l.signer, big.NewInt(2))

		if err != nil && !strings.Contains(err.Error(), "Transaction already in transactions pool") {
			log.Fatalf("Failed to set commission: %v", err)
		}
		// wait 1 sec
		time.Sleep(1 * time.Second)
	}
	l.state = State{
		epochDuration:                 epochDuration,
		protocolStartTimestamp:        epochStartTimestamp,
		isEpochRunning:                isEpochRunning,
		lastEpochTotalDelegatedAmount: lastEpochTotalDelegatedAmount,
		validatorStakes:               validatorStakes,
		validators:                    validators,
	}
}

func (l *Lara) StartEpoch() {
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
	}
	// wait 3 sec
	time.Sleep(3 * time.Second)
	l.SyncState()
	log.Infof("Started epoch: %s", l.state.protocolStartTimestamp)
}

func (l *Lara) EndEpoch() {
	opts := &bind.TransactOpts{
		From:     l.signer.From,
		Signer:   l.signer.Signer,
		GasLimit: 0,
		Context:  nil,
	}
	tx, err := l.contract.EndEpoch(opts)
	if err != nil {
		if strings.Contains(err.Error(), "Transaction already in transactions pool") {
			log.Warn("End epoch tx already in pool")
		} else {
			log.Fatalf("Failed to end epoch: %v", err)
		}
	}
	if tx != nil {
		log.Warnf("Ended epoch at timestamp: %d", tx.Time().Unix())
	}
	l.SyncState()

	// can be solved by indexing a separate event for this and taking the validator info at the event height
	// find the delegators of Lara and divide & disburse them if these amounts are too big
	l.Rebalance()
	l.SyncState()
}

func (l *Lara) GetState() State {
	return l.state
}

func (l *Lara) Rebalance() {
	tx, err := l.contract.Rebalance(nil)
	if err != nil {
		if strings.Contains(err.Error(), "Transaction already in transactions pool") {
			log.Warn("Rebalance tx already in pool")
		} else {
			log.Fatalf("Failed to rebalance: %v", err)
		}
	}
	if tx != nil {
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
