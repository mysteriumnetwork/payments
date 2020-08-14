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

// ChannelImplementationABI is the input ABI used to generate the binding from.
const ChannelImplementationABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"hermes\",\"type\":\"address\"}],\"name\":\"ChannelInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousDestination\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newDestination\",\"type\":\"address\"}],\"name\":\"DestinationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timelock\",\"type\":\"uint256\"}],\"name\":\"ExitRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FinalizeExit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalSettled\",\"type\":\"uint256\"}],\"name\":\"PromiseSettled\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"claimEthers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"claimTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exitRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timelock\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFundsDestination\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hermes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"settled\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_newDestination\",\"type\":\"address\"}],\"name\":\"setFundsDestination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20Token\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dex\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_identityHash\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_hermesId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_transactorFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_lock\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"settlePromise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_validUntil\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"requestExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalizeExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_newDestination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"setFundsDestinationByCheque\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ChannelImplementationBin is the compiled bytecode used for deploying new contracts.
var ChannelImplementationBin = "0x608060405234801561001057600080fd5b50611db7806100206000396000f3fe60806040526004361061010d5760003560e01c80637c3e610511610095578063f2fde38b11610064578063f2fde38b14610595578063f4b3a197146105c8578063f58c5b6e146105fe578063f7013ef614610613578063fc0c546a14610666576101d0565b80637c3e6105146104465780638da5cb5b1461050e578063d8092c9214610523578063df8de3e714610562576101d0565b8063570ca735116100dc578063570ca73514610310578063692058c2146103415780636931b550146103565780636f1746301461036b578063715018a614610431576101d0565b806307e8ec1f146101d5578063182f3488146101ec578063238e130a146102b4578063392e53cd146102e7576101d0565b366101d0576009546040516000916001600160a01b031690349083903690808383808284376040519201945060009350909150508083038185875af1925050503d8060008114610179576040519150601f19603f3d011682016040523d82523d6000602084013e61017e565b606091505b50509050806101cd576040805162461bcd60e51b81526020600482015260166024820152750a8f040eec2e640e4cad4cac6e8cac840c4f240888ab60531b604482015290519081900360640190fd5b50005b600080fd5b3480156101e157600080fd5b506101ea61067b565b005b3480156101f857600080fd5b506101ea6004803603606081101561020f57600080fd5b6001600160a01b038235169160208101359181019060608101604082013564010000000081111561023f57600080fd5b82018360208201111561025157600080fd5b8035906020019184600183028401116401000000008311171561027357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610827945050505050565b3480156102c057600080fd5b506101ea600480360360208110156102d757600080fd5b50356001600160a01b0316610af7565b3480156102f357600080fd5b506102fc610bd9565b604080519115158252519081900360200190f35b34801561031c57600080fd5b50610325610bea565b604080516001600160a01b039092168252519081900360200190f35b34801561034d57600080fd5b50610325610bf9565b34801561036257600080fd5b506101ea610c08565b34801561037757600080fd5b506101ea6004803603608081101561038e57600080fd5b813591602081013591604082013591908101906080810160608201356401000000008111156103bc57600080fd5b8201836020820111156103ce57600080fd5b803590602001918460018302840111640100000000831117156103f057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610c59945050505050565b34801561043d57600080fd5b506101ea610f76565b34801561045257600080fd5b506101ea6004803603606081101561046957600080fd5b6001600160a01b038235169160208101359181019060608101604082013564010000000081111561049957600080fd5b8201836020820111156104ab57600080fd5b803590602001918460018302840111640100000000831117156104cd57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611033945050505050565b34801561051a57600080fd5b50610325611203565b34801561052f57600080fd5b50610538611212565b604080516001600160a01b0394851681529290931660208301528183015290519081900360600190f35b34801561056e57600080fd5b506101ea6004803603602081101561058557600080fd5b50356001600160a01b031661122e565b3480156105a157600080fd5b506101ea600480360360208110156105b857600080fd5b50356001600160a01b0316611394565b3480156105d457600080fd5b506105dd6114a7565b604080519283526001600160a01b0390911660208301528051918290030190f35b34801561060a57600080fd5b506103256114b9565b34801561061f57600080fd5b506101ea600480360360a081101561063657600080fd5b506001600160a01b03813581169160208101358216916040820135811691606081013590911690608001356114c8565b34801561067257600080fd5b506103256117ec565b6003541580159061068e57506003544310155b6106c95760405162461bcd60e51b8152600401808060200182810382526039815260200180611d246039913960400191505060405180910390fd5b600254604080516370a0823160e01b815230600482015290516000926001600160a01b0316916370a08231916024808301926020929190829003018186803b15801561071457600080fd5b505afa158015610728573d6000803e3d6000fd5b505050506040513d602081101561073e57600080fd5b5051600254600480546040805163a9059cbb60e01b81526001600160a01b039283169381019390935260248301859052519394509091169163a9059cbb916044808201926020929091908290030181600087803b15801561079e57600080fd5b505af11580156107b2573d6000803e3d6000fd5b505050506040513d60208110156107c857600080fd5b505060408051808201825260008082526020918201819052600355600480546001600160a01b0319169055815183815291517f50128f92fd19060780780085c779f5ddebca701ad03dc303be5b0859863458249281900390910190a150565b60006108316117fb565b600354909150156108735760405162461bcd60e51b8152600401808060200182810382526039815260200180611bf26039913960400191505060405180910390fd5b4383116108b15760405162461bcd60e51b8152600401808060200182810382526038815260200180611b4d6038913960400191505060405180910390fd5b8281116108ef5760405162461bcd60e51b8152600401808060200182810382526032815260200180611cf26032913960400191505060405180910390fd5b6001600160a01b0384166109345760405162461bcd60e51b8152600401808060200182810382526021815260200180611c2b6021913960400191505060405180910390fd5b6008546001600160a01b03163314610a8a5760003090506000610a22846040518060400160405280600d81526020016c22bc34ba103932b8bab2b9ba1d60991b8152508489896040516020018085805190602001908083835b602083106109ac5780518252601f19909201916020918201910161098d565b6001836020036101000a038019825116818451168082178552505050505050905001846001600160a01b031660601b8152601401836001600160a01b031660601b81526014018281526020019450505050506040516020818303038152906040528051906020012061180390919063ffffffff16565b6008549091506001600160a01b03808316911614610a87576040805162461bcd60e51b815260206004820152601d60248201527f6861766520746f206265207369676e6564206279206f70657261746f72000000604482015290519081900360640190fd5b50505b6040805180820182528281526001600160a01b03861660209182018190526003849055600480546001600160a01b0319169091179055815183815291517fe60f0366d8d61555184ea027447889648bae94ebfb1202a39544b6b6803969db9281900390910190a150505050565b6000546001600160a01b0316331480610b1957506000546001600160a01b0316155b610b6a576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6001600160a01b038116610b7d57600080fd5b6001546040516001600160a01b038084169216907fe1a66d77649cf0a57b9937073549f30f1c82bb865aaf066d2f299e37a62c6aad90600090a3600180546001600160a01b0319166001600160a01b0392909216919091179055565b6008546001600160a01b0316151590565b6008546001600160a01b031681565b6009546001600160a01b031681565b6001546001600160a01b0316610c1d57600080fd5b6001546040516001600160a01b03909116904780156108fc02916000818181858888f19350505050158015610c56573d6000803e3d6000fd5b50565b604080516020808201859052825180830382018152828401845280519082012030606084018190526080840189905260a0840188905260c08085018390528551808603909101815260e090940190945282519290910191909120909190600090610cc39085611803565b6008549091506001600160a01b03808316911614610d125760405162461bcd60e51b8152600401808060200182810382526025815260200180611d5d6025913960400191505060405180910390fd5b600754600090610d239089906119ee565b905060008111610d645760405162461bcd60e51b8152600401808060200182810382526037815260200180611c966037913960400191505060405180910390fd5b600254604080516370a0823160e01b81526001600160a01b038681166004830152915160009392909216916370a0823191602480820192602092909190829003018186803b158015610db557600080fd5b505afa158015610dc9573d6000803e3d6000fd5b505050506040513d6020811015610ddf57600080fd5b5051905080821115610def578091505b600754610dfc9083611a37565b6007556002546006546001600160a01b039182169163a9059cbb9116610e22858c6119ee565b6040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050602060405180830381600087803b158015610e6857600080fd5b505af1158015610e7c573d6000803e3d6000fd5b505050506040513d6020811015610e9257600080fd5b50508715610f1b576002546040805163a9059cbb60e01b8152336004820152602481018b905290516001600160a01b039092169163a9059cbb916044808201926020929091908290030181600087803b158015610eee57600080fd5b505af1158015610f02573d6000803e3d6000fd5b505050506040513d6020811015610f1857600080fd5b50505b600654600754604080516001600160a01b0390931683526020830185905282810191909152517f50c3491624aa1825a7653df63d067fecd5c8634ba63c99c4a7cf04ff1436070b9181900360600190a1505050505050505050565b6000546001600160a01b0316331480610f9857506000546001600160a01b0316155b610fe9576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6001600160a01b03831661104657600080fd5b600a5482116110865760405162461bcd60e51b8152600401808060200182810382526025815260200180611ccd6025913960400191505060405180910390fd5b6000611152826040518060400160405280601681526020017529b2ba10333ab73239903232b9ba34b730ba34b7b71d60511b81525086866040516020018084805190602001908083835b602083106110ef5780518252601f1990920191602091820191016110d0565b6001836020036101000a038019825116818451168082178552505050505050905001836001600160a01b031660601b815260140182815260200193505050506040516020818303038152906040528051906020012061180390919063ffffffff16565b6008549091506001600160a01b038083169116146111a15760405162461bcd60e51b8152600401808060200182810382526024815260200180611b296024913960400191505060405180910390fd5b6001546040516001600160a01b038087169216907fe1a66d77649cf0a57b9937073549f30f1c82bb865aaf066d2f299e37a62c6aad90600090a35050600180546001600160a01b0319166001600160a01b039390931692909217909155600a55565b6000546001600160a01b031690565b6005546006546007546001600160a01b03928316929091169083565b6001546001600160a01b031661124357600080fd5b6002546001600160a01b03828116911614156112905760405162461bcd60e51b8152600401808060200182810382526025815260200180611bab6025913960400191505060405180910390fd5b6000816001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b1580156112df57600080fd5b505afa1580156112f3573d6000803e3d6000fd5b505050506040513d602081101561130957600080fd5b50516001546040805163a9059cbb60e01b81526001600160a01b0392831660048201526024810184905290519293509084169163a9059cbb916044808201926020929091908290030181600087803b15801561136457600080fd5b505af1158015611378573d6000803e3d6000fd5b505050506040513d602081101561138e57600080fd5b50505050565b6000546001600160a01b03163314806113b657506000546001600160a01b0316155b611407576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6001600160a01b03811661144c5760405162461bcd60e51b8152600401808060200182810382526026815260200180611b856026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b6003546004546001600160a01b031682565b6001546001600160a01b031690565b6114d0610bd9565b1561151b576040805162461bcd60e51b8152602060048201526016602482015275125cc8185b1c9958591e481a5b9a5d1a585b1a5e995960521b604482015290519081900360640190fd5b6001600160a01b03831661156f576040805162461bcd60e51b81526020600482015260166024820152754964656e746974792063616e2774206265207a65726f60501b604482015290519081900360640190fd5b6001600160a01b0382166115c3576040805162461bcd60e51b81526020600482015260166024820152754865726d657349442063616e2774206265207a65726f60501b604482015290519081900360640190fd5b6001600160a01b0385166116085760405162461bcd60e51b8152600401808060200182810382526028815260200180611c4c6028913960400191505060405180910390fd5b600280546001600160a01b038088166001600160a01b031992831617909255600980549287169290911691909117905580156116bf576002546040805163a9059cbb60e01b81523360048201526024810184905290516001600160a01b039092169163a9059cbb916044808201926020929091908290030181600087803b15801561169257600080fd5b505af11580156116a6573d6000803e3d6000fd5b505050506040513d60208110156116bc57600080fd5b50505b600880546001600160a01b0319166001600160a01b0385811691909117918290556116ea9116611394565b6040518060600160405280836001600160a01b031663e7f43c686040518163ffffffff1660e01b815260040160206040518083038186803b15801561172e57600080fd5b505afa158015611742573d6000803e3d6000fd5b505050506040513d602081101561175857600080fd5b50516001600160a01b039081168252848116602083810182905260006040948501528451600580549185166001600160a01b031992831617905585820151600680549186169190921617905593830151600755825191871682529281019290925280517f9a7def6556351196c74c99e1cc8dcd284e9da181ea854c3e6367cc9fad882a519281900390910190a15050505050565b6002546001600160a01b031681565b436146500190565b6000815160411461185b576040805162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015290519081900360640190fd5b60208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08211156118cc5760405162461bcd60e51b8152600401808060200182810382526022815260200180611bd06022913960400191505060405180910390fd5b8060ff16601b141580156118e457508060ff16601c14155b156119205760405162461bcd60e51b8152600401808060200182810382526022815260200180611c746022913960400191505060405180910390fd5b600060018783868660405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa15801561197c573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166119e4576040805162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015290519081900360640190fd5b9695505050505050565b6000611a3083836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250611a91565b9392505050565b600082820183811015611a30576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b60008184841115611b205760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611ae5578181015183820152602001611acd565b50505050905090810190601f168015611b125780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50505090039056fe6861766520746f206265207369676e65642062792070726f706572206964656e7469747976616c696420756e74696c206861766520746f2062652067726561746572207468616e2063757272656e7420626c6f636b206e756d6265724f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573736e617469766520746f6b656e2066756e64732063616e2774206265207265636f766572656445434453413a20696e76616c6964207369676e6174757265202773272076616c75656e657720657869742063616e20626520726571756573746564206f6e6c79207768656e206f6c64206f6e65207761732066696e616c6973656462656e65666963696172792063616e2774206265207a65726f2061646472657373546f6b656e2063616e2774206265206465706c6f796420696e746f207a65726f206164647265737345434453413a20696e76616c6964207369676e6174757265202776272076616c7565616d6f756e7420746f20736574746c652073686f756c642062652067726561746572207468617420616c726561647920736574746c65646e6f6e6365206861766520746f20626520626967676572207468616e206c617374206f6e6572657175657374206861766520746f2062652076616c69642073686f72746572207468616e2044454c41595f424c4f434b5365786974206861766520746f2062652072657175657374656420616e642074696d656c6f636b206861766520746f20626520696e20706173746861766520746f206265207369676e6564206279206368616e6e656c206f70657261746f72a264697066735822122080502b97debe93a7622593061b6263f6fa2a5f684a764a6bd3bbf1e77c4b723e64736f6c634300060c0033"

