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
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// EthMultiClient is a ethereum client that can reconnect.
type EthMultiClient struct {
	// timeout is the default amount of time to be assigned for a parent context
	// from a caller if it doesn't have one.
	timeout time.Duration

	// clients holds all the possible clients to call.
	clients []AddressableEthClientGetter

	// notifyDown is an optional channel.
	//
	// If given updates will be sent to this channel
	// when a certain client goes down.
	notifyDown *safeChannel

	mu sync.Mutex
}

type Notification struct {
	Address string
	Error   error
}

type safeChannel struct {
	ch chan<- Notification
	mu sync.Mutex
}

type doFunc func(ctx context.Context, c EtherClient)

var ErrClientNoConnection = errors.New("failed to connect to the eth client with a given address")

// NewEthMultiClient creates a new multi clients eth client.
func NewEthMultiClient(defaulTimeout time.Duration, clients []AddressableEthClientGetter) (*EthMultiClient, error) {
	if len(clients) == 0 {
		return nil, errors.New("expected more than 0 clients to use")
	}

	return &EthMultiClient{
		timeout: defaulTimeout,
		clients: clients,
		notifyDown: &safeChannel{
			ch: nil,
		},
	}, nil
}

// Client implements the EthClientGetter interface and returns itself as a EtherClient.
func (c *EthMultiClient) Client() EtherClient {
	return c
}

// NewEthMultiClientNotifyDown creates a new multi clients eth client.
//
// Channel `notifications` must be given and will be used to push notifications to the
// client if any nodes go down. The channel is closed when before the clients are closed.
func NewEthMultiClientNotifyDown(defaulTimeout time.Duration, clients []AddressableEthClientGetter, notifications chan<- Notification) (*EthMultiClient, error) {
	if len(clients) == 0 {
		return nil, errors.New("expected more than 0 clients to use")
	}

	return &EthMultiClient{
		timeout: defaulTimeout,
		clients: clients,
		notifyDown: &safeChannel{
			ch: notifications,
		},
	}, nil
}

// Close will close all the clients the multiclient instance holds.
func (c *EthMultiClient) Close() {
	c.closeNotify()
	c.doWithMultipleClients(context.Background(), func(ctx context.Context, c EtherClient) error {
		c.Close()
		return nil
	})
}

// ChainId retrieves the current chain ID for transaction replay protection.
func (c *EthMultiClient) ChainID(ctx context.Context) (*big.Int, error) {
	var res *big.Int
	return res, c.doWithMultipleClients(ctx, func(ctx context.Context, c EtherClient) error {
		val, err := c.ChainID(ctx)
		if err != nil {
			return err
		}

		res = val
		return nil
	})
}

// BlockByHash returns the given full block.
//
// Note that loading full blocks requires two requests. Use HeaderByHash
// if you don't need all transactions or uncle headers.
func (c *EthMultiClient) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	var res *types.Block
	return res, c.doWithMultipleClients(ctx, func(ctx context.Context, c EtherClient) error {
		val, err := c.BlockByHash(ctx, hash)
		if err != nil {
			return err
		}

		res = val
		return nil
	})
}

// BlockByNumber returns a block from the current canonical chain. If number is nil, the
// latest known block is returned.
//
// Note that loading full blocks requires two requests. Use HeaderByNumber
// if you don't need all transactions or uncle headers.
func (c *EthMultiClient) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	var res *types.Block
	return res, c.doWithMultipleClients(ctx, func(ctx context.Context, c EtherClient) error {
		val, err := c.BlockByNumber(ctx, number)
		if err != nil {
			return err
		}

		res = val
		return nil
	})
}

// BlockNumber returns the most recent block number
func (c *EthMultiClient) BlockNumber(ctx context.Context) (uint64, error) {
	var res uint64
	return res, c.doWithMultipleClients(ctx, func(ctx context.Context, c EtherClient) error {
		val, err := c.BlockNumber(ctx)
		if err != nil {
			return err
		}

		res = val
		return nil
	})
}

