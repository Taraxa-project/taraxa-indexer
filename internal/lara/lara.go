package lara

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type State struct {
	epochDuration                 uint64
	isEpochRunning                bool
	lastEpochTotalDelegatedAmount *big.Int
	validatorStakesTotal          *big.Int
	validatorStakes               map[string]*big.Int
	validators                    []common.Address
}

type Lara struct {
	deploymentAddress string
	eth               *ethclient.Client
	signer            *bind.TransactOpts
	chainID           *int
	contract          *bind.BoundContract
	state             State
}

// func MakeLara(blockchain_ws, signing_key, deployment_address string, chainID *big.Int) *Lara {
// 	l := new(Lara)
// 	l.eth = connect(blockchain_ws)
// 	l.signer = makeSigner(signing_key, chainID)
// 	l.deploymentAddress = deployment_address
// 	l.chainID = chainID
// 	l.contract = l.makeContract()
// 	return l
// }
