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

// RegistryABI is the input ABI used to generate the binding from.
const RegistryABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dexAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_regFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minimalHermesStake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_channelImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_hermesImplementation\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"identityHash\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"hermesId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"channelAddress\",\"type\":\"address\"}],\"name\":\"ConsumerChannelCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousDestination\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newDestination\",\"type\":\"address\"}],\"name\":\"DestinationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"hermesId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"hermesOperator\",\"type\":\"address\"}],\"name\":\"RegisteredHermes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"identityHash\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"hermesId\",\"type\":\"address\"}],\"name\":\"RegisteredIdentity\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"claimEthers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"claimTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFundsDestination\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"hermeses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"function()viewexternalreturns(uint256)\",\"name\":\"stake\",\"type\":\"function\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimalHermesStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registrationFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_newDestination\",\"type\":\"address\"}],\"name\":\"setFundsDestination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_hermesId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_transactorFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"registerIdentity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_hermesOperator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_hermesFee\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_maxStake\",\"type\":\"uint256\"}],\"name\":\"registerHermes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identity\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_hermesId\",\"type\":\"address\"}],\"name\":\"getChannelAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_hermesOperator\",\"type\":\"address\"}],\"name\":\"getHermesAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"}],\"name\":\"getProxyCode\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChannelImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getHermesImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identityHash\",\"type\":\"address\"}],\"name\":\"isRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_hermesId\",\"type\":\"address\"}],\"name\":\"isHermes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newFee\",\"type\":\"uint256\"}],\"name\":\"changeRegistrationFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"transferCollectedFeeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RegistryBin is the compiled bytecode used for deploying new contracts.
var RegistryBin = "0x608060405234801561001057600080fd5b50604051611cf8380380611cf8833981810160405260c081101561003357600080fd5b508051602082015160408301516060840151608085015160a0909501516004839055600582905593949293919290916001600160a01b03861661007557600080fd5b600280546001600160a01b0319166001600160a01b038881169190911790915585166100a057600080fd5b600380546001600160a01b039687166001600160a01b03199182161790915560068054938716938216939093179092556007805491909516911617909255505050611c08806100f06000396000f3fe6080604052600436106101445760003560e01c8063ab867213116100b6578063e32525371161006f578063e3252537146105a1578063e617aaac146105d4578063e64324fd1461060f578063f2fde38b14610658578063f58c5b6e1461068b578063fc0c546a146106a057610196565b8063ab86721314610342578063acc831d0146103ea578063c3c5a5471461041d578063cdd596e014610464578063cf10c96914610497578063df8de3e71461056e57610196565b8063692058c211610108578063692058c2146102a85780636931b550146102d9578063715018a6146102ee5780637c671a21146103035780638da5cb5b146103185780639936a87b1461032d57610196565b806303fb422f1461019b57806314c44e091461020d578063238e130a14610234578063500507691461026957806366cf58751461029357610196565b36610196576040805162461bcd60e51b815260206004820152601d60248201527f52656a656374696e672074782077697468206574686572732073656e74000000604482015290519081900360640190fd5b600080fd5b3480156101a757600080fd5b506101ce600480360360208110156101be57600080fd5b50356001600160a01b03166106b5565b60405180846001600160a01b03166001600160a01b03168152602001838363ffffffff169060201b1760401b8152602001935050505060405180910390f35b34801561021957600080fd5b506102226106ea565b60408051918252519081900360200190f35b34801561024057600080fd5b506102676004803603602081101561025757600080fd5b50356001600160a01b03166106f0565b005b34801561027557600080fd5b506102676004803603602081101561028c57600080fd5b50356107c0565b34801561029f57600080fd5b50610222610826565b3480156102b457600080fd5b506102bd61082c565b604080516001600160a01b039092168252519081900360200190f35b3480156102e557600080fd5b5061026761083b565b3480156102fa57600080fd5b5061026761088c565b34801561030f57600080fd5b506102bd610937565b34801561032457600080fd5b506102bd610946565b34801561033957600080fd5b506102bd610955565b34801561034e57600080fd5b506103756004803603602081101561036557600080fd5b50356001600160a01b0316610964565b6040805160208082528351818301528351919283929083019185019080838360005b838110156103af578181015183820152602001610397565b50505050905090810190601f1680156103dc5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156103f657600080fd5b506102bd6004803603602081101561040d57600080fd5b50356001600160a01b03166109e3565b34801561042957600080fd5b506104506004803603602081101561044057600080fd5b50356001600160a01b0316610a1a565b604080519115158252519081900360200190f35b34801561047057600080fd5b506104506004803603602081101561048757600080fd5b50356001600160a01b0316610a38565b3480156104a357600080fd5b50610267600480360360a08110156104ba57600080fd5b6001600160a01b038235811692602081013592604082013592606083013516919081019060a0810160808201356401000000008111156104f957600080fd5b82018360208201111561050b57600080fd5b8035906020019184600183028401116401000000008311171561052d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610a6a945050505050565b34801561057a57600080fd5b506102676004803603602081101561059157600080fd5b50356001600160a01b0316610fd6565b3480156105ad57600080fd5b50610267600480360360208110156105c457600080fd5b50356001600160a01b0316611137565b3480156105e057600080fd5b506102bd600480360360408110156105f757600080fd5b506001600160a01b03813581169160200135166112a3565b34801561061b57600080fd5b506102676004803603608081101561063257600080fd5b506001600160a01b038135169060208101359061ffff604082013516906060013561130a565b34801561066457600080fd5b506102676004803603602081101561067b57600080fd5b50356001600160a01b03166115fc565b34801561069757600080fd5b506102bd6116fd565b3480156106ac57600080fd5b506102bd61170c565b600860209081526000918252604090912080546001909101546001600160a01b039182169281901c9091169063ffffffff1683565b60045481565b6000546001600160a01b031633148061071257506000546001600160a01b0316155b610751576040805162461bcd60e51b81526020600482018190526024820152600080516020611b63833981519152604482015290519081900360640190fd5b6001600160a01b03811661076457600080fd5b6001546040516001600160a01b038084169216907fe1a66d77649cf0a57b9937073549f30f1c82bb865aaf066d2f299e37a62c6aad90600090a3600180546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b03163314806107e257506000546001600160a01b0316155b610821576040805162461bcd60e51b81526020600482018190526024820152600080516020611b63833981519152604482015290519081900360640190fd5b600455565b60055481565b6003546001600160a01b031681565b6001546001600160a01b031661085057600080fd5b6001546040516001600160a01b03909116904780156108fc02916000818181858888f19350505050158015610889573d6000803e3d6000fd5b50565b6000546001600160a01b03163314806108ae57506000546001600160a01b0316155b6108ed576040805162461bcd60e51b81526020600482018190526024820152600080516020611b63833981519152604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6006546001600160a01b031690565b6000546001600160a01b031690565b6007546001600160a01b031690565b606080604051806060016040528060378152602001611b0a603791399050606083901b60005b60148160ff1610156109da57818160ff16601481106109a557fe5b1a60f81b838260140160ff16815181106109bb57fe5b60200101906001600160f81b031916908160001a90535060010161098a565b50909392505050565b6000806109f66109f1610955565b610964565b80516020909101209050610a136001600160a01b0384168261171b565b9392505050565b6001600160a01b031660009081526009602052604090205460ff1690565b6001600160a01b0380821660009081526008602052604081205490911681610a5f826109e3565b3b1515949350505050565b610a7385611762565b610ac4576040805162461bcd60e51b815260206004820152601e60248201527f70726f766964656420686173206861766520746f206265206163746976650000604482015290519081900360640190fd5b6040805130606090811b6020808401919091526001600160601b031989831b8116603485015260488401899052606884018890529186901b90911660888301528251607c818403018152609c9092019092528051910120600090610b2e908363ffffffff6117e216565b90506001600160a01b038116610b7d576040805162461bcd60e51b815260206004820152600f60248201526e77726f6e67207369676e617475726560881b604482015290519081900360640190fd5b6000610ba485610b98886004546119c990919063ffffffff16565b9063ffffffff6119c916565b6002549091506001600160a01b03166370a08231610bc2848a6112a3565b6040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b158015610c0857600080fd5b505afa158015610c1c573d6000803e3d6000fd5b505050506040513d6020811015610c3257600080fd5b5051811115610c725760405162461bcd60e51b815260040180806020018281038252602a815260200180611b83602a913960400191505060405180910390fd5b604080516001600160601b0319606085811b82166020808501919091528b821b9092166034840152835160288185030181526048909301909352815191012090610cbd6109f1610937565b90506000610ccb8383611a23565b60025460035460408051637b809f7b60e11b81526001600160a01b039384166004820152918316602483015288831660448301528d8316606483015260848201889052519293509083169163f7013ef69160a48082019260009290919082900301818387803b158015610d3d57600080fd5b505af1158015610d51573d6000803e3d6000fd5b50505050600089118015610d6d57506001600160a01b03871615155b15610ea2576002546040805163095ea7b360e01b81526001600160a01b038d81166004830152602482018d90529151919092169163095ea7b39160448083019260209291908290030181600087803b158015610dc857600080fd5b505af1158015610ddc573d6000803e3d6000fd5b505050506040513d6020811015610df257600080fd5b5051610e2f5760405162461bcd60e51b815260040180806020018281038252602d815260200180611a64602d913960400191505060405180910390fd5b6040805163029e63c960e21b81526001600160a01b0387811660048301528981166024830152604482018c90529151918c1691630a798f249160648082019260009290919082900301818387803b158015610e8957600080fd5b505af1158015610e9d573d6000803e3d6000fd5b505050505b8715610f29576002546040805163a9059cbb60e01b8152336004820152602481018b905290516001600160a01b039092169163a9059cbb916044808201926020929091908290030181600087803b158015610efc57600080fd5b505af1158015610f10573d6000803e3d6000fd5b505050506040513d6020811015610f2657600080fd5b50505b604080516001600160a01b0383811682529151828d16928816917f2ed7bcf2ff03098102c7003d7ce2a633e4b49b8198b07de5383cdf4c0ab9228b919081900360200190a3610f7785610a1a565b610fca576001600160a01b03808616600081815260096020526040808220805460ff1916600117905551928d16927fefaf768237c22e140a862d5d375ad5c153479fac3f8bcf8b580a1651fd62c3ef9190a35b50505050505050505050565b6001546001600160a01b0316610feb57600080fd5b6002546001600160a01b03828116911614156110385760405162461bcd60e51b8152600401808060200182810382526025815260200180611a916025913960400191505060405180910390fd5b604080516370a0823160e01b815230600482015290516000916001600160a01b038416916370a0823191602480820192602092909190829003018186803b15801561108257600080fd5b505afa158015611096573d6000803e3d6000fd5b505050506040513d60208110156110ac57600080fd5b50516001546040805163a9059cbb60e01b81526001600160a01b0392831660048201526024810184905290519293509084169163a9059cbb916044808201926020929091908290030181600087803b15801561110757600080fd5b505af115801561111b573d6000803e3d6000fd5b505050506040513d602081101561113157600080fd5b50505050565b6000546001600160a01b031633148061115957506000546001600160a01b0316155b611198576040805162461bcd60e51b81526020600482018190526024820152600080516020611b63833981519152604482015290519081900360640190fd5b600254604080516370a0823160e01b815230600482015290516000926001600160a01b0316916370a08231916024808301926020929190829003018186803b1580156111e357600080fd5b505afa1580156111f7573d6000803e3d6000fd5b505050506040513d602081101561120d57600080fd5b505190508061124d5760405162461bcd60e51b8152600401808060200182810382526026815260200180611bad6026913960400191505060405180910390fd5b6002546040805163a9059cbb60e01b81526001600160a01b038581166004830152602482018590529151919092169163a9059cbb9160448083019260209291908290030181600087803b15801561110757600080fd5b6000806112b16109f1610937565b8051602091820120604080516001600160601b0319606089811b82168387015288901b1660348201528151602881830301815260489091019091528051920191909120909150611301818361171b565b95945050505050565b6001600160a01b038416611365576040805162461bcd60e51b815260206004820152601e60248201527f6f70657261746f722063616e2774206265207a65726f20616464726573730000604482015290519081900360640190fd5b6005548310156113a65760405162461bcd60e51b8152600401808060200182810382526032815260200180611ab66032913960400191505060405180910390fd5b60006113b1856109e3565b90506113bc81610a38565b1561140e576040805162461bcd60e51b815260206004820152601960248201527f6865726d657320616c7265616479207265676973746572656400000000000000604482015290519081900360640190fd5b600061142d866001600160a01b03166114286109f1610955565b611a23565b600254604080516323b872dd60e01b81523360048201526001600160a01b038085166024830152604482018a905291519394509116916323b872dd916064808201926020929091908290030181600087803b15801561148b57600080fd5b505af115801561149f573d6000803e3d6000fd5b505050506040513d60208110156114b557600080fd5b50506002546040805163fec8157d60e01b81526001600160a01b039283166004820152888316602482015261ffff871660448201526064810186905290519183169163fec8157d9160848082019260009290919082900301818387803b15801561151e57600080fd5b505af1158015611532573d6000803e3d6000fd5b50506040805180820182526001600160a01b03808b16808352640100000000600160c01b03602088811b821663fc0e3d9017861b8186019081528985166000818152600884528890209651875496166001600160a01b0319909616959095178655516001909501805495871c92831663ffffffff93909316929092176001600160c01b0319909516949094179055835190815292519094507f04924690bc091b8e80a2560b12c3fab5591a0e0114172cf52ccd7b771d49507b9350918290030190a2505050505050565b6000546001600160a01b031633148061161e57506000546001600160a01b0316155b61165d576040805162461bcd60e51b81526020600482018190526024820152600080516020611b63833981519152604482015290519081900360640190fd5b6001600160a01b0381166116a25760405162461bcd60e51b8152600401808060200182810382526026815260200180611a3e6026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b6001546001600160a01b031690565b6002546001600160a01b031681565b604080516001600160f81b03196020808301919091523060601b60218301526035820194909452605580820193909352815180820390930183526075019052805191012090565b600080826001600160a01b0316634e69d5606040518163ffffffff1660e01b815260040160206040518083038186803b15801561179e57600080fd5b505afa1580156117b2573d6000803e3d6000fd5b505050506040513d60208110156117c857600080fd5b5051905060008160038111156117da57fe5b149392505050565b6000815160411461183a576040805162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015290519081900360640190fd5b60208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08211156118ab5760405162461bcd60e51b8152600401808060200182810382526022815260200180611ae86022913960400191505060405180910390fd5b8060ff16601b141580156118c357508060ff16601c14155b156118ff5760405162461bcd60e51b8152600401808060200182810382526022815260200180611b416022913960400191505060405180910390fd5b60408051600080825260208083018085528a905260ff85168385015260608301879052608083018690529251909260019260a080820193601f1981019281900390910190855afa158015611957573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166119bf576040805162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015290519081900360640190fd5b9695505050505050565b600082820183811015610a13576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b600080838351602085016000f59050803b610a1357600080fdfe4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573736865726d65732073686f756c642067657420617070726f76616c20746f207472616e7366657220746f6b656e736e617469766520746f6b656e2066756e64732063616e2774206265207265636f76657265646865726d6573206861766520746f207374616b65206174206c65617374206d696e696d616c207374616b6520616d6f756e7445434453413a20696e76616c6964207369676e6174757265202773272076616c75653d602d80600a3d3981f3363d3d373d3d3d363d73bebebebebebebebebebebebebebebebebebebebe5af43d82803e903d91602b57fd5bf345434453413a20696e76616c6964207369676e6174757265202776272076616c75654f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726e6f7420656e6f756768742066756e647320696e206368616e6e656c20746f20636f7665722066656573636f6c6c6563746564206665652063616e6e6f74206265206c657373207468616e207a65726fa26469706673582212208ff953adb989f70a1beabbeabb366da58a8a754823cc7b0b72f4318eea69233364736f6c634300060a0033"

