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

// RootChainABI is the input ABI used to generate the binding from.
const RootChainABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"headerBlockId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"NewHeaderBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"headerBlockId\",\"type\":\"uint256\"}],\"name\":\"ResetHeaderBlock\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"CHAINID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"VOTE_TYPE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_nextHeaderBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentHeaderBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLastChildBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"headerBlocks\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"heimdallId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"networkId\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_heimdallId\",\"type\":\"string\"}],\"name\":\"setHeimdallId\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setNextHeaderBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"slash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256[3][]\",\"name\":\"sigs\",\"type\":\"uint256[3][]\"}],\"name\":\"submitCheckpoint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sigs\",\"type\":\"bytes\"}],\"name\":\"submitHeaderBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numDeposits\",\"type\":\"uint256\"}],\"name\":\"updateDepositId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RootChain is an auto generated Go binding around an Ethereum contract.
type RootChain struct {
	RootChainCaller     // Read-only binding to the contract
	RootChainTransactor // Write-only binding to the contract
	RootChainFilterer   // Log filterer for contract events
}

// RootChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type RootChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RootChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RootChainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RootChainSession struct {
	Contract     *RootChain        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RootChainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RootChainCallerSession struct {
	Contract *RootChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RootChainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RootChainTransactorSession struct {
	Contract     *RootChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RootChainRaw is an auto generated low-level Go binding around an Ethereum contract.
type RootChainRaw struct {
	Contract *RootChain // Generic contract binding to access the raw methods on
}

// RootChainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RootChainCallerRaw struct {
	Contract *RootChainCaller // Generic read-only contract binding to access the raw methods on
}

// RootChainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RootChainTransactorRaw struct {
	Contract *RootChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRootChain creates a new instance of RootChain, bound to a specific deployed contract.
func NewRootChain(address common.Address, backend bind.ContractBackend) (*RootChain, error) {
	contract, err := bindRootChain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RootChain{RootChainCaller: RootChainCaller{contract: contract}, RootChainTransactor: RootChainTransactor{contract: contract}, RootChainFilterer: RootChainFilterer{contract: contract}}, nil
}

// NewRootChainCaller creates a new read-only instance of RootChain, bound to a specific deployed contract.
func NewRootChainCaller(address common.Address, caller bind.ContractCaller) (*RootChainCaller, error) {
	contract, err := bindRootChain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RootChainCaller{contract: contract}, nil
}

// NewRootChainTransactor creates a new write-only instance of RootChain, bound to a specific deployed contract.
func NewRootChainTransactor(address common.Address, transactor bind.ContractTransactor) (*RootChainTransactor, error) {
	contract, err := bindRootChain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RootChainTransactor{contract: contract}, nil
}

// NewRootChainFilterer creates a new log filterer instance of RootChain, bound to a specific deployed contract.
func NewRootChainFilterer(address common.Address, filterer bind.ContractFilterer) (*RootChainFilterer, error) {
	contract, err := bindRootChain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RootChainFilterer{contract: contract}, nil
}

// bindRootChain binds a generic wrapper to an already deployed contract.
func bindRootChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RootChainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootChain *RootChainRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RootChain.Contract.RootChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootChain *RootChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.Contract.RootChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootChain *RootChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootChain.Contract.RootChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootChain *RootChainCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RootChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootChain *RootChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootChain *RootChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootChain.Contract.contract.Transact(opts, method, params...)
}

// CHAINID is a free data retrieval call binding the contract method 0xcc79f97b.
//
// Solidity: function CHAINID() view returns(uint256)
func (_RootChain *RootChainCaller) CHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "CHAINID")
	return *ret0, err
}

// CHAINID is a free data retrieval call binding the contract method 0xcc79f97b.
//
// Solidity: function CHAINID() view returns(uint256)
func (_RootChain *RootChainSession) CHAINID() (*big.Int, error) {
	return _RootChain.Contract.CHAINID(&_RootChain.CallOpts)
}

// CHAINID is a free data retrieval call binding the contract method 0xcc79f97b.
//
// Solidity: function CHAINID() view returns(uint256)
func (_RootChain *RootChainCallerSession) CHAINID() (*big.Int, error) {
	return _RootChain.Contract.CHAINID(&_RootChain.CallOpts)
}

// VOTETYPE is a free data retrieval call binding the contract method 0xd5b844eb.
//
// Solidity: function VOTE_TYPE() view returns(uint8)
func (_RootChain *RootChainCaller) VOTETYPE(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "VOTE_TYPE")
	return *ret0, err
}