// DeployChannelImplementation deploys a new Ethereum contract, binding an instance of ChannelImplementation to it.
func DeployChannelImplementation(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ChannelImplementation, error) {
	parsed, err := abi.JSON(strings.NewReader(ChannelImplementationABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChannelImplementationBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChannelImplementation{ChannelImplementationCaller: ChannelImplementationCaller{contract: contract}, ChannelImplementationTransactor: ChannelImplementationTransactor{contract: contract}, ChannelImplementationFilterer: ChannelImplementationFilterer{contract: contract}}, nil
}

// ChannelImplementation is an auto generated Go binding around an Ethereum contract.
type ChannelImplementation struct {
	ChannelImplementationCaller     // Read-only binding to the contract
	ChannelImplementationTransactor // Write-only binding to the contract
	ChannelImplementationFilterer   // Log filterer for contract events
}

// ChannelImplementationCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChannelImplementationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChannelImplementationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChannelImplementationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChannelImplementationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChannelImplementationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChannelImplementationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChannelImplementationSession struct {
	Contract     *ChannelImplementation // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ChannelImplementationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChannelImplementationCallerSession struct {
	Contract *ChannelImplementationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// ChannelImplementationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChannelImplementationTransactorSession struct {
	Contract     *ChannelImplementationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ChannelImplementationRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChannelImplementationRaw struct {
	Contract *ChannelImplementation // Generic contract binding to access the raw methods on
}

// ChannelImplementationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChannelImplementationCallerRaw struct {
	Contract *ChannelImplementationCaller // Generic read-only contract binding to access the raw methods on
}

// ChannelImplementationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChannelImplementationTransactorRaw struct {
	Contract *ChannelImplementationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChannelImplementation creates a new instance of ChannelImplementation, bound to a specific deployed contract.
func NewChannelImplementation(address common.Address, backend bind.ContractBackend) (*ChannelImplementation, error) {
	contract, err := bindChannelImplementation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChannelImplementation{ChannelImplementationCaller: ChannelImplementationCaller{contract: contract}, ChannelImplementationTransactor: ChannelImplementationTransactor{contract: contract}, ChannelImplementationFilterer: ChannelImplementationFilterer{contract: contract}}, nil
}

// NewChannelImplementationCaller creates a new read-only instance of ChannelImplementation, bound to a specific deployed contract.
func NewChannelImplementationCaller(address common.Address, caller bind.ContractCaller) (*ChannelImplementationCaller, error) {
	contract, err := bindChannelImplementation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChannelImplementationCaller{contract: contract}, nil
}

// NewChannelImplementationTransactor creates a new write-only instance of ChannelImplementation, bound to a specific deployed contract.
func NewChannelImplementationTransactor(address common.Address, transactor bind.ContractTransactor) (*ChannelImplementationTransactor, error) {
	contract, err := bindChannelImplementation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChannelImplementationTransactor{contract: contract}, nil
}

// NewChannelImplementationFilterer creates a new log filterer instance of ChannelImplementation, bound to a specific deployed contract.
func NewChannelImplementationFilterer(address common.Address, filterer bind.ContractFilterer) (*ChannelImplementationFilterer, error) {
	contract, err := bindChannelImplementation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChannelImplementationFilterer{contract: contract}, nil
}

// bindChannelImplementation binds a generic wrapper to an already deployed contract.
func bindChannelImplementation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChannelImplementationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChannelImplementation *ChannelImplementationRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChannelImplementation.Contract.ChannelImplementationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChannelImplementation *ChannelImplementationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChannelImplementation.Contract.ChannelImplementationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChannelImplementation *ChannelImplementationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChannelImplementation.Contract.ChannelImplementationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChannelImplementation *ChannelImplementationCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChannelImplementation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChannelImplementation *ChannelImplementationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChannelImplementation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChannelImplementation *ChannelImplementationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChannelImplementation.Contract.contract.Transact(opts, method, params...)
}

