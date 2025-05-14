package coingecko

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/mysteriumnetwork/payments/v3/exchange"
	"github.com/patrickmn/go-cache"
)

// API can be used to check exchange prices between different currencies using coingecko.
type API struct {
	baseURI string
	client  *http.Client
	cache   *cache.Cache
}

const DefaultGeckoURI = "https://api.coingecko.com/api/v3"

const name = "coingecko"

// NewAPI returns a new gecko client.
func NewAPI(baseURI string, cacheExpiration time.Duration) *API {
	if !strings.HasSuffix(baseURI, "/") {
		baseURI += "/"
	}
	return &API{
		baseURI: baseURI,
		client: &http.Client{
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout:   10 * time.Second,
					KeepAlive: 10 * time.Second,
				}).DialContext,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 4 * time.Second,
				ResponseHeaderTimeout: 3 * time.Second,
			},
			Timeout: 30 * time.Second,
		},
		cache: cache.New(cacheExpiration, 1*time.Minute),
	}
}

func (g *API) GetName() string {
	return name
}

// GetRateCacheWithFallback will try to get price from cache, if it can't find it, will try
// to look it up and save it to cache for later use.
func (g *API) GetRateCacheWithFallback(coins []exchange.Coin, vsCurrencies []exchange.Currency) (exchange.PriceResponse, error) {
	pc, err := g.GetRateCache(coins, vsCurrencies)
	if err == nil {
		return pc, nil
	}

	return g.GetRate(coins, vsCurrencies)
}

func (g *API) GetRate(coins []exchange.Coin, vsCurrencies []exchange.Currency) (exchange.PriceResponse, error) {
	path := fmt.Sprintf("simple/price?ids=%v&vs_currencies=%v", exchange.CoinsJoin(coins), exchange.CurrenciesJoin(vsCurrencies))
	url := fmt.Sprintf("%v%v", g.baseURI, path)
	resp, err := g.client.Get(url)
	if err != nil {
		return exchange.PriceResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return exchange.PriceResponse{}, fmt.Errorf("got an unexpected error code %v", resp.StatusCode)
	}

	var res exchange.PriceResponse
	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return exchange.PriceResponse{}, err
	}

	// Convert keys to ones that node understands.
	for k, v := range res {
		rates := make(exchange.Rates)
		for curr, value := range v {
			rates[curr.ToUpper()] = value
		}
		res[k] = rates
	}

	g.cache.Set(g.cachePriceKey(coins, vsCurrencies), res, cache.DefaultExpiration)

	return res, nil
}

// GetRateCache given coins and vsCurrencies returns latest response received from gecko for these values.
// Order of these values is important, they must match the ones given to GetCoinPriceInUSD()
func (g *API) GetRateCache(coins []exchange.Coin, vsCurrencies []exchange.Currency) (exchange.PriceResponse, error) {
	key := g.cachePriceKey(coins, vsCurrencies)

	obj, ok := g.cache.Get(key)
	if !ok {
		return exchange.PriceResponse{}, fmt.Errorf("nothing found in cache for key '%s'", key)
	}
	pr, ok := obj.(exchange.PriceResponse)
	if !ok {
		return exchange.PriceResponse{}, errors.New("cache is malformed, wrong object found")
	}

	return pr, nil
}

func (g *API) cachePriceKey(coins []exchange.Coin, vsCurrencies []exchange.Currency) string {
	return fmt.Sprintf("price:%s|%s", exchange.CoinsJoin(coins), exchange.CurrenciesJoin(vsCurrencies))
}
