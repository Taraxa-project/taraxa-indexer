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

// IApyOracleTentativeDelegation is an auto generated low-level Go binding around an user-defined struct.
type IApyOracleTentativeDelegation struct {
	Validator common.Address
	Amount    *big.Int
}

// LaraContractMetaData contains all meta data concerning the LaraContract contract.
var LaraContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sttaraToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dposContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_apyOracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_treasuryAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"CancelUndelegationFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"ConfirmUndelegationFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"DelegationFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"RedelegationFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"batch\",\"type\":\"uint32\"}],\"name\":\"RewardClaimFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"}],\"name\":\"StakeAmountTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sentAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetAmount\",\"type\":\"uint256\"}],\"name\":\"StakeValueTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"UndelegationFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AllRewardsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCommission\",\"type\":\"uint256\"}],\"name\":\"CommissionChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CommissionWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"CompoundChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Delegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalEpochDelegation\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalEpochReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"EpochEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalEpochDelegation\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"EpochStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"RewardsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"TaraSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newTreasury\",\"type\":\"address\"}],\"name\":\"TreasuryChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Undelegated\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"apyOracle\",\"outputs\":[{\"internalType\":\"contractIApyOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"cancelUndelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimableRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"commission\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"confirmUndelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"delegateStakeOfUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"delegateToDpos\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"delegateToValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remainingAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delegatedAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"delegators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dposContract\",\"outputs\":[{\"internalType\":\"contractDposInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getDelegatorAtIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getValidatorsForAmount\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIApyOracle.TentativeDelegation[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isCompounding\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEpochRunning\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"isValidatorRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastEpochTotalDelegatedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxValidatorStakeCapacity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minStakeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolStartTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"protocolTotalStakeAtValidator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"reDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"removeStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"requestUndelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_commission\",\"type\":\"uint256\"}],\"name\":\"setCommission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"setCompound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochDuration\",\"type\":\"uint256\"}],\"name\":\"setEpochDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxValidatorStakeCapacity\",\"type\":\"uint256\"}],\"name\":\"setMaxValidatorStakeCapacity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minStakeAmount\",\"type\":\"uint256\"}],\"name\":\"setMinStakeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_treasuryAddress\",\"type\":\"address\"}],\"name\":\"setTreasuryAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stTaraToken\",\"outputs\":[{\"internalType\":\"contractIstTara\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakedAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"undelegated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
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

