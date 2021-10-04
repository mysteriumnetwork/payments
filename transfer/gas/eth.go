package gas

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/mysteriumnetwork/payments/units"
)

// DefaultDefiPulseEndpointURI the default defipulse api endpoint.
const DefaultDefiPulseEndpointURI = "https://data-api.defipulse.com/api/v1/"

// EthStation represents the defi pulse api to retrive gas prices.
type EthStation struct {
	APIKey      string
	EndpointURI string
	upperBound  *big.Int

	client *http.Client
}

// NewEthStation returns a new instance of defi pulse api for gas price checks.
func NewEthStation(timeout time.Duration, apiKey, endpointURI string, upperBound *big.Int) *EthStation {
	endpoint := endpointURI
	if !strings.HasSuffix(endpoint, "/") {
		endpoint += "/"
	}

	return &EthStation{
		client: &http.Client{
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout:   10 * time.Second,
					KeepAlive: 5 * time.Second,
				}).DialContext,
				TLSHandshakeTimeout:   5 * time.Second,
				ExpectContinueTimeout: 4 * time.Second,
				ResponseHeaderTimeout: 3 * time.Second,
			},
			Timeout: timeout,
		},
		APIKey:      apiKey,
		EndpointURI: endpoint,
		upperBound:  upperBound,
	}
}

func (dpa *EthStation) GetGasPrices() (*GasPrices, error) {
	res, err := dpa.defiRequest()
	if err != nil {
		return nil, err
	}
	prices := GasPrices{
		SafeLow: dpa.result(res.SafeLow),
		Average: dpa.result(res.Average),
		Fast:    dpa.result(res.Fast),
	}
	return &prices, nil
}

func (dpa *EthStation) defiRequest() (*GasStationResponse, error) {
	if dpa.APIKey == "" {
		return nil, errors.New("no API key set, can't use eth gas station")
	}

	response, err := dpa.client.Get(fmt.Sprintf("%v%v%v", dpa.EndpointURI, "egs/api/ethgasAPI.json?api-key=", dpa.APIKey))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	resp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	res := GasStationResponse{}
	err = json.Unmarshal(resp, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (dpa *EthStation) result(p float64) *big.Int {
	pb := p / 10.0
	return priceMaxUpperBound(units.FloatGweiToBigIntWei(pb), dpa.upperBound)
}

// GasStationResponse returns the gas station response.
type GasStationResponse struct {
	Fast          float64       `json:"fast"`
	Fastest       float64       `json:"fastest"`
	SafeLow       float64       `json:"safeLow"`
	Average       float64       `json:"average"`
	BlockTime     float64       `json:"block_time"`
	BlockNum      int64         `json:"blockNum"`
	Speed         float64       `json:"speed"`
	SafeLowWait   float64       `json:"safeLowWait"`
	AvgWait       float64       `json:"avgWait"`
	FastWait      float64       `json:"fastWait"`
	FastestWait   float64       `json:"fastestWait"`
	GasPriceRange GasPriceRange `json:"gasPriceRange"`
}

// GasPriceRange the gas price range report.
type GasPriceRange struct {
	Num4   float64 `json:"4"`
	Num6   float64 `json:"6"`
	Num8   float64 `json:"8"`
	Num10  float64 `json:"10"`
	Num20  float64 `json:"20"`
	Num30  float64 `json:"30"`
	Num40  float64 `json:"40"`
	Num50  float64 `json:"50"`
	Num60  float64 `json:"60"`
	Num70  float64 `json:"70"`
	Num80  float64 `json:"80"`
	Num90  float64 `json:"90"`
	Num100 float64 `json:"100"`
	Num110 float64 `json:"110"`
	Num120 float64 `json:"120"`
	Num130 float64 `json:"130"`
	Num140 float64 `json:"140"`
	Num150 float64 `json:"150"`
	Num160 float64 `json:"160"`
	Num170 float64 `json:"170"`
	Num180 float64 `json:"180"`
	Num190 float64 `json:"190"`
	Num200 float64 `json:"200"`
	Num220 float64 `json:"220"`
	Num240 float64 `json:"240"`
	Num260 float64 `json:"260"`
	Num280 float64 `json:"280"`
	Num300 float64 `json:"300"`
	Num320 float64 `json:"320"`
	Num340 float64 `json:"340"`
	Num360 float64 `json:"360"`
	Num380 float64 `json:"380"`
	Num400 float64 `json:"400"`
	Num420 float64 `json:"420"`
	Num440 float64 `json:"440"`
	Num460 float64 `json:"460"`
	Num480 float64 `json:"480"`
	Num500 float64 `json:"500"`
	Num520 float64 `json:"520"`
	Num540 float64 `json:"540"`
	Num560 float64 `json:"560"`
	Num570 float64 `json:"570"`
	Num580 float64 `json:"580"`
	Num600 float64 `json:"600"`
	Num620 float64 `json:"620"`
	Num640 float64 `json:"640"`
	Num660 float64 `json:"660"`
	Num680 float64 `json:"680"`
	Num700 float64 `json:"700"`
	Num720 float64 `json:"720"`
	Num740 float64 `json:"740"`
	Num760 float64 `json:"760"`
	Num780 float64 `json:"780"`
	Num800 float64 `json:"800"`
	Num820 float64 `json:"820"`
}
