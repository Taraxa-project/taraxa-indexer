// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dpos_contract

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

// DposInterfaceDelegationData is an auto generated low-level Go binding around an user-defined struct.
type DposInterfaceDelegationData struct {
	Account    common.Address
	Delegation DposInterfaceDelegatorInfo
}

// DposInterfaceDelegatorInfo is an auto generated low-level Go binding around an user-defined struct.
type DposInterfaceDelegatorInfo struct {
	Stake   *big.Int
	Rewards *big.Int
}

// DposInterfaceUndelegationData is an auto generated low-level Go binding around an user-defined struct.
type DposInterfaceUndelegationData struct {
	Stake           *big.Int
	Block           uint64
	Validator       common.Address
	ValidatorExists bool
}

// DposInterfaceValidatorBasicInfo is an auto generated low-level Go binding around an user-defined struct.
type DposInterfaceValidatorBasicInfo struct {
	TotalStake           *big.Int
	CommissionReward     *big.Int
	Commission           uint16
	LastCommissionChange uint64
	UndelegationsCount   uint16
	Owner                common.Address
	Description          string
	Endpoint             string
}

// DposInterfaceValidatorData is an auto generated low-level Go binding around an user-defined struct.
type DposInterfaceValidatorData struct {
	Account common.Address
	Info    DposInterfaceValidatorBasicInfo
}

// DposContractMetaData contains all meta data concerning the DposContract contract.
var DposContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CommissionRewardsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"commission\",\"type\":\"uint16\"}],\"name\":\"CommissionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Delegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Redelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UndelegateCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UndelegateConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Undelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorInfoSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorRegistered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"cancelUndelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"batch\",\"type\":\"uint32\"}],\"name\":\"claimAllRewards\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"end\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"claimCommissionRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"confirmUndelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"batch\",\"type\":\"uint32\"}],\"name\":\"getDelegations\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"internalType\":\"structDposInterface.DelegatorInfo\",\"name\":\"delegation\",\"type\":\"tuple\"}],\"internalType\":\"structDposInterface.DelegationData[]\",\"name\":\"delegations\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"end\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"getTotalDelegation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total_delegation\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalEligibleVotesCount\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"batch\",\"type\":\"uint32\"}],\"name\":\"getUndelegations\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"block\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"validator_exists\",\"type\":\"bool\"}],\"internalType\":\"structDposInterface.UndelegationData[]\",\"name\":\"undelegations\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"end\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"getValidator\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"total_stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commission_reward\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"commission\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"last_commission_change\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"undelegations_count\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"}],\"internalType\":\"structDposInterface.ValidatorBasicInfo\",\"name\":\"validator_info\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"getValidatorEligibleVotesCount\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"batch\",\"type\":\"uint32\"}],\"name\":\"getValidators\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"total_stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commission_reward\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"commission\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"last_commission_change\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"undelegations_count\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"}],\"internalType\":\"structDposInterface.ValidatorBasicInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"internalType\":\"structDposInterface.ValidatorData[]\",\"name\":\"validators\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"end\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"batch\",\"type\":\"uint32\"}],\"name\":\"getValidatorsFor\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"total_stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commission_reward\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"commission\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"last_commission_change\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"undelegations_count\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"}],\"internalType\":\"structDposInterface.ValidatorBasicInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"internalType\":\"structDposInterface.ValidatorData[]\",\"name\":\"validators\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"end\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"isValidatorEligible\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validator_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"reDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"vrf_key\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"commission\",\"type\":\"uint16\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"}],\"name\":\"registerValidator\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"commission\",\"type\":\"uint16\"}],\"name\":\"setCommission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"}],\"name\":\"setValidatorInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"undelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DposContractABI is the input ABI used to generate the binding from.
// Deprecated: Use DposContractMetaData.ABI instead.
var DposContractABI = DposContractMetaData.ABI

// DposContract is an auto generated Go binding around an Ethereum contract.
type DposContract struct {
	DposContractCaller     // Read-only binding to the contract
	DposContractTransactor // Write-only binding to the contract
	DposContractFilterer   // Log filterer for contract events
}

// DposContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type DposContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DposContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DposContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DposContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DposContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DposContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DposContractSession struct {
	Contract     *DposContract     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DposContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DposContractCallerSession struct {
	Contract *DposContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// DposContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DposContractTransactorSession struct {
	Contract     *DposContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DposContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type DposContractRaw struct {
	Contract *DposContract // Generic contract binding to access the raw methods on
}

// DposContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DposContractCallerRaw struct {
	Contract *DposContractCaller // Generic read-only contract binding to access the raw methods on
}

// DposContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DposContractTransactorRaw struct {
	Contract *DposContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDposContract creates a new instance of DposContract, bound to a specific deployed contract.
func NewDposContract(address common.Address, backend bind.ContractBackend) (*DposContract, error) {
	contract, err := bindDposContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DposContract{DposContractCaller: DposContractCaller{contract: contract}, DposContractTransactor: DposContractTransactor{contract: contract}, DposContractFilterer: DposContractFilterer{contract: contract}}, nil
}

// NewDposContractCaller creates a new read-only instance of DposContract, bound to a specific deployed contract.
func NewDposContractCaller(address common.Address, caller bind.ContractCaller) (*DposContractCaller, error) {
	contract, err := bindDposContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DposContractCaller{contract: contract}, nil
}

// NewDposContractTransactor creates a new write-only instance of DposContract, bound to a specific deployed contract.
func NewDposContractTransactor(address common.Address, transactor bind.ContractTransactor) (*DposContractTransactor, error) {
	contract, err := bindDposContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DposContractTransactor{contract: contract}, nil
}

// NewDposContractFilterer creates a new log filterer instance of DposContract, bound to a specific deployed contract.
func NewDposContractFilterer(address common.Address, filterer bind.ContractFilterer) (*DposContractFilterer, error) {
	contract, err := bindDposContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DposContractFilterer{contract: contract}, nil
}

// bindDposContract binds a generic wrapper to an already deployed contract.
func bindDposContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DposContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DposContract *DposContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DposContract.Contract.DposContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DposContract *DposContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DposContract.Contract.DposContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DposContract *DposContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DposContract.Contract.DposContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DposContract *DposContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DposContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DposContract *DposContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DposContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DposContract *DposContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DposContract.Contract.contract.Transact(opts, method, params...)
}

