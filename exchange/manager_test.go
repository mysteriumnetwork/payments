package exchange

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManager(t *testing.T) {
	mockApi := newMockApi(PriceResponse{
		CoinETH: Rates{
			CurrencyUSD: 1,
		},
	})
	mockApi2 := newMockApi(PriceResponse{
		CoinBTC: Rates{
			CurrencyUSD: 2,
		},
	})
	resetApis := func() {
		mockApi.resetCalls()
		mockApi2.resetCalls()
	}

	manager := NewMultiManager([]API{mockApi, mockApi2})

	t.Run("get rate", func(t *testing.T) {
		res, err := manager.GetRate([]Coin{CoinETH}, []Currency{CurrencyUSD})
		assert.NoError(t, err)
		assert.Equal(t, 1., res[CoinETH][CurrencyUSD])
		assert.Equal(t, 1, mockApi.Calls["GetRate"])
		assert.Equal(t, 0, mockApi.Calls["GetRateCacheWithFallback"])
		assert.Equal(t, 0, mockApi.Calls["GetRateCache"])
		assert.Equal(t, 0, mockApi2.Calls["GetRate"])
		assert.Equal(t, 0, mockApi2.Calls["GetRateCacheWithFallback"])
		assert.Equal(t, 0, mockApi2.Calls["GetRateCache"])
		resetApis()

		res, err = manager.GetRate([]Coin{CoinBTC}, []Currency{CurrencyUSD})
		assert.NoError(t, err)
		assert.Equal(t, 2., res[CoinBTC][CurrencyUSD])
		assert.Equal(t, 1, mockApi.Calls["GetRate"])
		assert.Equal(t, 0, mockApi.Calls["GetRateCacheWithFallback"])
		assert.Equal(t, 0, mockApi.Calls["GetRateCache"])
		assert.Equal(t, 1, mockApi2.Calls["GetRate"])
		assert.Equal(t, 0, mockApi2.Calls["GetRateCacheWithFallback"])
		assert.Equal(t, 0, mockApi2.Calls["GetRateCache"])
		resetApis()

		_, err = manager.GetRate([]Coin{CoinMATIC}, []Currency{CurrencyUSD})
		assert.Error(t, err)
		assert.Equal(t, 1, mockApi.Calls["GetRate"])
		assert.Equal(t, 0, mockApi.Calls["GetRateCacheWithFallback"])
		assert.Equal(t, 0, mockApi.Calls["GetRateCache"])
		assert.Equal(t, 1, mockApi2.Calls["GetRate"])
		assert.Equal(t, 0, mockApi2.Calls["GetRateCacheWithFallback"])
		assert.Equal(t, 0, mockApi2.Calls["GetRateCache"])
		resetApis()
	})

	t.Run("get rate with fallback", func(t *testing.T) {
		res, err := manager.GetRateCacheWithFallback([]Coin{CoinETH}, []Currency{CurrencyUSD})
		assert.NoError(t, err)
		assert.Equal(t, 1., res[CoinETH][CurrencyUSD])
		assert.Equal(t, 0, mockApi.Calls["GetRate"])
		assert.Equal(t, 1, mockApi.Calls["GetRateCacheWithFallback"])
		assert.Equal(t, 0, mockApi.Calls["GetRateCache"])
		assert.Equal(t, 0, mockApi2.Calls["GetRate"])
		assert.Equal(t, 0, mockApi2.Calls["GetRateCacheWithFallback"])
		assert.Equal(t, 0, mockApi2.Calls["GetRateCache"])
		resetApis()

		res, err = manager.GetRateCacheWithFallback([]Coin{CoinBTC}, []Currency{CurrencyUSD})
		assert.NoError(t, err)
		assert.Equal(t, 2., res[CoinBTC][CurrencyUSD])
		assert.Equal(t, 0, mockApi.Calls["GetRate"])
		assert.Equal(t, 1, mockApi.Calls["GetRateCacheWithFallback"])
		assert.Equal(t, 0, mockApi.Calls["GetRateCache"])
		assert.Equal(t, 0, mockApi2.Calls["GetRate"])
		assert.Equal(t, 1, mockApi2.Calls["GetRateCacheWithFallback"])
		assert.Equal(t, 0, mockApi2.Calls["GetRateCache"])
		resetApis()

		_, err = manager.GetRateCacheWithFallback([]Coin{CoinMATIC}, []Currency{CurrencyUSD})
		assert.Error(t, err)
		assert.Equal(t, 0, mockApi.Calls["GetRate"])
		assert.Equal(t, 1, mockApi.Calls["GetRateCacheWithFallback"])
		assert.Equal(t, 0, mockApi.Calls["GetRateCache"])
		assert.Equal(t, 0, mockApi2.Calls["GetRate"])
		assert.Equal(t, 1, mockApi2.Calls["GetRateCacheWithFallback"])
		assert.Equal(t, 0, mockApi2.Calls["GetRateCache"])
		resetApis()
	})

	t.Run("get rate cache", func(t *testing.T) {
		res, err := manager.GetRateCache([]Coin{CoinETH}, []Currency{CurrencyUSD})
		assert.NoError(t, err)
		assert.Equal(t, 1., res[CoinETH][CurrencyUSD])
		assert.Equal(t, 0, mockApi.Calls["GetRate"])
		assert.Equal(t, 0, mockApi.Calls["GetRateCacheWithFallback"])
		assert.Equal(t, 1, mockApi.Calls["GetRateCache"])
		assert.Equal(t, 0, mockApi2.Calls["GetRate"])
		assert.Equal(t, 0, mockApi2.Calls["GetRateCacheWithFallback"])
		assert.Equal(t, 0, mockApi2.Calls["GetRateCache"])
		resetApis()

		res, err = manager.GetRateCache([]Coin{CoinBTC}, []Currency{CurrencyUSD})
		assert.NoError(t, err)
		assert.Equal(t, 2., res[CoinBTC][CurrencyUSD])
		assert.Equal(t, 0, mockApi.Calls["GetRate"])
		assert.Equal(t, 0, mockApi.Calls["GetRateCacheWithFallback"])
		assert.Equal(t, 1, mockApi.Calls["GetRateCache"])
		assert.Equal(t, 0, mockApi2.Calls["GetRate"])
		assert.Equal(t, 0, mockApi2.Calls["GetRateCacheWithFallback"])
		assert.Equal(t, 1, mockApi2.Calls["GetRateCache"])
		resetApis()

		_, err = manager.GetRateCache([]Coin{CoinMATIC}, []Currency{CurrencyUSD})
		assert.Error(t, err)
		assert.Equal(t, 0, mockApi.Calls["GetRate"])
		assert.Equal(t, 0, mockApi.Calls["GetRateCacheWithFallback"])
		assert.Equal(t, 1, mockApi.Calls["GetRateCache"])
		assert.Equal(t, 0, mockApi2.Calls["GetRate"])
		assert.Equal(t, 0, mockApi2.Calls["GetRateCacheWithFallback"])
		assert.Equal(t, 1, mockApi2.Calls["GetRateCache"])
		resetApis()
	})
}

