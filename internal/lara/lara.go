package lara

import (
	"context"
	"fmt"
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
	lastSnapshotBlock             *big.Int
	lastSnapshotID                *big.Int
	lastSnapshotIdDistributed     *big.Int
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
	graphQLEndpoint   string
}

func MakeLara(rpc *ethclient.Client, signing_key, deployment_address, oracle_address, graphQLEndpoint string, chainID int) *Lara {
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
	l.graphQLEndpoint = graphQLEndpoint
	l.SyncState()
	l.FetchAndDistributePastRewards()
	l.ScanAndDistributeRewards()
	return l
}

func (l *Lara) Run() {
	if l.Eth == nil {
		log.Fatalf("Eth client is nil")
	}
	ticker := time.NewTicker(10 * time.Second)
	for range ticker.C {
		ctx := context.Background()
		currentBlock, err := l.Eth.BlockNumber(ctx)
		if err != nil {
			log.Fatalf("Lara: Failed to get current block: %v", err)
		}
		// if we pass the time to end epoch
		expenctedSnapshotTime := l.state.lastSnapshotBlock.Int64() + l.state.epochDuration.Int64()
		expectedRebalanceTime := l.state.lastRebalance.Int64() + l.state.epochDuration.Int64()
		log.WithFields(log.Fields{"expectedSnapshotTime": expenctedSnapshotTime, "expectedRebalanceTime": expectedRebalanceTime, "currentBlock": currentBlock}).Info("LARA: ")
		l.SyncState()

		if int64(currentBlock) > expenctedSnapshotTime {
			// if the epoch is running
			// end the epoch
			l.Snapshot()
			// wait 3 sec
			time.Sleep(4 * time.Second)

			l.Compound()
		}
		if int64(currentBlock) > expectedRebalanceTime {
			log.Warnf("Triggering rebalance at block: %d, expected rebalance time: %d", currentBlock, expectedRebalanceTime)
			l.Rebalance()
		}
	}
}

func (l *Lara) IsSnapshotDistributedToUser(snapshotId *big.Int, userAddress common.Address) bool {
	opts := &bind.CallOpts{
		Pending:     false,
		From:        l.signer.From,
		BlockNumber: nil,
		Context:     nil,
	}
	laraDistributedAlready, err := l.contract.StakerSnapshotClaimed(opts, userAddress, snapshotId)
	if err != nil {
		log.Fatalf("Failed to get staker snapshot claimed: %v", err)
	}
	return laraDistributedAlready
}

func (l *Lara) GetRewardsPerSnapshot(snapshotId *big.Int) *big.Int {
	opts := &bind.CallOpts{
		Pending:     false,
		From:        l.signer.From,
		BlockNumber: nil,
		Context:     nil,
	}
	rewards, err := l.contract.RewardsPerSnapshot(opts, snapshotId)
	if err != nil {
		log.Fatalf("Failed to get rewards per snapshot: %v", err)
	}
	return rewards
}

func (l *Lara) DisburseRewardsBetweenHolders(snapshotId *big.Int) {
	rewards := l.GetRewardsPerSnapshot(snapshotId)
	if rewards.Cmp(big.NewInt(0)) == 0 {
		l.state.lastSnapshotIdDistributed = snapshotId
		log.WithFields(log.Fields{"snapshotID": snapshotId}).Info("LARA: No rewards to distribute")
		return
	}
	blockNumber, err := l.GetLastSnapshotIDUpdateTime(snapshotId)
	if err != nil {
		log.Fatalf("Failed to get block number: %v", err)
	}
	log.WithFields(log.Fields{"blockNumber": blockNumber, "snapshotID": snapshotId}).Info("LARA: Getting staked tara holders")
	holders := GetStakedTaraHolders(l, blockNumber)

	log.WithFields(log.Fields{"# of holders": len(holders), "snapshotID": snapshotId}).Info("LARA: Disbursing rewards to holders for snapshot")

	opts := &bind.TransactOpts{
		From:     l.signer.From,
		Signer:   l.signer.Signer,
		GasLimit: 0,
		Context:  nil,
	}

	for _, holder := range holders {
		holderAddress := common.HexToAddress(holder)

		isDistributed := l.IsSnapshotDistributedToUser(snapshotId, holderAddress)
		if isDistributed {
			log.WithFields(log.Fields{"holder": holder, "snapshotID": snapshotId}).Info("LARA: Snapshot already distributed to holder")
			continue
		}

		tx, err := l.contract.DistrbuteRewardsForSnapshot(opts, holderAddress, snapshotId)
		if err != nil {
			if strings.Contains(err.Error(), "Transaction already in transactions pool") {
				log.Warn("Disburse tx already in pool")
			} else {
				log.Fatalf("Failed to disburse rewards for snapshot: %v with address: %s and snapshotID: %s", err, holderAddress.Hex(), snapshotId.String())
			}
		}
		log.WithFields(log.Fields{"txhash": tx.Hash().Hex(), "holder": holder, "snapshotID": snapshotId}).Info("LARA: Disbursed rewards to holder")
	}
}

