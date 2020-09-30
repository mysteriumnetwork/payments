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

// OldMystTokenABI is the input ABI used to generate the binding from.
const OldMystTokenABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Upgrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"UpgradeAgentSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mintAgents\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintingFinished\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalUpgraded\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradeAgent\",\"outputs\":[{\"internalType\":\"contractIUpgradeAgent\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradeMaster\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"addApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"subApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"setMintAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"setUpgradeAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUpgradeState\",\"outputs\":[{\"internalType\":\"enumOldMystToken.UpgradeState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"master\",\"type\":\"address\"}],\"name\":\"setUpgradeMaster\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OldMystTokenBin is the compiled bytecode used for deploying new contracts.
var OldMystTokenBin = "0x60806040526001805460ff60a01b1916905534801561001d57600080fd5b5060018054600280546001600160a01b0319908116339081179092559182168117909116811782556000908152600560205260409020805460ff1916909117905561112c8061006d6000396000f3fe608060405234801561001057600080fd5b50600436106101585760003560e01c806370a08231116100c3578063c752ff621161007c578063c752ff6214610412578063d7e7088a1461041a578063dd62ed3e14610440578063e2301d021461046e578063f2fde38b1461049a578063ffeb7d75146104c057610158565b806370a082311461035b5780638444b391146103815780638da5cb5b146103aa57806395d89b41146103b2578063a9059cbb146103ba578063ac3cb72c146103e657610158565b806340c10f191161011557806340c10f191461029057806342c1867b146102be57806343214675146102e457806345977d03146103125780635de4ccb01461032f578063600440cb1461035357610158565b806305d2035b1461015d57806306fdde0314610179578063095ea7b3146101f657806318160ddd1461022257806323b872dd1461023c578063313ce56714610272575b600080fd5b6101656104e6565b604080519115158252519081900360200190f35b6101816104f6565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101bb5781810151838201526020016101a3565b50505050905090810190601f1680156101e85780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101656004803603604081101561020c57600080fd5b506001600160a01b038135169060200135610526565b61022a6105ca565b60408051918252519081900360200190f35b6101656004803603606081101561025257600080fd5b506001600160a01b038135811691602081013590911690604001356105d0565b61027a6106cd565b6040805160ff9092168252519081900360200190f35b6102bc600480360360408110156102a657600080fd5b506001600160a01b0381351690602001356106d2565b005b610165600480360360208110156102d457600080fd5b50356001600160a01b0316610892565b6102bc600480360360408110156102fa57600080fd5b506001600160a01b03813516906020013515156108a7565b6102bc6004803603602081101561032857600080fd5b5035610900565b610337610a45565b604080516001600160a01b039092168252519081900360200190f35b610337610a54565b61022a6004803603602081101561037157600080fd5b50356001600160a01b0316610a63565b610389610a7e565b6040518082600481111561039957fe5b815260200191505060405180910390f35b610337610aaf565b610181610abe565b610165600480360360408110156103d057600080fd5b506001600160a01b038135169060200135610adf565b610165600480360360408110156103fc57600080fd5b506001600160a01b038135169060200135610b95565b61022a610c3b565b6102bc6004803603602081101561043057600080fd5b50356001600160a01b0316610c41565b61022a6004803603604081101561045657600080fd5b506001600160a01b0381358116916020013516610e49565b6101656004803603604081101561048457600080fd5b506001600160a01b038135169060200135610e74565b6102bc600480360360208110156104b057600080fd5b50356001600160a01b0316610f4b565b6102bc600480360360208110156104d657600080fd5b50356001600160a01b0316610f90565b600154600160a01b900460ff1681565b604051806040016040528060148152602001732a32b9ba1026bcb9ba32b934bab6903a37b5b2b760611b81525081565b6000811580159061055957503360009081526007602090815260408083206001600160a01b038716845290915290205415155b1561056357600080fd5b3360008181526007602090815260408083206001600160a01b03881680855290835292819020869055805186815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a350600192915050565b60005481565b6001600160a01b0380841660009081526007602090815260408083203384528252808320549386168352600690915281205490919061060f9084610fdc565b6001600160a01b03808616600090815260066020526040808220939093559087168152205461063e9084611049565b6001600160a01b0386166000908152600660205260409020556106618184611049565b6001600160a01b03808716600081815260076020908152604080832033845282529182902094909455805187815290519288169391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a3506001949350505050565b600881565b3360009081526005602052604090205460ff166106ee57600080fd5b600154600160a01b900460ff161561070557600080fd5b8061070f57600080fd5b60005473__SafeMathLib___________________________6366098d4f9091836040518363ffffffff1660e01b8152600401808381526020018281526020019250505060206040518083038186803b15801561076a57600080fd5b505af415801561077e573d6000803e3d6000fd5b505050506040513d602081101561079457600080fd5b505160009081556001600160a01b0383168152600660209081526040918290205482516366098d4f60e01b8152600481019190915260248101849052915173__SafeMathLib___________________________926366098d4f926044808301939192829003018186803b15801561080a57600080fd5b505af415801561081e573d6000803e3d6000fd5b505050506040513d602081101561083457600080fd5b50516001600160a01b03831660008181526006602090815260409182902093909355805191825291810183905281517f30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe929181900390910190a15050565b60056020526000908152604090205460ff1681565b6001546001600160a01b031633146108be57600080fd5b600154600160a01b900460ff16156108d557600080fd5b6001600160a01b03919091166000908152600560205260409020805460ff1916911515919091179055565b600061090a610a7e565b9050600381600481111561091a57fe5b14806109315750600481600481111561092f57fe5b145b61093a57600080fd5b8161094457600080fd5b3360009081526006602052604090205461095e9083611049565b336000908152600660205260408120919091555461097c9083611049565b60005560045461098c9083610fdc565b60049081556003546040805163753e88e560e01b8152339381019390935260248301859052516001600160a01b039091169163753e88e591604480830192600092919082900301818387803b1580156109e457600080fd5b505af11580156109f8573d6000803e3d6000fd5b50506003546040805186815290516001600160a01b0390921693503392507f7e5c344a8141a805725cb476f76c6953b842222b967edd1f78ddb6e8b3f397ac919081900360200190a35050565b6003546001600160a01b031681565b6002546001600160a01b031681565b6001600160a01b031660009081526006602052604090205490565b6003546000906001600160a01b0316610a9957506002610aac565b600454610aa857506003610aac565b5060045b90565b6001546001600160a01b031681565b60405180604001604052806005815260200164135654d51560da1b81525081565b6000604036604414610af057600080fd5b33600090815260066020526040902054610b0a9084611049565b33600090815260066020526040808220929092556001600160a01b03861681522054610b369084610fdc565b6001600160a01b0385166000818152600660209081526040918290209390935580518681529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35060019392505050565b6000604036604414610ba657600080fd5b3360009081526007602090815260408083206001600160a01b0388168452909152902054610bd48185610fdc565b3360008181526007602090815260408083206001600160a01b038b168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a3506001949350505050565b60045481565b6002546001600160a01b03163314610c8a5760405162461bcd60e51b815260040180806020018281038252602a8152602001806110a7602a913960400191505060405180910390fd5b6001600160a01b038116610c9d57600080fd5b6004610ca7610a7e565b6004811115610cb257fe5b1415610cef5760405162461bcd60e51b81526004018080602001828103825260268152602001806110d16026913960400191505060405180910390fd5b600380546001600160a01b0319166001600160a01b038381169190911791829055604080516330e9ebd360e11b8152905192909116916361d3d7a691600480820192602092909190829003018186803b158015610d4b57600080fd5b505afa158015610d5f573d6000803e3d6000fd5b505050506040513d6020811015610d7557600080fd5b5051610d8057600080fd5b600054600360009054906101000a90046001600160a01b03166001600160a01b0316634b2ba0dd6040518163ffffffff1660e01b815260040160206040518083038186803b158015610dd157600080fd5b505afa158015610de5573d6000803e3d6000fd5b505050506040513d6020811015610dfb57600080fd5b505114610e0757600080fd5b600354604080516001600160a01b039092168252517f7845d5aa74cc410e35571258d954f23b82276e160fe8c188fa80566580f279cc9181900360200190a150565b6001600160a01b03918216600090815260076020908152604080832093909416825291909152205490565b6000604036604414610e8557600080fd5b3360009081526007602090815260408083206001600160a01b038816845290915290205480841115610eda573360009081526007602090815260408083206001600160a01b0389168452909152812055610ee4565b610bd48185611049565b3360008181526007602090815260408083206001600160a01b038a168085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a3506001949350505050565b6001546001600160a01b03163314610f6257600080fd5b6001600160a01b03811615610f8d57600180546001600160a01b0319166001600160a01b0383161790555b50565b6001600160a01b038116610fa357600080fd5b6002546001600160a01b03163314610fba57600080fd5b600280546001600160a01b0319166001600160a01b0392909216919091179055565b6000828201838110801590610ff15750828110155b611042576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b6000828211156110a0576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b5090039056fe4f6e6c792061206d61737465722063616e2064657369676e61746520746865206e657874206167656e74557067726164652068617320616c726561647920626567756e20666f7220616e206167656e74a26469706673582212205653e00f2013287a1490e48babc280a56f4c43452b6faf6d4beea13f3db01e0064736f6c634300060c0033"

