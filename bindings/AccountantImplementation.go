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

// AccountantImplementationABI is the input ABI used to generate the binding from.
const AccountantImplementationABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"AccountantClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newFee\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"validFromBlock\",\"type\":\"uint64\"}],\"name\":\"AccountantFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activationBlock\",\"type\":\"uint256\"}],\"name\":\"AccountantPunishmentActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"AccountantPunishmentDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newStake\",\"type\":\"uint256\"}],\"name\":\"AccountantStakeIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"}],\"name\":\"ChannelBalanceDecreaseRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"ChannelBalanceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newBeneficiary\",\"type\":\"address\"}],\"name\":\"ChannelBeneficiaryChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"initialBalance\",\"type\":\"uint256\"}],\"name\":\"ChannelOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ChannelOpeningActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ChannelOpeningPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousDestination\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newDestination\",\"type\":\"address\"}],\"name\":\"DestinationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"FundsWithdrawned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newMaxLoan\",\"type\":\"uint256\"}],\"name\":\"MaxLoanValueUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"NewAccountantOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loanAmount\",\"type\":\"uint256\"}],\"name\":\"NewLoan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalSettled\",\"type\":\"uint256\"}],\"name\":\"PromiseSettled\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"channels\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settled\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"loan\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUsedNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timelock\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimEthers\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"claimTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFundsDestination\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastFee\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"value\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"validFrom\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"previousFee\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"value\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"validFrom\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"punishment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"activationBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOperator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_party\",\"type\":\"address\"}],\"name\":\"getChannelId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStatus\",\"outputs\":[{\"internalType\":\"enumAccountantImplementation.Status\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_fee\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_maxLoan\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_party\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountToLend\",\"type\":\"uint256\"}],\"name\":\"openChannel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_transactorFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_lock\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"settlePromise\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_newBalance\",\"type\":\"uint256\"}],\"name\":\"updateChannelBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"}],\"name\":\"rebalanceChannel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"increaseLoan\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"decreaseLoan\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"resolveEmergency\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_party\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newBeneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"setBeneficiary\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOperator\",\"type\":\"address\"}],\"name\":\"setAccountantOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newMaxLoan\",\"type\":\"uint256\"}],\"name\":\"setMaxLoan\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_newFee\",\"type\":\"uint16\"}],\"name\":\"setAccountantFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"getAccountantFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_additionalStake\",\"type\":\"uint256\"}],\"name\":\"increaseAccountantStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"}],\"name\":\"isOpened\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isAccountantActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pauseChannelOpening\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"activateChannelOpening\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"availableBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minimalExpectedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"closeAccountant\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_newDestination\",\"type\":\"address\"}],\"name\":\"setFundsDestination\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AccountantImplementationBin is the compiled bytecode used for deploying new contracts.
var AccountantImplementationBin = "0x6080604052336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3615b83806100cf6000396000f3fe608060405234801561001057600080fd5b50600436106102535760003560e01c80638f32d59b11610146578063e7f43c68116100c3578063f58c5b6e11610087578063f58c5b6e14610bfd578063f7d7636914610c47578063fbb46b9814610c51578063fc0c546a14610c5b578063fc0e3d9014610ca5578063fec8157d14610cc357610253565b8063e7f43c6814610a9b578063eb295b2714610ae5578063efde05ec14610b3d578063f2fde38b14610b6b578063f3fef3a314610baf57610253565b8063ab2f0e511161010a578063ab2f0e5114610994578063be02c06c146109b2578063d0acce61146109d4578063df8de3e714610a16578063e1c6648714610a5a57610253565b80638f32d59b1461082657806394c7915d146108485780639801134e14610866578063a58b2b71146108a7578063aa606dee1461098a57610253565b806354ded42d116101d4578063715018a611610198578063715018a6146106c55780637a7ebd7b146106cf5780637c2be0a314610760578063800d6afb146107985780638da5cb5b146107dc57610253565b806354ded42d1461051a5780635ab1bd53146105f35780635f5794f11461063d5780636931b550146106835780636e17b0d71461068d57610253565b80631822af6f1161021b5780631822af6f1461042c578063238e130a1461045a578063392e53cd1461049e578063456e182f146104c05780634e69d560146104ee57610253565b80630684cd20146102585780630996fcbc1461027d5780630a798f24146102af5780630b3834ea1461031d57806315c73afd14610422575b600080fd5b610260610d3f565b604051808381526020018281526020019250505060405180910390f35b6102ad6004803603602081101561029357600080fd5b81019080803561ffff169060200190929190505050610d51565b005b61031b600480360360608110156102c557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506110fe565b005b6104206004803603608081101561033357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561039a57600080fd5b8201836020820111156103ac57600080fd5b803590602001918460018302840111640100000000831117156103ce57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050611654565b005b61042a6119c9565b005b6104586004803603602081101561044257600080fd5b8101908080359060200190929190505050611d94565b005b61049c6004803603602081101561047057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611fd9565b005b6104a6612179565b604051808215151515815260200191505060405180910390f35b6104ec600480360360208110156104d657600080fd5b81019080803590602001909291905050506121d2565b005b6104f6612333565b6040518082600381111561050657fe5b60ff16815260200191505060405180910390f35b6105f16004803603608081101561053057600080fd5b810190808035906020019092919080359060200190929190803590602001909291908035906020019064010000000081111561056b57600080fd5b82018360208201111561057d57600080fd5b8035906020019184600183028401116401000000008311171561059f57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929050505061234a565b005b6105fb612a08565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6106696004803603602081101561065357600080fd5b8101908080359060200190929190505050612a32565b604051808215151515815260200191505060405180910390f35b61068b612aa1565b005b6106c3600480360360408110156106a357600080fd5b810190808035906020019092919080359060200190929190505050612b7f565b005b6106cd612f78565b005b6106fb600480360360208110156106e557600080fd5b81019080803590602001909291905050506130b1565b604051808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001868152602001858152602001848152602001838152602001828152602001965050505050505060405180910390f35b6107966004803603604081101561077657600080fd5b81019080803590602001909291908035906020019092919050505061310d565b005b6107da600480360360208110156107ae57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506134c4565b005b6107e46136b4565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61082e6136dd565b604051808215151515815260200191505060405180910390f35b610850613734565b6040518082815260200191505060405180910390f35b61086e613760565b604051808361ffff1661ffff1681526020018267ffffffffffffffff1667ffffffffffffffff1681526020019250505060405180910390f35b610988600480360360a08110156108bd57600080fd5b81019080803590602001909291908035906020019092919080359060200190929190803590602001909291908035906020019064010000000081111561090257600080fd5b82018360208201111561091457600080fd5b8035906020019184600183028401116401000000008311171561093657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050613794565b005b610992613d48565b005b61099c613eb6565b6040518082815260200191505060405180910390f35b6109ba6140c9565b604051808215151515815260200191505060405180910390f35b610a00600480360360208110156109ea57600080fd5b8101908080359060200190929190505050614119565b6040518082815260200191505060405180910390f35b610a5860048036036020811015610a2c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506141e8565b005b610a62614490565b604051808361ffff1661ffff1681526020018267ffffffffffffffff1667ffffffffffffffff1681526020019250505060405180910390f35b610aa36144c4565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610b2760048036036020811015610afb57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506144ee565b6040518082815260200191505060405180910390f35b610b6960048036036020811015610b5357600080fd5b8101908080359060200190929190505050614584565b005b610bad60048036036020811015610b8157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050614890565b005b610bfb60048036036040811015610bc557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050614916565b005b610c05614bea565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610c4f614c14565b005b610c59614d91565b005b610c63614eff565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610cad614f25565b6040518082815260200191505060405180910390f35b610d3d60048036036080811015610cd957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803561ffff16906020019092919080359060200190929190505050614f2f565b005b600d8060000154908060010154905082565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610df7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001806158d56024913960400191505060405180910390fd5b600380811115610e0357fe5b610e0b612333565b6003811115610e1657fe5b1415610e8a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f6163636f756e74616e742073686f756c64206265206e6f7420636c6f7365640081525060200191505060405180910390fd5b6113888161ffff161115610f06576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f6665652063616e2774206265206269676765722074686174203530250000000081525060200191505060405180910390fd5b600a60000160029054906101000a900467ffffffffffffffff1667ffffffffffffffff16431015610f9f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f63616e27742075706461746520696e616374697665206665650000000000000081525060200191505060405180910390fd5b6000610fa9615379565b9050600a600b6000820160009054906101000a900461ffff168160000160006101000a81548161ffff021916908361ffff1602179055506000820160029054906101000a900467ffffffffffffffff168160000160026101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555090505060405180604001604052808361ffff1681526020018267ffffffffffffffff16815250600a60008201518160000160006101000a81548161ffff021916908361ffff16021790555060208201518160000160026101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055509050507e8b9bffa5c48d8c0b96ce879f8491c5605cc3d7a46a46711b522dbe6d4070ee8282604051808361ffff1661ffff1681526020018267ffffffffffffffff1667ffffffffffffffff1681526020019250505060405180910390a15050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146111c1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f6f6e6c792072656769737472792063616e206f70656e206368616e6e656c730081525060200191505060405180910390fd5b600060038111156111ce57fe5b6111d6612333565b60038111156111e157fe5b14611237576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180615ab46025913960400191505060405180910390fd5b60008330604051602001808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b81526014018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b8152601401925050506040516020818303038152906040528051906020012090506112d181612a32565b15611327576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001806158b46021913960400191505060405180910390fd5b82600c600083815260200190815260200160002060000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600c600083815260200190815260200160002060010181905550600082111561160f576007548211156113fb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526035815260200180615a7f6035913960400191505060405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330856040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b1580156114d857600080fd5b505af11580156114ec573d6000803e3d6000fd5b505050506040513d602081101561150257600080fd5b8101908080519060200190929190505050611585576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f746f6b656e207472616e736665722073686f756c64207375636365656400000081525060200191505060405180910390fd5b61159a8260055461538590919063ffffffff16565b60058190555081600c6000838152602001908152602001600020600301819055506115d08260065461538590919063ffffffff16565b600681905550807f9839fd1896801c6864456fe7cbd098b1e4a81dba19536764cea53a1fb07ed239836040518082815260200191505060405180910390a25b7fbe2e1f3a6197dfd16fa6830c4870364b618b8b288c21cbcfa4fdb5d7c6a5e45b8183604051808381526020018281526020019250505060405180910390a150505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156116da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001806159936021913960400191505060405180910390fd5b60006116e5856144ee565b90506000600c6000838152602001908152602001600020905061170782612a32565b611779576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f6368616e6e656c206861766520746f206265206f70656e65640000000000000081525060200191505060405180910390fd5b8573ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146119135780600401548411611808576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260298152602001806159446029913960400191505060405180910390fd5b838160040181905550600061188b84848888604051602001808481526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b815260140182815260200193505050506040516020818303038152906040528051906020012061540d90919063ffffffff16565b90508673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614611911576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180615a5d6022913960400191505060405180910390fd5b505b848160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f8756aa559142225f918d7584303ecfe48e75b454f6614d0fae9f0d6ca0a898cc8286604051808381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a1505050505050565b600260038111156119d657fe5b6119de612333565b60038111156119e957fe5b14611a3f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526029815260200180615a0c6029913960400191505060405180910390fd5b6000611a766004611a68612710611a5a600554612710615511565b61552c90919063ffffffff16565b6155bb90919063ffffffff16565b90506000600d60000154430390506000611ab06001610101611a9a85610101615511565b81611aa157fe5b0461564190919063ffffffff16565b90506000611ac784836155bb90919063ffffffff16565b9050611ae181600d6001015461538590919063ffffffff16565b600d600101819055506000611b1c611b00600854600d600101546156ca565b611b0e6005546006546156ca565b61538590919063ffffffff16565b90506000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b158015611bbf57600080fd5b505afa158015611bd3573d6000803e3d6000fd5b505050506040513d6020811015611be957600080fd5b810190808051906020019092919050505090506000828210611c0c576000611c20565b611c1f828461564190919063ffffffff16565b5b9050600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330846040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b158015611cff57600080fd5b505af1158015611d13573d6000803e3d6000fd5b505050506040513d6020811015611d2957600080fd5b8101908080519060200190929190505050506000600960006101000a81548160ff02191690836003811115611d5a57fe5b02179055507f58ef313a2eb2567f3b143ff20930622dd67a0de84902cc93b7ddddd72b7773ef60405160405180910390a150505050505050565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611e3a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001806158d56024913960400191505060405180910390fd5b80611e43613eb6565b1015611f82576000611e65611e56613eb6565b8361564190919063ffffffff16565b9050600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330846040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b158015611f4457600080fd5b505af1158015611f58573d6000803e3d6000fd5b505050506040513d6020811015611f6e57600080fd5b810190808051906020019092919050505050505b611f978160085461538590919063ffffffff16565b6008819055507f41a5bb80f9c1243f3d450690277c955ff8982168e34ed096457afdc31cefef7f6008546040518082815260200191505060405180910390a150565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461207f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001806158d56024913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156120b957600080fd5b8073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fe1a66d77649cf0a57b9937073549f30f1c82bb865aaf066d2f299e37a62c6aad60405160405180910390a380600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008073ffffffffffffffffffffffffffffffffffffffff16600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415905090565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612278576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001806158d56024913960400191505060405180910390fd5b6122806140c9565b6122f2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f6163636f756e74616e74206861766520746f206265206163746976650000000081525060200191505060405180910390fd5b806007819055507f2181d8ed90eadf541579998852434d883f30ace513163bbe0e4115b29afb517a816040518082815260200191505060405180910390a150565b6000600960009054906101000a900460ff16905090565b600061241d826040518060400160405280601381526020017f4c6f61642072657475726e2072657175657374000000000000000000000000008152508787876040516020018085805190602001908083835b602083106123bf578051825260208201915060208101905060208303925061239c565b6001836020036101000a0380198251168184511680821785525050505050509050018481526020018381526020018281526020019450505050506040516020818303038152906040528051906020012061540d90919063ffffffff16565b905084612429826144ee565b1461247f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180615a5d6022913960400191505060405180910390fd5b61248885612a32565b6124fa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f6368616e6e656c206861766520746f206265206f70656e65640000000000000081525060200191505060405180910390fd5b6000600c600087815260200190815260200160002090508060040154841161256d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260298152602001806159446029913960400191505060405180910390fd5b83816004018190555080600301548511156125f0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f63616e2774207769746864726177206d6f7265207468616e206c656e6465640081525060200191505060405180910390fd5b60006126008260010154876156e3565b9050600061261e82612610613734565b61564190919063ffffffff16565b90506000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b1580156126c157600080fd5b505afa1580156126d5573d6000803e3d6000fd5b505050506040513d60208110156126eb57600080fd5b810190808051906020019092919050505090508088118061271d57508161271b898361564190919063ffffffff16565b105b156127ab5761272a6140c9565b15612795576002600960006101000a81548160ff0219169083600381111561274e57fe5b021790555043600d600001819055507fb3e91d0895882cef621b468a8235043537ca2e4d8d91ee6587801041054107e1436040518082815260200191505060405180910390a15b6127a8828261564190919063ffffffff16565b97505b60006127c489866003015461564190919063ffffffff16565b9050600754811115612821576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526035815260200180615a7f6035913960400191505060405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8660000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168b6040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1580156128ee57600080fd5b505af1158015612902573d6000803e3d6000fd5b505050506040513d602081101561291857600080fd5b81019080805190602001909291905050505080856003018190555061294a84866001015461564190919063ffffffff16565b85600101819055506129678460055461564190919063ffffffff16565b6005819055506129828960065461564190919063ffffffff16565b600681905550897f2eb87c52e5d2d7395da77618aa1afe310b9cf13fa9f70e3199d100adda3a7aa586600101546040518082815260200191505060405180910390a2897f9839fd1896801c6864456fe7cbd098b1e4a81dba19536764cea53a1fb07ed239826040518082815260200191505060405180910390a250505050505050505050565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008073ffffffffffffffffffffffffffffffffffffffff16600c600084815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614159050919050565b600073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415612afd57600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc3073ffffffffffffffffffffffffffffffffffffffff16319081150290604051600060405180830381858888f19350505050158015612b7c573d6000803e3d6000fd5b50565b612b8882612a32565b612bfa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f6368616e6e656c206861766520746f206265206f70656e65640000000000000081525060200191505060405180910390fd5b600380811115612c0657fe5b612c0e612333565b6003811115612c1957fe5b1415612c8d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f6163636f756e74616e742073686f756c64206265206e6f7420636c6f7365640081525060200191505060405180910390fd5b6000600c600084815260200190815260200160002090506000612cbd83836003015461538590919063ffffffff16565b9050600754811115612d1a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526035815260200180615a7f6035913960400191505060405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330866040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b158015612df757600080fd5b505af1158015612e0b573d6000803e3d6000fd5b505050506040513d6020811015612e2157600080fd5b8101908080519060200190929190505050612ea4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f7472616e73666572206861766520746f206265207375636365737366756c6c0081525060200191505060405180910390fd5b612ecf612ebe83600101548361564190919063ffffffff16565b60055461538590919063ffffffff16565b600581905550612eea8360065461538590919063ffffffff16565b600681905550808260010181905550808260030181905550837f2eb87c52e5d2d7395da77618aa1afe310b9cf13fa9f70e3199d100adda3a7aa5826040518082815260200191505060405180910390a2837f9839fd1896801c6864456fe7cbd098b1e4a81dba19536764cea53a1fb07ed239846040518082815260200191505060405180910390a250505050565b612f806136dd565b612ff2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b600c6020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020154908060030154908060040154908060050154905086565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146131b3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001806158d56024913960400191505060405180910390fd5b6131bb6140c9565b61322d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f6163636f756e74616e74206861766520746f206265206163746976650000000081525060200191505060405180910390fd5b61323682612a32565b6132a8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f6368616e6e656c206861766520746f206265206f70656e65640000000000000081525060200191505060405180910390fd5b600c600083815260200190815260200160002060030154811015613317576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602681526020018061596d6026913960400191505060405180910390fd5b6000600c60008481526020019081526020016000209050600081600101548311156133d55761335382600101548461564190919063ffffffff16565b90508061335e613eb6565b10156133b5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806158696022913960400191505060405180910390fd5b6133ca8160055461538590919063ffffffff16565b600581905550613478565b600082600501541415613426576133ea615379565b8260050181905550837faf4c616dc7856b81dbc1346e5547f0a1d4f1553011653f920d1041f21540107560405160405180910390a250506134c0565b81600501544310156134395750506134c0565b61345083836001015461564190919063ffffffff16565b90506134678160055461564190919063ffffffff16565b600581905550600082600501819055505b828260010181905550837f2eb87c52e5d2d7395da77618aa1afe310b9cf13fa9f70e3199d100adda3a7aa583600101546040518082815260200191505060405180910390a250505b5050565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461356a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001806158d56024913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561360d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f63616e2774206265207a65726f2061646472657373000000000000000000000081525060200191505060405180910390fd5b80600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fa326c787c51b80574c7b572d0c9664e64f1107538b902f519a896901b413791881604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b600061375b60055461374d600854600d600101546156ca565b61538590919063ffffffff16565b905090565b600a8060000160009054906101000a900461ffff16908060000160029054906101000a900467ffffffffffffffff16905082565b6000600c60008781526020019081526020016000209050600073ffffffffffffffffffffffffffffffffffffffff168160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415613872576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f6368616e6e656c2073686f756c6420657869737400000000000000000000000081525060200191505060405180910390fd5b6000836040516020018082815260200191505060405160208183030381529060405280519060200120905060006138f08489898986604051602001808581526020018481526020018381526020018281526020019450505050506040516020818303038152906040528051906020012061540d90919063ffffffff16565b9050600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614613998576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180615b056025913960400191505060405180910390fd5b60006139b184600201548961564190919063ffffffff16565b905060008111613a0c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260378152602001806159d56037913960400191505060405180910390fd5b60008460010154905080821115613a21578091505b613a3882866002015461538590919063ffffffff16565b85600201819055506000613a4b83614119565b9050600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8760000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16613ad584613ac78e8961564190919063ffffffff16565b61564190919063ffffffff16565b6040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015613b3e57600080fd5b505af1158015613b52573d6000803e3d6000fd5b505050506040513d6020811015613b6857600080fd5b810190808051906020019092919050505050613b8d838361564190919063ffffffff16565b8660010181905550613baa8360055461564190919063ffffffff16565b6005819055506000891115613c9f57600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb338b6040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015613c6257600080fd5b505af1158015613c76573d6000803e3d6000fd5b505050506040513d6020811015613c8c57600080fd5b8101908080519060200190929190505050505b8a7fa5a1f05785a942c5f624cee545c68394881a83bcaf21a83f4d76a9e8240a56688760000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16858960020154604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001838152602001828152602001935050505060405180910390a25050505050505050505050565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614613dee576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001806158d56024913960400191505060405180910390fd5b60006003811115613dfb57fe5b613e03612333565b6003811115613e0e57fe5b14613e64576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180615ab46025913960400191505060405180910390fd5b6001600960006101000a81548160ff02191690836003811115613e8357fe5b02179055507f1f4cd5d6edef8a0c4dbe6d547fdc42e0f3575167257553271f2366f9d497f67e60405160405180910390a1565b600080613ee9613ecd600854600d600101546156ca565b613edb6005546006546156ca565b61538590919063ffffffff16565b9050600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b158015613f8a57600080fd5b505afa158015613f9e573d6000803e3d6000fd5b505050506040513d6020811015613fb457600080fd5b8101908080519060200190929190505050811115613fd65760009150506140c6565b6140c281600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561407957600080fd5b505afa15801561408d573d6000803e3d6000fd5b505050506040513d60208110156140a357600080fd5b810190808051906020019092919050505061564190919063ffffffff16565b9150505b90565b6000806140d4612333565b9050600260038111156140e357fe5b8160038111156140ef57fe5b14158015614113575060038081111561410457fe5b81600381111561411057fe5b14155b91505090565b6000614123615840565b600a60000160029054906101000a900467ffffffffffffffff1667ffffffffffffffff1643101561415557600b614158565b600a5b6040518060400160405290816000820160009054906101000a900461ffff1661ffff1661ffff1681526020016000820160029054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681525050905060646141d86064836000015161ffff168602816141d057fe5b046064615511565b816141df57fe5b04915050919050565b600073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141561424457600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156142eb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602581526020018061591f6025913960400191505060405180910390fd5b60008173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561436a57600080fd5b505afa15801561437e573d6000803e3d6000fd5b505050506040513d602081101561439457600080fd5b810190808051906020019092919050505090508173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561445057600080fd5b505af1158015614464573d6000803e3d6000fd5b505050506040513d602081101561447a57600080fd5b8101908080519060200190929190505050505050565b600b8060000160009054906101000a900461ffff16908060000160029054906101000a900467ffffffffffffffff16905082565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008130604051602001808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b81526014018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b815260140192505050604051602081830303815290604052805190602001209050919050565b61458c6140c9565b6145fe576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f6163636f756e74616e74206861766520746f206265206163746976650000000081525060200191505060405180910390fd5b6000600c600083815260200190815260200160002090508060010154816003015411614675576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602981526020018061588b6029913960400191505060405180910390fd5b60006146928260010154836003015461564190919063ffffffff16565b905060006146b0826146a2613734565b61538590919063ffffffff16565b90506000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561475357600080fd5b505afa158015614767573d6000803e3d6000fd5b505050506040513d602081101561477d57600080fd5b8101908080519060200190929190505050905081811015614813576002600960006101000a81548160ff021916908360038111156147b757fe5b021790555043600d600001819055506147d9818361564190919063ffffffff16565b92507fb3e91d0895882cef621b468a8235043537ca2e4d8d91ee6587801041054107e1436040518082815260200191505060405180910390a15b6148288360055461538590919063ffffffff16565b60058190555061484583856001015461538590919063ffffffff16565b8460010181905550847f2eb87c52e5d2d7395da77618aa1afe310b9cf13fa9f70e3199d100adda3a7aa585600101546040518082815260200191505060405180910390a25050505050565b6148986136dd565b61490a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b614913816156fc565b50565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146149bc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001806158d56024913960400191505060405180910390fd5b6149c46140c9565b614a36576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f6163636f756e74616e74206861766520746f206265206163746976650000000081525060200191505060405180910390fd5b80614a3f613eb6565b1015614a96576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602c815260200180615ad9602c913960400191505060405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015614b3f57600080fd5b505af1158015614b53573d6000803e3d6000fd5b505050506040513d6020811015614b6957600080fd5b8101908080519060200190929190505050507fa2e147ce2b7cb83d9c07e397bb806f23dd42c42e86ea45e1611d6e50eb1ec8bf8183604051808381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a15050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614614cba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001806158d56024913960400191505060405180910390fd5b614cc26140c9565b614d34576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f6163636f756e74616e742073686f756c6420626520616374697665000000000081525060200191505060405180910390fd5b6003600960006101000a81548160ff02191690836003811115614d5357fe5b02179055507f888906f0892e56365e679111a6f8ba7d0742bae70d0a532641cbf0da77d5af92436040518082815260200191505060405180910390a1565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614614e37576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001806158d56024913960400191505060405180910390fd5b60016003811115614e4457fe5b614e4c612333565b6003811115614e5757fe5b14614ead576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180615b2a6025913960400191505060405180910390fd5b6000600960006101000a81548160ff02191690836003811115614ecc57fe5b02179055507f2d8b6ec230798e206d536342a28b7b61cc8fcfafb1d27c11c5519b3c42eb7df860405160405180910390a1565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600854905090565b614f37612179565b15614faa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f6861766520746f206265206e6f7420696e697469616c697a656400000000000081525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141561504d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f6f70657261746f72206861766520746f2062652073657400000000000000000081525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614156150d3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526028815260200180615a356028913960400191505060405180910390fd5b6113888261ffff16111561514f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f6665652063616e277420626520626967676572207468616e203530250000000081525060200191505060405180910390fd5b83600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555033600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060405180604001604052808361ffff1681526020014367ffffffffffffffff16815250600a60008201518160000160006101000a81548161ffff021916908361ffff16021790555060208201518160000160026101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555090505080600781905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561533257600080fd5b505afa158015615346573d6000803e3d6000fd5b505050506040513d602081101561535c57600080fd5b810190808051906020019092919050505060088190555050505050565b60006146504301905090565b600080828401905083811015615403576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b60006041825114615421576000905061550b565b60008060006020850151925060408501519150606085015160001a90507f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08260001c1115615475576000935050505061550b565b601b8160ff161415801561548d5750601c8160ff1614155b1561549e576000935050505061550b565b60018682858560405160008152602001604052604051808581526020018460ff1660ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa1580156154fb573d6000803e3d6000fd5b5050506020604051035193505050505b92915050565b600081826001848601038161552257fe5b0402905092915050565b60008082116155a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f536166654d6174683a206469766973696f6e206279207a65726f00000000000081525060200191505060405180910390fd5b60008284816155ae57fe5b0490508091505092915050565b6000808314156155ce576000905061563b565b60008284029050828482816155df57fe5b0414615636576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001806159b46021913960400191505060405180910390fd5b809150505b92915050565b6000828211156156b9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081525060200191505060405180910390fd5b600082840390508091505092915050565b60008183116156d957816156db565b825b905092915050565b60008183106156f257816156f4565b825b905092915050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415615782576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260268152602001806158f96026913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6040518060400160405280600061ffff168152602001600067ffffffffffffffff168152509056fe73686f756c6420626520656e6f75676820617661696c61626c652062616c616e63656e65772062616c616e63652073686f756c642062652062696767657220746861742063757272656e746368616e6e656c206861766520746f206265206e6f74206f70656e6564207965746f6e6c79206f70657261746f722063616e2063616c6c20746869732066756e6374696f6e4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573736e617469766520746f6b656e2066756e64732063616e2774206265207265636f76657265646e6f6e6365206861766520746f20626520626967676572207468616e20616c7265616479207573656462616c616e63652063616e2774206265206c657373207468616e206c6f616e20616d6f756e7462656e65666963696172792063616e2774206265207a65726f2061646472657373536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77616d6f756e7420746f20736574746c652073686f756c642062652067726561746572207468617420616c726561647920736574746c65646163636f756e74616e742073686f756c6420626520696e2070756e6973686d656e7420737461747573746f6b656e2063616e2774206265206465706c6f796420696e746f207a65726f20616464726573736861766520746f206265207369676e6564206279206368616e6e656c207061727479616d6f756e7420746f206c656e642063616e2774206265206269676765722074686174206d6178696d616c6c7920616c6c6f7765646163636f756e74616e74206861766520746f20626520696e2061637469766520737461746573686f756c6420626520656e6f7567682066756e647320617661696c61626c6520746f2077697468647261776861766520746f206265207369676e6564206279206368616e6e656c206f70657261746f726163636f756e74616e74206861766520746f20626520696e20706175736564207374617465a265627a7a7231582037078915fefadde25c82fd956efcc487e62283c06424cee97497d74c1c42b3ec64736f6c634300050c0032"