// HeaderByHash returns the block header with the given hash.
func (c *EthMultiClient) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	var res *types.Header
	return res, c.doWithMultipleClients(ctx, func(ctx context.Context, c EtherClient) error {
		val, err := c.HeaderByHash(ctx, hash)
		if err != nil {
			return err
		}

		res = val
		return nil
	})
}

// HeaderByNumber returns a block header from the current canonical chain. If number is
// nil, the latest known header is returned.
func (c *EthMultiClient) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	var res *types.Header
	return res, c.doWithMultipleClients(ctx, func(ctx context.Context, c EtherClient) error {
		val, err := c.HeaderByNumber(ctx, number)
		if err != nil {
			return err
		}

		res = val
		return nil
	})
}

// TransactionByHash returns the transaction with the given hash.
func (c *EthMultiClient) TransactionByHash(ctx context.Context, hash common.Hash) (*types.Transaction, bool, error) {
	var tx *types.Transaction
	var isPending bool
	return tx, isPending, c.doWithMultipleClients(ctx, func(ctx context.Context, c EtherClient) error {
		t, is, err := c.TransactionByHash(ctx, hash)
		if err != nil {
			return err
		}

		tx = t
		isPending = is
		return nil
	})
}

// TransactionSender returns the sender address of the given transaction. The transaction
// must be known to the remote node and included in the blockchain at the given block and
// index. The sender is the one derived by the protocol at the time of inclusion.
//
// There is a fast-path for transactions retrieved by TransactionByHash and
// TransactionInBlock. Getting their sender address can be done without an RPC interaction.
func (c *EthMultiClient) TransactionSender(ctx context.Context, tx *types.Transaction, block common.Hash, index uint) (common.Address, error) {
	var res common.Address
	return res, c.doWithMultipleClients(ctx, func(ctx context.Context, c EtherClient) error {
		val, err := c.TransactionSender(ctx, tx, block, index)
		if err != nil {
			return err
		}

		res = val
		return nil
	})
}

// TransactionCount returns the total number of transactions in the given block.
func (c *EthMultiClient) TransactionCount(ctx context.Context, blockHash common.Hash) (uint, error) {
	var res uint
	return res, c.doWithMultipleClients(ctx, func(ctx context.Context, c EtherClient) error {
		val, err := c.TransactionCount(ctx, blockHash)
		if err != nil {
			return err
		}

		res = val
		return nil
	})
}

// TransactionInBlock returns a single transaction at index in the given block.
func (c *EthMultiClient) TransactionInBlock(ctx context.Context, blockHash common.Hash, index uint) (*types.Transaction, error) {
	var res *types.Transaction
	return res, c.doWithMultipleClients(ctx, func(ctx context.Context, c EtherClient) error {
		val, err := c.TransactionInBlock(ctx, blockHash, index)
		if err != nil {
			return err
		}

		res = val
		return nil
	})
}

// TransactionReceipt returns the receipt of a transaction by transaction hash.
// Note that the receipt is not available for pending transactions.
func (c *EthMultiClient) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	var res *types.Receipt
	return res, c.doWithMultipleClients(ctx, func(ctx context.Context, c EtherClient) error {
		val, err := c.TransactionReceipt(ctx, txHash)
		if err != nil {
			return err
		}

		res = val
		return nil
	})
}

// SyncProgress retrieves the current progress of the sync algorithm. If there's
// no sync currently running, it returns nil.
func (c *EthMultiClient) SyncProgress(ctx context.Context) (*ethereum.SyncProgress, error) {
	var res *ethereum.SyncProgress
	return res, c.doWithMultipleClients(ctx, func(ctx context.Context, c EtherClient) error {
		val, err := c.SyncProgress(ctx)
		if err != nil {
			return err
		}

		res = val
		return nil
	})
}

// SubscribeNewHead subscribes to notifications about the current blockchain head
// on the given channel.
func (c *EthMultiClient) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
	var res ethereum.Subscription
	return res, c.doWithMultipleClients(ctx, func(ctx context.Context, c EtherClient) error {
		val, err := c.SubscribeNewHead(ctx, ch)
		if err != nil {
			return err
		}

		res = val
		return nil
	})
}

