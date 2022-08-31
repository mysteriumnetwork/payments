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
	"fmt"
	"math/big"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/mysteriumnetwork/payments/client/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_EthMultiClient(t *testing.T) {
	t.Run("no clients produces an error", func(t *testing.T) {
		_, err := NewEthMultiClient(time.Second, []AddressableEthClientGetter{})
		assert.Error(t, err)
	})

	t.Run("only one client passed is used", func(t *testing.T) {
		// https://github.com/matryer/moq
		// moq -out ./client/mocks/etherclient.go  ./client EtherClient
		cl := &mocks.EtherClientMock{
			ChainIDFunc: func(ctx context.Context) (*big.Int, error) {
				return big.NewInt(1), nil
			},
		}
		getter := NewDefaultAddressableEthClientGetter("", cl)
		multi, err := NewEthMultiClient(time.Second, []AddressableEthClientGetter{getter})
		assert.NoError(t, err)

		ethc := multi.Client()
		assert.NotNil(t, ethc)

		ctx, cancel := context.WithTimeout(context.TODO(), time.Second*2)
		defer cancel()

		chainID, err := multi.ChainID(ctx)
		assert.NoError(t, err)
		assert.True(t, chainID.Cmp(big.NewInt(1)) == 0)
		assert.True(t, len(cl.ChainIDCalls()) == 1)
	})

	t.Run("client closes", func(t *testing.T) {
		closeCalls := atomic.Uint64{}
		cl := &mocks.EtherClientMock{
			CloseFunc: func() {
				closeCalls.Add(1)
			},
		}
		getter := NewDefaultAddressableEthClientGetter("", cl)
		multi, err := NewEthMultiClient(time.Second, []AddressableEthClientGetter{getter})
		assert.NoError(t, err)

		multi.Close()
		assert.Equal(t, uint64(1), closeCalls.Load())
	})

	t.Run("multiple clients close", func(t *testing.T) {
		closeCalls := atomic.Int64{}
		cl := &mocks.EtherClientMock{
			CloseFunc: func() {
				closeCalls.Add(1)
			},
		}
		cl2 := &mocks.EtherClientMock{
			CloseFunc: func() {
				closeCalls.Add(1)
			},
		}
		getter := NewDefaultAddressableEthClientGetter("", cl)
		getter2 := NewDefaultAddressableEthClientGetter("", cl2)
		multi, err := NewEthMultiClient(time.Second, []AddressableEthClientGetter{getter, getter2})
		assert.NoError(t, err)

		multi.Close()
		assert.Equal(t, int64(2), closeCalls.Load())
	})

	t.Run("calls correct function once", func(t *testing.T) {
		calls := make(map[string]int)
		callsLock := sync.Mutex{}
		cl := getMockClientWithFunctionCallCounter(calls, &callsLock)

		getter := NewDefaultAddressableEthClientGetter("", cl)
		multi, err := NewEthMultiClient(time.Second, []AddressableEthClientGetter{getter})
		assert.NoError(t, err)

		t.Run("calls all functions without errors", func(t *testing.T) {
			_, err = multi.ChainID(context.Background())
			assert.NoError(t, err)
			_, err = multi.BlockByHash(context.Background(), common.Hash{})
			assert.NoError(t, err)
			_, err = multi.BlockByNumber(context.Background(), nil)
			assert.NoError(t, err)
			_, err = multi.BlockNumber(context.Background())
			assert.NoError(t, err)
			_, err = multi.HeaderByHash(context.Background(), common.Hash{})
			assert.NoError(t, err)
			_, err = multi.HeaderByNumber(context.Background(), nil)
			assert.NoError(t, err)
			_, _, err = multi.TransactionByHash(context.Background(), common.Hash{})
			assert.NoError(t, err)
			_, err = multi.TransactionSender(context.Background(), nil, common.Hash{}, 0)
			assert.NoError(t, err)
			_, err = multi.TransactionCount(context.Background(), common.Hash{})
			assert.NoError(t, err)
			_, err = multi.TransactionInBlock(context.Background(), common.Hash{}, 0)
			assert.NoError(t, err)
			_, err = multi.TransactionReceipt(context.Background(), common.Hash{})
			assert.NoError(t, err)
			_, err = multi.SyncProgress(context.Background())
			assert.NoError(t, err)
			_, err = multi.SubscribeNewHead(context.Background(), nil)
			assert.NoError(t, err)
			_, err = multi.NetworkID(context.Background())
			assert.NoError(t, err)
			_, err = multi.BalanceAt(context.Background(), common.Address{}, nil)
			assert.NoError(t, err)
			_, err = multi.StorageAt(context.Background(), common.Address{}, common.Hash{}, nil)
			assert.NoError(t, err)
			_, err = multi.CodeAt(context.Background(), common.Address{}, nil)
			assert.NoError(t, err)
			_, err = multi.NonceAt(context.Background(), common.Address{}, nil)
			assert.NoError(t, err)
			_, err = multi.FilterLogs(context.Background(), ethereum.FilterQuery{})
			assert.NoError(t, err)
			_, err = multi.SubscribeFilterLogs(context.Background(), ethereum.FilterQuery{}, nil)
			assert.NoError(t, err)
			_, err = multi.PendingBalanceAt(context.Background(), common.Address{})
			assert.NoError(t, err)
			_, err = multi.PendingStorageAt(context.Background(), common.Address{}, common.Hash{})
			assert.NoError(t, err)
			_, err = multi.PendingCodeAt(context.Background(), common.Address{})
			assert.NoError(t, err)
			_, err = multi.PendingNonceAt(context.Background(), common.Address{})
			assert.NoError(t, err)
			_, err = multi.PendingTransactionCount(context.Background())
			assert.NoError(t, err)
			_, err = multi.CallContract(context.Background(), ethereum.CallMsg{}, nil)
			assert.NoError(t, err)
			_, err = multi.PendingCallContract(context.Background(), ethereum.CallMsg{})
			assert.NoError(t, err)
			_, err = multi.SuggestGasPrice(context.Background())
			assert.NoError(t, err)
			_, err = multi.EstimateGas(context.Background(), ethereum.CallMsg{})
			assert.NoError(t, err)
			_, err = multi.SuggestGasTipCap(context.Background())
			assert.NoError(t, err)
			err = multi.SendTransaction(context.Background(), &types.Transaction{})
			assert.NoError(t, err)

			multi.Close()
		})

		callsLock.Lock()
		defer callsLock.Unlock()
		assert.Len(t, calls, 32)
		for f, c := range calls {
			if c != 1 {
				assert.Fail(t, "%s should be called only once", f)
			}
		}
	})

	t.Run("two client but calls correct function once (except close)", func(t *testing.T) {
		calls := make(map[string]int)
		callsLock := sync.Mutex{}
		cl := getMockClientWithFunctionCallCounter(calls, &callsLock)

		cl2 := getMockClientWithFunctionCallCounter(calls, &callsLock)

		getter := NewDefaultAddressableEthClientGetter("", cl)
		getter2 := NewDefaultAddressableEthClientGetter("", cl2)
		multi, err := NewEthMultiClient(time.Second, []AddressableEthClientGetter{getter, getter2})
		assert.NoError(t, err)

		t.Run("calls all functions without errors", func(t *testing.T) {
			_, err = multi.ChainID(context.Background())
			assert.NoError(t, err)
			_, err = multi.BlockByHash(context.Background(), common.Hash{})
			assert.NoError(t, err)
			_, err = multi.BlockByNumber(context.Background(), nil)
			assert.NoError(t, err)
			_, err = multi.BlockNumber(context.Background())
			assert.NoError(t, err)
			_, err = multi.HeaderByHash(context.Background(), common.Hash{})
			assert.NoError(t, err)
			_, err = multi.HeaderByNumber(context.Background(), nil)
			assert.NoError(t, err)
			_, _, err = multi.TransactionByHash(context.Background(), common.Hash{})
			assert.NoError(t, err)
			_, err = multi.TransactionSender(context.Background(), nil, common.Hash{}, 0)
			assert.NoError(t, err)
			_, err = multi.TransactionCount(context.Background(), common.Hash{})
			assert.NoError(t, err)
			_, err = multi.TransactionInBlock(context.Background(), common.Hash{}, 0)
			assert.NoError(t, err)
			_, err = multi.TransactionReceipt(context.Background(), common.Hash{})
			assert.NoError(t, err)
			_, err = multi.SyncProgress(context.Background())
			assert.NoError(t, err)
			_, err = multi.SubscribeNewHead(context.Background(), nil)
			assert.NoError(t, err)
			_, err = multi.NetworkID(context.Background())
			assert.NoError(t, err)
			_, err = multi.BalanceAt(context.Background(), common.Address{}, nil)
			assert.NoError(t, err)
			_, err = multi.StorageAt(context.Background(), common.Address{}, common.Hash{}, nil)
			assert.NoError(t, err)
			_, err = multi.CodeAt(context.Background(), common.Address{}, nil)
			assert.NoError(t, err)
			_, err = multi.NonceAt(context.Background(), common.Address{}, nil)
			assert.NoError(t, err)
			_, err = multi.FilterLogs(context.Background(), ethereum.FilterQuery{})
			assert.NoError(t, err)
			_, err = multi.SubscribeFilterLogs(context.Background(), ethereum.FilterQuery{}, nil)
			assert.NoError(t, err)
			_, err = multi.PendingBalanceAt(context.Background(), common.Address{})
			assert.NoError(t, err)
			_, err = multi.PendingStorageAt(context.Background(), common.Address{}, common.Hash{})
			assert.NoError(t, err)
			_, err = multi.PendingCodeAt(context.Background(), common.Address{})
			assert.NoError(t, err)
			_, err = multi.PendingNonceAt(context.Background(), common.Address{})
			assert.NoError(t, err)
			_, err = multi.PendingTransactionCount(context.Background())
			assert.NoError(t, err)
			_, err = multi.CallContract(context.Background(), ethereum.CallMsg{}, nil)
			assert.NoError(t, err)
			_, err = multi.PendingCallContract(context.Background(), ethereum.CallMsg{})
			assert.NoError(t, err)
			_, err = multi.SuggestGasPrice(context.Background())
			assert.NoError(t, err)
			_, err = multi.EstimateGas(context.Background(), ethereum.CallMsg{})
			assert.NoError(t, err)
			_, err = multi.SuggestGasTipCap(context.Background())
			assert.NoError(t, err)
			err = multi.SendTransaction(context.Background(), &types.Transaction{})
			assert.NoError(t, err)

			multi.Close()
		})

		callsLock.Lock()
		defer callsLock.Unlock()
		assert.Len(t, calls, 32)
		for f, c := range calls {
			if f == "Close" {
				if c != 2 {
					assert.Fail(t, fmt.Sprintf("%s should be called only twice", f))
				}
				continue
			}
			if c != 1 {
				assert.Fail(t, fmt.Sprintf("%s should be called only once", f))
			}
		}
	})

	t.Run("context with no timeout gets one assigned", func(t *testing.T) {
		cl := &mocks.EtherClientMock{
			ChainIDFunc: func(ctx context.Context) (*big.Int, error) {
				ded, ok := ctx.Deadline()
				assert.True(t, ok)
				assert.NotZero(t, ded)
				assert.True(t, time.Until(ded) > time.Second)

				return big.NewInt(1), nil
			},
		}
		getter := NewDefaultAddressableEthClientGetter("", cl)
		multi, err := NewEthMultiClient(time.Second*2, []AddressableEthClientGetter{getter})
		assert.NoError(t, err)

		chainID, err := multi.ChainID(context.TODO())
		assert.NoError(t, err)
		assert.True(t, chainID.Cmp(big.NewInt(1)) == 0)
		assert.True(t, len(cl.ChainIDCalls()) == 1)
	})

	t.Run("two clients passed first one works, second never called", func(t *testing.T) {
		cl := &mocks.EtherClientMock{
			ChainIDFunc: func(ctx context.Context) (*big.Int, error) {
				return big.NewInt(1), nil
			},
		}
		cl2 := &mocks.EtherClientMock{
			ChainIDFunc: func(ctx context.Context) (*big.Int, error) {
				return big.NewInt(1), nil
			},
		}
		getter := NewDefaultAddressableEthClientGetter("", cl)
		getter2 := NewDefaultAddressableEthClientGetter("", cl2)

		multi, err := NewEthMultiClient(time.Second, []AddressableEthClientGetter{getter, getter2})
		assert.NoError(t, err)

		ctx, cancel := context.WithTimeout(context.TODO(), time.Second*2)
		defer cancel()

		chainID, err := multi.ChainID(ctx)
		assert.NoError(t, err)
		assert.True(t, chainID.Cmp(big.NewInt(1)) == 0)
		assert.True(t, len(cl.ChainIDCalls()) == 1)
		assert.True(t, len(cl2.ChainIDCalls()) == 0)
	})

	t.Run("two clients passed first breaks, second is used", func(t *testing.T) {
		cl := &mocks.EtherClientMock{
			ChainIDFunc: func(ctx context.Context) (*big.Int, error) {
				time.Sleep(time.Second / 6)
				return big.NewInt(1), ctx.Err()
			},
		}
		cl2 := &mocks.EtherClientMock{
			ChainIDFunc: func(ctx context.Context) (*big.Int, error) {
				return big.NewInt(2), nil
			},
		}
		getter := NewDefaultAddressableEthClientGetter("", cl)
		getter2 := NewDefaultAddressableEthClientGetter("", cl2)

		multi, err := NewEthMultiClient(time.Second, []AddressableEthClientGetter{getter, getter2})
		assert.NoError(t, err)

		ctx, cancel := context.WithTimeout(context.TODO(), time.Second/4)
		defer cancel()

		chainID, err := multi.ChainID(ctx)
		assert.NoError(t, err)
		assert.True(t, chainID.Cmp(big.NewInt(2)) == 0)
		assert.True(t, len(cl.ChainIDCalls()) == 1)
		assert.True(t, len(cl2.ChainIDCalls()) == 1)
	})

	t.Run("two clients passed first breaks, notification received", func(t *testing.T) {
		cl := &mocks.EtherClientMock{
			ChainIDFunc: func(ctx context.Context) (*big.Int, error) {
				time.Sleep(time.Second / 6)
				return big.NewInt(1), ctx.Err()
			},
		}
		cl2 := &mocks.EtherClientMock{
			ChainIDFunc: func(ctx context.Context) (*big.Int, error) {
				return big.NewInt(2), nil
			},
		}
		getter := NewDefaultAddressableEthClientGetter("first", cl)
		getter2 := NewDefaultAddressableEthClientGetter("second", cl2)

		notificationReceived := make(chan struct{})
		notify := make(chan Notification)
		testCtx, cancel := context.WithCancel(context.TODO())
		defer cancel()
		go func() {
			select {
			case <-testCtx.Done():
				return
			case result := <-notify:
				notificationReceived <- struct{}{}
				assert.Equal(t, "first", result.Address)
			}
		}()

		multi, err := NewEthMultiClientNotifyDown(time.Second, []AddressableEthClientGetter{getter, getter2}, notify)
		assert.NoError(t, err)

		ctx, cancel := context.WithTimeout(context.TODO(), time.Second/4)
		defer cancel()

		chainID, err := multi.ChainID(ctx)
		assert.NoError(t, err)
		assert.True(t, chainID.Cmp(big.NewInt(2)) == 0)
		assert.True(t, len(cl.ChainIDCalls()) == 1)
		assert.True(t, len(cl2.ChainIDCalls()) == 1)
		assert.Eventually(t, func() bool {
			select {
			case <-notificationReceived:
				return true
			default:
				return false
			}
		}, time.Second/2, time.Second/200)
	})

	t.Run("two clients, callin specific one works", func(t *testing.T) {
		cl := &mocks.EtherClientMock{
			BlockNumberFunc: func(ctx context.Context) (uint64, error) {
				return 1, nil
			},
		}
		cl2 := &mocks.EtherClientMock{
			BlockNumberFunc: func(ctx context.Context) (uint64, error) {
				return 2, nil
			},
		}
		getter := NewDefaultAddressableEthClientGetter("first", cl)
		getter2 := NewDefaultAddressableEthClientGetter("second", cl2)

		multi, err := NewEthMultiClient(time.Second, []AddressableEthClientGetter{getter, getter2})
		assert.NoError(t, err)

		var blockNumber uint64
		multi.CallSpecificClient("second", func(c EtherClient) error {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
			defer cancel()

			bn, err := c.BlockNumber(ctx)
			if err != nil {
				return err
			}

			blockNumber = bn
			return nil
		})

		assert.Equal(t, 2, int(blockNumber), "blocknumber should be received from second client")
	})

	t.Run("multiple clients sorting", func(t *testing.T) {
		cl := &mocks.EtherClientMock{
			ChainIDFunc: func(ctx context.Context) (*big.Int, error) {
				return big.NewInt(1), nil
			},
		}

		cl2 := &mocks.EtherClientMock{
			ChainIDFunc: func(ctx context.Context) (*big.Int, error) {
				return big.NewInt(2), nil
			},
		}

		cl3 := &mocks.EtherClientMock{
			ChainIDFunc: func(ctx context.Context) (*big.Int, error) {
				return big.NewInt(2), nil
			},
		}
		getter := NewDefaultAddressableEthClientGetter("first", cl)
		getter2 := NewDefaultAddressableEthClientGetter("second", cl2)
		getter3 := NewDefaultAddressableEthClientGetter("third", cl3)

		multi, err := NewEthMultiClient(time.Second*2, []AddressableEthClientGetter{getter, getter2, getter3})
		assert.NoError(t, err)

		order := multi.CurrentClientOrder()
		assert.Equal(t, []string{"first", "second", "third"}, order)

		err = multi.ReorderClients([]string{"first"})
		assert.Error(t, err)

		err = multi.ReorderClients([]string{"first", "second", "wrong"})
		assert.Error(t, err)

		err = multi.ReorderClients([]string{"second", "first", "third"})
		assert.NoError(t, err)

		order = multi.CurrentClientOrder()
		assert.Equal(t, []string{"second", "first", "third"}, order)
	})
}

