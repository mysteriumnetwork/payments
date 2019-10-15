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

// RegistryABI is the input ABI used to generate the binding from.
const RegistryABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dexAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_channelImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_accountantImplementation\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_regFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minimalAccountantStake\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"identityHash\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"accountantId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"channelAddress\",\"type\":\"address\"}],\"name\":\"ConsumerChannelCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousDestination\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newDestination\",\"type\":\"address\"}],\"name\":\"DestinationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"accountantId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"accountantOperator\",\"type\":\"address\"}],\"name\":\"RegisteredAccountant\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"identityHash\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"accountantId\",\"type\":\"address\"}],\"name\":\"RegisteredIdentity\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"accountantImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"accountants\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimEthers\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"claimTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFundsDestination\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minimalAccountantStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registrationFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_newDestination\",\"type\":\"address\"}],\"name\":\"setFundsDestination\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_accountantId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_loanAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"registerIdentity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_accountantOperator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_accountantFee\",\"type\":\"uint16\"}],\"name\":\"registerAccountant\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identityHash\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_accountantId\",\"type\":\"address\"}],\"name\":\"getChannelAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_accountantOperator\",\"type\":\"address\"}],\"name\":\"getAccountantAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"}],\"name\":\"getProxyCode\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identityHash\",\"type\":\"address\"}],\"name\":\"isRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_accountantId\",\"type\":\"address\"}],\"name\":\"isAccountant\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_accountantId\",\"type\":\"address\"}],\"name\":\"isActiveAccountant\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newFee\",\"type\":\"uint256\"}],\"name\":\"changeRegistrationFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"transferCollectedFeeTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RegistryBin is the compiled bytecode used for deploying new contracts.
const RegistryBin = `0x60806040523480156200001157600080fd5b50604051620030d9380380620030d9833981810160405260c08110156200003757600080fd5b81019080805190602001909291908051906020019092919080519060200190929190805190602001909291908051906020019092919080519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a38160048190555080600581905550600073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1614156200017f57600080fd5b85600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff161415620001fb57600080fd5b84600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614156200027757600080fd5b83600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415620002f357600080fd5b82600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050505050612d8f806200034a6000396000f3fe6080604052600436106101665760003560e01c80638f32d59b116100d1578063e32525371161008a578063f58c5b6e11610064578063f58c5b6e14610971578063f595cfd2146109c8578063f707fb4a14610a59578063fc0c546a14610af157610166565b8063e32525371461081e578063e617aaac1461086f578063f2fde38b1461092057610166565b80638f32d59b146104f85780639a3ce27414610527578063ab8672131461057e578063c3c5a54714610648578063cf10c969146106b1578063df8de3e7146107cd57610166565b80636931b550116101235780636931b550146103b45780636e82e4e1146103cb578063715018a614610434578063817b1cd21461044b578063824b09d6146104765780638da5cb5b146104a157610166565b806314c44e09146101d45780631a3d9a59146101ff578063238e130a146102685780632a33ddbd146102b95780635005076914610322578063692058c21461035d575b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f52656a656374696e672074782077697468206574686572732073656e7400000081525060200191505060405180910390fd5b3480156101e057600080fd5b506101e9610b48565b6040518082815260200191505060405180910390f35b34801561020b57600080fd5b5061024e6004803603602081101561022257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610b4e565b604051808215151515815260200191505060405180910390f35b34801561027457600080fd5b506102b76004803603602081101561028b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610b9d565b005b3480156102c557600080fd5b50610308600480360360208110156102dc57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610d11565b604051808215151515815260200191505060405180910390f35b34801561032e57600080fd5b5061035b6004803603602081101561034557600080fd5b8101908080359060200190929190505050610d9b565b005b34801561036957600080fd5b50610372610e1f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156103c057600080fd5b506103c9610e45565b005b3480156103d757600080fd5b50610432600480360360608110156103ee57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803561ffff169060200190929190505050610f23565b005b34801561044057600080fd5b5061044961138e565b005b34801561045757600080fd5b506104606114c7565b6040518082815260200191505060405180910390f35b34801561048257600080fd5b5061048b6114cd565b6040518082815260200191505060405180910390f35b3480156104ad57600080fd5b506104b66114d3565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561050457600080fd5b5061050d6114fc565b604051808215151515815260200191505060405180910390f35b34801561053357600080fd5b5061053c611553565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561058a57600080fd5b506105cd600480360360208110156105a157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611579565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561060d5780820151818401526020810190506105f2565b50505050905090810190601f16801561063a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561065457600080fd5b506106976004803603602081101561066b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061161d565b604051808215151515815260200191505060405180910390f35b3480156106bd57600080fd5b506107cb600480360360a08110156106d457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019064010000000081111561074557600080fd5b82018360208201111561075757600080fd5b8035906020019184600183028401116401000000008311171561077957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050611673565b005b3480156107d957600080fd5b5061081c600480360360208110156107f057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506120ee565b005b34801561082a57600080fd5b5061086d6004803603602081101561084157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612396565b005b34801561087b57600080fd5b506108de6004803603604081101561089257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506125e4565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561092c57600080fd5b5061096f6004803603602081101561094357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506126c0565b005b34801561097d57600080fd5b50610986612746565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156109d457600080fd5b50610a17600480360360208110156109eb57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612770565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015610a6557600080fd5b50610aa860048036036020811015610a7c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506127d2565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b348015610afd57600080fd5b50610b06612816565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b60045481565b600080600960008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015414159050919050565b610ba56114fc565b610c17576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610c5157600080fd5b8073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fe1a66d77649cf0a57b9937073549f30f1c82bb865aaf066d2f299e37a62c6aad60405160405180910390a380600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600080600960008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506000610d8482612770565b90506000813b905060008114159350505050919050565b610da36114fc565b610e15576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b8060048190555050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415610ea157600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc3073ffffffffffffffffffffffffffffffffffffffff16319081150290604051600060405180830381858888f19350505050158015610f20573d6000803e3d6000fd5b50565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610f5d57600080fd5b600554821015610fb8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526036815260200180612cfb6036913960400191505060405180910390fd5b6000610fc384612770565b9050610fce81610d11565b15610fd857600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330866040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b1580156110b557600080fd5b505af11580156110c9573d6000803e3d6000fd5b505050506040513d60208110156110df57600080fd5b8101908080519060200190929190505050506111068360065461283c90919063ffffffff16565b60068190555060006111508573ffffffffffffffffffffffffffffffffffffffff16600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166128c4565b90508073ffffffffffffffffffffffffffffffffffffffff16637ebef529600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1687866040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018261ffff1661ffff1681526020019350505050600060405180830381600087803b15801561123757600080fd5b505af115801561124b573d6000803e3d6000fd5b5050505060405180604001604052808673ffffffffffffffffffffffffffffffffffffffff16815260200185815250600960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101559050508073ffffffffffffffffffffffffffffffffffffffff167fc9f77826eb4f5ea1c94b1b3ff214796f59d1c5f610af8a3ff52cd83e3eebf64e86604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a25050505050565b6113966114fc565b611408576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60065481565b60055481565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b606080604051806060016040528060378152602001612c7e60379139905060008360601b905060008090505b60148160ff16101561161257818160ff16601481106115c057fe5b1a60f81b838260140160ff16815181106115d657fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535080806001019150506115a5565b508192505050919050565b6000600a60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b61167c85610b4e565b6116d1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180612cd66025913960400191505060405180910390fd5b60006117b9823088888888604051602001808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b81526014018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b81526014018481526020018381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b815260140195505050505050604051602081830303815290604052805190602001206128f690919063ffffffff16565b9050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561185e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f77726f6e67207369676e6174757265000000000000000000000000000000000081525060200191505060405180910390fd5b6000611887856118798860045461283c90919063ffffffff16565b61283c90919063ffffffff16565b9050600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a082316118d1848a6125e4565b6040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561193157600080fd5b505afa158015611945573d6000803e3d6000fd5b505050506040513d602081101561195b57600080fd5b81019080805190602001909291905050508111156119c4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a815260200180612d31602a913960400191505060405180910390fd5b60008288604051602001808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b81526014018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b8152601401925050506040516020818303038152906040528051906020012090506000611a868260001c600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166128c4565b90508073ffffffffffffffffffffffffffffffffffffffff1663f7013ef6600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16878d886040518663ffffffff1660e01b8152600401808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200195505050505050600060405180830381600087803b158015611bef57600080fd5b505af1158015611c03573d6000803e3d6000fd5b505050506000881115611d4a57600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663095ea7b38a8a6040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015611cb957600080fd5b505af1158015611ccd573d6000803e3d6000fd5b505050506040513d6020811015611ce357600080fd5b8101908080519060200190929190505050611d49576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526031815260200180612c286031913960400191505060405180910390fd5b5b600073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff161415611dd0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526021815260200180612cb56021913960400191505060405180910390fd5b8873ffffffffffffffffffffffffffffffffffffffff16630a798f2485888b6040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b158015611e8b57600080fd5b505af1158015611e9f573d6000803e3d6000fd5b505050506000871115611f9257600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33896040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015611f5557600080fd5b505af1158015611f69573d6000803e3d6000fd5b505050506040513d6020811015611f7f57600080fd5b8101908080519060200190929190505050505b8873ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f2ed7bcf2ff03098102c7003d7ce2a633e4b49b8198b07de5383cdf4c0ab9228b83604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a361202c8461161d565b6120e3576001600a60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508873ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fefaf768237c22e140a862d5d375ad5c153479fac3f8bcf8b580a1651fd62c3ef60405160405180910390a35b505050505050505050565b600073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141561214a57600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156121f1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180612c596025913960400191505060405180910390fd5b60008173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561227057600080fd5b505afa158015612284573d6000803e3d6000fd5b505050506040513d602081101561229a57600080fd5b810190808051906020019092919050505090508173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561235657600080fd5b505af115801561236a573d6000803e3d6000fd5b505050506040513d602081101561238057600080fd5b8101908080519060200190929190505050505050565b61239e6114fc565b612410576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b1580156124b157600080fd5b505afa1580156124c5573d6000803e3d6000fd5b505050506040513d60208110156124db57600080fd5b81019080805190602001909291905050509050600081116124fb57600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1580156125a457600080fd5b505af11580156125b8573d6000803e3d6000fd5b505050506040513d60208110156125ce57600080fd5b8101908080519060200190929190505050505050565b600080612612600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16611579565b80519060200120905060008484604051602001808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b81526014018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b8152601401925050506040516020818303038152906040528051906020012090506126b681836129fa565b9250505092915050565b6126c86114fc565b61273a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b61274381612abd565b50565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008061279e600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16611579565b8051906020012090506127ca8373ffffffffffffffffffffffffffffffffffffffff1660001b826129fa565b915050919050565b60096020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154905082565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000808284019050838110156128ba576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b60008060606128d284611579565b9050848151602083016000f59150813b6128eb57600080fd5b819250505092915050565b6000604182511461290a57600090506129f4565b60008060006020850151925060408501519150606085015160001a90507f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08260001c111561295e57600093505050506129f4565b601b8160ff16141580156129765750601c8160ff1614155b1561298757600093505050506129f4565b60018682858560405160008152602001604052604051808581526020018460ff1660ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa1580156129e4573d6000803e3d6000fd5b5050506020604051035193505050505b92915050565b600060ff60f81b30848460405160200180857effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526001018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b81526014018381526020018281526020019450505050506040516020818303038152906040528051906020012060001c905092915050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415612b43576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526026815260200180612c026026913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573736163636f756e74616e742073686f756c642067657420617070726f76616c20746f207472616e7366657220746f6b656e736e617469766520746f6b656e2066756e64732063616e2774206265207265636f76657265643d602d80600a3d3981f3363d3d373d3d3d363d73bebebebebebebebebebebebebebebebebebebebe5af43d82803e903d91602b57fd5bf362656e65666963696172792063616e2774206265207a65726f206164647265737370726f7669646564206163636f756e74616e74206861766520746f206265206163746976656163636f756e74616e74206861766520746f207374616b65206174206c65617374206d696e696d616c207374616b6520616d6f756e746e6f7420656e6f756768742066756e647320696e206368616e6e656c20746f20636f7665722066656573a265627a7a72315820d23ffb67c4fc828911e756952da1649d81c5e4feba08633ae46bf31897cefe2164736f6c634300050c0032`