// Dex is a free data retrieval call binding the contract method 0x692058c2.
//
// Solidity: function dex() view returns(address)
func (_ChannelImplementation *ChannelImplementationCaller) Dex(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChannelImplementation.contract.Call(opts, out, "dex")
	return *ret0, err
}

// Dex is a free data retrieval call binding the contract method 0x692058c2.
//
// Solidity: function dex() view returns(address)
func (_ChannelImplementation *ChannelImplementationSession) Dex() (common.Address, error) {
	return _ChannelImplementation.Contract.Dex(&_ChannelImplementation.CallOpts)
}

// Dex is a free data retrieval call binding the contract method 0x692058c2.
//
// Solidity: function dex() view returns(address)
func (_ChannelImplementation *ChannelImplementationCallerSession) Dex() (common.Address, error) {
	return _ChannelImplementation.Contract.Dex(&_ChannelImplementation.CallOpts)
}

// ExitRequest is a free data retrieval call binding the contract method 0xf4b3a197.
//
// Solidity: function exitRequest() view returns(uint256 timelock, address beneficiary)
func (_ChannelImplementation *ChannelImplementationCaller) ExitRequest(opts *bind.CallOpts) (struct {
	Timelock    *big.Int
	Beneficiary common.Address
}, error) {
	ret := new(struct {
		Timelock    *big.Int
		Beneficiary common.Address
	})
	out := ret
	err := _ChannelImplementation.contract.Call(opts, out, "exitRequest")
	return *ret, err
}

