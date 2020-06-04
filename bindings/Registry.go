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

// RegistryABI is the input ABI used to generate the binding from.
const RegistryABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dexAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_configAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_regFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minimalAccountantStake\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"identityHash\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"accountantId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"channelAddress\",\"type\":\"address\"}],\"name\":\"ConsumerChannelCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousDestination\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newDestination\",\"type\":\"address\"}],\"name\":\"DestinationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"accountantId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"accountantOperator\",\"type\":\"address\"}],\"name\":\"RegisteredAccountant\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"identityHash\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"accountantId\",\"type\":\"address\"}],\"name\":\"RegisteredIdentity\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"accountants\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"function()viewexternalreturns(uint256)\",\"name\":\"stake\",\"type\":\"function\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimEthers\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"claimTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"config\",\"outputs\":[{\"internalType\":\"contractConfig\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFundsDestination\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minimalAccountantStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registrationFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_newDestination\",\"type\":\"address\"}],\"name\":\"setFundsDestination\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_hermesId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_transactorFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"registerIdentity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_accountantOperator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_accountantFee\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_maxStake\",\"type\":\"uint256\"}],\"name\":\"registerAccountant\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identity\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_hermesId\",\"type\":\"address\"}],\"name\":\"getChannelAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_accountantOperator\",\"type\":\"address\"}],\"name\":\"getAccountantAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"}],\"name\":\"getProxyCode\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getChannelImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAccountantImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identityHash\",\"type\":\"address\"}],\"name\":\"isRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_accountantId\",\"type\":\"address\"}],\"name\":\"isAccountant\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newFee\",\"type\":\"uint256\"}],\"name\":\"changeRegistrationFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"transferCollectedFeeTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RegistryBin is the compiled bytecode used for deploying new contracts.
var RegistryBin = "0x608060405234801561001057600080fd5b50604051611ced380380611ced833981810160405260a081101561003357600080fd5b50805160208201516040808401516060850151608090950151600080546001600160a01b031916331780825593519596949592949391926001600160a01b0392909216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3600582905560068190556001600160a01b0385166100b957600080fd5b600280546001600160a01b0319166001600160a01b038781169190911790915584166100e457600080fd5b600380546001600160a01b0319166001600160a01b0386811691909117909155831661010f57600080fd5b5050600480546001600160a01b039092166001600160a01b03199092169190911790555050611baa806101436000396000f3fe6080604052600436106101665760003560e01c80638da5cb5b116100d1578063e32525371161008a578063f58c5b6e11610064578063f58c5b6e14610628578063f595cfd21461063d578063f707fb4a14610670578063fc0c546a146106e257610166565b8063e325253714610587578063e617aaac146105ba578063f2fde38b146105f557610166565b80638da5cb5b146103785780638f32d59b1461038d578063ab867213146103a2578063c3c5a5471461044a578063cf10c9691461047d578063df8de3e71461055457610166565b80636931b550116101235780636931b550146102fa5780636ef3b1971461030f578063715018a61461032457806379502c55146103395780637c671a211461034e578063824b09d61461036357610166565b806314c44e09146101b3578063238e130a146101da57806328a2276c1461020f5780632a33ddbd14610258578063500507691461029f578063692058c2146102c9575b6040805162461bcd60e51b815260206004820152601d60248201527f52656a656374696e672074782077697468206574686572732073656e74000000604482015290519081900360640190fd5b3480156101bf57600080fd5b506101c86106f7565b60408051918252519081900360200190f35b3480156101e657600080fd5b5061020d600480360360208110156101fd57600080fd5b50356001600160a01b03166106fd565b005b34801561021b57600080fd5b5061020d6004803603608081101561023257600080fd5b506001600160a01b038135169060208101359061ffff60408201351690606001356107b3565b34801561026457600080fd5b5061028b6004803603602081101561027b57600080fd5b50356001600160a01b0316610aaa565b604080519115158252519081900360200190f35b3480156102ab57600080fd5b5061020d600480360360208110156102c257600080fd5b5035610adc565b3480156102d557600080fd5b506102de610b28565b604080516001600160a01b039092168252519081900360200190f35b34801561030657600080fd5b5061020d610b37565b34801561031b57600080fd5b506102de610b88565b34801561033057600080fd5b5061020d610c28565b34801561034557600080fd5b506102de610cb9565b34801561035a57600080fd5b506102de610cc8565b34801561036f57600080fd5b506101c8610d37565b34801561038457600080fd5b506102de610d3d565b34801561039957600080fd5b5061028b610d4c565b3480156103ae57600080fd5b506103d5600480360360208110156103c557600080fd5b50356001600160a01b0316610d5d565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561040f5781810151838201526020016103f7565b50505050905090810190601f16801561043c5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561045657600080fd5b5061028b6004803603602081101561046d57600080fd5b50356001600160a01b0316610ddc565b34801561048957600080fd5b5061020d600480360360a08110156104a057600080fd5b6001600160a01b038235811692602081013592604082013592606083013516919081019060a0810160808201356401000000008111156104df57600080fd5b8201836020820111156104f157600080fd5b8035906020019184600183028401116401000000008311171561051357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610dfa945050505050565b34801561056057600080fd5b5061020d6004803603602081101561057757600080fd5b50356001600160a01b0316611366565b34801561059357600080fd5b5061020d600480360360208110156105aa57600080fd5b50356001600160a01b03166114c7565b3480156105c657600080fd5b506102de600480360360408110156105dd57600080fd5b506001600160a01b0381358116916020013516611619565b34801561060157600080fd5b5061020d6004803603602081101561061857600080fd5b50356001600160a01b0316611682565b34801561063457600080fd5b506102de6116d2565b34801561064957600080fd5b506102de6004803603602081101561066057600080fd5b50356001600160a01b03166116e1565b34801561067c57600080fd5b506106a36004803603602081101561069357600080fd5b50356001600160a01b0316611713565b60405180846001600160a01b03166001600160a01b03168152602001838363ffffffff169060201b1760401b8152602001935050505060405180910390f35b3480156106ee57600080fd5b506102de611748565b60055481565b610705610d4c565b610744576040805162461bcd60e51b81526020600482018190526024820152600080516020611ad0833981519152604482015290519081900360640190fd5b6001600160a01b03811661075757600080fd5b6001546040516001600160a01b038084169216907fe1a66d77649cf0a57b9937073549f30f1c82bb865aaf066d2f299e37a62c6aad90600090a3600180546001600160a01b0319166001600160a01b0392909216919091179055565b6001600160a01b03841661080e576040805162461bcd60e51b815260206004820152601e60248201527f6f70657261746f722063616e2774206265207a65726f20616464726573730000604482015290519081900360640190fd5b60065483101561084f5760405162461bcd60e51b8152600401808060200182810382526036815260200180611af06036913960400191505060405180910390fd5b600061085a856116e1565b905061086581610aaa565b156108b7576040805162461bcd60e51b815260206004820152601d60248201527f6163636f756e74616e7420616c72656164792072656769737465726564000000604482015290519081900360640190fd5b60006108db866001600160a01b03166108d66108d1610b88565b610d5d565b611757565b600254604080516323b872dd60e01b81523360048201526001600160a01b038085166024830152604482018a905291519394509116916323b872dd916064808201926020929091908290030181600087803b15801561093957600080fd5b505af115801561094d573d6000803e3d6000fd5b505050506040513d602081101561096357600080fd5b50506002546040805163fec8157d60e01b81526001600160a01b039283166004820152888316602482015261ffff871660448201526064810186905290519183169163fec8157d9160848082019260009290919082900301818387803b1580156109cc57600080fd5b505af11580156109e0573d6000803e3d6000fd5b50506040805180820182526001600160a01b03808b16808352640100000000600160c01b03602088811b821663fc0e3d9017861b8186019081528985166000818152600784528890209651875496166001600160a01b0319909616959095178655516001909501805495871c92831663ffffffff93909316929092176001600160c01b0319909516949094179055835190815292519094507fc9f77826eb4f5ea1c94b1b3ff214796f59d1c5f610af8a3ff52cd83e3eebf64e9350918290030190a2505050505050565b6001600160a01b0380821660009081526007602052604081205490911681610ad1826116e1565b3b1515949350505050565b610ae4610d4c565b610b23576040805162461bcd60e51b81526020600482018190526024820152600080516020611ad0833981519152604482015290519081900360640190fd5b600555565b6003546001600160a01b031681565b6001546001600160a01b0316610b4c57600080fd5b6001546040516001600160a01b03909116904780156108fc02916000818181858888f19350505050158015610b85573d6000803e3d6000fd5b50565b60048054604080516321f8a72160e01b81527f52948fa93a94851571e57fddc2be83c51e0a64bb5e9ca55f4f90439b9802b57593810193909352516000926001600160a01b03909216916321f8a721916024808301926020929190829003018186803b158015610bf757600080fd5b505afa158015610c0b573d6000803e3d6000fd5b505050506040513d6020811015610c2157600080fd5b5051905090565b610c30610d4c565b610c6f576040805162461bcd60e51b81526020600482018190526024820152600080516020611ad0833981519152604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6004546001600160a01b031681565b60048054604080516321f8a72160e01b81527f2ef7e7c50e1b6a574193d0d32b7c0456cf12390a0872cf00be4797e71c3756f793810193909352516000926001600160a01b03909216916321f8a721916024808301926020929190829003018186803b158015610bf757600080fd5b60065481565b6000546001600160a01b031690565b6000546001600160a01b0316331490565b606080604051806060016040528060378152602001611a99603791399050606083901b60005b60148160ff161015610dd357818160ff1660148110610d9e57fe5b1a60f81b838260140160ff1681518110610db457fe5b60200101906001600160f81b031916908160001a905350600101610d83565b50909392505050565b6001600160a01b031660009081526008602052604090205460ff1690565b610e0385611771565b610e54576040805162461bcd60e51b815260206004820152601e60248201527f70726f766964656420686173206861766520746f206265206163746976650000604482015290519081900360640190fd5b6040805130606090811b6020808401919091526001600160601b031989831b8116603485015260488401899052606884018890529186901b90911660888301528251607c818403018152609c9092019092528051910120600090610ebe908363ffffffff6117f116565b90506001600160a01b038116610f0d576040805162461bcd60e51b815260206004820152600f60248201526e77726f6e67207369676e617475726560881b604482015290519081900360640190fd5b6000610f3485610f28886005546118df90919063ffffffff16565b9063ffffffff6118df16565b6002549091506001600160a01b03166370a08231610f52848a611619565b6040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b158015610f9857600080fd5b505afa158015610fac573d6000803e3d6000fd5b505050506040513d6020811015610fc257600080fd5b50518111156110025760405162461bcd60e51b815260040180806020018281038252602a815260200180611b26602a913960400191505060405180910390fd5b604080516001600160601b0319606085811b82166020808501919091528b821b909216603484015283516028818503018152604890930190935281519101209061104d6108d1610cc8565b9050600061105b8383611757565b60025460035460408051637b809f7b60e11b81526001600160a01b039384166004820152918316602483015288831660448301528d8316606483015260848201889052519293509083169163f7013ef69160a48082019260009290919082900301818387803b1580156110cd57600080fd5b505af11580156110e1573d6000803e3d6000fd5b505050506000891180156110fd57506001600160a01b03871615155b15611232576002546040805163095ea7b360e01b81526001600160a01b038d81166004830152602482018d90529151919092169163095ea7b39160448083019260209291908290030181600087803b15801561115857600080fd5b505af115801561116c573d6000803e3d6000fd5b505050506040513d602081101561118257600080fd5b50516111bf5760405162461bcd60e51b815260040180806020018281038252602d815260200180611a47602d913960400191505060405180910390fd5b6040805163029e63c960e21b81526001600160a01b0387811660048301528981166024830152604482018c90529151918c1691630a798f249160648082019260009290919082900301818387803b15801561121957600080fd5b505af115801561122d573d6000803e3d6000fd5b505050505b87156112b9576002546040805163a9059cbb60e01b8152336004820152602481018b905290516001600160a01b039092169163a9059cbb916044808201926020929091908290030181600087803b15801561128c57600080fd5b505af11580156112a0573d6000803e3d6000fd5b505050506040513d60208110156112b657600080fd5b50505b604080516001600160a01b0383811682529151828d16928816917f2ed7bcf2ff03098102c7003d7ce2a633e4b49b8198b07de5383cdf4c0ab9228b919081900360200190a361130785610ddc565b61135a576001600160a01b03808616600081815260086020526040808220805460ff1916600117905551928d16927fefaf768237c22e140a862d5d375ad5c153479fac3f8bcf8b580a1651fd62c3ef9190a35b50505050505050505050565b6001546001600160a01b031661137b57600080fd5b6002546001600160a01b03828116911614156113c85760405162461bcd60e51b8152600401808060200182810382526025815260200180611a746025913960400191505060405180910390fd5b604080516370a0823160e01b815230600482015290516000916001600160a01b038416916370a0823191602480820192602092909190829003018186803b15801561141257600080fd5b505afa158015611426573d6000803e3d6000fd5b505050506040513d602081101561143c57600080fd5b50516001546040805163a9059cbb60e01b81526001600160a01b0392831660048201526024810184905290519293509084169163a9059cbb916044808201926020929091908290030181600087803b15801561149757600080fd5b505af11580156114ab573d6000803e3d6000fd5b505050506040513d60208110156114c157600080fd5b50505050565b6114cf610d4c565b61150e576040805162461bcd60e51b81526020600482018190526024820152600080516020611ad0833981519152604482015290519081900360640190fd5b600254604080516370a0823160e01b815230600482015290516000926001600160a01b0316916370a08231916024808301926020929190829003018186803b15801561155957600080fd5b505afa15801561156d573d6000803e3d6000fd5b505050506040513d602081101561158357600080fd5b50519050806115c35760405162461bcd60e51b8152600401808060200182810382526026815260200180611b506026913960400191505060405180910390fd5b6002546040805163a9059cbb60e01b81526001600160a01b038581166004830152602482018590529151919092169163a9059cbb9160448083019260209291908290030181600087803b15801561149757600080fd5b6000806116276108d1610cc8565b8051602091820120604080516001600160601b0319606089811b82168387015288901b16603482015281516028818303018152604890910190915280519201919091209091506116778183611939565b925050505b92915050565b61168a610d4c565b6116c9576040805162461bcd60e51b81526020600482018190526024820152600080516020611ad0833981519152604482015290519081900360640190fd5b610b8581611980565b6001546001600160a01b031690565b6000806116ef6108d1610b88565b8051602090910120905061170c6001600160a01b03841682611939565b9392505050565b600760209081526000918252604090912080546001909101546001600160a01b039182169281901c9091169063ffffffff1683565b6002546001600160a01b031681565b600080838351602085016000f59050803b61170c57600080fd5b600080826001600160a01b0316634e69d5606040518163ffffffff1660e01b815260040160206040518083038186803b1580156117ad57600080fd5b505afa1580156117c1573d6000803e3d6000fd5b505050506040513d60208110156117d757600080fd5b5051905060008160038111156117e957fe5b149392505050565b600081516041146118045750600061167c565b60208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a082111561184a576000935050505061167c565b8060ff16601b1415801561186257508060ff16601c14155b15611873576000935050505061167c565b6040805160008152602080820180845289905260ff8416828401526060820186905260808201859052915160019260a0808401939192601f1981019281900390910190855afa1580156118ca573d6000803e3d6000fd5b5050604051601f190151979650505050505050565b60008282018381101561170c576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b604080516001600160f81b03196020808301919091523060601b60218301526035820194909452605580820193909352815180820390930183526075019052805191012090565b6001600160a01b0381166119c55760405162461bcd60e51b8152600401808060200182810382526026815260200180611a216026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b039290921691909117905556fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573736865726d65732073686f756c642067657420617070726f76616c20746f207472616e7366657220746f6b656e736e617469766520746f6b656e2066756e64732063616e2774206265207265636f76657265643d602d80600a3d3981f3363d3d373d3d3d363d73bebebebebebebebebebebebebebebebebebebebe5af43d82803e903d91602b57fd5bf34f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726163636f756e74616e74206861766520746f207374616b65206174206c65617374206d696e696d616c207374616b6520616d6f756e746e6f7420656e6f756768742066756e647320696e206368616e6e656c20746f20636f7665722066656573636f6c6c6563746564206665652063616e6e6f74206265206c657373207468616e207a65726fa265627a7a72315820fd19007a68bcee723566aac5f2ffd2533ea75375652b4c6e1a5ce622095bf08464736f6c63430005110032"

