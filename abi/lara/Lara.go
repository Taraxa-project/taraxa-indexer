// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lara_contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// DposInterfaceUndelegationData is an auto generated low-level Go binding around an user-defined struct.
type DposInterfaceUndelegationData struct {
	Stake           *big.Int
	Block           uint64
	Validator       common.Address
	ValidatorExists bool
}

// LaraContractMetaData contains all meta data concerning the LaraContract contract.
var LaraContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"apyOracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIApyOracle\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelUndelegate\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"commission\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"compound\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"confirmUndelegate\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegators\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"dposContract\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractDposInterface\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"epochDuration\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_sttaraToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_dposContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_apyOracle\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_treasuryAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isValidatorRegistered\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastRebalance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastSnapshot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxValidatorStakeCapacity\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minStakeAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"protocolStartTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"protocolTotalStakeAtValidator\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"protocolValidatorRatingAtDelegation\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rebalance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestUndelegate\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"undelegation_ids\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCommission\",\"inputs\":[{\"name\":\"_commission\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEpochDuration\",\"inputs\":[{\"name\":\"_epochDuration\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxValidatorStakeCapacity\",\"inputs\":[{\"name\":\"_maxValidatorStakeCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinStakeAmount\",\"inputs\":[{\"name\":\"_minStakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTreasuryAddress\",\"inputs\":[{\"name\":\"_treasuryAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"snapshot\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stTaraToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIstTara\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stake\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"totalDelegated\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"treasuryAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"undelegated\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"undelegations\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"undelegation_data\",\"type\":\"tuple\",\"internalType\":\"structDposInterface.UndelegationData\",\"components\":[{\"name\":\"stake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"block\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validator_exists\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"undelegation_id\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"validators\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AllRewardsClaimed\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CommissionChanged\",\"inputs\":[{\"name\":\"newCommission\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CommissionWithdrawn\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RedelegationRewardsClaimed\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SnapshotTaken\",\"inputs\":[{\"name\":\"totalDelegation\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"totalRewards\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"nextSnapshotBlock\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Staked\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaraSent\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TreasuryChanged\",\"inputs\":[{\"name\":\"newTreasury\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Undelegated\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UndelegationCancelled\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UndelegationConfirmed\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CancelUndelegationFailed\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reason\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"ConfirmUndelegationFailed\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reason\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"DelegationFailed\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reason\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EpochDurationNotMet\",\"inputs\":[{\"name\":\"lastEpochStart\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currentBlockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"epochDuration\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"RedelegationFailed\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reason\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RewardClaimFailed\",\"inputs\":[{\"name\":\"reason\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"StakeAmountTooLow\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"StakeValueTooLow\",\"inputs\":[{\"name\":\"sentAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"targetAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"UndelegationFailed\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"UndelegationNotFound\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"UndelegationsNotMatching\",\"inputs\":[{\"name\":\"undelegations\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
}

// LaraContractABI is the input ABI used to generate the binding from.
// Deprecated: Use LaraContractMetaData.ABI instead.
var LaraContractABI = LaraContractMetaData.ABI

// LaraContract is an auto generated Go binding around an Ethereum contract.
type LaraContract struct {
	LaraContractCaller     // Read-only binding to the contract
	LaraContractTransactor // Write-only binding to the contract
	LaraContractFilterer   // Log filterer for contract events
}

// LaraContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type LaraContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LaraContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LaraContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LaraContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LaraContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LaraContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LaraContractSession struct {
	Contract     *LaraContract     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LaraContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LaraContractCallerSession struct {
	Contract *LaraContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// LaraContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LaraContractTransactorSession struct {
	Contract     *LaraContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// LaraContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type LaraContractRaw struct {
	Contract *LaraContract // Generic contract binding to access the raw methods on
}

// LaraContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LaraContractCallerRaw struct {
	Contract *LaraContractCaller // Generic read-only contract binding to access the raw methods on
}

// LaraContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LaraContractTransactorRaw struct {
	Contract *LaraContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLaraContract creates a new instance of LaraContract, bound to a specific deployed contract.
func NewLaraContract(address common.Address, backend bind.ContractBackend) (*LaraContract, error) {
	contract, err := bindLaraContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LaraContract{LaraContractCaller: LaraContractCaller{contract: contract}, LaraContractTransactor: LaraContractTransactor{contract: contract}, LaraContractFilterer: LaraContractFilterer{contract: contract}}, nil
}

// NewLaraContractCaller creates a new read-only instance of LaraContract, bound to a specific deployed contract.
func NewLaraContractCaller(address common.Address, caller bind.ContractCaller) (*LaraContractCaller, error) {
	contract, err := bindLaraContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LaraContractCaller{contract: contract}, nil
}

// NewLaraContractTransactor creates a new write-only instance of LaraContract, bound to a specific deployed contract.
func NewLaraContractTransactor(address common.Address, transactor bind.ContractTransactor) (*LaraContractTransactor, error) {
	contract, err := bindLaraContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LaraContractTransactor{contract: contract}, nil
}

// NewLaraContractFilterer creates a new log filterer instance of LaraContract, bound to a specific deployed contract.
func NewLaraContractFilterer(address common.Address, filterer bind.ContractFilterer) (*LaraContractFilterer, error) {
	contract, err := bindLaraContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LaraContractFilterer{contract: contract}, nil
}

// bindLaraContract binds a generic wrapper to an already deployed contract.
func bindLaraContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LaraContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LaraContract *LaraContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LaraContract.Contract.LaraContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LaraContract *LaraContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LaraContract.Contract.LaraContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LaraContract *LaraContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LaraContract.Contract.LaraContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LaraContract *LaraContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LaraContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LaraContract *LaraContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LaraContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LaraContract *LaraContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LaraContract.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_LaraContract *LaraContractCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LaraContract.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_LaraContract *LaraContractSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _LaraContract.Contract.UPGRADEINTERFACEVERSION(&_LaraContract.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_LaraContract *LaraContractCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _LaraContract.Contract.UPGRADEINTERFACEVERSION(&_LaraContract.CallOpts)
}

