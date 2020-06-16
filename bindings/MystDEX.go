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

// MystDEXABI is the input ABI used to generate the binding from.
const MystDEXABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousDestination\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newDestination\",\"type\":\"address\"}],\"name\":\"DestinationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"claimEthers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"claimTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFundsDestination\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialised\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_newDestination\",\"type\":\"address\"}],\"name\":\"setFundsDestination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dexOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newRate\",\"type\":\"uint256\"}],\"name\":\"setRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferEthers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferMyst\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MystDEXBin is the compiled bytecode used for deploying new contracts.
var MystDEXBin = "0x608060405234801561001057600080fd5b50610e79806100206000396000f3fe6080604052600436106100c65760003560e01c8063715018a61161007f578063df8de3e711610059578063df8de3e7146103ff578063f2fde38b14610432578063f58c5b6e14610465578063fc0c546a1461047a57610262565b8063715018a6146103765780638595d1491461038b5780638da5cb5b146103ce57610262565b806307003bb4146102675780631254e64d14610290578063238e130a146102cb57806334fcf437146102fe57806338d2e411146103285780636931b5501461036157610262565b3661026257600254600160a01b900460ff16610129576040805162461bcd60e51b815260206004820152601b60248201527f436f6e7472616374206973206e6f7420696e697469616c697365640000000000604482015290519081900360640190fd5b6000610158670de0b6b3a764000061014c6003543461048f90919063ffffffff16565b9063ffffffff6104da16565b600254604080516370a0823160e01b8152306004820152905192935083926001600160a01b03909216916370a0823191602480820192602092909190829003018186803b1580156101a857600080fd5b505afa1580156101bc573d6000803e3d6000fd5b505050506040513d60208110156101d257600080fd5b505110156101df57600080fd5b6002546040805163a9059cbb60e01b81523360048201526024810184905290516001600160a01b039092169163a9059cbb916044808201926020929091908290030181600087803b15801561023357600080fd5b505af1158015610247573d6000803e3d6000fd5b505050506040513d602081101561025d57600080fd5b505050005b600080fd5b34801561027357600080fd5b5061027c610533565b604080519115158252519081900360200190f35b34801561029c57600080fd5b506102c9600480360360408110156102b357600080fd5b506001600160a01b038135169060200135610543565b005b3480156102d757600080fd5b506102c9600480360360208110156102ee57600080fd5b50356001600160a01b0316610634565b34801561030a57600080fd5b506102c96004803603602081101561032157600080fd5b5035610704565b34801561033457600080fd5b506102c96004803603604081101561034b57600080fd5b506001600160a01b03813516906020013561076a565b34801561036d57600080fd5b506102c961091a565b34801561038257600080fd5b506102c961096b565b34801561039757600080fd5b506102c9600480360360608110156103ae57600080fd5b506001600160a01b03813581169160208101359091169060400135610a16565b3480156103da57600080fd5b506103e3610ab6565b604080516001600160a01b039092168252519081900360200190f35b34801561040b57600080fd5b506102c96004803603602081101561042257600080fd5b50356001600160a01b0316610ac5565b34801561043e57600080fd5b506102c96004803603602081101561045557600080fd5b50356001600160a01b0316610bf6565b34801561047157600080fd5b506103e3610cf7565b34801561048657600080fd5b506103e3610d06565b60006104d183836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f000000000000815250610d15565b90505b92915050565b6000826104e9575060006104d4565b828202828482816104f657fe5b04146104d15760405162461bcd60e51b8152600401808060200182810382526021815260200180610e036021913960400191505060405180910390fd5b600254600160a01b900460ff1681565b6000546001600160a01b031633148061056557506000546001600160a01b0316155b6105a4576040805162461bcd60e51b81526020600482018190526024820152600080516020610e24833981519152604482015290519081900360640190fd5b804710156105f9576040805162461bcd60e51b815260206004820152601860248201527f6e6f7420656e6f7567682065746865722062616c616e63650000000000000000604482015290519081900360640190fd5b6040516001600160a01b0383169082156108fc029083906000818181858888f1935050505015801561062f573d6000803e3d6000fd5b505050565b6000546001600160a01b031633148061065657506000546001600160a01b0316155b610695576040805162461bcd60e51b81526020600482018190526024820152600080516020610e24833981519152604482015290519081900360640190fd5b6001600160a01b0381166106a857600080fd5b6001546040516001600160a01b038084169216907fe1a66d77649cf0a57b9937073549f30f1c82bb865aaf066d2f299e37a62c6aad90600090a3600180546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b031633148061072657506000546001600160a01b0316155b610765576040805162461bcd60e51b81526020600482018190526024820152600080516020610e24833981519152604482015290519081900360640190fd5b600355565b6000546001600160a01b031633148061078c57506000546001600160a01b0316155b6107cb576040805162461bcd60e51b81526020600482018190526024820152600080516020610e24833981519152604482015290519081900360640190fd5b600254604080516370a0823160e01b8152306004820152905183926001600160a01b0316916370a08231916024808301926020929190829003018186803b15801561081557600080fd5b505afa158015610829573d6000803e3d6000fd5b505050506040513d602081101561083f57600080fd5b50511015610894576040805162461bcd60e51b815260206004820152601760248201527f6e6f7420656e6f756768206d7973742062616c616e6365000000000000000000604482015290519081900360640190fd5b6002546040805163a9059cbb60e01b81526001600160a01b038581166004830152602482018590529151919092169163a9059cbb9160448083019260209291908290030181600087803b1580156108ea57600080fd5b505af11580156108fe573d6000803e3d6000fd5b505050506040513d602081101561091457600080fd5b50505050565b6001546001600160a01b031661092f57600080fd5b6001546040516001600160a01b03909116904780156108fc02916000818181858888f19350505050158015610968573d6000803e3d6000fd5b50565b6000546001600160a01b031633148061098d57506000546001600160a01b0316155b6109cc576040805162461bcd60e51b81526020600482018190526024820152600080516020610e24833981519152604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b600254600160a01b900460ff1615610a75576040805162461bcd60e51b815260206004820152601f60248201527f436f6e747261637420697320616c726561647920696e697469616c6973656400604482015290519081900360640190fd5b610a7e83610bf6565b60028054600392909255600160a01b6001600160a01b03199092166001600160a01b03939093169290921760ff60a01b191617905550565b6000546001600160a01b031690565b6001546001600160a01b0316610ada57600080fd5b6002546001600160a01b0382811691161415610b275760405162461bcd60e51b8152600401808060200182810382526025815260200180610dde6025913960400191505060405180910390fd5b604080516370a0823160e01b815230600482015290516000916001600160a01b038416916370a0823191602480820192602092909190829003018186803b158015610b7157600080fd5b505afa158015610b85573d6000803e3d6000fd5b505050506040513d6020811015610b9b57600080fd5b50516001546040805163a9059cbb60e01b81526001600160a01b0392831660048201526024810184905290519293509084169163a9059cbb916044808201926020929091908290030181600087803b1580156108ea57600080fd5b6000546001600160a01b0316331480610c1857506000546001600160a01b0316155b610c57576040805162461bcd60e51b81526020600482018190526024820152600080516020610e24833981519152604482015290519081900360640190fd5b6001600160a01b038116610c9c5760405162461bcd60e51b8152600401808060200182810382526026815260200180610db86026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b6001546001600160a01b031690565b6002546001600160a01b031681565b60008183610da15760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610d66578181015183820152602001610d4e565b50505050905090810190601f168015610d935780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b506000838581610dad57fe5b049594505050505056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573736e617469766520746f6b656e2066756e64732063616e2774206265207265636f7665726564536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f774f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572a2646970667358221220b72fc1763123dc6df7917907c23740ac6967ede47dfe1aa2b9452361b729817464736f6c634300060a0033"