// DeployAccountantImplementation deploys a new Ethereum contract, binding an instance of AccountantImplementation to it.
func DeployAccountantImplementation(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AccountantImplementation, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountantImplementationABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AccountantImplementationBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AccountantImplementation{AccountantImplementationCaller: AccountantImplementationCaller{contract: contract}, AccountantImplementationTransactor: AccountantImplementationTransactor{contract: contract}, AccountantImplementationFilterer: AccountantImplementationFilterer{contract: contract}}, nil
}

// AccountantImplementation is an auto generated Go binding around an Ethereum contract.
type AccountantImplementation struct {
	AccountantImplementationCaller     // Read-only binding to the contract
	AccountantImplementationTransactor // Write-only binding to the contract
	AccountantImplementationFilterer   // Log filterer for contract events
}

// AccountantImplementationCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountantImplementationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountantImplementationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountantImplementationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountantImplementationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccountantImplementationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountantImplementationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountantImplementationSession struct {
	Contract     *AccountantImplementation // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AccountantImplementationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountantImplementationCallerSession struct {
	Contract *AccountantImplementationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// AccountantImplementationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountantImplementationTransactorSession struct {
	Contract     *AccountantImplementationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// AccountantImplementationRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountantImplementationRaw struct {
	Contract *AccountantImplementation // Generic contract binding to access the raw methods on
}

// AccountantImplementationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountantImplementationCallerRaw struct {
	Contract *AccountantImplementationCaller // Generic read-only contract binding to access the raw methods on
}

// AccountantImplementationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountantImplementationTransactorRaw struct {
	Contract *AccountantImplementationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccountantImplementation creates a new instance of AccountantImplementation, bound to a specific deployed contract.
func NewAccountantImplementation(address common.Address, backend bind.ContractBackend) (*AccountantImplementation, error) {
	contract, err := bindAccountantImplementation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccountantImplementation{AccountantImplementationCaller: AccountantImplementationCaller{contract: contract}, AccountantImplementationTransactor: AccountantImplementationTransactor{contract: contract}, AccountantImplementationFilterer: AccountantImplementationFilterer{contract: contract}}, nil
}

// NewAccountantImplementationCaller creates a new read-only instance of AccountantImplementation, bound to a specific deployed contract.
func NewAccountantImplementationCaller(address common.Address, caller bind.ContractCaller) (*AccountantImplementationCaller, error) {
	contract, err := bindAccountantImplementation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccountantImplementationCaller{contract: contract}, nil
}

// NewAccountantImplementationTransactor creates a new write-only instance of AccountantImplementation, bound to a specific deployed contract.
func NewAccountantImplementationTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountantImplementationTransactor, error) {
	contract, err := bindAccountantImplementation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccountantImplementationTransactor{contract: contract}, nil
}

// NewAccountantImplementationFilterer creates a new log filterer instance of AccountantImplementation, bound to a specific deployed contract.
func NewAccountantImplementationFilterer(address common.Address, filterer bind.ContractFilterer) (*AccountantImplementationFilterer, error) {
	contract, err := bindAccountantImplementation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccountantImplementationFilterer{contract: contract}, nil
}

