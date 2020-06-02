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
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/mysteriumnetwork/payments/bindings"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type blockchain interface {
	GetAccountantFee(accountantAddress common.Address) (uint16, error)
	CalculateAccountantFee(accountantAddress common.Address, value uint64) (*big.Int, error)
	IsRegisteredAsProvider(accountantAddress, registryAddress, addressToCheck common.Address) (bool, error)
	GetProviderChannel(accountantAddress common.Address, addressToCheck common.Address, pending bool) (ProviderChannel, error)
	IsRegistered(registryAddress, addressToCheck common.Address) (bool, error)
	SubscribeToPromiseSettledEvent(providerID, accountantID common.Address) (sink chan *bindings.AccountantImplementationPromiseSettled, cancel func(), err error)
	GetMystBalance(mystSCAddress, address common.Address) (*big.Int, error)
	SubscribeToConsumerBalanceEvent(channel, mystSCAddress common.Address, timeout time.Duration) (chan *bindings.MystTokenTransfer, func(), error)
	GetRegistrationFee(registryAddress common.Address) (*big.Int, error)
	RegisterIdentity(rr RegistrationRequest) (*types.Transaction, error)
	TransferMyst(req TransferRequest) (tx *types.Transaction, err error)
	IsAccountantRegistered(registryAddress, acccountantID common.Address) (bool, error)
	GetAccountantOperator(accountantID common.Address) (common.Address, error)
	SettleAndRebalance(req SettleAndRebalanceRequest) (*types.Transaction, error)
	SettleWithBeneficiary(req SettleWithBeneficiaryRequest) (*types.Transaction, error)
	GetConsumerChannelsAccountant(channelAddress common.Address) (ConsumersAccountant, error)
	GetConsumerChannelOperator(channelAddress common.Address) (common.Address, error)
	GetProviderChannelByID(acc common.Address, chID []byte) (ProviderChannel, error)
	SubscribeToIdentityRegistrationEvents(registryAddress common.Address, accountantIDs []common.Address) (sink chan *bindings.RegistryRegisteredIdentity, cancel func(), err error)
	SubscribeToConsumerChannelBalanceUpdate(mystSCAddress common.Address, channelAddresses []common.Address) (sink chan *bindings.MystTokenTransfer, cancel func(), err error)
	SubscribeToProviderChannelBalanceUpdate(accountantAddress common.Address, channelAddresses [][32]byte) (sink chan *bindings.AccountantImplementationChannelBalanceUpdated, cancel func(), err error)
	SettlePromise(req SettleRequest) (*types.Transaction, error)
	SubscribeToChannelOpenedEvents(accountantAddress common.Address) (sink chan *bindings.AccountantImplementationChannelOpened, cancel func(), err error)
	SubscribeToPromiseSettledEventByChannelID(accountantID common.Address, providerAddresses [][32]byte) (sink chan *bindings.AccountantImplementationPromiseSettled, cancel func(), err error)
	SubscribeToMystTokenTransfers(mystSCAddress common.Address) (chan *bindings.MystTokenTransfer, func(), error)
	NetworkID() (*big.Int, error)
	EstimateGas(msg ethereum.CallMsg) (uint64, error)
	GetConsumerChannel(addr common.Address, mystSCAddress common.Address) (ConsumerChannel, error)
	GetEthBalance(address common.Address) (*big.Int, error)
	TransferEth(etr EthTransferRequest) (*types.Transaction, error)
}

// BlockchainWithRetries takes in the plain blockchain implementation and exposes methods that will retry the underlying bc methods before giving up.
// This is required as the ethereum client will occasionally spit a TLS error if running for prolonged periods of time.
type BlockchainWithRetries struct {
	delay      time.Duration
	maxRetries int
	bc         blockchain
	stop       chan struct{}
	once       sync.Once
}

// ErrStopped represents an error when a call is interrupted
var ErrStopped = errors.New("call stopped")

// NewBlockchainWithRetries returns a new instance of blockchain with retries
func NewBlockchainWithRetries(bc blockchain, delay time.Duration, maxRetries int) *BlockchainWithRetries {
	return &BlockchainWithRetries{
		bc:         bc,
		delay:      delay,
		maxRetries: maxRetries,
	}
}