// DeployRegistry deploys a new Ethereum contract, binding an instance of Registry to it.
func DeployRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _tokenAddress common.Address, _dexAddress common.Address, _regFee *big.Int, _minimalHermesStake *big.Int, _channelImplementation common.Address, _hermesImplementation common.Address) (common.Address, *types.Transaction, *Registry, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RegistryBin), backend, _tokenAddress, _dexAddress, _regFee, _minimalHermesStake, _channelImplementation, _hermesImplementation)
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

// Dex is a free data retrieval call binding the contract method 0x692058c2.
//
// Solidity: function dex() view returns(address)
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
// Solidity: function dex() view returns(address)
func (_Registry *RegistrySession) Dex() (common.Address, error) {
	return _Registry.Contract.Dex(&_Registry.CallOpts)
}

// Dex is a free data retrieval call binding the contract method 0x692058c2.
//
// Solidity: function dex() view returns(address)
func (_Registry *RegistryCallerSession) Dex() (common.Address, error) {
	return _Registry.Contract.Dex(&_Registry.CallOpts)
}

// GetChannelAddress is a free data retrieval call binding the contract method 0xe617aaac.
//
// Solidity: function getChannelAddress(address _identity, address _hermesId) view returns(address)
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
// Solidity: function getChannelAddress(address _identity, address _hermesId) view returns(address)
func (_Registry *RegistrySession) GetChannelAddress(_identity common.Address, _hermesId common.Address) (common.Address, error) {
	return _Registry.Contract.GetChannelAddress(&_Registry.CallOpts, _identity, _hermesId)
}

