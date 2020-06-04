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

// AccountantImplementationABI is the input ABI used to generate the binding from.
const AccountantImplementationABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"AccountantClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newFee\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"validFromBlock\",\"type\":\"uint64\"}],\"name\":\"AccountantFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activationBlock\",\"type\":\"uint256\"}],\"name\":\"AccountantPunishmentActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"AccountantPunishmentDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newStake\",\"type\":\"uint256\"}],\"name\":\"AccountantStakeIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"}],\"name\":\"ChannelBalanceDecreaseRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"ChannelBalanceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newBeneficiary\",\"type\":\"address\"}],\"name\":\"ChannelBeneficiaryChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"initialBalance\",\"type\":\"uint256\"}],\"name\":\"ChannelOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ChannelOpeningActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ChannelOpeningPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousDestination\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newDestination\",\"type\":\"address\"}],\"name\":\"DestinationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"FundsWithdrawned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaxStake\",\"type\":\"uint256\"}],\"name\":\"MaxStakeValueUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMinStake\",\"type\":\"uint256\"}],\"name\":\"MinStakeValueUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"NewAccountantOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"}],\"name\":\"NewStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalSettled\",\"type\":\"uint256\"}],\"name\":\"PromiseSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newStakeGoal\",\"type\":\"uint256\"}],\"name\":\"StakeGoalUpdated\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"channels\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settled\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeGoal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUsedNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timelock\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimEthers\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"claimTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFundsDestination\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastFee\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"value\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"validFrom\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"previousFee\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"value\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"validFrom\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"punishment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"activationBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOperator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identity\",\"type\":\"address\"}],\"name\":\"getChannelId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAccountantStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStakeThresholds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStatus\",\"outputs\":[{\"internalType\":\"enumAccountantImplementation.Status\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_fee\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_maxStake\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identity\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountToLend\",\"type\":\"uint256\"}],\"name\":\"openChannel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identity\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_transactorFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_lock\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"settlePromise\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identity\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_transactorFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_lock\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"settleAndRebalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identity\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_transactorFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_lock\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_promiseSignature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_newBeneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"settleWithBeneficiary\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_transactorFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_lock\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_promiseSignature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_newStakeGoal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_goalChangeSignature\",\"type\":\"bytes\"}],\"name\":\"settleWithGoalIncrease\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_newBalance\",\"type\":\"uint256\"}],\"name\":\"updateChannelBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"}],\"name\":\"rebalanceChannel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_transactorFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_lock\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"settleIntoStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"increaseStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_transactorFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"decreaseStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_newStakeGoal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"updateStakeGoal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"resolveEmergency\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_newBeneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"setBeneficiary\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOperator\",\"type\":\"address\"}],\"name\":\"setAccountantOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newMaxStake\",\"type\":\"uint256\"}],\"name\":\"setMaxStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newMinStake\",\"type\":\"uint256\"}],\"name\":\"setMinStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_newFee\",\"type\":\"uint16\"}],\"name\":\"setAccountantFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"calculateAccountantFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_additionalStake\",\"type\":\"uint256\"}],\"name\":\"increaseAccountantStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"}],\"name\":\"isChannelOpened\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isAccountantActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pauseChannelOpening\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"activateChannelOpening\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"availableBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minimalExpectedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"closeAccountant\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"getStakeBack\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_newDestination\",\"type\":\"address\"}],\"name\":\"setFundsDestination\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AccountantImplementationBin is the compiled bytecode used for deploying new contracts.
var AccountantImplementationBin = "0x60806040819052600080546001600160a01b03191633178082556001600160a01b0316917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3614a1a806100576000396000f3fe608060405234801561001057600080fd5b50600436106102bb5760003560e01c806394c7915d11610182578063df8de3e7116100e9578063f2fde38b116100a2578063f7d763691161007c578063f7d7636914610de1578063fbb46b9814610de9578063fc0c546a14610df1578063fec8157d14610df9576102bb565b8063f2fde38b14610d87578063f3fef3a314610dad578063f58c5b6e14610dd9576102bb565b8063df8de3e714610c57578063e1c6648714610c7d578063e7f43c6814610c85578063ea48d8bf14610c8d578063eb295b2714610d44578063efde05ec14610d6a576102bb565b8063ac7d01b61161013b578063ac7d01b6146108aa578063b07d948314610966578063be02c06c14610a2b578063c062fe0614610a33578063cc44905414610aef578063d1439a4d14610c4f576102bb565b806394c7915d1461082c5780639801134e146108465780639ed9903e146108755780639fe827ed1461087d578063aa606dee1461089a578063ab2f0e51146108a2576102bb565b80635f59def0116102265780637a7ebd7b116101df5780637a7ebd7b146107575780637c2be0a3146107b6578063800d6afb146107d95780638c80fd90146107ff5780638da5cb5b1461081c5780638f32d59b14610824576102bb565b80635f59def0146106275780636138dda7146106e75780636931b5501461070d5780636e9094ea146107155780636fc1483714610732578063715018a61461074f576102bb565b8063238e130a11610278578063238e130a146104ad578063392e53cd146104d357806339f97626146104ef57806348d9f01e146105125780634e69d560146105d75780635ab1bd5314610603576102bb565b80630684cd20146102c05780630996fcbc146102e15780630a798f24146103045780630b9a91c11461033a57806315c73afd146104885780631822af6f14610490575b600080fd5b6102c8610e39565b6040805192835260208301919091528051918290030190f35b610302600480360360208110156102f757600080fd5b503561ffff16610e42565b005b6103026004803603606081101561031a57600080fd5b506001600160a01b03813581169160208101359091169060400135611068565b610302600480360361010081101561035157600080fd5b81359160208101359160408201359160608101359181019060a081016080820135600160201b81111561038357600080fd5b82018360208201111561039557600080fd5b803590602001918460018302840111600160201b831117156103b657600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295843595602086013595919450925060608101915060400135600160201b81111561041457600080fd5b82018360208201111561042657600080fd5b803590602001918460018302840111600160201b8311171561044757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611136945050505050565b610302611159565b610302600480360360208110156104a657600080fd5b50356113d8565b610302600480360360208110156104c357600080fd5b50356001600160a01b0316611522565b6104db6115da565b604080519115158252519081900360200190f35b6103026004803603604081101561050557600080fd5b50803590602001356115ec565b610302600480360360a081101561052857600080fd5b6001600160a01b038235169160208101359160408201359160608101359181019060a081016080820135600160201b81111561056357600080fd5b82018360208201111561057557600080fd5b803590602001918460018302840111600160201b8311171561059657600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506116fa945050505050565b6105df6117bc565b604051808260038111156105ef57fe5b60ff16815260200191505060405180910390f35b61060b6117c5565b604080516001600160a01b039092168252519081900360200190f35b6103026004803603608081101561063d57600080fd5b8135916001600160a01b036020820135169160408201359190810190608081016060820135600160201b81111561067357600080fd5b82018360208201111561068557600080fd5b803590602001918460018302840111600160201b831117156106a657600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506117d4945050505050565b610302600480360360208110156106fd57600080fd5b50356001600160a01b03166119c1565b610302611bc2565b6104db6004803603602081101561072b57600080fd5b5035611c13565b6103026004803603602081101561074857600080fd5b5035611c30565b610302611d0d565b6107746004803603602081101561076d57600080fd5b5035611db0565b604080516001600160a01b0390981688526020880196909652868601949094526060860192909252608085015260a084015260c0830152519081900360e00190f35b610302600480360360408110156107cc57600080fd5b5080359060200135611df7565b610302600480360360208110156107ef57600080fd5b50356001600160a01b0316612097565b6103026004803603602081101561081557600080fd5b5035612187565b61060b6122a4565b6104db6122b3565b6108346122c4565b60408051918252519081900360200190f35b61084e6122e2565b6040805161ffff909316835267ffffffffffffffff90911660208301528051918290030190f35b6102c86122fe565b6108346004803603602081101561089357600080fd5b5035612308565b610302612385565b610834612457565b610302600480360360a08110156108c057600080fd5b81359160208101359160408201359160608101359181019060a081016080820135600160201b8111156108f257600080fd5b82018360208201111561090457600080fd5b803590602001918460018302840111600160201b8311171561092557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550612550945050505050565b610302600480360360a081101561097c57600080fd5b6001600160a01b038235169160208101359160408201359160608101359181019060a081016080820135600160201b8111156109b757600080fd5b8201836020820111156109c957600080fd5b803590602001918460018302840111600160201b831117156109ea57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550612aff945050505050565b6104db612bc2565b610302600480360360a0811015610a4957600080fd5b81359160208101359160408201359160608101359181019060a081016080820135600160201b811115610a7b57600080fd5b820183602082011115610a8d57600080fd5b803590602001918460018302840111600160201b83111715610aae57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550612bfc945050505050565b6103026004803603610100811015610b0657600080fd5b6001600160a01b038235169160208101359160408201359160608101359181019060a081016080820135600160201b811115610b4157600080fd5b820183602082011115610b5357600080fd5b803590602001918460018302840111600160201b83111715610b7457600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092956001600160a01b0385351695602086013595919450925060608101915060400135600160201b811115610bdb57600080fd5b820183602082011115610bed57600080fd5b803590602001918460018302840111600160201b83111715610c0e57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550612eab945050505050565b610834612ef3565b61030260048036036020811015610c6d57600080fd5b50356001600160a01b0316612ef9565b61084e61302a565b61060b613046565b61030260048036036080811015610ca357600080fd5b81359160208101359160408201359190810190608081016060820135600160201b811115610cd057600080fd5b820183602082011115610ce257600080fd5b803590602001918460018302840111600160201b83111715610d0357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550613055945050505050565b61083460048036036020811015610d5a57600080fd5b50356001600160a01b031661328f565b61030260048036036020811015610d8057600080fd5b50356132d3565b61030260048036036020811015610d9d57600080fd5b50356001600160a01b03166134e3565b61030260048036036040811015610dc357600080fd5b506001600160a01b038135169060200135613545565b61060b6136f6565b610302613705565b6103026137f4565b61060b6138c3565b61030260048036036080811015610e0f57600080fd5b506001600160a01b03813581169160208101359091169061ffff60408201351690606001356138d2565b600f5460105482565b6004546001600160a01b03163314610e8b5760405162461bcd60e51b81526004018080602001828103825260248152602001806145df6024913960400191505060405180910390fd5b6003610e956117bc565b6003811115610ea057fe5b1415610ef3576040805162461bcd60e51b815260206004820152601f60248201527f6163636f756e74616e742073686f756c64206265206e6f7420636c6f73656400604482015290519081900360640190fd5b6113888161ffff161115610f4e576040805162461bcd60e51b815260206004820152601c60248201527f6665652063616e27742062652062696767657220746861742035302500000000604482015290519081900360640190fd5b600c5462010000900467ffffffffffffffff16431015610fb5576040805162461bcd60e51b815260206004820152601960248201527f63616e27742075706461746520696e6163746976652066656500000000000000604482015290519081900360640190fd5b6000610fbf613b28565b600c8054600d805467ffffffffffffffff620100008085048216810269ffffffffffffffff00001961ffff1994851661ffff80891691909117821692909217909555604080518082018252918b1680835293891660209283018190529283029490961683179094169290921790945582519384529083015280519293507e8b9bffa5c48d8c0b96ce879f8491c5605cc3d7a46a46711b522dbe6d4070ee92918290030190a15050565b6003546001600160a01b031633146110c7576040805162461bcd60e51b815260206004820152601f60248201527f6f6e6c792072656769737472792063616e206f70656e206368616e6e656c7300604482015290519081900360640190fd5b60006110d16117bc565b60038111156110dc57fe5b146111185760405162461bcd60e51b81526004018080602001828103825260218152602001806147526021913960400191505060405180910390fd5b60006111238461328f565b9050611130818484613b30565b50505050565b61114288848484613055565b61114f8888888888613bfe565b5050505050505050565b60026111636117bc565b600381111561116e57fe5b146111aa5760405162461bcd60e51b815260040180806020018281038252602981526020018061483d6029913960400191505060405180910390fd5b60006111dd60646111d16111ca6004600554613fde90919063ffffffff16565b6064614040565b9063ffffffff61405a16565b905060006111e96140c4565b600f54909150430360006112176001846112038582614040565b8161120a57fe5b049063ffffffff6140ca16565b9050600061122b828663ffffffff613fde16565b601054909150611241908263ffffffff61412716565b60108190556009546000916112739161125991614181565b611267600554600654614181565b9063ffffffff61412716565b600254604080516370a0823160e01b815230600482015290519293506000926001600160a01b03909216916370a0823191602480820192602092909190829003018186803b1580156112c457600080fd5b505afa1580156112d8573d6000803e3d6000fd5b505050506040513d60208110156112ee57600080fd5b505190506000828210611302576000611312565b611312838363ffffffff6140ca16565b600254604080516323b872dd60e01b81523360048201523060248201526044810184905290519293506001600160a01b03909116916323b872dd916064808201926020929091908290030181600087803b15801561136f57600080fd5b505af1158015611383573d6000803e3d6000fd5b505050506040513d602081101561139957600080fd5b5050600b805460ff191690556040517f58ef313a2eb2567f3b143ff20930622dd67a0de84902cc93b7ddddd72b7773ef90600090a15050505050505050565b6004546001600160a01b031633146114215760405162461bcd60e51b81526004018080602001828103825260248152602001806145df6024913960400191505060405180910390fd5b8061142a612457565b10156114d557600061144a61143d612457565b839063ffffffff6140ca16565b600254604080516323b872dd60e01b81523360048201523060248201526044810184905290519293506001600160a01b03909116916323b872dd916064808201926020929091908290030181600087803b1580156114a757600080fd5b505af11580156114bb573d6000803e3d6000fd5b505050506040513d60208110156114d157600080fd5b5050505b6009546114e8908263ffffffff61412716565b600981905560408051918252517f41a5bb80f9c1243f3d450690277c955ff8982168e34ed096457afdc31cefef7f9181900360200190a150565b6004546001600160a01b0316331461156b5760405162461bcd60e51b81526004018080602001828103825260248152602001806145df6024913960400191505060405180910390fd5b6001600160a01b03811661157e57600080fd5b6001546040516001600160a01b038084169216907fe1a66d77649cf0a57b9937073549f30f1c82bb865aaf066d2f299e37a62c6aad90600090a3600180546001600160a01b0319166001600160a01b0392909216919091179055565b6004546001600160a01b031615155b90565b6115f582611c13565b611641576040805162461bcd60e51b815260206004820152601860248201527718da185b9b995b081a185cc81d1bc81899481bdc195b995960421b604482015290519081900360640190fd5b600361164b6117bc565b600381111561165657fe5b14156116a9576040805162461bcd60e51b815260206004820152601f60248201527f6163636f756e74616e742073686f756c64206265206e6f7420636c6f73656400604482015290519081900360640190fd5b6116b582826000614197565b6000828152600e6020908152604091829020600381015460018201819055835190815292519092859260008051602061470783398151915292918290030190a2505050565b60006117058661328f565b905061171081611c13565b6117a75760035460408051633985eaab60e21b81526001600160a01b0389811660048301523060248301529151600093929092169163e617aaac91604480820192602092909190829003018186803b15801561176b57600080fd5b505afa15801561177f573d6000803e3d6000fd5b505050506040513d602081101561179557600080fd5b505190506117a582826000613b30565b505b6117b48186868686613bfe565b505050505050565b600b5460ff1690565b6003546001600160a01b031690565b6117dd84611c13565b611829576040805162461bcd60e51b815260206004820152601860248201527718da185b9b995b081a185cc81d1bc81899481bdc195b995960421b604482015290519081900360640190fd5b6001600160a01b03831661186e5760405162461bcd60e51b81526004018080602001828103825260218152602001806147736021913960400191505060405180910390fd5b6000848152600e60205260409020600581015483116118be5760405162461bcd60e51b81526004018080602001828103825260298152602001806146de6029913960400191505060405180910390fd5b6040805160208082018890526bffffffffffffffffffffffff19606088901b168284015260548083018790528351808403909101815260749092019092528051910120600090611914908463ffffffff6143b716565b9050856119208261328f565b1461195c5760405162461bcd60e51b81526004018080602001828103825260228152602001806148b76022913960400191505060405180910390fd5b6005820184905581546001600160a01b0319166001600160a01b038616908117835560408051888152602081019290925280517f8756aa559142225f918d7584303ecfe48e75b454f6614d0fae9f0d6ca0a898cc9281900390910190a1505050505050565b6004546001600160a01b03163314611a0a5760405162461bcd60e51b81526004018080602001828103825260248152602001806145df6024913960400191505060405180910390fd5b6003611a146117bc565b6003811115611a1f57fe5b14611a71576040805162461bcd60e51b815260206004820152601c60248201527f6163636f756e74616e74206861766520746f20626520636c6f73656400000000604482015290519081900360640190fd5b600a544311611ab15760405162461bcd60e51b81526004018080602001828103825260268152602001806147946026913960400191505060405180910390fd5b601054600254604080516370a0823160e01b81523060048201529051600093611b409390926001600160a01b03909116916370a0823191602480820192602092909190829003018186803b158015611b0857600080fd5b505afa158015611b1c573d6000803e3d6000fd5b505050506040513d6020811015611b3257600080fd5b50519063ffffffff6140ca16565b6002546040805163a9059cbb60e01b81526001600160a01b03868116600483015260248201859052915193945091169163a9059cbb916044808201926020929091908290030181600087803b158015611b9857600080fd5b505af1158015611bac573d6000803e3d6000fd5b505050506040513d602081101561113057600080fd5b6001546001600160a01b0316611bd757600080fd5b6001546040516001600160a01b03909116904780156108fc02916000818181858888f19350505050158015611c10573d6000803e3d6000fd5b50565b6000908152600e60205260409020546001600160a01b0316151590565b6004546001600160a01b03163314611c795760405162461bcd60e51b81526004018080602001828103825260248152602001806145df6024913960400191505060405180910390fd5b611c81612bc2565b611cd2576040805162461bcd60e51b815260206004820152601b60248201527f6163636f756e74616e742068617320746f206265206163746976650000000000604482015290519081900360640190fd5b60088190556040805182815290517f53f4fb18cb329155d5af04681c1d0846d0484d7de33791619c6988ca61910e3d9181900360200190a150565b611d156122b3565b611d66576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b600e6020526000908152604090208054600182015460028301546003840154600485015460058601546006909601546001600160a01b039095169593949293919290919087565b6004546001600160a01b03163314611e405760405162461bcd60e51b81526004018080602001828103825260248152602001806145df6024913960400191505060405180910390fd5b611e48612bc2565b611e99576040805162461bcd60e51b815260206004820152601b60248201527f6163636f756e74616e742068617320746f206265206163746976650000000000604482015290519081900360640190fd5b611ea282611c13565b611eee576040805162461bcd60e51b815260206004820152601860248201527718da185b9b995b081a185cc81d1bc81899481bdc195b995960421b604482015290519081900360640190fd5b6000828152600e6020526040902060030154811015611f3e5760405162461bcd60e51b81526004018080602001828103825260278152602001806145756027913960400191505060405180910390fd5b6000828152600e602052604081206001810154909190831115611fd4576001820154611f7190849063ffffffff6140ca16565b905080611f7c612457565b1015611fb95760405162461bcd60e51b815260040180806020018281038252602281526020018061459c6022913960400191505060405180910390fd5b600554611fcc908263ffffffff61412716565b600555612065565b600682015461201c57611fe5613b28565b600683015560405184907faf4c616dc7856b81dbc1346e5547f0a1d4f1553011653f920d1041f21540107590600090a25050612093565b816006015443101561202f575050612093565b6001820154612044908463ffffffff6140ca16565b60055490915061205a908263ffffffff6140ca16565b600555600060068301555b600182018390556040805184815290518591600080516020614707833981519152919081900360200190a250505b5050565b6004546001600160a01b031633146120e05760405162461bcd60e51b81526004018080602001828103825260248152602001806145df6024913960400191505060405180910390fd5b6001600160a01b038116612133576040805162461bcd60e51b815260206004820152601560248201527463616e2774206265207a65726f206164647265737360581b604482015290519081900360640190fd5b600480546001600160a01b0383166001600160a01b0319909116811790915560408051918252517fa326c787c51b80574c7b572d0c9664e64f1107538b902f519a896901b41379189181900360200190a150565b6004546001600160a01b031633146121d05760405162461bcd60e51b81526004018080602001828103825260248152602001806145df6024913960400191505060405180910390fd5b6121d8612bc2565b612229576040805162461bcd60e51b815260206004820152601b60248201527f6163636f756e74616e742068617320746f206265206163746976650000000000604482015290519081900360640190fd5b60085481106122695760405162461bcd60e51b815260040180806020018281038252602a81526020018061468f602a913960400191505060405180910390fd5b60078190556040805182815290517fb9e5e6e8db1283ee860f3856d8383e40665c58a5264ede5e6ed8ec1afb0312519181900360200190a150565b6000546001600160a01b031690565b6000546001600160a01b0316331490565b60006122dd600554611267600954600f60010154614181565b905090565b600c5461ffff81169062010000900467ffffffffffffffff1682565b6007546008549091565b600061231261455d565b600c5462010000900467ffffffffffffffff1643101561233357600d612336565b600c5b60408051808201909152905461ffff81168083526201000090910467ffffffffffffffff1660208301529091506064906123769082908602046064614040565b8161237d57fe5b049392505050565b6004546001600160a01b031633146123ce5760405162461bcd60e51b81526004018080602001828103825260248152602001806145df6024913960400191505060405180910390fd5b60006123d86117bc565b60038111156123e357fe5b1461241f5760405162461bcd60e51b81526004018080602001828103825260258152602001806149706025913960400191505060405180910390fd5b600b805460ff191660011790556040517f1f4cd5d6edef8a0c4dbe6d547fdc42e0f3575167257553271f2366f9d497f67e90600090a1565b60008061246e611259600954600f60010154614181565b600254604080516370a0823160e01b815230600482015290519293506001600160a01b03909116916370a0823191602480820192602092909190829003018186803b1580156124bc57600080fd5b505afa1580156124d0573d6000803e3d6000fd5b505050506040513d60208110156124e657600080fd5b50518111156124f95760009150506115e9565b600254604080516370a0823160e01b8152306004820152905161254a9284926001600160a01b03909116916370a0823191602480820192602092909190829003018186803b158015611b0857600080fd5b91505090565b6000612611826040518060400160405280601481526020017314dd185ad9481c995d1d5c9b881c995c5d595cdd60621b815250888888886040516020018086805190602001908083835b602083106125b95780518252601f19909201916020918201910161259a565b51815160209384036101000a60001901801990921691161790529201968752508581019490945250604080850192909252606080850191909152815180850390910181526080909301905281519101209190506143b7565b90508561261d8261328f565b146126595760405162461bcd60e51b81526004018080602001828103825260228152602001806148b76022913960400191505060405180910390fd5b61266286611c13565b6126ae576040805162461bcd60e51b815260206004820152601860248201527718da185b9b995b081a185cc81d1bc81899481bdc195b995960421b604482015290519081900360640190fd5b6000868152600e60205260409020600581015484116126fe5760405162461bcd60e51b81526004018080602001828103825260288152602001806148d96028913960400191505060405180910390fd5b6005810184905560038101548611156127485760405162461bcd60e51b815260040180806020018281038252602a815260200180614629602a913960400191505060405180910390fd5b848610156127875760405162461bcd60e51b815260040180806020018281038252602b8152602001806147ba602b913960400191505060405180910390fd5b60006127978260010154886144a5565b905060006127b3826127a76122c4565b9063ffffffff6140ca16565b600254604080516370a0823160e01b815230600482015290519293506000926001600160a01b03909216916370a0823191602480820192602092909190829003018186803b15801561280457600080fd5b505afa158015612818573d6000803e3d6000fd5b505050506040513d602081101561282e57600080fd5b505190508089118061284e57508161284c828b63ffffffff6140ca16565b105b156128ba5761285b612bc2565b156128a757600b805460ff1916600217905543600f81905560408051918252517fb3e91d0895882cef621b468a8235043537ca2e4d8d91ee6587801041054107e1916020908290030190a15b6128b7818363ffffffff6140ca16565b98505b60038401546000906128d2908b63ffffffff6140ca16565b90506008548111156129155760405162461bcd60e51b81526004018080602001828103825260338152602001806149016033913960400191505060405180910390fd5b881561299c576002546040805163a9059cbb60e01b8152336004820152602481018c905290516001600160a01b039092169163a9059cbb916044808201926020929091908290030181600087803b15801561296f57600080fd5b505af1158015612983573d6000803e3d6000fd5b505050506040513d602081101561299957600080fd5b50505b60025485546001600160a01b039182169163a9059cbb91166129c48d8d63ffffffff6140ca16565b6040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050602060405180830381600087803b158015612a1357600080fd5b505af1158015612a27573d6000803e3d6000fd5b505050506040513d6020811015612a3d57600080fd5b5050600385018190556001850154612a5b908563ffffffff6140ca16565b60018601556007546004860155600554612a7b908563ffffffff6140ca16565b600555600654612a91908b63ffffffff6140ca16565b600655600185015460408051918252518c91600080516020614707833981519152919081900360200190a26040805182815290518c917fc5f0715c45dab2e8f14871936119e3c64fd5841d397130c2d1db743d142522cb919081900360200190a25050505050505050505050565b6000612b0a8661328f565b9050612b1581611c13565b612bac5760035460408051633985eaab60e21b81526001600160a01b0389811660048301523060248301529151600093929092169163e617aaac91604480820192602092909190829003018186803b158015612b7057600080fd5b505afa158015612b84573d6000803e3d6000fd5b505050506040513d6020811015612b9a57600080fd5b50519050612baa82826000613b30565b505b612bb98186868686613bfe565b6117b4816132d3565b600080612bcd6117bc565b90506002816003811115612bdd57fe5b1415801561254a57506003816003811115612bf457fe5b141591505090565b612c0585611c13565b612c52576040805162461bcd60e51b815260206004820152601960248201527818da185b9b995b081a185d99481d1bc81899481bdc195b9959603a1b604482015290519081900360640190fd5b6000858152600e6020908152604080832081518084018790528251808203850181528184018452805190850120606082018b9052608082018a905260a0820189905260c08083018290528451808403909101815260e0909201909352805193019290922091929091612cca908563ffffffff6143b716565b6004549091506001600160a01b03808316911614612d2f576040805162461bcd60e51b815260206004820152601d60248201527f6861766520746f206265207369676e6564206279206f70657261746f72000000604482015290519081900360640190fd5b6000612d488460020154896140ca90919063ffffffff16565b905080871115612d895760405162461bcd60e51b815260040180806020018281038252603c815260200180614653603c913960400191505060405180910390fd5b612da489612d9d838a63ffffffff6140ca16565b6001614197565b6002840154612db9908263ffffffff61412716565b600285018190558454604080516001600160a01b039092168252602082018490528181019290925290518a917fa5a1f05785a942c5f624cee545c68394881a83bcaf21a83f4d76a9e8240a5668919081900360600190a28615612e97576002546040805163a9059cbb60e01b8152336004820152602481018a905290516001600160a01b039092169163a9059cbb916044808201926020929091908290030181600087803b158015612e6a57600080fd5b505af1158015612e7e573d6000803e3d6000fd5b505050506040513d6020811015612e9457600080fd5b50505b612ea0896132d3565b505050505050505050565b6000612eb68961328f565b9050612ec181611c13565b612ed157612ed181856000613b30565b612edd818585856117d4565b612eea8189898989613bfe565b612ea0816132d3565b60095490565b6001546001600160a01b0316612f0e57600080fd5b6002546001600160a01b0382811691161415612f5b5760405162461bcd60e51b81526004018080602001828103825260258152602001806146b96025913960400191505060405180910390fd5b604080516370a0823160e01b815230600482015290516000916001600160a01b038416916370a0823191602480820192602092909190829003018186803b158015612fa557600080fd5b505afa158015612fb9573d6000803e3d6000fd5b505050506040513d6020811015612fcf57600080fd5b50516001546040805163a9059cbb60e01b81526001600160a01b0392831660048201526024810184905290519293509084169163a9059cbb916044808201926020929091908290030181600087803b158015611b9857600080fd5b600d5461ffff81169062010000900467ffffffffffffffff1682565b6004546001600160a01b031690565b61305e84611c13565b6130ab576040805162461bcd60e51b815260206004820152601960248201527818da185b9b995b081a185d99481d1bc81899481bdc195b9959603a1b604482015290519081900360640190fd5b6007548310156130ec5760405162461bcd60e51b815260040180806020018281038252602b815260200180614727602b913960400191505060405180910390fd5b6000848152600e602052604090206005810154831161313c5760405162461bcd60e51b81526004018080602001828103825260298152602001806146de6029913960400191505060405180910390fd5b60006131fb836040518060400160405280601981526020017f5374616b6520676f616c207570646174652072657175657374000000000000008152508888886040516020018085805190602001908083835b602083106131ad5780518252601f19909201916020918201910161318e565b51815160209384036101000a600019018019909216911617905292019586525084810193909352506040808401919091528051808403820181526060909301905281519101209190506143b7565b9050856132078261328f565b146132435760405162461bcd60e51b81526004018080602001828103825260228152602001806148b76022913960400191505060405180910390fd5b600582018490556004820185905560408051868152905187917f5ec1b60bc705196753b2e21e144bbb185e6ef7264f9d79b99f04f1bc22198733919081900360200190a2505050505050565b60408051606092831b6bffffffffffffffffffffffff19166020808301919091523090931b6034820152815180820360280181526048909101909152805191012090565b6132db612bc2565b61332c576040805162461bcd60e51b815260206004820152601c60248201527f6163636f756e74616e74206861766520746f2062652061637469766500000000604482015290519081900360640190fd5b6000818152600e6020526040902060018101546003820154116133805760405162461bcd60e51b815260040180806020018281038252602981526020018061488e6029913960400191505060405180910390fd5b600061339d826001015483600301546140ca90919063ffffffff16565b905060006133ad826112676122c4565b600254604080516370a0823160e01b815230600482015290519293506000926001600160a01b03909216916370a0823191602480820192602092909190829003018186803b1580156133fe57600080fd5b505afa158015613412573d6000803e3d6000fd5b505050506040513d602081101561342857600080fd5b505190508181101561348757600b805460ff1916600217905543600f5561344f82826140ca565b6040805143815290519194507fb3e91d0895882cef621b468a8235043537ca2e4d8d91ee6587801041054107e1919081900360200190a15b60055461349a908463ffffffff61412716565b60055560018401546134b2908463ffffffff61412716565b6001850181905560408051918252518691600080516020614707833981519152919081900360200190a25050505050565b6134eb6122b3565b61353c576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b611c10816144b4565b6004546001600160a01b0316331461358e5760405162461bcd60e51b81526004018080602001828103825260248152602001806145df6024913960400191505060405180910390fd5b613596612bc2565b6135e7576040805162461bcd60e51b815260206004820152601c60248201527f6163636f756e74616e74206861766520746f2062652061637469766500000000604482015290519081900360640190fd5b806135f0612457565b101561362d5760405162461bcd60e51b815260040180806020018281038252602c815260200180614995602c913960400191505060405180910390fd5b6002546040805163a9059cbb60e01b81526001600160a01b038581166004830152602482018590529151919092169163a9059cbb9160448083019260209291908290030181600087803b15801561368357600080fd5b505af1158015613697573d6000803e3d6000fd5b505050506040513d60208110156136ad57600080fd5b5050604080518281526001600160a01b038416602082015281517fa2e147ce2b7cb83d9c07e397bb806f23dd42c42e86ea45e1611d6e50eb1ec8bf929181900390910190a15050565b6001546001600160a01b031690565b6004546001600160a01b0316331461374e5760405162461bcd60e51b81526004018080602001828103825260248152602001806145df6024913960400191505060405180910390fd5b613756612bc2565b6137a7576040805162461bcd60e51b815260206004820152601b60248201527f6163636f756e74616e742073686f756c64206265206163746976650000000000604482015290519081900360640190fd5b600b805460ff191660031790556137bc614554565b600a556040805143815290517f888906f0892e56365e679111a6f8ba7d0742bae70d0a532641cbf0da77d5af929181900360200190a1565b6004546001600160a01b0316331461383d5760405162461bcd60e51b81526004018080602001828103825260248152602001806145df6024913960400191505060405180910390fd5b60016138476117bc565b600381111561385257fe5b1461388e5760405162461bcd60e51b81526004018080602001828103825260258152602001806149c16025913960400191505060405180910390fd5b600b805460ff191690556040517f2d8b6ec230798e206d536342a28b7b61cc8fcfafb1d27c11c5519b3c42eb7df890600090a1565b6002546001600160a01b031681565b6138da6115da565b1561392c576040805162461bcd60e51b815260206004820152601a60248201527f6861766520746f206265206e6f7420696e697469616c697a6564000000000000604482015290519081900360640190fd5b6001600160a01b038316613987576040805162461bcd60e51b815260206004820152601760248201527f6f70657261746f72206861766520746f20626520736574000000000000000000604482015290519081900360640190fd5b6001600160a01b0384166139cc5760405162461bcd60e51b81526004018080602001828103825260288152602001806148666028913960400191505060405180910390fd5b6113888261ffff161115613a27576040805162461bcd60e51b815260206004820152601c60248201527f6665652063616e277420626520626967676572207468616e2035302500000000604482015290519081900360640190fd5b600280546001600160a01b03199081166001600160a01b03878116919091179283905560038054339084161790556004805490921686821617825560408051808201825261ffff87168082524367ffffffffffffffff166020928301819052600c805461ffff191690921769ffffffffffffffff00001916620100009091021790556305f5e100600755600886905581516370a0823160e01b81523094810194909452905193909116926370a08231926024808201939291829003018186803b158015613af357600080fd5b505afa158015613b07573d6000803e3d6000fd5b505050506040513d6020811015613b1d57600080fd5b505160095550505050565b436146500190565b613b3983611c13565b15613b755760405162461bcd60e51b81526004018080602001828103825260218152602001806145be6021913960400191505060405180910390fd5b8015613b8757613b8783826000614197565b6000838152600e602090815260409182902080546001600160a01b0319166001600160a01b0386161781556001810184905560075460048201558251868152918201849052825190927fbe2e1f3a6197dfd16fa6830c4870364b618b8b288c21cbcfa4fdb5d7c6a5e45b928290030190a150505050565b6000858152600e6020908152604080832081518084018790528251808203850181528184018452805190850120606082018b9052608082018a905260a0820189905260c08083018290528451808403909101815260e0909201909352805193019290922091929091613c76908563ffffffff6143b716565b6004549091506001600160a01b03808316911614613cdb576040805162461bcd60e51b815260206004820152601d60248201527f6861766520746f206265207369676e6564206279206f70657261746f72000000604482015290519081900360640190fd5b6000613cf48460020154896140ca90919063ffffffff16565b905060008111613d355760405162461bcd60e51b81526004018080602001828103825260378152602001806148066037913960400191505060405180910390fd5b80871115613d745760405162461bcd60e51b815260040180806020018281038252603c815260200180614653603c913960400191505060405180910390fd5b6000846004015485600301541015613d90578460040154613d96565b84600101545b905080821115613da4578091505b6002850154613db9908363ffffffff61412716565b60028601556000613dc983612308565b90506000613de1826127a7868d63ffffffff6140ca16565b9050866004015487600301541015613e40576000613e1a600a8304613e158a600301546008546140ca90919063ffffffff16565b6144a5565b9050613e288d826001614197565b613e38828263ffffffff6140ca16565b915050613e71565b6001870154613e55908563ffffffff6140ca16565b6001880155600554613e6d908563ffffffff6140ca16565b6005555b60025487546040805163a9059cbb60e01b81526001600160a01b039283166004820152602481018590529051919092169163a9059cbb9160448083019260209291908290030181600087803b158015613ec957600080fd5b505af1158015613edd573d6000803e3d6000fd5b505050506040513d6020811015613ef357600080fd5b50508915613f7c576002546040805163a9059cbb60e01b8152336004820152602481018d905290516001600160a01b039092169163a9059cbb916044808201926020929091908290030181600087803b158015613f4f57600080fd5b505af1158015613f63573d6000803e3d6000fd5b505050506040513d6020811015613f7957600080fd5b50505b86546002880154604080516001600160a01b0390931683526020830187905282810191909152518d917fa5a1f05785a942c5f624cee545c68394881a83bcaf21a83f4d76a9e8240a5668919081900360600190a2505050505050505050505050565b600082613fed5750600061403a565b82820282848281613ffa57fe5b04146140375760405162461bcd60e51b81526004018080602001828103825260218152602001806147e56021913960400191505060405180910390fd5b90505b92915050565b600081826001848601038161405157fe5b04029392505050565b60008082116140b0576040805162461bcd60e51b815260206004820152601a60248201527f536166654d6174683a206469766973696f6e206279207a65726f000000000000604482015290519081900360640190fd5b60008284816140bb57fe5b04949350505050565b61010190565b600082821115614121576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b600082820183811015614037576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b60008183116141905781614037565b5090919050565b600082116141ec576040805162461bcd60e51b815260206004820152601b60248201527f73686f756c64207374616b65206d6f7265207468616e207a65726f0000000000604482015290519081900360640190fd5b6000838152600e602052604081206003810154909190614212908563ffffffff61412716565b90506008548111156142555760405162461bcd60e51b815260040180806020018281038252603c815260200180614934603c913960400191505060405180910390fd5b8261433157600254604080516323b872dd60e01b81523360048201523060248201526044810187905290516001600160a01b03909216916323b872dd916064808201926020929091908290030181600087803b1580156142b457600080fd5b505af11580156142c8573d6000803e3d6000fd5b505050506040513d60208110156142de57600080fd5b5051614331576040805162461bcd60e51b815260206004820152601d60248201527f746f6b656e207472616e736665722073686f756c642073756363656564000000604482015290519081900360640190fd5b6003820181905560018201546143619061435290839063ffffffff6140ca16565b6005549063ffffffff61412716565b600555600654614377908563ffffffff61412716565b60065560408051828152905186917fc5f0715c45dab2e8f14871936119e3c64fd5841d397130c2d1db743d142522cb919081900360200190a25050505050565b600081516041146143ca5750600061403a565b60208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0821115614410576000935050505061403a565b8060ff16601b1415801561442857508060ff16601c14155b15614439576000935050505061403a565b6040805160008152602080820180845289905260ff8416828401526060820186905260808201859052915160019260a0808401939192601f1981019281900390910190855afa158015614490573d6000803e3d6000fd5b5050604051601f190151979650505050505050565b60008183106141905781614037565b6001600160a01b0381166144f95760405162461bcd60e51b81526004018080602001828103825260268152602001806146036026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b43621b77400190565b60408051808201909152600080825260208201529056fe62616c616e63652063616e2774206265206c657373207468616e207374616b6520616d6f756e7473686f756c6420626520656e6f75676820617661696c61626c652062616c616e63656368616e6e656c206861766520746f206265206e6f74206f70656e6564207965746f6e6c79206f70657261746f722063616e2063616c6c20746869732066756e6374696f6e4f776e61626c653a206e6577206f776e657220697320746865207a65726f206164647265737363616e2774207769746864726177206d6f7265207468616e207468652063757272656e74207374616b657472616e736163746f72206665652073686f756c6420626520657175616c20746f206f72206c657373207468616e205f756e70616964416d6f756e746d696e207374616b652068617320746f20626520736d616c6c6572207468616e206d6178207374616b656e617469766520746f6b656e2066756e64732063616e2774206265207265636f76657265646e6f6e6365206861766520746f20626520626967676572207468616e20616c726561647920757365642eb87c52e5d2d7395da77618aa1afe310b9cf13fa9f70e3199d100adda3a7aa57374616b6520676f616c2063616e2774206265206c657373207468616e206d696e696d616c207374616b656865726d6573206861766520746f20626520696e2061637469766520737461746562656e65666963696172792063616e2774206265207a65726f206164647265737374696d656c6f636b20706572696f64206861766520626520616c726561647920706173736564616d6f756e742073686f756c6420626520626967676572207468616e207472616e736163746f7220666565536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77616d6f756e7420746f20736574746c652073686f756c642062652067726561746572207468617420616c726561647920736574746c65646163636f756e74616e742073686f756c6420626520696e2070756e6973686d656e7420737461747573746f6b656e2063616e2774206265206465706c6f796420696e746f207a65726f20616464726573736e65772062616c616e63652073686f756c6420626520626967676572207468616e2063757272656e746861766520746f206265207369676e6564206279206368616e6e656c2070617274796e6f6e63652068617320746f20626520626967676572207468616e20616c72656164792075736564616d6f756e7420746f206c656e642063616e277420626520626967676572207468616e206d6178696d756d20616c6c6f776564746f74616c20616d6f756e7420746f207374616b652063616e277420626520626967676572207468616e206d6178696d616c6c7920616c6c6f7765646163636f756e74616e74206861766520746f20626520696e2061637469766520737461746573686f756c6420626520656e6f7567682066756e647320617661696c61626c6520746f2077697468647261776163636f756e74616e74206861766520746f20626520696e20706175736564207374617465a265627a7a72315820f3813336e19b940ec9599b9e564035f3fe756fed2a5dae434809368e9a7216df64736f6c63430005110032"

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

