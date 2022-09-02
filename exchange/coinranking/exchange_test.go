package coinranking

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mysteriumnetwork/payments/exchange"
	"github.com/stretchr/testify/assert"
)

func TestExchange(t *testing.T) {
	mockApi := newMockApi()
	mockApi.run()
	defer mockApi.stop()

	api := NewAPI("http://localhost:8081", "token", 200*time.Millisecond)

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
			assert.Len(t, rates, 1)
			assert.Len(t, rates[exchange.CoinETH], 0)
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
	r.GET("/coins", m.coins)
	m.server.Handler = r
	m.Calls = 0
}

func (m *mockApi) setResponse(response exchange.PriceResponse) {
	m.response = response
}

func (m *mockApi) setHandler(handler func(c *gin.Context)) {
	r := gin.Default()
	r.GET("/coins", handler)
	m.server.Handler = r
}

func (m *mockApi) run() {
	r := gin.Default()
	r.GET("/coins", m.coins)

	m.server = &http.Server{
		Addr:    ":8081",
		Handler: r,
	}

	go func() {
		m.server.ListenAndServe()
	}()
}

func (m *mockApi) stop() {
	m.server.Close()
}

func (m *mockApi) coins(c *gin.Context) {
	m.Calls++
	token := c.GetHeader("x-access-token")
	if token != "token" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	uuids := c.QueryArray("uuids[]")
	currencyUuid := c.Query("referenceCurrencyUuid")

	resp := coinsPricesResponse{
		Status: "success",
		Data: coinsPricesResponseData{
			Coins: []coinResponse{},
		},
	}
	for _, coinUuid := range uuids {
		coinAsCurrency, ok := uuidToCurrency(coinUuid)
		if !ok {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		coin, ok := exchange.CurrencyToCoin(coinAsCurrency)
		if !ok {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		rates, ok := m.response[coin]
		if ok {
			currency, ok := uuidToCurrency(currencyUuid)
			if !ok {
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			rate, ok := rates[currency]
			if !ok {
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}

			resp.Data.Coins = append(resp.Data.Coins, coinResponse{
				UUID:  coinUuid,
				Price: fmt.Sprintf("%f", rate),
			})
		}
	}
	c.JSON(http.StatusOK, resp)
}
