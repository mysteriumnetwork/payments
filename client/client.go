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
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/mysteriumnetwork/payments/bindings"
	"github.com/mysteriumnetwork/payments/crypto"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// DefaultBackoff is the default backoff for the client
const DefaultBackoff = time.Second * 3

// ProviderChannel represents the provider channel
type ProviderChannel struct {
	Settled       *big.Int
	Stake         *big.Int
	LastUsedNonce *big.Int
	Timelock      *big.Int
}

type ethClientGetter interface {
	Client() *ethclient.Client
}

// Blockchain contains all the useful blockchain utilities for the payment off chain messaging
type Blockchain struct {
	ethClient ethClientGetter
	bcTimeout time.Duration
	nonceFunc nonceFunc
}

type nonceFunc func(ctx context.Context, account common.Address) (uint64, error)

// NewBlockchain returns a new instance of blockchain
func NewBlockchain(ethClient ethClientGetter, timeout time.Duration) *Blockchain {
	return &Blockchain{
		ethClient: ethClient,
		bcTimeout: timeout,
		nonceFunc: func(ctx context.Context, account common.Address) (uint64, error) {
			return ethClient.Client().PendingNonceAt(ctx, account)
		},
	}
}

// NewBlockchainWithCustomNonceTracker returns a new instance of blockchain with the provided nonce tracking func
func NewBlockchainWithCustomNonceTracker(ethClient ethClientGetter, timeout time.Duration, nonceFunc nonceFunc) *Blockchain {
	return &Blockchain{
		ethClient: ethClient,
		bcTimeout: timeout,
		nonceFunc: nonceFunc,
	}
}

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
	PayAndSettle(psr PayAndSettleRequest) (*types.Transaction, error)
	IsChannelOpened(registryID, identity, hermesID common.Address) (bool, error)
}

// GetHermesFee fetches the hermes fee from blockchain
func (bc *Blockchain) GetHermesFee(hermesAddress common.Address) (uint16, error) {
	caller, err := bindings.NewHermesImplementationCaller(hermesAddress, bc.ethClient.Client())
	if err != nil {
		return 0, errors.Wrap(err, "could not create hermes implementation caller")
	}

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	res, err := caller.LastFee(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return 0, errors.Wrap(err, "could not get hermes fee")
	}

	return res.Value, err
}

