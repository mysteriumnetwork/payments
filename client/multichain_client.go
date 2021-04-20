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
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/mysteriumnetwork/payments/bindings"
	"github.com/pkg/errors"
)

type MultichainBlockchainClient struct {
	clients map[int64]BC
}

func NewMultichainBlockchainClient(clients map[int64]BC) *MultichainBlockchainClient {
	return &MultichainBlockchainClient{
		clients: clients,
	}
}

var ErrUnknownChain = errors.New("unknown chain")

func (mbc *MultichainBlockchainClient) getClientByChain(chainID int64) (BC, error) {
	if v, ok := mbc.clients[chainID]; ok {
		return v, nil
	}
	return nil, ErrUnknownChain
}

func (mbc *MultichainBlockchainClient) GetHermesFee(chainID int64, hermesAddress common.Address) (uint16, error) {
	bc, err := mbc.getClientByChain(chainID)
	if err != nil {
		return 0, err
	}

	return bc.GetHermesFee(hermesAddress)
}

func (mbc *MultichainBlockchainClient) CalculateHermesFee(chainID int64, hermesAddress common.Address, value *big.Int) (*big.Int, error) {
	bc, err := mbc.getClientByChain(chainID)
	if err != nil {
		return nil, err
	}

	return bc.CalculateHermesFee(hermesAddress, value)
}

func (mbc *MultichainBlockchainClient) IsRegisteredAsProvider(chainID int64, hermesAddress, registryAddress, addressToCheck common.Address) (bool, error) {
	bc, err := mbc.getClientByChain(chainID)
	if err != nil {
		return false, err
	}

	return bc.IsRegisteredAsProvider(hermesAddress, registryAddress, addressToCheck)
}

func (mbc *MultichainBlockchainClient) GetProviderChannel(chainID int64, hermesAddress common.Address, addressToCheck common.Address, pending bool) (ProviderChannel, error) {
	bc, err := mbc.getClientByChain(chainID)
	if err != nil {
		return ProviderChannel{}, err
	}

	return bc.GetProviderChannel(hermesAddress, addressToCheck, pending)
}

func (mbc *MultichainBlockchainClient) IsRegistered(chainID int64, registryAddress, addressToCheck common.Address) (bool, error) {
	bc, err := mbc.getClientByChain(chainID)
	if err != nil {
		return false, err
	}

	return bc.IsRegistered(registryAddress, addressToCheck)
}

func (mbc *MultichainBlockchainClient) SubscribeToPromiseSettledEvent(chainID int64, providerID, hermesID common.Address) (sink chan *bindings.HermesImplementationPromiseSettled, cancel func(), err error) {
	bc, err := mbc.getClientByChain(chainID)
	if err != nil {
		return nil, func() {}, err
	}

	return bc.SubscribeToPromiseSettledEvent(providerID, hermesID)
}

func (mbc *MultichainBlockchainClient) GetMystBalance(chainID int64, mystSCAddress, address common.Address) (*big.Int, error) {
	bc, err := mbc.getClientByChain(chainID)
	if err != nil {
		return nil, err
	}

	return bc.GetMystBalance(mystSCAddress, address)
}

func (mbc *MultichainBlockchainClient) SubscribeToConsumerBalanceEvent(chainID int64, channel, mystSCAddress common.Address, timeout time.Duration) (chan *bindings.MystTokenTransfer, func(), error) {
	bc, err := mbc.getClientByChain(chainID)
	if err != nil {
		return nil, func() {}, err
	}

	return bc.SubscribeToConsumerBalanceEvent(channel, mystSCAddress, timeout)
}

func (mbc *MultichainBlockchainClient) SubscribeToIdentityRegistrationEvents(chainID int64, registryAddress common.Address) (sink chan *bindings.RegistryRegisteredIdentity, cancel func(), err error) {
	bc, err := mbc.getClientByChain(chainID)
	if err != nil {
		return nil, func() {}, err
	}

	return bc.SubscribeToIdentityRegistrationEvents(registryAddress)
}

func (mbc *MultichainBlockchainClient) SuggestGasPrice(chainID int64) (*big.Int, error) {
	bc, err := mbc.getClientByChain(chainID)
	if err != nil {
		return nil, err
	}
	return bc.SuggestGasPrice()
}

func (mbc *MultichainBlockchainClient) SubscribeToConsumerChannelBalanceUpdate(chainID int64, mystSCAddress common.Address, channelAddresses []common.Address) (sink chan *bindings.MystTokenTransfer, cancel func(), err error) {
	bc, err := mbc.getClientByChain(chainID)
	if err != nil {
		return nil, func() {}, err
	}

	return bc.SubscribeToConsumerChannelBalanceUpdate(mystSCAddress, channelAddresses)
}
func (mbc *MultichainBlockchainClient) SubscribeToPromiseSettledEventByChannelID(chainID int64, hermesID common.Address, providerAddresses [][32]byte) (sink chan *bindings.HermesImplementationPromiseSettled, cancel func(), err error) {
	bc, err := mbc.getClientByChain(chainID)
	if err != nil {
		return nil, func() {}, err
	}

	return bc.SubscribeToPromiseSettledEventByChannelID(hermesID, providerAddresses)
}

