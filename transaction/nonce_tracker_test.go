package transaction

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/mysteriumnetwork/payments/client"
	"github.com/mysteriumnetwork/payments/client/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNonceTracker(t *testing.T) {
	var nt *NonceTracker

	cl := &mocks.EtherClientMock{
		PendingNonceAtFunc: func(ctx context.Context, address common.Address) (uint64, error) {
			return 0, nil
		},
	}
	getter := client.NewDefaultAddressableEthClientGetter("", cl)
	mbc := client.NewMultichainBlockchainClient(map[int64]client.BC{
		1: client.NewBlockchain(getter, time.Second),
	})
	mockStorage := &mockStorage{
		deliveries: []Delivery{},
	}

	t.Run("new", func(t *testing.T) {
		nt = NewNonceTracker(mbc, mockStorage)
	})

	t.Run("set next nonce", func(t *testing.T) {
		var nonce uint64 = 1000
		//should set 0 if no info
		nt.SetNextNonce(1, common.HexToAddress("0x1"), func(n uint64) error {
			nonce = n
			return nil
		})
		assert.Equal(t, 0, int(nonce))

		//should assign next
		nt.SetNextNonce(1, common.HexToAddress("0x1"), func(n uint64) error {
			nonce = n
			return nil
		})
		assert.Equal(t, 1, int(nonce))

		//should set pending nonce
		cl.PendingNonceAtFunc = func(ctx context.Context, address common.Address) (uint64, error) {
			return 30, nil
		}
		nt.SetNextNonce(1, common.HexToAddress("0x2"), func(n uint64) error {
			nonce = n
			return nil
		})
		assert.Equal(t, 30, int(nonce))

		//should set highest of pending and queued+1
		sender := common.HexToAddress("0x3")
		mockStorage.deliveries = append(mockStorage.deliveries, Delivery{Sender: sender, ChainID: 1, Nonce: 40})
		nt.SetNextNonce(1, sender, func(n uint64) error {
			nonce = n
			return nil
		})
		assert.Equal(t, 41, int(nonce))

		//should assign next
		nt.SetNextNonce(1, sender, func(n uint64) error {
			nonce = n
			return nil
		})
		assert.Equal(t, 42, int(nonce))

		nt.ForceReloadNonce(1, sender)

		//should assign hgihest again
		nt.SetNextNonce(1, sender, func(n uint64) error {
			nonce = n
			return nil
		})
		assert.Equal(t, 41, int(nonce))
	})

	t.Run("confirmed", func(t *testing.T) {
		cl.NonceAtFunc = func(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
			return 42, nil
		}

		sender := common.HexToAddress("0x4")
		nonce, err := nt.GetConfirmedNonce(1, sender)
		assert.NoError(t, err)
		assert.Equal(t, 42, int(nonce))

		cl.NonceAtFunc = func(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
			return 23, nil
		}
		nonce, err = nt.GetConfirmedNonce(1, sender)
		assert.NoError(t, err)
		assert.Equal(t, 23, int(nonce))
	})
}
