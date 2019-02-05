// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abigen

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

// ERC20AwareBalanceABI is the input ABI used to generate the binding from.
const ERC20AwareBalanceABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"ERC20Token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"erc20tokenAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ERC20AwareBalanceBin is the compiled bytecode used for deploying new contracts.
const ERC20AwareBalanceBin = `0x608060405234801561001057600080fd5b506040516020806101278339810180604052602081101561003057600080fd5b505160008054600160a060020a03909216600160a060020a031990921691909117905560c6806100616000396000f3fe608060405260043610603e5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416637a80760e81146043575b600080fd5b348015604e57600080fd5b506055607e565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b60005473ffffffffffffffffffffffffffffffffffffffff168156fea165627a7a723058203ea703367e19773ea454846b7a4935ad4b292c00142fa7363d5ad8895fd04d970029`

// DeployERC20AwareBalance deploys a new Ethereum contract, binding an instance of ERC20AwareBalance to it.
func DeployERC20AwareBalance(auth *bind.TransactOpts, backend bind.ContractBackend, erc20tokenAddress common.Address) (common.Address, *types.Transaction, *ERC20AwareBalance, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20AwareBalanceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20AwareBalanceBin), backend, erc20tokenAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20AwareBalance{ERC20AwareBalanceCaller: ERC20AwareBalanceCaller{contract: contract}, ERC20AwareBalanceTransactor: ERC20AwareBalanceTransactor{contract: contract}, ERC20AwareBalanceFilterer: ERC20AwareBalanceFilterer{contract: contract}}, nil
}

// ERC20AwareBalance is an auto generated Go binding around an Ethereum contract.
type ERC20AwareBalance struct {
	ERC20AwareBalanceCaller     // Read-only binding to the contract
	ERC20AwareBalanceTransactor // Write-only binding to the contract
	ERC20AwareBalanceFilterer   // Log filterer for contract events
}

// ERC20AwareBalanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20AwareBalanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20AwareBalanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20AwareBalanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20AwareBalanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20AwareBalanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20AwareBalanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20AwareBalanceSession struct {
	Contract     *ERC20AwareBalance // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ERC20AwareBalanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20AwareBalanceCallerSession struct {
	Contract *ERC20AwareBalanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ERC20AwareBalanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20AwareBalanceTransactorSession struct {
	Contract     *ERC20AwareBalanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ERC20AwareBalanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20AwareBalanceRaw struct {
	Contract *ERC20AwareBalance // Generic contract binding to access the raw methods on
}

// ERC20AwareBalanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20AwareBalanceCallerRaw struct {
	Contract *ERC20AwareBalanceCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20AwareBalanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20AwareBalanceTransactorRaw struct {
	Contract *ERC20AwareBalanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20AwareBalance creates a new instance of ERC20AwareBalance, bound to a specific deployed contract.
func NewERC20AwareBalance(address common.Address, backend bind.ContractBackend) (*ERC20AwareBalance, error) {
	contract, err := bindERC20AwareBalance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20AwareBalance{ERC20AwareBalanceCaller: ERC20AwareBalanceCaller{contract: contract}, ERC20AwareBalanceTransactor: ERC20AwareBalanceTransactor{contract: contract}, ERC20AwareBalanceFilterer: ERC20AwareBalanceFilterer{contract: contract}}, nil
}

// NewERC20AwareBalanceCaller creates a new read-only instance of ERC20AwareBalance, bound to a specific deployed contract.
func NewERC20AwareBalanceCaller(address common.Address, caller bind.ContractCaller) (*ERC20AwareBalanceCaller, error) {
	contract, err := bindERC20AwareBalance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20AwareBalanceCaller{contract: contract}, nil
}

// NewERC20AwareBalanceTransactor creates a new write-only instance of ERC20AwareBalance, bound to a specific deployed contract.
func NewERC20AwareBalanceTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20AwareBalanceTransactor, error) {
	contract, err := bindERC20AwareBalance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20AwareBalanceTransactor{contract: contract}, nil
}

// NewERC20AwareBalanceFilterer creates a new log filterer instance of ERC20AwareBalance, bound to a specific deployed contract.
func NewERC20AwareBalanceFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20AwareBalanceFilterer, error) {
	contract, err := bindERC20AwareBalance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20AwareBalanceFilterer{contract: contract}, nil
}

// bindERC20AwareBalance binds a generic wrapper to an already deployed contract.
func bindERC20AwareBalance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20AwareBalanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20AwareBalance *ERC20AwareBalanceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20AwareBalance.Contract.ERC20AwareBalanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20AwareBalance *ERC20AwareBalanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20AwareBalance.Contract.ERC20AwareBalanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20AwareBalance *ERC20AwareBalanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20AwareBalance.Contract.ERC20AwareBalanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20AwareBalance *ERC20AwareBalanceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20AwareBalance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20AwareBalance *ERC20AwareBalanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20AwareBalance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20AwareBalance *ERC20AwareBalanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20AwareBalance.Contract.contract.Transact(opts, method, params...)
}

// ERC20Token is a free data retrieval call binding the contract method 0x7a80760e.
//
// Solidity: function ERC20Token() constant returns(address)
func (_ERC20AwareBalance *ERC20AwareBalanceCaller) ERC20Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC20AwareBalance.contract.Call(opts, out, "ERC20Token")
	return *ret0, err
}

// ERC20Token is a free data retrieval call binding the contract method 0x7a80760e.
//
// Solidity: function ERC20Token() constant returns(address)
func (_ERC20AwareBalance *ERC20AwareBalanceSession) ERC20Token() (common.Address, error) {
	return _ERC20AwareBalance.Contract.ERC20Token(&_ERC20AwareBalance.CallOpts)
}

// ERC20Token is a free data retrieval call binding the contract method 0x7a80760e.
//
// Solidity: function ERC20Token() constant returns(address)
func (_ERC20AwareBalance *ERC20AwareBalanceCallerSession) ERC20Token() (common.Address, error) {
	return _ERC20AwareBalance.Contract.ERC20Token(&_ERC20AwareBalance.CallOpts)
}

