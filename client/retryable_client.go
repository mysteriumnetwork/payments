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

type BC interface {
	GetHermesFee(hermesAddress common.Address) (uint16, error)
	CalculateHermesFee(hermesAddress common.Address, value *big.Int) (*big.Int, error)
	IsRegisteredAsProvider(hermesAddress, registryAddress, addressToCheck common.Address) (bool, error)
	GetProviderChannel(hermesAddress common.Address, addressToCheck common.Address, pending bool) (ProviderChannel, error)
	IsRegistered(registryAddress, addressToCheck common.Address) (bool, error)
	SubscribeToPromiseSettledEvent(providerID, hermesID common.Address) (sink chan *bindings.HermesImplementationPromiseSettled, cancel func(), err error)
	GetMystBalance(mystSCAddress, address common.Address) (*big.Int, error)
	SubscribeToConsumerBalanceEvent(channel, mystSCAddress common.Address, timeout time.Duration) (chan *bindings.MystTokenTransfer, func(), error)
	RegisterIdentity(rr RegistrationRequest) (*types.Transaction, error)
	TransferMyst(req TransferRequest) (tx *types.Transaction, err error)
	IsHermesRegistered(registryAddress, acccountantID common.Address) (bool, error)
	GetHermesOperator(hermesID common.Address) (common.Address, error)
	SettleAndRebalance(req SettleAndRebalanceRequest) (*types.Transaction, error)
	SettleWithBeneficiary(req SettleWithBeneficiaryRequest) (*types.Transaction, error)
	GetConsumerChannelsHermes(channelAddress common.Address) (ConsumersHermes, error)
	GetConsumerChannelOperator(channelAddress common.Address) (common.Address, error)
	GetProviderChannelByID(acc common.Address, chID []byte) (ProviderChannel, error)
	SubscribeToIdentityRegistrationEvents(registryAddress common.Address) (sink chan *bindings.RegistryRegisteredIdentity, cancel func(), err error)
	SubscribeToConsumerChannelBalanceUpdate(mystSCAddress common.Address, channelAddresses []common.Address) (sink chan *bindings.MystTokenTransfer, cancel func(), err error)
	SettlePromise(req SettleRequest) (*types.Transaction, error)
	SubscribeToPromiseSettledEventByChannelID(hermesID common.Address, providerAddresses [][32]byte) (sink chan *bindings.HermesImplementationPromiseSettled, cancel func(), err error)
	SubscribeToMystTokenTransfers(mystSCAddress common.Address) (chan *bindings.MystTokenTransfer, func(), error)
	NetworkID() (*big.Int, error)
	GetConsumerChannel(addr common.Address, mystSCAddress common.Address) (ConsumerChannel, error)
	GetEthBalance(address common.Address) (*big.Int, error)
	TransferEth(etr EthTransferRequest) (*types.Transaction, error)
	GetHermessAvailableBalance(hermesAddress common.Address) (*big.Int, error)
	DecreaseProviderStake(req DecreaseProviderStakeRequest) (*types.Transaction, error)
	SettleIntoStake(req SettleIntoStakeRequest) (*types.Transaction, error)
	IncreaseProviderStake(req ProviderStakeIncreaseRequest) (*types.Transaction, error)
	TransactionReceipt(hash common.Hash) (*types.Receipt, error)
	GetHermesURL(registryID, hermesID common.Address) (string, error)
	GetStakeThresholds(hermesID common.Address) (min, max *big.Int, err error)
	GetBeneficiary(registryAddress, identity common.Address) (common.Address, error)
	SuggestGasPrice() (*big.Int, error)
	FilterLogs(q ethereum.FilterQuery) ([]types.Log, error)
	HeaderByNumber(number *big.Int) (*types.Header, error)
	GetLastRegistryNonce(registry common.Address) (*big.Int, error)
	SendTransaction(tx *types.Transaction) error
	IsHermesActive(hermesID common.Address) (bool, error)
	GetHermes(registryID, hermesID common.Address) (Hermes, error)
	GetChannelImplementationByVersion(registryID common.Address, version *big.Int) (common.Address, error)
	TransactionByHash(hash common.Hash) (*types.Transaction, bool, error)
	RewarderTotalPayoutsFor(rewarderAddress common.Address, payoutsFor common.Address) (*big.Int, error)
	RewarderAirDrop(req RewarderAirDrop) (*types.Transaction, error)
	RewarderUpdateRoot(req RewarderUpdateRoot) (*types.Transaction, error)
	RewarderTotalClaimed(rewarderAddress common.Address) (*big.Int, error)
}

