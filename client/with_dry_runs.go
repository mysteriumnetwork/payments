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

package client

import (
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/mysteriumnetwork/payments/bindings"
	"github.com/pkg/errors"
)

// WithDryRuns forces a dry run before running a write transaction on blockchain.
// Ethereum client will perform a dry run on a transaction with no gas limit set.
// This component will perform a dry run if and only if the gas limit is set to a non zero value.
// In this way, the dry run is always performed before sending the transaction to the network.
// For convenience, this component proxies read only calls to the underlying blockchain.
type WithDryRuns struct {
	bc blockchain
}

// NewWithDryRuns creates a new instance of client with dry runs.
func NewWithDryRuns(bc blockchain) *WithDryRuns {
	return &WithDryRuns{
		bc: bc,
	}
}

type gasLimitProvider interface {
	GetGasLimit() uint64
}

// The params for the dry run should be in the strict order as required by the smart contract.
func (cwdr *WithDryRuns) dryRun(glp gasLimitProvider, binding string, from, to common.Address, method string, params ...interface{}) (uint64, error) {
	// If the gas limit is set to 0, ethereum client will do the estimation for us.
	// We only force the estimation if the gas limit is set to a non zero value.
	if glp.GetGasLimit() == 0 {
		return 0, nil
	}
	parsed, err := abi.JSON(strings.NewReader(binding))
	if err != nil {
		return 0, errors.Wrap(err, "could not deserialize ABI")
	}
	input, err := parsed.Pack(method, params...)
	if err != nil {
		return 0, errors.Wrap(err, "could not pack input")
	}
	msg := ethereum.CallMsg{From: from, To: &to, Data: input}
	gas, err := cwdr.bc.EstimateGas(msg)
	return gas, errors.Wrap(err, "could not estimate gas")
}

// TransferMyst transfers myst
func (cwdr *WithDryRuns) TransferMyst(req TransferRequest) (tx *types.Transaction, err error) {
	_, err = cwdr.dryRun(
		req,
		bindings.MystTokenABI,
		req.Identity,
		req.MystAddress,
		"transfer",
		req.Recipient,
		req.Amount,
	)
	if err != nil {
		return nil, errors.Wrap(err, "transaction dry run failed")
	}

	return cwdr.bc.TransferMyst(req)
}

// SettlePromise is settling the given consumer issued promise
func (cwdr *WithDryRuns) SettlePromise(req SettleRequest) (*types.Transaction, error) {
	lock := [32]byte{}
	copy(lock[:], req.Promise.R)

	_, err := cwdr.dryRun(
		req,
		bindings.ChannelImplementationABI,
		req.Identity,
		req.ChannelID,
		"settlePromise",
		big.NewInt(0).SetUint64(req.Promise.Amount),
		big.NewInt(0).SetUint64(req.Promise.Fee),
		lock,
		req.Promise.Signature,
	)
	if err != nil {
		return nil, errors.Wrap(err, "transaction dry run failed")
	}

	return cwdr.bc.SettlePromise(req)
}

// RegisterIdentity registers the given identity on blockchain
func (cwdr *WithDryRuns) RegisterIdentity(rr RegistrationRequest) (*types.Transaction, error) {
	_, err := cwdr.dryRun(
		rr,
		bindings.RegistryABI,
		rr.Identity,
		rr.RegistryAddress,
		"registerIdentity",
		rr.HermesID,
		rr.Stake,
		rr.TransactorFee,
		rr.Beneficiary,
		rr.Signature,
	)
	if err != nil {
		return nil, errors.Wrap(err, "transaction dry run failed")
	}

	return cwdr.bc.RegisterIdentity(rr)
}

// SettleAndRebalance is settling given hermes issued promise
func (cwdr *WithDryRuns) SettleAndRebalance(req SettleAndRebalanceRequest) (*types.Transaction, error) {
	_, err := cwdr.dryRun(
		req,
		bindings.HermesImplementationABI,
		req.Identity,
		req.HermesID,
		"settleAndRebalance",
		req.ProviderID,
		big.NewInt(0).SetUint64(req.Promise.Amount),
		big.NewInt(0).SetUint64(req.Promise.Fee),
		toBytes32(req.Promise.R),
		req.Promise.Signature,
	)
	if err != nil {
		return nil, errors.Wrap(err, "transaction dry run failed")
	}
	return cwdr.bc.SettleAndRebalance(req)
}