// DeployRegistry deploys a new Ethereum contract, binding an instance of Registry to it.
func DeployRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _tokenAddress common.Address, _dexAddress common.Address, _channelImplementation common.Address, _accountantImplementation common.Address, _regFee *big.Int, _minimalAccountantStake *big.Int) (common.Address, *types.Transaction, *Registry, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RegistryBin), backend, _tokenAddress, _dexAddress, _channelImplementation, _accountantImplementation, _regFee, _minimalAccountantStake)
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

// AccountantImplementation is a free data retrieval call binding the contract method 0x9a3ce274.
//
// Solidity: function accountantImplementation() constant returns(address)
func (_Registry *RegistryCaller) AccountantImplementation(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "accountantImplementation")
	return *ret0, err
}

// AccountantImplementation is a free data retrieval call binding the contract method 0x9a3ce274.
//
// Solidity: function accountantImplementation() constant returns(address)
func (_Registry *RegistrySession) AccountantImplementation() (common.Address, error) {
	return _Registry.Contract.AccountantImplementation(&_Registry.CallOpts)
}

// AccountantImplementation is a free data retrieval call binding the contract method 0x9a3ce274.
//
// Solidity: function accountantImplementation() constant returns(address)
func (_Registry *RegistryCallerSession) AccountantImplementation() (common.Address, error) {
	return _Registry.Contract.AccountantImplementation(&_Registry.CallOpts)
}