// BlockchainWithRetries takes in the plain blockchain implementation and exposes methods that will retry the underlying bc methods before giving up.
// This is required as the ethereum client will occasionally spit a TLS error if running for prolonged periods of time.
type BlockchainWithRetries struct {
	delay      time.Duration
	maxRetries int
	bc         BC
	stop       chan struct{}
	once       sync.Once
}

// ErrStopped represents an error when a call is interrupted
var ErrStopped = errors.New("call stopped")

// NewBlockchainWithRetries returns a new instance of blockchain with retries
func NewBlockchainWithRetries(bc BC, delay time.Duration, maxRetries int) *BlockchainWithRetries {
	return &BlockchainWithRetries{
		bc:         bc,
		delay:      delay,
		maxRetries: maxRetries,
	}
}

// FilterLogs executes a filter query.
func (bwr *BlockchainWithRetries) FilterLogs(q ethereum.FilterQuery) ([]types.Log, error) {
	var res []types.Log
	err := bwr.callWithRetry(func() error {
		r, err := bwr.bc.FilterLogs(q)
		if err != nil {
			return errors.Wrap(err, "could not filter logs")
		}
		res = r
		return nil
	})
	return res, err
}

func (bwr *BlockchainWithRetries) GetLastRegistryNonce(registry common.Address) (*big.Int, error) {
	var res *big.Int
	err := bwr.callWithRetry(func() error {
		r, err := bwr.bc.GetLastRegistryNonce(registry)
		if err != nil {
			return errors.Wrap(err, "could not get registry nonce")
		}
		res = r
		return nil
	})
	return res, err
}

// HeaderByNumber returns a block header from the current canonical chain. If number is
// nil, the latest known header is returned.
func (bwr *BlockchainWithRetries) HeaderByNumber(number *big.Int) (*types.Header, error) {
	var res *types.Header
	err := bwr.callWithRetry(func() error {
		r, err := bwr.bc.HeaderByNumber(number)
		if err != nil {
			return errors.Wrap(err, "could not get header by number")
		}
		res = r
		return nil
	})
	return res, err
}