// DeployOldMystToken deploys a new Ethereum contract, binding an instance of OldMystToken to it.
func DeployOldMystToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OldMystToken, error) {
	parsed, err := abi.JSON(strings.NewReader(OldMystTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OldMystTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OldMystToken{OldMystTokenCaller: OldMystTokenCaller{contract: contract}, OldMystTokenTransactor: OldMystTokenTransactor{contract: contract}, OldMystTokenFilterer: OldMystTokenFilterer{contract: contract}}, nil
}

// OldMystToken is an auto generated Go binding around an Ethereum contract.
type OldMystToken struct {
	OldMystTokenCaller     // Read-only binding to the contract
	OldMystTokenTransactor // Write-only binding to the contract
	OldMystTokenFilterer   // Log filterer for contract events
}

// OldMystTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type OldMystTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OldMystTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OldMystTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OldMystTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OldMystTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OldMystTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OldMystTokenSession struct {
	Contract     *OldMystToken     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OldMystTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OldMystTokenCallerSession struct {
	Contract *OldMystTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// OldMystTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OldMystTokenTransactorSession struct {
	Contract     *OldMystTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// OldMystTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type OldMystTokenRaw struct {
	Contract *OldMystToken // Generic contract binding to access the raw methods on
}

// OldMystTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OldMystTokenCallerRaw struct {
	Contract *OldMystTokenCaller // Generic read-only contract binding to access the raw methods on
}

// OldMystTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OldMystTokenTransactorRaw struct {
	Contract *OldMystTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOldMystToken creates a new instance of OldMystToken, bound to a specific deployed contract.
func NewOldMystToken(address common.Address, backend bind.ContractBackend) (*OldMystToken, error) {
	contract, err := bindOldMystToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OldMystToken{OldMystTokenCaller: OldMystTokenCaller{contract: contract}, OldMystTokenTransactor: OldMystTokenTransactor{contract: contract}, OldMystTokenFilterer: OldMystTokenFilterer{contract: contract}}, nil
}

// NewOldMystTokenCaller creates a new read-only instance of OldMystToken, bound to a specific deployed contract.
func NewOldMystTokenCaller(address common.Address, caller bind.ContractCaller) (*OldMystTokenCaller, error) {
	contract, err := bindOldMystToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OldMystTokenCaller{contract: contract}, nil
}

// NewOldMystTokenTransactor creates a new write-only instance of OldMystToken, bound to a specific deployed contract.
func NewOldMystTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*OldMystTokenTransactor, error) {
	contract, err := bindOldMystToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OldMystTokenTransactor{contract: contract}, nil
}

