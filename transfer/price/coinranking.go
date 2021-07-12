/* Mysterium network payment library.
 *
 * Copyright (C) 2021 BlockDev AG
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package price

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// CoinRanking represents a CoinRanking REST API client.
type CoinRanking struct {
	baseURI     string
	client      *http.Client
	accessToken *string
}

// DefaultCoinRankingURI points to default CoinRanking api.
const DefaultCoinRankingURI = "https://api.coinranking.com/v2/"

// NewCoinRanking returns a new CoinRanking API client.
func NewCoinRanking(baseURI string, coinRankingAccessToken *string) *CoinRanking {
	if !strings.HasSuffix(baseURI, "/") {
		baseURI += "/"
	}
	return &CoinRanking{
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
			Timeout: 10 * time.Second,
		},
	}
}

type coinsPricesResponse struct {
	Status string `json:"status"`
	Data   struct {
		Coins []struct {
			UUID  string `json:"uuid"`
			Price string `json:"price"`
		} `json:"coins"`
	} `json:"data"`
}

var currencyToUuid = map[currency]string{
	ETH:   "razxDUgYGNAdQ",
	MYST:  "C5Hx25DA3vp8h",
	MATIC: "uW2tk-ILY0ii",
	USD:   "yhjMzLPhuIDl",
	EUR:   "5k-_VTxqtCEI",
}
var uuidToCurrency = map[string]currency{}

func init() {
	for k, v := range currencyToUuid {
		uuidToCurrency[v] = k
	}
}

func coinToUUID(coin currency) (string, bool) {
	u, ok := currencyToUuid[coin]
	return u, ok
}

func uuidToCoin(uuid string) (currency, bool) {
	c, ok := uuidToCurrency[uuid]
	return c, ok
}

// GetCoinPrice returns the coin price in the given currencies for given coins.
func (g *CoinRanking) GetCoinPrice(coins []string, vsCurrencies []string) (PriceResponse, error) {
	r := make(PriceResponse)
	for _, cur := range vsCurrencies {
		r[cur] = make(Prices)

		url := g.baseURI + "coins"
		req, _ := http.NewRequest(http.MethodGet, url, nil)
		if g.accessToken != nil && *g.accessToken != "" {
			req.Header.Add("x-access-token", *g.accessToken)
		}
		q := req.URL.Query()
		for _, co := range coins {
			u, ok := coinToUUID(currency(co))
			if !ok {
				return PriceResponse{}, errors.New("unknown coin: " + co)
			}
			q.Add("uuids[]", u)
		}

		curUuid, ok := coinToUUID(currency(cur))
		if !ok {
			return PriceResponse{}, errors.New("unknown currency")
		}
		q.Add("referenceCurrencyUuid", curUuid)
		req.URL.RawQuery = q.Encode()

		resp, err := g.client.Do(req)
		if err != nil {
			return PriceResponse{}, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := ioutil.ReadAll(resp.Body)
			return PriceResponse{}, fmt.Errorf("got an unexpected error code %v, body: %v", resp.StatusCode, string(body))
		}
		var res coinsPricesResponse
		if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
			return PriceResponse{}, err
		}

		for _, v := range res.Data.Coins {
			price, err := strconv.ParseFloat(v.Price, 64)
			if err != nil {
				return PriceResponse{}, err
			}

			coin, ok := uuidToCoin(v.UUID)
			if !ok {
				return PriceResponse{}, errors.New("unknown coin: " + v.UUID)
			}
			r[cur][string(coin)] = price
		}
	}
	return r, nil
}
