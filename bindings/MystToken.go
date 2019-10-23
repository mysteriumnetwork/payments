/*
 * Copyright (C) 2019 The "MysteriumNetwork/payments" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
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
var MystTokenBin = "0x60806040526000600160146101000a81548160ff02191690831515021790555034801561002b57600080fd5b5033600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555033600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555033600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555061223a806101566000396000f3fe608060405234801561001057600080fd5b50600436106101585760003560e01c806370a08231116100c3578063c752ff621161007c578063c752ff6214610709578063d7e7088a14610727578063dd62ed3e1461076b578063e2301d02146107e3578063f2fde38b14610849578063ffeb7d751461088d57610158565b806370a08231146104ec5780638444b391146105445780638da5cb5b1461057057806395d89b41146105ba578063a9059cbb1461063d578063ac3cb72c146106a357610158565b806340c10f191161011557806340c10f191461033057806342c1867b1461037e57806343214675146103da57806345977d031461042a5780635de4ccb014610458578063600440cb146104a257610158565b806305d2035b1461015d57806306fdde031461017f578063095ea7b31461020257806318160ddd1461026857806323b872dd14610286578063313ce5671461030c575b600080fd5b6101656108d1565b604051808215151515815260200191505060405180910390f35b6101876108e4565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101c75780820151818401526020810190506101ac565b50505050905090810190601f1680156101f45780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61024e6004803603604081101561021857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061091d565b604051808215151515815260200191505060405180910390f35b610270610aa6565b6040518082815260200191505060405180910390f35b6102f26004803603606081101561029c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610aac565b604051808215151515815260200191505060405180910390f35b610314610d41565b604051808260ff1660ff16815260200191505060405180910390f35b61037c6004803603604081101561034657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610d46565b005b6103c06004803603602081101561039457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610fe5565b604051808215151515815260200191505060405180910390f35b610428600480360360408110156103f057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803515159060200190929190505050611005565b005b6104566004803603602081101561044057600080fd5b81019080803590602001909291905050506110d4565b005b61046061132d565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6104aa611353565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61052e6004803603602081101561050257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611379565b6040518082815260200191505060405180910390f35b61054c6113c2565b6040518082600481111561055c57fe5b60ff16815260200191505060405180910390f35b61057861143f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6105c2611465565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156106025780820151818401526020810190506105e7565b50505050905090810190601f16801561062f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6106896004803603604081101561065357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061149e565b604051808215151515815260200191505060405180910390f35b6106ef600480360360408110156106b957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061163d565b604051808215151515815260200191505060405180910390f35b61071161184c565b6040518082815260200191505060405180910390f35b6107696004803603602081101561073d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611852565b005b6107cd6004803603604081101561078157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611bcd565b6040518082815260200191505060405180910390f35b61082f600480360360408110156107f957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611c54565b604051808215151515815260200191505060405180910390f35b61088b6004803603602081101561085f57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611ef3565b005b6108cf600480360360208110156108a357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611fc6565b005b600160149054906101000a900460ff1681565b6040518060400160405280601481526020017f54657374204d797374657269756d20746f6b656e00000000000000000000000081525081565b60008082141580156109ac57506000600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414155b156109b657600080fd5b81600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040518082815260200191505060405180910390a36001905092915050565b60005481565b600080600760008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050610b77600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548461209e565b600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610c03600660008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205484612132565b600660008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610c508184612132565b600760008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef856040518082815260200191505060405180910390a360019150509392505050565b600881565b600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610d9c57600080fd5b600160149054906101000a900460ff1615610db657600080fd5b6000811415610dc457600080fd5b60005473__SafeMathLib___________________________6366098d4f9091836040518363ffffffff1660e01b8152600401808381526020018281526020019250505060206040518083038186803b158015610e1f57600080fd5b505af4158015610e33573d6000803e3d6000fd5b505050506040513d6020811015610e4957600080fd5b8101908080519060200190929190505050600081905550600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205473__SafeMathLib___________________________6366098d4f9091836040518363ffffffff1660e01b8152600401808381526020018281526020019250505060206040518083038186803b158015610ef857600080fd5b505af4158015610f0c573d6000803e3d6000fd5b505050506040513d6020811015610f2257600080fd5b8101908080519060200190929190505050600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055507f30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe8282604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a15050565b60056020528060005260406000206000915054906101000a900460ff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461105f57600080fd5b600160149054906101000a900460ff161561107957600080fd5b80600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b60006110de6113c2565b9050600360048111156110ed57fe5b8160048111156110f957fe5b148061111a575060048081111561110c57fe5b81600481111561111857fe5b145b61112357600080fd5b600082141561113157600080fd5b61117a600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205483612132565b600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506111c960005483612132565b6000819055506111db6004548361209e565b600481905550600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663753e88e533846040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b15801561128a57600080fd5b505af115801561129e573d6000803e3d6000fd5b50505050600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f7e5c344a8141a805725cb476f76c6953b842222b967edd1f78ddb6e8b3f397ac846040518082815260200191505060405180910390a35050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60008073ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415611423576002905061143c565b60006004541415611437576003905061143c565b600490505b90565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6040518060400160405280600581526020017f4d5953545400000000000000000000000000000000000000000000000000000081525081565b60006040600481016000369050146114b557600080fd5b6114fe600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205484612132565b600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555061158a600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548461209e565b600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508373ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef856040518082815260200191505060405180910390a3600191505092915050565b600060406004810160003690501461165457600080fd5b6000600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506116df818561209e565b600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546040518082815260200191505060405180910390a360019250505092915050565b60045481565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146118f8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a8152602001806121b6602a913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561193257600080fd5b60048081111561193e57fe5b6119466113c2565b600481111561195157fe5b14156119a8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260268152602001806121e06026913960400191505060405180910390fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166361d3d7a66040518163ffffffff1660e01b815260040160206040518083038186803b158015611a5157600080fd5b505afa158015611a65573d6000803e3d6000fd5b505050506040513d6020811015611a7b57600080fd5b8101908080519060200190929190505050611a9557600080fd5b600054600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634b2ba0dd6040518163ffffffff1660e01b815260040160206040518083038186803b158015611b0057600080fd5b505afa158015611b14573d6000803e3d6000fd5b505050506040513d6020811015611b2a57600080fd5b810190808051906020019092919050505014611b4557600080fd5b7f7845d5aa74cc410e35571258d954f23b82276e160fe8c188fa80566580f279cc600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b6000600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b6000604060048101600036905014611c6b57600080fd5b6000600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905080841115611d7b576000600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550611e06565b611d858185612132565b600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b8473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546040518082815260200191505060405180910390a360019250505092915050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611f4d57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614611fc35780600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561200057600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461205a57600080fd5b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008082840190508381101580156120b65750828110155b612128576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b6000828211156121aa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081525060200191505060405180910390fd5b81830390509291505056fe4f6e6c792061206d61737465722063616e2064657369676e61746520746865206e657874206167656e74557067726164652068617320616c726561647920626567756e20666f7220616e206167656e74a265627a7a72315820921fdcca1b5830427c91dee56ee0c41251c5b969a877ee640ae3be4fff1a750664736f6c634300050c0032"

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