// bindAccountantImplementation binds a generic wrapper to an already deployed contract.
func bindAccountantImplementation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountantImplementationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountantImplementation *AccountantImplementationRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AccountantImplementation.Contract.AccountantImplementationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountantImplementation *AccountantImplementationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.AccountantImplementationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountantImplementation *AccountantImplementationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.AccountantImplementationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountantImplementation *AccountantImplementationCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AccountantImplementation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountantImplementation *AccountantImplementationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountantImplementation *AccountantImplementationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.contract.Transact(opts, method, params...)
}

// AvailableBalance is a free data retrieval call binding the contract method 0xab2f0e51.
//
// Solidity: function availableBalance() constant returns(uint256)
func (_AccountantImplementation *AccountantImplementationCaller) AvailableBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccountantImplementation.contract.Call(opts, out, "availableBalance")
	return *ret0, err
}

// AvailableBalance is a free data retrieval call binding the contract method 0xab2f0e51.
//
// Solidity: function availableBalance() constant returns(uint256)
func (_AccountantImplementation *AccountantImplementationSession) AvailableBalance() (*big.Int, error) {
	return _AccountantImplementation.Contract.AvailableBalance(&_AccountantImplementation.CallOpts)
}

// AvailableBalance is a free data retrieval call binding the contract method 0xab2f0e51.
//
// Solidity: function availableBalance() constant returns(uint256)
func (_AccountantImplementation *AccountantImplementationCallerSession) AvailableBalance() (*big.Int, error) {
	return _AccountantImplementation.Contract.AvailableBalance(&_AccountantImplementation.CallOpts)
}

// Channels is a free data retrieval call binding the contract method 0x7a7ebd7b.
//
// Solidity: function channels(bytes32 ) constant returns(address beneficiary, uint256 balance, uint256 settled, uint256 loan, uint256 lastUsedNonce, uint256 timelock)
func (_AccountantImplementation *AccountantImplementationCaller) Channels(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Beneficiary   common.Address
	Balance       *big.Int
	Settled       *big.Int
	Loan          *big.Int
	LastUsedNonce *big.Int
	Timelock      *big.Int
}, error) {
	ret := new(struct {
		Beneficiary   common.Address
		Balance       *big.Int
		Settled       *big.Int
		Loan          *big.Int
		LastUsedNonce *big.Int
		Timelock      *big.Int
	})
	out := ret
	err := _AccountantImplementation.contract.Call(opts, out, "channels", arg0)
	return *ret, err
}

// Channels is a free data retrieval call binding the contract method 0x7a7ebd7b.
//
// Solidity: function channels(bytes32 ) constant returns(address beneficiary, uint256 balance, uint256 settled, uint256 loan, uint256 lastUsedNonce, uint256 timelock)
func (_AccountantImplementation *AccountantImplementationSession) Channels(arg0 [32]byte) (struct {
	Beneficiary   common.Address
	Balance       *big.Int
	Settled       *big.Int
	Loan          *big.Int
	LastUsedNonce *big.Int
	Timelock      *big.Int
}, error) {
	return _AccountantImplementation.Contract.Channels(&_AccountantImplementation.CallOpts, arg0)
}

// Channels is a free data retrieval call binding the contract method 0x7a7ebd7b.
//
// Solidity: function channels(bytes32 ) constant returns(address beneficiary, uint256 balance, uint256 settled, uint256 loan, uint256 lastUsedNonce, uint256 timelock)
func (_AccountantImplementation *AccountantImplementationCallerSession) Channels(arg0 [32]byte) (struct {
	Beneficiary   common.Address
	Balance       *big.Int
	Settled       *big.Int
	Loan          *big.Int
	LastUsedNonce *big.Int
	Timelock      *big.Int
}, error) {
	return _AccountantImplementation.Contract.Channels(&_AccountantImplementation.CallOpts, arg0)
}

// GetAccountantFee is a free data retrieval call binding the contract method 0xd0acce61.
//
// Solidity: function getAccountantFee(uint256 _amount) constant returns(uint256)
func (_AccountantImplementation *AccountantImplementationCaller) GetAccountantFee(opts *bind.CallOpts, _amount *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccountantImplementation.contract.Call(opts, out, "getAccountantFee", _amount)
	return *ret0, err
}

// GetAccountantFee is a free data retrieval call binding the contract method 0xd0acce61.
//
// Solidity: function getAccountantFee(uint256 _amount) constant returns(uint256)
func (_AccountantImplementation *AccountantImplementationSession) GetAccountantFee(_amount *big.Int) (*big.Int, error) {
	return _AccountantImplementation.Contract.GetAccountantFee(&_AccountantImplementation.CallOpts, _amount)
}

// GetAccountantFee is a free data retrieval call binding the contract method 0xd0acce61.
//
// Solidity: function getAccountantFee(uint256 _amount) constant returns(uint256)
func (_AccountantImplementation *AccountantImplementationCallerSession) GetAccountantFee(_amount *big.Int) (*big.Int, error) {
	return _AccountantImplementation.Contract.GetAccountantFee(&_AccountantImplementation.CallOpts, _amount)
}

// GetChannelId is a free data retrieval call binding the contract method 0xeb295b27.
//
// Solidity: function getChannelId(address _party) constant returns(bytes32)
func (_AccountantImplementation *AccountantImplementationCaller) GetChannelId(opts *bind.CallOpts, _party common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _AccountantImplementation.contract.Call(opts, out, "getChannelId", _party)
	return *ret0, err
}

// GetChannelId is a free data retrieval call binding the contract method 0xeb295b27.
//
// Solidity: function getChannelId(address _party) constant returns(bytes32)
func (_AccountantImplementation *AccountantImplementationSession) GetChannelId(_party common.Address) ([32]byte, error) {
	return _AccountantImplementation.Contract.GetChannelId(&_AccountantImplementation.CallOpts, _party)
}

// GetChannelId is a free data retrieval call binding the contract method 0xeb295b27.
//
// Solidity: function getChannelId(address _party) constant returns(bytes32)
func (_AccountantImplementation *AccountantImplementationCallerSession) GetChannelId(_party common.Address) ([32]byte, error) {
	return _AccountantImplementation.Contract.GetChannelId(&_AccountantImplementation.CallOpts, _party)
}

// GetFundsDestination is a free data retrieval call binding the contract method 0xf58c5b6e.
//
// Solidity: function getFundsDestination() constant returns(address)
func (_AccountantImplementation *AccountantImplementationCaller) GetFundsDestination(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AccountantImplementation.contract.Call(opts, out, "getFundsDestination")
	return *ret0, err
}

// GetFundsDestination is a free data retrieval call binding the contract method 0xf58c5b6e.
//
// Solidity: function getFundsDestination() constant returns(address)
func (_AccountantImplementation *AccountantImplementationSession) GetFundsDestination() (common.Address, error) {
	return _AccountantImplementation.Contract.GetFundsDestination(&_AccountantImplementation.CallOpts)
}

// GetFundsDestination is a free data retrieval call binding the contract method 0xf58c5b6e.
//
// Solidity: function getFundsDestination() constant returns(address)
func (_AccountantImplementation *AccountantImplementationCallerSession) GetFundsDestination() (common.Address, error) {
	return _AccountantImplementation.Contract.GetFundsDestination(&_AccountantImplementation.CallOpts)
}

// GetOperator is a free data retrieval call binding the contract method 0xe7f43c68.
//
// Solidity: function getOperator() constant returns(address)
func (_AccountantImplementation *AccountantImplementationCaller) GetOperator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AccountantImplementation.contract.Call(opts, out, "getOperator")
	return *ret0, err
}

// GetOperator is a free data retrieval call binding the contract method 0xe7f43c68.
//
// Solidity: function getOperator() constant returns(address)
func (_AccountantImplementation *AccountantImplementationSession) GetOperator() (common.Address, error) {
	return _AccountantImplementation.Contract.GetOperator(&_AccountantImplementation.CallOpts)
}

// GetOperator is a free data retrieval call binding the contract method 0xe7f43c68.
//
// Solidity: function getOperator() constant returns(address)
func (_AccountantImplementation *AccountantImplementationCallerSession) GetOperator() (common.Address, error) {
	return _AccountantImplementation.Contract.GetOperator(&_AccountantImplementation.CallOpts)
}

// GetRegistry is a free data retrieval call binding the contract method 0x5ab1bd53.
//
// Solidity: function getRegistry() constant returns(address)
func (_AccountantImplementation *AccountantImplementationCaller) GetRegistry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AccountantImplementation.contract.Call(opts, out, "getRegistry")
	return *ret0, err
}

// GetRegistry is a free data retrieval call binding the contract method 0x5ab1bd53.
//
// Solidity: function getRegistry() constant returns(address)
func (_AccountantImplementation *AccountantImplementationSession) GetRegistry() (common.Address, error) {
	return _AccountantImplementation.Contract.GetRegistry(&_AccountantImplementation.CallOpts)
}

// GetRegistry is a free data retrieval call binding the contract method 0x5ab1bd53.
//
// Solidity: function getRegistry() constant returns(address)
func (_AccountantImplementation *AccountantImplementationCallerSession) GetRegistry() (common.Address, error) {
	return _AccountantImplementation.Contract.GetRegistry(&_AccountantImplementation.CallOpts)
}

// GetStake is a free data retrieval call binding the contract method 0xfc0e3d90.
//
// Solidity: function getStake() constant returns(uint256)
func (_AccountantImplementation *AccountantImplementationCaller) GetStake(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccountantImplementation.contract.Call(opts, out, "getStake")
	return *ret0, err
}

// GetStake is a free data retrieval call binding the contract method 0xfc0e3d90.
//
// Solidity: function getStake() constant returns(uint256)
func (_AccountantImplementation *AccountantImplementationSession) GetStake() (*big.Int, error) {
	return _AccountantImplementation.Contract.GetStake(&_AccountantImplementation.CallOpts)
}