// CalculateHermesFee calls blockchain for calculation of hermes fee
func (bc *Blockchain) CalculateHermesFee(hermesAddress common.Address, value *big.Int) (*big.Int, error) {
	caller, err := bindings.NewHermesImplementationCaller(hermesAddress, bc.ethClient.Client())
	if err != nil {
		return nil, errors.Wrap(err, "could not create hermes implementation caller")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return caller.CalculateHermesFee(&bind.CallOpts{
		Context: ctx,
	}, value)
}

// IsRegisteredAsProvider checks if the provider is registered with the hermes properly
func (bc *Blockchain) IsRegisteredAsProvider(hermesAddress, registryAddress, addressToCheck common.Address) (bool, error) {
	registered, err := bc.IsRegistered(registryAddress, addressToCheck)
	if err != nil {
		return false, errors.Wrap(err, "could not check registration status")
	}

	if !registered {
		return false, nil
	}

	res, err := bc.getProviderChannelStake(hermesAddress, addressToCheck)
	if err != nil {
		return false, errors.Wrap(err, "could not get provider channel stake amount")
	}

	return res.Cmp(big.NewInt(0)) == 1, nil
}

// SubscribeToMystTokenTransfers subscribes to myst token transfers
func (bc *Blockchain) SubscribeToMystTokenTransfers(mystSCAddress common.Address) (chan *bindings.MystTokenTransfer, func(), error) {
	sink := make(chan *bindings.MystTokenTransfer)
	mtc, err := bindings.NewMystTokenFilterer(mystSCAddress, bc.ethClient.Client())
	if err != nil {
		return sink, nil, err
	}
	sub := event.Resubscribe(DefaultBackoff, func(ctx context.Context) (event.Subscription, error) {
		return mtc.WatchTransfer(&bind.WatchOpts{
			Context: ctx,
		}, sink, []common.Address{}, []common.Address{})
	})
	go func() {
		subErr := <-sub.Err()
		if subErr != nil {
			log.Error().Err(err).Msg("subscription error")
		}
		close(sink)
	}()

	return sink, sub.Unsubscribe, nil
}

// SubscribeToConsumerBalanceEvent subscribes to balance change events in blockchain
func (bc *Blockchain) SubscribeToConsumerBalanceEvent(channel, mystSCAddress common.Address, timeout time.Duration) (chan *bindings.MystTokenTransfer, func(), error) {
	sink := make(chan *bindings.MystTokenTransfer)
	mtc, err := bindings.NewMystTokenFilterer(mystSCAddress, bc.ethClient.Client())
	if err != nil {
		return sink, nil, err
	}

	sub, err := mtc.WatchTransfer(&bind.WatchOpts{}, sink, []common.Address{}, []common.Address{channel})
	if err != nil {
		return sink, nil, err
	}

	go func() {
		select {
		case <-time.After(timeout):
			sub.Unsubscribe()
		}
	}()

	go func() {
		subErr := <-sub.Err()
		if subErr != nil {
			log.Error().Err(err).Msg("subscription error")
		}
		close(sink)
	}()
	return sink, sub.Unsubscribe, nil
}

// GetProviderChannel returns the provider channel
func (bc *Blockchain) GetProviderChannel(hermesAddress common.Address, addressToCheck common.Address, pending bool) (ProviderChannel, error) {
	addressBytes, err := bc.getProviderChannelAddressBytes(hermesAddress, addressToCheck)
	if err != nil {
		return ProviderChannel{}, errors.Wrap(err, "could not calculate provider channel address")
	}
	caller, err := bindings.NewHermesImplementationCaller(hermesAddress, bc.ethClient.Client())
	if err != nil {
		return ProviderChannel{}, errors.Wrap(err, "could not create hermes caller")
	}

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	ch, err := caller.Channels(&bind.CallOpts{
		Pending: pending,
		Context: ctx,
	}, addressBytes)
	return ch, errors.Wrap(err, "could not get provider channel from bc")
}

func (bc *Blockchain) getProviderChannelStake(hermesAddress common.Address, addressToCheck common.Address) (*big.Int, error) {
	ch, err := bc.GetProviderChannel(hermesAddress, addressToCheck, false)
	return ch.Stake, errors.Wrap(err, "could not get provider channel from bc")
}

func (bc *Blockchain) TransactionReceipt(hash common.Hash) (*types.Receipt, error) {
	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()
	return bc.ethClient.Client().TransactionReceipt(ctx, hash)
}

func (bc *Blockchain) getProviderChannelAddressBytes(hermesAddress, addressToCheck common.Address) ([32]byte, error) {
	addressBytes := [32]byte{}

	addr, err := crypto.GenerateProviderChannelID(addressToCheck.Hex(), hermesAddress.Hex())
	if err != nil {
		return addressBytes, errors.Wrap(err, "could not generate channel address")
	}

	copy(addressBytes[:], crypto.Pad(common.Hex2Bytes(strings.TrimPrefix(addr, "0x")), 32))

	return addressBytes, nil
}

// SubscribeToPromiseSettledEvent subscribes to promise settled events
func (bc *Blockchain) SubscribeToPromiseSettledEvent(providerID, hermesID common.Address) (sink chan *bindings.HermesImplementationPromiseSettled, cancel func(), err error) {
	addr, err := bc.getProviderChannelAddressBytes(hermesID, providerID)
	if err != nil {
		return sink, cancel, errors.Wrap(err, "could not get provider channel address")
	}
	return bc.SubscribeToPromiseSettledEventByChannelID(hermesID, [][32]byte{addr})
}

// IsRegistered checks wether the given identity is registered or not
func (bc *Blockchain) IsRegistered(registryAddress, addressToCheck common.Address) (bool, error) {
	caller, err := bindings.NewRegistryCaller(registryAddress, bc.ethClient.Client())
	if err != nil {
		return false, errors.Wrap(err, "could not create registry caller")
	}

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	res, err := caller.IsRegistered(&bind.CallOpts{
		Context: ctx,
	}, addressToCheck)
	return res, errors.Wrap(err, "could not check registration status")
}

// GetMystBalance returns myst balance
func (bc *Blockchain) GetMystBalance(mystAddress, identity common.Address) (*big.Int, error) {
	c, err := bindings.NewMystTokenCaller(mystAddress, bc.ethClient.Client())
	if err != nil {
		return nil, err
	}
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, bc.bcTimeout)
	defer cancel()
	return c.BalanceOf(&bind.CallOpts{
		Context: ctx,
	}, identity)
}

// RegistrationRequest contains all the parameters for the registration request
type RegistrationRequest struct {
	WriteRequest
	ChainID         int64
	HermesID        common.Address
	Stake           *big.Int
	TransactorFee   *big.Int
	Beneficiary     common.Address
	Signature       []byte
	RegistryAddress common.Address
	Nonce           *big.Int
}

func (r RegistrationRequest) toEstimator(ethClient ethClientGetter) (*bindings.ContractEstimator, error) {
	return bindings.NewContractEstimator(r.RegistryAddress, bindings.RegistryABI, ethClient.Client())
}

func (r RegistrationRequest) toEstimateOps() *bindings.EstimateOpts {
	return &bindings.EstimateOpts{
		From:   r.Identity,
		Method: "registerIdentity",
		Params: []interface{}{r.HermesID, r.Stake, r.TransactorFee, r.Beneficiary, r.Signature},
	}
}

