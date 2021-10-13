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

// CoolNumberContractMetaData contains all meta data concerning the CoolNumberContract contract.
var CoolNumberContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"coolNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_coolNumber\",\"type\":\"uint256\"}],\"name\":\"setCoolNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"1b975e43": "coolNumber()",
		"29fa5d3d": "setCoolNumber(uint256)",
	},
	Bin: "0x6080604052600a60005534801561001557600080fd5b5060ac806100246000396000f3fe6080604052348015600f57600080fd5b506004361060325760003560e01c80631b975e4314603757806329fa5d3d14604f575b600080fd5b603d606b565b60408051918252519081900360200190f35b606960048036036020811015606357600080fd5b50356071565b005b60005481565b60005556fea26469706673582212201824ab9e838017d3ef4c5ed38f81b4e95ec2c3630da74743b87c4305c975645964736f6c634300060c0033",
}

// CoolNumberContractABI is the input ABI used to generate the binding from.
// Deprecated: Use CoolNumberContractMetaData.ABI instead.
var CoolNumberContractABI = CoolNumberContractMetaData.ABI

// Deprecated: Use CoolNumberContractMetaData.Sigs instead.
// CoolNumberContractFuncSigs maps the 4-byte function signature to its string representation.
var CoolNumberContractFuncSigs = CoolNumberContractMetaData.Sigs

// CoolNumberContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CoolNumberContractMetaData.Bin instead.
var CoolNumberContractBin = CoolNumberContractMetaData.Bin

// DeployCoolNumberContract deploys a new Ethereum contract, binding an instance of CoolNumberContract to it.
func DeployCoolNumberContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CoolNumberContract, error) {
	parsed, err := CoolNumberContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CoolNumberContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CoolNumberContract{CoolNumberContractCaller: CoolNumberContractCaller{contract: contract}, CoolNumberContractTransactor: CoolNumberContractTransactor{contract: contract}, CoolNumberContractFilterer: CoolNumberContractFilterer{contract: contract}}, nil
}

// CoolNumberContract is an auto generated Go binding around an Ethereum contract.
type CoolNumberContract struct {
	CoolNumberContractCaller     // Read-only binding to the contract
	CoolNumberContractTransactor // Write-only binding to the contract
	CoolNumberContractFilterer   // Log filterer for contract events
}

// CoolNumberContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type CoolNumberContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoolNumberContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CoolNumberContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoolNumberContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CoolNumberContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoolNumberContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CoolNumberContractSession struct {
	Contract     *CoolNumberContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CoolNumberContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CoolNumberContractCallerSession struct {
	Contract *CoolNumberContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// CoolNumberContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CoolNumberContractTransactorSession struct {
	Contract     *CoolNumberContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// CoolNumberContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type CoolNumberContractRaw struct {
	Contract *CoolNumberContract // Generic contract binding to access the raw methods on
}

// CoolNumberContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CoolNumberContractCallerRaw struct {
	Contract *CoolNumberContractCaller // Generic read-only contract binding to access the raw methods on
}

// CoolNumberContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CoolNumberContractTransactorRaw struct {
	Contract *CoolNumberContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCoolNumberContract creates a new instance of CoolNumberContract, bound to a specific deployed contract.
func NewCoolNumberContract(address common.Address, backend bind.ContractBackend) (*CoolNumberContract, error) {
	contract, err := bindCoolNumberContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CoolNumberContract{CoolNumberContractCaller: CoolNumberContractCaller{contract: contract}, CoolNumberContractTransactor: CoolNumberContractTransactor{contract: contract}, CoolNumberContractFilterer: CoolNumberContractFilterer{contract: contract}}, nil
}

// NewCoolNumberContractCaller creates a new read-only instance of CoolNumberContract, bound to a specific deployed contract.
func NewCoolNumberContractCaller(address common.Address, caller bind.ContractCaller) (*CoolNumberContractCaller, error) {
	contract, err := bindCoolNumberContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CoolNumberContractCaller{contract: contract}, nil
}

// NewCoolNumberContractTransactor creates a new write-only instance of CoolNumberContract, bound to a specific deployed contract.
func NewCoolNumberContractTransactor(address common.Address, transactor bind.ContractTransactor) (*CoolNumberContractTransactor, error) {
	contract, err := bindCoolNumberContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CoolNumberContractTransactor{contract: contract}, nil
}

// NewCoolNumberContractFilterer creates a new log filterer instance of CoolNumberContract, bound to a specific deployed contract.
func NewCoolNumberContractFilterer(address common.Address, filterer bind.ContractFilterer) (*CoolNumberContractFilterer, error) {
	contract, err := bindCoolNumberContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CoolNumberContractFilterer{contract: contract}, nil
}

// bindCoolNumberContract binds a generic wrapper to an already deployed contract.
func bindCoolNumberContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CoolNumberContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CoolNumberContract *CoolNumberContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CoolNumberContract.Contract.CoolNumberContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CoolNumberContract *CoolNumberContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoolNumberContract.Contract.CoolNumberContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CoolNumberContract *CoolNumberContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CoolNumberContract.Contract.CoolNumberContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CoolNumberContract *CoolNumberContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CoolNumberContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CoolNumberContract *CoolNumberContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoolNumberContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CoolNumberContract *CoolNumberContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CoolNumberContract.Contract.contract.Transact(opts, method, params...)
}

// CoolNumber is a free data retrieval call binding the contract method 0x1b975e43.
//
// Solidity: function coolNumber() view returns(uint256)
func (_CoolNumberContract *CoolNumberContractCaller) CoolNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CoolNumberContract.contract.Call(opts, &out, "coolNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CoolNumber is a free data retrieval call binding the contract method 0x1b975e43.
//
// Solidity: function coolNumber() view returns(uint256)
func (_CoolNumberContract *CoolNumberContractSession) CoolNumber() (*big.Int, error) {
	return _CoolNumberContract.Contract.CoolNumber(&_CoolNumberContract.CallOpts)
}

// CoolNumber is a free data retrieval call binding the contract method 0x1b975e43.
//
// Solidity: function coolNumber() view returns(uint256)
func (_CoolNumberContract *CoolNumberContractCallerSession) CoolNumber() (*big.Int, error) {
	return _CoolNumberContract.Contract.CoolNumber(&_CoolNumberContract.CallOpts)
}

// SetCoolNumber is a paid mutator transaction binding the contract method 0x29fa5d3d.
//
// Solidity: function setCoolNumber(uint256 _coolNumber) returns()
func (_CoolNumberContract *CoolNumberContractTransactor) SetCoolNumber(opts *bind.TransactOpts, _coolNumber *big.Int) (*types.Transaction, error) {
	return _CoolNumberContract.contract.Transact(opts, "setCoolNumber", _coolNumber)
}

// SetCoolNumber is a paid mutator transaction binding the contract method 0x29fa5d3d.
//
// Solidity: function setCoolNumber(uint256 _coolNumber) returns()
func (_CoolNumberContract *CoolNumberContractSession) SetCoolNumber(_coolNumber *big.Int) (*types.Transaction, error) {
	return _CoolNumberContract.Contract.SetCoolNumber(&_CoolNumberContract.TransactOpts, _coolNumber)
}

// SetCoolNumber is a paid mutator transaction binding the contract method 0x29fa5d3d.
//
// Solidity: function setCoolNumber(uint256 _coolNumber) returns()
func (_CoolNumberContract *CoolNumberContractTransactorSession) SetCoolNumber(_coolNumber *big.Int) (*types.Transaction, error) {
	return _CoolNumberContract.Contract.SetCoolNumber(&_CoolNumberContract.TransactOpts, _coolNumber)
}
