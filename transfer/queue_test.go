package transfer

import (
	"errors"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	t.Run("green path", func(t *testing.T) {
		q := NewQueue(1, 0)
		go q.Run()
		defer q.Stop()
		fn := func() (*types.Transaction, error) {
			return types.NewTransaction(1, common.Address{}, big.NewInt(1), 10, big.NewInt(1), []byte{}), nil
		}
		resp := make(chan QueueResponse, 0)
		err := q.TransactionEnqueue(fn, resp)

		got := <-resp
		tx, err := got.Result()
		assert.NoError(t, err)
		assert.NotNil(t, tx)
		assert.Equal(t, uint64(0x1), tx.Nonce())
	})
	t.Run("exec produced an error", func(t *testing.T) {
		q := NewQueue(1, 0)
		go q.Run()
		defer q.Stop()
		fn := func() (*types.Transaction, error) {
			return nil, errors.New(":(")
		}
		resp := make(chan QueueResponse, 0)
		err := q.TransactionEnqueue(fn, resp)

		got := <-resp
		tx, err := got.Result()
		assert.Error(t, err)
		assert.Nil(t, tx)
	})
	t.Run("closed queue returns an error", func(t *testing.T) {
		q := NewQueue(1, 0)
		q.Stop()
		fn := func() (*types.Transaction, error) {
			return types.NewTransaction(1, common.Address{}, big.NewInt(1), 10, big.NewInt(1), []byte{}), nil
		}
		resp := make(chan QueueResponse, 0)
		err := q.TransactionEnqueue(fn, resp)
		assert.ErrorIs(t, err, ErrQueueClosed)
		assert.Len(t, resp, 0)
	})
}
