package coinranking

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/mysteriumnetwork/payments/exchange"
	"github.com/patrickmn/go-cache"
)

// API represents a Coinranking REST API client.
type API struct {
	baseURI     string
	client      *http.Client
	accessToken string
	cache       *cache.Cache
}

// DefaultCoinRankingURI points to default API api.
const DefaultCoinRankingURI = "https://api.coinranking.com/v2/"

const name = "coinranking"

// NewAPI returns a new API API client.
func NewAPI(baseURI string, coinRankingAccessToken string, cacheExpiration time.Duration) *API {
	if !strings.HasSuffix(baseURI, "/") {
		baseURI += "/"
	}
	return &API{
		baseURI:     baseURI,
		accessToken: coinRankingAccessToken,

		client: &http.Client{
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout:   10 * time.Second,
					KeepAlive: 10 * time.Second,
				}).DialContext,
				TLSHandshakeTimeout:   5 * time.Second,
				ExpectContinueTimeout: 4 * time.Second,
				ResponseHeaderTimeout: 3 * time.Second,
			},
			Timeout: 15 * time.Second,
		},
		cache: cache.New(cacheExpiration, 1*time.Minute),
	}
}

type coinsPricesResponse struct {
	Status string                  `json:"status"`
	Data   coinsPricesResponseData `json:"data"`
}

type coinsPricesResponseData struct {
	Coins []coinResponse `json:"coins"`
}

type coinResponse struct {
	UUID  string `json:"uuid"`
	Price string `json:"price"`
}

var currencyToUuid = map[exchange.Currency]string{
	exchange.CurrencyETH:   "razxDUgYGNAdQ",
	exchange.CurrencyMYST:  "C5Hx25DA3vp8h",
	exchange.CurrencyMATIC: "uW2tk-ILY0ii",
	exchange.CurrencyUSD:   "yhjMzLPhuIDl",
	exchange.CurrencyEUR:   "5k-_VTxqtCEI",
	exchange.CurrencyBTC:   "Qwsogvtv82FCd",
	exchange.CurrencyBCH:   "ZlZpzOJo43mIo",
	exchange.CurrencyDAI:   "MoTuySvg7",
	exchange.CurrencyLTC:   "D7B1x_ks7WhV5",
	exchange.CurrencyDOGE:  "a91GCGd_u96cF",
	exchange.CurrencyUSDT:  "HIVsRcGKkPFtW",
	exchange.CurrencyGBP:   "Hokyui45Z38f",
}
var uuidToCurrencyMap = map[string]exchange.Currency{
	"razxDUgYGNAdQ": exchange.CurrencyETH,
	"C5Hx25DA3vp8h": exchange.CurrencyMYST,
	"uW2tk-ILY0ii":  exchange.CurrencyMATIC,
	"yhjMzLPhuIDl":  exchange.CurrencyUSD,
	"5k-_VTxqtCEI":  exchange.CurrencyEUR,
	"Qwsogvtv82FCd": exchange.CurrencyBTC,
	"ZlZpzOJo43mIo": exchange.CurrencyBCH,
	"MoTuySvg7":     exchange.CurrencyDAI,
	"D7B1x_ks7WhV5": exchange.CurrencyLTC,
	"a91GCGd_u96cF": exchange.CurrencyDOGE,
	"HIVsRcGKkPFtW": exchange.CurrencyUSDT,
	"Hokyui45Z38f":  exchange.CurrencyGBP,
}

func currencyToUUID(currency exchange.Currency) (string, bool) {
	u, ok := currencyToUuid[currency]
	return u, ok
}

func uuidToCurrency(uuid string) (exchange.Currency, bool) {
	c, ok := uuidToCurrencyMap[uuid]
	return c, ok
}

func (cr *API) GetName() string {
	return name
}

func (cr *API) GetRateCacheWithFallback(coins []exchange.Coin, vsCurrencies []exchange.Currency) (exchange.PriceResponse, error) {
	pc, err := cr.GetRateCache(coins, vsCurrencies)
	if err == nil {
		return pc, nil
	}

	return cr.GetRate(coins, vsCurrencies)
}

func (cr *API) GetRate(coins []exchange.Coin, vsCurrencies []exchange.Currency) (exchange.PriceResponse, error) {
	r := make(exchange.PriceResponse)
	uiids := []string{}
	// create map for each coin and add coin uuid to array
	for _, co := range coins {
		c, ok := exchange.CoinToCurrency(co)
		if !ok {
			return exchange.PriceResponse{}, errors.New("unknown coin: " + string(co))
		}
		u, ok := currencyToUUID(c)
		if !ok {
			return exchange.PriceResponse{}, errors.New("unknown currency: " + string(u))
		}

		r[co] = make(exchange.Rates)
		uiids = append(uiids, u)
	}
	// get price of all coins against a single vsCurrency
	for _, cur := range vsCurrencies {
		url := cr.baseURI + "coins"
		req, _ := http.NewRequest(http.MethodGet, url, nil)
		if cr.accessToken != "" {
			req.Header.Add("x-access-token", cr.accessToken)
		}
		q := req.URL.Query()

		// request price for all coins
		for _, uiid := range uiids {
			q.Add("uuids[]", uiid)
		}
		curUuid, ok := currencyToUUID(cur)
		if !ok {
			return exchange.PriceResponse{}, errors.New("unknown currency: " + string(cur))
		}

		// set reference currency
		q.Add("referenceCurrencyUuid", curUuid)
		req.URL.RawQuery = q.Encode()

		resp, err := cr.client.Do(req)
		if err != nil {
			return exchange.PriceResponse{}, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			return exchange.PriceResponse{}, fmt.Errorf("got an unexpected error code %v, body: %v", resp.StatusCode, string(body))
		}
		var res coinsPricesResponse
		if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
			return exchange.PriceResponse{}, err
		}

		for _, v := range res.Data.Coins {
			price, err := strconv.ParseFloat(v.Price, 64)
			if err != nil {
				return exchange.PriceResponse{}, err
			}

			coinInCurrency, ok := uuidToCurrency(v.UUID)
			if !ok {
				return exchange.PriceResponse{}, errors.New("unknown uuid: " + v.UUID)
			}
			coin, ok := exchange.CurrencyToCoin(coinInCurrency)
			if !ok {
				return exchange.PriceResponse{}, errors.New("unknown currency: " + string(coinInCurrency))
			}
			r[coin][cur] = price
		}
	}

	cr.cache.Set(cr.cachePriceKey(coins, vsCurrencies), r, cache.DefaultExpiration)

	return r, nil
}

// GetRateCache given coins and vsCurrencies returns latest response received from gecko for these values.
// Order of these values is important, they must match the ones given to GetCoinPriceInUSD()
func (cr *API) GetRateCache(coins []exchange.Coin, vsCurrencies []exchange.Currency) (exchange.PriceResponse, error) {
	key := cr.cachePriceKey(coins, vsCurrencies)

	obj, ok := cr.cache.Get(key)
	if !ok {
		return exchange.PriceResponse{}, fmt.Errorf("nothing found in cache for key '%s'", key)
	}
	pr, ok := obj.(exchange.PriceResponse)
	if !ok {
		return exchange.PriceResponse{}, errors.New("cache is malformed, wrong object found")
	}

	return pr, nil
}

func (cr *API) cachePriceKey(coins []exchange.Coin, vsCurrencies []exchange.Currency) string {
	return fmt.Sprintf("price:%s|%s", exchange.CoinsJoin(coins), exchange.CurrenciesJoin(vsCurrencies))
}
