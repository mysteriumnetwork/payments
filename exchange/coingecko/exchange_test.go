package coingecko

import (
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mysteriumnetwork/payments/v3/exchange"
	"github.com/stretchr/testify/assert"
)

func TestExchange(t *testing.T) {
	mockApi := newMockApi()
	mockApi.run()
	defer mockApi.stop()

	api := NewAPI("http://localhost:8180", 200*time.Millisecond)

	t.Run("gets rates", func(t *testing.T) {
		mockApi.setResponse(exchange.PriceResponse{
			exchange.CoinETH: exchange.Rates{
				exchange.CurrencyUSD: 1500,
			},
		})
		rates, err := api.GetRate([]exchange.Coin{exchange.CoinETH}, []exchange.Currency{exchange.CurrencyUSD})
		assert.NoError(t, err)
		assert.Equal(t, 1500., rates[exchange.CoinETH][exchange.CurrencyUSD])
		assert.Len(t, rates, 1)
		assert.Len(t, rates[exchange.CoinETH], 1)
		assert.Equal(t, 1, mockApi.Calls)

		rates, err = api.GetRate([]exchange.Coin{exchange.CoinETH}, []exchange.Currency{exchange.CurrencyUSD})
		assert.NoError(t, err)
		assert.Equal(t, 1500., rates[exchange.CoinETH][exchange.CurrencyUSD])
		assert.Equal(t, 2, mockApi.Calls)

		mockApi.reset()

		t.Run("handles errors", func(t *testing.T) {
			mockApi.setHandler(func(c *gin.Context) {
				c.JSON(http.StatusAccepted, nil)
			})
			_, err = api.GetRate([]exchange.Coin{exchange.CoinETH}, []exchange.Currency{exchange.CurrencyUSD})
			assert.Error(t, err)
			mockApi.reset()

			mockApi.setHandler(func(c *gin.Context) {
				c.JSON(http.StatusOK, "{badjson: hi}")
			})
			_, err = api.GetRate([]exchange.Coin{exchange.CoinETH}, []exchange.Currency{exchange.CurrencyUSD})
			assert.Error(t, err)
			mockApi.reset()

			mockApi.setResponse(exchange.PriceResponse{
				exchange.CoinBTC: exchange.Rates{
					exchange.CurrencyUSD: 1500,
				},
			})
			rates, err = api.GetRate([]exchange.Coin{exchange.CoinETH}, []exchange.Currency{exchange.CurrencyUSD})
			assert.NoError(t, err)
			assert.Len(t, rates, 0)
			mockApi.reset()
		})
	})

	t.Run("cache", func(t *testing.T) {
		t.Run("as fallback", func(t *testing.T) {
			mockApi.setResponse(exchange.PriceResponse{
				exchange.CoinMATIC: exchange.Rates{
					exchange.CurrencyDAI: 10,
				},
			})
			rates, err := api.GetRateCacheWithFallback([]exchange.Coin{exchange.CoinMATIC}, []exchange.Currency{exchange.CurrencyDAI})
			assert.NoError(t, err)
			assert.Equal(t, 10., rates[exchange.CoinMATIC][exchange.CurrencyDAI])
			assert.Len(t, rates, 1)
			assert.Len(t, rates[exchange.CoinMATIC], 1)
			assert.Equal(t, 1, mockApi.Calls)

			rates, err = api.GetRateCacheWithFallback([]exchange.Coin{exchange.CoinMATIC}, []exchange.Currency{exchange.CurrencyDAI})
			assert.NoError(t, err)
			assert.Equal(t, 10., rates[exchange.CoinMATIC][exchange.CurrencyDAI])
			assert.Len(t, rates, 1)
			assert.Len(t, rates[exchange.CoinMATIC], 1)
			//used cache
			assert.Equal(t, 1, mockApi.Calls)

			//wait for cache to expire
			time.Sleep(201 * time.Millisecond)
			rates, err = api.GetRateCacheWithFallback([]exchange.Coin{exchange.CoinMATIC}, []exchange.Currency{exchange.CurrencyDAI})
			assert.NoError(t, err)
			assert.Equal(t, 10., rates[exchange.CoinMATIC][exchange.CurrencyDAI])
			assert.Len(t, rates, 1)
			assert.Len(t, rates[exchange.CoinMATIC], 1)
			assert.Equal(t, 2, mockApi.Calls)

			mockApi.reset()
		})

		t.Run("no fallback", func(t *testing.T) {
			mockApi.setResponse(exchange.PriceResponse{
				exchange.CoinMYST: exchange.Rates{
					exchange.CurrencyGBP: 3,
				},
			})
			rates, err := api.GetRate([]exchange.Coin{exchange.CoinMYST}, []exchange.Currency{exchange.CurrencyGBP})
			assert.NoError(t, err)
			assert.Equal(t, 3., rates[exchange.CoinMYST][exchange.CurrencyGBP])
			assert.Len(t, rates, 1)
			assert.Len(t, rates[exchange.CoinMYST], 1)
			assert.Equal(t, 1, mockApi.Calls)

			rates, err = api.GetRateCache([]exchange.Coin{exchange.CoinMYST}, []exchange.Currency{exchange.CurrencyGBP})
			assert.NoError(t, err)
			assert.Equal(t, 3., rates[exchange.CoinMYST][exchange.CurrencyGBP])
			assert.Len(t, rates, 1)
			assert.Len(t, rates[exchange.CoinMYST], 1)
			//used cache
			assert.Equal(t, 1, mockApi.Calls)

			//wait for cache to expire
			time.Sleep(201 * time.Millisecond)
			_, err = api.GetRateCache([]exchange.Coin{exchange.CoinMYST}, []exchange.Currency{exchange.CurrencyGBP})
			assert.Error(t, err)

			mockApi.reset()
		})
	})
}

type mockApi struct {
	server   *http.Server
	response exchange.PriceResponse
	Calls    int
}

func newMockApi() *mockApi {
	return &mockApi{
		response: make(exchange.PriceResponse),
		Calls:    0,
	}
}

func (m *mockApi) reset() {
	m.response = make(exchange.PriceResponse)
	r := gin.Default()
	r.GET("/simple/price", m.simplePrices)
	m.server.Handler = r
	m.Calls = 0
}

func (m *mockApi) setResponse(response exchange.PriceResponse) {
	m.response = response
}

func (m *mockApi) setHandler(handler func(c *gin.Context)) {
	r := gin.Default()
	r.GET("/simple/price", handler)
	m.server.Handler = r
}

func (m *mockApi) run() {
	r := gin.Default()
	r.GET("/simple/price", m.simplePrices)

	m.server = &http.Server{
		Addr:    ":8180",
		Handler: r,
	}

	go func() {
		m.server.ListenAndServe()
	}()
}

func (m *mockApi) stop() {
	m.server.Close()
}

func (m *mockApi) simplePrices(c *gin.Context) {
	m.Calls++
	ids := strings.Split(c.Query("ids"), ",")
	currencies := strings.Split(c.Query("vs_currencies"), ",")

	resp := make(exchange.PriceResponse)
	for _, id := range ids {
		coin := exchange.Coin(id)
		rates, ok := m.response[coin]
		if ok {
			ratesResp := make(map[exchange.Currency]float64)
			for _, curr := range currencies {
				currency := exchange.Currency(curr)
				rate, ok := rates[currency]
				if ok {
					ratesResp[currency] = rate
				}
			}
			resp[coin] = ratesResp
		}
	}
	c.JSON(http.StatusOK, resp)
}
