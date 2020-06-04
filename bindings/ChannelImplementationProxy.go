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

// ChannelImplementationProxyABI is the input ABI used to generate the binding from.
const ChannelImplementationProxyABI = "[{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// ChannelImplementationProxyBin is the compiled bytecode used for deploying new contracts.
var ChannelImplementationProxyBin = "0x608060405234801561001057600080fd5b5060f48061001f6000396000f3fe608060408190526321f8a72160e01b81527f48df65c92c1c0e8e19a219c69bfeb4cf7c1c123e0c266d555abb508d37c6d96e608452600090730a0aa1711df0a972655914244507d0f6fa852b6f906321f8a7219060a49060209060248186803b158015606a57600080fd5b505afa158015607d573d6000803e3d6000fd5b505050506040513d6020811015609257600080fd5b505160405190915036600082376000803683856127105a03f43d806000843e81801560bb578184f35b8184fdfea265627a7a7231582063b9545dff3e0070d96ad42546cc2fb122d6e9bcd9f6783fa724f8cb4886a23564736f6c63430005110032"

// DeployChannelImplementationProxy deploys a new Ethereum contract, binding an instance of ChannelImplementationProxy to it.
func DeployChannelImplementationProxy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ChannelImplementationProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(ChannelImplementationProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChannelImplementationProxyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChannelImplementationProxy{ChannelImplementationProxyCaller: ChannelImplementationProxyCaller{contract: contract}, ChannelImplementationProxyTransactor: ChannelImplementationProxyTransactor{contract: contract}, ChannelImplementationProxyFilterer: ChannelImplementationProxyFilterer{contract: contract}}, nil
}

// ChannelImplementationProxy is an auto generated Go binding around an Ethereum contract.
type ChannelImplementationProxy struct {
	ChannelImplementationProxyCaller     // Read-only binding to the contract
	ChannelImplementationProxyTransactor // Write-only binding to the contract
	ChannelImplementationProxyFilterer   // Log filterer for contract events
}

// ChannelImplementationProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChannelImplementationProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChannelImplementationProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChannelImplementationProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChannelImplementationProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChannelImplementationProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChannelImplementationProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChannelImplementationProxySession struct {
	Contract     *ChannelImplementationProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ChannelImplementationProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChannelImplementationProxyCallerSession struct {
	Contract *ChannelImplementationProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// ChannelImplementationProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChannelImplementationProxyTransactorSession struct {
	Contract     *ChannelImplementationProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// ChannelImplementationProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChannelImplementationProxyRaw struct {
	Contract *ChannelImplementationProxy // Generic contract binding to access the raw methods on
}

// ChannelImplementationProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChannelImplementationProxyCallerRaw struct {
	Contract *ChannelImplementationProxyCaller // Generic read-only contract binding to access the raw methods on
}

// ChannelImplementationProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChannelImplementationProxyTransactorRaw struct {
	Contract *ChannelImplementationProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChannelImplementationProxy creates a new instance of ChannelImplementationProxy, bound to a specific deployed contract.
func NewChannelImplementationProxy(address common.Address, backend bind.ContractBackend) (*ChannelImplementationProxy, error) {
	contract, err := bindChannelImplementationProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChannelImplementationProxy{ChannelImplementationProxyCaller: ChannelImplementationProxyCaller{contract: contract}, ChannelImplementationProxyTransactor: ChannelImplementationProxyTransactor{contract: contract}, ChannelImplementationProxyFilterer: ChannelImplementationProxyFilterer{contract: contract}}, nil
}

// NewChannelImplementationProxyCaller creates a new read-only instance of ChannelImplementationProxy, bound to a specific deployed contract.
func NewChannelImplementationProxyCaller(address common.Address, caller bind.ContractCaller) (*ChannelImplementationProxyCaller, error) {
	contract, err := bindChannelImplementationProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChannelImplementationProxyCaller{contract: contract}, nil
}

// NewChannelImplementationProxyTransactor creates a new write-only instance of ChannelImplementationProxy, bound to a specific deployed contract.
func NewChannelImplementationProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*ChannelImplementationProxyTransactor, error) {
	contract, err := bindChannelImplementationProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChannelImplementationProxyTransactor{contract: contract}, nil
}

// NewChannelImplementationProxyFilterer creates a new log filterer instance of ChannelImplementationProxy, bound to a specific deployed contract.
func NewChannelImplementationProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*ChannelImplementationProxyFilterer, error) {
	contract, err := bindChannelImplementationProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChannelImplementationProxyFilterer{contract: contract}, nil
}

// bindChannelImplementationProxy binds a generic wrapper to an already deployed contract.
func bindChannelImplementationProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChannelImplementationProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChannelImplementationProxy *ChannelImplementationProxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChannelImplementationProxy.Contract.ChannelImplementationProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChannelImplementationProxy *ChannelImplementationProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChannelImplementationProxy.Contract.ChannelImplementationProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChannelImplementationProxy *ChannelImplementationProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChannelImplementationProxy.Contract.ChannelImplementationProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChannelImplementationProxy *ChannelImplementationProxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChannelImplementationProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChannelImplementationProxy *ChannelImplementationProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChannelImplementationProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChannelImplementationProxy *ChannelImplementationProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChannelImplementationProxy.Contract.contract.Transact(opts, method, params...)
}