// Accountants is a free data retrieval call binding the contract method 0xf707fb4a.
//
// Solidity: function accountants(address ) constant returns(address operator, uint256 stake)
func (_Registry *RegistryCaller) Accountants(opts *bind.CallOpts, arg0 common.Address) (struct {
	Operator common.Address
	Stake    *big.Int
}, error) {
	ret := new(struct {
		Operator common.Address
		Stake    *big.Int
	})
	out := ret
	err := _Registry.contract.Call(opts, out, "accountants", arg0)
	return *ret, err
}

// Accountants is a free data retrieval call binding the contract method 0xf707fb4a.
//
// Solidity: function accountants(address ) constant returns(address operator, uint256 stake)
func (_Registry *RegistrySession) Accountants(arg0 common.Address) (struct {
	Operator common.Address
	Stake    *big.Int
}, error) {
	return _Registry.Contract.Accountants(&_Registry.CallOpts, arg0)
}

// Accountants is a free data retrieval call binding the contract method 0xf707fb4a.
//
// Solidity: function accountants(address ) constant returns(address operator, uint256 stake)
func (_Registry *RegistryCallerSession) Accountants(arg0 common.Address) (struct {
	Operator common.Address
	Stake    *big.Int
}, error) {
	return _Registry.Contract.Accountants(&_Registry.CallOpts, arg0)
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

// GetChannelAddress is a free data retrieval call binding the contract method 0xe617aaac.
//
// Solidity: function getChannelAddress(address _identityHash, address _accountantId) constant returns(address)
func (_Registry *RegistryCaller) GetChannelAddress(opts *bind.CallOpts, _identityHash common.Address, _accountantId common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "getChannelAddress", _identityHash, _accountantId)
	return *ret0, err
}

