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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// HermesContractABI is the input ABI used to generate the binding from.
const HermesContractABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_hermesFee\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_maxStake\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_party\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountToLend\",\"type\":\"uint256\"}],\"name\":\"openChannel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStatus\",\"outputs\":[{\"internalType\":\"enumHermesContract.Status\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// HermesContract is an auto generated Go binding around an Ethereum contract.
type HermesContract struct {
	HermesContractCaller     // Read-only binding to the contract
	HermesContractTransactor // Write-only binding to the contract
	HermesContractFilterer   // Log filterer for contract events
}

// HermesContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type HermesContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HermesContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HermesContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HermesContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HermesContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HermesContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HermesContractSession struct {
	Contract     *HermesContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HermesContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HermesContractCallerSession struct {
	Contract *HermesContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// HermesContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HermesContractTransactorSession struct {
	Contract     *HermesContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// HermesContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type HermesContractRaw struct {
	Contract *HermesContract // Generic contract binding to access the raw methods on
}

// HermesContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HermesContractCallerRaw struct {
	Contract *HermesContractCaller // Generic read-only contract binding to access the raw methods on
}

// HermesContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HermesContractTransactorRaw struct {
	Contract *HermesContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHermesContract creates a new instance of HermesContract, bound to a specific deployed contract.
func NewHermesContract(address common.Address, backend bind.ContractBackend) (*HermesContract, error) {
	contract, err := bindHermesContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HermesContract{HermesContractCaller: HermesContractCaller{contract: contract}, HermesContractTransactor: HermesContractTransactor{contract: contract}, HermesContractFilterer: HermesContractFilterer{contract: contract}}, nil
}

// NewHermesContractCaller creates a new read-only instance of HermesContract, bound to a specific deployed contract.
func NewHermesContractCaller(address common.Address, caller bind.ContractCaller) (*HermesContractCaller, error) {
	contract, err := bindHermesContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HermesContractCaller{contract: contract}, nil
}

// NewHermesContractTransactor creates a new write-only instance of HermesContract, bound to a specific deployed contract.
func NewHermesContractTransactor(address common.Address, transactor bind.ContractTransactor) (*HermesContractTransactor, error) {
	contract, err := bindHermesContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HermesContractTransactor{contract: contract}, nil
}

// NewHermesContractFilterer creates a new log filterer instance of HermesContract, bound to a specific deployed contract.
func NewHermesContractFilterer(address common.Address, filterer bind.ContractFilterer) (*HermesContractFilterer, error) {
	contract, err := bindHermesContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HermesContractFilterer{contract: contract}, nil
}

// bindHermesContract binds a generic wrapper to an already deployed contract.
func bindHermesContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HermesContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HermesContract *HermesContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HermesContract.Contract.HermesContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HermesContract *HermesContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HermesContract.Contract.HermesContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HermesContract *HermesContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HermesContract.Contract.HermesContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HermesContract *HermesContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HermesContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HermesContract *HermesContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HermesContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HermesContract *HermesContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HermesContract.Contract.contract.Transact(opts, method, params...)
}

// GetStake is a free data retrieval call binding the contract method 0xfc0e3d90.
//
// Solidity: function getStake() view returns(uint256)
func (_HermesContract *HermesContractCaller) GetStake(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HermesContract.contract.Call(opts, out, "getStake")
	return *ret0, err
}

// GetStake is a free data retrieval call binding the contract method 0xfc0e3d90.
//
// Solidity: function getStake() view returns(uint256)
func (_HermesContract *HermesContractSession) GetStake() (*big.Int, error) {
	return _HermesContract.Contract.GetStake(&_HermesContract.CallOpts)
}

// GetStake is a free data retrieval call binding the contract method 0xfc0e3d90.
//
// Solidity: function getStake() view returns(uint256)
func (_HermesContract *HermesContractCallerSession) GetStake() (*big.Int, error) {
	return _HermesContract.Contract.GetStake(&_HermesContract.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() view returns(uint8)
func (_HermesContract *HermesContractCaller) GetStatus(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _HermesContract.contract.Call(opts, out, "getStatus")
	return *ret0, err
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() view returns(uint8)
func (_HermesContract *HermesContractSession) GetStatus() (uint8, error) {
	return _HermesContract.Contract.GetStatus(&_HermesContract.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() view returns(uint8)
func (_HermesContract *HermesContractCallerSession) GetStatus() (uint8, error) {
	return _HermesContract.Contract.GetStatus(&_HermesContract.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xfec8157d.
//
// Solidity: function initialize(address _token, address _operator, uint16 _hermesFee, uint256 _maxStake) returns()
func (_HermesContract *HermesContractTransactor) Initialize(opts *bind.TransactOpts, _token common.Address, _operator common.Address, _hermesFee uint16, _maxStake *big.Int) (*types.Transaction, error) {
	return _HermesContract.contract.Transact(opts, "initialize", _token, _operator, _hermesFee, _maxStake)
}

// Initialize is a paid mutator transaction binding the contract method 0xfec8157d.
//
// Solidity: function initialize(address _token, address _operator, uint16 _hermesFee, uint256 _maxStake) returns()
func (_HermesContract *HermesContractSession) Initialize(_token common.Address, _operator common.Address, _hermesFee uint16, _maxStake *big.Int) (*types.Transaction, error) {
	return _HermesContract.Contract.Initialize(&_HermesContract.TransactOpts, _token, _operator, _hermesFee, _maxStake)
}

// Initialize is a paid mutator transaction binding the contract method 0xfec8157d.
//
// Solidity: function initialize(address _token, address _operator, uint16 _hermesFee, uint256 _maxStake) returns()
func (_HermesContract *HermesContractTransactorSession) Initialize(_token common.Address, _operator common.Address, _hermesFee uint16, _maxStake *big.Int) (*types.Transaction, error) {
	return _HermesContract.Contract.Initialize(&_HermesContract.TransactOpts, _token, _operator, _hermesFee, _maxStake)
}

// OpenChannel is a paid mutator transaction binding the contract method 0x0a798f24.
//
// Solidity: function openChannel(address _party, address _beneficiary, uint256 _amountToLend) returns()
func (_HermesContract *HermesContractTransactor) OpenChannel(opts *bind.TransactOpts, _party common.Address, _beneficiary common.Address, _amountToLend *big.Int) (*types.Transaction, error) {
	return _HermesContract.contract.Transact(opts, "openChannel", _party, _beneficiary, _amountToLend)
}

// OpenChannel is a paid mutator transaction binding the contract method 0x0a798f24.
//
// Solidity: function openChannel(address _party, address _beneficiary, uint256 _amountToLend) returns()
func (_HermesContract *HermesContractSession) OpenChannel(_party common.Address, _beneficiary common.Address, _amountToLend *big.Int) (*types.Transaction, error) {
	return _HermesContract.Contract.OpenChannel(&_HermesContract.TransactOpts, _party, _beneficiary, _amountToLend)
}

// OpenChannel is a paid mutator transaction binding the contract method 0x0a798f24.
//
// Solidity: function openChannel(address _party, address _beneficiary, uint256 _amountToLend) returns()
func (_HermesContract *HermesContractTransactorSession) OpenChannel(_party common.Address, _beneficiary common.Address, _amountToLend *big.Int) (*types.Transaction, error) {
	return _HermesContract.Contract.OpenChannel(&_HermesContract.TransactOpts, _party, _beneficiary, _amountToLend)
}