func (bwr *BlockchainWithRetries) callWithRetry(f func() error) error {
	for i := 1; i <= bwr.maxRetries; i++ {
		err := f()
		if err == nil {
			return nil
		}
		if i == bwr.maxRetries {
			return err
		}

		log.Warn().Err(err).Msgf("retry %v of %v", i+1, bwr.maxRetries)
		select {
		case <-bwr.stop:
			return ErrStopped
		case <-time.After(bwr.delay):
		}
	}
	return nil
}

// SubscribeToMystTokenTransfers subscribes to myst token transfer events
func (bwr *BlockchainWithRetries) SubscribeToMystTokenTransfers(mystSCAddress common.Address) (chan *bindings.MystTokenTransfer, func(), error) {
	var sink chan *bindings.MystTokenTransfer
	var cancel func()
	err := bwr.callWithRetry(func() error {
		s, c, err := bwr.bc.SubscribeToMystTokenTransfers(mystSCAddress)
		if err != nil {
			return errors.Wrap(err, "could not subscribe to settlement events")
		}
		sink = s
		cancel = c
		return nil
	})
	return sink, cancel, err
}

// GetAccountantFee fetches the accountant fee from blockchain
func (bwr *BlockchainWithRetries) GetAccountantFee(accountantAddress common.Address) (uint16, error) {
	var res uint16
	err := bwr.callWithRetry(func() error {
		r, err := bwr.bc.GetAccountantFee(accountantAddress)
		if err != nil {
			return errors.Wrap(err, "could not get accountant fee")
		}
		res = r
		return nil
	})
	return res, err
}

// CalculateAccountantFee fetches the accountant fee from blockchain
func (bwr *BlockchainWithRetries) CalculateAccountantFee(accountantAddress common.Address, value uint64) (*big.Int, error) {
	var res *big.Int
	err := bwr.callWithRetry(func() error {
		r, err := bwr.bc.CalculateAccountantFee(accountantAddress, value)
		if err != nil {
			return errors.Wrap(err, "could not calculate accountant fee")
		}
		res = r
		return nil
	})
	return res, err
}

// IsRegisteredAsProvider checks if the provider is registered with the accountant properly
func (bwr *BlockchainWithRetries) IsRegisteredAsProvider(accountantAddress, registryAddress, addressToCheck common.Address) (bool, error) {
	var res bool
	err := bwr.callWithRetry(func() error {
		r, err := bwr.bc.IsRegisteredAsProvider(accountantAddress, registryAddress, addressToCheck)
		if err != nil {
			return errors.Wrap(err, "could not check if registered as provider")
		}
		res = r
		return nil
	})
	return res, err
}

// GetProviderChannel returns the provider channel
func (bwr *BlockchainWithRetries) GetProviderChannel(accountantAddress, addressToCheck common.Address, pending bool) (ProviderChannel, error) {
	var res ProviderChannel
	err := bwr.callWithRetry(func() error {
		r, err := bwr.bc.GetProviderChannel(accountantAddress, addressToCheck, pending)
		if err != nil {
			return errors.Wrap(err, "could not get provider channel")
		}
		res = r
		return nil
	})

	return res, err
}

// SubscribeToPromiseSettledEvent subscribes to promise settled events
func (bwr *BlockchainWithRetries) SubscribeToPromiseSettledEvent(providerID, accountantID common.Address) (chan *bindings.AccountantImplementationPromiseSettled, func(), error) {
	var sink chan *bindings.AccountantImplementationPromiseSettled
	var cancel func()
	err := bwr.callWithRetry(func() error {
		s, c, err := bwr.bc.SubscribeToPromiseSettledEvent(providerID, accountantID)
		if err != nil {
			return errors.Wrap(err, "could not subscribe to settlement events")
		}
		sink = s
		cancel = c
		return nil
	})
	return sink, cancel, err
}