// VOTETYPE is a free data retrieval call binding the contract method 0xd5b844eb.
//
// Solidity: function VOTE_TYPE() view returns(uint8)
func (_RootChain *RootChainSession) VOTETYPE() (uint8, error) {
	return _RootChain.Contract.VOTETYPE(&_RootChain.CallOpts)
}

// VOTETYPE is a free data retrieval call binding the contract method 0xd5b844eb.
//
// Solidity: function VOTE_TYPE() view returns(uint8)
func (_RootChain *RootChainCallerSession) VOTETYPE() (uint8, error) {
	return _RootChain.Contract.VOTETYPE(&_RootChain.CallOpts)
}

// NextHeaderBlock is a free data retrieval call binding the contract method 0x8d978d88.
//
// Solidity: function _nextHeaderBlock() view returns(uint256)
func (_RootChain *RootChainCaller) NextHeaderBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "_nextHeaderBlock")
	return *ret0, err
}

// NextHeaderBlock is a free data retrieval call binding the contract method 0x8d978d88.
//
// Solidity: function _nextHeaderBlock() view returns(uint256)
func (_RootChain *RootChainSession) NextHeaderBlock() (*big.Int, error) {
	return _RootChain.Contract.NextHeaderBlock(&_RootChain.CallOpts)
}

// NextHeaderBlock is a free data retrieval call binding the contract method 0x8d978d88.
//
// Solidity: function _nextHeaderBlock() view returns(uint256)
func (_RootChain *RootChainCallerSession) NextHeaderBlock() (*big.Int, error) {
	return _RootChain.Contract.NextHeaderBlock(&_RootChain.CallOpts)
}

// CurrentHeaderBlock is a free data retrieval call binding the contract method 0xec7e4855.
//
// Solidity: function currentHeaderBlock() view returns(uint256)
func (_RootChain *RootChainCaller) CurrentHeaderBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "currentHeaderBlock")
	return *ret0, err
}

// CurrentHeaderBlock is a free data retrieval call binding the contract method 0xec7e4855.
//
// Solidity: function currentHeaderBlock() view returns(uint256)
func (_RootChain *RootChainSession) CurrentHeaderBlock() (*big.Int, error) {
	return _RootChain.Contract.CurrentHeaderBlock(&_RootChain.CallOpts)
}

// CurrentHeaderBlock is a free data retrieval call binding the contract method 0xec7e4855.
//
// Solidity: function currentHeaderBlock() view returns(uint256)
func (_RootChain *RootChainCallerSession) CurrentHeaderBlock() (*big.Int, error) {
	return _RootChain.Contract.CurrentHeaderBlock(&_RootChain.CallOpts)
}

// GetLastChildBlock is a free data retrieval call binding the contract method 0xb87e1b66.
//
// Solidity: function getLastChildBlock() view returns(uint256)
func (_RootChain *RootChainCaller) GetLastChildBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "getLastChildBlock")
	return *ret0, err
}

// GetLastChildBlock is a free data retrieval call binding the contract method 0xb87e1b66.
//
// Solidity: function getLastChildBlock() view returns(uint256)
func (_RootChain *RootChainSession) GetLastChildBlock() (*big.Int, error) {
	return _RootChain.Contract.GetLastChildBlock(&_RootChain.CallOpts)
}

// GetLastChildBlock is a free data retrieval call binding the contract method 0xb87e1b66.
//
// Solidity: function getLastChildBlock() view returns(uint256)
func (_RootChain *RootChainCallerSession) GetLastChildBlock() (*big.Int, error) {
	return _RootChain.Contract.GetLastChildBlock(&_RootChain.CallOpts)
}

// HeaderBlocks is a free data retrieval call binding the contract method 0x41539d4a.
//
// Solidity: function headerBlocks(uint256 ) view returns(bytes32 root, uint256 start, uint256 end, uint256 createdAt, address proposer)
func (_RootChain *RootChainCaller) HeaderBlocks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Root      [32]byte
	Start     *big.Int
	End       *big.Int
	CreatedAt *big.Int
	Proposer  common.Address
}, error) {
	ret := new(struct {
		Root      [32]byte
		Start     *big.Int
		End       *big.Int
		CreatedAt *big.Int
		Proposer  common.Address
	})
	out := ret
	err := _RootChain.contract.Call(opts, out, "headerBlocks", arg0)
	return *ret, err
}