// GetChannelAddress is a free data retrieval call binding the contract method 0xe617aaac.
//
// Solidity: function getChannelAddress(address _identityHash, address _accountantId) constant returns(address)
func (_Registry *RegistrySession) GetChannelAddress(_identityHash common.Address, _accountantId common.Address) (common.Address, error) {
	return _Registry.Contract.GetChannelAddress(&_Registry.CallOpts, _identityHash, _accountantId)
}

// GetChannelAddress is a free data retrieval call binding the contract method 0xe617aaac.
//
// Solidity: function getChannelAddress(address _identityHash, address _accountantId) constant returns(address)
func (_Registry *RegistryCallerSession) GetChannelAddress(_identityHash common.Address, _accountantId common.Address) (common.Address, error) {
	return _Registry.Contract.GetChannelAddress(&_Registry.CallOpts, _identityHash, _accountantId)
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

// IsActiveAccountant is a free data retrieval call binding the contract method 0x1a3d9a59.
//
// Solidity: function isActiveAccountant(address _accountantId) constant returns(bool)
func (_Registry *RegistryCaller) IsActiveAccountant(opts *bind.CallOpts, _accountantId common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "isActiveAccountant", _accountantId)
	return *ret0, err
}

// IsActiveAccountant is a free data retrieval call binding the contract method 0x1a3d9a59.
//
// Solidity: function isActiveAccountant(address _accountantId) constant returns(bool)
func (_Registry *RegistrySession) IsActiveAccountant(_accountantId common.Address) (bool, error) {
	return _Registry.Contract.IsActiveAccountant(&_Registry.CallOpts, _accountantId)
}

