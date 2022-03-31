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
	"github.com/ethereum/go-ethereum/event"
	"github.com/mysteriumnetwork/payments/bindings"
	"github.com/mysteriumnetwork/payments/bindings/rewarder"
	"github.com/mysteriumnetwork/payments/bindings/topperupper"
	"github.com/mysteriumnetwork/payments/bindings/uniswapv3"
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

// Blockchain contains all the useful blockchain utilities for the payment off chain messaging
type Blockchain struct {
	ethClient EthClientGetter
	bcTimeout time.Duration
	nonceFunc nonceFunc
}

type nonceFunc func(ctx context.Context, account common.Address) (uint64, error)

// NewBlockchain returns a new instance of blockchain
func NewBlockchain(ethClient EthClientGetter, timeout time.Duration) *Blockchain {
	return &Blockchain{
		ethClient: ethClient,
		bcTimeout: timeout,
		nonceFunc: func(ctx context.Context, account common.Address) (uint64, error) {
			return ethClient.Client().PendingNonceAt(ctx, account)
		},
	}
}

// NewBlockchainWithCustomNonceTracker returns a new instance of blockchain with the provided nonce tracking func
func NewBlockchainWithCustomNonceTracker(ethClient EthClientGetter, timeout time.Duration, nonceFunc nonceFunc) *Blockchain {
	return &Blockchain{
		ethClient: ethClient,
		bcTimeout: timeout,
		nonceFunc: nonceFunc,
	}
}