// DeployRegistry deploys a new Ethereum contract, binding an instance of Registry to it.
func DeployRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _tokenAddress common.Address, _dexAddress common.Address, _configAddress common.Address, _regFee *big.Int, _minimalAccountantStake *big.Int) (common.Address, *types.Transaction, *Registry, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RegistryBin), backend, _tokenAddress, _dexAddress, _configAddress, _regFee, _minimalAccountantStake)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// Registry is an auto generated Go binding around an Ethereum contract.
type Registry struct {
	RegistryCaller     // Read-only binding to the contract
	RegistryTransactor // Write-only binding to the contract
	RegistryFilterer   // Log filterer for contract events
}

// RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrySession struct {
	Contract     *Registry         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistryCallerSession struct {
	Contract *RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistryTransactorSession struct {
	Contract     *RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistryRaw struct {
	Contract *Registry // Generic contract binding to access the raw methods on
}

// RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistryCallerRaw struct {
	Contract *RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistryTransactorRaw struct {
	Contract *RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistry creates a new instance of Registry, bound to a specific deployed contract.
func NewRegistry(address common.Address, backend bind.ContractBackend) (*Registry, error) {
	contract, err := bindRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// NewRegistryCaller creates a new read-only instance of Registry, bound to a specific deployed contract.
func NewRegistryCaller(address common.Address, caller bind.ContractCaller) (*RegistryCaller, error) {
	contract, err := bindRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryCaller{contract: contract}, nil
}

// NewRegistryTransactor creates a new write-only instance of Registry, bound to a specific deployed contract.
func NewRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryTransactor, error) {
	contract, err := bindRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryTransactor{contract: contract}, nil
}

// NewRegistryFilterer creates a new log filterer instance of Registry, bound to a specific deployed contract.
func NewRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryFilterer, error) {
	contract, err := bindRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryFilterer{contract: contract}, nil
}

// bindRegistry binds a generic wrapper to an already deployed contract.
func bindRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transact(opts, method, params...)
}

// Accountants is a free data retrieval call binding the contract method 0xf707fb4a.
//
// Solidity: function accountants(address ) constant returns(address operator, function stake)
func (_Registry *RegistryCaller) Accountants(opts *bind.CallOpts, arg0 common.Address) (struct {
	Operator common.Address
	Stake    [24]byte
}, error) {
	ret := new(struct {
		Operator common.Address
		Stake    [24]byte
	})
	out := ret
	err := _Registry.contract.Call(opts, out, "accountants", arg0)
	return *ret, err
}

// Accountants is a free data retrieval call binding the contract method 0xf707fb4a.
//
// Solidity: function accountants(address ) constant returns(address operator, function stake)
func (_Registry *RegistrySession) Accountants(arg0 common.Address) (struct {
	Operator common.Address
	Stake    [24]byte
}, error) {
	return _Registry.Contract.Accountants(&_Registry.CallOpts, arg0)
}

// Accountants is a free data retrieval call binding the contract method 0xf707fb4a.
//
// Solidity: function accountants(address ) constant returns(address operator, function stake)
func (_Registry *RegistryCallerSession) Accountants(arg0 common.Address) (struct {
	Operator common.Address
	Stake    [24]byte
}, error) {
	return _Registry.Contract.Accountants(&_Registry.CallOpts, arg0)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() constant returns(address)
func (_Registry *RegistryCaller) Config(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "config")
	return *ret0, err
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() constant returns(address)
func (_Registry *RegistrySession) Config() (common.Address, error) {
	return _Registry.Contract.Config(&_Registry.CallOpts)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() constant returns(address)
func (_Registry *RegistryCallerSession) Config() (common.Address, error) {
	return _Registry.Contract.Config(&_Registry.CallOpts)
}

// Dex is a free data retrieval call binding the contract method 0x692058c2.
//
// Solidity: function dex() constant returns(address)
func (_Registry *RegistryCaller) Dex(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "dex")
	return *ret0, err
}

// Dex is a free data retrieval call binding the contract method 0x692058c2.
//
// Solidity: function dex() constant returns(address)
func (_Registry *RegistrySession) Dex() (common.Address, error) {
	return _Registry.Contract.Dex(&_Registry.CallOpts)
}

// Dex is a free data retrieval call binding the contract method 0x692058c2.
//
// Solidity: function dex() constant returns(address)
func (_Registry *RegistryCallerSession) Dex() (common.Address, error) {
	return _Registry.Contract.Dex(&_Registry.CallOpts)
}

// GetAccountantAddress is a free data retrieval call binding the contract method 0xf595cfd2.
//
// Solidity: function getAccountantAddress(address _accountantOperator) constant returns(address)
func (_Registry *RegistryCaller) GetAccountantAddress(opts *bind.CallOpts, _accountantOperator common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "getAccountantAddress", _accountantOperator)
	return *ret0, err
}

// GetAccountantAddress is a free data retrieval call binding the contract method 0xf595cfd2.
//
// Solidity: function getAccountantAddress(address _accountantOperator) constant returns(address)
func (_Registry *RegistrySession) GetAccountantAddress(_accountantOperator common.Address) (common.Address, error) {
	return _Registry.Contract.GetAccountantAddress(&_Registry.CallOpts, _accountantOperator)
}

// GetAccountantAddress is a free data retrieval call binding the contract method 0xf595cfd2.
//
// Solidity: function getAccountantAddress(address _accountantOperator) constant returns(address)
func (_Registry *RegistryCallerSession) GetAccountantAddress(_accountantOperator common.Address) (common.Address, error) {
	return _Registry.Contract.GetAccountantAddress(&_Registry.CallOpts, _accountantOperator)
}

// GetAccountantImplementation is a free data retrieval call binding the contract method 0x6ef3b197.
//
// Solidity: function getAccountantImplementation() constant returns(address)
func (_Registry *RegistryCaller) GetAccountantImplementation(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "getAccountantImplementation")
	return *ret0, err
}