// CalculateAccountantFee is a free data retrieval call binding the contract method 0x9fe827ed.
//
// Solidity: function calculateAccountantFee(uint256 _amount) constant returns(uint256)
func (_AccountantImplementation *AccountantImplementationCaller) CalculateAccountantFee(opts *bind.CallOpts, _amount *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccountantImplementation.contract.Call(opts, out, "calculateAccountantFee", _amount)
	return *ret0, err
}

// CalculateAccountantFee is a free data retrieval call binding the contract method 0x9fe827ed.
//
// Solidity: function calculateAccountantFee(uint256 _amount) constant returns(uint256)
func (_AccountantImplementation *AccountantImplementationSession) CalculateAccountantFee(_amount *big.Int) (*big.Int, error) {
	return _AccountantImplementation.Contract.CalculateAccountantFee(&_AccountantImplementation.CallOpts, _amount)
}

// CalculateAccountantFee is a free data retrieval call binding the contract method 0x9fe827ed.
//
// Solidity: function calculateAccountantFee(uint256 _amount) constant returns(uint256)
func (_AccountantImplementation *AccountantImplementationCallerSession) CalculateAccountantFee(_amount *big.Int) (*big.Int, error) {
	return _AccountantImplementation.Contract.CalculateAccountantFee(&_AccountantImplementation.CallOpts, _amount)
}

