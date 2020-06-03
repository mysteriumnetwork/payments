/* Mysterium network payment library.
 *
 * Copyright (C) 2020 BlockDev AG
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ConfigABI is the input ABI used to generate the binding from.
const ConfigABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"getConfig\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"res\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"res\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"getUint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"res\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"addConfig\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ConfigBin is the compiled bytecode used for deploying new contracts.
var ConfigBin = "0x608060405234801561001057600080fd5b506101d5806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806313af40351461006757806321f8a7211461008f5780636dd5b69d146100c85780638da5cb5b146100f0578063bd02d0f5146100c8578063c4909a70146100f8575b600080fd5b61008d6004803603602081101561007d57600080fd5b50356001600160a01b031661011b565b005b6100ac600480360360208110156100a557600080fd5b5035610168565b604080516001600160a01b039092168252519081900360200190f35b6100de600480360360208110156100a557600080fd5b60408051918252519081900360200190f35b6100ac61016c565b61008d6004803603604081101561010e57600080fd5b508035906020013561017b565b6000546001600160a01b031633148061013d57506000546001600160a01b0316155b61014657600080fd5b600080546001600160a01b0319166001600160a01b0392909216919091179055565b5490565b6000546001600160a01b031681565b6000546001600160a01b0316331461019257600080fd5b8161019c57600080fd5b905556fea265627a7a7231582074078bc0d1cbfdf84aceec13aa72eb9b7eb7d87bcf5e52688a6a1f595b7df79d64736f6c63430005110032"

// DeployConfig deploys a new Ethereum contract, binding an instance of Config to it.
func DeployConfig(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Config, error) {
	parsed, err := abi.JSON(strings.NewReader(ConfigABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ConfigBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Config{ConfigCaller: ConfigCaller{contract: contract}, ConfigTransactor: ConfigTransactor{contract: contract}, ConfigFilterer: ConfigFilterer{contract: contract}}, nil
}

// Config is an auto generated Go binding around an Ethereum contract.
type Config struct {
	ConfigCaller     // Read-only binding to the contract
	ConfigTransactor // Write-only binding to the contract
	ConfigFilterer   // Log filterer for contract events
}

// ConfigCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConfigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConfigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConfigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConfigFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConfigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConfigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConfigSession struct {
	Contract     *Config           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConfigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConfigCallerSession struct {
	Contract *ConfigCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ConfigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConfigTransactorSession struct {
	Contract     *ConfigTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConfigRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConfigRaw struct {
	Contract *Config // Generic contract binding to access the raw methods on
}

// ConfigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConfigCallerRaw struct {
	Contract *ConfigCaller // Generic read-only contract binding to access the raw methods on
}

// ConfigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConfigTransactorRaw struct {
	Contract *ConfigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConfig creates a new instance of Config, bound to a specific deployed contract.
func NewConfig(address common.Address, backend bind.ContractBackend) (*Config, error) {
	contract, err := bindConfig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Config{ConfigCaller: ConfigCaller{contract: contract}, ConfigTransactor: ConfigTransactor{contract: contract}, ConfigFilterer: ConfigFilterer{contract: contract}}, nil
}

// NewConfigCaller creates a new read-only instance of Config, bound to a specific deployed contract.
func NewConfigCaller(address common.Address, caller bind.ContractCaller) (*ConfigCaller, error) {
	contract, err := bindConfig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConfigCaller{contract: contract}, nil
}

// NewConfigTransactor creates a new write-only instance of Config, bound to a specific deployed contract.
func NewConfigTransactor(address common.Address, transactor bind.ContractTransactor) (*ConfigTransactor, error) {
	contract, err := bindConfig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConfigTransactor{contract: contract}, nil
}

// NewConfigFilterer creates a new log filterer instance of Config, bound to a specific deployed contract.
func NewConfigFilterer(address common.Address, filterer bind.ContractFilterer) (*ConfigFilterer, error) {
	contract, err := bindConfig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConfigFilterer{contract: contract}, nil
}

// bindConfig binds a generic wrapper to an already deployed contract.
func bindConfig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConfigABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Config *ConfigRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Config.Contract.ConfigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Config *ConfigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Config.Contract.ConfigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Config *ConfigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Config.Contract.ConfigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Config *ConfigCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Config.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Config *ConfigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Config.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Config *ConfigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Config.Contract.contract.Transact(opts, method, params...)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 slot) constant returns(address res)
func (_Config *ConfigCaller) GetAddress(opts *bind.CallOpts, slot [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Config.contract.Call(opts, out, "getAddress", slot)
	return *ret0, err
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 slot) constant returns(address res)
func (_Config *ConfigSession) GetAddress(slot [32]byte) (common.Address, error) {
	return _Config.Contract.GetAddress(&_Config.CallOpts, slot)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 slot) constant returns(address res)
func (_Config *ConfigCallerSession) GetAddress(slot [32]byte) (common.Address, error) {
	return _Config.Contract.GetAddress(&_Config.CallOpts, slot)
}

// GetConfig is a free data retrieval call binding the contract method 0x6dd5b69d.
//
// Solidity: function getConfig(bytes32 slot) constant returns(bytes32 res)
func (_Config *ConfigCaller) GetConfig(opts *bind.CallOpts, slot [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Config.contract.Call(opts, out, "getConfig", slot)
	return *ret0, err
}

// GetConfig is a free data retrieval call binding the contract method 0x6dd5b69d.
//
// Solidity: function getConfig(bytes32 slot) constant returns(bytes32 res)
func (_Config *ConfigSession) GetConfig(slot [32]byte) ([32]byte, error) {
	return _Config.Contract.GetConfig(&_Config.CallOpts, slot)
}

// GetConfig is a free data retrieval call binding the contract method 0x6dd5b69d.
//
// Solidity: function getConfig(bytes32 slot) constant returns(bytes32 res)
func (_Config *ConfigCallerSession) GetConfig(slot [32]byte) ([32]byte, error) {
	return _Config.Contract.GetConfig(&_Config.CallOpts, slot)
}

// GetUint is a free data retrieval call binding the contract method 0xbd02d0f5.
//
// Solidity: function getUint(bytes32 slot) constant returns(uint256 res)
func (_Config *ConfigCaller) GetUint(opts *bind.CallOpts, slot [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Config.contract.Call(opts, out, "getUint", slot)
	return *ret0, err
}

// GetUint is a free data retrieval call binding the contract method 0xbd02d0f5.
//
// Solidity: function getUint(bytes32 slot) constant returns(uint256 res)
func (_Config *ConfigSession) GetUint(slot [32]byte) (*big.Int, error) {
	return _Config.Contract.GetUint(&_Config.CallOpts, slot)
}

// GetUint is a free data retrieval call binding the contract method 0xbd02d0f5.
//
// Solidity: function getUint(bytes32 slot) constant returns(uint256 res)
func (_Config *ConfigCallerSession) GetUint(slot [32]byte) (*big.Int, error) {
	return _Config.Contract.GetUint(&_Config.CallOpts, slot)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Config *ConfigCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Config.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Config *ConfigSession) Owner() (common.Address, error) {
	return _Config.Contract.Owner(&_Config.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Config *ConfigCallerSession) Owner() (common.Address, error) {
	return _Config.Contract.Owner(&_Config.CallOpts)
}

// AddConfig is a paid mutator transaction binding the contract method 0xc4909a70.
//
// Solidity: function addConfig(bytes32 slot, bytes32 value) returns()
func (_Config *ConfigTransactor) AddConfig(opts *bind.TransactOpts, slot [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "addConfig", slot, value)
}

// AddConfig is a paid mutator transaction binding the contract method 0xc4909a70.
//
// Solidity: function addConfig(bytes32 slot, bytes32 value) returns()
func (_Config *ConfigSession) AddConfig(slot [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Config.Contract.AddConfig(&_Config.TransactOpts, slot, value)
}

// AddConfig is a paid mutator transaction binding the contract method 0xc4909a70.
//
// Solidity: function addConfig(bytes32 slot, bytes32 value) returns()
func (_Config *ConfigTransactorSession) AddConfig(slot [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Config.Contract.AddConfig(&_Config.TransactOpts, slot, value)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_Config *ConfigTransactor) SetOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "setOwner", _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_Config *ConfigSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _Config.Contract.SetOwner(&_Config.TransactOpts, _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_Config *ConfigTransactorSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _Config.Contract.SetOwner(&_Config.TransactOpts, _owner)
}