// GetAccountantImplementation is a free data retrieval call binding the contract method 0x6ef3b197.
//
// Solidity: function getAccountantImplementation() constant returns(address)
func (_Registry *RegistrySession) GetAccountantImplementation() (common.Address, error) {
	return _Registry.Contract.GetAccountantImplementation(&_Registry.CallOpts)
}

// GetAccountantImplementation is a free data retrieval call binding the contract method 0x6ef3b197.
//
// Solidity: function getAccountantImplementation() constant returns(address)
func (_Registry *RegistryCallerSession) GetAccountantImplementation() (common.Address, error) {
	return _Registry.Contract.GetAccountantImplementation(&_Registry.CallOpts)
}

// GetChannelAddress is a free data retrieval call binding the contract method 0xe617aaac.
//
// Solidity: function getChannelAddress(address _identity, address _hermesId) constant returns(address)
func (_Registry *RegistryCaller) GetChannelAddress(opts *bind.CallOpts, _identity common.Address, _hermesId common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "getChannelAddress", _identity, _hermesId)
	return *ret0, err
}

// GetChannelAddress is a free data retrieval call binding the contract method 0xe617aaac.
//
// Solidity: function getChannelAddress(address _identity, address _hermesId) constant returns(address)
func (_Registry *RegistrySession) GetChannelAddress(_identity common.Address, _hermesId common.Address) (common.Address, error) {
	return _Registry.Contract.GetChannelAddress(&_Registry.CallOpts, _identity, _hermesId)
}