// ERC20AwareRegistryABI is the input ABI used to generate the binding from.
const ERC20AwareRegistryABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"ERC20Token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"erc20tokenAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ERC20AwareRegistryBin is the compiled bytecode used for deploying new contracts.
const ERC20AwareRegistryBin = `0x608060405234801561001057600080fd5b506040516020806101278339810180604052602081101561003057600080fd5b505160008054600160a060020a03909216600160a060020a031990921691909117905560c6806100616000396000f3fe608060405260043610603e5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416637a80760e81146043575b600080fd5b348015604e57600080fd5b506055607e565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b60005473ffffffffffffffffffffffffffffffffffffffff168156fea165627a7a72305820658f63a9ee5644d3a1a8dbd20a28882efe3a0ab056c865e288336ea29d8e60530029`

// DeployERC20AwareRegistry deploys a new Ethereum contract, binding an instance of ERC20AwareRegistry to it.
func DeployERC20AwareRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, erc20tokenAddress common.Address) (common.Address, *types.Transaction, *ERC20AwareRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20AwareRegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20AwareRegistryBin), backend, erc20tokenAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20AwareRegistry{ERC20AwareRegistryCaller: ERC20AwareRegistryCaller{contract: contract}, ERC20AwareRegistryTransactor: ERC20AwareRegistryTransactor{contract: contract}, ERC20AwareRegistryFilterer: ERC20AwareRegistryFilterer{contract: contract}}, nil
}

// ERC20AwareRegistry is an auto generated Go binding around an Ethereum contract.
type ERC20AwareRegistry struct {
	ERC20AwareRegistryCaller     // Read-only binding to the contract
	ERC20AwareRegistryTransactor // Write-only binding to the contract
	ERC20AwareRegistryFilterer   // Log filterer for contract events
}

// ERC20AwareRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20AwareRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20AwareRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20AwareRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20AwareRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20AwareRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20AwareRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20AwareRegistrySession struct {
	Contract     *ERC20AwareRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ERC20AwareRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20AwareRegistryCallerSession struct {
	Contract *ERC20AwareRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ERC20AwareRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20AwareRegistryTransactorSession struct {
	Contract     *ERC20AwareRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ERC20AwareRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20AwareRegistryRaw struct {
	Contract *ERC20AwareRegistry // Generic contract binding to access the raw methods on
}

// ERC20AwareRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20AwareRegistryCallerRaw struct {
	Contract *ERC20AwareRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20AwareRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20AwareRegistryTransactorRaw struct {
	Contract *ERC20AwareRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20AwareRegistry creates a new instance of ERC20AwareRegistry, bound to a specific deployed contract.
func NewERC20AwareRegistry(address common.Address, backend bind.ContractBackend) (*ERC20AwareRegistry, error) {
	contract, err := bindERC20AwareRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20AwareRegistry{ERC20AwareRegistryCaller: ERC20AwareRegistryCaller{contract: contract}, ERC20AwareRegistryTransactor: ERC20AwareRegistryTransactor{contract: contract}, ERC20AwareRegistryFilterer: ERC20AwareRegistryFilterer{contract: contract}}, nil
}

// NewERC20AwareRegistryCaller creates a new read-only instance of ERC20AwareRegistry, bound to a specific deployed contract.
func NewERC20AwareRegistryCaller(address common.Address, caller bind.ContractCaller) (*ERC20AwareRegistryCaller, error) {
	contract, err := bindERC20AwareRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20AwareRegistryCaller{contract: contract}, nil
}

// NewERC20AwareRegistryTransactor creates a new write-only instance of ERC20AwareRegistry, bound to a specific deployed contract.
func NewERC20AwareRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20AwareRegistryTransactor, error) {
	contract, err := bindERC20AwareRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20AwareRegistryTransactor{contract: contract}, nil
}

// NewERC20AwareRegistryFilterer creates a new log filterer instance of ERC20AwareRegistry, bound to a specific deployed contract.
func NewERC20AwareRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20AwareRegistryFilterer, error) {
	contract, err := bindERC20AwareRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20AwareRegistryFilterer{contract: contract}, nil
}

// bindERC20AwareRegistry binds a generic wrapper to an already deployed contract.
func bindERC20AwareRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20AwareRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20AwareRegistry *ERC20AwareRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20AwareRegistry.Contract.ERC20AwareRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20AwareRegistry *ERC20AwareRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20AwareRegistry.Contract.ERC20AwareRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20AwareRegistry *ERC20AwareRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20AwareRegistry.Contract.ERC20AwareRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20AwareRegistry *ERC20AwareRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20AwareRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20AwareRegistry *ERC20AwareRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20AwareRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20AwareRegistry *ERC20AwareRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20AwareRegistry.Contract.contract.Transact(opts, method, params...)
}

// ERC20Token is a free data retrieval call binding the contract method 0x7a80760e.
//
// Solidity: function ERC20Token() constant returns(address)
func (_ERC20AwareRegistry *ERC20AwareRegistryCaller) ERC20Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC20AwareRegistry.contract.Call(opts, out, "ERC20Token")
	return *ret0, err
}

// ERC20Token is a free data retrieval call binding the contract method 0x7a80760e.
//
// Solidity: function ERC20Token() constant returns(address)
func (_ERC20AwareRegistry *ERC20AwareRegistrySession) ERC20Token() (common.Address, error) {
	return _ERC20AwareRegistry.Contract.ERC20Token(&_ERC20AwareRegistry.CallOpts)
}

// ERC20Token is a free data retrieval call binding the contract method 0x7a80760e.
//
// Solidity: function ERC20Token() constant returns(address)
func (_ERC20AwareRegistry *ERC20AwareRegistryCallerSession) ERC20Token() (common.Address, error) {
	return _ERC20AwareRegistry.Contract.ERC20Token(&_ERC20AwareRegistry.CallOpts)
}