// GetChannelAddress is a free data retrieval call binding the contract method 0xe617aaac.
//
// Solidity: function getChannelAddress(address _identity, address _hermesId) view returns(address)
func (_Registry *RegistryCallerSession) GetChannelAddress(_identity common.Address, _hermesId common.Address) (common.Address, error) {
	return _Registry.Contract.GetChannelAddress(&_Registry.CallOpts, _identity, _hermesId)
}

// GetChannelImplementation is a free data retrieval call binding the contract method 0x7c671a21.
//
// Solidity: function getChannelImplementation() view returns(address)
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
// Solidity: function getChannelImplementation() view returns(address)
func (_Registry *RegistrySession) GetChannelImplementation() (common.Address, error) {
	return _Registry.Contract.GetChannelImplementation(&_Registry.CallOpts)
}

// GetChannelImplementation is a free data retrieval call binding the contract method 0x7c671a21.
//
// Solidity: function getChannelImplementation() view returns(address)
func (_Registry *RegistryCallerSession) GetChannelImplementation() (common.Address, error) {
	return _Registry.Contract.GetChannelImplementation(&_Registry.CallOpts)
}

// GetFundsDestination is a free data retrieval call binding the contract method 0xf58c5b6e.
//
// Solidity: function getFundsDestination() view returns(address)
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
// Solidity: function getFundsDestination() view returns(address)
func (_Registry *RegistrySession) GetFundsDestination() (common.Address, error) {
	return _Registry.Contract.GetFundsDestination(&_Registry.CallOpts)
}