// GetChannelAddress is a free data retrieval call binding the contract method 0xe617aaac.
//
// Solidity: function getChannelAddress(address _identity, address _hermesId) constant returns(address)
func (_Registry *RegistryCallerSession) GetChannelAddress(_identity common.Address, _hermesId common.Address) (common.Address, error) {
	return _Registry.Contract.GetChannelAddress(&_Registry.CallOpts, _identity, _hermesId)
}

// GetChannelImplementation is a free data retrieval call binding the contract method 0x7c671a21.
//
// Solidity: function getChannelImplementation() constant returns(address)
func (_Registry *RegistryCaller) GetChannelImplementation(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "getChannelImplementation")
	return *ret0, err
}

// GetChannelImplementation is a free data retrieval call binding the contract method 0x7c671a21.
//
// Solidity: function getChannelImplementation() constant returns(address)
func (_Registry *RegistrySession) GetChannelImplementation() (common.Address, error) {
	return _Registry.Contract.GetChannelImplementation(&_Registry.CallOpts)
}

// GetChannelImplementation is a free data retrieval call binding the contract method 0x7c671a21.
//
// Solidity: function getChannelImplementation() constant returns(address)
func (_Registry *RegistryCallerSession) GetChannelImplementation() (common.Address, error) {
	return _Registry.Contract.GetChannelImplementation(&_Registry.CallOpts)
}

// GetFundsDestination is a free data retrieval call binding the contract method 0xf58c5b6e.
//
// Solidity: function getFundsDestination() constant returns(address)
func (_Registry *RegistryCaller) GetFundsDestination(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "getFundsDestination")
	return *ret0, err
}

// GetFundsDestination is a free data retrieval call binding the contract method 0xf58c5b6e.
//
// Solidity: function getFundsDestination() constant returns(address)
func (_Registry *RegistrySession) GetFundsDestination() (common.Address, error) {
	return _Registry.Contract.GetFundsDestination(&_Registry.CallOpts)
}

// GetFundsDestination is a free data retrieval call binding the contract method 0xf58c5b6e.
//
// Solidity: function getFundsDestination() constant returns(address)
func (_Registry *RegistryCallerSession) GetFundsDestination() (common.Address, error) {
	return _Registry.Contract.GetFundsDestination(&_Registry.CallOpts)
}

// GetProxyCode is a free data retrieval call binding the contract method 0xab867213.
//
// Solidity: function getProxyCode(address _implementation) constant returns(bytes)
func (_Registry *RegistryCaller) GetProxyCode(opts *bind.CallOpts, _implementation common.Address) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "getProxyCode", _implementation)
	return *ret0, err
}

// GetProxyCode is a free data retrieval call binding the contract method 0xab867213.
//
// Solidity: function getProxyCode(address _implementation) constant returns(bytes)
func (_Registry *RegistrySession) GetProxyCode(_implementation common.Address) ([]byte, error) {
	return _Registry.Contract.GetProxyCode(&_Registry.CallOpts, _implementation)
}