// Channels is a free data retrieval call binding the contract method 0x7a7ebd7b.
//
// Solidity: function channels(bytes32 ) constant returns(address beneficiary, uint256 balance, uint256 settled, uint256 stake, uint256 stakeGoal, uint256 lastUsedNonce, uint256 timelock)
func (_AccountantImplementation *AccountantImplementationCaller) Channels(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Beneficiary   common.Address
	Balance       *big.Int
	Settled       *big.Int
	Stake         *big.Int
	StakeGoal     *big.Int
	LastUsedNonce *big.Int
	Timelock      *big.Int
}, error) {
	ret := new(struct {
		Beneficiary   common.Address
		Balance       *big.Int
		Settled       *big.Int
		Stake         *big.Int
		StakeGoal     *big.Int
		LastUsedNonce *big.Int
		Timelock      *big.Int
	})
	out := ret
	err := _AccountantImplementation.contract.Call(opts, out, "channels", arg0)
	return *ret, err
}

// Channels is a free data retrieval call binding the contract method 0x7a7ebd7b.
//
// Solidity: function channels(bytes32 ) constant returns(address beneficiary, uint256 balance, uint256 settled, uint256 stake, uint256 stakeGoal, uint256 lastUsedNonce, uint256 timelock)
func (_AccountantImplementation *AccountantImplementationSession) Channels(arg0 [32]byte) (struct {
	Beneficiary   common.Address
	Balance       *big.Int
	Settled       *big.Int
	Stake         *big.Int
	StakeGoal     *big.Int
	LastUsedNonce *big.Int
	Timelock      *big.Int
}, error) {
	return _AccountantImplementation.Contract.Channels(&_AccountantImplementation.CallOpts, arg0)
}