// NewOldMystTokenFilterer creates a new log filterer instance of OldMystToken, bound to a specific deployed contract.
func NewOldMystTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*OldMystTokenFilterer, error) {
	contract, err := bindOldMystToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OldMystTokenFilterer{contract: contract}, nil
}

// bindOldMystToken binds a generic wrapper to an already deployed contract.
func bindOldMystToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OldMystTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OldMystToken *OldMystTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OldMystToken.Contract.OldMystTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OldMystToken *OldMystTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OldMystToken.Contract.OldMystTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OldMystToken *OldMystTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OldMystToken.Contract.OldMystTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OldMystToken *OldMystTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OldMystToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OldMystToken *OldMystTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OldMystToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OldMystToken *OldMystTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OldMystToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_OldMystToken *OldMystTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OldMystToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_OldMystToken *OldMystTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _OldMystToken.Contract.Allowance(&_OldMystToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_OldMystToken *OldMystTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _OldMystToken.Contract.Allowance(&_OldMystToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_OldMystToken *OldMystTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OldMystToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_OldMystToken *OldMystTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _OldMystToken.Contract.BalanceOf(&_OldMystToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_OldMystToken *OldMystTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _OldMystToken.Contract.BalanceOf(&_OldMystToken.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_OldMystToken *OldMystTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _OldMystToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_OldMystToken *OldMystTokenSession) Decimals() (uint8, error) {
	return _OldMystToken.Contract.Decimals(&_OldMystToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_OldMystToken *OldMystTokenCallerSession) Decimals() (uint8, error) {
	return _OldMystToken.Contract.Decimals(&_OldMystToken.CallOpts)
}

// GetUpgradeState is a free data retrieval call binding the contract method 0x8444b391.
//
// Solidity: function getUpgradeState() view returns(uint8)
func (_OldMystToken *OldMystTokenCaller) GetUpgradeState(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _OldMystToken.contract.Call(opts, out, "getUpgradeState")
	return *ret0, err
}

// GetUpgradeState is a free data retrieval call binding the contract method 0x8444b391.
//
// Solidity: function getUpgradeState() view returns(uint8)
func (_OldMystToken *OldMystTokenSession) GetUpgradeState() (uint8, error) {
	return _OldMystToken.Contract.GetUpgradeState(&_OldMystToken.CallOpts)
}

// GetUpgradeState is a free data retrieval call binding the contract method 0x8444b391.
//
// Solidity: function getUpgradeState() view returns(uint8)
func (_OldMystToken *OldMystTokenCallerSession) GetUpgradeState() (uint8, error) {
	return _OldMystToken.Contract.GetUpgradeState(&_OldMystToken.CallOpts)
}

// MintAgents is a free data retrieval call binding the contract method 0x42c1867b.
//
// Solidity: function mintAgents(address ) view returns(bool)
func (_OldMystToken *OldMystTokenCaller) MintAgents(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _OldMystToken.contract.Call(opts, out, "mintAgents", arg0)
	return *ret0, err
}

// MintAgents is a free data retrieval call binding the contract method 0x42c1867b.
//
// Solidity: function mintAgents(address ) view returns(bool)
func (_OldMystToken *OldMystTokenSession) MintAgents(arg0 common.Address) (bool, error) {
	return _OldMystToken.Contract.MintAgents(&_OldMystToken.CallOpts, arg0)
}

// MintAgents is a free data retrieval call binding the contract method 0x42c1867b.
//
// Solidity: function mintAgents(address ) view returns(bool)
func (_OldMystToken *OldMystTokenCallerSession) MintAgents(arg0 common.Address) (bool, error) {
	return _OldMystToken.Contract.MintAgents(&_OldMystToken.CallOpts, arg0)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() view returns(bool)
func (_OldMystToken *OldMystTokenCaller) MintingFinished(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _OldMystToken.contract.Call(opts, out, "mintingFinished")
	return *ret0, err
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() view returns(bool)
func (_OldMystToken *OldMystTokenSession) MintingFinished() (bool, error) {
	return _OldMystToken.Contract.MintingFinished(&_OldMystToken.CallOpts)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() view returns(bool)
func (_OldMystToken *OldMystTokenCallerSession) MintingFinished() (bool, error) {
	return _OldMystToken.Contract.MintingFinished(&_OldMystToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_OldMystToken *OldMystTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _OldMystToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_OldMystToken *OldMystTokenSession) Name() (string, error) {
	return _OldMystToken.Contract.Name(&_OldMystToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_OldMystToken *OldMystTokenCallerSession) Name() (string, error) {
	return _OldMystToken.Contract.Name(&_OldMystToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OldMystToken *OldMystTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _OldMystToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OldMystToken *OldMystTokenSession) Owner() (common.Address, error) {
	return _OldMystToken.Contract.Owner(&_OldMystToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OldMystToken *OldMystTokenCallerSession) Owner() (common.Address, error) {
	return _OldMystToken.Contract.Owner(&_OldMystToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_OldMystToken *OldMystTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _OldMystToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_OldMystToken *OldMystTokenSession) Symbol() (string, error) {
	return _OldMystToken.Contract.Symbol(&_OldMystToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_OldMystToken *OldMystTokenCallerSession) Symbol() (string, error) {
	return _OldMystToken.Contract.Symbol(&_OldMystToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_OldMystToken *OldMystTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OldMystToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_OldMystToken *OldMystTokenSession) TotalSupply() (*big.Int, error) {
	return _OldMystToken.Contract.TotalSupply(&_OldMystToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_OldMystToken *OldMystTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _OldMystToken.Contract.TotalSupply(&_OldMystToken.CallOpts)
}

// TotalUpgraded is a free data retrieval call binding the contract method 0xc752ff62.
//
// Solidity: function totalUpgraded() view returns(uint256)
func (_OldMystToken *OldMystTokenCaller) TotalUpgraded(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OldMystToken.contract.Call(opts, out, "totalUpgraded")
	return *ret0, err
}

// TotalUpgraded is a free data retrieval call binding the contract method 0xc752ff62.
//
// Solidity: function totalUpgraded() view returns(uint256)
func (_OldMystToken *OldMystTokenSession) TotalUpgraded() (*big.Int, error) {
	return _OldMystToken.Contract.TotalUpgraded(&_OldMystToken.CallOpts)
}

// TotalUpgraded is a free data retrieval call binding the contract method 0xc752ff62.
//
// Solidity: function totalUpgraded() view returns(uint256)
func (_OldMystToken *OldMystTokenCallerSession) TotalUpgraded() (*big.Int, error) {
	return _OldMystToken.Contract.TotalUpgraded(&_OldMystToken.CallOpts)
}

// UpgradeAgent is a free data retrieval call binding the contract method 0x5de4ccb0.
//
// Solidity: function upgradeAgent() view returns(address)
func (_OldMystToken *OldMystTokenCaller) UpgradeAgent(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _OldMystToken.contract.Call(opts, out, "upgradeAgent")
	return *ret0, err
}

// UpgradeAgent is a free data retrieval call binding the contract method 0x5de4ccb0.
//
// Solidity: function upgradeAgent() view returns(address)
func (_OldMystToken *OldMystTokenSession) UpgradeAgent() (common.Address, error) {
	return _OldMystToken.Contract.UpgradeAgent(&_OldMystToken.CallOpts)
}

// UpgradeAgent is a free data retrieval call binding the contract method 0x5de4ccb0.
//
// Solidity: function upgradeAgent() view returns(address)
func (_OldMystToken *OldMystTokenCallerSession) UpgradeAgent() (common.Address, error) {
	return _OldMystToken.Contract.UpgradeAgent(&_OldMystToken.CallOpts)
}

// UpgradeMaster is a free data retrieval call binding the contract method 0x600440cb.
//
// Solidity: function upgradeMaster() view returns(address)
func (_OldMystToken *OldMystTokenCaller) UpgradeMaster(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _OldMystToken.contract.Call(opts, out, "upgradeMaster")
	return *ret0, err
}

// UpgradeMaster is a free data retrieval call binding the contract method 0x600440cb.
//
// Solidity: function upgradeMaster() view returns(address)
func (_OldMystToken *OldMystTokenSession) UpgradeMaster() (common.Address, error) {
	return _OldMystToken.Contract.UpgradeMaster(&_OldMystToken.CallOpts)
}

// UpgradeMaster is a free data retrieval call binding the contract method 0x600440cb.
//
// Solidity: function upgradeMaster() view returns(address)
func (_OldMystToken *OldMystTokenCallerSession) UpgradeMaster() (common.Address, error) {
	return _OldMystToken.Contract.UpgradeMaster(&_OldMystToken.CallOpts)
}

// AddApproval is a paid mutator transaction binding the contract method 0xac3cb72c.
//
// Solidity: function addApproval(address _spender, uint256 _addedValue) returns(bool success)
func (_OldMystToken *OldMystTokenTransactor) AddApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _OldMystToken.contract.Transact(opts, "addApproval", _spender, _addedValue)
}

// AddApproval is a paid mutator transaction binding the contract method 0xac3cb72c.
//
// Solidity: function addApproval(address _spender, uint256 _addedValue) returns(bool success)
func (_OldMystToken *OldMystTokenSession) AddApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _OldMystToken.Contract.AddApproval(&_OldMystToken.TransactOpts, _spender, _addedValue)
}

// AddApproval is a paid mutator transaction binding the contract method 0xac3cb72c.
//
// Solidity: function addApproval(address _spender, uint256 _addedValue) returns(bool success)
func (_OldMystToken *OldMystTokenTransactorSession) AddApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _OldMystToken.Contract.AddApproval(&_OldMystToken.TransactOpts, _spender, _addedValue)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_OldMystToken *OldMystTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _OldMystToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_OldMystToken *OldMystTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _OldMystToken.Contract.Approve(&_OldMystToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_OldMystToken *OldMystTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _OldMystToken.Contract.Approve(&_OldMystToken.TransactOpts, _spender, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address receiver, uint256 amount) returns()
func (_OldMystToken *OldMystTokenTransactor) Mint(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OldMystToken.contract.Transact(opts, "mint", receiver, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address receiver, uint256 amount) returns()
func (_OldMystToken *OldMystTokenSession) Mint(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OldMystToken.Contract.Mint(&_OldMystToken.TransactOpts, receiver, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address receiver, uint256 amount) returns()
func (_OldMystToken *OldMystTokenTransactorSession) Mint(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OldMystToken.Contract.Mint(&_OldMystToken.TransactOpts, receiver, amount)
}

// SetMintAgent is a paid mutator transaction binding the contract method 0x43214675.
//
// Solidity: function setMintAgent(address addr, bool state) returns()
func (_OldMystToken *OldMystTokenTransactor) SetMintAgent(opts *bind.TransactOpts, addr common.Address, state bool) (*types.Transaction, error) {
	return _OldMystToken.contract.Transact(opts, "setMintAgent", addr, state)
}

// SetMintAgent is a paid mutator transaction binding the contract method 0x43214675.
//
// Solidity: function setMintAgent(address addr, bool state) returns()
func (_OldMystToken *OldMystTokenSession) SetMintAgent(addr common.Address, state bool) (*types.Transaction, error) {
	return _OldMystToken.Contract.SetMintAgent(&_OldMystToken.TransactOpts, addr, state)
}

// SetMintAgent is a paid mutator transaction binding the contract method 0x43214675.
//
// Solidity: function setMintAgent(address addr, bool state) returns()
func (_OldMystToken *OldMystTokenTransactorSession) SetMintAgent(addr common.Address, state bool) (*types.Transaction, error) {
	return _OldMystToken.Contract.SetMintAgent(&_OldMystToken.TransactOpts, addr, state)
}

// SetUpgradeAgent is a paid mutator transaction binding the contract method 0xd7e7088a.
//
// Solidity: function setUpgradeAgent(address agent) returns()
func (_OldMystToken *OldMystTokenTransactor) SetUpgradeAgent(opts *bind.TransactOpts, agent common.Address) (*types.Transaction, error) {
	return _OldMystToken.contract.Transact(opts, "setUpgradeAgent", agent)
}

// SetUpgradeAgent is a paid mutator transaction binding the contract method 0xd7e7088a.
//
// Solidity: function setUpgradeAgent(address agent) returns()
func (_OldMystToken *OldMystTokenSession) SetUpgradeAgent(agent common.Address) (*types.Transaction, error) {
	return _OldMystToken.Contract.SetUpgradeAgent(&_OldMystToken.TransactOpts, agent)
}

// SetUpgradeAgent is a paid mutator transaction binding the contract method 0xd7e7088a.
//
// Solidity: function setUpgradeAgent(address agent) returns()
func (_OldMystToken *OldMystTokenTransactorSession) SetUpgradeAgent(agent common.Address) (*types.Transaction, error) {
	return _OldMystToken.Contract.SetUpgradeAgent(&_OldMystToken.TransactOpts, agent)
}

// SetUpgradeMaster is a paid mutator transaction binding the contract method 0xffeb7d75.
//
// Solidity: function setUpgradeMaster(address master) returns()
func (_OldMystToken *OldMystTokenTransactor) SetUpgradeMaster(opts *bind.TransactOpts, master common.Address) (*types.Transaction, error) {
	return _OldMystToken.contract.Transact(opts, "setUpgradeMaster", master)
}

// SetUpgradeMaster is a paid mutator transaction binding the contract method 0xffeb7d75.
//
// Solidity: function setUpgradeMaster(address master) returns()
func (_OldMystToken *OldMystTokenSession) SetUpgradeMaster(master common.Address) (*types.Transaction, error) {
	return _OldMystToken.Contract.SetUpgradeMaster(&_OldMystToken.TransactOpts, master)
}

// SetUpgradeMaster is a paid mutator transaction binding the contract method 0xffeb7d75.
//
// Solidity: function setUpgradeMaster(address master) returns()
func (_OldMystToken *OldMystTokenTransactorSession) SetUpgradeMaster(master common.Address) (*types.Transaction, error) {
	return _OldMystToken.Contract.SetUpgradeMaster(&_OldMystToken.TransactOpts, master)
}

// SubApproval is a paid mutator transaction binding the contract method 0xe2301d02.
//
// Solidity: function subApproval(address _spender, uint256 _subtractedValue) returns(bool success)
func (_OldMystToken *OldMystTokenTransactor) SubApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _OldMystToken.contract.Transact(opts, "subApproval", _spender, _subtractedValue)
}

// SubApproval is a paid mutator transaction binding the contract method 0xe2301d02.
//
// Solidity: function subApproval(address _spender, uint256 _subtractedValue) returns(bool success)
func (_OldMystToken *OldMystTokenSession) SubApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _OldMystToken.Contract.SubApproval(&_OldMystToken.TransactOpts, _spender, _subtractedValue)
}

// SubApproval is a paid mutator transaction binding the contract method 0xe2301d02.
//
// Solidity: function subApproval(address _spender, uint256 _subtractedValue) returns(bool success)
func (_OldMystToken *OldMystTokenTransactorSession) SubApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _OldMystToken.Contract.SubApproval(&_OldMystToken.TransactOpts, _spender, _subtractedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_OldMystToken *OldMystTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _OldMystToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_OldMystToken *OldMystTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _OldMystToken.Contract.Transfer(&_OldMystToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_OldMystToken *OldMystTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _OldMystToken.Contract.Transfer(&_OldMystToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_OldMystToken *OldMystTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _OldMystToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_OldMystToken *OldMystTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _OldMystToken.Contract.TransferFrom(&_OldMystToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_OldMystToken *OldMystTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _OldMystToken.Contract.TransferFrom(&_OldMystToken.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OldMystToken *OldMystTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OldMystToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OldMystToken *OldMystTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OldMystToken.Contract.TransferOwnership(&_OldMystToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OldMystToken *OldMystTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OldMystToken.Contract.TransferOwnership(&_OldMystToken.TransactOpts, newOwner)
}

// Upgrade is a paid mutator transaction binding the contract method 0x45977d03.
//
// Solidity: function upgrade(uint256 value) returns()
func (_OldMystToken *OldMystTokenTransactor) Upgrade(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _OldMystToken.contract.Transact(opts, "upgrade", value)
}

// Upgrade is a paid mutator transaction binding the contract method 0x45977d03.
//
// Solidity: function upgrade(uint256 value) returns()
func (_OldMystToken *OldMystTokenSession) Upgrade(value *big.Int) (*types.Transaction, error) {
	return _OldMystToken.Contract.Upgrade(&_OldMystToken.TransactOpts, value)
}

// Upgrade is a paid mutator transaction binding the contract method 0x45977d03.
//
// Solidity: function upgrade(uint256 value) returns()
func (_OldMystToken *OldMystTokenTransactorSession) Upgrade(value *big.Int) (*types.Transaction, error) {
	return _OldMystToken.Contract.Upgrade(&_OldMystToken.TransactOpts, value)
}

// OldMystTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the OldMystToken contract.
type OldMystTokenApprovalIterator struct {
	Event *OldMystTokenApproval // Event containing the contract specifics and raw log

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
func (it *OldMystTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OldMystTokenApproval)
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
		it.Event = new(OldMystTokenApproval)
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
func (it *OldMystTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OldMystTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OldMystTokenApproval represents a Approval event raised by the OldMystToken contract.
type OldMystTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_OldMystToken *OldMystTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*OldMystTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _OldMystToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &OldMystTokenApprovalIterator{contract: _OldMystToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_OldMystToken *OldMystTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *OldMystTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _OldMystToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OldMystTokenApproval)
				if err := _OldMystToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_OldMystToken *OldMystTokenFilterer) ParseApproval(log types.Log) (*OldMystTokenApproval, error) {
	event := new(OldMystTokenApproval)
	if err := _OldMystToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OldMystTokenMintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the OldMystToken contract.
type OldMystTokenMintedIterator struct {
	Event *OldMystTokenMinted // Event containing the contract specifics and raw log

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
func (it *OldMystTokenMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OldMystTokenMinted)
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
		it.Event = new(OldMystTokenMinted)
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
func (it *OldMystTokenMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OldMystTokenMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OldMystTokenMinted represents a Minted event raised by the OldMystToken contract.
type OldMystTokenMinted struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe.
//
// Solidity: event Minted(address receiver, uint256 amount)
func (_OldMystToken *OldMystTokenFilterer) FilterMinted(opts *bind.FilterOpts) (*OldMystTokenMintedIterator, error) {

	logs, sub, err := _OldMystToken.contract.FilterLogs(opts, "Minted")
	if err != nil {
		return nil, err
	}
	return &OldMystTokenMintedIterator{contract: _OldMystToken.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe.
//
// Solidity: event Minted(address receiver, uint256 amount)
func (_OldMystToken *OldMystTokenFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *OldMystTokenMinted) (event.Subscription, error) {

	logs, sub, err := _OldMystToken.contract.WatchLogs(opts, "Minted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OldMystTokenMinted)
				if err := _OldMystToken.contract.UnpackLog(event, "Minted", log); err != nil {
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
func (_OldMystToken *OldMystTokenFilterer) ParseMinted(log types.Log) (*OldMystTokenMinted, error) {
	event := new(OldMystTokenMinted)
	if err := _OldMystToken.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OldMystTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the OldMystToken contract.
type OldMystTokenTransferIterator struct {
	Event *OldMystTokenTransfer // Event containing the contract specifics and raw log

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
func (it *OldMystTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OldMystTokenTransfer)
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
		it.Event = new(OldMystTokenTransfer)
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
func (it *OldMystTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OldMystTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OldMystTokenTransfer represents a Transfer event raised by the OldMystToken contract.
type OldMystTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_OldMystToken *OldMystTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OldMystTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OldMystToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OldMystTokenTransferIterator{contract: _OldMystToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_OldMystToken *OldMystTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *OldMystTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OldMystToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OldMystTokenTransfer)
				if err := _OldMystToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_OldMystToken *OldMystTokenFilterer) ParseTransfer(log types.Log) (*OldMystTokenTransfer, error) {
	event := new(OldMystTokenTransfer)
	if err := _OldMystToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OldMystTokenUpgradeIterator is returned from FilterUpgrade and is used to iterate over the raw logs and unpacked data for Upgrade events raised by the OldMystToken contract.
type OldMystTokenUpgradeIterator struct {
	Event *OldMystTokenUpgrade // Event containing the contract specifics and raw log

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
func (it *OldMystTokenUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OldMystTokenUpgrade)
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
		it.Event = new(OldMystTokenUpgrade)
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
func (it *OldMystTokenUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OldMystTokenUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OldMystTokenUpgrade represents a Upgrade event raised by the OldMystToken contract.
type OldMystTokenUpgrade struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpgrade is a free log retrieval operation binding the contract event 0x7e5c344a8141a805725cb476f76c6953b842222b967edd1f78ddb6e8b3f397ac.
//
// Solidity: event Upgrade(address indexed _from, address indexed _to, uint256 _value)
func (_OldMystToken *OldMystTokenFilterer) FilterUpgrade(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*OldMystTokenUpgradeIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _OldMystToken.contract.FilterLogs(opts, "Upgrade", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &OldMystTokenUpgradeIterator{contract: _OldMystToken.contract, event: "Upgrade", logs: logs, sub: sub}, nil
}

// WatchUpgrade is a free log subscription operation binding the contract event 0x7e5c344a8141a805725cb476f76c6953b842222b967edd1f78ddb6e8b3f397ac.
//
// Solidity: event Upgrade(address indexed _from, address indexed _to, uint256 _value)
func (_OldMystToken *OldMystTokenFilterer) WatchUpgrade(opts *bind.WatchOpts, sink chan<- *OldMystTokenUpgrade, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _OldMystToken.contract.WatchLogs(opts, "Upgrade", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OldMystTokenUpgrade)
				if err := _OldMystToken.contract.UnpackLog(event, "Upgrade", log); err != nil {
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
func (_OldMystToken *OldMystTokenFilterer) ParseUpgrade(log types.Log) (*OldMystTokenUpgrade, error) {
	event := new(OldMystTokenUpgrade)
	if err := _OldMystToken.contract.UnpackLog(event, "Upgrade", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OldMystTokenUpgradeAgentSetIterator is returned from FilterUpgradeAgentSet and is used to iterate over the raw logs and unpacked data for UpgradeAgentSet events raised by the OldMystToken contract.
type OldMystTokenUpgradeAgentSetIterator struct {
	Event *OldMystTokenUpgradeAgentSet // Event containing the contract specifics and raw log

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
func (it *OldMystTokenUpgradeAgentSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OldMystTokenUpgradeAgentSet)
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
		it.Event = new(OldMystTokenUpgradeAgentSet)
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
func (it *OldMystTokenUpgradeAgentSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OldMystTokenUpgradeAgentSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OldMystTokenUpgradeAgentSet represents a UpgradeAgentSet event raised by the OldMystToken contract.
type OldMystTokenUpgradeAgentSet struct {
	Agent common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpgradeAgentSet is a free log retrieval operation binding the contract event 0x7845d5aa74cc410e35571258d954f23b82276e160fe8c188fa80566580f279cc.
//
// Solidity: event UpgradeAgentSet(address agent)
func (_OldMystToken *OldMystTokenFilterer) FilterUpgradeAgentSet(opts *bind.FilterOpts) (*OldMystTokenUpgradeAgentSetIterator, error) {

	logs, sub, err := _OldMystToken.contract.FilterLogs(opts, "UpgradeAgentSet")
	if err != nil {
		return nil, err
	}
	return &OldMystTokenUpgradeAgentSetIterator{contract: _OldMystToken.contract, event: "UpgradeAgentSet", logs: logs, sub: sub}, nil
}

// WatchUpgradeAgentSet is a free log subscription operation binding the contract event 0x7845d5aa74cc410e35571258d954f23b82276e160fe8c188fa80566580f279cc.
//
// Solidity: event UpgradeAgentSet(address agent)
func (_OldMystToken *OldMystTokenFilterer) WatchUpgradeAgentSet(opts *bind.WatchOpts, sink chan<- *OldMystTokenUpgradeAgentSet) (event.Subscription, error) {

	logs, sub, err := _OldMystToken.contract.WatchLogs(opts, "UpgradeAgentSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OldMystTokenUpgradeAgentSet)
				if err := _OldMystToken.contract.UnpackLog(event, "UpgradeAgentSet", log); err != nil {
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
func (_OldMystToken *OldMystTokenFilterer) ParseUpgradeAgentSet(log types.Log) (*OldMystTokenUpgradeAgentSet, error) {
	event := new(OldMystTokenUpgradeAgentSet)
	if err := _OldMystToken.contract.UnpackLog(event, "UpgradeAgentSet", log); err != nil {
		return nil, err
	}
	return event, nil
}