// WriteRequest contains the required params for a write request
type WriteRequest struct {
	Identity common.Address
	Signer   bind.SignerFn
	GasLimit uint64
	GasPrice *big.Int
	Nonce    *big.Int
}

// getGasLimit returns the gas limit
func (wr WriteRequest) getGasLimit() uint64 {
	return wr.GasLimit
}

// RegisterIdentity registers the given identity on blockchain
func (bc *Blockchain) RegisterIdentity(rr RegistrationRequest) (*types.Transaction, error) {
	transactor, err := bindings.NewRegistryTransactor(rr.RegistryAddress, bc.ethClient.Client())
	if err != nil {
		return nil, err
	}
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, bc.bcTimeout)
	defer cancel()

	nonce := rr.Nonce
	if nonce == nil {
		nonceUint, err := bc.getNonce(rr.Identity)
		if err != nil {
			return nil, errors.Wrap(err, "could not get nonce")
		}
		nonce = big.NewInt(0).SetUint64(nonceUint)
	}
	tx, err := transactor.RegisterIdentity(&bind.TransactOpts{
		From:     rr.Identity,
		Signer:   rr.Signer,
		Context:  ctx,
		GasLimit: rr.GasLimit,
		GasPrice: rr.GasPrice,
		Nonce:    nonce,
	},
		rr.HermesID,
		rr.Stake,
		rr.TransactorFee,
		rr.Beneficiary,
		rr.Signature,
	)
	return tx, err
}

// PayAndSettleRequest allows to pay and settle and exit to l1 via this.
type PayAndSettleRequest struct {
	WriteRequest
	Beneficiary          common.Address
	HermesID             common.Address
	ProviderID           common.Address
	Promise              crypto.Promise
	BeneficiarySignature []byte
}

func (psr PayAndSettleRequest) toEstimator(ethClient ethClientGetter) (*bindings.ContractEstimator, error) {
	return bindings.NewContractEstimator(psr.HermesID, bindings.HermesImplementationABI, ethClient.Client())
}

func (psr PayAndSettleRequest) toEstimateOps() *bindings.EstimateOpts {
	return &bindings.EstimateOpts{
		From:   psr.Identity,
		Method: "payAndSettle",
		Params: []interface{}{psr.ProviderID, psr.Promise.Amount, psr.Promise.Fee, ToBytes32(psr.Promise.R), psr.Promise.Signature, psr.Beneficiary, psr.BeneficiarySignature},
	}
}

// PayAndSettle registers the given identity on blockchain.
func (bc *Blockchain) PayAndSettle(psr PayAndSettleRequest) (*types.Transaction, error) {
	transactor, err := bindings.NewHermesImplementationTransactor(psr.HermesID, bc.ethClient.Client())
	if err != nil {
		return nil, err
	}
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, bc.bcTimeout)
	defer cancel()

	nonce := psr.Nonce
	if nonce == nil {
		nonceUint, err := bc.getNonce(psr.Identity)
		if err != nil {
			return nil, errors.Wrap(err, "could not get nonce")
		}
		nonce = big.NewInt(0).SetUint64(nonceUint)
	}
	tx, err := transactor.PayAndSettle(&bind.TransactOpts{
		From:     psr.Identity,
		Signer:   psr.Signer,
		Context:  ctx,
		GasLimit: psr.GasLimit,
		GasPrice: psr.GasPrice,
		Nonce:    nonce,
	},
		psr.ProviderID,
		psr.Promise.Amount,
		psr.Promise.Fee,
		ToBytes32(psr.Promise.R),
		psr.Promise.Signature,
		psr.Beneficiary,
		psr.BeneficiarySignature,
	)
	return tx, err
}

// TransferRequest contains all the parameters for a transfer request
type TransferRequest struct {
	MystAddress common.Address
	Recipient   common.Address
	Amount      *big.Int
	WriteRequest
}

func (r TransferRequest) toEstimator(ethClient ethClientGetter) (*bindings.ContractEstimator, error) {
	return bindings.NewContractEstimator(r.MystAddress, bindings.MystTokenABI, ethClient.Client())
}

func (r TransferRequest) toEstimateOps() *bindings.EstimateOpts {
	return &bindings.EstimateOpts{
		From:   r.Identity,
		Method: "transfer",
		Params: []interface{}{r.Recipient, r.Amount},
	}
}

// TransferMyst transfers myst
func (bc *Blockchain) TransferMyst(req TransferRequest) (tx *types.Transaction, err error) {
	transactor, err := bindings.NewMystTokenTransactor(req.MystAddress, bc.ethClient.Client())
	if err != nil {
		return tx, err
	}

	nonce, err := bc.getNonce(req.Identity)
	if err != nil {
		return nil, errors.Wrap(err, "could not get nonce")
	}

	return transactor.Transfer(&bind.TransactOpts{
		From:     req.Identity,
		Signer:   req.Signer,
		GasPrice: req.GasPrice,
		GasLimit: req.GasLimit,
		Nonce:    big.NewInt(0).SetUint64(nonce),
	}, req.Recipient, req.Amount)
}