// NetworkID returns the network ID (also known as the chain ID) for this chain.
func (c *EthMultiClient) NetworkID(ctx context.Context) (*big.Int, error) {
	var res *big.Int
	return res, c.doWithMultipleClients(ctx, func(ctx context.Context, c EtherClient) error {
		val, err := c.NetworkID(ctx)
		if err != nil {
			return err
		}

		res = val
		return nil
	})
}

// BalanceAt returns the wei balance of the given account.
// The block number can be nil, in which case the balance is taken from the latest known block.
func (c *EthMultiClient) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	var res *big.Int
	return res, c.doWithMultipleClients(ctx, func(ctx context.Context, c EtherClient) error {
		val, err := c.BalanceAt(ctx, account, blockNumber)
		if err != nil {
			return err
		}

		res = val
		return nil
	})
}

// StorageAt returns the value of key in the contract storage of the given account.
// The block number can be nil, in which case the value is taken from the latest known block.
func (c *EthMultiClient) StorageAt(ctx context.Context, account common.Address, key common.Hash, blockNumber *big.Int) ([]byte, error) {
	var res []byte
	return res, c.doWithMultipleClients(ctx, func(ctx context.Context, c EtherClient) error {
		val, err := c.StorageAt(ctx, account, key, blockNumber)
		if err != nil {
			return err
		}

		res = val
		return nil
	})
}

// CodeAt returns the contract code of the given account.
// The block number can be nil, in which case the code is taken from the latest known block.
func (c *EthMultiClient) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	var res []byte
	return res, c.doWithMultipleClients(ctx, func(ctx context.Context, c EtherClient) error {
		val, err := c.CodeAt(ctx, account, blockNumber)
		if err != nil {
			return err
		}

		res = val
		return nil
	})
}

// NonceAt returns the account nonce of the given account.
// The block number can be nil, in which case the nonce is taken from the latest known block.
func (c *EthMultiClient) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	var res uint64
	return res, c.doWithMultipleClients(ctx, func(ctx context.Context, c EtherClient) error {
		val, err := c.NonceAt(ctx, account, blockNumber)
		if err != nil {
			return err
		}

		res = val
		return nil
	})
}

// FilterLogs executes a filter query.
func (c *EthMultiClient) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	var res []types.Log
	return res, c.doWithMultipleClients(ctx, func(ctx context.Context, c EtherClient) error {
		val, err := c.FilterLogs(ctx, q)
		if err != nil {
			return err
		}

		res = val
		return nil
	})
}

// SubscribeFilterLogs subscribes to the results of a streaming filter query.
func (c *EthMultiClient) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	var res ethereum.Subscription
	return res, c.doWithMultipleClients(ctx, func(ctx context.Context, c EtherClient) error {
		val, err := c.SubscribeFilterLogs(ctx, q, ch)
		if err != nil {
			return err
		}

		res = val
		return nil
	})
}

// PendingBalanceAt returns the wei balance of the given account in the pending state.
func (c *EthMultiClient) PendingBalanceAt(ctx context.Context, account common.Address) (*big.Int, error) {
	var res *big.Int
	return res, c.doWithMultipleClients(ctx, func(ctx context.Context, c EtherClient) error {
		val, err := c.PendingBalanceAt(ctx, account)
		if err != nil {
			return err
		}

		res = val
		return nil
	})
}

// PendingStorageAt returns the value of key in the contract storage of the given account in the pending state.
func (c *EthMultiClient) PendingStorageAt(ctx context.Context, account common.Address, key common.Hash) ([]byte, error) {
	var res []byte
	return res, c.doWithMultipleClients(ctx, func(ctx context.Context, c EtherClient) error {
		val, err := c.PendingStorageAt(ctx, account, key)
		if err != nil {
			return err
		}

		res = val
		return nil
	})
}