// Channels is a free data retrieval call binding the contract method 0x7a7ebd7b.
//
// Solidity: function channels(bytes32 ) constant returns(address beneficiary, uint256 balance, uint256 settled, uint256 stake, uint256 stakeGoal, uint256 lastUsedNonce, uint256 timelock)
func (_AccountantImplementation *AccountantImplementationCallerSession) Channels(arg0 [32]byte) (struct {
	Beneficiary   common.Address
	Balance       *big.Int
	Settled       *big.Int
	Stake         *big.Int
	StakeGoal     *big.Int
	LastUsedNonce *big.Int
	Timelock      *big.Int
}, error) {
	return _AccountantImplementation.Contract.Channels(&_AccountantImplementation.CallOpts, arg0)
}

// GetAccountantStake is a free data retrieval call binding the contract method 0xd1439a4d.
//
// Solidity: function getAccountantStake() constant returns(uint256)
func (_AccountantImplementation *AccountantImplementationCaller) GetAccountantStake(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccountantImplementation.contract.Call(opts, out, "getAccountantStake")
	return *ret0, err
}

// GetAccountantStake is a free data retrieval call binding the contract method 0xd1439a4d.
//
// Solidity: function getAccountantStake() constant returns(uint256)
func (_AccountantImplementation *AccountantImplementationSession) GetAccountantStake() (*big.Int, error) {
	return _AccountantImplementation.Contract.GetAccountantStake(&_AccountantImplementation.CallOpts)
}

