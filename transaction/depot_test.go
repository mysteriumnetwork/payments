package transaction

import (
	"fmt"
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/mysteriumnetwork/payments/transaction/gas"
	"github.com/stretchr/testify/assert"
)

func TestDepot(t *testing.T) {
	senderAddr := common.Address{}
	mockStorage := mockStorage{
		deliveries: []Delivery{},
	}
	mockNonceTracker := mockNonceTracker{
		nonces:     make(map[string]uint64),
		confirmAll: true,
		lock:       sync.Mutex{},
	}
	gasTipMultiplier := 1.1
	defaultPrice := big.NewInt(1)
	mockGasStation := mockGasStation{
		defaultPrice:   defaultPrice,
		defaultBaseFee: defaultPrice,
	}
	gasTracker := NewGasTracker(&mockGasStation, map[int64]GasIncreaseOpts{
		1: {
			Multiplier:       gasTipMultiplier,
			PriceLimit:       big.NewInt(1000),
			IncreaseInterval: time.Second,
			OverpayFor:       nil,
			OverpayByMul:     0,
		},
	}, GasTrackerSpeedMedium)
	mockCourier := mockCourier{
		lastDeliveredNonce: -1,
		calls:              0,
	}
	depot := NewDepot(&mockCourier, &mockStorage, &mockNonceTracker, gasTracker, DepotConfig{
		MaxNonDelivered: 5,
		Workers: []DepotWorker{
			{
				Address:         senderAddr,
				ChainID:         1,
				ProcessInterval: time.Millisecond * 10,
				ProcessCount:    3,
			},
		},
	})

	resetFunc := func() {
		mockStorage.reset()
		mockNonceTracker.reset()
		mockCourier.reset()
		mockGasStation.reset(defaultPrice, defaultPrice)
		mockNonceTracker.setConfirmAll(true)
	}

	t.Run("delivery", func(t *testing.T) {
		t.Run("runs", func(t *testing.T) {
			depot.Run()
		})

		t.Run("green path", func(t *testing.T) {
			defer resetFunc()

			var txType DeliverableType = "test"
			_, err := depot.EnqueueDelivery(DeliveryRequest{
				ChainID: 1,
				Sender:  senderAddr,
				Type:    txType,
				Data:    mockData{"tx1"},
			}, false)
			assert.NoError(t, err)

			assert.Equal(t, 1, mockStorage.length())

			assert.Eventually(t, func() bool {
				return mockStorage.get(0).State == DeliveryStateDelivered
			}, 2*time.Second, time.Millisecond*100)

			delivery := mockStorage.get(0)
			assert.True(t, delivery.BaseFee.Cmp(defaultPrice) == 0)
			assert.True(t, delivery.GasTip.Cmp(defaultPrice) == 0)
			assert.Equal(t, 1, int(delivery.ChainID))
			assert.Equal(t, 0, int(delivery.Nonce))
			assert.Equal(t, senderAddr, delivery.Sender)
			assert.Equal(t, txType, delivery.Type)
			assert.Equal(t, []byte("{\"data\":\"tx1\"}"), delivery.ShipmentData)
			txJson, err := types.NewTx(&types.DynamicFeeTx{
				ChainID:   big.NewInt(1),
				Nonce:     0,
				GasTipCap: defaultPrice,
				GasFeeCap: defaultPrice,
				Gas:       0,
			}).MarshalJSON()
			assert.NoError(t, err)
			assert.Equal(t, txJson, delivery.SentTransaction)
		})

		t.Run("delivers all if nonce in bc is bigger", func(t *testing.T) {
			defer resetFunc()

			//add 3 deliveries
			for i := 0; i < 3; i++ {
				_, err := depot.EnqueueDelivery(DeliveryRequest{
					ChainID: 1,
					Sender:  senderAddr,
					Type:    "test",
					Data:    mockData{fmt.Sprintf("tx%d", i)},
				}, false)
				assert.NoError(t, err)
			}

			assert.Equal(t, 3, mockStorage.length())
			assert.Eventually(t, func() bool {
				for _, d := range mockStorage.deliveries {
					if d.State != DeliveryStateDelivered {
						return false
					}
				}
				return true
			}, 2*time.Second, time.Millisecond*100)
			assert.Equal(t, 2, int(mockCourier.getLastDeliveredNonce()))
			assert.Equal(t, 3, int(mockCourier.getCalls()))
		})

		t.Run("not delivered until nonce is not confirmed", func(t *testing.T) {
			defer resetFunc()
			mockNonceTracker.setConfirmAll(false)

			_, err := depot.EnqueueDelivery(DeliveryRequest{
				ChainID: 1,
				Sender:  senderAddr,
				Type:    "test",
				Data:    "tx1",
			}, false)
			assert.NoError(t, err)

			assert.Equal(t, 1, mockStorage.length())
			assert.Eventually(t, func() bool {
				for _, d := range mockStorage.deliveries {
					if d.State != DeliveryStateSent {
						return false
					}
				}
				return true
			}, 2*time.Second, time.Millisecond*100)
			assert.Equal(t, 0, int(mockCourier.getLastDeliveredNonce()))
			assert.Equal(t, 1, int(mockCourier.getCalls()))
			assert.True(t, mockStorage.get(0).State == DeliveryStateSent)
		})

		t.Run("increases gas after some time if not confirmed", func(t *testing.T) {
			defer resetFunc()
			mockNonceTracker.setConfirmAll(false)

			newDefaultPrice := big.NewInt(100)
			mockGasStation.defaultBaseFee = newDefaultPrice
			mockGasStation.defaultPrice = newDefaultPrice

			_, err := depot.EnqueueDelivery(DeliveryRequest{
				ChainID: 1,
				Sender:  senderAddr,
				Type:    "test",
				Data:    "tx1",
			}, false)
			assert.NoError(t, err)

			var initalDelivery Delivery
			//sends it once
			assert.Eventually(t, func() bool {
				if mockCourier.getCalls() == 1 {
					initalDelivery = mockStorage.get(0)
					return true
				}
				return false
			}, 2*time.Second, time.Millisecond*100)

			increasedGasPrice := big.NewInt(500)
			var deliveryRetry Delivery
			//sends it again
			assert.Eventually(t, func() bool {
				if mockCourier.getCalls() == 2 {
					deliveryRetry = mockStorage.get(0)
					//simulating a gas spike, in this case the multiplier would not be enough.
					mockGasStation.defaultBaseFee = increasedGasPrice
					mockGasStation.defaultPrice = increasedGasPrice
					return true
				}
				return false
			}, 2*time.Second, time.Millisecond*100)

			var deliveryFinal Delivery
			//sends it again
			assert.Eventually(t, func() bool {
				if mockCourier.getCalls() == 3 {
					deliveryFinal = mockStorage.get(0)
					//confirm it once its sent again (should not send again)
					mockNonceTracker.setConfirmAll(true)
					return true
				}
				return false
			}, 2*time.Second, time.Millisecond*100)

			assert.Equal(t, 0, int(initalDelivery.Nonce))

			expectedTipRetry, _ := new(big.Float).Mul(
				big.NewFloat(gasTipMultiplier),
				new(big.Float).SetInt(newDefaultPrice),
			).Int(nil)
			assert.True(t, expectedTipRetry.Cmp(deliveryRetry.GasTip) == 0)
			assert.Equal(t, 0, int(deliveryRetry.Nonce))

			expectedTipFinal := increasedGasPrice
			assert.True(t, expectedTipFinal.Cmp(deliveryFinal.GasTip) == 0)
			assert.Equal(t, 0, int(deliveryFinal.Nonce))
		})

		t.Run("does not allow more than max non delivered", func(t *testing.T) {
			defer resetFunc()
			mockNonceTracker.setConfirmNone(true)

			//add 5 deliveries
			for i := 0; i < 5; i++ {
				_, err := depot.EnqueueDelivery(DeliveryRequest{
					ChainID: 1,
					Sender:  senderAddr,
					Type:    "test",
					Data:    mockData{fmt.Sprintf("tx%d", i)},
				}, false)
				assert.NoError(t, err)
			}

			assert.Equal(t, 5, mockStorage.length())

			// 3 sent and 2 are waiting
			assert.Eventually(t, func() bool {
				sent := 0
				waiting := 0
				for i := 0; i < mockStorage.length(); i++ {
					delivery := mockStorage.get(i)
					if delivery.State == DeliveryStateSent {
						sent++
					} else if delivery.State == DeliveryStateWaiting {
						waiting++
					} else {
						return false
					}
				}
				if sent == 3 && waiting == 2 {
					return true
				}
				return false
			}, 2*time.Second, time.Millisecond*100)

			_, err := depot.EnqueueDelivery(DeliveryRequest{
				ChainID: 1,
				Sender:  senderAddr,
				Type:    "test",
				Data:    mockData{"tx5"},
			}, false)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "cannot queue a new entry, max count of 5 reached")

			//can force enqueue
			_, err = depot.EnqueueDelivery(DeliveryRequest{
				ChainID: 1,
				Sender:  senderAddr,
				Type:    "test",
				Data:    mockData{"tx5"},
			}, true)
			assert.NoError(t, err)
			assert.Equal(t, 6, mockStorage.length())

			// 3 sent and 3 are waiting
			assert.Eventually(t, func() bool {
				sent := 0
				waiting := 0
				for i := 0; i < mockStorage.length(); i++ {
					delivery := mockStorage.get(i)
					if delivery.State == DeliveryStateSent {
						sent++
					} else if delivery.State == DeliveryStateWaiting {
						waiting++
					} else {
						return false
					}
				}
				if sent == 3 && waiting == 3 {
					return true
				}
				return false
			}, 2*time.Second, time.Millisecond*100)

			//confirm them
			mockNonceTracker.setConfirmAll(true)

			// all get delivered
			assert.Eventually(t, func() bool {
				for i := 0; i < mockStorage.length(); i++ {
					if mockStorage.get(i).State != DeliveryStateDelivered {
						return false
					}
				}
				return true
			}, 2*time.Second, time.Millisecond*100)

			//can enqueue again
			_, err = depot.EnqueueDelivery(DeliveryRequest{
				ChainID: 1,
				Sender:  senderAddr,
				Type:    "test",
				Data:    mockData{"tx6"},
			}, false)
			assert.NoError(t, err)
		})
	})

	t.Run("cleaner", func(t *testing.T) {
		t.Run("attaches and runs cleaner", func(t *testing.T) {
			depot.AttachCleaner(&mockStorage, DepotCleanupConfig{
				CleanupInterval:  10 * time.Millisecond,
				CleanupDaysLimit: 3,
				CleanupLimit:     2,
			})
			depot.RunCleaner()
		})

		t.Run("deletes old delivered", func(t *testing.T) {
			defer resetFunc()

			//add 3 deliveries
			for i := 0; i < 3; i++ {
				_, err := depot.EnqueueDelivery(DeliveryRequest{
					ChainID: 1,
					Sender:  senderAddr,
					Type:    "test",
					Data:    mockData{fmt.Sprintf("tx%d", i)},
				}, false)
				assert.NoError(t, err)
			}

			t.Run("has 3 deliveries in db", func(t *testing.T) {
				assert.Equal(t, 3, mockStorage.length())
			})

			t.Run("all deliveries are delivered", func(t *testing.T) {
				assert.Eventually(t, func() bool {
					for _, d := range mockStorage.deliveries {
						if d.State != DeliveryStateDelivered {
							return false
						}
					}
					return true
				}, 2*time.Second, time.Millisecond*100)
			})

			//set to 8 days ago
			mockStorage.addUpdatedUtc(0, -time.Hour*192)
			//set to 3 days and 10 millis ago
			mockStorage.addUpdatedUtc(1, -time.Hour*72-time.Millisecond*10)
			//set to 2 days and 23 hours ago
			mockStorage.addUpdatedUtc(2, -time.Hour*71)

			t.Run("should clean 2 of the 3 deliveries", func(t *testing.T) {
				assert.Eventually(t, func() bool {
					return mockStorage.length() == 1
				}, time.Second, time.Millisecond*100)
				assert.Equal(t, []byte("{\"data\":\"tx2\"}"), mockStorage.deliveries[0].ShipmentData)
			})

			mockStorage.addUpdatedUtc(0, -time.Hour*1-time.Millisecond*10)

			t.Run("should clean all of the deliveries", func(t *testing.T) {
				assert.Eventually(t, func() bool {
					return mockStorage.length() == 0
				}, time.Second, time.Millisecond*100)
			})
		})

		t.Run("does not remove undelivered", func(t *testing.T) {
			defer resetFunc()
			mockNonceTracker.setConfirmAll(false)

			//add 3 deliveries
			for i := 0; i < 3; i++ {
				_, err := depot.EnqueueDelivery(DeliveryRequest{
					ChainID: 1,
					Sender:  senderAddr,
					Type:    "test",
					Data:    mockData{fmt.Sprintf("tx%d", i)},
				}, false)
				assert.NoError(t, err)
			}

			t.Run("has 3 deliveries in db", func(t *testing.T) {
				assert.Equal(t, 3, mockStorage.length())
			})

			t.Run("all deliveries except one are delivered", func(t *testing.T) {
				assert.Eventually(t, func() bool {
					undelivered := 0
					for _, d := range mockStorage.deliveries {
						if d.State != DeliveryStateDelivered {
							undelivered += 1
						}
					}
					return undelivered == 1
				}, time.Second, time.Millisecond*100)
			})

			//set all to 8 days ago
			mockStorage.addUpdatedUtc(0, -time.Hour*192)
			mockStorage.addUpdatedUtc(1, -time.Hour*192)
			mockStorage.addUpdatedUtc(2, -time.Hour*192)

			t.Run("should clean 2 of the 3 deliveries", func(t *testing.T) {
				assert.Eventually(t, func() bool {
					return mockStorage.length() == 1
				}, time.Second, time.Millisecond*100)

				assert.Equal(t, []byte("{\"data\":\"tx2\"}"), mockStorage.get(0).ShipmentData)
				assert.True(t, mockStorage.deliveries[0].State == DeliveryStateSent)
			})
		})
	})

}

