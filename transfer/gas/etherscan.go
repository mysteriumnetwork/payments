package gas

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/mysteriumnetwork/payments/units"
	"github.com/rs/zerolog/log"
)

// DefaultEtherscanEndpointURI the default etherscan api endpoint.
const DefaultEtherscanEndpointURI = "https://api.etherscan.io/"

// EtherscanStation represents the etherscan api to retrive gas prices.
type EtherscanStation struct {
	APIKey      string
	EndpointURI string
	upperBound  *big.Int

	client *http.Client
}

// NewEtherscanStation returns a new instance of etherscan api for gas price checks.
func NewEtherscanStation(apiKey, endpointURI string, upperBound *big.Int) *EtherscanStation {
	endpoint := endpointURI
	if !strings.HasSuffix(endpoint, "/") {
		endpoint += "/"
	}

	return &EtherscanStation{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		EndpointURI: endpointURI,
		upperBound:  upperBound,
		APIKey:      apiKey,
	}
}

func (esa *EtherscanStation) GetAverageGasPrice() (*big.Int, error) {
	res, err := esa.request()
	if err != nil {
		return nil, err
	}
	value, err := strconv.ParseFloat(res.Result.ProposeGasPrice, 64)
	if err != nil {
		return nil, err
	}
	return esa.result(value), nil
}

func (esa *EtherscanStation) GetLowGasPrice() (*big.Int, error) {
	res, err := esa.request()
	if err != nil {
		return nil, err
	}
	value, err := strconv.ParseFloat(res.Result.SafeGasPrice, 64)
	if err != nil {
		return nil, err
	}
	return esa.result(value), nil
}

func (dpa *EtherscanStation) GetFastGasPrice() (*big.Int, error) {
	res, err := dpa.request()
	if err != nil {
		return nil, err
	}
	value, err := strconv.ParseFloat(res.Result.FastGasPrice, 64)
	if err != nil {
		return nil, err
	}
	return dpa.result(value), nil
}

func (esa *EtherscanStation) request() (*etherscanGasPriceResponse, error) {
	if esa.APIKey == "" {
		log.Warn().Msg("no API key set, rate is limited")
	}

	response, err := esa.client.Get(fmt.Sprintf("%v%v%v", esa.EndpointURI, "api?module=gastracker&action=gasoracle&apikey=", esa.APIKey))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	resp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(resp))

	var res etherscanGasPriceResponse
	err = json.Unmarshal(resp, &res)
	if err != nil {
		return nil, err
	}

	if res.Status != "1" {
		return nil, errors.New("etherscan API request failed")
	}

	return &res, nil
}

func (esa *EtherscanStation) result(price float64) *big.Int {
	bp := units.FloatGweiToBigIntWei(price)
	return priceMaxUpperBound(bp, esa.upperBound)
}

// etherscanGasPriceResponse returns the gas station response.
type etherscanGasPriceResponse struct {
	Status  string         `json:"status"`
	Message string         `json:"message"`
	Result  gasPriceResult `json:"result"`
}

// gasPriceResult the gas prices for the last block.
type gasPriceResult struct {
	LastBlock       string `json:"LastBlock"`
	SafeGasPrice    string `json:"SafeGasPrice"`
	ProposeGasPrice string `json:"ProposeGasPrice"`
	FastGasPrice    string `json:"FastGasPrice"`
	SuggestBaseFee  string `json:"suggestBaseFee"`
	GasUsedRatio    string `json:"gasUsedRatio"`
}