// IsHermesRegistered checks if given hermes is registered and returns true or false.
func (bc *Blockchain) IsHermesRegistered(registryAddress, acccountantID common.Address) (bool, error) {
	caller, err := bindings.NewRegistryCaller(registryAddress, bc.ethClient.Client())
	if err != nil {
		return false, err
	}
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, bc.bcTimeout)
	defer cancel()
	return caller.IsHermes(&bind.CallOpts{
		Pending: false,
		Context: ctx,
	}, acccountantID)
}

// ProviderStakeIncreaseRequest represents all the parameters required for stake increase.
type ProviderStakeIncreaseRequest struct {
	WriteRequest
	ChannelID [32]byte
	HermesID  common.Address
	Amount    *big.Int
}

func (r ProviderStakeIncreaseRequest) toEstimator(ethClient ethClientGetter) (*bindings.ContractEstimator, error) {
	return bindings.NewContractEstimator(r.HermesID, bindings.HermesImplementationABI, ethClient.Client())
}

func (r ProviderStakeIncreaseRequest) toEstimateOps() *bindings.EstimateOpts {
	return &bindings.EstimateOpts{
		From:   r.Identity,
		Method: "increaseStake",
		Params: []interface{}{r.ChannelID, r.Amount},
	}
}

// IncreaseProviderStake increases the provider stake.
func (bc *Blockchain) IncreaseProviderStake(req ProviderStakeIncreaseRequest) (*types.Transaction, error) {
	t, err := bindings.NewHermesImplementationTransactor(req.HermesID, bc.ethClient.Client())
	if err != nil {
		return nil, err
	}

	transactor, cancel, err := bc.getTransactorFromRequest(req.WriteRequest)
	defer cancel()
	if err != nil {
		return nil, fmt.Errorf("could not get transactor: %w", err)
	}

	return t.IncreaseStake(transactor, req.ChannelID, req.Amount)
}

// SettleIntoStakeRequest represents all the parameters required for settling into stake.
type SettleIntoStakeRequest struct {
	WriteRequest
	Promise    crypto.Promise
	HermesID   common.Address
	ProviderID common.Address
}

func (r SettleIntoStakeRequest) toEstimator(ethClient ethClientGetter) (*bindings.ContractEstimator, error) {
	return bindings.NewContractEstimator(r.HermesID, bindings.HermesImplementationABI, ethClient.Client())
}

func (r SettleIntoStakeRequest) toEstimateOps() *bindings.EstimateOpts {
	return &bindings.EstimateOpts{
		From:   r.Identity,
		Method: "settleIntoStake",
		Params: []interface{}{
			r.ProviderID,
			r.Promise.Amount,
			r.Promise.Fee,
			ToBytes32(r.Promise.R),
			r.Promise.Signature,
		},
	}
}

// SettleIntoStake settles the hermes promise into stake increase.
func (bc *Blockchain) SettleIntoStake(req SettleIntoStakeRequest) (*types.Transaction, error) {
	t, err := bindings.NewHermesImplementationTransactor(req.HermesID, bc.ethClient.Client())
	if err != nil {
		return nil, err
	}

	transactor, cancel, err := bc.getTransactorFromRequest(req.WriteRequest)
	defer cancel()
	if err != nil {
		return nil, fmt.Errorf("could not get transactor: %w", err)
	}

	amount := req.Promise.Amount
	fee := req.Promise.Fee
	lock := [32]byte{}
	copy(lock[:], req.Promise.R)

	channelID := [32]byte{}
	copy(channelID[:], req.Promise.ChannelID)

	return t.SettleIntoStake(transactor, req.ProviderID, amount, fee, lock, req.Promise.Signature)
}

// DecreaseProviderStakeRequest represents all the parameters required for decreasing provider stake.
type DecreaseProviderStakeRequest struct {
	WriteRequest
	Request    crypto.DecreaseProviderStakeRequest
	ProviderID common.Address
}

func (r DecreaseProviderStakeRequest) toEstimator(ethClient ethClientGetter) (*bindings.ContractEstimator, error) {
	return bindings.NewContractEstimator(r.Request.HermesID, bindings.HermesImplementationABI, ethClient.Client())
}

func (r DecreaseProviderStakeRequest) toEstimateOps() *bindings.EstimateOpts {
	return &bindings.EstimateOpts{
		From:   r.Identity,
		Method: "decreaseStake",
		Params: []interface{}{r.ProviderID, r.Request.Amount, r.Request.TransactorFee, r.Request.Signature},
	}
}