func (l *Lara) Compound() {
	opts := &bind.TransactOpts{
		From:     l.signer.From,
		Signer:   l.signer.Signer,
		GasLimit: 0,
		Context:  nil,
	}
	laraEthBalance, err := l.Eth.BalanceAt(context.Background(), common.HexToAddress(l.deploymentAddress), nil)
	if err != nil {
		log.Fatalf("Failed to get lara eth balance: %v", err)
	}
	tx, err := l.contract.Compound(opts, laraEthBalance)
	if err != nil {
		if strings.Contains(err.Error(), "Transaction already in transactions pool") {
			log.Warn("Compound tx already in pool")
		} else {
			if strings.Contains(err.Error(), "No nodes available for delegation") {
				log.Warn("No nodes available for delegation")
			} else {
				log.Fatalf("Failed to compound: %v", err)
			}
		}
	} else {
		log.WithFields(log.Fields{"compoundedTaraAmount": laraEthBalance, "txhash": tx.Hash().Hex()}).Info("LARA COMPOUNDED: ")
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

	lastSnapshotBlock, err := l.contract.LastSnapshotBlock(opts)
	if err != nil {
		log.Fatalf("Failed to get last snapshot: %v", err)
	}

	lastSnapshotID, err := l.contract.LastSnapshotId(opts)
	if err != nil {
		log.Fatalf("Failed to get last snapshot ID: %v", err)
	}

	lastRebalanceBlock, err := l.contract.LastRebalance(opts)
	if err != nil {
		log.Fatalf("Failed to get last rebalance block: %v", err)
	}

	lastEpochTotalDelegatedAmount, err := l.dpos.GetTotalDelegation(opts, common.HexToAddress(l.deploymentAddress))
	if err != nil {
		log.Fatalf("Failed to get last epoch total delegated amount: %v", err)
	}

	// fetch validators from 0 until reverted and put in map
	pos := big.NewInt(0)
	validatorStakes := make(map[common.Address]*big.Int)
	validators := make([]oracle.NodeData, 0)
	for {
		validatorsFromDpos, err := l.dpos.GetDelegations(opts, common.HexToAddress(l.deploymentAddress), uint32(pos.Uint64()))
		if err != nil {
			if strings.Contains(err.Error(), "reverted") {
				break
			} else {
				log.Fatalf("Failed to get validator: %v", err)
				break
			}
		}
		for _, delegation := range validatorsFromDpos.Delegations {
			validator := delegation.Account
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
		if !validatorsFromDpos.End {
			pos.Add(pos, big.NewInt(1))
		} else {
			break
		}
	}

	//set commission to 8%
	commission, err := l.contract.Commission(opts)
	if err != nil {
		log.Fatalf("Failed to get commission: %v", err)
	}
	if commission.Cmp(big.NewInt(8)) != 0 {
		_, err = l.contract.SetCommission(l.signer, big.NewInt(8))

		if err != nil && !strings.Contains(err.Error(), "Transaction already in transactions pool") {
			log.Fatalf("Failed to set commission: %v", err)
		}
		// wait 1 sec
		time.Sleep(1 * time.Second)
	}
	l.state = State{
		epochDuration:                 epochDuration,
		lastSnapshotBlock:             lastSnapshotBlock,
		lastSnapshotID:                lastSnapshotID,
		lastRebalance:                 lastRebalanceBlock,
		lastEpochTotalDelegatedAmount: lastEpochTotalDelegatedAmount,
		validatorStakes:               validatorStakes,
		validators:                    validators,
		lastSnapshotIdDistributed:     big.NewInt(1),
	}
	nextSnapshot := big.NewInt(0).Add(lastSnapshotBlock, epochDuration)
	currentBlock, err := l.Eth.BlockNumber(context.Background())
	if err != nil {
		log.Fatalf("SyncState: Failed to get current block: %v", err)
	}
	log.WithFields(log.Fields{"currentBlock": currentBlock, "lastRebalance": l.state.lastRebalance, "lastSnapshotBlock": l.state.lastSnapshotBlock, "nextSnapshotBlock": nextSnapshot, "nodesDelegatedTo": len(l.state.validators), "totalDelegated": l.state.lastEpochTotalDelegatedAmount}).Info("LARA STATE: ")
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
			log.Warn("Snapshot tx already in pool")
		} else if strings.Contains(err.Error(), "EpochDurationNotMet") {
			log.Warn("Epoch duration not met")
		} else {
			log.Warnf("Failed to make snapshot: %v", err)
		}
	}
	// wait 4 secs ~ 1 block
	time.Sleep(4 * time.Second)

	log.Warnf("Snapshot hash: %s", tx.Hash().Hex())

	l.state.isMakingSnapshot = false
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
	if err != nil {
		if strings.Contains(err.Error(), "Transaction already in transactions pool") {
			log.Warn("Rebalance tx already in pool")
		} else {
			log.Warnf("Failed to rebalance: %v", err)
		}
	}
	if tx != nil {
		l.state.isRebalancing = false
		log.WithFields(log.Fields{"Timestamp": tx.Time().Unix(), "hash": tx.Hash().Hex()}).Warn("Rebalanced")
	}
}

