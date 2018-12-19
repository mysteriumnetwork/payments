// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package utils

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

// TestUtilsContractABI is the input ABI used to generate the binding from.
const TestUtilsContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"param\",\"type\":\"bytes32\"}],\"name\":\"GiveMeData\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TestUtilsContractBin is the compiled bytecode used for deploying new contracts.
const TestUtilsContractBin = `0x608060405234801561001057600080fd5b50610111806100206000396000f300608060405260043610603e5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416635aaa775e81146043575b600080fd5b348015604e57600080fd5b5060586004356091565b6040805173ffffffffffffffffffffffffffffffffffffffff909416845260ff9283166020850152911682820152519081900360600190f35b6000806000609d8460aa565b9250925092509193909250565b907f0100000000000000000000000000000000000000000000000000000000000000600a83901a810281900491600b84901a820291909104905600a165627a7a72305820a7f9663a7c161da55d464d8854af862c7fe9201e52dee860357cd451b74080fc0029`

// DeployTestUtilsContract deploys a new Ethereum contract, binding an instance of TestUtilsContract to it.
func DeployTestUtilsContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestUtilsContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestUtilsContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TestUtilsContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestUtilsContract{TestUtilsContractCaller: TestUtilsContractCaller{contract: contract}, TestUtilsContractTransactor: TestUtilsContractTransactor{contract: contract}, TestUtilsContractFilterer: TestUtilsContractFilterer{contract: contract}}, nil
}

// TestUtilsContract is an auto generated Go binding around an Ethereum contract.
type TestUtilsContract struct {
	TestUtilsContractCaller     // Read-only binding to the contract
	TestUtilsContractTransactor // Write-only binding to the contract
	TestUtilsContractFilterer   // Log filterer for contract events
}

// TestUtilsContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestUtilsContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestUtilsContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestUtilsContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestUtilsContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestUtilsContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestUtilsContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestUtilsContractSession struct {
	Contract     *TestUtilsContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TestUtilsContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestUtilsContractCallerSession struct {
	Contract *TestUtilsContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// TestUtilsContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestUtilsContractTransactorSession struct {
	Contract     *TestUtilsContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// TestUtilsContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestUtilsContractRaw struct {
	Contract *TestUtilsContract // Generic contract binding to access the raw methods on
}

// TestUtilsContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestUtilsContractCallerRaw struct {
	Contract *TestUtilsContractCaller // Generic read-only contract binding to access the raw methods on
}

// TestUtilsContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestUtilsContractTransactorRaw struct {
	Contract *TestUtilsContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestUtilsContract creates a new instance of TestUtilsContract, bound to a specific deployed contract.
func NewTestUtilsContract(address common.Address, backend bind.ContractBackend) (*TestUtilsContract, error) {
	contract, err := bindTestUtilsContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestUtilsContract{TestUtilsContractCaller: TestUtilsContractCaller{contract: contract}, TestUtilsContractTransactor: TestUtilsContractTransactor{contract: contract}, TestUtilsContractFilterer: TestUtilsContractFilterer{contract: contract}}, nil
}

// NewTestUtilsContractCaller creates a new read-only instance of TestUtilsContract, bound to a specific deployed contract.
func NewTestUtilsContractCaller(address common.Address, caller bind.ContractCaller) (*TestUtilsContractCaller, error) {
	contract, err := bindTestUtilsContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestUtilsContractCaller{contract: contract}, nil
}

// NewTestUtilsContractTransactor creates a new write-only instance of TestUtilsContract, bound to a specific deployed contract.
func NewTestUtilsContractTransactor(address common.Address, transactor bind.ContractTransactor) (*TestUtilsContractTransactor, error) {
	contract, err := bindTestUtilsContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestUtilsContractTransactor{contract: contract}, nil
}

// NewTestUtilsContractFilterer creates a new log filterer instance of TestUtilsContract, bound to a specific deployed contract.
func NewTestUtilsContractFilterer(address common.Address, filterer bind.ContractFilterer) (*TestUtilsContractFilterer, error) {
	contract, err := bindTestUtilsContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestUtilsContractFilterer{contract: contract}, nil
}

// bindTestUtilsContract binds a generic wrapper to an already deployed contract.
func bindTestUtilsContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestUtilsContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestUtilsContract *TestUtilsContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TestUtilsContract.Contract.TestUtilsContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestUtilsContract *TestUtilsContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestUtilsContract.Contract.TestUtilsContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestUtilsContract *TestUtilsContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestUtilsContract.Contract.TestUtilsContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestUtilsContract *TestUtilsContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TestUtilsContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestUtilsContract *TestUtilsContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestUtilsContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestUtilsContract *TestUtilsContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestUtilsContract.Contract.contract.Transact(opts, method, params...)
}