func (bc *Blockchain) GetBeneficiary(registryAddress, identity common.Address) (common.Address, error) {
	caller, err := bindings.NewRegistryCaller(registryAddress, bc.ethClient.Client())
	if err != nil {
		return common.Address{}, err
	}

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, bc.bcTimeout)
	defer cancel()
	return caller.GetBeneficiary(&bind.CallOpts{
		Pending: false,
		Context: ctx,
	}, identity)
}

// DecreaseProviderStake decreases provider stake.
func (bc *Blockchain) DecreaseProviderStake(req DecreaseProviderStakeRequest) (*types.Transaction, error) {
	t, err := bindings.NewHermesImplementationTransactor(req.Request.HermesID, bc.ethClient.Client())
	if err != nil {
		return nil, err
	}

	transactor, cancel, err := bc.getTransactorFromRequest(req.WriteRequest)
	defer cancel()
	if err != nil {
		return nil, fmt.Errorf("could not get transactor: %w", err)
	}

	return t.DecreaseStake(transactor, req.ProviderID, req.Request.Amount, req.Request.TransactorFee, req.Request.Signature)
}

func (bc *Blockchain) getTransactorFromRequest(req WriteRequest) (*bind.TransactOpts, func(), error) {
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, bc.bcTimeout)

	if req.Nonce == nil {
		nonce, err := bc.getNonce(req.Identity)
		if err != nil {
			return nil, cancel, errors.Wrap(err, "could not get nonce")
		}
		req.Nonce = big.NewInt(0).SetUint64(nonce)
	}

	return &bind.TransactOpts{
		From:     req.Identity,
		Signer:   req.Signer,
		Context:  ctx,
		GasLimit: req.GasLimit,
		GasPrice: req.GasPrice,
		Nonce:    req.Nonce,
	}, cancel, nil
}

// GetHermesOperator returns operator address of given hermes
func (bc *Blockchain) GetHermesOperator(hermesID common.Address) (common.Address, error) {
	caller, err := bindings.NewHermesImplementationCaller(hermesID, bc.ethClient.Client())
	if err != nil {
		return common.Address{}, err
	}

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, bc.bcTimeout)
	defer cancel()
	return caller.GetOperator(&bind.CallOpts{
		Pending: false,
		Context: ctx,
	})
}

// IsHermesActive determines if hermes is active or not.
func (bc *Blockchain) IsHermesActive(hermesID common.Address) (bool, error) {
	caller, err := bindings.NewHermesImplementationCaller(hermesID, bc.ethClient.Client())
	if err != nil {
		return false, err
	}

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, bc.bcTimeout)
	defer cancel()
	return caller.IsHermesActive(&bind.CallOpts{
		Pending: false,
		Context: ctx,
	})
}

// SettleAndRebalanceRequest represents all the parameters required for settle and rebalance
type SettleAndRebalanceRequest struct {
	WriteRequest
	HermesID   common.Address
	ProviderID common.Address
	Promise    crypto.Promise
}

func (r SettleAndRebalanceRequest) toEstimator(ethClient ethClientGetter) (*bindings.ContractEstimator, error) {
	return bindings.NewContractEstimator(r.HermesID, bindings.HermesImplementationABI, ethClient.Client())
}

func (r SettleAndRebalanceRequest) toEstimateOps() *bindings.EstimateOpts {
	return &bindings.EstimateOpts{
		From:   r.Identity,
		Method: "settlePromise",
		Params: []interface{}{r.ProviderID, r.Promise.Amount, r.Promise.Fee, ToBytes32(r.Promise.R), r.Promise.Signature},
	}
}

// SettleAndRebalance is settling given hermes issued promise
func (bc *Blockchain) SettleAndRebalance(req SettleAndRebalanceRequest) (*types.Transaction, error) {
	transactor, err := bindings.NewHermesImplementationTransactor(req.HermesID, bc.ethClient.Client())
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	nonce, err := bc.getNonce(req.Identity)
	if err != nil {
		return nil, errors.Wrap(err, "could not get nonce")
	}

	return transactor.SettlePromise(&bind.TransactOpts{
		From:     req.Identity,
		Signer:   req.Signer,
		Context:  ctx,
		GasLimit: req.GasLimit,
		GasPrice: req.GasPrice,
		Nonce:    big.NewInt(0).SetUint64(nonce),
	},
		req.ProviderID,
		req.Promise.Amount,
		req.Promise.Fee,
		ToBytes32(req.Promise.R),
		req.Promise.Signature,
	)
}

func ToBytes32(arr []byte) (res [32]byte) {
	copy(res[:], arr)
	return res
}

// GetLastRegistryNonce returns the last registry nonce.
func (bc *Blockchain) GetLastRegistryNonce(registry common.Address) (*big.Int, error) {
	caller, err := bindings.NewRegistryCaller(registry, bc.ethClient.Client())
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	return caller.LastNonce(&bind.CallOpts{
		Pending: false,
		Context: ctx,
	})
}

