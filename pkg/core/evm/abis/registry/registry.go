// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package registry

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
)

// AbiMetaData contains all meta data concerning the Abi contract.
var AbiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"channel\",\"type\":\"string\"}],\"name\":\"addToWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"getWhitelistedChannels\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"channels\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"channel\",\"type\":\"string\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"channel\",\"type\":\"string\"}],\"name\":\"removeFromWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userChannels\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AbiABI is the input ABI used to generate the binding from.
// Deprecated: Use AbiMetaData.ABI instead.
var AbiABI = AbiMetaData.ABI

// Abi is an auto generated Go binding around an Ethereum contract.
type Abi struct {
	AbiCaller     // Read-only binding to the contract
	AbiTransactor // Write-only binding to the contract
	AbiFilterer   // Log filterer for contract events
}

// AbiCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbiSession struct {
	Contract     *Abi              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbiCallerSession struct {
	Contract *AbiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AbiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbiTransactorSession struct {
	Contract     *AbiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbiRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbiRaw struct {
	Contract *Abi // Generic contract binding to access the raw methods on
}

// AbiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbiCallerRaw struct {
	Contract *AbiCaller // Generic read-only contract binding to access the raw methods on
}

// AbiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbiTransactorRaw struct {
	Contract *AbiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbi creates a new instance of Abi, bound to a specific deployed contract.
func NewAbi(address common.Address, backend bind.ContractBackend) (*Abi, error) {
	contract, err := bindAbi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Abi{AbiCaller: AbiCaller{contract: contract}, AbiTransactor: AbiTransactor{contract: contract}, AbiFilterer: AbiFilterer{contract: contract}}, nil
}

// NewAbiCaller creates a new read-only instance of Abi, bound to a specific deployed contract.
func NewAbiCaller(address common.Address, caller bind.ContractCaller) (*AbiCaller, error) {
	contract, err := bindAbi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbiCaller{contract: contract}, nil
}

// NewAbiTransactor creates a new write-only instance of Abi, bound to a specific deployed contract.
func NewAbiTransactor(address common.Address, transactor bind.ContractTransactor) (*AbiTransactor, error) {
	contract, err := bindAbi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbiTransactor{contract: contract}, nil
}

// NewAbiFilterer creates a new log filterer instance of Abi, bound to a specific deployed contract.
func NewAbiFilterer(address common.Address, filterer bind.ContractFilterer) (*AbiFilterer, error) {
	contract, err := bindAbi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbiFilterer{contract: contract}, nil
}

// bindAbi binds a generic wrapper to an already deployed contract.
func bindAbi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AbiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abi *AbiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abi.Contract.AbiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abi *AbiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.Contract.AbiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abi *AbiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abi.Contract.AbiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abi *AbiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abi *AbiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abi *AbiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abi.Contract.contract.Transact(opts, method, params...)
}