type mockStorage struct {
	deliveries []Delivery
	lock       sync.Mutex
}

func (m *mockStorage) GetOrderedDeliveryRequests(count uint, chainID int64, sender common.Address) ([]Delivery, error) {
	res := []Delivery{}
	if count == 0 {
		return res, nil
	}
	var lastNonce uint64 = 0

	m.lock.Lock()
	defer m.lock.Unlock()

	for _, d := range m.deliveries {
		if d.Sender == sender && d.ChainID == chainID && d.State != DeliveryStateDelivered {
			if d.Nonce < lastNonce {
				return nil, fmt.Errorf("invalid nonce found")
			}
			lastNonce = d.Nonce
			res = append(res, d)
			if len(res) >= int(count) {
				return res, nil
			}
		}
	}
	return res, nil
}

func (m *mockStorage) GetLastQueuedDelivery(chainID int64, sender common.Address) (*Delivery, error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	for i := len(m.deliveries) - 1; i >= 0; i-- {
		if m.deliveries[i].Sender == sender && m.deliveries[i].ChainID == chainID {
			res := m.deliveries[i]
			return &res, nil
		}
	}
	return nil, nil
}

func (m *mockStorage) GetLastDelivered(chainID int64, sender common.Address) (*Delivery, error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	for _, d := range m.deliveries {
		if d.Sender == sender && d.ChainID == chainID {
			return &d, nil
		}
	}
	return nil, nil
}