// ExitRequest is a free data retrieval call binding the contract method 0xf4b3a197.
//
// Solidity: function exitRequest() view returns(uint256 timelock, address beneficiary)
func (_ChannelImplementation *ChannelImplementationSession) ExitRequest() (struct {
	Timelock    *big.Int
	Beneficiary common.Address
}, error) {
	return _ChannelImplementation.Contract.ExitRequest(&_ChannelImplementation.CallOpts)
}

// ExitRequest is a free data retrieval call binding the contract method 0xf4b3a197.
//
// Solidity: function exitRequest() view returns(uint256 timelock, address beneficiary)
func (_ChannelImplementation *ChannelImplementationCallerSession) ExitRequest() (struct {
	Timelock    *big.Int
	Beneficiary common.Address
}, error) {
	return _ChannelImplementation.Contract.ExitRequest(&_ChannelImplementation.CallOpts)
}

// GetFundsDestination is a free data retrieval call binding the contract method 0xf58c5b6e.
//
// Solidity: function getFundsDestination() view returns(address)
func (_ChannelImplementation *ChannelImplementationCaller) GetFundsDestination(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChannelImplementation.contract.Call(opts, out, "getFundsDestination")
	return *ret0, err
}

// GetFundsDestination is a free data retrieval call binding the contract method 0xf58c5b6e.
//
// Solidity: function getFundsDestination() view returns(address)
func (_ChannelImplementation *ChannelImplementationSession) GetFundsDestination() (common.Address, error) {
	return _ChannelImplementation.Contract.GetFundsDestination(&_ChannelImplementation.CallOpts)
}

// GetFundsDestination is a free data retrieval call binding the contract method 0xf58c5b6e.
//
// Solidity: function getFundsDestination() view returns(address)
func (_ChannelImplementation *ChannelImplementationCallerSession) GetFundsDestination() (common.Address, error) {
	return _ChannelImplementation.Contract.GetFundsDestination(&_ChannelImplementation.CallOpts)
}

// Hermes is a free data retrieval call binding the contract method 0xd8092c92.
//
// Solidity: function hermes() view returns(address operator, address contractAddress, uint256 settled)
func (_ChannelImplementation *ChannelImplementationCaller) Hermes(opts *bind.CallOpts) (struct {
	Operator        common.Address
	ContractAddress common.Address
	Settled         *big.Int
}, error) {
	ret := new(struct {
		Operator        common.Address
		ContractAddress common.Address
		Settled         *big.Int
	})
	out := ret
	err := _ChannelImplementation.contract.Call(opts, out, "hermes")
	return *ret, err
}

// Hermes is a free data retrieval call binding the contract method 0xd8092c92.
//
// Solidity: function hermes() view returns(address operator, address contractAddress, uint256 settled)
func (_ChannelImplementation *ChannelImplementationSession) Hermes() (struct {
	Operator        common.Address
	ContractAddress common.Address
	Settled         *big.Int
}, error) {
	return _ChannelImplementation.Contract.Hermes(&_ChannelImplementation.CallOpts)
}

