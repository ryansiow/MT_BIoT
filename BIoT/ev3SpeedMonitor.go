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

// EV3SpeedMonitorMetaData contains all meta data concerning the EV3SpeedMonitor contract.
var EV3SpeedMonitorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_deviceId\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"deviceId\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"running\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"speed\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_deviceId\",\"type\":\"string\"}],\"name\":\"start\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_speed\",\"type\":\"uint32\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"44cf9b31": "deviceId()",
		"d85bd526": "running()",
		"c20906ac": "speed()",
		"258e60b6": "start(string)",
		"07da68f5": "stop()",
		"c580032f": "update(uint32)",
	},
	Bin: "0x60806040526001805464ffffffffff1916905534801561001e57600080fd5b506040516105273803806105278339818101604052602081101561004157600080fd5b810190808051604051939291908464010000000082111561006157600080fd5b90830190602082018581111561007657600080fd5b825164010000000081118282018810171561009057600080fd5b82525081516020918201929091019080838360005b838110156100bd5781810151838201526020016100a5565b50505050905090810190601f1680156100ea5780820380516001836020036101000a031916815260200191505b50604052505081516101049150600090602084019061010b565b505061019e565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061014c57805160ff1916838001178555610179565b82800160010185558215610179579182015b8281111561017957825182559160200191906001019061015e565b50610185929150610189565b5090565b5b80821115610185576000815560010161018a565b61037a806101ad6000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806307da68f514610067578063258e60b61461007157806344cf9b31146100e1578063c20906ac1461015e578063c580032f1461017f578063d85bd526146101a2575b600080fd5b61006f6101be565b005b61006f6004803603602081101561008757600080fd5b8101906020810181356401000000008111156100a257600080fd5b8201836020820111156100b457600080fd5b803590602001918460018302840111640100000000831117156100d657600080fd5b5090925090506101ca565b6100e96101e7565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561012357818101518382015260200161010b565b50505050905090810190601f1680156101505780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610166610275565b6040805163ffffffff9092168252519081900360200190f35b61006f6004803603602081101561019557600080fd5b503563ffffffff16610286565b6101aa6102a8565b604080519115158252519081900360200190f35b6001805460ff19169055565b6101d6600083836102b1565b50506001805460ff19168117905550565b6000805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561026d5780601f106102425761010080835404028352916020019161026d565b820191906000526020600020905b81548152906001019060200180831161025057829003601f168201915b505050505081565b600154610100900463ffffffff1681565b6001805463ffffffff9092166101000264ffffffff0019909216919091179055565b60015460ff1681565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106102f25782800160ff1982351617855561031f565b8280016001018555821561031f579182015b8281111561031f578235825591602001919060010190610304565b5061032b92915061032f565b5090565b5b8082111561032b576000815560010161033056fea26469706673582212207b31099c047debc7b20f00604dbe071c390b5583952c0277d571f20862d10d5964736f6c634300060c0033",
}

// EV3SpeedMonitorABI is the input ABI used to generate the binding from.
// Deprecated: Use EV3SpeedMonitorMetaData.ABI instead.
var EV3SpeedMonitorABI = EV3SpeedMonitorMetaData.ABI

// Deprecated: Use EV3SpeedMonitorMetaData.Sigs instead.
// EV3SpeedMonitorFuncSigs maps the 4-byte function signature to its string representation.
var EV3SpeedMonitorFuncSigs = EV3SpeedMonitorMetaData.Sigs

// EV3SpeedMonitorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EV3SpeedMonitorMetaData.Bin instead.
var EV3SpeedMonitorBin = EV3SpeedMonitorMetaData.Bin

// DeployEV3SpeedMonitor deploys a new Ethereum contract, binding an instance of EV3SpeedMonitor to it.
func DeployEV3SpeedMonitor(auth *bind.TransactOpts, backend bind.ContractBackend, _deviceId string) (common.Address, *types.Transaction, *EV3SpeedMonitor, error) {
	parsed, err := EV3SpeedMonitorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EV3SpeedMonitorBin), backend, _deviceId)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EV3SpeedMonitor{EV3SpeedMonitorCaller: EV3SpeedMonitorCaller{contract: contract}, EV3SpeedMonitorTransactor: EV3SpeedMonitorTransactor{contract: contract}, EV3SpeedMonitorFilterer: EV3SpeedMonitorFilterer{contract: contract}}, nil
}

// EV3SpeedMonitor is an auto generated Go binding around an Ethereum contract.
type EV3SpeedMonitor struct {
	EV3SpeedMonitorCaller     // Read-only binding to the contract
	EV3SpeedMonitorTransactor // Write-only binding to the contract
	EV3SpeedMonitorFilterer   // Log filterer for contract events
}