// HeaderBlocks is a free data retrieval call binding the contract method 0x41539d4a.
//
// Solidity: function headerBlocks(uint256 ) view returns(bytes32 root, uint256 start, uint256 end, uint256 createdAt, address proposer)
func (_RootChain *RootChainSession) HeaderBlocks(arg0 *big.Int) (struct {
	Root      [32]byte
	Start     *big.Int
	End       *big.Int
	CreatedAt *big.Int
	Proposer  common.Address
}, error) {
	return _RootChain.Contract.HeaderBlocks(&_RootChain.CallOpts, arg0)
}

// HeaderBlocks is a free data retrieval call binding the contract method 0x41539d4a.
//
// Solidity: function headerBlocks(uint256 ) view returns(bytes32 root, uint256 start, uint256 end, uint256 createdAt, address proposer)
func (_RootChain *RootChainCallerSession) HeaderBlocks(arg0 *big.Int) (struct {
	Root      [32]byte
	Start     *big.Int
	End       *big.Int
	CreatedAt *big.Int
	Proposer  common.Address
}, error) {
	return _RootChain.Contract.HeaderBlocks(&_RootChain.CallOpts, arg0)
}

// HeimdallId is a free data retrieval call binding the contract method 0xfbc3dd36.
//
// Solidity: function heimdallId() view returns(bytes32)
func (_RootChain *RootChainCaller) HeimdallId(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "heimdallId")
	return *ret0, err
}

// HeimdallId is a free data retrieval call binding the contract method 0xfbc3dd36.
//
// Solidity: function heimdallId() view returns(bytes32)
func (_RootChain *RootChainSession) HeimdallId() ([32]byte, error) {
	return _RootChain.Contract.HeimdallId(&_RootChain.CallOpts)
}

// HeimdallId is a free data retrieval call binding the contract method 0xfbc3dd36.
//
// Solidity: function heimdallId() view returns(bytes32)
func (_RootChain *RootChainCallerSession) HeimdallId() ([32]byte, error) {
	return _RootChain.Contract.HeimdallId(&_RootChain.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_RootChain *RootChainCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_RootChain *RootChainSession) IsOwner() (bool, error) {
	return _RootChain.Contract.IsOwner(&_RootChain.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_RootChain *RootChainCallerSession) IsOwner() (bool, error) {
	return _RootChain.Contract.IsOwner(&_RootChain.CallOpts)
}

// NetworkId is a free data retrieval call binding the contract method 0x9025e64c.
//
// Solidity: function networkId() view returns(bytes)
func (_RootChain *RootChainCaller) NetworkId(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "networkId")
	return *ret0, err
}

// NetworkId is a free data retrieval call binding the contract method 0x9025e64c.
//
// Solidity: function networkId() view returns(bytes)
func (_RootChain *RootChainSession) NetworkId() ([]byte, error) {
	return _RootChain.Contract.NetworkId(&_RootChain.CallOpts)
}

// NetworkId is a free data retrieval call binding the contract method 0x9025e64c.
//
// Solidity: function networkId() view returns(bytes)
func (_RootChain *RootChainCallerSession) NetworkId() ([]byte, error) {
	return _RootChain.Contract.NetworkId(&_RootChain.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RootChain *RootChainCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RootChain *RootChainSession) Owner() (common.Address, error) {
	return _RootChain.Contract.Owner(&_RootChain.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RootChain *RootChainCallerSession) Owner() (common.Address, error) {
	return _RootChain.Contract.Owner(&_RootChain.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RootChain *RootChainTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RootChain *RootChainSession) RenounceOwnership() (*types.Transaction, error) {
	return _RootChain.Contract.RenounceOwnership(&_RootChain.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RootChain *RootChainTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RootChain.Contract.RenounceOwnership(&_RootChain.TransactOpts)
}

// SetHeimdallId is a paid mutator transaction binding the contract method 0xea0688b3.
//
// Solidity: function setHeimdallId(string _heimdallId) returns()
func (_RootChain *RootChainTransactor) SetHeimdallId(opts *bind.TransactOpts, _heimdallId string) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "setHeimdallId", _heimdallId)
}

// SetHeimdallId is a paid mutator transaction binding the contract method 0xea0688b3.
//
// Solidity: function setHeimdallId(string _heimdallId) returns()
func (_RootChain *RootChainSession) SetHeimdallId(_heimdallId string) (*types.Transaction, error) {
	return _RootChain.Contract.SetHeimdallId(&_RootChain.TransactOpts, _heimdallId)
}

// SetHeimdallId is a paid mutator transaction binding the contract method 0xea0688b3.
//
// Solidity: function setHeimdallId(string _heimdallId) returns()
func (_RootChain *RootChainTransactorSession) SetHeimdallId(_heimdallId string) (*types.Transaction, error) {
	return _RootChain.Contract.SetHeimdallId(&_RootChain.TransactOpts, _heimdallId)
}

// SetNextHeaderBlock is a paid mutator transaction binding the contract method 0xcf24a0ea.
//
// Solidity: function setNextHeaderBlock(uint256 _value) returns()
func (_RootChain *RootChainTransactor) SetNextHeaderBlock(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "setNextHeaderBlock", _value)
}

// SetNextHeaderBlock is a paid mutator transaction binding the contract method 0xcf24a0ea.
//
// Solidity: function setNextHeaderBlock(uint256 _value) returns()
func (_RootChain *RootChainSession) SetNextHeaderBlock(_value *big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.SetNextHeaderBlock(&_RootChain.TransactOpts, _value)
}

// SetNextHeaderBlock is a paid mutator transaction binding the contract method 0xcf24a0ea.
//
// Solidity: function setNextHeaderBlock(uint256 _value) returns()
func (_RootChain *RootChainTransactorSession) SetNextHeaderBlock(_value *big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.SetNextHeaderBlock(&_RootChain.TransactOpts, _value)
}

// Slash is a paid mutator transaction binding the contract method 0x2da25de3.
//
// Solidity: function slash() returns()
func (_RootChain *RootChainTransactor) Slash(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "slash")
}

// Slash is a paid mutator transaction binding the contract method 0x2da25de3.
//
// Solidity: function slash() returns()
func (_RootChain *RootChainSession) Slash() (*types.Transaction, error) {
	return _RootChain.Contract.Slash(&_RootChain.TransactOpts)
}

// Slash is a paid mutator transaction binding the contract method 0x2da25de3.
//
// Solidity: function slash() returns()
func (_RootChain *RootChainTransactorSession) Slash() (*types.Transaction, error) {
	return _RootChain.Contract.Slash(&_RootChain.TransactOpts)
}

// SubmitCheckpoint is a paid mutator transaction binding the contract method 0x4e43e495.
//
// Solidity: function submitCheckpoint(bytes data, uint256[3][] sigs) returns()
func (_RootChain *RootChainTransactor) SubmitCheckpoint(opts *bind.TransactOpts, data []byte, sigs [][3]*big.Int) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "submitCheckpoint", data, sigs)
}