func (m *mockStorage) GetNonDeliveredCount(chainID int64, sender common.Address) (uint, error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	var count uint = 0
	for _, d := range m.deliveries {
		if d.Sender == sender && d.ChainID == chainID && d.State != DeliveryStateDelivered {
			count += 1
		}
	}
	return count, nil
}

func (m *mockStorage) getIndexById(id string) int {
	for i, d := range m.deliveries {
		if d.UniqueID == id {
			return i
		}
	}
	return -1
}

func (m *mockStorage) UpsertDeliveryRequest(tx Delivery) error {
	if tx.GasPrice == nil {
		tx.GasPrice = big.NewInt(0)
	}
	if tx.ShipmentData == nil {
		tx.ShipmentData = []byte("")
	}
	if tx.SentTransaction == nil {
		tx.SentTransaction = []byte("")
	}

	m.lock.Lock()
	defer m.lock.Unlock()

	index := m.getIndexById(tx.UniqueID)
	if index == -1 {
		m.deliveries = append(m.deliveries, tx)
	} else {
		m.deliveries[index] = tx
	}

	return nil
}

func (m *mockStorage) GetDeliveredDeliveryRequests(chainID int64, sender common.Address, olderThan time.Time, limit uint64) ([]Delivery, error) {
	res := []Delivery{}
	if limit == 0 {
		return res, nil
	}

	m.lock.Lock()
	defer m.lock.Unlock()

	for _, d := range m.deliveries {
		if d.Sender == sender && d.ChainID == chainID && d.State == DeliveryStateDelivered && d.UpdateUTC.Before(olderThan) {
			res = append(res, d)
			if len(res) >= int(limit) {
				return res, nil
			}
		}
	}
	return res, nil
}