// IsActiveAccountant is a free data retrieval call binding the contract method 0x1a3d9a59.
//
// Solidity: function isActiveAccountant(address _accountantId) constant returns(bool)
func (_Registry *RegistryCallerSession) IsActiveAccountant(_accountantId common.Address) (bool, error) {
	return _Registry.Contract.IsActiveAccountant(&_Registry.CallOpts, _accountantId)
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

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() constant returns(uint256)
func (_Registry *RegistryCaller) TotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "totalStaked")
	return *ret0, err
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() constant returns(uint256)
func (_Registry *RegistrySession) TotalStaked() (*big.Int, error) {
	return _Registry.Contract.TotalStaked(&_Registry.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() constant returns(uint256)
func (_Registry *RegistryCallerSession) TotalStaked() (*big.Int, error) {
	return _Registry.Contract.TotalStaked(&_Registry.CallOpts)
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

// RegisterAccountant is a paid mutator transaction binding the contract method 0x6e82e4e1.
//
// Solidity: function registerAccountant(address _accountantOperator, uint256 _stakeAmount, uint16 _accountantFee) returns()
func (_Registry *RegistryTransactor) RegisterAccountant(opts *bind.TransactOpts, _accountantOperator common.Address, _stakeAmount *big.Int, _accountantFee uint16) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "registerAccountant", _accountantOperator, _stakeAmount, _accountantFee)
}

// RegisterAccountant is a paid mutator transaction binding the contract method 0x6e82e4e1.
//
// Solidity: function registerAccountant(address _accountantOperator, uint256 _stakeAmount, uint16 _accountantFee) returns()
func (_Registry *RegistrySession) RegisterAccountant(_accountantOperator common.Address, _stakeAmount *big.Int, _accountantFee uint16) (*types.Transaction, error) {
	return _Registry.Contract.RegisterAccountant(&_Registry.TransactOpts, _accountantOperator, _stakeAmount, _accountantFee)
}

// RegisterAccountant is a paid mutator transaction binding the contract method 0x6e82e4e1.
//
// Solidity: function registerAccountant(address _accountantOperator, uint256 _stakeAmount, uint16 _accountantFee) returns()
func (_Registry *RegistryTransactorSession) RegisterAccountant(_accountantOperator common.Address, _stakeAmount *big.Int, _accountantFee uint16) (*types.Transaction, error) {
	return _Registry.Contract.RegisterAccountant(&_Registry.TransactOpts, _accountantOperator, _stakeAmount, _accountantFee)
}

// RegisterIdentity is a paid mutator transaction binding the contract method 0xcf10c969.
//
// Solidity: function registerIdentity(address _accountantId, uint256 _loanAmount, uint256 _fee, address _beneficiary, bytes _signature) returns()
func (_Registry *RegistryTransactor) RegisterIdentity(opts *bind.TransactOpts, _accountantId common.Address, _loanAmount *big.Int, _fee *big.Int, _beneficiary common.Address, _signature []byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "registerIdentity", _accountantId, _loanAmount, _fee, _beneficiary, _signature)
}

// RegisterIdentity is a paid mutator transaction binding the contract method 0xcf10c969.
//
// Solidity: function registerIdentity(address _accountantId, uint256 _loanAmount, uint256 _fee, address _beneficiary, bytes _signature) returns()
func (_Registry *RegistrySession) RegisterIdentity(_accountantId common.Address, _loanAmount *big.Int, _fee *big.Int, _beneficiary common.Address, _signature []byte) (*types.Transaction, error) {
	return _Registry.Contract.RegisterIdentity(&_Registry.TransactOpts, _accountantId, _loanAmount, _fee, _beneficiary, _signature)
}

// RegisterIdentity is a paid mutator transaction binding the contract method 0xcf10c969.
//
// Solidity: function registerIdentity(address _accountantId, uint256 _loanAmount, uint256 _fee, address _beneficiary, bytes _signature) returns()
func (_Registry *RegistryTransactorSession) RegisterIdentity(_accountantId common.Address, _loanAmount *big.Int, _fee *big.Int, _beneficiary common.Address, _signature []byte) (*types.Transaction, error) {
	return _Registry.Contract.RegisterIdentity(&_Registry.TransactOpts, _accountantId, _loanAmount, _fee, _beneficiary, _signature)
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