// SubmitCheckpoint is a paid mutator transaction binding the contract method 0x4e43e495.
//
// Solidity: function submitCheckpoint(bytes data, uint256[3][] sigs) returns()
func (_RootChain *RootChainSession) SubmitCheckpoint(data []byte, sigs [][3]*big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.SubmitCheckpoint(&_RootChain.TransactOpts, data, sigs)
}

// SubmitCheckpoint is a paid mutator transaction binding the contract method 0x4e43e495.
//
// Solidity: function submitCheckpoint(bytes data, uint256[3][] sigs) returns()
func (_RootChain *RootChainTransactorSession) SubmitCheckpoint(data []byte, sigs [][3]*big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.SubmitCheckpoint(&_RootChain.TransactOpts, data, sigs)
}

// SubmitHeaderBlock is a paid mutator transaction binding the contract method 0x6a791f11.
//
// Solidity: function submitHeaderBlock(bytes data, bytes sigs) returns()
func (_RootChain *RootChainTransactor) SubmitHeaderBlock(opts *bind.TransactOpts, data []byte, sigs []byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "submitHeaderBlock", data, sigs)
}

// SubmitHeaderBlock is a paid mutator transaction binding the contract method 0x6a791f11.
//
// Solidity: function submitHeaderBlock(bytes data, bytes sigs) returns()
func (_RootChain *RootChainSession) SubmitHeaderBlock(data []byte, sigs []byte) (*types.Transaction, error) {
	return _RootChain.Contract.SubmitHeaderBlock(&_RootChain.TransactOpts, data, sigs)
}

