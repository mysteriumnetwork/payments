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

// AccountantImplementationProxyABI is the input ABI used to generate the binding from.
const AccountantImplementationProxyABI = "[{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// AccountantImplementationProxyBin is the compiled bytecode used for deploying new contracts.
var AccountantImplementationProxyBin = "0x608060405234801561001057600080fd5b5060f48061001f6000396000f3fe608060408190526321f8a72160e01b81527fe6906d4b6048dd18329c27945d05f766dd19b003dc60f82fd4037c490ee55be0608452600090730a0aa1711df0a972655914244507d0f6fa852b6f906321f8a7219060a49060209060248186803b158015606a57600080fd5b505afa158015607d573d6000803e3d6000fd5b505050506040513d6020811015609257600080fd5b505160405190915036600082376000803683856127105a03f43d806000843e81801560bb578184f35b8184fdfea265627a7a72315820b0faa46352f514aacb3c1ebb08166b9b46da8899f4704658f72f7ad73ec8e31964736f6c63430005110032"

// DeployAccountantImplementationProxy deploys a new Ethereum contract, binding an instance of AccountantImplementationProxy to it.
func DeployAccountantImplementationProxy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AccountantImplementationProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountantImplementationProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AccountantImplementationProxyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AccountantImplementationProxy{AccountantImplementationProxyCaller: AccountantImplementationProxyCaller{contract: contract}, AccountantImplementationProxyTransactor: AccountantImplementationProxyTransactor{contract: contract}, AccountantImplementationProxyFilterer: AccountantImplementationProxyFilterer{contract: contract}}, nil
}

// AccountantImplementationProxy is an auto generated Go binding around an Ethereum contract.
type AccountantImplementationProxy struct {
	AccountantImplementationProxyCaller     // Read-only binding to the contract
	AccountantImplementationProxyTransactor // Write-only binding to the contract
	AccountantImplementationProxyFilterer   // Log filterer for contract events
}

// AccountantImplementationProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountantImplementationProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountantImplementationProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountantImplementationProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountantImplementationProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccountantImplementationProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountantImplementationProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountantImplementationProxySession struct {
	Contract     *AccountantImplementationProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// AccountantImplementationProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountantImplementationProxyCallerSession struct {
	Contract *AccountantImplementationProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// AccountantImplementationProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountantImplementationProxyTransactorSession struct {
	Contract     *AccountantImplementationProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// AccountantImplementationProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountantImplementationProxyRaw struct {
	Contract *AccountantImplementationProxy // Generic contract binding to access the raw methods on
}

// AccountantImplementationProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountantImplementationProxyCallerRaw struct {
	Contract *AccountantImplementationProxyCaller // Generic read-only contract binding to access the raw methods on
}

// AccountantImplementationProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountantImplementationProxyTransactorRaw struct {
	Contract *AccountantImplementationProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccountantImplementationProxy creates a new instance of AccountantImplementationProxy, bound to a specific deployed contract.
func NewAccountantImplementationProxy(address common.Address, backend bind.ContractBackend) (*AccountantImplementationProxy, error) {
	contract, err := bindAccountantImplementationProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccountantImplementationProxy{AccountantImplementationProxyCaller: AccountantImplementationProxyCaller{contract: contract}, AccountantImplementationProxyTransactor: AccountantImplementationProxyTransactor{contract: contract}, AccountantImplementationProxyFilterer: AccountantImplementationProxyFilterer{contract: contract}}, nil
}

// NewAccountantImplementationProxyCaller creates a new read-only instance of AccountantImplementationProxy, bound to a specific deployed contract.
func NewAccountantImplementationProxyCaller(address common.Address, caller bind.ContractCaller) (*AccountantImplementationProxyCaller, error) {
	contract, err := bindAccountantImplementationProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccountantImplementationProxyCaller{contract: contract}, nil
}

// NewAccountantImplementationProxyTransactor creates a new write-only instance of AccountantImplementationProxy, bound to a specific deployed contract.
func NewAccountantImplementationProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountantImplementationProxyTransactor, error) {
	contract, err := bindAccountantImplementationProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccountantImplementationProxyTransactor{contract: contract}, nil
}

// NewAccountantImplementationProxyFilterer creates a new log filterer instance of AccountantImplementationProxy, bound to a specific deployed contract.
func NewAccountantImplementationProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*AccountantImplementationProxyFilterer, error) {
	contract, err := bindAccountantImplementationProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccountantImplementationProxyFilterer{contract: contract}, nil
}

// bindAccountantImplementationProxy binds a generic wrapper to an already deployed contract.
func bindAccountantImplementationProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountantImplementationProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountantImplementationProxy *AccountantImplementationProxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AccountantImplementationProxy.Contract.AccountantImplementationProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountantImplementationProxy *AccountantImplementationProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountantImplementationProxy.Contract.AccountantImplementationProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountantImplementationProxy *AccountantImplementationProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountantImplementationProxy.Contract.AccountantImplementationProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountantImplementationProxy *AccountantImplementationProxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AccountantImplementationProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountantImplementationProxy *AccountantImplementationProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountantImplementationProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountantImplementationProxy *AccountantImplementationProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountantImplementationProxy.Contract.contract.Transact(opts, method, params...)
}