// IdentityPromisesABI is the input ABI used to generate the binding from.
const IdentityPromisesABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"registrationFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"receiverAndSigns\",\"type\":\"bytes32\"},{\"name\":\"extraDataHash\",\"type\":\"bytes32\"},{\"name\":\"seq\",\"type\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"sender_R\",\"type\":\"bytes32\"},{\"name\":\"sender_S\",\"type\":\"bytes32\"},{\"name\":\"receiver_R\",\"type\":\"bytes32\"},{\"name\":\"receiver_S\",\"type\":\"bytes32\"}],\"name\":\"clearPromise\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"changeRegistrationFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERC20Token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"}],\"name\":\"getPublicKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pubKeyPart1\",\"type\":\"bytes32\"},{\"name\":\"pubKeyPart2\",\"type\":\"bytes32\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"RegisterIdentity\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"}],\"name\":\"isRegistered\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"topUp\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"transferCollectedFeeTo\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"collectedFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"clearedPromises\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"erc20address\",\"type\":\"address\"},{\"name\":\"registrationFee\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"seqNo\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PromiseCleared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalBalance\",\"type\":\"uint256\"}],\"name\":\"ToppedUp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalBalance\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"identity\",\"type\":\"address\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// IdentityPromisesBin is the compiled bytecode used for deploying new contracts.
const IdentityPromisesBin = `0x608060405234801561001057600080fd5b506040516040806112d68339810180604052604081101561003057600080fd5b50805160209091015160008054600160a060020a031916331780825560405184928392839286928492600160a060020a0316917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a360018054600160a060020a03928316600160a060020a0319918216179091556003929092556005805494909116939091169290921790915550505050611203806100d36000396000f3fe6080604052600436106100f05763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166314c44e0981146100f557806327e235e31461011c5780634fe0d3321461014f57806350050769146101b85780635ebfdfc6146101e4578063715018a6146102235780637a80760e14610238578063857cdbb8146102695780638da5cb5b146102b55780638f32d59b146102ca578063a17fc7f3146102df578063c3c5a54714610324578063d6f7ddf914610357578063e325253714610390578063e811f50a146103c3578063f2fde38b146103d8578063f6c1a23e1461040b575b600080fd5b34801561010157600080fd5b5061010a610446565b60408051918252519081900360200190f35b34801561012857600080fd5b5061010a6004803603602081101561013f57600080fd5b5035600160a060020a031661044c565b34801561015b57600080fd5b506101a4600480360361010081101561017357600080fd5b5080359060208101359060408101359060608101359060808101359060a08101359060c08101359060e0013561045e565b604080519115158252519081900360200190f35b3480156101c457600080fd5b506101e2600480360360208110156101db57600080fd5b5035610781565b005b3480156101f057600080fd5b506101a46004803603608081101561020757600080fd5b5080359060ff6020820135169060408101359060600135610799565b34801561022f57600080fd5b506101e26109fc565b34801561024457600080fd5b5061024d610a66565b60408051600160a060020a039092168252519081900360200190f35b34801561027557600080fd5b5061029c6004803603602081101561028c57600080fd5b5035600160a060020a0316610a75565b6040805192835260208301919091528051918290030190f35b3480156102c157600080fd5b5061024d610a98565b3480156102d657600080fd5b506101a4610aa7565b3480156102eb57600080fd5b506101a4600480360360a081101561030257600080fd5b5080359060208101359060ff6040820135169060608101359060800135610ab8565b34801561033057600080fd5b506101a46004803603602081101561034757600080fd5b5035600160a060020a0316610d50565b34801561036357600080fd5b506101a46004803603604081101561037a57600080fd5b50600160a060020a038135169060200135610d83565b34801561039c57600080fd5b506101a4600480360360208110156103b357600080fd5b5035600160a060020a0316610ea9565b3480156103cf57600080fd5b5061010a610f8a565b3480156103e457600080fd5b506101e2600480360360208110156103fb57600080fd5b5035600160a060020a0316610f90565b34801561041757600080fd5b5061010a6004803603604081101561042e57600080fd5b50600160a060020a0381358116916020013516610faf565b60035481565b60066020526000908152604090205481565b60008060008061046d8c610fcc565b92509250925060006040805190810160405280600e81526020017f497373756572207072656669783a0000000000000000000000000000000000008152508c858d8d6040516020018086805190602001908083835b602083106104e15780518252601f1990920191602091820191016104c2565b6001836020036101000a03801982511681845116808217855250505050505090500185815260200184600160a060020a0316600160a060020a03166c01000000000000000000000000028152601401838152602001828152602001955050505050506040516020818303038152906040528051906020012090506000600182858c8c60405160008152602001604052604051808581526020018460ff1660ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa1580156105ba573d6000803e3d6000fd5b505050602060405103519050600060016040805190810160405280601081526020017f5265636569766572207072656669783a0000000000000000000000000000000081525084846040516020018084805190602001908083835b602083106106345780518252601f199092019160209182019101610615565b51815160209384036101000a6000190180199092169116179052920194855250600160a060020a03929092166c01000000000000000000000000028383015250604080516014818503018152603484018083528151918401919091206000909152605484018083525260ff89166074840152609483018e905260b483018d90525160d480840194509192601f19820192908290030190855afa1580156106de573d6000803e3d6000fd5b5050604051601f190151915050600160a060020a038216151561070057600080fd5b600160a060020a038116151561071557600080fd5b600160a060020a038181169087161461072d57600080fd5b61073682610d50565b151561074157600080fd5b61074a86610d50565b151561075557600080fd5b60008c1161076257600080fd5b61076e86838f8f611007565b9f9e505050505050505050505050505050565b610789610aa7565b151561079457600080fd5b600355565b600060606040805190810160405280601181526020017f576974686472617720726571756573743a000000000000000000000000000000815250866040516020018083805190602001908083835b602083106108065780518252601f1990920191602091820191016107e7565b51815160209384036101000a60001901801990921691161790529201938452506040805180850381528483018083528151828501206000918290528387018085525260ff8c166060870152608086018b905260a086018a9052915190965090945060019360c08082019450601f19830192918290030190855afa158015610891573d6000803e3d6000fd5b5050604051601f190151915050600160a060020a03811615156108b357600080fd5b600087116108c057600080fd5b600160a060020a0381166000908152600660205260409020548711156108e557600080fd5b600160a060020a03808216600090815260066020908152604080832080548c9003905560055481517fa9059cbb000000000000000000000000000000000000000000000000000000008152336004820152602481018d9052915194169363a9059cbb93604480840194938390030190829087803b15801561096557600080fd5b505af1158015610979573d6000803e3d6000fd5b505050506040513d602081101561098f57600080fd5b5051151561099c57600080fd5b600160a060020a0381166000818152600660209081526040918290205482518b81529182015281517f92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6929181900390910190a25060019695505050505050565b610a04610aa7565b1515610a0f57600080fd5b60008054604051600160a060020a03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a36000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600554600160a060020a031681565b600160a060020a0316600090815260026020526040902080546001909101549091565b600054600160a060020a031690565b600054600160a060020a0316331490565b6040805160208082018890528183018790528251808303840181526060830180855281519183019190912060a08401855260108083527f5265676973746572207072656669783a000000000000000000000000000000006080909501948552945160009591948694600194938d938d9392019182918083835b60208310610b505780518252601f199092019160209182019101610b31565b51815160209384036101000a6000190180199092169116179052920194855250838101929092525060408051808403830181528184018083528151918401919091206000909152606084018083525260ff8b16608084015260a083018a905260c083018990525160e080840194509192601f19820192908290030190855afa158015610be0573d6000803e3d6000fd5b5050604051601f190151915050600160a060020a0381161515610c0257600080fd5b600160a060020a0382811690821614610c1a57600080fd5b610c2381610d50565b15610c2d57600080fd5b600380546004805490910181556040805180820182528b815260208082018c8152600160a060020a038088166000908152600284528581209451855591516001948501559254955484517f23b872dd00000000000000000000000000000000000000000000000000000000815233968101969096523060248701526044860152925194909116936323b872dd93606480820194918390030190829087803b158015610cd757600080fd5b505af1158015610ceb573d6000803e3d6000fd5b505050506040513d6020811015610d0157600080fd5b50511515610d0e57600080fd5b604051600160a060020a038216907f2d3734a8e47ac8316e500ac231c90a6e1848ca2285f40d07eaa52005e4b3a0e990600090a2506001979650505050505050565b600160a060020a0381166000908152600260205260408120805482108015610d7c575060018101546000105b9392505050565b6000808211610d9157600080fd5b600160a060020a03808416600090815260066020908152604080832080548701905560055481517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526044810188905291519416936323b872dd93606480840194938390030190829087803b158015610e1657600080fd5b505af1158015610e2a573d6000803e3d6000fd5b505050506040513d6020811015610e4057600080fd5b50511515610e4d57600080fd5b600160a060020a0383166000818152600660209081526040918290205482518681529182015281517fbdde76a89d276b6d334c784be5c2d00d6c2219d11a2f2c80e00b85144845ab4d929181900390910190a250600192915050565b6000610eb3610aa7565b1515610ebe57600080fd5b600454600010610ecd57600080fd5b600480546000808355600154604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a038881169682019690965260248101859052905193949091169263a9059cbb92604480840193602093929083900390910190829087803b158015610f4a57600080fd5b505af1158015610f5e573d6000803e3d6000fd5b505050506040513d6020811015610f7457600080fd5b50511515610f8157600080fd5b50600192915050565b60045481565b610f98610aa7565b1515610fa357600080fd5b610fac816110dc565b50565b600760209081526000928352604080842090915290825290205481565b907f0100000000000000000000000000000000000000000000000000000000000000600a83901a810281900491600b84901a82029190910490565b600160a060020a038084166000908152600760209081526040808320938816835292905290812054831161103a57600080fd5b600160a060020a038085166000908152600760209081526040808320938916835292905290812084905561106f858785611159565b905082811461107d57600080fd5b85600160a060020a031685600160a060020a03167f3d35f4a81f47a50c001e0e834334665e3a64680385c0a8d2fdfe2e0d209645768686604051808381526020018281526020019250505060405180910390a350600195945050505050565b600160a060020a03811615156110f157600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600080821161116757600080fd5b600160a060020a03841660009081526006602052604081205461118a90846111c1565b600160a060020a03808716600090815260066020526040808220805485900390559187168152208054820190559150509392505050565b60008183106111d05781610d7c565b509091905056fea165627a7a7230582045dd81e9ec29ee6c4b3b94d075fccfe744c2896c2e4439c3b471e0909a8ee8f70029`

