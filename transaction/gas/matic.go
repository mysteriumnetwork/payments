package gas

import (
	"encoding/json"
	"io/ioutil"
	"math/big"
	"net/http"
	"time"

	"github.com/mysteriumnetwork/payments/units"
)

// DefaultMaticStationURI is the default gas station URL that can be used in matic gas station.
// Default URL is for mainnet of matic gas station service.
const DefaultMaticStationURI = "https://gasstation-mainnet.matic.network"

// MaticStation represents matic gas station api.
type MaticStation struct {
	apiURL     string
	client     *http.Client
	upperBound *big.Int
}

type maticGasPriceResp struct {
	SafeLow     float64 `json:"safeLow"`
	Standard    float64 `json:"standard"`
	Fast        float64 `json:"fast"`
	Fastest     float64 `json:"fastest"`
	BlockTime   int     `json:"blockTime"`
	BlockNumber int     `json:"blockNumber"`
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
		SafeLow: m.result(resp.SafeLow),
		Average: m.result(resp.Standard),
		Fast:    m.result(resp.Fast),
	}
	return &prices, nil
}

func (m *MaticStation) result(price float64) *big.Int {
	bp := units.FloatGweiToBigIntWei(price)
	return priceMaxUpperBound(bp, m.upperBound)
}

func (m *MaticStation) request() (*maticGasPriceResp, error) {
	resp, err := m.client.Get(m.apiURL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var price maticGasPriceResp
	if err := json.Unmarshal(body, &price); err != nil {
		return nil, err
	}

	return &price, nil
}
