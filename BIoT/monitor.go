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

// MonitorMetaData contains all meta data concerning the Monitor contract.
var MonitorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_deviceId\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"deviceId\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rotations\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"running\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_deviceId\",\"type\":\"string\"}],\"name\":\"start\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_delta\",\"type\":\"uint32\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"44cf9b31": "deviceId()",
		"430425d8": "rotations()",
		"d85bd526": "running()",
		"258e60b6": "start(string)",
		"07da68f5": "stop()",
		"c580032f": "update(uint32)",
	},
	Bin: "0x60806040526001805464ffffffffff1916905534801561001e57600080fd5b5060405161052f38038061052f8339818101604052602081101561004157600080fd5b810190808051604051939291908464010000000082111561006157600080fd5b90830190602082018581111561007657600080fd5b825164010000000081118282018810171561009057600080fd5b82525081516020918201929091019080838360005b838110156100bd5781810151838201526020016100a5565b50505050905090810190601f1680156100ea5780820380516001836020036101000a031916815260200191505b50604052505081516101049150600090602084019061010b565b505061019e565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061014c57805160ff1916838001178555610179565b82800160010185558215610179579182015b8281111561017957825182559160200191906001019061015e565b50610185929150610189565b5090565b5b80821115610185576000815560010161018a565b610382806101ad6000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806307da68f514610067578063258e60b614610071578063430425d8146100e157806344cf9b3114610102578063c580032f1461017f578063d85bd526146101a2575b600080fd5b61006f6101be565b005b61006f6004803603602081101561008757600080fd5b8101906020810181356401000000008111156100a257600080fd5b8201836020820111156100b457600080fd5b803590602001918460018302840111640100000000831117156100d657600080fd5b5090925090506101ca565b6100e96101e7565b6040805163ffffffff9092168252519081900360200190f35b61010a6101f8565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561014457818101518382015260200161012c565b50505050905090810190601f1680156101715780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61006f6004803603602081101561019557600080fd5b503563ffffffff16610286565b6101aa6102b0565b604080519115158252519081900360200190f35b6001805460ff19169055565b6101d6600083836102b9565b50506001805460ff19168117905550565b600154610100900463ffffffff1681565b6000805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561027e5780601f106102535761010080835404028352916020019161027e565b820191906000526020600020905b81548152906001019060200180831161026157829003601f168201915b505050505081565b6001805463ffffffff61010080830482169094011690920264ffffffff0019909216919091179055565b60015460ff1681565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106102fa5782800160ff19823516178555610327565b82800160010185558215610327579182015b8281111561032757823582559160200191906001019061030c565b50610333929150610337565b5090565b5b80821115610333576000815560010161033856fea264697066735822122056d4f4f766e842ecfefa097bbf5f899880bd9ab99caa7f2320a16fb0db4d592764736f6c634300060c0033",
}

// MonitorABI is the input ABI used to generate the binding from.
// Deprecated: Use MonitorMetaData.ABI instead.
var MonitorABI = MonitorMetaData.ABI

// Deprecated: Use MonitorMetaData.Sigs instead.
// MonitorFuncSigs maps the 4-byte function signature to its string representation.
var MonitorFuncSigs = MonitorMetaData.Sigs

// MonitorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MonitorMetaData.Bin instead.
var MonitorBin = MonitorMetaData.Bin

// DeployMonitor deploys a new Ethereum contract, binding an instance of Monitor to it.
func DeployMonitor(auth *bind.TransactOpts, backend bind.ContractBackend, _deviceId string) (common.Address, *types.Transaction, *Monitor, error) {
	parsed, err := MonitorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MonitorBin), backend, _deviceId)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Monitor{MonitorCaller: MonitorCaller{contract: contract}, MonitorTransactor: MonitorTransactor{contract: contract}, MonitorFilterer: MonitorFilterer{contract: contract}}, nil
}

// Monitor is an auto generated Go binding around an Ethereum contract.
type Monitor struct {
	MonitorCaller     // Read-only binding to the contract
	MonitorTransactor // Write-only binding to the contract
	MonitorFilterer   // Log filterer for contract events
}