// SubmitHeaderBlock is a paid mutator transaction binding the contract method 0x6a791f11.
//
// Solidity: function submitHeaderBlock(bytes data, bytes sigs) returns()
func (_RootChain *RootChainTransactorSession) SubmitHeaderBlock(data []byte, sigs []byte) (*types.Transaction, error) {
	return _RootChain.Contract.SubmitHeaderBlock(&_RootChain.TransactOpts, data, sigs)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RootChain *RootChainTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RootChain *RootChainSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RootChain.Contract.TransferOwnership(&_RootChain.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RootChain *RootChainTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RootChain.Contract.TransferOwnership(&_RootChain.TransactOpts, newOwner)
}

// UpdateDepositId is a paid mutator transaction binding the contract method 0x5391f483.
//
// Solidity: function updateDepositId(uint256 numDeposits) returns(uint256 depositId)
func (_RootChain *RootChainTransactor) UpdateDepositId(opts *bind.TransactOpts, numDeposits *big.Int) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "updateDepositId", numDeposits)
}

// UpdateDepositId is a paid mutator transaction binding the contract method 0x5391f483.
//
// Solidity: function updateDepositId(uint256 numDeposits) returns(uint256 depositId)
func (_RootChain *RootChainSession) UpdateDepositId(numDeposits *big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.UpdateDepositId(&_RootChain.TransactOpts, numDeposits)
}

// UpdateDepositId is a paid mutator transaction binding the contract method 0x5391f483.
//
// Solidity: function updateDepositId(uint256 numDeposits) returns(uint256 depositId)
func (_RootChain *RootChainTransactorSession) UpdateDepositId(numDeposits *big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.UpdateDepositId(&_RootChain.TransactOpts, numDeposits)
}

