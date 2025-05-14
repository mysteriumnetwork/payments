package gas

import (
	"encoding/json"
	"io"
	"math/big"
	"net/http"
	"time"

	"github.com/mysteriumnetwork/payments/v3/units"
)

// DefaultMaticStationURI is the default gas station URL that can be used in matic gas station.
// Default URL is for mainnet of matic gas station service.
const DefaultMaticStationURI = "https://gasstation.polygon.technology/v2"

// MaticStation represents matic gas station api.
type MaticStation struct {
	apiURL     string
	client     *http.Client
	upperBound *big.Int
}

type maticGasPriceResp struct {
	BlockNumber      int64   `json:"blockNumber"`
	BlockTime        int64   `json:"blockTime"`
	EstimatedBaseFee float64 `json:"estimatedBaseFee"`
	Fast             struct {
		MaxFee         float64 `json:"maxFee"`
		MaxPriorityFee float64 `json:"maxPriorityFee"`
	} `json:"fast"`
	SafeLow struct {
		MaxFee         float64 `json:"maxFee"`
		MaxPriorityFee float64 `json:"maxPriorityFee"`
	} `json:"safeLow"`
	Standard struct {
		MaxFee         float64 `json:"maxFee"`
		MaxPriorityFee float64 `json:"maxPriorityFee"`
	} `json:"standard"`
}

// NewMaticStation returns a new instance of matic gas station which can be used for gas price checks.
func NewMaticStation(apiURL string, upperBound *big.Int) *MaticStation {
	return &MaticStation{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		apiURL:     apiURL,
		upperBound: upperBound,
	}
}

func (m *MaticStation) GetGasPrices() (*GasPrices, error) {
	resp, err := m.request()
	if err != nil {
		return nil, err
	}
	prices := GasPrices{
		SafeLow: m.result(resp.SafeLow.MaxPriorityFee),
		Average: m.result(resp.Standard.MaxPriorityFee),
		Fast:    m.result(resp.Fast.MaxPriorityFee),

		BaseFee: units.FloatGweiToBigIntWei(resp.EstimatedBaseFee),
	}
	return &prices, nil
}

func (m *MaticStation) result(price float64) *big.Int {
	if price <= 30 {
		price = 31
	}

	bp := units.FloatGweiToBigIntWei(price)
	return priceMaxUpperBound(bp, m.upperBound)
}

func (m *MaticStation) request() (*maticGasPriceResp, error) {
	resp, err := m.client.Get(m.apiURL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var price maticGasPriceResp
	if err := json.Unmarshal(body, &price); err != nil {
		return nil, err
	}

	return &price, nil
}