// MonitorCaller is an auto generated read-only Go binding around an Ethereum contract.
type MonitorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MonitorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MonitorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MonitorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MonitorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MonitorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MonitorSession struct {
	Contract     *Monitor          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MonitorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MonitorCallerSession struct {
	Contract *MonitorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MonitorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MonitorTransactorSession struct {
	Contract     *MonitorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MonitorRaw is an auto generated low-level Go binding around an Ethereum contract.
type MonitorRaw struct {
	Contract *Monitor // Generic contract binding to access the raw methods on
}

// MonitorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MonitorCallerRaw struct {
	Contract *MonitorCaller // Generic read-only contract binding to access the raw methods on
}

// MonitorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MonitorTransactorRaw struct {
	Contract *MonitorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMonitor creates a new instance of Monitor, bound to a specific deployed contract.
func NewMonitor(address common.Address, backend bind.ContractBackend) (*Monitor, error) {
	contract, err := bindMonitor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Monitor{MonitorCaller: MonitorCaller{contract: contract}, MonitorTransactor: MonitorTransactor{contract: contract}, MonitorFilterer: MonitorFilterer{contract: contract}}, nil
}

// NewMonitorCaller creates a new read-only instance of Monitor, bound to a specific deployed contract.
func NewMonitorCaller(address common.Address, caller bind.ContractCaller) (*MonitorCaller, error) {
	contract, err := bindMonitor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MonitorCaller{contract: contract}, nil
}

// NewMonitorTransactor creates a new write-only instance of Monitor, bound to a specific deployed contract.
func NewMonitorTransactor(address common.Address, transactor bind.ContractTransactor) (*MonitorTransactor, error) {
	contract, err := bindMonitor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MonitorTransactor{contract: contract}, nil
}

// NewMonitorFilterer creates a new log filterer instance of Monitor, bound to a specific deployed contract.
func NewMonitorFilterer(address common.Address, filterer bind.ContractFilterer) (*MonitorFilterer, error) {
	contract, err := bindMonitor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MonitorFilterer{contract: contract}, nil
}

// bindMonitor binds a generic wrapper to an already deployed contract.
func bindMonitor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MonitorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Monitor *MonitorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Monitor.Contract.MonitorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Monitor *MonitorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Monitor.Contract.MonitorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Monitor *MonitorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Monitor.Contract.MonitorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Monitor *MonitorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Monitor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Monitor *MonitorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Monitor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Monitor *MonitorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Monitor.Contract.contract.Transact(opts, method, params...)
}

// DeviceId is a free data retrieval call binding the contract method 0x44cf9b31.
//
// Solidity: function deviceId() view returns(string)
func (_Monitor *MonitorCaller) DeviceId(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Monitor.contract.Call(opts, &out, "deviceId")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DeviceId is a free data retrieval call binding the contract method 0x44cf9b31.
//
// Solidity: function deviceId() view returns(string)
func (_Monitor *MonitorSession) DeviceId() (string, error) {
	return _Monitor.Contract.DeviceId(&_Monitor.CallOpts)
}

// DeviceId is a free data retrieval call binding the contract method 0x44cf9b31.
//
// Solidity: function deviceId() view returns(string)
func (_Monitor *MonitorCallerSession) DeviceId() (string, error) {
	return _Monitor.Contract.DeviceId(&_Monitor.CallOpts)
}

// Rotations is a free data retrieval call binding the contract method 0x430425d8.
//
// Solidity: function rotations() view returns(uint32)
func (_Monitor *MonitorCaller) Rotations(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Monitor.contract.Call(opts, &out, "rotations")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Rotations is a free data retrieval call binding the contract method 0x430425d8.
//
// Solidity: function rotations() view returns(uint32)
func (_Monitor *MonitorSession) Rotations() (uint32, error) {
	return _Monitor.Contract.Rotations(&_Monitor.CallOpts)
}

// Rotations is a free data retrieval call binding the contract method 0x430425d8.
//
// Solidity: function rotations() view returns(uint32)
func (_Monitor *MonitorCallerSession) Rotations() (uint32, error) {
	return _Monitor.Contract.Rotations(&_Monitor.CallOpts)
}

// Running is a free data retrieval call binding the contract method 0xd85bd526.
//
// Solidity: function running() view returns(bool)
func (_Monitor *MonitorCaller) Running(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Monitor.contract.Call(opts, &out, "running")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Running is a free data retrieval call binding the contract method 0xd85bd526.
//
// Solidity: function running() view returns(bool)
func (_Monitor *MonitorSession) Running() (bool, error) {
	return _Monitor.Contract.Running(&_Monitor.CallOpts)
}

// Running is a free data retrieval call binding the contract method 0xd85bd526.
//
// Solidity: function running() view returns(bool)
func (_Monitor *MonitorCallerSession) Running() (bool, error) {
	return _Monitor.Contract.Running(&_Monitor.CallOpts)
}

// Start is a paid mutator transaction binding the contract method 0x258e60b6.
//
// Solidity: function start(string _deviceId) returns()
func (_Monitor *MonitorTransactor) Start(opts *bind.TransactOpts, _deviceId string) (*types.Transaction, error) {
	return _Monitor.contract.Transact(opts, "start", _deviceId)
}

// Start is a paid mutator transaction binding the contract method 0x258e60b6.
//
// Solidity: function start(string _deviceId) returns()
func (_Monitor *MonitorSession) Start(_deviceId string) (*types.Transaction, error) {
	return _Monitor.Contract.Start(&_Monitor.TransactOpts, _deviceId)
}

// Start is a paid mutator transaction binding the contract method 0x258e60b6.
//
// Solidity: function start(string _deviceId) returns()
func (_Monitor *MonitorTransactorSession) Start(_deviceId string) (*types.Transaction, error) {
	return _Monitor.Contract.Start(&_Monitor.TransactOpts, _deviceId)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Monitor *MonitorTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Monitor.contract.Transact(opts, "stop")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Monitor *MonitorSession) Stop() (*types.Transaction, error) {
	return _Monitor.Contract.Stop(&_Monitor.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Monitor *MonitorTransactorSession) Stop() (*types.Transaction, error) {
	return _Monitor.Contract.Stop(&_Monitor.TransactOpts)
}

// Update is a paid mutator transaction binding the contract method 0xc580032f.
//
// Solidity: function update(uint32 _delta) returns()
func (_Monitor *MonitorTransactor) Update(opts *bind.TransactOpts, _delta uint32) (*types.Transaction, error) {
	return _Monitor.contract.Transact(opts, "update", _delta)
}

// Update is a paid mutator transaction binding the contract method 0xc580032f.
//
// Solidity: function update(uint32 _delta) returns()
func (_Monitor *MonitorSession) Update(_delta uint32) (*types.Transaction, error) {
	return _Monitor.Contract.Update(&_Monitor.TransactOpts, _delta)
}

// Update is a paid mutator transaction binding the contract method 0xc580032f.
//
// Solidity: function update(uint32 _delta) returns()
func (_Monitor *MonitorTransactorSession) Update(_delta uint32) (*types.Transaction, error) {
	return _Monitor.Contract.Update(&_Monitor.TransactOpts, _delta)
}