func (l *Lara) GetLastSnapshotIDUpdateTime(snapshotID *big.Int) (uint64, error) {
	iter, err := l.contract.FilterSnapshotTaken(&bind.FilterOpts{}, []*big.Int{snapshotID}, nil, nil)
	if err != nil {
		return 0, err
	}
	defer iter.Close()

	if iter.Next() {
		return iter.Event.Raw.BlockNumber, nil
	}

	return 0, fmt.Errorf("no SnapshotTaken event found for snapshotID %s", snapshotID.String())
}

func (l *Lara) ScanAndDistributeRewards() {
	go func() {
		logs := make(chan *lara_contract.LaraContractSnapshotTaken)
		sub, err := l.contract.WatchSnapshotTaken(nil, logs, nil, nil, nil)
		if err != nil {
			log.Fatalf("Failed to watch SnapshotTaken events: %v", err)
		}
		defer sub.Unsubscribe()

		log.Info("Successfully started subscribing to SnapshotTaken events")

		for {
			select {
			case err := <-sub.Err():
				log.Errorf("Error in SnapshotTaken subscription: %v", err)
				return
			case event := <-logs:
				log.Infof("New SnapshotTaken event: SnapshotID %s", event.SnapshotId)
				l.distributeRewardsUpToSnapshot(event.SnapshotId)
			}
		}
	}()
}

func (l *Lara) distributeRewardsUpToSnapshot(latestSnapshotId *big.Int) {
	for snapshotId := new(big.Int).Add(l.state.lastSnapshotIdDistributed, big.NewInt(1)); snapshotId.Cmp(latestSnapshotId) <= 0; snapshotId.Add(snapshotId, big.NewInt(1)) {
		log.Infof("Checking rewards distribution for SnapshotID %s", snapshotId)

		l.distributeRewardsForSnapshot(snapshotId)
		l.state.lastSnapshotIdDistributed = snapshotId
	}
}

func (l *Lara) FetchAndDistributePastRewards() {
	go func() {
		log.Info("Starting to fetch and distribute past rewards")

		// Get the latest block number
		latestBlock, err := l.Eth.BlockNumber(context.Background())
		if err != nil {
			log.Errorf("Failed to get latest block number: %v", err)
			return
		}

		// Create a filter for SnapshotTaken events from block 0 to the latest block
		filterOpts := &bind.FilterOpts{
			Start:   0,
			End:     &latestBlock,
			Context: context.Background(),
		}

		// Filter for SnapshotTaken events
		iter, err := l.contract.FilterSnapshotTaken(filterOpts, nil, nil, nil)
		if err != nil {
			log.Errorf("Failed to filter SnapshotTaken events: %v", err)
			return
		}
		defer iter.Close()

		for iter.Next() {
			event := iter.Event
			log.Infof("Processing past SnapshotTaken event: SnapshotID %s", event.SnapshotId)

			// Check and distribute rewards for this snapshot
			l.distributeRewardsUpToSnapshot(event.SnapshotId)
		}

		if err := iter.Error(); err != nil {
			log.Errorf("Error iterating through SnapshotTaken events: %v", err)
		}

		log.Info("Finished processing past SnapshotTaken events")
	}()
}

func (l *Lara) distributeRewardsForSnapshot(snapshotId *big.Int) {
	log.Infof("Checking rewards distribution for SnapshotID %s", snapshotId)

	blockNumber, err := l.GetLastSnapshotIDUpdateTime(snapshotId)
	if err != nil {
		log.Errorf("Failed to get block number for SnapshotID %s: %v", snapshotId, err)
		return
	}

	holders := GetStakedTaraHolders(l, blockNumber)
	rewardsDistributed := false

	for _, holder := range holders {
		holderAddress := common.HexToAddress(holder)
		isDistributed := l.IsSnapshotDistributedToUser(snapshotId, holderAddress)
		if !isDistributed {
			log.Infof("Distributing rewards for SnapshotID %s to holder %s", snapshotId, holder)
			l.DisburseRewardsBetweenHolders(snapshotId)
			rewardsDistributed = true
			break
		}
	}

	if !rewardsDistributed {
		log.Infof("All rewards already distributed for SnapshotID %s", snapshotId)
	}
}