// GetHermesFee fetches the hermes fee from blockchain
func (cwdr *WithDryRuns) GetHermesFee(hermesAddress common.Address) (uint16, error) {
	return cwdr.bc.GetHermesFee(hermesAddress)
}

// CalculateHermesFee fetches the hermes fee from blockchain
func (cwdr *WithDryRuns) CalculateHermesFee(hermesAddress common.Address, value uint64) (*big.Int, error) {
	return cwdr.bc.CalculateHermesFee(hermesAddress, value)
}

// IsRegisteredAsProvider checks if the provider is registered with the hermes properly
func (cwdr *WithDryRuns) IsRegisteredAsProvider(hermesAddress, registryAddress, addressToCheck common.Address) (bool, error) {
	return cwdr.bc.IsRegisteredAsProvider(hermesAddress, registryAddress, addressToCheck)
}

// GetProviderChannel returns the provider channel
func (cwdr *WithDryRuns) GetProviderChannel(hermesAddress common.Address, addressToCheck common.Address, pending bool) (ProviderChannel, error) {
	return cwdr.bc.GetProviderChannel(hermesAddress, addressToCheck, pending)
}

// IsRegistered checks wether the given identity is registered or not
func (cwdr *WithDryRuns) IsRegistered(registryAddress, addressToCheck common.Address) (bool, error) {
	return cwdr.bc.IsRegistered(registryAddress, addressToCheck)
}

// SubscribeToPromiseSettledEvent subscribes to promise settled events
func (cwdr *WithDryRuns) SubscribeToPromiseSettledEvent(providerID, hermesID common.Address) (sink chan *bindings.HermesImplementationPromiseSettled, cancel func(), err error) {
	return cwdr.bc.SubscribeToPromiseSettledEvent(providerID, hermesID)
}

// GetMystBalance returns the balance in myst
func (cwdr *WithDryRuns) GetMystBalance(mystSCAddress, address common.Address) (*big.Int, error) {
	return cwdr.bc.GetMystBalance(mystSCAddress, address)
}

// SubscribeToConsumerBalanceEvent subscribes to the consumer balance change events
func (cwdr *WithDryRuns) SubscribeToConsumerBalanceEvent(channel, mystSCAddress common.Address, timeout time.Duration) (chan *bindings.MystTokenTransfer, func(), error) {
	return cwdr.bc.SubscribeToConsumerBalanceEvent(channel, mystSCAddress, timeout)
}

// GetRegistrationFee returns the registration fee
func (cwdr *WithDryRuns) GetRegistrationFee(registryAddress common.Address) (*big.Int, error) {
	return cwdr.bc.GetRegistrationFee(registryAddress)
}

// IsHermesRegistered checks if given hermes is registered and returns true or false.
func (cwdr *WithDryRuns) IsHermesRegistered(registryAddress, acccountantID common.Address) (bool, error) {
	return cwdr.bc.IsHermesRegistered(registryAddress, acccountantID)
}

// GetHermesOperator returns operator address of given hermes
func (cwdr *WithDryRuns) GetHermesOperator(hermesID common.Address) (common.Address, error) {
	return cwdr.bc.GetHermesOperator(hermesID)
}

// GetConsumerChannelsHermes returns the consumer channels hermes
func (cwdr *WithDryRuns) GetConsumerChannelsHermes(channelAddress common.Address) (ConsumersHermes, error) {
	return cwdr.bc.GetConsumerChannelsHermes(channelAddress)
}

// SubscribeToMystTokenTransfers subscribes to myst token transfers
func (cwdr *WithDryRuns) SubscribeToMystTokenTransfers(mystSCAddress common.Address) (chan *bindings.MystTokenTransfer, func(), error) {
	return cwdr.bc.SubscribeToMystTokenTransfers(mystSCAddress)
}