func (mbc *MultichainBlockchainClient) SubscribeToMystTokenTransfers(chainID int64, mystSCAddress common.Address) (chan *bindings.MystTokenTransfer, func(), error) {
	bc, err := mbc.getClientByChain(chainID)
	if err != nil {
		return nil, func() {}, err
	}

	return bc.SubscribeToMystTokenTransfers(mystSCAddress)
}

func (mbc *MultichainBlockchainClient) RegisterIdentity(chainID int64, rr RegistrationRequest) (*types.Transaction, error) {
	bc, err := mbc.getClientByChain(chainID)
	if err != nil {
		return nil, err
	}

	return bc.RegisterIdentity(rr)
}

func (mbc *MultichainBlockchainClient) TransferMyst(chainID int64, req TransferRequest) (tx *types.Transaction, err error) {
	bc, err := mbc.getClientByChain(chainID)
	if err != nil {
		return nil, err
	}

	return bc.TransferMyst(req)
}

func (mbc *MultichainBlockchainClient) IsHermesRegistered(chainID int64, registryAddress, hermesID common.Address) (bool, error) {
	bc, err := mbc.getClientByChain(chainID)
	if err != nil {
		return false, err
	}

	return bc.IsHermesRegistered(registryAddress, hermesID)
}

func (mbc *MultichainBlockchainClient) GetHermesOperator(chainID int64, hermesID common.Address) (common.Address, error) {
	bc, err := mbc.getClientByChain(chainID)
	if err != nil {
		return common.Address{}, err
	}

	return bc.GetHermesOperator(hermesID)
}

func (mbc *MultichainBlockchainClient) SettleAndRebalance(req SettleAndRebalanceRequest) (*types.Transaction, error) {
	bc, err := mbc.getClientByChain(req.Promise.ChainID)
	if err != nil {
		return nil, err
	}

	return bc.SettleAndRebalance(req)
}

func (mbc *MultichainBlockchainClient) GetLastRegistryNonce(chainID int64, registry common.Address) (*big.Int, error) {
	bc, err := mbc.getClientByChain(chainID)
	if err != nil {
		return nil, err
	}
	return bc.GetLastRegistryNonce(registry)
}

func (mbc *MultichainBlockchainClient) GetBeneficiary(chainID int64, registryAddress, identity common.Address) (common.Address, error) {
	bc, err := mbc.getClientByChain(chainID)
	if err != nil {
		return common.Address{}, err
	}

	return bc.GetBeneficiary(registryAddress, identity)
}

func (mbc *MultichainBlockchainClient) SettleWithBeneficiary(req SettleWithBeneficiaryRequest) (*types.Transaction, error) {
	bc, err := mbc.getClientByChain(req.Promise.ChainID)
	if err != nil {
		return nil, err
	}

	return bc.SettleWithBeneficiary(req)
}

func (mbc *MultichainBlockchainClient) GetConsumerChannelsHermes(chainID int64, channelAddress common.Address) (ConsumersHermes, error) {
	bc, err := mbc.getClientByChain(chainID)
	if err != nil {
		return ConsumersHermes{}, err
	}

	return bc.GetConsumerChannelsHermes(channelAddress)
}

func (mbc *MultichainBlockchainClient) GetConsumerChannelOperator(chainID int64, channelAddress common.Address) (common.Address, error) {
	bc, err := mbc.getClientByChain(chainID)
	if err != nil {
		return common.Address{}, err
	}

	return bc.GetConsumerChannelOperator(channelAddress)
}

func (mbc *MultichainBlockchainClient) GetProviderChannelByID(chainID int64, acc common.Address, chID []byte) (ProviderChannel, error) {
	bc, err := mbc.getClientByChain(chainID)
	if err != nil {
		return ProviderChannel{}, err
	}

	return bc.GetProviderChannelByID(acc, chID)
}

func (mbc *MultichainBlockchainClient) SettlePromise(req SettleRequest) (*types.Transaction, error) {
	bc, err := mbc.getClientByChain(req.Promise.ChainID)
	if err != nil {
		return nil, err
	}

	return bc.SettlePromise(req)
}

// NetworkID method is really not that useful, as chain id == networkid
func (mbc *MultichainBlockchainClient) NetworkID(chainID int64) (*big.Int, error) {
	bc, err := mbc.getClientByChain(chainID)
	if err != nil {
		return nil, err
	}

	return bc.NetworkID()
}

func (mbc *MultichainBlockchainClient) GetConsumerChannel(chainID int64, addr common.Address, mystSCAddress common.Address) (ConsumerChannel, error) {
	bc, err := mbc.getClientByChain(chainID)
	if err != nil {
		return ConsumerChannel{}, err
	}

	return bc.GetConsumerChannel(addr, mystSCAddress)
}

func (mbc *MultichainBlockchainClient) GetEthBalance(chainID int64, address common.Address) (*big.Int, error) {
	bc, err := mbc.getClientByChain(chainID)
	if err != nil {
		return nil, err
	}

	return bc.GetEthBalance(address)
}