var errNotFound = fmt.Errorf("not found")

type mockApi struct {
	response    PriceResponse
	ReturnError error
	Calls       map[string]int
}

func newMockApi(response PriceResponse) *mockApi {
	return &mockApi{
		response: response,
		Calls:    make(map[string]int),
	}
}

// GetRateCacheWithFallback returns a rate from the default response.
func (m *mockApi) GetRateCacheWithFallback(coins []Coin, vsCurrencies []Currency) (PriceResponse, error) {
	res, err := m.GetRate(coins, vsCurrencies)
	m.Calls["GetRateCacheWithFallback"]++
	m.Calls["GetRate"]--
	return res, err
}

// GetRate returns a rate from the default response.
func (m *mockApi) GetRate(coins []Coin, vsCurrencies []Currency) (PriceResponse, error) {
	m.Calls["GetRate"]++
	if m.ReturnError != nil {
		return nil, m.ReturnError
	}
	res := make(PriceResponse)
	for _, c := range coins {
		rates, ok := m.response[c]
		if ok {
			coinRates := make(Rates)
			for _, curr := range vsCurrencies {
				rate, ok := rates[curr]
				if ok {
					coinRates[curr] = rate
				} else {
					return nil, errNotFound
				}
			}
			res[c] = coinRates
		} else {
			return nil, errNotFound
		}
	}
	return res, nil
}

// GetRateCache returns a rate from the default response.
func (m *mockApi) GetRateCache(coins []Coin, vsCurrencies []Currency) (PriceResponse, error) {
	res, err := m.GetRate(coins, vsCurrencies)
	m.Calls["GetRateCache"]++
	m.Calls["GetRate"]--
	return res, err
}

// GetName returns a name for a exchange provider
func (m *mockApi) GetName() string {
	return "mock_api"
}

func (m *mockApi) resetCalls() {
	m.Calls = make(map[string]int)
}
