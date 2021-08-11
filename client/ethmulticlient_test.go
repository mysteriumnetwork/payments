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
	"math/big"
	"testing"
	"time"

	"github.com/mysteriumnetwork/payments/client/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_EthMultiClient(t *testing.T) {
	t.Run("no clients produces an error", func(t *testing.T) {
		_, err := NewEthMultiClient(time.Second, []AddressableEthClientGetter{})
		assert.Error(t, err)
	})

	t.Run("only one client passed is used", func(t *testing.T) {
		cl := &mocks.EtherClientMock{
			ChainIDFunc: func(ctx context.Context) (*big.Int, error) {
				return big.NewInt(1), nil
			},
		}
		getter := NewDefaultAddressableEthClientGetter("", cl)
		multi, err := NewEthMultiClient(time.Second, []AddressableEthClientGetter{getter})
		assert.NoError(t, err)

		ctx, cancel := context.WithTimeout(context.TODO(), time.Second*2)
		defer cancel()

		chainID, err := multi.ChainID(ctx)
		assert.NoError(t, err)
		assert.True(t, chainID.Cmp(big.NewInt(1)) == 0)
		assert.True(t, len(cl.ChainIDCalls()) == 1)
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
		notify := make(chan Notification, 0)
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