// SubscribeToConsumerBalanceEvent subscribes to the consumer balance change events
func (bwr *BlockchainWithRetries) SubscribeToConsumerBalanceEvent(channel, mystSCAddress common.Address, timeout time.Duration) (chan *bindings.MystTokenTransfer, func(), error) {
	var sink chan *bindings.MystTokenTransfer
	var cancel func()
	err := bwr.callWithRetry(func() error {
		s, c, err := bwr.bc.SubscribeToConsumerBalanceEvent(channel, mystSCAddress, timeout)
		if err != nil {
			return errors.Wrap(err, "could not subscribe to settlement events")
		}
		sink = s
		cancel = c
		return nil
	})
	return sink, cancel, err
}

// IsRegistered checks wether the given identity is registered or not
func (bwr *BlockchainWithRetries) IsRegistered(registryAddress, addressToCheck common.Address) (bool, error) {
	var res bool
	err := bwr.callWithRetry(func() error {
		r, err := bwr.bc.IsRegistered(registryAddress, addressToCheck)
		if err != nil {
			return errors.Wrap(err, "check registration status")
		}
		res = r
		return nil
	})
	return res, err
}

// GetMystBalance returns the balance in myst
func (bwr *BlockchainWithRetries) GetMystBalance(mystSCAddress, channel common.Address) (*big.Int, error) {
	var res *big.Int
	err := bwr.callWithRetry(func() error {
		result, bcErr := bwr.bc.GetMystBalance(mystSCAddress, channel)
		if bcErr != nil {
			return errors.Wrap(bcErr, "could not get myst balance")
		}
		res = result
		return nil
	})
	return res, err
}

// GetRegistrationFee returns the registration fee
func (bwr *BlockchainWithRetries) GetRegistrationFee(registryAddress common.Address) (*big.Int, error) {
	var res *big.Int
	err := bwr.callWithRetry(func() error {
		var bcErr error
		res, bcErr = bwr.bc.GetRegistrationFee(registryAddress)
		return bcErr
	})
	return res, err
}

// RegisterIdentity registers the given identity on blockchain
func (bwr *BlockchainWithRetries) RegisterIdentity(rr RegistrationRequest) (*types.Transaction, error) {
	var res *types.Transaction
	err := bwr.callWithRetry(func() error {
		result, bcErr := bwr.bc.RegisterIdentity(rr)
		if bcErr != nil {
			return errors.Wrap(bcErr, "could not register identity")
		}
		res = result
		return nil
	})
	return res, err
}

// TransferMyst transfers myst to the provided address
func (bwr *BlockchainWithRetries) TransferMyst(req TransferRequest) (tx *types.Transaction, err error) {
	err = bwr.callWithRetry(func() error {
		result, bcErr := bwr.bc.TransferMyst(req)
		if bcErr != nil {
			return errors.Wrap(bcErr, "could not transfer myst")
		}
		tx = result
		return nil
	})
	return tx, err
}

// IsAccountantRegistered checks if given accountant is registered and returns true or false.
func (bwr *BlockchainWithRetries) IsAccountantRegistered(registryAddress, acccountantID common.Address) (bool, error) {
	var res bool
	err := bwr.callWithRetry(func() error {
		result, bcErr := bwr.bc.IsAccountantRegistered(registryAddress, acccountantID)
		if bcErr != nil {
			return errors.Wrap(bcErr, "could not check if accountant is registered")
		}
		res = result
		return nil
	})
	return res, err
}

// GetAccountantOperator returns operator address of given accountant
func (bwr *BlockchainWithRetries) GetAccountantOperator(accountantID common.Address) (common.Address, error) {
	var res common.Address
	err := bwr.callWithRetry(func() error {
		result, bcErr := bwr.bc.GetAccountantOperator(accountantID)
		if bcErr != nil {
			return errors.Wrap(bcErr, "could not get accountant operator")
		}
		res = result
		return nil
	})
	return res, err
}

// SettleAndRebalance is settling given accountant issued promise
func (bwr *BlockchainWithRetries) SettleAndRebalance(req SettleAndRebalanceRequest) (*types.Transaction, error) {
	var res *types.Transaction
	err := bwr.callWithRetry(func() error {
		result, bcErr := bwr.bc.SettleAndRebalance(req)
		if bcErr != nil {
			return errors.Wrap(bcErr, "could not settle and rebalance")
		}
		res = result
		return nil
	})
	return res, err
}

