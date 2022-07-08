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
		nonces: make(map[string]uint64),
	}
	gasTracker := NewGasTracker(&mockGasStation{
		defaultPrice:   big.NewInt(1),
		defaultBaseFee: big.NewInt(1),
	}, map[int64]GasIncreaseOpts{
		1: {
			Multiplier:       1.1,
			PriceLimit:       big.NewInt(4),
			IncreaseInterval: time.Second,
			OverpayFor:       nil,
			OverpayByMul:     0,
		},
	}, GasTrackerSpeedMedium)
	depot := NewDepot(&mockCourier{}, &mockStorage, &mockNonceTracker, gasTracker, DepotConfig{
		MaxNonDelivered: 5,
		Workers: []DepotWorker{
			{
				Address:         senderAddr,
				ChainID:         1,
				ProcessInterval: time.Millisecond * 10,
				ProcessCount:    2,
			},
		},
	})
	depot.Run()

	t.Run("attaches and runs cleaner", func(t *testing.T) {
		depot.AttachCleaner(&mockStorage, DepotCleanupConfig{
			CleanupInterval:  10 * time.Millisecond,
			CleanupDaysLimit: 3,
			CleanupLimit:     2,
		})
		depot.RunCleaner()
	})

	t.Run("deletes old delivered", func(t *testing.T) {
		mockNonceTracker.confirmAll = true

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
			assert.Len(t, mockStorage.deliveries, 3)
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
				return len(mockStorage.deliveries) == 1
			}, time.Second, time.Millisecond*100)
			assert.Equal(t, mockStorage.deliveries[0].ShipmentData, []byte("{\"data\":\"tx2\"}"))
		})

		mockStorage.addUpdatedUtc(0, -time.Hour*1-time.Millisecond*10)

		t.Run("should clean all of the deliveries", func(t *testing.T) {
			assert.Eventually(t, func() bool {
				return len(mockStorage.deliveries) == 0
			}, time.Second, time.Millisecond*100)
		})

		mockStorage.reset()
		mockNonceTracker.reset()
	})

	t.Run("does not remove undelivered", func(t *testing.T) {
		mockNonceTracker.confirmAll = false

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
			assert.Len(t, mockStorage.deliveries, 3)
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
				return len(mockStorage.deliveries) == 1
			}, time.Second, time.Millisecond*100)

			assert.Equal(t, mockStorage.get(0).ShipmentData, []byte("{\"data\":\"tx2\"}"))
			assert.Equal(t, string(DeliveryStateSent), string(mockStorage.deliveries[0].State))
		})

		mockStorage.reset()
		mockNonceTracker.reset()
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

func (m *mockStorage) reset() {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.deliveries = []Delivery{}
}

type mockCourier struct{}

func (m *mockCourier) DeliverTransaction(tx Delivery) (*types.Transaction, error) {
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

type mockNonceTracker struct {
	lock       sync.Mutex
	nonces     map[string]uint64
	confirmAll bool
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
	if m.confirmAll {
		return nonce + 1, nil
	}
	return nonce, nil
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

type mockData struct {
	Data string `json:"data"`
}
