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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dataFeed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dpos\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"apy\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pbftCount\",\"type\":\"uint256\"}],\"name\":\"NodeDataUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_dataFeed\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_dpos\",\"outputs\":[{\"internalType\":\"contractDposInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rating\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"fromBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"toBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"rank\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"apy\",\"type\":\"uint16\"}],\"internalType\":\"structIApyOracle.NodeData[]\",\"name\":\"data\",\"type\":\"tuple[]\"}],\"name\":\"batchUpdateNodeData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDataFeedAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"getNodeData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rating\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"fromBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"toBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"rank\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"apy\",\"type\":\"uint16\"}],\"internalType\":\"structIApyOracle.NodeData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getNodesForDelegation\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rating\",\"type\":\"uint256\"}],\"internalType\":\"structIApyOracle.TentativeDelegation[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rating\",\"type\":\"uint256\"}],\"internalType\":\"structIApyOracle.TentativeDelegation[]\",\"name\":\"currentValidators\",\"type\":\"tuple[]\"}],\"name\":\"getRebalanceList\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toRating\",\"type\":\"uint256\"}],\"internalType\":\"structIApyOracle.TentativeReDelegation[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lara\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxValidatorStakeCapacity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nodes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rating\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"fromBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"toBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"rank\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"apy\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nodesList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lara\",\"type\":\"address\"}],\"name\":\"setLara\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"name\":\"setMaxValidatorStakeCapacity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"updateNodeCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rating\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"fromBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"toBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"rank\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"apy\",\"type\":\"uint16\"}],\"internalType\":\"structIApyOracle.NodeData\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"updateNodeData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// DataFeed is a free data retrieval call binding the contract method 0x486146d4.
//
// Solidity: function _dataFeed() view returns(address)
func (_ApyOracle *ApyOracleCaller) DataFeed(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ApyOracle.contract.Call(opts, &out, "_dataFeed")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DataFeed is a free data retrieval call binding the contract method 0x486146d4.
//
// Solidity: function _dataFeed() view returns(address)
func (_ApyOracle *ApyOracleSession) DataFeed() (common.Address, error) {
	return _ApyOracle.Contract.DataFeed(&_ApyOracle.CallOpts)
}

// DataFeed is a free data retrieval call binding the contract method 0x486146d4.
//
// Solidity: function _dataFeed() view returns(address)
func (_ApyOracle *ApyOracleCallerSession) DataFeed() (common.Address, error) {
	return _ApyOracle.Contract.DataFeed(&_ApyOracle.CallOpts)
}

// Dpos is a free data retrieval call binding the contract method 0x4cb1a454.
//
// Solidity: function _dpos() view returns(address)
func (_ApyOracle *ApyOracleCaller) Dpos(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ApyOracle.contract.Call(opts, &out, "_dpos")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dpos is a free data retrieval call binding the contract method 0x4cb1a454.
//
// Solidity: function _dpos() view returns(address)
func (_ApyOracle *ApyOracleSession) Dpos() (common.Address, error) {
	return _ApyOracle.Contract.Dpos(&_ApyOracle.CallOpts)
}

// Dpos is a free data retrieval call binding the contract method 0x4cb1a454.
//
// Solidity: function _dpos() view returns(address)
func (_ApyOracle *ApyOracleCallerSession) Dpos() (common.Address, error) {
	return _ApyOracle.Contract.Dpos(&_ApyOracle.CallOpts)
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