// ApyOracle is a free data retrieval call binding the contract method 0x627ed636.
//
// Solidity: function apyOracle() view returns(address)
func (_LaraContract *LaraContractCaller) ApyOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LaraContract.contract.Call(opts, &out, "apyOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ApyOracle is a free data retrieval call binding the contract method 0x627ed636.
//
// Solidity: function apyOracle() view returns(address)
func (_LaraContract *LaraContractSession) ApyOracle() (common.Address, error) {
	return _LaraContract.Contract.ApyOracle(&_LaraContract.CallOpts)
}

// ApyOracle is a free data retrieval call binding the contract method 0x627ed636.
//
// Solidity: function apyOracle() view returns(address)
func (_LaraContract *LaraContractCallerSession) ApyOracle() (common.Address, error) {
	return _LaraContract.Contract.ApyOracle(&_LaraContract.CallOpts)
}

// Commission is a free data retrieval call binding the contract method 0xe1489191.
//
// Solidity: function commission() view returns(uint256)
func (_LaraContract *LaraContractCaller) Commission(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LaraContract.contract.Call(opts, &out, "commission")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Commission is a free data retrieval call binding the contract method 0xe1489191.
//
// Solidity: function commission() view returns(uint256)
func (_LaraContract *LaraContractSession) Commission() (*big.Int, error) {
	return _LaraContract.Contract.Commission(&_LaraContract.CallOpts)
}

// Commission is a free data retrieval call binding the contract method 0xe1489191.
//
// Solidity: function commission() view returns(uint256)
func (_LaraContract *LaraContractCallerSession) Commission() (*big.Int, error) {
	return _LaraContract.Contract.Commission(&_LaraContract.CallOpts)
}

// Delegators is a free data retrieval call binding the contract method 0x5be612c7.
//
// Solidity: function delegators(uint256 ) view returns(address)
func (_LaraContract *LaraContractCaller) Delegators(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LaraContract.contract.Call(opts, &out, "delegators", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegators is a free data retrieval call binding the contract method 0x5be612c7.
//
// Solidity: function delegators(uint256 ) view returns(address)
func (_LaraContract *LaraContractSession) Delegators(arg0 *big.Int) (common.Address, error) {
	return _LaraContract.Contract.Delegators(&_LaraContract.CallOpts, arg0)
}

// Delegators is a free data retrieval call binding the contract method 0x5be612c7.
//
// Solidity: function delegators(uint256 ) view returns(address)
func (_LaraContract *LaraContractCallerSession) Delegators(arg0 *big.Int) (common.Address, error) {
	return _LaraContract.Contract.Delegators(&_LaraContract.CallOpts, arg0)
}

// DposContract is a free data retrieval call binding the contract method 0xe1fb9ae2.
//
// Solidity: function dposContract() view returns(address)
func (_LaraContract *LaraContractCaller) DposContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LaraContract.contract.Call(opts, &out, "dposContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DposContract is a free data retrieval call binding the contract method 0xe1fb9ae2.
//
// Solidity: function dposContract() view returns(address)
func (_LaraContract *LaraContractSession) DposContract() (common.Address, error) {
	return _LaraContract.Contract.DposContract(&_LaraContract.CallOpts)
}

// DposContract is a free data retrieval call binding the contract method 0xe1fb9ae2.
//
// Solidity: function dposContract() view returns(address)
func (_LaraContract *LaraContractCallerSession) DposContract() (common.Address, error) {
	return _LaraContract.Contract.DposContract(&_LaraContract.CallOpts)
}

// EpochDuration is a free data retrieval call binding the contract method 0x4ff0876a.
//
// Solidity: function epochDuration() view returns(uint256)
func (_LaraContract *LaraContractCaller) EpochDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LaraContract.contract.Call(opts, &out, "epochDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochDuration is a free data retrieval call binding the contract method 0x4ff0876a.
//
// Solidity: function epochDuration() view returns(uint256)
func (_LaraContract *LaraContractSession) EpochDuration() (*big.Int, error) {
	return _LaraContract.Contract.EpochDuration(&_LaraContract.CallOpts)
}

// EpochDuration is a free data retrieval call binding the contract method 0x4ff0876a.
//
// Solidity: function epochDuration() view returns(uint256)
func (_LaraContract *LaraContractCallerSession) EpochDuration() (*big.Int, error) {
	return _LaraContract.Contract.EpochDuration(&_LaraContract.CallOpts)
}

// IsValidatorRegistered is a free data retrieval call binding the contract method 0xd04a68c7.
//
// Solidity: function isValidatorRegistered(address validator) view returns(bool)
func (_LaraContract *LaraContractCaller) IsValidatorRegistered(opts *bind.CallOpts, validator common.Address) (bool, error) {
	var out []interface{}
	err := _LaraContract.contract.Call(opts, &out, "isValidatorRegistered", validator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidatorRegistered is a free data retrieval call binding the contract method 0xd04a68c7.
//
// Solidity: function isValidatorRegistered(address validator) view returns(bool)
func (_LaraContract *LaraContractSession) IsValidatorRegistered(validator common.Address) (bool, error) {
	return _LaraContract.Contract.IsValidatorRegistered(&_LaraContract.CallOpts, validator)
}

// IsValidatorRegistered is a free data retrieval call binding the contract method 0xd04a68c7.
//
// Solidity: function isValidatorRegistered(address validator) view returns(bool)
func (_LaraContract *LaraContractCallerSession) IsValidatorRegistered(validator common.Address) (bool, error) {
	return _LaraContract.Contract.IsValidatorRegistered(&_LaraContract.CallOpts, validator)
}

// LastRebalance is a free data retrieval call binding the contract method 0x106b9ca1.
//
// Solidity: function lastRebalance() view returns(uint256)
func (_LaraContract *LaraContractCaller) LastRebalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LaraContract.contract.Call(opts, &out, "lastRebalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRebalance is a free data retrieval call binding the contract method 0x106b9ca1.
//
// Solidity: function lastRebalance() view returns(uint256)
func (_LaraContract *LaraContractSession) LastRebalance() (*big.Int, error) {
	return _LaraContract.Contract.LastRebalance(&_LaraContract.CallOpts)
}

// LastRebalance is a free data retrieval call binding the contract method 0x106b9ca1.
//
// Solidity: function lastRebalance() view returns(uint256)
func (_LaraContract *LaraContractCallerSession) LastRebalance() (*big.Int, error) {
	return _LaraContract.Contract.LastRebalance(&_LaraContract.CallOpts)
}

// LastSnapshot is a free data retrieval call binding the contract method 0xfb861ac1.
//
// Solidity: function lastSnapshot() view returns(uint256)
func (_LaraContract *LaraContractCaller) LastSnapshot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LaraContract.contract.Call(opts, &out, "lastSnapshot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastSnapshot is a free data retrieval call binding the contract method 0xfb861ac1.
//
// Solidity: function lastSnapshot() view returns(uint256)
func (_LaraContract *LaraContractSession) LastSnapshot() (*big.Int, error) {
	return _LaraContract.Contract.LastSnapshot(&_LaraContract.CallOpts)
}

// LastSnapshot is a free data retrieval call binding the contract method 0xfb861ac1.
//
// Solidity: function lastSnapshot() view returns(uint256)
func (_LaraContract *LaraContractCallerSession) LastSnapshot() (*big.Int, error) {
	return _LaraContract.Contract.LastSnapshot(&_LaraContract.CallOpts)
}

// MaxValidatorStakeCapacity is a free data retrieval call binding the contract method 0x2a8cf87f.
//
// Solidity: function maxValidatorStakeCapacity() view returns(uint256)
func (_LaraContract *LaraContractCaller) MaxValidatorStakeCapacity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LaraContract.contract.Call(opts, &out, "maxValidatorStakeCapacity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxValidatorStakeCapacity is a free data retrieval call binding the contract method 0x2a8cf87f.
//
// Solidity: function maxValidatorStakeCapacity() view returns(uint256)
func (_LaraContract *LaraContractSession) MaxValidatorStakeCapacity() (*big.Int, error) {
	return _LaraContract.Contract.MaxValidatorStakeCapacity(&_LaraContract.CallOpts)
}

// MaxValidatorStakeCapacity is a free data retrieval call binding the contract method 0x2a8cf87f.
//
// Solidity: function maxValidatorStakeCapacity() view returns(uint256)
func (_LaraContract *LaraContractCallerSession) MaxValidatorStakeCapacity() (*big.Int, error) {
	return _LaraContract.Contract.MaxValidatorStakeCapacity(&_LaraContract.CallOpts)
}

// MinStakeAmount is a free data retrieval call binding the contract method 0xf1887684.
//
// Solidity: function minStakeAmount() view returns(uint256)
func (_LaraContract *LaraContractCaller) MinStakeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LaraContract.contract.Call(opts, &out, "minStakeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStakeAmount is a free data retrieval call binding the contract method 0xf1887684.
//
// Solidity: function minStakeAmount() view returns(uint256)
func (_LaraContract *LaraContractSession) MinStakeAmount() (*big.Int, error) {
	return _LaraContract.Contract.MinStakeAmount(&_LaraContract.CallOpts)
}

// MinStakeAmount is a free data retrieval call binding the contract method 0xf1887684.
//
// Solidity: function minStakeAmount() view returns(uint256)
func (_LaraContract *LaraContractCallerSession) MinStakeAmount() (*big.Int, error) {
	return _LaraContract.Contract.MinStakeAmount(&_LaraContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LaraContract *LaraContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LaraContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LaraContract *LaraContractSession) Owner() (common.Address, error) {
	return _LaraContract.Contract.Owner(&_LaraContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LaraContract *LaraContractCallerSession) Owner() (common.Address, error) {
	return _LaraContract.Contract.Owner(&_LaraContract.CallOpts)
}

// ProtocolStartTimestamp is a free data retrieval call binding the contract method 0x64956417.
//
// Solidity: function protocolStartTimestamp() view returns(uint256)
func (_LaraContract *LaraContractCaller) ProtocolStartTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LaraContract.contract.Call(opts, &out, "protocolStartTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolStartTimestamp is a free data retrieval call binding the contract method 0x64956417.
//
// Solidity: function protocolStartTimestamp() view returns(uint256)
func (_LaraContract *LaraContractSession) ProtocolStartTimestamp() (*big.Int, error) {
	return _LaraContract.Contract.ProtocolStartTimestamp(&_LaraContract.CallOpts)
}

// ProtocolStartTimestamp is a free data retrieval call binding the contract method 0x64956417.
//
// Solidity: function protocolStartTimestamp() view returns(uint256)
func (_LaraContract *LaraContractCallerSession) ProtocolStartTimestamp() (*big.Int, error) {
	return _LaraContract.Contract.ProtocolStartTimestamp(&_LaraContract.CallOpts)
}

// ProtocolTotalStakeAtValidator is a free data retrieval call binding the contract method 0xf553d398.
//
// Solidity: function protocolTotalStakeAtValidator(address ) view returns(uint256)
func (_LaraContract *LaraContractCaller) ProtocolTotalStakeAtValidator(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LaraContract.contract.Call(opts, &out, "protocolTotalStakeAtValidator", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolTotalStakeAtValidator is a free data retrieval call binding the contract method 0xf553d398.
//
// Solidity: function protocolTotalStakeAtValidator(address ) view returns(uint256)
func (_LaraContract *LaraContractSession) ProtocolTotalStakeAtValidator(arg0 common.Address) (*big.Int, error) {
	return _LaraContract.Contract.ProtocolTotalStakeAtValidator(&_LaraContract.CallOpts, arg0)
}

// ProtocolTotalStakeAtValidator is a free data retrieval call binding the contract method 0xf553d398.
//
// Solidity: function protocolTotalStakeAtValidator(address ) view returns(uint256)
func (_LaraContract *LaraContractCallerSession) ProtocolTotalStakeAtValidator(arg0 common.Address) (*big.Int, error) {
	return _LaraContract.Contract.ProtocolTotalStakeAtValidator(&_LaraContract.CallOpts, arg0)
}

// ProtocolValidatorRatingAtDelegation is a free data retrieval call binding the contract method 0xddb63cc8.
//
// Solidity: function protocolValidatorRatingAtDelegation(address ) view returns(uint256)
func (_LaraContract *LaraContractCaller) ProtocolValidatorRatingAtDelegation(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LaraContract.contract.Call(opts, &out, "protocolValidatorRatingAtDelegation", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolValidatorRatingAtDelegation is a free data retrieval call binding the contract method 0xddb63cc8.
//
// Solidity: function protocolValidatorRatingAtDelegation(address ) view returns(uint256)
func (_LaraContract *LaraContractSession) ProtocolValidatorRatingAtDelegation(arg0 common.Address) (*big.Int, error) {
	return _LaraContract.Contract.ProtocolValidatorRatingAtDelegation(&_LaraContract.CallOpts, arg0)
}

// ProtocolValidatorRatingAtDelegation is a free data retrieval call binding the contract method 0xddb63cc8.
//
// Solidity: function protocolValidatorRatingAtDelegation(address ) view returns(uint256)
func (_LaraContract *LaraContractCallerSession) ProtocolValidatorRatingAtDelegation(arg0 common.Address) (*big.Int, error) {
	return _LaraContract.Contract.ProtocolValidatorRatingAtDelegation(&_LaraContract.CallOpts, arg0)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LaraContract *LaraContractCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LaraContract.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LaraContract *LaraContractSession) ProxiableUUID() ([32]byte, error) {
	return _LaraContract.Contract.ProxiableUUID(&_LaraContract.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LaraContract *LaraContractCallerSession) ProxiableUUID() ([32]byte, error) {
	return _LaraContract.Contract.ProxiableUUID(&_LaraContract.CallOpts)
}

// StTaraToken is a free data retrieval call binding the contract method 0x021b7a81.
//
// Solidity: function stTaraToken() view returns(address)
func (_LaraContract *LaraContractCaller) StTaraToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LaraContract.contract.Call(opts, &out, "stTaraToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StTaraToken is a free data retrieval call binding the contract method 0x021b7a81.
//
// Solidity: function stTaraToken() view returns(address)
func (_LaraContract *LaraContractSession) StTaraToken() (common.Address, error) {
	return _LaraContract.Contract.StTaraToken(&_LaraContract.CallOpts)
}

// StTaraToken is a free data retrieval call binding the contract method 0x021b7a81.
//
// Solidity: function stTaraToken() view returns(address)
func (_LaraContract *LaraContractCallerSession) StTaraToken() (common.Address, error) {
	return _LaraContract.Contract.StTaraToken(&_LaraContract.CallOpts)
}

// TotalDelegated is a free data retrieval call binding the contract method 0x80d04de8.
//
// Solidity: function totalDelegated() view returns(uint256)
func (_LaraContract *LaraContractCaller) TotalDelegated(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LaraContract.contract.Call(opts, &out, "totalDelegated")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDelegated is a free data retrieval call binding the contract method 0x80d04de8.
//
// Solidity: function totalDelegated() view returns(uint256)
func (_LaraContract *LaraContractSession) TotalDelegated() (*big.Int, error) {
	return _LaraContract.Contract.TotalDelegated(&_LaraContract.CallOpts)
}

// TotalDelegated is a free data retrieval call binding the contract method 0x80d04de8.
//
// Solidity: function totalDelegated() view returns(uint256)
func (_LaraContract *LaraContractCallerSession) TotalDelegated() (*big.Int, error) {
	return _LaraContract.Contract.TotalDelegated(&_LaraContract.CallOpts)
}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_LaraContract *LaraContractCaller) TreasuryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LaraContract.contract.Call(opts, &out, "treasuryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_LaraContract *LaraContractSession) TreasuryAddress() (common.Address, error) {
	return _LaraContract.Contract.TreasuryAddress(&_LaraContract.CallOpts)
}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_LaraContract *LaraContractCallerSession) TreasuryAddress() (common.Address, error) {
	return _LaraContract.Contract.TreasuryAddress(&_LaraContract.CallOpts)
}

// Undelegated is a free data retrieval call binding the contract method 0x53013f29.
//
// Solidity: function undelegated(address ) view returns(uint256)
func (_LaraContract *LaraContractCaller) Undelegated(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LaraContract.contract.Call(opts, &out, "undelegated", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Undelegated is a free data retrieval call binding the contract method 0x53013f29.
//
// Solidity: function undelegated(address ) view returns(uint256)
func (_LaraContract *LaraContractSession) Undelegated(arg0 common.Address) (*big.Int, error) {
	return _LaraContract.Contract.Undelegated(&_LaraContract.CallOpts, arg0)
}

// Undelegated is a free data retrieval call binding the contract method 0x53013f29.
//
// Solidity: function undelegated(address ) view returns(uint256)
func (_LaraContract *LaraContractCallerSession) Undelegated(arg0 common.Address) (*big.Int, error) {
	return _LaraContract.Contract.Undelegated(&_LaraContract.CallOpts, arg0)
}

// Undelegations is a free data retrieval call binding the contract method 0xb9e60f17.
//
// Solidity: function undelegations(address , uint64 ) view returns((uint256,uint64,address,bool) undelegation_data, uint64 undelegation_id)
func (_LaraContract *LaraContractCaller) Undelegations(opts *bind.CallOpts, arg0 common.Address, arg1 uint64) (struct {
	UndelegationData DposInterfaceUndelegationData
	UndelegationId   uint64
}, error) {
	var out []interface{}
	err := _LaraContract.contract.Call(opts, &out, "undelegations", arg0, arg1)

	outstruct := new(struct {
		UndelegationData DposInterfaceUndelegationData
		UndelegationId   uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UndelegationData = *abi.ConvertType(out[0], new(DposInterfaceUndelegationData)).(*DposInterfaceUndelegationData)
	outstruct.UndelegationId = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// Undelegations is a free data retrieval call binding the contract method 0xb9e60f17.
//
// Solidity: function undelegations(address , uint64 ) view returns((uint256,uint64,address,bool) undelegation_data, uint64 undelegation_id)
func (_LaraContract *LaraContractSession) Undelegations(arg0 common.Address, arg1 uint64) (struct {
	UndelegationData DposInterfaceUndelegationData
	UndelegationId   uint64
}, error) {
	return _LaraContract.Contract.Undelegations(&_LaraContract.CallOpts, arg0, arg1)
}

// Undelegations is a free data retrieval call binding the contract method 0xb9e60f17.
//
// Solidity: function undelegations(address , uint64 ) view returns((uint256,uint64,address,bool) undelegation_data, uint64 undelegation_id)
func (_LaraContract *LaraContractCallerSession) Undelegations(arg0 common.Address, arg1 uint64) (struct {
	UndelegationData DposInterfaceUndelegationData
	UndelegationId   uint64
}, error) {
	return _LaraContract.Contract.Undelegations(&_LaraContract.CallOpts, arg0, arg1)
}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 ) view returns(address)
func (_LaraContract *LaraContractCaller) Validators(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LaraContract.contract.Call(opts, &out, "validators", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 ) view returns(address)
func (_LaraContract *LaraContractSession) Validators(arg0 *big.Int) (common.Address, error) {
	return _LaraContract.Contract.Validators(&_LaraContract.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 ) view returns(address)
func (_LaraContract *LaraContractCallerSession) Validators(arg0 *big.Int) (common.Address, error) {
	return _LaraContract.Contract.Validators(&_LaraContract.CallOpts, arg0)
}

// CancelUndelegate is a paid mutator transaction binding the contract method 0x6342306d.
//
// Solidity: function cancelUndelegate(uint64 id) returns()
func (_LaraContract *LaraContractTransactor) CancelUndelegate(opts *bind.TransactOpts, id uint64) (*types.Transaction, error) {
	return _LaraContract.contract.Transact(opts, "cancelUndelegate", id)
}

// CancelUndelegate is a paid mutator transaction binding the contract method 0x6342306d.
//
// Solidity: function cancelUndelegate(uint64 id) returns()
func (_LaraContract *LaraContractSession) CancelUndelegate(id uint64) (*types.Transaction, error) {
	return _LaraContract.Contract.CancelUndelegate(&_LaraContract.TransactOpts, id)
}

// CancelUndelegate is a paid mutator transaction binding the contract method 0x6342306d.
//
// Solidity: function cancelUndelegate(uint64 id) returns()
func (_LaraContract *LaraContractTransactorSession) CancelUndelegate(id uint64) (*types.Transaction, error) {
	return _LaraContract.Contract.CancelUndelegate(&_LaraContract.TransactOpts, id)
}

// Compound is a paid mutator transaction binding the contract method 0xaa5f7e26.
//
// Solidity: function compound(uint256 amount) returns()
func (_LaraContract *LaraContractTransactor) Compound(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _LaraContract.contract.Transact(opts, "compound", amount)
}

// Compound is a paid mutator transaction binding the contract method 0xaa5f7e26.
//
// Solidity: function compound(uint256 amount) returns()
func (_LaraContract *LaraContractSession) Compound(amount *big.Int) (*types.Transaction, error) {
	return _LaraContract.Contract.Compound(&_LaraContract.TransactOpts, amount)
}

// Compound is a paid mutator transaction binding the contract method 0xaa5f7e26.
//
// Solidity: function compound(uint256 amount) returns()
func (_LaraContract *LaraContractTransactorSession) Compound(amount *big.Int) (*types.Transaction, error) {
	return _LaraContract.Contract.Compound(&_LaraContract.TransactOpts, amount)
}

// ConfirmUndelegate is a paid mutator transaction binding the contract method 0xdfb6bdf5.
//
// Solidity: function confirmUndelegate(uint64 id) returns()
func (_LaraContract *LaraContractTransactor) ConfirmUndelegate(opts *bind.TransactOpts, id uint64) (*types.Transaction, error) {
	return _LaraContract.contract.Transact(opts, "confirmUndelegate", id)
}

// ConfirmUndelegate is a paid mutator transaction binding the contract method 0xdfb6bdf5.
//
// Solidity: function confirmUndelegate(uint64 id) returns()
func (_LaraContract *LaraContractSession) ConfirmUndelegate(id uint64) (*types.Transaction, error) {
	return _LaraContract.Contract.ConfirmUndelegate(&_LaraContract.TransactOpts, id)
}

// ConfirmUndelegate is a paid mutator transaction binding the contract method 0xdfb6bdf5.
//
// Solidity: function confirmUndelegate(uint64 id) returns()
func (_LaraContract *LaraContractTransactorSession) ConfirmUndelegate(id uint64) (*types.Transaction, error) {
	return _LaraContract.Contract.ConfirmUndelegate(&_LaraContract.TransactOpts, id)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _sttaraToken, address _dposContract, address _apyOracle, address _treasuryAddress) returns()
func (_LaraContract *LaraContractTransactor) Initialize(opts *bind.TransactOpts, _sttaraToken common.Address, _dposContract common.Address, _apyOracle common.Address, _treasuryAddress common.Address) (*types.Transaction, error) {
	return _LaraContract.contract.Transact(opts, "initialize", _sttaraToken, _dposContract, _apyOracle, _treasuryAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _sttaraToken, address _dposContract, address _apyOracle, address _treasuryAddress) returns()
func (_LaraContract *LaraContractSession) Initialize(_sttaraToken common.Address, _dposContract common.Address, _apyOracle common.Address, _treasuryAddress common.Address) (*types.Transaction, error) {
	return _LaraContract.Contract.Initialize(&_LaraContract.TransactOpts, _sttaraToken, _dposContract, _apyOracle, _treasuryAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _sttaraToken, address _dposContract, address _apyOracle, address _treasuryAddress) returns()
func (_LaraContract *LaraContractTransactorSession) Initialize(_sttaraToken common.Address, _dposContract common.Address, _apyOracle common.Address, _treasuryAddress common.Address) (*types.Transaction, error) {
	return _LaraContract.Contract.Initialize(&_LaraContract.TransactOpts, _sttaraToken, _dposContract, _apyOracle, _treasuryAddress)
}

// Rebalance is a paid mutator transaction binding the contract method 0x7d7c2a1c.
//
// Solidity: function rebalance() returns()
func (_LaraContract *LaraContractTransactor) Rebalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LaraContract.contract.Transact(opts, "rebalance")
}

// Rebalance is a paid mutator transaction binding the contract method 0x7d7c2a1c.
//
// Solidity: function rebalance() returns()
func (_LaraContract *LaraContractSession) Rebalance() (*types.Transaction, error) {
	return _LaraContract.Contract.Rebalance(&_LaraContract.TransactOpts)
}

// Rebalance is a paid mutator transaction binding the contract method 0x7d7c2a1c.
//
// Solidity: function rebalance() returns()
func (_LaraContract *LaraContractTransactorSession) Rebalance() (*types.Transaction, error) {
	return _LaraContract.Contract.Rebalance(&_LaraContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LaraContract *LaraContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LaraContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LaraContract *LaraContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _LaraContract.Contract.RenounceOwnership(&_LaraContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LaraContract *LaraContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LaraContract.Contract.RenounceOwnership(&_LaraContract.TransactOpts)
}

// RequestUndelegate is a paid mutator transaction binding the contract method 0xf86bc80c.
//
// Solidity: function requestUndelegate(uint256 amount) returns(uint64[] undelegation_ids)
func (_LaraContract *LaraContractTransactor) RequestUndelegate(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _LaraContract.contract.Transact(opts, "requestUndelegate", amount)
}

// RequestUndelegate is a paid mutator transaction binding the contract method 0xf86bc80c.
//
// Solidity: function requestUndelegate(uint256 amount) returns(uint64[] undelegation_ids)
func (_LaraContract *LaraContractSession) RequestUndelegate(amount *big.Int) (*types.Transaction, error) {
	return _LaraContract.Contract.RequestUndelegate(&_LaraContract.TransactOpts, amount)
}

// RequestUndelegate is a paid mutator transaction binding the contract method 0xf86bc80c.
//
// Solidity: function requestUndelegate(uint256 amount) returns(uint64[] undelegation_ids)
func (_LaraContract *LaraContractTransactorSession) RequestUndelegate(amount *big.Int) (*types.Transaction, error) {
	return _LaraContract.Contract.RequestUndelegate(&_LaraContract.TransactOpts, amount)
}

// SetCommission is a paid mutator transaction binding the contract method 0x355e6b43.
//
// Solidity: function setCommission(uint256 _commission) returns()
func (_LaraContract *LaraContractTransactor) SetCommission(opts *bind.TransactOpts, _commission *big.Int) (*types.Transaction, error) {
	return _LaraContract.contract.Transact(opts, "setCommission", _commission)
}

// SetCommission is a paid mutator transaction binding the contract method 0x355e6b43.
//
// Solidity: function setCommission(uint256 _commission) returns()
func (_LaraContract *LaraContractSession) SetCommission(_commission *big.Int) (*types.Transaction, error) {
	return _LaraContract.Contract.SetCommission(&_LaraContract.TransactOpts, _commission)
}

// SetCommission is a paid mutator transaction binding the contract method 0x355e6b43.
//
// Solidity: function setCommission(uint256 _commission) returns()
func (_LaraContract *LaraContractTransactorSession) SetCommission(_commission *big.Int) (*types.Transaction, error) {
	return _LaraContract.Contract.SetCommission(&_LaraContract.TransactOpts, _commission)
}

// SetEpochDuration is a paid mutator transaction binding the contract method 0x30024dfe.
//
// Solidity: function setEpochDuration(uint256 _epochDuration) returns()
func (_LaraContract *LaraContractTransactor) SetEpochDuration(opts *bind.TransactOpts, _epochDuration *big.Int) (*types.Transaction, error) {
	return _LaraContract.contract.Transact(opts, "setEpochDuration", _epochDuration)
}

// SetEpochDuration is a paid mutator transaction binding the contract method 0x30024dfe.
//
// Solidity: function setEpochDuration(uint256 _epochDuration) returns()
func (_LaraContract *LaraContractSession) SetEpochDuration(_epochDuration *big.Int) (*types.Transaction, error) {
	return _LaraContract.Contract.SetEpochDuration(&_LaraContract.TransactOpts, _epochDuration)
}

// SetEpochDuration is a paid mutator transaction binding the contract method 0x30024dfe.
//
// Solidity: function setEpochDuration(uint256 _epochDuration) returns()
func (_LaraContract *LaraContractTransactorSession) SetEpochDuration(_epochDuration *big.Int) (*types.Transaction, error) {
	return _LaraContract.Contract.SetEpochDuration(&_LaraContract.TransactOpts, _epochDuration)
}

// SetMaxValidatorStakeCapacity is a paid mutator transaction binding the contract method 0x6d2d8519.
//
// Solidity: function setMaxValidatorStakeCapacity(uint256 _maxValidatorStakeCapacity) returns()
func (_LaraContract *LaraContractTransactor) SetMaxValidatorStakeCapacity(opts *bind.TransactOpts, _maxValidatorStakeCapacity *big.Int) (*types.Transaction, error) {
	return _LaraContract.contract.Transact(opts, "setMaxValidatorStakeCapacity", _maxValidatorStakeCapacity)
}

// SetMaxValidatorStakeCapacity is a paid mutator transaction binding the contract method 0x6d2d8519.
//
// Solidity: function setMaxValidatorStakeCapacity(uint256 _maxValidatorStakeCapacity) returns()
func (_LaraContract *LaraContractSession) SetMaxValidatorStakeCapacity(_maxValidatorStakeCapacity *big.Int) (*types.Transaction, error) {
	return _LaraContract.Contract.SetMaxValidatorStakeCapacity(&_LaraContract.TransactOpts, _maxValidatorStakeCapacity)
}

// SetMaxValidatorStakeCapacity is a paid mutator transaction binding the contract method 0x6d2d8519.
//
// Solidity: function setMaxValidatorStakeCapacity(uint256 _maxValidatorStakeCapacity) returns()
func (_LaraContract *LaraContractTransactorSession) SetMaxValidatorStakeCapacity(_maxValidatorStakeCapacity *big.Int) (*types.Transaction, error) {
	return _LaraContract.Contract.SetMaxValidatorStakeCapacity(&_LaraContract.TransactOpts, _maxValidatorStakeCapacity)
}

// SetMinStakeAmount is a paid mutator transaction binding the contract method 0xeb4af045.
//
// Solidity: function setMinStakeAmount(uint256 _minStakeAmount) returns()
func (_LaraContract *LaraContractTransactor) SetMinStakeAmount(opts *bind.TransactOpts, _minStakeAmount *big.Int) (*types.Transaction, error) {
	return _LaraContract.contract.Transact(opts, "setMinStakeAmount", _minStakeAmount)
}

// SetMinStakeAmount is a paid mutator transaction binding the contract method 0xeb4af045.
//
// Solidity: function setMinStakeAmount(uint256 _minStakeAmount) returns()
func (_LaraContract *LaraContractSession) SetMinStakeAmount(_minStakeAmount *big.Int) (*types.Transaction, error) {
	return _LaraContract.Contract.SetMinStakeAmount(&_LaraContract.TransactOpts, _minStakeAmount)
}

// SetMinStakeAmount is a paid mutator transaction binding the contract method 0xeb4af045.
//
// Solidity: function setMinStakeAmount(uint256 _minStakeAmount) returns()
func (_LaraContract *LaraContractTransactorSession) SetMinStakeAmount(_minStakeAmount *big.Int) (*types.Transaction, error) {
	return _LaraContract.Contract.SetMinStakeAmount(&_LaraContract.TransactOpts, _minStakeAmount)
}

// SetTreasuryAddress is a paid mutator transaction binding the contract method 0x6605bfda.
//
// Solidity: function setTreasuryAddress(address _treasuryAddress) returns()
func (_LaraContract *LaraContractTransactor) SetTreasuryAddress(opts *bind.TransactOpts, _treasuryAddress common.Address) (*types.Transaction, error) {
	return _LaraContract.contract.Transact(opts, "setTreasuryAddress", _treasuryAddress)
}

// SetTreasuryAddress is a paid mutator transaction binding the contract method 0x6605bfda.
//
// Solidity: function setTreasuryAddress(address _treasuryAddress) returns()
func (_LaraContract *LaraContractSession) SetTreasuryAddress(_treasuryAddress common.Address) (*types.Transaction, error) {
	return _LaraContract.Contract.SetTreasuryAddress(&_LaraContract.TransactOpts, _treasuryAddress)
}

// SetTreasuryAddress is a paid mutator transaction binding the contract method 0x6605bfda.
//
// Solidity: function setTreasuryAddress(address _treasuryAddress) returns()
func (_LaraContract *LaraContractTransactorSession) SetTreasuryAddress(_treasuryAddress common.Address) (*types.Transaction, error) {
	return _LaraContract.Contract.SetTreasuryAddress(&_LaraContract.TransactOpts, _treasuryAddress)
}

// Snapshot is a paid mutator transaction binding the contract method 0x9711715a.
//
// Solidity: function snapshot() returns()
func (_LaraContract *LaraContractTransactor) Snapshot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LaraContract.contract.Transact(opts, "snapshot")
}

// Snapshot is a paid mutator transaction binding the contract method 0x9711715a.
//
// Solidity: function snapshot() returns()
func (_LaraContract *LaraContractSession) Snapshot() (*types.Transaction, error) {
	return _LaraContract.Contract.Snapshot(&_LaraContract.TransactOpts)
}

// Snapshot is a paid mutator transaction binding the contract method 0x9711715a.
//
// Solidity: function snapshot() returns()
func (_LaraContract *LaraContractTransactorSession) Snapshot() (*types.Transaction, error) {
	return _LaraContract.Contract.Snapshot(&_LaraContract.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) payable returns(uint256)
func (_LaraContract *LaraContractTransactor) Stake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _LaraContract.contract.Transact(opts, "stake", amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) payable returns(uint256)
func (_LaraContract *LaraContractSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _LaraContract.Contract.Stake(&_LaraContract.TransactOpts, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) payable returns(uint256)
func (_LaraContract *LaraContractTransactorSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _LaraContract.Contract.Stake(&_LaraContract.TransactOpts, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LaraContract *LaraContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LaraContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LaraContract *LaraContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LaraContract.Contract.TransferOwnership(&_LaraContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LaraContract *LaraContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LaraContract.Contract.TransferOwnership(&_LaraContract.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LaraContract *LaraContractTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LaraContract.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LaraContract *LaraContractSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LaraContract.Contract.UpgradeToAndCall(&_LaraContract.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LaraContract *LaraContractTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LaraContract.Contract.UpgradeToAndCall(&_LaraContract.TransactOpts, newImplementation, data)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_LaraContract *LaraContractTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _LaraContract.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_LaraContract *LaraContractSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _LaraContract.Contract.Fallback(&_LaraContract.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_LaraContract *LaraContractTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _LaraContract.Contract.Fallback(&_LaraContract.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LaraContract *LaraContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LaraContract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LaraContract *LaraContractSession) Receive() (*types.Transaction, error) {
	return _LaraContract.Contract.Receive(&_LaraContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LaraContract *LaraContractTransactorSession) Receive() (*types.Transaction, error) {
	return _LaraContract.Contract.Receive(&_LaraContract.TransactOpts)
}

// LaraContractAllRewardsClaimedIterator is returned from FilterAllRewardsClaimed and is used to iterate over the raw logs and unpacked data for AllRewardsClaimed events raised by the LaraContract contract.
type LaraContractAllRewardsClaimedIterator struct {
	Event *LaraContractAllRewardsClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LaraContractAllRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaraContractAllRewardsClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LaraContractAllRewardsClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LaraContractAllRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaraContractAllRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaraContractAllRewardsClaimed represents a AllRewardsClaimed event raised by the LaraContract contract.
type LaraContractAllRewardsClaimed struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAllRewardsClaimed is a free log retrieval operation binding the contract event 0x06b7f38a79869900bd1aadf75f7322983f44648a0899421e4b8ade76235f63c3.
//
// Solidity: event AllRewardsClaimed(uint256 indexed amount)
func (_LaraContract *LaraContractFilterer) FilterAllRewardsClaimed(opts *bind.FilterOpts, amount []*big.Int) (*LaraContractAllRewardsClaimedIterator, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _LaraContract.contract.FilterLogs(opts, "AllRewardsClaimed", amountRule)
	if err != nil {
		return nil, err
	}
	return &LaraContractAllRewardsClaimedIterator{contract: _LaraContract.contract, event: "AllRewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchAllRewardsClaimed is a free log subscription operation binding the contract event 0x06b7f38a79869900bd1aadf75f7322983f44648a0899421e4b8ade76235f63c3.
//
// Solidity: event AllRewardsClaimed(uint256 indexed amount)
func (_LaraContract *LaraContractFilterer) WatchAllRewardsClaimed(opts *bind.WatchOpts, sink chan<- *LaraContractAllRewardsClaimed, amount []*big.Int) (event.Subscription, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _LaraContract.contract.WatchLogs(opts, "AllRewardsClaimed", amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaraContractAllRewardsClaimed)
				if err := _LaraContract.contract.UnpackLog(event, "AllRewardsClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAllRewardsClaimed is a log parse operation binding the contract event 0x06b7f38a79869900bd1aadf75f7322983f44648a0899421e4b8ade76235f63c3.
//
// Solidity: event AllRewardsClaimed(uint256 indexed amount)
func (_LaraContract *LaraContractFilterer) ParseAllRewardsClaimed(log types.Log) (*LaraContractAllRewardsClaimed, error) {
	event := new(LaraContractAllRewardsClaimed)
	if err := _LaraContract.contract.UnpackLog(event, "AllRewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaraContractCommissionChangedIterator is returned from FilterCommissionChanged and is used to iterate over the raw logs and unpacked data for CommissionChanged events raised by the LaraContract contract.
type LaraContractCommissionChangedIterator struct {
	Event *LaraContractCommissionChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LaraContractCommissionChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaraContractCommissionChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LaraContractCommissionChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LaraContractCommissionChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaraContractCommissionChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaraContractCommissionChanged represents a CommissionChanged event raised by the LaraContract contract.
type LaraContractCommissionChanged struct {
	NewCommission *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCommissionChanged is a free log retrieval operation binding the contract event 0x839e4456845dbc05c7d8638cf0b0976161331b5f9163980d71d9a6444a326c61.
//
// Solidity: event CommissionChanged(uint256 indexed newCommission)
func (_LaraContract *LaraContractFilterer) FilterCommissionChanged(opts *bind.FilterOpts, newCommission []*big.Int) (*LaraContractCommissionChangedIterator, error) {

	var newCommissionRule []interface{}
	for _, newCommissionItem := range newCommission {
		newCommissionRule = append(newCommissionRule, newCommissionItem)
	}

	logs, sub, err := _LaraContract.contract.FilterLogs(opts, "CommissionChanged", newCommissionRule)
	if err != nil {
		return nil, err
	}
	return &LaraContractCommissionChangedIterator{contract: _LaraContract.contract, event: "CommissionChanged", logs: logs, sub: sub}, nil
}

// WatchCommissionChanged is a free log subscription operation binding the contract event 0x839e4456845dbc05c7d8638cf0b0976161331b5f9163980d71d9a6444a326c61.
//
// Solidity: event CommissionChanged(uint256 indexed newCommission)
func (_LaraContract *LaraContractFilterer) WatchCommissionChanged(opts *bind.WatchOpts, sink chan<- *LaraContractCommissionChanged, newCommission []*big.Int) (event.Subscription, error) {

	var newCommissionRule []interface{}
	for _, newCommissionItem := range newCommission {
		newCommissionRule = append(newCommissionRule, newCommissionItem)
	}

	logs, sub, err := _LaraContract.contract.WatchLogs(opts, "CommissionChanged", newCommissionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaraContractCommissionChanged)
				if err := _LaraContract.contract.UnpackLog(event, "CommissionChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCommissionChanged is a log parse operation binding the contract event 0x839e4456845dbc05c7d8638cf0b0976161331b5f9163980d71d9a6444a326c61.
//
// Solidity: event CommissionChanged(uint256 indexed newCommission)
func (_LaraContract *LaraContractFilterer) ParseCommissionChanged(log types.Log) (*LaraContractCommissionChanged, error) {
	event := new(LaraContractCommissionChanged)
	if err := _LaraContract.contract.UnpackLog(event, "CommissionChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaraContractCommissionWithdrawnIterator is returned from FilterCommissionWithdrawn and is used to iterate over the raw logs and unpacked data for CommissionWithdrawn events raised by the LaraContract contract.
type LaraContractCommissionWithdrawnIterator struct {
	Event *LaraContractCommissionWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LaraContractCommissionWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaraContractCommissionWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LaraContractCommissionWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LaraContractCommissionWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaraContractCommissionWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaraContractCommissionWithdrawn represents a CommissionWithdrawn event raised by the LaraContract contract.
type LaraContractCommissionWithdrawn struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCommissionWithdrawn is a free log retrieval operation binding the contract event 0xd244b5a3b2e3977ecffe1a5e5ab7661aadfecbae24be711b7a72bb42bd1b2db0.
//
// Solidity: event CommissionWithdrawn(address indexed user, uint256 indexed amount)
func (_LaraContract *LaraContractFilterer) FilterCommissionWithdrawn(opts *bind.FilterOpts, user []common.Address, amount []*big.Int) (*LaraContractCommissionWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _LaraContract.contract.FilterLogs(opts, "CommissionWithdrawn", userRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &LaraContractCommissionWithdrawnIterator{contract: _LaraContract.contract, event: "CommissionWithdrawn", logs: logs, sub: sub}, nil
}

// WatchCommissionWithdrawn is a free log subscription operation binding the contract event 0xd244b5a3b2e3977ecffe1a5e5ab7661aadfecbae24be711b7a72bb42bd1b2db0.
//
// Solidity: event CommissionWithdrawn(address indexed user, uint256 indexed amount)
func (_LaraContract *LaraContractFilterer) WatchCommissionWithdrawn(opts *bind.WatchOpts, sink chan<- *LaraContractCommissionWithdrawn, user []common.Address, amount []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _LaraContract.contract.WatchLogs(opts, "CommissionWithdrawn", userRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaraContractCommissionWithdrawn)
				if err := _LaraContract.contract.UnpackLog(event, "CommissionWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCommissionWithdrawn is a log parse operation binding the contract event 0xd244b5a3b2e3977ecffe1a5e5ab7661aadfecbae24be711b7a72bb42bd1b2db0.
//
// Solidity: event CommissionWithdrawn(address indexed user, uint256 indexed amount)
func (_LaraContract *LaraContractFilterer) ParseCommissionWithdrawn(log types.Log) (*LaraContractCommissionWithdrawn, error) {
	event := new(LaraContractCommissionWithdrawn)
	if err := _LaraContract.contract.UnpackLog(event, "CommissionWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaraContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the LaraContract contract.
type LaraContractInitializedIterator struct {
	Event *LaraContractInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LaraContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaraContractInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LaraContractInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LaraContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaraContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaraContractInitialized represents a Initialized event raised by the LaraContract contract.
type LaraContractInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_LaraContract *LaraContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*LaraContractInitializedIterator, error) {

	logs, sub, err := _LaraContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LaraContractInitializedIterator{contract: _LaraContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_LaraContract *LaraContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LaraContractInitialized) (event.Subscription, error) {

	logs, sub, err := _LaraContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaraContractInitialized)
				if err := _LaraContract.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_LaraContract *LaraContractFilterer) ParseInitialized(log types.Log) (*LaraContractInitialized, error) {
	event := new(LaraContractInitialized)
	if err := _LaraContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaraContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LaraContract contract.
type LaraContractOwnershipTransferredIterator struct {
	Event *LaraContractOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LaraContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaraContractOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LaraContractOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LaraContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaraContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaraContractOwnershipTransferred represents a OwnershipTransferred event raised by the LaraContract contract.
type LaraContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LaraContract *LaraContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LaraContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LaraContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LaraContractOwnershipTransferredIterator{contract: _LaraContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LaraContract *LaraContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LaraContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LaraContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaraContractOwnershipTransferred)
				if err := _LaraContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LaraContract *LaraContractFilterer) ParseOwnershipTransferred(log types.Log) (*LaraContractOwnershipTransferred, error) {
	event := new(LaraContractOwnershipTransferred)
	if err := _LaraContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaraContractRedelegationRewardsClaimedIterator is returned from FilterRedelegationRewardsClaimed and is used to iterate over the raw logs and unpacked data for RedelegationRewardsClaimed events raised by the LaraContract contract.
type LaraContractRedelegationRewardsClaimedIterator struct {
	Event *LaraContractRedelegationRewardsClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LaraContractRedelegationRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaraContractRedelegationRewardsClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LaraContractRedelegationRewardsClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LaraContractRedelegationRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaraContractRedelegationRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaraContractRedelegationRewardsClaimed represents a RedelegationRewardsClaimed event raised by the LaraContract contract.
type LaraContractRedelegationRewardsClaimed struct {
	Amount    *big.Int
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRedelegationRewardsClaimed is a free log retrieval operation binding the contract event 0x126041a9ce96bf0b59451f9688c03fa384b673c2a8ba3c8dc59adc393a69862e.
//
// Solidity: event RedelegationRewardsClaimed(uint256 indexed amount, address indexed validator)
func (_LaraContract *LaraContractFilterer) FilterRedelegationRewardsClaimed(opts *bind.FilterOpts, amount []*big.Int, validator []common.Address) (*LaraContractRedelegationRewardsClaimedIterator, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _LaraContract.contract.FilterLogs(opts, "RedelegationRewardsClaimed", amountRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &LaraContractRedelegationRewardsClaimedIterator{contract: _LaraContract.contract, event: "RedelegationRewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchRedelegationRewardsClaimed is a free log subscription operation binding the contract event 0x126041a9ce96bf0b59451f9688c03fa384b673c2a8ba3c8dc59adc393a69862e.
//
// Solidity: event RedelegationRewardsClaimed(uint256 indexed amount, address indexed validator)
func (_LaraContract *LaraContractFilterer) WatchRedelegationRewardsClaimed(opts *bind.WatchOpts, sink chan<- *LaraContractRedelegationRewardsClaimed, amount []*big.Int, validator []common.Address) (event.Subscription, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _LaraContract.contract.WatchLogs(opts, "RedelegationRewardsClaimed", amountRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaraContractRedelegationRewardsClaimed)
				if err := _LaraContract.contract.UnpackLog(event, "RedelegationRewardsClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRedelegationRewardsClaimed is a log parse operation binding the contract event 0x126041a9ce96bf0b59451f9688c03fa384b673c2a8ba3c8dc59adc393a69862e.
//
// Solidity: event RedelegationRewardsClaimed(uint256 indexed amount, address indexed validator)
func (_LaraContract *LaraContractFilterer) ParseRedelegationRewardsClaimed(log types.Log) (*LaraContractRedelegationRewardsClaimed, error) {
	event := new(LaraContractRedelegationRewardsClaimed)
	if err := _LaraContract.contract.UnpackLog(event, "RedelegationRewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaraContractSnapshotTakenIterator is returned from FilterSnapshotTaken and is used to iterate over the raw logs and unpacked data for SnapshotTaken events raised by the LaraContract contract.
type LaraContractSnapshotTakenIterator struct {
	Event *LaraContractSnapshotTaken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LaraContractSnapshotTakenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaraContractSnapshotTaken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LaraContractSnapshotTaken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LaraContractSnapshotTakenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaraContractSnapshotTakenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaraContractSnapshotTaken represents a SnapshotTaken event raised by the LaraContract contract.
type LaraContractSnapshotTaken struct {
	TotalDelegation   *big.Int
	TotalRewards      *big.Int
	NextSnapshotBlock *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSnapshotTaken is a free log retrieval operation binding the contract event 0x7be9d0a76c3aa32b2063c1e71a2737740371887411d23841efd15985cce48f53.
//
// Solidity: event SnapshotTaken(uint256 indexed totalDelegation, uint256 indexed totalRewards, uint256 indexed nextSnapshotBlock)
func (_LaraContract *LaraContractFilterer) FilterSnapshotTaken(opts *bind.FilterOpts, totalDelegation []*big.Int, totalRewards []*big.Int, nextSnapshotBlock []*big.Int) (*LaraContractSnapshotTakenIterator, error) {

	var totalDelegationRule []interface{}
	for _, totalDelegationItem := range totalDelegation {
		totalDelegationRule = append(totalDelegationRule, totalDelegationItem)
	}
	var totalRewardsRule []interface{}
	for _, totalRewardsItem := range totalRewards {
		totalRewardsRule = append(totalRewardsRule, totalRewardsItem)
	}
	var nextSnapshotBlockRule []interface{}
	for _, nextSnapshotBlockItem := range nextSnapshotBlock {
		nextSnapshotBlockRule = append(nextSnapshotBlockRule, nextSnapshotBlockItem)
	}

	logs, sub, err := _LaraContract.contract.FilterLogs(opts, "SnapshotTaken", totalDelegationRule, totalRewardsRule, nextSnapshotBlockRule)
	if err != nil {
		return nil, err
	}
	return &LaraContractSnapshotTakenIterator{contract: _LaraContract.contract, event: "SnapshotTaken", logs: logs, sub: sub}, nil
}

// WatchSnapshotTaken is a free log subscription operation binding the contract event 0x7be9d0a76c3aa32b2063c1e71a2737740371887411d23841efd15985cce48f53.
//
// Solidity: event SnapshotTaken(uint256 indexed totalDelegation, uint256 indexed totalRewards, uint256 indexed nextSnapshotBlock)
func (_LaraContract *LaraContractFilterer) WatchSnapshotTaken(opts *bind.WatchOpts, sink chan<- *LaraContractSnapshotTaken, totalDelegation []*big.Int, totalRewards []*big.Int, nextSnapshotBlock []*big.Int) (event.Subscription, error) {

	var totalDelegationRule []interface{}
	for _, totalDelegationItem := range totalDelegation {
		totalDelegationRule = append(totalDelegationRule, totalDelegationItem)
	}
	var totalRewardsRule []interface{}
	for _, totalRewardsItem := range totalRewards {
		totalRewardsRule = append(totalRewardsRule, totalRewardsItem)
	}
	var nextSnapshotBlockRule []interface{}
	for _, nextSnapshotBlockItem := range nextSnapshotBlock {
		nextSnapshotBlockRule = append(nextSnapshotBlockRule, nextSnapshotBlockItem)
	}

	logs, sub, err := _LaraContract.contract.WatchLogs(opts, "SnapshotTaken", totalDelegationRule, totalRewardsRule, nextSnapshotBlockRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaraContractSnapshotTaken)
				if err := _LaraContract.contract.UnpackLog(event, "SnapshotTaken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSnapshotTaken is a log parse operation binding the contract event 0x7be9d0a76c3aa32b2063c1e71a2737740371887411d23841efd15985cce48f53.
//
// Solidity: event SnapshotTaken(uint256 indexed totalDelegation, uint256 indexed totalRewards, uint256 indexed nextSnapshotBlock)
func (_LaraContract *LaraContractFilterer) ParseSnapshotTaken(log types.Log) (*LaraContractSnapshotTaken, error) {
	event := new(LaraContractSnapshotTaken)
	if err := _LaraContract.contract.UnpackLog(event, "SnapshotTaken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaraContractStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the LaraContract contract.
type LaraContractStakedIterator struct {
	Event *LaraContractStaked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LaraContractStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaraContractStaked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LaraContractStaked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LaraContractStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaraContractStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaraContractStaked represents a Staked event raised by the LaraContract contract.
type LaraContractStaked struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 indexed amount)
func (_LaraContract *LaraContractFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address, amount []*big.Int) (*LaraContractStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _LaraContract.contract.FilterLogs(opts, "Staked", userRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &LaraContractStakedIterator{contract: _LaraContract.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 indexed amount)
func (_LaraContract *LaraContractFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *LaraContractStaked, user []common.Address, amount []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _LaraContract.contract.WatchLogs(opts, "Staked", userRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaraContractStaked)
				if err := _LaraContract.contract.UnpackLog(event, "Staked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStaked is a log parse operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 indexed amount)
func (_LaraContract *LaraContractFilterer) ParseStaked(log types.Log) (*LaraContractStaked, error) {
	event := new(LaraContractStaked)
	if err := _LaraContract.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaraContractTaraSentIterator is returned from FilterTaraSent and is used to iterate over the raw logs and unpacked data for TaraSent events raised by the LaraContract contract.
type LaraContractTaraSentIterator struct {
	Event *LaraContractTaraSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LaraContractTaraSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaraContractTaraSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LaraContractTaraSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LaraContractTaraSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaraContractTaraSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaraContractTaraSent represents a TaraSent event raised by the LaraContract contract.
type LaraContractTaraSent struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTaraSent is a free log retrieval operation binding the contract event 0xcc5583b88329e9a0fa4480cb58b74a074292da12cb9926181098e98e4043acc8.
//
// Solidity: event TaraSent(address indexed user, uint256 indexed amount)
func (_LaraContract *LaraContractFilterer) FilterTaraSent(opts *bind.FilterOpts, user []common.Address, amount []*big.Int) (*LaraContractTaraSentIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _LaraContract.contract.FilterLogs(opts, "TaraSent", userRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &LaraContractTaraSentIterator{contract: _LaraContract.contract, event: "TaraSent", logs: logs, sub: sub}, nil
}

// WatchTaraSent is a free log subscription operation binding the contract event 0xcc5583b88329e9a0fa4480cb58b74a074292da12cb9926181098e98e4043acc8.
//
// Solidity: event TaraSent(address indexed user, uint256 indexed amount)
func (_LaraContract *LaraContractFilterer) WatchTaraSent(opts *bind.WatchOpts, sink chan<- *LaraContractTaraSent, user []common.Address, amount []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _LaraContract.contract.WatchLogs(opts, "TaraSent", userRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaraContractTaraSent)
				if err := _LaraContract.contract.UnpackLog(event, "TaraSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaraSent is a log parse operation binding the contract event 0xcc5583b88329e9a0fa4480cb58b74a074292da12cb9926181098e98e4043acc8.
//
// Solidity: event TaraSent(address indexed user, uint256 indexed amount)
func (_LaraContract *LaraContractFilterer) ParseTaraSent(log types.Log) (*LaraContractTaraSent, error) {
	event := new(LaraContractTaraSent)
	if err := _LaraContract.contract.UnpackLog(event, "TaraSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaraContractTreasuryChangedIterator is returned from FilterTreasuryChanged and is used to iterate over the raw logs and unpacked data for TreasuryChanged events raised by the LaraContract contract.
type LaraContractTreasuryChangedIterator struct {
	Event *LaraContractTreasuryChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LaraContractTreasuryChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaraContractTreasuryChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LaraContractTreasuryChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LaraContractTreasuryChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaraContractTreasuryChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaraContractTreasuryChanged represents a TreasuryChanged event raised by the LaraContract contract.
type LaraContractTreasuryChanged struct {
	NewTreasury common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTreasuryChanged is a free log retrieval operation binding the contract event 0xc714d22a2f08b695f81e7c707058db484aa5b4d6b4c9fd64beb10fe85832f608.
//
// Solidity: event TreasuryChanged(address indexed newTreasury)
func (_LaraContract *LaraContractFilterer) FilterTreasuryChanged(opts *bind.FilterOpts, newTreasury []common.Address) (*LaraContractTreasuryChangedIterator, error) {

	var newTreasuryRule []interface{}
	for _, newTreasuryItem := range newTreasury {
		newTreasuryRule = append(newTreasuryRule, newTreasuryItem)
	}

	logs, sub, err := _LaraContract.contract.FilterLogs(opts, "TreasuryChanged", newTreasuryRule)
	if err != nil {
		return nil, err
	}
	return &LaraContractTreasuryChangedIterator{contract: _LaraContract.contract, event: "TreasuryChanged", logs: logs, sub: sub}, nil
}

// WatchTreasuryChanged is a free log subscription operation binding the contract event 0xc714d22a2f08b695f81e7c707058db484aa5b4d6b4c9fd64beb10fe85832f608.
//
// Solidity: event TreasuryChanged(address indexed newTreasury)
func (_LaraContract *LaraContractFilterer) WatchTreasuryChanged(opts *bind.WatchOpts, sink chan<- *LaraContractTreasuryChanged, newTreasury []common.Address) (event.Subscription, error) {

	var newTreasuryRule []interface{}
	for _, newTreasuryItem := range newTreasury {
		newTreasuryRule = append(newTreasuryRule, newTreasuryItem)
	}

	logs, sub, err := _LaraContract.contract.WatchLogs(opts, "TreasuryChanged", newTreasuryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaraContractTreasuryChanged)
				if err := _LaraContract.contract.UnpackLog(event, "TreasuryChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTreasuryChanged is a log parse operation binding the contract event 0xc714d22a2f08b695f81e7c707058db484aa5b4d6b4c9fd64beb10fe85832f608.
//
// Solidity: event TreasuryChanged(address indexed newTreasury)
func (_LaraContract *LaraContractFilterer) ParseTreasuryChanged(log types.Log) (*LaraContractTreasuryChanged, error) {
	event := new(LaraContractTreasuryChanged)
	if err := _LaraContract.contract.UnpackLog(event, "TreasuryChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaraContractUndelegatedIterator is returned from FilterUndelegated and is used to iterate over the raw logs and unpacked data for Undelegated events raised by the LaraContract contract.
type LaraContractUndelegatedIterator struct {
	Event *LaraContractUndelegated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LaraContractUndelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaraContractUndelegated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LaraContractUndelegated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LaraContractUndelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaraContractUndelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaraContractUndelegated represents a Undelegated event raised by the LaraContract contract.
type LaraContractUndelegated struct {
	Id        uint64
	User      common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUndelegated is a free log retrieval operation binding the contract event 0x3f38af2f79cac1d5234e4c00197dfeff662a8c780ce3d810e05e00cccf4d0696.
//
// Solidity: event Undelegated(uint64 indexed id, address indexed user, address indexed validator, uint256 amount)
func (_LaraContract *LaraContractFilterer) FilterUndelegated(opts *bind.FilterOpts, id []uint64, user []common.Address, validator []common.Address) (*LaraContractUndelegatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _LaraContract.contract.FilterLogs(opts, "Undelegated", idRule, userRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &LaraContractUndelegatedIterator{contract: _LaraContract.contract, event: "Undelegated", logs: logs, sub: sub}, nil
}

// WatchUndelegated is a free log subscription operation binding the contract event 0x3f38af2f79cac1d5234e4c00197dfeff662a8c780ce3d810e05e00cccf4d0696.
//
// Solidity: event Undelegated(uint64 indexed id, address indexed user, address indexed validator, uint256 amount)
func (_LaraContract *LaraContractFilterer) WatchUndelegated(opts *bind.WatchOpts, sink chan<- *LaraContractUndelegated, id []uint64, user []common.Address, validator []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _LaraContract.contract.WatchLogs(opts, "Undelegated", idRule, userRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaraContractUndelegated)
				if err := _LaraContract.contract.UnpackLog(event, "Undelegated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUndelegated is a log parse operation binding the contract event 0x3f38af2f79cac1d5234e4c00197dfeff662a8c780ce3d810e05e00cccf4d0696.
//
// Solidity: event Undelegated(uint64 indexed id, address indexed user, address indexed validator, uint256 amount)
func (_LaraContract *LaraContractFilterer) ParseUndelegated(log types.Log) (*LaraContractUndelegated, error) {
	event := new(LaraContractUndelegated)
	if err := _LaraContract.contract.UnpackLog(event, "Undelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaraContractUndelegationCancelledIterator is returned from FilterUndelegationCancelled and is used to iterate over the raw logs and unpacked data for UndelegationCancelled events raised by the LaraContract contract.
type LaraContractUndelegationCancelledIterator struct {
	Event *LaraContractUndelegationCancelled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LaraContractUndelegationCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaraContractUndelegationCancelled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LaraContractUndelegationCancelled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LaraContractUndelegationCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaraContractUndelegationCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaraContractUndelegationCancelled represents a UndelegationCancelled event raised by the LaraContract contract.
type LaraContractUndelegationCancelled struct {
	Id   uint64
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUndelegationCancelled is a free log retrieval operation binding the contract event 0x85ed4768aab5e673ef5298e3bb26c5be7762d867bc6cd32a3519a5e787157dde.
//
// Solidity: event UndelegationCancelled(uint64 indexed id, address indexed user)
func (_LaraContract *LaraContractFilterer) FilterUndelegationCancelled(opts *bind.FilterOpts, id []uint64, user []common.Address) (*LaraContractUndelegationCancelledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LaraContract.contract.FilterLogs(opts, "UndelegationCancelled", idRule, userRule)
	if err != nil {
		return nil, err
	}
	return &LaraContractUndelegationCancelledIterator{contract: _LaraContract.contract, event: "UndelegationCancelled", logs: logs, sub: sub}, nil
}

// WatchUndelegationCancelled is a free log subscription operation binding the contract event 0x85ed4768aab5e673ef5298e3bb26c5be7762d867bc6cd32a3519a5e787157dde.
//
// Solidity: event UndelegationCancelled(uint64 indexed id, address indexed user)
func (_LaraContract *LaraContractFilterer) WatchUndelegationCancelled(opts *bind.WatchOpts, sink chan<- *LaraContractUndelegationCancelled, id []uint64, user []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LaraContract.contract.WatchLogs(opts, "UndelegationCancelled", idRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaraContractUndelegationCancelled)
				if err := _LaraContract.contract.UnpackLog(event, "UndelegationCancelled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUndelegationCancelled is a log parse operation binding the contract event 0x85ed4768aab5e673ef5298e3bb26c5be7762d867bc6cd32a3519a5e787157dde.
//
// Solidity: event UndelegationCancelled(uint64 indexed id, address indexed user)
func (_LaraContract *LaraContractFilterer) ParseUndelegationCancelled(log types.Log) (*LaraContractUndelegationCancelled, error) {
	event := new(LaraContractUndelegationCancelled)
	if err := _LaraContract.contract.UnpackLog(event, "UndelegationCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaraContractUndelegationConfirmedIterator is returned from FilterUndelegationConfirmed and is used to iterate over the raw logs and unpacked data for UndelegationConfirmed events raised by the LaraContract contract.
type LaraContractUndelegationConfirmedIterator struct {
	Event *LaraContractUndelegationConfirmed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LaraContractUndelegationConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaraContractUndelegationConfirmed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LaraContractUndelegationConfirmed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LaraContractUndelegationConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaraContractUndelegationConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaraContractUndelegationConfirmed represents a UndelegationConfirmed event raised by the LaraContract contract.
type LaraContractUndelegationConfirmed struct {
	Id   uint64
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUndelegationConfirmed is a free log retrieval operation binding the contract event 0x72d055944733b65c37b436eaa8614bdf49b8ce59f4703025ea6cdff08df69bf8.
//
// Solidity: event UndelegationConfirmed(uint64 indexed id, address indexed user)
func (_LaraContract *LaraContractFilterer) FilterUndelegationConfirmed(opts *bind.FilterOpts, id []uint64, user []common.Address) (*LaraContractUndelegationConfirmedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LaraContract.contract.FilterLogs(opts, "UndelegationConfirmed", idRule, userRule)
	if err != nil {
		return nil, err
	}
	return &LaraContractUndelegationConfirmedIterator{contract: _LaraContract.contract, event: "UndelegationConfirmed", logs: logs, sub: sub}, nil
}

// WatchUndelegationConfirmed is a free log subscription operation binding the contract event 0x72d055944733b65c37b436eaa8614bdf49b8ce59f4703025ea6cdff08df69bf8.
//
// Solidity: event UndelegationConfirmed(uint64 indexed id, address indexed user)
func (_LaraContract *LaraContractFilterer) WatchUndelegationConfirmed(opts *bind.WatchOpts, sink chan<- *LaraContractUndelegationConfirmed, id []uint64, user []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LaraContract.contract.WatchLogs(opts, "UndelegationConfirmed", idRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaraContractUndelegationConfirmed)
				if err := _LaraContract.contract.UnpackLog(event, "UndelegationConfirmed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUndelegationConfirmed is a log parse operation binding the contract event 0x72d055944733b65c37b436eaa8614bdf49b8ce59f4703025ea6cdff08df69bf8.
//
// Solidity: event UndelegationConfirmed(uint64 indexed id, address indexed user)
func (_LaraContract *LaraContractFilterer) ParseUndelegationConfirmed(log types.Log) (*LaraContractUndelegationConfirmed, error) {
	event := new(LaraContractUndelegationConfirmed)
	if err := _LaraContract.contract.UnpackLog(event, "UndelegationConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaraContractUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the LaraContract contract.
type LaraContractUpgradedIterator struct {
	Event *LaraContractUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LaraContractUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaraContractUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LaraContractUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LaraContractUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaraContractUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaraContractUpgraded represents a Upgraded event raised by the LaraContract contract.
type LaraContractUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_LaraContract *LaraContractFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*LaraContractUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _LaraContract.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &LaraContractUpgradedIterator{contract: _LaraContract.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_LaraContract *LaraContractFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *LaraContractUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _LaraContract.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaraContractUpgraded)
				if err := _LaraContract.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_LaraContract *LaraContractFilterer) ParseUpgraded(log types.Log) (*LaraContractUpgraded, error) {
	event := new(LaraContractUpgraded)
	if err := _LaraContract.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