// GetProxyCode is a free data retrieval call binding the contract method 0xab867213.
//
// Solidity: function getProxyCode(address _implementation) constant returns(bytes)
func (_Registry *RegistryCallerSession) GetProxyCode(_implementation common.Address) ([]byte, error) {
	return _Registry.Contract.GetProxyCode(&_Registry.CallOpts, _implementation)
}

// IsAccountant is a free data retrieval call binding the contract method 0x2a33ddbd.
//
// Solidity: function isAccountant(address _accountantId) constant returns(bool)
func (_Registry *RegistryCaller) IsAccountant(opts *bind.CallOpts, _accountantId common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "isAccountant", _accountantId)
	return *ret0, err
}

// IsAccountant is a free data retrieval call binding the contract method 0x2a33ddbd.
//
// Solidity: function isAccountant(address _accountantId) constant returns(bool)
func (_Registry *RegistrySession) IsAccountant(_accountantId common.Address) (bool, error) {
	return _Registry.Contract.IsAccountant(&_Registry.CallOpts, _accountantId)
}

// IsAccountant is a free data retrieval call binding the contract method 0x2a33ddbd.
//
// Solidity: function isAccountant(address _accountantId) constant returns(bool)
func (_Registry *RegistryCallerSession) IsAccountant(_accountantId common.Address) (bool, error) {
	return _Registry.Contract.IsAccountant(&_Registry.CallOpts, _accountantId)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Registry *RegistryCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Registry *RegistrySession) IsOwner() (bool, error) {
	return _Registry.Contract.IsOwner(&_Registry.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Registry *RegistryCallerSession) IsOwner() (bool, error) {
	return _Registry.Contract.IsOwner(&_Registry.CallOpts)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address _identityHash) constant returns(bool)
func (_Registry *RegistryCaller) IsRegistered(opts *bind.CallOpts, _identityHash common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "isRegistered", _identityHash)
	return *ret0, err
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address _identityHash) constant returns(bool)
func (_Registry *RegistrySession) IsRegistered(_identityHash common.Address) (bool, error) {
	return _Registry.Contract.IsRegistered(&_Registry.CallOpts, _identityHash)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address _identityHash) constant returns(bool)
func (_Registry *RegistryCallerSession) IsRegistered(_identityHash common.Address) (bool, error) {
	return _Registry.Contract.IsRegistered(&_Registry.CallOpts, _identityHash)
}

// MinimalAccountantStake is a free data retrieval call binding the contract method 0x824b09d6.
//
// Solidity: function minimalAccountantStake() constant returns(uint256)
func (_Registry *RegistryCaller) MinimalAccountantStake(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "minimalAccountantStake")
	return *ret0, err
}

// MinimalAccountantStake is a free data retrieval call binding the contract method 0x824b09d6.
//
// Solidity: function minimalAccountantStake() constant returns(uint256)
func (_Registry *RegistrySession) MinimalAccountantStake() (*big.Int, error) {
	return _Registry.Contract.MinimalAccountantStake(&_Registry.CallOpts)
}

// MinimalAccountantStake is a free data retrieval call binding the contract method 0x824b09d6.
//
// Solidity: function minimalAccountantStake() constant returns(uint256)
func (_Registry *RegistryCallerSession) MinimalAccountantStake() (*big.Int, error) {
	return _Registry.Contract.MinimalAccountantStake(&_Registry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Registry *RegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Registry *RegistrySession) Owner() (common.Address, error) {
	return _Registry.Contract.Owner(&_Registry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Registry *RegistryCallerSession) Owner() (common.Address, error) {
	return _Registry.Contract.Owner(&_Registry.CallOpts)
}

// RegistrationFee is a free data retrieval call binding the contract method 0x14c44e09.
//
// Solidity: function registrationFee() constant returns(uint256)
func (_Registry *RegistryCaller) RegistrationFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "registrationFee")
	return *ret0, err
}

// RegistrationFee is a free data retrieval call binding the contract method 0x14c44e09.
//
// Solidity: function registrationFee() constant returns(uint256)
func (_Registry *RegistrySession) RegistrationFee() (*big.Int, error) {
	return _Registry.Contract.RegistrationFee(&_Registry.CallOpts)
}

// RegistrationFee is a free data retrieval call binding the contract method 0x14c44e09.
//
// Solidity: function registrationFee() constant returns(uint256)
func (_Registry *RegistryCallerSession) RegistrationFee() (*big.Int, error) {
	return _Registry.Contract.RegistrationFee(&_Registry.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_Registry *RegistryCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_Registry *RegistrySession) Token() (common.Address, error) {
	return _Registry.Contract.Token(&_Registry.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_Registry *RegistryCallerSession) Token() (common.Address, error) {
	return _Registry.Contract.Token(&_Registry.CallOpts)
}

// ChangeRegistrationFee is a paid mutator transaction binding the contract method 0x50050769.
//
// Solidity: function changeRegistrationFee(uint256 _newFee) returns()
func (_Registry *RegistryTransactor) ChangeRegistrationFee(opts *bind.TransactOpts, _newFee *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "changeRegistrationFee", _newFee)
}

// ChangeRegistrationFee is a paid mutator transaction binding the contract method 0x50050769.
//
// Solidity: function changeRegistrationFee(uint256 _newFee) returns()
func (_Registry *RegistrySession) ChangeRegistrationFee(_newFee *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.ChangeRegistrationFee(&_Registry.TransactOpts, _newFee)
}

// ChangeRegistrationFee is a paid mutator transaction binding the contract method 0x50050769.
//
// Solidity: function changeRegistrationFee(uint256 _newFee) returns()
func (_Registry *RegistryTransactorSession) ChangeRegistrationFee(_newFee *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.ChangeRegistrationFee(&_Registry.TransactOpts, _newFee)
}

// ClaimEthers is a paid mutator transaction binding the contract method 0x6931b550.
//
// Solidity: function claimEthers() returns()
func (_Registry *RegistryTransactor) ClaimEthers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "claimEthers")
}

// ClaimEthers is a paid mutator transaction binding the contract method 0x6931b550.
//
// Solidity: function claimEthers() returns()
func (_Registry *RegistrySession) ClaimEthers() (*types.Transaction, error) {
	return _Registry.Contract.ClaimEthers(&_Registry.TransactOpts)
}

