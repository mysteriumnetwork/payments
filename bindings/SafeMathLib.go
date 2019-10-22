/*
 * Copyright (C) 2019 The "MysteriumNetwork/payments" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
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

// SafeMathLibABI is the input ABI used to generate the binding from.
const SafeMathLibABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"times\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"minus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"plus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// SafeMathLibBin is the compiled bytecode used for deploying new contracts.
var SafeMathLibBin = "0x61031f610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061004b5760003560e01c80631d3b9edf1461005057806366098d4f1461009c578063f4f3bdc1146100e8575b600080fd5b6100866004803603604081101561006657600080fd5b810190808035906020019092919080359060200190929190505050610134565b6040518082815260200191505060405180910390f35b6100d2600480360360408110156100b257600080fd5b8101908080359060200190929190803590602001909291905050506101b2565b6040518082815260200191505060405180910390f35b61011e600480360360408110156100fe57600080fd5b810190808035906020019092919080359060200190929190505050610246565b6040518082815260200191505060405180910390f35b6000808284029050600084148061015357508284828161015057fe5b04145b6101a8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001806102ca6021913960400191505060405180910390fd5b8091505092915050565b60008082840190508381101580156101ca5750828110155b61023c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b6000828211156102be576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081525060200191505060405180910390fd5b81830390509291505056fe536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77a265627a7a723158201371c066e7508ce1296e8973a258bdb0b97919c380ce1aeca24ccd49ec2e03e964736f6c634300050c0032"

// DeploySafeMathLib deploys a new Ethereum contract, binding an instance of SafeMathLib to it.
func DeploySafeMathLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMathLib, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathLibABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMathLib{SafeMathLibCaller: SafeMathLibCaller{contract: contract}, SafeMathLibTransactor: SafeMathLibTransactor{contract: contract}, SafeMathLibFilterer: SafeMathLibFilterer{contract: contract}}, nil
}

// SafeMathLib is an auto generated Go binding around an Ethereum contract.
type SafeMathLib struct {
	SafeMathLibCaller     // Read-only binding to the contract
	SafeMathLibTransactor // Write-only binding to the contract
	SafeMathLibFilterer   // Log filterer for contract events
}

// SafeMathLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathLibSession struct {
	Contract     *SafeMathLib      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathLibCallerSession struct {
	Contract *SafeMathLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SafeMathLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathLibTransactorSession struct {
	Contract     *SafeMathLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SafeMathLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathLibRaw struct {
	Contract *SafeMathLib // Generic contract binding to access the raw methods on
}

// SafeMathLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathLibCallerRaw struct {
	Contract *SafeMathLibCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathLibTransactorRaw struct {
	Contract *SafeMathLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMathLib creates a new instance of SafeMathLib, bound to a specific deployed contract.
func NewSafeMathLib(address common.Address, backend bind.ContractBackend) (*SafeMathLib, error) {
	contract, err := bindSafeMathLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMathLib{SafeMathLibCaller: SafeMathLibCaller{contract: contract}, SafeMathLibTransactor: SafeMathLibTransactor{contract: contract}, SafeMathLibFilterer: SafeMathLibFilterer{contract: contract}}, nil
}

// NewSafeMathLibCaller creates a new read-only instance of SafeMathLib, bound to a specific deployed contract.
func NewSafeMathLibCaller(address common.Address, caller bind.ContractCaller) (*SafeMathLibCaller, error) {
	contract, err := bindSafeMathLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathLibCaller{contract: contract}, nil
}

// NewSafeMathLibTransactor creates a new write-only instance of SafeMathLib, bound to a specific deployed contract.
func NewSafeMathLibTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathLibTransactor, error) {
	contract, err := bindSafeMathLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathLibTransactor{contract: contract}, nil
}

// NewSafeMathLibFilterer creates a new log filterer instance of SafeMathLib, bound to a specific deployed contract.
func NewSafeMathLibFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathLibFilterer, error) {
	contract, err := bindSafeMathLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathLibFilterer{contract: contract}, nil
}

// bindSafeMathLib binds a generic wrapper to an already deployed contract.
func bindSafeMathLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMathLib *SafeMathLibRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMathLib.Contract.SafeMathLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMathLib *SafeMathLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMathLib.Contract.SafeMathLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMathLib *SafeMathLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMathLib.Contract.SafeMathLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMathLib *SafeMathLibCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMathLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMathLib *SafeMathLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMathLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMathLib *SafeMathLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMathLib.Contract.contract.Transact(opts, method, params...)
}

// Minus is a free data retrieval call binding the contract method 0xf4f3bdc1.
//
// Solidity: function minus(uint256 a, uint256 b) constant returns(uint256)
func (_SafeMathLib *SafeMathLibCaller) Minus(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeMathLib.contract.Call(opts, out, "minus", a, b)
	return *ret0, err
}

// Minus is a free data retrieval call binding the contract method 0xf4f3bdc1.
//
// Solidity: function minus(uint256 a, uint256 b) constant returns(uint256)
func (_SafeMathLib *SafeMathLibSession) Minus(a *big.Int, b *big.Int) (*big.Int, error) {
	return _SafeMathLib.Contract.Minus(&_SafeMathLib.CallOpts, a, b)
}

// Minus is a free data retrieval call binding the contract method 0xf4f3bdc1.
//
// Solidity: function minus(uint256 a, uint256 b) constant returns(uint256)
func (_SafeMathLib *SafeMathLibCallerSession) Minus(a *big.Int, b *big.Int) (*big.Int, error) {
	return _SafeMathLib.Contract.Minus(&_SafeMathLib.CallOpts, a, b)
}

// Plus is a free data retrieval call binding the contract method 0x66098d4f.
//
// Solidity: function plus(uint256 a, uint256 b) constant returns(uint256)
func (_SafeMathLib *SafeMathLibCaller) Plus(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeMathLib.contract.Call(opts, out, "plus", a, b)
	return *ret0, err
}

// Plus is a free data retrieval call binding the contract method 0x66098d4f.
//
// Solidity: function plus(uint256 a, uint256 b) constant returns(uint256)
func (_SafeMathLib *SafeMathLibSession) Plus(a *big.Int, b *big.Int) (*big.Int, error) {
	return _SafeMathLib.Contract.Plus(&_SafeMathLib.CallOpts, a, b)
}

// Plus is a free data retrieval call binding the contract method 0x66098d4f.
//
// Solidity: function plus(uint256 a, uint256 b) constant returns(uint256)
func (_SafeMathLib *SafeMathLibCallerSession) Plus(a *big.Int, b *big.Int) (*big.Int, error) {
	return _SafeMathLib.Contract.Plus(&_SafeMathLib.CallOpts, a, b)
}

// Times is a free data retrieval call binding the contract method 0x1d3b9edf.
//
// Solidity: function times(uint256 a, uint256 b) constant returns(uint256)
func (_SafeMathLib *SafeMathLibCaller) Times(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeMathLib.contract.Call(opts, out, "times", a, b)
	return *ret0, err
}

// Times is a free data retrieval call binding the contract method 0x1d3b9edf.
//
// Solidity: function times(uint256 a, uint256 b) constant returns(uint256)
func (_SafeMathLib *SafeMathLibSession) Times(a *big.Int, b *big.Int) (*big.Int, error) {
	return _SafeMathLib.Contract.Times(&_SafeMathLib.CallOpts, a, b)
}

// Times is a free data retrieval call binding the contract method 0x1d3b9edf.
//
// Solidity: function times(uint256 a, uint256 b) constant returns(uint256)
func (_SafeMathLib *SafeMathLibCallerSession) Times(a *big.Int, b *big.Int) (*big.Int, error) {
	return _SafeMathLib.Contract.Times(&_SafeMathLib.CallOpts, a, b)
}