// GetAccountantStake is a free data retrieval call binding the contract method 0xd1439a4d.
//
// Solidity: function getAccountantStake() constant returns(uint256)
func (_AccountantImplementation *AccountantImplementationCallerSession) GetAccountantStake() (*big.Int, error) {
	return _AccountantImplementation.Contract.GetAccountantStake(&_AccountantImplementation.CallOpts)
}

// GetChannelId is a free data retrieval call binding the contract method 0xeb295b27.
//
// Solidity: function getChannelId(address _identity) constant returns(bytes32)
func (_AccountantImplementation *AccountantImplementationCaller) GetChannelId(opts *bind.CallOpts, _identity common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _AccountantImplementation.contract.Call(opts, out, "getChannelId", _identity)
	return *ret0, err
}

// GetChannelId is a free data retrieval call binding the contract method 0xeb295b27.
//
// Solidity: function getChannelId(address _identity) constant returns(bytes32)
func (_AccountantImplementation *AccountantImplementationSession) GetChannelId(_identity common.Address) ([32]byte, error) {
	return _AccountantImplementation.Contract.GetChannelId(&_AccountantImplementation.CallOpts, _identity)
}

// GetChannelId is a free data retrieval call binding the contract method 0xeb295b27.
//
// Solidity: function getChannelId(address _identity) constant returns(bytes32)
func (_AccountantImplementation *AccountantImplementationCallerSession) GetChannelId(_identity common.Address) ([32]byte, error) {
	return _AccountantImplementation.Contract.GetChannelId(&_AccountantImplementation.CallOpts, _identity)
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

// GetStakeThresholds is a free data retrieval call binding the contract method 0x9ed9903e.
//
// Solidity: function getStakeThresholds() constant returns(uint256, uint256)
func (_AccountantImplementation *AccountantImplementationCaller) GetStakeThresholds(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _AccountantImplementation.contract.Call(opts, out, "getStakeThresholds")
	return *ret0, *ret1, err
}

// GetStakeThresholds is a free data retrieval call binding the contract method 0x9ed9903e.
//
// Solidity: function getStakeThresholds() constant returns(uint256, uint256)
func (_AccountantImplementation *AccountantImplementationSession) GetStakeThresholds() (*big.Int, *big.Int, error) {
	return _AccountantImplementation.Contract.GetStakeThresholds(&_AccountantImplementation.CallOpts)
}

// GetStakeThresholds is a free data retrieval call binding the contract method 0x9ed9903e.
//
// Solidity: function getStakeThresholds() constant returns(uint256, uint256)
func (_AccountantImplementation *AccountantImplementationCallerSession) GetStakeThresholds() (*big.Int, *big.Int, error) {
	return _AccountantImplementation.Contract.GetStakeThresholds(&_AccountantImplementation.CallOpts)
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

// IsChannelOpened is a free data retrieval call binding the contract method 0x6e9094ea.
//
// Solidity: function isChannelOpened(bytes32 _channelId) constant returns(bool)
func (_AccountantImplementation *AccountantImplementationCaller) IsChannelOpened(opts *bind.CallOpts, _channelId [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AccountantImplementation.contract.Call(opts, out, "isChannelOpened", _channelId)
	return *ret0, err
}

// IsChannelOpened is a free data retrieval call binding the contract method 0x6e9094ea.
//
// Solidity: function isChannelOpened(bytes32 _channelId) constant returns(bool)
func (_AccountantImplementation *AccountantImplementationSession) IsChannelOpened(_channelId [32]byte) (bool, error) {
	return _AccountantImplementation.Contract.IsChannelOpened(&_AccountantImplementation.CallOpts, _channelId)
}

// IsChannelOpened is a free data retrieval call binding the contract method 0x6e9094ea.
//
// Solidity: function isChannelOpened(bytes32 _channelId) constant returns(bool)
func (_AccountantImplementation *AccountantImplementationCallerSession) IsChannelOpened(_channelId [32]byte) (bool, error) {
	return _AccountantImplementation.Contract.IsChannelOpened(&_AccountantImplementation.CallOpts, _channelId)
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

// DecreaseStake is a paid mutator transaction binding the contract method 0xac7d01b6.
//
// Solidity: function decreaseStake(bytes32 _channelId, uint256 _amount, uint256 _transactorFee, uint256 _nonce, bytes _signature) returns()
func (_AccountantImplementation *AccountantImplementationTransactor) DecreaseStake(opts *bind.TransactOpts, _channelId [32]byte, _amount *big.Int, _transactorFee *big.Int, _nonce *big.Int, _signature []byte) (*types.Transaction, error) {
	return _AccountantImplementation.contract.Transact(opts, "decreaseStake", _channelId, _amount, _transactorFee, _nonce, _signature)
}

// DecreaseStake is a paid mutator transaction binding the contract method 0xac7d01b6.
//
// Solidity: function decreaseStake(bytes32 _channelId, uint256 _amount, uint256 _transactorFee, uint256 _nonce, bytes _signature) returns()
func (_AccountantImplementation *AccountantImplementationSession) DecreaseStake(_channelId [32]byte, _amount *big.Int, _transactorFee *big.Int, _nonce *big.Int, _signature []byte) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.DecreaseStake(&_AccountantImplementation.TransactOpts, _channelId, _amount, _transactorFee, _nonce, _signature)
}

// DecreaseStake is a paid mutator transaction binding the contract method 0xac7d01b6.
//
// Solidity: function decreaseStake(bytes32 _channelId, uint256 _amount, uint256 _transactorFee, uint256 _nonce, bytes _signature) returns()
func (_AccountantImplementation *AccountantImplementationTransactorSession) DecreaseStake(_channelId [32]byte, _amount *big.Int, _transactorFee *big.Int, _nonce *big.Int, _signature []byte) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.DecreaseStake(&_AccountantImplementation.TransactOpts, _channelId, _amount, _transactorFee, _nonce, _signature)
}

// GetStakeBack is a paid mutator transaction binding the contract method 0x6138dda7.
//
// Solidity: function getStakeBack(address _beneficiary) returns()
func (_AccountantImplementation *AccountantImplementationTransactor) GetStakeBack(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _AccountantImplementation.contract.Transact(opts, "getStakeBack", _beneficiary)
}

// GetStakeBack is a paid mutator transaction binding the contract method 0x6138dda7.
//
// Solidity: function getStakeBack(address _beneficiary) returns()
func (_AccountantImplementation *AccountantImplementationSession) GetStakeBack(_beneficiary common.Address) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.GetStakeBack(&_AccountantImplementation.TransactOpts, _beneficiary)
}

// GetStakeBack is a paid mutator transaction binding the contract method 0x6138dda7.
//
// Solidity: function getStakeBack(address _beneficiary) returns()
func (_AccountantImplementation *AccountantImplementationTransactorSession) GetStakeBack(_beneficiary common.Address) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.GetStakeBack(&_AccountantImplementation.TransactOpts, _beneficiary)
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

// IncreaseStake is a paid mutator transaction binding the contract method 0x39f97626.
//
// Solidity: function increaseStake(bytes32 _channelId, uint256 _amount) returns()
func (_AccountantImplementation *AccountantImplementationTransactor) IncreaseStake(opts *bind.TransactOpts, _channelId [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _AccountantImplementation.contract.Transact(opts, "increaseStake", _channelId, _amount)
}

// IncreaseStake is a paid mutator transaction binding the contract method 0x39f97626.
//
// Solidity: function increaseStake(bytes32 _channelId, uint256 _amount) returns()
func (_AccountantImplementation *AccountantImplementationSession) IncreaseStake(_channelId [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.IncreaseStake(&_AccountantImplementation.TransactOpts, _channelId, _amount)
}

// IncreaseStake is a paid mutator transaction binding the contract method 0x39f97626.
//
// Solidity: function increaseStake(bytes32 _channelId, uint256 _amount) returns()
func (_AccountantImplementation *AccountantImplementationTransactorSession) IncreaseStake(_channelId [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.IncreaseStake(&_AccountantImplementation.TransactOpts, _channelId, _amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xfec8157d.
//
// Solidity: function initialize(address _token, address _operator, uint16 _fee, uint256 _maxStake) returns()
func (_AccountantImplementation *AccountantImplementationTransactor) Initialize(opts *bind.TransactOpts, _token common.Address, _operator common.Address, _fee uint16, _maxStake *big.Int) (*types.Transaction, error) {
	return _AccountantImplementation.contract.Transact(opts, "initialize", _token, _operator, _fee, _maxStake)
}

// Initialize is a paid mutator transaction binding the contract method 0xfec8157d.
//
// Solidity: function initialize(address _token, address _operator, uint16 _fee, uint256 _maxStake) returns()
func (_AccountantImplementation *AccountantImplementationSession) Initialize(_token common.Address, _operator common.Address, _fee uint16, _maxStake *big.Int) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.Initialize(&_AccountantImplementation.TransactOpts, _token, _operator, _fee, _maxStake)
}

// Initialize is a paid mutator transaction binding the contract method 0xfec8157d.
//
// Solidity: function initialize(address _token, address _operator, uint16 _fee, uint256 _maxStake) returns()
func (_AccountantImplementation *AccountantImplementationTransactorSession) Initialize(_token common.Address, _operator common.Address, _fee uint16, _maxStake *big.Int) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.Initialize(&_AccountantImplementation.TransactOpts, _token, _operator, _fee, _maxStake)
}

// OpenChannel is a paid mutator transaction binding the contract method 0x0a798f24.
//
// Solidity: function openChannel(address _identity, address _beneficiary, uint256 _amountToLend) returns()
func (_AccountantImplementation *AccountantImplementationTransactor) OpenChannel(opts *bind.TransactOpts, _identity common.Address, _beneficiary common.Address, _amountToLend *big.Int) (*types.Transaction, error) {
	return _AccountantImplementation.contract.Transact(opts, "openChannel", _identity, _beneficiary, _amountToLend)
}

// OpenChannel is a paid mutator transaction binding the contract method 0x0a798f24.
//
// Solidity: function openChannel(address _identity, address _beneficiary, uint256 _amountToLend) returns()
func (_AccountantImplementation *AccountantImplementationSession) OpenChannel(_identity common.Address, _beneficiary common.Address, _amountToLend *big.Int) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.OpenChannel(&_AccountantImplementation.TransactOpts, _identity, _beneficiary, _amountToLend)
}

// OpenChannel is a paid mutator transaction binding the contract method 0x0a798f24.
//
// Solidity: function openChannel(address _identity, address _beneficiary, uint256 _amountToLend) returns()
func (_AccountantImplementation *AccountantImplementationTransactorSession) OpenChannel(_identity common.Address, _beneficiary common.Address, _amountToLend *big.Int) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.OpenChannel(&_AccountantImplementation.TransactOpts, _identity, _beneficiary, _amountToLend)
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

// SetBeneficiary is a paid mutator transaction binding the contract method 0x5f59def0.
//
// Solidity: function setBeneficiary(bytes32 _channelId, address _newBeneficiary, uint256 _nonce, bytes _signature) returns()
func (_AccountantImplementation *AccountantImplementationTransactor) SetBeneficiary(opts *bind.TransactOpts, _channelId [32]byte, _newBeneficiary common.Address, _nonce *big.Int, _signature []byte) (*types.Transaction, error) {
	return _AccountantImplementation.contract.Transact(opts, "setBeneficiary", _channelId, _newBeneficiary, _nonce, _signature)
}

// SetBeneficiary is a paid mutator transaction binding the contract method 0x5f59def0.
//
// Solidity: function setBeneficiary(bytes32 _channelId, address _newBeneficiary, uint256 _nonce, bytes _signature) returns()
func (_AccountantImplementation *AccountantImplementationSession) SetBeneficiary(_channelId [32]byte, _newBeneficiary common.Address, _nonce *big.Int, _signature []byte) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.SetBeneficiary(&_AccountantImplementation.TransactOpts, _channelId, _newBeneficiary, _nonce, _signature)
}

// SetBeneficiary is a paid mutator transaction binding the contract method 0x5f59def0.
//
// Solidity: function setBeneficiary(bytes32 _channelId, address _newBeneficiary, uint256 _nonce, bytes _signature) returns()
func (_AccountantImplementation *AccountantImplementationTransactorSession) SetBeneficiary(_channelId [32]byte, _newBeneficiary common.Address, _nonce *big.Int, _signature []byte) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.SetBeneficiary(&_AccountantImplementation.TransactOpts, _channelId, _newBeneficiary, _nonce, _signature)
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

// SetMaxStake is a paid mutator transaction binding the contract method 0x6fc14837.
//
// Solidity: function setMaxStake(uint256 _newMaxStake) returns()
func (_AccountantImplementation *AccountantImplementationTransactor) SetMaxStake(opts *bind.TransactOpts, _newMaxStake *big.Int) (*types.Transaction, error) {
	return _AccountantImplementation.contract.Transact(opts, "setMaxStake", _newMaxStake)
}

// SetMaxStake is a paid mutator transaction binding the contract method 0x6fc14837.
//
// Solidity: function setMaxStake(uint256 _newMaxStake) returns()
func (_AccountantImplementation *AccountantImplementationSession) SetMaxStake(_newMaxStake *big.Int) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.SetMaxStake(&_AccountantImplementation.TransactOpts, _newMaxStake)
}

// SetMaxStake is a paid mutator transaction binding the contract method 0x6fc14837.
//
// Solidity: function setMaxStake(uint256 _newMaxStake) returns()
func (_AccountantImplementation *AccountantImplementationTransactorSession) SetMaxStake(_newMaxStake *big.Int) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.SetMaxStake(&_AccountantImplementation.TransactOpts, _newMaxStake)
}

