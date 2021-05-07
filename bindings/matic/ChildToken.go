/* Mysterium network payment library.
 *
 * Copyright (C) 2021 BlockDev AG
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

package matic

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

// ChildTokenABI is the input ABI used to generate the binding from.
const ChildTokenABI = "[{\"type\":\"constructor\",\"stateMutability\":\"nonpayable\",\"inputs\":[]},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"spender\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MetaTransactionExecuted\",\"inputs\":[{\"type\":\"address\",\"name\":\"userAddress\",\"internalType\":\"address\",\"indexed\":false},{\"type\":\"address\",\"name\":\"relayerAddress\",\"internalType\":\"addresspayable\",\"indexed\":false},{\"type\":\"bytes\",\"name\":\"functionSignature\",\"internalType\":\"bytes\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"internalType\":\"bytes32\",\"indexed\":true},{\"type\":\"bytes32\",\"name\":\"previousAdminRole\",\"internalType\":\"bytes32\",\"indexed\":true},{\"type\":\"bytes32\",\"name\":\"newAdminRole\",\"internalType\":\"bytes32\",\"indexed\":true}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"internalType\":\"bytes32\",\"indexed\":true},{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"sender\",\"internalType\":\"address\",\"indexed\":true}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"internalType\":\"bytes32\",\"indexed\":true},{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"sender\",\"internalType\":\"address\",\"indexed\":true}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"CHILD_CHAIN_ID\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bytes\",\"name\":\"\",\"internalType\":\"bytes\"}],\"name\":\"CHILD_CHAIN_ID_BYTES\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"\",\"internalType\":\"bytes32\"}],\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"\",\"internalType\":\"bytes32\"}],\"name\":\"DEPOSITOR_ROLE\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"name\":\"ERC712_VERSION\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"ROOT_CHAIN_ID\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bytes\",\"name\":\"\",\"internalType\":\"bytes\"}],\"name\":\"ROOT_CHAIN_ID_BYTES\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"allowance\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"spender\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"approve\",\"inputs\":[{\"type\":\"address\",\"name\":\"spender\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"balanceOf\",\"inputs\":[{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"uint8\"}],\"name\":\"decimals\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"decreaseAllowance\",\"inputs\":[{\"type\":\"address\",\"name\":\"spender\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"subtractedValue\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"deposit\",\"inputs\":[{\"type\":\"address\",\"name\":\"user\",\"internalType\":\"address\"},{\"type\":\"bytes\",\"name\":\"depositData\",\"internalType\":\"bytes\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[{\"type\":\"bytes\",\"name\":\"\",\"internalType\":\"bytes\"}],\"name\":\"executeMetaTransaction\",\"inputs\":[{\"type\":\"address\",\"name\":\"userAddress\",\"internalType\":\"address\"},{\"type\":\"bytes\",\"name\":\"functionSignature\",\"internalType\":\"bytes\"},{\"type\":\"bytes32\",\"name\":\"sigR\",\"internalType\":\"bytes32\"},{\"type\":\"bytes32\",\"name\":\"sigS\",\"internalType\":\"bytes32\"},{\"type\":\"uint8\",\"name\":\"sigV\",\"internalType\":\"uint8\"}]},{\"type\":\"function\",\"stateMutability\":\"pure\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getChainId\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"\",\"internalType\":\"bytes32\"}],\"name\":\"getDomainSeperator\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"nonce\",\"internalType\":\"uint256\"}],\"name\":\"getNonce\",\"inputs\":[{\"type\":\"address\",\"name\":\"user\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"\",\"internalType\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"getRoleMember\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"internalType\":\"bytes32\"},{\"type\":\"uint256\",\"name\":\"index\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getRoleMemberCount\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"grantRole\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"internalType\":\"bytes32\"},{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"hasRole\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"internalType\":\"bytes32\"},{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"increaseAllowance\",\"inputs\":[{\"type\":\"address\",\"name\":\"spender\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"addedValue\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"initialize\",\"inputs\":[{\"type\":\"string\",\"name\":\"name_\",\"internalType\":\"string\"},{\"type\":\"string\",\"name\":\"symbol_\",\"internalType\":\"string\"},{\"type\":\"uint8\",\"name\":\"decimals_\",\"internalType\":\"uint8\"},{\"type\":\"address\",\"name\":\"childChainManager\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"name\":\"name\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"renounceRole\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"internalType\":\"bytes32\"},{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"revokeRole\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"internalType\":\"bytes32\"},{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"name\":\"symbol\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"totalSupply\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"transfer\",\"inputs\":[{\"type\":\"address\",\"name\":\"recipient\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"transferFrom\",\"inputs\":[{\"type\":\"address\",\"name\":\"sender\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"recipient\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"withdraw\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\"}]}]"

// ChildToken is an auto generated Go binding around an Ethereum contract.
type ChildToken struct {
	ChildTokenCaller     // Read-only binding to the contract
	ChildTokenTransactor // Write-only binding to the contract
	ChildTokenFilterer   // Log filterer for contract events
}

// ChildTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChildTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChildTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChildTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChildTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChildTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChildTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChildTokenSession struct {
	Contract     *ChildToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChildTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChildTokenCallerSession struct {
	Contract *ChildTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ChildTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChildTokenTransactorSession struct {
	Contract     *ChildTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ChildTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChildTokenRaw struct {
	Contract *ChildToken // Generic contract binding to access the raw methods on
}

// ChildTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChildTokenCallerRaw struct {
	Contract *ChildTokenCaller // Generic read-only contract binding to access the raw methods on
}

// ChildTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChildTokenTransactorRaw struct {
	Contract *ChildTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChildToken creates a new instance of ChildToken, bound to a specific deployed contract.
func NewChildToken(address common.Address, backend bind.ContractBackend) (*ChildToken, error) {
	contract, err := bindChildToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChildToken{ChildTokenCaller: ChildTokenCaller{contract: contract}, ChildTokenTransactor: ChildTokenTransactor{contract: contract}, ChildTokenFilterer: ChildTokenFilterer{contract: contract}}, nil
}

// NewChildTokenCaller creates a new read-only instance of ChildToken, bound to a specific deployed contract.
func NewChildTokenCaller(address common.Address, caller bind.ContractCaller) (*ChildTokenCaller, error) {
	contract, err := bindChildToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChildTokenCaller{contract: contract}, nil
}

// NewChildTokenTransactor creates a new write-only instance of ChildToken, bound to a specific deployed contract.
func NewChildTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ChildTokenTransactor, error) {
	contract, err := bindChildToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChildTokenTransactor{contract: contract}, nil
}

// NewChildTokenFilterer creates a new log filterer instance of ChildToken, bound to a specific deployed contract.
func NewChildTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ChildTokenFilterer, error) {
	contract, err := bindChildToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChildTokenFilterer{contract: contract}, nil
}

// bindChildToken binds a generic wrapper to an already deployed contract.
func bindChildToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChildTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChildToken *ChildTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChildToken.Contract.ChildTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChildToken *ChildTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChildToken.Contract.ChildTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChildToken *ChildTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChildToken.Contract.ChildTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChildToken *ChildTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChildToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChildToken *ChildTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChildToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChildToken *ChildTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChildToken.Contract.contract.Transact(opts, method, params...)
}

// CHILDCHAINID is a free data retrieval call binding the contract method 0x626381a0.
//
// Solidity: function CHILD_CHAIN_ID() view returns(uint256)
func (_ChildToken *ChildTokenCaller) CHILDCHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChildToken.contract.Call(opts, &out, "CHILD_CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CHILDCHAINID is a free data retrieval call binding the contract method 0x626381a0.
//
// Solidity: function CHILD_CHAIN_ID() view returns(uint256)
func (_ChildToken *ChildTokenSession) CHILDCHAINID() (*big.Int, error) {
	return _ChildToken.Contract.CHILDCHAINID(&_ChildToken.CallOpts)
}

// CHILDCHAINID is a free data retrieval call binding the contract method 0x626381a0.
//
// Solidity: function CHILD_CHAIN_ID() view returns(uint256)
func (_ChildToken *ChildTokenCallerSession) CHILDCHAINID() (*big.Int, error) {
	return _ChildToken.Contract.CHILDCHAINID(&_ChildToken.CallOpts)
}

// CHILDCHAINIDBYTES is a free data retrieval call binding the contract method 0x0b54817c.
//
// Solidity: function CHILD_CHAIN_ID_BYTES() view returns(bytes)
func (_ChildToken *ChildTokenCaller) CHILDCHAINIDBYTES(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _ChildToken.contract.Call(opts, &out, "CHILD_CHAIN_ID_BYTES")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// CHILDCHAINIDBYTES is a free data retrieval call binding the contract method 0x0b54817c.
//
// Solidity: function CHILD_CHAIN_ID_BYTES() view returns(bytes)
func (_ChildToken *ChildTokenSession) CHILDCHAINIDBYTES() ([]byte, error) {
	return _ChildToken.Contract.CHILDCHAINIDBYTES(&_ChildToken.CallOpts)
}

// CHILDCHAINIDBYTES is a free data retrieval call binding the contract method 0x0b54817c.
//
// Solidity: function CHILD_CHAIN_ID_BYTES() view returns(bytes)
func (_ChildToken *ChildTokenCallerSession) CHILDCHAINIDBYTES() ([]byte, error) {
	return _ChildToken.Contract.CHILDCHAINIDBYTES(&_ChildToken.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ChildToken *ChildTokenCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ChildToken.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ChildToken *ChildTokenSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ChildToken.Contract.DEFAULTADMINROLE(&_ChildToken.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ChildToken *ChildTokenCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ChildToken.Contract.DEFAULTADMINROLE(&_ChildToken.CallOpts)
}

// DEPOSITORROLE is a free data retrieval call binding the contract method 0xa3b0b5a3.
//
// Solidity: function DEPOSITOR_ROLE() view returns(bytes32)
func (_ChildToken *ChildTokenCaller) DEPOSITORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ChildToken.contract.Call(opts, &out, "DEPOSITOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEPOSITORROLE is a free data retrieval call binding the contract method 0xa3b0b5a3.
//
// Solidity: function DEPOSITOR_ROLE() view returns(bytes32)
func (_ChildToken *ChildTokenSession) DEPOSITORROLE() ([32]byte, error) {
	return _ChildToken.Contract.DEPOSITORROLE(&_ChildToken.CallOpts)
}

// DEPOSITORROLE is a free data retrieval call binding the contract method 0xa3b0b5a3.
//
// Solidity: function DEPOSITOR_ROLE() view returns(bytes32)
func (_ChildToken *ChildTokenCallerSession) DEPOSITORROLE() ([32]byte, error) {
	return _ChildToken.Contract.DEPOSITORROLE(&_ChildToken.CallOpts)
}

// ERC712VERSION is a free data retrieval call binding the contract method 0x0f7e5970.
//
// Solidity: function ERC712_VERSION() view returns(string)
func (_ChildToken *ChildTokenCaller) ERC712VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ChildToken.contract.Call(opts, &out, "ERC712_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ERC712VERSION is a free data retrieval call binding the contract method 0x0f7e5970.
//
// Solidity: function ERC712_VERSION() view returns(string)
func (_ChildToken *ChildTokenSession) ERC712VERSION() (string, error) {
	return _ChildToken.Contract.ERC712VERSION(&_ChildToken.CallOpts)
}

// ERC712VERSION is a free data retrieval call binding the contract method 0x0f7e5970.
//
// Solidity: function ERC712_VERSION() view returns(string)
func (_ChildToken *ChildTokenCallerSession) ERC712VERSION() (string, error) {
	return _ChildToken.Contract.ERC712VERSION(&_ChildToken.CallOpts)
}

// ROOTCHAINID is a free data retrieval call binding the contract method 0x8acfcaf7.
//
// Solidity: function ROOT_CHAIN_ID() view returns(uint256)
func (_ChildToken *ChildTokenCaller) ROOTCHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChildToken.contract.Call(opts, &out, "ROOT_CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ROOTCHAINID is a free data retrieval call binding the contract method 0x8acfcaf7.
//
// Solidity: function ROOT_CHAIN_ID() view returns(uint256)
func (_ChildToken *ChildTokenSession) ROOTCHAINID() (*big.Int, error) {
	return _ChildToken.Contract.ROOTCHAINID(&_ChildToken.CallOpts)
}

// ROOTCHAINID is a free data retrieval call binding the contract method 0x8acfcaf7.
//
// Solidity: function ROOT_CHAIN_ID() view returns(uint256)
func (_ChildToken *ChildTokenCallerSession) ROOTCHAINID() (*big.Int, error) {
	return _ChildToken.Contract.ROOTCHAINID(&_ChildToken.CallOpts)
}

// ROOTCHAINIDBYTES is a free data retrieval call binding the contract method 0x0dd7531a.
//
// Solidity: function ROOT_CHAIN_ID_BYTES() view returns(bytes)
func (_ChildToken *ChildTokenCaller) ROOTCHAINIDBYTES(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _ChildToken.contract.Call(opts, &out, "ROOT_CHAIN_ID_BYTES")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ROOTCHAINIDBYTES is a free data retrieval call binding the contract method 0x0dd7531a.
//
// Solidity: function ROOT_CHAIN_ID_BYTES() view returns(bytes)
func (_ChildToken *ChildTokenSession) ROOTCHAINIDBYTES() ([]byte, error) {
	return _ChildToken.Contract.ROOTCHAINIDBYTES(&_ChildToken.CallOpts)
}

// ROOTCHAINIDBYTES is a free data retrieval call binding the contract method 0x0dd7531a.
//
// Solidity: function ROOT_CHAIN_ID_BYTES() view returns(bytes)
func (_ChildToken *ChildTokenCallerSession) ROOTCHAINIDBYTES() ([]byte, error) {
	return _ChildToken.Contract.ROOTCHAINIDBYTES(&_ChildToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ChildToken *ChildTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ChildToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ChildToken *ChildTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ChildToken.Contract.Allowance(&_ChildToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ChildToken *ChildTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ChildToken.Contract.Allowance(&_ChildToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ChildToken *ChildTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ChildToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ChildToken *ChildTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ChildToken.Contract.BalanceOf(&_ChildToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ChildToken *ChildTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ChildToken.Contract.BalanceOf(&_ChildToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ChildToken *ChildTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ChildToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ChildToken *ChildTokenSession) Decimals() (uint8, error) {
	return _ChildToken.Contract.Decimals(&_ChildToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ChildToken *ChildTokenCallerSession) Decimals() (uint8, error) {
	return _ChildToken.Contract.Decimals(&_ChildToken.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() pure returns(uint256)
func (_ChildToken *ChildTokenCaller) GetChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChildToken.contract.Call(opts, &out, "getChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() pure returns(uint256)
func (_ChildToken *ChildTokenSession) GetChainId() (*big.Int, error) {
	return _ChildToken.Contract.GetChainId(&_ChildToken.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() pure returns(uint256)
func (_ChildToken *ChildTokenCallerSession) GetChainId() (*big.Int, error) {
	return _ChildToken.Contract.GetChainId(&_ChildToken.CallOpts)
}

// GetDomainSeperator is a free data retrieval call binding the contract method 0x20379ee5.
//
// Solidity: function getDomainSeperator() view returns(bytes32)
func (_ChildToken *ChildTokenCaller) GetDomainSeperator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ChildToken.contract.Call(opts, &out, "getDomainSeperator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDomainSeperator is a free data retrieval call binding the contract method 0x20379ee5.
//
// Solidity: function getDomainSeperator() view returns(bytes32)
func (_ChildToken *ChildTokenSession) GetDomainSeperator() ([32]byte, error) {
	return _ChildToken.Contract.GetDomainSeperator(&_ChildToken.CallOpts)
}

// GetDomainSeperator is a free data retrieval call binding the contract method 0x20379ee5.
//
// Solidity: function getDomainSeperator() view returns(bytes32)
func (_ChildToken *ChildTokenCallerSession) GetDomainSeperator() ([32]byte, error) {
	return _ChildToken.Contract.GetDomainSeperator(&_ChildToken.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address user) view returns(uint256 nonce)
func (_ChildToken *ChildTokenCaller) GetNonce(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ChildToken.contract.Call(opts, &out, "getNonce", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address user) view returns(uint256 nonce)
func (_ChildToken *ChildTokenSession) GetNonce(user common.Address) (*big.Int, error) {
	return _ChildToken.Contract.GetNonce(&_ChildToken.CallOpts, user)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address user) view returns(uint256 nonce)
func (_ChildToken *ChildTokenCallerSession) GetNonce(user common.Address) (*big.Int, error) {
	return _ChildToken.Contract.GetNonce(&_ChildToken.CallOpts, user)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ChildToken *ChildTokenCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ChildToken.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ChildToken *ChildTokenSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ChildToken.Contract.GetRoleAdmin(&_ChildToken.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ChildToken *ChildTokenCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ChildToken.Contract.GetRoleAdmin(&_ChildToken.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ChildToken *ChildTokenCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ChildToken.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ChildToken *ChildTokenSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _ChildToken.Contract.GetRoleMember(&_ChildToken.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ChildToken *ChildTokenCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _ChildToken.Contract.GetRoleMember(&_ChildToken.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ChildToken *ChildTokenCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ChildToken.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ChildToken *ChildTokenSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _ChildToken.Contract.GetRoleMemberCount(&_ChildToken.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ChildToken *ChildTokenCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _ChildToken.Contract.GetRoleMemberCount(&_ChildToken.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ChildToken *ChildTokenCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ChildToken.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ChildToken *ChildTokenSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ChildToken.Contract.HasRole(&_ChildToken.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ChildToken *ChildTokenCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ChildToken.Contract.HasRole(&_ChildToken.CallOpts, role, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ChildToken *ChildTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ChildToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ChildToken *ChildTokenSession) Name() (string, error) {
	return _ChildToken.Contract.Name(&_ChildToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ChildToken *ChildTokenCallerSession) Name() (string, error) {
	return _ChildToken.Contract.Name(&_ChildToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ChildToken *ChildTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ChildToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ChildToken *ChildTokenSession) Symbol() (string, error) {
	return _ChildToken.Contract.Symbol(&_ChildToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ChildToken *ChildTokenCallerSession) Symbol() (string, error) {
	return _ChildToken.Contract.Symbol(&_ChildToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ChildToken *ChildTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChildToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ChildToken *ChildTokenSession) TotalSupply() (*big.Int, error) {
	return _ChildToken.Contract.TotalSupply(&_ChildToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ChildToken *ChildTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _ChildToken.Contract.TotalSupply(&_ChildToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ChildToken *ChildTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChildToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ChildToken *ChildTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChildToken.Contract.Approve(&_ChildToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ChildToken *ChildTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChildToken.Contract.Approve(&_ChildToken.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ChildToken *ChildTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ChildToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ChildToken *ChildTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ChildToken.Contract.DecreaseAllowance(&_ChildToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ChildToken *ChildTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ChildToken.Contract.DecreaseAllowance(&_ChildToken.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0xcf2c52cb.
//
// Solidity: function deposit(address user, bytes depositData) returns()
func (_ChildToken *ChildTokenTransactor) Deposit(opts *bind.TransactOpts, user common.Address, depositData []byte) (*types.Transaction, error) {
	return _ChildToken.contract.Transact(opts, "deposit", user, depositData)
}

// Deposit is a paid mutator transaction binding the contract method 0xcf2c52cb.
//
// Solidity: function deposit(address user, bytes depositData) returns()
func (_ChildToken *ChildTokenSession) Deposit(user common.Address, depositData []byte) (*types.Transaction, error) {
	return _ChildToken.Contract.Deposit(&_ChildToken.TransactOpts, user, depositData)
}

// Deposit is a paid mutator transaction binding the contract method 0xcf2c52cb.
//
// Solidity: function deposit(address user, bytes depositData) returns()
func (_ChildToken *ChildTokenTransactorSession) Deposit(user common.Address, depositData []byte) (*types.Transaction, error) {
	return _ChildToken.Contract.Deposit(&_ChildToken.TransactOpts, user, depositData)
}

// ExecuteMetaTransaction is a paid mutator transaction binding the contract method 0x0c53c51c.
//
// Solidity: function executeMetaTransaction(address userAddress, bytes functionSignature, bytes32 sigR, bytes32 sigS, uint8 sigV) payable returns(bytes)
func (_ChildToken *ChildTokenTransactor) ExecuteMetaTransaction(opts *bind.TransactOpts, userAddress common.Address, functionSignature []byte, sigR [32]byte, sigS [32]byte, sigV uint8) (*types.Transaction, error) {
	return _ChildToken.contract.Transact(opts, "executeMetaTransaction", userAddress, functionSignature, sigR, sigS, sigV)
}

// ExecuteMetaTransaction is a paid mutator transaction binding the contract method 0x0c53c51c.
//
// Solidity: function executeMetaTransaction(address userAddress, bytes functionSignature, bytes32 sigR, bytes32 sigS, uint8 sigV) payable returns(bytes)
func (_ChildToken *ChildTokenSession) ExecuteMetaTransaction(userAddress common.Address, functionSignature []byte, sigR [32]byte, sigS [32]byte, sigV uint8) (*types.Transaction, error) {
	return _ChildToken.Contract.ExecuteMetaTransaction(&_ChildToken.TransactOpts, userAddress, functionSignature, sigR, sigS, sigV)
}

// ExecuteMetaTransaction is a paid mutator transaction binding the contract method 0x0c53c51c.
//
// Solidity: function executeMetaTransaction(address userAddress, bytes functionSignature, bytes32 sigR, bytes32 sigS, uint8 sigV) payable returns(bytes)
func (_ChildToken *ChildTokenTransactorSession) ExecuteMetaTransaction(userAddress common.Address, functionSignature []byte, sigR [32]byte, sigS [32]byte, sigV uint8) (*types.Transaction, error) {
	return _ChildToken.Contract.ExecuteMetaTransaction(&_ChildToken.TransactOpts, userAddress, functionSignature, sigR, sigS, sigV)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ChildToken *ChildTokenTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ChildToken.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ChildToken *ChildTokenSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ChildToken.Contract.GrantRole(&_ChildToken.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ChildToken *ChildTokenTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ChildToken.Contract.GrantRole(&_ChildToken.TransactOpts, role, account)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ChildToken *ChildTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ChildToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ChildToken *ChildTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ChildToken.Contract.IncreaseAllowance(&_ChildToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ChildToken *ChildTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ChildToken.Contract.IncreaseAllowance(&_ChildToken.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0xde7ea79d.
//
// Solidity: function initialize(string name_, string symbol_, uint8 decimals_, address childChainManager) returns()
func (_ChildToken *ChildTokenTransactor) Initialize(opts *bind.TransactOpts, name_ string, symbol_ string, decimals_ uint8, childChainManager common.Address) (*types.Transaction, error) {
	return _ChildToken.contract.Transact(opts, "initialize", name_, symbol_, decimals_, childChainManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xde7ea79d.
//
// Solidity: function initialize(string name_, string symbol_, uint8 decimals_, address childChainManager) returns()
func (_ChildToken *ChildTokenSession) Initialize(name_ string, symbol_ string, decimals_ uint8, childChainManager common.Address) (*types.Transaction, error) {
	return _ChildToken.Contract.Initialize(&_ChildToken.TransactOpts, name_, symbol_, decimals_, childChainManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xde7ea79d.
//
// Solidity: function initialize(string name_, string symbol_, uint8 decimals_, address childChainManager) returns()
func (_ChildToken *ChildTokenTransactorSession) Initialize(name_ string, symbol_ string, decimals_ uint8, childChainManager common.Address) (*types.Transaction, error) {
	return _ChildToken.Contract.Initialize(&_ChildToken.TransactOpts, name_, symbol_, decimals_, childChainManager)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ChildToken *ChildTokenTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ChildToken.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ChildToken *ChildTokenSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ChildToken.Contract.RenounceRole(&_ChildToken.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ChildToken *ChildTokenTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ChildToken.Contract.RenounceRole(&_ChildToken.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ChildToken *ChildTokenTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ChildToken.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ChildToken *ChildTokenSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ChildToken.Contract.RevokeRole(&_ChildToken.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ChildToken *ChildTokenTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ChildToken.Contract.RevokeRole(&_ChildToken.TransactOpts, role, account)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ChildToken *ChildTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChildToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ChildToken *ChildTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChildToken.Contract.Transfer(&_ChildToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ChildToken *ChildTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChildToken.Contract.Transfer(&_ChildToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ChildToken *ChildTokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChildToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ChildToken *ChildTokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChildToken.Contract.TransferFrom(&_ChildToken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ChildToken *ChildTokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChildToken.Contract.TransferFrom(&_ChildToken.TransactOpts, sender, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_ChildToken *ChildTokenTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ChildToken.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_ChildToken *ChildTokenSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _ChildToken.Contract.Withdraw(&_ChildToken.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_ChildToken *ChildTokenTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _ChildToken.Contract.Withdraw(&_ChildToken.TransactOpts, amount)
}

// ChildTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ChildToken contract.
type ChildTokenApprovalIterator struct {
	Event *ChildTokenApproval // Event containing the contract specifics and raw log

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
func (it *ChildTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChildTokenApproval)
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
		it.Event = new(ChildTokenApproval)
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
func (it *ChildTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChildTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChildTokenApproval represents a Approval event raised by the ChildToken contract.
type ChildTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ChildToken *ChildTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ChildTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ChildToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ChildTokenApprovalIterator{contract: _ChildToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ChildToken *ChildTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ChildTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ChildToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChildTokenApproval)
				if err := _ChildToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ChildToken *ChildTokenFilterer) ParseApproval(log types.Log) (*ChildTokenApproval, error) {
	event := new(ChildTokenApproval)
	if err := _ChildToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChildTokenMetaTransactionExecutedIterator is returned from FilterMetaTransactionExecuted and is used to iterate over the raw logs and unpacked data for MetaTransactionExecuted events raised by the ChildToken contract.
type ChildTokenMetaTransactionExecutedIterator struct {
	Event *ChildTokenMetaTransactionExecuted // Event containing the contract specifics and raw log

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
func (it *ChildTokenMetaTransactionExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChildTokenMetaTransactionExecuted)
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
		it.Event = new(ChildTokenMetaTransactionExecuted)
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
func (it *ChildTokenMetaTransactionExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChildTokenMetaTransactionExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChildTokenMetaTransactionExecuted represents a MetaTransactionExecuted event raised by the ChildToken contract.
type ChildTokenMetaTransactionExecuted struct {
	UserAddress       common.Address
	RelayerAddress    common.Address
	FunctionSignature []byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterMetaTransactionExecuted is a free log retrieval operation binding the contract event 0x5845892132946850460bff5a0083f71031bc5bf9aadcd40f1de79423eac9b10b.
//
// Solidity: event MetaTransactionExecuted(address userAddress, address relayerAddress, bytes functionSignature)
func (_ChildToken *ChildTokenFilterer) FilterMetaTransactionExecuted(opts *bind.FilterOpts) (*ChildTokenMetaTransactionExecutedIterator, error) {

	logs, sub, err := _ChildToken.contract.FilterLogs(opts, "MetaTransactionExecuted")
	if err != nil {
		return nil, err
	}
	return &ChildTokenMetaTransactionExecutedIterator{contract: _ChildToken.contract, event: "MetaTransactionExecuted", logs: logs, sub: sub}, nil
}

// WatchMetaTransactionExecuted is a free log subscription operation binding the contract event 0x5845892132946850460bff5a0083f71031bc5bf9aadcd40f1de79423eac9b10b.
//
// Solidity: event MetaTransactionExecuted(address userAddress, address relayerAddress, bytes functionSignature)
func (_ChildToken *ChildTokenFilterer) WatchMetaTransactionExecuted(opts *bind.WatchOpts, sink chan<- *ChildTokenMetaTransactionExecuted) (event.Subscription, error) {

	logs, sub, err := _ChildToken.contract.WatchLogs(opts, "MetaTransactionExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChildTokenMetaTransactionExecuted)
				if err := _ChildToken.contract.UnpackLog(event, "MetaTransactionExecuted", log); err != nil {
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

// ParseMetaTransactionExecuted is a log parse operation binding the contract event 0x5845892132946850460bff5a0083f71031bc5bf9aadcd40f1de79423eac9b10b.
//
// Solidity: event MetaTransactionExecuted(address userAddress, address relayerAddress, bytes functionSignature)
func (_ChildToken *ChildTokenFilterer) ParseMetaTransactionExecuted(log types.Log) (*ChildTokenMetaTransactionExecuted, error) {
	event := new(ChildTokenMetaTransactionExecuted)
	if err := _ChildToken.contract.UnpackLog(event, "MetaTransactionExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChildTokenRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ChildToken contract.
type ChildTokenRoleAdminChangedIterator struct {
	Event *ChildTokenRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ChildTokenRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChildTokenRoleAdminChanged)
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
		it.Event = new(ChildTokenRoleAdminChanged)
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
func (it *ChildTokenRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChildTokenRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChildTokenRoleAdminChanged represents a RoleAdminChanged event raised by the ChildToken contract.
type ChildTokenRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ChildToken *ChildTokenFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ChildTokenRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ChildToken.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ChildTokenRoleAdminChangedIterator{contract: _ChildToken.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ChildToken *ChildTokenFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ChildTokenRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ChildToken.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChildTokenRoleAdminChanged)
				if err := _ChildToken.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ChildToken *ChildTokenFilterer) ParseRoleAdminChanged(log types.Log) (*ChildTokenRoleAdminChanged, error) {
	event := new(ChildTokenRoleAdminChanged)
	if err := _ChildToken.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChildTokenRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ChildToken contract.
type ChildTokenRoleGrantedIterator struct {
	Event *ChildTokenRoleGranted // Event containing the contract specifics and raw log

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
func (it *ChildTokenRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChildTokenRoleGranted)
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
		it.Event = new(ChildTokenRoleGranted)
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
func (it *ChildTokenRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChildTokenRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChildTokenRoleGranted represents a RoleGranted event raised by the ChildToken contract.
type ChildTokenRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ChildToken *ChildTokenFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ChildTokenRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ChildToken.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ChildTokenRoleGrantedIterator{contract: _ChildToken.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ChildToken *ChildTokenFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ChildTokenRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ChildToken.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChildTokenRoleGranted)
				if err := _ChildToken.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ChildToken *ChildTokenFilterer) ParseRoleGranted(log types.Log) (*ChildTokenRoleGranted, error) {
	event := new(ChildTokenRoleGranted)
	if err := _ChildToken.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChildTokenRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ChildToken contract.
type ChildTokenRoleRevokedIterator struct {
	Event *ChildTokenRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ChildTokenRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChildTokenRoleRevoked)
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
		it.Event = new(ChildTokenRoleRevoked)
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
func (it *ChildTokenRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChildTokenRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChildTokenRoleRevoked represents a RoleRevoked event raised by the ChildToken contract.
type ChildTokenRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ChildToken *ChildTokenFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ChildTokenRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ChildToken.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ChildTokenRoleRevokedIterator{contract: _ChildToken.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ChildToken *ChildTokenFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ChildTokenRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ChildToken.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChildTokenRoleRevoked)
				if err := _ChildToken.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ChildToken *ChildTokenFilterer) ParseRoleRevoked(log types.Log) (*ChildTokenRoleRevoked, error) {
	event := new(ChildTokenRoleRevoked)
	if err := _ChildToken.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChildTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ChildToken contract.
type ChildTokenTransferIterator struct {
	Event *ChildTokenTransfer // Event containing the contract specifics and raw log

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
func (it *ChildTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChildTokenTransfer)
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
		it.Event = new(ChildTokenTransfer)
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
func (it *ChildTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChildTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChildTokenTransfer represents a Transfer event raised by the ChildToken contract.
type ChildTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ChildToken *ChildTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ChildTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ChildToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ChildTokenTransferIterator{contract: _ChildToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ChildToken *ChildTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ChildTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ChildToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChildTokenTransfer)
				if err := _ChildToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ChildToken *ChildTokenFilterer) ParseTransfer(log types.Log) (*ChildTokenTransfer, error) {
	event := new(ChildTokenTransfer)
	if err := _ChildToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