func (m *mockStorage) DeleteDelivery(delivery ...Delivery) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	for _, d := range delivery {
		index := m.getIndexById(d.UniqueID)
		if index == -1 {
			return fmt.Errorf("not found")
		}
		m.deliveries = append(m.deliveries[:index], m.deliveries[index+1:]...)
	}
	return nil
}

func (m *mockStorage) get(index int) Delivery {
	m.lock.Lock()
	defer m.lock.Unlock()

	if index >= len(m.deliveries) {
		return Delivery{}
	}
	return m.deliveries[index]
}

func (m *mockStorage) addUpdatedUtc(index int, duration time.Duration) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	if index >= len(m.deliveries) {
		return fmt.Errorf("out of range")
	}
	m.deliveries[index].UpdateUTC = m.deliveries[index].UpdateUTC.Add(duration)
	return nil
}

func (m *mockStorage) length() int {
	m.lock.Lock()
	defer m.lock.Unlock()

	return len(m.deliveries)
}

func (m *mockStorage) reset() {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.deliveries = []Delivery{}
}

type mockCourier struct {
	lastDeliveredNonce int64
	calls              uint64
	lock               sync.Mutex
}

func (m *mockCourier) DeliverTransaction(tx Delivery) (*types.Transaction, error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if m.lastDeliveredNonce == -1 || tx.Nonce > uint64(m.lastDeliveredNonce) {
		m.lastDeliveredNonce = int64(tx.Nonce)
	}
	m.calls += 1
	return types.NewTx(&types.DynamicFeeTx{
		ChainID:   big.NewInt(tx.ChainID),
		Nonce:     tx.Nonce,
		GasTipCap: tx.GasTip,
		GasFeeCap: tx.BaseFee,
		Gas:       tx.GasPrice.Uint64(),
	}), nil
}