// SetMinStake is a paid mutator transaction binding the contract method 0x8c80fd90.
//
// Solidity: function setMinStake(uint256 _newMinStake) returns()
func (_AccountantImplementation *AccountantImplementationTransactor) SetMinStake(opts *bind.TransactOpts, _newMinStake *big.Int) (*types.Transaction, error) {
	return _AccountantImplementation.contract.Transact(opts, "setMinStake", _newMinStake)
}

// SetMinStake is a paid mutator transaction binding the contract method 0x8c80fd90.
//
// Solidity: function setMinStake(uint256 _newMinStake) returns()
func (_AccountantImplementation *AccountantImplementationSession) SetMinStake(_newMinStake *big.Int) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.SetMinStake(&_AccountantImplementation.TransactOpts, _newMinStake)
}

// SetMinStake is a paid mutator transaction binding the contract method 0x8c80fd90.
//
// Solidity: function setMinStake(uint256 _newMinStake) returns()
func (_AccountantImplementation *AccountantImplementationTransactorSession) SetMinStake(_newMinStake *big.Int) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.SetMinStake(&_AccountantImplementation.TransactOpts, _newMinStake)
}

// SettleAndRebalance is a paid mutator transaction binding the contract method 0xb07d9483.
//
// Solidity: function settleAndRebalance(address _identity, uint256 _amount, uint256 _transactorFee, bytes32 _lock, bytes _signature) returns()
func (_AccountantImplementation *AccountantImplementationTransactor) SettleAndRebalance(opts *bind.TransactOpts, _identity common.Address, _amount *big.Int, _transactorFee *big.Int, _lock [32]byte, _signature []byte) (*types.Transaction, error) {
	return _AccountantImplementation.contract.Transact(opts, "settleAndRebalance", _identity, _amount, _transactorFee, _lock, _signature)
}

// SettleAndRebalance is a paid mutator transaction binding the contract method 0xb07d9483.
//
// Solidity: function settleAndRebalance(address _identity, uint256 _amount, uint256 _transactorFee, bytes32 _lock, bytes _signature) returns()
func (_AccountantImplementation *AccountantImplementationSession) SettleAndRebalance(_identity common.Address, _amount *big.Int, _transactorFee *big.Int, _lock [32]byte, _signature []byte) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.SettleAndRebalance(&_AccountantImplementation.TransactOpts, _identity, _amount, _transactorFee, _lock, _signature)
}

// SettleAndRebalance is a paid mutator transaction binding the contract method 0xb07d9483.
//
// Solidity: function settleAndRebalance(address _identity, uint256 _amount, uint256 _transactorFee, bytes32 _lock, bytes _signature) returns()
func (_AccountantImplementation *AccountantImplementationTransactorSession) SettleAndRebalance(_identity common.Address, _amount *big.Int, _transactorFee *big.Int, _lock [32]byte, _signature []byte) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.SettleAndRebalance(&_AccountantImplementation.TransactOpts, _identity, _amount, _transactorFee, _lock, _signature)
}

// SettleIntoStake is a paid mutator transaction binding the contract method 0xc062fe06.
//
// Solidity: function settleIntoStake(bytes32 _channelId, uint256 _amount, uint256 _transactorFee, bytes32 _lock, bytes _signature) returns()
func (_AccountantImplementation *AccountantImplementationTransactor) SettleIntoStake(opts *bind.TransactOpts, _channelId [32]byte, _amount *big.Int, _transactorFee *big.Int, _lock [32]byte, _signature []byte) (*types.Transaction, error) {
	return _AccountantImplementation.contract.Transact(opts, "settleIntoStake", _channelId, _amount, _transactorFee, _lock, _signature)
}

// SettleIntoStake is a paid mutator transaction binding the contract method 0xc062fe06.
//
// Solidity: function settleIntoStake(bytes32 _channelId, uint256 _amount, uint256 _transactorFee, bytes32 _lock, bytes _signature) returns()
func (_AccountantImplementation *AccountantImplementationSession) SettleIntoStake(_channelId [32]byte, _amount *big.Int, _transactorFee *big.Int, _lock [32]byte, _signature []byte) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.SettleIntoStake(&_AccountantImplementation.TransactOpts, _channelId, _amount, _transactorFee, _lock, _signature)
}