// GetProviderChannelByID returns the given channel information
func (bwr *BlockchainWithRetries) GetProviderChannelByID(acc common.Address, chID []byte) (ProviderChannel, error) {
	var res ProviderChannel
	err := bwr.callWithRetry(func() error {
		result, bcErr := bwr.bc.GetProviderChannelByID(acc, chID)
		if bcErr != nil {
			return errors.Wrap(bcErr, "could not register identity")
		}
		res = result
		return nil
	})
	return res, err
}

// GetConsumerChannelsAccountant returns the consumer channels accountant
func (bwr *BlockchainWithRetries) GetConsumerChannelsAccountant(channelAddress common.Address) (ConsumersAccountant, error) {
	var res ConsumersAccountant
	err := bwr.callWithRetry(func() error {
		result, bcErr := bwr.bc.GetConsumerChannelsAccountant(channelAddress)
		if bcErr != nil {
			return errors.Wrap(bcErr, "could not get consumers accountant")
		}
		res = result
		return nil
	})
	return res, err
}

// GetConsumerChannelOperator returns the consumer channel operator/identity
func (bwr *BlockchainWithRetries) GetConsumerChannelOperator(channelAddress common.Address) (common.Address, error) {
	var res common.Address
	err := bwr.callWithRetry(func() error {
		result, bcErr := bwr.bc.GetConsumerChannelOperator(channelAddress)
		if bcErr != nil {
			return errors.Wrap(bcErr, "could not get consumer's operator")
		}
		res = result
		return nil
	})
	return res, err
}

// SubscribeToIdentityRegistrationEvents subscribes to identity registration events
func (bwr *BlockchainWithRetries) SubscribeToIdentityRegistrationEvents(registryAddress common.Address, accountantIDs []common.Address) (chan *bindings.RegistryRegisteredIdentity, func(), error) {
	var sink chan *bindings.RegistryRegisteredIdentity
	var cancel func()
	err := bwr.callWithRetry(func() error {
		s, c, err := bwr.bc.SubscribeToIdentityRegistrationEvents(registryAddress, accountantIDs)
		if err != nil {
			return errors.Wrap(err, "could not subscribe to registration events")
		}
		sink = s
		cancel = c
		return nil
	})
	return sink, cancel, err
}

// SubscribeToConsumerChannelBalanceUpdate subscribes to consumer channel balance update events
func (bwr *BlockchainWithRetries) SubscribeToConsumerChannelBalanceUpdate(mystSCAddress common.Address, channelAddresses []common.Address) (chan *bindings.MystTokenTransfer, func(), error) {
	var sink chan *bindings.MystTokenTransfer
	var cancel func()
	err := bwr.callWithRetry(func() error {
		s, c, err := bwr.bc.SubscribeToConsumerChannelBalanceUpdate(mystSCAddress, channelAddresses)
		if err != nil {
			return errors.Wrap(err, "could not subscribe to channel balance events")
		}
		sink = s
		cancel = c
		return nil
	})
	return sink, cancel, err
}

// SubscribeToProviderChannelBalanceUpdate subscribes to provider channel balance update events
func (bwr *BlockchainWithRetries) SubscribeToProviderChannelBalanceUpdate(accountantAddress common.Address, channelAddresses [][32]byte) (chan *bindings.AccountantImplementationChannelBalanceUpdated, func(), error) {
	var sink chan *bindings.AccountantImplementationChannelBalanceUpdated
	var cancel func()
	err := bwr.callWithRetry(func() error {
		s, c, err := bwr.bc.SubscribeToProviderChannelBalanceUpdate(accountantAddress, channelAddresses)
		if err != nil {
			return errors.Wrap(err, "could not subscribe to channel balance events")
		}
		sink = s
		cancel = c
		return nil
	})
	return sink, cancel, err
}

// SettlePromise is settling the given consumer issued promise
func (bwr *BlockchainWithRetries) SettlePromise(req SettleRequest) (*types.Transaction, error) {
	var res *types.Transaction
	err := bwr.callWithRetry(func() error {
		result, bcErr := bwr.bc.SettlePromise(req)
		if bcErr != nil {
			return errors.Wrap(bcErr, "could not settle promise")
		}
		res = result
		return nil
	})
	return res, err
}

