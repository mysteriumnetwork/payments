// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rewarder

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

// RewarderABI is the input ABI used to generate the binding from.
const RewarderABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_custody\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalDropped\",\"type\":\"uint256\"}],\"name\":\"Airdrop\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalUnclaimed\",\"type\":\"uint256\"}],\"name\":\"ClaimedChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalAmount\",\"type\":\"uint256\"}],\"name\":\"RootUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_beneficiaries\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_totalEarnings\",\"type\":\"uint256[]\"}],\"name\":\"airdrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_totalEarned\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_proof\",\"type\":\"bytes32[]\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"claimRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"custody\",\"outputs\":[{\"internalType\":\"contractCustody\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_totalEarned\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_proof\",\"type\":\"bytes32[]\"}],\"name\":\"isValidProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRootBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_erc20\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"recoverTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalClaimed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalPayoutsFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_claimRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalReward\",\"type\":\"uint256\"}],\"name\":\"updateRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Rewarder is an auto generated Go binding around an Ethereum contract.
type Rewarder struct {
	RewarderCaller     // Read-only binding to the contract
	RewarderTransactor // Write-only binding to the contract
	RewarderFilterer   // Log filterer for contract events
}

// RewarderCaller is an auto generated read-only Go binding around an Ethereum contract.
type RewarderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewarderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RewarderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewarderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RewarderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewarderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RewarderSession struct {
	Contract     *Rewarder         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RewarderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RewarderCallerSession struct {
	Contract *RewarderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RewarderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RewarderTransactorSession struct {
	Contract     *RewarderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RewarderRaw is an auto generated low-level Go binding around an Ethereum contract.
type RewarderRaw struct {
	Contract *Rewarder // Generic contract binding to access the raw methods on
}

// RewarderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RewarderCallerRaw struct {
	Contract *RewarderCaller // Generic read-only contract binding to access the raw methods on
}

// RewarderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RewarderTransactorRaw struct {
	Contract *RewarderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRewarder creates a new instance of Rewarder, bound to a specific deployed contract.
func NewRewarder(address common.Address, backend bind.ContractBackend) (*Rewarder, error) {
	contract, err := bindRewarder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rewarder{RewarderCaller: RewarderCaller{contract: contract}, RewarderTransactor: RewarderTransactor{contract: contract}, RewarderFilterer: RewarderFilterer{contract: contract}}, nil
}

// NewRewarderCaller creates a new read-only instance of Rewarder, bound to a specific deployed contract.
func NewRewarderCaller(address common.Address, caller bind.ContractCaller) (*RewarderCaller, error) {
	contract, err := bindRewarder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RewarderCaller{contract: contract}, nil
}

// NewRewarderTransactor creates a new write-only instance of Rewarder, bound to a specific deployed contract.
func NewRewarderTransactor(address common.Address, transactor bind.ContractTransactor) (*RewarderTransactor, error) {
	contract, err := bindRewarder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RewarderTransactor{contract: contract}, nil
}

// NewRewarderFilterer creates a new log filterer instance of Rewarder, bound to a specific deployed contract.
func NewRewarderFilterer(address common.Address, filterer bind.ContractFilterer) (*RewarderFilterer, error) {
	contract, err := bindRewarder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RewarderFilterer{contract: contract}, nil
}

// bindRewarder binds a generic wrapper to an already deployed contract.
func bindRewarder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RewarderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rewarder *RewarderRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Rewarder.Contract.RewarderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rewarder *RewarderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewarder.Contract.RewarderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rewarder *RewarderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rewarder.Contract.RewarderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rewarder *RewarderCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Rewarder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rewarder *RewarderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewarder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rewarder *RewarderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rewarder.Contract.contract.Transact(opts, method, params...)
}

// ClaimRoots is a free data retrieval call binding the contract method 0x7558b6ef.
//
// Solidity: function claimRoots(uint256 ) view returns(bytes32)
func (_Rewarder *RewarderCaller) ClaimRoots(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Rewarder.contract.Call(opts, out, "claimRoots", arg0)
	return *ret0, err
}

// ClaimRoots is a free data retrieval call binding the contract method 0x7558b6ef.
//
// Solidity: function claimRoots(uint256 ) view returns(bytes32)
func (_Rewarder *RewarderSession) ClaimRoots(arg0 *big.Int) ([32]byte, error) {
	return _Rewarder.Contract.ClaimRoots(&_Rewarder.CallOpts, arg0)
}

// ClaimRoots is a free data retrieval call binding the contract method 0x7558b6ef.
//
// Solidity: function claimRoots(uint256 ) view returns(bytes32)
func (_Rewarder *RewarderCallerSession) ClaimRoots(arg0 *big.Int) ([32]byte, error) {
	return _Rewarder.Contract.ClaimRoots(&_Rewarder.CallOpts, arg0)
}

// Custody is a free data retrieval call binding the contract method 0xdda79b75.
//
// Solidity: function custody() view returns(address)
func (_Rewarder *RewarderCaller) Custody(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Rewarder.contract.Call(opts, out, "custody")
	return *ret0, err
}

// Custody is a free data retrieval call binding the contract method 0xdda79b75.
//
// Solidity: function custody() view returns(address)
func (_Rewarder *RewarderSession) Custody() (common.Address, error) {
	return _Rewarder.Contract.Custody(&_Rewarder.CallOpts)
}

// Custody is a free data retrieval call binding the contract method 0xdda79b75.
//
// Solidity: function custody() view returns(address)
func (_Rewarder *RewarderCallerSession) Custody() (common.Address, error) {
	return _Rewarder.Contract.Custody(&_Rewarder.CallOpts)
}

// IsValidProof is a free data retrieval call binding the contract method 0xca5a2780.
//
// Solidity: function isValidProof(address _recipient, uint256 _totalEarned, uint256 _blockNumber, bytes32[] _proof) view returns(bool)
func (_Rewarder *RewarderCaller) IsValidProof(opts *bind.CallOpts, _recipient common.Address, _totalEarned *big.Int, _blockNumber *big.Int, _proof [][32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Rewarder.contract.Call(opts, out, "isValidProof", _recipient, _totalEarned, _blockNumber, _proof)
	return *ret0, err
}

// IsValidProof is a free data retrieval call binding the contract method 0xca5a2780.
//
// Solidity: function isValidProof(address _recipient, uint256 _totalEarned, uint256 _blockNumber, bytes32[] _proof) view returns(bool)
func (_Rewarder *RewarderSession) IsValidProof(_recipient common.Address, _totalEarned *big.Int, _blockNumber *big.Int, _proof [][32]byte) (bool, error) {
	return _Rewarder.Contract.IsValidProof(&_Rewarder.CallOpts, _recipient, _totalEarned, _blockNumber, _proof)
}

// IsValidProof is a free data retrieval call binding the contract method 0xca5a2780.
//
// Solidity: function isValidProof(address _recipient, uint256 _totalEarned, uint256 _blockNumber, bytes32[] _proof) view returns(bool)
func (_Rewarder *RewarderCallerSession) IsValidProof(_recipient common.Address, _totalEarned *big.Int, _blockNumber *big.Int, _proof [][32]byte) (bool, error) {
	return _Rewarder.Contract.IsValidProof(&_Rewarder.CallOpts, _recipient, _totalEarned, _blockNumber, _proof)
}

// LastRootBlock is a free data retrieval call binding the contract method 0x8e59d422.
//
// Solidity: function lastRootBlock() view returns(uint256)
func (_Rewarder *RewarderCaller) LastRootBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rewarder.contract.Call(opts, out, "lastRootBlock")
	return *ret0, err
}

// LastRootBlock is a free data retrieval call binding the contract method 0x8e59d422.
//
// Solidity: function lastRootBlock() view returns(uint256)
func (_Rewarder *RewarderSession) LastRootBlock() (*big.Int, error) {
	return _Rewarder.Contract.LastRootBlock(&_Rewarder.CallOpts)
}

// LastRootBlock is a free data retrieval call binding the contract method 0x8e59d422.
//
// Solidity: function lastRootBlock() view returns(uint256)
func (_Rewarder *RewarderCallerSession) LastRootBlock() (*big.Int, error) {
	return _Rewarder.Contract.LastRootBlock(&_Rewarder.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rewarder *RewarderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Rewarder.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rewarder *RewarderSession) Owner() (common.Address, error) {
	return _Rewarder.Contract.Owner(&_Rewarder.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rewarder *RewarderCallerSession) Owner() (common.Address, error) {
	return _Rewarder.Contract.Owner(&_Rewarder.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Rewarder *RewarderCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Rewarder.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Rewarder *RewarderSession) Token() (common.Address, error) {
	return _Rewarder.Contract.Token(&_Rewarder.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Rewarder *RewarderCallerSession) Token() (common.Address, error) {
	return _Rewarder.Contract.Token(&_Rewarder.CallOpts)
}

// TotalClaimed is a free data retrieval call binding the contract method 0xd54ad2a1.
//
// Solidity: function totalClaimed() view returns(uint256)
func (_Rewarder *RewarderCaller) TotalClaimed(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rewarder.contract.Call(opts, out, "totalClaimed")
	return *ret0, err
}

// TotalClaimed is a free data retrieval call binding the contract method 0xd54ad2a1.
//
// Solidity: function totalClaimed() view returns(uint256)
func (_Rewarder *RewarderSession) TotalClaimed() (*big.Int, error) {
	return _Rewarder.Contract.TotalClaimed(&_Rewarder.CallOpts)
}

// TotalClaimed is a free data retrieval call binding the contract method 0xd54ad2a1.
//
// Solidity: function totalClaimed() view returns(uint256)
func (_Rewarder *RewarderCallerSession) TotalClaimed() (*big.Int, error) {
	return _Rewarder.Contract.TotalClaimed(&_Rewarder.CallOpts)
}

// TotalPayoutsFor is a free data retrieval call binding the contract method 0x4d676fa7.
//
// Solidity: function totalPayoutsFor(address ) view returns(uint256)
func (_Rewarder *RewarderCaller) TotalPayoutsFor(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rewarder.contract.Call(opts, out, "totalPayoutsFor", arg0)
	return *ret0, err
}

// TotalPayoutsFor is a free data retrieval call binding the contract method 0x4d676fa7.
//
// Solidity: function totalPayoutsFor(address ) view returns(uint256)
func (_Rewarder *RewarderSession) TotalPayoutsFor(arg0 common.Address) (*big.Int, error) {
	return _Rewarder.Contract.TotalPayoutsFor(&_Rewarder.CallOpts, arg0)
}

// TotalPayoutsFor is a free data retrieval call binding the contract method 0x4d676fa7.
//
// Solidity: function totalPayoutsFor(address ) view returns(uint256)
func (_Rewarder *RewarderCallerSession) TotalPayoutsFor(arg0 common.Address) (*big.Int, error) {
	return _Rewarder.Contract.TotalPayoutsFor(&_Rewarder.CallOpts, arg0)
}

// Airdrop is a paid mutator transaction binding the contract method 0x67243482.
//
// Solidity: function airdrop(address[] _beneficiaries, uint256[] _totalEarnings) returns()
func (_Rewarder *RewarderTransactor) Airdrop(opts *bind.TransactOpts, _beneficiaries []common.Address, _totalEarnings []*big.Int) (*types.Transaction, error) {
	return _Rewarder.contract.Transact(opts, "airdrop", _beneficiaries, _totalEarnings)
}

// Airdrop is a paid mutator transaction binding the contract method 0x67243482.
//
// Solidity: function airdrop(address[] _beneficiaries, uint256[] _totalEarnings) returns()
func (_Rewarder *RewarderSession) Airdrop(_beneficiaries []common.Address, _totalEarnings []*big.Int) (*types.Transaction, error) {
	return _Rewarder.Contract.Airdrop(&_Rewarder.TransactOpts, _beneficiaries, _totalEarnings)
}

// Airdrop is a paid mutator transaction binding the contract method 0x67243482.
//
// Solidity: function airdrop(address[] _beneficiaries, uint256[] _totalEarnings) returns()
func (_Rewarder *RewarderTransactorSession) Airdrop(_beneficiaries []common.Address, _totalEarnings []*big.Int) (*types.Transaction, error) {
	return _Rewarder.Contract.Airdrop(&_Rewarder.TransactOpts, _beneficiaries, _totalEarnings)
}

// Claim is a paid mutator transaction binding the contract method 0x172bd6de.
//
// Solidity: function claim(address _recipient, uint256 _totalEarned, uint256 _blockNumber, bytes32[] _proof) returns()
func (_Rewarder *RewarderTransactor) Claim(opts *bind.TransactOpts, _recipient common.Address, _totalEarned *big.Int, _blockNumber *big.Int, _proof [][32]byte) (*types.Transaction, error) {
	return _Rewarder.contract.Transact(opts, "claim", _recipient, _totalEarned, _blockNumber, _proof)
}

// Claim is a paid mutator transaction binding the contract method 0x172bd6de.
//
// Solidity: function claim(address _recipient, uint256 _totalEarned, uint256 _blockNumber, bytes32[] _proof) returns()
func (_Rewarder *RewarderSession) Claim(_recipient common.Address, _totalEarned *big.Int, _blockNumber *big.Int, _proof [][32]byte) (*types.Transaction, error) {
	return _Rewarder.Contract.Claim(&_Rewarder.TransactOpts, _recipient, _totalEarned, _blockNumber, _proof)
}

// Claim is a paid mutator transaction binding the contract method 0x172bd6de.
//
// Solidity: function claim(address _recipient, uint256 _totalEarned, uint256 _blockNumber, bytes32[] _proof) returns()
func (_Rewarder *RewarderTransactorSession) Claim(_recipient common.Address, _totalEarned *big.Int, _blockNumber *big.Int, _proof [][32]byte) (*types.Transaction, error) {
	return _Rewarder.Contract.Claim(&_Rewarder.TransactOpts, _recipient, _totalEarned, _blockNumber, _proof)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0x056097ac.
//
// Solidity: function recoverTokens(address _erc20, address _to) returns()
func (_Rewarder *RewarderTransactor) RecoverTokens(opts *bind.TransactOpts, _erc20 common.Address, _to common.Address) (*types.Transaction, error) {
	return _Rewarder.contract.Transact(opts, "recoverTokens", _erc20, _to)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0x056097ac.
//
// Solidity: function recoverTokens(address _erc20, address _to) returns()
func (_Rewarder *RewarderSession) RecoverTokens(_erc20 common.Address, _to common.Address) (*types.Transaction, error) {
	return _Rewarder.Contract.RecoverTokens(&_Rewarder.TransactOpts, _erc20, _to)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0x056097ac.
//
// Solidity: function recoverTokens(address _erc20, address _to) returns()
func (_Rewarder *RewarderTransactorSession) RecoverTokens(_erc20 common.Address, _to common.Address) (*types.Transaction, error) {
	return _Rewarder.Contract.RecoverTokens(&_Rewarder.TransactOpts, _erc20, _to)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Rewarder *RewarderTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewarder.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Rewarder *RewarderSession) RenounceOwnership() (*types.Transaction, error) {
	return _Rewarder.Contract.RenounceOwnership(&_Rewarder.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Rewarder *RewarderTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Rewarder.Contract.RenounceOwnership(&_Rewarder.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rewarder *RewarderTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Rewarder.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rewarder *RewarderSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Rewarder.Contract.TransferOwnership(&_Rewarder.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rewarder *RewarderTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Rewarder.Contract.TransferOwnership(&_Rewarder.TransactOpts, newOwner)
}

// UpdateRoot is a paid mutator transaction binding the contract method 0x12b6af24.
//
// Solidity: function updateRoot(bytes32 _claimRoot, uint256 _blockNumber, uint256 _totalReward) returns()
func (_Rewarder *RewarderTransactor) UpdateRoot(opts *bind.TransactOpts, _claimRoot [32]byte, _blockNumber *big.Int, _totalReward *big.Int) (*types.Transaction, error) {
	return _Rewarder.contract.Transact(opts, "updateRoot", _claimRoot, _blockNumber, _totalReward)
}

// UpdateRoot is a paid mutator transaction binding the contract method 0x12b6af24.
//
// Solidity: function updateRoot(bytes32 _claimRoot, uint256 _blockNumber, uint256 _totalReward) returns()
func (_Rewarder *RewarderSession) UpdateRoot(_claimRoot [32]byte, _blockNumber *big.Int, _totalReward *big.Int) (*types.Transaction, error) {
	return _Rewarder.Contract.UpdateRoot(&_Rewarder.TransactOpts, _claimRoot, _blockNumber, _totalReward)
}

// UpdateRoot is a paid mutator transaction binding the contract method 0x12b6af24.
//
// Solidity: function updateRoot(bytes32 _claimRoot, uint256 _blockNumber, uint256 _totalReward) returns()
func (_Rewarder *RewarderTransactorSession) UpdateRoot(_claimRoot [32]byte, _blockNumber *big.Int, _totalReward *big.Int) (*types.Transaction, error) {
	return _Rewarder.Contract.UpdateRoot(&_Rewarder.TransactOpts, _claimRoot, _blockNumber, _totalReward)
}

// RewarderAirdropIterator is returned from FilterAirdrop and is used to iterate over the raw logs and unpacked data for Airdrop events raised by the Rewarder contract.
type RewarderAirdropIterator struct {
	Event *RewarderAirdrop // Event containing the contract specifics and raw log

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
func (it *RewarderAirdropIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewarderAirdrop)
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
		it.Event = new(RewarderAirdrop)
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
func (it *RewarderAirdropIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewarderAirdropIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewarderAirdrop represents a Airdrop event raised by the Rewarder contract.
type RewarderAirdrop struct {
	TotalDropped *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAirdrop is a free log retrieval operation binding the contract event 0xd0ecdf4854f39daba34ba0e2c1ed0132a7023a5434bffc90b24f3335fb90e5e3.
//
// Solidity: event Airdrop(uint256 totalDropped)
func (_Rewarder *RewarderFilterer) FilterAirdrop(opts *bind.FilterOpts) (*RewarderAirdropIterator, error) {

	logs, sub, err := _Rewarder.contract.FilterLogs(opts, "Airdrop")
	if err != nil {
		return nil, err
	}
	return &RewarderAirdropIterator{contract: _Rewarder.contract, event: "Airdrop", logs: logs, sub: sub}, nil
}

// WatchAirdrop is a free log subscription operation binding the contract event 0xd0ecdf4854f39daba34ba0e2c1ed0132a7023a5434bffc90b24f3335fb90e5e3.
//
// Solidity: event Airdrop(uint256 totalDropped)
func (_Rewarder *RewarderFilterer) WatchAirdrop(opts *bind.WatchOpts, sink chan<- *RewarderAirdrop) (event.Subscription, error) {

	logs, sub, err := _Rewarder.contract.WatchLogs(opts, "Airdrop")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewarderAirdrop)
				if err := _Rewarder.contract.UnpackLog(event, "Airdrop", log); err != nil {
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

// ParseAirdrop is a log parse operation binding the contract event 0xd0ecdf4854f39daba34ba0e2c1ed0132a7023a5434bffc90b24f3335fb90e5e3.
//
// Solidity: event Airdrop(uint256 totalDropped)
func (_Rewarder *RewarderFilterer) ParseAirdrop(log types.Log) (*RewarderAirdrop, error) {
	event := new(RewarderAirdrop)
	if err := _Rewarder.contract.UnpackLog(event, "Airdrop", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RewarderClaimedChangedIterator is returned from FilterClaimedChanged and is used to iterate over the raw logs and unpacked data for ClaimedChanged events raised by the Rewarder contract.
type RewarderClaimedChangedIterator struct {
	Event *RewarderClaimedChanged // Event containing the contract specifics and raw log

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
func (it *RewarderClaimedChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewarderClaimedChanged)
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
		it.Event = new(RewarderClaimedChanged)
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
func (it *RewarderClaimedChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewarderClaimedChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewarderClaimedChanged represents a ClaimedChanged event raised by the Rewarder contract.
type RewarderClaimedChanged struct {
	TotalUnclaimed *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterClaimedChanged is a free log retrieval operation binding the contract event 0xc58f861f50ca7e2541250d781773a4614f235d08e72ecc0d2f25028b70a1b665.
//
// Solidity: event ClaimedChanged(uint256 totalUnclaimed)
func (_Rewarder *RewarderFilterer) FilterClaimedChanged(opts *bind.FilterOpts) (*RewarderClaimedChangedIterator, error) {

	logs, sub, err := _Rewarder.contract.FilterLogs(opts, "ClaimedChanged")
	if err != nil {
		return nil, err
	}
	return &RewarderClaimedChangedIterator{contract: _Rewarder.contract, event: "ClaimedChanged", logs: logs, sub: sub}, nil
}

// WatchClaimedChanged is a free log subscription operation binding the contract event 0xc58f861f50ca7e2541250d781773a4614f235d08e72ecc0d2f25028b70a1b665.
//
// Solidity: event ClaimedChanged(uint256 totalUnclaimed)
func (_Rewarder *RewarderFilterer) WatchClaimedChanged(opts *bind.WatchOpts, sink chan<- *RewarderClaimedChanged) (event.Subscription, error) {

	logs, sub, err := _Rewarder.contract.WatchLogs(opts, "ClaimedChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewarderClaimedChanged)
				if err := _Rewarder.contract.UnpackLog(event, "ClaimedChanged", log); err != nil {
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

// ParseClaimedChanged is a log parse operation binding the contract event 0xc58f861f50ca7e2541250d781773a4614f235d08e72ecc0d2f25028b70a1b665.
//
// Solidity: event ClaimedChanged(uint256 totalUnclaimed)
func (_Rewarder *RewarderFilterer) ParseClaimedChanged(log types.Log) (*RewarderClaimedChanged, error) {
	event := new(RewarderClaimedChanged)
	if err := _Rewarder.contract.UnpackLog(event, "ClaimedChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RewarderOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Rewarder contract.
type RewarderOwnershipTransferredIterator struct {
	Event *RewarderOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RewarderOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewarderOwnershipTransferred)
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
		it.Event = new(RewarderOwnershipTransferred)
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
func (it *RewarderOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewarderOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewarderOwnershipTransferred represents a OwnershipTransferred event raised by the Rewarder contract.
type RewarderOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Rewarder *RewarderFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RewarderOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Rewarder.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RewarderOwnershipTransferredIterator{contract: _Rewarder.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Rewarder *RewarderFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RewarderOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Rewarder.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewarderOwnershipTransferred)
				if err := _Rewarder.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Rewarder *RewarderFilterer) ParseOwnershipTransferred(log types.Log) (*RewarderOwnershipTransferred, error) {
	event := new(RewarderOwnershipTransferred)
	if err := _Rewarder.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RewarderRootUpdatedIterator is returned from FilterRootUpdated and is used to iterate over the raw logs and unpacked data for RootUpdated events raised by the Rewarder contract.
type RewarderRootUpdatedIterator struct {
	Event *RewarderRootUpdated // Event containing the contract specifics and raw log

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
func (it *RewarderRootUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewarderRootUpdated)
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
		it.Event = new(RewarderRootUpdated)
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
func (it *RewarderRootUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewarderRootUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewarderRootUpdated represents a RootUpdated event raised by the Rewarder contract.
type RewarderRootUpdated struct {
	Root        [32]byte
	BlockNumber *big.Int
	TotalAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRootUpdated is a free log retrieval operation binding the contract event 0x156798a72d63b37f86ff1ecc41eec4f30e3c7b325c8a410c2671f2e7fc0c30f3.
//
// Solidity: event RootUpdated(bytes32 root, uint256 blockNumber, uint256 _totalAmount)
func (_Rewarder *RewarderFilterer) FilterRootUpdated(opts *bind.FilterOpts) (*RewarderRootUpdatedIterator, error) {

	logs, sub, err := _Rewarder.contract.FilterLogs(opts, "RootUpdated")
	if err != nil {
		return nil, err
	}
	return &RewarderRootUpdatedIterator{contract: _Rewarder.contract, event: "RootUpdated", logs: logs, sub: sub}, nil
}

// WatchRootUpdated is a free log subscription operation binding the contract event 0x156798a72d63b37f86ff1ecc41eec4f30e3c7b325c8a410c2671f2e7fc0c30f3.
//
// Solidity: event RootUpdated(bytes32 root, uint256 blockNumber, uint256 _totalAmount)
func (_Rewarder *RewarderFilterer) WatchRootUpdated(opts *bind.WatchOpts, sink chan<- *RewarderRootUpdated) (event.Subscription, error) {

	logs, sub, err := _Rewarder.contract.WatchLogs(opts, "RootUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewarderRootUpdated)
				if err := _Rewarder.contract.UnpackLog(event, "RootUpdated", log); err != nil {
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

// ParseRootUpdated is a log parse operation binding the contract event 0x156798a72d63b37f86ff1ecc41eec4f30e3c7b325c8a410c2671f2e7fc0c30f3.
//
// Solidity: event RootUpdated(bytes32 root, uint256 blockNumber, uint256 _totalAmount)
func (_Rewarder *RewarderFilterer) ParseRootUpdated(log types.Log) (*RewarderRootUpdated, error) {
	event := new(RewarderRootUpdated)
	if err := _Rewarder.contract.UnpackLog(event, "RootUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}