// ClaimableRewards is a free data retrieval call binding the contract method 0xdc01f60d.
//
// Solidity: function claimableRewards(address ) view returns(uint256)
func (_LaraContract *LaraContractCaller) ClaimableRewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LaraContract.contract.Call(opts, &out, "claimableRewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimableRewards is a free data retrieval call binding the contract method 0xdc01f60d.
//
// Solidity: function claimableRewards(address ) view returns(uint256)
func (_LaraContract *LaraContractSession) ClaimableRewards(arg0 common.Address) (*big.Int, error) {
	return _LaraContract.Contract.ClaimableRewards(&_LaraContract.CallOpts, arg0)
}

// ClaimableRewards is a free data retrieval call binding the contract method 0xdc01f60d.
//
// Solidity: function claimableRewards(address ) view returns(uint256)
func (_LaraContract *LaraContractCallerSession) ClaimableRewards(arg0 common.Address) (*big.Int, error) {
	return _LaraContract.Contract.ClaimableRewards(&_LaraContract.CallOpts, arg0)
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

// DelegatedAmounts is a free data retrieval call binding the contract method 0xf472b8d7.
//
// Solidity: function delegatedAmounts(address ) view returns(uint256)
func (_LaraContract *LaraContractCaller) DelegatedAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LaraContract.contract.Call(opts, &out, "delegatedAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegatedAmounts is a free data retrieval call binding the contract method 0xf472b8d7.
//
// Solidity: function delegatedAmounts(address ) view returns(uint256)
func (_LaraContract *LaraContractSession) DelegatedAmounts(arg0 common.Address) (*big.Int, error) {
	return _LaraContract.Contract.DelegatedAmounts(&_LaraContract.CallOpts, arg0)
}

// DelegatedAmounts is a free data retrieval call binding the contract method 0xf472b8d7.
//
// Solidity: function delegatedAmounts(address ) view returns(uint256)
func (_LaraContract *LaraContractCallerSession) DelegatedAmounts(arg0 common.Address) (*big.Int, error) {
	return _LaraContract.Contract.DelegatedAmounts(&_LaraContract.CallOpts, arg0)
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

// GetDelegatorAtIndex is a free data retrieval call binding the contract method 0x34b3e1d3.
//
// Solidity: function getDelegatorAtIndex(uint256 index) view returns(address)
func (_LaraContract *LaraContractCaller) GetDelegatorAtIndex(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LaraContract.contract.Call(opts, &out, "getDelegatorAtIndex", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDelegatorAtIndex is a free data retrieval call binding the contract method 0x34b3e1d3.
//
// Solidity: function getDelegatorAtIndex(uint256 index) view returns(address)
func (_LaraContract *LaraContractSession) GetDelegatorAtIndex(index *big.Int) (common.Address, error) {
	return _LaraContract.Contract.GetDelegatorAtIndex(&_LaraContract.CallOpts, index)
}

// GetDelegatorAtIndex is a free data retrieval call binding the contract method 0x34b3e1d3.
//
// Solidity: function getDelegatorAtIndex(uint256 index) view returns(address)
func (_LaraContract *LaraContractCallerSession) GetDelegatorAtIndex(index *big.Int) (common.Address, error) {
	return _LaraContract.Contract.GetDelegatorAtIndex(&_LaraContract.CallOpts, index)
}

// IsCompounding is a free data retrieval call binding the contract method 0xf042d758.
//
// Solidity: function isCompounding(address ) view returns(bool)
func (_LaraContract *LaraContractCaller) IsCompounding(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _LaraContract.contract.Call(opts, &out, "isCompounding", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCompounding is a free data retrieval call binding the contract method 0xf042d758.
//
// Solidity: function isCompounding(address ) view returns(bool)
func (_LaraContract *LaraContractSession) IsCompounding(arg0 common.Address) (bool, error) {
	return _LaraContract.Contract.IsCompounding(&_LaraContract.CallOpts, arg0)
}

// IsCompounding is a free data retrieval call binding the contract method 0xf042d758.
//
// Solidity: function isCompounding(address ) view returns(bool)
func (_LaraContract *LaraContractCallerSession) IsCompounding(arg0 common.Address) (bool, error) {
	return _LaraContract.Contract.IsCompounding(&_LaraContract.CallOpts, arg0)
}

// IsEpochRunning is a free data retrieval call binding the contract method 0xc5c75098.
//
// Solidity: function isEpochRunning() view returns(bool)
func (_LaraContract *LaraContractCaller) IsEpochRunning(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LaraContract.contract.Call(opts, &out, "isEpochRunning")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEpochRunning is a free data retrieval call binding the contract method 0xc5c75098.
//
// Solidity: function isEpochRunning() view returns(bool)
func (_LaraContract *LaraContractSession) IsEpochRunning() (bool, error) {
	return _LaraContract.Contract.IsEpochRunning(&_LaraContract.CallOpts)
}

// IsEpochRunning is a free data retrieval call binding the contract method 0xc5c75098.
//
// Solidity: function isEpochRunning() view returns(bool)
func (_LaraContract *LaraContractCallerSession) IsEpochRunning() (bool, error) {
	return _LaraContract.Contract.IsEpochRunning(&_LaraContract.CallOpts)
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

// LastEpochTotalDelegatedAmount is a free data retrieval call binding the contract method 0x9a99670e.
//
// Solidity: function lastEpochTotalDelegatedAmount() view returns(uint256)
func (_LaraContract *LaraContractCaller) LastEpochTotalDelegatedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LaraContract.contract.Call(opts, &out, "lastEpochTotalDelegatedAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastEpochTotalDelegatedAmount is a free data retrieval call binding the contract method 0x9a99670e.
//
// Solidity: function lastEpochTotalDelegatedAmount() view returns(uint256)
func (_LaraContract *LaraContractSession) LastEpochTotalDelegatedAmount() (*big.Int, error) {
	return _LaraContract.Contract.LastEpochTotalDelegatedAmount(&_LaraContract.CallOpts)
}

// LastEpochTotalDelegatedAmount is a free data retrieval call binding the contract method 0x9a99670e.
//
// Solidity: function lastEpochTotalDelegatedAmount() view returns(uint256)
func (_LaraContract *LaraContractCallerSession) LastEpochTotalDelegatedAmount() (*big.Int, error) {
	return _LaraContract.Contract.LastEpochTotalDelegatedAmount(&_LaraContract.CallOpts)
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

// StakedAmounts is a free data retrieval call binding the contract method 0x10c1c103.
//
// Solidity: function stakedAmounts(address ) view returns(uint256)
func (_LaraContract *LaraContractCaller) StakedAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LaraContract.contract.Call(opts, &out, "stakedAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakedAmounts is a free data retrieval call binding the contract method 0x10c1c103.
//
// Solidity: function stakedAmounts(address ) view returns(uint256)
func (_LaraContract *LaraContractSession) StakedAmounts(arg0 common.Address) (*big.Int, error) {
	return _LaraContract.Contract.StakedAmounts(&_LaraContract.CallOpts, arg0)
}

// StakedAmounts is a free data retrieval call binding the contract method 0x10c1c103.
//
// Solidity: function stakedAmounts(address ) view returns(uint256)
func (_LaraContract *LaraContractCallerSession) StakedAmounts(arg0 common.Address) (*big.Int, error) {
	return _LaraContract.Contract.StakedAmounts(&_LaraContract.CallOpts, arg0)
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

// CancelUndelegate is a paid mutator transaction binding the contract method 0x3c52e53c.
//
// Solidity: function cancelUndelegate(address validator, uint256 amount) returns()
func (_LaraContract *LaraContractTransactor) CancelUndelegate(opts *bind.TransactOpts, validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LaraContract.contract.Transact(opts, "cancelUndelegate", validator, amount)
}

// CancelUndelegate is a paid mutator transaction binding the contract method 0x3c52e53c.
//
// Solidity: function cancelUndelegate(address validator, uint256 amount) returns()
func (_LaraContract *LaraContractSession) CancelUndelegate(validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LaraContract.Contract.CancelUndelegate(&_LaraContract.TransactOpts, validator, amount)
}

// CancelUndelegate is a paid mutator transaction binding the contract method 0x3c52e53c.
//
// Solidity: function cancelUndelegate(address validator, uint256 amount) returns()
func (_LaraContract *LaraContractTransactorSession) CancelUndelegate(validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LaraContract.Contract.CancelUndelegate(&_LaraContract.TransactOpts, validator, amount)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() returns()
func (_LaraContract *LaraContractTransactor) ClaimRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LaraContract.contract.Transact(opts, "claimRewards")
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() returns()
func (_LaraContract *LaraContractSession) ClaimRewards() (*types.Transaction, error) {
	return _LaraContract.Contract.ClaimRewards(&_LaraContract.TransactOpts)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() returns()
func (_LaraContract *LaraContractTransactorSession) ClaimRewards() (*types.Transaction, error) {
	return _LaraContract.Contract.ClaimRewards(&_LaraContract.TransactOpts)
}

// ConfirmUndelegate is a paid mutator transaction binding the contract method 0x689ad336.
//
// Solidity: function confirmUndelegate(address validator, uint256 amount) returns()
func (_LaraContract *LaraContractTransactor) ConfirmUndelegate(opts *bind.TransactOpts, validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LaraContract.contract.Transact(opts, "confirmUndelegate", validator, amount)
}

// ConfirmUndelegate is a paid mutator transaction binding the contract method 0x689ad336.
//
// Solidity: function confirmUndelegate(address validator, uint256 amount) returns()
func (_LaraContract *LaraContractSession) ConfirmUndelegate(validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LaraContract.Contract.ConfirmUndelegate(&_LaraContract.TransactOpts, validator, amount)
}

// ConfirmUndelegate is a paid mutator transaction binding the contract method 0x689ad336.
//
// Solidity: function confirmUndelegate(address validator, uint256 amount) returns()
func (_LaraContract *LaraContractTransactorSession) ConfirmUndelegate(validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LaraContract.Contract.ConfirmUndelegate(&_LaraContract.TransactOpts, validator, amount)
}

// DelegateStakeOfUser is a paid mutator transaction binding the contract method 0x18bd5305.
//
// Solidity: function delegateStakeOfUser(address user) returns()
func (_LaraContract *LaraContractTransactor) DelegateStakeOfUser(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _LaraContract.contract.Transact(opts, "delegateStakeOfUser", user)
}

// DelegateStakeOfUser is a paid mutator transaction binding the contract method 0x18bd5305.
//
// Solidity: function delegateStakeOfUser(address user) returns()
func (_LaraContract *LaraContractSession) DelegateStakeOfUser(user common.Address) (*types.Transaction, error) {
	return _LaraContract.Contract.DelegateStakeOfUser(&_LaraContract.TransactOpts, user)
}

// DelegateStakeOfUser is a paid mutator transaction binding the contract method 0x18bd5305.
//
// Solidity: function delegateStakeOfUser(address user) returns()
func (_LaraContract *LaraContractTransactorSession) DelegateStakeOfUser(user common.Address) (*types.Transaction, error) {
	return _LaraContract.Contract.DelegateStakeOfUser(&_LaraContract.TransactOpts, user)
}

// DelegateToDpos is a paid mutator transaction binding the contract method 0xdb22009c.
//
// Solidity: function delegateToDpos(address validator) payable returns()
func (_LaraContract *LaraContractTransactor) DelegateToDpos(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _LaraContract.contract.Transact(opts, "delegateToDpos", validator)
}

// DelegateToDpos is a paid mutator transaction binding the contract method 0xdb22009c.
//
// Solidity: function delegateToDpos(address validator) payable returns()
func (_LaraContract *LaraContractSession) DelegateToDpos(validator common.Address) (*types.Transaction, error) {
	return _LaraContract.Contract.DelegateToDpos(&_LaraContract.TransactOpts, validator)
}

// DelegateToDpos is a paid mutator transaction binding the contract method 0xdb22009c.
//
// Solidity: function delegateToDpos(address validator) payable returns()
func (_LaraContract *LaraContractTransactorSession) DelegateToDpos(validator common.Address) (*types.Transaction, error) {
	return _LaraContract.Contract.DelegateToDpos(&_LaraContract.TransactOpts, validator)
}

// DelegateToValidators is a paid mutator transaction binding the contract method 0xccc5b2bd.
//
// Solidity: function delegateToValidators(uint256 amount) returns(uint256 remainingAmount)
func (_LaraContract *LaraContractTransactor) DelegateToValidators(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _LaraContract.contract.Transact(opts, "delegateToValidators", amount)
}

// DelegateToValidators is a paid mutator transaction binding the contract method 0xccc5b2bd.
//
// Solidity: function delegateToValidators(uint256 amount) returns(uint256 remainingAmount)
func (_LaraContract *LaraContractSession) DelegateToValidators(amount *big.Int) (*types.Transaction, error) {
	return _LaraContract.Contract.DelegateToValidators(&_LaraContract.TransactOpts, amount)
}

// DelegateToValidators is a paid mutator transaction binding the contract method 0xccc5b2bd.
//
// Solidity: function delegateToValidators(uint256 amount) returns(uint256 remainingAmount)
func (_LaraContract *LaraContractTransactorSession) DelegateToValidators(amount *big.Int) (*types.Transaction, error) {
	return _LaraContract.Contract.DelegateToValidators(&_LaraContract.TransactOpts, amount)
}

// EndEpoch is a paid mutator transaction binding the contract method 0x0b9663db.
//
// Solidity: function endEpoch() returns()
func (_LaraContract *LaraContractTransactor) EndEpoch(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LaraContract.contract.Transact(opts, "endEpoch")
}

// EndEpoch is a paid mutator transaction binding the contract method 0x0b9663db.
//
// Solidity: function endEpoch() returns()
func (_LaraContract *LaraContractSession) EndEpoch() (*types.Transaction, error) {
	return _LaraContract.Contract.EndEpoch(&_LaraContract.TransactOpts)
}

// EndEpoch is a paid mutator transaction binding the contract method 0x0b9663db.
//
// Solidity: function endEpoch() returns()
func (_LaraContract *LaraContractTransactorSession) EndEpoch() (*types.Transaction, error) {
	return _LaraContract.Contract.EndEpoch(&_LaraContract.TransactOpts)
}

// GetValidatorsForAmount is a paid mutator transaction binding the contract method 0x621c067d.
//
// Solidity: function getValidatorsForAmount(uint256 amount) returns((address,uint256)[])
func (_LaraContract *LaraContractTransactor) GetValidatorsForAmount(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _LaraContract.contract.Transact(opts, "getValidatorsForAmount", amount)
}

// GetValidatorsForAmount is a paid mutator transaction binding the contract method 0x621c067d.
//
// Solidity: function getValidatorsForAmount(uint256 amount) returns((address,uint256)[])
func (_LaraContract *LaraContractSession) GetValidatorsForAmount(amount *big.Int) (*types.Transaction, error) {
	return _LaraContract.Contract.GetValidatorsForAmount(&_LaraContract.TransactOpts, amount)
}

// GetValidatorsForAmount is a paid mutator transaction binding the contract method 0x621c067d.
//
// Solidity: function getValidatorsForAmount(uint256 amount) returns((address,uint256)[])
func (_LaraContract *LaraContractTransactorSession) GetValidatorsForAmount(amount *big.Int) (*types.Transaction, error) {
	return _LaraContract.Contract.GetValidatorsForAmount(&_LaraContract.TransactOpts, amount)
}

// ReDelegate is a paid mutator transaction binding the contract method 0x703812cc.
//
// Solidity: function reDelegate(address from, address to, uint256 amount) returns()
func (_LaraContract *LaraContractTransactor) ReDelegate(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LaraContract.contract.Transact(opts, "reDelegate", from, to, amount)
}

// ReDelegate is a paid mutator transaction binding the contract method 0x703812cc.
//
// Solidity: function reDelegate(address from, address to, uint256 amount) returns()
func (_LaraContract *LaraContractSession) ReDelegate(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LaraContract.Contract.ReDelegate(&_LaraContract.TransactOpts, from, to, amount)
}

// ReDelegate is a paid mutator transaction binding the contract method 0x703812cc.
//
// Solidity: function reDelegate(address from, address to, uint256 amount) returns()
func (_LaraContract *LaraContractTransactorSession) ReDelegate(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LaraContract.Contract.ReDelegate(&_LaraContract.TransactOpts, from, to, amount)
}

// RemoveStake is a paid mutator transaction binding the contract method 0x939624ab.
//
// Solidity: function removeStake(uint256 amount) returns()
func (_LaraContract *LaraContractTransactor) RemoveStake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _LaraContract.contract.Transact(opts, "removeStake", amount)
}

// RemoveStake is a paid mutator transaction binding the contract method 0x939624ab.
//
// Solidity: function removeStake(uint256 amount) returns()
func (_LaraContract *LaraContractSession) RemoveStake(amount *big.Int) (*types.Transaction, error) {
	return _LaraContract.Contract.RemoveStake(&_LaraContract.TransactOpts, amount)
}

// RemoveStake is a paid mutator transaction binding the contract method 0x939624ab.
//
// Solidity: function removeStake(uint256 amount) returns()
func (_LaraContract *LaraContractTransactorSession) RemoveStake(amount *big.Int) (*types.Transaction, error) {
	return _LaraContract.Contract.RemoveStake(&_LaraContract.TransactOpts, amount)
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
// Solidity: function requestUndelegate(uint256 amount) returns()
func (_LaraContract *LaraContractTransactor) RequestUndelegate(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _LaraContract.contract.Transact(opts, "requestUndelegate", amount)
}

// RequestUndelegate is a paid mutator transaction binding the contract method 0xf86bc80c.
//
// Solidity: function requestUndelegate(uint256 amount) returns()
func (_LaraContract *LaraContractSession) RequestUndelegate(amount *big.Int) (*types.Transaction, error) {
	return _LaraContract.Contract.RequestUndelegate(&_LaraContract.TransactOpts, amount)
}

// RequestUndelegate is a paid mutator transaction binding the contract method 0xf86bc80c.
//
// Solidity: function requestUndelegate(uint256 amount) returns()
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

// SetCompound is a paid mutator transaction binding the contract method 0xabfdc73b.
//
// Solidity: function setCompound(bool value) returns()
func (_LaraContract *LaraContractTransactor) SetCompound(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _LaraContract.contract.Transact(opts, "setCompound", value)
}

// SetCompound is a paid mutator transaction binding the contract method 0xabfdc73b.
//
// Solidity: function setCompound(bool value) returns()
func (_LaraContract *LaraContractSession) SetCompound(value bool) (*types.Transaction, error) {
	return _LaraContract.Contract.SetCompound(&_LaraContract.TransactOpts, value)
}

// SetCompound is a paid mutator transaction binding the contract method 0xabfdc73b.
//
// Solidity: function setCompound(bool value) returns()
func (_LaraContract *LaraContractTransactorSession) SetCompound(value bool) (*types.Transaction, error) {
	return _LaraContract.Contract.SetCompound(&_LaraContract.TransactOpts, value)
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

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) payable returns()
func (_LaraContract *LaraContractTransactor) Stake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _LaraContract.contract.Transact(opts, "stake", amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) payable returns()
func (_LaraContract *LaraContractSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _LaraContract.Contract.Stake(&_LaraContract.TransactOpts, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) payable returns()
func (_LaraContract *LaraContractTransactorSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _LaraContract.Contract.Stake(&_LaraContract.TransactOpts, amount)
}

// StartEpoch is a paid mutator transaction binding the contract method 0xa2c8b177.
//
// Solidity: function startEpoch() returns()
func (_LaraContract *LaraContractTransactor) StartEpoch(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LaraContract.contract.Transact(opts, "startEpoch")
}

// StartEpoch is a paid mutator transaction binding the contract method 0xa2c8b177.
//
// Solidity: function startEpoch() returns()
func (_LaraContract *LaraContractSession) StartEpoch() (*types.Transaction, error) {
	return _LaraContract.Contract.StartEpoch(&_LaraContract.TransactOpts)
}

// StartEpoch is a paid mutator transaction binding the contract method 0xa2c8b177.
//
// Solidity: function startEpoch() returns()
func (_LaraContract *LaraContractTransactorSession) StartEpoch() (*types.Transaction, error) {
	return _LaraContract.Contract.StartEpoch(&_LaraContract.TransactOpts)
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
// Solidity: event AllRewardsClaimed(uint256 amount)
func (_LaraContract *LaraContractFilterer) FilterAllRewardsClaimed(opts *bind.FilterOpts) (*LaraContractAllRewardsClaimedIterator, error) {

	logs, sub, err := _LaraContract.contract.FilterLogs(opts, "AllRewardsClaimed")
	if err != nil {
		return nil, err
	}
	return &LaraContractAllRewardsClaimedIterator{contract: _LaraContract.contract, event: "AllRewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchAllRewardsClaimed is a free log subscription operation binding the contract event 0x06b7f38a79869900bd1aadf75f7322983f44648a0899421e4b8ade76235f63c3.
//
// Solidity: event AllRewardsClaimed(uint256 amount)
func (_LaraContract *LaraContractFilterer) WatchAllRewardsClaimed(opts *bind.WatchOpts, sink chan<- *LaraContractAllRewardsClaimed) (event.Subscription, error) {

	logs, sub, err := _LaraContract.contract.WatchLogs(opts, "AllRewardsClaimed")
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
// Solidity: event AllRewardsClaimed(uint256 amount)
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
// Solidity: event CommissionChanged(uint256 newCommission)
func (_LaraContract *LaraContractFilterer) FilterCommissionChanged(opts *bind.FilterOpts) (*LaraContractCommissionChangedIterator, error) {

	logs, sub, err := _LaraContract.contract.FilterLogs(opts, "CommissionChanged")
	if err != nil {
		return nil, err
	}
	return &LaraContractCommissionChangedIterator{contract: _LaraContract.contract, event: "CommissionChanged", logs: logs, sub: sub}, nil
}

// WatchCommissionChanged is a free log subscription operation binding the contract event 0x839e4456845dbc05c7d8638cf0b0976161331b5f9163980d71d9a6444a326c61.
//
// Solidity: event CommissionChanged(uint256 newCommission)
func (_LaraContract *LaraContractFilterer) WatchCommissionChanged(opts *bind.WatchOpts, sink chan<- *LaraContractCommissionChanged) (event.Subscription, error) {

	logs, sub, err := _LaraContract.contract.WatchLogs(opts, "CommissionChanged")
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
// Solidity: event CommissionChanged(uint256 newCommission)
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
// Solidity: event CommissionWithdrawn(address indexed user, uint256 amount)
func (_LaraContract *LaraContractFilterer) FilterCommissionWithdrawn(opts *bind.FilterOpts, user []common.Address) (*LaraContractCommissionWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LaraContract.contract.FilterLogs(opts, "CommissionWithdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &LaraContractCommissionWithdrawnIterator{contract: _LaraContract.contract, event: "CommissionWithdrawn", logs: logs, sub: sub}, nil
}

// WatchCommissionWithdrawn is a free log subscription operation binding the contract event 0xd244b5a3b2e3977ecffe1a5e5ab7661aadfecbae24be711b7a72bb42bd1b2db0.
//
// Solidity: event CommissionWithdrawn(address indexed user, uint256 amount)
func (_LaraContract *LaraContractFilterer) WatchCommissionWithdrawn(opts *bind.WatchOpts, sink chan<- *LaraContractCommissionWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LaraContract.contract.WatchLogs(opts, "CommissionWithdrawn", userRule)
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
// Solidity: event CommissionWithdrawn(address indexed user, uint256 amount)
func (_LaraContract *LaraContractFilterer) ParseCommissionWithdrawn(log types.Log) (*LaraContractCommissionWithdrawn, error) {
	event := new(LaraContractCommissionWithdrawn)
	if err := _LaraContract.contract.UnpackLog(event, "CommissionWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaraContractCompoundChangedIterator is returned from FilterCompoundChanged and is used to iterate over the raw logs and unpacked data for CompoundChanged events raised by the LaraContract contract.
type LaraContractCompoundChangedIterator struct {
	Event *LaraContractCompoundChanged // Event containing the contract specifics and raw log

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
func (it *LaraContractCompoundChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaraContractCompoundChanged)
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
		it.Event = new(LaraContractCompoundChanged)
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
func (it *LaraContractCompoundChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaraContractCompoundChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaraContractCompoundChanged represents a CompoundChanged event raised by the LaraContract contract.
type LaraContractCompoundChanged struct {
	User  common.Address
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCompoundChanged is a free log retrieval operation binding the contract event 0x9aff58f0c4166e57f91e07ecea456a8048481c3d9666af0dfea960fc32bd6dba.
//
// Solidity: event CompoundChanged(address indexed user, bool value)
func (_LaraContract *LaraContractFilterer) FilterCompoundChanged(opts *bind.FilterOpts, user []common.Address) (*LaraContractCompoundChangedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LaraContract.contract.FilterLogs(opts, "CompoundChanged", userRule)
	if err != nil {
		return nil, err
	}
	return &LaraContractCompoundChangedIterator{contract: _LaraContract.contract, event: "CompoundChanged", logs: logs, sub: sub}, nil
}

// WatchCompoundChanged is a free log subscription operation binding the contract event 0x9aff58f0c4166e57f91e07ecea456a8048481c3d9666af0dfea960fc32bd6dba.
//
// Solidity: event CompoundChanged(address indexed user, bool value)
func (_LaraContract *LaraContractFilterer) WatchCompoundChanged(opts *bind.WatchOpts, sink chan<- *LaraContractCompoundChanged, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LaraContract.contract.WatchLogs(opts, "CompoundChanged", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaraContractCompoundChanged)
				if err := _LaraContract.contract.UnpackLog(event, "CompoundChanged", log); err != nil {
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

// ParseCompoundChanged is a log parse operation binding the contract event 0x9aff58f0c4166e57f91e07ecea456a8048481c3d9666af0dfea960fc32bd6dba.
//
// Solidity: event CompoundChanged(address indexed user, bool value)
func (_LaraContract *LaraContractFilterer) ParseCompoundChanged(log types.Log) (*LaraContractCompoundChanged, error) {
	event := new(LaraContractCompoundChanged)
	if err := _LaraContract.contract.UnpackLog(event, "CompoundChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaraContractDelegatedIterator is returned from FilterDelegated and is used to iterate over the raw logs and unpacked data for Delegated events raised by the LaraContract contract.
type LaraContractDelegatedIterator struct {
	Event *LaraContractDelegated // Event containing the contract specifics and raw log

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
func (it *LaraContractDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaraContractDelegated)
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
		it.Event = new(LaraContractDelegated)
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
func (it *LaraContractDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaraContractDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaraContractDelegated represents a Delegated event raised by the LaraContract contract.
type LaraContractDelegated struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDelegated is a free log retrieval operation binding the contract event 0x83b3f5ce88736f0128f880f5cac19836da52ea5c5ca7704c7b38f3b06fffd7ab.
//
// Solidity: event Delegated(address indexed user, uint256 amount)
func (_LaraContract *LaraContractFilterer) FilterDelegated(opts *bind.FilterOpts, user []common.Address) (*LaraContractDelegatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LaraContract.contract.FilterLogs(opts, "Delegated", userRule)
	if err != nil {
		return nil, err
	}
	return &LaraContractDelegatedIterator{contract: _LaraContract.contract, event: "Delegated", logs: logs, sub: sub}, nil
}

// WatchDelegated is a free log subscription operation binding the contract event 0x83b3f5ce88736f0128f880f5cac19836da52ea5c5ca7704c7b38f3b06fffd7ab.
//
// Solidity: event Delegated(address indexed user, uint256 amount)
func (_LaraContract *LaraContractFilterer) WatchDelegated(opts *bind.WatchOpts, sink chan<- *LaraContractDelegated, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LaraContract.contract.WatchLogs(opts, "Delegated", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaraContractDelegated)
				if err := _LaraContract.contract.UnpackLog(event, "Delegated", log); err != nil {
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

// ParseDelegated is a log parse operation binding the contract event 0x83b3f5ce88736f0128f880f5cac19836da52ea5c5ca7704c7b38f3b06fffd7ab.
//
// Solidity: event Delegated(address indexed user, uint256 amount)
func (_LaraContract *LaraContractFilterer) ParseDelegated(log types.Log) (*LaraContractDelegated, error) {
	event := new(LaraContractDelegated)
	if err := _LaraContract.contract.UnpackLog(event, "Delegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaraContractEpochEndedIterator is returned from FilterEpochEnded and is used to iterate over the raw logs and unpacked data for EpochEnded events raised by the LaraContract contract.
type LaraContractEpochEndedIterator struct {
	Event *LaraContractEpochEnded // Event containing the contract specifics and raw log

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
func (it *LaraContractEpochEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaraContractEpochEnded)
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
		it.Event = new(LaraContractEpochEnded)
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
func (it *LaraContractEpochEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaraContractEpochEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaraContractEpochEnded represents a EpochEnded event raised by the LaraContract contract.
type LaraContractEpochEnded struct {
	TotalEpochDelegation *big.Int
	TotalEpochReward     *big.Int
	Timestamp            *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterEpochEnded is a free log retrieval operation binding the contract event 0x469dcd0c8170c8faf3678737ed5cc64107d6d30846f97438a0feefacbed00ec2.
//
// Solidity: event EpochEnded(uint256 totalEpochDelegation, uint256 totalEpochReward, uint256 timestamp)
func (_LaraContract *LaraContractFilterer) FilterEpochEnded(opts *bind.FilterOpts) (*LaraContractEpochEndedIterator, error) {

	logs, sub, err := _LaraContract.contract.FilterLogs(opts, "EpochEnded")
	if err != nil {
		return nil, err
	}
	return &LaraContractEpochEndedIterator{contract: _LaraContract.contract, event: "EpochEnded", logs: logs, sub: sub}, nil
}

// WatchEpochEnded is a free log subscription operation binding the contract event 0x469dcd0c8170c8faf3678737ed5cc64107d6d30846f97438a0feefacbed00ec2.
//
// Solidity: event EpochEnded(uint256 totalEpochDelegation, uint256 totalEpochReward, uint256 timestamp)
func (_LaraContract *LaraContractFilterer) WatchEpochEnded(opts *bind.WatchOpts, sink chan<- *LaraContractEpochEnded) (event.Subscription, error) {

	logs, sub, err := _LaraContract.contract.WatchLogs(opts, "EpochEnded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaraContractEpochEnded)
				if err := _LaraContract.contract.UnpackLog(event, "EpochEnded", log); err != nil {
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

// ParseEpochEnded is a log parse operation binding the contract event 0x469dcd0c8170c8faf3678737ed5cc64107d6d30846f97438a0feefacbed00ec2.
//
// Solidity: event EpochEnded(uint256 totalEpochDelegation, uint256 totalEpochReward, uint256 timestamp)
func (_LaraContract *LaraContractFilterer) ParseEpochEnded(log types.Log) (*LaraContractEpochEnded, error) {
	event := new(LaraContractEpochEnded)
	if err := _LaraContract.contract.UnpackLog(event, "EpochEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaraContractEpochStartedIterator is returned from FilterEpochStarted and is used to iterate over the raw logs and unpacked data for EpochStarted events raised by the LaraContract contract.
type LaraContractEpochStartedIterator struct {
	Event *LaraContractEpochStarted // Event containing the contract specifics and raw log

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
func (it *LaraContractEpochStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaraContractEpochStarted)
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
		it.Event = new(LaraContractEpochStarted)
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
func (it *LaraContractEpochStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaraContractEpochStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaraContractEpochStarted represents a EpochStarted event raised by the LaraContract contract.
type LaraContractEpochStarted struct {
	TotalEpochDelegation *big.Int
	Timestamp            *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterEpochStarted is a free log retrieval operation binding the contract event 0x41787f1277821474072e18df95f0bd9ed9f117003aa97732ebbd737225b32b02.
//
// Solidity: event EpochStarted(uint256 totalEpochDelegation, uint256 timestamp)
func (_LaraContract *LaraContractFilterer) FilterEpochStarted(opts *bind.FilterOpts) (*LaraContractEpochStartedIterator, error) {

	logs, sub, err := _LaraContract.contract.FilterLogs(opts, "EpochStarted")
	if err != nil {
		return nil, err
	}
	return &LaraContractEpochStartedIterator{contract: _LaraContract.contract, event: "EpochStarted", logs: logs, sub: sub}, nil
}

// WatchEpochStarted is a free log subscription operation binding the contract event 0x41787f1277821474072e18df95f0bd9ed9f117003aa97732ebbd737225b32b02.
//
// Solidity: event EpochStarted(uint256 totalEpochDelegation, uint256 timestamp)
func (_LaraContract *LaraContractFilterer) WatchEpochStarted(opts *bind.WatchOpts, sink chan<- *LaraContractEpochStarted) (event.Subscription, error) {

	logs, sub, err := _LaraContract.contract.WatchLogs(opts, "EpochStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaraContractEpochStarted)
				if err := _LaraContract.contract.UnpackLog(event, "EpochStarted", log); err != nil {
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

// ParseEpochStarted is a log parse operation binding the contract event 0x41787f1277821474072e18df95f0bd9ed9f117003aa97732ebbd737225b32b02.
//
// Solidity: event EpochStarted(uint256 totalEpochDelegation, uint256 timestamp)
func (_LaraContract *LaraContractFilterer) ParseEpochStarted(log types.Log) (*LaraContractEpochStarted, error) {
	event := new(LaraContractEpochStarted)
	if err := _LaraContract.contract.UnpackLog(event, "EpochStarted", log); err != nil {
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

// LaraContractRewardsClaimedIterator is returned from FilterRewardsClaimed and is used to iterate over the raw logs and unpacked data for RewardsClaimed events raised by the LaraContract contract.
type LaraContractRewardsClaimedIterator struct {
	Event *LaraContractRewardsClaimed // Event containing the contract specifics and raw log

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
func (it *LaraContractRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaraContractRewardsClaimed)
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
		it.Event = new(LaraContractRewardsClaimed)
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
func (it *LaraContractRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaraContractRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaraContractRewardsClaimed represents a RewardsClaimed event raised by the LaraContract contract.
type LaraContractRewardsClaimed struct {
	User      common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRewardsClaimed is a free log retrieval operation binding the contract event 0xdacbdde355ba930696a362ea6738feb9f8bd52dfb3d81947558fd3217e23e325.
//
// Solidity: event RewardsClaimed(address indexed user, uint256 amount, uint256 timestamp)
func (_LaraContract *LaraContractFilterer) FilterRewardsClaimed(opts *bind.FilterOpts, user []common.Address) (*LaraContractRewardsClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LaraContract.contract.FilterLogs(opts, "RewardsClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return &LaraContractRewardsClaimedIterator{contract: _LaraContract.contract, event: "RewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardsClaimed is a free log subscription operation binding the contract event 0xdacbdde355ba930696a362ea6738feb9f8bd52dfb3d81947558fd3217e23e325.
//
// Solidity: event RewardsClaimed(address indexed user, uint256 amount, uint256 timestamp)
func (_LaraContract *LaraContractFilterer) WatchRewardsClaimed(opts *bind.WatchOpts, sink chan<- *LaraContractRewardsClaimed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LaraContract.contract.WatchLogs(opts, "RewardsClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaraContractRewardsClaimed)
				if err := _LaraContract.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
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

// ParseRewardsClaimed is a log parse operation binding the contract event 0xdacbdde355ba930696a362ea6738feb9f8bd52dfb3d81947558fd3217e23e325.
//
// Solidity: event RewardsClaimed(address indexed user, uint256 amount, uint256 timestamp)
func (_LaraContract *LaraContractFilterer) ParseRewardsClaimed(log types.Log) (*LaraContractRewardsClaimed, error) {
	event := new(LaraContractRewardsClaimed)
	if err := _LaraContract.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaraContractStakeRemovedIterator is returned from FilterStakeRemoved and is used to iterate over the raw logs and unpacked data for StakeRemoved events raised by the LaraContract contract.
type LaraContractStakeRemovedIterator struct {
	Event *LaraContractStakeRemoved // Event containing the contract specifics and raw log

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
func (it *LaraContractStakeRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaraContractStakeRemoved)
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
		it.Event = new(LaraContractStakeRemoved)
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
func (it *LaraContractStakeRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaraContractStakeRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaraContractStakeRemoved represents a StakeRemoved event raised by the LaraContract contract.
type LaraContractStakeRemoved struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakeRemoved is a free log retrieval operation binding the contract event 0xa018dcbc822f59fb0d0c3e7a86c8e4259b9676cdea9e5fc26279b9c4c5d86eef.
//
// Solidity: event StakeRemoved(address indexed user, uint256 amount)
func (_LaraContract *LaraContractFilterer) FilterStakeRemoved(opts *bind.FilterOpts, user []common.Address) (*LaraContractStakeRemovedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LaraContract.contract.FilterLogs(opts, "StakeRemoved", userRule)
	if err != nil {
		return nil, err
	}
	return &LaraContractStakeRemovedIterator{contract: _LaraContract.contract, event: "StakeRemoved", logs: logs, sub: sub}, nil
}

// WatchStakeRemoved is a free log subscription operation binding the contract event 0xa018dcbc822f59fb0d0c3e7a86c8e4259b9676cdea9e5fc26279b9c4c5d86eef.
//
// Solidity: event StakeRemoved(address indexed user, uint256 amount)
func (_LaraContract *LaraContractFilterer) WatchStakeRemoved(opts *bind.WatchOpts, sink chan<- *LaraContractStakeRemoved, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LaraContract.contract.WatchLogs(opts, "StakeRemoved", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaraContractStakeRemoved)
				if err := _LaraContract.contract.UnpackLog(event, "StakeRemoved", log); err != nil {
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

// ParseStakeRemoved is a log parse operation binding the contract event 0xa018dcbc822f59fb0d0c3e7a86c8e4259b9676cdea9e5fc26279b9c4c5d86eef.
//
// Solidity: event StakeRemoved(address indexed user, uint256 amount)
func (_LaraContract *LaraContractFilterer) ParseStakeRemoved(log types.Log) (*LaraContractStakeRemoved, error) {
	event := new(LaraContractStakeRemoved)
	if err := _LaraContract.contract.UnpackLog(event, "StakeRemoved", log); err != nil {
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
// Solidity: event Staked(address indexed user, uint256 amount)
func (_LaraContract *LaraContractFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*LaraContractStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LaraContract.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &LaraContractStakedIterator{contract: _LaraContract.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_LaraContract *LaraContractFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *LaraContractStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LaraContract.contract.WatchLogs(opts, "Staked", userRule)
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
// Solidity: event Staked(address indexed user, uint256 amount)
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
	User        common.Address
	Amount      *big.Int
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTaraSent is a free log retrieval operation binding the contract event 0xb47c5984477bc8ec64a806f6a426296660111ec77d665d05e7d0c582ca518168.
//
// Solidity: event TaraSent(address indexed user, uint256 amount, uint256 blockNumber)
func (_LaraContract *LaraContractFilterer) FilterTaraSent(opts *bind.FilterOpts, user []common.Address) (*LaraContractTaraSentIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LaraContract.contract.FilterLogs(opts, "TaraSent", userRule)
	if err != nil {
		return nil, err
	}
	return &LaraContractTaraSentIterator{contract: _LaraContract.contract, event: "TaraSent", logs: logs, sub: sub}, nil
}

// WatchTaraSent is a free log subscription operation binding the contract event 0xb47c5984477bc8ec64a806f6a426296660111ec77d665d05e7d0c582ca518168.
//
// Solidity: event TaraSent(address indexed user, uint256 amount, uint256 blockNumber)
func (_LaraContract *LaraContractFilterer) WatchTaraSent(opts *bind.WatchOpts, sink chan<- *LaraContractTaraSent, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LaraContract.contract.WatchLogs(opts, "TaraSent", userRule)
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

// ParseTaraSent is a log parse operation binding the contract event 0xb47c5984477bc8ec64a806f6a426296660111ec77d665d05e7d0c582ca518168.
//
// Solidity: event TaraSent(address indexed user, uint256 amount, uint256 blockNumber)
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
	User      common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUndelegated is a free log retrieval operation binding the contract event 0x4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c.
//
// Solidity: event Undelegated(address indexed user, address indexed validator, uint256 amount)
func (_LaraContract *LaraContractFilterer) FilterUndelegated(opts *bind.FilterOpts, user []common.Address, validator []common.Address) (*LaraContractUndelegatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _LaraContract.contract.FilterLogs(opts, "Undelegated", userRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &LaraContractUndelegatedIterator{contract: _LaraContract.contract, event: "Undelegated", logs: logs, sub: sub}, nil
}

// WatchUndelegated is a free log subscription operation binding the contract event 0x4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c.
//
// Solidity: event Undelegated(address indexed user, address indexed validator, uint256 amount)
func (_LaraContract *LaraContractFilterer) WatchUndelegated(opts *bind.WatchOpts, sink chan<- *LaraContractUndelegated, user []common.Address, validator []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _LaraContract.contract.WatchLogs(opts, "Undelegated", userRule, validatorRule)
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

// ParseUndelegated is a log parse operation binding the contract event 0x4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c.
//
// Solidity: event Undelegated(address indexed user, address indexed validator, uint256 amount)
func (_LaraContract *LaraContractFilterer) ParseUndelegated(log types.Log) (*LaraContractUndelegated, error) {
	event := new(LaraContractUndelegated)
	if err := _LaraContract.contract.UnpackLog(event, "Undelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