// GetFundsDestination is a free data retrieval call binding the contract method 0xf58c5b6e.
//
// Solidity: function getFundsDestination() view returns(address)
func (_Registry *RegistryCallerSession) GetFundsDestination() (common.Address, error) {
	return _Registry.Contract.GetFundsDestination(&_Registry.CallOpts)
}

// GetHermesAddress is a free data retrieval call binding the contract method 0xacc831d0.
//
// Solidity: function getHermesAddress(address _hermesOperator) view returns(address)
func (_Registry *RegistryCaller) GetHermesAddress(opts *bind.CallOpts, _hermesOperator common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "getHermesAddress", _hermesOperator)
	return *ret0, err
}

// GetHermesAddress is a free data retrieval call binding the contract method 0xacc831d0.
//
// Solidity: function getHermesAddress(address _hermesOperator) view returns(address)
func (_Registry *RegistrySession) GetHermesAddress(_hermesOperator common.Address) (common.Address, error) {
	return _Registry.Contract.GetHermesAddress(&_Registry.CallOpts, _hermesOperator)
}

// GetHermesAddress is a free data retrieval call binding the contract method 0xacc831d0.
//
// Solidity: function getHermesAddress(address _hermesOperator) view returns(address)
func (_Registry *RegistryCallerSession) GetHermesAddress(_hermesOperator common.Address) (common.Address, error) {
	return _Registry.Contract.GetHermesAddress(&_Registry.CallOpts, _hermesOperator)
}

