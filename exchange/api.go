package exchange

// API reprenets an implementation of a exchange API.
type API interface {
	// GetRateCacheWithFallback tries to return exchange rate from cache, if no
	// cache is found it will look up the values in the exchange.
	GetRateCacheWithFallback(coins []Coin, vsCurrencies []Currency) (PriceResponse, error)

	// GetRate returns a rate from the exchange skipping cache.
	GetRate(coins []Coin, vsCurrencies []Currency) (PriceResponse, error)

	// GetRateCache returns a rate for a given currencies from cache only.
	// If the rate is not found in cache it will return an error.
	GetRateCache(coins []Coin, vsCurrencies []Currency) (PriceResponse, error)

	// GetName returns a name for a exchange provider
	GetName() string
}