// ClaimEthers is a paid mutator transaction binding the contract method 0x6931b550.
//
// Solidity: function claimEthers() returns()
func (_Registry *RegistryTransactorSession) ClaimEthers() (*types.Transaction, error) {
	return _Registry.Contract.ClaimEthers(&_Registry.TransactOpts)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0xdf8de3e7.
//
// Solidity: function claimTokens(address _token) returns()
func (_Registry *RegistryTransactor) ClaimTokens(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "claimTokens", _token)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0xdf8de3e7.
//
// Solidity: function claimTokens(address _token) returns()
func (_Registry *RegistrySession) ClaimTokens(_token common.Address) (*types.Transaction, error) {
	return _Registry.Contract.ClaimTokens(&_Registry.TransactOpts, _token)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0xdf8de3e7.
//
// Solidity: function claimTokens(address _token) returns()
func (_Registry *RegistryTransactorSession) ClaimTokens(_token common.Address) (*types.Transaction, error) {
	return _Registry.Contract.ClaimTokens(&_Registry.TransactOpts, _token)
}

// RegisterAccountant is a paid mutator transaction binding the contract method 0x28a2276c.
//
// Solidity: function registerAccountant(address _accountantOperator, uint256 _stakeAmount, uint16 _accountantFee, uint256 _maxStake) returns()
func (_Registry *RegistryTransactor) RegisterAccountant(opts *bind.TransactOpts, _accountantOperator common.Address, _stakeAmount *big.Int, _accountantFee uint16, _maxStake *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "registerAccountant", _accountantOperator, _stakeAmount, _accountantFee, _maxStake)
}

// RegisterAccountant is a paid mutator transaction binding the contract method 0x28a2276c.
//
// Solidity: function registerAccountant(address _accountantOperator, uint256 _stakeAmount, uint16 _accountantFee, uint256 _maxStake) returns()
func (_Registry *RegistrySession) RegisterAccountant(_accountantOperator common.Address, _stakeAmount *big.Int, _accountantFee uint16, _maxStake *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.RegisterAccountant(&_Registry.TransactOpts, _accountantOperator, _stakeAmount, _accountantFee, _maxStake)
}

// RegisterAccountant is a paid mutator transaction binding the contract method 0x28a2276c.
//
// Solidity: function registerAccountant(address _accountantOperator, uint256 _stakeAmount, uint16 _accountantFee, uint256 _maxStake) returns()
func (_Registry *RegistryTransactorSession) RegisterAccountant(_accountantOperator common.Address, _stakeAmount *big.Int, _accountantFee uint16, _maxStake *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.RegisterAccountant(&_Registry.TransactOpts, _accountantOperator, _stakeAmount, _accountantFee, _maxStake)
}

// RegisterIdentity is a paid mutator transaction binding the contract method 0xcf10c969.
//
// Solidity: function registerIdentity(address _hermesId, uint256 _stakeAmount, uint256 _transactorFee, address _beneficiary, bytes _signature) returns()
func (_Registry *RegistryTransactor) RegisterIdentity(opts *bind.TransactOpts, _hermesId common.Address, _stakeAmount *big.Int, _transactorFee *big.Int, _beneficiary common.Address, _signature []byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "registerIdentity", _hermesId, _stakeAmount, _transactorFee, _beneficiary, _signature)
}

// RegisterIdentity is a paid mutator transaction binding the contract method 0xcf10c969.
//
// Solidity: function registerIdentity(address _hermesId, uint256 _stakeAmount, uint256 _transactorFee, address _beneficiary, bytes _signature) returns()
func (_Registry *RegistrySession) RegisterIdentity(_hermesId common.Address, _stakeAmount *big.Int, _transactorFee *big.Int, _beneficiary common.Address, _signature []byte) (*types.Transaction, error) {
	return _Registry.Contract.RegisterIdentity(&_Registry.TransactOpts, _hermesId, _stakeAmount, _transactorFee, _beneficiary, _signature)
}

