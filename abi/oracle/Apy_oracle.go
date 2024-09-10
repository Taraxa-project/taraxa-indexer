// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package apy_oracle

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

// IApyOracleNodeData is an auto generated low-level Go binding around an user-defined struct.
type IApyOracleNodeData struct {
	Rating    *big.Int
	Account   common.Address
	FromBlock uint64
	ToBlock   uint64
	Rank      uint16
	Apy       uint16
}

// IApyOracleTentativeDelegation is an auto generated low-level Go binding around an user-defined struct.
type IApyOracleTentativeDelegation struct {
	Validator common.Address
	Amount    *big.Int
	Rating    *big.Int
}

// IApyOracleTentativeReDelegation is an auto generated low-level Go binding around an user-defined struct.
type IApyOracleTentativeReDelegation struct {
	From     common.Address
	To       common.Address
	Amount   *big.Int
	ToRating *big.Int
}

// ApyOracleMetaData contains all meta data concerning the ApyOracle contract.
var ApyOracleMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DATA_FEED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DPOS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractDposInterface\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"batchUpdateNodeData\",\"inputs\":[{\"name\":\"data\",\"type\":\"tuple[]\",\"internalType\":\"structIApyOracle.NodeData[]\",\"components\":[{\"name\":\"rating\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fromBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"toBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rank\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"apy\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"distrbutedRewards\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDataFeedAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodeCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodeData\",\"inputs\":[{\"name\":\"node\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIApyOracle.NodeData\",\"components\":[{\"name\":\"rating\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fromBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"toBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rank\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"apy\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodesForDelegation\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIApyOracle.TentativeDelegation[]\",\"components\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rating\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRebalanceList\",\"inputs\":[{\"name\":\"currentValidators\",\"type\":\"tuple[]\",\"internalType\":\"structIApyOracle.TentativeDelegation[]\",\"components\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rating\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIApyOracle.TentativeReDelegation[]\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"toRating\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"dataFeed\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dpos\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lara\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"logRewardDistribution\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"snapshotId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"maxValidatorStakeCapacity\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nodeCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nodes\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"rating\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fromBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"toBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rank\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"apy\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nodesList\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setLara\",\"inputs\":[{\"name\":\"_lara\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxValidatorStakeCapacity\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateNodeCount\",\"inputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateNodeData\",\"inputs\":[{\"name\":\"node\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"tuple\",\"internalType\":\"structIApyOracle.NodeData\",\"components\":[{\"name\":\"rating\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fromBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"toBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rank\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"apy\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxValidatorStakeUpdated\",\"inputs\":[{\"name\":\"maxValidatorStake\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeDataUpdated\",\"inputs\":[{\"name\":\"node\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"apy\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"pbftCount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
}

// ApyOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use ApyOracleMetaData.ABI instead.
var ApyOracleABI = ApyOracleMetaData.ABI

// ApyOracle is an auto generated Go binding around an Ethereum contract.
type ApyOracle struct {
	ApyOracleCaller     // Read-only binding to the contract
	ApyOracleTransactor // Write-only binding to the contract
	ApyOracleFilterer   // Log filterer for contract events
}

// ApyOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApyOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApyOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApyOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApyOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApyOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApyOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApyOracleSession struct {
	Contract     *ApyOracle        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApyOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApyOracleCallerSession struct {
	Contract *ApyOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ApyOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApyOracleTransactorSession struct {
	Contract     *ApyOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ApyOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApyOracleRaw struct {
	Contract *ApyOracle // Generic contract binding to access the raw methods on
}

// ApyOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApyOracleCallerRaw struct {
	Contract *ApyOracleCaller // Generic read-only contract binding to access the raw methods on
}

// ApyOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApyOracleTransactorRaw struct {
	Contract *ApyOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApyOracle creates a new instance of ApyOracle, bound to a specific deployed contract.
func NewApyOracle(address common.Address, backend bind.ContractBackend) (*ApyOracle, error) {
	contract, err := bindApyOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ApyOracle{ApyOracleCaller: ApyOracleCaller{contract: contract}, ApyOracleTransactor: ApyOracleTransactor{contract: contract}, ApyOracleFilterer: ApyOracleFilterer{contract: contract}}, nil
}

// NewApyOracleCaller creates a new read-only instance of ApyOracle, bound to a specific deployed contract.
func NewApyOracleCaller(address common.Address, caller bind.ContractCaller) (*ApyOracleCaller, error) {
	contract, err := bindApyOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApyOracleCaller{contract: contract}, nil
}

// NewApyOracleTransactor creates a new write-only instance of ApyOracle, bound to a specific deployed contract.
func NewApyOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*ApyOracleTransactor, error) {
	contract, err := bindApyOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApyOracleTransactor{contract: contract}, nil
}

// NewApyOracleFilterer creates a new log filterer instance of ApyOracle, bound to a specific deployed contract.
func NewApyOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*ApyOracleFilterer, error) {
	contract, err := bindApyOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApyOracleFilterer{contract: contract}, nil
}

