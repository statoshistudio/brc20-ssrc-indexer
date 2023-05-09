// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stake

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

// StakeMetaData contains all meta data concerning the Stake contract.
var StakeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"token\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"messages\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"}],\"name\":\"MessagePurchase\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"MessageTokenRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"StakeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"UnStakeEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddresses\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"}],\"name\":\"buyMessages\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"name\":\"enableWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddresses\",\"type\":\"address\"}],\"name\":\"getNodeLevel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"locked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nodeAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nodeTokenPerMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"orders\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"messageAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"registerNodeAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stake\",\"type\":\"uint256\"}],\"name\":\"setMinStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"setTokenPerMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakeAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakeBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userMessages\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StakeABI is the input ABI used to generate the binding from.
// Deprecated: Use StakeMetaData.ABI instead.
var StakeABI = StakeMetaData.ABI

// Stake is an auto generated Go binding around an Ethereum contract.
type Stake struct {
	StakeCaller     // Read-only binding to the contract
	StakeTransactor // Write-only binding to the contract
	StakeFilterer   // Log filterer for contract events
}

// StakeCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakeSession struct {
	Contract     *Stake            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakeCallerSession struct {
	Contract *StakeCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StakeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakeTransactorSession struct {
	Contract     *StakeTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakeRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakeRaw struct {
	Contract *Stake // Generic contract binding to access the raw methods on
}

// StakeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakeCallerRaw struct {
	Contract *StakeCaller // Generic read-only contract binding to access the raw methods on
}

// StakeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakeTransactorRaw struct {
	Contract *StakeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStake creates a new instance of Stake, bound to a specific deployed contract.
func NewStake(address common.Address, backend bind.ContractBackend) (*Stake, error) {
	contract, err := bindStake(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stake{StakeCaller: StakeCaller{contract: contract}, StakeTransactor: StakeTransactor{contract: contract}, StakeFilterer: StakeFilterer{contract: contract}}, nil
}

// NewStakeCaller creates a new read-only instance of Stake, bound to a specific deployed contract.
func NewStakeCaller(address common.Address, caller bind.ContractCaller) (*StakeCaller, error) {
	contract, err := bindStake(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakeCaller{contract: contract}, nil
}

// NewStakeTransactor creates a new write-only instance of Stake, bound to a specific deployed contract.
func NewStakeTransactor(address common.Address, transactor bind.ContractTransactor) (*StakeTransactor, error) {
	contract, err := bindStake(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakeTransactor{contract: contract}, nil
}

// NewStakeFilterer creates a new log filterer instance of Stake, bound to a specific deployed contract.
func NewStakeFilterer(address common.Address, filterer bind.ContractFilterer) (*StakeFilterer, error) {
	contract, err := bindStake(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakeFilterer{contract: contract}, nil
}

// bindStake binds a generic wrapper to an already deployed contract.
func bindStake(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stake *StakeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stake.Contract.StakeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stake *StakeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stake.Contract.StakeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stake *StakeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stake.Contract.StakeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stake *StakeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stake.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stake *StakeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stake.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stake *StakeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stake.Contract.contract.Transact(opts, method, params...)
}

// GetNodeLevel is a free data retrieval call binding the contract method 0x295b87a8.
//
// Solidity: function getNodeLevel(address _nodeAddresses) view returns(uint256)
func (_Stake *StakeCaller) GetNodeLevel(opts *bind.CallOpts, _nodeAddresses common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "getNodeLevel", _nodeAddresses)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeLevel is a free data retrieval call binding the contract method 0x295b87a8.
//
// Solidity: function getNodeLevel(address _nodeAddresses) view returns(uint256)
func (_Stake *StakeSession) GetNodeLevel(_nodeAddresses common.Address) (*big.Int, error) {
	return _Stake.Contract.GetNodeLevel(&_Stake.CallOpts, _nodeAddresses)
}

// GetNodeLevel is a free data retrieval call binding the contract method 0x295b87a8.
//
// Solidity: function getNodeLevel(address _nodeAddresses) view returns(uint256)
func (_Stake *StakeCallerSession) GetNodeLevel(_nodeAddresses common.Address) (*big.Int, error) {
	return _Stake.Contract.GetNodeLevel(&_Stake.CallOpts, _nodeAddresses)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_Stake *StakeCaller) Locked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "locked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_Stake *StakeSession) Locked() (bool, error) {
	return _Stake.Contract.Locked(&_Stake.CallOpts)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_Stake *StakeCallerSession) Locked() (bool, error) {
	return _Stake.Contract.Locked(&_Stake.CallOpts)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() view returns(uint256)
func (_Stake *StakeCaller) MinStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "minStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() view returns(uint256)
func (_Stake *StakeSession) MinStake() (*big.Int, error) {
	return _Stake.Contract.MinStake(&_Stake.CallOpts)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() view returns(uint256)
func (_Stake *StakeCallerSession) MinStake() (*big.Int, error) {
	return _Stake.Contract.MinStake(&_Stake.CallOpts)
}

// NodeAddresses is a free data retrieval call binding the contract method 0xb58a76c0.
//
// Solidity: function nodeAddresses(address ) view returns(address)
func (_Stake *StakeCaller) NodeAddresses(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "nodeAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NodeAddresses is a free data retrieval call binding the contract method 0xb58a76c0.
//
// Solidity: function nodeAddresses(address ) view returns(address)
func (_Stake *StakeSession) NodeAddresses(arg0 common.Address) (common.Address, error) {
	return _Stake.Contract.NodeAddresses(&_Stake.CallOpts, arg0)
}

// NodeAddresses is a free data retrieval call binding the contract method 0xb58a76c0.
//
// Solidity: function nodeAddresses(address ) view returns(address)
func (_Stake *StakeCallerSession) NodeAddresses(arg0 common.Address) (common.Address, error) {
	return _Stake.Contract.NodeAddresses(&_Stake.CallOpts, arg0)
}

// NodeTokenPerMessage is a free data retrieval call binding the contract method 0x76be2d21.
//
// Solidity: function nodeTokenPerMessage(address ) view returns(uint256)
func (_Stake *StakeCaller) NodeTokenPerMessage(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "nodeTokenPerMessage", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NodeTokenPerMessage is a free data retrieval call binding the contract method 0x76be2d21.
//
// Solidity: function nodeTokenPerMessage(address ) view returns(uint256)
func (_Stake *StakeSession) NodeTokenPerMessage(arg0 common.Address) (*big.Int, error) {
	return _Stake.Contract.NodeTokenPerMessage(&_Stake.CallOpts, arg0)
}

// NodeTokenPerMessage is a free data retrieval call binding the contract method 0x76be2d21.
//
// Solidity: function nodeTokenPerMessage(address ) view returns(uint256)
func (_Stake *StakeCallerSession) NodeTokenPerMessage(arg0 common.Address) (*big.Int, error) {
	return _Stake.Contract.NodeTokenPerMessage(&_Stake.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0x9c3f1e90.
//
// Solidity: function orders(bytes32 ) view returns(address nodeAddress, address buyer, uint256 tokenAmount, uint256 messageAmount)
func (_Stake *StakeCaller) Orders(opts *bind.CallOpts, arg0 [32]byte) (struct {
	NodeAddress   common.Address
	Buyer         common.Address
	TokenAmount   *big.Int
	MessageAmount *big.Int
}, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "orders", arg0)

	outstruct := new(struct {
		NodeAddress   common.Address
		Buyer         common.Address
		TokenAmount   *big.Int
		MessageAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NodeAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Buyer = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TokenAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MessageAmount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Orders is a free data retrieval call binding the contract method 0x9c3f1e90.
//
// Solidity: function orders(bytes32 ) view returns(address nodeAddress, address buyer, uint256 tokenAmount, uint256 messageAmount)
func (_Stake *StakeSession) Orders(arg0 [32]byte) (struct {
	NodeAddress   common.Address
	Buyer         common.Address
	TokenAmount   *big.Int
	MessageAmount *big.Int
}, error) {
	return _Stake.Contract.Orders(&_Stake.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0x9c3f1e90.
//
// Solidity: function orders(bytes32 ) view returns(address nodeAddress, address buyer, uint256 tokenAmount, uint256 messageAmount)
func (_Stake *StakeCallerSession) Orders(arg0 [32]byte) (struct {
	NodeAddress   common.Address
	Buyer         common.Address
	TokenAmount   *big.Int
	MessageAmount *big.Int
}, error) {
	return _Stake.Contract.Orders(&_Stake.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stake *StakeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stake *StakeSession) Owner() (common.Address, error) {
	return _Stake.Contract.Owner(&_Stake.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stake *StakeCallerSession) Owner() (common.Address, error) {
	return _Stake.Contract.Owner(&_Stake.CallOpts)
}

// StakeAddresses is a free data retrieval call binding the contract method 0x6a7cce86.
//
// Solidity: function stakeAddresses(address ) view returns(address)
func (_Stake *StakeCaller) StakeAddresses(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "stakeAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeAddresses is a free data retrieval call binding the contract method 0x6a7cce86.
//
// Solidity: function stakeAddresses(address ) view returns(address)
func (_Stake *StakeSession) StakeAddresses(arg0 common.Address) (common.Address, error) {
	return _Stake.Contract.StakeAddresses(&_Stake.CallOpts, arg0)
}

// StakeAddresses is a free data retrieval call binding the contract method 0x6a7cce86.
//
// Solidity: function stakeAddresses(address ) view returns(address)
func (_Stake *StakeCallerSession) StakeAddresses(arg0 common.Address) (common.Address, error) {
	return _Stake.Contract.StakeAddresses(&_Stake.CallOpts, arg0)
}

// StakeBalance is a free data retrieval call binding the contract method 0x4e7c57a6.
//
// Solidity: function stakeBalance(address ) view returns(uint256)
func (_Stake *StakeCaller) StakeBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "stakeBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeBalance is a free data retrieval call binding the contract method 0x4e7c57a6.
//
// Solidity: function stakeBalance(address ) view returns(uint256)
func (_Stake *StakeSession) StakeBalance(arg0 common.Address) (*big.Int, error) {
	return _Stake.Contract.StakeBalance(&_Stake.CallOpts, arg0)
}

// StakeBalance is a free data retrieval call binding the contract method 0x4e7c57a6.
//
// Solidity: function stakeBalance(address ) view returns(uint256)
func (_Stake *StakeCallerSession) StakeBalance(arg0 common.Address) (*big.Int, error) {
	return _Stake.Contract.StakeBalance(&_Stake.CallOpts, arg0)
}

// UserMessages is a free data retrieval call binding the contract method 0xa7ba2df2.
//
// Solidity: function userMessages(address , address ) view returns(uint256)
func (_Stake *StakeCaller) UserMessages(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "userMessages", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserMessages is a free data retrieval call binding the contract method 0xa7ba2df2.
//
// Solidity: function userMessages(address , address ) view returns(uint256)
func (_Stake *StakeSession) UserMessages(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Stake.Contract.UserMessages(&_Stake.CallOpts, arg0, arg1)
}

// UserMessages is a free data retrieval call binding the contract method 0xa7ba2df2.
//
// Solidity: function userMessages(address , address ) view returns(uint256)
func (_Stake *StakeCallerSession) UserMessages(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Stake.Contract.UserMessages(&_Stake.CallOpts, arg0, arg1)
}

// WithdrawalEnabled is a free data retrieval call binding the contract method 0xf8ea5daf.
//
// Solidity: function withdrawalEnabled() view returns(bool)
func (_Stake *StakeCaller) WithdrawalEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "withdrawalEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WithdrawalEnabled is a free data retrieval call binding the contract method 0xf8ea5daf.
//
// Solidity: function withdrawalEnabled() view returns(bool)
func (_Stake *StakeSession) WithdrawalEnabled() (bool, error) {
	return _Stake.Contract.WithdrawalEnabled(&_Stake.CallOpts)
}

// WithdrawalEnabled is a free data retrieval call binding the contract method 0xf8ea5daf.
//
// Solidity: function withdrawalEnabled() view returns(bool)
func (_Stake *StakeCallerSession) WithdrawalEnabled() (bool, error) {
	return _Stake.Contract.WithdrawalEnabled(&_Stake.CallOpts)
}

// BuyMessages is a paid mutator transaction binding the contract method 0xe26074f7.
//
// Solidity: function buyMessages(address _nodeAddresses, uint256 tokens, bytes32 nonce) returns()
func (_Stake *StakeTransactor) BuyMessages(opts *bind.TransactOpts, _nodeAddresses common.Address, tokens *big.Int, nonce [32]byte) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "buyMessages", _nodeAddresses, tokens, nonce)
}

// BuyMessages is a paid mutator transaction binding the contract method 0xe26074f7.
//
// Solidity: function buyMessages(address _nodeAddresses, uint256 tokens, bytes32 nonce) returns()
func (_Stake *StakeSession) BuyMessages(_nodeAddresses common.Address, tokens *big.Int, nonce [32]byte) (*types.Transaction, error) {
	return _Stake.Contract.BuyMessages(&_Stake.TransactOpts, _nodeAddresses, tokens, nonce)
}

// BuyMessages is a paid mutator transaction binding the contract method 0xe26074f7.
//
// Solidity: function buyMessages(address _nodeAddresses, uint256 tokens, bytes32 nonce) returns()
func (_Stake *StakeTransactorSession) BuyMessages(_nodeAddresses common.Address, tokens *big.Int, nonce [32]byte) (*types.Transaction, error) {
	return _Stake.Contract.BuyMessages(&_Stake.TransactOpts, _nodeAddresses, tokens, nonce)
}

// EnableWithdrawal is a paid mutator transaction binding the contract method 0x5636548f.
//
// Solidity: function enableWithdrawal(bool _enabled) returns()
func (_Stake *StakeTransactor) EnableWithdrawal(opts *bind.TransactOpts, _enabled bool) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "enableWithdrawal", _enabled)
}

// EnableWithdrawal is a paid mutator transaction binding the contract method 0x5636548f.
//
// Solidity: function enableWithdrawal(bool _enabled) returns()
func (_Stake *StakeSession) EnableWithdrawal(_enabled bool) (*types.Transaction, error) {
	return _Stake.Contract.EnableWithdrawal(&_Stake.TransactOpts, _enabled)
}

// EnableWithdrawal is a paid mutator transaction binding the contract method 0x5636548f.
//
// Solidity: function enableWithdrawal(bool _enabled) returns()
func (_Stake *StakeTransactorSession) EnableWithdrawal(_enabled bool) (*types.Transaction, error) {
	return _Stake.Contract.EnableWithdrawal(&_Stake.TransactOpts, _enabled)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _address) returns()
func (_Stake *StakeTransactor) Initialize(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "initialize", _address)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _address) returns()
func (_Stake *StakeSession) Initialize(_address common.Address) (*types.Transaction, error) {
	return _Stake.Contract.Initialize(&_Stake.TransactOpts, _address)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _address) returns()
func (_Stake *StakeTransactorSession) Initialize(_address common.Address) (*types.Transaction, error) {
	return _Stake.Contract.Initialize(&_Stake.TransactOpts, _address)
}

// RegisterNodeAccount is a paid mutator transaction binding the contract method 0x433eeb98.
//
// Solidity: function registerNodeAccount(address nodeAddress) returns()
func (_Stake *StakeTransactor) RegisterNodeAccount(opts *bind.TransactOpts, nodeAddress common.Address) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "registerNodeAccount", nodeAddress)
}

// RegisterNodeAccount is a paid mutator transaction binding the contract method 0x433eeb98.
//
// Solidity: function registerNodeAccount(address nodeAddress) returns()
func (_Stake *StakeSession) RegisterNodeAccount(nodeAddress common.Address) (*types.Transaction, error) {
	return _Stake.Contract.RegisterNodeAccount(&_Stake.TransactOpts, nodeAddress)
}

// RegisterNodeAccount is a paid mutator transaction binding the contract method 0x433eeb98.
//
// Solidity: function registerNodeAccount(address nodeAddress) returns()
func (_Stake *StakeTransactorSession) RegisterNodeAccount(nodeAddress common.Address) (*types.Transaction, error) {
	return _Stake.Contract.RegisterNodeAccount(&_Stake.TransactOpts, nodeAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Stake *StakeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Stake *StakeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Stake.Contract.RenounceOwnership(&_Stake.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Stake *StakeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Stake.Contract.RenounceOwnership(&_Stake.TransactOpts)
}

// SetMinStake is a paid mutator transaction binding the contract method 0x8c80fd90.
//
// Solidity: function setMinStake(uint256 _stake) returns()
func (_Stake *StakeTransactor) SetMinStake(opts *bind.TransactOpts, _stake *big.Int) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "setMinStake", _stake)
}

// SetMinStake is a paid mutator transaction binding the contract method 0x8c80fd90.
//
// Solidity: function setMinStake(uint256 _stake) returns()
func (_Stake *StakeSession) SetMinStake(_stake *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.SetMinStake(&_Stake.TransactOpts, _stake)
}

// SetMinStake is a paid mutator transaction binding the contract method 0x8c80fd90.
//
// Solidity: function setMinStake(uint256 _stake) returns()
func (_Stake *StakeTransactorSession) SetMinStake(_stake *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.SetMinStake(&_Stake.TransactOpts, _stake)
}

// SetTokenPerMessage is a paid mutator transaction binding the contract method 0xaf14be57.
//
// Solidity: function setTokenPerMessage(uint256 rate) returns()
func (_Stake *StakeTransactor) SetTokenPerMessage(opts *bind.TransactOpts, rate *big.Int) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "setTokenPerMessage", rate)
}

// SetTokenPerMessage is a paid mutator transaction binding the contract method 0xaf14be57.
//
// Solidity: function setTokenPerMessage(uint256 rate) returns()
func (_Stake *StakeSession) SetTokenPerMessage(rate *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.SetTokenPerMessage(&_Stake.TransactOpts, rate)
}

// SetTokenPerMessage is a paid mutator transaction binding the contract method 0xaf14be57.
//
// Solidity: function setTokenPerMessage(uint256 rate) returns()
func (_Stake *StakeTransactorSession) SetTokenPerMessage(rate *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.SetTokenPerMessage(&_Stake.TransactOpts, rate)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Stake *StakeTransactor) Stake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "stake", amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Stake *StakeSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.Stake(&_Stake.TransactOpts, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Stake *StakeTransactorSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.Stake(&_Stake.TransactOpts, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Stake *StakeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Stake *StakeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Stake.Contract.TransferOwnership(&_Stake.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Stake *StakeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Stake.Contract.TransferOwnership(&_Stake.TransactOpts, newOwner)
}

// UnStake is a paid mutator transaction binding the contract method 0x73cf575a.
//
// Solidity: function unStake() returns()
func (_Stake *StakeTransactor) UnStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "unStake")
}

// UnStake is a paid mutator transaction binding the contract method 0x73cf575a.
//
// Solidity: function unStake() returns()
func (_Stake *StakeSession) UnStake() (*types.Transaction, error) {
	return _Stake.Contract.UnStake(&_Stake.TransactOpts)
}

// UnStake is a paid mutator transaction binding the contract method 0x73cf575a.
//
// Solidity: function unStake() returns()
func (_Stake *StakeTransactorSession) UnStake() (*types.Transaction, error) {
	return _Stake.Contract.UnStake(&_Stake.TransactOpts)
}

// StakeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Stake contract.
type StakeInitializedIterator struct {
	Event *StakeInitialized // Event containing the contract specifics and raw log

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
func (it *StakeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeInitialized)
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
		it.Event = new(StakeInitialized)
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
func (it *StakeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeInitialized represents a Initialized event raised by the Stake contract.
type StakeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Stake *StakeFilterer) FilterInitialized(opts *bind.FilterOpts) (*StakeInitializedIterator, error) {

	logs, sub, err := _Stake.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StakeInitializedIterator{contract: _Stake.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Stake *StakeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StakeInitialized) (event.Subscription, error) {

	logs, sub, err := _Stake.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeInitialized)
				if err := _Stake.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Stake *StakeFilterer) ParseInitialized(log types.Log) (*StakeInitialized, error) {
	event := new(StakeInitialized)
	if err := _Stake.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeMessagePurchaseIterator is returned from FilterMessagePurchase and is used to iterate over the raw logs and unpacked data for MessagePurchase events raised by the Stake contract.
type StakeMessagePurchaseIterator struct {
	Event *StakeMessagePurchase // Event containing the contract specifics and raw log

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
func (it *StakeMessagePurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeMessagePurchase)
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
		it.Event = new(StakeMessagePurchase)
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
func (it *StakeMessagePurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeMessagePurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeMessagePurchase represents a MessagePurchase event raised by the Stake contract.
type StakeMessagePurchase struct {
	Node     common.Address
	Buyer    common.Address
	Token    *big.Int
	Messages *big.Int
	Nonce    [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMessagePurchase is a free log retrieval operation binding the contract event 0x134f5d697fdb606afe1a229994a1339c4d3d0984e77e6c6ef90a4dd61f3bede1.
//
// Solidity: event MessagePurchase(address indexed node, address indexed buyer, uint256 token, uint256 messages, bytes32 indexed nonce)
func (_Stake *StakeFilterer) FilterMessagePurchase(opts *bind.FilterOpts, node []common.Address, buyer []common.Address, nonce [][32]byte) (*StakeMessagePurchaseIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _Stake.contract.FilterLogs(opts, "MessagePurchase", nodeRule, buyerRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &StakeMessagePurchaseIterator{contract: _Stake.contract, event: "MessagePurchase", logs: logs, sub: sub}, nil
}

// WatchMessagePurchase is a free log subscription operation binding the contract event 0x134f5d697fdb606afe1a229994a1339c4d3d0984e77e6c6ef90a4dd61f3bede1.
//
// Solidity: event MessagePurchase(address indexed node, address indexed buyer, uint256 token, uint256 messages, bytes32 indexed nonce)
func (_Stake *StakeFilterer) WatchMessagePurchase(opts *bind.WatchOpts, sink chan<- *StakeMessagePurchase, node []common.Address, buyer []common.Address, nonce [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _Stake.contract.WatchLogs(opts, "MessagePurchase", nodeRule, buyerRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeMessagePurchase)
				if err := _Stake.contract.UnpackLog(event, "MessagePurchase", log); err != nil {
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

// ParseMessagePurchase is a log parse operation binding the contract event 0x134f5d697fdb606afe1a229994a1339c4d3d0984e77e6c6ef90a4dd61f3bede1.
//
// Solidity: event MessagePurchase(address indexed node, address indexed buyer, uint256 token, uint256 messages, bytes32 indexed nonce)
func (_Stake *StakeFilterer) ParseMessagePurchase(log types.Log) (*StakeMessagePurchase, error) {
	event := new(StakeMessagePurchase)
	if err := _Stake.contract.UnpackLog(event, "MessagePurchase", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeMessageTokenRateIterator is returned from FilterMessageTokenRate and is used to iterate over the raw logs and unpacked data for MessageTokenRate events raised by the Stake contract.
type StakeMessageTokenRateIterator struct {
	Event *StakeMessageTokenRate // Event containing the contract specifics and raw log

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
func (it *StakeMessageTokenRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeMessageTokenRate)
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
		it.Event = new(StakeMessageTokenRate)
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
func (it *StakeMessageTokenRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeMessageTokenRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeMessageTokenRate represents a MessageTokenRate event raised by the Stake contract.
type StakeMessageTokenRate struct {
	Node common.Address
	Rate *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMessageTokenRate is a free log retrieval operation binding the contract event 0x4112b5f837ca529fae87c04e8f274182d1397401721765ea2e68e0dc804e2e47.
//
// Solidity: event MessageTokenRate(address indexed node, uint256 rate)
func (_Stake *StakeFilterer) FilterMessageTokenRate(opts *bind.FilterOpts, node []common.Address) (*StakeMessageTokenRateIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Stake.contract.FilterLogs(opts, "MessageTokenRate", nodeRule)
	if err != nil {
		return nil, err
	}
	return &StakeMessageTokenRateIterator{contract: _Stake.contract, event: "MessageTokenRate", logs: logs, sub: sub}, nil
}

// WatchMessageTokenRate is a free log subscription operation binding the contract event 0x4112b5f837ca529fae87c04e8f274182d1397401721765ea2e68e0dc804e2e47.
//
// Solidity: event MessageTokenRate(address indexed node, uint256 rate)
func (_Stake *StakeFilterer) WatchMessageTokenRate(opts *bind.WatchOpts, sink chan<- *StakeMessageTokenRate, node []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Stake.contract.WatchLogs(opts, "MessageTokenRate", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeMessageTokenRate)
				if err := _Stake.contract.UnpackLog(event, "MessageTokenRate", log); err != nil {
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

// ParseMessageTokenRate is a log parse operation binding the contract event 0x4112b5f837ca529fae87c04e8f274182d1397401721765ea2e68e0dc804e2e47.
//
// Solidity: event MessageTokenRate(address indexed node, uint256 rate)
func (_Stake *StakeFilterer) ParseMessageTokenRate(log types.Log) (*StakeMessageTokenRate, error) {
	event := new(StakeMessageTokenRate)
	if err := _Stake.contract.UnpackLog(event, "MessageTokenRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Stake contract.
type StakeOwnershipTransferredIterator struct {
	Event *StakeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeOwnershipTransferred)
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
		it.Event = new(StakeOwnershipTransferred)
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
func (it *StakeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeOwnershipTransferred represents a OwnershipTransferred event raised by the Stake contract.
type StakeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Stake *StakeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Stake.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakeOwnershipTransferredIterator{contract: _Stake.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Stake *StakeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Stake.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeOwnershipTransferred)
				if err := _Stake.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Stake *StakeFilterer) ParseOwnershipTransferred(log types.Log) (*StakeOwnershipTransferred, error) {
	event := new(StakeOwnershipTransferred)
	if err := _Stake.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeStakeEventIterator is returned from FilterStakeEvent and is used to iterate over the raw logs and unpacked data for StakeEvent events raised by the Stake contract.
type StakeStakeEventIterator struct {
	Event *StakeStakeEvent // Event containing the contract specifics and raw log

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
func (it *StakeStakeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeStakeEvent)
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
		it.Event = new(StakeStakeEvent)
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
func (it *StakeStakeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeStakeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeStakeEvent represents a StakeEvent event raised by the Stake contract.
type StakeStakeEvent struct {
	Account   common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStakeEvent is a free log retrieval operation binding the contract event 0x9dbaf9c586508abc91d6ee4e67d3c7a82ccb09bca5d9fe2c3b690f27b7e0a256.
//
// Solidity: event StakeEvent(address indexed account, uint256 amount, uint256 timestamp)
func (_Stake *StakeFilterer) FilterStakeEvent(opts *bind.FilterOpts, account []common.Address) (*StakeStakeEventIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Stake.contract.FilterLogs(opts, "StakeEvent", accountRule)
	if err != nil {
		return nil, err
	}
	return &StakeStakeEventIterator{contract: _Stake.contract, event: "StakeEvent", logs: logs, sub: sub}, nil
}

// WatchStakeEvent is a free log subscription operation binding the contract event 0x9dbaf9c586508abc91d6ee4e67d3c7a82ccb09bca5d9fe2c3b690f27b7e0a256.
//
// Solidity: event StakeEvent(address indexed account, uint256 amount, uint256 timestamp)
func (_Stake *StakeFilterer) WatchStakeEvent(opts *bind.WatchOpts, sink chan<- *StakeStakeEvent, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Stake.contract.WatchLogs(opts, "StakeEvent", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeStakeEvent)
				if err := _Stake.contract.UnpackLog(event, "StakeEvent", log); err != nil {
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

// ParseStakeEvent is a log parse operation binding the contract event 0x9dbaf9c586508abc91d6ee4e67d3c7a82ccb09bca5d9fe2c3b690f27b7e0a256.
//
// Solidity: event StakeEvent(address indexed account, uint256 amount, uint256 timestamp)
func (_Stake *StakeFilterer) ParseStakeEvent(log types.Log) (*StakeStakeEvent, error) {
	event := new(StakeStakeEvent)
	if err := _Stake.contract.UnpackLog(event, "StakeEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeUnStakeEventIterator is returned from FilterUnStakeEvent and is used to iterate over the raw logs and unpacked data for UnStakeEvent events raised by the Stake contract.
type StakeUnStakeEventIterator struct {
	Event *StakeUnStakeEvent // Event containing the contract specifics and raw log

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
func (it *StakeUnStakeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeUnStakeEvent)
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
		it.Event = new(StakeUnStakeEvent)
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
func (it *StakeUnStakeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeUnStakeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeUnStakeEvent represents a UnStakeEvent event raised by the Stake contract.
type StakeUnStakeEvent struct {
	Account   common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnStakeEvent is a free log retrieval operation binding the contract event 0xa96db5cb7927744735ed3891b328387fe2917486b42976ded70e42a827396ad7.
//
// Solidity: event UnStakeEvent(address indexed account, uint256 amount, uint256 timestamp)
func (_Stake *StakeFilterer) FilterUnStakeEvent(opts *bind.FilterOpts, account []common.Address) (*StakeUnStakeEventIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Stake.contract.FilterLogs(opts, "UnStakeEvent", accountRule)
	if err != nil {
		return nil, err
	}
	return &StakeUnStakeEventIterator{contract: _Stake.contract, event: "UnStakeEvent", logs: logs, sub: sub}, nil
}

// WatchUnStakeEvent is a free log subscription operation binding the contract event 0xa96db5cb7927744735ed3891b328387fe2917486b42976ded70e42a827396ad7.
//
// Solidity: event UnStakeEvent(address indexed account, uint256 amount, uint256 timestamp)
func (_Stake *StakeFilterer) WatchUnStakeEvent(opts *bind.WatchOpts, sink chan<- *StakeUnStakeEvent, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Stake.contract.WatchLogs(opts, "UnStakeEvent", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeUnStakeEvent)
				if err := _Stake.contract.UnpackLog(event, "UnStakeEvent", log); err != nil {
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

// ParseUnStakeEvent is a log parse operation binding the contract event 0xa96db5cb7927744735ed3891b328387fe2917486b42976ded70e42a827396ad7.
//
// Solidity: event UnStakeEvent(address indexed account, uint256 amount, uint256 timestamp)
func (_Stake *StakeFilterer) ParseUnStakeEvent(log types.Log) (*StakeUnStakeEvent, error) {
	event := new(StakeUnStakeEvent)
	if err := _Stake.contract.UnpackLog(event, "UnStakeEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