// PendingCodeAt returns the contract code of the given account in the pending state.
func (c *EthMultiClient) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	var res []byte
	return res, c.doWithMultipleClients(ctx, func(ctx context.Context, c EtherClient) error {
		val, err := c.PendingCodeAt(ctx, account)
		if err != nil {
			return err
		}

		res = val
		return nil
	})
}

// PendingNonceAt returns the account nonce of the given account in the pending state.
// This is the nonce that should be used for the next transaction.
func (c *EthMultiClient) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	var res uint64
	return res, c.doWithMultipleClients(ctx, func(ctx context.Context, c EtherClient) error {
		val, err := c.PendingNonceAt(ctx, account)
		if err != nil {
			return err
		}

		res = val
		return nil
	})
}

// PendingTransactionCount returns the total number of transactions in the pending state.
func (c *EthMultiClient) PendingTransactionCount(ctx context.Context) (uint, error) {
	var res uint
	return res, c.doWithMultipleClients(ctx, func(ctx context.Context, c EtherClient) error {
		val, err := c.PendingTransactionCount(ctx)
		if err != nil {
			return err
		}

		res = val
		return nil
	})
}

// CallContract executes a message call transaction, which is directly executed in the VM
// of the node, but never mined into the blockchain.
//
// blockNumber selects the block height at which the call runs. It can be nil, in which
// case the code is taken from the latest known block. Note that state from very old
// blocks might not be available.
func (c *EthMultiClient) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	var res []byte
	return res, c.doWithMultipleClients(ctx, func(ctx context.Context, c EtherClient) error {
		val, err := c.CallContract(ctx, msg, blockNumber)
		if err != nil {
			return err
		}

		res = val
		return nil
	})
}

// PendingCallContract executes a message call transaction using the EVM.
// The state seen by the contract call is the pending state.
func (c *EthMultiClient) PendingCallContract(ctx context.Context, msg ethereum.CallMsg) ([]byte, error) {
	var res []byte
	return res, c.doWithMultipleClients(ctx, func(ctx context.Context, c EtherClient) error {
		val, err := c.PendingCallContract(ctx, msg)
		if err != nil {
			return err
		}

		res = val
		return nil
	})
}

// SuggestGasPrice retrieves the currently suggested gas price to allow a timely
// execution of a transaction.
func (c *EthMultiClient) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	var res *big.Int
	return res, c.doWithMultipleClients(ctx, func(ctx context.Context, c EtherClient) error {
		val, err := c.SuggestGasPrice(ctx)
		if err != nil {
			return err
		}

		res = val
		return nil
	})
}

// EstimateGas tries to estimate the gas needed to execute a specific transaction based on
// the current pending state of the backend blockchain. There is no guarantee that this is
// the true gas limit requirement as other transactions may be added or removed by miners,
// but it should provide a basis for setting a reasonable default.
func (c *EthMultiClient) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	var res uint64
	return res, c.doWithMultipleClients(ctx, func(ctx context.Context, c EtherClient) error {
		val, err := c.EstimateGas(ctx, msg)
		if err != nil {
			return err
		}

		res = val
		return nil
	})
}

// SendTransaction injects a signed transaction into the pending pool for execution.
//
// If the transaction was a contract creation use the TransactionReceipt method to get the
// contract address after the transaction has been mined.
func (c *EthMultiClient) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return c.doWithMultipleClients(ctx, func(ctx context.Context, c EtherClient) error {
		return c.SendTransaction(ctx, tx)
	})
}