// DeployIdentityPromises deploys a new Ethereum contract, binding an instance of IdentityPromises to it.
func DeployIdentityPromises(auth *bind.TransactOpts, backend bind.ContractBackend, erc20address common.Address, registrationFee *big.Int) (common.Address, *types.Transaction, *IdentityPromises, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityPromisesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IdentityPromisesBin), backend, erc20address, registrationFee)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IdentityPromises{IdentityPromisesCaller: IdentityPromisesCaller{contract: contract}, IdentityPromisesTransactor: IdentityPromisesTransactor{contract: contract}, IdentityPromisesFilterer: IdentityPromisesFilterer{contract: contract}}, nil
}

// IdentityPromises is an auto generated Go binding around an Ethereum contract.
type IdentityPromises struct {
	IdentityPromisesCaller     // Read-only binding to the contract
	IdentityPromisesTransactor // Write-only binding to the contract
	IdentityPromisesFilterer   // Log filterer for contract events
}

// IdentityPromisesCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdentityPromisesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityPromisesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdentityPromisesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityPromisesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdentityPromisesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityPromisesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdentityPromisesSession struct {
	Contract     *IdentityPromises // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdentityPromisesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdentityPromisesCallerSession struct {
	Contract *IdentityPromisesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IdentityPromisesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdentityPromisesTransactorSession struct {
	Contract     *IdentityPromisesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IdentityPromisesRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdentityPromisesRaw struct {
	Contract *IdentityPromises // Generic contract binding to access the raw methods on
}

// IdentityPromisesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdentityPromisesCallerRaw struct {
	Contract *IdentityPromisesCaller // Generic read-only contract binding to access the raw methods on
}

// IdentityPromisesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdentityPromisesTransactorRaw struct {
	Contract *IdentityPromisesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdentityPromises creates a new instance of IdentityPromises, bound to a specific deployed contract.