// EV3SpeedMonitorCaller is an auto generated read-only Go binding around an Ethereum contract.
type EV3SpeedMonitorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EV3SpeedMonitorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EV3SpeedMonitorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EV3SpeedMonitorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EV3SpeedMonitorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EV3SpeedMonitorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EV3SpeedMonitorSession struct {
	Contract     *EV3SpeedMonitor  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EV3SpeedMonitorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EV3SpeedMonitorCallerSession struct {
	Contract *EV3SpeedMonitorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// EV3SpeedMonitorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EV3SpeedMonitorTransactorSession struct {
	Contract     *EV3SpeedMonitorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// EV3SpeedMonitorRaw is an auto generated low-level Go binding around an Ethereum contract.
type EV3SpeedMonitorRaw struct {
	Contract *EV3SpeedMonitor // Generic contract binding to access the raw methods on
}

// EV3SpeedMonitorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EV3SpeedMonitorCallerRaw struct {
	Contract *EV3SpeedMonitorCaller // Generic read-only contract binding to access the raw methods on
}

// EV3SpeedMonitorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EV3SpeedMonitorTransactorRaw struct {
	Contract *EV3SpeedMonitorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEV3SpeedMonitor creates a new instance of EV3SpeedMonitor, bound to a specific deployed contract.
func NewEV3SpeedMonitor(address common.Address, backend bind.ContractBackend) (*EV3SpeedMonitor, error) {
	contract, err := bindEV3SpeedMonitor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EV3SpeedMonitor{EV3SpeedMonitorCaller: EV3SpeedMonitorCaller{contract: contract}, EV3SpeedMonitorTransactor: EV3SpeedMonitorTransactor{contract: contract}, EV3SpeedMonitorFilterer: EV3SpeedMonitorFilterer{contract: contract}}, nil
}

// NewEV3SpeedMonitorCaller creates a new read-only instance of EV3SpeedMonitor, bound to a specific deployed contract.
func NewEV3SpeedMonitorCaller(address common.Address, caller bind.ContractCaller) (*EV3SpeedMonitorCaller, error) {
	contract, err := bindEV3SpeedMonitor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EV3SpeedMonitorCaller{contract: contract}, nil
}

// NewEV3SpeedMonitorTransactor creates a new write-only instance of EV3SpeedMonitor, bound to a specific deployed contract.
func NewEV3SpeedMonitorTransactor(address common.Address, transactor bind.ContractTransactor) (*EV3SpeedMonitorTransactor, error) {
	contract, err := bindEV3SpeedMonitor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EV3SpeedMonitorTransactor{contract: contract}, nil
}

// NewEV3SpeedMonitorFilterer creates a new log filterer instance of EV3SpeedMonitor, bound to a specific deployed contract.
func NewEV3SpeedMonitorFilterer(address common.Address, filterer bind.ContractFilterer) (*EV3SpeedMonitorFilterer, error) {
	contract, err := bindEV3SpeedMonitor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EV3SpeedMonitorFilterer{contract: contract}, nil
}

// bindEV3SpeedMonitor binds a generic wrapper to an already deployed contract.
func bindEV3SpeedMonitor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EV3SpeedMonitorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EV3SpeedMonitor *EV3SpeedMonitorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EV3SpeedMonitor.Contract.EV3SpeedMonitorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EV3SpeedMonitor *EV3SpeedMonitorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EV3SpeedMonitor.Contract.EV3SpeedMonitorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EV3SpeedMonitor *EV3SpeedMonitorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EV3SpeedMonitor.Contract.EV3SpeedMonitorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EV3SpeedMonitor *EV3SpeedMonitorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EV3SpeedMonitor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EV3SpeedMonitor *EV3SpeedMonitorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EV3SpeedMonitor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EV3SpeedMonitor *EV3SpeedMonitorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EV3SpeedMonitor.Contract.contract.Transact(opts, method, params...)
}