func (m *mockCourier) CanDeliver(DeliverableType) bool {
	return true
}

func (m *mockCourier) reset() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.lastDeliveredNonce = -1
	m.calls = 0
}

func (m *mockCourier) getLastDeliveredNonce() int64 {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.lastDeliveredNonce
}

func (m *mockCourier) getCalls() uint64 {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.calls
}

type mockNonceTracker struct {
	lock        sync.Mutex
	nonces      map[string]uint64
	confirmAll  bool
	confirmNone bool
}

func (m *mockNonceTracker) SetNextNonce(chainID int64, account common.Address, fn nonceSetFn) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	key := fmt.Sprintf("%s:%d", account.Hex(), chainID)
	if v, ok := m.nonces[key]; ok {
		v++
		if err := fn(v); err != nil {
			return err
		}
		m.nonces[key] = v
		return nil
	}
	if err := fn(0); err != nil {
		return err
	}
	m.nonces[key] = 0
	return nil
}

func (m *mockNonceTracker) GetConfirmedNonce(chainID int64, account common.Address) (uint64, error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	var nonce uint64 = 0
	key := fmt.Sprintf("%s:%d", account.Hex(), chainID)
	if v, ok := m.nonces[key]; ok {
		nonce = v
	}
	if m.confirmNone && !m.confirmAll {
		return 0, nil
	}
	if m.confirmAll && !m.confirmNone {
		return nonce + 1, nil
	}
	return nonce, nil
}

func (m *mockNonceTracker) setConfirmAll(value bool) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.confirmAll = value
	if value {
		m.confirmNone = false
	}
}

func (m *mockNonceTracker) setConfirmNone(value bool) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.confirmNone = value
	if value {
		m.confirmAll = false
	}
}

func (m *mockNonceTracker) reset() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.nonces = make(map[string]uint64)
}

type mockGasStation struct {
	defaultPrice   *big.Int
	defaultBaseFee *big.Int
}

func (m *mockGasStation) GetGasPrices(chainID int64) (*gas.GasPrices, error) {
	return &gas.GasPrices{
		SafeLow: m.defaultPrice,
		Average: m.defaultPrice,
		Fast:    m.defaultPrice,
		BaseFee: m.defaultBaseFee,
	}, nil
}

func (m *mockGasStation) reset(baseFee, gasPrice *big.Int) {
	m.defaultBaseFee = baseFee
	m.defaultPrice = gasPrice
}

type mockData struct {
	Data string `json:"data"`
}