// GetStake is a free data retrieval call binding the contract method 0xfc0e3d90.
//
// Solidity: function getStake() constant returns(uint256)
func (_AccountantImplementation *AccountantImplementationCallerSession) GetStake() (*big.Int, error) {
	return _AccountantImplementation.Contract.GetStake(&_AccountantImplementation.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(uint8)
func (_AccountantImplementation *AccountantImplementationCaller) GetStatus(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _AccountantImplementation.contract.Call(opts, out, "getStatus")
	return *ret0, err
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(uint8)
func (_AccountantImplementation *AccountantImplementationSession) GetStatus() (uint8, error) {
	return _AccountantImplementation.Contract.GetStatus(&_AccountantImplementation.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(uint8)
func (_AccountantImplementation *AccountantImplementationCallerSession) GetStatus() (uint8, error) {
	return _AccountantImplementation.Contract.GetStatus(&_AccountantImplementation.CallOpts)
}

// IsAccountantActive is a free data retrieval call binding the contract method 0xbe02c06c.
//
// Solidity: function isAccountantActive() constant returns(bool)
func (_AccountantImplementation *AccountantImplementationCaller) IsAccountantActive(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AccountantImplementation.contract.Call(opts, out, "isAccountantActive")
	return *ret0, err
}

// IsAccountantActive is a free data retrieval call binding the contract method 0xbe02c06c.
//
// Solidity: function isAccountantActive() constant returns(bool)
func (_AccountantImplementation *AccountantImplementationSession) IsAccountantActive() (bool, error) {
	return _AccountantImplementation.Contract.IsAccountantActive(&_AccountantImplementation.CallOpts)
}

// IsAccountantActive is a free data retrieval call binding the contract method 0xbe02c06c.
//
// Solidity: function isAccountantActive() constant returns(bool)
func (_AccountantImplementation *AccountantImplementationCallerSession) IsAccountantActive() (bool, error) {
	return _AccountantImplementation.Contract.IsAccountantActive(&_AccountantImplementation.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() constant returns(bool)
func (_AccountantImplementation *AccountantImplementationCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AccountantImplementation.contract.Call(opts, out, "isInitialized")
	return *ret0, err
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() constant returns(bool)
func (_AccountantImplementation *AccountantImplementationSession) IsInitialized() (bool, error) {
	return _AccountantImplementation.Contract.IsInitialized(&_AccountantImplementation.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() constant returns(bool)
func (_AccountantImplementation *AccountantImplementationCallerSession) IsInitialized() (bool, error) {
	return _AccountantImplementation.Contract.IsInitialized(&_AccountantImplementation.CallOpts)
}

// IsOpened is a free data retrieval call binding the contract method 0x5f5794f1.
//
// Solidity: function isOpened(bytes32 _channelId) constant returns(bool)
func (_AccountantImplementation *AccountantImplementationCaller) IsOpened(opts *bind.CallOpts, _channelId [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AccountantImplementation.contract.Call(opts, out, "isOpened", _channelId)
	return *ret0, err
}

// IsOpened is a free data retrieval call binding the contract method 0x5f5794f1.
//
// Solidity: function isOpened(bytes32 _channelId) constant returns(bool)
func (_AccountantImplementation *AccountantImplementationSession) IsOpened(_channelId [32]byte) (bool, error) {
	return _AccountantImplementation.Contract.IsOpened(&_AccountantImplementation.CallOpts, _channelId)
}

// IsOpened is a free data retrieval call binding the contract method 0x5f5794f1.
//
// Solidity: function isOpened(bytes32 _channelId) constant returns(bool)
func (_AccountantImplementation *AccountantImplementationCallerSession) IsOpened(_channelId [32]byte) (bool, error) {
	return _AccountantImplementation.Contract.IsOpened(&_AccountantImplementation.CallOpts, _channelId)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_AccountantImplementation *AccountantImplementationCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AccountantImplementation.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_AccountantImplementation *AccountantImplementationSession) IsOwner() (bool, error) {
	return _AccountantImplementation.Contract.IsOwner(&_AccountantImplementation.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_AccountantImplementation *AccountantImplementationCallerSession) IsOwner() (bool, error) {
	return _AccountantImplementation.Contract.IsOwner(&_AccountantImplementation.CallOpts)
}

// LastFee is a free data retrieval call binding the contract method 0x9801134e.
//
// Solidity: function lastFee() constant returns(uint16 value, uint64 validFrom)
func (_AccountantImplementation *AccountantImplementationCaller) LastFee(opts *bind.CallOpts) (struct {
	Value     uint16
	ValidFrom uint64
}, error) {
	ret := new(struct {
		Value     uint16
		ValidFrom uint64
	})
	out := ret
	err := _AccountantImplementation.contract.Call(opts, out, "lastFee")
	return *ret, err
}

// LastFee is a free data retrieval call binding the contract method 0x9801134e.
//
// Solidity: function lastFee() constant returns(uint16 value, uint64 validFrom)
func (_AccountantImplementation *AccountantImplementationSession) LastFee() (struct {
	Value     uint16
	ValidFrom uint64
}, error) {
	return _AccountantImplementation.Contract.LastFee(&_AccountantImplementation.CallOpts)
}

// LastFee is a free data retrieval call binding the contract method 0x9801134e.
//
// Solidity: function lastFee() constant returns(uint16 value, uint64 validFrom)
func (_AccountantImplementation *AccountantImplementationCallerSession) LastFee() (struct {
	Value     uint16
	ValidFrom uint64
}, error) {
	return _AccountantImplementation.Contract.LastFee(&_AccountantImplementation.CallOpts)
}

// MinimalExpectedBalance is a free data retrieval call binding the contract method 0x94c7915d.
//
// Solidity: function minimalExpectedBalance() constant returns(uint256)
func (_AccountantImplementation *AccountantImplementationCaller) MinimalExpectedBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccountantImplementation.contract.Call(opts, out, "minimalExpectedBalance")
	return *ret0, err
}

// MinimalExpectedBalance is a free data retrieval call binding the contract method 0x94c7915d.
//
// Solidity: function minimalExpectedBalance() constant returns(uint256)
func (_AccountantImplementation *AccountantImplementationSession) MinimalExpectedBalance() (*big.Int, error) {
	return _AccountantImplementation.Contract.MinimalExpectedBalance(&_AccountantImplementation.CallOpts)
}

// MinimalExpectedBalance is a free data retrieval call binding the contract method 0x94c7915d.
//
// Solidity: function minimalExpectedBalance() constant returns(uint256)
func (_AccountantImplementation *AccountantImplementationCallerSession) MinimalExpectedBalance() (*big.Int, error) {
	return _AccountantImplementation.Contract.MinimalExpectedBalance(&_AccountantImplementation.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AccountantImplementation *AccountantImplementationCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AccountantImplementation.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AccountantImplementation *AccountantImplementationSession) Owner() (common.Address, error) {
	return _AccountantImplementation.Contract.Owner(&_AccountantImplementation.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AccountantImplementation *AccountantImplementationCallerSession) Owner() (common.Address, error) {
	return _AccountantImplementation.Contract.Owner(&_AccountantImplementation.CallOpts)
}

// PreviousFee is a free data retrieval call binding the contract method 0xe1c66487.
//
// Solidity: function previousFee() constant returns(uint16 value, uint64 validFrom)
func (_AccountantImplementation *AccountantImplementationCaller) PreviousFee(opts *bind.CallOpts) (struct {
	Value     uint16
	ValidFrom uint64
}, error) {
	ret := new(struct {
		Value     uint16
		ValidFrom uint64
	})
	out := ret
	err := _AccountantImplementation.contract.Call(opts, out, "previousFee")
	return *ret, err
}

// PreviousFee is a free data retrieval call binding the contract method 0xe1c66487.
//
// Solidity: function previousFee() constant returns(uint16 value, uint64 validFrom)
func (_AccountantImplementation *AccountantImplementationSession) PreviousFee() (struct {
	Value     uint16
	ValidFrom uint64
}, error) {
	return _AccountantImplementation.Contract.PreviousFee(&_AccountantImplementation.CallOpts)
}

// PreviousFee is a free data retrieval call binding the contract method 0xe1c66487.
//
// Solidity: function previousFee() constant returns(uint16 value, uint64 validFrom)
func (_AccountantImplementation *AccountantImplementationCallerSession) PreviousFee() (struct {
	Value     uint16
	ValidFrom uint64
}, error) {
	return _AccountantImplementation.Contract.PreviousFee(&_AccountantImplementation.CallOpts)
}

// Punishment is a free data retrieval call binding the contract method 0x0684cd20.
//
// Solidity: function punishment() constant returns(uint256 activationBlock, uint256 amount)
func (_AccountantImplementation *AccountantImplementationCaller) Punishment(opts *bind.CallOpts) (struct {
	ActivationBlock *big.Int
	Amount          *big.Int
}, error) {
	ret := new(struct {
		ActivationBlock *big.Int
		Amount          *big.Int
	})
	out := ret
	err := _AccountantImplementation.contract.Call(opts, out, "punishment")
	return *ret, err
}

// Punishment is a free data retrieval call binding the contract method 0x0684cd20.
//
// Solidity: function punishment() constant returns(uint256 activationBlock, uint256 amount)
func (_AccountantImplementation *AccountantImplementationSession) Punishment() (struct {
	ActivationBlock *big.Int
	Amount          *big.Int
}, error) {
	return _AccountantImplementation.Contract.Punishment(&_AccountantImplementation.CallOpts)
}

// Punishment is a free data retrieval call binding the contract method 0x0684cd20.
//
// Solidity: function punishment() constant returns(uint256 activationBlock, uint256 amount)
func (_AccountantImplementation *AccountantImplementationCallerSession) Punishment() (struct {
	ActivationBlock *big.Int
	Amount          *big.Int
}, error) {
	return _AccountantImplementation.Contract.Punishment(&_AccountantImplementation.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_AccountantImplementation *AccountantImplementationCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AccountantImplementation.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_AccountantImplementation *AccountantImplementationSession) Token() (common.Address, error) {
	return _AccountantImplementation.Contract.Token(&_AccountantImplementation.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_AccountantImplementation *AccountantImplementationCallerSession) Token() (common.Address, error) {
	return _AccountantImplementation.Contract.Token(&_AccountantImplementation.CallOpts)
}

// ActivateChannelOpening is a paid mutator transaction binding the contract method 0xfbb46b98.
//
// Solidity: function activateChannelOpening() returns()
func (_AccountantImplementation *AccountantImplementationTransactor) ActivateChannelOpening(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountantImplementation.contract.Transact(opts, "activateChannelOpening")
}

// ActivateChannelOpening is a paid mutator transaction binding the contract method 0xfbb46b98.
//
// Solidity: function activateChannelOpening() returns()
func (_AccountantImplementation *AccountantImplementationSession) ActivateChannelOpening() (*types.Transaction, error) {
	return _AccountantImplementation.Contract.ActivateChannelOpening(&_AccountantImplementation.TransactOpts)
}

// ActivateChannelOpening is a paid mutator transaction binding the contract method 0xfbb46b98.
//
// Solidity: function activateChannelOpening() returns()
func (_AccountantImplementation *AccountantImplementationTransactorSession) ActivateChannelOpening() (*types.Transaction, error) {
	return _AccountantImplementation.Contract.ActivateChannelOpening(&_AccountantImplementation.TransactOpts)
}

// ClaimEthers is a paid mutator transaction binding the contract method 0x6931b550.
//
// Solidity: function claimEthers() returns()
func (_AccountantImplementation *AccountantImplementationTransactor) ClaimEthers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountantImplementation.contract.Transact(opts, "claimEthers")
}

// ClaimEthers is a paid mutator transaction binding the contract method 0x6931b550.
//
// Solidity: function claimEthers() returns()
func (_AccountantImplementation *AccountantImplementationSession) ClaimEthers() (*types.Transaction, error) {
	return _AccountantImplementation.Contract.ClaimEthers(&_AccountantImplementation.TransactOpts)
}

// ClaimEthers is a paid mutator transaction binding the contract method 0x6931b550.
//
// Solidity: function claimEthers() returns()
func (_AccountantImplementation *AccountantImplementationTransactorSession) ClaimEthers() (*types.Transaction, error) {
	return _AccountantImplementation.Contract.ClaimEthers(&_AccountantImplementation.TransactOpts)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0xdf8de3e7.
//
// Solidity: function claimTokens(address _token) returns()
func (_AccountantImplementation *AccountantImplementationTransactor) ClaimTokens(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _AccountantImplementation.contract.Transact(opts, "claimTokens", _token)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0xdf8de3e7.
//
// Solidity: function claimTokens(address _token) returns()
func (_AccountantImplementation *AccountantImplementationSession) ClaimTokens(_token common.Address) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.ClaimTokens(&_AccountantImplementation.TransactOpts, _token)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0xdf8de3e7.
//
// Solidity: function claimTokens(address _token) returns()
func (_AccountantImplementation *AccountantImplementationTransactorSession) ClaimTokens(_token common.Address) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.ClaimTokens(&_AccountantImplementation.TransactOpts, _token)
}

// CloseAccountant is a paid mutator transaction binding the contract method 0xf7d76369.
//
// Solidity: function closeAccountant() returns()
func (_AccountantImplementation *AccountantImplementationTransactor) CloseAccountant(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountantImplementation.contract.Transact(opts, "closeAccountant")
}

// CloseAccountant is a paid mutator transaction binding the contract method 0xf7d76369.
//
// Solidity: function closeAccountant() returns()
func (_AccountantImplementation *AccountantImplementationSession) CloseAccountant() (*types.Transaction, error) {
	return _AccountantImplementation.Contract.CloseAccountant(&_AccountantImplementation.TransactOpts)
}

// CloseAccountant is a paid mutator transaction binding the contract method 0xf7d76369.
//
// Solidity: function closeAccountant() returns()
func (_AccountantImplementation *AccountantImplementationTransactorSession) CloseAccountant() (*types.Transaction, error) {
	return _AccountantImplementation.Contract.CloseAccountant(&_AccountantImplementation.TransactOpts)
}

// DecreaseLoan is a paid mutator transaction binding the contract method 0x54ded42d.
//
// Solidity: function decreaseLoan(bytes32 _channelId, uint256 _amount, uint256 _nonce, bytes _signature) returns()
func (_AccountantImplementation *AccountantImplementationTransactor) DecreaseLoan(opts *bind.TransactOpts, _channelId [32]byte, _amount *big.Int, _nonce *big.Int, _signature []byte) (*types.Transaction, error) {
	return _AccountantImplementation.contract.Transact(opts, "decreaseLoan", _channelId, _amount, _nonce, _signature)
}

// DecreaseLoan is a paid mutator transaction binding the contract method 0x54ded42d.
//
// Solidity: function decreaseLoan(bytes32 _channelId, uint256 _amount, uint256 _nonce, bytes _signature) returns()
func (_AccountantImplementation *AccountantImplementationSession) DecreaseLoan(_channelId [32]byte, _amount *big.Int, _nonce *big.Int, _signature []byte) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.DecreaseLoan(&_AccountantImplementation.TransactOpts, _channelId, _amount, _nonce, _signature)
}

// DecreaseLoan is a paid mutator transaction binding the contract method 0x54ded42d.
//
// Solidity: function decreaseLoan(bytes32 _channelId, uint256 _amount, uint256 _nonce, bytes _signature) returns()
func (_AccountantImplementation *AccountantImplementationTransactorSession) DecreaseLoan(_channelId [32]byte, _amount *big.Int, _nonce *big.Int, _signature []byte) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.DecreaseLoan(&_AccountantImplementation.TransactOpts, _channelId, _amount, _nonce, _signature)
}

// IncreaseAccountantStake is a paid mutator transaction binding the contract method 0x1822af6f.
//
// Solidity: function increaseAccountantStake(uint256 _additionalStake) returns()
func (_AccountantImplementation *AccountantImplementationTransactor) IncreaseAccountantStake(opts *bind.TransactOpts, _additionalStake *big.Int) (*types.Transaction, error) {
	return _AccountantImplementation.contract.Transact(opts, "increaseAccountantStake", _additionalStake)
}

// IncreaseAccountantStake is a paid mutator transaction binding the contract method 0x1822af6f.
//
// Solidity: function increaseAccountantStake(uint256 _additionalStake) returns()
func (_AccountantImplementation *AccountantImplementationSession) IncreaseAccountantStake(_additionalStake *big.Int) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.IncreaseAccountantStake(&_AccountantImplementation.TransactOpts, _additionalStake)
}

// IncreaseAccountantStake is a paid mutator transaction binding the contract method 0x1822af6f.
//
// Solidity: function increaseAccountantStake(uint256 _additionalStake) returns()
func (_AccountantImplementation *AccountantImplementationTransactorSession) IncreaseAccountantStake(_additionalStake *big.Int) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.IncreaseAccountantStake(&_AccountantImplementation.TransactOpts, _additionalStake)
}

// IncreaseLoan is a paid mutator transaction binding the contract method 0x6e17b0d7.
//
// Solidity: function increaseLoan(bytes32 _channelId, uint256 _amount) returns()
func (_AccountantImplementation *AccountantImplementationTransactor) IncreaseLoan(opts *bind.TransactOpts, _channelId [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _AccountantImplementation.contract.Transact(opts, "increaseLoan", _channelId, _amount)
}

// IncreaseLoan is a paid mutator transaction binding the contract method 0x6e17b0d7.
//
// Solidity: function increaseLoan(bytes32 _channelId, uint256 _amount) returns()
func (_AccountantImplementation *AccountantImplementationSession) IncreaseLoan(_channelId [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.IncreaseLoan(&_AccountantImplementation.TransactOpts, _channelId, _amount)
}

// IncreaseLoan is a paid mutator transaction binding the contract method 0x6e17b0d7.
//
// Solidity: function increaseLoan(bytes32 _channelId, uint256 _amount) returns()
func (_AccountantImplementation *AccountantImplementationTransactorSession) IncreaseLoan(_channelId [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.IncreaseLoan(&_AccountantImplementation.TransactOpts, _channelId, _amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xfec8157d.
//
// Solidity: function initialize(address _token, address _operator, uint16 _fee, uint256 _maxLoan) returns()
func (_AccountantImplementation *AccountantImplementationTransactor) Initialize(opts *bind.TransactOpts, _token common.Address, _operator common.Address, _fee uint16, _maxLoan *big.Int) (*types.Transaction, error) {
	return _AccountantImplementation.contract.Transact(opts, "initialize", _token, _operator, _fee, _maxLoan)
}

// Initialize is a paid mutator transaction binding the contract method 0xfec8157d.
//
// Solidity: function initialize(address _token, address _operator, uint16 _fee, uint256 _maxLoan) returns()
func (_AccountantImplementation *AccountantImplementationSession) Initialize(_token common.Address, _operator common.Address, _fee uint16, _maxLoan *big.Int) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.Initialize(&_AccountantImplementation.TransactOpts, _token, _operator, _fee, _maxLoan)
}

// Initialize is a paid mutator transaction binding the contract method 0xfec8157d.
//
// Solidity: function initialize(address _token, address _operator, uint16 _fee, uint256 _maxLoan) returns()
func (_AccountantImplementation *AccountantImplementationTransactorSession) Initialize(_token common.Address, _operator common.Address, _fee uint16, _maxLoan *big.Int) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.Initialize(&_AccountantImplementation.TransactOpts, _token, _operator, _fee, _maxLoan)
}

// OpenChannel is a paid mutator transaction binding the contract method 0x0a798f24.
//
// Solidity: function openChannel(address _party, address _beneficiary, uint256 _amountToLend) returns()
func (_AccountantImplementation *AccountantImplementationTransactor) OpenChannel(opts *bind.TransactOpts, _party common.Address, _beneficiary common.Address, _amountToLend *big.Int) (*types.Transaction, error) {
	return _AccountantImplementation.contract.Transact(opts, "openChannel", _party, _beneficiary, _amountToLend)
}

// OpenChannel is a paid mutator transaction binding the contract method 0x0a798f24.
//
// Solidity: function openChannel(address _party, address _beneficiary, uint256 _amountToLend) returns()
func (_AccountantImplementation *AccountantImplementationSession) OpenChannel(_party common.Address, _beneficiary common.Address, _amountToLend *big.Int) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.OpenChannel(&_AccountantImplementation.TransactOpts, _party, _beneficiary, _amountToLend)
}

// OpenChannel is a paid mutator transaction binding the contract method 0x0a798f24.
//
// Solidity: function openChannel(address _party, address _beneficiary, uint256 _amountToLend) returns()
func (_AccountantImplementation *AccountantImplementationTransactorSession) OpenChannel(_party common.Address, _beneficiary common.Address, _amountToLend *big.Int) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.OpenChannel(&_AccountantImplementation.TransactOpts, _party, _beneficiary, _amountToLend)
}

// PauseChannelOpening is a paid mutator transaction binding the contract method 0xaa606dee.
//
// Solidity: function pauseChannelOpening() returns()
func (_AccountantImplementation *AccountantImplementationTransactor) PauseChannelOpening(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountantImplementation.contract.Transact(opts, "pauseChannelOpening")
}

// PauseChannelOpening is a paid mutator transaction binding the contract method 0xaa606dee.
//
// Solidity: function pauseChannelOpening() returns()
func (_AccountantImplementation *AccountantImplementationSession) PauseChannelOpening() (*types.Transaction, error) {
	return _AccountantImplementation.Contract.PauseChannelOpening(&_AccountantImplementation.TransactOpts)
}

// PauseChannelOpening is a paid mutator transaction binding the contract method 0xaa606dee.
//
// Solidity: function pauseChannelOpening() returns()
func (_AccountantImplementation *AccountantImplementationTransactorSession) PauseChannelOpening() (*types.Transaction, error) {
	return _AccountantImplementation.Contract.PauseChannelOpening(&_AccountantImplementation.TransactOpts)
}

// RebalanceChannel is a paid mutator transaction binding the contract method 0xefde05ec.
//
// Solidity: function rebalanceChannel(bytes32 _channelId) returns()
func (_AccountantImplementation *AccountantImplementationTransactor) RebalanceChannel(opts *bind.TransactOpts, _channelId [32]byte) (*types.Transaction, error) {
	return _AccountantImplementation.contract.Transact(opts, "rebalanceChannel", _channelId)
}

// RebalanceChannel is a paid mutator transaction binding the contract method 0xefde05ec.
//
// Solidity: function rebalanceChannel(bytes32 _channelId) returns()
func (_AccountantImplementation *AccountantImplementationSession) RebalanceChannel(_channelId [32]byte) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.RebalanceChannel(&_AccountantImplementation.TransactOpts, _channelId)
}

// RebalanceChannel is a paid mutator transaction binding the contract method 0xefde05ec.
//
// Solidity: function rebalanceChannel(bytes32 _channelId) returns()
func (_AccountantImplementation *AccountantImplementationTransactorSession) RebalanceChannel(_channelId [32]byte) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.RebalanceChannel(&_AccountantImplementation.TransactOpts, _channelId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AccountantImplementation *AccountantImplementationTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountantImplementation.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AccountantImplementation *AccountantImplementationSession) RenounceOwnership() (*types.Transaction, error) {
	return _AccountantImplementation.Contract.RenounceOwnership(&_AccountantImplementation.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AccountantImplementation *AccountantImplementationTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AccountantImplementation.Contract.RenounceOwnership(&_AccountantImplementation.TransactOpts)
}

// ResolveEmergency is a paid mutator transaction binding the contract method 0x15c73afd.
//
// Solidity: function resolveEmergency() returns()
func (_AccountantImplementation *AccountantImplementationTransactor) ResolveEmergency(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountantImplementation.contract.Transact(opts, "resolveEmergency")
}

// ResolveEmergency is a paid mutator transaction binding the contract method 0x15c73afd.
//
// Solidity: function resolveEmergency() returns()
func (_AccountantImplementation *AccountantImplementationSession) ResolveEmergency() (*types.Transaction, error) {
	return _AccountantImplementation.Contract.ResolveEmergency(&_AccountantImplementation.TransactOpts)
}

// ResolveEmergency is a paid mutator transaction binding the contract method 0x15c73afd.
//
// Solidity: function resolveEmergency() returns()
func (_AccountantImplementation *AccountantImplementationTransactorSession) ResolveEmergency() (*types.Transaction, error) {
	return _AccountantImplementation.Contract.ResolveEmergency(&_AccountantImplementation.TransactOpts)
}

// SetAccountantFee is a paid mutator transaction binding the contract method 0x0996fcbc.
//
// Solidity: function setAccountantFee(uint16 _newFee) returns()
func (_AccountantImplementation *AccountantImplementationTransactor) SetAccountantFee(opts *bind.TransactOpts, _newFee uint16) (*types.Transaction, error) {
	return _AccountantImplementation.contract.Transact(opts, "setAccountantFee", _newFee)
}

// SetAccountantFee is a paid mutator transaction binding the contract method 0x0996fcbc.
//
// Solidity: function setAccountantFee(uint16 _newFee) returns()
func (_AccountantImplementation *AccountantImplementationSession) SetAccountantFee(_newFee uint16) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.SetAccountantFee(&_AccountantImplementation.TransactOpts, _newFee)
}

// SetAccountantFee is a paid mutator transaction binding the contract method 0x0996fcbc.
//
// Solidity: function setAccountantFee(uint16 _newFee) returns()
func (_AccountantImplementation *AccountantImplementationTransactorSession) SetAccountantFee(_newFee uint16) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.SetAccountantFee(&_AccountantImplementation.TransactOpts, _newFee)
}

// SetAccountantOperator is a paid mutator transaction binding the contract method 0x800d6afb.
//
// Solidity: function setAccountantOperator(address _newOperator) returns()
func (_AccountantImplementation *AccountantImplementationTransactor) SetAccountantOperator(opts *bind.TransactOpts, _newOperator common.Address) (*types.Transaction, error) {
	return _AccountantImplementation.contract.Transact(opts, "setAccountantOperator", _newOperator)
}

// SetAccountantOperator is a paid mutator transaction binding the contract method 0x800d6afb.
//
// Solidity: function setAccountantOperator(address _newOperator) returns()
func (_AccountantImplementation *AccountantImplementationSession) SetAccountantOperator(_newOperator common.Address) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.SetAccountantOperator(&_AccountantImplementation.TransactOpts, _newOperator)
}

// SetAccountantOperator is a paid mutator transaction binding the contract method 0x800d6afb.
//
// Solidity: function setAccountantOperator(address _newOperator) returns()
func (_AccountantImplementation *AccountantImplementationTransactorSession) SetAccountantOperator(_newOperator common.Address) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.SetAccountantOperator(&_AccountantImplementation.TransactOpts, _newOperator)
}

// SetBeneficiary is a paid mutator transaction binding the contract method 0x0b3834ea.
//
// Solidity: function setBeneficiary(address _party, address _newBeneficiary, uint256 _nonce, bytes _signature) returns()
func (_AccountantImplementation *AccountantImplementationTransactor) SetBeneficiary(opts *bind.TransactOpts, _party common.Address, _newBeneficiary common.Address, _nonce *big.Int, _signature []byte) (*types.Transaction, error) {
	return _AccountantImplementation.contract.Transact(opts, "setBeneficiary", _party, _newBeneficiary, _nonce, _signature)
}

// SetBeneficiary is a paid mutator transaction binding the contract method 0x0b3834ea.
//
// Solidity: function setBeneficiary(address _party, address _newBeneficiary, uint256 _nonce, bytes _signature) returns()
func (_AccountantImplementation *AccountantImplementationSession) SetBeneficiary(_party common.Address, _newBeneficiary common.Address, _nonce *big.Int, _signature []byte) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.SetBeneficiary(&_AccountantImplementation.TransactOpts, _party, _newBeneficiary, _nonce, _signature)
}

// SetBeneficiary is a paid mutator transaction binding the contract method 0x0b3834ea.
//
// Solidity: function setBeneficiary(address _party, address _newBeneficiary, uint256 _nonce, bytes _signature) returns()
func (_AccountantImplementation *AccountantImplementationTransactorSession) SetBeneficiary(_party common.Address, _newBeneficiary common.Address, _nonce *big.Int, _signature []byte) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.SetBeneficiary(&_AccountantImplementation.TransactOpts, _party, _newBeneficiary, _nonce, _signature)
}

// SetFundsDestination is a paid mutator transaction binding the contract method 0x238e130a.
//
// Solidity: function setFundsDestination(address _newDestination) returns()
func (_AccountantImplementation *AccountantImplementationTransactor) SetFundsDestination(opts *bind.TransactOpts, _newDestination common.Address) (*types.Transaction, error) {
	return _AccountantImplementation.contract.Transact(opts, "setFundsDestination", _newDestination)
}

// SetFundsDestination is a paid mutator transaction binding the contract method 0x238e130a.
//
// Solidity: function setFundsDestination(address _newDestination) returns()
func (_AccountantImplementation *AccountantImplementationSession) SetFundsDestination(_newDestination common.Address) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.SetFundsDestination(&_AccountantImplementation.TransactOpts, _newDestination)
}

// SetFundsDestination is a paid mutator transaction binding the contract method 0x238e130a.
//
// Solidity: function setFundsDestination(address _newDestination) returns()
func (_AccountantImplementation *AccountantImplementationTransactorSession) SetFundsDestination(_newDestination common.Address) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.SetFundsDestination(&_AccountantImplementation.TransactOpts, _newDestination)
}

// SetMaxLoan is a paid mutator transaction binding the contract method 0x456e182f.
//
// Solidity: function setMaxLoan(uint256 _newMaxLoan) returns()
func (_AccountantImplementation *AccountantImplementationTransactor) SetMaxLoan(opts *bind.TransactOpts, _newMaxLoan *big.Int) (*types.Transaction, error) {
	return _AccountantImplementation.contract.Transact(opts, "setMaxLoan", _newMaxLoan)
}

// SetMaxLoan is a paid mutator transaction binding the contract method 0x456e182f.
//
// Solidity: function setMaxLoan(uint256 _newMaxLoan) returns()
func (_AccountantImplementation *AccountantImplementationSession) SetMaxLoan(_newMaxLoan *big.Int) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.SetMaxLoan(&_AccountantImplementation.TransactOpts, _newMaxLoan)
}

// SetMaxLoan is a paid mutator transaction binding the contract method 0x456e182f.
//
// Solidity: function setMaxLoan(uint256 _newMaxLoan) returns()
func (_AccountantImplementation *AccountantImplementationTransactorSession) SetMaxLoan(_newMaxLoan *big.Int) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.SetMaxLoan(&_AccountantImplementation.TransactOpts, _newMaxLoan)
}

// SettlePromise is a paid mutator transaction binding the contract method 0xa58b2b71.
//
// Solidity: function settlePromise(bytes32 _channelId, uint256 _amount, uint256 _transactorFee, bytes32 _lock, bytes _signature) returns()
func (_AccountantImplementation *AccountantImplementationTransactor) SettlePromise(opts *bind.TransactOpts, _channelId [32]byte, _amount *big.Int, _transactorFee *big.Int, _lock [32]byte, _signature []byte) (*types.Transaction, error) {
	return _AccountantImplementation.contract.Transact(opts, "settlePromise", _channelId, _amount, _transactorFee, _lock, _signature)
}

// SettlePromise is a paid mutator transaction binding the contract method 0xa58b2b71.
//
// Solidity: function settlePromise(bytes32 _channelId, uint256 _amount, uint256 _transactorFee, bytes32 _lock, bytes _signature) returns()
func (_AccountantImplementation *AccountantImplementationSession) SettlePromise(_channelId [32]byte, _amount *big.Int, _transactorFee *big.Int, _lock [32]byte, _signature []byte) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.SettlePromise(&_AccountantImplementation.TransactOpts, _channelId, _amount, _transactorFee, _lock, _signature)
}

// SettlePromise is a paid mutator transaction binding the contract method 0xa58b2b71.
//
// Solidity: function settlePromise(bytes32 _channelId, uint256 _amount, uint256 _transactorFee, bytes32 _lock, bytes _signature) returns()
func (_AccountantImplementation *AccountantImplementationTransactorSession) SettlePromise(_channelId [32]byte, _amount *big.Int, _transactorFee *big.Int, _lock [32]byte, _signature []byte) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.SettlePromise(&_AccountantImplementation.TransactOpts, _channelId, _amount, _transactorFee, _lock, _signature)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AccountantImplementation *AccountantImplementationTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AccountantImplementation.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AccountantImplementation *AccountantImplementationSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.TransferOwnership(&_AccountantImplementation.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AccountantImplementation *AccountantImplementationTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.TransferOwnership(&_AccountantImplementation.TransactOpts, newOwner)
}

// UpdateChannelBalance is a paid mutator transaction binding the contract method 0x7c2be0a3.
//
// Solidity: function updateChannelBalance(bytes32 _channelId, uint256 _newBalance) returns()
func (_AccountantImplementation *AccountantImplementationTransactor) UpdateChannelBalance(opts *bind.TransactOpts, _channelId [32]byte, _newBalance *big.Int) (*types.Transaction, error) {
	return _AccountantImplementation.contract.Transact(opts, "updateChannelBalance", _channelId, _newBalance)
}

// UpdateChannelBalance is a paid mutator transaction binding the contract method 0x7c2be0a3.
//
// Solidity: function updateChannelBalance(bytes32 _channelId, uint256 _newBalance) returns()
func (_AccountantImplementation *AccountantImplementationSession) UpdateChannelBalance(_channelId [32]byte, _newBalance *big.Int) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.UpdateChannelBalance(&_AccountantImplementation.TransactOpts, _channelId, _newBalance)
}

