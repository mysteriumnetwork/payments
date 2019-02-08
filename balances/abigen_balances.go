// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package balances

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
const ERC20ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// ERC20Bin is the compiled bytecode used for deploying new contracts.
const ERC20Bin = `0x608060405234801561001057600080fd5b5061062d806100206000396000f3fe60806040526004361061008d5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b3811461009257806318160ddd146100df57806323b872dd14610106578063395093511461014957806370a0823114610182578063a457c2d7146101b5578063a9059cbb146101ee578063dd62ed3e14610227575b600080fd5b34801561009e57600080fd5b506100cb600480360360408110156100b557600080fd5b50600160a060020a038135169060200135610262565b604080519115158252519081900360200190f35b3480156100eb57600080fd5b506100f46102e0565b60408051918252519081900360200190f35b34801561011257600080fd5b506100cb6004803603606081101561012957600080fd5b50600160a060020a038135811691602081013590911690604001356102e6565b34801561015557600080fd5b506100cb6004803603604081101561016c57600080fd5b50600160a060020a0381351690602001356103af565b34801561018e57600080fd5b506100f4600480360360208110156101a557600080fd5b5035600160a060020a031661045f565b3480156101c157600080fd5b506100cb600480360360408110156101d857600080fd5b50600160a060020a03813516906020013561047a565b3480156101fa57600080fd5b506100cb6004803603604081101561021157600080fd5b50600160a060020a0381351690602001356104c5565b34801561023357600080fd5b506100f46004803603604081101561024a57600080fd5b50600160a060020a03813581169160200135166104db565b6000600160a060020a038316151561027957600080fd5b336000818152600160209081526040808320600160a060020a03881680855290835292819020869055805186815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a350600192915050565b60025490565b600160a060020a038316600090815260016020908152604080832033845290915281205461031a908363ffffffff61050616565b600160a060020a038516600090815260016020908152604080832033845290915290205561034984848461051b565b600160a060020a0384166000818152600160209081526040808320338085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b6000600160a060020a03831615156103c657600080fd5b336000908152600160209081526040808320600160a060020a03871684529091529020546103fa908363ffffffff6105e816565b336000818152600160209081526040808320600160a060020a0389168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a350600192915050565b600160a060020a031660009081526020819052604090205490565b6000600160a060020a038316151561049157600080fd5b336000908152600160209081526040808320600160a060020a03871684529091529020546103fa908363ffffffff61050616565b60006104d233848461051b565b50600192915050565b600160a060020a03918216600090815260016020908152604080832093909416825291909152205490565b60008282111561051557600080fd5b50900390565b600160a060020a038216151561053057600080fd5b600160a060020a038316600090815260208190526040902054610559908263ffffffff61050616565b600160a060020a03808516600090815260208190526040808220939093559084168152205461058e908263ffffffff6105e816565b600160a060020a038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b6000828201838110156105fa57600080fd5b939250505056fea165627a7a72305820a4f2199d497827ec00eec7d5de0217d8904730cf174875a557aeb2f98daca1630029`

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
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
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
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC20 *ERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, owner)
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
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20 *ERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20 *ERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
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

// ERC20AwareBalanceABI is the input ABI used to generate the binding from.
const ERC20AwareBalanceABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"ERC20Token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"erc20tokenAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ERC20AwareBalanceBin is the compiled bytecode used for deploying new contracts.
const ERC20AwareBalanceBin = `0x608060405234801561001057600080fd5b506040516020806101278339810180604052602081101561003057600080fd5b505160008054600160a060020a03909216600160a060020a031990921691909117905560c6806100616000396000f3fe608060405260043610603e5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416637a80760e81146043575b600080fd5b348015604e57600080fd5b506055607e565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b60005473ffffffffffffffffffffffffffffffffffffffff168156fea165627a7a723058206779e9c6f3c42c6a895631d862888d82433f8e4a72c00633442047e1594d1b390029`

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
const ERC20AwareRegistryBin = `0x608060405234801561001057600080fd5b506040516020806101278339810180604052602081101561003057600080fd5b505160008054600160a060020a03909216600160a060020a031990921691909117905560c6806100616000396000f3fe608060405260043610603e5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416637a80760e81146043575b600080fd5b348015604e57600080fd5b506055607e565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b60005473ffffffffffffffffffffffffffffffffffffffff168156fea165627a7a7230582014c7c327dc3b276cb2bb339620678fb8b00275b1e6711d3d97ad8fbe6f56f9a60029`

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
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
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
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) constant returns(uint256)
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
// Solidity: function balanceOf(address who) constant returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(who common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) constant returns(uint256)
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
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
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

// IdentityBalancesABI is the input ABI used to generate the binding from.
const IdentityBalancesABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERC20Token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"topUp\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalBalance\",\"type\":\"uint256\"}],\"name\":\"ToppedUp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalBalance\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"}]"