// ReorderClients reorders clients to use in BC calls by the given address slice.
// If there are a lot of clients this function could be expensive and it will block all other
// calls or usage of these clients while reordering.
func (c *EthMultiClient) ReorderClients(addresses []string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if len(c.clients) != len(addresses) {
		return fmt.Errorf("can't reorder: given %d addresses to use when reordering but have %d clients", len(addresses), len(c.clients))
	}

	newClients := make([]AddressableEthClientGetter, len(c.clients))
	for i, addr := range addresses {
		found := false
		for _, cl := range c.clients {
			if addr == cl.Address() {
				newClients[i] = cl
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("can't re-order: no client with address: %s", addr)
		}
	}

	c.clients = newClients
	return nil
}

// CurrentClientOrder returns the current orders of all clients the multiclient uses.
func (c *EthMultiClient) CurrentClientOrder() []string {
	c.mu.Lock()
	defer c.mu.Unlock()

	result := make([]string, len(c.clients))
	for i, c := range c.clients {
		result[i] = c.Address()
	}
	return result
}

// CallSpecificClient allows to call a spefific client by a given address.
func (c *EthMultiClient) CallSpecificClient(address string, call func(c EtherClient) error) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	var client EtherClient
	for _, cl := range c.clients {
		if cl.Address() == address {
			client = cl.Client()
			break
		}
	}

	if client == nil {
		return errors.New("client not found")
	}

	return call(client)
}

// doWithMultipleClientsCtx will execute a given function with all clients.
//
// If parent context is cancel or receives a timeout, all calls will be also cancels and the
// function will return.
func (c *EthMultiClient) doWithMultipleClients(ctx context.Context, do func(ctx context.Context, c EtherClient) error) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	ctxs, cancel := c.produceCtxs(ctx)
	defer cancel()

	done := make(chan struct{}, 0)
	defer close(done)

	go func() {
		select {
		case <-done:
			return
		case <-ctx.Done():
			cancel()
		}
	}()

	var err error
	for i, cl := range c.clients {
		select {
		case <-ctx.Done():
			return context.DeadlineExceeded
		default:
			childCtx := ctxs[i]
			if childCtx.Err() != nil && ctx.Err() != nil {
				return ctx.Err()
			}

			err = do(childCtx, cl.Client())
			if err != nil {
				if errors.Is(err, context.DeadlineExceeded) {
					// If parent context wasn't canceled that means that we've timed out on this node.
					if ctx.Err() == nil {
						c.notifyDowntime(cl.Address())
					}
					continue
				}
				return err
			}

			return nil
		}
	}

	return err
}

func (c *EthMultiClient) notifyDowntime(address string) {
	c.notifyDown.mu.Lock()
	defer c.notifyDown.mu.Unlock()

	// If no channel is given, just return.
	if c.notifyDown.ch == nil {
		return
	}

	payload := Notification{
		Address: address,
		Error:   ErrClientNoConnection,
	}
	// If channel is full, drop the notification.
	select {
	case c.notifyDown.ch <- payload:
	default:
	}
}

func (c *EthMultiClient) closeNotify() {
	c.notifyDown.mu.Lock()
	defer c.notifyDown.mu.Unlock()

	if c.notifyDown.ch != nil {
		close(c.notifyDown.ch)
		// Set equal to `nil` so that any hanging queries dont
		// get broadcasted as failure on client close.
		c.notifyDown.ch = nil
	}
}

// produceCtxs produces contexts for each client.
func (c *EthMultiClient) produceCtxs(ctx context.Context) ([]context.Context, func()) {
	singleCtxDuration := c.childTimeout(ctx)
	ctxs := make([]context.Context, len(c.clients))
	cancels := make([]func(), len(c.clients))

	for i := range ctxs {
		ctxs[i], cancels[i] = context.WithTimeout(context.Background(), singleCtxDuration*time.Duration((i+1)))
	}

	cancelFn := func() {
		for _, fn := range cancels {
			fn()
		}
	}

	return ctxs, cancelFn
}

// childTimeout produces an average duration to use for each child context.
//
// It rounds the extracted duration down to miliseconds elimiating
// the posiblity to get context duration like: 1.99999999999s
func (c *EthMultiClient) childTimeout(ctx context.Context) time.Duration {
	deadline, ok := ctx.Deadline()
	if !ok {
		deadline = time.Now().Add(c.timeout)
	}
	res := time.Until(deadline) / time.Duration(len(c.clients))

	// avoid errors if time is actually less
	if res > time.Millisecond {
		res = res.Round(time.Millisecond)
	}

	return res

}