// DeviceId is a free data retrieval call binding the contract method 0x44cf9b31.
//
// Solidity: function deviceId() view returns(string)
func (_EV3SpeedMonitor *EV3SpeedMonitorCaller) DeviceId(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EV3SpeedMonitor.contract.Call(opts, &out, "deviceId")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DeviceId is a free data retrieval call binding the contract method 0x44cf9b31.
//
// Solidity: function deviceId() view returns(string)
func (_EV3SpeedMonitor *EV3SpeedMonitorSession) DeviceId() (string, error) {
	return _EV3SpeedMonitor.Contract.DeviceId(&_EV3SpeedMonitor.CallOpts)
}

// DeviceId is a free data retrieval call binding the contract method 0x44cf9b31.
//
// Solidity: function deviceId() view returns(string)
func (_EV3SpeedMonitor *EV3SpeedMonitorCallerSession) DeviceId() (string, error) {
	return _EV3SpeedMonitor.Contract.DeviceId(&_EV3SpeedMonitor.CallOpts)
}

// Running is a free data retrieval call binding the contract method 0xd85bd526.
//
// Solidity: function running() view returns(bool)
func (_EV3SpeedMonitor *EV3SpeedMonitorCaller) Running(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EV3SpeedMonitor.contract.Call(opts, &out, "running")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Running is a free data retrieval call binding the contract method 0xd85bd526.
//
// Solidity: function running() view returns(bool)
func (_EV3SpeedMonitor *EV3SpeedMonitorSession) Running() (bool, error) {
	return _EV3SpeedMonitor.Contract.Running(&_EV3SpeedMonitor.CallOpts)
}

// Running is a free data retrieval call binding the contract method 0xd85bd526.
//
// Solidity: function running() view returns(bool)
func (_EV3SpeedMonitor *EV3SpeedMonitorCallerSession) Running() (bool, error) {
	return _EV3SpeedMonitor.Contract.Running(&_EV3SpeedMonitor.CallOpts)
}

// Speed is a free data retrieval call binding the contract method 0xc20906ac.
//
// Solidity: function speed() view returns(uint32)
func (_EV3SpeedMonitor *EV3SpeedMonitorCaller) Speed(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _EV3SpeedMonitor.contract.Call(opts, &out, "speed")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Speed is a free data retrieval call binding the contract method 0xc20906ac.
//
// Solidity: function speed() view returns(uint32)
func (_EV3SpeedMonitor *EV3SpeedMonitorSession) Speed() (uint32, error) {
	return _EV3SpeedMonitor.Contract.Speed(&_EV3SpeedMonitor.CallOpts)
}

// Speed is a free data retrieval call binding the contract method 0xc20906ac.
//
// Solidity: function speed() view returns(uint32)
func (_EV3SpeedMonitor *EV3SpeedMonitorCallerSession) Speed() (uint32, error) {
	return _EV3SpeedMonitor.Contract.Speed(&_EV3SpeedMonitor.CallOpts)
}

// Start is a paid mutator transaction binding the contract method 0x258e60b6.
//
// Solidity: function start(string _deviceId) returns()
func (_EV3SpeedMonitor *EV3SpeedMonitorTransactor) Start(opts *bind.TransactOpts, _deviceId string) (*types.Transaction, error) {
	return _EV3SpeedMonitor.contract.Transact(opts, "start", _deviceId)
}

// Start is a paid mutator transaction binding the contract method 0x258e60b6.
//
// Solidity: function start(string _deviceId) returns()
func (_EV3SpeedMonitor *EV3SpeedMonitorSession) Start(_deviceId string) (*types.Transaction, error) {
	return _EV3SpeedMonitor.Contract.Start(&_EV3SpeedMonitor.TransactOpts, _deviceId)
}

// Start is a paid mutator transaction binding the contract method 0x258e60b6.
//
// Solidity: function start(string _deviceId) returns()
func (_EV3SpeedMonitor *EV3SpeedMonitorTransactorSession) Start(_deviceId string) (*types.Transaction, error) {
	return _EV3SpeedMonitor.Contract.Start(&_EV3SpeedMonitor.TransactOpts, _deviceId)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_EV3SpeedMonitor *EV3SpeedMonitorTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EV3SpeedMonitor.contract.Transact(opts, "stop")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_EV3SpeedMonitor *EV3SpeedMonitorSession) Stop() (*types.Transaction, error) {
	return _EV3SpeedMonitor.Contract.Stop(&_EV3SpeedMonitor.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_EV3SpeedMonitor *EV3SpeedMonitorTransactorSession) Stop() (*types.Transaction, error) {
	return _EV3SpeedMonitor.Contract.Stop(&_EV3SpeedMonitor.TransactOpts)
}

// Update is a paid mutator transaction binding the contract method 0xc580032f.
//
// Solidity: function update(uint32 _speed) returns()
func (_EV3SpeedMonitor *EV3SpeedMonitorTransactor) Update(opts *bind.TransactOpts, _speed uint32) (*types.Transaction, error) {
	return _EV3SpeedMonitor.contract.Transact(opts, "update", _speed)
}

// Update is a paid mutator transaction binding the contract method 0xc580032f.
//
// Solidity: function update(uint32 _speed) returns()
func (_EV3SpeedMonitor *EV3SpeedMonitorSession) Update(_speed uint32) (*types.Transaction, error) {
	return _EV3SpeedMonitor.Contract.Update(&_EV3SpeedMonitor.TransactOpts, _speed)
}

// Update is a paid mutator transaction binding the contract method 0xc580032f.
//
// Solidity: function update(uint32 _speed) returns()
func (_EV3SpeedMonitor *EV3SpeedMonitorTransactorSession) Update(_speed uint32) (*types.Transaction, error) {
	return _EV3SpeedMonitor.Contract.Update(&_EV3SpeedMonitor.TransactOpts, _speed)
}