func (mbc *MultichainBlockchainClient) TransactionReceipt(chainID int64, hash common.Hash) (*types.Receipt, error) {
	bc, err := mbc.getClientByChain(chainID)
	if err != nil {
		return nil, err
	}
	return bc.TransactionReceipt(hash)
}

func (mbc *MultichainBlockchainClient) TransferEth(chainID int64, etr EthTransferRequest) (*types.Transaction, error) {
	bc, err := mbc.getClientByChain(chainID)
	if err != nil {
		return nil, err
	}

	return bc.TransferEth(etr)
}

func (mbc *MultichainBlockchainClient) GetHermessAvailableBalance(chainID int64, hermesAddress common.Address) (*big.Int, error) {
	bc, err := mbc.getClientByChain(chainID)
	if err != nil {
		return nil, err
	}

	return bc.GetHermessAvailableBalance(hermesAddress)
}

func (mbc *MultichainBlockchainClient) DecreaseProviderStake(chainID int64, req DecreaseProviderStakeRequest) (*types.Transaction, error) {
	bc, err := mbc.getClientByChain(chainID)
	if err != nil {
		return nil, err
	}

	return bc.DecreaseProviderStake(req)
}

func (mbc *MultichainBlockchainClient) SettleIntoStake(req SettleIntoStakeRequest) (*types.Transaction, error) {
	bc, err := mbc.getClientByChain(req.Promise.ChainID)
	if err != nil {
		return nil, err
	}

	return bc.SettleIntoStake(req)
}

func (mbc *MultichainBlockchainClient) IncreaseProviderStake(chainID int64, req ProviderStakeIncreaseRequest) (*types.Transaction, error) {
	bc, err := mbc.getClientByChain(chainID)
	if err != nil {
		return nil, err
	}

	return bc.IncreaseProviderStake(req)
}

func (mbc *MultichainBlockchainClient) GetHermesURL(chainID int64, registryID, hermesID common.Address) (string, error) {
	bc, err := mbc.getClientByChain(chainID)
	if err != nil {
		return "", err
	}

	return bc.GetHermesURL(registryID, hermesID)
}

// GetStakeThresholds returns the stake tresholds for the given hermes.
func (mbc *MultichainBlockchainClient) GetStakeThresholds(chainID int64, hermesID common.Address) (min, max *big.Int, err error) {
	bc, err := mbc.getClientByChain(chainID)
	if err != nil {
		return nil, nil, err
	}

	return bc.GetStakeThresholds(hermesID)
}

func (mbc *MultichainBlockchainClient) GetHermes(chainID int64, registryID, hermesID common.Address) (Hermes, error) {
	bc, err := mbc.getClientByChain(chainID)
	if err != nil {
		return Hermes{}, err
	}

	return bc.GetHermes(registryID, hermesID)
}

func (mbc *MultichainBlockchainClient) GetChannelImplementationByVersion(chainID int64, registryID common.Address, version *big.Int) (common.Address, error) {
	bc, err := mbc.getClientByChain(chainID)
	if err != nil {
		return common.Address{}, err
	}

	return bc.GetChannelImplementationByVersion(registryID, version)
}

func (mbc *MultichainBlockchainClient) IsChannelOpened(chainID int64, registryID, identity, hermesID common.Address) (bool, error) {
	bc, err := mbc.getClientByChain(chainID)
	if err != nil {
		return false, err
	}

	return bc.IsChannelOpened(registryID, identity, hermesID)
}

// FilterLogs executes a filter query.
func (mbc *MultichainBlockchainClient) FilterLogs(chainID int64, q ethereum.FilterQuery) ([]types.Log, error) {
	bc, err := mbc.getClientByChain(chainID)
	if err != nil {
		return nil, err
	}
	return bc.FilterLogs(q)
}

// HeaderByNumber returns a block header from the current canonical chain. If number is
// nil, the latest known header is returned.
func (mbc *MultichainBlockchainClient) HeaderByNumber(chainID int64, number *big.Int) (*types.Header, error) {
	bc, err := mbc.getClientByChain(chainID)
	if err != nil {
		return nil, err
	}
	return bc.HeaderByNumber(number)
}

func (mbc *MultichainBlockchainClient) SendTransaction(chainID int64, tx *types.Transaction) error {
	bc, err := mbc.getClientByChain(chainID)
	if err != nil {
		return err
	}
	return bc.SendTransaction(tx)
}

func (mbc *MultichainBlockchainClient) IsHermesActive(chainID int64, hermesID common.Address) (bool, error) {
	bc, err := mbc.getClientByChain(chainID)
	if err != nil {
		return false, err
	}
	return bc.IsHermesActive(hermesID)
}

func (mbc *MultichainBlockchainClient) PayAndSettle(chainID int64, psr PayAndSettleRequest) (*types.Transaction, error) {
	bc, err := mbc.getClientByChain(chainID)
	if err != nil {
		return nil, err
	}
	return bc.PayAndSettle(psr)
}
