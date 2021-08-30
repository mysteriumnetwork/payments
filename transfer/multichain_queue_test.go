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

package transfer

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
)

func TestMultichainQueue(t *testing.T) {
	t.Run("green path", func(t *testing.T) {
		mq := NewMultichainQueue()
		q1 := NewMockQueue()
		q2 := NewMockQueue()
		mq.AddQueue(1, q1)
		mq.AddQueue(2, q2)

		fn1 := func() (*types.Transaction, error) {
			return types.NewTransaction(1, common.Address{}, big.NewInt(1), 10, big.NewInt(1), []byte{}), nil
		}
		resp1 := make(chan QueueResponse, 1)

		fn2 := func() (*types.Transaction, error) {
			return types.NewTransaction(2, common.Address{}, big.NewInt(1), 10, big.NewInt(1), []byte{}), nil
		}
		resp2 := make(chan QueueResponse, 1)

		mq.TransactionEnqueue(1, fn1, resp1)
		mq.TransactionEnqueue(2, fn2, resp2)

		got1 := <-resp1
		got2 := <-resp2
		tx1, err1 := got1.Result()
		tx2, err2 := got2.Result()
		assert.NoError(t, err1)
		assert.NotNil(t, tx1)
		assert.Equal(t, uint64(0x1), tx1.Nonce())
		assert.Equal(t, tx1.Nonce(), q1.lastTx.Nonce())
		assert.NoError(t, err2)
		assert.NotNil(t, tx2)
		assert.Equal(t, uint64(0x2), tx2.Nonce())
		assert.Equal(t, tx2.Nonce(), q2.lastTx.Nonce())
	})
	t.Run("chain does not exist", func(t *testing.T) {
		mq := NewMultichainQueue()
		q1 := NewQueue(1)
		q2 := NewQueue(1)
		mq.AddQueue(1, q1)
		mq.AddQueue(2, q2)
		go q1.Run()
		defer q1.Stop()
		go q2.Run()
		defer q2.Stop()
		fn := func() (*types.Transaction, error) {
			return types.NewTransaction(1, common.Address{}, big.NewInt(1), 10, big.NewInt(1), []byte{}), nil
		}
		resp := make(chan QueueResponse, 0)
		err := mq.TransactionEnqueue(3, fn, resp)
		assert.ErrorIs(t, err, ErrMissingQueueForChainId)
		assert.Len(t, resp, 0)
	})
}

type QueueMock struct {
	lastTx *types.Transaction
}

func NewMockQueue() *QueueMock {
	return &QueueMock{
		lastTx: nil,
	}
}

func (mock *QueueMock) TransactionEnqueue(exec TransactionSendFn, resp chan<- QueueResponse) error {
	tx, err := exec()
	resp <- QueueResponse{
		tx:  tx,
		err: err,
	}
	mock.lastTx = tx
	close(resp)
	return nil
}
