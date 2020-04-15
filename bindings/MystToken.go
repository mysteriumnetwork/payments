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

// MystTokenABI is the input ABI used to generate the binding from.
const MystTokenABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Upgrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"UpgradeAgentSet\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mintAgents\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mintingFinished\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalUpgraded\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"upgradeAgent\",\"outputs\":[{\"internalType\":\"contractUpgradeAgent\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"upgradeMaster\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"addApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"subApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"setMintAgent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"upgrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"setUpgradeAgent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getUpgradeState\",\"outputs\":[{\"internalType\":\"enumMystToken.UpgradeState\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"master\",\"type\":\"address\"}],\"name\":\"setUpgradeMaster\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MystTokenBin is the compiled bytecode used for deploying new contracts.
var MystTokenBin = "0x60806040526001805460ff60a01b1916905534801561001d57600080fd5b5060018054600280546001600160a01b0319908116339081179092559182168117909116811782556000908152600560205260409020805460ff1916909117905561112e8061006d6000396000f3fe608060405234801561001057600080fd5b50600436106101585760003560e01c806370a08231116100c3578063c752ff621161007c578063c752ff6214610415578063d7e7088a1461041d578063dd62ed3e14610443578063e2301d0214610471578063f2fde38b1461049d578063ffeb7d75146104c357610158565b806370a082311461035b5780638444b391146103815780638da5cb5b146103ad57806395d89b41146103b5578063a9059cbb146103bd578063ac3cb72c146103e957610158565b806340c10f191161011557806340c10f191461029057806342c1867b146102be57806343214675146102e457806345977d03146103125780635de4ccb01461032f578063600440cb1461035357610158565b806305d2035b1461015d57806306fdde0314610179578063095ea7b3146101f657806318160ddd1461022257806323b872dd1461023c578063313ce56714610272575b600080fd5b6101656104e9565b604080519115158252519081900360200190f35b6101816104f9565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101bb5781810151838201526020016101a3565b50505050905090810190601f1680156101e85780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101656004803603604081101561020c57600080fd5b506001600160a01b038135169060200135610529565b61022a6105cd565b60408051918252519081900360200190f35b6101656004803603606081101561025257600080fd5b506001600160a01b038135811691602081013590911690604001356105d3565b61027a6106d0565b6040805160ff9092168252519081900360200190f35b6102bc600480360360408110156102a657600080fd5b506001600160a01b0381351690602001356106d5565b005b610165600480360360208110156102d457600080fd5b50356001600160a01b0316610895565b6102bc600480360360408110156102fa57600080fd5b506001600160a01b03813516906020013515156108aa565b6102bc6004803603602081101561032857600080fd5b5035610903565b610337610a48565b604080516001600160a01b039092168252519081900360200190f35b610337610a57565b61022a6004803603602081101561037157600080fd5b50356001600160a01b0316610a66565b610389610a81565b6040518082600481111561039957fe5b60ff16815260200191505060405180910390f35b610337610ab2565b610181610ac1565b610165600480360360408110156103d357600080fd5b506001600160a01b038135169060200135610ae2565b610165600480360360408110156103ff57600080fd5b506001600160a01b038135169060200135610b98565b61022a610c3e565b6102bc6004803603602081101561043357600080fd5b50356001600160a01b0316610c44565b61022a6004803603604081101561045957600080fd5b506001600160a01b0381358116916020013516610e4c565b6101656004803603604081101561048757600080fd5b506001600160a01b038135169060200135610e77565b6102bc600480360360208110156104b357600080fd5b50356001600160a01b0316610f4e565b6102bc600480360360208110156104d957600080fd5b50356001600160a01b0316610f93565b600154600160a01b900460ff1681565b604051806040016040528060148152602001732a32b9ba1026bcb9ba32b934bab6903a37b5b2b760611b81525081565b6000811580159061055c57503360009081526007602090815260408083206001600160a01b038716845290915290205415155b1561056657600080fd5b3360008181526007602090815260408083206001600160a01b03881680855290835292819020869055805186815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a350600192915050565b60005481565b6001600160a01b038084166000908152600760209081526040808320338452825280832054938616835260069091528120549091906106129084610fdf565b6001600160a01b038086166000908152600660205260408082209390935590871681522054610641908461104c565b6001600160a01b038616600090815260066020526040902055610664818461104c565b6001600160a01b03808716600081815260076020908152604080832033845282529182902094909455805187815290519288169391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a3506001949350505050565b600881565b3360009081526005602052604090205460ff166106f157600080fd5b600154600160a01b900460ff161561070857600080fd5b8061071257600080fd5b60005473__SafeMathLib___________________________6366098d4f9091836040518363ffffffff1660e01b8152600401808381526020018281526020019250505060206040518083038186803b15801561076d57600080fd5b505af4158015610781573d6000803e3d6000fd5b505050506040513d602081101561079757600080fd5b505160009081556001600160a01b0383168152600660209081526040918290205482516366098d4f60e01b8152600481019190915260248101849052915173__SafeMathLib___________________________926366098d4f926044808301939192829003018186803b15801561080d57600080fd5b505af4158015610821573d6000803e3d6000fd5b505050506040513d602081101561083757600080fd5b50516001600160a01b03831660008181526006602090815260409182902093909355805191825291810183905281517f30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe929181900390910190a15050565b60056020526000908152604090205460ff1681565b6001546001600160a01b031633146108c157600080fd5b600154600160a01b900460ff16156108d857600080fd5b6001600160a01b03919091166000908152600560205260409020805460ff1916911515919091179055565b600061090d610a81565b9050600381600481111561091d57fe5b14806109345750600481600481111561093257fe5b145b61093d57600080fd5b8161094757600080fd5b33600090815260066020526040902054610961908361104c565b336000908152600660205260408120919091555461097f908361104c565b60005560045461098f9083610fdf565b60049081556003546040805163753e88e560e01b8152339381019390935260248301859052516001600160a01b039091169163753e88e591604480830192600092919082900301818387803b1580156109e757600080fd5b505af11580156109fb573d6000803e3d6000fd5b50506003546040805186815290516001600160a01b0390921693503392507f7e5c344a8141a805725cb476f76c6953b842222b967edd1f78ddb6e8b3f397ac919081900360200190a35050565b6003546001600160a01b031681565b6002546001600160a01b031681565b6001600160a01b031660009081526006602052604090205490565b6003546000906001600160a01b0316610a9c57506002610aaf565b600454610aab57506003610aaf565b5060045b90565b6001546001600160a01b031681565b60405180604001604052806005815260200164135654d51560da1b81525081565b6000604036604414610af357600080fd5b33600090815260066020526040902054610b0d908461104c565b33600090815260066020526040808220929092556001600160a01b03861681522054610b399084610fdf565b6001600160a01b0385166000818152600660209081526040918290209390935580518681529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35060019392505050565b6000604036604414610ba957600080fd5b3360009081526007602090815260408083206001600160a01b0388168452909152902054610bd78185610fdf565b3360008181526007602090815260408083206001600160a01b038b168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a3506001949350505050565b60045481565b6002546001600160a01b03163314610c8d5760405162461bcd60e51b815260040180806020018281038252602a8152602001806110aa602a913960400191505060405180910390fd5b6001600160a01b038116610ca057600080fd5b6004610caa610a81565b6004811115610cb557fe5b1415610cf25760405162461bcd60e51b81526004018080602001828103825260268152602001806110d46026913960400191505060405180910390fd5b600380546001600160a01b0319166001600160a01b038381169190911791829055604080516330e9ebd360e11b8152905192909116916361d3d7a691600480820192602092909190829003018186803b158015610d4e57600080fd5b505afa158015610d62573d6000803e3d6000fd5b505050506040513d6020811015610d7857600080fd5b5051610d8357600080fd5b600054600360009054906101000a90046001600160a01b03166001600160a01b0316634b2ba0dd6040518163ffffffff1660e01b815260040160206040518083038186803b158015610dd457600080fd5b505afa158015610de8573d6000803e3d6000fd5b505050506040513d6020811015610dfe57600080fd5b505114610e0a57600080fd5b600354604080516001600160a01b039092168252517f7845d5aa74cc410e35571258d954f23b82276e160fe8c188fa80566580f279cc9181900360200190a150565b6001600160a01b03918216600090815260076020908152604080832093909416825291909152205490565b6000604036604414610e8857600080fd5b3360009081526007602090815260408083206001600160a01b038816845290915290205480841115610edd573360009081526007602090815260408083206001600160a01b0389168452909152812055610ee7565b610bd7818561104c565b3360008181526007602090815260408083206001600160a01b038a168085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a3506001949350505050565b6001546001600160a01b03163314610f6557600080fd5b6001600160a01b03811615610f9057600180546001600160a01b0319166001600160a01b0383161790555b50565b6001600160a01b038116610fa657600080fd5b6002546001600160a01b03163314610fbd57600080fd5b600280546001600160a01b0319166001600160a01b0392909216919091179055565b6000828201838110801590610ff45750828110155b611045576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b6000828211156110a3576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b5090039056fe4f6e6c792061206d61737465722063616e2064657369676e61746520746865206e657874206167656e74557067726164652068617320616c726561647920626567756e20666f7220616e206167656e74a265627a7a72315820984fd1cfb4f5d3ddc2a400584563c99e06d1b838e894379adf7e27153d5fae3164736f6c634300050c0032"

