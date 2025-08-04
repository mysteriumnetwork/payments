package gas

import (
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strconv"
	"time"

	"github.com/mysteriumnetwork/payments/v3/units"
	"github.com/rs/zerolog/log"
)

// PolygonscanStation represents the polygonscan api to retrive polygonscan prices.
type PolygonscanStation struct {
	apiConfig  EtherscanApiConfig
	upperBound *big.Int

	client *http.Client
}

// NewPolygonscanStation returns a new instance of polygonscan api for polygonscan price checks.
func NewPolygonscanStation(timeout time.Duration, apiConfig EtherscanApiConfig, upperBound *big.Int) *PolygonscanStation {
	return &PolygonscanStation{
		client: &http.Client{
			Timeout: timeout,
		},
		upperBound: upperBound,
		apiConfig:  apiConfig,
	}
}

func (esa *PolygonscanStation) GetGasPrices() (*GasPrices, error) {
	res, err := esa.request()
	if err != nil {
		return nil, err
	}
	average, err := strconv.ParseFloat(res.Result.ProposeGasPrice, 64)
	if err != nil {
		return nil, err
	}
	safeLow, err := strconv.ParseFloat(res.Result.SafeGasPrice, 64)
	if err != nil {
		return nil, err
	}
	fast, err := strconv.ParseFloat(res.Result.FastGasPrice, 64)
	if err != nil {
		return nil, err
	}
	base, err := strconv.ParseFloat(res.Result.SuggestBaseFee, 64)
	if err != nil {
		return nil, err
	}
	prices := GasPrices{
		SafeLow: esa.result(safeLow),
		Average: esa.result(average),
		Fast:    esa.result(fast),

		BaseFee: units.FloatGweiToBigIntWei(base),
	}
	return &prices, nil
}

func (esa *PolygonscanStation) request() (*polygonscanGasPriceResponse, error) {
	if esa.apiConfig.ApiKey == "" {
		log.Warn().Msg("no API key set, rate is limited")
	}

	response, err := esa.client.Get(fmt.Sprintf("%vapi?chainid=%d&apikey=%s&module=gastracker&action=gasoracle", esa.apiConfig.URL, esa.apiConfig.ChainID, esa.apiConfig.ApiKey))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	resp, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var res polygonscanGasPriceResponse
	err = json.Unmarshal(resp, &res)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body from polygonscan with an error: %w and body: %s", err, string(resp))
	}

	if res.Status != "1" {
		return nil, fmt.Errorf("polygonscan api failed with message: %s", res.Message)
	}

	return &res, nil
}

func (esa *PolygonscanStation) result(price float64) *big.Int {
	if price <= 30 {
		price = 31
	}

	bp := units.FloatGweiToBigIntWei(price)
	return priceMaxUpperBound(bp, esa.upperBound)
}

// polygonscanGasPriceResponse returns the polygonscan station response.
type polygonscanGasPriceResponse struct {
	Status  string                 `json:"status"`
	Message string                 `json:"message"`
	Result  polygonscanPriceResult `json:"result"`
}

// polygonscanPriceResult the polygonscan prices for the last block.
type polygonscanPriceResult struct {
	LastBlock      string `json:"LastBlock"`
	SuggestBaseFee string `json:"suggestBaseFee"`

	FastGasPrice    string `json:"FastGasPrice"`
	ProposeGasPrice string `json:"ProposeGasPrice"`
	SafeGasPrice    string `json:"SafeGasPrice"`

	UsdPrice     string `json:"UsdPrice"`
	GasUsedRatio string `json:"gasUsedRatio"`
}