func (bc *Blockchain) makeTransactOpts(ctx context.Context, rr *WriteRequest) (*bind.TransactOpts, error) {
	if rr.Nonce == nil {
		nonceUint, err := bc.getNonce(rr.Identity)
		if err != nil {
			return nil, errors.Wrap(err, "could not get nonce")
		}
		rr.Nonce = big.NewInt(0).SetUint64(nonceUint)
	}

	return rr.toTransactOpts(ctx), nil
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

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
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

// GetProviderChannel returns the provider channel
func (bc *Blockchain) GetProvidersWithdrawalChannel(hermesAddress common.Address, addressToCheck common.Address, pending bool) (ProviderChannel, error) {
	addressBytes, err := bc.getProviderChannelAddressForWithdrawalBytes(hermesAddress, addressToCheck)
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

func (bc *Blockchain) TransactionByHash(hash common.Hash) (*types.Transaction, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()
	return bc.ethClient.Client().TransactionByHash(ctx, hash)
}

func (bc *Blockchain) TransactionReceipt(hash common.Hash) (*types.Receipt, error) {
	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()
	return bc.ethClient.Client().TransactionReceipt(ctx, hash)
}

func (bc *Blockchain) PendingNonceAt(account common.Address) (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()
	return bc.ethClient.Client().PendingNonceAt(ctx, account)
}

func (bc *Blockchain) NonceAt(account common.Address, blockNum *big.Int) (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()
	return bc.ethClient.Client().NonceAt(ctx, account, blockNum)
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

func (bc *Blockchain) getProviderChannelAddressForWithdrawalBytes(hermesAddress, addressToCheck common.Address) ([32]byte, error) {
	addressBytes := [32]byte{}

	addr, err := crypto.GenerateProviderChannelIDForPayAndSettle(addressToCheck.Hex(), hermesAddress.Hex())
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

// SubscribeToWithdrawalPromiseSettledEvent subscribes to withdrawal promise settled events
func (bc *Blockchain) SubscribeToWithdrawalPromiseSettledEvent(providerID, hermesID common.Address) (sink chan *bindings.HermesImplementationPromiseSettled, cancel func(), err error) {
	addr, err := bc.getProviderChannelAddressForWithdrawalBytes(hermesID, providerID)
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

func (r RegistrationRequest) toEstimator(ethClient EthClientGetter) (*bindings.ContractEstimator, error) {
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
	Nonce    *big.Int
	Identity common.Address
	Signer   bind.SignerFn

	// GasTip is the amount of tokens to tip for miners. Required for EIP-1559.
	GasTip *big.Int
	// BaseFee is the amount which has and will be paid for the transaction. Required for EIP-1559.
	BaseFee  *big.Int
	GasLimit uint64

	// GasPrice is the legacy gas price pre london hardfork.
	GasPrice *big.Int
}

func (wr *WriteRequest) toTransactOpts(ctx context.Context) *bind.TransactOpts {
	if wr.GasPrice != nil && wr.GasTip != nil {
		panic("could not convert write request to transact opts: can't set both GasTip and GasPrice")
	}

	to := &bind.TransactOpts{
		Context:  ctx,
		From:     wr.Identity,
		Signer:   wr.Signer,
		GasLimit: wr.GasLimit,
		Nonce:    wr.Nonce,
	}

	// Support pre EIP-1559 transactions
	if wr.GasPrice != nil && wr.GasPrice.Cmp(big.NewInt(0)) > 0 {
		to.GasPrice = wr.GasPrice
		return to
	}

	if wr.GasTip != nil && wr.BaseFee != nil {
		to.GasFeeCap = new(big.Int).Add(wr.GasTip, wr.BaseFee)
	}
	to.GasTipCap = wr.GasTip

	return to
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

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	to, err := bc.makeTransactOpts(ctx, &rr.WriteRequest)
	if err != nil {
		return nil, err
	}

	tx, err := transactor.RegisterIdentity(
		to,
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

func (psr PayAndSettleRequest) toEstimator(ethClient EthClientGetter) (*bindings.ContractEstimator, error) {
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

	to, err := bc.makeTransactOpts(ctx, &psr.WriteRequest)
	if err != nil {
		return nil, err
	}

	tx, err := transactor.PayAndSettle(
		to,
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

func (r TransferRequest) toEstimator(ethClient EthClientGetter) (*bindings.ContractEstimator, error) {
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

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	to, err := bc.makeTransactOpts(ctx, &req.WriteRequest)
	if err != nil {
		return nil, err
	}

	return transactor.Transfer(to, req.Recipient, req.Amount)
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

func (r ProviderStakeIncreaseRequest) toEstimator(ethClient EthClientGetter) (*bindings.ContractEstimator, error) {
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

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	to, err := bc.makeTransactOpts(ctx, &req.WriteRequest)
	if err != nil {
		return nil, err
	}

	return t.IncreaseStake(to, req.ChannelID, req.Amount)
}

// SettleIntoStakeRequest represents all the parameters required for settling into stake.
type SettleIntoStakeRequest struct {
	WriteRequest
	Promise    crypto.Promise
	HermesID   common.Address
	ProviderID common.Address
}

func (r SettleIntoStakeRequest) toEstimator(ethClient EthClientGetter) (*bindings.ContractEstimator, error) {
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

	amount := req.Promise.Amount
	fee := req.Promise.Fee
	lock := [32]byte{}
	copy(lock[:], req.Promise.R)

	channelID := [32]byte{}
	copy(channelID[:], req.Promise.ChannelID)

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	to, err := bc.makeTransactOpts(ctx, &req.WriteRequest)
	if err != nil {
		return nil, err
	}

	return t.SettleIntoStake(to, req.ProviderID, amount, fee, lock, req.Promise.Signature)
}

// DecreaseProviderStakeRequest represents all the parameters required for decreasing provider stake.
type DecreaseProviderStakeRequest struct {
	WriteRequest
	Request    crypto.DecreaseProviderStakeRequest
	ProviderID common.Address
}

func (r DecreaseProviderStakeRequest) toEstimator(ethClient EthClientGetter) (*bindings.ContractEstimator, error) {
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

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	transactor, err := bc.makeTransactOpts(ctx, &req.WriteRequest)
	if err != nil {
		return nil, fmt.Errorf("could not get transactor: %w", err)
	}

	return t.DecreaseStake(transactor, req.ProviderID, req.Request.Amount, req.Request.TransactorFee, req.Request.Signature)
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

func (r SettleAndRebalanceRequest) toEstimator(ethClient EthClientGetter) (*bindings.ContractEstimator, error) {
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

	to, err := bc.makeTransactOpts(ctx, &req.WriteRequest)
	if err != nil {
		return nil, err
	}

	return transactor.SettlePromise(
		to,
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

func (r SettleRequest) toEstimator(ethClient EthClientGetter) (*bindings.ContractEstimator, error) {
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

	if req.Nonce == nil {
		nonce, err := bc.getNonce(req.Identity)
		if err != nil {
			return nil, errors.Wrap(err, "could not get nonce")
		}
		req.Nonce = new(big.Int).SetUint64(nonce)
	}

	to, err := bc.makeTransactOpts(ctx, &req.WriteRequest)
	if err != nil {
		return nil, err

	}
	return transactor.SettlePromise(
		to, amount, fee, lock, req.Promise.Signature,
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

// GetHermesRegistry returns the registry address of a given hermes.
func (bc *Blockchain) GetHermesRegistry(hermesID common.Address) (common.Address, error) {
	caller, err := bindings.NewHermesImplementationCaller(hermesID, bc.ethClient.Client())
	if err != nil {
		return common.Address{}, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	return caller.GetRegistry(&bind.CallOpts{
		Context: ctx,
	})
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
		}, sink, []common.Address{}, providerAddresses, []common.Address{})
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

// FilterPromiseSettledEventByChannelID filters promise settled events
func (bc *Blockchain) FilterPromiseSettledEventByChannelID(from uint64, to *uint64, hermesID common.Address, providerAddresses [][32]byte) ([]bindings.HermesImplementationPromiseSettled, error) {
	caller, err := bindings.NewHermesImplementationFilterer(hermesID, bc.ethClient.Client())
	if err != nil {
		return nil, errors.Wrap(err, "could not create hermes caller")
	}
	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()
	iter, err := caller.FilterPromiseSettled(&bind.FilterOpts{
		Start:   from,
		End:     to,
		Context: ctx,
	}, []common.Address{}, providerAddresses, []common.Address{})
	if err != nil {
		return nil, err
	}
	res := make([]bindings.HermesImplementationPromiseSettled, 0)
	for iter.Next() {
		ev := iter.Event
		res = append(res, *ev)
	}
	return res, nil
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

	to, err := bc.makeTransactOpts(ctx, &etr.WriteRequest)
	if err != nil {
		return nil, err
	}

	var tx *types.Transaction
	if to.GasPrice != nil && to.GasPrice.Cmp(big.NewInt(0)) > 0 {
		tx = types.NewTransaction(to.Nonce.Uint64(), etr.To, etr.Amount, to.GasLimit, to.GasPrice, nil)
	} else {
		tx = types.NewTx(&types.DynamicFeeTx{
			Nonce:     to.Nonce.Uint64(),
			To:        &etr.To,
			Value:     etr.Amount,
			Gas:       to.GasLimit,
			GasFeeCap: to.GasFeeCap,
			GasTipCap: to.GasTipCap,
			Data:      nil,
		})
	}

	signedTx, err := etr.Signer(etr.Identity, tx)
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

// BlockNumber returns the last known block number
func (bc *Blockchain) BlockNumber() (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()
	return bc.ethClient.Client().BlockNumber(ctx)
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

func (r SettleWithBeneficiaryRequest) toEstimator(ethClient EthClientGetter) (*bindings.ContractEstimator, error) {
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

	to, err := bc.makeTransactOpts(ctx, &req.WriteRequest)
	if err != nil {
		return nil, err
	}

	return transactor.SettleWithBeneficiary(
		to,
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

type RewarderUpdateRoot struct {
	WriteRequest
	RewaderID   common.Address
	ClaimRoot   []byte
	BlockNumber *big.Int
	TotalReward *big.Int
}

func (bc *Blockchain) RewarderUpdateRoot(req RewarderUpdateRoot) (*types.Transaction, error) {
	transactor, err := rewarder.NewRewarderTransactor(req.RewaderID, bc.ethClient.Client())
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	to, err := bc.makeTransactOpts(ctx, &req.WriteRequest)
	if err != nil {
		return nil, err
	}

	return transactor.UpdateRoot(
		to,
		ToBytes32(req.ClaimRoot),
		req.BlockNumber,
		req.TotalReward,
	)
}

type RewarderAirDrop struct {
	WriteRequest
	RewaderID     common.Address
	Beneficiaries []common.Address
	TotalEarnings []*big.Int
}

func (bc *Blockchain) RewarderAirDrop(req RewarderAirDrop) (*types.Transaction, error) {
	transactor, err := rewarder.NewRewarderTransactor(req.RewaderID, bc.ethClient.Client())
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	to, err := bc.makeTransactOpts(ctx, &req.WriteRequest)
	if err != nil {
		return nil, err
	}

	return transactor.Airdrop(
		to,
		req.Beneficiaries,
		req.TotalEarnings,
	)
}

func (bc *Blockchain) RewarderTotalPayoutsFor(rewarderAddress common.Address, payoutsFor common.Address) (*big.Int, error) {
	caller, err := rewarder.NewRewarderCaller(rewarderAddress, bc.ethClient.Client())
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	return caller.TotalPayoutsFor(&bind.CallOpts{
		Context: ctx,
	}, payoutsFor)
}

// RewarderTotalClaimed is a free lookup in the blockchain for the total amount of claimed tokens in the blockchain.
func (bc *Blockchain) RewarderTotalClaimed(rewarderAddress common.Address) (*big.Int, error) {
	caller, err := rewarder.NewRewarderCaller(rewarderAddress, bc.ethClient.Client())
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	return caller.TotalClaimed(&bind.CallOpts{
		Context: ctx,
	})
}

type CustodyTokensTransfer struct {
	WriteRequest
	CustodyAddress common.Address
	Recipients     []common.Address
	Amounts        []*big.Int
}

func (bc *Blockchain) CustodyTransferTokens(req CustodyTokensTransfer) (*types.Transaction, error) {
	transactor, err := rewarder.NewCustodyTransactor(req.CustodyAddress, bc.ethClient.Client())
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	to, err := bc.makeTransactOpts(ctx, &req.WriteRequest)
	if err != nil {
		return nil, err
	}

	return transactor.Payout(
		to,
		req.Recipients,
		req.Amounts,
	)
}

type ApprovedAddress struct {
	Native       *big.Int
	Token        *big.Int
	BlocksWindow *big.Int
}

func (bc *Blockchain) TopperupperApprovedAddress(topperupperAddress common.Address, forAddress common.Address) (*ApprovedAddress, error) {
	caller, err := topperupper.NewTopperUpperCaller(topperupperAddress, bc.ethClient.Client())
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	res, err := caller.ApprovedAddresses(&bind.CallOpts{
		Context: ctx,
	}, forAddress)
	if err != nil {
		return nil, err
	}

	return &ApprovedAddress{
		Native:       res.Native,
		Token:        res.Token,
		BlocksWindow: res.BlocksWindow,
	}, nil
}

type CurrentLimits struct {
	Limit           *big.Int
	ValidUntilBlock *big.Int
}

func (bc *Blockchain) TopperupperNativeLimits(topperupperAddress common.Address, forAddress common.Address) (*CurrentLimits, error) {
	caller, err := topperupper.NewTopperUpperCaller(topperupperAddress, bc.ethClient.Client())
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	res, err := caller.NativeLimits(&bind.CallOpts{
		Context: ctx,
	}, forAddress)
	if err != nil {
		return nil, err
	}

	return &CurrentLimits{
		Limit:           res.Amount,
		ValidUntilBlock: res.ValidTill,
	}, nil
}

func (bc *Blockchain) TopperupperTokenLimits(topperupperAddress common.Address, forAddress common.Address) (*CurrentLimits, error) {
	caller, err := topperupper.NewTopperUpperCaller(topperupperAddress, bc.ethClient.Client())
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	res, err := caller.TokenLimits(&bind.CallOpts{
		Context: ctx,
	}, forAddress)
	if err != nil {
		return nil, err
	}

	return &CurrentLimits{
		Limit:           res.Amount,
		ValidUntilBlock: res.ValidTill,
	}, nil
}

type TopperupperApproveAddressesReq struct {
	WriteRequest
	ContractAddress common.Address
	Address         common.Address
	LimitsNative    *big.Int
	LimitsToken     *big.Int
	BlockWindow     *big.Int
}

func (bc *Blockchain) TopperupperApproveAddresses(req TopperupperApproveAddressesReq) (*types.Transaction, error) {
	transactor, err := topperupper.NewTopperUpperTransactor(req.ContractAddress, bc.ethClient.Client())
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	to, err := bc.makeTransactOpts(ctx, &req.WriteRequest)
	if err != nil {
		return nil, err
	}

	return transactor.ApproveAddresses(
		to,
		[]common.Address{req.Address},
		[]*big.Int{req.LimitsNative},
		[]*big.Int{req.LimitsToken},
		[]*big.Int{req.BlockWindow},
	)
}

type TopperupperModeratorsReq struct {
	WriteRequest
	ContractAddress common.Address
	Managers        []common.Address
}

func (bc *Blockchain) TopperupperSetManagers(req TopperupperModeratorsReq) (*types.Transaction, error) {
	transactor, err := topperupper.NewTopperUpperTransactor(req.ContractAddress, bc.ethClient.Client())
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	to, err := bc.makeTransactOpts(ctx, &req.WriteRequest)
	if err != nil {
		return nil, err
	}

	return transactor.SetManagers(
		to,
		req.Managers,
	)
}

type TopperupperTopupNativeReq struct {
	WriteRequest
	ContractAddress common.Address
	To              common.Address
	Amount          *big.Int
}

func (bc *Blockchain) TopperupperTopupNative(req TopperupperTopupNativeReq) (*types.Transaction, error) {
	transactor, err := topperupper.NewTopperUpperTransactor(req.ContractAddress, bc.ethClient.Client())
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	to, err := bc.makeTransactOpts(ctx, &req.WriteRequest)
	if err != nil {
		return nil, err
	}

	return transactor.TopupNative(
		to,
		req.To,
		req.Amount,
	)
}

type TopperupperTopupTokenReq struct {
	WriteRequest
	ContractAddress common.Address
	To              common.Address
	Amount          *big.Int
}

func (bc *Blockchain) TopperupperTopupToken(req TopperupperTopupTokenReq) (*types.Transaction, error) {
	transactor, err := topperupper.NewTopperUpperTransactor(req.ContractAddress, bc.ethClient.Client())
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	to, err := bc.makeTransactOpts(ctx, &req.WriteRequest)
	if err != nil {
		return nil, err
	}

	return transactor.TopupToken(
		to,
		req.To,
		req.Amount,
	)
}

type MystApproveReq struct {
	WriteRequest
	MystAddress common.Address
	Spender     common.Address
	Amount      *big.Int
}

func (bc *Blockchain) MystTokenApprove(req MystApproveReq) (*types.Transaction, error) {
	txer, err := bindings.NewMystTokenTransactor(req.MystAddress, bc.ethClient.Client())
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	to, err := bc.makeTransactOpts(ctx, &req.WriteRequest)
	if err != nil {
		return nil, err
	}

	return txer.Approve(to, req.Spender, req.Amount)
}

func (bc *Blockchain) MystAllowance(mystTokenAddress, holder, spender common.Address) (*big.Int, error) {
	caller, err := bindings.NewMystTokenCaller(mystTokenAddress, bc.ethClient.Client())
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	return caller.Allowance(&bind.CallOpts{
		Context: ctx,
	}, holder, spender)
}

type UniswapExactInputSingleReq struct {
	WriteRequest
	SwapRouterAddress common.Address

	TokenIn  common.Address
	TokenOut common.Address
	Fee      *big.Int

	Recipient       common.Address
	DeadlineSeconds uint64

	AmountIn         *big.Int
	AmountOutMinimum *big.Int
}

func (bc *Blockchain) UniswapV3ExactInputSingle(req UniswapExactInputSingleReq) (*types.Transaction, error) {
	txer, err := uniswapv3.NewSwapRouterTransactor(req.SwapRouterAddress, bc.ethClient.Client())
	if err != nil {
		return nil, err
	}

	ctx1, cancel1 := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel1()

	b, err := bc.ethClient.Client().BlockByNumber(ctx1, nil)
	if err != nil {
		return nil, err
	}

	ctx2, cancel2 := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel2()

	to, err := bc.makeTransactOpts(ctx2, &req.WriteRequest)
	if err != nil {
		return nil, err
	}

	return txer.ExactInputSingle(
		to,
		uniswapv3.ISwapRouterExactInputSingleParams{
			TokenIn:  req.TokenIn,
			TokenOut: req.TokenOut,
			Fee:      req.Fee,

			Recipient: req.Recipient,
			Deadline:  big.NewInt(0).SetUint64(b.Time() + req.DeadlineSeconds),

			AmountIn:         req.AmountIn,
			AmountOutMinimum: req.AmountOutMinimum,

			SqrtPriceLimitX96: big.NewInt(0), // We can mostly ignore it for our purposes.
		})
}

type SwapTokenPair struct {
	Token0 common.Address
	Token1 common.Address
}

func (bc *Blockchain) UniswapV3TokenPair(poolAddress common.Address) (*SwapTokenPair, error) {
	caller, err := uniswapv3.NewPoolCaller(poolAddress, bc.ethClient.Client())
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	res := &SwapTokenPair{}

	res.Token0, err = caller.Token0(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}

	res.Token1, err = caller.Token1(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (bc *Blockchain) UniswapV3PoolFee(poolAddress common.Address) (*big.Int, error) {
	caller, err := uniswapv3.NewPoolCaller(poolAddress, bc.ethClient.Client())
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	return caller.Fee(&bind.CallOpts{
		Context: ctx,
	})
}

func (bc *Blockchain) WMaticBalance(holder, wmaticAddress common.Address) (*big.Int, error) {
	caller, err := uniswapv3.NewWmaticCaller(wmaticAddress, bc.ethClient.Client())
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	return caller.BalanceOf(&bind.CallOpts{
		Context: ctx,
	}, holder)
}

type WMaticWithdrawReq struct {
	WriteRequest
	WMaticAddress common.Address

	// Amount describes how much should be withdrawn.
	// If `nil` all balance is withdrawn.
	Amount *big.Int
}

func (bc *Blockchain) WMaticWithdraw(req WMaticWithdrawReq) (*types.Transaction, error) {
	caller, err := uniswapv3.NewWmaticTransactor(req.WMaticAddress, bc.ethClient.Client())
	if err != nil {
		return nil, err
	}

	amount := req.Amount
	if amount == nil {
		all, err := bc.WMaticBalance(req.Identity, req.WMaticAddress)
		if err != nil {
			return nil, fmt.Errorf("tried to pull wmatic balance, but failed: %w", err)
		}

		amount = all
	}

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	to, err := bc.makeTransactOpts(ctx, &req.WriteRequest)
	if err != nil {
		return nil, err
	}

	return caller.Withdraw(to, amount)
}