// UpdateChannelBalance is a paid mutator transaction binding the contract method 0x7c2be0a3.
//
// Solidity: function updateChannelBalance(bytes32 _channelId, uint256 _newBalance) returns()
func (_AccountantImplementation *AccountantImplementationTransactorSession) UpdateChannelBalance(_channelId [32]byte, _newBalance *big.Int) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.UpdateChannelBalance(&_AccountantImplementation.TransactOpts, _channelId, _newBalance)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _beneficiary, uint256 _amount) returns()
func (_AccountantImplementation *AccountantImplementationTransactor) Withdraw(opts *bind.TransactOpts, _beneficiary common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AccountantImplementation.contract.Transact(opts, "withdraw", _beneficiary, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _beneficiary, uint256 _amount) returns()
func (_AccountantImplementation *AccountantImplementationSession) Withdraw(_beneficiary common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.Withdraw(&_AccountantImplementation.TransactOpts, _beneficiary, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _beneficiary, uint256 _amount) returns()
func (_AccountantImplementation *AccountantImplementationTransactorSession) Withdraw(_beneficiary common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.Withdraw(&_AccountantImplementation.TransactOpts, _beneficiary, _amount)
}

// AccountantImplementationAccountantClosedIterator is returned from FilterAccountantClosed and is used to iterate over the raw logs and unpacked data for AccountantClosed events raised by the AccountantImplementation contract.
type AccountantImplementationAccountantClosedIterator struct {
	Event *AccountantImplementationAccountantClosed // Event containing the contract specifics and raw log

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
func (it *AccountantImplementationAccountantClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountantImplementationAccountantClosed)
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
		it.Event = new(AccountantImplementationAccountantClosed)
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
func (it *AccountantImplementationAccountantClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountantImplementationAccountantClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountantImplementationAccountantClosed represents a AccountantClosed event raised by the AccountantImplementation contract.
type AccountantImplementationAccountantClosed struct {
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAccountantClosed is a free log retrieval operation binding the contract event 0x888906f0892e56365e679111a6f8ba7d0742bae70d0a532641cbf0da77d5af92.
//
// Solidity: event AccountantClosed(uint256 blockNumber)
func (_AccountantImplementation *AccountantImplementationFilterer) FilterAccountantClosed(opts *bind.FilterOpts) (*AccountantImplementationAccountantClosedIterator, error) {

	logs, sub, err := _AccountantImplementation.contract.FilterLogs(opts, "AccountantClosed")
	if err != nil {
		return nil, err
	}
	return &AccountantImplementationAccountantClosedIterator{contract: _AccountantImplementation.contract, event: "AccountantClosed", logs: logs, sub: sub}, nil
}

// WatchAccountantClosed is a free log subscription operation binding the contract event 0x888906f0892e56365e679111a6f8ba7d0742bae70d0a532641cbf0da77d5af92.
//
// Solidity: event AccountantClosed(uint256 blockNumber)
func (_AccountantImplementation *AccountantImplementationFilterer) WatchAccountantClosed(opts *bind.WatchOpts, sink chan<- *AccountantImplementationAccountantClosed) (event.Subscription, error) {

	logs, sub, err := _AccountantImplementation.contract.WatchLogs(opts, "AccountantClosed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountantImplementationAccountantClosed)
				if err := _AccountantImplementation.contract.UnpackLog(event, "AccountantClosed", log); err != nil {
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

// ParseAccountantClosed is a log parse operation binding the contract event 0x888906f0892e56365e679111a6f8ba7d0742bae70d0a532641cbf0da77d5af92.
//
// Solidity: event AccountantClosed(uint256 blockNumber)
func (_AccountantImplementation *AccountantImplementationFilterer) ParseAccountantClosed(log types.Log) (*AccountantImplementationAccountantClosed, error) {
	event := new(AccountantImplementationAccountantClosed)
	if err := _AccountantImplementation.contract.UnpackLog(event, "AccountantClosed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountantImplementationAccountantFeeUpdatedIterator is returned from FilterAccountantFeeUpdated and is used to iterate over the raw logs and unpacked data for AccountantFeeUpdated events raised by the AccountantImplementation contract.
type AccountantImplementationAccountantFeeUpdatedIterator struct {
	Event *AccountantImplementationAccountantFeeUpdated // Event containing the contract specifics and raw log

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
func (it *AccountantImplementationAccountantFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountantImplementationAccountantFeeUpdated)
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
		it.Event = new(AccountantImplementationAccountantFeeUpdated)
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
func (it *AccountantImplementationAccountantFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountantImplementationAccountantFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountantImplementationAccountantFeeUpdated represents a AccountantFeeUpdated event raised by the AccountantImplementation contract.
type AccountantImplementationAccountantFeeUpdated struct {
	NewFee         uint16
	ValidFromBlock uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAccountantFeeUpdated is a free log retrieval operation binding the contract event 0x008b9bffa5c48d8c0b96ce879f8491c5605cc3d7a46a46711b522dbe6d4070ee.
//
// Solidity: event AccountantFeeUpdated(uint16 newFee, uint64 validFromBlock)
func (_AccountantImplementation *AccountantImplementationFilterer) FilterAccountantFeeUpdated(opts *bind.FilterOpts) (*AccountantImplementationAccountantFeeUpdatedIterator, error) {

	logs, sub, err := _AccountantImplementation.contract.FilterLogs(opts, "AccountantFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &AccountantImplementationAccountantFeeUpdatedIterator{contract: _AccountantImplementation.contract, event: "AccountantFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchAccountantFeeUpdated is a free log subscription operation binding the contract event 0x008b9bffa5c48d8c0b96ce879f8491c5605cc3d7a46a46711b522dbe6d4070ee.
//
// Solidity: event AccountantFeeUpdated(uint16 newFee, uint64 validFromBlock)
func (_AccountantImplementation *AccountantImplementationFilterer) WatchAccountantFeeUpdated(opts *bind.WatchOpts, sink chan<- *AccountantImplementationAccountantFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _AccountantImplementation.contract.WatchLogs(opts, "AccountantFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountantImplementationAccountantFeeUpdated)
				if err := _AccountantImplementation.contract.UnpackLog(event, "AccountantFeeUpdated", log); err != nil {
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

// ParseAccountantFeeUpdated is a log parse operation binding the contract event 0x008b9bffa5c48d8c0b96ce879f8491c5605cc3d7a46a46711b522dbe6d4070ee.
//
// Solidity: event AccountantFeeUpdated(uint16 newFee, uint64 validFromBlock)
func (_AccountantImplementation *AccountantImplementationFilterer) ParseAccountantFeeUpdated(log types.Log) (*AccountantImplementationAccountantFeeUpdated, error) {
	event := new(AccountantImplementationAccountantFeeUpdated)
	if err := _AccountantImplementation.contract.UnpackLog(event, "AccountantFeeUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountantImplementationAccountantPunishmentActivatedIterator is returned from FilterAccountantPunishmentActivated and is used to iterate over the raw logs and unpacked data for AccountantPunishmentActivated events raised by the AccountantImplementation contract.
type AccountantImplementationAccountantPunishmentActivatedIterator struct {
	Event *AccountantImplementationAccountantPunishmentActivated // Event containing the contract specifics and raw log

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
func (it *AccountantImplementationAccountantPunishmentActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountantImplementationAccountantPunishmentActivated)
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
		it.Event = new(AccountantImplementationAccountantPunishmentActivated)
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
func (it *AccountantImplementationAccountantPunishmentActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountantImplementationAccountantPunishmentActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountantImplementationAccountantPunishmentActivated represents a AccountantPunishmentActivated event raised by the AccountantImplementation contract.
type AccountantImplementationAccountantPunishmentActivated struct {
	ActivationBlock *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAccountantPunishmentActivated is a free log retrieval operation binding the contract event 0xb3e91d0895882cef621b468a8235043537ca2e4d8d91ee6587801041054107e1.
//
// Solidity: event AccountantPunishmentActivated(uint256 activationBlock)
func (_AccountantImplementation *AccountantImplementationFilterer) FilterAccountantPunishmentActivated(opts *bind.FilterOpts) (*AccountantImplementationAccountantPunishmentActivatedIterator, error) {

	logs, sub, err := _AccountantImplementation.contract.FilterLogs(opts, "AccountantPunishmentActivated")
	if err != nil {
		return nil, err
	}
	return &AccountantImplementationAccountantPunishmentActivatedIterator{contract: _AccountantImplementation.contract, event: "AccountantPunishmentActivated", logs: logs, sub: sub}, nil
}

// WatchAccountantPunishmentActivated is a free log subscription operation binding the contract event 0xb3e91d0895882cef621b468a8235043537ca2e4d8d91ee6587801041054107e1.
//
// Solidity: event AccountantPunishmentActivated(uint256 activationBlock)
func (_AccountantImplementation *AccountantImplementationFilterer) WatchAccountantPunishmentActivated(opts *bind.WatchOpts, sink chan<- *AccountantImplementationAccountantPunishmentActivated) (event.Subscription, error) {

	logs, sub, err := _AccountantImplementation.contract.WatchLogs(opts, "AccountantPunishmentActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountantImplementationAccountantPunishmentActivated)
				if err := _AccountantImplementation.contract.UnpackLog(event, "AccountantPunishmentActivated", log); err != nil {
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

// ParseAccountantPunishmentActivated is a log parse operation binding the contract event 0xb3e91d0895882cef621b468a8235043537ca2e4d8d91ee6587801041054107e1.
//
// Solidity: event AccountantPunishmentActivated(uint256 activationBlock)
func (_AccountantImplementation *AccountantImplementationFilterer) ParseAccountantPunishmentActivated(log types.Log) (*AccountantImplementationAccountantPunishmentActivated, error) {
	event := new(AccountantImplementationAccountantPunishmentActivated)
	if err := _AccountantImplementation.contract.UnpackLog(event, "AccountantPunishmentActivated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountantImplementationAccountantPunishmentDeactivatedIterator is returned from FilterAccountantPunishmentDeactivated and is used to iterate over the raw logs and unpacked data for AccountantPunishmentDeactivated events raised by the AccountantImplementation contract.
type AccountantImplementationAccountantPunishmentDeactivatedIterator struct {
	Event *AccountantImplementationAccountantPunishmentDeactivated // Event containing the contract specifics and raw log

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
func (it *AccountantImplementationAccountantPunishmentDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountantImplementationAccountantPunishmentDeactivated)
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
		it.Event = new(AccountantImplementationAccountantPunishmentDeactivated)
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
func (it *AccountantImplementationAccountantPunishmentDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountantImplementationAccountantPunishmentDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountantImplementationAccountantPunishmentDeactivated represents a AccountantPunishmentDeactivated event raised by the AccountantImplementation contract.
type AccountantImplementationAccountantPunishmentDeactivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAccountantPunishmentDeactivated is a free log retrieval operation binding the contract event 0x58ef313a2eb2567f3b143ff20930622dd67a0de84902cc93b7ddddd72b7773ef.
//
// Solidity: event AccountantPunishmentDeactivated()
func (_AccountantImplementation *AccountantImplementationFilterer) FilterAccountantPunishmentDeactivated(opts *bind.FilterOpts) (*AccountantImplementationAccountantPunishmentDeactivatedIterator, error) {

	logs, sub, err := _AccountantImplementation.contract.FilterLogs(opts, "AccountantPunishmentDeactivated")
	if err != nil {
		return nil, err
	}
	return &AccountantImplementationAccountantPunishmentDeactivatedIterator{contract: _AccountantImplementation.contract, event: "AccountantPunishmentDeactivated", logs: logs, sub: sub}, nil
}

// WatchAccountantPunishmentDeactivated is a free log subscription operation binding the contract event 0x58ef313a2eb2567f3b143ff20930622dd67a0de84902cc93b7ddddd72b7773ef.
//
// Solidity: event AccountantPunishmentDeactivated()
func (_AccountantImplementation *AccountantImplementationFilterer) WatchAccountantPunishmentDeactivated(opts *bind.WatchOpts, sink chan<- *AccountantImplementationAccountantPunishmentDeactivated) (event.Subscription, error) {

	logs, sub, err := _AccountantImplementation.contract.WatchLogs(opts, "AccountantPunishmentDeactivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountantImplementationAccountantPunishmentDeactivated)
				if err := _AccountantImplementation.contract.UnpackLog(event, "AccountantPunishmentDeactivated", log); err != nil {
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

// ParseAccountantPunishmentDeactivated is a log parse operation binding the contract event 0x58ef313a2eb2567f3b143ff20930622dd67a0de84902cc93b7ddddd72b7773ef.
//
// Solidity: event AccountantPunishmentDeactivated()
func (_AccountantImplementation *AccountantImplementationFilterer) ParseAccountantPunishmentDeactivated(log types.Log) (*AccountantImplementationAccountantPunishmentDeactivated, error) {
	event := new(AccountantImplementationAccountantPunishmentDeactivated)
	if err := _AccountantImplementation.contract.UnpackLog(event, "AccountantPunishmentDeactivated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountantImplementationAccountantStakeIncreasedIterator is returned from FilterAccountantStakeIncreased and is used to iterate over the raw logs and unpacked data for AccountantStakeIncreased events raised by the AccountantImplementation contract.
type AccountantImplementationAccountantStakeIncreasedIterator struct {
	Event *AccountantImplementationAccountantStakeIncreased // Event containing the contract specifics and raw log

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
func (it *AccountantImplementationAccountantStakeIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountantImplementationAccountantStakeIncreased)
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
		it.Event = new(AccountantImplementationAccountantStakeIncreased)
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
func (it *AccountantImplementationAccountantStakeIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountantImplementationAccountantStakeIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountantImplementationAccountantStakeIncreased represents a AccountantStakeIncreased event raised by the AccountantImplementation contract.
type AccountantImplementationAccountantStakeIncreased struct {
	NewStake *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAccountantStakeIncreased is a free log retrieval operation binding the contract event 0x41a5bb80f9c1243f3d450690277c955ff8982168e34ed096457afdc31cefef7f.
//
// Solidity: event AccountantStakeIncreased(uint256 newStake)
func (_AccountantImplementation *AccountantImplementationFilterer) FilterAccountantStakeIncreased(opts *bind.FilterOpts) (*AccountantImplementationAccountantStakeIncreasedIterator, error) {

	logs, sub, err := _AccountantImplementation.contract.FilterLogs(opts, "AccountantStakeIncreased")
	if err != nil {
		return nil, err
	}
	return &AccountantImplementationAccountantStakeIncreasedIterator{contract: _AccountantImplementation.contract, event: "AccountantStakeIncreased", logs: logs, sub: sub}, nil
}

// WatchAccountantStakeIncreased is a free log subscription operation binding the contract event 0x41a5bb80f9c1243f3d450690277c955ff8982168e34ed096457afdc31cefef7f.
//
// Solidity: event AccountantStakeIncreased(uint256 newStake)
func (_AccountantImplementation *AccountantImplementationFilterer) WatchAccountantStakeIncreased(opts *bind.WatchOpts, sink chan<- *AccountantImplementationAccountantStakeIncreased) (event.Subscription, error) {

	logs, sub, err := _AccountantImplementation.contract.WatchLogs(opts, "AccountantStakeIncreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountantImplementationAccountantStakeIncreased)
				if err := _AccountantImplementation.contract.UnpackLog(event, "AccountantStakeIncreased", log); err != nil {
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

// ParseAccountantStakeIncreased is a log parse operation binding the contract event 0x41a5bb80f9c1243f3d450690277c955ff8982168e34ed096457afdc31cefef7f.
//
// Solidity: event AccountantStakeIncreased(uint256 newStake)
func (_AccountantImplementation *AccountantImplementationFilterer) ParseAccountantStakeIncreased(log types.Log) (*AccountantImplementationAccountantStakeIncreased, error) {
	event := new(AccountantImplementationAccountantStakeIncreased)
	if err := _AccountantImplementation.contract.UnpackLog(event, "AccountantStakeIncreased", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountantImplementationChannelBalanceDecreaseRequestedIterator is returned from FilterChannelBalanceDecreaseRequested and is used to iterate over the raw logs and unpacked data for ChannelBalanceDecreaseRequested events raised by the AccountantImplementation contract.
type AccountantImplementationChannelBalanceDecreaseRequestedIterator struct {
	Event *AccountantImplementationChannelBalanceDecreaseRequested // Event containing the contract specifics and raw log

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
func (it *AccountantImplementationChannelBalanceDecreaseRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountantImplementationChannelBalanceDecreaseRequested)
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
		it.Event = new(AccountantImplementationChannelBalanceDecreaseRequested)
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
func (it *AccountantImplementationChannelBalanceDecreaseRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountantImplementationChannelBalanceDecreaseRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountantImplementationChannelBalanceDecreaseRequested represents a ChannelBalanceDecreaseRequested event raised by the AccountantImplementation contract.
type AccountantImplementationChannelBalanceDecreaseRequested struct {
	ChannelId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChannelBalanceDecreaseRequested is a free log retrieval operation binding the contract event 0xaf4c616dc7856b81dbc1346e5547f0a1d4f1553011653f920d1041f215401075.
//
// Solidity: event ChannelBalanceDecreaseRequested(bytes32 indexed channelId)
func (_AccountantImplementation *AccountantImplementationFilterer) FilterChannelBalanceDecreaseRequested(opts *bind.FilterOpts, channelId [][32]byte) (*AccountantImplementationChannelBalanceDecreaseRequestedIterator, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _AccountantImplementation.contract.FilterLogs(opts, "ChannelBalanceDecreaseRequested", channelIdRule)
	if err != nil {
		return nil, err
	}
	return &AccountantImplementationChannelBalanceDecreaseRequestedIterator{contract: _AccountantImplementation.contract, event: "ChannelBalanceDecreaseRequested", logs: logs, sub: sub}, nil
}

// WatchChannelBalanceDecreaseRequested is a free log subscription operation binding the contract event 0xaf4c616dc7856b81dbc1346e5547f0a1d4f1553011653f920d1041f215401075.
//
// Solidity: event ChannelBalanceDecreaseRequested(bytes32 indexed channelId)
func (_AccountantImplementation *AccountantImplementationFilterer) WatchChannelBalanceDecreaseRequested(opts *bind.WatchOpts, sink chan<- *AccountantImplementationChannelBalanceDecreaseRequested, channelId [][32]byte) (event.Subscription, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _AccountantImplementation.contract.WatchLogs(opts, "ChannelBalanceDecreaseRequested", channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountantImplementationChannelBalanceDecreaseRequested)
				if err := _AccountantImplementation.contract.UnpackLog(event, "ChannelBalanceDecreaseRequested", log); err != nil {
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

// ParseChannelBalanceDecreaseRequested is a log parse operation binding the contract event 0xaf4c616dc7856b81dbc1346e5547f0a1d4f1553011653f920d1041f215401075.
//
// Solidity: event ChannelBalanceDecreaseRequested(bytes32 indexed channelId)
func (_AccountantImplementation *AccountantImplementationFilterer) ParseChannelBalanceDecreaseRequested(log types.Log) (*AccountantImplementationChannelBalanceDecreaseRequested, error) {
	event := new(AccountantImplementationChannelBalanceDecreaseRequested)
	if err := _AccountantImplementation.contract.UnpackLog(event, "ChannelBalanceDecreaseRequested", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountantImplementationChannelBalanceUpdatedIterator is returned from FilterChannelBalanceUpdated and is used to iterate over the raw logs and unpacked data for ChannelBalanceUpdated events raised by the AccountantImplementation contract.
type AccountantImplementationChannelBalanceUpdatedIterator struct {
	Event *AccountantImplementationChannelBalanceUpdated // Event containing the contract specifics and raw log

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
func (it *AccountantImplementationChannelBalanceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountantImplementationChannelBalanceUpdated)
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
		it.Event = new(AccountantImplementationChannelBalanceUpdated)
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
func (it *AccountantImplementationChannelBalanceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountantImplementationChannelBalanceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountantImplementationChannelBalanceUpdated represents a ChannelBalanceUpdated event raised by the AccountantImplementation contract.
type AccountantImplementationChannelBalanceUpdated struct {
	ChannelId  [32]byte
	NewBalance *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterChannelBalanceUpdated is a free log retrieval operation binding the contract event 0x2eb87c52e5d2d7395da77618aa1afe310b9cf13fa9f70e3199d100adda3a7aa5.
//
// Solidity: event ChannelBalanceUpdated(bytes32 indexed channelId, uint256 newBalance)
func (_AccountantImplementation *AccountantImplementationFilterer) FilterChannelBalanceUpdated(opts *bind.FilterOpts, channelId [][32]byte) (*AccountantImplementationChannelBalanceUpdatedIterator, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _AccountantImplementation.contract.FilterLogs(opts, "ChannelBalanceUpdated", channelIdRule)
	if err != nil {
		return nil, err
	}
	return &AccountantImplementationChannelBalanceUpdatedIterator{contract: _AccountantImplementation.contract, event: "ChannelBalanceUpdated", logs: logs, sub: sub}, nil
}

// WatchChannelBalanceUpdated is a free log subscription operation binding the contract event 0x2eb87c52e5d2d7395da77618aa1afe310b9cf13fa9f70e3199d100adda3a7aa5.
//
// Solidity: event ChannelBalanceUpdated(bytes32 indexed channelId, uint256 newBalance)
func (_AccountantImplementation *AccountantImplementationFilterer) WatchChannelBalanceUpdated(opts *bind.WatchOpts, sink chan<- *AccountantImplementationChannelBalanceUpdated, channelId [][32]byte) (event.Subscription, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _AccountantImplementation.contract.WatchLogs(opts, "ChannelBalanceUpdated", channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountantImplementationChannelBalanceUpdated)
				if err := _AccountantImplementation.contract.UnpackLog(event, "ChannelBalanceUpdated", log); err != nil {
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

// ParseChannelBalanceUpdated is a log parse operation binding the contract event 0x2eb87c52e5d2d7395da77618aa1afe310b9cf13fa9f70e3199d100adda3a7aa5.
//
// Solidity: event ChannelBalanceUpdated(bytes32 indexed channelId, uint256 newBalance)
func (_AccountantImplementation *AccountantImplementationFilterer) ParseChannelBalanceUpdated(log types.Log) (*AccountantImplementationChannelBalanceUpdated, error) {
	event := new(AccountantImplementationChannelBalanceUpdated)
	if err := _AccountantImplementation.contract.UnpackLog(event, "ChannelBalanceUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountantImplementationChannelBeneficiaryChangedIterator is returned from FilterChannelBeneficiaryChanged and is used to iterate over the raw logs and unpacked data for ChannelBeneficiaryChanged events raised by the AccountantImplementation contract.
type AccountantImplementationChannelBeneficiaryChangedIterator struct {
	Event *AccountantImplementationChannelBeneficiaryChanged // Event containing the contract specifics and raw log

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
func (it *AccountantImplementationChannelBeneficiaryChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountantImplementationChannelBeneficiaryChanged)
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
		it.Event = new(AccountantImplementationChannelBeneficiaryChanged)
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
func (it *AccountantImplementationChannelBeneficiaryChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountantImplementationChannelBeneficiaryChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountantImplementationChannelBeneficiaryChanged represents a ChannelBeneficiaryChanged event raised by the AccountantImplementation contract.
type AccountantImplementationChannelBeneficiaryChanged struct {
	ChannelId      [32]byte
	NewBeneficiary common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterChannelBeneficiaryChanged is a free log retrieval operation binding the contract event 0x8756aa559142225f918d7584303ecfe48e75b454f6614d0fae9f0d6ca0a898cc.
//
// Solidity: event ChannelBeneficiaryChanged(bytes32 channelId, address newBeneficiary)
func (_AccountantImplementation *AccountantImplementationFilterer) FilterChannelBeneficiaryChanged(opts *bind.FilterOpts) (*AccountantImplementationChannelBeneficiaryChangedIterator, error) {

	logs, sub, err := _AccountantImplementation.contract.FilterLogs(opts, "ChannelBeneficiaryChanged")
	if err != nil {
		return nil, err
	}
	return &AccountantImplementationChannelBeneficiaryChangedIterator{contract: _AccountantImplementation.contract, event: "ChannelBeneficiaryChanged", logs: logs, sub: sub}, nil
}

// WatchChannelBeneficiaryChanged is a free log subscription operation binding the contract event 0x8756aa559142225f918d7584303ecfe48e75b454f6614d0fae9f0d6ca0a898cc.
//
// Solidity: event ChannelBeneficiaryChanged(bytes32 channelId, address newBeneficiary)
func (_AccountantImplementation *AccountantImplementationFilterer) WatchChannelBeneficiaryChanged(opts *bind.WatchOpts, sink chan<- *AccountantImplementationChannelBeneficiaryChanged) (event.Subscription, error) {

	logs, sub, err := _AccountantImplementation.contract.WatchLogs(opts, "ChannelBeneficiaryChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountantImplementationChannelBeneficiaryChanged)
				if err := _AccountantImplementation.contract.UnpackLog(event, "ChannelBeneficiaryChanged", log); err != nil {
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

// ParseChannelBeneficiaryChanged is a log parse operation binding the contract event 0x8756aa559142225f918d7584303ecfe48e75b454f6614d0fae9f0d6ca0a898cc.
//
// Solidity: event ChannelBeneficiaryChanged(bytes32 channelId, address newBeneficiary)
func (_AccountantImplementation *AccountantImplementationFilterer) ParseChannelBeneficiaryChanged(log types.Log) (*AccountantImplementationChannelBeneficiaryChanged, error) {
	event := new(AccountantImplementationChannelBeneficiaryChanged)
	if err := _AccountantImplementation.contract.UnpackLog(event, "ChannelBeneficiaryChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountantImplementationChannelOpenedIterator is returned from FilterChannelOpened and is used to iterate over the raw logs and unpacked data for ChannelOpened events raised by the AccountantImplementation contract.
type AccountantImplementationChannelOpenedIterator struct {
	Event *AccountantImplementationChannelOpened // Event containing the contract specifics and raw log

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
func (it *AccountantImplementationChannelOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountantImplementationChannelOpened)
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
		it.Event = new(AccountantImplementationChannelOpened)
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
func (it *AccountantImplementationChannelOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountantImplementationChannelOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountantImplementationChannelOpened represents a ChannelOpened event raised by the AccountantImplementation contract.
type AccountantImplementationChannelOpened struct {
	ChannelId      [32]byte
	InitialBalance *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterChannelOpened is a free log retrieval operation binding the contract event 0xbe2e1f3a6197dfd16fa6830c4870364b618b8b288c21cbcfa4fdb5d7c6a5e45b.
//
// Solidity: event ChannelOpened(bytes32 channelId, uint256 initialBalance)
func (_AccountantImplementation *AccountantImplementationFilterer) FilterChannelOpened(opts *bind.FilterOpts) (*AccountantImplementationChannelOpenedIterator, error) {

	logs, sub, err := _AccountantImplementation.contract.FilterLogs(opts, "ChannelOpened")
	if err != nil {
		return nil, err
	}
	return &AccountantImplementationChannelOpenedIterator{contract: _AccountantImplementation.contract, event: "ChannelOpened", logs: logs, sub: sub}, nil
}

// WatchChannelOpened is a free log subscription operation binding the contract event 0xbe2e1f3a6197dfd16fa6830c4870364b618b8b288c21cbcfa4fdb5d7c6a5e45b.
//
// Solidity: event ChannelOpened(bytes32 channelId, uint256 initialBalance)
func (_AccountantImplementation *AccountantImplementationFilterer) WatchChannelOpened(opts *bind.WatchOpts, sink chan<- *AccountantImplementationChannelOpened) (event.Subscription, error) {

	logs, sub, err := _AccountantImplementation.contract.WatchLogs(opts, "ChannelOpened")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountantImplementationChannelOpened)
				if err := _AccountantImplementation.contract.UnpackLog(event, "ChannelOpened", log); err != nil {
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

// ParseChannelOpened is a log parse operation binding the contract event 0xbe2e1f3a6197dfd16fa6830c4870364b618b8b288c21cbcfa4fdb5d7c6a5e45b.
//
// Solidity: event ChannelOpened(bytes32 channelId, uint256 initialBalance)
func (_AccountantImplementation *AccountantImplementationFilterer) ParseChannelOpened(log types.Log) (*AccountantImplementationChannelOpened, error) {
	event := new(AccountantImplementationChannelOpened)
	if err := _AccountantImplementation.contract.UnpackLog(event, "ChannelOpened", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountantImplementationChannelOpeningActivatedIterator is returned from FilterChannelOpeningActivated and is used to iterate over the raw logs and unpacked data for ChannelOpeningActivated events raised by the AccountantImplementation contract.
type AccountantImplementationChannelOpeningActivatedIterator struct {
	Event *AccountantImplementationChannelOpeningActivated // Event containing the contract specifics and raw log

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
func (it *AccountantImplementationChannelOpeningActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountantImplementationChannelOpeningActivated)
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
		it.Event = new(AccountantImplementationChannelOpeningActivated)
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
func (it *AccountantImplementationChannelOpeningActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountantImplementationChannelOpeningActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountantImplementationChannelOpeningActivated represents a ChannelOpeningActivated event raised by the AccountantImplementation contract.
type AccountantImplementationChannelOpeningActivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChannelOpeningActivated is a free log retrieval operation binding the contract event 0x2d8b6ec230798e206d536342a28b7b61cc8fcfafb1d27c11c5519b3c42eb7df8.
//
// Solidity: event ChannelOpeningActivated()
func (_AccountantImplementation *AccountantImplementationFilterer) FilterChannelOpeningActivated(opts *bind.FilterOpts) (*AccountantImplementationChannelOpeningActivatedIterator, error) {

	logs, sub, err := _AccountantImplementation.contract.FilterLogs(opts, "ChannelOpeningActivated")
	if err != nil {
		return nil, err
	}
	return &AccountantImplementationChannelOpeningActivatedIterator{contract: _AccountantImplementation.contract, event: "ChannelOpeningActivated", logs: logs, sub: sub}, nil
}

// WatchChannelOpeningActivated is a free log subscription operation binding the contract event 0x2d8b6ec230798e206d536342a28b7b61cc8fcfafb1d27c11c5519b3c42eb7df8.
//
// Solidity: event ChannelOpeningActivated()
func (_AccountantImplementation *AccountantImplementationFilterer) WatchChannelOpeningActivated(opts *bind.WatchOpts, sink chan<- *AccountantImplementationChannelOpeningActivated) (event.Subscription, error) {

	logs, sub, err := _AccountantImplementation.contract.WatchLogs(opts, "ChannelOpeningActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountantImplementationChannelOpeningActivated)
				if err := _AccountantImplementation.contract.UnpackLog(event, "ChannelOpeningActivated", log); err != nil {
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

// ParseChannelOpeningActivated is a log parse operation binding the contract event 0x2d8b6ec230798e206d536342a28b7b61cc8fcfafb1d27c11c5519b3c42eb7df8.
//
// Solidity: event ChannelOpeningActivated()
func (_AccountantImplementation *AccountantImplementationFilterer) ParseChannelOpeningActivated(log types.Log) (*AccountantImplementationChannelOpeningActivated, error) {
	event := new(AccountantImplementationChannelOpeningActivated)
	if err := _AccountantImplementation.contract.UnpackLog(event, "ChannelOpeningActivated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountantImplementationChannelOpeningPausedIterator is returned from FilterChannelOpeningPaused and is used to iterate over the raw logs and unpacked data for ChannelOpeningPaused events raised by the AccountantImplementation contract.
type AccountantImplementationChannelOpeningPausedIterator struct {
	Event *AccountantImplementationChannelOpeningPaused // Event containing the contract specifics and raw log

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
func (it *AccountantImplementationChannelOpeningPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountantImplementationChannelOpeningPaused)
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
		it.Event = new(AccountantImplementationChannelOpeningPaused)
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
func (it *AccountantImplementationChannelOpeningPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountantImplementationChannelOpeningPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountantImplementationChannelOpeningPaused represents a ChannelOpeningPaused event raised by the AccountantImplementation contract.
type AccountantImplementationChannelOpeningPaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChannelOpeningPaused is a free log retrieval operation binding the contract event 0x1f4cd5d6edef8a0c4dbe6d547fdc42e0f3575167257553271f2366f9d497f67e.
//
// Solidity: event ChannelOpeningPaused()
func (_AccountantImplementation *AccountantImplementationFilterer) FilterChannelOpeningPaused(opts *bind.FilterOpts) (*AccountantImplementationChannelOpeningPausedIterator, error) {

	logs, sub, err := _AccountantImplementation.contract.FilterLogs(opts, "ChannelOpeningPaused")
	if err != nil {
		return nil, err
	}
	return &AccountantImplementationChannelOpeningPausedIterator{contract: _AccountantImplementation.contract, event: "ChannelOpeningPaused", logs: logs, sub: sub}, nil
}

// WatchChannelOpeningPaused is a free log subscription operation binding the contract event 0x1f4cd5d6edef8a0c4dbe6d547fdc42e0f3575167257553271f2366f9d497f67e.
//
// Solidity: event ChannelOpeningPaused()
func (_AccountantImplementation *AccountantImplementationFilterer) WatchChannelOpeningPaused(opts *bind.WatchOpts, sink chan<- *AccountantImplementationChannelOpeningPaused) (event.Subscription, error) {

	logs, sub, err := _AccountantImplementation.contract.WatchLogs(opts, "ChannelOpeningPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountantImplementationChannelOpeningPaused)
				if err := _AccountantImplementation.contract.UnpackLog(event, "ChannelOpeningPaused", log); err != nil {
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

// ParseChannelOpeningPaused is a log parse operation binding the contract event 0x1f4cd5d6edef8a0c4dbe6d547fdc42e0f3575167257553271f2366f9d497f67e.
//
// Solidity: event ChannelOpeningPaused()
func (_AccountantImplementation *AccountantImplementationFilterer) ParseChannelOpeningPaused(log types.Log) (*AccountantImplementationChannelOpeningPaused, error) {
	event := new(AccountantImplementationChannelOpeningPaused)
	if err := _AccountantImplementation.contract.UnpackLog(event, "ChannelOpeningPaused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountantImplementationDestinationChangedIterator is returned from FilterDestinationChanged and is used to iterate over the raw logs and unpacked data for DestinationChanged events raised by the AccountantImplementation contract.
type AccountantImplementationDestinationChangedIterator struct {
	Event *AccountantImplementationDestinationChanged // Event containing the contract specifics and raw log

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
func (it *AccountantImplementationDestinationChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountantImplementationDestinationChanged)
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
		it.Event = new(AccountantImplementationDestinationChanged)
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
func (it *AccountantImplementationDestinationChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountantImplementationDestinationChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountantImplementationDestinationChanged represents a DestinationChanged event raised by the AccountantImplementation contract.
type AccountantImplementationDestinationChanged struct {
	PreviousDestination common.Address
	NewDestination      common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDestinationChanged is a free log retrieval operation binding the contract event 0xe1a66d77649cf0a57b9937073549f30f1c82bb865aaf066d2f299e37a62c6aad.
//
// Solidity: event DestinationChanged(address indexed previousDestination, address indexed newDestination)
func (_AccountantImplementation *AccountantImplementationFilterer) FilterDestinationChanged(opts *bind.FilterOpts, previousDestination []common.Address, newDestination []common.Address) (*AccountantImplementationDestinationChangedIterator, error) {

	var previousDestinationRule []interface{}
	for _, previousDestinationItem := range previousDestination {
		previousDestinationRule = append(previousDestinationRule, previousDestinationItem)
	}
	var newDestinationRule []interface{}
	for _, newDestinationItem := range newDestination {
		newDestinationRule = append(newDestinationRule, newDestinationItem)
	}

	logs, sub, err := _AccountantImplementation.contract.FilterLogs(opts, "DestinationChanged", previousDestinationRule, newDestinationRule)
	if err != nil {
		return nil, err
	}
	return &AccountantImplementationDestinationChangedIterator{contract: _AccountantImplementation.contract, event: "DestinationChanged", logs: logs, sub: sub}, nil
}

// WatchDestinationChanged is a free log subscription operation binding the contract event 0xe1a66d77649cf0a57b9937073549f30f1c82bb865aaf066d2f299e37a62c6aad.
//
// Solidity: event DestinationChanged(address indexed previousDestination, address indexed newDestination)
func (_AccountantImplementation *AccountantImplementationFilterer) WatchDestinationChanged(opts *bind.WatchOpts, sink chan<- *AccountantImplementationDestinationChanged, previousDestination []common.Address, newDestination []common.Address) (event.Subscription, error) {

	var previousDestinationRule []interface{}
	for _, previousDestinationItem := range previousDestination {
		previousDestinationRule = append(previousDestinationRule, previousDestinationItem)
	}
	var newDestinationRule []interface{}
	for _, newDestinationItem := range newDestination {
		newDestinationRule = append(newDestinationRule, newDestinationItem)
	}

	logs, sub, err := _AccountantImplementation.contract.WatchLogs(opts, "DestinationChanged", previousDestinationRule, newDestinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountantImplementationDestinationChanged)
				if err := _AccountantImplementation.contract.UnpackLog(event, "DestinationChanged", log); err != nil {
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
func (_AccountantImplementation *AccountantImplementationFilterer) ParseDestinationChanged(log types.Log) (*AccountantImplementationDestinationChanged, error) {
	event := new(AccountantImplementationDestinationChanged)
	if err := _AccountantImplementation.contract.UnpackLog(event, "DestinationChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountantImplementationFundsWithdrawnedIterator is returned from FilterFundsWithdrawned and is used to iterate over the raw logs and unpacked data for FundsWithdrawned events raised by the AccountantImplementation contract.
type AccountantImplementationFundsWithdrawnedIterator struct {
	Event *AccountantImplementationFundsWithdrawned // Event containing the contract specifics and raw log

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
func (it *AccountantImplementationFundsWithdrawnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountantImplementationFundsWithdrawned)
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
		it.Event = new(AccountantImplementationFundsWithdrawned)
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
func (it *AccountantImplementationFundsWithdrawnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountantImplementationFundsWithdrawnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountantImplementationFundsWithdrawned represents a FundsWithdrawned event raised by the AccountantImplementation contract.
type AccountantImplementationFundsWithdrawned struct {
	Amount      *big.Int
	Beneficiary common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFundsWithdrawned is a free log retrieval operation binding the contract event 0xa2e147ce2b7cb83d9c07e397bb806f23dd42c42e86ea45e1611d6e50eb1ec8bf.
//
// Solidity: event FundsWithdrawned(uint256 amount, address beneficiary)
func (_AccountantImplementation *AccountantImplementationFilterer) FilterFundsWithdrawned(opts *bind.FilterOpts) (*AccountantImplementationFundsWithdrawnedIterator, error) {

	logs, sub, err := _AccountantImplementation.contract.FilterLogs(opts, "FundsWithdrawned")
	if err != nil {
		return nil, err
	}
	return &AccountantImplementationFundsWithdrawnedIterator{contract: _AccountantImplementation.contract, event: "FundsWithdrawned", logs: logs, sub: sub}, nil
}

// WatchFundsWithdrawned is a free log subscription operation binding the contract event 0xa2e147ce2b7cb83d9c07e397bb806f23dd42c42e86ea45e1611d6e50eb1ec8bf.
//
// Solidity: event FundsWithdrawned(uint256 amount, address beneficiary)
func (_AccountantImplementation *AccountantImplementationFilterer) WatchFundsWithdrawned(opts *bind.WatchOpts, sink chan<- *AccountantImplementationFundsWithdrawned) (event.Subscription, error) {

	logs, sub, err := _AccountantImplementation.contract.WatchLogs(opts, "FundsWithdrawned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountantImplementationFundsWithdrawned)
				if err := _AccountantImplementation.contract.UnpackLog(event, "FundsWithdrawned", log); err != nil {
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

// ParseFundsWithdrawned is a log parse operation binding the contract event 0xa2e147ce2b7cb83d9c07e397bb806f23dd42c42e86ea45e1611d6e50eb1ec8bf.
//
// Solidity: event FundsWithdrawned(uint256 amount, address beneficiary)
func (_AccountantImplementation *AccountantImplementationFilterer) ParseFundsWithdrawned(log types.Log) (*AccountantImplementationFundsWithdrawned, error) {
	event := new(AccountantImplementationFundsWithdrawned)
	if err := _AccountantImplementation.contract.UnpackLog(event, "FundsWithdrawned", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountantImplementationMaxLoanValueUpdatedIterator is returned from FilterMaxLoanValueUpdated and is used to iterate over the raw logs and unpacked data for MaxLoanValueUpdated events raised by the AccountantImplementation contract.
type AccountantImplementationMaxLoanValueUpdatedIterator struct {
	Event *AccountantImplementationMaxLoanValueUpdated // Event containing the contract specifics and raw log

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
func (it *AccountantImplementationMaxLoanValueUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountantImplementationMaxLoanValueUpdated)
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
		it.Event = new(AccountantImplementationMaxLoanValueUpdated)
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
func (it *AccountantImplementationMaxLoanValueUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountantImplementationMaxLoanValueUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountantImplementationMaxLoanValueUpdated represents a MaxLoanValueUpdated event raised by the AccountantImplementation contract.
type AccountantImplementationMaxLoanValueUpdated struct {
	NewMaxLoan *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMaxLoanValueUpdated is a free log retrieval operation binding the contract event 0x2181d8ed90eadf541579998852434d883f30ace513163bbe0e4115b29afb517a.
//
// Solidity: event MaxLoanValueUpdated(uint256 _newMaxLoan)
func (_AccountantImplementation *AccountantImplementationFilterer) FilterMaxLoanValueUpdated(opts *bind.FilterOpts) (*AccountantImplementationMaxLoanValueUpdatedIterator, error) {

	logs, sub, err := _AccountantImplementation.contract.FilterLogs(opts, "MaxLoanValueUpdated")
	if err != nil {
		return nil, err
	}
	return &AccountantImplementationMaxLoanValueUpdatedIterator{contract: _AccountantImplementation.contract, event: "MaxLoanValueUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxLoanValueUpdated is a free log subscription operation binding the contract event 0x2181d8ed90eadf541579998852434d883f30ace513163bbe0e4115b29afb517a.
//
// Solidity: event MaxLoanValueUpdated(uint256 _newMaxLoan)
func (_AccountantImplementation *AccountantImplementationFilterer) WatchMaxLoanValueUpdated(opts *bind.WatchOpts, sink chan<- *AccountantImplementationMaxLoanValueUpdated) (event.Subscription, error) {

	logs, sub, err := _AccountantImplementation.contract.WatchLogs(opts, "MaxLoanValueUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountantImplementationMaxLoanValueUpdated)
				if err := _AccountantImplementation.contract.UnpackLog(event, "MaxLoanValueUpdated", log); err != nil {
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

// ParseMaxLoanValueUpdated is a log parse operation binding the contract event 0x2181d8ed90eadf541579998852434d883f30ace513163bbe0e4115b29afb517a.
//
// Solidity: event MaxLoanValueUpdated(uint256 _newMaxLoan)
func (_AccountantImplementation *AccountantImplementationFilterer) ParseMaxLoanValueUpdated(log types.Log) (*AccountantImplementationMaxLoanValueUpdated, error) {
	event := new(AccountantImplementationMaxLoanValueUpdated)
	if err := _AccountantImplementation.contract.UnpackLog(event, "MaxLoanValueUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountantImplementationNewAccountantOperatorIterator is returned from FilterNewAccountantOperator and is used to iterate over the raw logs and unpacked data for NewAccountantOperator events raised by the AccountantImplementation contract.
type AccountantImplementationNewAccountantOperatorIterator struct {
	Event *AccountantImplementationNewAccountantOperator // Event containing the contract specifics and raw log

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
func (it *AccountantImplementationNewAccountantOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountantImplementationNewAccountantOperator)
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
		it.Event = new(AccountantImplementationNewAccountantOperator)
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
func (it *AccountantImplementationNewAccountantOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountantImplementationNewAccountantOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountantImplementationNewAccountantOperator represents a NewAccountantOperator event raised by the AccountantImplementation contract.
type AccountantImplementationNewAccountantOperator struct {
	NewOperator common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewAccountantOperator is a free log retrieval operation binding the contract event 0xa326c787c51b80574c7b572d0c9664e64f1107538b902f519a896901b4137918.
//
// Solidity: event NewAccountantOperator(address newOperator)
func (_AccountantImplementation *AccountantImplementationFilterer) FilterNewAccountantOperator(opts *bind.FilterOpts) (*AccountantImplementationNewAccountantOperatorIterator, error) {

	logs, sub, err := _AccountantImplementation.contract.FilterLogs(opts, "NewAccountantOperator")
	if err != nil {
		return nil, err
	}
	return &AccountantImplementationNewAccountantOperatorIterator{contract: _AccountantImplementation.contract, event: "NewAccountantOperator", logs: logs, sub: sub}, nil
}

// WatchNewAccountantOperator is a free log subscription operation binding the contract event 0xa326c787c51b80574c7b572d0c9664e64f1107538b902f519a896901b4137918.
//
// Solidity: event NewAccountantOperator(address newOperator)
func (_AccountantImplementation *AccountantImplementationFilterer) WatchNewAccountantOperator(opts *bind.WatchOpts, sink chan<- *AccountantImplementationNewAccountantOperator) (event.Subscription, error) {

	logs, sub, err := _AccountantImplementation.contract.WatchLogs(opts, "NewAccountantOperator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountantImplementationNewAccountantOperator)
				if err := _AccountantImplementation.contract.UnpackLog(event, "NewAccountantOperator", log); err != nil {
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

// ParseNewAccountantOperator is a log parse operation binding the contract event 0xa326c787c51b80574c7b572d0c9664e64f1107538b902f519a896901b4137918.
//
// Solidity: event NewAccountantOperator(address newOperator)
func (_AccountantImplementation *AccountantImplementationFilterer) ParseNewAccountantOperator(log types.Log) (*AccountantImplementationNewAccountantOperator, error) {
	event := new(AccountantImplementationNewAccountantOperator)
	if err := _AccountantImplementation.contract.UnpackLog(event, "NewAccountantOperator", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountantImplementationNewLoanIterator is returned from FilterNewLoan and is used to iterate over the raw logs and unpacked data for NewLoan events raised by the AccountantImplementation contract.
type AccountantImplementationNewLoanIterator struct {
	Event *AccountantImplementationNewLoan // Event containing the contract specifics and raw log

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
func (it *AccountantImplementationNewLoanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountantImplementationNewLoan)
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
		it.Event = new(AccountantImplementationNewLoan)
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
func (it *AccountantImplementationNewLoanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountantImplementationNewLoanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountantImplementationNewLoan represents a NewLoan event raised by the AccountantImplementation contract.
type AccountantImplementationNewLoan struct {
	ChannelId  [32]byte
	LoanAmount *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewLoan is a free log retrieval operation binding the contract event 0x9839fd1896801c6864456fe7cbd098b1e4a81dba19536764cea53a1fb07ed239.
//
// Solidity: event NewLoan(bytes32 indexed channelId, uint256 loanAmount)
func (_AccountantImplementation *AccountantImplementationFilterer) FilterNewLoan(opts *bind.FilterOpts, channelId [][32]byte) (*AccountantImplementationNewLoanIterator, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _AccountantImplementation.contract.FilterLogs(opts, "NewLoan", channelIdRule)
	if err != nil {
		return nil, err
	}
	return &AccountantImplementationNewLoanIterator{contract: _AccountantImplementation.contract, event: "NewLoan", logs: logs, sub: sub}, nil
}

// WatchNewLoan is a free log subscription operation binding the contract event 0x9839fd1896801c6864456fe7cbd098b1e4a81dba19536764cea53a1fb07ed239.
//
// Solidity: event NewLoan(bytes32 indexed channelId, uint256 loanAmount)
func (_AccountantImplementation *AccountantImplementationFilterer) WatchNewLoan(opts *bind.WatchOpts, sink chan<- *AccountantImplementationNewLoan, channelId [][32]byte) (event.Subscription, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _AccountantImplementation.contract.WatchLogs(opts, "NewLoan", channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountantImplementationNewLoan)
				if err := _AccountantImplementation.contract.UnpackLog(event, "NewLoan", log); err != nil {
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

// ParseNewLoan is a log parse operation binding the contract event 0x9839fd1896801c6864456fe7cbd098b1e4a81dba19536764cea53a1fb07ed239.
//
// Solidity: event NewLoan(bytes32 indexed channelId, uint256 loanAmount)
func (_AccountantImplementation *AccountantImplementationFilterer) ParseNewLoan(log types.Log) (*AccountantImplementationNewLoan, error) {
	event := new(AccountantImplementationNewLoan)
	if err := _AccountantImplementation.contract.UnpackLog(event, "NewLoan", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountantImplementationOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AccountantImplementation contract.
type AccountantImplementationOwnershipTransferredIterator struct {
	Event *AccountantImplementationOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AccountantImplementationOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountantImplementationOwnershipTransferred)
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
		it.Event = new(AccountantImplementationOwnershipTransferred)
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
func (it *AccountantImplementationOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountantImplementationOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountantImplementationOwnershipTransferred represents a OwnershipTransferred event raised by the AccountantImplementation contract.
type AccountantImplementationOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AccountantImplementation *AccountantImplementationFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AccountantImplementationOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AccountantImplementation.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AccountantImplementationOwnershipTransferredIterator{contract: _AccountantImplementation.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AccountantImplementation *AccountantImplementationFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AccountantImplementationOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AccountantImplementation.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountantImplementationOwnershipTransferred)
				if err := _AccountantImplementation.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AccountantImplementation *AccountantImplementationFilterer) ParseOwnershipTransferred(log types.Log) (*AccountantImplementationOwnershipTransferred, error) {
	event := new(AccountantImplementationOwnershipTransferred)
	if err := _AccountantImplementation.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountantImplementationPromiseSettledIterator is returned from FilterPromiseSettled and is used to iterate over the raw logs and unpacked data for PromiseSettled events raised by the AccountantImplementation contract.
type AccountantImplementationPromiseSettledIterator struct {
	Event *AccountantImplementationPromiseSettled // Event containing the contract specifics and raw log

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
func (it *AccountantImplementationPromiseSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountantImplementationPromiseSettled)
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
		it.Event = new(AccountantImplementationPromiseSettled)
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
func (it *AccountantImplementationPromiseSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountantImplementationPromiseSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountantImplementationPromiseSettled represents a PromiseSettled event raised by the AccountantImplementation contract.
type AccountantImplementationPromiseSettled struct {
	ChannelId    [32]byte
	Beneficiary  common.Address
	Amount       *big.Int
	TotalSettled *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPromiseSettled is a free log retrieval operation binding the contract event 0xa5a1f05785a942c5f624cee545c68394881a83bcaf21a83f4d76a9e8240a5668.
//
// Solidity: event PromiseSettled(bytes32 indexed channelId, address beneficiary, uint256 amount, uint256 totalSettled)
func (_AccountantImplementation *AccountantImplementationFilterer) FilterPromiseSettled(opts *bind.FilterOpts, channelId [][32]byte) (*AccountantImplementationPromiseSettledIterator, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _AccountantImplementation.contract.FilterLogs(opts, "PromiseSettled", channelIdRule)
	if err != nil {
		return nil, err
	}
	return &AccountantImplementationPromiseSettledIterator{contract: _AccountantImplementation.contract, event: "PromiseSettled", logs: logs, sub: sub}, nil
}

// WatchPromiseSettled is a free log subscription operation binding the contract event 0xa5a1f05785a942c5f624cee545c68394881a83bcaf21a83f4d76a9e8240a5668.
//
// Solidity: event PromiseSettled(bytes32 indexed channelId, address beneficiary, uint256 amount, uint256 totalSettled)
func (_AccountantImplementation *AccountantImplementationFilterer) WatchPromiseSettled(opts *bind.WatchOpts, sink chan<- *AccountantImplementationPromiseSettled, channelId [][32]byte) (event.Subscription, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _AccountantImplementation.contract.WatchLogs(opts, "PromiseSettled", channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountantImplementationPromiseSettled)
				if err := _AccountantImplementation.contract.UnpackLog(event, "PromiseSettled", log); err != nil {
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

// ParsePromiseSettled is a log parse operation binding the contract event 0xa5a1f05785a942c5f624cee545c68394881a83bcaf21a83f4d76a9e8240a5668.
//
// Solidity: event PromiseSettled(bytes32 indexed channelId, address beneficiary, uint256 amount, uint256 totalSettled)
func (_AccountantImplementation *AccountantImplementationFilterer) ParsePromiseSettled(log types.Log) (*AccountantImplementationPromiseSettled, error) {
	event := new(AccountantImplementationPromiseSettled)
	if err := _AccountantImplementation.contract.UnpackLog(event, "PromiseSettled", log); err != nil {
		return nil, err
	}
	return event, nil
}