// GetWhitelistedChannels is a free data retrieval call binding the contract method 0x834266ab.
//
// Solidity: function getWhitelistedChannels(address account, uint256 startIndex, uint256 length) view returns(string[] channels, uint256 total)
func (_Abi *AbiCaller) GetWhitelistedChannels(opts *bind.CallOpts, account common.Address, startIndex *big.Int, length *big.Int) (struct {
	Channels []string
	Total    *big.Int
}, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "getWhitelistedChannels", account, startIndex, length)

	outstruct := new(struct {
		Channels []string
		Total    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Channels = *abi.ConvertType(out[0], new([]string)).(*[]string)
	outstruct.Total = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetWhitelistedChannels is a free data retrieval call binding the contract method 0x834266ab.
//
// Solidity: function getWhitelistedChannels(address account, uint256 startIndex, uint256 length) view returns(string[] channels, uint256 total)
func (_Abi *AbiSession) GetWhitelistedChannels(account common.Address, startIndex *big.Int, length *big.Int) (struct {
	Channels []string
	Total    *big.Int
}, error) {
	return _Abi.Contract.GetWhitelistedChannels(&_Abi.CallOpts, account, startIndex, length)
}

// GetWhitelistedChannels is a free data retrieval call binding the contract method 0x834266ab.
//
// Solidity: function getWhitelistedChannels(address account, uint256 startIndex, uint256 length) view returns(string[] channels, uint256 total)
func (_Abi *AbiCallerSession) GetWhitelistedChannels(account common.Address, startIndex *big.Int, length *big.Int) (struct {
	Channels []string
	Total    *big.Int
}, error) {
	return _Abi.Contract.GetWhitelistedChannels(&_Abi.CallOpts, account, startIndex, length)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0xc1233d61.
//
// Solidity: function isWhitelisted(address user, string channel) view returns(bool)
func (_Abi *AbiCaller) IsWhitelisted(opts *bind.CallOpts, user common.Address, channel string) (bool, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "isWhitelisted", user, channel)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelisted is a free data retrieval call binding the contract method 0xc1233d61.
//
// Solidity: function isWhitelisted(address user, string channel) view returns(bool)
func (_Abi *AbiSession) IsWhitelisted(user common.Address, channel string) (bool, error) {
	return _Abi.Contract.IsWhitelisted(&_Abi.CallOpts, user, channel)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0xc1233d61.
//
// Solidity: function isWhitelisted(address user, string channel) view returns(bool)
func (_Abi *AbiCallerSession) IsWhitelisted(user common.Address, channel string) (bool, error) {
	return _Abi.Contract.IsWhitelisted(&_Abi.CallOpts, user, channel)
}

// UserChannels is a free data retrieval call binding the contract method 0x85e684f3.
//
// Solidity: function userChannels(address , uint256 ) view returns(string)
func (_Abi *AbiCaller) UserChannels(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (string, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "userChannels", arg0, arg1)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UserChannels is a free data retrieval call binding the contract method 0x85e684f3.
//
// Solidity: function userChannels(address , uint256 ) view returns(string)
func (_Abi *AbiSession) UserChannels(arg0 common.Address, arg1 *big.Int) (string, error) {
	return _Abi.Contract.UserChannels(&_Abi.CallOpts, arg0, arg1)
}

// UserChannels is a free data retrieval call binding the contract method 0x85e684f3.
//
// Solidity: function userChannels(address , uint256 ) view returns(string)
func (_Abi *AbiCallerSession) UserChannels(arg0 common.Address, arg1 *big.Int) (string, error) {
	return _Abi.Contract.UserChannels(&_Abi.CallOpts, arg0, arg1)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0x73e08a47.
//
// Solidity: function addToWhitelist(string channel) returns()
func (_Abi *AbiTransactor) AddToWhitelist(opts *bind.TransactOpts, channel string) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "addToWhitelist", channel)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0x73e08a47.
//
// Solidity: function addToWhitelist(string channel) returns()
func (_Abi *AbiSession) AddToWhitelist(channel string) (*types.Transaction, error) {
	return _Abi.Contract.AddToWhitelist(&_Abi.TransactOpts, channel)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0x73e08a47.
//
// Solidity: function addToWhitelist(string channel) returns()
func (_Abi *AbiTransactorSession) AddToWhitelist(channel string) (*types.Transaction, error) {
	return _Abi.Contract.AddToWhitelist(&_Abi.TransactOpts, channel)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8cbab21f.
//
// Solidity: function removeFromWhitelist(string channel) returns()
func (_Abi *AbiTransactor) RemoveFromWhitelist(opts *bind.TransactOpts, channel string) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "removeFromWhitelist", channel)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8cbab21f.
//
// Solidity: function removeFromWhitelist(string channel) returns()
func (_Abi *AbiSession) RemoveFromWhitelist(channel string) (*types.Transaction, error) {
	return _Abi.Contract.RemoveFromWhitelist(&_Abi.TransactOpts, channel)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8cbab21f.
//
// Solidity: function removeFromWhitelist(string channel) returns()
func (_Abi *AbiTransactorSession) RemoveFromWhitelist(channel string) (*types.Transaction, error) {
	return _Abi.Contract.RemoveFromWhitelist(&_Abi.TransactOpts, channel)
}