// RootChainNewHeaderBlockIterator is returned from FilterNewHeaderBlock and is used to iterate over the raw logs and unpacked data for NewHeaderBlock events raised by the RootChain contract.
type RootChainNewHeaderBlockIterator struct {
	Event *RootChainNewHeaderBlock // Event containing the contract specifics and raw log

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
func (it *RootChainNewHeaderBlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainNewHeaderBlock)
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
		it.Event = new(RootChainNewHeaderBlock)
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
func (it *RootChainNewHeaderBlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainNewHeaderBlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainNewHeaderBlock represents a NewHeaderBlock event raised by the RootChain contract.
type RootChainNewHeaderBlock struct {
	Proposer      common.Address
	HeaderBlockId *big.Int
	Reward        *big.Int
	Start         *big.Int
	End           *big.Int
	Root          [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewHeaderBlock is a free log retrieval operation binding the contract event 0xba5de06d22af2685c6c7765f60067f7d2b08c2d29f53cdf14d67f6d1c9bfb527.
//
// Solidity: event NewHeaderBlock(address indexed proposer, uint256 indexed headerBlockId, uint256 indexed reward, uint256 start, uint256 end, bytes32 root)
func (_RootChain *RootChainFilterer) FilterNewHeaderBlock(opts *bind.FilterOpts, proposer []common.Address, headerBlockId []*big.Int, reward []*big.Int) (*RootChainNewHeaderBlockIterator, error) {

	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}
	var headerBlockIdRule []interface{}
	for _, headerBlockIdItem := range headerBlockId {
		headerBlockIdRule = append(headerBlockIdRule, headerBlockIdItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "NewHeaderBlock", proposerRule, headerBlockIdRule, rewardRule)
	if err != nil {
		return nil, err
	}
	return &RootChainNewHeaderBlockIterator{contract: _RootChain.contract, event: "NewHeaderBlock", logs: logs, sub: sub}, nil
}

// WatchNewHeaderBlock is a free log subscription operation binding the contract event 0xba5de06d22af2685c6c7765f60067f7d2b08c2d29f53cdf14d67f6d1c9bfb527.
//
// Solidity: event NewHeaderBlock(address indexed proposer, uint256 indexed headerBlockId, uint256 indexed reward, uint256 start, uint256 end, bytes32 root)
func (_RootChain *RootChainFilterer) WatchNewHeaderBlock(opts *bind.WatchOpts, sink chan<- *RootChainNewHeaderBlock, proposer []common.Address, headerBlockId []*big.Int, reward []*big.Int) (event.Subscription, error) {

	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}
	var headerBlockIdRule []interface{}
	for _, headerBlockIdItem := range headerBlockId {
		headerBlockIdRule = append(headerBlockIdRule, headerBlockIdItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "NewHeaderBlock", proposerRule, headerBlockIdRule, rewardRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainNewHeaderBlock)
				if err := _RootChain.contract.UnpackLog(event, "NewHeaderBlock", log); err != nil {
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

// ParseNewHeaderBlock is a log parse operation binding the contract event 0xba5de06d22af2685c6c7765f60067f7d2b08c2d29f53cdf14d67f6d1c9bfb527.
//
// Solidity: event NewHeaderBlock(address indexed proposer, uint256 indexed headerBlockId, uint256 indexed reward, uint256 start, uint256 end, bytes32 root)
func (_RootChain *RootChainFilterer) ParseNewHeaderBlock(log types.Log) (*RootChainNewHeaderBlock, error) {
	event := new(RootChainNewHeaderBlock)
	if err := _RootChain.contract.UnpackLog(event, "NewHeaderBlock", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RootChainOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RootChain contract.
type RootChainOwnershipTransferredIterator struct {
	Event *RootChainOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RootChainOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainOwnershipTransferred)
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
		it.Event = new(RootChainOwnershipTransferred)
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
func (it *RootChainOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainOwnershipTransferred represents a OwnershipTransferred event raised by the RootChain contract.
type RootChainOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RootChain *RootChainFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RootChainOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RootChainOwnershipTransferredIterator{contract: _RootChain.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RootChain *RootChainFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RootChainOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainOwnershipTransferred)
				if err := _RootChain.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RootChain *RootChainFilterer) ParseOwnershipTransferred(log types.Log) (*RootChainOwnershipTransferred, error) {
	event := new(RootChainOwnershipTransferred)
	if err := _RootChain.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RootChainResetHeaderBlockIterator is returned from FilterResetHeaderBlock and is used to iterate over the raw logs and unpacked data for ResetHeaderBlock events raised by the RootChain contract.
type RootChainResetHeaderBlockIterator struct {
	Event *RootChainResetHeaderBlock // Event containing the contract specifics and raw log

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
func (it *RootChainResetHeaderBlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainResetHeaderBlock)
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
		it.Event = new(RootChainResetHeaderBlock)
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
func (it *RootChainResetHeaderBlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainResetHeaderBlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainResetHeaderBlock represents a ResetHeaderBlock event raised by the RootChain contract.
type RootChainResetHeaderBlock struct {
	Proposer      common.Address
	HeaderBlockId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterResetHeaderBlock is a free log retrieval operation binding the contract event 0xca1d8316287f938830e225956a7bb10fd5a1a1506dd2eb3a476751a488117205.
//
// Solidity: event ResetHeaderBlock(address indexed proposer, uint256 indexed headerBlockId)
func (_RootChain *RootChainFilterer) FilterResetHeaderBlock(opts *bind.FilterOpts, proposer []common.Address, headerBlockId []*big.Int) (*RootChainResetHeaderBlockIterator, error) {

	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}
	var headerBlockIdRule []interface{}
	for _, headerBlockIdItem := range headerBlockId {
		headerBlockIdRule = append(headerBlockIdRule, headerBlockIdItem)
	}

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "ResetHeaderBlock", proposerRule, headerBlockIdRule)
	if err != nil {
		return nil, err
	}
	return &RootChainResetHeaderBlockIterator{contract: _RootChain.contract, event: "ResetHeaderBlock", logs: logs, sub: sub}, nil
}

// WatchResetHeaderBlock is a free log subscription operation binding the contract event 0xca1d8316287f938830e225956a7bb10fd5a1a1506dd2eb3a476751a488117205.
//
// Solidity: event ResetHeaderBlock(address indexed proposer, uint256 indexed headerBlockId)
func (_RootChain *RootChainFilterer) WatchResetHeaderBlock(opts *bind.WatchOpts, sink chan<- *RootChainResetHeaderBlock, proposer []common.Address, headerBlockId []*big.Int) (event.Subscription, error) {

	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}
	var headerBlockIdRule []interface{}
	for _, headerBlockIdItem := range headerBlockId {
		headerBlockIdRule = append(headerBlockIdRule, headerBlockIdItem)
	}

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "ResetHeaderBlock", proposerRule, headerBlockIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainResetHeaderBlock)
				if err := _RootChain.contract.UnpackLog(event, "ResetHeaderBlock", log); err != nil {
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

// ParseResetHeaderBlock is a log parse operation binding the contract event 0xca1d8316287f938830e225956a7bb10fd5a1a1506dd2eb3a476751a488117205.
//
// Solidity: event ResetHeaderBlock(address indexed proposer, uint256 indexed headerBlockId)
func (_RootChain *RootChainFilterer) ParseResetHeaderBlock(log types.Log) (*RootChainResetHeaderBlock, error) {
	event := new(RootChainResetHeaderBlock)
	if err := _RootChain.contract.UnpackLog(event, "ResetHeaderBlock", log); err != nil {
		return nil, err
	}
	return event, nil
}