// GetHermesImplementation is a free data retrieval call binding the contract method 0x9936a87b.
//
// Solidity: function getHermesImplementation() view returns(address)
func (_Registry *RegistryCaller) GetHermesImplementation(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "getHermesImplementation")
	return *ret0, err
}

// GetHermesImplementation is a free data retrieval call binding the contract method 0x9936a87b.
//
// Solidity: function getHermesImplementation() view returns(address)
func (_Registry *RegistrySession) GetHermesImplementation() (common.Address, error) {
	return _Registry.Contract.GetHermesImplementation(&_Registry.CallOpts)
}

// GetHermesImplementation is a free data retrieval call binding the contract method 0x9936a87b.
//
// Solidity: function getHermesImplementation() view returns(address)
func (_Registry *RegistryCallerSession) GetHermesImplementation() (common.Address, error) {
	return _Registry.Contract.GetHermesImplementation(&_Registry.CallOpts)
}

// GetProxyCode is a free data retrieval call binding the contract method 0xab867213.
//
// Solidity: function getProxyCode(address _implementation) pure returns(bytes)
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
// Solidity: function getProxyCode(address _implementation) pure returns(bytes)
func (_Registry *RegistrySession) GetProxyCode(_implementation common.Address) ([]byte, error) {
	return _Registry.Contract.GetProxyCode(&_Registry.CallOpts, _implementation)
}

// GetProxyCode is a free data retrieval call binding the contract method 0xab867213.
//
// Solidity: function getProxyCode(address _implementation) pure returns(bytes)
func (_Registry *RegistryCallerSession) GetProxyCode(_implementation common.Address) ([]byte, error) {
	return _Registry.Contract.GetProxyCode(&_Registry.CallOpts, _implementation)
}

// Hermeses is a free data retrieval call binding the contract method 0x03fb422f.
//
// Solidity: function hermeses(address ) view returns(address operator, function stake)
func (_Registry *RegistryCaller) Hermeses(opts *bind.CallOpts, arg0 common.Address) (struct {
	Operator common.Address
	Stake    [24]byte
}, error) {
	ret := new(struct {
		Operator common.Address
		Stake    [24]byte
	})
	out := ret
	err := _Registry.contract.Call(opts, out, "hermeses", arg0)
	return *ret, err
}

