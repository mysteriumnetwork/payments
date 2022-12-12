/*
 * Copyright (C) 2021 The "MysteriumNetwork/payments" Authors.
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

package client

import (
	"context"
	"errors"
	"math/big"
	"net"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/mysteriumnetwork/payments/crypto"

	"github.com/mysteriumnetwork/payments/bindings"
)

type BC interface {
	CalculateHermesFee(hermesAddress crypto.Address, value *big.Int) (*big.Int, error)
	IsHermesActive(hermesID crypto.Address) (bool, error)
	IsHermesRegistered(registryAddress, acccountantID crypto.Address) (bool, error)

	GetHermesFee(hermesAddress crypto.Address) (uint16, error)
	GetHermessAvailableBalance(hermesAddress crypto.Address) (*big.Int, error)
	GetHermesURL(registryID, hermesID crypto.Address) (string, error)
	GetHermes(registryID, hermesID crypto.Address) (Hermes, error)
	GetHermesRegistry(hermesID crypto.Address) (crypto.Address, error)
	GetHermesOperator(hermesID crypto.Address) (crypto.Address, error)
	GetProviderChannel(hermesAddress crypto.Address, addressToCheck crypto.Address, pending bool) (ProviderChannel, error)
	GetMystBalance(mystSCAddress, address crypto.Address) (*big.Int, error)
	GetConsumerChannelsHermes(channelAddress crypto.Address) (ConsumersHermes, error)
	GetConsumerChannelOperator(channelAddress crypto.Address) (crypto.Address, error)
	GetProviderChannelByID(acc crypto.Address, chID []byte) (ProviderChannel, error)
	GetConsumerChannel(addr crypto.Address, mystSCAddress crypto.Address) (ConsumerChannel, error)
	GetEthBalance(address crypto.Address) (*big.Int, error)
	GetStakeThresholds(hermesID crypto.Address) (min, max *big.Int, err error)
	GetBeneficiary(registryAddress, identity crypto.Address) (crypto.Address, error)
	GetLastRegistryNonce(registry crypto.Address) (*big.Int, error)
	GetChannelImplementationByVersion(registryID crypto.Address, version *big.Int) (crypto.Address, error)
	GetProvidersWithdrawalChannel(hermesAddress crypto.Address, addressToCheck crypto.Address, pending bool) (ProviderChannel, error)
	SubscribeToWithdrawalPromiseSettledEvent(providerID, hermesID crypto.Address) (sink chan *bindings.HermesImplementationPromiseSettled, cancel func(), err error)

	IsRegisteredAsProvider(hermesAddress, registryAddress, addressToCheck crypto.Address) (bool, error)
	IsRegistered(registryAddress, addressToCheck crypto.Address) (bool, error)
	RegisterIdentity(rr RegistrationRequest) (*types.Transaction, error)
	OpenConsumerChannel(req OpenConsumerChannelRequest) (*types.Transaction, error)
	IncreaseProviderStake(req ProviderStakeIncreaseRequest) (*types.Transaction, error)

	SettleAndRebalance(req SettleAndRebalanceRequest) (*types.Transaction, error)
	SettleWithBeneficiary(req SettleWithBeneficiaryRequest) (*types.Transaction, error)
	SettlePromise(req SettleRequest) (*types.Transaction, error)
	DecreaseProviderStake(req DecreaseProviderStakeRequest) (*types.Transaction, error)
	SettleIntoStake(req SettleIntoStakeRequest) (*types.Transaction, error)

	SubscribeToPromiseSettledEvent(providerID, hermesID crypto.Address) (sink chan *bindings.HermesImplementationPromiseSettled, cancel func(), err error)
	SubscribeToConsumerBalanceEvent(channel, mystSCAddress crypto.Address, timeout time.Duration) (chan *bindings.MystTokenTransfer, func(), error)
	SubscribeToIdentityRegistrationEvents(registryAddress crypto.Address) (sink chan *bindings.RegistryRegisteredIdentity, cancel func(), err error)
	SubscribeToConsumerChannelBalanceUpdate(mystSCAddress crypto.Address, channelAddresses []crypto.Address) (sink chan *bindings.MystTokenTransfer, cancel func(), err error)
	SubscribeToPromiseSettledEventByChannelID(hermesID crypto.Address, providerAddresses [][32]byte) (sink chan *bindings.HermesImplementationPromiseSettled, cancel func(), err error)
	SubscribeToMystTokenTransfers(mystSCAddress crypto.Address) (chan *bindings.MystTokenTransfer, func(), error)
	FilterLogs(q ethereum.FilterQuery) ([]types.Log, error)
	FilterHermesRegistered(from uint64, to *uint64, registryID crypto.Address) ([]bindings.RegistryRegisteredHermes, error)
	FilterHermesURLUpdated(from uint64, to *uint64, registryID crypto.Address) ([]bindings.RegistryHermesURLUpdated, error)

	NetworkID() (*big.Int, error)
	BlockNumber() (uint64, error)
	SuggestGasPrice() (*big.Int, error)
	HeaderByNumber(number *big.Int) (*types.Header, error)
	PendingNonceAt(account crypto.Address) (uint64, error)
	NonceAt(account crypto.Address, blockNum *big.Int) (uint64, error)

	TransferMyst(req TransferRequest) (tx *types.Transaction, err error)
	TransferEth(etr EthTransferRequest) (*types.Transaction, error)

	SendTransaction(tx *types.Transaction) error
	TransactionReceipt(hash common.Hash) (*types.Receipt, error)
	TransactionByHash(hash common.Hash) (*types.Transaction, bool, error)

	RewarderTotalPayoutsFor(rewarderAddress crypto.Address, payoutsFor crypto.Address) (*big.Int, error)
	RewarderAirDrop(req RewarderAirDrop) (*types.Transaction, error)
	RewarderUpdateRoot(req RewarderUpdateRoot) (*types.Transaction, error)
	RewarderTotalClaimed(rewarderAddress crypto.Address) (*big.Int, error)
	CustodyTransferTokens(req CustodyTokensTransfer) (*types.Transaction, error)

	PayAndSettle(psr PayAndSettleRequest) (*types.Transaction, error)
	IsChannelOpened(registryID, identity, hermesID crypto.Address) (bool, error)
	FilterPromiseSettledEventByChannelID(from uint64, to *uint64, hermesID crypto.Address, providerAddresses [][32]byte) ([]bindings.HermesImplementationPromiseSettled, error)

	TopperupperTopupToken(req TopperupperTopupTokenReq) (*types.Transaction, error)
	TopperupperTopupNative(req TopperupperTopupNativeReq) (*types.Transaction, error)
	TopperupperSetManagers(req TopperupperModeratorsReq) (*types.Transaction, error)
	TopperupperApproveAddresses(req TopperupperApproveAddressesReq) (*types.Transaction, error)
	TopperupperApprovedAddress(topperupperAddress crypto.Address, forAddress crypto.Address) (*ApprovedAddress, error)
	TopperupperNativeLimits(topperupperAddress crypto.Address, forAddress crypto.Address) (*CurrentLimits, error)
	TopperupperTokenLimits(topperupperAddress crypto.Address, forAddress crypto.Address) (*CurrentLimits, error)

	MystTokenApprove(req MystApproveReq) (*types.Transaction, error)
	MystAllowance(mystTokenAddress, holder, spender crypto.Address) (*big.Int, error)
	UniswapV3ExactInputSingle(req UniswapExactInputSingleReq) (*types.Transaction, error)
	UniswapV3TokenPair(poolAddress crypto.Address) (*SwapTokenPair, error)
	UniswapV3PoolFee(poolAddress crypto.Address) (*big.Int, error)
	WMaticBalance(holder, wmaticAddress crypto.Address) (*big.Int, error)
	WMaticWithdraw(req WMaticWithdrawReq) (*types.Transaction, error)
}

// EtherClient interface implements all methods required for a EtherClient to work
type EtherClient interface {
	Close()
	// ChainId retrieves the current chain ID for transaction replay protection.
	ChainID(ctx context.Context) (*big.Int, error)
	// BlockByHash returns the given full block.
	//
	// Note that loading full blocks requires two requests. Use HeaderByHash
	// if you don't need all transactions or uncle headers.
	BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error)
	// BlockByNumber returns a block from the current canonical chain. If number is nil, the
	// latest known block is returned.
	//
	// Note that loading full blocks requires two requests. Use HeaderByNumber
	// if you don't need all transactions or uncle headers.
	BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error)
	// BlockNumber returns the most recent block number
	BlockNumber(ctx context.Context) (uint64, error)
	// HeaderByHash returns the block header with the given hash.
	HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error)
	// HeaderByNumber returns a block header from the current canonical chain. If number is
	// nil, the latest known header is returned.
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
	// TransactionByHash returns the transaction with the given hash.
	TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error)
	// TransactionSender returns the sender address of the given transaction. The transaction
	// must be known to the remote node and included in the blockchain at the given block and
	// index. The sender is the one derived by the protocol at the time of inclusion.
	//
	// There is a fast-path for transactions retrieved by TransactionByHash and
	// TransactionInBlock. Getting their sender address can be done without an RPC interaction.
	TransactionSender(ctx context.Context, tx *types.Transaction, block common.Hash, index uint) (common.Address, error)
	// TransactionCount returns the total number of transactions in the given block.
	TransactionCount(ctx context.Context, blockHash common.Hash) (uint, error)
	// TransactionInBlock returns a single transaction at index in the given block.
	TransactionInBlock(ctx context.Context, blockHash common.Hash, index uint) (*types.Transaction, error)
	// TransactionReceipt returns the receipt of a transaction by transaction hash.
	// Note that the receipt is not available for pending transactions.
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
	// SyncProgress retrieves the current progress of the sync algorithm. If there's
	// no sync currently running, it returns nil.
	SyncProgress(ctx context.Context) (*ethereum.SyncProgress, error)
	// SubscribeNewHead subscribes to notifications about the current blockchain head
	// on the given channel.
	SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error)
	// NetworkID returns the network ID (also known as the chain ID) for this chain.
	NetworkID(ctx context.Context) (*big.Int, error)
	// BalanceAt returns the wei balance of the given account.
	// The block number can be nil, in which case the balance is taken from the latest known block.
	BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error)
	// StorageAt returns the value of key in the contract storage of the given account.
	// The block number can be nil, in which case the value is taken from the latest known block.
	StorageAt(ctx context.Context, account common.Address, key common.Hash, blockNumber *big.Int) ([]byte, error)
	// CodeAt returns the contract code of the given account.
	// The block number can be nil, in which case the code is taken from the latest known block.
	CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error)
	// NonceAt returns the account nonce of the given account.
	// The block number can be nil, in which case the nonce is taken from the latest known block.
	NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error)
	// FilterLogs executes a filter query.
	FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error)
	// SubscribeFilterLogs subscribes to the results of a streaming filter query.
	SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error)
	// PendingBalanceAt returns the wei balance of the given account in the pending state.
	PendingBalanceAt(ctx context.Context, account common.Address) (*big.Int, error)
	// PendingStorageAt returns the value of key in the contract storage of the given account in the pending state.
	PendingStorageAt(ctx context.Context, account common.Address, key common.Hash) ([]byte, error)
	// PendingCodeAt returns the contract code of the given account in the pending state.
	PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error)
	// PendingNonceAt returns the account nonce of the given account in the pending state.
	// This is the nonce that should be used for the next transaction.
	PendingNonceAt(ctx context.Context, account common.Address) (uint64, error)
	// PendingTransactionCount returns the total number of transactions in the pending state.
	PendingTransactionCount(ctx context.Context) (uint, error)
	// CallContract executes a message call transaction, which is directly executed in the VM
	// of the node, but never mined into the blockchain.
	//
	// blockNumber selects the block height at which the call runs. It can be nil, in which
	// case the code is taken from the latest known block. Note that state from very old
	// blocks might not be available.
	CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error)
	// PendingCallContract executes a message call transaction using the EVM.
	// The state seen by the contract call is the pending state.
	PendingCallContract(ctx context.Context, msg ethereum.CallMsg) ([]byte, error)
	// SuggestGasPrice retrieves the currently suggested gas price to allow a timely
	// execution of a transaction.
	SuggestGasPrice(ctx context.Context) (*big.Int, error)
	// SuggestGasTipCap retrieves the currently suggested 1559 priority fee to allow
	// a timely execution of a transaction.
	SuggestGasTipCap(ctx context.Context) (*big.Int, error)
	// EstimateGas tries to estimate the gas needed to execute a specific transaction based on
	// the current pending state of the backend blockchain. There is no guarantee that this is
	// the true gas limit requirement as other transactions may be added or removed by miners,
	// but it should provide a basis for setting a reasonable default.
	EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error)
	// SendTransaction injects a signed transaction into the pending pool for execution.
	//
	// If the transaction was a contract creation use the TransactionReceipt method to get the
	// contract address after the transaction has been mined.
	SendTransaction(ctx context.Context, tx *types.Transaction) error
}

// IsErrConnectionFailed returns true if the error represents
// a failure to reach blockchain client.
func IsErrConnectionFailed(err error) bool {
	if errors.Is(err, context.DeadlineExceeded) {
		return true
	}

	if errors.Is(err, context.Canceled) {
		return true
	}

	var e net.Error
	if errors.As(err, &e) {
		return e.Timeout()
	}

	return false
}

// EthClientGetter wraps any eth client.
type EthClientGetter interface {
	// Client returns a client to use for making call to the eth blockchain.
	Client() EtherClient
}

// AddressableEthClientGetter wraps any eth client and is able to return the address
// used to create the client and the client itself.
type AddressableEthClientGetter interface {
	EthClientGetter
	// Address returns an address which was used to create a particular client
	Address() string
}

type DefaultEthClientGetter struct {
	client EtherClient
}

func NewDefaultEthClientGetter(cl EtherClient) EthClientGetter {
	return &DefaultEthClientGetter{
		client: cl,
	}
}

func (eth *DefaultEthClientGetter) Client() EtherClient {
	return eth.client
}

type DefaultAddressableEthClientGetter struct {
	address string
	*DefaultEthClientGetter
}

func NewDefaultAddressableEthClientGetter(address string, cl EtherClient) AddressableEthClientGetter {
	return &DefaultAddressableEthClientGetter{
		DefaultEthClientGetter: &DefaultEthClientGetter{
			client: cl,
		},
		address: address,
	}
}

func (eth *DefaultAddressableEthClientGetter) Address() string {
	return eth.address
}