// GiveMeData is a free data retrieval call binding the contract method 0x5aaa775e.
//
// Solidity: function GiveMeData(param bytes32) constant returns(address, uint8, uint8)
func (_TestUtilsContract *TestUtilsContractCaller) GiveMeData(opts *bind.CallOpts, param [32]byte) (common.Address, uint8, uint8, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(uint8)
		ret2 = new(uint8)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _TestUtilsContract.contract.Call(opts, out, "GiveMeData", param)
	return *ret0, *ret1, *ret2, err
}

// GiveMeData is a free data retrieval call binding the contract method 0x5aaa775e.
//
// Solidity: function GiveMeData(param bytes32) constant returns(address, uint8, uint8)
func (_TestUtilsContract *TestUtilsContractSession) GiveMeData(param [32]byte) (common.Address, uint8, uint8, error) {
	return _TestUtilsContract.Contract.GiveMeData(&_TestUtilsContract.CallOpts, param)
}

// GiveMeData is a free data retrieval call binding the contract method 0x5aaa775e.
//
// Solidity: function GiveMeData(param bytes32) constant returns(address, uint8, uint8)
func (_TestUtilsContract *TestUtilsContractCallerSession) GiveMeData(param [32]byte) (common.Address, uint8, uint8, error) {
	return _TestUtilsContract.Contract.GiveMeData(&_TestUtilsContract.CallOpts, param)
}

// UtilsABI is the input ABI used to generate the binding from.
const UtilsABI = "[]"

// UtilsBin is the compiled bytecode used for deploying new contracts.
const UtilsBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820c232048537555b7e051e77aa2a4d096e5c271d87135a1f904651e9c48b5f6bbf0029`

// DeployUtils deploys a new Ethereum contract, binding an instance of Utils to it.
func DeployUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Utils, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Utils{UtilsCaller: UtilsCaller{contract: contract}, UtilsTransactor: UtilsTransactor{contract: contract}, UtilsFilterer: UtilsFilterer{contract: contract}}, nil
}

// Utils is an auto generated Go binding around an Ethereum contract.
type Utils struct {
	UtilsCaller     // Read-only binding to the contract
	UtilsTransactor // Write-only binding to the contract
	UtilsFilterer   // Log filterer for contract events
}

// UtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type UtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UtilsSession struct {
	Contract     *Utils            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UtilsCallerSession struct {
	Contract *UtilsCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UtilsTransactorSession struct {
	Contract     *UtilsTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type UtilsRaw struct {
	Contract *Utils // Generic contract binding to access the raw methods on
}

// UtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UtilsCallerRaw struct {
	Contract *UtilsCaller // Generic read-only contract binding to access the raw methods on
}

// UtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UtilsTransactorRaw struct {
	Contract *UtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUtils creates a new instance of Utils, bound to a specific deployed contract.
func NewUtils(address common.Address, backend bind.ContractBackend) (*Utils, error) {
	contract, err := bindUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Utils{UtilsCaller: UtilsCaller{contract: contract}, UtilsTransactor: UtilsTransactor{contract: contract}, UtilsFilterer: UtilsFilterer{contract: contract}}, nil
}

// NewUtilsCaller creates a new read-only instance of Utils, bound to a specific deployed contract.
func NewUtilsCaller(address common.Address, caller bind.ContractCaller) (*UtilsCaller, error) {
	contract, err := bindUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UtilsCaller{contract: contract}, nil
}

// NewUtilsTransactor creates a new write-only instance of Utils, bound to a specific deployed contract.
func NewUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*UtilsTransactor, error) {
	contract, err := bindUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UtilsTransactor{contract: contract}, nil
}

// NewUtilsFilterer creates a new log filterer instance of Utils, bound to a specific deployed contract.
func NewUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*UtilsFilterer, error) {
	contract, err := bindUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UtilsFilterer{contract: contract}, nil
}

// bindUtils binds a generic wrapper to an already deployed contract.
func bindUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Utils *UtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Utils.Contract.UtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Utils *UtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Utils.Contract.UtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Utils *UtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Utils.Contract.UtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Utils *UtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Utils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Utils *UtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Utils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Utils *UtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Utils.Contract.contract.Transact(opts, method, params...)
}