// Hermeses is a free data retrieval call binding the contract method 0x03fb422f.
//
// Solidity: function hermeses(address ) view returns(address operator, function stake)
func (_Registry *RegistrySession) Hermeses(arg0 common.Address) (struct {
	Operator common.Address
	Stake    [24]byte
}, error) {
	return _Registry.Contract.Hermeses(&_Registry.CallOpts, arg0)
}

// Hermeses is a free data retrieval call binding the contract method 0x03fb422f.
//
// Solidity: function hermeses(address ) view returns(address operator, function stake)
func (_Registry *RegistryCallerSession) Hermeses(arg0 common.Address) (struct {
	Operator common.Address
	Stake    [24]byte
}, error) {
	return _Registry.Contract.Hermeses(&_Registry.CallOpts, arg0)
}

// IsHermes is a free data retrieval call binding the contract method 0xcdd596e0.
//
// Solidity: function isHermes(address _hermesId) view returns(bool)
func (_Registry *RegistryCaller) IsHermes(opts *bind.CallOpts, _hermesId common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "isHermes", _hermesId)
	return *ret0, err
}

// IsHermes is a free data retrieval call binding the contract method 0xcdd596e0.
//
// Solidity: function isHermes(address _hermesId) view returns(bool)
func (_Registry *RegistrySession) IsHermes(_hermesId common.Address) (bool, error) {
	return _Registry.Contract.IsHermes(&_Registry.CallOpts, _hermesId)
}

// IsHermes is a free data retrieval call binding the contract method 0xcdd596e0.
//
// Solidity: function isHermes(address _hermesId) view returns(bool)
func (_Registry *RegistryCallerSession) IsHermes(_hermesId common.Address) (bool, error) {
	return _Registry.Contract.IsHermes(&_Registry.CallOpts, _hermesId)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address _identityHash) view returns(bool)
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
// Solidity: function isRegistered(address _identityHash) view returns(bool)
func (_Registry *RegistrySession) IsRegistered(_identityHash common.Address) (bool, error) {
	return _Registry.Contract.IsRegistered(&_Registry.CallOpts, _identityHash)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address _identityHash) view returns(bool)
func (_Registry *RegistryCallerSession) IsRegistered(_identityHash common.Address) (bool, error) {
	return _Registry.Contract.IsRegistered(&_Registry.CallOpts, _identityHash)
}

// MinimalHermesStake is a free data retrieval call binding the contract method 0x66cf5875.
//
// Solidity: function minimalHermesStake() view returns(uint256)
func (_Registry *RegistryCaller) MinimalHermesStake(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "minimalHermesStake")
	return *ret0, err
}

// MinimalHermesStake is a free data retrieval call binding the contract method 0x66cf5875.
//
// Solidity: function minimalHermesStake() view returns(uint256)
func (_Registry *RegistrySession) MinimalHermesStake() (*big.Int, error) {
	return _Registry.Contract.MinimalHermesStake(&_Registry.CallOpts)
}