// IdentityBalancesBin is the compiled bytecode used for deploying new contracts.
const IdentityBalancesBin = `0x608060405234801561001057600080fd5b5060405160208061059e8339810180604052602081101561003057600080fd5b505160008054600160a060020a03909216600160a060020a031990921691909117905561053c806100626000396000f3fe6080604052600436106100615763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166327e235e381146100665780635ebfdfc6146100ab5780637a80760e146100fe578063d6f7ddf91461012f575b600080fd5b34801561007257600080fd5b506100996004803603602081101561008957600080fd5b5035600160a060020a0316610168565b60408051918252519081900360200190f35b3480156100b757600080fd5b506100ea600480360360808110156100ce57600080fd5b5080359060ff602082013516906040810135906060013561017a565b604080519115158252519081900360200190f35b34801561010a57600080fd5b506101136103dc565b60408051600160a060020a039092168252519081900360200190f35b34801561013b57600080fd5b506100ea6004803603604081101561015257600080fd5b50600160a060020a0381351690602001356103eb565b60016020526000908152604090205481565b600060606040805190810160405280601181526020017f576974686472617720726571756573743a000000000000000000000000000000815250866040516020018083805190602001908083835b602083106101e75780518252601f1990920191602091820191016101c8565b51815160209384036101000a60001901801990921691161790529201938452506040805180850381528483018083528151828501206000918290528387018085525260ff8c166060870152608086018b905260a086018a9052915190965090945060019360c08082019450601f19830192918290030190855afa158015610272573d6000803e3d6000fd5b5050604051601f190151915050600160a060020a038116151561029457600080fd5b600087116102a157600080fd5b600160a060020a0381166000908152600160205260409020548711156102c657600080fd5b600160a060020a03808216600090815260016020908152604080832080548c90039055825481517fa9059cbb000000000000000000000000000000000000000000000000000000008152336004820152602481018d9052915194169363a9059cbb93604480840194938390030190829087803b15801561034557600080fd5b505af1158015610359573d6000803e3d6000fd5b505050506040513d602081101561036f57600080fd5b5051151561037c57600080fd5b600160a060020a0381166000818152600160209081526040918290205482518b81529182015281517f92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6929181900390910190a25060019695505050505050565b600054600160a060020a031681565b60008082116103f957600080fd5b600160a060020a038084166000908152600160209081526040808320805487019055825481517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526044810188905291519416936323b872dd93606480840194938390030190829087803b15801561047d57600080fd5b505af1158015610491573d6000803e3d6000fd5b505050506040513d60208110156104a757600080fd5b505115156104b457600080fd5b600160a060020a0383166000818152600160209081526040918290205482518681529182015281517fbdde76a89d276b6d334c784be5c2d00d6c2219d11a2f2c80e00b85144845ab4d929181900390910190a25060019291505056fea165627a7a723058204a4e9437f6562bcbf5e85011b5fb4c9d5fd92afd4a7f70d412ab741ce9a83b3b0029`

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
// Solidity: function balances(address ) constant returns(uint256)
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
// Solidity: function balances(address ) constant returns(uint256)
func (_IdentityBalances *IdentityBalancesSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _IdentityBalances.Contract.Balances(&_IdentityBalances.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) constant returns(uint256)
func (_IdentityBalances *IdentityBalancesCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _IdentityBalances.Contract.Balances(&_IdentityBalances.CallOpts, arg0)
}

// TopUp is a paid mutator transaction binding the contract method 0xd6f7ddf9.
//
// Solidity: function topUp(address identity, uint256 amount) returns(bool)
func (_IdentityBalances *IdentityBalancesTransactor) TopUp(opts *bind.TransactOpts, identity common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IdentityBalances.contract.Transact(opts, "topUp", identity, amount)
}

// TopUp is a paid mutator transaction binding the contract method 0xd6f7ddf9.
//
// Solidity: function topUp(address identity, uint256 amount) returns(bool)
func (_IdentityBalances *IdentityBalancesSession) TopUp(identity common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IdentityBalances.Contract.TopUp(&_IdentityBalances.TransactOpts, identity, amount)
}

// TopUp is a paid mutator transaction binding the contract method 0xd6f7ddf9.
//
// Solidity: function topUp(address identity, uint256 amount) returns(bool)
func (_IdentityBalances *IdentityBalancesTransactorSession) TopUp(identity common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IdentityBalances.Contract.TopUp(&_IdentityBalances.TransactOpts, identity, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5ebfdfc6.
//
// Solidity: function withdraw(uint256 amount, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_IdentityBalances *IdentityBalancesTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IdentityBalances.contract.Transact(opts, "withdraw", amount, v, r, s)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5ebfdfc6.
//
// Solidity: function withdraw(uint256 amount, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_IdentityBalances *IdentityBalancesSession) Withdraw(amount *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IdentityBalances.Contract.Withdraw(&_IdentityBalances.TransactOpts, amount, v, r, s)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5ebfdfc6.
//
// Solidity: function withdraw(uint256 amount, uint8 v, bytes32 r, bytes32 s) returns(bool)
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
// Solidity: event ToppedUp(address indexed identity, uint256 amount, uint256 totalBalance)
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
// Solidity: event ToppedUp(address indexed identity, uint256 amount, uint256 totalBalance)
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
// Solidity: event Withdrawn(address indexed identity, uint256 amount, uint256 totalBalance)
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
// Solidity: event Withdrawn(address indexed identity, uint256 amount, uint256 totalBalance)
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

// MathABI is the input ABI used to generate the binding from.
const MathABI = "[]"

// MathBin is the compiled bytecode used for deploying new contracts.
const MathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a7230582010775f97cd9c7fd51f186d5d9fe7fcdd078b433e889e907c87e4eab2dbed13d00029`

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

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a7230582003d7c93596189e7e13f3648db06d1403c31116ffcce7cf93c0ab07bc8a36cc920029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}
