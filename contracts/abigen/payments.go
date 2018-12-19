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

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// IERC20Bin is the compiled bytecode used for deploying new contracts.
const IERC20Bin = `0x`

// DeployIERC20 deploys a new Ethereum contract, binding an instance of IERC20 to it.
func DeployIERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(who common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_IERC20 *IERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
const IdentityPromisesABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"registrationFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"receiverAndSigns\",\"type\":\"bytes32\"},{\"name\":\"extraDataHash\",\"type\":\"bytes32\"},{\"name\":\"seq\",\"type\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"sender_R\",\"type\":\"bytes32\"},{\"name\":\"sender_S\",\"type\":\"bytes32\"},{\"name\":\"receiver_R\",\"type\":\"bytes32\"},{\"name\":\"receiver_S\",\"type\":\"bytes32\"}],\"name\":\"clearPromise\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"changeRegistrationFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERC20Token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"}],\"name\":\"getPublicKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pubKeyPart1\",\"type\":\"bytes32\"},{\"name\":\"pubKeyPart2\",\"type\":\"bytes32\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"RegisterIdentity\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"}],\"name\":\"isRegistered\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"topUp\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"transferCollectedFeeTo\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"collectedFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"clearedPromises\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"erc20address\",\"type\":\"address\"},{\"name\":\"registrationFee\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"seqNo\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PromiseCleared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalBalance\",\"type\":\"uint256\"}],\"name\":\"ToppedUp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalBalance\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"identity\",\"type\":\"address\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// IdentityPromisesBin is the compiled bytecode used for deploying new contracts.
const IdentityPromisesBin = `0x608060405234801561001057600080fd5b5060405160408061139e833981016040819052815160209092015160008054600160a060020a031916331780825591928492839285928492600160a060020a031691907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a360018054600160a060020a03909216600160a060020a0319909216919091179055600355505050506112ef806100af6000396000f3006080604052600436106100f05763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166314c44e0981146100f557806327e235e31461011c5780634fe0d3321461013d578063500507691461017e5780635ebfdfc614610198578063715018a6146101bc5780637a80760e146101d1578063857cdbb8146102025780638da5cb5b1461023c5780638f32d59b14610251578063a17fc7f314610266578063c3c5a5471461028d578063d6f7ddf9146102ae578063e3252537146102d2578063e811f50a146102f3578063f2fde38b14610308578063f6c1a23e14610329575b600080fd5b34801561010157600080fd5b5061010a610350565b60408051918252519081900360200190f35b34801561012857600080fd5b5061010a600160a060020a0360043516610356565b34801561014957600080fd5b5061016a60043560243560443560643560843560a43560c43560e435610368565b604080519115158252519081900360200190f35b34801561018a57600080fd5b50610196600435610768565b005b3480156101a457600080fd5b5061016a60043560ff60243516604435606435610780565b3480156101c857600080fd5b50610196610a3d565b3480156101dd57600080fd5b506101e6610aa7565b60408051600160a060020a039092168252519081900360200190f35b34801561020e57600080fd5b50610223600160a060020a0360043516610ab6565b6040805192835260208301919091528051918290030190f35b34801561024857600080fd5b506101e6610ad9565b34801561025d57600080fd5b5061016a610ae8565b34801561027257600080fd5b5061016a60043560243560ff60443516606435608435610af9565b34801561029957600080fd5b5061016a600160a060020a0360043516610e37565b3480156102ba57600080fd5b5061016a600160a060020a0360043516602435610e6a565b3480156102de57600080fd5b5061016a600160a060020a0360043516610f90565b3480156102ff57600080fd5b5061010a611073565b34801561031457600080fd5b50610196600160a060020a0360043516611079565b34801561033557600080fd5b5061010a600160a060020a0360043581169060243516611098565b60035481565b60056020526000908152604090205481565b600080600080600080600061037c8f6110b5565b9550955095506040805190810160405280600e81526020017f497373756572207072656669783a0000000000000000000000000000000000008152508e878f8f6040516020018086805190602001908083835b602083106103ee5780518252601f1990920191602091820191016103cf565b51815160209384036101000a6000190180199092169116179052920196875250600160a060020a03949094166c010000000000000000000000000285850152506034840191909152605480840191909152604080518085039092018252607490930192839052805190935082918401908083835b602083106104815780518252601f199092019160209182019101610462565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209250600183868d8d604051600081526020016040526040518085600019166000191681526020018460ff1660ff1681526020018360001916600019168152602001826000191660001916815260200194505050505060206040516020810390808403906000865af1158015610527573d6000803e3d6000fd5b50505060206040510351915060016040805190810160405280601081526020017f5265636569766572207072656669783a0000000000000000000000000000000081525084846040516020018084805190602001908083835b6020831061059f5780518252601f199092019160209182019101610580565b51815160209384036101000a6000190180199092169116179052920194855250600160a060020a03929092166c0100000000000000000000000002838301525060408051601481850301815260349093019081905282519293509182918401908083835b602083106106225780518252601f199092019160209182019101610603565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020858b8b604051600081526020016040526040518085600019166000191681526020018460ff1660ff1681526020018360001916600019168152602001826000191660001916815260200194505050505060206040516020810390808403906000865af11580156106c3573d6000803e3d6000fd5b5050604051601f1901519150506000600160a060020a038316116106e657600080fd5b6000600160a060020a038216116106fc57600080fd5b600160a060020a038181169087161461071457600080fd5b61071d82610e37565b151561072857600080fd5b61073186610e37565b151561073c57600080fd5b60008c1161074957600080fd5b61075586838f8f6110f0565b9f9e505050505050505050505050505050565b610770610ae8565b151561077b57600080fd5b600355565b6000606060006040805190810160405280601181526020017f576974686472617720726571756573743a000000000000000000000000000000815250876040516020018083805190602001908083835b602083106107ef5780518252601f1990920191602091820191016107d0565b51815160209384036101000a600019018019909216911617905292019384525060408051808503815293820190819052835193965060019450869390925082918401908083835b602083106108555780518252601f199092019160209182019101610836565b51815160209384036101000a60001901801990921691161790526040805192909401829003822060008084528383018087529190915260ff8e1683860152606083018d9052608083018c9052935160a08084019750919550601f1981019492819003909101925090865af11580156108d1573d6000803e3d6000fd5b5050604051601f1901519150506000600160a060020a038216116108f457600080fd5b6000871161090157600080fd5b600160a060020a03811660009081526005602052604090205487111561092657600080fd5b600160a060020a03808216600090815260056020908152604080832080548c9003905560015481517fa9059cbb000000000000000000000000000000000000000000000000000000008152336004820152602481018d9052915194169363a9059cbb93604480840194938390030190829087803b1580156109a657600080fd5b505af11580156109ba573d6000803e3d6000fd5b505050506040513d60208110156109d057600080fd5b505115156109dd57600080fd5b600160a060020a0381166000818152600560209081526040918290205482518b81529182015281517f92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6929181900390910190a25060019695505050505050565b610a45610ae8565b1515610a5057600080fd5b60008054604051600160a060020a03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a36000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600154600160a060020a031681565b600160a060020a0316600090815260026020526040902080546001909101549091565b600054600160a060020a031690565b600054600160a060020a0316331490565b604080516020808201889052818301879052825180830384018152606090920192839052815160009384938493909282918401908083835b60208310610b505780518252601f199092019160209182019101610b31565b51815160209384036101000a600019018019909216911617905260408051929094018290038220828501855260108084527f5265676973746572207072656669783a000000000000000000000000000000008484019081529551919950600197509295508e948e94509101918291908083835b60208310610be25780518252601f199092019160209182019101610bc3565b51815160209384036101000a6000190180199092169116179052920194855250838101929092525060408051808403830181529281019081905282519293509182918401908083835b60208310610c4a5780518252601f199092019160209182019101610c2b565b51815160209384036101000a60001901801990921691161790526040805192909401829003822060008084528383018087529190915260ff8e1683860152606083018d9052608083018c9052935160a08084019750919550601f1981019492819003909101925090865af1158015610cc6573d6000803e3d6000fd5b5050604051601f1901519150506000600160a060020a03821611610ce957600080fd5b600160a060020a0382811690821614610d0157600080fd5b610d0a81610e37565b15610d1457600080fd5b600380546004805490910181556040805180820182528b815260208082018c8152600160a060020a038088166000908152600284528581209451855591516001948501559254955484517f23b872dd00000000000000000000000000000000000000000000000000000000815233968101969096523060248701526044860152925194909116936323b872dd93606480820194918390030190829087803b158015610dbe57600080fd5b505af1158015610dd2573d6000803e3d6000fd5b505050506040513d6020811015610de857600080fd5b50511515610df557600080fd5b604051600160a060020a038216907f2d3734a8e47ac8316e500ac231c90a6e1848ca2285f40d07eaa52005e4b3a0e990600090a2506001979650505050505050565b600160a060020a0381166000908152600260205260408120805482108015610e63575060018101546000105b9392505050565b6000808211610e7857600080fd5b600160a060020a03808416600090815260056020908152604080832080548701905560015481517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526044810188905291519416936323b872dd93606480840194938390030190829087803b158015610efd57600080fd5b505af1158015610f11573d6000803e3d6000fd5b505050506040513d6020811015610f2757600080fd5b50511515610f3457600080fd5b600160a060020a0383166000818152600560209081526040918290205482518681529182015281517fbdde76a89d276b6d334c784be5c2d00d6c2219d11a2f2c80e00b85144845ab4d929181900390910190a250600192915050565b600080610f9b610ae8565b1515610fa657600080fd5b600454600010610fb557600080fd5b50600480546000808355600154604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a038881169682019690965260248101859052905193949091169263a9059cbb92604480840193602093929083900390910190829087803b15801561103357600080fd5b505af1158015611047573d6000803e3d6000fd5b505050506040513d602081101561105d57600080fd5b5051151561106a57600080fd5b50600192915050565b60045481565b611081610ae8565b151561108c57600080fd5b611095816111c5565b50565b600660209081526000928352604080842090915290825290205481565b907f0100000000000000000000000000000000000000000000000000000000000000600a83901a810281900491600b84901a82029190910490565b600160a060020a0380841660009081526006602090815260408083209388168352929052908120548190841161112557600080fd5b600160a060020a038086166000908152600660209081526040808320938a16835292905220849055611158858785611242565b905082811461116657600080fd5b85600160a060020a031685600160a060020a03167f3d35f4a81f47a50c001e0e834334665e3a64680385c0a8d2fdfe2e0d209645768686604051808381526020018281526020019250505060405180910390a350600195945050505050565b600160a060020a03811615156111da57600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60008080831161125157600080fd5b600160a060020a03851660009081526005602052604090205461127490846112ad565b600160a060020a039586166000908152600560205260408082208054849003905595909616865293909420805484019055509092915050565b60008183106112bc5781610e63565b50909190505600a165627a7a72305820d785f760de027fb8fa797cf9812b4f5c12e11b25d6139bdf1121396394e657930029`

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