func (bwr *BlockchainWithRetries) SuggestGasPrice() (*big.Int, error) {
	var res *big.Int
	err := bwr.callWithRetry(func() error {
		r, err := bwr.bc.SuggestGasPrice()
		if err != nil {
			return errors.Wrap(err, "could not get gas price")
		}
		res = r
		return nil
	})
	return res, err
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

// GetHermesFee fetches the hermes fee from blockchain
func (bwr *BlockchainWithRetries) GetHermesFee(hermesAddress common.Address) (uint16, error) {
	var res uint16
	err := bwr.callWithRetry(func() error {
		r, err := bwr.bc.GetHermesFee(hermesAddress)
		if err != nil {
			return errors.Wrap(err, "could not get hermes fee")
		}
		res = r
		return nil
	})
	return res, err
}

func (bwr *BlockchainWithRetries) GetBeneficiary(registryAddress, identity common.Address) (common.Address, error) {
	var res common.Address
	err := bwr.callWithRetry(func() error {
		r, err := bwr.bc.GetBeneficiary(registryAddress, identity)
		if err != nil {
			return errors.Wrap(err, "could not get beneficiary")
		}
		res = r
		return nil
	})
	return res, err
}

// CalculateHermesFee fetches the hermes fee from blockchain
func (bwr *BlockchainWithRetries) CalculateHermesFee(hermesAddress common.Address, value *big.Int) (*big.Int, error) {
	var res *big.Int
	err := bwr.callWithRetry(func() error {
		r, err := bwr.bc.CalculateHermesFee(hermesAddress, value)
		if err != nil {
			return errors.Wrap(err, "could not calculate hermes fee")
		}
		res = r
		return nil
	})
	return res, err
}

// IsRegisteredAsProvider checks if the provider is registered with the hermes properly
func (bwr *BlockchainWithRetries) IsRegisteredAsProvider(hermesAddress, registryAddress, addressToCheck common.Address) (bool, error) {
	var res bool
	err := bwr.callWithRetry(func() error {
		r, err := bwr.bc.IsRegisteredAsProvider(hermesAddress, registryAddress, addressToCheck)
		if err != nil {
			return errors.Wrap(err, "could not check if registered as provider")
		}
		res = r
		return nil
	})
	return res, err
}

// GetProviderChannel returns the provider channel
func (bwr *BlockchainWithRetries) GetProviderChannel(hermesAddress, addressToCheck common.Address, pending bool) (ProviderChannel, error) {
	var res ProviderChannel
	err := bwr.callWithRetry(func() error {
		r, err := bwr.bc.GetProviderChannel(hermesAddress, addressToCheck, pending)
		if err != nil {
			return errors.Wrap(err, "could not get provider channel")
		}
		res = r
		return nil
	})

	return res, err
}

// SubscribeToPromiseSettledEvent subscribes to promise settled events
func (bwr *BlockchainWithRetries) SubscribeToPromiseSettledEvent(providerID, hermesID common.Address) (chan *bindings.HermesImplementationPromiseSettled, func(), error) {
	var sink chan *bindings.HermesImplementationPromiseSettled
	var cancel func()
	err := bwr.callWithRetry(func() error {
		s, c, err := bwr.bc.SubscribeToPromiseSettledEvent(providerID, hermesID)
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

// IsHermesRegistered checks if given hermes is registered and returns true or false.
func (bwr *BlockchainWithRetries) IsHermesRegistered(registryAddress, acccountantID common.Address) (bool, error) {
	var res bool
	err := bwr.callWithRetry(func() error {
		result, bcErr := bwr.bc.IsHermesRegistered(registryAddress, acccountantID)
		if bcErr != nil {
			return errors.Wrap(bcErr, "could not check if hermes is registered")
		}
		res = result
		return nil
	})
	return res, err
}

// GetHermesOperator returns operator address of given hermes
func (bwr *BlockchainWithRetries) GetHermesOperator(hermesID common.Address) (common.Address, error) {
	var res common.Address
	err := bwr.callWithRetry(func() error {
		result, bcErr := bwr.bc.GetHermesOperator(hermesID)
		if bcErr != nil {
			return errors.Wrap(bcErr, "could not get hermes operator")
		}
		res = result
		return nil
	})
	return res, err
}

// SettleAndRebalance is settling given hermes issued promise
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

// GetConsumerChannelsHermes returns the consumer channels hermes
func (bwr *BlockchainWithRetries) GetConsumerChannelsHermes(channelAddress common.Address) (ConsumersHermes, error) {
	var res ConsumersHermes
	err := bwr.callWithRetry(func() error {
		result, bcErr := bwr.bc.GetConsumerChannelsHermes(channelAddress)
		if bcErr != nil {
			return errors.Wrap(bcErr, "could not get consumers hermes")
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

func (bwr *BlockchainWithRetries) TransactionReceipt(hash common.Hash) (*types.Receipt, error) {
	var res *types.Receipt
	err := bwr.callWithRetry(func() error {
		result, bcErr := bwr.bc.TransactionReceipt(hash)
		if bcErr != nil {
			return errors.Wrap(bcErr, "could not get transaction receipt")
		}
		res = result
		return nil
	})
	return res, err
}

// SubscribeToIdentityRegistrationEvents subscribes to identity registration events
func (bwr *BlockchainWithRetries) SubscribeToIdentityRegistrationEvents(registryAddress common.Address) (chan *bindings.RegistryRegisteredIdentity, func(), error) {
	var sink chan *bindings.RegistryRegisteredIdentity
	var cancel func()
	err := bwr.callWithRetry(func() error {
		s, c, err := bwr.bc.SubscribeToIdentityRegistrationEvents(registryAddress)
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

// GetHermessAvailableBalance returns the balance that is available for hermes.
func (bwr *BlockchainWithRetries) GetHermessAvailableBalance(hermesAddress common.Address) (*big.Int, error) {
	var res *big.Int
	err := bwr.callWithRetry(func() error {
		result, bcErr := bwr.bc.GetHermessAvailableBalance(hermesAddress)
		if bcErr != nil {
			return errors.Wrap(bcErr, "could not get balance")
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

// GetHermesURL returns the hermes URL.
func (bwr *BlockchainWithRetries) GetHermesURL(registryID, hermesID common.Address) (string, error) {
	var res string
	err := bwr.callWithRetry(func() error {
		result, bcErr := bwr.bc.GetHermesURL(registryID, hermesID)
		if bcErr != nil {
			return errors.Wrap(bcErr, "could not get hermes url")
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

// SubscribeToPromiseSettledEventByChannelID subscribes to promise settled events
func (bwr *BlockchainWithRetries) SubscribeToPromiseSettledEventByChannelID(hermesID common.Address, providerAddresses [][32]byte) (chan *bindings.HermesImplementationPromiseSettled, func(), error) {
	var sink chan *bindings.HermesImplementationPromiseSettled
	var cancel func()
	err := bwr.callWithRetry(func() error {
		s, c, err := bwr.bc.SubscribeToPromiseSettledEventByChannelID(hermesID, providerAddresses)
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

// DecreaseProviderStake decreases provider stake.
func (bwr *BlockchainWithRetries) DecreaseProviderStake(req DecreaseProviderStakeRequest) (*types.Transaction, error) {
	var res *types.Transaction
	err := bwr.callWithRetry(func() error {
		result, bcErr := bwr.bc.DecreaseProviderStake(req)
		if bcErr != nil {
			return errors.Wrap(bcErr, "could not set beneficiary")
		}
		res = result
		return nil
	})
	return res, err
}

// SettleIntoStake settles the hermes promise into stake increase.
func (bwr *BlockchainWithRetries) SettleIntoStake(req SettleIntoStakeRequest) (*types.Transaction, error) {
	var res *types.Transaction
	err := bwr.callWithRetry(func() error {
		result, bcErr := bwr.bc.SettleIntoStake(req)
		if bcErr != nil {
			return errors.Wrap(bcErr, "could not set beneficiary")
		}
		res = result
		return nil
	})
	return res, err
}

// IncreaseProviderStake increases the provider stake.
func (bwr *BlockchainWithRetries) IncreaseProviderStake(req ProviderStakeIncreaseRequest) (*types.Transaction, error) {
	var res *types.Transaction
	err := bwr.callWithRetry(func() error {
		result, bcErr := bwr.bc.IncreaseProviderStake(req)
		if bcErr != nil {
			return errors.Wrap(bcErr, "could not set beneficiary")
		}
		res = result
		return nil
	})
	return res, err
}

// GetStakeThresholds returns the stake tresholds for the given hermes.
func (bwr *BlockchainWithRetries) GetStakeThresholds(hermesID common.Address) (min, max *big.Int, err error) {
	err = bwr.callWithRetry(func() error {
		m, ma, bcErr := bwr.bc.GetStakeThresholds(hermesID)
		if bcErr != nil {
			return errors.Wrap(bcErr, "could not set beneficiary")
		}
		min = m
		max = ma
		return nil
	})
	return min, max, err
}

func (bwr *BlockchainWithRetries) SendTransaction(tx *types.Transaction) error {
	return bwr.callWithRetry(func() error {
		if err := bwr.bc.SendTransaction(tx); err != nil {
			return errors.Wrap(err, "could not send transaction to bc")
		}
		return nil
	})
}

func (bwr *BlockchainWithRetries) IsHermesActive(hermesID common.Address) (bool, error) {
	res := false
	err := bwr.callWithRetry(func() error {
		r, err := bwr.bc.IsHermesActive(hermesID)
		if err != nil {
			return err
		}
		res = r
		return nil
	})
	return res, err
}

func (bwr *BlockchainWithRetries) GetChannelImplementationByVersion(registryID common.Address, version *big.Int) (common.Address, error) {
	var res common.Address
	err := bwr.callWithRetry(func() error {
		addr, err := bwr.bc.GetChannelImplementationByVersion(registryID, version)
		if err != nil {
			return errors.Wrap(err, "could not get channel implementation version")
		}
		res = addr
		return nil
	})
	return res, err
}

func (bwr *BlockchainWithRetries) GetHermes(registryID, hermesID common.Address) (Hermes, error) {
	var res Hermes
	err := bwr.callWithRetry(func() error {
		h, err := bwr.bc.GetHermes(registryID, hermesID)
		if err != nil {
			return errors.Wrap(err, "could not get hermes")
		}
		res = h
		return nil
	})
	return res, err
}

func (bwr *BlockchainWithRetries) TransactionByHash(hash common.Hash) (*types.Transaction, bool, error) {
	var res *types.Transaction
	var ok bool
	err := bwr.callWithRetry(func() error {
		h, k, err := bwr.bc.TransactionByHash(hash)
		if err != nil {
			return errors.Wrap(err, "could not get transaction by hash")
		}
		res = h
		ok = k
		return nil
	})
	return res, ok, err
}

func (bwr *BlockchainWithRetries) RewarderTotalPayoutsFor(rewarderAddress common.Address, payoutsFor common.Address) (*big.Int, error) {
	var total *big.Int
	err := bwr.callWithRetry(func() error {
		t, err := bwr.bc.RewarderTotalPayoutsFor(rewarderAddress, payoutsFor)
		if err != nil {
			return errors.Wrap(err, "could get total payouts for address")
		}
		total = t
		return nil
	})
	return total, err
}

func (bwr *BlockchainWithRetries) RewarderAirDrop(req RewarderAirDrop) (*types.Transaction, error) {
	var res *types.Transaction
	err := bwr.callWithRetry(func() error {
		tx, err := bwr.bc.RewarderAirDrop(req)
		if err != nil {
			return errors.Wrap(err, "could send air drop request")
		}
		res = tx
		return nil
	})
	return res, err
}

func (bwr *BlockchainWithRetries) RewarderUpdateRoot(req RewarderUpdateRoot) (*types.Transaction, error) {
	var res *types.Transaction
	err := bwr.callWithRetry(func() error {
		tx, err := bwr.bc.RewarderUpdateRoot(req)
		if err != nil {
			return errors.Wrap(err, "could send update root request")
		}
		res = tx
		return nil
	})
	return res, err
}

// RewarderTotalClaimed is a free lookup in the blockchain for the total amount of claimed tokens in the blockchain.
func (bwr *BlockchainWithRetries) RewarderTotalClaimed(chainID int64, rewarderAddress common.Address) (*big.Int, error) {
	var total *big.Int
	err := bwr.callWithRetry(func() error {
		t, err := bwr.bc.RewarderTotalClaimed(rewarderAddress)
		if err != nil {
			return errors.Wrap(err, "could get total claimed")
		}
		total = t
		return nil
	})
	return total, err
}