// GetConsumerChannelOperator returns the consumer channel operator/identity
func (cwdr *WithDryRuns) GetConsumerChannelOperator(channelAddress common.Address) (common.Address, error) {
	return cwdr.bc.GetConsumerChannelOperator(channelAddress)
}

// GetConsumerChannel returns the consumer channel
func (cwdr *WithDryRuns) GetConsumerChannel(addr common.Address, mystSCAddress common.Address) (ConsumerChannel, error) {
	return cwdr.bc.GetConsumerChannel(addr, mystSCAddress)
}

// GetProviderChannelByID returns the given channel information
func (cwdr *WithDryRuns) GetProviderChannelByID(acc common.Address, chID []byte) (ProviderChannel, error) {
	return cwdr.bc.GetProviderChannelByID(acc, chID)
}

// SubscribeToIdentityRegistrationEvents subscribes to identity registration events
func (cwdr *WithDryRuns) SubscribeToIdentityRegistrationEvents(registryAddress common.Address, hermesIDs []common.Address) (sink chan *bindings.RegistryRegisteredIdentity, cancel func(), err error) {
	return cwdr.bc.SubscribeToIdentityRegistrationEvents(registryAddress, hermesIDs)
}

// SubscribeToConsumerChannelBalanceUpdate subscribes to consumer channel balance update events
func (cwdr *WithDryRuns) SubscribeToConsumerChannelBalanceUpdate(mystSCAddress common.Address, channelAddresses []common.Address) (sink chan *bindings.MystTokenTransfer, cancel func(), err error) {
	return cwdr.bc.SubscribeToConsumerChannelBalanceUpdate(mystSCAddress, channelAddresses)
}

// SubscribeToProviderChannelBalanceUpdate subscribes to provider channel balance update events
func (cwdr *WithDryRuns) SubscribeToProviderChannelBalanceUpdate(hermesAddress common.Address, channelAddresses [][32]byte) (sink chan *bindings.HermesImplementationChannelBalanceUpdated, cancel func(), err error) {
	return cwdr.bc.SubscribeToProviderChannelBalanceUpdate(hermesAddress, channelAddresses)
}

// SubscribeToChannelOpenedEvents subscribes to provider channel opened events
func (cwdr *WithDryRuns) SubscribeToChannelOpenedEvents(hermesAddress common.Address) (sink chan *bindings.HermesImplementationChannelOpened, cancel func(), err error) {
	return cwdr.bc.SubscribeToChannelOpenedEvents(hermesAddress)
}

// SubscribeToPromiseSettledEventByChannelID subscribes to promise settled events
func (cwdr *WithDryRuns) SubscribeToPromiseSettledEventByChannelID(hermesID common.Address, providerAddresses [][32]byte) (sink chan *bindings.HermesImplementationPromiseSettled, cancel func(), err error) {
	return cwdr.bc.SubscribeToPromiseSettledEventByChannelID(hermesID, providerAddresses)
}

// NetworkID returns the network id
func (cwdr *WithDryRuns) NetworkID() (*big.Int, error) {
	return cwdr.bc.NetworkID()
}

// SettleWithBeneficiary sets new beneficiary and settling given hermes issued promise into it.
func (cwdr *WithDryRuns) SettleWithBeneficiary(req SettleWithBeneficiaryRequest) (*types.Transaction, error) {
	_, err := cwdr.dryRun(
		req,
		bindings.HermesImplementationABI,
		req.Identity,
		req.HermesID,
		"settleWithBeneficiary",
		req.ProviderID,
		big.NewInt(0).SetUint64(req.Promise.Amount),
		big.NewInt(0).SetUint64(req.Promise.Fee),
		toBytes32(req.Promise.R),
		req.Promise.Signature,
		req.Beneficiary,
		big.NewInt(0).SetUint64(req.Nonce),
		req.Signature,
	)
	if err != nil {
		return nil, errors.Wrap(err, "transaction dry run failed")
	}
	return cwdr.bc.SettleWithBeneficiary(req)
}