func NewIdentityPromises(address common.Address, backend bind.ContractBackend) (*IdentityPromises, error) {
	contract, err := bindIdentityPromises(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IdentityPromises{IdentityPromisesCaller: IdentityPromisesCaller{contract: contract}, IdentityPromisesTransactor: IdentityPromisesTransactor{contract: contract}, IdentityPromisesFilterer: IdentityPromisesFilterer{contract: contract}}, nil
}

// NewIdentityPromisesCaller creates a new read-only instance of IdentityPromises, bound to a specific deployed contract.
func NewIdentityPromisesCaller(address common.Address, caller bind.ContractCaller) (*IdentityPromisesCaller, error) {
	contract, err := bindIdentityPromises(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityPromisesCaller{contract: contract}, nil
}

// NewIdentityPromisesTransactor creates a new write-only instance of IdentityPromises, bound to a specific deployed contract.
func NewIdentityPromisesTransactor(address common.Address, transactor bind.ContractTransactor) (*IdentityPromisesTransactor, error) {
	contract, err := bindIdentityPromises(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityPromisesTransactor{contract: contract}, nil
}

// NewIdentityPromisesFilterer creates a new log filterer instance of IdentityPromises, bound to a specific deployed contract.
func NewIdentityPromisesFilterer(address common.Address, filterer bind.ContractFilterer) (*IdentityPromisesFilterer, error) {
	contract, err := bindIdentityPromises(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdentityPromisesFilterer{contract: contract}, nil
}

// bindIdentityPromises binds a generic wrapper to an already deployed contract.
func bindIdentityPromises(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityPromisesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityPromises *IdentityPromisesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IdentityPromises.Contract.IdentityPromisesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityPromises *IdentityPromisesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityPromises.Contract.IdentityPromisesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityPromises *IdentityPromisesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityPromises.Contract.IdentityPromisesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityPromises *IdentityPromisesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IdentityPromises.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityPromises *IdentityPromisesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityPromises.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityPromises *IdentityPromisesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityPromises.Contract.contract.Transact(opts, method, params...)
}

// ERC20Token is a free data retrieval call binding the contract method 0x7a80760e.
//
// Solidity: function ERC20Token() constant returns(address)
func (_IdentityPromises *IdentityPromisesCaller) ERC20Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IdentityPromises.contract.Call(opts, out, "ERC20Token")
	return *ret0, err
}

// ERC20Token is a free data retrieval call binding the contract method 0x7a80760e.
//
// Solidity: function ERC20Token() constant returns(address)
func (_IdentityPromises *IdentityPromisesSession) ERC20Token() (common.Address, error) {
	return _IdentityPromises.Contract.ERC20Token(&_IdentityPromises.CallOpts)
}

// ERC20Token is a free data retrieval call binding the contract method 0x7a80760e.
//
// Solidity: function ERC20Token() constant returns(address)
func (_IdentityPromises *IdentityPromisesCallerSession) ERC20Token() (common.Address, error) {
	return _IdentityPromises.Contract.ERC20Token(&_IdentityPromises.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) constant returns(uint256)
func (_IdentityPromises *IdentityPromisesCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IdentityPromises.contract.Call(opts, out, "balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) constant returns(uint256)
func (_IdentityPromises *IdentityPromisesSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _IdentityPromises.Contract.Balances(&_IdentityPromises.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) constant returns(uint256)
func (_IdentityPromises *IdentityPromisesCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _IdentityPromises.Contract.Balances(&_IdentityPromises.CallOpts, arg0)
}

// ClearedPromises is a free data retrieval call binding the contract method 0xf6c1a23e.
//
// Solidity: function clearedPromises(address , address ) constant returns(uint256)
func (_IdentityPromises *IdentityPromisesCaller) ClearedPromises(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IdentityPromises.contract.Call(opts, out, "clearedPromises", arg0, arg1)
	return *ret0, err
}

// ClearedPromises is a free data retrieval call binding the contract method 0xf6c1a23e.
//
// Solidity: function clearedPromises(address , address ) constant returns(uint256)
func (_IdentityPromises *IdentityPromisesSession) ClearedPromises(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _IdentityPromises.Contract.ClearedPromises(&_IdentityPromises.CallOpts, arg0, arg1)
}

// ClearedPromises is a free data retrieval call binding the contract method 0xf6c1a23e.
//
// Solidity: function clearedPromises(address , address ) constant returns(uint256)
func (_IdentityPromises *IdentityPromisesCallerSession) ClearedPromises(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _IdentityPromises.Contract.ClearedPromises(&_IdentityPromises.CallOpts, arg0, arg1)
}

// CollectedFee is a free data retrieval call binding the contract method 0xe811f50a.
//
// Solidity: function collectedFee() constant returns(uint256)
func (_IdentityPromises *IdentityPromisesCaller) CollectedFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IdentityPromises.contract.Call(opts, out, "collectedFee")
	return *ret0, err
}

// CollectedFee is a free data retrieval call binding the contract method 0xe811f50a.
//
// Solidity: function collectedFee() constant returns(uint256)
func (_IdentityPromises *IdentityPromisesSession) CollectedFee() (*big.Int, error) {
	return _IdentityPromises.Contract.CollectedFee(&_IdentityPromises.CallOpts)
}

// CollectedFee is a free data retrieval call binding the contract method 0xe811f50a.
//
// Solidity: function collectedFee() constant returns(uint256)
func (_IdentityPromises *IdentityPromisesCallerSession) CollectedFee() (*big.Int, error) {
	return _IdentityPromises.Contract.CollectedFee(&_IdentityPromises.CallOpts)
}

// GetPublicKey is a free data retrieval call binding the contract method 0x857cdbb8.
//
// Solidity: function getPublicKey(address identity) constant returns(bytes32, bytes32)
func (_IdentityPromises *IdentityPromisesCaller) GetPublicKey(opts *bind.CallOpts, identity common.Address) ([32]byte, [32]byte, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _IdentityPromises.contract.Call(opts, out, "getPublicKey", identity)
	return *ret0, *ret1, err
}

// GetPublicKey is a free data retrieval call binding the contract method 0x857cdbb8.
//
// Solidity: function getPublicKey(address identity) constant returns(bytes32, bytes32)
func (_IdentityPromises *IdentityPromisesSession) GetPublicKey(identity common.Address) ([32]byte, [32]byte, error) {
	return _IdentityPromises.Contract.GetPublicKey(&_IdentityPromises.CallOpts, identity)
}

// GetPublicKey is a free data retrieval call binding the contract method 0x857cdbb8.
//
// Solidity: function getPublicKey(address identity) constant returns(bytes32, bytes32)
func (_IdentityPromises *IdentityPromisesCallerSession) GetPublicKey(identity common.Address) ([32]byte, [32]byte, error) {
	return _IdentityPromises.Contract.GetPublicKey(&_IdentityPromises.CallOpts, identity)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_IdentityPromises *IdentityPromisesCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IdentityPromises.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_IdentityPromises *IdentityPromisesSession) IsOwner() (bool, error) {
	return _IdentityPromises.Contract.IsOwner(&_IdentityPromises.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_IdentityPromises *IdentityPromisesCallerSession) IsOwner() (bool, error) {
	return _IdentityPromises.Contract.IsOwner(&_IdentityPromises.CallOpts)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address identity) constant returns(bool)
func (_IdentityPromises *IdentityPromisesCaller) IsRegistered(opts *bind.CallOpts, identity common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IdentityPromises.contract.Call(opts, out, "isRegistered", identity)
	return *ret0, err
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address identity) constant returns(bool)
func (_IdentityPromises *IdentityPromisesSession) IsRegistered(identity common.Address) (bool, error) {
	return _IdentityPromises.Contract.IsRegistered(&_IdentityPromises.CallOpts, identity)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address identity) constant returns(bool)
func (_IdentityPromises *IdentityPromisesCallerSession) IsRegistered(identity common.Address) (bool, error) {
	return _IdentityPromises.Contract.IsRegistered(&_IdentityPromises.CallOpts, identity)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_IdentityPromises *IdentityPromisesCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IdentityPromises.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_IdentityPromises *IdentityPromisesSession) Owner() (common.Address, error) {
	return _IdentityPromises.Contract.Owner(&_IdentityPromises.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_IdentityPromises *IdentityPromisesCallerSession) Owner() (common.Address, error) {
	return _IdentityPromises.Contract.Owner(&_IdentityPromises.CallOpts)
}

// RegistrationFee is a free data retrieval call binding the contract method 0x14c44e09.
//
// Solidity: function registrationFee() constant returns(uint256)
func (_IdentityPromises *IdentityPromisesCaller) RegistrationFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IdentityPromises.contract.Call(opts, out, "registrationFee")
	return *ret0, err
}

// RegistrationFee is a free data retrieval call binding the contract method 0x14c44e09.
//
// Solidity: function registrationFee() constant returns(uint256)
func (_IdentityPromises *IdentityPromisesSession) RegistrationFee() (*big.Int, error) {
	return _IdentityPromises.Contract.RegistrationFee(&_IdentityPromises.CallOpts)
}

// RegistrationFee is a free data retrieval call binding the contract method 0x14c44e09.
//
// Solidity: function registrationFee() constant returns(uint256)
func (_IdentityPromises *IdentityPromisesCallerSession) RegistrationFee() (*big.Int, error) {
	return _IdentityPromises.Contract.RegistrationFee(&_IdentityPromises.CallOpts)
}

// RegisterIdentity is a paid mutator transaction binding the contract method 0xa17fc7f3.
//
// Solidity: function RegisterIdentity(bytes32 pubKeyPart1, bytes32 pubKeyPart2, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_IdentityPromises *IdentityPromisesTransactor) RegisterIdentity(opts *bind.TransactOpts, pubKeyPart1 [32]byte, pubKeyPart2 [32]byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IdentityPromises.contract.Transact(opts, "RegisterIdentity", pubKeyPart1, pubKeyPart2, v, r, s)
}

// RegisterIdentity is a paid mutator transaction binding the contract method 0xa17fc7f3.
//
// Solidity: function RegisterIdentity(bytes32 pubKeyPart1, bytes32 pubKeyPart2, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_IdentityPromises *IdentityPromisesSession) RegisterIdentity(pubKeyPart1 [32]byte, pubKeyPart2 [32]byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IdentityPromises.Contract.RegisterIdentity(&_IdentityPromises.TransactOpts, pubKeyPart1, pubKeyPart2, v, r, s)
}

// RegisterIdentity is a paid mutator transaction binding the contract method 0xa17fc7f3.
//
// Solidity: function RegisterIdentity(bytes32 pubKeyPart1, bytes32 pubKeyPart2, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_IdentityPromises *IdentityPromisesTransactorSession) RegisterIdentity(pubKeyPart1 [32]byte, pubKeyPart2 [32]byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IdentityPromises.Contract.RegisterIdentity(&_IdentityPromises.TransactOpts, pubKeyPart1, pubKeyPart2, v, r, s)
}

// ChangeRegistrationFee is a paid mutator transaction binding the contract method 0x50050769.
//
// Solidity: function changeRegistrationFee(uint256 newFee) returns()
func (_IdentityPromises *IdentityPromisesTransactor) ChangeRegistrationFee(opts *bind.TransactOpts, newFee *big.Int) (*types.Transaction, error) {
	return _IdentityPromises.contract.Transact(opts, "changeRegistrationFee", newFee)
}

// ChangeRegistrationFee is a paid mutator transaction binding the contract method 0x50050769.
//
// Solidity: function changeRegistrationFee(uint256 newFee) returns()
func (_IdentityPromises *IdentityPromisesSession) ChangeRegistrationFee(newFee *big.Int) (*types.Transaction, error) {
	return _IdentityPromises.Contract.ChangeRegistrationFee(&_IdentityPromises.TransactOpts, newFee)
}

// ChangeRegistrationFee is a paid mutator transaction binding the contract method 0x50050769.
//
// Solidity: function changeRegistrationFee(uint256 newFee) returns()
func (_IdentityPromises *IdentityPromisesTransactorSession) ChangeRegistrationFee(newFee *big.Int) (*types.Transaction, error) {
	return _IdentityPromises.Contract.ChangeRegistrationFee(&_IdentityPromises.TransactOpts, newFee)
}

// ClearPromise is a paid mutator transaction binding the contract method 0x4fe0d332.
//
// Solidity: function clearPromise(bytes32 receiverAndSigns, bytes32 extraDataHash, uint256 seq, uint256 amount, bytes32 sender_R, bytes32 sender_S, bytes32 receiver_R, bytes32 receiver_S) returns(bool)
func (_IdentityPromises *IdentityPromisesTransactor) ClearPromise(opts *bind.TransactOpts, receiverAndSigns [32]byte, extraDataHash [32]byte, seq *big.Int, amount *big.Int, sender_R [32]byte, sender_S [32]byte, receiver_R [32]byte, receiver_S [32]byte) (*types.Transaction, error) {
	return _IdentityPromises.contract.Transact(opts, "clearPromise", receiverAndSigns, extraDataHash, seq, amount, sender_R, sender_S, receiver_R, receiver_S)
}

// ClearPromise is a paid mutator transaction binding the contract method 0x4fe0d332.
//
// Solidity: function clearPromise(bytes32 receiverAndSigns, bytes32 extraDataHash, uint256 seq, uint256 amount, bytes32 sender_R, bytes32 sender_S, bytes32 receiver_R, bytes32 receiver_S) returns(bool)
func (_IdentityPromises *IdentityPromisesSession) ClearPromise(receiverAndSigns [32]byte, extraDataHash [32]byte, seq *big.Int, amount *big.Int, sender_R [32]byte, sender_S [32]byte, receiver_R [32]byte, receiver_S [32]byte) (*types.Transaction, error) {
	return _IdentityPromises.Contract.ClearPromise(&_IdentityPromises.TransactOpts, receiverAndSigns, extraDataHash, seq, amount, sender_R, sender_S, receiver_R, receiver_S)
}

// ClearPromise is a paid mutator transaction binding the contract method 0x4fe0d332.
//
// Solidity: function clearPromise(bytes32 receiverAndSigns, bytes32 extraDataHash, uint256 seq, uint256 amount, bytes32 sender_R, bytes32 sender_S, bytes32 receiver_R, bytes32 receiver_S) returns(bool)
func (_IdentityPromises *IdentityPromisesTransactorSession) ClearPromise(receiverAndSigns [32]byte, extraDataHash [32]byte, seq *big.Int, amount *big.Int, sender_R [32]byte, sender_S [32]byte, receiver_R [32]byte, receiver_S [32]byte) (*types.Transaction, error) {
	return _IdentityPromises.Contract.ClearPromise(&_IdentityPromises.TransactOpts, receiverAndSigns, extraDataHash, seq, amount, sender_R, sender_S, receiver_R, receiver_S)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IdentityPromises *IdentityPromisesTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityPromises.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IdentityPromises *IdentityPromisesSession) RenounceOwnership() (*types.Transaction, error) {
	return _IdentityPromises.Contract.RenounceOwnership(&_IdentityPromises.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IdentityPromises *IdentityPromisesTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _IdentityPromises.Contract.RenounceOwnership(&_IdentityPromises.TransactOpts)
}

// TopUp is a paid mutator transaction binding the contract method 0xd6f7ddf9.
//
// Solidity: function topUp(address identity, uint256 amount) returns(bool)
func (_IdentityPromises *IdentityPromisesTransactor) TopUp(opts *bind.TransactOpts, identity common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IdentityPromises.contract.Transact(opts, "topUp", identity, amount)
}

// TopUp is a paid mutator transaction binding the contract method 0xd6f7ddf9.
//
// Solidity: function topUp(address identity, uint256 amount) returns(bool)
func (_IdentityPromises *IdentityPromisesSession) TopUp(identity common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IdentityPromises.Contract.TopUp(&_IdentityPromises.TransactOpts, identity, amount)
}

// TopUp is a paid mutator transaction binding the contract method 0xd6f7ddf9.
//
// Solidity: function topUp(address identity, uint256 amount) returns(bool)
func (_IdentityPromises *IdentityPromisesTransactorSession) TopUp(identity common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IdentityPromises.Contract.TopUp(&_IdentityPromises.TransactOpts, identity, amount)
}

// TransferCollectedFeeTo is a paid mutator transaction binding the contract method 0xe3252537.
//
// Solidity: function transferCollectedFeeTo(address receiver) returns(bool)
func (_IdentityPromises *IdentityPromisesTransactor) TransferCollectedFeeTo(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _IdentityPromises.contract.Transact(opts, "transferCollectedFeeTo", receiver)
}

// TransferCollectedFeeTo is a paid mutator transaction binding the contract method 0xe3252537.
//
// Solidity: function transferCollectedFeeTo(address receiver) returns(bool)
func (_IdentityPromises *IdentityPromisesSession) TransferCollectedFeeTo(receiver common.Address) (*types.Transaction, error) {
	return _IdentityPromises.Contract.TransferCollectedFeeTo(&_IdentityPromises.TransactOpts, receiver)
}

// TransferCollectedFeeTo is a paid mutator transaction binding the contract method 0xe3252537.
//
// Solidity: function transferCollectedFeeTo(address receiver) returns(bool)
func (_IdentityPromises *IdentityPromisesTransactorSession) TransferCollectedFeeTo(receiver common.Address) (*types.Transaction, error) {
	return _IdentityPromises.Contract.TransferCollectedFeeTo(&_IdentityPromises.TransactOpts, receiver)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IdentityPromises *IdentityPromisesTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IdentityPromises.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IdentityPromises *IdentityPromisesSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IdentityPromises.Contract.TransferOwnership(&_IdentityPromises.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IdentityPromises *IdentityPromisesTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IdentityPromises.Contract.TransferOwnership(&_IdentityPromises.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5ebfdfc6.
//
// Solidity: function withdraw(uint256 amount, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_IdentityPromises *IdentityPromisesTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IdentityPromises.contract.Transact(opts, "withdraw", amount, v, r, s)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5ebfdfc6.
//
// Solidity: function withdraw(uint256 amount, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_IdentityPromises *IdentityPromisesSession) Withdraw(amount *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IdentityPromises.Contract.Withdraw(&_IdentityPromises.TransactOpts, amount, v, r, s)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5ebfdfc6.
//
// Solidity: function withdraw(uint256 amount, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_IdentityPromises *IdentityPromisesTransactorSession) Withdraw(amount *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IdentityPromises.Contract.Withdraw(&_IdentityPromises.TransactOpts, amount, v, r, s)
}

// IdentityPromisesOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the IdentityPromises contract.
type IdentityPromisesOwnershipTransferredIterator struct {
	Event *IdentityPromisesOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *IdentityPromisesOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityPromisesOwnershipTransferred)
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
		it.Event = new(IdentityPromisesOwnershipTransferred)
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
func (it *IdentityPromisesOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityPromisesOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityPromisesOwnershipTransferred represents a OwnershipTransferred event raised by the IdentityPromises contract.
type IdentityPromisesOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IdentityPromises *IdentityPromisesFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IdentityPromisesOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IdentityPromises.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IdentityPromisesOwnershipTransferredIterator{contract: _IdentityPromises.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IdentityPromises *IdentityPromisesFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IdentityPromisesOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IdentityPromises.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityPromisesOwnershipTransferred)
				if err := _IdentityPromises.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// IdentityPromisesPromiseClearedIterator is returned from FilterPromiseCleared and is used to iterate over the raw logs and unpacked data for PromiseCleared events raised by the IdentityPromises contract.
type IdentityPromisesPromiseClearedIterator struct {
	Event *IdentityPromisesPromiseCleared // Event containing the contract specifics and raw log

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
func (it *IdentityPromisesPromiseClearedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityPromisesPromiseCleared)
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
		it.Event = new(IdentityPromisesPromiseCleared)
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
func (it *IdentityPromisesPromiseClearedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityPromisesPromiseClearedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityPromisesPromiseCleared represents a PromiseCleared event raised by the IdentityPromises contract.
type IdentityPromisesPromiseCleared struct {
	From   common.Address
	To     common.Address
	SeqNo  *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPromiseCleared is a free log retrieval operation binding the contract event 0x3d35f4a81f47a50c001e0e834334665e3a64680385c0a8d2fdfe2e0d20964576.
//
// Solidity: event PromiseCleared(address indexed from, address indexed to, uint256 seqNo, uint256 amount)
func (_IdentityPromises *IdentityPromisesFilterer) FilterPromiseCleared(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IdentityPromisesPromiseClearedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IdentityPromises.contract.FilterLogs(opts, "PromiseCleared", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IdentityPromisesPromiseClearedIterator{contract: _IdentityPromises.contract, event: "PromiseCleared", logs: logs, sub: sub}, nil
}

// WatchPromiseCleared is a free log subscription operation binding the contract event 0x3d35f4a81f47a50c001e0e834334665e3a64680385c0a8d2fdfe2e0d20964576.
//
// Solidity: event PromiseCleared(address indexed from, address indexed to, uint256 seqNo, uint256 amount)
func (_IdentityPromises *IdentityPromisesFilterer) WatchPromiseCleared(opts *bind.WatchOpts, sink chan<- *IdentityPromisesPromiseCleared, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IdentityPromises.contract.WatchLogs(opts, "PromiseCleared", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityPromisesPromiseCleared)
				if err := _IdentityPromises.contract.UnpackLog(event, "PromiseCleared", log); err != nil {
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

// IdentityPromisesRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the IdentityPromises contract.
type IdentityPromisesRegisteredIterator struct {
	Event *IdentityPromisesRegistered // Event containing the contract specifics and raw log

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
func (it *IdentityPromisesRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityPromisesRegistered)
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
		it.Event = new(IdentityPromisesRegistered)
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
func (it *IdentityPromisesRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityPromisesRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityPromisesRegistered represents a Registered event raised by the IdentityPromises contract.
type IdentityPromisesRegistered struct {
	Identity common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0x2d3734a8e47ac8316e500ac231c90a6e1848ca2285f40d07eaa52005e4b3a0e9.
//
// Solidity: event Registered(address indexed identity)
func (_IdentityPromises *IdentityPromisesFilterer) FilterRegistered(opts *bind.FilterOpts, identity []common.Address) (*IdentityPromisesRegisteredIterator, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _IdentityPromises.contract.FilterLogs(opts, "Registered", identityRule)
	if err != nil {
		return nil, err
	}
	return &IdentityPromisesRegisteredIterator{contract: _IdentityPromises.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0x2d3734a8e47ac8316e500ac231c90a6e1848ca2285f40d07eaa52005e4b3a0e9.
//
// Solidity: event Registered(address indexed identity)
func (_IdentityPromises *IdentityPromisesFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *IdentityPromisesRegistered, identity []common.Address) (event.Subscription, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _IdentityPromises.contract.WatchLogs(opts, "Registered", identityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityPromisesRegistered)
				if err := _IdentityPromises.contract.UnpackLog(event, "Registered", log); err != nil {
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

// IdentityPromisesToppedUpIterator is returned from FilterToppedUp and is used to iterate over the raw logs and unpacked data for ToppedUp events raised by the IdentityPromises contract.
type IdentityPromisesToppedUpIterator struct {
	Event *IdentityPromisesToppedUp // Event containing the contract specifics and raw log

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
func (it *IdentityPromisesToppedUpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityPromisesToppedUp)
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
		it.Event = new(IdentityPromisesToppedUp)
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
func (it *IdentityPromisesToppedUpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityPromisesToppedUpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityPromisesToppedUp represents a ToppedUp event raised by the IdentityPromises contract.
type IdentityPromisesToppedUp struct {
	Identity     common.Address
	Amount       *big.Int
	TotalBalance *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterToppedUp is a free log retrieval operation binding the contract event 0xbdde76a89d276b6d334c784be5c2d00d6c2219d11a2f2c80e00b85144845ab4d.
//
// Solidity: event ToppedUp(address indexed identity, uint256 amount, uint256 totalBalance)
func (_IdentityPromises *IdentityPromisesFilterer) FilterToppedUp(opts *bind.FilterOpts, identity []common.Address) (*IdentityPromisesToppedUpIterator, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _IdentityPromises.contract.FilterLogs(opts, "ToppedUp", identityRule)
	if err != nil {
		return nil, err
	}
	return &IdentityPromisesToppedUpIterator{contract: _IdentityPromises.contract, event: "ToppedUp", logs: logs, sub: sub}, nil
}

// WatchToppedUp is a free log subscription operation binding the contract event 0xbdde76a89d276b6d334c784be5c2d00d6c2219d11a2f2c80e00b85144845ab4d.
//
// Solidity: event ToppedUp(address indexed identity, uint256 amount, uint256 totalBalance)
func (_IdentityPromises *IdentityPromisesFilterer) WatchToppedUp(opts *bind.WatchOpts, sink chan<- *IdentityPromisesToppedUp, identity []common.Address) (event.Subscription, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _IdentityPromises.contract.WatchLogs(opts, "ToppedUp", identityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityPromisesToppedUp)
				if err := _IdentityPromises.contract.UnpackLog(event, "ToppedUp", log); err != nil {
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

// IdentityPromisesWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the IdentityPromises contract.
type IdentityPromisesWithdrawnIterator struct {
	Event *IdentityPromisesWithdrawn // Event containing the contract specifics and raw log

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
func (it *IdentityPromisesWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityPromisesWithdrawn)
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
		it.Event = new(IdentityPromisesWithdrawn)
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
func (it *IdentityPromisesWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityPromisesWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityPromisesWithdrawn represents a Withdrawn event raised by the IdentityPromises contract.
type IdentityPromisesWithdrawn struct {
	Identity     common.Address
	Amount       *big.Int
	TotalBalance *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: event Withdrawn(address indexed identity, uint256 amount, uint256 totalBalance)
func (_IdentityPromises *IdentityPromisesFilterer) FilterWithdrawn(opts *bind.FilterOpts, identity []common.Address) (*IdentityPromisesWithdrawnIterator, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _IdentityPromises.contract.FilterLogs(opts, "Withdrawn", identityRule)
	if err != nil {
		return nil, err
	}
	return &IdentityPromisesWithdrawnIterator{contract: _IdentityPromises.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: event Withdrawn(address indexed identity, uint256 amount, uint256 totalBalance)
func (_IdentityPromises *IdentityPromisesFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *IdentityPromisesWithdrawn, identity []common.Address) (event.Subscription, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _IdentityPromises.contract.WatchLogs(opts, "Withdrawn", identityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityPromisesWithdrawn)
				if err := _IdentityPromises.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
