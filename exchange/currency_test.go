package exchange

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriceResponse(t *testing.T) {
	pr := PriceResponse{
		CoinETH: Rates{
			CurrencyBTC: 1,
			CurrencyUSD: 2000,
		},
		CoinMATIC: Rates{
			CurrencyEUR: 10,
		},
		CoinMYST: Rates{
			CurrencyUSD: 10,
		},
	}

	t.Run("get rate in usd", func(t *testing.T) {
		rate, ok := pr.GetRateInUSD(CoinETH)
		assert.True(t, ok)
		assert.Equal(t, 2000., rate)

		_, ok = pr.GetRateInUSD(CoinMATIC)
		assert.False(t, ok)

		rate, ok = pr.GetRateInUSD(CoinMYST)
		assert.True(t, ok)
		assert.Equal(t, 10., rate)
	})

	t.Run("get rate", func(t *testing.T) {
		_, ok := pr.GetRates(CoinBTC)
		assert.False(t, ok)

		ratesEth, ok := pr.GetRates(CoinETH)
		assert.True(t, ok)
		rateEthUsd, ok := ratesEth.GetRate(CurrencyUSD)
		assert.True(t, ok)
		assert.Equal(t, 2000., rateEthUsd)
		rateEthUsd, ok = ratesEth.GetRate("usd")
		assert.True(t, ok)
		assert.Equal(t, 2000., rateEthUsd)
		rateEthUsd, ok = ratesEth.GetRate("uSD")
		assert.True(t, ok)
		assert.Equal(t, 2000., rateEthUsd)
	})

	t.Run("join", func(t *testing.T) {
		res := CoinsJoin([]Coin{CoinBTC, CoinETH, CoinMYST})
		assert.Equal(t, "bitcoin,ethereum,mysterium", res)

		res = CurrenciesJoin([]Currency{CurrencyGBP, CurrencyBTC, CurrencyUSD})
		assert.Equal(t, "GBP,BTC,USD", res)
	})
}