// RegisterIdentity is a paid mutator transaction binding the contract method 0xcf10c969.
//
// Solidity: function registerIdentity(address _hermesId, uint256 _stakeAmount, uint256 _transactorFee, address _beneficiary, bytes _signature) returns()
func (_Registry *RegistryTransactorSession) RegisterIdentity(_hermesId common.Address, _stakeAmount *big.Int, _transactorFee *big.Int, _beneficiary common.Address, _signature []byte) (*types.Transaction, error) {
	return _Registry.Contract.RegisterIdentity(&_Registry.TransactOpts, _hermesId, _stakeAmount, _transactorFee, _beneficiary, _signature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registry *RegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registry *RegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _Registry.Contract.RenounceOwnership(&_Registry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registry *RegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Registry.Contract.RenounceOwnership(&_Registry.TransactOpts)
}

// SetFundsDestination is a paid mutator transaction binding the contract method 0x238e130a.
//
// Solidity: function setFundsDestination(address _newDestination) returns()
func (_Registry *RegistryTransactor) SetFundsDestination(opts *bind.TransactOpts, _newDestination common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setFundsDestination", _newDestination)
}

// SetFundsDestination is a paid mutator transaction binding the contract method 0x238e130a.
//
// Solidity: function setFundsDestination(address _newDestination) returns()
func (_Registry *RegistrySession) SetFundsDestination(_newDestination common.Address) (*types.Transaction, error) {
	return _Registry.Contract.SetFundsDestination(&_Registry.TransactOpts, _newDestination)
}

// SetFundsDestination is a paid mutator transaction binding the contract method 0x238e130a.
//
// Solidity: function setFundsDestination(address _newDestination) returns()
func (_Registry *RegistryTransactorSession) SetFundsDestination(_newDestination common.Address) (*types.Transaction, error) {
	return _Registry.Contract.SetFundsDestination(&_Registry.TransactOpts, _newDestination)
}

// TransferCollectedFeeTo is a paid mutator transaction binding the contract method 0xe3252537.
//
// Solidity: function transferCollectedFeeTo(address _beneficiary) returns()
func (_Registry *RegistryTransactor) TransferCollectedFeeTo(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "transferCollectedFeeTo", _beneficiary)
}

// TransferCollectedFeeTo is a paid mutator transaction binding the contract method 0xe3252537.
//
// Solidity: function transferCollectedFeeTo(address _beneficiary) returns()
func (_Registry *RegistrySession) TransferCollectedFeeTo(_beneficiary common.Address) (*types.Transaction, error) {
	return _Registry.Contract.TransferCollectedFeeTo(&_Registry.TransactOpts, _beneficiary)
}

// TransferCollectedFeeTo is a paid mutator transaction binding the contract method 0xe3252537.
//
// Solidity: function transferCollectedFeeTo(address _beneficiary) returns()
func (_Registry *RegistryTransactorSession) TransferCollectedFeeTo(_beneficiary common.Address) (*types.Transaction, error) {
	return _Registry.Contract.TransferCollectedFeeTo(&_Registry.TransactOpts, _beneficiary)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registry *RegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registry *RegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Registry.Contract.TransferOwnership(&_Registry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registry *RegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Registry.Contract.TransferOwnership(&_Registry.TransactOpts, newOwner)
}

// RegistryConsumerChannelCreatedIterator is returned from FilterConsumerChannelCreated and is used to iterate over the raw logs and unpacked data for ConsumerChannelCreated events raised by the Registry contract.
type RegistryConsumerChannelCreatedIterator struct {
	Event *RegistryConsumerChannelCreated // Event containing the contract specifics and raw log

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
func (it *RegistryConsumerChannelCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryConsumerChannelCreated)
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
		it.Event = new(RegistryConsumerChannelCreated)
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
func (it *RegistryConsumerChannelCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryConsumerChannelCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryConsumerChannelCreated represents a ConsumerChannelCreated event raised by the Registry contract.
type RegistryConsumerChannelCreated struct {
	IdentityHash   common.Address
	AccountantId   common.Address
	ChannelAddress common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterConsumerChannelCreated is a free log retrieval operation binding the contract event 0x2ed7bcf2ff03098102c7003d7ce2a633e4b49b8198b07de5383cdf4c0ab9228b.
//
// Solidity: event ConsumerChannelCreated(address indexed identityHash, address indexed accountantId, address channelAddress)
func (_Registry *RegistryFilterer) FilterConsumerChannelCreated(opts *bind.FilterOpts, identityHash []common.Address, accountantId []common.Address) (*RegistryConsumerChannelCreatedIterator, error) {

	var identityHashRule []interface{}
	for _, identityHashItem := range identityHash {
		identityHashRule = append(identityHashRule, identityHashItem)
	}
	var accountantIdRule []interface{}
	for _, accountantIdItem := range accountantId {
		accountantIdRule = append(accountantIdRule, accountantIdItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "ConsumerChannelCreated", identityHashRule, accountantIdRule)
	if err != nil {
		return nil, err
	}
	return &RegistryConsumerChannelCreatedIterator{contract: _Registry.contract, event: "ConsumerChannelCreated", logs: logs, sub: sub}, nil
}

// WatchConsumerChannelCreated is a free log subscription operation binding the contract event 0x2ed7bcf2ff03098102c7003d7ce2a633e4b49b8198b07de5383cdf4c0ab9228b.
//
// Solidity: event ConsumerChannelCreated(address indexed identityHash, address indexed accountantId, address channelAddress)
func (_Registry *RegistryFilterer) WatchConsumerChannelCreated(opts *bind.WatchOpts, sink chan<- *RegistryConsumerChannelCreated, identityHash []common.Address, accountantId []common.Address) (event.Subscription, error) {

	var identityHashRule []interface{}
	for _, identityHashItem := range identityHash {
		identityHashRule = append(identityHashRule, identityHashItem)
	}
	var accountantIdRule []interface{}
	for _, accountantIdItem := range accountantId {
		accountantIdRule = append(accountantIdRule, accountantIdItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "ConsumerChannelCreated", identityHashRule, accountantIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryConsumerChannelCreated)
				if err := _Registry.contract.UnpackLog(event, "ConsumerChannelCreated", log); err != nil {
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

// ParseConsumerChannelCreated is a log parse operation binding the contract event 0x2ed7bcf2ff03098102c7003d7ce2a633e4b49b8198b07de5383cdf4c0ab9228b.
//
// Solidity: event ConsumerChannelCreated(address indexed identityHash, address indexed accountantId, address channelAddress)
func (_Registry *RegistryFilterer) ParseConsumerChannelCreated(log types.Log) (*RegistryConsumerChannelCreated, error) {
	event := new(RegistryConsumerChannelCreated)
	if err := _Registry.contract.UnpackLog(event, "ConsumerChannelCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RegistryDestinationChangedIterator is returned from FilterDestinationChanged and is used to iterate over the raw logs and unpacked data for DestinationChanged events raised by the Registry contract.
type RegistryDestinationChangedIterator struct {
	Event *RegistryDestinationChanged // Event containing the contract specifics and raw log

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
func (it *RegistryDestinationChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryDestinationChanged)
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
		it.Event = new(RegistryDestinationChanged)
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
func (it *RegistryDestinationChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryDestinationChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryDestinationChanged represents a DestinationChanged event raised by the Registry contract.
type RegistryDestinationChanged struct {
	PreviousDestination common.Address
	NewDestination      common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDestinationChanged is a free log retrieval operation binding the contract event 0xe1a66d77649cf0a57b9937073549f30f1c82bb865aaf066d2f299e37a62c6aad.
//
// Solidity: event DestinationChanged(address indexed previousDestination, address indexed newDestination)
func (_Registry *RegistryFilterer) FilterDestinationChanged(opts *bind.FilterOpts, previousDestination []common.Address, newDestination []common.Address) (*RegistryDestinationChangedIterator, error) {

	var previousDestinationRule []interface{}
	for _, previousDestinationItem := range previousDestination {
		previousDestinationRule = append(previousDestinationRule, previousDestinationItem)
	}
	var newDestinationRule []interface{}
	for _, newDestinationItem := range newDestination {
		newDestinationRule = append(newDestinationRule, newDestinationItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "DestinationChanged", previousDestinationRule, newDestinationRule)
	if err != nil {
		return nil, err
	}
	return &RegistryDestinationChangedIterator{contract: _Registry.contract, event: "DestinationChanged", logs: logs, sub: sub}, nil
}

// WatchDestinationChanged is a free log subscription operation binding the contract event 0xe1a66d77649cf0a57b9937073549f30f1c82bb865aaf066d2f299e37a62c6aad.
//
// Solidity: event DestinationChanged(address indexed previousDestination, address indexed newDestination)
func (_Registry *RegistryFilterer) WatchDestinationChanged(opts *bind.WatchOpts, sink chan<- *RegistryDestinationChanged, previousDestination []common.Address, newDestination []common.Address) (event.Subscription, error) {

	var previousDestinationRule []interface{}
	for _, previousDestinationItem := range previousDestination {
		previousDestinationRule = append(previousDestinationRule, previousDestinationItem)
	}
	var newDestinationRule []interface{}
	for _, newDestinationItem := range newDestination {
		newDestinationRule = append(newDestinationRule, newDestinationItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "DestinationChanged", previousDestinationRule, newDestinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryDestinationChanged)
				if err := _Registry.contract.UnpackLog(event, "DestinationChanged", log); err != nil {
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

// ParseDestinationChanged is a log parse operation binding the contract event 0xe1a66d77649cf0a57b9937073549f30f1c82bb865aaf066d2f299e37a62c6aad.
//
// Solidity: event DestinationChanged(address indexed previousDestination, address indexed newDestination)
func (_Registry *RegistryFilterer) ParseDestinationChanged(log types.Log) (*RegistryDestinationChanged, error) {
	event := new(RegistryDestinationChanged)
	if err := _Registry.contract.UnpackLog(event, "DestinationChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Registry contract.
type RegistryOwnershipTransferredIterator struct {
	Event *RegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryOwnershipTransferred)
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
		it.Event = new(RegistryOwnershipTransferred)
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
func (it *RegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryOwnershipTransferred represents a OwnershipTransferred event raised by the Registry contract.
type RegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Registry *RegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RegistryOwnershipTransferredIterator{contract: _Registry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Registry *RegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryOwnershipTransferred)
				if err := _Registry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Registry *RegistryFilterer) ParseOwnershipTransferred(log types.Log) (*RegistryOwnershipTransferred, error) {
	event := new(RegistryOwnershipTransferred)
	if err := _Registry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RegistryRegisteredAccountantIterator is returned from FilterRegisteredAccountant and is used to iterate over the raw logs and unpacked data for RegisteredAccountant events raised by the Registry contract.
type RegistryRegisteredAccountantIterator struct {
	Event *RegistryRegisteredAccountant // Event containing the contract specifics and raw log

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
func (it *RegistryRegisteredAccountantIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryRegisteredAccountant)
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
		it.Event = new(RegistryRegisteredAccountant)
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
func (it *RegistryRegisteredAccountantIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryRegisteredAccountantIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryRegisteredAccountant represents a RegisteredAccountant event raised by the Registry contract.
type RegistryRegisteredAccountant struct {
	AccountantId       common.Address
	AccountantOperator common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRegisteredAccountant is a free log retrieval operation binding the contract event 0xc9f77826eb4f5ea1c94b1b3ff214796f59d1c5f610af8a3ff52cd83e3eebf64e.
//
// Solidity: event RegisteredAccountant(address indexed accountantId, address accountantOperator)
func (_Registry *RegistryFilterer) FilterRegisteredAccountant(opts *bind.FilterOpts, accountantId []common.Address) (*RegistryRegisteredAccountantIterator, error) {

	var accountantIdRule []interface{}
	for _, accountantIdItem := range accountantId {
		accountantIdRule = append(accountantIdRule, accountantIdItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "RegisteredAccountant", accountantIdRule)
	if err != nil {
		return nil, err
	}
	return &RegistryRegisteredAccountantIterator{contract: _Registry.contract, event: "RegisteredAccountant", logs: logs, sub: sub}, nil
}

// WatchRegisteredAccountant is a free log subscription operation binding the contract event 0xc9f77826eb4f5ea1c94b1b3ff214796f59d1c5f610af8a3ff52cd83e3eebf64e.
//
// Solidity: event RegisteredAccountant(address indexed accountantId, address accountantOperator)
func (_Registry *RegistryFilterer) WatchRegisteredAccountant(opts *bind.WatchOpts, sink chan<- *RegistryRegisteredAccountant, accountantId []common.Address) (event.Subscription, error) {

	var accountantIdRule []interface{}
	for _, accountantIdItem := range accountantId {
		accountantIdRule = append(accountantIdRule, accountantIdItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "RegisteredAccountant", accountantIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryRegisteredAccountant)
				if err := _Registry.contract.UnpackLog(event, "RegisteredAccountant", log); err != nil {
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

// ParseRegisteredAccountant is a log parse operation binding the contract event 0xc9f77826eb4f5ea1c94b1b3ff214796f59d1c5f610af8a3ff52cd83e3eebf64e.
//
// Solidity: event RegisteredAccountant(address indexed accountantId, address accountantOperator)
func (_Registry *RegistryFilterer) ParseRegisteredAccountant(log types.Log) (*RegistryRegisteredAccountant, error) {
	event := new(RegistryRegisteredAccountant)
	if err := _Registry.contract.UnpackLog(event, "RegisteredAccountant", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RegistryRegisteredIdentityIterator is returned from FilterRegisteredIdentity and is used to iterate over the raw logs and unpacked data for RegisteredIdentity events raised by the Registry contract.
type RegistryRegisteredIdentityIterator struct {
	Event *RegistryRegisteredIdentity // Event containing the contract specifics and raw log

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
func (it *RegistryRegisteredIdentityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryRegisteredIdentity)
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
		it.Event = new(RegistryRegisteredIdentity)
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
func (it *RegistryRegisteredIdentityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryRegisteredIdentityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryRegisteredIdentity represents a RegisteredIdentity event raised by the Registry contract.
type RegistryRegisteredIdentity struct {
	IdentityHash common.Address
	AccountantId common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRegisteredIdentity is a free log retrieval operation binding the contract event 0xefaf768237c22e140a862d5d375ad5c153479fac3f8bcf8b580a1651fd62c3ef.
//
// Solidity: event RegisteredIdentity(address indexed identityHash, address indexed accountantId)
func (_Registry *RegistryFilterer) FilterRegisteredIdentity(opts *bind.FilterOpts, identityHash []common.Address, accountantId []common.Address) (*RegistryRegisteredIdentityIterator, error) {

	var identityHashRule []interface{}
	for _, identityHashItem := range identityHash {
		identityHashRule = append(identityHashRule, identityHashItem)
	}
	var accountantIdRule []interface{}
	for _, accountantIdItem := range accountantId {
		accountantIdRule = append(accountantIdRule, accountantIdItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "RegisteredIdentity", identityHashRule, accountantIdRule)
	if err != nil {
		return nil, err
	}
	return &RegistryRegisteredIdentityIterator{contract: _Registry.contract, event: "RegisteredIdentity", logs: logs, sub: sub}, nil
}

// WatchRegisteredIdentity is a free log subscription operation binding the contract event 0xefaf768237c22e140a862d5d375ad5c153479fac3f8bcf8b580a1651fd62c3ef.
//
// Solidity: event RegisteredIdentity(address indexed identityHash, address indexed accountantId)
func (_Registry *RegistryFilterer) WatchRegisteredIdentity(opts *bind.WatchOpts, sink chan<- *RegistryRegisteredIdentity, identityHash []common.Address, accountantId []common.Address) (event.Subscription, error) {

	var identityHashRule []interface{}
	for _, identityHashItem := range identityHash {
		identityHashRule = append(identityHashRule, identityHashItem)
	}
	var accountantIdRule []interface{}
	for _, accountantIdItem := range accountantId {
		accountantIdRule = append(accountantIdRule, accountantIdItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "RegisteredIdentity", identityHashRule, accountantIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryRegisteredIdentity)
				if err := _Registry.contract.UnpackLog(event, "RegisteredIdentity", log); err != nil {
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

// ParseRegisteredIdentity is a log parse operation binding the contract event 0xefaf768237c22e140a862d5d375ad5c153479fac3f8bcf8b580a1651fd62c3ef.
//
// Solidity: event RegisteredIdentity(address indexed identityHash, address indexed accountantId)
func (_Registry *RegistryFilterer) ParseRegisteredIdentity(log types.Log) (*RegistryRegisteredIdentity, error) {
	event := new(RegistryRegisteredIdentity)
	if err := _Registry.contract.UnpackLog(event, "RegisteredIdentity", log); err != nil {
		return nil, err
	}
	return event, nil
}