// GetEthBalance gets the current ethereum balance for the address.
func (bwr *BlockchainWithRetries) GetEthBalance(address common.Address) (*big.Int, error) {
	var res *big.Int
	err := bwr.callWithRetry(func() error {
		result, bcErr := bwr.bc.GetEthBalance(address)
		if bcErr != nil {
			return errors.Wrap(bcErr, "could not get balance")
		}
		res = result
		return nil
	})
	return res, err
}

// TransferEth transfers ethereum to the given address.
func (bwr *BlockchainWithRetries) TransferEth(etr EthTransferRequest) (*types.Transaction, error) {
	var res *types.Transaction
	err := bwr.callWithRetry(func() error {
		result, bcErr := bwr.bc.TransferEth(etr)
		if bcErr != nil {
			return errors.Wrap(bcErr, "could not transfer ethereum")
		}
		res = result
		return nil
	})
	return res, err
}

// SubscribeToChannelOpenedEvents subscribes to provider channel opened events
func (bwr *BlockchainWithRetries) SubscribeToChannelOpenedEvents(accountantAddress common.Address) (chan *bindings.AccountantImplementationChannelOpened, func(), error) {
	var sink chan *bindings.AccountantImplementationChannelOpened
	var cancel func()
	err := bwr.callWithRetry(func() error {
		s, c, err := bwr.bc.SubscribeToChannelOpenedEvents(accountantAddress)
		if err != nil {
			return errors.Wrap(err, "could not subscribe to channel opened events")
		}
		sink = s
		cancel = c
		return nil
	})
	return sink, cancel, err
}

// SubscribeToPromiseSettledEventByChannelID subscribes to promise settled events
func (bwr *BlockchainWithRetries) SubscribeToPromiseSettledEventByChannelID(accountantID common.Address, providerAddresses [][32]byte) (chan *bindings.AccountantImplementationPromiseSettled, func(), error) {
	var sink chan *bindings.AccountantImplementationPromiseSettled
	var cancel func()
	err := bwr.callWithRetry(func() error {
		s, c, err := bwr.bc.SubscribeToPromiseSettledEventByChannelID(accountantID, providerAddresses)
		if err != nil {
			return errors.Wrap(err, "could not subscribe to settlement events")
		}
		sink = s
		cancel = c
		return nil
	})
	return sink, cancel, err
}

// GetConsumerChannel returns the consumer channel
func (bwr *BlockchainWithRetries) GetConsumerChannel(addr common.Address, mystSCAddress common.Address) (ConsumerChannel, error) {
	var res ConsumerChannel
	err := bwr.callWithRetry(func() error {
		result, bcErr := bwr.bc.GetConsumerChannel(addr, mystSCAddress)
		if bcErr != nil {
			return errors.Wrap(bcErr, "could not get consumers channel")
		}
		res = result
		return nil
	})
	return res, err
}

// NetworkID returns the network id
func (bwr *BlockchainWithRetries) NetworkID() (*big.Int, error) {
	var res *big.Int
	err := bwr.callWithRetry(func() error {
		result, bcErr := bwr.bc.NetworkID()
		if bcErr != nil {
			return errors.Wrap(bcErr, "could not get network ID")
		}
		res = result
		return nil
	})
	return res, err
}

// EstimateGas proxies the estimate gas call to the underlying blockchain since no network calls are performed.
func (bwr *BlockchainWithRetries) EstimateGas(msg ethereum.CallMsg) (uint64, error) {
	return bwr.bc.EstimateGas(msg)
}

// Stop stops the blockchain with retries aborting any waits for retries
func (bwr *BlockchainWithRetries) Stop() {
	bwr.once.Do(func() {
		close(bwr.stop)
	})
}

// SettleWithBeneficiary is setting new beneficiary and settling latest promise balance into new beneficiary address.
func (bwr *BlockchainWithRetries) SettleWithBeneficiary(req SettleWithBeneficiaryRequest) (*types.Transaction, error) {
	var res *types.Transaction
	err := bwr.callWithRetry(func() error {
		result, bcErr := bwr.bc.SettleWithBeneficiary(req)
		if bcErr != nil {
			return errors.Wrap(bcErr, "could not set beneficiary")
		}
		res = result
		return nil
	})
	return res, err
}
