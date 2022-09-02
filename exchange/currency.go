package exchange

import (
	"errors"
	"strings"
)

// Coin represents a coin, example: bitcoin
type Coin string

// Currency represents a currency, example: btc
type Currency string

const (
	CoinMATIC Coin = "matic-network"
	CoinETH   Coin = "ethereum"
	CoinBTC   Coin = "bitcoin"
	CoinMYST  Coin = "mysterium"

	CurrencyGBP   Currency = "GBP"
	CurrencyUSD   Currency = "USD"
	CurrencyEUR   Currency = "EUR"
	CurrencyBTC   Currency = "BTC"
	CurrencyETH   Currency = "ETH"
	CurrencyMYST  Currency = "MYST"
	CurrencyBCH   Currency = "BCH"
	CurrencyDAI   Currency = "DAI"
	CurrencyLTC   Currency = "LTC"
	CurrencyUSDT  Currency = "USDT"
	CurrencyDOGE  Currency = "DOGE"
	CurrencyMATIC Currency = "MATIC"
)

var fiatCurrencyMap = map[Currency]bool{
	CurrencyEUR: true,
	CurrencyUSD: true,
	CurrencyGBP: true,
}

var currencyToCoinMap = map[Currency]Coin{
	CurrencyMATIC: CoinMATIC,
	CurrencyETH:   CoinETH,
	CurrencyBTC:   CoinBTC,
	CurrencyMYST:  CoinMYST,
}
var coinToCurrencyMap = map[Coin]Currency{
	CoinMATIC: CurrencyMATIC,
	CoinETH:   CurrencyETH,
	CoinBTC:   CurrencyBTC,
	CoinMYST:  CurrencyMYST,
}

func CoinToCurrency(coin Coin) (Currency, bool) {
	u, ok := coinToCurrencyMap[coin]
	return u, ok
}

func CurrencyToCoin(currency Currency) (Coin, bool) {
	u, ok := currencyToCoinMap[currency]
	return u, ok
}

// Valid validates the Currency.
func (c Currency) Valid() error {
	switch c {
	case
		CurrencyMYST,
		CurrencyBCH,
		CurrencyBTC,
		CurrencyDAI,
		CurrencyEUR,
		CurrencyETH,
		CurrencyLTC,
		CurrencyUSD,
		CurrencyDOGE,
		CurrencyUSDT:
		return nil
	}
	return errors.New("invalid currency type")
}

func (c Currency) IsFiat() bool {
	return fiatCurrencyMap[c]
}

func (c Currency) ToUpper() Currency {
	return Currency(strings.ToUpper(string(c)))
}

type PriceResponse map[Coin]Rates

type Rates map[Currency]float64

// GetRateInUSD returns the price in USD
func (pr PriceResponse) GetRateInUSD(c Coin) (float64, bool) {
	v, ok := pr[c]
	if !ok {
		return 0, false
	}
	p, ok := v[CurrencyUSD]
	return p, ok
}

// GetRates returns rates for a given coin
func (pr PriceResponse) GetRates(c Coin) (Rates, bool) {
	v, ok := pr[c]
	return v, ok
}

func (r Rates) GetRate(c Currency) (float64, bool) {
	v, ok := r[c]
	if !ok {
		v, ok = r[Currency(strings.ToUpper(string(c)))]
		return v, ok
	}
	return v, ok
}

func CoinsJoin(currencies []Coin) string {
	s := make([]string, len(currencies))
	for i, v := range currencies {
		s[i] = string(v)
	}

	return strings.Join(s, ",")
}

func CurrenciesJoin(currencies []Currency) string {
	s := make([]string, len(currencies))
	for i, v := range currencies {
		s[i] = string(v)
	}

	return strings.Join(s, ",")
}