// Hermes is a free data retrieval call binding the contract method 0xd8092c92.
//
// Solidity: function hermes() view returns(address operator, address contractAddress, uint256 settled)
func (_ChannelImplementation *ChannelImplementationCallerSession) Hermes() (struct {
	Operator        common.Address
	ContractAddress common.Address
	Settled         *big.Int
}, error) {
	return _ChannelImplementation.Contract.Hermes(&_ChannelImplementation.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_ChannelImplementation *ChannelImplementationCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ChannelImplementation.contract.Call(opts, out, "isInitialized")
	return *ret0, err
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_ChannelImplementation *ChannelImplementationSession) IsInitialized() (bool, error) {
	return _ChannelImplementation.Contract.IsInitialized(&_ChannelImplementation.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_ChannelImplementation *ChannelImplementationCallerSession) IsInitialized() (bool, error) {
	return _ChannelImplementation.Contract.IsInitialized(&_ChannelImplementation.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_ChannelImplementation *ChannelImplementationCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChannelImplementation.contract.Call(opts, out, "operator")
	return *ret0, err
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_ChannelImplementation *ChannelImplementationSession) Operator() (common.Address, error) {
	return _ChannelImplementation.Contract.Operator(&_ChannelImplementation.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_ChannelImplementation *ChannelImplementationCallerSession) Operator() (common.Address, error) {
	return _ChannelImplementation.Contract.Operator(&_ChannelImplementation.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChannelImplementation *ChannelImplementationCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChannelImplementation.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChannelImplementation *ChannelImplementationSession) Owner() (common.Address, error) {
	return _ChannelImplementation.Contract.Owner(&_ChannelImplementation.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChannelImplementation *ChannelImplementationCallerSession) Owner() (common.Address, error) {
	return _ChannelImplementation.Contract.Owner(&_ChannelImplementation.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ChannelImplementation *ChannelImplementationCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChannelImplementation.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ChannelImplementation *ChannelImplementationSession) Token() (common.Address, error) {
	return _ChannelImplementation.Contract.Token(&_ChannelImplementation.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ChannelImplementation *ChannelImplementationCallerSession) Token() (common.Address, error) {
	return _ChannelImplementation.Contract.Token(&_ChannelImplementation.CallOpts)
}

// ClaimEthers is a paid mutator transaction binding the contract method 0x6931b550.
//
// Solidity: function claimEthers() returns()
func (_ChannelImplementation *ChannelImplementationTransactor) ClaimEthers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChannelImplementation.contract.Transact(opts, "claimEthers")
}

// ClaimEthers is a paid mutator transaction binding the contract method 0x6931b550.
//
// Solidity: function claimEthers() returns()
func (_ChannelImplementation *ChannelImplementationSession) ClaimEthers() (*types.Transaction, error) {
	return _ChannelImplementation.Contract.ClaimEthers(&_ChannelImplementation.TransactOpts)
}

// ClaimEthers is a paid mutator transaction binding the contract method 0x6931b550.
//
// Solidity: function claimEthers() returns()
func (_ChannelImplementation *ChannelImplementationTransactorSession) ClaimEthers() (*types.Transaction, error) {
	return _ChannelImplementation.Contract.ClaimEthers(&_ChannelImplementation.TransactOpts)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0xdf8de3e7.
//
// Solidity: function claimTokens(address _token) returns()
func (_ChannelImplementation *ChannelImplementationTransactor) ClaimTokens(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _ChannelImplementation.contract.Transact(opts, "claimTokens", _token)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0xdf8de3e7.
//
// Solidity: function claimTokens(address _token) returns()
func (_ChannelImplementation *ChannelImplementationSession) ClaimTokens(_token common.Address) (*types.Transaction, error) {
	return _ChannelImplementation.Contract.ClaimTokens(&_ChannelImplementation.TransactOpts, _token)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0xdf8de3e7.
//
// Solidity: function claimTokens(address _token) returns()
func (_ChannelImplementation *ChannelImplementationTransactorSession) ClaimTokens(_token common.Address) (*types.Transaction, error) {
	return _ChannelImplementation.Contract.ClaimTokens(&_ChannelImplementation.TransactOpts, _token)
}

// FinalizeExit is a paid mutator transaction binding the contract method 0x07e8ec1f.
//
// Solidity: function finalizeExit() returns()
func (_ChannelImplementation *ChannelImplementationTransactor) FinalizeExit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChannelImplementation.contract.Transact(opts, "finalizeExit")
}

// FinalizeExit is a paid mutator transaction binding the contract method 0x07e8ec1f.
//
// Solidity: function finalizeExit() returns()
func (_ChannelImplementation *ChannelImplementationSession) FinalizeExit() (*types.Transaction, error) {
	return _ChannelImplementation.Contract.FinalizeExit(&_ChannelImplementation.TransactOpts)
}

// FinalizeExit is a paid mutator transaction binding the contract method 0x07e8ec1f.
//
// Solidity: function finalizeExit() returns()
func (_ChannelImplementation *ChannelImplementationTransactorSession) FinalizeExit() (*types.Transaction, error) {
	return _ChannelImplementation.Contract.FinalizeExit(&_ChannelImplementation.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xf7013ef6.
//
// Solidity: function initialize(address _token, address _dex, address _identityHash, address _hermesId, uint256 _fee) returns()
func (_ChannelImplementation *ChannelImplementationTransactor) Initialize(opts *bind.TransactOpts, _token common.Address, _dex common.Address, _identityHash common.Address, _hermesId common.Address, _fee *big.Int) (*types.Transaction, error) {
	return _ChannelImplementation.contract.Transact(opts, "initialize", _token, _dex, _identityHash, _hermesId, _fee)
}

// Initialize is a paid mutator transaction binding the contract method 0xf7013ef6.
//
// Solidity: function initialize(address _token, address _dex, address _identityHash, address _hermesId, uint256 _fee) returns()
func (_ChannelImplementation *ChannelImplementationSession) Initialize(_token common.Address, _dex common.Address, _identityHash common.Address, _hermesId common.Address, _fee *big.Int) (*types.Transaction, error) {
	return _ChannelImplementation.Contract.Initialize(&_ChannelImplementation.TransactOpts, _token, _dex, _identityHash, _hermesId, _fee)
}

// Initialize is a paid mutator transaction binding the contract method 0xf7013ef6.
//
// Solidity: function initialize(address _token, address _dex, address _identityHash, address _hermesId, uint256 _fee) returns()
func (_ChannelImplementation *ChannelImplementationTransactorSession) Initialize(_token common.Address, _dex common.Address, _identityHash common.Address, _hermesId common.Address, _fee *big.Int) (*types.Transaction, error) {
	return _ChannelImplementation.Contract.Initialize(&_ChannelImplementation.TransactOpts, _token, _dex, _identityHash, _hermesId, _fee)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ChannelImplementation *ChannelImplementationTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChannelImplementation.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ChannelImplementation *ChannelImplementationSession) RenounceOwnership() (*types.Transaction, error) {
	return _ChannelImplementation.Contract.RenounceOwnership(&_ChannelImplementation.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ChannelImplementation *ChannelImplementationTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ChannelImplementation.Contract.RenounceOwnership(&_ChannelImplementation.TransactOpts)
}

// RequestExit is a paid mutator transaction binding the contract method 0x182f3488.
//
// Solidity: function requestExit(address _beneficiary, uint256 _validUntil, bytes _signature) returns()
func (_ChannelImplementation *ChannelImplementationTransactor) RequestExit(opts *bind.TransactOpts, _beneficiary common.Address, _validUntil *big.Int, _signature []byte) (*types.Transaction, error) {
	return _ChannelImplementation.contract.Transact(opts, "requestExit", _beneficiary, _validUntil, _signature)
}

// RequestExit is a paid mutator transaction binding the contract method 0x182f3488.
//
// Solidity: function requestExit(address _beneficiary, uint256 _validUntil, bytes _signature) returns()
func (_ChannelImplementation *ChannelImplementationSession) RequestExit(_beneficiary common.Address, _validUntil *big.Int, _signature []byte) (*types.Transaction, error) {
	return _ChannelImplementation.Contract.RequestExit(&_ChannelImplementation.TransactOpts, _beneficiary, _validUntil, _signature)
}

// RequestExit is a paid mutator transaction binding the contract method 0x182f3488.
//
// Solidity: function requestExit(address _beneficiary, uint256 _validUntil, bytes _signature) returns()
func (_ChannelImplementation *ChannelImplementationTransactorSession) RequestExit(_beneficiary common.Address, _validUntil *big.Int, _signature []byte) (*types.Transaction, error) {
	return _ChannelImplementation.Contract.RequestExit(&_ChannelImplementation.TransactOpts, _beneficiary, _validUntil, _signature)
}

// SetFundsDestination is a paid mutator transaction binding the contract method 0x238e130a.
//
// Solidity: function setFundsDestination(address _newDestination) returns()
func (_ChannelImplementation *ChannelImplementationTransactor) SetFundsDestination(opts *bind.TransactOpts, _newDestination common.Address) (*types.Transaction, error) {
	return _ChannelImplementation.contract.Transact(opts, "setFundsDestination", _newDestination)
}

// SetFundsDestination is a paid mutator transaction binding the contract method 0x238e130a.
//
// Solidity: function setFundsDestination(address _newDestination) returns()
func (_ChannelImplementation *ChannelImplementationSession) SetFundsDestination(_newDestination common.Address) (*types.Transaction, error) {
	return _ChannelImplementation.Contract.SetFundsDestination(&_ChannelImplementation.TransactOpts, _newDestination)
}

// SetFundsDestination is a paid mutator transaction binding the contract method 0x238e130a.
//
// Solidity: function setFundsDestination(address _newDestination) returns()
func (_ChannelImplementation *ChannelImplementationTransactorSession) SetFundsDestination(_newDestination common.Address) (*types.Transaction, error) {
	return _ChannelImplementation.Contract.SetFundsDestination(&_ChannelImplementation.TransactOpts, _newDestination)
}

// SetFundsDestinationByCheque is a paid mutator transaction binding the contract method 0x7c3e6105.
//
// Solidity: function setFundsDestinationByCheque(address _newDestination, uint256 _nonce, bytes _signature) returns()
func (_ChannelImplementation *ChannelImplementationTransactor) SetFundsDestinationByCheque(opts *bind.TransactOpts, _newDestination common.Address, _nonce *big.Int, _signature []byte) (*types.Transaction, error) {
	return _ChannelImplementation.contract.Transact(opts, "setFundsDestinationByCheque", _newDestination, _nonce, _signature)
}

// SetFundsDestinationByCheque is a paid mutator transaction binding the contract method 0x7c3e6105.
//
// Solidity: function setFundsDestinationByCheque(address _newDestination, uint256 _nonce, bytes _signature) returns()
func (_ChannelImplementation *ChannelImplementationSession) SetFundsDestinationByCheque(_newDestination common.Address, _nonce *big.Int, _signature []byte) (*types.Transaction, error) {
	return _ChannelImplementation.Contract.SetFundsDestinationByCheque(&_ChannelImplementation.TransactOpts, _newDestination, _nonce, _signature)
}

// SetFundsDestinationByCheque is a paid mutator transaction binding the contract method 0x7c3e6105.
//
// Solidity: function setFundsDestinationByCheque(address _newDestination, uint256 _nonce, bytes _signature) returns()
func (_ChannelImplementation *ChannelImplementationTransactorSession) SetFundsDestinationByCheque(_newDestination common.Address, _nonce *big.Int, _signature []byte) (*types.Transaction, error) {
	return _ChannelImplementation.Contract.SetFundsDestinationByCheque(&_ChannelImplementation.TransactOpts, _newDestination, _nonce, _signature)
}

// SettlePromise is a paid mutator transaction binding the contract method 0x6f174630.
//
// Solidity: function settlePromise(uint256 _amount, uint256 _transactorFee, bytes32 _lock, bytes _signature) returns()
func (_ChannelImplementation *ChannelImplementationTransactor) SettlePromise(opts *bind.TransactOpts, _amount *big.Int, _transactorFee *big.Int, _lock [32]byte, _signature []byte) (*types.Transaction, error) {
	return _ChannelImplementation.contract.Transact(opts, "settlePromise", _amount, _transactorFee, _lock, _signature)
}

// SettlePromise is a paid mutator transaction binding the contract method 0x6f174630.
//
// Solidity: function settlePromise(uint256 _amount, uint256 _transactorFee, bytes32 _lock, bytes _signature) returns()
func (_ChannelImplementation *ChannelImplementationSession) SettlePromise(_amount *big.Int, _transactorFee *big.Int, _lock [32]byte, _signature []byte) (*types.Transaction, error) {
	return _ChannelImplementation.Contract.SettlePromise(&_ChannelImplementation.TransactOpts, _amount, _transactorFee, _lock, _signature)
}

// SettlePromise is a paid mutator transaction binding the contract method 0x6f174630.
//
// Solidity: function settlePromise(uint256 _amount, uint256 _transactorFee, bytes32 _lock, bytes _signature) returns()
func (_ChannelImplementation *ChannelImplementationTransactorSession) SettlePromise(_amount *big.Int, _transactorFee *big.Int, _lock [32]byte, _signature []byte) (*types.Transaction, error) {
	return _ChannelImplementation.Contract.SettlePromise(&_ChannelImplementation.TransactOpts, _amount, _transactorFee, _lock, _signature)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ChannelImplementation *ChannelImplementationTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ChannelImplementation.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ChannelImplementation *ChannelImplementationSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ChannelImplementation.Contract.TransferOwnership(&_ChannelImplementation.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ChannelImplementation *ChannelImplementationTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ChannelImplementation.Contract.TransferOwnership(&_ChannelImplementation.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ChannelImplementation *ChannelImplementationTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChannelImplementation.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ChannelImplementation *ChannelImplementationSession) Receive() (*types.Transaction, error) {
	return _ChannelImplementation.Contract.Receive(&_ChannelImplementation.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ChannelImplementation *ChannelImplementationTransactorSession) Receive() (*types.Transaction, error) {
	return _ChannelImplementation.Contract.Receive(&_ChannelImplementation.TransactOpts)
}

// ChannelImplementationChannelInitialisedIterator is returned from FilterChannelInitialised and is used to iterate over the raw logs and unpacked data for ChannelInitialised events raised by the ChannelImplementation contract.
type ChannelImplementationChannelInitialisedIterator struct {
	Event *ChannelImplementationChannelInitialised // Event containing the contract specifics and raw log

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
func (it *ChannelImplementationChannelInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChannelImplementationChannelInitialised)
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
		it.Event = new(ChannelImplementationChannelInitialised)
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
func (it *ChannelImplementationChannelInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChannelImplementationChannelInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChannelImplementationChannelInitialised represents a ChannelInitialised event raised by the ChannelImplementation contract.
type ChannelImplementationChannelInitialised struct {
	Operator common.Address
	Hermes   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChannelInitialised is a free log retrieval operation binding the contract event 0x9a7def6556351196c74c99e1cc8dcd284e9da181ea854c3e6367cc9fad882a51.
//
// Solidity: event ChannelInitialised(address operator, address hermes)
func (_ChannelImplementation *ChannelImplementationFilterer) FilterChannelInitialised(opts *bind.FilterOpts) (*ChannelImplementationChannelInitialisedIterator, error) {

	logs, sub, err := _ChannelImplementation.contract.FilterLogs(opts, "ChannelInitialised")
	if err != nil {
		return nil, err
	}
	return &ChannelImplementationChannelInitialisedIterator{contract: _ChannelImplementation.contract, event: "ChannelInitialised", logs: logs, sub: sub}, nil
}

// WatchChannelInitialised is a free log subscription operation binding the contract event 0x9a7def6556351196c74c99e1cc8dcd284e9da181ea854c3e6367cc9fad882a51.
//
// Solidity: event ChannelInitialised(address operator, address hermes)
func (_ChannelImplementation *ChannelImplementationFilterer) WatchChannelInitialised(opts *bind.WatchOpts, sink chan<- *ChannelImplementationChannelInitialised) (event.Subscription, error) {

	logs, sub, err := _ChannelImplementation.contract.WatchLogs(opts, "ChannelInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChannelImplementationChannelInitialised)
				if err := _ChannelImplementation.contract.UnpackLog(event, "ChannelInitialised", log); err != nil {
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

// ParseChannelInitialised is a log parse operation binding the contract event 0x9a7def6556351196c74c99e1cc8dcd284e9da181ea854c3e6367cc9fad882a51.
//
// Solidity: event ChannelInitialised(address operator, address hermes)
func (_ChannelImplementation *ChannelImplementationFilterer) ParseChannelInitialised(log types.Log) (*ChannelImplementationChannelInitialised, error) {
	event := new(ChannelImplementationChannelInitialised)
	if err := _ChannelImplementation.contract.UnpackLog(event, "ChannelInitialised", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChannelImplementationDestinationChangedIterator is returned from FilterDestinationChanged and is used to iterate over the raw logs and unpacked data for DestinationChanged events raised by the ChannelImplementation contract.
type ChannelImplementationDestinationChangedIterator struct {
	Event *ChannelImplementationDestinationChanged // Event containing the contract specifics and raw log

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
func (it *ChannelImplementationDestinationChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChannelImplementationDestinationChanged)
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
		it.Event = new(ChannelImplementationDestinationChanged)
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
func (it *ChannelImplementationDestinationChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChannelImplementationDestinationChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChannelImplementationDestinationChanged represents a DestinationChanged event raised by the ChannelImplementation contract.
type ChannelImplementationDestinationChanged struct {
	PreviousDestination common.Address
	NewDestination      common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDestinationChanged is a free log retrieval operation binding the contract event 0xe1a66d77649cf0a57b9937073549f30f1c82bb865aaf066d2f299e37a62c6aad.
//
// Solidity: event DestinationChanged(address indexed previousDestination, address indexed newDestination)
func (_ChannelImplementation *ChannelImplementationFilterer) FilterDestinationChanged(opts *bind.FilterOpts, previousDestination []common.Address, newDestination []common.Address) (*ChannelImplementationDestinationChangedIterator, error) {

	var previousDestinationRule []interface{}
	for _, previousDestinationItem := range previousDestination {
		previousDestinationRule = append(previousDestinationRule, previousDestinationItem)
	}
	var newDestinationRule []interface{}
	for _, newDestinationItem := range newDestination {
		newDestinationRule = append(newDestinationRule, newDestinationItem)
	}

	logs, sub, err := _ChannelImplementation.contract.FilterLogs(opts, "DestinationChanged", previousDestinationRule, newDestinationRule)
	if err != nil {
		return nil, err
	}
	return &ChannelImplementationDestinationChangedIterator{contract: _ChannelImplementation.contract, event: "DestinationChanged", logs: logs, sub: sub}, nil
}

// WatchDestinationChanged is a free log subscription operation binding the contract event 0xe1a66d77649cf0a57b9937073549f30f1c82bb865aaf066d2f299e37a62c6aad.
//
// Solidity: event DestinationChanged(address indexed previousDestination, address indexed newDestination)
func (_ChannelImplementation *ChannelImplementationFilterer) WatchDestinationChanged(opts *bind.WatchOpts, sink chan<- *ChannelImplementationDestinationChanged, previousDestination []common.Address, newDestination []common.Address) (event.Subscription, error) {

	var previousDestinationRule []interface{}
	for _, previousDestinationItem := range previousDestination {
		previousDestinationRule = append(previousDestinationRule, previousDestinationItem)
	}
	var newDestinationRule []interface{}
	for _, newDestinationItem := range newDestination {
		newDestinationRule = append(newDestinationRule, newDestinationItem)
	}

	logs, sub, err := _ChannelImplementation.contract.WatchLogs(opts, "DestinationChanged", previousDestinationRule, newDestinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChannelImplementationDestinationChanged)
				if err := _ChannelImplementation.contract.UnpackLog(event, "DestinationChanged", log); err != nil {
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
func (_ChannelImplementation *ChannelImplementationFilterer) ParseDestinationChanged(log types.Log) (*ChannelImplementationDestinationChanged, error) {
	event := new(ChannelImplementationDestinationChanged)
	if err := _ChannelImplementation.contract.UnpackLog(event, "DestinationChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChannelImplementationExitRequestedIterator is returned from FilterExitRequested and is used to iterate over the raw logs and unpacked data for ExitRequested events raised by the ChannelImplementation contract.
type ChannelImplementationExitRequestedIterator struct {
	Event *ChannelImplementationExitRequested // Event containing the contract specifics and raw log

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
func (it *ChannelImplementationExitRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChannelImplementationExitRequested)
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
		it.Event = new(ChannelImplementationExitRequested)
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
func (it *ChannelImplementationExitRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChannelImplementationExitRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChannelImplementationExitRequested represents a ExitRequested event raised by the ChannelImplementation contract.
type ChannelImplementationExitRequested struct {
	Timelock *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterExitRequested is a free log retrieval operation binding the contract event 0xe60f0366d8d61555184ea027447889648bae94ebfb1202a39544b6b6803969db.
//
// Solidity: event ExitRequested(uint256 timelock)
func (_ChannelImplementation *ChannelImplementationFilterer) FilterExitRequested(opts *bind.FilterOpts) (*ChannelImplementationExitRequestedIterator, error) {

	logs, sub, err := _ChannelImplementation.contract.FilterLogs(opts, "ExitRequested")
	if err != nil {
		return nil, err
	}
	return &ChannelImplementationExitRequestedIterator{contract: _ChannelImplementation.contract, event: "ExitRequested", logs: logs, sub: sub}, nil
}

// WatchExitRequested is a free log subscription operation binding the contract event 0xe60f0366d8d61555184ea027447889648bae94ebfb1202a39544b6b6803969db.
//
// Solidity: event ExitRequested(uint256 timelock)
func (_ChannelImplementation *ChannelImplementationFilterer) WatchExitRequested(opts *bind.WatchOpts, sink chan<- *ChannelImplementationExitRequested) (event.Subscription, error) {

	logs, sub, err := _ChannelImplementation.contract.WatchLogs(opts, "ExitRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChannelImplementationExitRequested)
				if err := _ChannelImplementation.contract.UnpackLog(event, "ExitRequested", log); err != nil {
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

// ParseExitRequested is a log parse operation binding the contract event 0xe60f0366d8d61555184ea027447889648bae94ebfb1202a39544b6b6803969db.
//
// Solidity: event ExitRequested(uint256 timelock)
func (_ChannelImplementation *ChannelImplementationFilterer) ParseExitRequested(log types.Log) (*ChannelImplementationExitRequested, error) {
	event := new(ChannelImplementationExitRequested)
	if err := _ChannelImplementation.contract.UnpackLog(event, "ExitRequested", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChannelImplementationFinalizeExitIterator is returned from FilterFinalizeExit and is used to iterate over the raw logs and unpacked data for FinalizeExit events raised by the ChannelImplementation contract.
type ChannelImplementationFinalizeExitIterator struct {
	Event *ChannelImplementationFinalizeExit // Event containing the contract specifics and raw log

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
func (it *ChannelImplementationFinalizeExitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChannelImplementationFinalizeExit)
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
		it.Event = new(ChannelImplementationFinalizeExit)
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
func (it *ChannelImplementationFinalizeExitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChannelImplementationFinalizeExitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChannelImplementationFinalizeExit represents a FinalizeExit event raised by the ChannelImplementation contract.
type ChannelImplementationFinalizeExit struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFinalizeExit is a free log retrieval operation binding the contract event 0x50128f92fd19060780780085c779f5ddebca701ad03dc303be5b085986345824.
//
// Solidity: event FinalizeExit(uint256 amount)
func (_ChannelImplementation *ChannelImplementationFilterer) FilterFinalizeExit(opts *bind.FilterOpts) (*ChannelImplementationFinalizeExitIterator, error) {

	logs, sub, err := _ChannelImplementation.contract.FilterLogs(opts, "FinalizeExit")
	if err != nil {
		return nil, err
	}
	return &ChannelImplementationFinalizeExitIterator{contract: _ChannelImplementation.contract, event: "FinalizeExit", logs: logs, sub: sub}, nil
}

// WatchFinalizeExit is a free log subscription operation binding the contract event 0x50128f92fd19060780780085c779f5ddebca701ad03dc303be5b085986345824.
//
// Solidity: event FinalizeExit(uint256 amount)
func (_ChannelImplementation *ChannelImplementationFilterer) WatchFinalizeExit(opts *bind.WatchOpts, sink chan<- *ChannelImplementationFinalizeExit) (event.Subscription, error) {

	logs, sub, err := _ChannelImplementation.contract.WatchLogs(opts, "FinalizeExit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChannelImplementationFinalizeExit)
				if err := _ChannelImplementation.contract.UnpackLog(event, "FinalizeExit", log); err != nil {
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

// ParseFinalizeExit is a log parse operation binding the contract event 0x50128f92fd19060780780085c779f5ddebca701ad03dc303be5b085986345824.
//
// Solidity: event FinalizeExit(uint256 amount)
func (_ChannelImplementation *ChannelImplementationFilterer) ParseFinalizeExit(log types.Log) (*ChannelImplementationFinalizeExit, error) {
	event := new(ChannelImplementationFinalizeExit)
	if err := _ChannelImplementation.contract.UnpackLog(event, "FinalizeExit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChannelImplementationOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ChannelImplementation contract.
type ChannelImplementationOwnershipTransferredIterator struct {
	Event *ChannelImplementationOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ChannelImplementationOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChannelImplementationOwnershipTransferred)
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
		it.Event = new(ChannelImplementationOwnershipTransferred)
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
func (it *ChannelImplementationOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChannelImplementationOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChannelImplementationOwnershipTransferred represents a OwnershipTransferred event raised by the ChannelImplementation contract.
type ChannelImplementationOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ChannelImplementation *ChannelImplementationFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ChannelImplementationOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ChannelImplementation.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ChannelImplementationOwnershipTransferredIterator{contract: _ChannelImplementation.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ChannelImplementation *ChannelImplementationFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ChannelImplementationOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ChannelImplementation.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChannelImplementationOwnershipTransferred)
				if err := _ChannelImplementation.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ChannelImplementation *ChannelImplementationFilterer) ParseOwnershipTransferred(log types.Log) (*ChannelImplementationOwnershipTransferred, error) {
	event := new(ChannelImplementationOwnershipTransferred)
	if err := _ChannelImplementation.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChannelImplementationPromiseSettledIterator is returned from FilterPromiseSettled and is used to iterate over the raw logs and unpacked data for PromiseSettled events raised by the ChannelImplementation contract.
type ChannelImplementationPromiseSettledIterator struct {
	Event *ChannelImplementationPromiseSettled // Event containing the contract specifics and raw log

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
func (it *ChannelImplementationPromiseSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChannelImplementationPromiseSettled)
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
		it.Event = new(ChannelImplementationPromiseSettled)
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
func (it *ChannelImplementationPromiseSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChannelImplementationPromiseSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChannelImplementationPromiseSettled represents a PromiseSettled event raised by the ChannelImplementation contract.
type ChannelImplementationPromiseSettled struct {
	Beneficiary  common.Address
	Amount       *big.Int
	TotalSettled *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPromiseSettled is a free log retrieval operation binding the contract event 0x50c3491624aa1825a7653df63d067fecd5c8634ba63c99c4a7cf04ff1436070b.
//
// Solidity: event PromiseSettled(address beneficiary, uint256 amount, uint256 totalSettled)
func (_ChannelImplementation *ChannelImplementationFilterer) FilterPromiseSettled(opts *bind.FilterOpts) (*ChannelImplementationPromiseSettledIterator, error) {

	logs, sub, err := _ChannelImplementation.contract.FilterLogs(opts, "PromiseSettled")
	if err != nil {
		return nil, err
	}
	return &ChannelImplementationPromiseSettledIterator{contract: _ChannelImplementation.contract, event: "PromiseSettled", logs: logs, sub: sub}, nil
}

// WatchPromiseSettled is a free log subscription operation binding the contract event 0x50c3491624aa1825a7653df63d067fecd5c8634ba63c99c4a7cf04ff1436070b.
//
// Solidity: event PromiseSettled(address beneficiary, uint256 amount, uint256 totalSettled)
func (_ChannelImplementation *ChannelImplementationFilterer) WatchPromiseSettled(opts *bind.WatchOpts, sink chan<- *ChannelImplementationPromiseSettled) (event.Subscription, error) {

	logs, sub, err := _ChannelImplementation.contract.WatchLogs(opts, "PromiseSettled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChannelImplementationPromiseSettled)
				if err := _ChannelImplementation.contract.UnpackLog(event, "PromiseSettled", log); err != nil {
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

// ParsePromiseSettled is a log parse operation binding the contract event 0x50c3491624aa1825a7653df63d067fecd5c8634ba63c99c4a7cf04ff1436070b.
//
// Solidity: event PromiseSettled(address beneficiary, uint256 amount, uint256 totalSettled)
func (_ChannelImplementation *ChannelImplementationFilterer) ParsePromiseSettled(log types.Log) (*ChannelImplementationPromiseSettled, error) {
	event := new(ChannelImplementationPromiseSettled)
	if err := _ChannelImplementation.contract.UnpackLog(event, "PromiseSettled", log); err != nil {
		return nil, err
	}
	return event, nil
}
