// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// RandomNumberMetaData contains all meta data concerning the RandomNumber contract.
var RandomNumberMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"randomNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_randomNumber\",\"type\":\"uint256\"}],\"name\":\"setRandomNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"ccbac9f5": "randomNumber()",
		"d6bfea28": "setRandomNumber(uint256)",
	},
	Bin: "0x6080604052600a60005534801561001557600080fd5b5060ac806100246000396000f3fe6080604052348015600f57600080fd5b506004361060325760003560e01c8063ccbac9f5146037578063d6bfea2814604f575b600080fd5b603d606b565b60408051918252519081900360200190f35b606960048036036020811015606357600080fd5b50356071565b005b60005481565b60005556fea26469706673582212208b9b9c9f8ab9f349898920af232b57ccc2147fb80da4d5c14c350d3e4b79c8a564736f6c634300060c0033",
}

// RandomNumberABI is the input ABI used to generate the binding from.
// Deprecated: Use RandomNumberMetaData.ABI instead.
var RandomNumberABI = RandomNumberMetaData.ABI

// Deprecated: Use RandomNumberMetaData.Sigs instead.
// RandomNumberFuncSigs maps the 4-byte function signature to its string representation.
var RandomNumberFuncSigs = RandomNumberMetaData.Sigs

// RandomNumberBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RandomNumberMetaData.Bin instead.
var RandomNumberBin = RandomNumberMetaData.Bin

// DeployRandomNumber deploys a new Ethereum contract, binding an instance of RandomNumber to it.
func DeployRandomNumber(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RandomNumber, error) {
	parsed, err := RandomNumberMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RandomNumberBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RandomNumber{RandomNumberCaller: RandomNumberCaller{contract: contract}, RandomNumberTransactor: RandomNumberTransactor{contract: contract}, RandomNumberFilterer: RandomNumberFilterer{contract: contract}}, nil
}

// RandomNumber is an auto generated Go binding around an Ethereum contract.
type RandomNumber struct {
	RandomNumberCaller     // Read-only binding to the contract
	RandomNumberTransactor // Write-only binding to the contract
	RandomNumberFilterer   // Log filterer for contract events
}

// RandomNumberCaller is an auto generated read-only Go binding around an Ethereum contract.
type RandomNumberCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomNumberTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RandomNumberTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomNumberFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RandomNumberFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomNumberSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RandomNumberSession struct {
	Contract     *RandomNumber     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RandomNumberCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RandomNumberCallerSession struct {
	Contract *RandomNumberCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RandomNumberTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RandomNumberTransactorSession struct {
	Contract     *RandomNumberTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RandomNumberRaw is an auto generated low-level Go binding around an Ethereum contract.
type RandomNumberRaw struct {
	Contract *RandomNumber // Generic contract binding to access the raw methods on
}

// RandomNumberCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RandomNumberCallerRaw struct {
	Contract *RandomNumberCaller // Generic read-only contract binding to access the raw methods on
}

// RandomNumberTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RandomNumberTransactorRaw struct {
	Contract *RandomNumberTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRandomNumber creates a new instance of RandomNumber, bound to a specific deployed contract.
func NewRandomNumber(address common.Address, backend bind.ContractBackend) (*RandomNumber, error) {
	contract, err := bindRandomNumber(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RandomNumber{RandomNumberCaller: RandomNumberCaller{contract: contract}, RandomNumberTransactor: RandomNumberTransactor{contract: contract}, RandomNumberFilterer: RandomNumberFilterer{contract: contract}}, nil
}

// NewRandomNumberCaller creates a new read-only instance of RandomNumber, bound to a specific deployed contract.
func NewRandomNumberCaller(address common.Address, caller bind.ContractCaller) (*RandomNumberCaller, error) {
	contract, err := bindRandomNumber(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RandomNumberCaller{contract: contract}, nil
}

// NewRandomNumberTransactor creates a new write-only instance of RandomNumber, bound to a specific deployed contract.
func NewRandomNumberTransactor(address common.Address, transactor bind.ContractTransactor) (*RandomNumberTransactor, error) {
	contract, err := bindRandomNumber(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RandomNumberTransactor{contract: contract}, nil
}

// NewRandomNumberFilterer creates a new log filterer instance of RandomNumber, bound to a specific deployed contract.
func NewRandomNumberFilterer(address common.Address, filterer bind.ContractFilterer) (*RandomNumberFilterer, error) {
	contract, err := bindRandomNumber(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RandomNumberFilterer{contract: contract}, nil
}

// bindRandomNumber binds a generic wrapper to an already deployed contract.
func bindRandomNumber(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RandomNumberABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RandomNumber *RandomNumberRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RandomNumber.Contract.RandomNumberCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RandomNumber *RandomNumberRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RandomNumber.Contract.RandomNumberTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RandomNumber *RandomNumberRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RandomNumber.Contract.RandomNumberTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RandomNumber *RandomNumberCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RandomNumber.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RandomNumber *RandomNumberTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RandomNumber.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RandomNumber *RandomNumberTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RandomNumber.Contract.contract.Transact(opts, method, params...)
}

// RandomNumber is a free data retrieval call binding the contract method 0xccbac9f5.
//
// Solidity: function randomNumber() view returns(uint256)
func (_RandomNumber *RandomNumberCaller) RandomNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RandomNumber.contract.Call(opts, &out, "randomNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RandomNumber is a free data retrieval call binding the contract method 0xccbac9f5.
//
// Solidity: function randomNumber() view returns(uint256)
func (_RandomNumber *RandomNumberSession) RandomNumber() (*big.Int, error) {
	return _RandomNumber.Contract.RandomNumber(&_RandomNumber.CallOpts)
}

// RandomNumber is a free data retrieval call binding the contract method 0xccbac9f5.
//
// Solidity: function randomNumber() view returns(uint256)
func (_RandomNumber *RandomNumberCallerSession) RandomNumber() (*big.Int, error) {
	return _RandomNumber.Contract.RandomNumber(&_RandomNumber.CallOpts)
}

// SetRandomNumber is a paid mutator transaction binding the contract method 0xd6bfea28.
//
// Solidity: function setRandomNumber(uint256 _randomNumber) returns()
func (_RandomNumber *RandomNumberTransactor) SetRandomNumber(opts *bind.TransactOpts, _randomNumber *big.Int) (*types.Transaction, error) {
	return _RandomNumber.contract.Transact(opts, "setRandomNumber", _randomNumber)
}

// SetRandomNumber is a paid mutator transaction binding the contract method 0xd6bfea28.
//
// Solidity: function setRandomNumber(uint256 _randomNumber) returns()
func (_RandomNumber *RandomNumberSession) SetRandomNumber(_randomNumber *big.Int) (*types.Transaction, error) {
	return _RandomNumber.Contract.SetRandomNumber(&_RandomNumber.TransactOpts, _randomNumber)
}

// SetRandomNumber is a paid mutator transaction binding the contract method 0xd6bfea28.
//
// Solidity: function setRandomNumber(uint256 _randomNumber) returns()
func (_RandomNumber *RandomNumberTransactorSession) SetRandomNumber(_randomNumber *big.Int) (*types.Transaction, error) {
	return _RandomNumber.Contract.SetRandomNumber(&_RandomNumber.TransactOpts, _randomNumber)
}