// DeployMystToken deploys a new Ethereum contract, binding an instance of MystToken to it.
func DeployMystToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MystToken, error) {
	parsed, err := abi.JSON(strings.NewReader(MystTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MystTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MystToken{MystTokenCaller: MystTokenCaller{contract: contract}, MystTokenTransactor: MystTokenTransactor{contract: contract}, MystTokenFilterer: MystTokenFilterer{contract: contract}}, nil
}

// MystToken is an auto generated Go binding around an Ethereum contract.
type MystToken struct {
	MystTokenCaller     // Read-only binding to the contract
	MystTokenTransactor // Write-only binding to the contract
	MystTokenFilterer   // Log filterer for contract events
}

// MystTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type MystTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MystTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MystTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MystTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MystTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MystTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MystTokenSession struct {
	Contract     *MystToken        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MystTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MystTokenCallerSession struct {
	Contract *MystTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MystTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MystTokenTransactorSession struct {
	Contract     *MystTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MystTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type MystTokenRaw struct {
	Contract *MystToken // Generic contract binding to access the raw methods on
}

// MystTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MystTokenCallerRaw struct {
	Contract *MystTokenCaller // Generic read-only contract binding to access the raw methods on
}

// MystTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MystTokenTransactorRaw struct {
	Contract *MystTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMystToken creates a new instance of MystToken, bound to a specific deployed contract.
func NewMystToken(address common.Address, backend bind.ContractBackend) (*MystToken, error) {
	contract, err := bindMystToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MystToken{MystTokenCaller: MystTokenCaller{contract: contract}, MystTokenTransactor: MystTokenTransactor{contract: contract}, MystTokenFilterer: MystTokenFilterer{contract: contract}}, nil
}

// NewMystTokenCaller creates a new read-only instance of MystToken, bound to a specific deployed contract.
func NewMystTokenCaller(address common.Address, caller bind.ContractCaller) (*MystTokenCaller, error) {
	contract, err := bindMystToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MystTokenCaller{contract: contract}, nil
}

// NewMystTokenTransactor creates a new write-only instance of MystToken, bound to a specific deployed contract.
func NewMystTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*MystTokenTransactor, error) {
	contract, err := bindMystToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MystTokenTransactor{contract: contract}, nil
}

// NewMystTokenFilterer creates a new log filterer instance of MystToken, bound to a specific deployed contract.
func NewMystTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*MystTokenFilterer, error) {
	contract, err := bindMystToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MystTokenFilterer{contract: contract}, nil
}

// bindMystToken binds a generic wrapper to an already deployed contract.
func bindMystToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MystTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MystToken *MystTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MystToken.Contract.MystTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MystToken *MystTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MystToken.Contract.MystTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MystToken *MystTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MystToken.Contract.MystTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MystToken *MystTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MystToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MystToken *MystTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MystToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MystToken *MystTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MystToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256 remaining)
func (_MystToken *MystTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MystToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256 remaining)
func (_MystToken *MystTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _MystToken.Contract.Allowance(&_MystToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256 remaining)
func (_MystToken *MystTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _MystToken.Contract.Allowance(&_MystToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 balance)
func (_MystToken *MystTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MystToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 balance)
func (_MystToken *MystTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _MystToken.Contract.BalanceOf(&_MystToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 balance)
func (_MystToken *MystTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _MystToken.Contract.BalanceOf(&_MystToken.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MystToken *MystTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _MystToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MystToken *MystTokenSession) Decimals() (uint8, error) {
	return _MystToken.Contract.Decimals(&_MystToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MystToken *MystTokenCallerSession) Decimals() (uint8, error) {
	return _MystToken.Contract.Decimals(&_MystToken.CallOpts)
}

// GetUpgradeState is a free data retrieval call binding the contract method 0x8444b391.
//
// Solidity: function getUpgradeState() constant returns(uint8)
func (_MystToken *MystTokenCaller) GetUpgradeState(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _MystToken.contract.Call(opts, out, "getUpgradeState")
	return *ret0, err
}

// GetUpgradeState is a free data retrieval call binding the contract method 0x8444b391.
//
// Solidity: function getUpgradeState() constant returns(uint8)
func (_MystToken *MystTokenSession) GetUpgradeState() (uint8, error) {
	return _MystToken.Contract.GetUpgradeState(&_MystToken.CallOpts)
}

// GetUpgradeState is a free data retrieval call binding the contract method 0x8444b391.
//
// Solidity: function getUpgradeState() constant returns(uint8)
func (_MystToken *MystTokenCallerSession) GetUpgradeState() (uint8, error) {
	return _MystToken.Contract.GetUpgradeState(&_MystToken.CallOpts)
}

// MintAgents is a free data retrieval call binding the contract method 0x42c1867b.
//
// Solidity: function mintAgents(address ) constant returns(bool)
func (_MystToken *MystTokenCaller) MintAgents(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MystToken.contract.Call(opts, out, "mintAgents", arg0)
	return *ret0, err
}

// MintAgents is a free data retrieval call binding the contract method 0x42c1867b.
//
// Solidity: function mintAgents(address ) constant returns(bool)
func (_MystToken *MystTokenSession) MintAgents(arg0 common.Address) (bool, error) {
	return _MystToken.Contract.MintAgents(&_MystToken.CallOpts, arg0)
}

// MintAgents is a free data retrieval call binding the contract method 0x42c1867b.
//
// Solidity: function mintAgents(address ) constant returns(bool)
func (_MystToken *MystTokenCallerSession) MintAgents(arg0 common.Address) (bool, error) {
	return _MystToken.Contract.MintAgents(&_MystToken.CallOpts, arg0)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_MystToken *MystTokenCaller) MintingFinished(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MystToken.contract.Call(opts, out, "mintingFinished")
	return *ret0, err
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_MystToken *MystTokenSession) MintingFinished() (bool, error) {
	return _MystToken.Contract.MintingFinished(&_MystToken.CallOpts)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_MystToken *MystTokenCallerSession) MintingFinished() (bool, error) {
	return _MystToken.Contract.MintingFinished(&_MystToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MystToken *MystTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MystToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MystToken *MystTokenSession) Name() (string, error) {
	return _MystToken.Contract.Name(&_MystToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MystToken *MystTokenCallerSession) Name() (string, error) {
	return _MystToken.Contract.Name(&_MystToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MystToken *MystTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MystToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MystToken *MystTokenSession) Owner() (common.Address, error) {
	return _MystToken.Contract.Owner(&_MystToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MystToken *MystTokenCallerSession) Owner() (common.Address, error) {
	return _MystToken.Contract.Owner(&_MystToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MystToken *MystTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MystToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MystToken *MystTokenSession) Symbol() (string, error) {
	return _MystToken.Contract.Symbol(&_MystToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MystToken *MystTokenCallerSession) Symbol() (string, error) {
	return _MystToken.Contract.Symbol(&_MystToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MystToken *MystTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MystToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MystToken *MystTokenSession) TotalSupply() (*big.Int, error) {
	return _MystToken.Contract.TotalSupply(&_MystToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MystToken *MystTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _MystToken.Contract.TotalSupply(&_MystToken.CallOpts)
}

// TotalUpgraded is a free data retrieval call binding the contract method 0xc752ff62.
//
// Solidity: function totalUpgraded() constant returns(uint256)
func (_MystToken *MystTokenCaller) TotalUpgraded(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MystToken.contract.Call(opts, out, "totalUpgraded")
	return *ret0, err
}

// TotalUpgraded is a free data retrieval call binding the contract method 0xc752ff62.
//
// Solidity: function totalUpgraded() constant returns(uint256)
func (_MystToken *MystTokenSession) TotalUpgraded() (*big.Int, error) {
	return _MystToken.Contract.TotalUpgraded(&_MystToken.CallOpts)
}

// TotalUpgraded is a free data retrieval call binding the contract method 0xc752ff62.
//
// Solidity: function totalUpgraded() constant returns(uint256)
func (_MystToken *MystTokenCallerSession) TotalUpgraded() (*big.Int, error) {
	return _MystToken.Contract.TotalUpgraded(&_MystToken.CallOpts)
}

// UpgradeAgent is a free data retrieval call binding the contract method 0x5de4ccb0.
//
// Solidity: function upgradeAgent() constant returns(address)
func (_MystToken *MystTokenCaller) UpgradeAgent(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MystToken.contract.Call(opts, out, "upgradeAgent")
	return *ret0, err
}

// UpgradeAgent is a free data retrieval call binding the contract method 0x5de4ccb0.
//
// Solidity: function upgradeAgent() constant returns(address)
func (_MystToken *MystTokenSession) UpgradeAgent() (common.Address, error) {
	return _MystToken.Contract.UpgradeAgent(&_MystToken.CallOpts)
}

// UpgradeAgent is a free data retrieval call binding the contract method 0x5de4ccb0.
//
// Solidity: function upgradeAgent() constant returns(address)
func (_MystToken *MystTokenCallerSession) UpgradeAgent() (common.Address, error) {
	return _MystToken.Contract.UpgradeAgent(&_MystToken.CallOpts)
}

// UpgradeMaster is a free data retrieval call binding the contract method 0x600440cb.
//
// Solidity: function upgradeMaster() constant returns(address)
func (_MystToken *MystTokenCaller) UpgradeMaster(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MystToken.contract.Call(opts, out, "upgradeMaster")
	return *ret0, err
}

// UpgradeMaster is a free data retrieval call binding the contract method 0x600440cb.
//
// Solidity: function upgradeMaster() constant returns(address)
func (_MystToken *MystTokenSession) UpgradeMaster() (common.Address, error) {
	return _MystToken.Contract.UpgradeMaster(&_MystToken.CallOpts)
}

// UpgradeMaster is a free data retrieval call binding the contract method 0x600440cb.
//
// Solidity: function upgradeMaster() constant returns(address)
func (_MystToken *MystTokenCallerSession) UpgradeMaster() (common.Address, error) {
	return _MystToken.Contract.UpgradeMaster(&_MystToken.CallOpts)
}

// AddApproval is a paid mutator transaction binding the contract method 0xac3cb72c.
//
// Solidity: function addApproval(address _spender, uint256 _addedValue) returns(bool success)
func (_MystToken *MystTokenTransactor) AddApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _MystToken.contract.Transact(opts, "addApproval", _spender, _addedValue)
}

// AddApproval is a paid mutator transaction binding the contract method 0xac3cb72c.
//
// Solidity: function addApproval(address _spender, uint256 _addedValue) returns(bool success)
func (_MystToken *MystTokenSession) AddApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.AddApproval(&_MystToken.TransactOpts, _spender, _addedValue)
}

// AddApproval is a paid mutator transaction binding the contract method 0xac3cb72c.
//
// Solidity: function addApproval(address _spender, uint256 _addedValue) returns(bool success)
func (_MystToken *MystTokenTransactorSession) AddApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.AddApproval(&_MystToken.TransactOpts, _spender, _addedValue)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_MystToken *MystTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MystToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_MystToken *MystTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.Approve(&_MystToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_MystToken *MystTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.Approve(&_MystToken.TransactOpts, _spender, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address receiver, uint256 amount) returns()
func (_MystToken *MystTokenTransactor) Mint(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MystToken.contract.Transact(opts, "mint", receiver, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address receiver, uint256 amount) returns()
func (_MystToken *MystTokenSession) Mint(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.Mint(&_MystToken.TransactOpts, receiver, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address receiver, uint256 amount) returns()
func (_MystToken *MystTokenTransactorSession) Mint(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.Mint(&_MystToken.TransactOpts, receiver, amount)
}

// SetMintAgent is a paid mutator transaction binding the contract method 0x43214675.
//
// Solidity: function setMintAgent(address addr, bool state) returns()
func (_MystToken *MystTokenTransactor) SetMintAgent(opts *bind.TransactOpts, addr common.Address, state bool) (*types.Transaction, error) {
	return _MystToken.contract.Transact(opts, "setMintAgent", addr, state)
}

// SetMintAgent is a paid mutator transaction binding the contract method 0x43214675.
//
// Solidity: function setMintAgent(address addr, bool state) returns()
func (_MystToken *MystTokenSession) SetMintAgent(addr common.Address, state bool) (*types.Transaction, error) {
	return _MystToken.Contract.SetMintAgent(&_MystToken.TransactOpts, addr, state)
}

// SetMintAgent is a paid mutator transaction binding the contract method 0x43214675.
//
// Solidity: function setMintAgent(address addr, bool state) returns()
func (_MystToken *MystTokenTransactorSession) SetMintAgent(addr common.Address, state bool) (*types.Transaction, error) {
	return _MystToken.Contract.SetMintAgent(&_MystToken.TransactOpts, addr, state)
}

// SetUpgradeAgent is a paid mutator transaction binding the contract method 0xd7e7088a.
//
// Solidity: function setUpgradeAgent(address agent) returns()
func (_MystToken *MystTokenTransactor) SetUpgradeAgent(opts *bind.TransactOpts, agent common.Address) (*types.Transaction, error) {
	return _MystToken.contract.Transact(opts, "setUpgradeAgent", agent)
}

// SetUpgradeAgent is a paid mutator transaction binding the contract method 0xd7e7088a.
//
// Solidity: function setUpgradeAgent(address agent) returns()
func (_MystToken *MystTokenSession) SetUpgradeAgent(agent common.Address) (*types.Transaction, error) {
	return _MystToken.Contract.SetUpgradeAgent(&_MystToken.TransactOpts, agent)
}

// SetUpgradeAgent is a paid mutator transaction binding the contract method 0xd7e7088a.
//
// Solidity: function setUpgradeAgent(address agent) returns()
func (_MystToken *MystTokenTransactorSession) SetUpgradeAgent(agent common.Address) (*types.Transaction, error) {
	return _MystToken.Contract.SetUpgradeAgent(&_MystToken.TransactOpts, agent)
}

// SetUpgradeMaster is a paid mutator transaction binding the contract method 0xffeb7d75.
//
// Solidity: function setUpgradeMaster(address master) returns()
func (_MystToken *MystTokenTransactor) SetUpgradeMaster(opts *bind.TransactOpts, master common.Address) (*types.Transaction, error) {
	return _MystToken.contract.Transact(opts, "setUpgradeMaster", master)
}

// SetUpgradeMaster is a paid mutator transaction binding the contract method 0xffeb7d75.
//
// Solidity: function setUpgradeMaster(address master) returns()
func (_MystToken *MystTokenSession) SetUpgradeMaster(master common.Address) (*types.Transaction, error) {
	return _MystToken.Contract.SetUpgradeMaster(&_MystToken.TransactOpts, master)
}

// SetUpgradeMaster is a paid mutator transaction binding the contract method 0xffeb7d75.
//
// Solidity: function setUpgradeMaster(address master) returns()
func (_MystToken *MystTokenTransactorSession) SetUpgradeMaster(master common.Address) (*types.Transaction, error) {
	return _MystToken.Contract.SetUpgradeMaster(&_MystToken.TransactOpts, master)
}

// SubApproval is a paid mutator transaction binding the contract method 0xe2301d02.
//
// Solidity: function subApproval(address _spender, uint256 _subtractedValue) returns(bool success)
func (_MystToken *MystTokenTransactor) SubApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _MystToken.contract.Transact(opts, "subApproval", _spender, _subtractedValue)
}

// SubApproval is a paid mutator transaction binding the contract method 0xe2301d02.
//
// Solidity: function subApproval(address _spender, uint256 _subtractedValue) returns(bool success)
func (_MystToken *MystTokenSession) SubApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.SubApproval(&_MystToken.TransactOpts, _spender, _subtractedValue)
}

// SubApproval is a paid mutator transaction binding the contract method 0xe2301d02.
//
// Solidity: function subApproval(address _spender, uint256 _subtractedValue) returns(bool success)
func (_MystToken *MystTokenTransactorSession) SubApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.SubApproval(&_MystToken.TransactOpts, _spender, _subtractedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_MystToken *MystTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MystToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_MystToken *MystTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.Transfer(&_MystToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_MystToken *MystTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.Transfer(&_MystToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_MystToken *MystTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MystToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_MystToken *MystTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.TransferFrom(&_MystToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_MystToken *MystTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.TransferFrom(&_MystToken.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MystToken *MystTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MystToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MystToken *MystTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MystToken.Contract.TransferOwnership(&_MystToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MystToken *MystTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MystToken.Contract.TransferOwnership(&_MystToken.TransactOpts, newOwner)
}

// Upgrade is a paid mutator transaction binding the contract method 0x45977d03.
//
// Solidity: function upgrade(uint256 value) returns()
func (_MystToken *MystTokenTransactor) Upgrade(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _MystToken.contract.Transact(opts, "upgrade", value)
}

// Upgrade is a paid mutator transaction binding the contract method 0x45977d03.
//
// Solidity: function upgrade(uint256 value) returns()
func (_MystToken *MystTokenSession) Upgrade(value *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.Upgrade(&_MystToken.TransactOpts, value)
}

// Upgrade is a paid mutator transaction binding the contract method 0x45977d03.
//
// Solidity: function upgrade(uint256 value) returns()
func (_MystToken *MystTokenTransactorSession) Upgrade(value *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.Upgrade(&_MystToken.TransactOpts, value)
}

// MystTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MystToken contract.
type MystTokenApprovalIterator struct {
	Event *MystTokenApproval // Event containing the contract specifics and raw log

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
func (it *MystTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MystTokenApproval)
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
		it.Event = new(MystTokenApproval)
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
func (it *MystTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MystTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MystTokenApproval represents a Approval event raised by the MystToken contract.
type MystTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MystToken *MystTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MystTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MystToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MystTokenApprovalIterator{contract: _MystToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MystToken *MystTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MystTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MystToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MystTokenApproval)
				if err := _MystToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_MystToken *MystTokenFilterer) ParseApproval(log types.Log) (*MystTokenApproval, error) {
	event := new(MystTokenApproval)
	if err := _MystToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MystTokenMintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the MystToken contract.
type MystTokenMintedIterator struct {
	Event *MystTokenMinted // Event containing the contract specifics and raw log

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
func (it *MystTokenMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MystTokenMinted)
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
		it.Event = new(MystTokenMinted)
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
func (it *MystTokenMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MystTokenMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MystTokenMinted represents a Minted event raised by the MystToken contract.
type MystTokenMinted struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe.
//
// Solidity: event Minted(address receiver, uint256 amount)
func (_MystToken *MystTokenFilterer) FilterMinted(opts *bind.FilterOpts) (*MystTokenMintedIterator, error) {

	logs, sub, err := _MystToken.contract.FilterLogs(opts, "Minted")
	if err != nil {
		return nil, err
	}
	return &MystTokenMintedIterator{contract: _MystToken.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe.
//
// Solidity: event Minted(address receiver, uint256 amount)
func (_MystToken *MystTokenFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *MystTokenMinted) (event.Subscription, error) {

	logs, sub, err := _MystToken.contract.WatchLogs(opts, "Minted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MystTokenMinted)
				if err := _MystToken.contract.UnpackLog(event, "Minted", log); err != nil {
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

// ParseMinted is a log parse operation binding the contract event 0x30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe.
//
// Solidity: event Minted(address receiver, uint256 amount)
func (_MystToken *MystTokenFilterer) ParseMinted(log types.Log) (*MystTokenMinted, error) {
	event := new(MystTokenMinted)
	if err := _MystToken.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MystTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MystToken contract.
type MystTokenTransferIterator struct {
	Event *MystTokenTransfer // Event containing the contract specifics and raw log

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
func (it *MystTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MystTokenTransfer)
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
		it.Event = new(MystTokenTransfer)
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
func (it *MystTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MystTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MystTokenTransfer represents a Transfer event raised by the MystToken contract.
type MystTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MystToken *MystTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MystTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MystToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MystTokenTransferIterator{contract: _MystToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MystToken *MystTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MystTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MystToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MystTokenTransfer)
				if err := _MystToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_MystToken *MystTokenFilterer) ParseTransfer(log types.Log) (*MystTokenTransfer, error) {
	event := new(MystTokenTransfer)
	if err := _MystToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MystTokenUpgradeIterator is returned from FilterUpgrade and is used to iterate over the raw logs and unpacked data for Upgrade events raised by the MystToken contract.
type MystTokenUpgradeIterator struct {
	Event *MystTokenUpgrade // Event containing the contract specifics and raw log

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
func (it *MystTokenUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MystTokenUpgrade)
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
		it.Event = new(MystTokenUpgrade)
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
func (it *MystTokenUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MystTokenUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MystTokenUpgrade represents a Upgrade event raised by the MystToken contract.
type MystTokenUpgrade struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpgrade is a free log retrieval operation binding the contract event 0x7e5c344a8141a805725cb476f76c6953b842222b967edd1f78ddb6e8b3f397ac.
//
// Solidity: event Upgrade(address indexed _from, address indexed _to, uint256 _value)
func (_MystToken *MystTokenFilterer) FilterUpgrade(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*MystTokenUpgradeIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _MystToken.contract.FilterLogs(opts, "Upgrade", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &MystTokenUpgradeIterator{contract: _MystToken.contract, event: "Upgrade", logs: logs, sub: sub}, nil
}

// WatchUpgrade is a free log subscription operation binding the contract event 0x7e5c344a8141a805725cb476f76c6953b842222b967edd1f78ddb6e8b3f397ac.
//
// Solidity: event Upgrade(address indexed _from, address indexed _to, uint256 _value)
func (_MystToken *MystTokenFilterer) WatchUpgrade(opts *bind.WatchOpts, sink chan<- *MystTokenUpgrade, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _MystToken.contract.WatchLogs(opts, "Upgrade", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MystTokenUpgrade)
				if err := _MystToken.contract.UnpackLog(event, "Upgrade", log); err != nil {
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

// ParseUpgrade is a log parse operation binding the contract event 0x7e5c344a8141a805725cb476f76c6953b842222b967edd1f78ddb6e8b3f397ac.
//
// Solidity: event Upgrade(address indexed _from, address indexed _to, uint256 _value)
func (_MystToken *MystTokenFilterer) ParseUpgrade(log types.Log) (*MystTokenUpgrade, error) {
	event := new(MystTokenUpgrade)
	if err := _MystToken.contract.UnpackLog(event, "Upgrade", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MystTokenUpgradeAgentSetIterator is returned from FilterUpgradeAgentSet and is used to iterate over the raw logs and unpacked data for UpgradeAgentSet events raised by the MystToken contract.
type MystTokenUpgradeAgentSetIterator struct {
	Event *MystTokenUpgradeAgentSet // Event containing the contract specifics and raw log

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
func (it *MystTokenUpgradeAgentSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MystTokenUpgradeAgentSet)
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
		it.Event = new(MystTokenUpgradeAgentSet)
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
func (it *MystTokenUpgradeAgentSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MystTokenUpgradeAgentSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MystTokenUpgradeAgentSet represents a UpgradeAgentSet event raised by the MystToken contract.
type MystTokenUpgradeAgentSet struct {
	Agent common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpgradeAgentSet is a free log retrieval operation binding the contract event 0x7845d5aa74cc410e35571258d954f23b82276e160fe8c188fa80566580f279cc.
//
// Solidity: event UpgradeAgentSet(address agent)
func (_MystToken *MystTokenFilterer) FilterUpgradeAgentSet(opts *bind.FilterOpts) (*MystTokenUpgradeAgentSetIterator, error) {

	logs, sub, err := _MystToken.contract.FilterLogs(opts, "UpgradeAgentSet")
	if err != nil {
		return nil, err
	}
	return &MystTokenUpgradeAgentSetIterator{contract: _MystToken.contract, event: "UpgradeAgentSet", logs: logs, sub: sub}, nil
}

// WatchUpgradeAgentSet is a free log subscription operation binding the contract event 0x7845d5aa74cc410e35571258d954f23b82276e160fe8c188fa80566580f279cc.
//
// Solidity: event UpgradeAgentSet(address agent)
func (_MystToken *MystTokenFilterer) WatchUpgradeAgentSet(opts *bind.WatchOpts, sink chan<- *MystTokenUpgradeAgentSet) (event.Subscription, error) {

	logs, sub, err := _MystToken.contract.WatchLogs(opts, "UpgradeAgentSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MystTokenUpgradeAgentSet)
				if err := _MystToken.contract.UnpackLog(event, "UpgradeAgentSet", log); err != nil {
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

// ParseUpgradeAgentSet is a log parse operation binding the contract event 0x7845d5aa74cc410e35571258d954f23b82276e160fe8c188fa80566580f279cc.
//
// Solidity: event UpgradeAgentSet(address agent)
func (_MystToken *MystTokenFilterer) ParseUpgradeAgentSet(log types.Log) (*MystTokenUpgradeAgentSet, error) {
	event := new(MystTokenUpgradeAgentSet)
	if err := _MystToken.contract.UnpackLog(event, "UpgradeAgentSet", log); err != nil {
		return nil, err
	}
	return event, nil
}