// GetProviderChannelByID returns the given provider channel information
func (bc *Blockchain) GetProviderChannelByID(acc common.Address, chID []byte) (ProviderChannel, error) {
	caller, err := bindings.NewHermesImplementationCaller(acc, bc.ethClient.Client())
	if err != nil {
		return ProviderChannel{}, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	return caller.Channels(&bind.CallOpts{
		Pending: false,
		Context: ctx,
	}, ToBytes32(chID))
}

// ConsumersHermes represents the consumers hermes
type ConsumersHermes struct {
	Operator        common.Address
	ContractAddress common.Address
	Settled         *big.Int
}

// GetConsumerChannelsHermes returns the consumer channels hermes
func (bc *Blockchain) GetConsumerChannelsHermes(channelAddress common.Address) (ConsumersHermes, error) {
	c, err := bindings.NewChannelImplementationCaller(channelAddress, bc.ethClient.Client())
	if err != nil {
		return ConsumersHermes{}, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	return c.Hermes(&bind.CallOpts{Context: ctx})
}

// GetConsumerChannelOperator returns the consumer channel operator/identity
func (bc *Blockchain) GetConsumerChannelOperator(channelAddress common.Address) (common.Address, error) {
	c, err := bindings.NewChannelImplementationCaller(channelAddress, bc.ethClient.Client())
	if err != nil {
		return common.Address{}, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	return c.Operator(&bind.CallOpts{Context: ctx})
}

// SubscribeToIdentityRegistrationEvents subscribes to identity registration events
func (bc *Blockchain) SubscribeToIdentityRegistrationEvents(registryAddress common.Address) (sink chan *bindings.RegistryRegisteredIdentity, cancel func(), err error) {
	filterer, err := bindings.NewRegistryFilterer(registryAddress, bc.ethClient.Client())
	if err != nil {
		return sink, cancel, errors.Wrap(err, "could not create registry filterer")
	}
	sink = make(chan *bindings.RegistryRegisteredIdentity)
	sub := event.Resubscribe(DefaultBackoff, func(ctx context.Context) (event.Subscription, error) {
		return filterer.WatchRegisteredIdentity(&bind.WatchOpts{
			Context: ctx,
		}, sink, nil)
	})
	go func() {
		subErr := <-sub.Err()
		if subErr != nil {
			log.Error().Err(err).Msg("subscription error")
		}
		close(sink)
	}()
	return sink, sub.Unsubscribe, nil
}

// SubscribeToConsumerChannelBalanceUpdate subscribes to consumer channel balance update events
func (bc *Blockchain) SubscribeToConsumerChannelBalanceUpdate(mystSCAddress common.Address, channelAddresses []common.Address) (sink chan *bindings.MystTokenTransfer, cancel func(), err error) {
	filterer, err := bindings.NewMystTokenFilterer(mystSCAddress, bc.ethClient.Client())
	if err != nil {
		return sink, cancel, errors.Wrap(err, "could not create myst token filterer")
	}

	sink = make(chan *bindings.MystTokenTransfer)
	sub := event.Resubscribe(DefaultBackoff, func(ctx context.Context) (event.Subscription, error) {
		return filterer.WatchTransfer(&bind.WatchOpts{
			Context: ctx,
		}, sink, nil, channelAddresses)
	})
	go func() {
		subErr := <-sub.Err()
		if subErr != nil {
			log.Error().Err(err).Msg("subscription error")
		}
		close(sink)
	}()
	return sink, sub.Unsubscribe, nil
}

// SettleRequest represents all the parameters required for settle
type SettleRequest struct {
	WriteRequest
	ChannelID common.Address
	Promise   crypto.Promise
}

func (r SettleRequest) toEstimator(ethClient ethClientGetter) (*bindings.ContractEstimator, error) {
	return bindings.NewContractEstimator(r.ChannelID, bindings.ChannelImplementationABI, ethClient.Client())
}

func (r SettleRequest) toEstimateOps() *bindings.EstimateOpts {
	lock := [32]byte{}
	copy(lock[:], r.Promise.R)

	return &bindings.EstimateOpts{
		From:   r.Identity,
		Method: "settlePromise",
		Params: []interface{}{r.Promise.Amount, r.Promise.Fee, lock, r.Promise.Signature},
	}
}

// SettlePromise is settling the given consumer issued promise
func (bc *Blockchain) SettlePromise(req SettleRequest) (*types.Transaction, error) {
	transactor, err := bindings.NewChannelImplementationTransactor(req.ChannelID, bc.ethClient.Client())
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	amount := req.Promise.Amount
	fee := req.Promise.Fee
	lock := [32]byte{}
	copy(lock[:], req.Promise.R)

	nonce, err := bc.getNonce(req.Identity)
	if err != nil {
		return nil, errors.Wrap(err, "could not get nonce")
	}

	return transactor.SettlePromise(&bind.TransactOpts{
		From:     req.Identity,
		Signer:   req.Signer,
		Context:  ctx,
		GasLimit: req.GasLimit,
		GasPrice: req.GasPrice,
		Nonce:    big.NewInt(0).SetUint64(nonce),
	},
		amount, fee, lock, req.Promise.Signature,
	)
}

func (bc *Blockchain) getNonce(identity common.Address) (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()
	return bc.nonceFunc(ctx, identity)
}

// GetHermesURL gets the hermes url from BC.
func (bc *Blockchain) GetHermesURL(registryID, hermesID common.Address) (string, error) {
	caller, err := bindings.NewRegistryCaller(registryID, bc.ethClient.Client())
	if err != nil {
		return "", fmt.Errorf("could not create new registry caller %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	urlBytes, err := caller.GetHermesURL(
		&bind.CallOpts{
			Context: ctx,
		},
		hermesID,
	)
	return string(urlBytes), err
}

// Hermes represents the hermes in registry.
type Hermes struct {
	Operator common.Address
	ImplVer  *big.Int
	Stake    *big.Int
	URL      string
}

// GetHermes returns hermes info from registry
func (bc *Blockchain) GetHermes(registryID, hermesID common.Address) (Hermes, error) {
	caller, err := bindings.NewRegistryCaller(registryID, bc.ethClient.Client())
	if err != nil {
		return Hermes{}, fmt.Errorf("could not create new registry caller %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	status, err := caller.GetHermes(&bind.CallOpts{
		Context: ctx,
	}, hermesID)
	if err != nil {
		return Hermes{}, err
	}

	return Hermes{
		Operator: status.Operator,
		ImplVer:  status.ImplVer,
		Stake:    big.NewInt(0).SetBytes(status.Stake[:]),
		URL:      string(status.Url),
	}, nil
}

// GetChannelImplementationByVersion returns the channel implementation for the specified version.
func (bc *Blockchain) GetChannelImplementationByVersion(registryID common.Address, version *big.Int) (common.Address, error) {
	caller, err := bindings.NewRegistryCaller(registryID, bc.ethClient.Client())
	if err != nil {
		return common.Address{}, fmt.Errorf("could not create new registry caller %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	return caller.GetChannelImplementation(&bind.CallOpts{
		Context: ctx,
	}, version)
}

func (bc *Blockchain) IsChannelOpened(registryID, identity, hermesID common.Address) (bool, error) {
	caller, err := bindings.NewRegistryCaller(registryID, bc.ethClient.Client())
	if err != nil {
		return false, fmt.Errorf("could not create new registry caller %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	return caller.IsChannelOpened(&bind.CallOpts{
		Context: ctx,
	}, identity, hermesID)
}

// SubscribeToPromiseSettledEventByChannelID subscribes to promise settled events
func (bc *Blockchain) SubscribeToPromiseSettledEventByChannelID(hermesID common.Address, providerAddresses [][32]byte) (sink chan *bindings.HermesImplementationPromiseSettled, cancel func(), err error) {
	caller, err := bindings.NewHermesImplementationFilterer(hermesID, bc.ethClient.Client())
	if err != nil {
		return sink, cancel, errors.Wrap(err, "could not create hermes caller")
	}
	sink = make(chan *bindings.HermesImplementationPromiseSettled)

	sub := event.Resubscribe(DefaultBackoff, func(ctx context.Context) (event.Subscription, error) {
		return caller.WatchPromiseSettled(&bind.WatchOpts{
			Context: ctx,
		}, sink, providerAddresses, []common.Address{})
	})
	go func() {
		subErr := <-sub.Err()
		if subErr != nil {
			log.Error().Err(err).Msg("subscription error")
		}
		cancel()
		close(sink)
	}()

	return sink, sub.Unsubscribe, nil
}

// GetEthBalance gets the current ethereum balance for the address.
func (bc *Blockchain) GetEthBalance(address common.Address) (*big.Int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()
	return bc.ethClient.Client().BalanceAt(ctx, address, nil)
}

// EthTransferRequest represents the ethereum transfer request input parameters.
type EthTransferRequest struct {
	WriteRequest
	To     common.Address
	Amount *big.Int
}

// TransferEth transfers ethereum to the given address.
func (bc *Blockchain) TransferEth(etr EthTransferRequest) (*types.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	id, err := bc.NetworkID()
	if err != nil {
		return nil, fmt.Errorf("could not get network id: %w", err)
	}

	nonceUint, err := bc.getNonce(etr.Identity)
	if err != nil {
		return nil, fmt.Errorf("could not get nonce: %w", err)
	}

	tx := types.NewTransaction(nonceUint, etr.To, etr.Amount, etr.GasLimit, etr.GasPrice, nil)
	signedTx, err := etr.Signer(types.NewEIP155Signer(id), etr.Identity, tx)
	if err != nil {
		return nil, fmt.Errorf("could not sign tx: %w", err)
	}

	err = bc.ethClient.Client().SendTransaction(ctx, signedTx)
	if err != nil {
		return nil, fmt.Errorf("could not send transaction: %w", err)
	}

	return signedTx, err
}

// FilterLogs executes a filter query.
func (bc *Blockchain) FilterLogs(q ethereum.FilterQuery) ([]types.Log, error) {
	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()
	return bc.ethClient.Client().FilterLogs(ctx, q)
}

// HeaderByNumber returns a block header from the current canonical chain. If number is
// nil, the latest known header is returned.
func (bc *Blockchain) HeaderByNumber(number *big.Int) (*types.Header, error) {
	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()
	return bc.ethClient.Client().HeaderByNumber(ctx, number)
}

func (bc *Blockchain) SuggestGasPrice() (*big.Int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()
	return bc.ethClient.Client().SuggestGasPrice(ctx)
}

// NetworkID returns the network id
func (bc *Blockchain) NetworkID() (*big.Int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()
	return bc.ethClient.Client().NetworkID(ctx)
}

// ConsumerChannel represents the consumer channel
type ConsumerChannel struct {
	Settled *big.Int
	Balance *big.Int
}

// GetConsumerChannel returns the consumer channel
func (bc *Blockchain) GetConsumerChannel(addr common.Address, mystSCAddress common.Address) (ConsumerChannel, error) {
	ad := common.BytesToAddress(addr.Bytes())
	party, err := bc.GetConsumerChannelsHermes(ad)
	if err != nil {
		return ConsumerChannel{}, err
	}

	balance, err := bc.GetMystBalance(mystSCAddress, ad)
	return ConsumerChannel{
		Settled: party.Settled,
		Balance: balance,
	}, err
}

// SettleWithBeneficiaryRequest represents all the parameters required for setting new beneficiary via transactor.
type SettleWithBeneficiaryRequest struct {
	WriteRequest
	Promise     crypto.Promise
	HermesID    common.Address
	ProviderID  common.Address
	Beneficiary common.Address
	Signature   []byte
}

func (r SettleWithBeneficiaryRequest) toEstimator(ethClient ethClientGetter) (*bindings.ContractEstimator, error) {
	return bindings.NewContractEstimator(r.HermesID, bindings.HermesImplementationABI, ethClient.Client())
}

func (r SettleWithBeneficiaryRequest) toEstimateOps() *bindings.EstimateOpts {
	return &bindings.EstimateOpts{
		From:   r.Identity,
		Method: "settleWithBeneficiary",
		Params: []interface{}{
			r.ProviderID,
			r.Promise.Amount,
			r.Promise.Fee,
			ToBytes32(r.Promise.R),
			r.Promise.Signature,
			r.Beneficiary,
			r.Signature,
		},
	}
}

// GetHermessAvailableBalance returns the balance that is available for hermes.
func (bc *Blockchain) GetHermessAvailableBalance(hermesAddress common.Address) (*big.Int, error) {
	caller, err := bindings.NewHermesImplementationCaller(hermesAddress, bc.ethClient.Client())
	if err != nil {
		return nil, errors.Wrap(err, "could not create hermes implementation caller")
	}

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	res, err := caller.AvailableBalance(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return nil, errors.Wrap(err, "could not get hermes available balance")
	}

	return res, nil
}

// SettleWithBeneficiary sets new beneficiary for the provided identity and settles lastest promise into new beneficiary address.
func (bc *Blockchain) SettleWithBeneficiary(req SettleWithBeneficiaryRequest) (*types.Transaction, error) {
	transactor, err := bindings.NewHermesImplementationTransactor(req.HermesID, bc.ethClient.Client())
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	nonce, err := bc.getNonce(req.Identity)
	if err != nil {
		return nil, errors.Wrap(err, "could not get nonce")
	}

	return transactor.SettleWithBeneficiary(&bind.TransactOpts{
		From:     req.Identity,
		Signer:   req.Signer,
		Context:  ctx,
		GasLimit: req.GasLimit,
		GasPrice: req.GasPrice,
		Nonce:    big.NewInt(0).SetUint64(nonce),
	},
		req.ProviderID,
		req.Promise.Amount,
		req.Promise.Fee,
		ToBytes32(req.Promise.R),
		req.Promise.Signature,
		req.Beneficiary,
		req.Signature,
	)
}

// GetStakeThresholds returns the stake tresholds for the given hermes.
func (bc *Blockchain) GetStakeThresholds(hermesID common.Address) (min, max *big.Int, err error) {
	caller, err := bindings.NewHermesImplementationCaller(hermesID, bc.ethClient.Client())
	if err != nil {
		return nil, nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	return caller.GetStakeThresholds(&bind.CallOpts{
		Context: ctx,
	})
}

// SendTransaction sends a transaction to the blockchain.
func (bc *Blockchain) SendTransaction(tx *types.Transaction) error {
	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	return bc.ethClient.Client().SendTransaction(ctx, tx)
}