// SettleIntoStake is a paid mutator transaction binding the contract method 0xc062fe06.
//
// Solidity: function settleIntoStake(bytes32 _channelId, uint256 _amount, uint256 _transactorFee, bytes32 _lock, bytes _signature) returns()
func (_AccountantImplementation *AccountantImplementationTransactorSession) SettleIntoStake(_channelId [32]byte, _amount *big.Int, _transactorFee *big.Int, _lock [32]byte, _signature []byte) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.SettleIntoStake(&_AccountantImplementation.TransactOpts, _channelId, _amount, _transactorFee, _lock, _signature)
}

// SettlePromise is a paid mutator transaction binding the contract method 0x48d9f01e.
//
// Solidity: function settlePromise(address _identity, uint256 _amount, uint256 _transactorFee, bytes32 _lock, bytes _signature) returns()
func (_AccountantImplementation *AccountantImplementationTransactor) SettlePromise(opts *bind.TransactOpts, _identity common.Address, _amount *big.Int, _transactorFee *big.Int, _lock [32]byte, _signature []byte) (*types.Transaction, error) {
	return _AccountantImplementation.contract.Transact(opts, "settlePromise", _identity, _amount, _transactorFee, _lock, _signature)
}

// SettlePromise is a paid mutator transaction binding the contract method 0x48d9f01e.
//
// Solidity: function settlePromise(address _identity, uint256 _amount, uint256 _transactorFee, bytes32 _lock, bytes _signature) returns()
func (_AccountantImplementation *AccountantImplementationSession) SettlePromise(_identity common.Address, _amount *big.Int, _transactorFee *big.Int, _lock [32]byte, _signature []byte) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.SettlePromise(&_AccountantImplementation.TransactOpts, _identity, _amount, _transactorFee, _lock, _signature)
}

// SettlePromise is a paid mutator transaction binding the contract method 0x48d9f01e.
//
// Solidity: function settlePromise(address _identity, uint256 _amount, uint256 _transactorFee, bytes32 _lock, bytes _signature) returns()
func (_AccountantImplementation *AccountantImplementationTransactorSession) SettlePromise(_identity common.Address, _amount *big.Int, _transactorFee *big.Int, _lock [32]byte, _signature []byte) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.SettlePromise(&_AccountantImplementation.TransactOpts, _identity, _amount, _transactorFee, _lock, _signature)
}

// SettleWithBeneficiary is a paid mutator transaction binding the contract method 0xcc449054.
//
// Solidity: function settleWithBeneficiary(address _identity, uint256 _amount, uint256 _transactorFee, bytes32 _lock, bytes _promiseSignature, address _newBeneficiary, uint256 _nonce, bytes _signature) returns()
func (_AccountantImplementation *AccountantImplementationTransactor) SettleWithBeneficiary(opts *bind.TransactOpts, _identity common.Address, _amount *big.Int, _transactorFee *big.Int, _lock [32]byte, _promiseSignature []byte, _newBeneficiary common.Address, _nonce *big.Int, _signature []byte) (*types.Transaction, error) {
	return _AccountantImplementation.contract.Transact(opts, "settleWithBeneficiary", _identity, _amount, _transactorFee, _lock, _promiseSignature, _newBeneficiary, _nonce, _signature)
}

// SettleWithBeneficiary is a paid mutator transaction binding the contract method 0xcc449054.
//
// Solidity: function settleWithBeneficiary(address _identity, uint256 _amount, uint256 _transactorFee, bytes32 _lock, bytes _promiseSignature, address _newBeneficiary, uint256 _nonce, bytes _signature) returns()
func (_AccountantImplementation *AccountantImplementationSession) SettleWithBeneficiary(_identity common.Address, _amount *big.Int, _transactorFee *big.Int, _lock [32]byte, _promiseSignature []byte, _newBeneficiary common.Address, _nonce *big.Int, _signature []byte) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.SettleWithBeneficiary(&_AccountantImplementation.TransactOpts, _identity, _amount, _transactorFee, _lock, _promiseSignature, _newBeneficiary, _nonce, _signature)
}

// SettleWithBeneficiary is a paid mutator transaction binding the contract method 0xcc449054.
//
// Solidity: function settleWithBeneficiary(address _identity, uint256 _amount, uint256 _transactorFee, bytes32 _lock, bytes _promiseSignature, address _newBeneficiary, uint256 _nonce, bytes _signature) returns()
func (_AccountantImplementation *AccountantImplementationTransactorSession) SettleWithBeneficiary(_identity common.Address, _amount *big.Int, _transactorFee *big.Int, _lock [32]byte, _promiseSignature []byte, _newBeneficiary common.Address, _nonce *big.Int, _signature []byte) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.SettleWithBeneficiary(&_AccountantImplementation.TransactOpts, _identity, _amount, _transactorFee, _lock, _promiseSignature, _newBeneficiary, _nonce, _signature)
}

// SettleWithGoalIncrease is a paid mutator transaction binding the contract method 0x0b9a91c1.
//
// Solidity: function settleWithGoalIncrease(bytes32 _channelId, uint256 _amount, uint256 _transactorFee, bytes32 _lock, bytes _promiseSignature, uint256 _newStakeGoal, uint256 _nonce, bytes _goalChangeSignature) returns()
func (_AccountantImplementation *AccountantImplementationTransactor) SettleWithGoalIncrease(opts *bind.TransactOpts, _channelId [32]byte, _amount *big.Int, _transactorFee *big.Int, _lock [32]byte, _promiseSignature []byte, _newStakeGoal *big.Int, _nonce *big.Int, _goalChangeSignature []byte) (*types.Transaction, error) {
	return _AccountantImplementation.contract.Transact(opts, "settleWithGoalIncrease", _channelId, _amount, _transactorFee, _lock, _promiseSignature, _newStakeGoal, _nonce, _goalChangeSignature)
}

// SettleWithGoalIncrease is a paid mutator transaction binding the contract method 0x0b9a91c1.
//
// Solidity: function settleWithGoalIncrease(bytes32 _channelId, uint256 _amount, uint256 _transactorFee, bytes32 _lock, bytes _promiseSignature, uint256 _newStakeGoal, uint256 _nonce, bytes _goalChangeSignature) returns()
func (_AccountantImplementation *AccountantImplementationSession) SettleWithGoalIncrease(_channelId [32]byte, _amount *big.Int, _transactorFee *big.Int, _lock [32]byte, _promiseSignature []byte, _newStakeGoal *big.Int, _nonce *big.Int, _goalChangeSignature []byte) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.SettleWithGoalIncrease(&_AccountantImplementation.TransactOpts, _channelId, _amount, _transactorFee, _lock, _promiseSignature, _newStakeGoal, _nonce, _goalChangeSignature)
}

// SettleWithGoalIncrease is a paid mutator transaction binding the contract method 0x0b9a91c1.
//
// Solidity: function settleWithGoalIncrease(bytes32 _channelId, uint256 _amount, uint256 _transactorFee, bytes32 _lock, bytes _promiseSignature, uint256 _newStakeGoal, uint256 _nonce, bytes _goalChangeSignature) returns()
func (_AccountantImplementation *AccountantImplementationTransactorSession) SettleWithGoalIncrease(_channelId [32]byte, _amount *big.Int, _transactorFee *big.Int, _lock [32]byte, _promiseSignature []byte, _newStakeGoal *big.Int, _nonce *big.Int, _goalChangeSignature []byte) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.SettleWithGoalIncrease(&_AccountantImplementation.TransactOpts, _channelId, _amount, _transactorFee, _lock, _promiseSignature, _newStakeGoal, _nonce, _goalChangeSignature)
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

// UpdateStakeGoal is a paid mutator transaction binding the contract method 0xea48d8bf.
//
// Solidity: function updateStakeGoal(bytes32 _channelId, uint256 _newStakeGoal, uint256 _nonce, bytes _signature) returns()
func (_AccountantImplementation *AccountantImplementationTransactor) UpdateStakeGoal(opts *bind.TransactOpts, _channelId [32]byte, _newStakeGoal *big.Int, _nonce *big.Int, _signature []byte) (*types.Transaction, error) {
	return _AccountantImplementation.contract.Transact(opts, "updateStakeGoal", _channelId, _newStakeGoal, _nonce, _signature)
}

// UpdateStakeGoal is a paid mutator transaction binding the contract method 0xea48d8bf.
//
// Solidity: function updateStakeGoal(bytes32 _channelId, uint256 _newStakeGoal, uint256 _nonce, bytes _signature) returns()
func (_AccountantImplementation *AccountantImplementationSession) UpdateStakeGoal(_channelId [32]byte, _newStakeGoal *big.Int, _nonce *big.Int, _signature []byte) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.UpdateStakeGoal(&_AccountantImplementation.TransactOpts, _channelId, _newStakeGoal, _nonce, _signature)
}