// MinimalHermesStake is a free data retrieval call binding the contract method 0x66cf5875.
//
// Solidity: function minimalHermesStake() view returns(uint256)
func (_Registry *RegistryCallerSession) MinimalHermesStake() (*big.Int, error) {
	return _Registry.Contract.MinimalHermesStake(&_Registry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
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
// Solidity: function owner() view returns(address)
func (_Registry *RegistrySession) Owner() (common.Address, error) {
	return _Registry.Contract.Owner(&_Registry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registry *RegistryCallerSession) Owner() (common.Address, error) {
	return _Registry.Contract.Owner(&_Registry.CallOpts)
}

// RegistrationFee is a free data retrieval call binding the contract method 0x14c44e09.
//
// Solidity: function registrationFee() view returns(uint256)
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
// Solidity: function registrationFee() view returns(uint256)
func (_Registry *RegistrySession) RegistrationFee() (*big.Int, error) {
	return _Registry.Contract.RegistrationFee(&_Registry.CallOpts)
}

// RegistrationFee is a free data retrieval call binding the contract method 0x14c44e09.
//
// Solidity: function registrationFee() view returns(uint256)
func (_Registry *RegistryCallerSession) RegistrationFee() (*big.Int, error) {
	return _Registry.Contract.RegistrationFee(&_Registry.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
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
// Solidity: function token() view returns(address)
func (_Registry *RegistrySession) Token() (common.Address, error) {
	return _Registry.Contract.Token(&_Registry.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
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

// RegisterHermes is a paid mutator transaction binding the contract method 0xe64324fd.
//
// Solidity: function registerHermes(address _hermesOperator, uint256 _stakeAmount, uint16 _hermesFee, uint256 _maxStake) returns()
func (_Registry *RegistryTransactor) RegisterHermes(opts *bind.TransactOpts, _hermesOperator common.Address, _stakeAmount *big.Int, _hermesFee uint16, _maxStake *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "registerHermes", _hermesOperator, _stakeAmount, _hermesFee, _maxStake)
}

// RegisterHermes is a paid mutator transaction binding the contract method 0xe64324fd.
//
// Solidity: function registerHermes(address _hermesOperator, uint256 _stakeAmount, uint16 _hermesFee, uint256 _maxStake) returns()
func (_Registry *RegistrySession) RegisterHermes(_hermesOperator common.Address, _stakeAmount *big.Int, _hermesFee uint16, _maxStake *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.RegisterHermes(&_Registry.TransactOpts, _hermesOperator, _stakeAmount, _hermesFee, _maxStake)
}

// RegisterHermes is a paid mutator transaction binding the contract method 0xe64324fd.
//
// Solidity: function registerHermes(address _hermesOperator, uint256 _stakeAmount, uint16 _hermesFee, uint256 _maxStake) returns()
func (_Registry *RegistryTransactorSession) RegisterHermes(_hermesOperator common.Address, _stakeAmount *big.Int, _hermesFee uint16, _maxStake *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.RegisterHermes(&_Registry.TransactOpts, _hermesOperator, _stakeAmount, _hermesFee, _maxStake)
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

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Registry *RegistryTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Registry *RegistrySession) Receive() (*types.Transaction, error) {
	return _Registry.Contract.Receive(&_Registry.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Registry *RegistryTransactorSession) Receive() (*types.Transaction, error) {
	return _Registry.Contract.Receive(&_Registry.TransactOpts)
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
	HermesId       common.Address
	ChannelAddress common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterConsumerChannelCreated is a free log retrieval operation binding the contract event 0x2ed7bcf2ff03098102c7003d7ce2a633e4b49b8198b07de5383cdf4c0ab9228b.
//
// Solidity: event ConsumerChannelCreated(address indexed identityHash, address indexed hermesId, address channelAddress)
func (_Registry *RegistryFilterer) FilterConsumerChannelCreated(opts *bind.FilterOpts, identityHash []common.Address, hermesId []common.Address) (*RegistryConsumerChannelCreatedIterator, error) {

	var identityHashRule []interface{}
	for _, identityHashItem := range identityHash {
		identityHashRule = append(identityHashRule, identityHashItem)
	}
	var hermesIdRule []interface{}
	for _, hermesIdItem := range hermesId {
		hermesIdRule = append(hermesIdRule, hermesIdItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "ConsumerChannelCreated", identityHashRule, hermesIdRule)
	if err != nil {
		return nil, err
	}
	return &RegistryConsumerChannelCreatedIterator{contract: _Registry.contract, event: "ConsumerChannelCreated", logs: logs, sub: sub}, nil
}

// WatchConsumerChannelCreated is a free log subscription operation binding the contract event 0x2ed7bcf2ff03098102c7003d7ce2a633e4b49b8198b07de5383cdf4c0ab9228b.
//
// Solidity: event ConsumerChannelCreated(address indexed identityHash, address indexed hermesId, address channelAddress)
func (_Registry *RegistryFilterer) WatchConsumerChannelCreated(opts *bind.WatchOpts, sink chan<- *RegistryConsumerChannelCreated, identityHash []common.Address, hermesId []common.Address) (event.Subscription, error) {

	var identityHashRule []interface{}
	for _, identityHashItem := range identityHash {
		identityHashRule = append(identityHashRule, identityHashItem)
	}
	var hermesIdRule []interface{}
	for _, hermesIdItem := range hermesId {
		hermesIdRule = append(hermesIdRule, hermesIdItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "ConsumerChannelCreated", identityHashRule, hermesIdRule)
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
// Solidity: event ConsumerChannelCreated(address indexed identityHash, address indexed hermesId, address channelAddress)
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

// RegistryRegisteredHermesIterator is returned from FilterRegisteredHermes and is used to iterate over the raw logs and unpacked data for RegisteredHermes events raised by the Registry contract.
type RegistryRegisteredHermesIterator struct {
	Event *RegistryRegisteredHermes // Event containing the contract specifics and raw log

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
func (it *RegistryRegisteredHermesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryRegisteredHermes)
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
		it.Event = new(RegistryRegisteredHermes)
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
func (it *RegistryRegisteredHermesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryRegisteredHermesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryRegisteredHermes represents a RegisteredHermes event raised by the Registry contract.
type RegistryRegisteredHermes struct {
	HermesId       common.Address
	HermesOperator common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRegisteredHermes is a free log retrieval operation binding the contract event 0x04924690bc091b8e80a2560b12c3fab5591a0e0114172cf52ccd7b771d49507b.
//
// Solidity: event RegisteredHermes(address indexed hermesId, address hermesOperator)
func (_Registry *RegistryFilterer) FilterRegisteredHermes(opts *bind.FilterOpts, hermesId []common.Address) (*RegistryRegisteredHermesIterator, error) {

	var hermesIdRule []interface{}
	for _, hermesIdItem := range hermesId {
		hermesIdRule = append(hermesIdRule, hermesIdItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "RegisteredHermes", hermesIdRule)
	if err != nil {
		return nil, err
	}
	return &RegistryRegisteredHermesIterator{contract: _Registry.contract, event: "RegisteredHermes", logs: logs, sub: sub}, nil
}

// WatchRegisteredHermes is a free log subscription operation binding the contract event 0x04924690bc091b8e80a2560b12c3fab5591a0e0114172cf52ccd7b771d49507b.
//
// Solidity: event RegisteredHermes(address indexed hermesId, address hermesOperator)
func (_Registry *RegistryFilterer) WatchRegisteredHermes(opts *bind.WatchOpts, sink chan<- *RegistryRegisteredHermes, hermesId []common.Address) (event.Subscription, error) {

	var hermesIdRule []interface{}
	for _, hermesIdItem := range hermesId {
		hermesIdRule = append(hermesIdRule, hermesIdItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "RegisteredHermes", hermesIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryRegisteredHermes)
				if err := _Registry.contract.UnpackLog(event, "RegisteredHermes", log); err != nil {
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

// ParseRegisteredHermes is a log parse operation binding the contract event 0x04924690bc091b8e80a2560b12c3fab5591a0e0114172cf52ccd7b771d49507b.
//
// Solidity: event RegisteredHermes(address indexed hermesId, address hermesOperator)
func (_Registry *RegistryFilterer) ParseRegisteredHermes(log types.Log) (*RegistryRegisteredHermes, error) {
	event := new(RegistryRegisteredHermes)
	if err := _Registry.contract.UnpackLog(event, "RegisteredHermes", log); err != nil {
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
	HermesId     common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRegisteredIdentity is a free log retrieval operation binding the contract event 0xefaf768237c22e140a862d5d375ad5c153479fac3f8bcf8b580a1651fd62c3ef.
//
// Solidity: event RegisteredIdentity(address indexed identityHash, address indexed hermesId)
func (_Registry *RegistryFilterer) FilterRegisteredIdentity(opts *bind.FilterOpts, identityHash []common.Address, hermesId []common.Address) (*RegistryRegisteredIdentityIterator, error) {

	var identityHashRule []interface{}
	for _, identityHashItem := range identityHash {
		identityHashRule = append(identityHashRule, identityHashItem)
	}
	var hermesIdRule []interface{}
	for _, hermesIdItem := range hermesId {
		hermesIdRule = append(hermesIdRule, hermesIdItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "RegisteredIdentity", identityHashRule, hermesIdRule)
	if err != nil {
		return nil, err
	}
	return &RegistryRegisteredIdentityIterator{contract: _Registry.contract, event: "RegisteredIdentity", logs: logs, sub: sub}, nil
}

// WatchRegisteredIdentity is a free log subscription operation binding the contract event 0xefaf768237c22e140a862d5d375ad5c153479fac3f8bcf8b580a1651fd62c3ef.
//
// Solidity: event RegisteredIdentity(address indexed identityHash, address indexed hermesId)
func (_Registry *RegistryFilterer) WatchRegisteredIdentity(opts *bind.WatchOpts, sink chan<- *RegistryRegisteredIdentity, identityHash []common.Address, hermesId []common.Address) (event.Subscription, error) {

	var identityHashRule []interface{}
	for _, identityHashItem := range identityHash {
		identityHashRule = append(identityHashRule, identityHashItem)
	}
	var hermesIdRule []interface{}
	for _, hermesIdItem := range hermesId {
		hermesIdRule = append(hermesIdRule, hermesIdItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "RegisteredIdentity", identityHashRule, hermesIdRule)
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
// Solidity: event RegisteredIdentity(address indexed identityHash, address indexed hermesId)
func (_Registry *RegistryFilterer) ParseRegisteredIdentity(log types.Log) (*RegistryRegisteredIdentity, error) {
	event := new(RegistryRegisteredIdentity)
	if err := _Registry.contract.UnpackLog(event, "RegisteredIdentity", log); err != nil {
		return nil, err
	}
	return event, nil
}
