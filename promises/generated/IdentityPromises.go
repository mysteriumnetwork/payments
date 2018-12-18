// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package generated

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

// ERC20ABI is the input ABI used to generate the binding from.
const ERC20ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// ERC20Bin is the compiled bytecode used for deploying new contracts.
const ERC20Bin = `0x`

// DeployERC20 deploys a new Ethereum contract, binding an instance of ERC20 to it.
func DeployERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_ERC20 *ERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_ERC20 *ERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20 *ERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct {
	Event *ERC20Approval // Event containing the contract specifics and raw log

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
func (it *ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
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
		it.Event = new(ERC20Approval)
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
func (it *ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ApprovalIterator{contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

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
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
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
		it.Event = new(ERC20Transfer)
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
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ERC20AwareABI is the input ABI used to generate the binding from.
const ERC20AwareABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"ERC20Token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"erc20tokenAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ERC20AwareBin is the compiled bytecode used for deploying new contracts.
const ERC20AwareBin = `0x608060405234801561001057600080fd5b50604051602080610117833981016040525160008054600160a060020a03909216600160a060020a031990921691909117905560c6806100516000396000f300608060405260043610603e5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416637a80760e81146043575b600080fd5b348015604e57600080fd5b506055607e565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b60005473ffffffffffffffffffffffffffffffffffffffff16815600a165627a7a7230582008a06d24cc1a85b57dc4b23176a06d710a9e3213ebd9c982b7fc46b4fc6c76190029`

// DeployERC20Aware deploys a new Ethereum contract, binding an instance of ERC20Aware to it.
func DeployERC20Aware(auth *bind.TransactOpts, backend bind.ContractBackend, erc20tokenAddress common.Address) (common.Address, *types.Transaction, *ERC20Aware, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20AwareABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20AwareBin), backend, erc20tokenAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Aware{ERC20AwareCaller: ERC20AwareCaller{contract: contract}, ERC20AwareTransactor: ERC20AwareTransactor{contract: contract}, ERC20AwareFilterer: ERC20AwareFilterer{contract: contract}}, nil
}

// ERC20Aware is an auto generated Go binding around an Ethereum contract.
type ERC20Aware struct {
	ERC20AwareCaller     // Read-only binding to the contract
	ERC20AwareTransactor // Write-only binding to the contract
	ERC20AwareFilterer   // Log filterer for contract events
}

// ERC20AwareCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20AwareCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20AwareTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20AwareTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20AwareFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20AwareFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20AwareSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20AwareSession struct {
	Contract     *ERC20Aware       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20AwareCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20AwareCallerSession struct {
	Contract *ERC20AwareCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ERC20AwareTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20AwareTransactorSession struct {
	Contract     *ERC20AwareTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ERC20AwareRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20AwareRaw struct {
	Contract *ERC20Aware // Generic contract binding to access the raw methods on
}

// ERC20AwareCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20AwareCallerRaw struct {
	Contract *ERC20AwareCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20AwareTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20AwareTransactorRaw struct {
	Contract *ERC20AwareTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Aware creates a new instance of ERC20Aware, bound to a specific deployed contract.
func NewERC20Aware(address common.Address, backend bind.ContractBackend) (*ERC20Aware, error) {
	contract, err := bindERC20Aware(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Aware{ERC20AwareCaller: ERC20AwareCaller{contract: contract}, ERC20AwareTransactor: ERC20AwareTransactor{contract: contract}, ERC20AwareFilterer: ERC20AwareFilterer{contract: contract}}, nil
}

// NewERC20AwareCaller creates a new read-only instance of ERC20Aware, bound to a specific deployed contract.
func NewERC20AwareCaller(address common.Address, caller bind.ContractCaller) (*ERC20AwareCaller, error) {
	contract, err := bindERC20Aware(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20AwareCaller{contract: contract}, nil
}

// NewERC20AwareTransactor creates a new write-only instance of ERC20Aware, bound to a specific deployed contract.
func NewERC20AwareTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20AwareTransactor, error) {
	contract, err := bindERC20Aware(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20AwareTransactor{contract: contract}, nil
}

// NewERC20AwareFilterer creates a new log filterer instance of ERC20Aware, bound to a specific deployed contract.
func NewERC20AwareFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20AwareFilterer, error) {
	contract, err := bindERC20Aware(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20AwareFilterer{contract: contract}, nil
}

// bindERC20Aware binds a generic wrapper to an already deployed contract.
func bindERC20Aware(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20AwareABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Aware *ERC20AwareRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Aware.Contract.ERC20AwareCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Aware *ERC20AwareRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Aware.Contract.ERC20AwareTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Aware *ERC20AwareRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Aware.Contract.ERC20AwareTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Aware *ERC20AwareCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Aware.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Aware *ERC20AwareTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Aware.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Aware *ERC20AwareTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Aware.Contract.contract.Transact(opts, method, params...)
}

// ERC20Token is a free data retrieval call binding the contract method 0x7a80760e.
//
// Solidity: function ERC20Token() constant returns(address)
func (_ERC20Aware *ERC20AwareCaller) ERC20Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC20Aware.contract.Call(opts, out, "ERC20Token")
	return *ret0, err
}

// ERC20Token is a free data retrieval call binding the contract method 0x7a80760e.
//
// Solidity: function ERC20Token() constant returns(address)
func (_ERC20Aware *ERC20AwareSession) ERC20Token() (common.Address, error) {
	return _ERC20Aware.Contract.ERC20Token(&_ERC20Aware.CallOpts)
}

// ERC20Token is a free data retrieval call binding the contract method 0x7a80760e.
//
// Solidity: function ERC20Token() constant returns(address)
func (_ERC20Aware *ERC20AwareCallerSession) ERC20Token() (common.Address, error) {
	return _ERC20Aware.Contract.ERC20Token(&_ERC20Aware.CallOpts)
}

// ERC20BasicABI is the input ABI used to generate the binding from.
const ERC20BasicABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// ERC20BasicBin is the compiled bytecode used for deploying new contracts.
const ERC20BasicBin = `0x`

// DeployERC20Basic deploys a new Ethereum contract, binding an instance of ERC20Basic to it.
func DeployERC20Basic(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20Basic, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BasicABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20BasicBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Basic{ERC20BasicCaller: ERC20BasicCaller{contract: contract}, ERC20BasicTransactor: ERC20BasicTransactor{contract: contract}, ERC20BasicFilterer: ERC20BasicFilterer{contract: contract}}, nil
}

// ERC20Basic is an auto generated Go binding around an Ethereum contract.
type ERC20Basic struct {
	ERC20BasicCaller     // Read-only binding to the contract
	ERC20BasicTransactor // Write-only binding to the contract
	ERC20BasicFilterer   // Log filterer for contract events
}

// ERC20BasicCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20BasicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20BasicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20BasicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20BasicSession struct {
	Contract     *ERC20Basic       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20BasicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20BasicCallerSession struct {
	Contract *ERC20BasicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ERC20BasicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20BasicTransactorSession struct {
	Contract     *ERC20BasicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ERC20BasicRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20BasicRaw struct {
	Contract *ERC20Basic // Generic contract binding to access the raw methods on
}

// ERC20BasicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20BasicCallerRaw struct {
	Contract *ERC20BasicCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20BasicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20BasicTransactorRaw struct {
	Contract *ERC20BasicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Basic creates a new instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20Basic(address common.Address, backend bind.ContractBackend) (*ERC20Basic, error) {
	contract, err := bindERC20Basic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Basic{ERC20BasicCaller: ERC20BasicCaller{contract: contract}, ERC20BasicTransactor: ERC20BasicTransactor{contract: contract}, ERC20BasicFilterer: ERC20BasicFilterer{contract: contract}}, nil
}

// NewERC20BasicCaller creates a new read-only instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicCaller(address common.Address, caller bind.ContractCaller) (*ERC20BasicCaller, error) {
	contract, err := bindERC20Basic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicCaller{contract: contract}, nil
}

// NewERC20BasicTransactor creates a new write-only instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20BasicTransactor, error) {
	contract, err := bindERC20Basic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicTransactor{contract: contract}, nil
}

// NewERC20BasicFilterer creates a new log filterer instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20BasicFilterer, error) {
	contract, err := bindERC20Basic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicFilterer{contract: contract}, nil
}

// bindERC20Basic binds a generic wrapper to an already deployed contract.
func bindERC20Basic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BasicABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Basic *ERC20BasicRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Basic.Contract.ERC20BasicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Basic *ERC20BasicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Basic.Contract.ERC20BasicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Basic *ERC20BasicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Basic.Contract.ERC20BasicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Basic *ERC20BasicCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Basic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Basic *ERC20BasicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Basic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Basic *ERC20BasicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Basic.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20Basic *ERC20BasicCaller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Basic.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20Basic *ERC20BasicSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20Basic.Contract.BalanceOf(&_ERC20Basic.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20Basic *ERC20BasicCallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20Basic.Contract.BalanceOf(&_ERC20Basic.CallOpts, who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Basic *ERC20BasicCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Basic.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Basic *ERC20BasicSession) TotalSupply() (*big.Int, error) {
	return _ERC20Basic.Contract.TotalSupply(&_ERC20Basic.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Basic *ERC20BasicCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Basic.Contract.TotalSupply(&_ERC20Basic.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20Basic *ERC20BasicTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20Basic *ERC20BasicSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.Contract.Transfer(&_ERC20Basic.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20Basic *ERC20BasicTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.Contract.Transfer(&_ERC20Basic.TransactOpts, to, value)
}

// ERC20BasicTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Basic contract.
type ERC20BasicTransferIterator struct {
	Event *ERC20BasicTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20BasicTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BasicTransfer)
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
		it.Event = new(ERC20BasicTransfer)
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
func (it *ERC20BasicTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BasicTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BasicTransfer represents a Transfer event raised by the ERC20Basic contract.
type ERC20BasicTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20Basic *ERC20BasicFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20BasicTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Basic.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicTransferIterator{contract: _ERC20Basic.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20Basic *ERC20BasicFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20BasicTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Basic.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BasicTransfer)
				if err := _ERC20Basic.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// IdentityBalancesABI is the input ABI used to generate the binding from.
const IdentityBalancesABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERC20Token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"topUp\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalBalance\",\"type\":\"uint256\"}],\"name\":\"ToppedUp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalBalance\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"}]"

// IdentityBalancesBin is the compiled bytecode used for deploying new contracts.
const IdentityBalancesBin = `0x608060405234801561001057600080fd5b506040516020806105a6833981016040525160008054600160a060020a03909216600160a060020a0319909216919091179055610554806100526000396000f3006080604052600436106100615763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166327e235e381146100665780635ebfdfc6146100995780637a80760e146100d1578063d6f7ddf914610102575b600080fd5b34801561007257600080fd5b50610087600160a060020a0360043516610126565b60408051918252519081900360200190f35b3480156100a557600080fd5b506100bd60043560ff60243516604435606435610138565b604080519115158252519081900360200190f35b3480156100dd57600080fd5b506100e66103f4565b60408051600160a060020a039092168252519081900360200190f35b34801561010e57600080fd5b506100bd600160a060020a0360043516602435610403565b60016020526000908152604090205481565b6000606060006040805190810160405280601181526020017f576974686472617720726571756573743a000000000000000000000000000000815250876040516020018083805190602001908083835b602083106101a75780518252601f199092019160209182019101610188565b51815160209384036101000a600019018019909216911617905292019384525060408051808503815293820190819052835193965060019450869390925082918401908083835b6020831061020d5780518252601f1990920191602091820191016101ee565b51815160209384036101000a60001901801990921691161790526040805192909401829003822060008084528383018087529190915260ff8e1683860152606083018d9052608083018c9052935160a08084019750919550601f1981019492819003909101925090865af1158015610289573d6000803e3d6000fd5b5050604051601f1901519150506000600160a060020a038216116102ac57600080fd5b600087116102b957600080fd5b600160a060020a0381166000908152600160205260409020548711156102de57600080fd5b600160a060020a03808216600090815260016020908152604080832080548c90039055825481517fa9059cbb000000000000000000000000000000000000000000000000000000008152336004820152602481018d9052915194169363a9059cbb93604480840194938390030190829087803b15801561035d57600080fd5b505af1158015610371573d6000803e3d6000fd5b505050506040513d602081101561038757600080fd5b5051151561039457600080fd5b600160a060020a0381166000818152600160209081526040918290205482518b81529182015281517f92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6929181900390910190a25060019695505050505050565b600054600160a060020a031681565b600080821161041157600080fd5b600160a060020a038084166000908152600160209081526040808320805487019055825481517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526044810188905291519416936323b872dd93606480840194938390030190829087803b15801561049557600080fd5b505af11580156104a9573d6000803e3d6000fd5b505050506040513d60208110156104bf57600080fd5b505115156104cc57600080fd5b600160a060020a0383166000818152600160209081526040918290205482518681529182015281517fbdde76a89d276b6d334c784be5c2d00d6c2219d11a2f2c80e00b85144845ab4d929181900390910190a2506001929150505600a165627a7a7230582083c0897f5216e30de8a7a09b3f1dca69fd6d3ff8a40b61b8d0156d9da9003a5f0029`

// DeployIdentityBalances deploys a new Ethereum contract, binding an instance of IdentityBalances to it.
func DeployIdentityBalances(auth *bind.TransactOpts, backend bind.ContractBackend, tokenAddress common.Address) (common.Address, *types.Transaction, *IdentityBalances, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityBalancesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IdentityBalancesBin), backend, tokenAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IdentityBalances{IdentityBalancesCaller: IdentityBalancesCaller{contract: contract}, IdentityBalancesTransactor: IdentityBalancesTransactor{contract: contract}, IdentityBalancesFilterer: IdentityBalancesFilterer{contract: contract}}, nil
}

// IdentityBalances is an auto generated Go binding around an Ethereum contract.
type IdentityBalances struct {
	IdentityBalancesCaller     // Read-only binding to the contract
	IdentityBalancesTransactor // Write-only binding to the contract
	IdentityBalancesFilterer   // Log filterer for contract events
}

// IdentityBalancesCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdentityBalancesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityBalancesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdentityBalancesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityBalancesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdentityBalancesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityBalancesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdentityBalancesSession struct {
	Contract     *IdentityBalances // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdentityBalancesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdentityBalancesCallerSession struct {
	Contract *IdentityBalancesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IdentityBalancesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdentityBalancesTransactorSession struct {
	Contract     *IdentityBalancesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IdentityBalancesRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdentityBalancesRaw struct {
	Contract *IdentityBalances // Generic contract binding to access the raw methods on
}

// IdentityBalancesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdentityBalancesCallerRaw struct {
	Contract *IdentityBalancesCaller // Generic read-only contract binding to access the raw methods on
}

// IdentityBalancesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdentityBalancesTransactorRaw struct {
	Contract *IdentityBalancesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdentityBalances creates a new instance of IdentityBalances, bound to a specific deployed contract.
func NewIdentityBalances(address common.Address, backend bind.ContractBackend) (*IdentityBalances, error) {
	contract, err := bindIdentityBalances(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IdentityBalances{IdentityBalancesCaller: IdentityBalancesCaller{contract: contract}, IdentityBalancesTransactor: IdentityBalancesTransactor{contract: contract}, IdentityBalancesFilterer: IdentityBalancesFilterer{contract: contract}}, nil
}

// NewIdentityBalancesCaller creates a new read-only instance of IdentityBalances, bound to a specific deployed contract.
func NewIdentityBalancesCaller(address common.Address, caller bind.ContractCaller) (*IdentityBalancesCaller, error) {
	contract, err := bindIdentityBalances(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityBalancesCaller{contract: contract}, nil
}

// NewIdentityBalancesTransactor creates a new write-only instance of IdentityBalances, bound to a specific deployed contract.
func NewIdentityBalancesTransactor(address common.Address, transactor bind.ContractTransactor) (*IdentityBalancesTransactor, error) {
	contract, err := bindIdentityBalances(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityBalancesTransactor{contract: contract}, nil
}

// NewIdentityBalancesFilterer creates a new log filterer instance of IdentityBalances, bound to a specific deployed contract.
func NewIdentityBalancesFilterer(address common.Address, filterer bind.ContractFilterer) (*IdentityBalancesFilterer, error) {
	contract, err := bindIdentityBalances(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdentityBalancesFilterer{contract: contract}, nil
}

// bindIdentityBalances binds a generic wrapper to an already deployed contract.
func bindIdentityBalances(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityBalancesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityBalances *IdentityBalancesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IdentityBalances.Contract.IdentityBalancesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityBalances *IdentityBalancesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityBalances.Contract.IdentityBalancesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityBalances *IdentityBalancesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityBalances.Contract.IdentityBalancesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityBalances *IdentityBalancesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IdentityBalances.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityBalances *IdentityBalancesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityBalances.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityBalances *IdentityBalancesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityBalances.Contract.contract.Transact(opts, method, params...)
}

// ERC20Token is a free data retrieval call binding the contract method 0x7a80760e.
//
// Solidity: function ERC20Token() constant returns(address)
func (_IdentityBalances *IdentityBalancesCaller) ERC20Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IdentityBalances.contract.Call(opts, out, "ERC20Token")
	return *ret0, err
}

// ERC20Token is a free data retrieval call binding the contract method 0x7a80760e.
//
// Solidity: function ERC20Token() constant returns(address)
func (_IdentityBalances *IdentityBalancesSession) ERC20Token() (common.Address, error) {
	return _IdentityBalances.Contract.ERC20Token(&_IdentityBalances.CallOpts)
}

// ERC20Token is a free data retrieval call binding the contract method 0x7a80760e.
//
// Solidity: function ERC20Token() constant returns(address)
func (_IdentityBalances *IdentityBalancesCallerSession) ERC20Token() (common.Address, error) {
	return _IdentityBalances.Contract.ERC20Token(&_IdentityBalances.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_IdentityBalances *IdentityBalancesCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IdentityBalances.contract.Call(opts, out, "balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_IdentityBalances *IdentityBalancesSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _IdentityBalances.Contract.Balances(&_IdentityBalances.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_IdentityBalances *IdentityBalancesCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _IdentityBalances.Contract.Balances(&_IdentityBalances.CallOpts, arg0)
}

// TopUp is a paid mutator transaction binding the contract method 0xd6f7ddf9.
//
// Solidity: function topUp(identity address, amount uint256) returns(bool)
func (_IdentityBalances *IdentityBalancesTransactor) TopUp(opts *bind.TransactOpts, identity common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IdentityBalances.contract.Transact(opts, "topUp", identity, amount)
}

// TopUp is a paid mutator transaction binding the contract method 0xd6f7ddf9.
//
// Solidity: function topUp(identity address, amount uint256) returns(bool)
func (_IdentityBalances *IdentityBalancesSession) TopUp(identity common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IdentityBalances.Contract.TopUp(&_IdentityBalances.TransactOpts, identity, amount)
}

// TopUp is a paid mutator transaction binding the contract method 0xd6f7ddf9.
//
// Solidity: function topUp(identity address, amount uint256) returns(bool)
func (_IdentityBalances *IdentityBalancesTransactorSession) TopUp(identity common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IdentityBalances.Contract.TopUp(&_IdentityBalances.TransactOpts, identity, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5ebfdfc6.
//
// Solidity: function withdraw(amount uint256, v uint8, r bytes32, s bytes32) returns(bool)
func (_IdentityBalances *IdentityBalancesTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IdentityBalances.contract.Transact(opts, "withdraw", amount, v, r, s)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5ebfdfc6.
//
// Solidity: function withdraw(amount uint256, v uint8, r bytes32, s bytes32) returns(bool)
func (_IdentityBalances *IdentityBalancesSession) Withdraw(amount *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IdentityBalances.Contract.Withdraw(&_IdentityBalances.TransactOpts, amount, v, r, s)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5ebfdfc6.
//
// Solidity: function withdraw(amount uint256, v uint8, r bytes32, s bytes32) returns(bool)
func (_IdentityBalances *IdentityBalancesTransactorSession) Withdraw(amount *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IdentityBalances.Contract.Withdraw(&_IdentityBalances.TransactOpts, amount, v, r, s)
}

// IdentityBalancesToppedUpIterator is returned from FilterToppedUp and is used to iterate over the raw logs and unpacked data for ToppedUp events raised by the IdentityBalances contract.
type IdentityBalancesToppedUpIterator struct {
	Event *IdentityBalancesToppedUp // Event containing the contract specifics and raw log

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
func (it *IdentityBalancesToppedUpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityBalancesToppedUp)
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
		it.Event = new(IdentityBalancesToppedUp)
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
func (it *IdentityBalancesToppedUpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityBalancesToppedUpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityBalancesToppedUp represents a ToppedUp event raised by the IdentityBalances contract.
type IdentityBalancesToppedUp struct {
	Identity     common.Address
	Amount       *big.Int
	TotalBalance *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterToppedUp is a free log retrieval operation binding the contract event 0xbdde76a89d276b6d334c784be5c2d00d6c2219d11a2f2c80e00b85144845ab4d.
//
// Solidity: e ToppedUp(identity indexed address, amount uint256, totalBalance uint256)
func (_IdentityBalances *IdentityBalancesFilterer) FilterToppedUp(opts *bind.FilterOpts, identity []common.Address) (*IdentityBalancesToppedUpIterator, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _IdentityBalances.contract.FilterLogs(opts, "ToppedUp", identityRule)
	if err != nil {
		return nil, err
	}
	return &IdentityBalancesToppedUpIterator{contract: _IdentityBalances.contract, event: "ToppedUp", logs: logs, sub: sub}, nil
}

// WatchToppedUp is a free log subscription operation binding the contract event 0xbdde76a89d276b6d334c784be5c2d00d6c2219d11a2f2c80e00b85144845ab4d.
//
// Solidity: e ToppedUp(identity indexed address, amount uint256, totalBalance uint256)
func (_IdentityBalances *IdentityBalancesFilterer) WatchToppedUp(opts *bind.WatchOpts, sink chan<- *IdentityBalancesToppedUp, identity []common.Address) (event.Subscription, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _IdentityBalances.contract.WatchLogs(opts, "ToppedUp", identityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityBalancesToppedUp)
				if err := _IdentityBalances.contract.UnpackLog(event, "ToppedUp", log); err != nil {
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

// IdentityBalancesWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the IdentityBalances contract.
type IdentityBalancesWithdrawnIterator struct {
	Event *IdentityBalancesWithdrawn // Event containing the contract specifics and raw log

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
func (it *IdentityBalancesWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityBalancesWithdrawn)
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
		it.Event = new(IdentityBalancesWithdrawn)
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
func (it *IdentityBalancesWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityBalancesWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityBalancesWithdrawn represents a Withdrawn event raised by the IdentityBalances contract.
type IdentityBalancesWithdrawn struct {
	Identity     common.Address
	Amount       *big.Int
	TotalBalance *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: e Withdrawn(identity indexed address, amount uint256, totalBalance uint256)
func (_IdentityBalances *IdentityBalancesFilterer) FilterWithdrawn(opts *bind.FilterOpts, identity []common.Address) (*IdentityBalancesWithdrawnIterator, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _IdentityBalances.contract.FilterLogs(opts, "Withdrawn", identityRule)
	if err != nil {
		return nil, err
	}
	return &IdentityBalancesWithdrawnIterator{contract: _IdentityBalances.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: e Withdrawn(identity indexed address, amount uint256, totalBalance uint256)
func (_IdentityBalances *IdentityBalancesFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *IdentityBalancesWithdrawn, identity []common.Address) (event.Subscription, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _IdentityBalances.contract.WatchLogs(opts, "Withdrawn", identityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityBalancesWithdrawn)
				if err := _IdentityBalances.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// IdentityPromisesABI is the input ABI used to generate the binding from.
const IdentityPromisesABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"registrationFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"receiverAndSigns\",\"type\":\"bytes32\"},{\"name\":\"extraDataHash\",\"type\":\"bytes32\"},{\"name\":\"seq\",\"type\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"sender_R\",\"type\":\"bytes32\"},{\"name\":\"sender_S\",\"type\":\"bytes32\"},{\"name\":\"receiver_R\",\"type\":\"bytes32\"},{\"name\":\"receiver_S\",\"type\":\"bytes32\"}],\"name\":\"clearPromise\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"changeRegistrationFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERC20Token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"}],\"name\":\"getPublicKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pubKeyPart1\",\"type\":\"bytes32\"},{\"name\":\"pubKeyPart2\",\"type\":\"bytes32\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"RegisterIdentity\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"}],\"name\":\"isRegistered\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"topUp\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"transferCollectedFeeTo\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"collectedFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"clearedPromises\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"erc20address\",\"type\":\"address\"},{\"name\":\"registrationFee\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"seqNo\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PromiseCleared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalBalance\",\"type\":\"uint256\"}],\"name\":\"ToppedUp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalBalance\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"identity\",\"type\":\"address\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// IdentityPromisesBin is the compiled bytecode used for deploying new contracts.
const IdentityPromisesBin = `0x608060405234801561001057600080fd5b5060405160408061129e83398101604052805160209091015160008054600160a060020a0319908116331790915560018054600160a060020a0390941693909116929092179091556003556112348061006a6000396000f3006080604052600436106100da5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166314c44e0981146100df57806327e235e3146101065780634fe0d3321461012757806350050769146101685780635ebfdfc6146101825780637a80760e146101a6578063857cdbb8146101d75780638da5cb5b14610211578063a17fc7f314610226578063c3c5a5471461024d578063d6f7ddf91461026e578063e325253714610292578063e811f50a146102b3578063f2fde38b146102c8578063f6c1a23e146102e9575b600080fd5b3480156100eb57600080fd5b506100f4610310565b60408051918252519081900360200190f35b34801561011257600080fd5b506100f4600160a060020a0360043516610316565b34801561013357600080fd5b5061015460043560243560443560643560843560a43560c43560e435610328565b604080519115158252519081900360200190f35b34801561017457600080fd5b50610180600435610728565b005b34801561018e57600080fd5b5061015460043560ff60243516604435606435610744565b3480156101b257600080fd5b506101bb610a01565b60408051600160a060020a039092168252519081900360200190f35b3480156101e357600080fd5b506101f8600160a060020a0360043516610a10565b6040805192835260208301919091528051918290030190f35b34801561021d57600080fd5b506101bb610a33565b34801561023257600080fd5b5061015460043560243560ff60443516606435608435610a42565b34801561025957600080fd5b50610154600160a060020a0360043516610d80565b34801561027a57600080fd5b50610154600160a060020a0360043516602435610db3565b34801561029e57600080fd5b50610154600160a060020a0360043516610ed9565b3480156102bf57600080fd5b506100f4610fc0565b3480156102d457600080fd5b50610180600160a060020a0360043516610fc6565b3480156102f557600080fd5b506100f4600160a060020a036004358116906024351661105a565b60035481565b60056020526000908152604090205481565b600080600080600080600061033c8f611077565b9550955095506040805190810160405280600e81526020017f497373756572207072656669783a0000000000000000000000000000000000008152508e878f8f6040516020018086805190602001908083835b602083106103ae5780518252601f19909201916020918201910161038f565b51815160209384036101000a6000190180199092169116179052920196875250600160a060020a03949094166c010000000000000000000000000285850152506034840191909152605480840191909152604080518085039092018252607490930192839052805190935082918401908083835b602083106104415780518252601f199092019160209182019101610422565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209250600183868d8d604051600081526020016040526040518085600019166000191681526020018460ff1660ff1681526020018360001916600019168152602001826000191660001916815260200194505050505060206040516020810390808403906000865af11580156104e7573d6000803e3d6000fd5b50505060206040510351915060016040805190810160405280601081526020017f5265636569766572207072656669783a0000000000000000000000000000000081525084846040516020018084805190602001908083835b6020831061055f5780518252601f199092019160209182019101610540565b51815160209384036101000a6000190180199092169116179052920194855250600160a060020a03929092166c0100000000000000000000000002838301525060408051601481850301815260349093019081905282519293509182918401908083835b602083106105e25780518252601f1990920191602091820191016105c3565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020858b8b604051600081526020016040526040518085600019166000191681526020018460ff1660ff1681526020018360001916600019168152602001826000191660001916815260200194505050505060206040516020810390808403906000865af1158015610683573d6000803e3d6000fd5b5050604051601f1901519150506000600160a060020a038316116106a657600080fd5b6000600160a060020a038216116106bc57600080fd5b600160a060020a03818116908716146106d457600080fd5b6106dd82610d80565b15156106e857600080fd5b6106f186610d80565b15156106fc57600080fd5b60008c1161070957600080fd5b61071586838f8f6110b2565b9f9e505050505050505050505050505050565b600054600160a060020a0316331461073f57600080fd5b600355565b6000606060006040805190810160405280601181526020017f576974686472617720726571756573743a000000000000000000000000000000815250876040516020018083805190602001908083835b602083106107b35780518252601f199092019160209182019101610794565b51815160209384036101000a600019018019909216911617905292019384525060408051808503815293820190819052835193965060019450869390925082918401908083835b602083106108195780518252601f1990920191602091820191016107fa565b51815160209384036101000a60001901801990921691161790526040805192909401829003822060008084528383018087529190915260ff8e1683860152606083018d9052608083018c9052935160a08084019750919550601f1981019492819003909101925090865af1158015610895573d6000803e3d6000fd5b5050604051601f1901519150506000600160a060020a038216116108b857600080fd5b600087116108c557600080fd5b600160a060020a0381166000908152600560205260409020548711156108ea57600080fd5b600160a060020a03808216600090815260056020908152604080832080548c9003905560015481517fa9059cbb000000000000000000000000000000000000000000000000000000008152336004820152602481018d9052915194169363a9059cbb93604480840194938390030190829087803b15801561096a57600080fd5b505af115801561097e573d6000803e3d6000fd5b505050506040513d602081101561099457600080fd5b505115156109a157600080fd5b600160a060020a0381166000818152600560209081526040918290205482518b81529182015281517f92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6929181900390910190a25060019695505050505050565b600154600160a060020a031681565b600160a060020a0316600090815260026020526040902080546001909101549091565b600054600160a060020a031681565b604080516020808201889052818301879052825180830384018152606090920192839052815160009384938493909282918401908083835b60208310610a995780518252601f199092019160209182019101610a7a565b51815160209384036101000a600019018019909216911617905260408051929094018290038220828501855260108084527f5265676973746572207072656669783a000000000000000000000000000000008484019081529551919950600197509295508e948e94509101918291908083835b60208310610b2b5780518252601f199092019160209182019101610b0c565b51815160209384036101000a6000190180199092169116179052920194855250838101929092525060408051808403830181529281019081905282519293509182918401908083835b60208310610b935780518252601f199092019160209182019101610b74565b51815160209384036101000a60001901801990921691161790526040805192909401829003822060008084528383018087529190915260ff8e1683860152606083018d9052608083018c9052935160a08084019750919550601f1981019492819003909101925090865af1158015610c0f573d6000803e3d6000fd5b5050604051601f1901519150506000600160a060020a03821611610c3257600080fd5b600160a060020a0382811690821614610c4a57600080fd5b610c5381610d80565b15610c5d57600080fd5b600380546004805490910181556040805180820182528b815260208082018c8152600160a060020a038088166000908152600284528581209451855591516001948501559254955484517f23b872dd00000000000000000000000000000000000000000000000000000000815233968101969096523060248701526044860152925194909116936323b872dd93606480820194918390030190829087803b158015610d0757600080fd5b505af1158015610d1b573d6000803e3d6000fd5b505050506040513d6020811015610d3157600080fd5b50511515610d3e57600080fd5b604051600160a060020a038216907f2d3734a8e47ac8316e500ac231c90a6e1848ca2285f40d07eaa52005e4b3a0e990600090a2506001979650505050505050565b600160a060020a0381166000908152600260205260408120805482108015610dac575060018101546000105b9392505050565b6000808211610dc157600080fd5b600160a060020a03808416600090815260056020908152604080832080548701905560015481517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526044810188905291519416936323b872dd93606480840194938390030190829087803b158015610e4657600080fd5b505af1158015610e5a573d6000803e3d6000fd5b505050506040513d6020811015610e7057600080fd5b50511515610e7d57600080fd5b600160a060020a0383166000818152600560209081526040918290205482518681529182015281517fbdde76a89d276b6d334c784be5c2d00d6c2219d11a2f2c80e00b85144845ab4d929181900390910190a250600192915050565b600080548190600160a060020a03163314610ef357600080fd5b600454600010610f0257600080fd5b50600480546000808355600154604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a038881169682019690965260248101859052905193949091169263a9059cbb92604480840193602093929083900390910190829087803b158015610f8057600080fd5b505af1158015610f94573d6000803e3d6000fd5b505050506040513d6020811015610faa57600080fd5b50511515610fb757600080fd5b50600192915050565b60045481565b600054600160a060020a03163314610fdd57600080fd5b600160a060020a0381161515610ff257600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600660209081526000928352604080842090915290825290205481565b907f0100000000000000000000000000000000000000000000000000000000000000600a83901a810281900491600b84901a82029190910490565b600160a060020a038084166000908152600660209081526040808320938816835292905290812054819084116110e757600080fd5b600160a060020a038086166000908152600660209081526040808320938a1683529290522084905561111a858785611187565b905082811461112857600080fd5b85600160a060020a031685600160a060020a03167f3d35f4a81f47a50c001e0e834334665e3a64680385c0a8d2fdfe2e0d209645768686604051808381526020018281526020019250505060405180910390a350600195945050505050565b60008080831161119657600080fd5b600160a060020a0385166000908152600560205260409020546111b990846111f2565b600160a060020a039586166000908152600560205260408082208054849003905595909616865293909420805484019055509092915050565b60008183106112015781610dac565b50909190505600a165627a7a72305820337bf536fd14c4ff30415cb65efbb9bbba8aa906872756f3e2d5039750021b850029`

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
// Solidity: function balances( address) constant returns(uint256)
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
// Solidity: function balances( address) constant returns(uint256)
func (_IdentityPromises *IdentityPromisesSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _IdentityPromises.Contract.Balances(&_IdentityPromises.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_IdentityPromises *IdentityPromisesCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _IdentityPromises.Contract.Balances(&_IdentityPromises.CallOpts, arg0)
}

// ClearedPromises is a free data retrieval call binding the contract method 0xf6c1a23e.
//
// Solidity: function clearedPromises( address,  address) constant returns(uint256)
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
// Solidity: function clearedPromises( address,  address) constant returns(uint256)
func (_IdentityPromises *IdentityPromisesSession) ClearedPromises(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _IdentityPromises.Contract.ClearedPromises(&_IdentityPromises.CallOpts, arg0, arg1)
}

// ClearedPromises is a free data retrieval call binding the contract method 0xf6c1a23e.
//
// Solidity: function clearedPromises( address,  address) constant returns(uint256)
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
// Solidity: function getPublicKey(identity address) constant returns(bytes32, bytes32)
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
// Solidity: function getPublicKey(identity address) constant returns(bytes32, bytes32)
func (_IdentityPromises *IdentityPromisesSession) GetPublicKey(identity common.Address) ([32]byte, [32]byte, error) {
	return _IdentityPromises.Contract.GetPublicKey(&_IdentityPromises.CallOpts, identity)
}

// GetPublicKey is a free data retrieval call binding the contract method 0x857cdbb8.
//
// Solidity: function getPublicKey(identity address) constant returns(bytes32, bytes32)
func (_IdentityPromises *IdentityPromisesCallerSession) GetPublicKey(identity common.Address) ([32]byte, [32]byte, error) {
	return _IdentityPromises.Contract.GetPublicKey(&_IdentityPromises.CallOpts, identity)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(identity address) constant returns(bool)
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
// Solidity: function isRegistered(identity address) constant returns(bool)
func (_IdentityPromises *IdentityPromisesSession) IsRegistered(identity common.Address) (bool, error) {
	return _IdentityPromises.Contract.IsRegistered(&_IdentityPromises.CallOpts, identity)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(identity address) constant returns(bool)
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
// Solidity: function RegisterIdentity(pubKeyPart1 bytes32, pubKeyPart2 bytes32, v uint8, r bytes32, s bytes32) returns(bool)
func (_IdentityPromises *IdentityPromisesTransactor) RegisterIdentity(opts *bind.TransactOpts, pubKeyPart1 [32]byte, pubKeyPart2 [32]byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IdentityPromises.contract.Transact(opts, "RegisterIdentity", pubKeyPart1, pubKeyPart2, v, r, s)
}

// RegisterIdentity is a paid mutator transaction binding the contract method 0xa17fc7f3.
//
// Solidity: function RegisterIdentity(pubKeyPart1 bytes32, pubKeyPart2 bytes32, v uint8, r bytes32, s bytes32) returns(bool)
func (_IdentityPromises *IdentityPromisesSession) RegisterIdentity(pubKeyPart1 [32]byte, pubKeyPart2 [32]byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IdentityPromises.Contract.RegisterIdentity(&_IdentityPromises.TransactOpts, pubKeyPart1, pubKeyPart2, v, r, s)
}

// RegisterIdentity is a paid mutator transaction binding the contract method 0xa17fc7f3.
//
// Solidity: function RegisterIdentity(pubKeyPart1 bytes32, pubKeyPart2 bytes32, v uint8, r bytes32, s bytes32) returns(bool)
func (_IdentityPromises *IdentityPromisesTransactorSession) RegisterIdentity(pubKeyPart1 [32]byte, pubKeyPart2 [32]byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IdentityPromises.Contract.RegisterIdentity(&_IdentityPromises.TransactOpts, pubKeyPart1, pubKeyPart2, v, r, s)
}

// ChangeRegistrationFee is a paid mutator transaction binding the contract method 0x50050769.
//
// Solidity: function changeRegistrationFee(newFee uint256) returns()
func (_IdentityPromises *IdentityPromisesTransactor) ChangeRegistrationFee(opts *bind.TransactOpts, newFee *big.Int) (*types.Transaction, error) {
	return _IdentityPromises.contract.Transact(opts, "changeRegistrationFee", newFee)
}

// ChangeRegistrationFee is a paid mutator transaction binding the contract method 0x50050769.
//
// Solidity: function changeRegistrationFee(newFee uint256) returns()
func (_IdentityPromises *IdentityPromisesSession) ChangeRegistrationFee(newFee *big.Int) (*types.Transaction, error) {
	return _IdentityPromises.Contract.ChangeRegistrationFee(&_IdentityPromises.TransactOpts, newFee)
}

// ChangeRegistrationFee is a paid mutator transaction binding the contract method 0x50050769.
//
// Solidity: function changeRegistrationFee(newFee uint256) returns()
func (_IdentityPromises *IdentityPromisesTransactorSession) ChangeRegistrationFee(newFee *big.Int) (*types.Transaction, error) {
	return _IdentityPromises.Contract.ChangeRegistrationFee(&_IdentityPromises.TransactOpts, newFee)
}

// ClearPromise is a paid mutator transaction binding the contract method 0x4fe0d332.
//
// Solidity: function clearPromise(receiverAndSigns bytes32, extraDataHash bytes32, seq uint256, amount uint256, sender_R bytes32, sender_S bytes32, receiver_R bytes32, receiver_S bytes32) returns(bool)
func (_IdentityPromises *IdentityPromisesTransactor) ClearPromise(opts *bind.TransactOpts, receiverAndSigns [32]byte, extraDataHash [32]byte, seq *big.Int, amount *big.Int, sender_R [32]byte, sender_S [32]byte, receiver_R [32]byte, receiver_S [32]byte) (*types.Transaction, error) {
	return _IdentityPromises.contract.Transact(opts, "clearPromise", receiverAndSigns, extraDataHash, seq, amount, sender_R, sender_S, receiver_R, receiver_S)
}

// ClearPromise is a paid mutator transaction binding the contract method 0x4fe0d332.
//
// Solidity: function clearPromise(receiverAndSigns bytes32, extraDataHash bytes32, seq uint256, amount uint256, sender_R bytes32, sender_S bytes32, receiver_R bytes32, receiver_S bytes32) returns(bool)
func (_IdentityPromises *IdentityPromisesSession) ClearPromise(receiverAndSigns [32]byte, extraDataHash [32]byte, seq *big.Int, amount *big.Int, sender_R [32]byte, sender_S [32]byte, receiver_R [32]byte, receiver_S [32]byte) (*types.Transaction, error) {
	return _IdentityPromises.Contract.ClearPromise(&_IdentityPromises.TransactOpts, receiverAndSigns, extraDataHash, seq, amount, sender_R, sender_S, receiver_R, receiver_S)
}

// ClearPromise is a paid mutator transaction binding the contract method 0x4fe0d332.
//
// Solidity: function clearPromise(receiverAndSigns bytes32, extraDataHash bytes32, seq uint256, amount uint256, sender_R bytes32, sender_S bytes32, receiver_R bytes32, receiver_S bytes32) returns(bool)
func (_IdentityPromises *IdentityPromisesTransactorSession) ClearPromise(receiverAndSigns [32]byte, extraDataHash [32]byte, seq *big.Int, amount *big.Int, sender_R [32]byte, sender_S [32]byte, receiver_R [32]byte, receiver_S [32]byte) (*types.Transaction, error) {
	return _IdentityPromises.Contract.ClearPromise(&_IdentityPromises.TransactOpts, receiverAndSigns, extraDataHash, seq, amount, sender_R, sender_S, receiver_R, receiver_S)
}

// TopUp is a paid mutator transaction binding the contract method 0xd6f7ddf9.
//
// Solidity: function topUp(identity address, amount uint256) returns(bool)
func (_IdentityPromises *IdentityPromisesTransactor) TopUp(opts *bind.TransactOpts, identity common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IdentityPromises.contract.Transact(opts, "topUp", identity, amount)
}

// TopUp is a paid mutator transaction binding the contract method 0xd6f7ddf9.
//
// Solidity: function topUp(identity address, amount uint256) returns(bool)
func (_IdentityPromises *IdentityPromisesSession) TopUp(identity common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IdentityPromises.Contract.TopUp(&_IdentityPromises.TransactOpts, identity, amount)
}

// TopUp is a paid mutator transaction binding the contract method 0xd6f7ddf9.
//
// Solidity: function topUp(identity address, amount uint256) returns(bool)
func (_IdentityPromises *IdentityPromisesTransactorSession) TopUp(identity common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IdentityPromises.Contract.TopUp(&_IdentityPromises.TransactOpts, identity, amount)
}

// TransferCollectedFeeTo is a paid mutator transaction binding the contract method 0xe3252537.
//
// Solidity: function transferCollectedFeeTo(receiver address) returns(bool)
func (_IdentityPromises *IdentityPromisesTransactor) TransferCollectedFeeTo(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _IdentityPromises.contract.Transact(opts, "transferCollectedFeeTo", receiver)
}

// TransferCollectedFeeTo is a paid mutator transaction binding the contract method 0xe3252537.
//
// Solidity: function transferCollectedFeeTo(receiver address) returns(bool)
func (_IdentityPromises *IdentityPromisesSession) TransferCollectedFeeTo(receiver common.Address) (*types.Transaction, error) {
	return _IdentityPromises.Contract.TransferCollectedFeeTo(&_IdentityPromises.TransactOpts, receiver)
}

// TransferCollectedFeeTo is a paid mutator transaction binding the contract method 0xe3252537.
//
// Solidity: function transferCollectedFeeTo(receiver address) returns(bool)
func (_IdentityPromises *IdentityPromisesTransactorSession) TransferCollectedFeeTo(receiver common.Address) (*types.Transaction, error) {
	return _IdentityPromises.Contract.TransferCollectedFeeTo(&_IdentityPromises.TransactOpts, receiver)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_IdentityPromises *IdentityPromisesTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IdentityPromises.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_IdentityPromises *IdentityPromisesSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IdentityPromises.Contract.TransferOwnership(&_IdentityPromises.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_IdentityPromises *IdentityPromisesTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IdentityPromises.Contract.TransferOwnership(&_IdentityPromises.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5ebfdfc6.
//
// Solidity: function withdraw(amount uint256, v uint8, r bytes32, s bytes32) returns(bool)
func (_IdentityPromises *IdentityPromisesTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IdentityPromises.contract.Transact(opts, "withdraw", amount, v, r, s)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5ebfdfc6.
//
// Solidity: function withdraw(amount uint256, v uint8, r bytes32, s bytes32) returns(bool)
func (_IdentityPromises *IdentityPromisesSession) Withdraw(amount *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IdentityPromises.Contract.Withdraw(&_IdentityPromises.TransactOpts, amount, v, r, s)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5ebfdfc6.
//
// Solidity: function withdraw(amount uint256, v uint8, r bytes32, s bytes32) returns(bool)
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
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
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
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
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
// Solidity: e PromiseCleared(from indexed address, to indexed address, seqNo uint256, amount uint256)
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
// Solidity: e PromiseCleared(from indexed address, to indexed address, seqNo uint256, amount uint256)
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
// Solidity: e Registered(identity indexed address)
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
// Solidity: e Registered(identity indexed address)
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
// Solidity: e ToppedUp(identity indexed address, amount uint256, totalBalance uint256)
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
// Solidity: e ToppedUp(identity indexed address, amount uint256, totalBalance uint256)
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
// Solidity: e Withdrawn(identity indexed address, amount uint256, totalBalance uint256)
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
// Solidity: e Withdrawn(identity indexed address, amount uint256, totalBalance uint256)
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

// IdentityRegistryABI is the input ABI used to generate the binding from.
const IdentityRegistryABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"registrationFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"changeRegistrationFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERC20Token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"}],\"name\":\"getPublicKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pubKeyPart1\",\"type\":\"bytes32\"},{\"name\":\"pubKeyPart2\",\"type\":\"bytes32\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"RegisterIdentity\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"}],\"name\":\"isRegistered\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"transferCollectedFeeTo\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"collectedFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\"},{\"name\":\"regFee\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"identity\",\"type\":\"address\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// IdentityRegistryBin is the compiled bytecode used for deploying new contracts.
const IdentityRegistryBin = `0x608060405234801561001057600080fd5b5060405160408061080783398101604052805160209091015160008054600160a060020a0319908116331790915560018054600160a060020a03909416939091169290921790915560035561079d8061006a6000396000f3006080604052600436106100a35763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166314c44e0981146100a857806350050769146100cf5780637a80760e146100e9578063857cdbb81461011a5780638da5cb5b14610154578063a17fc7f314610169578063c3c5a547146101a4578063e3252537146101c5578063e811f50a146101e6578063f2fde38b146101fb575b600080fd5b3480156100b457600080fd5b506100bd61021c565b60408051918252519081900360200190f35b3480156100db57600080fd5b506100e7600435610222565b005b3480156100f557600080fd5b506100fe61023e565b60408051600160a060020a039092168252519081900360200190f35b34801561012657600080fd5b5061013b600160a060020a036004351661024d565b6040805192835260208301919091528051918290030190f35b34801561016057600080fd5b506100fe610270565b34801561017557600080fd5b5061019060043560243560ff6044351660643560843561027f565b604080519115158252519081900360200190f35b3480156101b057600080fd5b50610190600160a060020a03600435166105bd565b3480156101d157600080fd5b50610190600160a060020a03600435166105f0565b3480156101f257600080fd5b506100bd6106d7565b34801561020757600080fd5b506100e7600160a060020a03600435166106dd565b60035481565b600054600160a060020a0316331461023957600080fd5b600355565b600154600160a060020a031681565b600160a060020a0316600090815260026020526040902080546001909101549091565b600054600160a060020a031681565b604080516020808201889052818301879052825180830384018152606090920192839052815160009384938493909282918401908083835b602083106102d65780518252601f1990920191602091820191016102b7565b51815160209384036101000a600019018019909216911617905260408051929094018290038220828501855260108084527f5265676973746572207072656669783a000000000000000000000000000000008484019081529551919950600197509295508e948e94509101918291908083835b602083106103685780518252601f199092019160209182019101610349565b51815160209384036101000a6000190180199092169116179052920194855250838101929092525060408051808403830181529281019081905282519293509182918401908083835b602083106103d05780518252601f1990920191602091820191016103b1565b51815160209384036101000a60001901801990921691161790526040805192909401829003822060008084528383018087529190915260ff8e1683860152606083018d9052608083018c9052935160a08084019750919550601f1981019492819003909101925090865af115801561044c573d6000803e3d6000fd5b5050604051601f1901519150506000600160a060020a0382161161046f57600080fd5b600160a060020a038281169082161461048757600080fd5b610490816105bd565b1561049a57600080fd5b600380546004805490910181556040805180820182528b815260208082018c8152600160a060020a038088166000908152600284528581209451855591516001948501559254955484517f23b872dd00000000000000000000000000000000000000000000000000000000815233968101969096523060248701526044860152925194909116936323b872dd93606480820194918390030190829087803b15801561054457600080fd5b505af1158015610558573d6000803e3d6000fd5b505050506040513d602081101561056e57600080fd5b5051151561057b57600080fd5b604051600160a060020a038216907f2d3734a8e47ac8316e500ac231c90a6e1848ca2285f40d07eaa52005e4b3a0e990600090a2506001979650505050505050565b600160a060020a03811660009081526002602052604081208054821080156105e9575060018101546000105b9392505050565b600080548190600160a060020a0316331461060a57600080fd5b60045460001061061957600080fd5b50600480546000808355600154604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a038881169682019690965260248101859052905193949091169263a9059cbb92604480840193602093929083900390910190829087803b15801561069757600080fd5b505af11580156106ab573d6000803e3d6000fd5b505050506040513d60208110156106c157600080fd5b505115156106ce57600080fd5b50600192915050565b60045481565b600054600160a060020a031633146106f457600080fd5b600160a060020a038116151561070957600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a7230582054a0a21d434504a59d905cd3a53afc5884d987f35d15969325f60728a0efa6820029`

// DeployIdentityRegistry deploys a new Ethereum contract, binding an instance of IdentityRegistry to it.
func DeployIdentityRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, tokenAddress common.Address, regFee *big.Int) (common.Address, *types.Transaction, *IdentityRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityRegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IdentityRegistryBin), backend, tokenAddress, regFee)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IdentityRegistry{IdentityRegistryCaller: IdentityRegistryCaller{contract: contract}, IdentityRegistryTransactor: IdentityRegistryTransactor{contract: contract}, IdentityRegistryFilterer: IdentityRegistryFilterer{contract: contract}}, nil
}

// IdentityRegistry is an auto generated Go binding around an Ethereum contract.
type IdentityRegistry struct {
	IdentityRegistryCaller     // Read-only binding to the contract
	IdentityRegistryTransactor // Write-only binding to the contract
	IdentityRegistryFilterer   // Log filterer for contract events
}

// IdentityRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdentityRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdentityRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdentityRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdentityRegistrySession struct {
	Contract     *IdentityRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdentityRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdentityRegistryCallerSession struct {
	Contract *IdentityRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IdentityRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdentityRegistryTransactorSession struct {
	Contract     *IdentityRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IdentityRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdentityRegistryRaw struct {
	Contract *IdentityRegistry // Generic contract binding to access the raw methods on
}

// IdentityRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdentityRegistryCallerRaw struct {
	Contract *IdentityRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// IdentityRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdentityRegistryTransactorRaw struct {
	Contract *IdentityRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdentityRegistry creates a new instance of IdentityRegistry, bound to a specific deployed contract.
func NewIdentityRegistry(address common.Address, backend bind.ContractBackend) (*IdentityRegistry, error) {
	contract, err := bindIdentityRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistry{IdentityRegistryCaller: IdentityRegistryCaller{contract: contract}, IdentityRegistryTransactor: IdentityRegistryTransactor{contract: contract}, IdentityRegistryFilterer: IdentityRegistryFilterer{contract: contract}}, nil
}

// NewIdentityRegistryCaller creates a new read-only instance of IdentityRegistry, bound to a specific deployed contract.
func NewIdentityRegistryCaller(address common.Address, caller bind.ContractCaller) (*IdentityRegistryCaller, error) {
	contract, err := bindIdentityRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryCaller{contract: contract}, nil
}

// NewIdentityRegistryTransactor creates a new write-only instance of IdentityRegistry, bound to a specific deployed contract.
func NewIdentityRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*IdentityRegistryTransactor, error) {
	contract, err := bindIdentityRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryTransactor{contract: contract}, nil
}

// NewIdentityRegistryFilterer creates a new log filterer instance of IdentityRegistry, bound to a specific deployed contract.
func NewIdentityRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*IdentityRegistryFilterer, error) {
	contract, err := bindIdentityRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryFilterer{contract: contract}, nil
}

// bindIdentityRegistry binds a generic wrapper to an already deployed contract.
func bindIdentityRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityRegistry *IdentityRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IdentityRegistry.Contract.IdentityRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityRegistry *IdentityRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.IdentityRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityRegistry *IdentityRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.IdentityRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityRegistry *IdentityRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IdentityRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityRegistry *IdentityRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityRegistry *IdentityRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.contract.Transact(opts, method, params...)
}

// ERC20Token is a free data retrieval call binding the contract method 0x7a80760e.
//
// Solidity: function ERC20Token() constant returns(address)
func (_IdentityRegistry *IdentityRegistryCaller) ERC20Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IdentityRegistry.contract.Call(opts, out, "ERC20Token")
	return *ret0, err
}

// ERC20Token is a free data retrieval call binding the contract method 0x7a80760e.
//
// Solidity: function ERC20Token() constant returns(address)
func (_IdentityRegistry *IdentityRegistrySession) ERC20Token() (common.Address, error) {
	return _IdentityRegistry.Contract.ERC20Token(&_IdentityRegistry.CallOpts)
}

// ERC20Token is a free data retrieval call binding the contract method 0x7a80760e.
//
// Solidity: function ERC20Token() constant returns(address)
func (_IdentityRegistry *IdentityRegistryCallerSession) ERC20Token() (common.Address, error) {
	return _IdentityRegistry.Contract.ERC20Token(&_IdentityRegistry.CallOpts)
}

// CollectedFee is a free data retrieval call binding the contract method 0xe811f50a.
//
// Solidity: function collectedFee() constant returns(uint256)
func (_IdentityRegistry *IdentityRegistryCaller) CollectedFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IdentityRegistry.contract.Call(opts, out, "collectedFee")
	return *ret0, err
}

// CollectedFee is a free data retrieval call binding the contract method 0xe811f50a.
//
// Solidity: function collectedFee() constant returns(uint256)
func (_IdentityRegistry *IdentityRegistrySession) CollectedFee() (*big.Int, error) {
	return _IdentityRegistry.Contract.CollectedFee(&_IdentityRegistry.CallOpts)
}

// CollectedFee is a free data retrieval call binding the contract method 0xe811f50a.
//
// Solidity: function collectedFee() constant returns(uint256)
func (_IdentityRegistry *IdentityRegistryCallerSession) CollectedFee() (*big.Int, error) {
	return _IdentityRegistry.Contract.CollectedFee(&_IdentityRegistry.CallOpts)
}

// GetPublicKey is a free data retrieval call binding the contract method 0x857cdbb8.
//
// Solidity: function getPublicKey(identity address) constant returns(bytes32, bytes32)
func (_IdentityRegistry *IdentityRegistryCaller) GetPublicKey(opts *bind.CallOpts, identity common.Address) ([32]byte, [32]byte, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _IdentityRegistry.contract.Call(opts, out, "getPublicKey", identity)
	return *ret0, *ret1, err
}

// GetPublicKey is a free data retrieval call binding the contract method 0x857cdbb8.
//
// Solidity: function getPublicKey(identity address) constant returns(bytes32, bytes32)
func (_IdentityRegistry *IdentityRegistrySession) GetPublicKey(identity common.Address) ([32]byte, [32]byte, error) {
	return _IdentityRegistry.Contract.GetPublicKey(&_IdentityRegistry.CallOpts, identity)
}

// GetPublicKey is a free data retrieval call binding the contract method 0x857cdbb8.
//
// Solidity: function getPublicKey(identity address) constant returns(bytes32, bytes32)
func (_IdentityRegistry *IdentityRegistryCallerSession) GetPublicKey(identity common.Address) ([32]byte, [32]byte, error) {
	return _IdentityRegistry.Contract.GetPublicKey(&_IdentityRegistry.CallOpts, identity)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(identity address) constant returns(bool)
func (_IdentityRegistry *IdentityRegistryCaller) IsRegistered(opts *bind.CallOpts, identity common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IdentityRegistry.contract.Call(opts, out, "isRegistered", identity)
	return *ret0, err
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(identity address) constant returns(bool)
func (_IdentityRegistry *IdentityRegistrySession) IsRegistered(identity common.Address) (bool, error) {
	return _IdentityRegistry.Contract.IsRegistered(&_IdentityRegistry.CallOpts, identity)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(identity address) constant returns(bool)
func (_IdentityRegistry *IdentityRegistryCallerSession) IsRegistered(identity common.Address) (bool, error) {
	return _IdentityRegistry.Contract.IsRegistered(&_IdentityRegistry.CallOpts, identity)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_IdentityRegistry *IdentityRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IdentityRegistry.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_IdentityRegistry *IdentityRegistrySession) Owner() (common.Address, error) {
	return _IdentityRegistry.Contract.Owner(&_IdentityRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_IdentityRegistry *IdentityRegistryCallerSession) Owner() (common.Address, error) {
	return _IdentityRegistry.Contract.Owner(&_IdentityRegistry.CallOpts)
}

// RegistrationFee is a free data retrieval call binding the contract method 0x14c44e09.
//
// Solidity: function registrationFee() constant returns(uint256)
func (_IdentityRegistry *IdentityRegistryCaller) RegistrationFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IdentityRegistry.contract.Call(opts, out, "registrationFee")
	return *ret0, err
}

// RegistrationFee is a free data retrieval call binding the contract method 0x14c44e09.
//
// Solidity: function registrationFee() constant returns(uint256)
func (_IdentityRegistry *IdentityRegistrySession) RegistrationFee() (*big.Int, error) {
	return _IdentityRegistry.Contract.RegistrationFee(&_IdentityRegistry.CallOpts)
}

// RegistrationFee is a free data retrieval call binding the contract method 0x14c44e09.
//
// Solidity: function registrationFee() constant returns(uint256)
func (_IdentityRegistry *IdentityRegistryCallerSession) RegistrationFee() (*big.Int, error) {
	return _IdentityRegistry.Contract.RegistrationFee(&_IdentityRegistry.CallOpts)
}

// RegisterIdentity is a paid mutator transaction binding the contract method 0xa17fc7f3.
//
// Solidity: function RegisterIdentity(pubKeyPart1 bytes32, pubKeyPart2 bytes32, v uint8, r bytes32, s bytes32) returns(bool)
func (_IdentityRegistry *IdentityRegistryTransactor) RegisterIdentity(opts *bind.TransactOpts, pubKeyPart1 [32]byte, pubKeyPart2 [32]byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "RegisterIdentity", pubKeyPart1, pubKeyPart2, v, r, s)
}

// RegisterIdentity is a paid mutator transaction binding the contract method 0xa17fc7f3.
//
// Solidity: function RegisterIdentity(pubKeyPart1 bytes32, pubKeyPart2 bytes32, v uint8, r bytes32, s bytes32) returns(bool)
func (_IdentityRegistry *IdentityRegistrySession) RegisterIdentity(pubKeyPart1 [32]byte, pubKeyPart2 [32]byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.RegisterIdentity(&_IdentityRegistry.TransactOpts, pubKeyPart1, pubKeyPart2, v, r, s)
}

// RegisterIdentity is a paid mutator transaction binding the contract method 0xa17fc7f3.
//
// Solidity: function RegisterIdentity(pubKeyPart1 bytes32, pubKeyPart2 bytes32, v uint8, r bytes32, s bytes32) returns(bool)
func (_IdentityRegistry *IdentityRegistryTransactorSession) RegisterIdentity(pubKeyPart1 [32]byte, pubKeyPart2 [32]byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.RegisterIdentity(&_IdentityRegistry.TransactOpts, pubKeyPart1, pubKeyPart2, v, r, s)
}

// ChangeRegistrationFee is a paid mutator transaction binding the contract method 0x50050769.
//
// Solidity: function changeRegistrationFee(newFee uint256) returns()
func (_IdentityRegistry *IdentityRegistryTransactor) ChangeRegistrationFee(opts *bind.TransactOpts, newFee *big.Int) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "changeRegistrationFee", newFee)
}

// ChangeRegistrationFee is a paid mutator transaction binding the contract method 0x50050769.
//
// Solidity: function changeRegistrationFee(newFee uint256) returns()
func (_IdentityRegistry *IdentityRegistrySession) ChangeRegistrationFee(newFee *big.Int) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.ChangeRegistrationFee(&_IdentityRegistry.TransactOpts, newFee)
}

// ChangeRegistrationFee is a paid mutator transaction binding the contract method 0x50050769.
//
// Solidity: function changeRegistrationFee(newFee uint256) returns()
func (_IdentityRegistry *IdentityRegistryTransactorSession) ChangeRegistrationFee(newFee *big.Int) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.ChangeRegistrationFee(&_IdentityRegistry.TransactOpts, newFee)
}

// TransferCollectedFeeTo is a paid mutator transaction binding the contract method 0xe3252537.
//
// Solidity: function transferCollectedFeeTo(receiver address) returns(bool)
func (_IdentityRegistry *IdentityRegistryTransactor) TransferCollectedFeeTo(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "transferCollectedFeeTo", receiver)
}

// TransferCollectedFeeTo is a paid mutator transaction binding the contract method 0xe3252537.
//
// Solidity: function transferCollectedFeeTo(receiver address) returns(bool)
func (_IdentityRegistry *IdentityRegistrySession) TransferCollectedFeeTo(receiver common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.TransferCollectedFeeTo(&_IdentityRegistry.TransactOpts, receiver)
}

// TransferCollectedFeeTo is a paid mutator transaction binding the contract method 0xe3252537.
//
// Solidity: function transferCollectedFeeTo(receiver address) returns(bool)
func (_IdentityRegistry *IdentityRegistryTransactorSession) TransferCollectedFeeTo(receiver common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.TransferCollectedFeeTo(&_IdentityRegistry.TransactOpts, receiver)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_IdentityRegistry *IdentityRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_IdentityRegistry *IdentityRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.TransferOwnership(&_IdentityRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_IdentityRegistry *IdentityRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.TransferOwnership(&_IdentityRegistry.TransactOpts, newOwner)
}

// IdentityRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the IdentityRegistry contract.
type IdentityRegistryOwnershipTransferredIterator struct {
	Event *IdentityRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryOwnershipTransferred)
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
		it.Event = new(IdentityRegistryOwnershipTransferred)
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
func (it *IdentityRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the IdentityRegistry contract.
type IdentityRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_IdentityRegistry *IdentityRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IdentityRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IdentityRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryOwnershipTransferredIterator{contract: _IdentityRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_IdentityRegistry *IdentityRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IdentityRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IdentityRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryOwnershipTransferred)
				if err := _IdentityRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// IdentityRegistryRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the IdentityRegistry contract.
type IdentityRegistryRegisteredIterator struct {
	Event *IdentityRegistryRegistered // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryRegistered)
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
		it.Event = new(IdentityRegistryRegistered)
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
func (it *IdentityRegistryRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryRegistered represents a Registered event raised by the IdentityRegistry contract.
type IdentityRegistryRegistered struct {
	Identity common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0x2d3734a8e47ac8316e500ac231c90a6e1848ca2285f40d07eaa52005e4b3a0e9.
//
// Solidity: e Registered(identity indexed address)
func (_IdentityRegistry *IdentityRegistryFilterer) FilterRegistered(opts *bind.FilterOpts, identity []common.Address) (*IdentityRegistryRegisteredIterator, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _IdentityRegistry.contract.FilterLogs(opts, "Registered", identityRule)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryRegisteredIterator{contract: _IdentityRegistry.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0x2d3734a8e47ac8316e500ac231c90a6e1848ca2285f40d07eaa52005e4b3a0e9.
//
// Solidity: e Registered(identity indexed address)
func (_IdentityRegistry *IdentityRegistryFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *IdentityRegistryRegistered, identity []common.Address) (event.Subscription, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _IdentityRegistry.contract.WatchLogs(opts, "Registered", identityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryRegistered)
				if err := _IdentityRegistry.contract.UnpackLog(event, "Registered", log); err != nil {
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

// MathABI is the input ABI used to generate the binding from.
const MathABI = "[]"

// MathBin is the compiled bytecode used for deploying new contracts.
const MathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820b834e945bcdec8f7927c0c2a68f2c1f4614752ca04b678ba1801abca0d6ed6260029`

// DeployMath deploys a new Ethereum contract, binding an instance of Math to it.
func DeployMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Math, error) {
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// Math is an auto generated Go binding around an Ethereum contract.
type Math struct {
	MathCaller     // Read-only binding to the contract
	MathTransactor // Write-only binding to the contract
	MathFilterer   // Log filterer for contract events
}

// MathCaller is an auto generated read-only Go binding around an Ethereum contract.
type MathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MathSession struct {
	Contract     *Math             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MathCallerSession struct {
	Contract *MathCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MathTransactorSession struct {
	Contract     *MathTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathRaw is an auto generated low-level Go binding around an Ethereum contract.
type MathRaw struct {
	Contract *Math // Generic contract binding to access the raw methods on
}

// MathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MathCallerRaw struct {
	Contract *MathCaller // Generic read-only contract binding to access the raw methods on
}

// MathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MathTransactorRaw struct {
	Contract *MathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMath creates a new instance of Math, bound to a specific deployed contract.
func NewMath(address common.Address, backend bind.ContractBackend) (*Math, error) {
	contract, err := bindMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// NewMathCaller creates a new read-only instance of Math, bound to a specific deployed contract.
func NewMathCaller(address common.Address, caller bind.ContractCaller) (*MathCaller, error) {
	contract, err := bindMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MathCaller{contract: contract}, nil
}

// NewMathTransactor creates a new write-only instance of Math, bound to a specific deployed contract.
func NewMathTransactor(address common.Address, transactor bind.ContractTransactor) (*MathTransactor, error) {
	contract, err := bindMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MathTransactor{contract: contract}, nil
}

// NewMathFilterer creates a new log filterer instance of Math, bound to a specific deployed contract.
func NewMathFilterer(address common.Address, filterer bind.ContractFilterer) (*MathFilterer, error) {
	contract, err := bindMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MathFilterer{contract: contract}, nil
}

// bindMath binds a generic wrapper to an already deployed contract.
func bindMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Math.Contract.MathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Math.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.contract.Transact(opts, method, params...)
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x608060405234801561001057600080fd5b5060008054600160a060020a03191633179055610173806100326000396000f30060806040526004361061004b5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416638da5cb5b8114610050578063f2fde38b14610081575b600080fd5b34801561005c57600080fd5b506100656100a4565b60408051600160a060020a039092168252519081900360200190f35b34801561008d57600080fd5b506100a2600160a060020a03600435166100b3565b005b600054600160a060020a031681565b600054600160a060020a031633146100ca57600080fd5b600160a060020a03811615156100df57600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820de42094f2b951d27237c26ecde08ca4e127a70368d1d9a9c32eb30a984179ae20029`

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
