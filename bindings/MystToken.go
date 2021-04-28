/* Mysterium network payment library.
 *
 * Copyright (C) 2021 BlockDev AG
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

// MystTokenABI is the input ABI used to generate the binding from.
const MystTokenABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousDestination\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newDestination\",\"type\":\"address\"}],\"name\":\"FundsRecoveryDestinationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Upgrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"UpgradeAgentSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"master\",\"type\":\"address\"}],\"name\":\"UpgradeMasterSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isUpgradeAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenHolder\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"originalToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"originalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"upgradeFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradeMaster\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradeAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalUpgraded\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newUpgradeMaster\",\"type\":\"address\"}],\"name\":\"setUpgradeMaster\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"setUpgradeAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUpgradeState\",\"outputs\":[{\"internalType\":\"enumMystToken.UpgradeState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDestination\",\"type\":\"address\"}],\"name\":\"setFundsDestination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFundsDestination\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"claimTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MystTokenBin is the compiled bytecode used for deploying new contracts.
var MystTokenBin = "0x60c06040523480156200001157600080fd5b5060405162001f9d38038062001f9d833981810160405260208110156200003757600080fd5b50516001600160601b0319606082901b16608052604080516318160ddd60e01b815290516001600160a01b038316916318160ddd916004808301926020929190829003018186803b1580156200008c57600080fd5b505afa158015620000a1573d6000803e3d6000fd5b505050506040513d6020811015620000b857600080fd5b505160a052620000c7620001d3565b600080546001600160a01b0319166001600160a01b039290921691909117905560408051808201825260098152684d797374657269756d60b81b602091820152815180830190925260018252603160f81b9101527f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7f5a6b6ed89c9a1a7ef403de5fac83d634f25f39cd4318f8dabc0fdddabbba7bd17fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc662000188620001d7565b6040805160208082019690965280820194909452606084019290925260808301523060a0808401919091528151808403909101815260c09092019052805191012060055550620001db565b3390565b4690565b60805160601c60a051611d9b6200020260003980610b3052508061066a5250611d9b6000f3fe608060405234801561001057600080fd5b50600436106101e55760003560e01c8063600440cb1161010f578063a9059cbb116100a2578063dd62ed3e11610071578063dd62ed3e146105a3578063df8de3e7146105d1578063f58c5b6e146105f7578063ffeb7d75146105ff576101e5565b8063a9059cbb146104f8578063c752ff6214610524578063d505accf1461052c578063d7e7088a1461057d576101e5565b80637ecebe00116100de5780637ecebe00146104755780638444b3911461049b57806395d89b41146104c4578063a457c2d7146104cc576101e5565b8063600440cb1461041357806361d3d7a61461041b57806370a0823114610423578063753e88e514610449576101e5565b8063313ce5671161018757806342966c681161015657806342966c68146103c957806345977d03146103e65780634b2ba0dd146104035780635de4ccb01461040b576101e5565b8063313ce5671461034b5780633644e51514610369578063395093511461037157806340c10f191461039d576101e5565b806318160ddd116101c357806318160ddd146102cb578063238e130a146102e557806323b872dd1461030d57806330adf81f14610343576101e5565b806306fdde03146101ea578063095ea7b3146102675780630e7c1cb5146102a7575b600080fd5b6101f2610625565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561022c578181015183820152602001610214565b50505050905090810190601f1680156102595780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102936004803603604081101561027d57600080fd5b506001600160a01b03813516906020013561064a565b604080519115158252519081900360200190f35b6102af610668565b604080516001600160a01b039092168252519081900360200190f35b6102d361068d565b60408051918252519081900360200190f35b61030b600480360360208110156102fb57600080fd5b50356001600160a01b0316610693565b005b6102936004803603606081101561032357600080fd5b506001600160a01b03813581169160208101359091169060400135610781565b6102d361088c565b6103536108b0565b6040805160ff9092168252519081900360200190f35b6102d36108b5565b6102936004803603604081101561038757600080fd5b506001600160a01b0381351690602001356108bb565b61030b600480360360408110156103b357600080fd5b506001600160a01b038135169060200135610909565b61030b600480360360208110156103df57600080fd5b5035610986565b61030b600480360360208110156103fc57600080fd5b503561099a565b6102d3610b2e565b6102af610b52565b6102af610b61565b610293610b70565b6102d36004803603602081101561043957600080fd5b50356001600160a01b0316610b75565b61030b6004803603604081101561045f57600080fd5b506001600160a01b038135169060200135610b90565b6102d36004803603602081101561048b57600080fd5b50356001600160a01b0316610c5b565b6104a3610c6d565b604051808260058111156104b357fe5b815260200191505060405180910390f35b6101f2610cb3565b610293600480360360408110156104e257600080fd5b506001600160a01b038135169060200135610cd3565b6102936004803603604081101561050e57600080fd5b506001600160a01b038135169060200135610d3b565b6102d3610d4f565b61030b600480360360e081101561054257600080fd5b506001600160a01b03813581169160208101359091169060408101359060608101359060ff6080820135169060a08101359060c00135610d55565b61030b6004803603602081101561059357600080fd5b50356001600160a01b0316610f59565b6102d3600480360360408110156105b957600080fd5b506001600160a01b0381358116916020013516611227565b61030b600480360360208110156105e757600080fd5b50356001600160a01b0316611252565b6102af61136b565b61030b6004803603602081101561061557600080fd5b50356001600160a01b031661137a565b604051806040016040528060098152602001684d797374657269756d60b81b81525081565b600061065e61065761145c565b8484611460565b5060015b92915050565b7f00000000000000000000000000000000000000000000000000000000000000005b90565b60045490565b6000546001600160a01b03166106a761145c565b6001600160a01b0316146106ec5760405162461bcd60e51b815260040180806020018281038252602d815260200180611ba2602d913960400191505060405180910390fd5b6001600160a01b0381166107315760405162461bcd60e51b815260040180806020018281038252602e815260200180611c1e602e913960400191505060405180910390fd5b600880546001600160a01b0319166001600160a01b0383811691821792839055604051919216907f2e1db88922daae16be4e3c1a1f4bfab0cf6741938844967bd985ac8b2a12c80490600090a350565b60006001600160a01b0384166107c85760405162461bcd60e51b8152600401808060200182810382526024815260200180611b026024913960400191505060405180910390fd5b60006107d261145c565b9050806001600160a01b0316856001600160a01b03161415801561081d57506001600160a01b0380861660009081526007602090815260408083209385168352929052205460001914155b1561087657610876858261087186604051806060016040528060278152602001611b26602791396001600160a01b03808c166000908152600760209081526040808320938b1683529290522054919061154c565b611460565b6108818585856115e3565b506001949350505050565b7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c981565b601281565b60055481565b600061065e6108c861145c565b8461087185600760006108d961145c565b6001600160a01b03908116825260208083019390935260409182016000908120918c1681529252902054906116c9565b6000546001600160a01b031661091d61145c565b6001600160a01b031614610978576040805162461bcd60e51b815260206004820152601c60248201527f4d5953543a206f6e6c792061206d61737465722063616e206d696e7400000000604482015290519081900360640190fd5b610982828261172a565b5050565b61099761099161145c565b8261184f565b50565b60006109a4610c6d565b905060038160058111156109b457fe5b14806109cb575060048160058111156109c957fe5b145b610a065760405162461bcd60e51b8152600401808060200182810382526025815260200180611bf96025913960400191505060405180910390fd5b81610a425760405162461bcd60e51b815260040180806020018281038252602d815260200180611ceb602d913960400191505060405180910390fd5b6000610a4c61145c565b9050610a58818461184f565b600254610a6590846116c9565b6002556001546040805163753e88e560e01b81526001600160a01b038481166004830152602482018790529151919092169163753e88e591604480830192600092919082900301818387803b158015610abd57600080fd5b505af1158015610ad1573d6000803e3d6000fd5b50505050806001600160a01b03167f7e5c344a8141a805725cb476f76c6953b842222b967edd1f78ddb6e8b3f397ac610b08610b52565b604080516001600160a01b039092168252602082018790528051918290030190a2505050565b7f000000000000000000000000000000000000000000000000000000000000000090565b6001546001600160a01b031690565b6000546001600160a01b031690565b600181565b6001600160a01b031660009081526003602052604090205490565b610b98610668565b6001600160a01b0316336001600160a01b031614610be75760405162461bcd60e51b8152600401808060200182810382526028815260200180611ab96028913960400191505060405180910390fd5b610bff82610bfa836402540be400611994565b61172a565b610c166402540be400610c10610b2e565b90611994565b610c1e61068d565b11156109825760405162461bcd60e51b8152600401808060200182810382526032815260200180611c4c6032913960400191505060405180910390fd5b60066020526000908152604090205481565b6001546000906001600160a01b0316610c885750600261068a565b600254610c975750600361068a565b610c9f61068d565b610cab5750600561068a565b50600461068a565b60405180604001604052806004815260200163135654d560e21b81525081565b600061065e610ce061145c565b8461087185604051806060016040528060258152602001611d186025913960076000610d0a61145c565b6001600160a01b03908116825260208083019390935260409182016000908120918d1681529252902054919061154c565b600061065e610d4861145c565b84846115e3565b60025490565b42841015610da1576040805162461bcd60e51b8152602060048201526014602482015273135654d50e8814195c9b5a5d08195e1c1a5c995960621b604482015290519081900360640190fd5b6005546001600160a01b0380891660008181526006602090815260408083208054600180820190925582517f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98186015280840196909652958d166060860152608085018c905260a085019590955260c08085018b90528151808603909101815260e08501825280519083012061190160f01b6101008601526101028501969096526101228085019690965280518085039096018652610142840180825286519683019690962095839052610162840180825286905260ff89166101828501526101a284018890526101c28401879052519193926101e280820193601f1981019281900390910190855afa158015610ebc573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811615801590610ef25750886001600160a01b0316816001600160a01b0316145b610f43576040805162461bcd60e51b815260206004820152601760248201527f4d5953543a20696e76616c6964207369676e6174757265000000000000000000604482015290519081900360640190fd5b610f4e898989611460565b505050505050505050565b6000546001600160a01b0316610f6d61145c565b6001600160a01b031614610fb25760405162461bcd60e51b8152600401808060200182810382526030815260200180611a896030913960400191505060405180910390fd5b6001600160a01b038116610ff75760405162461bcd60e51b8152600401808060200182810382526029815260200180611cc26029913960400191505060405180910390fd5b6004611001610c6d565b600581111561100c57fe5b141561105f576040805162461bcd60e51b815260206004820152601f60248201527f4d5953543a20757067726164652068617320616c726561647920626567756e00604482015290519081900360640190fd5b600180546001600160a01b0319166001600160a01b038381169190911791829055604080516330e9ebd360e11b8152905192909116916361d3d7a691600480820192602092909190829003018186803b1580156110bb57600080fd5b505afa1580156110cf573d6000803e3d6000fd5b505050506040513d60208110156110e557600080fd5b50516111225760405162461bcd60e51b8152600401808060200182810382526034815260200180611b6e6034913960400191505060405180910390fd5b61112a61068d565b600160009054906101000a90046001600160a01b03166001600160a01b0316634b2ba0dd6040518163ffffffff1660e01b815260040160206040518083038186803b15801561117857600080fd5b505afa15801561118c573d6000803e3d6000fd5b505050506040513d60208110156111a257600080fd5b5051146111e05760405162461bcd60e51b8152600401808060200182810382526034815260200180611a556034913960400191505060405180910390fd5b7f7845d5aa74cc410e35571258d954f23b82276e160fe8c188fa80566580f279cc611209610b52565b604080516001600160a01b039092168252519081900360200190a150565b6001600160a01b03918216600090815260076020908152604080832093909416825291909152205490565b6008546001600160a01b031661126757600080fd5b6000816001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b1580156112b657600080fd5b505afa1580156112ca573d6000803e3d6000fd5b505050506040513d60208110156112e057600080fd5b50516008546040805163a9059cbb60e01b81526001600160a01b0392831660048201526024810184905290519293509084169163a9059cbb916044808201926020929091908290030181600087803b15801561133b57600080fd5b505af115801561134f573d6000803e3d6000fd5b505050506040513d602081101561136557600080fd5b50505050565b6008546001600160a01b031690565b6001600160a01b0381166113bf5760405162461bcd60e51b815260040180806020018281038252602a815260200180611bcf602a913960400191505060405180910390fd5b6000546001600160a01b03166113d361145c565b6001600160a01b0316146114185760405162461bcd60e51b8152600401808060200182810382526029815260200180611d3d6029913960400191505060405180910390fd5b600080546001600160a01b0319166001600160a01b0383161790557f0bae748e6d38d2b1532af619519837d91d74845ad693f6f229677b4ac20b2d50611209610b61565b3390565b6001600160a01b0383166114a55760405162461bcd60e51b8152600401808060200182810382526023815260200180611c9f6023913960400191505060405180910390fd5b6001600160a01b0382166114ea5760405162461bcd60e51b8152600401808060200182810382526021815260200180611ae16021913960400191505060405180910390fd5b6001600160a01b03808416600081815260076020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b600081848411156115db5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156115a0578181015183820152602001611588565b50505050905090810190601f1680156115cd5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b6001600160a01b038216611600576115fb838261184f565b6116c4565b61163d81604051806060016040528060258152602001611a30602591396001600160a01b038616600090815260036020526040902054919061154c565b6001600160a01b03808516600090815260036020526040808220939093559084168152205461166c90826116c9565b6001600160a01b0380841660008181526003602090815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a35b505050565b600082820183811015611723576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b6001600160a01b038216611785576040805162461bcd60e51b815260206004820152601e60248201527f4d5953543a206d696e7420746f20746865207a65726f20616464726573730000604482015290519081900360640190fd5b60045461179290826116c9565b6004556001600160a01b0382166000908152600360205260409020546117b890826116c9565b6001600160a01b038316600081815260036020908152604091829020939093558051848152905191927f30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe92918290030190a26040805182815290516001600160a01b038416916000917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35050565b6001600160a01b0382166118aa576040805162461bcd60e51b815260206004820181905260248201527f4d5953543a206275726e2066726f6d20746865207a65726f2061646472657373604482015290519081900360640190fd5b6118e781604051806060016040528060218152602001611c7e602191396001600160a01b038516600090815260036020526040902054919061154c565b6001600160a01b03831660009081526003602052604090205560045461190d90826119ed565b6004556040805182815290516000916001600160a01b038516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a36040805182815290516001600160a01b038416917f696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7919081900360200190a25050565b6000826119a357506000610662565b828202828482816119b057fe5b04146117235760405162461bcd60e51b8152600401808060200182810382526021815260200180611b4d6021913960400191505060405180910390fd5b600061172383836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081525061154c56fe4d5953543a207472616e7366657220616d6f756e7420657863656564732062616c616e63654d5953543a2075706772616465206167656e742073686f756c64206b6e6f7720746f6b656e277320746f74616c20737570706c794d5953543a206f6e6c792061206d61737465722063616e2064657369676e61746520746865206e657874206167656e746f6e6c79206f726967696e616c20746f6b656e2063616e2063616c6c207570677261646546726f6d4d5953543a20617070726f766520746f20746865207a65726f20616464726573734d5953543a207472616e736665722066726f6d20746865207a65726f20616464726573734d5953543a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e6365536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f774d5953543a206167656e742073686f756c6420696d706c656d656e742049557067726164654167656e7420696e746572666163654d5953543a206f6e6c792061206d61737465722063616e207365742066756e64732064657374696e6174696f6e4d5953543a2075706772616465206d61737465722063616e2774206265207a65726f20616464726573734d5953543a20746f6b656e206973206e6f7420696e20757067726164696e672073746174654d5953543a2066756e64732064657374696e6174696f6e2063616e2774206265207a65726f20616464726565737363616e206e6f74206d696e74206d6f726520746f6b656e73207468616e20696e206f726967696e616c20636f6e74726163744d5953543a206275726e20616d6f756e7420657863656564732062616c616e63654d5953543a20617070726f76652066726f6d20746865207a65726f20616464726573734d5953543a2075706772616465206167656e742063616e2774206265207a65726f20616464726573734d5953543a2075706772616461626c6520616d6f756e742073686f756c64206265206d6f7265207468616e203045524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726f4d5953543a206f6e6c792075706772616465206d61737465722063616e20736574206e6577206f6e65a26469706673582212201cda6413f6ed0a2088f3c1e72f0dfbc0313b610e07cc6c11806a4dcd5271815064736f6c63430007040033"

// DeployMystToken deploys a new Ethereum contract, binding an instance of MystToken to it.
func DeployMystToken(auth *bind.TransactOpts, backend bind.ContractBackend, tokenAddress common.Address) (common.Address, *types.Transaction, *MystToken, error) {
	parsed, err := abi.JSON(strings.NewReader(MystTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MystTokenBin), backend, tokenAddress)
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
func (_MystToken *MystTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_MystToken *MystTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_MystToken *MystTokenCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MystToken.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_MystToken *MystTokenSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _MystToken.Contract.DOMAINSEPARATOR(&_MystToken.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_MystToken *MystTokenCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _MystToken.Contract.DOMAINSEPARATOR(&_MystToken.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_MystToken *MystTokenCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MystToken.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_MystToken *MystTokenSession) PERMITTYPEHASH() ([32]byte, error) {
	return _MystToken.Contract.PERMITTYPEHASH(&_MystToken.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_MystToken *MystTokenCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _MystToken.Contract.PERMITTYPEHASH(&_MystToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_MystToken *MystTokenCaller) Allowance(opts *bind.CallOpts, holder common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MystToken.contract.Call(opts, &out, "allowance", holder, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_MystToken *MystTokenSession) Allowance(holder common.Address, spender common.Address) (*big.Int, error) {
	return _MystToken.Contract.Allowance(&_MystToken.CallOpts, holder, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_MystToken *MystTokenCallerSession) Allowance(holder common.Address, spender common.Address) (*big.Int, error) {
	return _MystToken.Contract.Allowance(&_MystToken.CallOpts, holder, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenHolder) view returns(uint256)
func (_MystToken *MystTokenCaller) BalanceOf(opts *bind.CallOpts, tokenHolder common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MystToken.contract.Call(opts, &out, "balanceOf", tokenHolder)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenHolder) view returns(uint256)
func (_MystToken *MystTokenSession) BalanceOf(tokenHolder common.Address) (*big.Int, error) {
	return _MystToken.Contract.BalanceOf(&_MystToken.CallOpts, tokenHolder)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenHolder) view returns(uint256)
func (_MystToken *MystTokenCallerSession) BalanceOf(tokenHolder common.Address) (*big.Int, error) {
	return _MystToken.Contract.BalanceOf(&_MystToken.CallOpts, tokenHolder)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MystToken *MystTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MystToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MystToken *MystTokenSession) Decimals() (uint8, error) {
	return _MystToken.Contract.Decimals(&_MystToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MystToken *MystTokenCallerSession) Decimals() (uint8, error) {
	return _MystToken.Contract.Decimals(&_MystToken.CallOpts)
}

// GetFundsDestination is a free data retrieval call binding the contract method 0xf58c5b6e.
//
// Solidity: function getFundsDestination() view returns(address)
func (_MystToken *MystTokenCaller) GetFundsDestination(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MystToken.contract.Call(opts, &out, "getFundsDestination")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFundsDestination is a free data retrieval call binding the contract method 0xf58c5b6e.
//
// Solidity: function getFundsDestination() view returns(address)
func (_MystToken *MystTokenSession) GetFundsDestination() (common.Address, error) {
	return _MystToken.Contract.GetFundsDestination(&_MystToken.CallOpts)
}

// GetFundsDestination is a free data retrieval call binding the contract method 0xf58c5b6e.
//
// Solidity: function getFundsDestination() view returns(address)
func (_MystToken *MystTokenCallerSession) GetFundsDestination() (common.Address, error) {
	return _MystToken.Contract.GetFundsDestination(&_MystToken.CallOpts)
}

// GetUpgradeState is a free data retrieval call binding the contract method 0x8444b391.
//
// Solidity: function getUpgradeState() view returns(uint8)
func (_MystToken *MystTokenCaller) GetUpgradeState(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MystToken.contract.Call(opts, &out, "getUpgradeState")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetUpgradeState is a free data retrieval call binding the contract method 0x8444b391.
//
// Solidity: function getUpgradeState() view returns(uint8)
func (_MystToken *MystTokenSession) GetUpgradeState() (uint8, error) {
	return _MystToken.Contract.GetUpgradeState(&_MystToken.CallOpts)
}

// GetUpgradeState is a free data retrieval call binding the contract method 0x8444b391.
//
// Solidity: function getUpgradeState() view returns(uint8)
func (_MystToken *MystTokenCallerSession) GetUpgradeState() (uint8, error) {
	return _MystToken.Contract.GetUpgradeState(&_MystToken.CallOpts)
}

// IsUpgradeAgent is a free data retrieval call binding the contract method 0x61d3d7a6.
//
// Solidity: function isUpgradeAgent() view returns(bool)
func (_MystToken *MystTokenCaller) IsUpgradeAgent(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MystToken.contract.Call(opts, &out, "isUpgradeAgent")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUpgradeAgent is a free data retrieval call binding the contract method 0x61d3d7a6.
//
// Solidity: function isUpgradeAgent() view returns(bool)
func (_MystToken *MystTokenSession) IsUpgradeAgent() (bool, error) {
	return _MystToken.Contract.IsUpgradeAgent(&_MystToken.CallOpts)
}

// IsUpgradeAgent is a free data retrieval call binding the contract method 0x61d3d7a6.
//
// Solidity: function isUpgradeAgent() view returns(bool)
func (_MystToken *MystTokenCallerSession) IsUpgradeAgent() (bool, error) {
	return _MystToken.Contract.IsUpgradeAgent(&_MystToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MystToken *MystTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MystToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MystToken *MystTokenSession) Name() (string, error) {
	return _MystToken.Contract.Name(&_MystToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MystToken *MystTokenCallerSession) Name() (string, error) {
	return _MystToken.Contract.Name(&_MystToken.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_MystToken *MystTokenCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MystToken.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_MystToken *MystTokenSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _MystToken.Contract.Nonces(&_MystToken.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_MystToken *MystTokenCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _MystToken.Contract.Nonces(&_MystToken.CallOpts, arg0)
}

// OriginalSupply is a free data retrieval call binding the contract method 0x4b2ba0dd.
//
// Solidity: function originalSupply() view returns(uint256)
func (_MystToken *MystTokenCaller) OriginalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MystToken.contract.Call(opts, &out, "originalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OriginalSupply is a free data retrieval call binding the contract method 0x4b2ba0dd.
//
// Solidity: function originalSupply() view returns(uint256)
func (_MystToken *MystTokenSession) OriginalSupply() (*big.Int, error) {
	return _MystToken.Contract.OriginalSupply(&_MystToken.CallOpts)
}

// OriginalSupply is a free data retrieval call binding the contract method 0x4b2ba0dd.
//
// Solidity: function originalSupply() view returns(uint256)
func (_MystToken *MystTokenCallerSession) OriginalSupply() (*big.Int, error) {
	return _MystToken.Contract.OriginalSupply(&_MystToken.CallOpts)
}

// OriginalToken is a free data retrieval call binding the contract method 0x0e7c1cb5.
//
// Solidity: function originalToken() view returns(address)
func (_MystToken *MystTokenCaller) OriginalToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MystToken.contract.Call(opts, &out, "originalToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OriginalToken is a free data retrieval call binding the contract method 0x0e7c1cb5.
//
// Solidity: function originalToken() view returns(address)
func (_MystToken *MystTokenSession) OriginalToken() (common.Address, error) {
	return _MystToken.Contract.OriginalToken(&_MystToken.CallOpts)
}

// OriginalToken is a free data retrieval call binding the contract method 0x0e7c1cb5.
//
// Solidity: function originalToken() view returns(address)
func (_MystToken *MystTokenCallerSession) OriginalToken() (common.Address, error) {
	return _MystToken.Contract.OriginalToken(&_MystToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MystToken *MystTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MystToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MystToken *MystTokenSession) Symbol() (string, error) {
	return _MystToken.Contract.Symbol(&_MystToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MystToken *MystTokenCallerSession) Symbol() (string, error) {
	return _MystToken.Contract.Symbol(&_MystToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MystToken *MystTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MystToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MystToken *MystTokenSession) TotalSupply() (*big.Int, error) {
	return _MystToken.Contract.TotalSupply(&_MystToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MystToken *MystTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _MystToken.Contract.TotalSupply(&_MystToken.CallOpts)
}

// TotalUpgraded is a free data retrieval call binding the contract method 0xc752ff62.
//
// Solidity: function totalUpgraded() view returns(uint256)
func (_MystToken *MystTokenCaller) TotalUpgraded(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MystToken.contract.Call(opts, &out, "totalUpgraded")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalUpgraded is a free data retrieval call binding the contract method 0xc752ff62.
//
// Solidity: function totalUpgraded() view returns(uint256)
func (_MystToken *MystTokenSession) TotalUpgraded() (*big.Int, error) {
	return _MystToken.Contract.TotalUpgraded(&_MystToken.CallOpts)
}

// TotalUpgraded is a free data retrieval call binding the contract method 0xc752ff62.
//
// Solidity: function totalUpgraded() view returns(uint256)
func (_MystToken *MystTokenCallerSession) TotalUpgraded() (*big.Int, error) {
	return _MystToken.Contract.TotalUpgraded(&_MystToken.CallOpts)
}

// UpgradeAgent is a free data retrieval call binding the contract method 0x5de4ccb0.
//
// Solidity: function upgradeAgent() view returns(address)
func (_MystToken *MystTokenCaller) UpgradeAgent(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MystToken.contract.Call(opts, &out, "upgradeAgent")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UpgradeAgent is a free data retrieval call binding the contract method 0x5de4ccb0.
//
// Solidity: function upgradeAgent() view returns(address)
func (_MystToken *MystTokenSession) UpgradeAgent() (common.Address, error) {
	return _MystToken.Contract.UpgradeAgent(&_MystToken.CallOpts)
}

// UpgradeAgent is a free data retrieval call binding the contract method 0x5de4ccb0.
//
// Solidity: function upgradeAgent() view returns(address)
func (_MystToken *MystTokenCallerSession) UpgradeAgent() (common.Address, error) {
	return _MystToken.Contract.UpgradeAgent(&_MystToken.CallOpts)
}

// UpgradeMaster is a free data retrieval call binding the contract method 0x600440cb.
//
// Solidity: function upgradeMaster() view returns(address)
func (_MystToken *MystTokenCaller) UpgradeMaster(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MystToken.contract.Call(opts, &out, "upgradeMaster")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UpgradeMaster is a free data retrieval call binding the contract method 0x600440cb.
//
// Solidity: function upgradeMaster() view returns(address)
func (_MystToken *MystTokenSession) UpgradeMaster() (common.Address, error) {
	return _MystToken.Contract.UpgradeMaster(&_MystToken.CallOpts)
}

// UpgradeMaster is a free data retrieval call binding the contract method 0x600440cb.
//
// Solidity: function upgradeMaster() view returns(address)
func (_MystToken *MystTokenCallerSession) UpgradeMaster() (common.Address, error) {
	return _MystToken.Contract.UpgradeMaster(&_MystToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MystToken *MystTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MystToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MystToken *MystTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.Approve(&_MystToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MystToken *MystTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.Approve(&_MystToken.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_MystToken *MystTokenTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _MystToken.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_MystToken *MystTokenSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.Burn(&_MystToken.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_MystToken *MystTokenTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.Burn(&_MystToken.TransactOpts, amount)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0xdf8de3e7.
//
// Solidity: function claimTokens(address token) returns()
func (_MystToken *MystTokenTransactor) ClaimTokens(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _MystToken.contract.Transact(opts, "claimTokens", token)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0xdf8de3e7.
//
// Solidity: function claimTokens(address token) returns()
func (_MystToken *MystTokenSession) ClaimTokens(token common.Address) (*types.Transaction, error) {
	return _MystToken.Contract.ClaimTokens(&_MystToken.TransactOpts, token)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0xdf8de3e7.
//
// Solidity: function claimTokens(address token) returns()
func (_MystToken *MystTokenTransactorSession) ClaimTokens(token common.Address) (*types.Transaction, error) {
	return _MystToken.Contract.ClaimTokens(&_MystToken.TransactOpts, token)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_MystToken *MystTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MystToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_MystToken *MystTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.DecreaseAllowance(&_MystToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_MystToken *MystTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.DecreaseAllowance(&_MystToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_MystToken *MystTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MystToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_MystToken *MystTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.IncreaseAllowance(&_MystToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_MystToken *MystTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.IncreaseAllowance(&_MystToken.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_MystToken *MystTokenTransactor) Mint(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MystToken.contract.Transact(opts, "mint", _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_MystToken *MystTokenSession) Mint(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.Mint(&_MystToken.TransactOpts, _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_MystToken *MystTokenTransactorSession) Mint(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.Mint(&_MystToken.TransactOpts, _account, _amount)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address holder, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_MystToken *MystTokenTransactor) Permit(opts *bind.TransactOpts, holder common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _MystToken.contract.Transact(opts, "permit", holder, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address holder, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_MystToken *MystTokenSession) Permit(holder common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _MystToken.Contract.Permit(&_MystToken.TransactOpts, holder, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address holder, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_MystToken *MystTokenTransactorSession) Permit(holder common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _MystToken.Contract.Permit(&_MystToken.TransactOpts, holder, spender, value, deadline, v, r, s)
}

// SetFundsDestination is a paid mutator transaction binding the contract method 0x238e130a.
//
// Solidity: function setFundsDestination(address newDestination) returns()
func (_MystToken *MystTokenTransactor) SetFundsDestination(opts *bind.TransactOpts, newDestination common.Address) (*types.Transaction, error) {
	return _MystToken.contract.Transact(opts, "setFundsDestination", newDestination)
}

// SetFundsDestination is a paid mutator transaction binding the contract method 0x238e130a.
//
// Solidity: function setFundsDestination(address newDestination) returns()
func (_MystToken *MystTokenSession) SetFundsDestination(newDestination common.Address) (*types.Transaction, error) {
	return _MystToken.Contract.SetFundsDestination(&_MystToken.TransactOpts, newDestination)
}

// SetFundsDestination is a paid mutator transaction binding the contract method 0x238e130a.
//
// Solidity: function setFundsDestination(address newDestination) returns()
func (_MystToken *MystTokenTransactorSession) SetFundsDestination(newDestination common.Address) (*types.Transaction, error) {
	return _MystToken.Contract.SetFundsDestination(&_MystToken.TransactOpts, newDestination)
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
// Solidity: function setUpgradeMaster(address newUpgradeMaster) returns()
func (_MystToken *MystTokenTransactor) SetUpgradeMaster(opts *bind.TransactOpts, newUpgradeMaster common.Address) (*types.Transaction, error) {
	return _MystToken.contract.Transact(opts, "setUpgradeMaster", newUpgradeMaster)
}

// SetUpgradeMaster is a paid mutator transaction binding the contract method 0xffeb7d75.
//
// Solidity: function setUpgradeMaster(address newUpgradeMaster) returns()
func (_MystToken *MystTokenSession) SetUpgradeMaster(newUpgradeMaster common.Address) (*types.Transaction, error) {
	return _MystToken.Contract.SetUpgradeMaster(&_MystToken.TransactOpts, newUpgradeMaster)
}

// SetUpgradeMaster is a paid mutator transaction binding the contract method 0xffeb7d75.
//
// Solidity: function setUpgradeMaster(address newUpgradeMaster) returns()
func (_MystToken *MystTokenTransactorSession) SetUpgradeMaster(newUpgradeMaster common.Address) (*types.Transaction, error) {
	return _MystToken.Contract.SetUpgradeMaster(&_MystToken.TransactOpts, newUpgradeMaster)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_MystToken *MystTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MystToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_MystToken *MystTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.Transfer(&_MystToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_MystToken *MystTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.Transfer(&_MystToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address holder, address recipient, uint256 amount) returns(bool)
func (_MystToken *MystTokenTransactor) TransferFrom(opts *bind.TransactOpts, holder common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MystToken.contract.Transact(opts, "transferFrom", holder, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address holder, address recipient, uint256 amount) returns(bool)
func (_MystToken *MystTokenSession) TransferFrom(holder common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.TransferFrom(&_MystToken.TransactOpts, holder, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address holder, address recipient, uint256 amount) returns(bool)
func (_MystToken *MystTokenTransactorSession) TransferFrom(holder common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.TransferFrom(&_MystToken.TransactOpts, holder, recipient, amount)
}

// Upgrade is a paid mutator transaction binding the contract method 0x45977d03.
//
// Solidity: function upgrade(uint256 amount) returns()
func (_MystToken *MystTokenTransactor) Upgrade(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _MystToken.contract.Transact(opts, "upgrade", amount)
}

// Upgrade is a paid mutator transaction binding the contract method 0x45977d03.
//
// Solidity: function upgrade(uint256 amount) returns()
func (_MystToken *MystTokenSession) Upgrade(amount *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.Upgrade(&_MystToken.TransactOpts, amount)
}

// Upgrade is a paid mutator transaction binding the contract method 0x45977d03.
//
// Solidity: function upgrade(uint256 amount) returns()
func (_MystToken *MystTokenTransactorSession) Upgrade(amount *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.Upgrade(&_MystToken.TransactOpts, amount)
}

// UpgradeFrom is a paid mutator transaction binding the contract method 0x753e88e5.
//
// Solidity: function upgradeFrom(address _account, uint256 _value) returns()
func (_MystToken *MystTokenTransactor) UpgradeFrom(opts *bind.TransactOpts, _account common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MystToken.contract.Transact(opts, "upgradeFrom", _account, _value)
}

// UpgradeFrom is a paid mutator transaction binding the contract method 0x753e88e5.
//
// Solidity: function upgradeFrom(address _account, uint256 _value) returns()
func (_MystToken *MystTokenSession) UpgradeFrom(_account common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.UpgradeFrom(&_MystToken.TransactOpts, _account, _value)
}

// UpgradeFrom is a paid mutator transaction binding the contract method 0x753e88e5.
//
// Solidity: function upgradeFrom(address _account, uint256 _value) returns()
func (_MystToken *MystTokenTransactorSession) UpgradeFrom(_account common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.UpgradeFrom(&_MystToken.TransactOpts, _account, _value)
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
	event.Raw = log
	return event, nil
}

// MystTokenBurnedIterator is returned from FilterBurned and is used to iterate over the raw logs and unpacked data for Burned events raised by the MystToken contract.
type MystTokenBurnedIterator struct {
	Event *MystTokenBurned // Event containing the contract specifics and raw log

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
func (it *MystTokenBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MystTokenBurned)
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
		it.Event = new(MystTokenBurned)
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
func (it *MystTokenBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MystTokenBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MystTokenBurned represents a Burned event raised by the MystToken contract.
type MystTokenBurned struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurned is a free log retrieval operation binding the contract event 0x696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7.
//
// Solidity: event Burned(address indexed from, uint256 amount)
func (_MystToken *MystTokenFilterer) FilterBurned(opts *bind.FilterOpts, from []common.Address) (*MystTokenBurnedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _MystToken.contract.FilterLogs(opts, "Burned", fromRule)
	if err != nil {
		return nil, err
	}
	return &MystTokenBurnedIterator{contract: _MystToken.contract, event: "Burned", logs: logs, sub: sub}, nil
}

// WatchBurned is a free log subscription operation binding the contract event 0x696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7.
//
// Solidity: event Burned(address indexed from, uint256 amount)
func (_MystToken *MystTokenFilterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *MystTokenBurned, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _MystToken.contract.WatchLogs(opts, "Burned", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MystTokenBurned)
				if err := _MystToken.contract.UnpackLog(event, "Burned", log); err != nil {
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

// ParseBurned is a log parse operation binding the contract event 0x696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7.
//
// Solidity: event Burned(address indexed from, uint256 amount)
func (_MystToken *MystTokenFilterer) ParseBurned(log types.Log) (*MystTokenBurned, error) {
	event := new(MystTokenBurned)
	if err := _MystToken.contract.UnpackLog(event, "Burned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MystTokenFundsRecoveryDestinationChangedIterator is returned from FilterFundsRecoveryDestinationChanged and is used to iterate over the raw logs and unpacked data for FundsRecoveryDestinationChanged events raised by the MystToken contract.
type MystTokenFundsRecoveryDestinationChangedIterator struct {
	Event *MystTokenFundsRecoveryDestinationChanged // Event containing the contract specifics and raw log

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
func (it *MystTokenFundsRecoveryDestinationChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MystTokenFundsRecoveryDestinationChanged)
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
		it.Event = new(MystTokenFundsRecoveryDestinationChanged)
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
func (it *MystTokenFundsRecoveryDestinationChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MystTokenFundsRecoveryDestinationChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MystTokenFundsRecoveryDestinationChanged represents a FundsRecoveryDestinationChanged event raised by the MystToken contract.
type MystTokenFundsRecoveryDestinationChanged struct {
	PreviousDestination common.Address
	NewDestination      common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterFundsRecoveryDestinationChanged is a free log retrieval operation binding the contract event 0x2e1db88922daae16be4e3c1a1f4bfab0cf6741938844967bd985ac8b2a12c804.
//
// Solidity: event FundsRecoveryDestinationChanged(address indexed previousDestination, address indexed newDestination)
func (_MystToken *MystTokenFilterer) FilterFundsRecoveryDestinationChanged(opts *bind.FilterOpts, previousDestination []common.Address, newDestination []common.Address) (*MystTokenFundsRecoveryDestinationChangedIterator, error) {

	var previousDestinationRule []interface{}
	for _, previousDestinationItem := range previousDestination {
		previousDestinationRule = append(previousDestinationRule, previousDestinationItem)
	}
	var newDestinationRule []interface{}
	for _, newDestinationItem := range newDestination {
		newDestinationRule = append(newDestinationRule, newDestinationItem)
	}

	logs, sub, err := _MystToken.contract.FilterLogs(opts, "FundsRecoveryDestinationChanged", previousDestinationRule, newDestinationRule)
	if err != nil {
		return nil, err
	}
	return &MystTokenFundsRecoveryDestinationChangedIterator{contract: _MystToken.contract, event: "FundsRecoveryDestinationChanged", logs: logs, sub: sub}, nil
}

// WatchFundsRecoveryDestinationChanged is a free log subscription operation binding the contract event 0x2e1db88922daae16be4e3c1a1f4bfab0cf6741938844967bd985ac8b2a12c804.
//
// Solidity: event FundsRecoveryDestinationChanged(address indexed previousDestination, address indexed newDestination)
func (_MystToken *MystTokenFilterer) WatchFundsRecoveryDestinationChanged(opts *bind.WatchOpts, sink chan<- *MystTokenFundsRecoveryDestinationChanged, previousDestination []common.Address, newDestination []common.Address) (event.Subscription, error) {

	var previousDestinationRule []interface{}
	for _, previousDestinationItem := range previousDestination {
		previousDestinationRule = append(previousDestinationRule, previousDestinationItem)
	}
	var newDestinationRule []interface{}
	for _, newDestinationItem := range newDestination {
		newDestinationRule = append(newDestinationRule, newDestinationItem)
	}

	logs, sub, err := _MystToken.contract.WatchLogs(opts, "FundsRecoveryDestinationChanged", previousDestinationRule, newDestinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MystTokenFundsRecoveryDestinationChanged)
				if err := _MystToken.contract.UnpackLog(event, "FundsRecoveryDestinationChanged", log); err != nil {
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

// ParseFundsRecoveryDestinationChanged is a log parse operation binding the contract event 0x2e1db88922daae16be4e3c1a1f4bfab0cf6741938844967bd985ac8b2a12c804.
//
// Solidity: event FundsRecoveryDestinationChanged(address indexed previousDestination, address indexed newDestination)
func (_MystToken *MystTokenFilterer) ParseFundsRecoveryDestinationChanged(log types.Log) (*MystTokenFundsRecoveryDestinationChanged, error) {
	event := new(MystTokenFundsRecoveryDestinationChanged)
	if err := _MystToken.contract.UnpackLog(event, "FundsRecoveryDestinationChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
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
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe.
//
// Solidity: event Minted(address indexed to, uint256 amount)
func (_MystToken *MystTokenFilterer) FilterMinted(opts *bind.FilterOpts, to []common.Address) (*MystTokenMintedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MystToken.contract.FilterLogs(opts, "Minted", toRule)
	if err != nil {
		return nil, err
	}
	return &MystTokenMintedIterator{contract: _MystToken.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe.
//
// Solidity: event Minted(address indexed to, uint256 amount)
func (_MystToken *MystTokenFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *MystTokenMinted, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MystToken.contract.WatchLogs(opts, "Minted", toRule)
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
// Solidity: event Minted(address indexed to, uint256 amount)
func (_MystToken *MystTokenFilterer) ParseMinted(log types.Log) (*MystTokenMinted, error) {
	event := new(MystTokenMinted)
	if err := _MystToken.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
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
	event.Raw = log
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
	Agent common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpgrade is a free log retrieval operation binding the contract event 0x7e5c344a8141a805725cb476f76c6953b842222b967edd1f78ddb6e8b3f397ac.
//
// Solidity: event Upgrade(address indexed from, address agent, uint256 _value)
func (_MystToken *MystTokenFilterer) FilterUpgrade(opts *bind.FilterOpts, from []common.Address) (*MystTokenUpgradeIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _MystToken.contract.FilterLogs(opts, "Upgrade", fromRule)
	if err != nil {
		return nil, err
	}
	return &MystTokenUpgradeIterator{contract: _MystToken.contract, event: "Upgrade", logs: logs, sub: sub}, nil
}

// WatchUpgrade is a free log subscription operation binding the contract event 0x7e5c344a8141a805725cb476f76c6953b842222b967edd1f78ddb6e8b3f397ac.
//
// Solidity: event Upgrade(address indexed from, address agent, uint256 _value)
func (_MystToken *MystTokenFilterer) WatchUpgrade(opts *bind.WatchOpts, sink chan<- *MystTokenUpgrade, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _MystToken.contract.WatchLogs(opts, "Upgrade", fromRule)
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
// Solidity: event Upgrade(address indexed from, address agent, uint256 _value)
func (_MystToken *MystTokenFilterer) ParseUpgrade(log types.Log) (*MystTokenUpgrade, error) {
	event := new(MystTokenUpgrade)
	if err := _MystToken.contract.UnpackLog(event, "Upgrade", log); err != nil {
		return nil, err
	}
	event.Raw = log
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
	event.Raw = log
	return event, nil
}

// MystTokenUpgradeMasterSetIterator is returned from FilterUpgradeMasterSet and is used to iterate over the raw logs and unpacked data for UpgradeMasterSet events raised by the MystToken contract.
type MystTokenUpgradeMasterSetIterator struct {
	Event *MystTokenUpgradeMasterSet // Event containing the contract specifics and raw log

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
func (it *MystTokenUpgradeMasterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MystTokenUpgradeMasterSet)
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
		it.Event = new(MystTokenUpgradeMasterSet)
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
func (it *MystTokenUpgradeMasterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MystTokenUpgradeMasterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MystTokenUpgradeMasterSet represents a UpgradeMasterSet event raised by the MystToken contract.
type MystTokenUpgradeMasterSet struct {
	Master common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpgradeMasterSet is a free log retrieval operation binding the contract event 0x0bae748e6d38d2b1532af619519837d91d74845ad693f6f229677b4ac20b2d50.
//
// Solidity: event UpgradeMasterSet(address master)
func (_MystToken *MystTokenFilterer) FilterUpgradeMasterSet(opts *bind.FilterOpts) (*MystTokenUpgradeMasterSetIterator, error) {

	logs, sub, err := _MystToken.contract.FilterLogs(opts, "UpgradeMasterSet")
	if err != nil {
		return nil, err
	}
	return &MystTokenUpgradeMasterSetIterator{contract: _MystToken.contract, event: "UpgradeMasterSet", logs: logs, sub: sub}, nil
}

// WatchUpgradeMasterSet is a free log subscription operation binding the contract event 0x0bae748e6d38d2b1532af619519837d91d74845ad693f6f229677b4ac20b2d50.
//
// Solidity: event UpgradeMasterSet(address master)
func (_MystToken *MystTokenFilterer) WatchUpgradeMasterSet(opts *bind.WatchOpts, sink chan<- *MystTokenUpgradeMasterSet) (event.Subscription, error) {

	logs, sub, err := _MystToken.contract.WatchLogs(opts, "UpgradeMasterSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MystTokenUpgradeMasterSet)
				if err := _MystToken.contract.UnpackLog(event, "UpgradeMasterSet", log); err != nil {
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

// ParseUpgradeMasterSet is a log parse operation binding the contract event 0x0bae748e6d38d2b1532af619519837d91d74845ad693f6f229677b4ac20b2d50.
//
// Solidity: event UpgradeMasterSet(address master)
func (_MystToken *MystTokenFilterer) ParseUpgradeMasterSet(log types.Log) (*MystTokenUpgradeMasterSet, error) {
	event := new(MystTokenUpgradeMasterSet)
	if err := _MystToken.contract.UnpackLog(event, "UpgradeMasterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