// bindApyOracle binds a generic wrapper to an already deployed contract.
func bindApyOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ApyOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ApyOracle *ApyOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ApyOracle.Contract.ApyOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ApyOracle *ApyOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ApyOracle.Contract.ApyOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ApyOracle *ApyOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ApyOracle.Contract.ApyOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ApyOracle *ApyOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ApyOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ApyOracle *ApyOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ApyOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ApyOracle *ApyOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ApyOracle.Contract.contract.Transact(opts, method, params...)
}

// DATAFEED is a free data retrieval call binding the contract method 0x21db50a1.
//
// Solidity: function DATA_FEED() view returns(address)
func (_ApyOracle *ApyOracleCaller) DATAFEED(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ApyOracle.contract.Call(opts, &out, "DATA_FEED")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DATAFEED is a free data retrieval call binding the contract method 0x21db50a1.
//
// Solidity: function DATA_FEED() view returns(address)
func (_ApyOracle *ApyOracleSession) DATAFEED() (common.Address, error) {
	return _ApyOracle.Contract.DATAFEED(&_ApyOracle.CallOpts)
}

// DATAFEED is a free data retrieval call binding the contract method 0x21db50a1.
//
// Solidity: function DATA_FEED() view returns(address)
func (_ApyOracle *ApyOracleCallerSession) DATAFEED() (common.Address, error) {
	return _ApyOracle.Contract.DATAFEED(&_ApyOracle.CallOpts)
}

// DPOS is a free data retrieval call binding the contract method 0x6943935e.
//
// Solidity: function DPOS() view returns(address)
func (_ApyOracle *ApyOracleCaller) DPOS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ApyOracle.contract.Call(opts, &out, "DPOS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DPOS is a free data retrieval call binding the contract method 0x6943935e.
//
// Solidity: function DPOS() view returns(address)
func (_ApyOracle *ApyOracleSession) DPOS() (common.Address, error) {
	return _ApyOracle.Contract.DPOS(&_ApyOracle.CallOpts)
}

// DPOS is a free data retrieval call binding the contract method 0x6943935e.
//
// Solidity: function DPOS() view returns(address)
func (_ApyOracle *ApyOracleCallerSession) DPOS() (common.Address, error) {
	return _ApyOracle.Contract.DPOS(&_ApyOracle.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ApyOracle *ApyOracleCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ApyOracle.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ApyOracle *ApyOracleSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ApyOracle.Contract.UPGRADEINTERFACEVERSION(&_ApyOracle.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ApyOracle *ApyOracleCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ApyOracle.Contract.UPGRADEINTERFACEVERSION(&_ApyOracle.CallOpts)
}

// DistrbutedRewards is a free data retrieval call binding the contract method 0x2f78ebd6.
//
// Solidity: function distrbutedRewards(address , uint256 ) view returns(uint256)
func (_ApyOracle *ApyOracleCaller) DistrbutedRewards(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ApyOracle.contract.Call(opts, &out, "distrbutedRewards", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DistrbutedRewards is a free data retrieval call binding the contract method 0x2f78ebd6.
//
// Solidity: function distrbutedRewards(address , uint256 ) view returns(uint256)
func (_ApyOracle *ApyOracleSession) DistrbutedRewards(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _ApyOracle.Contract.DistrbutedRewards(&_ApyOracle.CallOpts, arg0, arg1)
}

// DistrbutedRewards is a free data retrieval call binding the contract method 0x2f78ebd6.
//
// Solidity: function distrbutedRewards(address , uint256 ) view returns(uint256)
func (_ApyOracle *ApyOracleCallerSession) DistrbutedRewards(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _ApyOracle.Contract.DistrbutedRewards(&_ApyOracle.CallOpts, arg0, arg1)
}

// GetDataFeedAddress is a free data retrieval call binding the contract method 0xa50d1850.
//
// Solidity: function getDataFeedAddress() view returns(address)
func (_ApyOracle *ApyOracleCaller) GetDataFeedAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ApyOracle.contract.Call(opts, &out, "getDataFeedAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDataFeedAddress is a free data retrieval call binding the contract method 0xa50d1850.
//
// Solidity: function getDataFeedAddress() view returns(address)
func (_ApyOracle *ApyOracleSession) GetDataFeedAddress() (common.Address, error) {
	return _ApyOracle.Contract.GetDataFeedAddress(&_ApyOracle.CallOpts)
}

// GetDataFeedAddress is a free data retrieval call binding the contract method 0xa50d1850.
//
// Solidity: function getDataFeedAddress() view returns(address)
func (_ApyOracle *ApyOracleCallerSession) GetDataFeedAddress() (common.Address, error) {
	return _ApyOracle.Contract.GetDataFeedAddress(&_ApyOracle.CallOpts)
}

// GetNodeCount is a free data retrieval call binding the contract method 0x39bf397e.
//
// Solidity: function getNodeCount() view returns(uint256)
func (_ApyOracle *ApyOracleCaller) GetNodeCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ApyOracle.contract.Call(opts, &out, "getNodeCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeCount is a free data retrieval call binding the contract method 0x39bf397e.
//
// Solidity: function getNodeCount() view returns(uint256)
func (_ApyOracle *ApyOracleSession) GetNodeCount() (*big.Int, error) {
	return _ApyOracle.Contract.GetNodeCount(&_ApyOracle.CallOpts)
}

// GetNodeCount is a free data retrieval call binding the contract method 0x39bf397e.
//
// Solidity: function getNodeCount() view returns(uint256)
func (_ApyOracle *ApyOracleCallerSession) GetNodeCount() (*big.Int, error) {
	return _ApyOracle.Contract.GetNodeCount(&_ApyOracle.CallOpts)
}

// GetNodeData is a free data retrieval call binding the contract method 0x38fc1090.
//
// Solidity: function getNodeData(address node) view returns((uint256,address,uint64,uint64,uint16,uint16))
func (_ApyOracle *ApyOracleCaller) GetNodeData(opts *bind.CallOpts, node common.Address) (IApyOracleNodeData, error) {
	var out []interface{}
	err := _ApyOracle.contract.Call(opts, &out, "getNodeData", node)

	if err != nil {
		return *new(IApyOracleNodeData), err
	}

	out0 := *abi.ConvertType(out[0], new(IApyOracleNodeData)).(*IApyOracleNodeData)

	return out0, err

}

// GetNodeData is a free data retrieval call binding the contract method 0x38fc1090.
//
// Solidity: function getNodeData(address node) view returns((uint256,address,uint64,uint64,uint16,uint16))
func (_ApyOracle *ApyOracleSession) GetNodeData(node common.Address) (IApyOracleNodeData, error) {
	return _ApyOracle.Contract.GetNodeData(&_ApyOracle.CallOpts, node)
}

// GetNodeData is a free data retrieval call binding the contract method 0x38fc1090.
//
// Solidity: function getNodeData(address node) view returns((uint256,address,uint64,uint64,uint16,uint16))
func (_ApyOracle *ApyOracleCallerSession) GetNodeData(node common.Address) (IApyOracleNodeData, error) {
	return _ApyOracle.Contract.GetNodeData(&_ApyOracle.CallOpts, node)
}

// Lara is a free data retrieval call binding the contract method 0x07f71eb2.
//
// Solidity: function lara() view returns(address)
func (_ApyOracle *ApyOracleCaller) Lara(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ApyOracle.contract.Call(opts, &out, "lara")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Lara is a free data retrieval call binding the contract method 0x07f71eb2.
//
// Solidity: function lara() view returns(address)
func (_ApyOracle *ApyOracleSession) Lara() (common.Address, error) {
	return _ApyOracle.Contract.Lara(&_ApyOracle.CallOpts)
}

// Lara is a free data retrieval call binding the contract method 0x07f71eb2.
//
// Solidity: function lara() view returns(address)
func (_ApyOracle *ApyOracleCallerSession) Lara() (common.Address, error) {
	return _ApyOracle.Contract.Lara(&_ApyOracle.CallOpts)
}

// MaxValidatorStakeCapacity is a free data retrieval call binding the contract method 0x2a8cf87f.
//
// Solidity: function maxValidatorStakeCapacity() view returns(uint256)
func (_ApyOracle *ApyOracleCaller) MaxValidatorStakeCapacity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ApyOracle.contract.Call(opts, &out, "maxValidatorStakeCapacity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxValidatorStakeCapacity is a free data retrieval call binding the contract method 0x2a8cf87f.
//
// Solidity: function maxValidatorStakeCapacity() view returns(uint256)
func (_ApyOracle *ApyOracleSession) MaxValidatorStakeCapacity() (*big.Int, error) {
	return _ApyOracle.Contract.MaxValidatorStakeCapacity(&_ApyOracle.CallOpts)
}

// MaxValidatorStakeCapacity is a free data retrieval call binding the contract method 0x2a8cf87f.
//
// Solidity: function maxValidatorStakeCapacity() view returns(uint256)
func (_ApyOracle *ApyOracleCallerSession) MaxValidatorStakeCapacity() (*big.Int, error) {
	return _ApyOracle.Contract.MaxValidatorStakeCapacity(&_ApyOracle.CallOpts)
}

// NodeCount is a free data retrieval call binding the contract method 0x6da49b83.
//
// Solidity: function nodeCount() view returns(uint256)
func (_ApyOracle *ApyOracleCaller) NodeCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ApyOracle.contract.Call(opts, &out, "nodeCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NodeCount is a free data retrieval call binding the contract method 0x6da49b83.
//
// Solidity: function nodeCount() view returns(uint256)
func (_ApyOracle *ApyOracleSession) NodeCount() (*big.Int, error) {
	return _ApyOracle.Contract.NodeCount(&_ApyOracle.CallOpts)
}

// NodeCount is a free data retrieval call binding the contract method 0x6da49b83.
//
// Solidity: function nodeCount() view returns(uint256)
func (_ApyOracle *ApyOracleCallerSession) NodeCount() (*big.Int, error) {
	return _ApyOracle.Contract.NodeCount(&_ApyOracle.CallOpts)
}

// Nodes is a free data retrieval call binding the contract method 0x189a5a17.
//
// Solidity: function nodes(address ) view returns(uint256 rating, address account, uint64 fromBlock, uint64 toBlock, uint16 rank, uint16 apy)
func (_ApyOracle *ApyOracleCaller) Nodes(opts *bind.CallOpts, arg0 common.Address) (struct {
	Rating    *big.Int
	Account   common.Address
	FromBlock uint64
	ToBlock   uint64
	Rank      uint16
	Apy       uint16
}, error) {
	var out []interface{}
	err := _ApyOracle.contract.Call(opts, &out, "nodes", arg0)

	outstruct := new(struct {
		Rating    *big.Int
		Account   common.Address
		FromBlock uint64
		ToBlock   uint64
		Rank      uint16
		Apy       uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Rating = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Account = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.FromBlock = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.ToBlock = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.Rank = *abi.ConvertType(out[4], new(uint16)).(*uint16)
	outstruct.Apy = *abi.ConvertType(out[5], new(uint16)).(*uint16)

	return *outstruct, err

}

// Nodes is a free data retrieval call binding the contract method 0x189a5a17.
//
// Solidity: function nodes(address ) view returns(uint256 rating, address account, uint64 fromBlock, uint64 toBlock, uint16 rank, uint16 apy)
func (_ApyOracle *ApyOracleSession) Nodes(arg0 common.Address) (struct {
	Rating    *big.Int
	Account   common.Address
	FromBlock uint64
	ToBlock   uint64
	Rank      uint16
	Apy       uint16
}, error) {
	return _ApyOracle.Contract.Nodes(&_ApyOracle.CallOpts, arg0)
}

// Nodes is a free data retrieval call binding the contract method 0x189a5a17.
//
// Solidity: function nodes(address ) view returns(uint256 rating, address account, uint64 fromBlock, uint64 toBlock, uint16 rank, uint16 apy)
func (_ApyOracle *ApyOracleCallerSession) Nodes(arg0 common.Address) (struct {
	Rating    *big.Int
	Account   common.Address
	FromBlock uint64
	ToBlock   uint64
	Rank      uint16
	Apy       uint16
}, error) {
	return _ApyOracle.Contract.Nodes(&_ApyOracle.CallOpts, arg0)
}

// NodesList is a free data retrieval call binding the contract method 0x76eb6d65.
//
// Solidity: function nodesList(uint256 ) view returns(address)
func (_ApyOracle *ApyOracleCaller) NodesList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ApyOracle.contract.Call(opts, &out, "nodesList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NodesList is a free data retrieval call binding the contract method 0x76eb6d65.
//
// Solidity: function nodesList(uint256 ) view returns(address)
func (_ApyOracle *ApyOracleSession) NodesList(arg0 *big.Int) (common.Address, error) {
	return _ApyOracle.Contract.NodesList(&_ApyOracle.CallOpts, arg0)
}

// NodesList is a free data retrieval call binding the contract method 0x76eb6d65.
//
// Solidity: function nodesList(uint256 ) view returns(address)
func (_ApyOracle *ApyOracleCallerSession) NodesList(arg0 *big.Int) (common.Address, error) {
	return _ApyOracle.Contract.NodesList(&_ApyOracle.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ApyOracle *ApyOracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ApyOracle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ApyOracle *ApyOracleSession) Owner() (common.Address, error) {
	return _ApyOracle.Contract.Owner(&_ApyOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ApyOracle *ApyOracleCallerSession) Owner() (common.Address, error) {
	return _ApyOracle.Contract.Owner(&_ApyOracle.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ApyOracle *ApyOracleCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ApyOracle.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ApyOracle *ApyOracleSession) ProxiableUUID() ([32]byte, error) {
	return _ApyOracle.Contract.ProxiableUUID(&_ApyOracle.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ApyOracle *ApyOracleCallerSession) ProxiableUUID() ([32]byte, error) {
	return _ApyOracle.Contract.ProxiableUUID(&_ApyOracle.CallOpts)
}

// BatchUpdateNodeData is a paid mutator transaction binding the contract method 0x1caa3b52.
//
// Solidity: function batchUpdateNodeData((uint256,address,uint64,uint64,uint16,uint16)[] data) returns()
func (_ApyOracle *ApyOracleTransactor) BatchUpdateNodeData(opts *bind.TransactOpts, data []IApyOracleNodeData) (*types.Transaction, error) {
	return _ApyOracle.contract.Transact(opts, "batchUpdateNodeData", data)
}

// BatchUpdateNodeData is a paid mutator transaction binding the contract method 0x1caa3b52.
//
// Solidity: function batchUpdateNodeData((uint256,address,uint64,uint64,uint16,uint16)[] data) returns()
func (_ApyOracle *ApyOracleSession) BatchUpdateNodeData(data []IApyOracleNodeData) (*types.Transaction, error) {
	return _ApyOracle.Contract.BatchUpdateNodeData(&_ApyOracle.TransactOpts, data)
}

// BatchUpdateNodeData is a paid mutator transaction binding the contract method 0x1caa3b52.
//
// Solidity: function batchUpdateNodeData((uint256,address,uint64,uint64,uint16,uint16)[] data) returns()
func (_ApyOracle *ApyOracleTransactorSession) BatchUpdateNodeData(data []IApyOracleNodeData) (*types.Transaction, error) {
	return _ApyOracle.Contract.BatchUpdateNodeData(&_ApyOracle.TransactOpts, data)
}

// GetNodesForDelegation is a paid mutator transaction binding the contract method 0xe87eafa4.
//
// Solidity: function getNodesForDelegation(uint256 amount) returns((address,uint256,uint256)[])
func (_ApyOracle *ApyOracleTransactor) GetNodesForDelegation(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ApyOracle.contract.Transact(opts, "getNodesForDelegation", amount)
}

// GetNodesForDelegation is a paid mutator transaction binding the contract method 0xe87eafa4.
//
// Solidity: function getNodesForDelegation(uint256 amount) returns((address,uint256,uint256)[])
func (_ApyOracle *ApyOracleSession) GetNodesForDelegation(amount *big.Int) (*types.Transaction, error) {
	return _ApyOracle.Contract.GetNodesForDelegation(&_ApyOracle.TransactOpts, amount)
}

// GetNodesForDelegation is a paid mutator transaction binding the contract method 0xe87eafa4.
//
// Solidity: function getNodesForDelegation(uint256 amount) returns((address,uint256,uint256)[])
func (_ApyOracle *ApyOracleTransactorSession) GetNodesForDelegation(amount *big.Int) (*types.Transaction, error) {
	return _ApyOracle.Contract.GetNodesForDelegation(&_ApyOracle.TransactOpts, amount)
}

// GetRebalanceList is a paid mutator transaction binding the contract method 0x9a0e3695.
//
// Solidity: function getRebalanceList((address,uint256,uint256)[] currentValidators) returns((address,address,uint256,uint256)[])
func (_ApyOracle *ApyOracleTransactor) GetRebalanceList(opts *bind.TransactOpts, currentValidators []IApyOracleTentativeDelegation) (*types.Transaction, error) {
	return _ApyOracle.contract.Transact(opts, "getRebalanceList", currentValidators)
}

// GetRebalanceList is a paid mutator transaction binding the contract method 0x9a0e3695.
//
// Solidity: function getRebalanceList((address,uint256,uint256)[] currentValidators) returns((address,address,uint256,uint256)[])
func (_ApyOracle *ApyOracleSession) GetRebalanceList(currentValidators []IApyOracleTentativeDelegation) (*types.Transaction, error) {
	return _ApyOracle.Contract.GetRebalanceList(&_ApyOracle.TransactOpts, currentValidators)
}

// GetRebalanceList is a paid mutator transaction binding the contract method 0x9a0e3695.
//
// Solidity: function getRebalanceList((address,uint256,uint256)[] currentValidators) returns((address,address,uint256,uint256)[])
func (_ApyOracle *ApyOracleTransactorSession) GetRebalanceList(currentValidators []IApyOracleTentativeDelegation) (*types.Transaction, error) {
	return _ApyOracle.Contract.GetRebalanceList(&_ApyOracle.TransactOpts, currentValidators)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address dataFeed, address dpos) returns()
func (_ApyOracle *ApyOracleTransactor) Initialize(opts *bind.TransactOpts, dataFeed common.Address, dpos common.Address) (*types.Transaction, error) {
	return _ApyOracle.contract.Transact(opts, "initialize", dataFeed, dpos)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address dataFeed, address dpos) returns()
func (_ApyOracle *ApyOracleSession) Initialize(dataFeed common.Address, dpos common.Address) (*types.Transaction, error) {
	return _ApyOracle.Contract.Initialize(&_ApyOracle.TransactOpts, dataFeed, dpos)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address dataFeed, address dpos) returns()
func (_ApyOracle *ApyOracleTransactorSession) Initialize(dataFeed common.Address, dpos common.Address) (*types.Transaction, error) {
	return _ApyOracle.Contract.Initialize(&_ApyOracle.TransactOpts, dataFeed, dpos)
}

// LogRewardDistribution is a paid mutator transaction binding the contract method 0x38c67884.
//
// Solidity: function logRewardDistribution(address staker, uint256 snapshotId, uint256 amount) returns()
func (_ApyOracle *ApyOracleTransactor) LogRewardDistribution(opts *bind.TransactOpts, staker common.Address, snapshotId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _ApyOracle.contract.Transact(opts, "logRewardDistribution", staker, snapshotId, amount)
}

// LogRewardDistribution is a paid mutator transaction binding the contract method 0x38c67884.
//
// Solidity: function logRewardDistribution(address staker, uint256 snapshotId, uint256 amount) returns()
func (_ApyOracle *ApyOracleSession) LogRewardDistribution(staker common.Address, snapshotId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _ApyOracle.Contract.LogRewardDistribution(&_ApyOracle.TransactOpts, staker, snapshotId, amount)
}

// LogRewardDistribution is a paid mutator transaction binding the contract method 0x38c67884.
//
// Solidity: function logRewardDistribution(address staker, uint256 snapshotId, uint256 amount) returns()
func (_ApyOracle *ApyOracleTransactorSession) LogRewardDistribution(staker common.Address, snapshotId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _ApyOracle.Contract.LogRewardDistribution(&_ApyOracle.TransactOpts, staker, snapshotId, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ApyOracle *ApyOracleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ApyOracle.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ApyOracle *ApyOracleSession) RenounceOwnership() (*types.Transaction, error) {
	return _ApyOracle.Contract.RenounceOwnership(&_ApyOracle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ApyOracle *ApyOracleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ApyOracle.Contract.RenounceOwnership(&_ApyOracle.TransactOpts)
}

// SetLara is a paid mutator transaction binding the contract method 0xe2b27fe9.
//
// Solidity: function setLara(address _lara) returns()
func (_ApyOracle *ApyOracleTransactor) SetLara(opts *bind.TransactOpts, _lara common.Address) (*types.Transaction, error) {
	return _ApyOracle.contract.Transact(opts, "setLara", _lara)
}

// SetLara is a paid mutator transaction binding the contract method 0xe2b27fe9.
//
// Solidity: function setLara(address _lara) returns()
func (_ApyOracle *ApyOracleSession) SetLara(_lara common.Address) (*types.Transaction, error) {
	return _ApyOracle.Contract.SetLara(&_ApyOracle.TransactOpts, _lara)
}

// SetLara is a paid mutator transaction binding the contract method 0xe2b27fe9.
//
// Solidity: function setLara(address _lara) returns()
func (_ApyOracle *ApyOracleTransactorSession) SetLara(_lara common.Address) (*types.Transaction, error) {
	return _ApyOracle.Contract.SetLara(&_ApyOracle.TransactOpts, _lara)
}

// SetMaxValidatorStakeCapacity is a paid mutator transaction binding the contract method 0x6d2d8519.
//
// Solidity: function setMaxValidatorStakeCapacity(uint256 capacity) returns()
func (_ApyOracle *ApyOracleTransactor) SetMaxValidatorStakeCapacity(opts *bind.TransactOpts, capacity *big.Int) (*types.Transaction, error) {
	return _ApyOracle.contract.Transact(opts, "setMaxValidatorStakeCapacity", capacity)
}

// SetMaxValidatorStakeCapacity is a paid mutator transaction binding the contract method 0x6d2d8519.
//
// Solidity: function setMaxValidatorStakeCapacity(uint256 capacity) returns()
func (_ApyOracle *ApyOracleSession) SetMaxValidatorStakeCapacity(capacity *big.Int) (*types.Transaction, error) {
	return _ApyOracle.Contract.SetMaxValidatorStakeCapacity(&_ApyOracle.TransactOpts, capacity)
}

// SetMaxValidatorStakeCapacity is a paid mutator transaction binding the contract method 0x6d2d8519.
//
// Solidity: function setMaxValidatorStakeCapacity(uint256 capacity) returns()
func (_ApyOracle *ApyOracleTransactorSession) SetMaxValidatorStakeCapacity(capacity *big.Int) (*types.Transaction, error) {
	return _ApyOracle.Contract.SetMaxValidatorStakeCapacity(&_ApyOracle.TransactOpts, capacity)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ApyOracle *ApyOracleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ApyOracle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ApyOracle *ApyOracleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ApyOracle.Contract.TransferOwnership(&_ApyOracle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ApyOracle *ApyOracleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ApyOracle.Contract.TransferOwnership(&_ApyOracle.TransactOpts, newOwner)
}

// UpdateNodeCount is a paid mutator transaction binding the contract method 0xfe329fd9.
//
// Solidity: function updateNodeCount(uint256 count) returns()
func (_ApyOracle *ApyOracleTransactor) UpdateNodeCount(opts *bind.TransactOpts, count *big.Int) (*types.Transaction, error) {
	return _ApyOracle.contract.Transact(opts, "updateNodeCount", count)
}

// UpdateNodeCount is a paid mutator transaction binding the contract method 0xfe329fd9.
//
// Solidity: function updateNodeCount(uint256 count) returns()
func (_ApyOracle *ApyOracleSession) UpdateNodeCount(count *big.Int) (*types.Transaction, error) {
	return _ApyOracle.Contract.UpdateNodeCount(&_ApyOracle.TransactOpts, count)
}

// UpdateNodeCount is a paid mutator transaction binding the contract method 0xfe329fd9.
//
// Solidity: function updateNodeCount(uint256 count) returns()
func (_ApyOracle *ApyOracleTransactorSession) UpdateNodeCount(count *big.Int) (*types.Transaction, error) {
	return _ApyOracle.Contract.UpdateNodeCount(&_ApyOracle.TransactOpts, count)
}

// UpdateNodeData is a paid mutator transaction binding the contract method 0xa4216ca4.
//
// Solidity: function updateNodeData(address node, (uint256,address,uint64,uint64,uint16,uint16) data) returns()
func (_ApyOracle *ApyOracleTransactor) UpdateNodeData(opts *bind.TransactOpts, node common.Address, data IApyOracleNodeData) (*types.Transaction, error) {
	return _ApyOracle.contract.Transact(opts, "updateNodeData", node, data)
}

// UpdateNodeData is a paid mutator transaction binding the contract method 0xa4216ca4.
//
// Solidity: function updateNodeData(address node, (uint256,address,uint64,uint64,uint16,uint16) data) returns()
func (_ApyOracle *ApyOracleSession) UpdateNodeData(node common.Address, data IApyOracleNodeData) (*types.Transaction, error) {
	return _ApyOracle.Contract.UpdateNodeData(&_ApyOracle.TransactOpts, node, data)
}

// UpdateNodeData is a paid mutator transaction binding the contract method 0xa4216ca4.
//
// Solidity: function updateNodeData(address node, (uint256,address,uint64,uint64,uint16,uint16) data) returns()
func (_ApyOracle *ApyOracleTransactorSession) UpdateNodeData(node common.Address, data IApyOracleNodeData) (*types.Transaction, error) {
	return _ApyOracle.Contract.UpdateNodeData(&_ApyOracle.TransactOpts, node, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ApyOracle *ApyOracleTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ApyOracle.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ApyOracle *ApyOracleSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ApyOracle.Contract.UpgradeToAndCall(&_ApyOracle.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ApyOracle *ApyOracleTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ApyOracle.Contract.UpgradeToAndCall(&_ApyOracle.TransactOpts, newImplementation, data)
}

// ApyOracleInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ApyOracle contract.
type ApyOracleInitializedIterator struct {
	Event *ApyOracleInitialized // Event containing the contract specifics and raw log

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
func (it *ApyOracleInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApyOracleInitialized)
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
		it.Event = new(ApyOracleInitialized)
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
func (it *ApyOracleInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApyOracleInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApyOracleInitialized represents a Initialized event raised by the ApyOracle contract.
type ApyOracleInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ApyOracle *ApyOracleFilterer) FilterInitialized(opts *bind.FilterOpts) (*ApyOracleInitializedIterator, error) {

	logs, sub, err := _ApyOracle.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ApyOracleInitializedIterator{contract: _ApyOracle.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ApyOracle *ApyOracleFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ApyOracleInitialized) (event.Subscription, error) {

	logs, sub, err := _ApyOracle.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApyOracleInitialized)
				if err := _ApyOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ApyOracle *ApyOracleFilterer) ParseInitialized(log types.Log) (*ApyOracleInitialized, error) {
	event := new(ApyOracleInitialized)
	if err := _ApyOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApyOracleMaxValidatorStakeUpdatedIterator is returned from FilterMaxValidatorStakeUpdated and is used to iterate over the raw logs and unpacked data for MaxValidatorStakeUpdated events raised by the ApyOracle contract.
type ApyOracleMaxValidatorStakeUpdatedIterator struct {
	Event *ApyOracleMaxValidatorStakeUpdated // Event containing the contract specifics and raw log

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
func (it *ApyOracleMaxValidatorStakeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApyOracleMaxValidatorStakeUpdated)
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
		it.Event = new(ApyOracleMaxValidatorStakeUpdated)
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
func (it *ApyOracleMaxValidatorStakeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApyOracleMaxValidatorStakeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApyOracleMaxValidatorStakeUpdated represents a MaxValidatorStakeUpdated event raised by the ApyOracle contract.
type ApyOracleMaxValidatorStakeUpdated struct {
	MaxValidatorStake *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterMaxValidatorStakeUpdated is a free log retrieval operation binding the contract event 0xa7810aa6979e862a514c8eb823c7dfecdeae735972b8bef8f76e2ae460f53e15.
//
// Solidity: event MaxValidatorStakeUpdated(uint256 maxValidatorStake)
func (_ApyOracle *ApyOracleFilterer) FilterMaxValidatorStakeUpdated(opts *bind.FilterOpts) (*ApyOracleMaxValidatorStakeUpdatedIterator, error) {

	logs, sub, err := _ApyOracle.contract.FilterLogs(opts, "MaxValidatorStakeUpdated")
	if err != nil {
		return nil, err
	}
	return &ApyOracleMaxValidatorStakeUpdatedIterator{contract: _ApyOracle.contract, event: "MaxValidatorStakeUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxValidatorStakeUpdated is a free log subscription operation binding the contract event 0xa7810aa6979e862a514c8eb823c7dfecdeae735972b8bef8f76e2ae460f53e15.
//
// Solidity: event MaxValidatorStakeUpdated(uint256 maxValidatorStake)
func (_ApyOracle *ApyOracleFilterer) WatchMaxValidatorStakeUpdated(opts *bind.WatchOpts, sink chan<- *ApyOracleMaxValidatorStakeUpdated) (event.Subscription, error) {

	logs, sub, err := _ApyOracle.contract.WatchLogs(opts, "MaxValidatorStakeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApyOracleMaxValidatorStakeUpdated)
				if err := _ApyOracle.contract.UnpackLog(event, "MaxValidatorStakeUpdated", log); err != nil {
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

// ParseMaxValidatorStakeUpdated is a log parse operation binding the contract event 0xa7810aa6979e862a514c8eb823c7dfecdeae735972b8bef8f76e2ae460f53e15.
//
// Solidity: event MaxValidatorStakeUpdated(uint256 maxValidatorStake)
func (_ApyOracle *ApyOracleFilterer) ParseMaxValidatorStakeUpdated(log types.Log) (*ApyOracleMaxValidatorStakeUpdated, error) {
	event := new(ApyOracleMaxValidatorStakeUpdated)
	if err := _ApyOracle.contract.UnpackLog(event, "MaxValidatorStakeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApyOracleNodeDataUpdatedIterator is returned from FilterNodeDataUpdated and is used to iterate over the raw logs and unpacked data for NodeDataUpdated events raised by the ApyOracle contract.
type ApyOracleNodeDataUpdatedIterator struct {
	Event *ApyOracleNodeDataUpdated // Event containing the contract specifics and raw log

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
func (it *ApyOracleNodeDataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApyOracleNodeDataUpdated)
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
		it.Event = new(ApyOracleNodeDataUpdated)
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
func (it *ApyOracleNodeDataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApyOracleNodeDataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApyOracleNodeDataUpdated represents a NodeDataUpdated event raised by the ApyOracle contract.
type ApyOracleNodeDataUpdated struct {
	Node      common.Address
	Apy       uint16
	PbftCount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNodeDataUpdated is a free log retrieval operation binding the contract event 0xf39bff1d19208897d6a4b343f9307b7c7d2df66be3d167cea7aa23e73a1f8896.
//
// Solidity: event NodeDataUpdated(address indexed node, uint16 apy, uint256 pbftCount)
func (_ApyOracle *ApyOracleFilterer) FilterNodeDataUpdated(opts *bind.FilterOpts, node []common.Address) (*ApyOracleNodeDataUpdatedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ApyOracle.contract.FilterLogs(opts, "NodeDataUpdated", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ApyOracleNodeDataUpdatedIterator{contract: _ApyOracle.contract, event: "NodeDataUpdated", logs: logs, sub: sub}, nil
}

// WatchNodeDataUpdated is a free log subscription operation binding the contract event 0xf39bff1d19208897d6a4b343f9307b7c7d2df66be3d167cea7aa23e73a1f8896.
//
// Solidity: event NodeDataUpdated(address indexed node, uint16 apy, uint256 pbftCount)
func (_ApyOracle *ApyOracleFilterer) WatchNodeDataUpdated(opts *bind.WatchOpts, sink chan<- *ApyOracleNodeDataUpdated, node []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ApyOracle.contract.WatchLogs(opts, "NodeDataUpdated", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApyOracleNodeDataUpdated)
				if err := _ApyOracle.contract.UnpackLog(event, "NodeDataUpdated", log); err != nil {
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

// ParseNodeDataUpdated is a log parse operation binding the contract event 0xf39bff1d19208897d6a4b343f9307b7c7d2df66be3d167cea7aa23e73a1f8896.
//
// Solidity: event NodeDataUpdated(address indexed node, uint16 apy, uint256 pbftCount)
func (_ApyOracle *ApyOracleFilterer) ParseNodeDataUpdated(log types.Log) (*ApyOracleNodeDataUpdated, error) {
	event := new(ApyOracleNodeDataUpdated)
	if err := _ApyOracle.contract.UnpackLog(event, "NodeDataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApyOracleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ApyOracle contract.
type ApyOracleOwnershipTransferredIterator struct {
	Event *ApyOracleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ApyOracleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApyOracleOwnershipTransferred)
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
		it.Event = new(ApyOracleOwnershipTransferred)
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
func (it *ApyOracleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApyOracleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApyOracleOwnershipTransferred represents a OwnershipTransferred event raised by the ApyOracle contract.
type ApyOracleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ApyOracle *ApyOracleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ApyOracleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ApyOracle.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ApyOracleOwnershipTransferredIterator{contract: _ApyOracle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ApyOracle *ApyOracleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ApyOracleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ApyOracle.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApyOracleOwnershipTransferred)
				if err := _ApyOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ApyOracle *ApyOracleFilterer) ParseOwnershipTransferred(log types.Log) (*ApyOracleOwnershipTransferred, error) {
	event := new(ApyOracleOwnershipTransferred)
	if err := _ApyOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApyOracleUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ApyOracle contract.
type ApyOracleUpgradedIterator struct {
	Event *ApyOracleUpgraded // Event containing the contract specifics and raw log

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
func (it *ApyOracleUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApyOracleUpgraded)
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
		it.Event = new(ApyOracleUpgraded)
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
func (it *ApyOracleUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApyOracleUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApyOracleUpgraded represents a Upgraded event raised by the ApyOracle contract.
type ApyOracleUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ApyOracle *ApyOracleFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ApyOracleUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ApyOracle.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ApyOracleUpgradedIterator{contract: _ApyOracle.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ApyOracle *ApyOracleFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ApyOracleUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ApyOracle.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApyOracleUpgraded)
				if err := _ApyOracle.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_ApyOracle *ApyOracleFilterer) ParseUpgraded(log types.Log) (*ApyOracleUpgraded, error) {
	event := new(ApyOracleUpgraded)
	if err := _ApyOracle.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
