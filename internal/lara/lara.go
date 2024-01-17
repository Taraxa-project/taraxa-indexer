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
	lastSnapshot                  *big.Int
	lastRebalance                 *big.Int
	lastEpochTotalDelegatedAmount *big.Int
	validatorStakes               map[common.Address]*big.Int
	validators                    []oracle.NodeData
	isMakingSnapshot              bool
	isRebalancing                 bool
}
type Lara struct {
	deploymentAddress string
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
	l.chainID = &chainID
	contract, err := lara_contract.NewLaraContract(common.HexToAddress(l.deploymentAddress), l.Eth)
	if err != nil {
		log.Fatalf("Failed to create contract: %v", err)
	}
	l.oracle, err = apy_oracle.NewApyOracle(common.HexToAddress(oracle_address), l.Eth)
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
	ticker := time.NewTicker(1 * time.Minute)
	for range ticker.C {
		ctx := context.Background()
		currentBlock, err := l.Eth.BlockNumber(ctx)
		if err != nil {
			log.Fatalf("Lara: Failed to get current block: %v", err)
		}
		// if we pass the time to end epoch
		expenctedSnapshotTime := l.state.lastSnapshot.Int64() + l.state.epochDuration.Int64()
		l.SyncState()
		if int64(currentBlock) > expenctedSnapshotTime {
			// if the epoch is running
			// end the epoch
			l.Snapshot()
			time.Sleep(4 * time.Second)
			l.Rebalance()
			// wait 3 sec
			time.Sleep(4 * time.Second)
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

	lastSnapshotBlock, err := l.contract.LastSnapshot(opts)
	if err != nil {
		log.Fatalf("Failed to get last snapshot: %v", err)
	}

	lastRebalanceBlock, err := l.contract.LastRebalance(opts)
	if err != nil {
		log.Fatalf("Failed to get last rebalance block: %v", err)
	}

	lastEpochTotalDelegatedAmount, err := l.contract.TotalDelegated(opts)
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
		lastSnapshot:                  lastSnapshotBlock,
		lastRebalance:                 lastRebalanceBlock,
		lastEpochTotalDelegatedAmount: lastEpochTotalDelegatedAmount,
		validatorStakes:               validatorStakes,
		validators:                    validators,
	}
	nextSnapshot := big.NewInt(0).Add(lastSnapshotBlock, epochDuration)
	currentBlock, err := l.Eth.BlockNumber(context.Background())
	if err != nil {
		log.Fatalf("SyncState: Failed to get current block: %v", err)
	}
	log.WithFields(log.Fields{"currentBlock": currentBlock, "lastRebalance": l.state.lastRebalance, "lastSnapshotBlock": l.state.lastSnapshot, "nextSnapshotBlock": nextSnapshot, "nodesDelegatedTo": len(l.state.validators), "totalDelegated": l.state.lastEpochTotalDelegatedAmount}).Info("LARA STATE: ")
}

func (l *Lara) Snapshot() {
	if l.state.isMakingSnapshot {
		log.Warn("WARN: PENDING SNPAHSOT")
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
	l.state.isMakingSnapshot = true
	tx, err := l.contract.Snapshot(opts)
	timer := time.NewTimer(5 * time.Second)
	if err != nil {
		if strings.Contains(err.Error(), "Transaction already in transactions pool") {
		} else {
			log.Fatalf("Failed to make snapshot: %v", err)
		}
	}
	// wait 4 secs ~ 1 block
	time.Sleep(4 * time.Second)

	if tx != nil {
		receipt, err := l.Eth.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				log.Warn("WARN: SNAPSHOT NOT FOUND")
			} else {
				log.Fatalf("Failed to get receipt: %v", err)
			}
		} else {
			log.Warnf("Made snapshot at timestamp: %d, hash: %s", receipt.BlockNumber, tx.Hash().Hex())
		}
		l.state.isMakingSnapshot = false
	}
	// wait 3 sec
	time.Sleep(3 * time.Second)
	l.SyncState()
	time.Sleep(3 * time.Second)
	if l.state.isMakingSnapshot {
		log.Warnf("Snapshot not made after : %d", timer.C)
		timer.Stop()
	}
}

func (l *Lara) GetState() State {
	return l.state
}

func (l *Lara) Rebalance() {
	if l.state.isRebalancing {
		log.Warn("WARN: PENDING REBALANCE")
		return
	}
	if l.state.isMakingSnapshot {
		log.Warn("WARN: SNAPSHOT IN PROGRESS")
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
	log.Printf("rebalance tx: %s", tx.Hash().Hex())
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