// DeployMystDEX deploys a new Ethereum contract, binding an instance of MystDEX to it.
func DeployMystDEX(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MystDEX, error) {
	parsed, err := abi.JSON(strings.NewReader(MystDEXABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MystDEXBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MystDEX{MystDEXCaller: MystDEXCaller{contract: contract}, MystDEXTransactor: MystDEXTransactor{contract: contract}, MystDEXFilterer: MystDEXFilterer{contract: contract}}, nil
}

// MystDEX is an auto generated Go binding around an Ethereum contract.
type MystDEX struct {
	MystDEXCaller     // Read-only binding to the contract
	MystDEXTransactor // Write-only binding to the contract
	MystDEXFilterer   // Log filterer for contract events
}

// MystDEXCaller is an auto generated read-only Go binding around an Ethereum contract.
type MystDEXCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MystDEXTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MystDEXTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MystDEXFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MystDEXFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MystDEXSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MystDEXSession struct {
	Contract     *MystDEX          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MystDEXCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MystDEXCallerSession struct {
	Contract *MystDEXCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MystDEXTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MystDEXTransactorSession struct {
	Contract     *MystDEXTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MystDEXRaw is an auto generated low-level Go binding around an Ethereum contract.
type MystDEXRaw struct {
	Contract *MystDEX // Generic contract binding to access the raw methods on
}

// MystDEXCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MystDEXCallerRaw struct {
	Contract *MystDEXCaller // Generic read-only contract binding to access the raw methods on
}

// MystDEXTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MystDEXTransactorRaw struct {
	Contract *MystDEXTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMystDEX creates a new instance of MystDEX, bound to a specific deployed contract.
func NewMystDEX(address common.Address, backend bind.ContractBackend) (*MystDEX, error) {
	contract, err := bindMystDEX(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MystDEX{MystDEXCaller: MystDEXCaller{contract: contract}, MystDEXTransactor: MystDEXTransactor{contract: contract}, MystDEXFilterer: MystDEXFilterer{contract: contract}}, nil
}

// NewMystDEXCaller creates a new read-only instance of MystDEX, bound to a specific deployed contract.
func NewMystDEXCaller(address common.Address, caller bind.ContractCaller) (*MystDEXCaller, error) {
	contract, err := bindMystDEX(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MystDEXCaller{contract: contract}, nil
}

// NewMystDEXTransactor creates a new write-only instance of MystDEX, bound to a specific deployed contract.
func NewMystDEXTransactor(address common.Address, transactor bind.ContractTransactor) (*MystDEXTransactor, error) {
	contract, err := bindMystDEX(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MystDEXTransactor{contract: contract}, nil
}

// NewMystDEXFilterer creates a new log filterer instance of MystDEX, bound to a specific deployed contract.
func NewMystDEXFilterer(address common.Address, filterer bind.ContractFilterer) (*MystDEXFilterer, error) {
	contract, err := bindMystDEX(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MystDEXFilterer{contract: contract}, nil
}

// bindMystDEX binds a generic wrapper to an already deployed contract.
func bindMystDEX(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MystDEXABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MystDEX *MystDEXRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MystDEX.Contract.MystDEXCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MystDEX *MystDEXRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MystDEX.Contract.MystDEXTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MystDEX *MystDEXRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MystDEX.Contract.MystDEXTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MystDEX *MystDEXCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MystDEX.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MystDEX *MystDEXTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MystDEX.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MystDEX *MystDEXTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MystDEX.Contract.contract.Transact(opts, method, params...)
}

// GetFundsDestination is a free data retrieval call binding the contract method 0xf58c5b6e.
//
// Solidity: function getFundsDestination() view returns(address)
func (_MystDEX *MystDEXCaller) GetFundsDestination(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MystDEX.contract.Call(opts, out, "getFundsDestination")
	return *ret0, err
}

// GetFundsDestination is a free data retrieval call binding the contract method 0xf58c5b6e.
//
// Solidity: function getFundsDestination() view returns(address)
func (_MystDEX *MystDEXSession) GetFundsDestination() (common.Address, error) {
	return _MystDEX.Contract.GetFundsDestination(&_MystDEX.CallOpts)
}

// GetFundsDestination is a free data retrieval call binding the contract method 0xf58c5b6e.
//
// Solidity: function getFundsDestination() view returns(address)
func (_MystDEX *MystDEXCallerSession) GetFundsDestination() (common.Address, error) {
	return _MystDEX.Contract.GetFundsDestination(&_MystDEX.CallOpts)
}

// Initialised is a free data retrieval call binding the contract method 0x07003bb4.
//
// Solidity: function initialised() view returns(bool)
func (_MystDEX *MystDEXCaller) Initialised(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MystDEX.contract.Call(opts, out, "initialised")
	return *ret0, err
}

// Initialised is a free data retrieval call binding the contract method 0x07003bb4.
//
// Solidity: function initialised() view returns(bool)
func (_MystDEX *MystDEXSession) Initialised() (bool, error) {
	return _MystDEX.Contract.Initialised(&_MystDEX.CallOpts)
}

// Initialised is a free data retrieval call binding the contract method 0x07003bb4.
//
// Solidity: function initialised() view returns(bool)
func (_MystDEX *MystDEXCallerSession) Initialised() (bool, error) {
	return _MystDEX.Contract.Initialised(&_MystDEX.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MystDEX *MystDEXCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MystDEX.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MystDEX *MystDEXSession) Owner() (common.Address, error) {
	return _MystDEX.Contract.Owner(&_MystDEX.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MystDEX *MystDEXCallerSession) Owner() (common.Address, error) {
	return _MystDEX.Contract.Owner(&_MystDEX.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_MystDEX *MystDEXCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MystDEX.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_MystDEX *MystDEXSession) Token() (common.Address, error) {
	return _MystDEX.Contract.Token(&_MystDEX.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_MystDEX *MystDEXCallerSession) Token() (common.Address, error) {
	return _MystDEX.Contract.Token(&_MystDEX.CallOpts)
}

// ClaimEthers is a paid mutator transaction binding the contract method 0x6931b550.
//
// Solidity: function claimEthers() returns()
func (_MystDEX *MystDEXTransactor) ClaimEthers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MystDEX.contract.Transact(opts, "claimEthers")
}

// ClaimEthers is a paid mutator transaction binding the contract method 0x6931b550.
//
// Solidity: function claimEthers() returns()
func (_MystDEX *MystDEXSession) ClaimEthers() (*types.Transaction, error) {
	return _MystDEX.Contract.ClaimEthers(&_MystDEX.TransactOpts)
}

// ClaimEthers is a paid mutator transaction binding the contract method 0x6931b550.
//
// Solidity: function claimEthers() returns()
func (_MystDEX *MystDEXTransactorSession) ClaimEthers() (*types.Transaction, error) {
	return _MystDEX.Contract.ClaimEthers(&_MystDEX.TransactOpts)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0xdf8de3e7.
//
// Solidity: function claimTokens(address _token) returns()
func (_MystDEX *MystDEXTransactor) ClaimTokens(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _MystDEX.contract.Transact(opts, "claimTokens", _token)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0xdf8de3e7.
//
// Solidity: function claimTokens(address _token) returns()
func (_MystDEX *MystDEXSession) ClaimTokens(_token common.Address) (*types.Transaction, error) {
	return _MystDEX.Contract.ClaimTokens(&_MystDEX.TransactOpts, _token)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0xdf8de3e7.
//
// Solidity: function claimTokens(address _token) returns()
func (_MystDEX *MystDEXTransactorSession) ClaimTokens(_token common.Address) (*types.Transaction, error) {
	return _MystDEX.Contract.ClaimTokens(&_MystDEX.TransactOpts, _token)
}

// Initialise is a paid mutator transaction binding the contract method 0x8595d149.
//
// Solidity: function initialise(address _dexOwner, address _token, uint256 _rate) returns()
func (_MystDEX *MystDEXTransactor) Initialise(opts *bind.TransactOpts, _dexOwner common.Address, _token common.Address, _rate *big.Int) (*types.Transaction, error) {
	return _MystDEX.contract.Transact(opts, "initialise", _dexOwner, _token, _rate)
}

// Initialise is a paid mutator transaction binding the contract method 0x8595d149.
//
// Solidity: function initialise(address _dexOwner, address _token, uint256 _rate) returns()
func (_MystDEX *MystDEXSession) Initialise(_dexOwner common.Address, _token common.Address, _rate *big.Int) (*types.Transaction, error) {
	return _MystDEX.Contract.Initialise(&_MystDEX.TransactOpts, _dexOwner, _token, _rate)
}

// Initialise is a paid mutator transaction binding the contract method 0x8595d149.
//
// Solidity: function initialise(address _dexOwner, address _token, uint256 _rate) returns()
func (_MystDEX *MystDEXTransactorSession) Initialise(_dexOwner common.Address, _token common.Address, _rate *big.Int) (*types.Transaction, error) {
	return _MystDEX.Contract.Initialise(&_MystDEX.TransactOpts, _dexOwner, _token, _rate)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MystDEX *MystDEXTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MystDEX.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MystDEX *MystDEXSession) RenounceOwnership() (*types.Transaction, error) {
	return _MystDEX.Contract.RenounceOwnership(&_MystDEX.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MystDEX *MystDEXTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MystDEX.Contract.RenounceOwnership(&_MystDEX.TransactOpts)
}

// SetFundsDestination is a paid mutator transaction binding the contract method 0x238e130a.
//
// Solidity: function setFundsDestination(address _newDestination) returns()
func (_MystDEX *MystDEXTransactor) SetFundsDestination(opts *bind.TransactOpts, _newDestination common.Address) (*types.Transaction, error) {
	return _MystDEX.contract.Transact(opts, "setFundsDestination", _newDestination)
}

// SetFundsDestination is a paid mutator transaction binding the contract method 0x238e130a.
//
// Solidity: function setFundsDestination(address _newDestination) returns()
func (_MystDEX *MystDEXSession) SetFundsDestination(_newDestination common.Address) (*types.Transaction, error) {
	return _MystDEX.Contract.SetFundsDestination(&_MystDEX.TransactOpts, _newDestination)
}

// SetFundsDestination is a paid mutator transaction binding the contract method 0x238e130a.
//
// Solidity: function setFundsDestination(address _newDestination) returns()
func (_MystDEX *MystDEXTransactorSession) SetFundsDestination(_newDestination common.Address) (*types.Transaction, error) {
	return _MystDEX.Contract.SetFundsDestination(&_MystDEX.TransactOpts, _newDestination)
}

// SetRate is a paid mutator transaction binding the contract method 0x34fcf437.
//
// Solidity: function setRate(uint256 _newRate) returns()
func (_MystDEX *MystDEXTransactor) SetRate(opts *bind.TransactOpts, _newRate *big.Int) (*types.Transaction, error) {
	return _MystDEX.contract.Transact(opts, "setRate", _newRate)
}

// SetRate is a paid mutator transaction binding the contract method 0x34fcf437.
//
// Solidity: function setRate(uint256 _newRate) returns()
func (_MystDEX *MystDEXSession) SetRate(_newRate *big.Int) (*types.Transaction, error) {
	return _MystDEX.Contract.SetRate(&_MystDEX.TransactOpts, _newRate)
}

// SetRate is a paid mutator transaction binding the contract method 0x34fcf437.
//
// Solidity: function setRate(uint256 _newRate) returns()
func (_MystDEX *MystDEXTransactorSession) SetRate(_newRate *big.Int) (*types.Transaction, error) {
	return _MystDEX.Contract.SetRate(&_MystDEX.TransactOpts, _newRate)
}

// TransferEthers is a paid mutator transaction binding the contract method 0x1254e64d.
//
// Solidity: function transferEthers(address _to, uint256 _amount) returns()
func (_MystDEX *MystDEXTransactor) TransferEthers(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MystDEX.contract.Transact(opts, "transferEthers", _to, _amount)
}

// TransferEthers is a paid mutator transaction binding the contract method 0x1254e64d.
//
// Solidity: function transferEthers(address _to, uint256 _amount) returns()
func (_MystDEX *MystDEXSession) TransferEthers(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MystDEX.Contract.TransferEthers(&_MystDEX.TransactOpts, _to, _amount)
}

// TransferEthers is a paid mutator transaction binding the contract method 0x1254e64d.
//
// Solidity: function transferEthers(address _to, uint256 _amount) returns()
func (_MystDEX *MystDEXTransactorSession) TransferEthers(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MystDEX.Contract.TransferEthers(&_MystDEX.TransactOpts, _to, _amount)
}

// TransferMyst is a paid mutator transaction binding the contract method 0x38d2e411.
//
// Solidity: function transferMyst(address _to, uint256 _amount) returns()
func (_MystDEX *MystDEXTransactor) TransferMyst(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MystDEX.contract.Transact(opts, "transferMyst", _to, _amount)
}

// TransferMyst is a paid mutator transaction binding the contract method 0x38d2e411.
//
// Solidity: function transferMyst(address _to, uint256 _amount) returns()
func (_MystDEX *MystDEXSession) TransferMyst(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MystDEX.Contract.TransferMyst(&_MystDEX.TransactOpts, _to, _amount)
}

// TransferMyst is a paid mutator transaction binding the contract method 0x38d2e411.
//
// Solidity: function transferMyst(address _to, uint256 _amount) returns()
func (_MystDEX *MystDEXTransactorSession) TransferMyst(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MystDEX.Contract.TransferMyst(&_MystDEX.TransactOpts, _to, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MystDEX *MystDEXTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MystDEX.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MystDEX *MystDEXSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MystDEX.Contract.TransferOwnership(&_MystDEX.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MystDEX *MystDEXTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MystDEX.Contract.TransferOwnership(&_MystDEX.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MystDEX *MystDEXTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MystDEX.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MystDEX *MystDEXSession) Receive() (*types.Transaction, error) {
	return _MystDEX.Contract.Receive(&_MystDEX.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MystDEX *MystDEXTransactorSession) Receive() (*types.Transaction, error) {
	return _MystDEX.Contract.Receive(&_MystDEX.TransactOpts)
}

// MystDEXDestinationChangedIterator is returned from FilterDestinationChanged and is used to iterate over the raw logs and unpacked data for DestinationChanged events raised by the MystDEX contract.
type MystDEXDestinationChangedIterator struct {
	Event *MystDEXDestinationChanged // Event containing the contract specifics and raw log

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
func (it *MystDEXDestinationChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MystDEXDestinationChanged)
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
		it.Event = new(MystDEXDestinationChanged)
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
func (it *MystDEXDestinationChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MystDEXDestinationChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MystDEXDestinationChanged represents a DestinationChanged event raised by the MystDEX contract.
type MystDEXDestinationChanged struct {
	PreviousDestination common.Address
	NewDestination      common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDestinationChanged is a free log retrieval operation binding the contract event 0xe1a66d77649cf0a57b9937073549f30f1c82bb865aaf066d2f299e37a62c6aad.
//
// Solidity: event DestinationChanged(address indexed previousDestination, address indexed newDestination)
func (_MystDEX *MystDEXFilterer) FilterDestinationChanged(opts *bind.FilterOpts, previousDestination []common.Address, newDestination []common.Address) (*MystDEXDestinationChangedIterator, error) {

	var previousDestinationRule []interface{}
	for _, previousDestinationItem := range previousDestination {
		previousDestinationRule = append(previousDestinationRule, previousDestinationItem)
	}
	var newDestinationRule []interface{}
	for _, newDestinationItem := range newDestination {
		newDestinationRule = append(newDestinationRule, newDestinationItem)
	}

	logs, sub, err := _MystDEX.contract.FilterLogs(opts, "DestinationChanged", previousDestinationRule, newDestinationRule)
	if err != nil {
		return nil, err
	}
	return &MystDEXDestinationChangedIterator{contract: _MystDEX.contract, event: "DestinationChanged", logs: logs, sub: sub}, nil
}

// WatchDestinationChanged is a free log subscription operation binding the contract event 0xe1a66d77649cf0a57b9937073549f30f1c82bb865aaf066d2f299e37a62c6aad.
//
// Solidity: event DestinationChanged(address indexed previousDestination, address indexed newDestination)
func (_MystDEX *MystDEXFilterer) WatchDestinationChanged(opts *bind.WatchOpts, sink chan<- *MystDEXDestinationChanged, previousDestination []common.Address, newDestination []common.Address) (event.Subscription, error) {

	var previousDestinationRule []interface{}
	for _, previousDestinationItem := range previousDestination {
		previousDestinationRule = append(previousDestinationRule, previousDestinationItem)
	}
	var newDestinationRule []interface{}
	for _, newDestinationItem := range newDestination {
		newDestinationRule = append(newDestinationRule, newDestinationItem)
	}

	logs, sub, err := _MystDEX.contract.WatchLogs(opts, "DestinationChanged", previousDestinationRule, newDestinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MystDEXDestinationChanged)
				if err := _MystDEX.contract.UnpackLog(event, "DestinationChanged", log); err != nil {
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

// ParseDestinationChanged is a log parse operation binding the contract event 0xe1a66d77649cf0a57b9937073549f30f1c82bb865aaf066d2f299e37a62c6aad.
//
// Solidity: event DestinationChanged(address indexed previousDestination, address indexed newDestination)
func (_MystDEX *MystDEXFilterer) ParseDestinationChanged(log types.Log) (*MystDEXDestinationChanged, error) {
	event := new(MystDEXDestinationChanged)
	if err := _MystDEX.contract.UnpackLog(event, "DestinationChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MystDEXOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MystDEX contract.
type MystDEXOwnershipTransferredIterator struct {
	Event *MystDEXOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MystDEXOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MystDEXOwnershipTransferred)
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
		it.Event = new(MystDEXOwnershipTransferred)
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
func (it *MystDEXOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MystDEXOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MystDEXOwnershipTransferred represents a OwnershipTransferred event raised by the MystDEX contract.
type MystDEXOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MystDEX *MystDEXFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MystDEXOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MystDEX.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MystDEXOwnershipTransferredIterator{contract: _MystDEX.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MystDEX *MystDEXFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MystDEXOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MystDEX.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MystDEXOwnershipTransferred)
				if err := _MystDEX.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MystDEX *MystDEXFilterer) ParseOwnershipTransferred(log types.Log) (*MystDEXOwnershipTransferred, error) {
	event := new(MystDEXOwnershipTransferred)
	if err := _MystDEX.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