func getMockClientWithFunctionCallCounter(counterMap map[string]int, mapLock *sync.Mutex) *mocks.EtherClientMock {
	return &mocks.EtherClientMock{
		CloseFunc: func() {
			lockAndIncrement(counterMap, mapLock, "Close")
		},
		ChainIDFunc: func(ctx context.Context) (*big.Int, error) {
			lockAndIncrement(counterMap, mapLock, "ChainID")
			return nil, nil
		},
		BlockByHashFunc: func(ctx context.Context, hash common.Hash) (*types.Block, error) {
			lockAndIncrement(counterMap, mapLock, "BlockByHash")
			return nil, nil
		},
		BlockByNumberFunc: func(ctx context.Context, number *big.Int) (*types.Block, error) {
			lockAndIncrement(counterMap, mapLock, "BlockByNumber")
			return nil, nil
		},
		BlockNumberFunc: func(ctx context.Context) (uint64, error) {
			lockAndIncrement(counterMap, mapLock, "BlockNumber")
			return 0, nil
		},
		HeaderByHashFunc: func(ctx context.Context, hash common.Hash) (*types.Header, error) {
			lockAndIncrement(counterMap, mapLock, "HeaderByHash")
			return nil, nil
		},
		HeaderByNumberFunc: func(ctx context.Context, number *big.Int) (*types.Header, error) {
			lockAndIncrement(counterMap, mapLock, "HeaderByNumber")
			return nil, nil
		},
		TransactionByHashFunc: func(ctx context.Context, hash common.Hash) (*types.Transaction, bool, error) {
			lockAndIncrement(counterMap, mapLock, "TransactionByHash")
			return nil, false, nil
		},
		TransactionSenderFunc: func(ctx context.Context, tx *types.Transaction, block common.Hash, index uint) (common.Address, error) {
			lockAndIncrement(counterMap, mapLock, "TransactionSender")
			return common.Address{}, nil
		},
		TransactionCountFunc: func(ctx context.Context, blockHash common.Hash) (uint, error) {
			lockAndIncrement(counterMap, mapLock, "TransactionCount")
			return 0, nil
		},
		TransactionInBlockFunc: func(ctx context.Context, blockHash common.Hash, index uint) (*types.Transaction, error) {
			lockAndIncrement(counterMap, mapLock, "TransactionInBlock")
			return nil, nil
		},
		TransactionReceiptFunc: func(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
			lockAndIncrement(counterMap, mapLock, "TransactionReceipt")
			return nil, nil
		},
		SyncProgressFunc: func(ctx context.Context) (*ethereum.SyncProgress, error) {
			lockAndIncrement(counterMap, mapLock, "SyncProgress")
			return nil, nil
		},
		SubscribeNewHeadFunc: func(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
			lockAndIncrement(counterMap, mapLock, "SubscribeNewHead")
			return nil, nil
		},
		NetworkIDFunc: func(ctx context.Context) (*big.Int, error) {
			lockAndIncrement(counterMap, mapLock, "NetworkID")
			return nil, nil
		},
		BalanceAtFunc: func(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
			lockAndIncrement(counterMap, mapLock, "BalanceAt")
			return nil, nil
		},
		StorageAtFunc: func(ctx context.Context, account common.Address, key common.Hash, blockNumber *big.Int) ([]byte, error) {
			lockAndIncrement(counterMap, mapLock, "StorageAt")
			return nil, nil
		},
		CodeAtFunc: func(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
			lockAndIncrement(counterMap, mapLock, "CodeAt")
			return nil, nil
		},
		NonceAtFunc: func(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
			lockAndIncrement(counterMap, mapLock, "NonceAt")
			return 0, nil
		},
		FilterLogsFunc: func(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
			lockAndIncrement(counterMap, mapLock, "FilterLogs")
			return nil, nil
		},
		SubscribeFilterLogsFunc: func(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
			lockAndIncrement(counterMap, mapLock, "SubscribeFilterLogs")
			return nil, nil
		},
		PendingBalanceAtFunc: func(ctx context.Context, account common.Address) (*big.Int, error) {
			lockAndIncrement(counterMap, mapLock, "PendingBalanceAt")
			return nil, nil
		},
		PendingStorageAtFunc: func(ctx context.Context, account common.Address, key common.Hash) ([]byte, error) {
			lockAndIncrement(counterMap, mapLock, "PendingStorageAt")
			return nil, nil
		},
		PendingCodeAtFunc: func(ctx context.Context, account common.Address) ([]byte, error) {
			lockAndIncrement(counterMap, mapLock, "PendingCodeAt")
			return nil, nil
		},
		PendingNonceAtFunc: func(ctx context.Context, account common.Address) (uint64, error) {
			lockAndIncrement(counterMap, mapLock, "PendingNonceAt")
			return 0, nil
		},
		PendingTransactionCountFunc: func(ctx context.Context) (uint, error) {
			lockAndIncrement(counterMap, mapLock, "PendingTransactionCount")
			return 0, nil
		},
		CallContractFunc: func(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
			lockAndIncrement(counterMap, mapLock, "CallContract")
			return nil, nil
		},
		PendingCallContractFunc: func(ctx context.Context, msg ethereum.CallMsg) ([]byte, error) {
			lockAndIncrement(counterMap, mapLock, "PendingCallContract")
			return nil, nil
		},
		SuggestGasPriceFunc: func(ctx context.Context) (*big.Int, error) {
			lockAndIncrement(counterMap, mapLock, "SuggestGasPrice")
			return nil, nil
		},
		EstimateGasFunc: func(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
			lockAndIncrement(counterMap, mapLock, "EstimateGas")
			return 0, nil
		},
		SuggestGasTipCapFunc: func(ctx context.Context) (*big.Int, error) {
			lockAndIncrement(counterMap, mapLock, "SuggestGasTipCap")
			return nil, nil
		},
		SendTransactionFunc: func(ctx context.Context, tx *types.Transaction) error {
			lockAndIncrement(counterMap, mapLock, "SendTransaction")
			return nil
		},
	}
}

func lockAndIncrement(callsCounter map[string]int, lock *sync.Mutex, key string) {
	lock.Lock()
	defer lock.Unlock()
	callsCounter[key]++
}