// UpdateStakeGoal is a paid mutator transaction binding the contract method 0xea48d8bf.
//
// Solidity: function updateStakeGoal(bytes32 _channelId, uint256 _newStakeGoal, uint256 _nonce, bytes _signature) returns()
func (_AccountantImplementation *AccountantImplementationTransactorSession) UpdateStakeGoal(_channelId [32]byte, _newStakeGoal *big.Int, _nonce *big.Int, _signature []byte) (*types.Transaction, error) {
	return _AccountantImplementation.Contract.UpdateStakeGoal(&_AccountantImplementation.TransactOpts, _channelId, _newStakeGoal, _nonce, _signature)
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

// AccountantImplementationMaxStakeValueUpdatedIterator is returned from FilterMaxStakeValueUpdated and is used to iterate over the raw logs and unpacked data for MaxStakeValueUpdated events raised by the AccountantImplementation contract.
type AccountantImplementationMaxStakeValueUpdatedIterator struct {
	Event *AccountantImplementationMaxStakeValueUpdated // Event containing the contract specifics and raw log

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
func (it *AccountantImplementationMaxStakeValueUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountantImplementationMaxStakeValueUpdated)
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
		it.Event = new(AccountantImplementationMaxStakeValueUpdated)
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
func (it *AccountantImplementationMaxStakeValueUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountantImplementationMaxStakeValueUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountantImplementationMaxStakeValueUpdated represents a MaxStakeValueUpdated event raised by the AccountantImplementation contract.
type AccountantImplementationMaxStakeValueUpdated struct {
	NewMaxStake *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMaxStakeValueUpdated is a free log retrieval operation binding the contract event 0x53f4fb18cb329155d5af04681c1d0846d0484d7de33791619c6988ca61910e3d.
//
// Solidity: event MaxStakeValueUpdated(uint256 newMaxStake)
func (_AccountantImplementation *AccountantImplementationFilterer) FilterMaxStakeValueUpdated(opts *bind.FilterOpts) (*AccountantImplementationMaxStakeValueUpdatedIterator, error) {

	logs, sub, err := _AccountantImplementation.contract.FilterLogs(opts, "MaxStakeValueUpdated")
	if err != nil {
		return nil, err
	}
	return &AccountantImplementationMaxStakeValueUpdatedIterator{contract: _AccountantImplementation.contract, event: "MaxStakeValueUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxStakeValueUpdated is a free log subscription operation binding the contract event 0x53f4fb18cb329155d5af04681c1d0846d0484d7de33791619c6988ca61910e3d.
//
// Solidity: event MaxStakeValueUpdated(uint256 newMaxStake)
func (_AccountantImplementation *AccountantImplementationFilterer) WatchMaxStakeValueUpdated(opts *bind.WatchOpts, sink chan<- *AccountantImplementationMaxStakeValueUpdated) (event.Subscription, error) {

	logs, sub, err := _AccountantImplementation.contract.WatchLogs(opts, "MaxStakeValueUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountantImplementationMaxStakeValueUpdated)
				if err := _AccountantImplementation.contract.UnpackLog(event, "MaxStakeValueUpdated", log); err != nil {
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

// ParseMaxStakeValueUpdated is a log parse operation binding the contract event 0x53f4fb18cb329155d5af04681c1d0846d0484d7de33791619c6988ca61910e3d.
//
// Solidity: event MaxStakeValueUpdated(uint256 newMaxStake)
func (_AccountantImplementation *AccountantImplementationFilterer) ParseMaxStakeValueUpdated(log types.Log) (*AccountantImplementationMaxStakeValueUpdated, error) {
	event := new(AccountantImplementationMaxStakeValueUpdated)
	if err := _AccountantImplementation.contract.UnpackLog(event, "MaxStakeValueUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountantImplementationMinStakeValueUpdatedIterator is returned from FilterMinStakeValueUpdated and is used to iterate over the raw logs and unpacked data for MinStakeValueUpdated events raised by the AccountantImplementation contract.
type AccountantImplementationMinStakeValueUpdatedIterator struct {
	Event *AccountantImplementationMinStakeValueUpdated // Event containing the contract specifics and raw log

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
func (it *AccountantImplementationMinStakeValueUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountantImplementationMinStakeValueUpdated)
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
		it.Event = new(AccountantImplementationMinStakeValueUpdated)
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
func (it *AccountantImplementationMinStakeValueUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountantImplementationMinStakeValueUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountantImplementationMinStakeValueUpdated represents a MinStakeValueUpdated event raised by the AccountantImplementation contract.
type AccountantImplementationMinStakeValueUpdated struct {
	NewMinStake *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMinStakeValueUpdated is a free log retrieval operation binding the contract event 0xb9e5e6e8db1283ee860f3856d8383e40665c58a5264ede5e6ed8ec1afb031251.
//
// Solidity: event MinStakeValueUpdated(uint256 newMinStake)
func (_AccountantImplementation *AccountantImplementationFilterer) FilterMinStakeValueUpdated(opts *bind.FilterOpts) (*AccountantImplementationMinStakeValueUpdatedIterator, error) {

	logs, sub, err := _AccountantImplementation.contract.FilterLogs(opts, "MinStakeValueUpdated")
	if err != nil {
		return nil, err
	}
	return &AccountantImplementationMinStakeValueUpdatedIterator{contract: _AccountantImplementation.contract, event: "MinStakeValueUpdated", logs: logs, sub: sub}, nil
}

// WatchMinStakeValueUpdated is a free log subscription operation binding the contract event 0xb9e5e6e8db1283ee860f3856d8383e40665c58a5264ede5e6ed8ec1afb031251.
//
// Solidity: event MinStakeValueUpdated(uint256 newMinStake)
func (_AccountantImplementation *AccountantImplementationFilterer) WatchMinStakeValueUpdated(opts *bind.WatchOpts, sink chan<- *AccountantImplementationMinStakeValueUpdated) (event.Subscription, error) {

	logs, sub, err := _AccountantImplementation.contract.WatchLogs(opts, "MinStakeValueUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountantImplementationMinStakeValueUpdated)
				if err := _AccountantImplementation.contract.UnpackLog(event, "MinStakeValueUpdated", log); err != nil {
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

// ParseMinStakeValueUpdated is a log parse operation binding the contract event 0xb9e5e6e8db1283ee860f3856d8383e40665c58a5264ede5e6ed8ec1afb031251.
//
// Solidity: event MinStakeValueUpdated(uint256 newMinStake)
func (_AccountantImplementation *AccountantImplementationFilterer) ParseMinStakeValueUpdated(log types.Log) (*AccountantImplementationMinStakeValueUpdated, error) {
	event := new(AccountantImplementationMinStakeValueUpdated)
	if err := _AccountantImplementation.contract.UnpackLog(event, "MinStakeValueUpdated", log); err != nil {
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

// AccountantImplementationNewStakeIterator is returned from FilterNewStake and is used to iterate over the raw logs and unpacked data for NewStake events raised by the AccountantImplementation contract.
type AccountantImplementationNewStakeIterator struct {
	Event *AccountantImplementationNewStake // Event containing the contract specifics and raw log

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
func (it *AccountantImplementationNewStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountantImplementationNewStake)
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
		it.Event = new(AccountantImplementationNewStake)
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
func (it *AccountantImplementationNewStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountantImplementationNewStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountantImplementationNewStake represents a NewStake event raised by the AccountantImplementation contract.
type AccountantImplementationNewStake struct {
	ChannelId   [32]byte
	StakeAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewStake is a free log retrieval operation binding the contract event 0xc5f0715c45dab2e8f14871936119e3c64fd5841d397130c2d1db743d142522cb.
//
// Solidity: event NewStake(bytes32 indexed channelId, uint256 stakeAmount)
func (_AccountantImplementation *AccountantImplementationFilterer) FilterNewStake(opts *bind.FilterOpts, channelId [][32]byte) (*AccountantImplementationNewStakeIterator, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _AccountantImplementation.contract.FilterLogs(opts, "NewStake", channelIdRule)
	if err != nil {
		return nil, err
	}
	return &AccountantImplementationNewStakeIterator{contract: _AccountantImplementation.contract, event: "NewStake", logs: logs, sub: sub}, nil
}

// WatchNewStake is a free log subscription operation binding the contract event 0xc5f0715c45dab2e8f14871936119e3c64fd5841d397130c2d1db743d142522cb.
//
// Solidity: event NewStake(bytes32 indexed channelId, uint256 stakeAmount)
func (_AccountantImplementation *AccountantImplementationFilterer) WatchNewStake(opts *bind.WatchOpts, sink chan<- *AccountantImplementationNewStake, channelId [][32]byte) (event.Subscription, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _AccountantImplementation.contract.WatchLogs(opts, "NewStake", channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountantImplementationNewStake)
				if err := _AccountantImplementation.contract.UnpackLog(event, "NewStake", log); err != nil {
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

// ParseNewStake is a log parse operation binding the contract event 0xc5f0715c45dab2e8f14871936119e3c64fd5841d397130c2d1db743d142522cb.
//
// Solidity: event NewStake(bytes32 indexed channelId, uint256 stakeAmount)
func (_AccountantImplementation *AccountantImplementationFilterer) ParseNewStake(log types.Log) (*AccountantImplementationNewStake, error) {
	event := new(AccountantImplementationNewStake)
	if err := _AccountantImplementation.contract.UnpackLog(event, "NewStake", log); err != nil {
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

// AccountantImplementationStakeGoalUpdatedIterator is returned from FilterStakeGoalUpdated and is used to iterate over the raw logs and unpacked data for StakeGoalUpdated events raised by the AccountantImplementation contract.
type AccountantImplementationStakeGoalUpdatedIterator struct {
	Event *AccountantImplementationStakeGoalUpdated // Event containing the contract specifics and raw log

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
func (it *AccountantImplementationStakeGoalUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountantImplementationStakeGoalUpdated)
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
		it.Event = new(AccountantImplementationStakeGoalUpdated)
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
func (it *AccountantImplementationStakeGoalUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountantImplementationStakeGoalUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountantImplementationStakeGoalUpdated represents a StakeGoalUpdated event raised by the AccountantImplementation contract.
type AccountantImplementationStakeGoalUpdated struct {
	ChannelId    [32]byte
	NewStakeGoal *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStakeGoalUpdated is a free log retrieval operation binding the contract event 0x5ec1b60bc705196753b2e21e144bbb185e6ef7264f9d79b99f04f1bc22198733.
//
// Solidity: event StakeGoalUpdated(bytes32 indexed channelId, uint256 newStakeGoal)
func (_AccountantImplementation *AccountantImplementationFilterer) FilterStakeGoalUpdated(opts *bind.FilterOpts, channelId [][32]byte) (*AccountantImplementationStakeGoalUpdatedIterator, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _AccountantImplementation.contract.FilterLogs(opts, "StakeGoalUpdated", channelIdRule)
	if err != nil {
		return nil, err
	}
	return &AccountantImplementationStakeGoalUpdatedIterator{contract: _AccountantImplementation.contract, event: "StakeGoalUpdated", logs: logs, sub: sub}, nil
}

// WatchStakeGoalUpdated is a free log subscription operation binding the contract event 0x5ec1b60bc705196753b2e21e144bbb185e6ef7264f9d79b99f04f1bc22198733.
//
// Solidity: event StakeGoalUpdated(bytes32 indexed channelId, uint256 newStakeGoal)
func (_AccountantImplementation *AccountantImplementationFilterer) WatchStakeGoalUpdated(opts *bind.WatchOpts, sink chan<- *AccountantImplementationStakeGoalUpdated, channelId [][32]byte) (event.Subscription, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _AccountantImplementation.contract.WatchLogs(opts, "StakeGoalUpdated", channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountantImplementationStakeGoalUpdated)
				if err := _AccountantImplementation.contract.UnpackLog(event, "StakeGoalUpdated", log); err != nil {
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

// ParseStakeGoalUpdated is a log parse operation binding the contract event 0x5ec1b60bc705196753b2e21e144bbb185e6ef7264f9d79b99f04f1bc22198733.
//
// Solidity: event StakeGoalUpdated(bytes32 indexed channelId, uint256 newStakeGoal)
func (_AccountantImplementation *AccountantImplementationFilterer) ParseStakeGoalUpdated(log types.Log) (*AccountantImplementationStakeGoalUpdated, error) {
	event := new(AccountantImplementationStakeGoalUpdated)
	if err := _AccountantImplementation.contract.UnpackLog(event, "StakeGoalUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}
