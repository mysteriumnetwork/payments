package transaction

import (
	"errors"
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
)

func TestDepotReceiptChecker(t *testing.T) {
	senderAddr := common.Address{}
	defaultPrice := big.NewInt(1)

	newDepot := func(rc DepotReceiptChecker, grace time.Duration) (*Depot, *mockStorage) {
		storage := &mockStorage{deliveries: []Delivery{}}
		nonces := &mockNonceTracker{nonces: make(map[string]uint64), confirmAll: true}
		gasTracker := NewGasTracker(&mockGasStation{defaultPrice: defaultPrice, defaultBaseFee: defaultPrice}, map[int64]GasIncreaseOpts{
			chainId: {Multiplier: 1.1, PriceLimit: big.NewInt(1000), IncreaseInterval: time.Second},
		}, GasTrackerSpeedMedium)
		d := NewDepot(&mockCourier{lastDeliveredNonce: -1}, storage, nonces, gasTracker, DepotConfig{
			MaxNonDelivered: 5,
			Workers: []DepotWorker{{
				Address:         senderAddr,
				ChainID:         chainId,
				ProcessInterval: 10 * time.Millisecond,
				ProcessCount:    3,
			}},
		})
		if rc != nil {
			d.AttachReceiptChecker(rc, grace)
		}
		return d, storage
	}

	enqueue := func(t *testing.T, d *Depot) {
		_, err := d.EnqueueDelivery(DeliveryRequest{
			ChainID: chainId,
			Sender:  senderAddr,
			Type:    "test",
			Data:    mockData{"tx"},
		}, false)
		assert.NoError(t, err)
	}

	stateIs := func(s *mockStorage, state DeliveryState) func() bool {
		return func() bool { return s.length() == 1 && DeliveryState(s.get(0).State) == state }
	}

	t.Run("marks delivered when receipt succeeded", func(t *testing.T) {
		d, s := newDepot(&mockReceipts{status: types.ReceiptStatusSuccessful}, 0)
		d.Run()
		defer d.Stop()

		enqueue(t, d)
		assert.Eventually(t, stateIs(s, DeliveryStateDelivered), 2*time.Second, 20*time.Millisecond)
	})

	t.Run("marks reverted when receipt failed", func(t *testing.T) {
		d, s := newDepot(&mockReceipts{status: types.ReceiptStatusFailed}, 0)
		d.Run()
		defer d.Stop()

		enqueue(t, d)
		assert.Eventually(t, stateIs(s, DeliveryStateReverted), 2*time.Second, 20*time.Millisecond)
	})

	t.Run("assumes delivered when receipt missing after grace", func(t *testing.T) {
		d, s := newDepot(&mockReceipts{err: ethereum.NotFound}, 0)
		d.Run()
		defer d.Stop()

		enqueue(t, d)
		assert.Eventually(t, stateIs(s, DeliveryStateDelivered), 2*time.Second, 20*time.Millisecond)
	})

	t.Run("keeps tracking when receipt missing within grace", func(t *testing.T) {
		rc := &mockReceipts{err: ethereum.NotFound}
		d, s := newDepot(rc, time.Hour)
		d.Run()
		defer d.Stop()

		enqueue(t, d)
		assert.Eventually(t, func() bool { return rc.getCalls() > 2 }, 2*time.Second, 20*time.Millisecond)
		assert.Equal(t, DeliveryState(DeliveryStateSent), s.get(0).State)
	})

	t.Run("keeps tracking on receipt lookup error", func(t *testing.T) {
		rc := &mockReceipts{err: errors.New("rpc down")}
		d, s := newDepot(rc, 0)
		d.Run()
		defer d.Stop()

		enqueue(t, d)
		assert.Eventually(t, func() bool { return rc.getCalls() > 2 }, 2*time.Second, 20*time.Millisecond)
		assert.Equal(t, DeliveryState(DeliveryStateSent), s.get(0).State)
	})

	t.Run("no checker keeps old behaviour", func(t *testing.T) {
		d, s := newDepot(nil, 0)
		d.Run()
		defer d.Stop()

		enqueue(t, d)
		assert.Eventually(t, stateIs(s, DeliveryStateDelivered), 2*time.Second, 20*time.Millisecond)
	})
}

type mockReceipts struct {
	lock   sync.Mutex
	status uint64
	err    error
	calls  int
}

func (m *mockReceipts) TransactionReceipt(chainID int64, hash common.Hash) (*types.Receipt, error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.calls++
	if m.err != nil {
		return nil, m.err
	}
	return &types.Receipt{Status: m.status, TxHash: hash}, nil
}

func (m *mockReceipts) getCalls() int {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.calls
}