// GetDelegations is a free data retrieval call binding the contract method 0x8b49d394.
//
// Solidity: function getDelegations(address delegator, uint32 batch) view returns((address,(uint256,uint256))[] delegations, bool end)
func (_DposContract *DposContractCaller) GetDelegations(opts *bind.CallOpts, delegator common.Address, batch uint32) (struct {
	Delegations []DposInterfaceDelegationData
	End         bool
}, error) {
	var out []interface{}
	err := _DposContract.contract.Call(opts, &out, "getDelegations", delegator, batch)

	outstruct := new(struct {
		Delegations []DposInterfaceDelegationData
		End         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Delegations = *abi.ConvertType(out[0], new([]DposInterfaceDelegationData)).(*[]DposInterfaceDelegationData)
	outstruct.End = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetDelegations is a free data retrieval call binding the contract method 0x8b49d394.
//
// Solidity: function getDelegations(address delegator, uint32 batch) view returns((address,(uint256,uint256))[] delegations, bool end)
func (_DposContract *DposContractSession) GetDelegations(delegator common.Address, batch uint32) (struct {
	Delegations []DposInterfaceDelegationData
	End         bool
}, error) {
	return _DposContract.Contract.GetDelegations(&_DposContract.CallOpts, delegator, batch)
}

// GetDelegations is a free data retrieval call binding the contract method 0x8b49d394.
//
// Solidity: function getDelegations(address delegator, uint32 batch) view returns((address,(uint256,uint256))[] delegations, bool end)
func (_DposContract *DposContractCallerSession) GetDelegations(delegator common.Address, batch uint32) (struct {
	Delegations []DposInterfaceDelegationData
	End         bool
}, error) {
	return _DposContract.Contract.GetDelegations(&_DposContract.CallOpts, delegator, batch)
}

// GetTotalDelegation is a free data retrieval call binding the contract method 0xfc5e7e09.
//
// Solidity: function getTotalDelegation(address delegator) view returns(uint256 total_delegation)
func (_DposContract *DposContractCaller) GetTotalDelegation(opts *bind.CallOpts, delegator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DposContract.contract.Call(opts, &out, "getTotalDelegation", delegator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalDelegation is a free data retrieval call binding the contract method 0xfc5e7e09.
//
// Solidity: function getTotalDelegation(address delegator) view returns(uint256 total_delegation)
func (_DposContract *DposContractSession) GetTotalDelegation(delegator common.Address) (*big.Int, error) {
	return _DposContract.Contract.GetTotalDelegation(&_DposContract.CallOpts, delegator)
}

// GetTotalDelegation is a free data retrieval call binding the contract method 0xfc5e7e09.
//
// Solidity: function getTotalDelegation(address delegator) view returns(uint256 total_delegation)
func (_DposContract *DposContractCallerSession) GetTotalDelegation(delegator common.Address) (*big.Int, error) {
	return _DposContract.Contract.GetTotalDelegation(&_DposContract.CallOpts, delegator)
}

// GetTotalEligibleVotesCount is a free data retrieval call binding the contract method 0xde8e4b50.
//
// Solidity: function getTotalEligibleVotesCount() view returns(uint64)
func (_DposContract *DposContractCaller) GetTotalEligibleVotesCount(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _DposContract.contract.Call(opts, &out, "getTotalEligibleVotesCount")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetTotalEligibleVotesCount is a free data retrieval call binding the contract method 0xde8e4b50.
//
// Solidity: function getTotalEligibleVotesCount() view returns(uint64)
func (_DposContract *DposContractSession) GetTotalEligibleVotesCount() (uint64, error) {
	return _DposContract.Contract.GetTotalEligibleVotesCount(&_DposContract.CallOpts)
}

// GetTotalEligibleVotesCount is a free data retrieval call binding the contract method 0xde8e4b50.
//
// Solidity: function getTotalEligibleVotesCount() view returns(uint64)
func (_DposContract *DposContractCallerSession) GetTotalEligibleVotesCount() (uint64, error) {
	return _DposContract.Contract.GetTotalEligibleVotesCount(&_DposContract.CallOpts)
}

// GetUndelegations is a free data retrieval call binding the contract method 0x4edd9943.
//
// Solidity: function getUndelegations(address delegator, uint32 batch) view returns((uint256,uint64,address,bool)[] undelegations, bool end)
func (_DposContract *DposContractCaller) GetUndelegations(opts *bind.CallOpts, delegator common.Address, batch uint32) (struct {
	Undelegations []DposInterfaceUndelegationData
	End           bool
}, error) {
	var out []interface{}
	err := _DposContract.contract.Call(opts, &out, "getUndelegations", delegator, batch)

	outstruct := new(struct {
		Undelegations []DposInterfaceUndelegationData
		End           bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Undelegations = *abi.ConvertType(out[0], new([]DposInterfaceUndelegationData)).(*[]DposInterfaceUndelegationData)
	outstruct.End = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetUndelegations is a free data retrieval call binding the contract method 0x4edd9943.
//
// Solidity: function getUndelegations(address delegator, uint32 batch) view returns((uint256,uint64,address,bool)[] undelegations, bool end)
func (_DposContract *DposContractSession) GetUndelegations(delegator common.Address, batch uint32) (struct {
	Undelegations []DposInterfaceUndelegationData
	End           bool
}, error) {
	return _DposContract.Contract.GetUndelegations(&_DposContract.CallOpts, delegator, batch)
}

// GetUndelegations is a free data retrieval call binding the contract method 0x4edd9943.
//
// Solidity: function getUndelegations(address delegator, uint32 batch) view returns((uint256,uint64,address,bool)[] undelegations, bool end)
func (_DposContract *DposContractCallerSession) GetUndelegations(delegator common.Address, batch uint32) (struct {
	Undelegations []DposInterfaceUndelegationData
	End           bool
}, error) {
	return _DposContract.Contract.GetUndelegations(&_DposContract.CallOpts, delegator, batch)
}

// GetValidator is a free data retrieval call binding the contract method 0x1904bb2e.
//
// Solidity: function getValidator(address validator) view returns((uint256,uint256,uint16,uint64,uint16,address,string,string) validator_info)
func (_DposContract *DposContractCaller) GetValidator(opts *bind.CallOpts, validator common.Address) (DposInterfaceValidatorBasicInfo, error) {
	var out []interface{}
	err := _DposContract.contract.Call(opts, &out, "getValidator", validator)

	if err != nil {
		return *new(DposInterfaceValidatorBasicInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(DposInterfaceValidatorBasicInfo)).(*DposInterfaceValidatorBasicInfo)

	return out0, err

}

// GetValidator is a free data retrieval call binding the contract method 0x1904bb2e.
//
// Solidity: function getValidator(address validator) view returns((uint256,uint256,uint16,uint64,uint16,address,string,string) validator_info)
func (_DposContract *DposContractSession) GetValidator(validator common.Address) (DposInterfaceValidatorBasicInfo, error) {
	return _DposContract.Contract.GetValidator(&_DposContract.CallOpts, validator)
}

// GetValidator is a free data retrieval call binding the contract method 0x1904bb2e.
//
// Solidity: function getValidator(address validator) view returns((uint256,uint256,uint16,uint64,uint16,address,string,string) validator_info)
func (_DposContract *DposContractCallerSession) GetValidator(validator common.Address) (DposInterfaceValidatorBasicInfo, error) {
	return _DposContract.Contract.GetValidator(&_DposContract.CallOpts, validator)
}

// GetValidatorEligibleVotesCount is a free data retrieval call binding the contract method 0x618e3862.
//
// Solidity: function getValidatorEligibleVotesCount(address validator) view returns(uint64)
func (_DposContract *DposContractCaller) GetValidatorEligibleVotesCount(opts *bind.CallOpts, validator common.Address) (uint64, error) {
	var out []interface{}
	err := _DposContract.contract.Call(opts, &out, "getValidatorEligibleVotesCount", validator)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetValidatorEligibleVotesCount is a free data retrieval call binding the contract method 0x618e3862.
//
// Solidity: function getValidatorEligibleVotesCount(address validator) view returns(uint64)
func (_DposContract *DposContractSession) GetValidatorEligibleVotesCount(validator common.Address) (uint64, error) {
	return _DposContract.Contract.GetValidatorEligibleVotesCount(&_DposContract.CallOpts, validator)
}

// GetValidatorEligibleVotesCount is a free data retrieval call binding the contract method 0x618e3862.
//
// Solidity: function getValidatorEligibleVotesCount(address validator) view returns(uint64)
func (_DposContract *DposContractCallerSession) GetValidatorEligibleVotesCount(validator common.Address) (uint64, error) {
	return _DposContract.Contract.GetValidatorEligibleVotesCount(&_DposContract.CallOpts, validator)
}

// GetValidators is a free data retrieval call binding the contract method 0x19d8024f.
//
// Solidity: function getValidators(uint32 batch) view returns((address,(uint256,uint256,uint16,uint64,uint16,address,string,string))[] validators, bool end)
func (_DposContract *DposContractCaller) GetValidators(opts *bind.CallOpts, batch uint32) (struct {
	Validators []DposInterfaceValidatorData
	End        bool
}, error) {
	var out []interface{}
	err := _DposContract.contract.Call(opts, &out, "getValidators", batch)

	outstruct := new(struct {
		Validators []DposInterfaceValidatorData
		End        bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Validators = *abi.ConvertType(out[0], new([]DposInterfaceValidatorData)).(*[]DposInterfaceValidatorData)
	outstruct.End = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetValidators is a free data retrieval call binding the contract method 0x19d8024f.
//
// Solidity: function getValidators(uint32 batch) view returns((address,(uint256,uint256,uint16,uint64,uint16,address,string,string))[] validators, bool end)
func (_DposContract *DposContractSession) GetValidators(batch uint32) (struct {
	Validators []DposInterfaceValidatorData
	End        bool
}, error) {
	return _DposContract.Contract.GetValidators(&_DposContract.CallOpts, batch)
}

// GetValidators is a free data retrieval call binding the contract method 0x19d8024f.
//
// Solidity: function getValidators(uint32 batch) view returns((address,(uint256,uint256,uint16,uint64,uint16,address,string,string))[] validators, bool end)
func (_DposContract *DposContractCallerSession) GetValidators(batch uint32) (struct {
	Validators []DposInterfaceValidatorData
	End        bool
}, error) {
	return _DposContract.Contract.GetValidators(&_DposContract.CallOpts, batch)
}

// GetValidatorsFor is a free data retrieval call binding the contract method 0x724ac6b0.
//
// Solidity: function getValidatorsFor(address owner, uint32 batch) view returns((address,(uint256,uint256,uint16,uint64,uint16,address,string,string))[] validators, bool end)
func (_DposContract *DposContractCaller) GetValidatorsFor(opts *bind.CallOpts, owner common.Address, batch uint32) (struct {
	Validators []DposInterfaceValidatorData
	End        bool
}, error) {
	var out []interface{}
	err := _DposContract.contract.Call(opts, &out, "getValidatorsFor", owner, batch)

	outstruct := new(struct {
		Validators []DposInterfaceValidatorData
		End        bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Validators = *abi.ConvertType(out[0], new([]DposInterfaceValidatorData)).(*[]DposInterfaceValidatorData)
	outstruct.End = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetValidatorsFor is a free data retrieval call binding the contract method 0x724ac6b0.
//
// Solidity: function getValidatorsFor(address owner, uint32 batch) view returns((address,(uint256,uint256,uint16,uint64,uint16,address,string,string))[] validators, bool end)
func (_DposContract *DposContractSession) GetValidatorsFor(owner common.Address, batch uint32) (struct {
	Validators []DposInterfaceValidatorData
	End        bool
}, error) {
	return _DposContract.Contract.GetValidatorsFor(&_DposContract.CallOpts, owner, batch)
}

// GetValidatorsFor is a free data retrieval call binding the contract method 0x724ac6b0.
//
// Solidity: function getValidatorsFor(address owner, uint32 batch) view returns((address,(uint256,uint256,uint16,uint64,uint16,address,string,string))[] validators, bool end)
func (_DposContract *DposContractCallerSession) GetValidatorsFor(owner common.Address, batch uint32) (struct {
	Validators []DposInterfaceValidatorData
	End        bool
}, error) {
	return _DposContract.Contract.GetValidatorsFor(&_DposContract.CallOpts, owner, batch)
}

// IsValidatorEligible is a free data retrieval call binding the contract method 0xf3094e90.
//
// Solidity: function isValidatorEligible(address validator) view returns(bool)
func (_DposContract *DposContractCaller) IsValidatorEligible(opts *bind.CallOpts, validator common.Address) (bool, error) {
	var out []interface{}
	err := _DposContract.contract.Call(opts, &out, "isValidatorEligible", validator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidatorEligible is a free data retrieval call binding the contract method 0xf3094e90.
//
// Solidity: function isValidatorEligible(address validator) view returns(bool)
func (_DposContract *DposContractSession) IsValidatorEligible(validator common.Address) (bool, error) {
	return _DposContract.Contract.IsValidatorEligible(&_DposContract.CallOpts, validator)
}

// IsValidatorEligible is a free data retrieval call binding the contract method 0xf3094e90.
//
// Solidity: function isValidatorEligible(address validator) view returns(bool)
func (_DposContract *DposContractCallerSession) IsValidatorEligible(validator common.Address) (bool, error) {
	return _DposContract.Contract.IsValidatorEligible(&_DposContract.CallOpts, validator)
}

// CancelUndelegate is a paid mutator transaction binding the contract method 0x399ff554.
//
// Solidity: function cancelUndelegate(address validator) returns()
func (_DposContract *DposContractTransactor) CancelUndelegate(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _DposContract.contract.Transact(opts, "cancelUndelegate", validator)
}

// CancelUndelegate is a paid mutator transaction binding the contract method 0x399ff554.
//
// Solidity: function cancelUndelegate(address validator) returns()
func (_DposContract *DposContractSession) CancelUndelegate(validator common.Address) (*types.Transaction, error) {
	return _DposContract.Contract.CancelUndelegate(&_DposContract.TransactOpts, validator)
}

// CancelUndelegate is a paid mutator transaction binding the contract method 0x399ff554.
//
// Solidity: function cancelUndelegate(address validator) returns()
func (_DposContract *DposContractTransactorSession) CancelUndelegate(validator common.Address) (*types.Transaction, error) {
	return _DposContract.Contract.CancelUndelegate(&_DposContract.TransactOpts, validator)
}

// ClaimAllRewards is a paid mutator transaction binding the contract method 0x09b72e00.
//
// Solidity: function claimAllRewards(uint32 batch) returns(bool end)
func (_DposContract *DposContractTransactor) ClaimAllRewards(opts *bind.TransactOpts, batch uint32) (*types.Transaction, error) {
	return _DposContract.contract.Transact(opts, "claimAllRewards", batch)
}

// ClaimAllRewards is a paid mutator transaction binding the contract method 0x09b72e00.
//
// Solidity: function claimAllRewards(uint32 batch) returns(bool end)
func (_DposContract *DposContractSession) ClaimAllRewards(batch uint32) (*types.Transaction, error) {
	return _DposContract.Contract.ClaimAllRewards(&_DposContract.TransactOpts, batch)
}

// ClaimAllRewards is a paid mutator transaction binding the contract method 0x09b72e00.
//
// Solidity: function claimAllRewards(uint32 batch) returns(bool end)
func (_DposContract *DposContractTransactorSession) ClaimAllRewards(batch uint32) (*types.Transaction, error) {
	return _DposContract.Contract.ClaimAllRewards(&_DposContract.TransactOpts, batch)
}

// ClaimCommissionRewards is a paid mutator transaction binding the contract method 0xd0eebfe2.
//
// Solidity: function claimCommissionRewards(address validator) returns()
func (_DposContract *DposContractTransactor) ClaimCommissionRewards(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _DposContract.contract.Transact(opts, "claimCommissionRewards", validator)
}

// ClaimCommissionRewards is a paid mutator transaction binding the contract method 0xd0eebfe2.
//
// Solidity: function claimCommissionRewards(address validator) returns()
func (_DposContract *DposContractSession) ClaimCommissionRewards(validator common.Address) (*types.Transaction, error) {
	return _DposContract.Contract.ClaimCommissionRewards(&_DposContract.TransactOpts, validator)
}

// ClaimCommissionRewards is a paid mutator transaction binding the contract method 0xd0eebfe2.
//
// Solidity: function claimCommissionRewards(address validator) returns()
func (_DposContract *DposContractTransactorSession) ClaimCommissionRewards(validator common.Address) (*types.Transaction, error) {
	return _DposContract.Contract.ClaimCommissionRewards(&_DposContract.TransactOpts, validator)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xef5cfb8c.
//
// Solidity: function claimRewards(address validator) returns()
func (_DposContract *DposContractTransactor) ClaimRewards(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _DposContract.contract.Transact(opts, "claimRewards", validator)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xef5cfb8c.
//
// Solidity: function claimRewards(address validator) returns()
func (_DposContract *DposContractSession) ClaimRewards(validator common.Address) (*types.Transaction, error) {
	return _DposContract.Contract.ClaimRewards(&_DposContract.TransactOpts, validator)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xef5cfb8c.
//
// Solidity: function claimRewards(address validator) returns()
func (_DposContract *DposContractTransactorSession) ClaimRewards(validator common.Address) (*types.Transaction, error) {
	return _DposContract.Contract.ClaimRewards(&_DposContract.TransactOpts, validator)
}

// ConfirmUndelegate is a paid mutator transaction binding the contract method 0x45a02561.
//
// Solidity: function confirmUndelegate(address validator) returns()
func (_DposContract *DposContractTransactor) ConfirmUndelegate(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _DposContract.contract.Transact(opts, "confirmUndelegate", validator)
}

// ConfirmUndelegate is a paid mutator transaction binding the contract method 0x45a02561.
//
// Solidity: function confirmUndelegate(address validator) returns()
func (_DposContract *DposContractSession) ConfirmUndelegate(validator common.Address) (*types.Transaction, error) {
	return _DposContract.Contract.ConfirmUndelegate(&_DposContract.TransactOpts, validator)
}

// ConfirmUndelegate is a paid mutator transaction binding the contract method 0x45a02561.
//
// Solidity: function confirmUndelegate(address validator) returns()
func (_DposContract *DposContractTransactorSession) ConfirmUndelegate(validator common.Address) (*types.Transaction, error) {
	return _DposContract.Contract.ConfirmUndelegate(&_DposContract.TransactOpts, validator)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address validator) payable returns()
func (_DposContract *DposContractTransactor) Delegate(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _DposContract.contract.Transact(opts, "delegate", validator)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address validator) payable returns()
func (_DposContract *DposContractSession) Delegate(validator common.Address) (*types.Transaction, error) {
	return _DposContract.Contract.Delegate(&_DposContract.TransactOpts, validator)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address validator) payable returns()
func (_DposContract *DposContractTransactorSession) Delegate(validator common.Address) (*types.Transaction, error) {
	return _DposContract.Contract.Delegate(&_DposContract.TransactOpts, validator)
}

// ReDelegate is a paid mutator transaction binding the contract method 0x703812cc.
//
// Solidity: function reDelegate(address validator_from, address validator_to, uint256 amount) returns()
func (_DposContract *DposContractTransactor) ReDelegate(opts *bind.TransactOpts, validator_from common.Address, validator_to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DposContract.contract.Transact(opts, "reDelegate", validator_from, validator_to, amount)
}

// ReDelegate is a paid mutator transaction binding the contract method 0x703812cc.
//
// Solidity: function reDelegate(address validator_from, address validator_to, uint256 amount) returns()
func (_DposContract *DposContractSession) ReDelegate(validator_from common.Address, validator_to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DposContract.Contract.ReDelegate(&_DposContract.TransactOpts, validator_from, validator_to, amount)
}

// ReDelegate is a paid mutator transaction binding the contract method 0x703812cc.
//
// Solidity: function reDelegate(address validator_from, address validator_to, uint256 amount) returns()
func (_DposContract *DposContractTransactorSession) ReDelegate(validator_from common.Address, validator_to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DposContract.Contract.ReDelegate(&_DposContract.TransactOpts, validator_from, validator_to, amount)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0xd6fdc127.
//
// Solidity: function registerValidator(address validator, bytes proof, bytes vrf_key, uint16 commission, string description, string endpoint) payable returns()
func (_DposContract *DposContractTransactor) RegisterValidator(opts *bind.TransactOpts, validator common.Address, proof []byte, vrf_key []byte, commission uint16, description string, endpoint string) (*types.Transaction, error) {
	return _DposContract.contract.Transact(opts, "registerValidator", validator, proof, vrf_key, commission, description, endpoint)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0xd6fdc127.
//
// Solidity: function registerValidator(address validator, bytes proof, bytes vrf_key, uint16 commission, string description, string endpoint) payable returns()
func (_DposContract *DposContractSession) RegisterValidator(validator common.Address, proof []byte, vrf_key []byte, commission uint16, description string, endpoint string) (*types.Transaction, error) {
	return _DposContract.Contract.RegisterValidator(&_DposContract.TransactOpts, validator, proof, vrf_key, commission, description, endpoint)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0xd6fdc127.
//
// Solidity: function registerValidator(address validator, bytes proof, bytes vrf_key, uint16 commission, string description, string endpoint) payable returns()
func (_DposContract *DposContractTransactorSession) RegisterValidator(validator common.Address, proof []byte, vrf_key []byte, commission uint16, description string, endpoint string) (*types.Transaction, error) {
	return _DposContract.Contract.RegisterValidator(&_DposContract.TransactOpts, validator, proof, vrf_key, commission, description, endpoint)
}

// SetCommission is a paid mutator transaction binding the contract method 0xf000322c.
//
// Solidity: function setCommission(address validator, uint16 commission) returns()
func (_DposContract *DposContractTransactor) SetCommission(opts *bind.TransactOpts, validator common.Address, commission uint16) (*types.Transaction, error) {
	return _DposContract.contract.Transact(opts, "setCommission", validator, commission)
}

// SetCommission is a paid mutator transaction binding the contract method 0xf000322c.
//
// Solidity: function setCommission(address validator, uint16 commission) returns()
func (_DposContract *DposContractSession) SetCommission(validator common.Address, commission uint16) (*types.Transaction, error) {
	return _DposContract.Contract.SetCommission(&_DposContract.TransactOpts, validator, commission)
}

// SetCommission is a paid mutator transaction binding the contract method 0xf000322c.
//
// Solidity: function setCommission(address validator, uint16 commission) returns()
func (_DposContract *DposContractTransactorSession) SetCommission(validator common.Address, commission uint16) (*types.Transaction, error) {
	return _DposContract.Contract.SetCommission(&_DposContract.TransactOpts, validator, commission)
}

// SetValidatorInfo is a paid mutator transaction binding the contract method 0x0babea4c.
//
// Solidity: function setValidatorInfo(address validator, string description, string endpoint) returns()
func (_DposContract *DposContractTransactor) SetValidatorInfo(opts *bind.TransactOpts, validator common.Address, description string, endpoint string) (*types.Transaction, error) {
	return _DposContract.contract.Transact(opts, "setValidatorInfo", validator, description, endpoint)
}

// SetValidatorInfo is a paid mutator transaction binding the contract method 0x0babea4c.
//
// Solidity: function setValidatorInfo(address validator, string description, string endpoint) returns()
func (_DposContract *DposContractSession) SetValidatorInfo(validator common.Address, description string, endpoint string) (*types.Transaction, error) {
	return _DposContract.Contract.SetValidatorInfo(&_DposContract.TransactOpts, validator, description, endpoint)
}

// SetValidatorInfo is a paid mutator transaction binding the contract method 0x0babea4c.
//
// Solidity: function setValidatorInfo(address validator, string description, string endpoint) returns()
func (_DposContract *DposContractTransactorSession) SetValidatorInfo(validator common.Address, description string, endpoint string) (*types.Transaction, error) {
	return _DposContract.Contract.SetValidatorInfo(&_DposContract.TransactOpts, validator, description, endpoint)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address validator, uint256 amount) returns()
func (_DposContract *DposContractTransactor) Undelegate(opts *bind.TransactOpts, validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DposContract.contract.Transact(opts, "undelegate", validator, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address validator, uint256 amount) returns()
func (_DposContract *DposContractSession) Undelegate(validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DposContract.Contract.Undelegate(&_DposContract.TransactOpts, validator, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address validator, uint256 amount) returns()
func (_DposContract *DposContractTransactorSession) Undelegate(validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DposContract.Contract.Undelegate(&_DposContract.TransactOpts, validator, amount)
}

// DposContractCommissionRewardsClaimedIterator is returned from FilterCommissionRewardsClaimed and is used to iterate over the raw logs and unpacked data for CommissionRewardsClaimed events raised by the DposContract contract.
type DposContractCommissionRewardsClaimedIterator struct {
	Event *DposContractCommissionRewardsClaimed // Event containing the contract specifics and raw log

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
func (it *DposContractCommissionRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DposContractCommissionRewardsClaimed)
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
		it.Event = new(DposContractCommissionRewardsClaimed)
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
func (it *DposContractCommissionRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DposContractCommissionRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DposContractCommissionRewardsClaimed represents a CommissionRewardsClaimed event raised by the DposContract contract.
type DposContractCommissionRewardsClaimed struct {
	Account   common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCommissionRewardsClaimed is a free log retrieval operation binding the contract event 0xf0ec9e0f6add850a1738c5822244e26ffc3d1f14da7537aa240582b25af12ad0.
//
// Solidity: event CommissionRewardsClaimed(address indexed account, address indexed validator, uint256 amount)
func (_DposContract *DposContractFilterer) FilterCommissionRewardsClaimed(opts *bind.FilterOpts, account []common.Address, validator []common.Address) (*DposContractCommissionRewardsClaimedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposContract.contract.FilterLogs(opts, "CommissionRewardsClaimed", accountRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &DposContractCommissionRewardsClaimedIterator{contract: _DposContract.contract, event: "CommissionRewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchCommissionRewardsClaimed is a free log subscription operation binding the contract event 0xf0ec9e0f6add850a1738c5822244e26ffc3d1f14da7537aa240582b25af12ad0.
//
// Solidity: event CommissionRewardsClaimed(address indexed account, address indexed validator, uint256 amount)
func (_DposContract *DposContractFilterer) WatchCommissionRewardsClaimed(opts *bind.WatchOpts, sink chan<- *DposContractCommissionRewardsClaimed, account []common.Address, validator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposContract.contract.WatchLogs(opts, "CommissionRewardsClaimed", accountRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DposContractCommissionRewardsClaimed)
				if err := _DposContract.contract.UnpackLog(event, "CommissionRewardsClaimed", log); err != nil {
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

// ParseCommissionRewardsClaimed is a log parse operation binding the contract event 0xf0ec9e0f6add850a1738c5822244e26ffc3d1f14da7537aa240582b25af12ad0.
//
// Solidity: event CommissionRewardsClaimed(address indexed account, address indexed validator, uint256 amount)
func (_DposContract *DposContractFilterer) ParseCommissionRewardsClaimed(log types.Log) (*DposContractCommissionRewardsClaimed, error) {
	event := new(DposContractCommissionRewardsClaimed)
	if err := _DposContract.contract.UnpackLog(event, "CommissionRewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DposContractCommissionSetIterator is returned from FilterCommissionSet and is used to iterate over the raw logs and unpacked data for CommissionSet events raised by the DposContract contract.
type DposContractCommissionSetIterator struct {
	Event *DposContractCommissionSet // Event containing the contract specifics and raw log

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
func (it *DposContractCommissionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DposContractCommissionSet)
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
		it.Event = new(DposContractCommissionSet)
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
func (it *DposContractCommissionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DposContractCommissionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DposContractCommissionSet represents a CommissionSet event raised by the DposContract contract.
type DposContractCommissionSet struct {
	Validator  common.Address
	Commission uint16
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCommissionSet is a free log retrieval operation binding the contract event 0xc909daf778d180f43dac53b55d0de934d2f1e0b70412ca274982e4e6e894eb1a.
//
// Solidity: event CommissionSet(address indexed validator, uint16 commission)
func (_DposContract *DposContractFilterer) FilterCommissionSet(opts *bind.FilterOpts, validator []common.Address) (*DposContractCommissionSetIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposContract.contract.FilterLogs(opts, "CommissionSet", validatorRule)
	if err != nil {
		return nil, err
	}
	return &DposContractCommissionSetIterator{contract: _DposContract.contract, event: "CommissionSet", logs: logs, sub: sub}, nil
}

// WatchCommissionSet is a free log subscription operation binding the contract event 0xc909daf778d180f43dac53b55d0de934d2f1e0b70412ca274982e4e6e894eb1a.
//
// Solidity: event CommissionSet(address indexed validator, uint16 commission)
func (_DposContract *DposContractFilterer) WatchCommissionSet(opts *bind.WatchOpts, sink chan<- *DposContractCommissionSet, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposContract.contract.WatchLogs(opts, "CommissionSet", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DposContractCommissionSet)
				if err := _DposContract.contract.UnpackLog(event, "CommissionSet", log); err != nil {
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

// ParseCommissionSet is a log parse operation binding the contract event 0xc909daf778d180f43dac53b55d0de934d2f1e0b70412ca274982e4e6e894eb1a.
//
// Solidity: event CommissionSet(address indexed validator, uint16 commission)
func (_DposContract *DposContractFilterer) ParseCommissionSet(log types.Log) (*DposContractCommissionSet, error) {
	event := new(DposContractCommissionSet)
	if err := _DposContract.contract.UnpackLog(event, "CommissionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DposContractDelegatedIterator is returned from FilterDelegated and is used to iterate over the raw logs and unpacked data for Delegated events raised by the DposContract contract.
type DposContractDelegatedIterator struct {
	Event *DposContractDelegated // Event containing the contract specifics and raw log

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
func (it *DposContractDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DposContractDelegated)
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
		it.Event = new(DposContractDelegated)
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
func (it *DposContractDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DposContractDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DposContractDelegated represents a Delegated event raised by the DposContract contract.
type DposContractDelegated struct {
	Delegator common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegated is a free log retrieval operation binding the contract event 0xe5541a6b6103d4fa7e021ed54fad39c66f27a76bd13d374cf6240ae6bd0bb72b.
//
// Solidity: event Delegated(address indexed delegator, address indexed validator, uint256 amount)
func (_DposContract *DposContractFilterer) FilterDelegated(opts *bind.FilterOpts, delegator []common.Address, validator []common.Address) (*DposContractDelegatedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposContract.contract.FilterLogs(opts, "Delegated", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &DposContractDelegatedIterator{contract: _DposContract.contract, event: "Delegated", logs: logs, sub: sub}, nil
}

// WatchDelegated is a free log subscription operation binding the contract event 0xe5541a6b6103d4fa7e021ed54fad39c66f27a76bd13d374cf6240ae6bd0bb72b.
//
// Solidity: event Delegated(address indexed delegator, address indexed validator, uint256 amount)
func (_DposContract *DposContractFilterer) WatchDelegated(opts *bind.WatchOpts, sink chan<- *DposContractDelegated, delegator []common.Address, validator []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposContract.contract.WatchLogs(opts, "Delegated", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DposContractDelegated)
				if err := _DposContract.contract.UnpackLog(event, "Delegated", log); err != nil {
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

// ParseDelegated is a log parse operation binding the contract event 0xe5541a6b6103d4fa7e021ed54fad39c66f27a76bd13d374cf6240ae6bd0bb72b.
//
// Solidity: event Delegated(address indexed delegator, address indexed validator, uint256 amount)
func (_DposContract *DposContractFilterer) ParseDelegated(log types.Log) (*DposContractDelegated, error) {
	event := new(DposContractDelegated)
	if err := _DposContract.contract.UnpackLog(event, "Delegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DposContractRedelegatedIterator is returned from FilterRedelegated and is used to iterate over the raw logs and unpacked data for Redelegated events raised by the DposContract contract.
type DposContractRedelegatedIterator struct {
	Event *DposContractRedelegated // Event containing the contract specifics and raw log

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
func (it *DposContractRedelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DposContractRedelegated)
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
		it.Event = new(DposContractRedelegated)
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
func (it *DposContractRedelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DposContractRedelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DposContractRedelegated represents a Redelegated event raised by the DposContract contract.
type DposContractRedelegated struct {
	Delegator common.Address
	From      common.Address
	To        common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRedelegated is a free log retrieval operation binding the contract event 0x12e144c27d0bad08abc77c66a640b5cf15a03a93f6582f40de6932b033a5fa5e.
//
// Solidity: event Redelegated(address indexed delegator, address indexed from, address indexed to, uint256 amount)
func (_DposContract *DposContractFilterer) FilterRedelegated(opts *bind.FilterOpts, delegator []common.Address, from []common.Address, to []common.Address) (*DposContractRedelegatedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DposContract.contract.FilterLogs(opts, "Redelegated", delegatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DposContractRedelegatedIterator{contract: _DposContract.contract, event: "Redelegated", logs: logs, sub: sub}, nil
}

// WatchRedelegated is a free log subscription operation binding the contract event 0x12e144c27d0bad08abc77c66a640b5cf15a03a93f6582f40de6932b033a5fa5e.
//
// Solidity: event Redelegated(address indexed delegator, address indexed from, address indexed to, uint256 amount)
func (_DposContract *DposContractFilterer) WatchRedelegated(opts *bind.WatchOpts, sink chan<- *DposContractRedelegated, delegator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DposContract.contract.WatchLogs(opts, "Redelegated", delegatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DposContractRedelegated)
				if err := _DposContract.contract.UnpackLog(event, "Redelegated", log); err != nil {
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

// ParseRedelegated is a log parse operation binding the contract event 0x12e144c27d0bad08abc77c66a640b5cf15a03a93f6582f40de6932b033a5fa5e.
//
// Solidity: event Redelegated(address indexed delegator, address indexed from, address indexed to, uint256 amount)
func (_DposContract *DposContractFilterer) ParseRedelegated(log types.Log) (*DposContractRedelegated, error) {
	event := new(DposContractRedelegated)
	if err := _DposContract.contract.UnpackLog(event, "Redelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DposContractRewardsClaimedIterator is returned from FilterRewardsClaimed and is used to iterate over the raw logs and unpacked data for RewardsClaimed events raised by the DposContract contract.
type DposContractRewardsClaimedIterator struct {
	Event *DposContractRewardsClaimed // Event containing the contract specifics and raw log

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
func (it *DposContractRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DposContractRewardsClaimed)
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
		it.Event = new(DposContractRewardsClaimed)
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
func (it *DposContractRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DposContractRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DposContractRewardsClaimed represents a RewardsClaimed event raised by the DposContract contract.
type DposContractRewardsClaimed struct {
	Account   common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRewardsClaimed is a free log retrieval operation binding the contract event 0x9310ccfcb8de723f578a9e4282ea9f521f05ae40dc08f3068dfad528a65ee3c7.
//
// Solidity: event RewardsClaimed(address indexed account, address indexed validator, uint256 amount)
func (_DposContract *DposContractFilterer) FilterRewardsClaimed(opts *bind.FilterOpts, account []common.Address, validator []common.Address) (*DposContractRewardsClaimedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposContract.contract.FilterLogs(opts, "RewardsClaimed", accountRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &DposContractRewardsClaimedIterator{contract: _DposContract.contract, event: "RewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardsClaimed is a free log subscription operation binding the contract event 0x9310ccfcb8de723f578a9e4282ea9f521f05ae40dc08f3068dfad528a65ee3c7.
//
// Solidity: event RewardsClaimed(address indexed account, address indexed validator, uint256 amount)
func (_DposContract *DposContractFilterer) WatchRewardsClaimed(opts *bind.WatchOpts, sink chan<- *DposContractRewardsClaimed, account []common.Address, validator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposContract.contract.WatchLogs(opts, "RewardsClaimed", accountRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DposContractRewardsClaimed)
				if err := _DposContract.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
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

// ParseRewardsClaimed is a log parse operation binding the contract event 0x9310ccfcb8de723f578a9e4282ea9f521f05ae40dc08f3068dfad528a65ee3c7.
//
// Solidity: event RewardsClaimed(address indexed account, address indexed validator, uint256 amount)
func (_DposContract *DposContractFilterer) ParseRewardsClaimed(log types.Log) (*DposContractRewardsClaimed, error) {
	event := new(DposContractRewardsClaimed)
	if err := _DposContract.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DposContractUndelegateCanceledIterator is returned from FilterUndelegateCanceled and is used to iterate over the raw logs and unpacked data for UndelegateCanceled events raised by the DposContract contract.
type DposContractUndelegateCanceledIterator struct {
	Event *DposContractUndelegateCanceled // Event containing the contract specifics and raw log

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
func (it *DposContractUndelegateCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DposContractUndelegateCanceled)
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
		it.Event = new(DposContractUndelegateCanceled)
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
func (it *DposContractUndelegateCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DposContractUndelegateCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DposContractUndelegateCanceled represents a UndelegateCanceled event raised by the DposContract contract.
type DposContractUndelegateCanceled struct {
	Delegator common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUndelegateCanceled is a free log retrieval operation binding the contract event 0xfc25f8a919d19f2c2dfce21115718abc9ef2b1e0c9218a488f614c75be4184b7.
//
// Solidity: event UndelegateCanceled(address indexed delegator, address indexed validator, uint256 amount)
func (_DposContract *DposContractFilterer) FilterUndelegateCanceled(opts *bind.FilterOpts, delegator []common.Address, validator []common.Address) (*DposContractUndelegateCanceledIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposContract.contract.FilterLogs(opts, "UndelegateCanceled", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &DposContractUndelegateCanceledIterator{contract: _DposContract.contract, event: "UndelegateCanceled", logs: logs, sub: sub}, nil
}

// WatchUndelegateCanceled is a free log subscription operation binding the contract event 0xfc25f8a919d19f2c2dfce21115718abc9ef2b1e0c9218a488f614c75be4184b7.
//
// Solidity: event UndelegateCanceled(address indexed delegator, address indexed validator, uint256 amount)
func (_DposContract *DposContractFilterer) WatchUndelegateCanceled(opts *bind.WatchOpts, sink chan<- *DposContractUndelegateCanceled, delegator []common.Address, validator []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposContract.contract.WatchLogs(opts, "UndelegateCanceled", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DposContractUndelegateCanceled)
				if err := _DposContract.contract.UnpackLog(event, "UndelegateCanceled", log); err != nil {
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

// ParseUndelegateCanceled is a log parse operation binding the contract event 0xfc25f8a919d19f2c2dfce21115718abc9ef2b1e0c9218a488f614c75be4184b7.
//
// Solidity: event UndelegateCanceled(address indexed delegator, address indexed validator, uint256 amount)
func (_DposContract *DposContractFilterer) ParseUndelegateCanceled(log types.Log) (*DposContractUndelegateCanceled, error) {
	event := new(DposContractUndelegateCanceled)
	if err := _DposContract.contract.UnpackLog(event, "UndelegateCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DposContractUndelegateConfirmedIterator is returned from FilterUndelegateConfirmed and is used to iterate over the raw logs and unpacked data for UndelegateConfirmed events raised by the DposContract contract.
type DposContractUndelegateConfirmedIterator struct {
	Event *DposContractUndelegateConfirmed // Event containing the contract specifics and raw log

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
func (it *DposContractUndelegateConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DposContractUndelegateConfirmed)
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
		it.Event = new(DposContractUndelegateConfirmed)
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
func (it *DposContractUndelegateConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DposContractUndelegateConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DposContractUndelegateConfirmed represents a UndelegateConfirmed event raised by the DposContract contract.
type DposContractUndelegateConfirmed struct {
	Delegator common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUndelegateConfirmed is a free log retrieval operation binding the contract event 0xf8bef3a6fe3b4c932b5b51c6472a89f171d039f4bacf18cff632208938bf0426.
//
// Solidity: event UndelegateConfirmed(address indexed delegator, address indexed validator, uint256 amount)
func (_DposContract *DposContractFilterer) FilterUndelegateConfirmed(opts *bind.FilterOpts, delegator []common.Address, validator []common.Address) (*DposContractUndelegateConfirmedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposContract.contract.FilterLogs(opts, "UndelegateConfirmed", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &DposContractUndelegateConfirmedIterator{contract: _DposContract.contract, event: "UndelegateConfirmed", logs: logs, sub: sub}, nil
}

// WatchUndelegateConfirmed is a free log subscription operation binding the contract event 0xf8bef3a6fe3b4c932b5b51c6472a89f171d039f4bacf18cff632208938bf0426.
//
// Solidity: event UndelegateConfirmed(address indexed delegator, address indexed validator, uint256 amount)
func (_DposContract *DposContractFilterer) WatchUndelegateConfirmed(opts *bind.WatchOpts, sink chan<- *DposContractUndelegateConfirmed, delegator []common.Address, validator []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposContract.contract.WatchLogs(opts, "UndelegateConfirmed", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DposContractUndelegateConfirmed)
				if err := _DposContract.contract.UnpackLog(event, "UndelegateConfirmed", log); err != nil {
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

// ParseUndelegateConfirmed is a log parse operation binding the contract event 0xf8bef3a6fe3b4c932b5b51c6472a89f171d039f4bacf18cff632208938bf0426.
//
// Solidity: event UndelegateConfirmed(address indexed delegator, address indexed validator, uint256 amount)
func (_DposContract *DposContractFilterer) ParseUndelegateConfirmed(log types.Log) (*DposContractUndelegateConfirmed, error) {
	event := new(DposContractUndelegateConfirmed)
	if err := _DposContract.contract.UnpackLog(event, "UndelegateConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DposContractUndelegatedIterator is returned from FilterUndelegated and is used to iterate over the raw logs and unpacked data for Undelegated events raised by the DposContract contract.
type DposContractUndelegatedIterator struct {
	Event *DposContractUndelegated // Event containing the contract specifics and raw log

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
func (it *DposContractUndelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DposContractUndelegated)
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
		it.Event = new(DposContractUndelegated)
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
func (it *DposContractUndelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DposContractUndelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DposContractUndelegated represents a Undelegated event raised by the DposContract contract.
type DposContractUndelegated struct {
	Delegator common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUndelegated is a free log retrieval operation binding the contract event 0x4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c.
//
// Solidity: event Undelegated(address indexed delegator, address indexed validator, uint256 amount)
func (_DposContract *DposContractFilterer) FilterUndelegated(opts *bind.FilterOpts, delegator []common.Address, validator []common.Address) (*DposContractUndelegatedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposContract.contract.FilterLogs(opts, "Undelegated", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &DposContractUndelegatedIterator{contract: _DposContract.contract, event: "Undelegated", logs: logs, sub: sub}, nil
}

// WatchUndelegated is a free log subscription operation binding the contract event 0x4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c.
//
// Solidity: event Undelegated(address indexed delegator, address indexed validator, uint256 amount)
func (_DposContract *DposContractFilterer) WatchUndelegated(opts *bind.WatchOpts, sink chan<- *DposContractUndelegated, delegator []common.Address, validator []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposContract.contract.WatchLogs(opts, "Undelegated", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DposContractUndelegated)
				if err := _DposContract.contract.UnpackLog(event, "Undelegated", log); err != nil {
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
// Solidity: event Undelegated(address indexed delegator, address indexed validator, uint256 amount)
func (_DposContract *DposContractFilterer) ParseUndelegated(log types.Log) (*DposContractUndelegated, error) {
	event := new(DposContractUndelegated)
	if err := _DposContract.contract.UnpackLog(event, "Undelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DposContractValidatorInfoSetIterator is returned from FilterValidatorInfoSet and is used to iterate over the raw logs and unpacked data for ValidatorInfoSet events raised by the DposContract contract.
type DposContractValidatorInfoSetIterator struct {
	Event *DposContractValidatorInfoSet // Event containing the contract specifics and raw log

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
func (it *DposContractValidatorInfoSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DposContractValidatorInfoSet)
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
		it.Event = new(DposContractValidatorInfoSet)
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
func (it *DposContractValidatorInfoSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DposContractValidatorInfoSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DposContractValidatorInfoSet represents a ValidatorInfoSet event raised by the DposContract contract.
type DposContractValidatorInfoSet struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorInfoSet is a free log retrieval operation binding the contract event 0x7aa20e1f59764c9066578febd688a51375adbd654aff86cef56593a17a99071d.
//
// Solidity: event ValidatorInfoSet(address indexed validator)
func (_DposContract *DposContractFilterer) FilterValidatorInfoSet(opts *bind.FilterOpts, validator []common.Address) (*DposContractValidatorInfoSetIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposContract.contract.FilterLogs(opts, "ValidatorInfoSet", validatorRule)
	if err != nil {
		return nil, err
	}
	return &DposContractValidatorInfoSetIterator{contract: _DposContract.contract, event: "ValidatorInfoSet", logs: logs, sub: sub}, nil
}

// WatchValidatorInfoSet is a free log subscription operation binding the contract event 0x7aa20e1f59764c9066578febd688a51375adbd654aff86cef56593a17a99071d.
//
// Solidity: event ValidatorInfoSet(address indexed validator)
func (_DposContract *DposContractFilterer) WatchValidatorInfoSet(opts *bind.WatchOpts, sink chan<- *DposContractValidatorInfoSet, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposContract.contract.WatchLogs(opts, "ValidatorInfoSet", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DposContractValidatorInfoSet)
				if err := _DposContract.contract.UnpackLog(event, "ValidatorInfoSet", log); err != nil {
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

// ParseValidatorInfoSet is a log parse operation binding the contract event 0x7aa20e1f59764c9066578febd688a51375adbd654aff86cef56593a17a99071d.
//
// Solidity: event ValidatorInfoSet(address indexed validator)
func (_DposContract *DposContractFilterer) ParseValidatorInfoSet(log types.Log) (*DposContractValidatorInfoSet, error) {
	event := new(DposContractValidatorInfoSet)
	if err := _DposContract.contract.UnpackLog(event, "ValidatorInfoSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DposContractValidatorRegisteredIterator is returned from FilterValidatorRegistered and is used to iterate over the raw logs and unpacked data for ValidatorRegistered events raised by the DposContract contract.
type DposContractValidatorRegisteredIterator struct {
	Event *DposContractValidatorRegistered // Event containing the contract specifics and raw log

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
func (it *DposContractValidatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DposContractValidatorRegistered)
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
		it.Event = new(DposContractValidatorRegistered)
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
func (it *DposContractValidatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DposContractValidatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DposContractValidatorRegistered represents a ValidatorRegistered event raised by the DposContract contract.
type DposContractValidatorRegistered struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorRegistered is a free log retrieval operation binding the contract event 0xd09501348473474a20c772c79c653e1fd7e8b437e418fe235d277d2c88853251.
//
// Solidity: event ValidatorRegistered(address indexed validator)
func (_DposContract *DposContractFilterer) FilterValidatorRegistered(opts *bind.FilterOpts, validator []common.Address) (*DposContractValidatorRegisteredIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposContract.contract.FilterLogs(opts, "ValidatorRegistered", validatorRule)
	if err != nil {
		return nil, err
	}
	return &DposContractValidatorRegisteredIterator{contract: _DposContract.contract, event: "ValidatorRegistered", logs: logs, sub: sub}, nil
}

// WatchValidatorRegistered is a free log subscription operation binding the contract event 0xd09501348473474a20c772c79c653e1fd7e8b437e418fe235d277d2c88853251.
//
// Solidity: event ValidatorRegistered(address indexed validator)
func (_DposContract *DposContractFilterer) WatchValidatorRegistered(opts *bind.WatchOpts, sink chan<- *DposContractValidatorRegistered, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposContract.contract.WatchLogs(opts, "ValidatorRegistered", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DposContractValidatorRegistered)
				if err := _DposContract.contract.UnpackLog(event, "ValidatorRegistered", log); err != nil {
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

// ParseValidatorRegistered is a log parse operation binding the contract event 0xd09501348473474a20c772c79c653e1fd7e8b437e418fe235d277d2c88853251.
//
// Solidity: event ValidatorRegistered(address indexed validator)
func (_DposContract *DposContractFilterer) ParseValidatorRegistered(log types.Log) (*DposContractValidatorRegistered, error) {
	event := new(DposContractValidatorRegistered)
	if err := _DposContract.contract.UnpackLog(event, "ValidatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
